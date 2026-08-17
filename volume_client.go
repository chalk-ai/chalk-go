package chalk

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"connectrpc.com/connect"
	"github.com/chalk-ai/chalk-go/auth"
	"github.com/chalk-ai/chalk-go/config"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	volumev2 "github.com/chalk-ai/chalk-go/gen/chalk/volume/v2"
	"github.com/chalk-ai/chalk-go/gen/chalk/volume/v2/volumev2connect"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
	"github.com/fxamacker/cbor/v2"
	"github.com/zeebo/blake3"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/encoding/protodelim"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// VolumeProgressFunc reports a landed object's size and whether object storage
// already had it (a dedupe hit).
type VolumeProgressFunc func(bytes uint64, alreadyExisted bool)

// VolumeClient provides access to Chalk volumes.
type VolumeClient interface {
	// Volumes.
	CreateVolume(ctx context.Context, params CreateVolumeParams) (*volumev2.CreateVolumeResponse, error)
	GetVolume(ctx context.Context, volume VolumeRef, selector *volumev2.VersionSelector) (*volumev2.GetVolumeResponse, error)
	ListVolumes(ctx context.Context, params ListVolumesParams) (*volumev2.ListVolumesResponse, error)
	DeleteVolume(ctx context.Context, volume VolumeRef) error
	ListVolumeVersions(ctx context.Context, params ListVolumeVersionsParams) (*volumev2.ListVolumeVersionsResponse, error)

	// Refs.
	CreateRef(ctx context.Context, volume VolumeRef, name string, fromVersionID uint64) (*volumev2.CreateRefResponse, error)
	ListRefs(ctx context.Context, volume VolumeRef) (*volumev2.ListRefsResponse, error)
	DeleteRef(ctx context.Context, volume VolumeRef, name string) error

	// Files.
	ListFiles(ctx context.Context, params ListFilesParams) (*volumev2.ListFilesResponse, error)
	GetFile(ctx context.Context, params GetFileParams) (*volumev2.GetFileResponse, error)

	// Commits.
	GetCommitStatus(ctx context.Context, volume VolumeRef, commitID string, includeDiff bool) (*volumev2.GetCommitStatusResponse, error)
	StageFiles(ctx context.Context, volume VolumeRef, files []VolumeUploadFile, config VolumeUploadConfig, onProgress VolumeProgressFunc) ([]VolumeUploadedFile, error)
	BuildPathIntent(ctx context.Context, volume VolumeRef, upserts []VolumeUploadedFile, removes []VolumeRemovePath, ref string) (*volumev2.CommitIntent, error)
	CommitPathDeltas(ctx context.Context, volume VolumeRef, upserts []VolumeUploadedFile, removes []VolumeRemovePath, opts VolumeCommitOptions) (*volumev2.CommitStatus, error)

	// Objects.
	RequestUploadURLs(ctx context.Context, volume VolumeRef, objects []*volumev2.UploadedObjectReference) ([]*volumev2.UploadURLItem, error)
	AllocateInodeRange(ctx context.Context, volume VolumeRef, count uint64, mountID string) (*volumev2.AllocateInodeRangeResponse, error)

	// Upload and download.
	UploadFiles(ctx context.Context, request VolumeUploadRequest, onProgress VolumeProgressFunc) ([]*volumev2.CommitStatus, error)
	UploadDirectory(ctx context.Context, volumeName string, dir string, config VolumeUploadConfig, ref string) ([]*volumev2.CommitStatus, error)
	RemoveFiles(ctx context.Context, volumeName string, paths []VolumeRemovePath, opts VolumeCommitOptions) (*volumev2.CommitStatus, error)
	DownloadBytes(ctx context.Context, request VolumeDownloadRequest, onProgress func(uint64)) ([]byte, *volumev2.FileInfo, error)
	DownloadToFile(ctx context.Context, request VolumeDownloadRequest, localPath string, onProgress func(uint64)) (*volumev2.FileInfo, error)
	DownloadToWriter(ctx context.Context, request VolumeDownloadRequest, writer io.Writer, onProgress func(uint64)) (*volumev2.FileInfo, error)
	DownloadToDirectory(ctx context.Context, volumeName string, localDir string, selector *volumev2.VersionSelector, config VolumeDownloadConfig) error
}

// VolumeClientConfig configures a volume client.
type VolumeClientConfig struct {
	ClientId                   string
	ClientSecret               string
	ApiServer                  string
	EnvironmentId              string
	ConfigDir                  *string
	HTTPClient                 connect.HTTPClient
	Interceptors               []connect.Interceptor
	JWT                        *serverv1.GetTokenResponse
	Timeout                    time.Duration
	SkipEnvironmentNameMapping bool
	SkipEngineMapping          bool
	CommitAuthor               string
}

// VolumeRef identifies a volume by name or id.
type VolumeRef struct {
	Name string
	ID   string
}

// CreateVolumeParams configures CreateVolume.
type CreateVolumeParams struct {
	Name string
	Kind volumev2.VolumeKind
}

// ListVolumesParams configures ListVolumes.
type ListVolumesParams struct {
	Limit      int32
	Cursor     string
	NamePrefix string
	Kind       *volumev2.VolumeKind
}

// ListVolumeVersionsParams configures ListVolumeVersions.
type ListVolumeVersionsParams struct {
	Volume VolumeRef
	Limit  int32
	Cursor string
	Ref    *string
}

// ListFilesParams configures ListFiles.
type ListFilesParams struct {
	Volume    VolumeRef
	Path      string
	Recursive bool
	Limit     int32
	Cursor    string
	Selector  *volumev2.VersionSelector
}

// GetFileParams configures GetFile.
type GetFileParams struct {
	Volume      VolumeRef
	Path        string
	Selector    *volumev2.VersionSelector
	IfNoneMatch string
}

// VolumeCommitOptions controls how staged deltas land as a version.
type VolumeCommitOptions struct {
	// Ref is the target ref. Empty means "main".
	Ref string
	// CommitID is the idempotency key. When set, the caller owns ordering: the
	// base version and the id are left alone across retries.
	CommitID string
	// MaxCommitRetries bounds CommitVersion retries on REBASE_REQUIRED.
	MaxCommitRetries int
}

// volumeDataSegmentBytes is the default chunk size and pack cap.
const volumeDataSegmentBytes = 16 * 1024 * 1024

// VolumeUploadedIntentMinEntries is the delta count at which a commit uploads
// its deltas as an intent object instead of inlining them.
const VolumeUploadedIntentMinEntries = 10_000

// volumeIntentFrameMaxEntries caps deltas per intent frame so the server
// decodes and drops one frame at a time.
const volumeIntentFrameMaxEntries = 1000

// VolumeUploadConfig controls upload sizing, concurrency, and retries.
type VolumeUploadConfig struct {
	ChunkSize           uint64
	MaxPackBytes        uint64
	FileConcurrency     int
	ChunkConcurrency    int
	MaxChunkRetries     int
	MaxRateLimitRetries int
	MaxCommitRetries    int
}

// DefaultVolumeUploadConfig returns upload settings matching the reference client.
func DefaultVolumeUploadConfig() VolumeUploadConfig {
	return VolumeUploadConfig{
		ChunkSize:           volumeDataSegmentBytes,
		MaxPackBytes:        volumeDataSegmentBytes,
		FileConcurrency:     4,
		ChunkConcurrency:    16,
		MaxChunkRetries:     3,
		MaxRateLimitRetries: 8,
		MaxCommitRetries:    3,
	}
}

// VolumeDownloadConfig controls object-store download concurrency.
type VolumeDownloadConfig struct {
	ChunkConcurrency int
	MaxChunkRetries  int
}

// DefaultVolumeDownloadConfig returns download settings matching the reference client.
func DefaultVolumeDownloadConfig() VolumeDownloadConfig {
	return VolumeDownloadConfig{
		ChunkConcurrency: 32,
		MaxChunkRetries:  3,
	}
}

// VolumeUploadRequest is a complete upload operation against one volume.
type VolumeUploadRequest struct {
	VolumeName string
	Files      []VolumeUploadFile
	Config     VolumeUploadConfig
	Observer   VolumeUploadObserver
	// Ref is the target ref. Empty means "main".
	Ref string
}

// VolumeUploadObserver receives concurrent file-level upload events.
type VolumeUploadObserver interface {
	FileStarted(path string, size uint64)
	FileProgress(path string, bytes uint64, alreadyExisted bool)
	FileCompleted(path string, fullyDeduplicated bool)
	FileFailed(path string, err error)
}

// VolumeUploadFile is one file to commit into a volume.
type VolumeUploadFile struct {
	Path     string
	Content  VolumeUploadContent
	Metadata *volumev2.FileMetadata
}

// VolumeUploadedFile is one file's content, uploaded and ready to commit.
type VolumeUploadedFile struct {
	Path            string
	ContentRef      *volumev2.ContentRef
	Metadata        *volumev2.FileMetadata
	UploadedObjects []*volumev2.UploadedObjectReference
}

// VolumeUploadContent is the source for one uploaded file.
type VolumeUploadContent struct {
	bytes     []byte
	localPath string
}

// VolumeUploadBytes creates in-memory upload content.
func VolumeUploadBytes(data []byte) VolumeUploadContent {
	return VolumeUploadContent{bytes: append([]byte(nil), data...)}
}

// VolumeUploadLocalPath creates lazy local-file upload content.
func VolumeUploadLocalPath(path string) VolumeUploadContent {
	return VolumeUploadContent{localPath: path}
}

// VolumeRemovePath describes a path deletion in a volume.
type VolumeRemovePath struct {
	Path      string
	Recursive bool
}

// VolumeDownloadRequest is a complete download operation for one file.
type VolumeDownloadRequest struct {
	VolumeName string
	Path       string
	Selector   *volumev2.VersionSelector
	Config     VolumeDownloadConfig
}

type volumeClientImpl struct {
	httpClient   connect.HTTPClient
	apiServer    string
	tokenManager *auth.Manager
	envID        string
	timeout      *time.Duration
	author       string
	authorOnce   sync.Once

	// volumeIDs caches name -> volume_id so requests skip the name lookup.
	volumeIDs sync.Map

	rpc volumev2connect.VolumeServiceClient
}

// NewVolumeClient creates a volume client using Chalk auth and config resolution.
func NewVolumeClient(ctx context.Context, configs ...*VolumeClientConfig) (VolumeClient, error) {
	var cfg *VolumeClientConfig
	switch len(configs) {
	case 0:
		cfg = &VolumeClientConfig{}
	case 1:
		cfg = configs[0]
	default:
		return nil, errors.Newf("expected at most one VolumeClientConfig, got %d", len(configs))
	}
	if cfg == nil {
		return nil, errors.New("volume client config must not be nil")
	}
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	manager, err := config.NewManager(ctx, &config.ManagerInputs{
		APIServer:       cfg.ApiServer,
		ClientId:        config.ClientId(cfg.ClientId),
		ClientSecret:    config.ClientSecret(cfg.ClientSecret),
		EnvironmentId:   cfg.EnvironmentId,
		ConfigDir:       cfg.ConfigDir,
		GRPCQueryServer: "",
	})
	if err != nil {
		return nil, errors.Wrap(err, "getting resolved config")
	}
	var timeout *time.Duration
	if cfg.Timeout != 0 {
		timeout = &cfg.Timeout
	}
	tokenManager, err := auth.NewManager(ctx, &auth.Inputs{
		Token:                      cfg.JWT,
		HttpClient:                 cfg.HTTPClient,
		Config:                     manager,
		Timeout:                    timeout,
		SkipEnvironmentNameMapping: cfg.SkipEnvironmentNameMapping,
		SkipEngineMapping:          cfg.SkipEngineMapping,
	})
	if err != nil {
		return nil, errors.Wrap(err, "initializing token manager")
	}

	envID := manager.EnvironmentId.Value
	apiServer := manager.GetAPIServer().Value
	client := &volumeClientImpl{
		httpClient:   cfg.HTTPClient,
		apiServer:    apiServer,
		tokenManager: tokenManager,
		envID:        envID,
		timeout:      timeout,
		author:       cfg.CommitAuthor,
	}
	client.rpc = volumev2connect.NewVolumeServiceClient(
		cfg.HTTPClient,
		apiServer,
		// Retry outside auth so every attempt carries a fresh token.
		connect.WithInterceptors(
			connect.UnaryInterceptorFunc(volumeRetryInterceptor()),
			connect.UnaryInterceptorFunc(client.authInterceptor()),
		),
		connect.WithInterceptors(cfg.Interceptors...),
	)
	return client, nil
}

func (c *volumeClientImpl) authInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if c.timeout != nil {
				if _, ok := ctx.Deadline(); !ok {
					var cancel context.CancelFunc
					ctx, cancel = context.WithTimeout(ctx, *c.timeout)
					defer cancel()
				}
			}
			if err := c.addAuthHeaders(ctx, req.Header()); err != nil {
				return nil, errors.Wrap(err, "error refreshing config")
			}
			return next(ctx, req)
		}
	}
}

// volumeRetryInterceptor retries RPCs that return Unavailable.
func volumeRetryInterceptor() connect.UnaryInterceptorFunc {
	const maxAttempts = 3
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			for attempt := 1; ; attempt++ {
				res, err := next(ctx, req)
				if err == nil || attempt >= maxAttempts || connect.CodeOf(err) != connect.CodeUnavailable {
					return res, err
				}
				if waitErr := waitVolumeRetry(ctx, time.Duration(rateLimitBackoffMS(attempt-1))*time.Millisecond); waitErr != nil {
					return res, waitErr
				}
			}
		}
	}
}

func (c *volumeClientImpl) addAuthHeaders(ctx context.Context, header http.Header) error {
	header.Set("x-chalk-server", "go-api")
	header.Set("User-Agent", internal.UserAgent())
	if c.envID != "" {
		header.Set("x-chalk-env-id", c.envID)
	}
	token, err := c.tokenManager.GetJWT(ctx, time.Now().Add(time.Minute))
	if err != nil {
		return err
	}
	header.Set("Authorization", "Bearer "+token.AccessToken)
	return nil
}

func rpcMsg[T any](res *connect.Response[T], err error) (*T, error) {
	if err != nil {
		return nil, err
	}
	return res.Msg, nil
}

// volumeRef builds a VolumeRef carrying the cached volume_id, if known.
func (c *volumeClientImpl) volumeRef(name string) *volumev2.VolumeRef {
	return c.volumeRefProto(VolumeRef{Name: name})
}

func (c *volumeClientImpl) volumeRefProto(v VolumeRef) *volumev2.VolumeRef {
	id := v.ID
	if id == "" && v.Name != "" {
		if cached, ok := c.volumeIDs.Load(v.Name); ok {
			id, _ = cached.(string)
		}
	}
	return &volumev2.VolumeRef{VolumeId: id, Name: v.Name}
}

func (c *volumeClientImpl) cacheVolumeID(name string, id string) {
	if name == "" || id == "" {
		return
	}
	c.volumeIDs.Store(name, id)
}

func (c *volumeClientImpl) evictVolumeID(name string) {
	c.volumeIDs.Delete(name)
}

// volumeRPC issues an RPC carrying a VolumeRef. NotFound against a cached
// volume_id means the cache is stale: evict, and retry by name if idempotent.
func volumeRPC[Req, Resp any](
	ctx context.Context,
	c *volumeClientImpl,
	req *Req,
	volumeOf func(*Req) *volumev2.VolumeRef,
	retryOnStale bool,
	call func(context.Context, *connect.Request[Req]) (*connect.Response[Resp], error),
) (*Resp, error) {
	var cachedName string
	if volume := volumeOf(req); volume.GetVolumeId() != "" {
		cachedName = volume.GetName()
	}
	res, err := call(ctx, connect.NewRequest(req))
	if err == nil || cachedName == "" || connect.CodeOf(err) != connect.CodeNotFound {
		return rpcMsg(res, err)
	}
	c.evictVolumeID(cachedName)
	if !retryOnStale {
		return nil, err
	}
	volumeOf(req).VolumeId = ""
	return rpcMsg(call(ctx, connect.NewRequest(req)))
}

func (c *volumeClientImpl) CreateVolume(ctx context.Context, params CreateVolumeParams) (*volumev2.CreateVolumeResponse, error) {
	res, err := rpcMsg(c.rpc.CreateVolume(ctx, connect.NewRequest(&volumev2.CreateVolumeRequest{
		Name:       params.Name,
		VolumeType: params.Kind,
	})))
	if err != nil {
		return nil, err
	}
	c.cacheVolumeID(res.GetVolume().GetName(), res.GetVolume().GetVolumeId())
	return res, nil
}

func (c *volumeClientImpl) GetVolume(ctx context.Context, volume VolumeRef, selector *volumev2.VersionSelector) (*volumev2.GetVolumeResponse, error) {
	res, err := volumeRPC(ctx, c, &volumev2.GetVolumeRequest{
		Volume:   c.volumeRefProto(volume),
		Selector: selector,
	}, func(r *volumev2.GetVolumeRequest) *volumev2.VolumeRef { return r.Volume }, true, c.rpc.GetVolume)
	if err != nil {
		return nil, err
	}
	c.cacheVolumeID(res.GetVolume().GetName(), res.GetVolume().GetVolumeId())
	return res, nil
}

func (c *volumeClientImpl) ListVolumes(ctx context.Context, params ListVolumesParams) (*volumev2.ListVolumesResponse, error) {
	return rpcMsg(c.rpc.ListVolumes(ctx, connect.NewRequest(&volumev2.ListVolumesRequest{
		Limit:      params.Limit,
		Cursor:     params.Cursor,
		NamePrefix: params.NamePrefix,
		VolumeKind: params.Kind,
	})))
}

func (c *volumeClientImpl) DeleteVolume(ctx context.Context, volume VolumeRef) error {
	_, err := volumeRPC(ctx, c, &volumev2.DeleteVolumeRequest{Volume: c.volumeRefProto(volume)},
		func(r *volumev2.DeleteVolumeRequest) *volumev2.VolumeRef { return r.Volume }, false, c.rpc.DeleteVolume)
	c.evictVolumeID(volume.Name)
	return err
}

func (c *volumeClientImpl) ListVolumeVersions(ctx context.Context, params ListVolumeVersionsParams) (*volumev2.ListVolumeVersionsResponse, error) {
	return volumeRPC(ctx, c, &volumev2.ListVolumeVersionsRequest{
		Volume: c.volumeRefProto(params.Volume),
		Limit:  params.Limit,
		Cursor: params.Cursor,
		Ref:    params.Ref,
	}, func(r *volumev2.ListVolumeVersionsRequest) *volumev2.VolumeRef { return r.Volume }, true, c.rpc.ListVolumeVersions)
}

func (c *volumeClientImpl) CreateRef(ctx context.Context, volume VolumeRef, name string, fromVersionID uint64) (*volumev2.CreateRefResponse, error) {
	return volumeRPC(ctx, c, &volumev2.CreateRefRequest{
		Volume:        c.volumeRefProto(volume),
		Name:          name,
		FromVersionId: fromVersionID,
	}, func(r *volumev2.CreateRefRequest) *volumev2.VolumeRef { return r.Volume }, false, c.rpc.CreateRef)
}

func (c *volumeClientImpl) ListRefs(ctx context.Context, volume VolumeRef) (*volumev2.ListRefsResponse, error) {
	return volumeRPC(ctx, c, &volumev2.ListRefsRequest{Volume: c.volumeRefProto(volume)},
		func(r *volumev2.ListRefsRequest) *volumev2.VolumeRef { return r.Volume }, true, c.rpc.ListRefs)
}

func (c *volumeClientImpl) DeleteRef(ctx context.Context, volume VolumeRef, name string) error {
	_, err := volumeRPC(ctx, c, &volumev2.DeleteRefRequest{Volume: c.volumeRefProto(volume), Name: name},
		func(r *volumev2.DeleteRefRequest) *volumev2.VolumeRef { return r.Volume }, false, c.rpc.DeleteRef)
	return err
}

func (c *volumeClientImpl) GetCommitStatus(ctx context.Context, volume VolumeRef, commitID string, includeDiff bool) (*volumev2.GetCommitStatusResponse, error) {
	res, err := volumeRPC(ctx, c, &volumev2.GetCommitStatusRequest{
		Volume:      c.volumeRefProto(volume),
		CommitId:    commitID,
		IncludeDiff: includeDiff,
	}, func(r *volumev2.GetCommitStatusRequest) *volumev2.VolumeRef { return r.Volume }, true, c.rpc.GetCommitStatus)
	if err != nil {
		return nil, err
	}
	c.cacheVolumeID(res.GetStatus().GetVolume().GetName(), res.GetStatus().GetVolume().GetVolumeId())
	return res, nil
}

func (c *volumeClientImpl) AllocateInodeRange(ctx context.Context, volume VolumeRef, count uint64, mountID string) (*volumev2.AllocateInodeRangeResponse, error) {
	return volumeRPC(ctx, c, &volumev2.AllocateInodeRangeRequest{
		Volume:  c.volumeRefProto(volume),
		Count:   count,
		MountId: mountID,
	}, func(r *volumev2.AllocateInodeRangeRequest) *volumev2.VolumeRef { return r.Volume }, false, c.rpc.AllocateInodeRange)
}

func (c *volumeClientImpl) ListFiles(ctx context.Context, params ListFilesParams) (*volumev2.ListFilesResponse, error) {
	return volumeRPC(ctx, c, &volumev2.ListFilesRequest{
		Volume:    c.volumeRefProto(params.Volume),
		Path:      params.Path,
		Recursive: params.Recursive,
		Limit:     params.Limit,
		Cursor:    params.Cursor,
		Selector:  params.Selector,
	}, func(r *volumev2.ListFilesRequest) *volumev2.VolumeRef { return r.Volume }, true, c.rpc.ListFiles)
}

func (c *volumeClientImpl) GetFile(ctx context.Context, params GetFileParams) (*volumev2.GetFileResponse, error) {
	return c.getFile(ctx, &volumev2.GetFileRequest{
		Volume:      c.volumeRefProto(params.Volume),
		Path:        params.Path,
		Selector:    params.Selector,
		IfNoneMatch: params.IfNoneMatch,
	})
}

func (c *volumeClientImpl) getFile(ctx context.Context, req *volumev2.GetFileRequest) (*volumev2.GetFileResponse, error) {
	return volumeRPC(ctx, c, req,
		func(r *volumev2.GetFileRequest) *volumev2.VolumeRef { return r.Volume }, true, c.rpc.GetFile)
}

func (c *volumeClientImpl) RequestUploadURLs(ctx context.Context, volume VolumeRef, objects []*volumev2.UploadedObjectReference) ([]*volumev2.UploadURLItem, error) {
	return c.requestUploadURLs(ctx, c.volumeRefProto(volume), objects)
}

func (c *volumeClientImpl) requestUploadURLs(ctx context.Context, volume *volumev2.VolumeRef, objects []*volumev2.UploadedObjectReference) ([]*volumev2.UploadURLItem, error) {
	res, err := volumeRPC(ctx, c, &volumev2.RequestUploadURLsRequest{
		Volume:  volume,
		Objects: objects,
	}, func(r *volumev2.RequestUploadURLsRequest) *volumev2.VolumeRef { return r.Volume }, false, c.rpc.RequestUploadURLs)
	if err != nil {
		return nil, err
	}
	return res.GetUrls(), nil
}

func (c *volumeClientImpl) commitVersion(ctx context.Context, intent *volumev2.CommitIntent) (*volumev2.CommitStatus, error) {
	res, err := volumeRPC(ctx, c, &volumev2.CommitVersionRequest{Intent: intent},
		func(r *volumev2.CommitVersionRequest) *volumev2.VolumeRef { return r.GetIntent().GetVolume() },
		false, c.rpc.CommitVersion)
	if err != nil {
		return nil, err
	}
	status := res.GetStatus()
	if status == nil {
		return nil, fmt.Errorf("CommitVersion returned no status")
	}
	c.cacheVolumeID(status.GetVolume().GetName(), status.GetVolume().GetVolumeId())
	return status, nil
}

func VolumeVersionSelector(versionID uint64) *volumev2.VersionSelector {
	return &volumev2.VersionSelector{Selector: &volumev2.VersionSelector_VersionId{VersionId: versionID}}
}

func VolumeRefSelector(ref string) *volumev2.VersionSelector {
	return &volumev2.VersionSelector{Selector: &volumev2.VersionSelector_Ref{Ref: ref}}
}

// --- Upload -----------------------------------------------------------------

func (c *volumeClientImpl) UploadFiles(ctx context.Context, request VolumeUploadRequest, onProgress VolumeProgressFunc) ([]*volumev2.CommitStatus, error) {
	cfg := request.Config.withDefaults()
	volume := c.volumeRef(request.VolumeName)
	uploaded, err := c.stageFiles(ctx, volume, request.Files, cfg, onProgress, request.Observer)
	if err != nil {
		return nil, err
	}
	if len(uploaded) == 0 {
		return nil, nil
	}
	status, err := c.commitPathDeltas(ctx, volume, uploaded, nil, VolumeCommitOptions{
		Ref:              request.Ref,
		MaxCommitRetries: cfg.MaxCommitRetries,
	})
	if err != nil {
		return nil, err
	}
	return []*volumev2.CommitStatus{status}, nil
}

func (c *volumeClientImpl) UploadDirectory(ctx context.Context, volumeName string, dir string, config VolumeUploadConfig, ref string) ([]*volumev2.CommitStatus, error) {
	files, err := collectVolumeLocalFiles(dir)
	if err != nil {
		return nil, err
	}
	return c.UploadFiles(ctx, VolumeUploadRequest{
		VolumeName: volumeName,
		Files:      files,
		Config:     config,
		Ref:        ref,
	}, nil)
}

func (c *volumeClientImpl) StageFiles(ctx context.Context, volume VolumeRef, files []VolumeUploadFile, config VolumeUploadConfig, onProgress VolumeProgressFunc) ([]VolumeUploadedFile, error) {
	return c.stageFiles(ctx, c.volumeRefProto(volume), files, config.withDefaults(), onProgress, nil)
}

// stageFiles uploads files without committing them.
func (c *volumeClientImpl) stageFiles(ctx context.Context, volume *volumev2.VolumeRef, files []VolumeUploadFile, cfg VolumeUploadConfig, onProgress VolumeProgressFunc, observer VolumeUploadObserver) ([]VolumeUploadedFile, error) {
	if onProgress == nil {
		onProgress = func(uint64, bool) {}
	}
	throttle := &volumeThrottle{}
	return c.uploadFiles(ctx, volume, files, cfg, throttle, onProgress, observer)
}

func (c *volumeClientImpl) RemoveFiles(ctx context.Context, volumeName string, paths []VolumeRemovePath, opts VolumeCommitOptions) (*volumev2.CommitStatus, error) {
	if opts.MaxCommitRetries <= 0 {
		opts.MaxCommitRetries = DefaultVolumeUploadConfig().MaxCommitRetries
	}
	return c.commitPathDeltas(ctx, c.volumeRef(volumeName), nil, volumeRemoveDeltas(paths), opts)
}

type sizedVolumeUploadFile struct {
	file VolumeUploadFile
	size uint64
}

func (c *volumeClientImpl) uploadFiles(ctx context.Context, volume *volumev2.VolumeRef, files []VolumeUploadFile, cfg VolumeUploadConfig, throttle *volumeThrottle, onProgress VolumeProgressFunc, observer VolumeUploadObserver) ([]VolumeUploadedFile, error) {
	sized := make([]sizedVolumeUploadFile, len(files))
	var sizeGroup errgroup.Group
	sizeGroup.SetLimit(max(cfg.FileConcurrency, 1))
	for i := range files {
		sizeGroup.Go(func() error {
			size, err := files[i].Content.size()
			if err != nil {
				return err
			}
			sized[i] = sizedVolumeUploadFile{file: files[i], size: size}
			return nil
		})
	}
	if err := sizeGroup.Wait(); err != nil {
		return nil, err
	}

	// Packable when a pack holding it alone stays under the cap.
	var perFile []sizedVolumeUploadFile
	var packable []sizedVolumeUploadFile
	for _, sf := range sized {
		if sf.size > 0 && newDataPackBuilder().fits(sf.size, cfg.MaxPackBytes) {
			packable = append(packable, sf)
		} else {
			perFile = append(perFile, sf)
		}
	}

	var mu sync.Mutex
	var out []VolumeUploadedFile
	uploadGroup, uploadCtx := errgroup.WithContext(ctx)
	uploadGroup.SetLimit(max(cfg.FileConcurrency, 1))
	for _, sf := range perFile {
		uploadGroup.Go(func() error {
			if observer != nil {
				observer.FileStarted(sf.file.Path, sf.size)
			}
			var fullyDeduplicated atomic.Bool
			fullyDeduplicated.Store(sf.size > 0)
			fileProgress := func(bytes uint64, alreadyExisted bool) {
				onProgress(bytes, alreadyExisted)
				if !alreadyExisted {
					fullyDeduplicated.Store(false)
				}
				if observer != nil {
					observer.FileProgress(sf.file.Path, bytes, alreadyExisted)
				}
			}
			uploaded, err := c.uploadOneFile(uploadCtx, volume, sf.file, sf.size, cfg, throttle, fileProgress)
			if err != nil {
				if observer != nil {
					observer.FileFailed(sf.file.Path, err)
				}
				return err
			}
			if observer != nil {
				observer.FileCompleted(sf.file.Path, fullyDeduplicated.Load())
			}
			mu.Lock()
			out = append(out, uploaded)
			mu.Unlock()
			return nil
		})
	}
	uploadGroup.Go(func() error {
		uploaded, err := c.packAndUpload(uploadCtx, volume, packable, cfg, throttle, onProgress, observer)
		if err != nil {
			return err
		}
		mu.Lock()
		out = append(out, uploaded...)
		mu.Unlock()
		return nil
	})
	if err := uploadGroup.Wait(); err != nil {
		return nil, err
	}
	return out, nil
}

// volumeUploadPipelineShape splits a chunk concurrency budget into
// (pipelineDepth, window). Their product never exceeds chunkConcurrency.
func volumeUploadPipelineShape(chunkConcurrency int) (int, int) {
	limit := max(chunkConcurrency, 1)
	if limit == 1 {
		return 1, 1
	}
	return 2, limit / 2
}

// uploadOneFile chunks a file too large to pack. Windows pipeline, so the next
// window reads while the previous uploads, and bytes are released once hashed.
// Peak memory stays near chunkConcurrency * chunkSize, not the file size.
func (c *volumeClientImpl) uploadOneFile(ctx context.Context, volume *volumev2.VolumeRef, file VolumeUploadFile, size uint64, cfg VolumeUploadConfig, throttle *volumeThrottle, onProgress VolumeProgressFunc) (VolumeUploadedFile, error) {
	metadata, err := file.metadata()
	if err != nil {
		return VolumeUploadedFile{}, err
	}
	if size == 0 {
		onProgress(0, false)
		return VolumeUploadedFile{
			Path:       file.Path,
			ContentRef: emptyVolumeContentRef(),
			Metadata:   metadata,
		}, nil
	}

	depth, windowSize := volumeUploadPipelineShape(cfg.ChunkConcurrency)
	slices := computeVolumeSlices(size, cfg.ChunkSize)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Capacity plus the window being handed off caps windows holding bytes at depth.
	queue := make(chan chan volumeWindowResult, depth-1)
	go func() {
		defer close(queue)
		for start := 0; start < len(slices); start += windowSize {
			window := slices[start:min(start+windowSize, len(slices))]
			done := make(chan volumeWindowResult, 1)
			go func() {
				done <- c.uploadWindow(ctx, volume, file.Content, window, cfg, throttle, onProgress)
			}()
			select {
			case queue <- done:
			case <-ctx.Done():
				return
			}
		}
	}()

	// blake3.New, not a zero-value Hasher: the zero value digests to nothing.
	hasher := blake3.New()
	chunks := make([]*volumev2.ChunkRef, 0, len(slices))
	uploaded := make([]*volumev2.UploadedObjectReference, 0, len(slices))
	for done := range queue {
		result := <-done
		if result.err != nil {
			return VolumeUploadedFile{}, result.err
		}
		for _, data := range result.bytes {
			hasher.Write(data)
		}
		chunks = append(chunks, result.chunks...)
		uploaded = append(uploaded, result.uploaded...)
	}
	if err := ctx.Err(); err != nil {
		return VolumeUploadedFile{}, err
	}
	return VolumeUploadedFile{
		Path:            file.Path,
		ContentRef:      chunkedVolumeContentRef(hex.EncodeToString(hasher.Sum(nil)), size, chunks),
		Metadata:        metadata,
		UploadedObjects: uploaded,
	}, nil
}

// volumeWindowResult carries one window's refs plus its bytes in offset order,
// which the caller folds into the whole-file hash and then drops.
type volumeWindowResult struct {
	chunks   []*volumev2.ChunkRef
	uploaded []*volumev2.UploadedObjectReference
	bytes    [][]byte
	err      error
}

func (c *volumeClientImpl) uploadWindow(ctx context.Context, volume *volumev2.VolumeRef, content VolumeUploadContent, window []volumeSlice, cfg VolumeUploadConfig, throttle *volumeThrottle, onProgress VolumeProgressFunc) volumeWindowResult {
	objects := make([]plannedVolumeObject, len(window))
	var eg errgroup.Group
	eg.SetLimit(max(cfg.ChunkConcurrency, 1))
	for i := range window {
		eg.Go(func() error {
			data, err := content.readChunk(window[i].offset, window[i].size)
			if err != nil {
				return err
			}
			hash := blake3Hex(data)
			objects[i] = plannedVolumeObject{
				relativeKey: chunkRelativeObjectKey(hash),
				hash:        hash,
				bytes:       data,
				kind:        volumev2.UploadedObjectKind_UPLOADED_OBJECT_KIND_CHUNK,
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return volumeWindowResult{err: err}
	}

	landed, err := c.putObjects(ctx, volume, objects, cfg, throttle, onProgress)
	if err != nil {
		return volumeWindowResult{err: err}
	}
	result := volumeWindowResult{
		chunks:   make([]*volumev2.ChunkRef, len(window)),
		uploaded: make([]*volumev2.UploadedObjectReference, len(window)),
		bytes:    make([][]byte, len(window)),
	}
	for i, object := range objects {
		result.chunks[i] = &volumev2.ChunkRef{
			ObjectKey: landed[i],
			Hash:      object.hash,
			Size:      window[i].size,
			Offset:    window[i].offset,
		}
		result.uploaded[i] = uploadedVolumeObjectRef(object.relativeKey, object.hash, window[i].size, object.kind)
		result.bytes[i] = object.bytes
	}
	return result
}

type plannedVolumeObject struct {
	relativeKey string
	hash        string
	bytes       []byte
	kind        volumev2.UploadedObjectKind
}

// putObjects mints signed URLs in one RPC, lands the objects storage lacks, and
// returns each object's key.
func (c *volumeClientImpl) putObjects(ctx context.Context, volume *volumev2.VolumeRef, objects []plannedVolumeObject, cfg VolumeUploadConfig, throttle *volumeThrottle, onProgress VolumeProgressFunc) ([]string, error) {
	if len(objects) == 0 {
		return nil, nil
	}
	requested := make([]*volumev2.UploadedObjectReference, len(objects))
	for i, object := range objects {
		requested[i] = uploadedVolumeObjectRef(object.relativeKey, object.hash, uint64(len(object.bytes)), object.kind)
	}
	urls, err := c.requestUploadURLs(ctx, volume, requested)
	if err != nil {
		return nil, err
	}
	if len(urls) != len(objects) {
		return nil, fmt.Errorf("RequestUploadURLs returned %d items for %d objects", len(urls), len(objects))
	}

	landed := make([]string, len(objects))
	eg, egCtx := errgroup.WithContext(ctx)
	eg.SetLimit(max(cfg.ChunkConcurrency, 1))
	for i := range objects {
		eg.Go(func() error {
			planned, url := objects[i], urls[i]
			object := uploadedVolumeObjectRef(url.ObjectKey, planned.hash, uint64(len(planned.bytes)), planned.kind)
			if !url.AlreadyExists {
				label := "upload chunk"
				if planned.kind == volumev2.UploadedObjectKind_UPLOADED_OBJECT_KIND_PACK {
					label = "upload pack"
				}
				if _, err := c.signedRequest(egCtx, volumeSignedRequest{
					label:   label,
					method:  http.MethodPut,
					url:     url.SignedUploadUri,
					headers: octetStreamHeaders(),
					body:    planned.bytes,
				}, volumeRetryPolicy{
					maxAttempts:         cfg.MaxChunkRetries,
					throttle:            throttle,
					maxRateLimitRetries: cfg.MaxRateLimitRetries,
				}, func(ctx context.Context) (string, error) {
					return c.refreshUploadURL(ctx, volume, object)
				}); err != nil {
					return err
				}
			}
			onProgress(object.ContentSize, url.AlreadyExists)
			landed[i] = object.ObjectKey
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return landed, nil
}

func (c *volumeClientImpl) refreshUploadURL(ctx context.Context, volume *volumev2.VolumeRef, object *volumev2.UploadedObjectReference) (string, error) {
	refreshed, err := c.requestUploadURLs(ctx, volume, []*volumev2.UploadedObjectReference{object})
	if err != nil {
		return "", err
	}
	if len(refreshed) == 0 {
		return "", fmt.Errorf("URL refresh returned no items")
	}
	return refreshed[0].SignedUploadUri, nil
}

type volumeSlice struct {
	offset uint64
	size   uint64
}

func computeVolumeSlices(total uint64, chunkSize uint64) []volumeSlice {
	chunk := chunkSize
	if chunk == 0 {
		chunk = 1
	}
	var slices []volumeSlice
	for offset := uint64(0); offset < total; {
		size := min(chunk, total-offset)
		slices = append(slices, volumeSlice{offset: offset, size: size})
		offset += size
	}
	return slices
}

type packMember struct {
	path     string
	size     uint64
	metadata *volumev2.FileMetadata
	hash     [32]byte
}

type volumeReadResult struct {
	file sizedVolumeUploadFile
	data []byte
	err  error
}

// packAndUpload bin-packs small files into CDP1 packs in arrival order. Reads
// are bounded read-ahead and sealed packs upload concurrently, so peak bytes
// stay near fileConcurrency * maxPackBytes instead of the whole upload.
func (c *volumeClientImpl) packAndUpload(ctx context.Context, volume *volumev2.VolumeRef, files []sizedVolumeUploadFile, cfg VolumeUploadConfig, throttle *volumeThrottle, onProgress VolumeProgressFunc, observer VolumeUploadObserver) ([]VolumeUploadedFile, error) {
	if len(files) == 0 {
		return nil, nil
	}
	concurrency := max(cfg.FileConcurrency, 1)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Read ahead in arrival order, so pack membership stays deterministic.
	queue := make(chan chan volumeReadResult, concurrency-1)
	go func() {
		defer close(queue)
		for _, sf := range files {
			done := make(chan volumeReadResult, 1)
			go func() {
				data, err := sf.file.Content.readChunk(0, sf.size)
				done <- volumeReadResult{file: sf, data: data, err: err}
			}()
			select {
			case queue <- done:
			case <-ctx.Done():
				return
			}
		}
	}()

	var mu sync.Mutex
	var out []VolumeUploadedFile
	// eg.Go blocks once `concurrency` packs are in flight, back-pressuring reads.
	eg, egCtx := errgroup.WithContext(ctx)
	eg.SetLimit(concurrency)
	flush := func(builder *dataPackBuilder, members []packMember) {
		eg.Go(func() error {
			uploaded, err := c.uploadOnePack(egCtx, volume, builder, members, cfg, throttle, onProgress, observer)
			if err != nil {
				if observer != nil {
					for _, member := range members {
						observer.FileFailed(member.path, err)
					}
				}
				return err
			}
			mu.Lock()
			out = append(out, uploaded...)
			mu.Unlock()
			return nil
		})
	}

	var readErr error
	builder := newDataPackBuilder()
	var members []packMember
	for done := range queue {
		if egCtx.Err() != nil {
			break // a pack upload failed; eg.Wait reports it
		}
		read := <-done
		if observer != nil {
			observer.FileStarted(read.file.file.Path, read.file.size)
		}
		if read.err != nil {
			if observer != nil {
				observer.FileFailed(read.file.file.Path, read.err)
			}
			readErr = read.err
			break
		}
		metadata, err := read.file.file.metadata()
		if err != nil {
			if observer != nil {
				observer.FileFailed(read.file.file.Path, err)
			}
			readErr = err
			break
		}
		if !builder.isEmpty() && !builder.fits(read.file.size, cfg.MaxPackBytes) {
			flush(builder, members)
			builder, members = newDataPackBuilder(), nil
		}
		hash := blake3Sum(read.data)
		builder.append(hash, read.data)
		members = append(members, packMember{path: read.file.file.Path, size: read.file.size, metadata: metadata, hash: hash})
	}
	if readErr != nil {
		cancel()
	} else if !builder.isEmpty() {
		flush(builder, members)
	}

	waitErr := eg.Wait()
	if readErr != nil {
		return nil, readErr
	}
	if waitErr != nil {
		return nil, waitErr
	}
	return out, nil
}

func (c *volumeClientImpl) uploadOnePack(ctx context.Context, volume *volumev2.VolumeRef, builder *dataPackBuilder, members []packMember, cfg VolumeUploadConfig, throttle *volumeThrottle, onProgress VolumeProgressFunc, observer VolumeUploadObserver) ([]VolumeUploadedFile, error) {
	sealed, err := builder.seal()
	if err != nil {
		return nil, err
	}
	packID := hashHex(sealed.chunkID)
	relativeKey := chunkRelativeObjectKey(packID)
	fullyDeduplicated := false
	landed, err := c.putObjects(ctx, volume, []plannedVolumeObject{{
		relativeKey: relativeKey,
		hash:        packID,
		bytes:       sealed.bytes,
		kind:        volumev2.UploadedObjectKind_UPLOADED_OBJECT_KIND_PACK,
	}}, cfg, throttle, func(bytes uint64, alreadyExisted bool) {
		fullyDeduplicated = alreadyExisted
		onProgress(bytes, alreadyExisted)
	})
	if err != nil {
		return nil, err
	}
	if len(landed) == 0 {
		return nil, fmt.Errorf("RequestUploadURLs returned no item for pack")
	}
	packRef := uploadedVolumeObjectRef(relativeKey, packID, uint64(len(sealed.bytes)),
		volumev2.UploadedObjectKind_UPLOADED_OBJECT_KIND_PACK)

	out := make([]VolumeUploadedFile, 0, len(members))
	for _, member := range members {
		entry, ok := sealed.entryFor(member.hash)
		if !ok {
			return nil, fmt.Errorf("sealed pack is missing a member entry")
		}
		out = append(out, VolumeUploadedFile{
			Path:            member.path,
			ContentRef:      packedVolumeContentRef(landed[0], packID, entry.offset, entry.length, hashHex(member.hash)),
			Metadata:        member.metadata,
			UploadedObjects: []*volumev2.UploadedObjectReference{packRef},
		})
		if observer != nil {
			observer.FileProgress(member.path, member.size, fullyDeduplicated)
			observer.FileCompleted(member.path, fullyDeduplicated)
		}
	}
	return out, nil
}

// --- Commit -----------------------------------------------------------------

func (c *volumeClientImpl) BuildPathIntent(ctx context.Context, volume VolumeRef, upserts []VolumeUploadedFile, removes []VolumeRemovePath, ref string) (*volumev2.CommitIntent, error) {
	return c.buildPathIntent(ctx, c.volumeRefProto(volume), upserts, volumeRemoveDeltas(removes), ref)
}

func (c *volumeClientImpl) buildPathIntent(ctx context.Context, volume *volumev2.VolumeRef, upserts []VolumeUploadedFile, removes []*volumev2.PathRemoveDelta, ref string) (*volumev2.CommitIntent, error) {
	deltas, err := volumePathDeltas(upserts)
	if err != nil {
		return nil, err
	}
	refName := volumeRefName(ref)
	return &volumev2.CommitIntent{
		Volume:                   volume,
		Ref:                      &refName,
		UploadedObjectReferences: dedupeVolumeObjectRefs(upserts),
		Author:                   c.resolveCommitAuthor(ctx),
		Deltas: &volumev2.CommitIntent_PathDeltas{PathDeltas: &volumev2.PathDeltaList{
			Upserts: deltas,
			Removes: removes,
		}},
	}, nil
}

// buildUploadedPathIntent puts deltas in an intent object rather than the RPC
// message. The object is keyed by commit id, so the id is fixed up front.
func (c *volumeClientImpl) buildUploadedPathIntent(ctx context.Context, volume *volumev2.VolumeRef, commitID string, upserts []VolumeUploadedFile, removes []*volumev2.PathRemoveDelta, ref string) (*volumev2.CommitIntent, error) {
	deltas, err := volumePathDeltas(upserts)
	if err != nil {
		return nil, err
	}
	frames, err := encodeVolumePathIntentFrames(deltas, removes)
	if err != nil {
		return nil, err
	}
	reference, err := c.uploadIntentObject(ctx, volume, commitID, frames)
	if err != nil {
		return nil, err
	}
	refName := volumeRefName(ref)
	return &volumev2.CommitIntent{
		Volume:                   volume,
		CommitId:                 commitID,
		Ref:                      &refName,
		UploadedObjectReferences: dedupeVolumeObjectRefs(upserts),
		Author:                   c.resolveCommitAuthor(ctx),
		Deltas:                   &volumev2.CommitIntent_UploadedDeltas{UploadedDeltas: reference},
	}, nil
}

// uploadIntentObject lands framed deltas as commitID's intent object.
func (c *volumeClientImpl) uploadIntentObject(ctx context.Context, volume *volumev2.VolumeRef, commitID string, frames []byte) (*volumev2.UploadedObjectReference, error) {
	if commitID == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("uploaded intents require a non-empty commit_id"))
	}
	object := uploadedVolumeObjectRef(commitID, blake3Hex(frames), uint64(len(frames)),
		volumev2.UploadedObjectKind_UPLOADED_OBJECT_KIND_INTENT)
	urls, err := c.requestUploadURLs(ctx, volume, []*volumev2.UploadedObjectReference{object})
	if err != nil {
		return nil, err
	}
	if len(urls) == 0 {
		return nil, fmt.Errorf("RequestUploadURLs returned no items")
	}
	if !urls[0].AlreadyExists {
		if _, err := c.signedRequest(ctx, volumeSignedRequest{
			label:   "upload intent",
			method:  http.MethodPut,
			url:     urls[0].SignedUploadUri,
			headers: octetStreamHeaders(),
			body:    frames,
		}, volumeRetryPolicy{maxAttempts: 3}, func(ctx context.Context) (string, error) {
			return c.refreshUploadURL(ctx, volume, object)
		}); err != nil {
			return nil, err
		}
	}
	return object, nil
}

// encodeVolumePathIntentFrames encodes path deltas as length-delimited
// CommitDeltasObject frames. Empty deltas still emit one frame.
func encodeVolumePathIntentFrames(upserts []*volumev2.PathFileDelta, removes []*volumev2.PathRemoveDelta) ([]byte, error) {
	var frames []*volumev2.PathDeltaList
	for start := 0; start < len(upserts); start += volumeIntentFrameMaxEntries {
		frames = append(frames, &volumev2.PathDeltaList{
			Upserts: upserts[start:min(start+volumeIntentFrameMaxEntries, len(upserts))],
		})
	}
	for start := 0; start < len(removes); start += volumeIntentFrameMaxEntries {
		frames = append(frames, &volumev2.PathDeltaList{
			Removes: removes[start:min(start+volumeIntentFrameMaxEntries, len(removes))],
		})
	}
	if len(frames) == 0 {
		frames = append(frames, &volumev2.PathDeltaList{})
	}
	var out bytes.Buffer
	for _, frame := range frames {
		if _, err := protodelim.MarshalTo(&out, &volumev2.CommitDeltasObject{
			Deltas: &volumev2.CommitDeltasObject_PathDeltas{PathDeltas: frame},
		}); err != nil {
			return nil, err
		}
	}
	return out.Bytes(), nil
}

func (c *volumeClientImpl) CommitPathDeltas(ctx context.Context, volume VolumeRef, upserts []VolumeUploadedFile, removes []VolumeRemovePath, opts VolumeCommitOptions) (*volumev2.CommitStatus, error) {
	return c.commitPathDeltas(ctx, c.volumeRefProto(volume), upserts, volumeRemoveDeltas(removes), opts)
}

func (c *volumeClientImpl) commitPathDeltas(ctx context.Context, volume *volumev2.VolumeRef, upserts []VolumeUploadedFile, removes []*volumev2.PathRemoveDelta, opts VolumeCommitOptions) (*volumev2.CommitStatus, error) {
	callerOwnsOrdering := opts.CommitID != ""
	useUploaded := len(upserts)+len(removes) >= VolumeUploadedIntentMinEntries

	var intent *volumev2.CommitIntent
	var err error
	if useUploaded {
		commitID := opts.CommitID
		if commitID == "" {
			commitID = newVolumeCommitID()
		}
		intent, err = c.buildUploadedPathIntent(ctx, volume, commitID, upserts, removes, opts.Ref)
	} else {
		intent, err = c.buildPathIntent(ctx, volume, upserts, removes, opts.Ref)
		if err == nil {
			intent.CommitId = opts.CommitID
		}
	}
	if err != nil {
		return nil, err
	}

	maxCommitRetries := max(opts.MaxCommitRetries, 1)
	for attempt := 0; attempt < maxCommitRetries; attempt++ {
		// Each attempt gets a fresh envelope sharing the delta slices, so a sent
		// intent is never mutated underneath the RPC layer.
		attemptIntent := &volumev2.CommitIntent{
			Volume:                   intent.Volume,
			CommitId:                 intent.CommitId,
			Ref:                      intent.Ref,
			UploadedObjectReferences: intent.UploadedObjectReferences,
			Author:                   intent.Author,
			Deltas:                   intent.Deltas,
		}
		if !callerOwnsOrdering {
			base, err := c.tipVersion(ctx, volume, opts.Ref)
			if err != nil {
				return nil, err
			}
			if !useUploaded {
				attemptIntent.CommitId = newVolumeCommitID()
			}
			versionID, sequenceNumber := base.VersionId, base.SequenceNumber
			attemptIntent.BaseVersionId = &versionID
			attemptIntent.BaseSequenceNumber = &sequenceNumber
		}
		status, err := c.commitVersion(ctx, attemptIntent)
		if err != nil {
			return nil, err
		}
		switch status.Result {
		case volumev2.CommitResult_COMMIT_RESULT_COMMITTED:
			return status, nil
		case volumev2.CommitResult_COMMIT_RESULT_REBASE_REQUIRED:
			continue
		default:
			return nil, connect.NewError(connect.CodeFailedPrecondition, fmt.Errorf("commit failed: %s", status.Result.String()))
		}
	}
	return nil, connect.NewError(connect.CodeAborted, fmt.Errorf("commit rebased %d times without landing", maxCommitRetries))
}

// tipVersion returns the ref tip, for stamping a commit intent's base.
func (c *volumeClientImpl) tipVersion(ctx context.Context, volume *volumev2.VolumeRef, ref string) (*volumev2.VersionInfo, error) {
	var selector *volumev2.VersionSelector
	if ref != "" {
		selector = VolumeRefSelector(ref)
	}
	res, err := volumeRPC(ctx, c, &volumev2.GetVolumeRequest{Volume: volume, Selector: selector},
		func(r *volumev2.GetVolumeRequest) *volumev2.VolumeRef { return r.Volume }, true, c.rpc.GetVolume)
	if err != nil {
		return nil, err
	}
	c.cacheVolumeID(res.GetVolume().GetName(), res.GetVolume().GetVolumeId())
	version := res.GetVersion()
	if version == nil {
		return nil, fmt.Errorf("GetVolume returned no version on the volume tip")
	}
	return version, nil
}

func volumeRefName(ref string) string {
	if ref == "" {
		return "main"
	}
	return ref
}

func volumeRemoveDeltas(paths []VolumeRemovePath) []*volumev2.PathRemoveDelta {
	if len(paths) == 0 {
		return nil
	}
	out := make([]*volumev2.PathRemoveDelta, 0, len(paths))
	for _, path := range paths {
		out = append(out, &volumev2.PathRemoveDelta{Path: path.Path, Recursive: path.Recursive})
	}
	return out
}

// volumePathDeltas turns staged upserts into path deltas, rejecting inline
// content, which is only valid for symlinks.
func volumePathDeltas(upserts []VolumeUploadedFile) ([]*volumev2.PathFileDelta, error) {
	deltas := make([]*volumev2.PathFileDelta, 0, len(upserts))
	for _, file := range upserts {
		if file.ContentRef == nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("regular file %q has no content reference", file.Path))
		}
		if file.ContentRef.GetInline() != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("inline content is only valid for symlinks, not regular file %q", file.Path))
		}
		deltas = append(deltas, pathFileDelta(file))
	}
	return deltas, nil
}

func (c *volumeClientImpl) resolveCommitAuthor(ctx context.Context) string {
	if c.author != "" {
		return c.author
	}
	c.authorOnce.Do(func() {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, strings.TrimRight(c.apiServer, "/")+"/v1/who-am-i", nil)
		if err != nil {
			return
		}
		if err := c.addAuthHeaders(ctx, req.Header); err != nil {
			return
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return
		}
		var body struct {
			User          string `json:"user"`
			EnvironmentID string `json:"environment_id"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			return
		}
		envID := c.envID
		if envID == "" {
			envID = body.EnvironmentID
		}
		if envID != "" && body.User != "" {
			c.author = fmt.Sprintf("chalk:%s:agent:%s", envID, body.User)
		}
	})
	return c.author
}

// --- Signed object-store requests -------------------------------------------

// volumeThrottle shares rate-limit backoff across in-flight signed requests.
// Concurrent observers converge on the largest pending deadline.
type volumeThrottle struct {
	pauseUntilMS atomic.Int64
}

func (t *volumeThrottle) waitTurn(ctx context.Context) error {
	for {
		now := time.Now().UnixMilli()
		until := t.pauseUntilMS.Load()
		if until <= now {
			return nil
		}
		if err := waitVolumeRetry(ctx, time.Duration(until-now)*time.Millisecond); err != nil {
			return err
		}
	}
}

func (t *volumeThrottle) observeRateLimit(delayMS uint64) {
	target := time.Now().UnixMilli() + int64(delayMS)
	for {
		current := t.pauseUntilMS.Load()
		if target <= current || t.pauseUntilMS.CompareAndSwap(current, target) {
			return
		}
	}
}

type volumeSignedRequest struct {
	label   string
	method  string
	url     string
	headers http.Header
	body    []byte
}

// volumeRetryPolicy schedules retries for a signed-URL request. Uploads share a
// throttle for 429/503; downloads leave it nil and retry like any failure.
type volumeRetryPolicy struct {
	maxAttempts         int
	throttle            *volumeThrottle
	maxRateLimitRetries int
}

// signedRequest issues one signed object-store request. A rejected signature
// (400/401/403) is re-minted once via refresh; other failures retry linearly.
func (c *volumeClientImpl) signedRequest(ctx context.Context, request volumeSignedRequest, policy volumeRetryPolicy, refresh func(context.Context) (string, error)) ([]byte, error) {
	maxAttempts := max(policy.maxAttempts, 1)
	url := request.url
	refreshed := false
	rateAttempts := 0
	var lastErr error
	for attempt := 0; attempt < maxAttempts; {
		if policy.throttle != nil {
			if err := policy.throttle.waitTurn(ctx); err != nil {
				return nil, err
			}
		}
		body, status, err := c.doSignedRequest(ctx, request.method, url, request.headers, request.body)
		switch {
		case err == nil && status >= 200 && status < 300:
			return body, nil
		case err == nil && (status == http.StatusBadRequest || status == http.StatusUnauthorized || status == http.StatusForbidden) && !refreshed:
			refreshed = true
			next, refreshErr := refresh(ctx)
			if refreshErr != nil {
				return nil, refreshErr
			}
			url = next
			continue
		case err == nil && (status == http.StatusTooManyRequests || status == http.StatusServiceUnavailable) && policy.throttle != nil:
			if rateAttempts >= policy.maxRateLimitRetries {
				return nil, fmt.Errorf("%s failed: HTTP %d after %d rate-limit backoffs", request.label, status, rateAttempts)
			}
			policy.throttle.observeRateLimit(rateLimitBackoffMS(rateAttempts))
			rateAttempts++
			continue
		case err != nil:
			lastErr = err
		default:
			lastErr = fmt.Errorf("%s failed: HTTP %d", request.label, status)
		}
		attempt++
		if attempt < maxAttempts {
			if err := waitVolumeRetry(ctx, time.Duration(250*attempt)*time.Millisecond); err != nil {
				return nil, err
			}
		}
	}
	return nil, lastErr
}

func (c *volumeClientImpl) doSignedRequest(ctx context.Context, method string, url string, headers http.Header, body []byte) ([]byte, int, error) {
	var reader io.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	}
	req, err := http.NewRequestWithContext(ctx, method, url, reader)
	if err != nil {
		return nil, 0, err
	}
	req.Header = cloneHeader(headers)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return data, resp.StatusCode, nil
}

// --- Download ---------------------------------------------------------------

func (c *volumeClientImpl) DownloadBytes(ctx context.Context, request VolumeDownloadRequest, onProgress func(uint64)) ([]byte, *volumev2.FileInfo, error) {
	var mu sync.Mutex
	var out []byte
	info, err := c.downloadFile(ctx, request, func(offset uint64, data []byte) error {
		mu.Lock()
		defer mu.Unlock()
		start, ok := uint64ToInt(offset)
		if !ok {
			return fmt.Errorf("download offset %d exceeds int", offset)
		}
		end, ok := checkedIntAdd(start, len(data))
		if !ok {
			return fmt.Errorf("download range at offset %d exceeds int", offset)
		}
		if len(out) < end {
			next := make([]byte, end)
			copy(next, out)
			out = next
		}
		copy(out[start:end], data)
		return nil
	}, onProgress, false)
	if err != nil {
		return nil, nil, err
	}
	if info != nil {
		size, ok := uint64ToInt(info.Size)
		if !ok {
			return nil, nil, fmt.Errorf("download size %d exceeds int", info.Size)
		}
		if short := size - len(out); len(out) < size {
			out = append(out, make([]byte, short)...)
		}
		out = out[:size]
	}
	return out, info, nil
}

func (c *volumeClientImpl) DownloadToFile(ctx context.Context, request VolumeDownloadRequest, localPath string, onProgress func(uint64)) (*volumev2.FileInfo, error) {
	if err := os.MkdirAll(filepath.Dir(localPath), 0o755); err != nil {
		return nil, err
	}
	file, err := os.Create(localPath)
	if err != nil {
		return nil, err
	}
	var mu sync.Mutex
	info, err := c.downloadFile(ctx, request, func(offset uint64, data []byte) error {
		mu.Lock()
		defer mu.Unlock()
		_, err := file.WriteAt(data, int64(offset))
		return err
	}, onProgress, false)
	if err != nil {
		_ = file.Close()
		return nil, err
	}
	if err := file.Close(); err != nil {
		return nil, err
	}
	if info != nil {
		if err := applyVolumeFileInfoMetadata(localPath, info); err != nil {
			return nil, err
		}
	}
	return info, nil
}

// DownloadToWriter downloads one file in order with bounded buffering.
func (c *volumeClientImpl) DownloadToWriter(ctx context.Context, request VolumeDownloadRequest, writer io.Writer, onProgress func(uint64)) (*volumev2.FileInfo, error) {
	return c.downloadFile(ctx, request, func(_ uint64, data []byte) error {
		return writeVolumeAll(writer, data)
	}, onProgress, true)
}

func (c *volumeClientImpl) DownloadToDirectory(ctx context.Context, volumeName string, localDir string, selector *volumev2.VersionSelector, config VolumeDownloadConfig) error {
	files, err := c.listFilesRecursive(ctx, volumeName, selector)
	if err != nil {
		return err
	}
	eg, egCtx := errgroup.WithContext(ctx)
	eg.SetLimit(4)
	for _, info := range files {
		eg.Go(func() error {
			rel, err := safeVolumeRelativePath(info.Path)
			if err != nil {
				return err
			}
			localPath := filepath.Join(localDir, rel)
			switch info.Kind {
			case volumev2.FileKind_FILE_KIND_FILE:
				_, err := c.DownloadToFile(egCtx, VolumeDownloadRequest{
					VolumeName: volumeName,
					Path:       info.Path,
					Selector:   selector,
					Config:     config,
				}, localPath, nil)
				return err
			case volumev2.FileKind_FILE_KIND_DIRECTORY:
				if err := os.MkdirAll(localPath, 0o755); err != nil {
					return err
				}
				return applyVolumeFileInfoMetadata(localPath, info)
			}
			return nil
		})
	}
	return eg.Wait()
}

func (c *volumeClientImpl) downloadFile(ctx context.Context, request VolumeDownloadRequest, onWrite func(uint64, []byte) error, onProgress func(uint64), ordered bool) (*volumev2.FileInfo, error) {
	cfg := request.Config.withDefaults()
	if onProgress == nil {
		onProgress = func(uint64) {}
	}
	policy := volumeRetryPolicy{maxAttempts: cfg.MaxChunkRetries}
	fileReq := &volumev2.GetFileRequest{
		Volume:   c.volumeRef(request.VolumeName),
		Path:     request.Path,
		Selector: request.Selector,
	}
	resp, err := c.getFile(ctx, fileReq)
	if err != nil {
		return nil, err
	}
	info := resp.GetFile()
	switch content := resp.GetContent().(type) {
	case *volumev2.GetFileResponse_Data:
		if err := onWrite(0, content.Data); err != nil {
			return nil, err
		}
		onProgress(uint64(len(content.Data)))
	case *volumev2.GetFileResponse_Chunked:
		pinned, err := pinnedVolumeFileRequest(fileReq, resp.GetVersion())
		if err != nil {
			return nil, err
		}
		if ordered {
			if err := c.downloadChunksOrdered(ctx, content.Chunked.Chunks, pinned, policy, cfg.ChunkConcurrency, onWrite, onProgress); err != nil {
				return nil, err
			}
			break
		}
		eg, egCtx := errgroup.WithContext(ctx)
		eg.SetLimit(max(cfg.ChunkConcurrency, 1))
		for _, chunk := range content.Chunked.Chunks {
			eg.Go(func() error {
				data, err := c.downloadChunk(egCtx, pinned, chunk, policy)
				if err != nil {
					return err
				}
				if err := onWrite(chunk.Offset, data); err != nil {
					return err
				}
				onProgress(chunk.Size)
				return nil
			})
		}
		if err := eg.Wait(); err != nil {
			return nil, err
		}
	case *volumev2.GetFileResponse_Packed:
		pack := content.Packed.GetPack()
		if pack == nil {
			return nil, fmt.Errorf("packed content missing pack entry")
		}
		if pack.Size == 0 {
			if err := onWrite(0, nil); err != nil {
				return nil, err
			}
			onProgress(0)
			return info, nil
		}
		pinned, err := pinnedVolumeFileRequest(fileReq, resp.GetVersion())
		if err != nil {
			return nil, err
		}
		end, ok := checkedInclusiveRangeEnd(pack.Offset, pack.Size)
		if !ok {
			return nil, fmt.Errorf("packed range overflows uint64")
		}
		data, err := c.signedRequest(ctx, volumeSignedRequest{
			label:   "download pack",
			method:  http.MethodGet,
			url:     pack.SignedDownloadUri,
			headers: http.Header{"Range": []string{fmt.Sprintf("bytes=%d-%d", pack.Offset, end)}},
		}, policy, func(ctx context.Context) (string, error) {
			return c.refreshPackedURL(ctx, pinned)
		})
		if err != nil {
			return nil, err
		}
		if uint64(len(data)) != pack.Size {
			return nil, fmt.Errorf("packed range at offset %d returned %d bytes, expected %d", pack.Offset, len(data), pack.Size)
		}
		if info != nil && info.Hash != "" && blake3Hex(data) != info.Hash {
			return nil, connect.NewError(connect.CodeDataLoss, fmt.Errorf("packed range at offset %d failed hash validation", pack.Offset))
		}
		if err := onWrite(0, data); err != nil {
			return nil, err
		}
		onProgress(uint64(len(data)))
	}
	return info, nil
}

func (c *volumeClientImpl) downloadChunksOrdered(ctx context.Context, chunks []*volumev2.SignedChunkRef, pinned *volumev2.GetFileRequest, policy volumeRetryPolicy, concurrency int, onWrite func(uint64, []byte) error, onProgress func(uint64)) error {
	ordered := append([]*volumev2.SignedChunkRef(nil), chunks...)
	sort.Slice(ordered, func(i, j int) bool { return ordered[i].Offset < ordered[j].Offset })
	windowSize := max(concurrency, 1)
	for start := 0; start < len(ordered); start += windowSize {
		window := ordered[start:min(start+windowSize, len(ordered))]
		data := make([][]byte, len(window))
		eg, egCtx := errgroup.WithContext(ctx)
		for i, chunk := range window {
			eg.Go(func() error {
				var err error
				data[i], err = c.downloadChunk(egCtx, pinned, chunk, policy)
				return err
			})
		}
		if err := eg.Wait(); err != nil {
			return err
		}
		for i, chunk := range window {
			if err := onWrite(chunk.Offset, data[i]); err != nil {
				return err
			}
			onProgress(chunk.Size)
		}
	}
	return nil
}

func (c *volumeClientImpl) downloadChunk(ctx context.Context, pinned *volumev2.GetFileRequest, chunk *volumev2.SignedChunkRef, policy volumeRetryPolicy) ([]byte, error) {
	data, err := c.signedRequest(ctx, volumeSignedRequest{
		label:  "download chunk",
		method: http.MethodGet,
		url:    chunk.SignedDownloadUri,
	}, policy, func(ctx context.Context) (string, error) {
		return c.refreshChunkURL(ctx, pinned, chunk)
	})
	if err != nil {
		return nil, err
	}
	if uint64(len(data)) != chunk.Size {
		return nil, fmt.Errorf("download chunk at offset %d returned %d bytes, expected %d", chunk.Offset, len(data), chunk.Size)
	}
	if chunk.Hash != "" && blake3Hex(data) != chunk.Hash {
		return nil, connect.NewError(connect.CodeDataLoss, fmt.Errorf("download chunk at offset %d failed hash validation", chunk.Offset))
	}
	return data, nil
}

func writeVolumeAll(writer io.Writer, data []byte) error {
	for len(data) > 0 {
		written, err := writer.Write(data)
		if err != nil {
			return err
		}
		if written == 0 {
			return io.ErrShortWrite
		}
		data = data[written:]
	}
	return nil
}

func (c *volumeClientImpl) refreshChunkURL(ctx context.Context, pinned *volumev2.GetFileRequest, chunk *volumev2.SignedChunkRef) (string, error) {
	res, err := c.getFile(ctx, pinned)
	if err != nil {
		return "", err
	}
	chunked := res.GetChunked()
	if chunked == nil {
		return "", fmt.Errorf("URL refresh returned non-chunked content")
	}
	for _, candidate := range chunked.Chunks {
		if candidate.Offset == chunk.Offset && (chunk.Hash == "" || candidate.Hash == chunk.Hash) {
			return candidate.SignedDownloadUri, nil
		}
	}
	return "", fmt.Errorf("URL refresh returned no matching chunk at offset %d", chunk.Offset)
}

func (c *volumeClientImpl) refreshPackedURL(ctx context.Context, pinned *volumev2.GetFileRequest) (string, error) {
	res, err := c.getFile(ctx, pinned)
	if err != nil {
		return "", err
	}
	refreshed := res.GetPacked()
	if refreshed == nil || refreshed.GetPack() == nil {
		return "", fmt.Errorf("URL refresh returned non-packed content")
	}
	return refreshed.GetPack().SignedDownloadUri, nil
}

func pinnedVolumeFileRequest(request *volumev2.GetFileRequest, version *volumev2.VersionInfo) (*volumev2.GetFileRequest, error) {
	if version == nil {
		return nil, fmt.Errorf("GetFile returned remote content without a version")
	}
	return &volumev2.GetFileRequest{
		Volume:      request.GetVolume(),
		Path:        request.GetPath(),
		Selector:    VolumeVersionSelector(version.VersionId),
		IfNoneMatch: request.GetIfNoneMatch(),
	}, nil
}

func (c *volumeClientImpl) listFilesRecursive(ctx context.Context, volumeName string, selector *volumev2.VersionSelector) ([]*volumev2.FileInfo, error) {
	var out []*volumev2.FileInfo
	stack := []string{""}
	for len(stack) > 0 {
		dir := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cursor := ""
		for {
			res, err := c.ListFiles(ctx, ListFilesParams{
				Volume:   VolumeRef{Name: volumeName},
				Path:     dir,
				Cursor:   cursor,
				Selector: selector,
			})
			if err != nil {
				return nil, err
			}
			for _, file := range res.Files {
				if file.Kind == volumev2.FileKind_FILE_KIND_DIRECTORY {
					stack = append(stack, file.Path)
				}
				out = append(out, file)
			}
			cursor = res.NextCursor
			if cursor == "" {
				break
			}
		}
	}
	return out, nil
}

// --- Local content ----------------------------------------------------------

func (c VolumeUploadContent) size() (uint64, error) {
	if c.localPath == "" {
		return uint64(len(c.bytes)), nil
	}
	info, err := os.Stat(c.localPath)
	if err != nil {
		return 0, err
	}
	return uint64(info.Size()), nil
}

func (c VolumeUploadContent) readChunk(offset uint64, length uint64) ([]byte, error) {
	if c.localPath == "" {
		end, ok := checkedUint64Add(offset, length)
		if !ok || end > uint64(len(c.bytes)) {
			return nil, fmt.Errorf("read_chunk out of range: %d+%d > %d", offset, length, len(c.bytes))
		}
		return append([]byte(nil), c.bytes[offset:end]...), nil
	}
	lengthInt, ok := uint64ToInt(length)
	if !ok {
		return nil, fmt.Errorf("read_chunk length %d exceeds int", length)
	}
	if offset > uint64(^uint64(0)>>1) {
		return nil, fmt.Errorf("read_chunk offset %d exceeds int64", offset)
	}
	file, err := os.Open(c.localPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buf := make([]byte, lengthInt)
	if _, err := file.ReadAt(buf, int64(offset)); err != nil {
		return nil, err
	}
	return buf, nil
}

func (f VolumeUploadFile) metadata() (*volumev2.FileMetadata, error) {
	if f.Metadata != nil || f.Content.localPath == "" {
		return f.Metadata, nil
	}
	info, err := os.Stat(f.Content.localPath)
	if err != nil {
		return nil, err
	}
	mode := uint32(info.Mode().Perm())
	return &volumev2.FileMetadata{
		Mode:      &mode,
		UpdatedAt: timestamppb.New(info.ModTime()),
	}, nil
}

func collectVolumeLocalFiles(root string) ([]VolumeUploadFile, error) {
	var out []VolumeUploadFile
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		returnedPath := filepath.ToSlash(rel)
		out = append(out, VolumeUploadFile{
			Path:    returnedPath,
			Content: VolumeUploadLocalPath(path),
		})
		return nil
	})
	return out, err
}

func safeVolumeRelativePath(path string) (string, error) {
	clean := filepath.Clean(path)
	if filepath.IsAbs(clean) || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) || strings.HasPrefix(filepath.ToSlash(clean), "../") {
		return "", fmt.Errorf("volume path escapes download directory: %s", path)
	}
	return clean, nil
}

func applyVolumeFileInfoMetadata(path string, info *volumev2.FileInfo) error {
	if info == nil {
		return nil
	}
	if runtime.GOOS != "windows" && info.Mode != nil {
		if err := os.Chmod(path, os.FileMode(info.GetMode()&0o7777)); err != nil {
			return err
		}
	}
	if info.UpdatedAt != nil {
		t := info.UpdatedAt.AsTime()
		if err := os.Chtimes(path, t, t); err != nil {
			return err
		}
	}
	return nil
}

// --- Content refs -----------------------------------------------------------

func emptyVolumeContentRef() *volumev2.ContentRef {
	return &volumev2.ContentRef{Content: &volumev2.ContentRef_Empty{Empty: &volumev2.EmptyFileContent{}}}
}

func chunkedVolumeContentRef(hash string, size uint64, chunks []*volumev2.ChunkRef) *volumev2.ContentRef {
	return &volumev2.ContentRef{Content: &volumev2.ContentRef_Chunked{Chunked: &volumev2.ChunkedContentRef{Hash: hash, Size: size, Chunks: chunks}}}
}

func packedVolumeContentRef(objectKey string, packID string, offset uint64, length uint32, hash string) *volumev2.ContentRef {
	return &volumev2.ContentRef{Content: &volumev2.ContentRef_Packed{Packed: &volumev2.PackedContentRef{
		Hash: hash,
		Size: uint64(length),
		Pack: &volumev2.PackEntryRef{ObjectKey: objectKey, PackId: packID, Offset: offset, Size: uint64(length)},
	}}}
}

func pathFileDelta(file VolumeUploadedFile) *volumev2.PathFileDelta {
	return &volumev2.PathFileDelta{
		Path: file.Path,
		Node: &volumev2.FileNode{
			Metadata: file.Metadata,
			Node: &volumev2.FileNode_File{File: &volumev2.RegularFileNode{
				Content: file.ContentRef,
			}},
		},
		Mode: volumev2.PathWriteMode_PATH_WRITE_MODE_UPSERT,
	}
}

func uploadedVolumeObjectRef(objectKey string, hash string, size uint64, kind volumev2.UploadedObjectKind) *volumev2.UploadedObjectReference {
	return &volumev2.UploadedObjectReference{ObjectKey: objectKey, Hash: hash, ContentSize: size, Kind: kind}
}

func dedupeVolumeObjectRefs(files []VolumeUploadedFile) []*volumev2.UploadedObjectReference {
	var refs []*volumev2.UploadedObjectReference
	for _, file := range files {
		for _, ref := range file.UploadedObjects {
			if ref != nil {
				refs = append(refs, ref)
			}
		}
	}
	sort.Slice(refs, func(i, j int) bool { return refs[i].ObjectKey < refs[j].ObjectKey })
	out := refs[:0]
	var last string
	for i, ref := range refs {
		if i == 0 || ref.ObjectKey != last {
			out = append(out, ref)
			last = ref.ObjectKey
		}
	}
	return out
}

func chunkRelativeObjectKey(hash string) string {
	if len(hash) < 2 {
		return hash
	}
	return hash[:2] + "/" + hash
}

func blake3Sum(data []byte) [32]byte {
	return blake3.Sum256(data)
}

func blake3Hex(data []byte) string {
	return hashHex(blake3Sum(data))
}

func hashHex(hash [32]byte) string {
	return hex.EncodeToString(hash[:])
}

// newVolumeCommitID returns a UUIDv7 idempotency key, time-ordered by prefix.
func newVolumeCommitID() string {
	var raw [16]byte
	_, _ = rand.Read(raw[:])
	ms := time.Now().UnixMilli()
	binary.BigEndian.PutUint16(raw[0:], uint16(ms>>32))
	binary.BigEndian.PutUint32(raw[2:], uint32(ms))
	raw[6] = (raw[6] & 0x0f) | 0x70 // version 7
	raw[8] = (raw[8] & 0x3f) | 0x80 // RFC 4122 variant
	return fmt.Sprintf("%x-%x-%x-%x-%x", raw[0:4], raw[4:6], raw[6:8], raw[8:10], raw[10:16])
}

func octetStreamHeaders() http.Header {
	return http.Header{"Content-Type": []string{"application/octet-stream"}}
}

func cloneHeader(header http.Header) http.Header {
	if header == nil {
		return http.Header{}
	}
	return header.Clone()
}

func rateLimitBackoffMS(attempt int) uint64 {
	exp := min(5, attempt)
	delay := uint64(1000) << exp
	if delay > 30000 {
		return 30000
	}
	return delay
}

func waitVolumeRetry(ctx context.Context, delay time.Duration) error {
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func checkedInclusiveRangeEnd(offset uint64, size uint64) (uint64, bool) {
	if size == 0 {
		return offset, true
	}
	end := offset + size - 1
	return end, end >= offset
}

func checkedUint64Add(a uint64, b uint64) (uint64, bool) {
	sum := a + b
	return sum, sum >= a
}

func uint64ToInt(value uint64) (int, bool) {
	converted := int(value)
	return converted, converted >= 0 && uint64(converted) == value
}

func checkedIntAdd(a int, b int) (int, bool) {
	if b > int(^uint(0)>>1)-a {
		return 0, false
	}
	return a + b, true
}

func (cfg VolumeUploadConfig) withDefaults() VolumeUploadConfig {
	defaults := DefaultVolumeUploadConfig()
	if cfg.ChunkSize == 0 {
		cfg.ChunkSize = defaults.ChunkSize
	}
	if cfg.MaxPackBytes == 0 {
		cfg.MaxPackBytes = defaults.MaxPackBytes
	}
	if cfg.FileConcurrency <= 0 {
		cfg.FileConcurrency = defaults.FileConcurrency
	}
	if cfg.ChunkConcurrency <= 0 {
		cfg.ChunkConcurrency = defaults.ChunkConcurrency
	}
	if cfg.MaxChunkRetries <= 0 {
		cfg.MaxChunkRetries = defaults.MaxChunkRetries
	}
	if cfg.MaxRateLimitRetries <= 0 {
		cfg.MaxRateLimitRetries = defaults.MaxRateLimitRetries
	}
	if cfg.MaxCommitRetries <= 0 {
		cfg.MaxCommitRetries = defaults.MaxCommitRetries
	}
	return cfg
}

func (cfg VolumeDownloadConfig) withDefaults() VolumeDownloadConfig {
	defaults := DefaultVolumeDownloadConfig()
	if cfg.ChunkConcurrency <= 0 {
		cfg.ChunkConcurrency = defaults.ChunkConcurrency
	}
	if cfg.MaxChunkRetries <= 0 {
		cfg.MaxChunkRetries = defaults.MaxChunkRetries
	}
	return cfg
}

// --- CDP1 data packs --------------------------------------------------------

var (
	cdp1Magic = []byte("CDP1")
	cdp1CBOR  = mustCBORMode()
)

const cdp1PrefixBytes = 8

type dataPackBuilder struct {
	body    []byte
	entries []packRelEntry
	seen    map[[32]byte]struct{}
}

type packRelEntry struct {
	hash   [32]byte
	offset uint64
	length uint32
}

type sealedPackEntry struct {
	hash   [32]byte
	offset uint64
	length uint32
}

type sealedDataPack struct {
	bytes   []byte
	chunkID [32]byte
	entries []sealedPackEntry
}

type packHeader struct {
	Entries      []packHeaderEntry `cbor:"entries"`
	TotalBodyLen uint64            `cbor:"total_body_len"`
}

type packHeaderEntry struct {
	Hash   []byte `cbor:"hash"`
	Length uint32 `cbor:"length"`
	Offset uint64 `cbor:"offset"`
}

func mustCBORMode() cbor.EncMode {
	mode, err := cbor.CanonicalEncOptions().EncMode()
	if err != nil {
		panic(err)
	}
	return mode
}

func newDataPackBuilder() *dataPackBuilder {
	return &dataPackBuilder{seen: make(map[[32]byte]struct{})}
}

func (b *dataPackBuilder) isEmpty() bool {
	return len(b.entries) == 0
}

func (b *dataPackBuilder) objectLen() uint64 {
	header, err := encodePackHeader(b.entries, uint64(len(b.body)))
	if err != nil {
		panic(err)
	}
	return uint64(cdp1PrefixBytes+len(header)) + uint64(len(b.body))
}

func (b *dataPackBuilder) fits(length uint64, maxObjectBytes uint64) bool {
	if length > uint64(^uint32(0)) {
		return false
	}
	entries := append([]packRelEntry(nil), b.entries...)
	entries = append(entries, packRelEntry{
		offset: uint64(len(b.body)),
		length: uint32(length),
	})
	header, err := encodePackHeader(entries, uint64(len(b.body))+length)
	if err != nil {
		return false
	}
	return uint64(cdp1PrefixBytes+len(header))+uint64(len(b.body))+length <= maxObjectBytes
}

func (b *dataPackBuilder) append(hash [32]byte, data []byte) {
	if _, ok := b.seen[hash]; ok {
		return
	}
	b.seen[hash] = struct{}{}
	b.entries = append(b.entries, packRelEntry{
		hash:   hash,
		offset: uint64(len(b.body)),
		length: uint32(len(data)),
	})
	b.body = append(b.body, data...)
}

func (b *dataPackBuilder) seal() (sealedDataPack, error) {
	sort.Slice(b.entries, func(i, j int) bool {
		return bytes.Compare(b.entries[i].hash[:], b.entries[j].hash[:]) < 0
	})
	header, err := encodePackHeader(b.entries, uint64(len(b.body)))
	if err != nil {
		return sealedDataPack{}, err
	}
	bodyStart := uint64(cdp1PrefixBytes + len(header))

	out := make([]byte, 0, cdp1PrefixBytes+len(header)+len(b.body))
	out = append(out, cdp1Magic...)
	out = binary.LittleEndian.AppendUint32(out, uint32(len(header)))
	out = append(out, header...)
	out = append(out, b.body...)

	entries := make([]sealedPackEntry, 0, len(b.entries))
	for _, entry := range b.entries {
		entries = append(entries, sealedPackEntry{
			hash:   entry.hash,
			offset: bodyStart + entry.offset,
			length: entry.length,
		})
	}
	return sealedDataPack{
		bytes:   out,
		chunkID: blake3Sum(out),
		entries: entries,
	}, nil
}

func (p sealedDataPack) entryFor(hash [32]byte) (sealedPackEntry, bool) {
	i := sort.Search(len(p.entries), func(i int) bool {
		return bytes.Compare(p.entries[i].hash[:], hash[:]) >= 0
	})
	if i < len(p.entries) && p.entries[i].hash == hash {
		return p.entries[i], true
	}
	return sealedPackEntry{}, false
}

func encodePackHeader(entries []packRelEntry, totalBodyLen uint64) ([]byte, error) {
	header := packHeader{
		Entries:      make([]packHeaderEntry, 0, len(entries)),
		TotalBodyLen: totalBodyLen,
	}
	for _, entry := range entries {
		hash := make([]byte, len(entry.hash))
		copy(hash, entry.hash[:])
		header.Entries = append(header.Entries, packHeaderEntry{
			Hash:   hash,
			Length: entry.length,
			Offset: entry.offset,
		})
	}
	return cdp1CBOR.Marshal(header)
}

func parseDataPackHeader(data []byte) ([]sealedPackEntry, int, error) {
	if len(data) < cdp1PrefixBytes {
		return nil, 0, fmt.Errorf("pack too short: %d bytes", len(data))
	}
	if !bytes.Equal(data[:len(cdp1Magic)], cdp1Magic) {
		return nil, 0, fmt.Errorf("bad pack magic: %x", data[:len(cdp1Magic)])
	}
	headerLen := int(binary.LittleEndian.Uint32(data[4:8]))
	bodyStart := cdp1PrefixBytes + headerLen
	if bodyStart > len(data) {
		return nil, 0, fmt.Errorf("header length %d exceeds data %d", headerLen, len(data))
	}

	var header packHeader
	if err := cbor.Unmarshal(data[cdp1PrefixBytes:bodyStart], &header); err != nil {
		return nil, 0, err
	}
	entries := make([]sealedPackEntry, 0, len(header.Entries))
	for _, entry := range header.Entries {
		if len(entry.Hash) != 32 {
			return nil, 0, fmt.Errorf("entry hash is %d bytes, want 32", len(entry.Hash))
		}
		var hash [32]byte
		copy(hash[:], entry.Hash)
		entries = append(entries, sealedPackEntry{
			hash:   hash,
			offset: uint64(bodyStart) + entry.Offset,
			length: entry.Length,
		})
	}
	return entries, bodyStart, nil
}
