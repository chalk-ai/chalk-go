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
	"google.golang.org/protobuf/types/known/timestamppb"
)

// VolumeClient provides access to Chalk's versioned volume service.
type VolumeClient interface {
	CreateVolume(ctx context.Context, name string) (*volumev2.CreateVolumeResponse, error)
	GetVolume(ctx context.Context, volume VolumeRef, selector *volumev2.VersionSelector) (*volumev2.GetVolumeResponse, error)
	ListVolumes(ctx context.Context, limit int32, cursor string) (*volumev2.ListVolumesResponse, error)
	DeleteVolume(ctx context.Context, volume VolumeRef) error
	ListVolumeVersions(ctx context.Context, volume VolumeRef, limit int32, cursor string) (*volumev2.ListVolumeVersionsResponse, error)
	ListFiles(ctx context.Context, params ListVolumeFilesParams) (*volumev2.ListFilesResponse, error)
	GetFile(ctx context.Context, volume VolumeRef, path string, selector *volumev2.VersionSelector) (*volumev2.GetFileResponse, error)
	UploadFiles(ctx context.Context, request VolumeUploadRequest, onProgress func(uint64)) ([]*volumev2.CommitStatus, error)
	UploadDirectory(ctx context.Context, volumeName string, dir string, config VolumeUploadConfig) ([]*volumev2.CommitStatus, error)
	RemoveFiles(ctx context.Context, volumeName string, paths []VolumeRemovePath, maxCommitRetries int) (*volumev2.CommitStatus, error)
	DownloadBytes(ctx context.Context, request VolumeDownloadRequest, onProgress func(uint64)) ([]byte, *volumev2.FileInfo, error)
	DownloadToFile(ctx context.Context, request VolumeDownloadRequest, localPath string, onProgress func(uint64)) (*volumev2.FileInfo, error)
	DownloadToDirectory(ctx context.Context, volumeName string, localDir string, selector *volumev2.VersionSelector, config VolumeDownloadConfig) error
}

// VolumeClientConfig configures a versioned volume client.
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

// ListVolumeFilesParams configures ListFiles.
type ListVolumeFilesParams struct {
	Volume    VolumeRef
	Path      string
	Recursive bool
	Limit     int32
	Cursor    string
	Selector  *volumev2.VersionSelector
}

// VolumeUploadConfig controls upload batching and object-store concurrency.
type VolumeUploadConfig struct {
	ChunkSize           uint64
	MaxPackBytes        uint64
	FileConcurrency     int
	ChunkConcurrency    int
	BatchSize           int
	MaxChunkRetries     int
	MaxRateLimitRetries int
	MaxCommitRetries    int
}

// DefaultVolumeUploadConfig returns upload settings matching the reference client.
func DefaultVolumeUploadConfig() VolumeUploadConfig {
	return VolumeUploadConfig{
		ChunkSize:           8 * 1024 * 1024,
		MaxPackBytes:        16 * 1024 * 1024,
		FileConcurrency:     4,
		ChunkConcurrency:    32,
		BatchSize:           500,
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
}

// VolumeUploadFile is one file to commit into a volume.
type VolumeUploadFile struct {
	Path     string
	Content  VolumeUploadContent
	Metadata *volumev2.FileMetadata
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

	rpc volumev2connect.VolumeServiceClient
}

// NewVolumeClient creates a versioned volume client using Chalk auth/config resolution.
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
		connect.WithInterceptors(connect.UnaryInterceptorFunc(client.authInterceptor())),
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
	header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	return nil
}

func rpcMsg[T any](res *connect.Response[T], err error) (*T, error) {
	if err != nil {
		return nil, err
	}
	return res.Msg, nil
}

func (c *volumeClientImpl) CreateVolume(ctx context.Context, name string) (*volumev2.CreateVolumeResponse, error) {
	return rpcMsg(c.rpc.CreateVolume(ctx, connect.NewRequest(&volumev2.CreateVolumeRequest{Name: name})))
}

func (c *volumeClientImpl) GetVolume(ctx context.Context, volume VolumeRef, selector *volumev2.VersionSelector) (*volumev2.GetVolumeResponse, error) {
	return rpcMsg(c.rpc.GetVolume(ctx, connect.NewRequest(&volumev2.GetVolumeRequest{
		Volume:   volume.toProto(),
		Selector: selector,
	})))
}

func (c *volumeClientImpl) ListVolumes(ctx context.Context, limit int32, cursor string) (*volumev2.ListVolumesResponse, error) {
	return rpcMsg(c.rpc.ListVolumes(ctx, connect.NewRequest(&volumev2.ListVolumesRequest{
		Limit:  limit,
		Cursor: cursor,
	})))
}

func (c *volumeClientImpl) DeleteVolume(ctx context.Context, volume VolumeRef) error {
	_, err := c.rpc.DeleteVolume(ctx, connect.NewRequest(&volumev2.DeleteVolumeRequest{Volume: volume.toProto()}))
	return err
}

func (c *volumeClientImpl) ListVolumeVersions(ctx context.Context, volume VolumeRef, limit int32, cursor string) (*volumev2.ListVolumeVersionsResponse, error) {
	return rpcMsg(c.rpc.ListVolumeVersions(ctx, connect.NewRequest(&volumev2.ListVolumeVersionsRequest{
		Volume: volume.toProto(),
		Limit:  limit,
		Cursor: cursor,
	})))
}

func (c *volumeClientImpl) ListFiles(ctx context.Context, params ListVolumeFilesParams) (*volumev2.ListFilesResponse, error) {
	return rpcMsg(c.rpc.ListFiles(ctx, connect.NewRequest(&volumev2.ListFilesRequest{
		Volume:    params.Volume.toProto(),
		Path:      params.Path,
		Recursive: params.Recursive,
		Limit:     params.Limit,
		Cursor:    params.Cursor,
		Selector:  params.Selector,
	})))
}

func (c *volumeClientImpl) GetFile(ctx context.Context, volume VolumeRef, path string, selector *volumev2.VersionSelector) (*volumev2.GetFileResponse, error) {
	return rpcMsg(c.rpc.GetFile(ctx, connect.NewRequest(&volumev2.GetFileRequest{
		Volume:   volume.toProto(),
		Path:     path,
		Selector: selector,
	})))
}

func (v VolumeRef) toProto() *volumev2.VolumeRef {
	return &volumev2.VolumeRef{VolumeId: v.ID, Name: v.Name}
}

func VolumeVersionSelector(versionID uint64) *volumev2.VersionSelector {
	return &volumev2.VersionSelector{Selector: &volumev2.VersionSelector_VersionId{VersionId: versionID}}
}

func VolumeRefSelector(ref string) *volumev2.VersionSelector {
	return &volumev2.VersionSelector{Selector: &volumev2.VersionSelector_Ref{Ref: ref}}
}

func (c *volumeClientImpl) UploadFiles(ctx context.Context, request VolumeUploadRequest, onProgress func(uint64)) ([]*volumev2.CommitStatus, error) {
	cfg := request.Config.withDefaults()
	if onProgress == nil {
		onProgress = func(uint64) {}
	}
	volume := &volumev2.VolumeRef{Name: request.VolumeName}
	var statuses []*volumev2.CommitStatus
	for start := 0; start < len(request.Files); start += cfg.BatchSize {
		end := min(start+cfg.BatchSize, len(request.Files))
		uploaded, err := c.uploadBatchFiles(ctx, volume, request.Files[start:end], cfg, onProgress)
		if err != nil {
			return nil, err
		}
		status, err := c.commitPathDeltas(ctx, volume, uploaded, nil, cfg.MaxCommitRetries)
		if err != nil {
			return nil, err
		}
		statuses = append(statuses, status)
	}
	return statuses, nil
}

func (c *volumeClientImpl) UploadDirectory(ctx context.Context, volumeName string, dir string, config VolumeUploadConfig) ([]*volumev2.CommitStatus, error) {
	files, err := collectVolumeLocalFiles(dir)
	if err != nil {
		return nil, err
	}
	return c.UploadFiles(ctx, VolumeUploadRequest{VolumeName: volumeName, Files: files, Config: config}, nil)
}

func (c *volumeClientImpl) RemoveFiles(ctx context.Context, volumeName string, paths []VolumeRemovePath, maxCommitRetries int) (*volumev2.CommitStatus, error) {
	removes := make([]*volumev2.PathRemoveDelta, 0, len(paths))
	for _, path := range paths {
		removes = append(removes, &volumev2.PathRemoveDelta{Path: path.Path, Recursive: path.Recursive})
	}
	if maxCommitRetries <= 0 {
		maxCommitRetries = DefaultVolumeUploadConfig().MaxCommitRetries
	}
	return c.commitPathDeltas(ctx, &volumev2.VolumeRef{Name: volumeName}, nil, removes, maxCommitRetries)
}

type uploadedVolumeFile struct {
	path            string
	contentRef      *volumev2.ContentRef
	metadata        *volumev2.FileMetadata
	uploadedObjects []*volumev2.UploadedObjectReference
}

type sizedVolumeUploadFile struct {
	file VolumeUploadFile
	size uint64
}

func (c *volumeClientImpl) uploadBatchFiles(ctx context.Context, volume *volumev2.VolumeRef, files []VolumeUploadFile, cfg VolumeUploadConfig, onProgress func(uint64)) ([]uploadedVolumeFile, error) {
	sized := make([]sizedVolumeUploadFile, len(files))
	eg, egCtx := errgroup.WithContext(ctx)
	eg.SetLimit(maxInt(cfg.FileConcurrency, 1))
	for i := range files {
		eg.Go(func() error {
			size, err := files[i].Content.size()
			if err != nil {
				return err
			}
			sized[i] = sizedVolumeUploadFile{file: files[i], size: size}
			return egCtx.Err()
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

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
	var out []uploadedVolumeFile
	eg, egCtx = errgroup.WithContext(ctx)
	eg.SetLimit(maxInt(cfg.FileConcurrency, 1))
	for _, sf := range perFile {
		eg.Go(func() error {
			uploaded, err := c.uploadOneFile(egCtx, volume, sf.file, sf.size, cfg, onProgress)
			if err != nil {
				return err
			}
			mu.Lock()
			out = append(out, uploaded)
			mu.Unlock()
			return nil
		})
	}
	eg.Go(func() error {
		uploaded, err := c.packAndUpload(egCtx, volume, packable, cfg, onProgress)
		if err != nil {
			return err
		}
		mu.Lock()
		out = append(out, uploaded...)
		mu.Unlock()
		return nil
	})
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeClientImpl) uploadOneFile(ctx context.Context, volume *volumev2.VolumeRef, file VolumeUploadFile, size uint64, cfg VolumeUploadConfig, onProgress func(uint64)) (uploadedVolumeFile, error) {
	metadata, err := file.metadata()
	if err != nil {
		return uploadedVolumeFile{}, err
	}
	if size == 0 {
		onProgress(0)
		return uploadedVolumeFile{
			path:       file.Path,
			contentRef: emptyVolumeContentRef(),
			metadata:   metadata,
		}, nil
	}

	slices := computeVolumeSlices(size, cfg.ChunkSize)
	prepared := make([]preparedVolumeSlice, len(slices))
	eg, egCtx := errgroup.WithContext(ctx)
	eg.SetLimit(maxInt(cfg.ChunkConcurrency, 1))
	for i := range slices {
		eg.Go(func() error {
			bytes, err := file.Content.readChunk(slices[i].offset, slices[i].size)
			if err != nil {
				return err
			}
			prepared[i] = preparedVolumeSlice{
				slice: slices[i],
				bytes: bytes,
				hash:  blake3Hex(bytes),
			}
			return egCtx.Err()
		})
	}
	if err := eg.Wait(); err != nil {
		return uploadedVolumeFile{}, err
	}

	objects := make([]*volumev2.UploadedObjectReference, 0, len(prepared))
	for _, p := range prepared {
		objects = append(objects, uploadedVolumeObjectRef(chunkRelativeObjectKey(p.hash), p.hash, p.slice.size, volumev2.UploadedObjectKind_UPLOADED_OBJECT_KIND_CHUNK))
	}
	urls, err := c.requestUploadURLs(ctx, volume, objects)
	if err != nil {
		return uploadedVolumeFile{}, err
	}
	if len(urls) != len(prepared) {
		return uploadedVolumeFile{}, fmt.Errorf("RequestUploadURLs returned %d items for %d chunks", len(urls), len(prepared))
	}

	var fileHasher blake3.Hasher
	var uploadedObjects []*volumev2.UploadedObjectReference
	chunks := make([]*volumev2.ChunkRef, len(prepared))
	for _, p := range prepared {
		fileHasher.Write(p.bytes)
	}
	fileHash := hex.EncodeToString(fileHasher.Sum(nil))

	eg, egCtx = errgroup.WithContext(ctx)
	eg.SetLimit(maxInt(cfg.ChunkConcurrency, 1))
	for i := range prepared {
		eg.Go(func() error {
			p := prepared[i]
			url := urls[i]
			object := uploadedVolumeObjectRef(url.ObjectKey, p.hash, p.slice.size, volumev2.UploadedObjectKind_UPLOADED_OBJECT_KIND_CHUNK)
			if !url.AlreadyExists {
				if err := c.signedRequestWithRetry(egCtx, http.MethodPut, url.SignedUploadUri, octetStreamHeaders(), bytes.NewReader(p.bytes), cfg.MaxChunkRetries, cfg.MaxRateLimitRetries, func(ctx context.Context) (string, error) {
					refreshed, err := c.requestUploadURLs(ctx, volume, []*volumev2.UploadedObjectReference{object})
					if err != nil {
						return "", err
					}
					if len(refreshed) == 0 {
						return "", fmt.Errorf("URL refresh returned no items")
					}
					return refreshed[0].SignedUploadUri, nil
				}); err != nil {
					return err
				}
			}
			onProgress(p.slice.size)
			chunks[i] = &volumev2.ChunkRef{ObjectKey: object.ObjectKey, Hash: p.hash, Size: p.slice.size, Offset: p.slice.offset}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return uploadedVolumeFile{}, err
	}
	for i := range prepared {
		p := prepared[i]
		uploadedObjects = append(uploadedObjects, uploadedVolumeObjectRef(urls[i].ObjectKey, p.hash, p.slice.size, volumev2.UploadedObjectKind_UPLOADED_OBJECT_KIND_CHUNK))
	}
	return uploadedVolumeFile{
		path:            file.Path,
		contentRef:      chunkedVolumeContentRef(fileHash, size, chunks),
		metadata:        metadata,
		uploadedObjects: uploadedObjects,
	}, nil
}

type preparedVolumeSlice struct {
	slice volumeSlice
	bytes []byte
	hash  string
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

func (c *volumeClientImpl) packAndUpload(ctx context.Context, volume *volumev2.VolumeRef, files []sizedVolumeUploadFile, cfg VolumeUploadConfig, onProgress func(uint64)) ([]uploadedVolumeFile, error) {
	if len(files) == 0 {
		return nil, nil
	}
	var out []uploadedVolumeFile
	builder := newDataPackBuilder()
	var members []packMember
	for _, sf := range files {
		data, err := sf.file.Content.readChunk(0, sf.size)
		if err != nil {
			return nil, err
		}
		hash := blake3Sum(data)
		if !builder.isEmpty() && !builder.fits(sf.size, cfg.MaxPackBytes) {
			uploaded, err := c.uploadOnePack(ctx, volume, builder, members, cfg, onProgress)
			if err != nil {
				return nil, err
			}
			out = append(out, uploaded...)
			builder = newDataPackBuilder()
			members = nil
		}
		metadata, err := sf.file.metadata()
		if err != nil {
			return nil, err
		}
		builder.append(hash, data)
		members = append(members, packMember{path: sf.file.Path, metadata: metadata, hash: hash})
	}
	if !builder.isEmpty() {
		uploaded, err := c.uploadOnePack(ctx, volume, builder, members, cfg, onProgress)
		if err != nil {
			return nil, err
		}
		out = append(out, uploaded...)
	}
	return out, nil
}

type packMember struct {
	path     string
	metadata *volumev2.FileMetadata
	hash     [32]byte
}

func (c *volumeClientImpl) uploadOnePack(ctx context.Context, volume *volumev2.VolumeRef, builder *dataPackBuilder, members []packMember, cfg VolumeUploadConfig, onProgress func(uint64)) ([]uploadedVolumeFile, error) {
	sealed, err := builder.seal()
	if err != nil {
		return nil, err
	}
	packID := hashHex(sealed.chunkID)
	packObject := uploadedVolumeObjectRef(chunkRelativeObjectKey(packID), packID, uint64(len(sealed.bytes)), volumev2.UploadedObjectKind_UPLOADED_OBJECT_KIND_PACK)
	urls, err := c.requestUploadURLs(ctx, volume, []*volumev2.UploadedObjectReference{packObject})
	if err != nil {
		return nil, err
	}
	if len(urls) == 0 {
		return nil, fmt.Errorf("RequestUploadURLs returned no item for pack")
	}
	if !urls[0].AlreadyExists {
		if err := c.signedRequestWithRetry(ctx, http.MethodPut, urls[0].SignedUploadUri, octetStreamHeaders(), bytes.NewReader(sealed.bytes), cfg.MaxChunkRetries, cfg.MaxRateLimitRetries, func(ctx context.Context) (string, error) {
			refreshed, err := c.requestUploadURLs(ctx, volume, []*volumev2.UploadedObjectReference{packObject})
			if err != nil {
				return "", err
			}
			if len(refreshed) == 0 {
				return "", fmt.Errorf("URL refresh returned no items")
			}
			return refreshed[0].SignedUploadUri, nil
		}); err != nil {
			return nil, err
		}
	}
	onProgress(uint64(len(sealed.bytes)))

	out := make([]uploadedVolumeFile, 0, len(members))
	for _, member := range members {
		entry, ok := sealed.entryFor(member.hash)
		if !ok {
			return nil, fmt.Errorf("sealed pack is missing a member entry")
		}
		out = append(out, uploadedVolumeFile{
			path:            member.path,
			contentRef:      packedVolumeContentRef(packObject.ObjectKey, packID, entry.offset, entry.length, hashHex(member.hash)),
			metadata:        member.metadata,
			uploadedObjects: []*volumev2.UploadedObjectReference{packObject},
		})
	}
	return out, nil
}

func (c *volumeClientImpl) requestUploadURLs(ctx context.Context, volume *volumev2.VolumeRef, objects []*volumev2.UploadedObjectReference) ([]*volumev2.UploadURLItem, error) {
	res, err := c.rpc.RequestUploadURLs(ctx, connect.NewRequest(&volumev2.RequestUploadURLsRequest{
		Volume:  volume,
		Objects: objects,
	}))
	if err != nil {
		return nil, err
	}
	return res.Msg.GetUrls(), nil
}

func (c *volumeClientImpl) commitPathDeltas(ctx context.Context, volume *volumev2.VolumeRef, upserts []uploadedVolumeFile, removes []*volumev2.PathRemoveDelta, maxCommitRetries int) (*volumev2.CommitStatus, error) {
	deltas := make([]*volumev2.PathFileDelta, 0, len(upserts))
	for _, file := range upserts {
		if file.contentRef.GetInline() != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("inline content is only valid for symlinks, not regular file %q", file.path))
		}
		deltas = append(deltas, pathFileDelta(file))
	}
	refs := dedupeVolumeObjectRefs(upserts)
	author := c.resolveCommitAuthor(ctx)
	if maxCommitRetries <= 0 {
		maxCommitRetries = 1
	}
	ref := "main"
	for attempt := 0; attempt < maxCommitRetries; attempt++ {
		base, err := c.rpc.GetVolume(ctx, connect.NewRequest(&volumev2.GetVolumeRequest{Volume: volume}))
		if err != nil {
			return nil, err
		}
		version := base.Msg.GetVersion()
		if version == nil {
			return nil, fmt.Errorf("GetVolume returned no version on the volume tip")
		}
		baseVersionID := version.VersionId
		baseSequenceNumber := version.SequenceNumber
		commitID := newVolumeCommitID()
		res, err := c.rpc.CommitVersion(ctx, connect.NewRequest(&volumev2.CommitVersionRequest{
			Intent: &volumev2.CommitIntent{
				Volume:                   volume,
				CommitId:                 commitID,
				Ref:                      &ref,
				BaseVersionId:            &baseVersionID,
				BaseSequenceNumber:       &baseSequenceNumber,
				UploadedObjectReferences: refs,
				Author:                   author,
				Deltas: &volumev2.CommitIntent_PathDeltas{PathDeltas: &volumev2.PathDeltaList{
					Upserts: deltas,
					Removes: removes,
				}},
			},
		}))
		if err != nil {
			return nil, err
		}
		status := res.Msg.GetStatus()
		if status == nil {
			return nil, fmt.Errorf("CommitVersion returned no status")
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

func (c *volumeClientImpl) signedRequestWithRetry(ctx context.Context, method string, url string, headers http.Header, body io.ReadSeeker, maxAttempts int, maxRateLimitRetries int, refresh func(context.Context) (string, error)) error {
	if maxAttempts <= 0 {
		maxAttempts = 1
	}
	var lastErr error
	refreshed := false
	rateAttempts := 0
	for attempt := 0; attempt < maxAttempts; {
		if body != nil {
			if _, err := body.Seek(0, io.SeekStart); err != nil {
				return err
			}
		}
		req, err := http.NewRequestWithContext(ctx, method, url, body)
		if err != nil {
			return err
		}
		req.Header = cloneHeader(headers)
		resp, err := c.httpClient.Do(req)
		if err == nil && resp != nil {
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
		}
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return nil
		}
		if err == nil && (resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden) && !refreshed {
			refreshed = true
			url, err = refresh(ctx)
			if err != nil {
				return err
			}
			continue
		}
		if err == nil && (resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusServiceUnavailable) {
			if rateAttempts >= maxRateLimitRetries {
				return fmt.Errorf("%s signed URL request failed: HTTP %d after %d rate-limit backoffs", method, resp.StatusCode, rateAttempts)
			}
			time.Sleep(time.Duration(rateLimitBackoffMS(rateAttempts)) * time.Millisecond)
			rateAttempts++
			continue
		}
		if err != nil {
			lastErr = err
		} else {
			lastErr = fmt.Errorf("%s signed URL request failed: HTTP %d", method, resp.StatusCode)
		}
		attempt++
		if attempt < maxAttempts {
			time.Sleep(time.Duration(250*attempt) * time.Millisecond)
		}
	}
	return lastErr
}

func (c *volumeClientImpl) signedGetWithRetry(ctx context.Context, url string, headers http.Header, maxAttempts int, refresh func(context.Context) (string, error)) ([]byte, error) {
	if maxAttempts <= 0 {
		maxAttempts = 1
	}
	var lastErr error
	refreshed := false
	for attempt := 0; attempt < maxAttempts; {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
		req.Header = cloneHeader(headers)
		resp, err := c.httpClient.Do(req)
		if err == nil && resp != nil {
			defer resp.Body.Close()
		}
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return io.ReadAll(resp.Body)
		}
		if err == nil && (resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden) && !refreshed {
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
			refreshed = true
			url, err = refresh(ctx)
			if err != nil {
				return nil, err
			}
			continue
		}
		if err != nil {
			lastErr = err
		} else {
			io.Copy(io.Discard, resp.Body)
			lastErr = fmt.Errorf("GET signed URL request failed: HTTP %d", resp.StatusCode)
		}
		attempt++
		if attempt < maxAttempts {
			time.Sleep(time.Duration(250*attempt) * time.Millisecond)
		}
	}
	return nil, lastErr
}

func (c *volumeClientImpl) DownloadBytes(ctx context.Context, request VolumeDownloadRequest, onProgress func(uint64)) ([]byte, *volumev2.FileInfo, error) {
	var mu sync.Mutex
	var out []byte
	info, err := c.downloadFile(ctx, request, func(offset uint64, data []byte) error {
		mu.Lock()
		defer mu.Unlock()
		end := int(offset) + len(data)
		if len(out) < end {
			next := make([]byte, end)
			copy(next, out)
			out = next
		}
		copy(out[int(offset):end], data)
		return nil
	}, onProgress)
	if err != nil {
		return nil, nil, err
	}
	if info != nil && uint64(len(out)) > info.Size {
		out = out[:info.Size]
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
	defer file.Close()
	var mu sync.Mutex
	info, err := c.downloadFile(ctx, request, func(offset uint64, data []byte) error {
		mu.Lock()
		defer mu.Unlock()
		_, err := file.WriteAt(data, int64(offset))
		return err
	}, onProgress)
	if err != nil {
		return nil, err
	}
	if info != nil {
		applyVolumeFileInfoMetadata(localPath, info)
	}
	return info, nil
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
				applyVolumeFileInfoMetadata(localPath, info)
			}
			return nil
		})
	}
	return eg.Wait()
}

func (c *volumeClientImpl) downloadFile(ctx context.Context, request VolumeDownloadRequest, onWrite func(uint64, []byte) error, onProgress func(uint64)) (*volumev2.FileInfo, error) {
	cfg := request.Config.withDefaults()
	if onProgress == nil {
		onProgress = func(uint64) {}
	}
	fileReq := &volumev2.GetFileRequest{
		Volume:   (&VolumeRef{Name: request.VolumeName}).toProto(),
		Path:     request.Path,
		Selector: request.Selector,
	}
	resp, err := c.rpc.GetFile(ctx, connect.NewRequest(fileReq))
	if err != nil {
		return nil, err
	}
	info := resp.Msg.GetFile()
	switch content := resp.Msg.GetContent().(type) {
	case *volumev2.GetFileResponse_Data:
		if err := onWrite(0, content.Data); err != nil {
			return nil, err
		}
		onProgress(uint64(len(content.Data)))
	case *volumev2.GetFileResponse_Chunked:
		pinned, err := pinnedVolumeFileRequest(fileReq, resp.Msg.GetVersion())
		if err != nil {
			return nil, err
		}
		eg, egCtx := errgroup.WithContext(ctx)
		eg.SetLimit(maxInt(cfg.ChunkConcurrency, 1))
		for _, chunk := range content.Chunked.Chunks {
			eg.Go(func() error {
				data, err := c.signedGetWithRetry(egCtx, chunk.SignedDownloadUri, nil, cfg.MaxChunkRetries, func(ctx context.Context) (string, error) {
					res, err := c.rpc.GetFile(ctx, connect.NewRequest(pinned))
					if err != nil {
						return "", err
					}
					chunked := res.Msg.GetChunked()
					if chunked == nil {
						return "", fmt.Errorf("URL refresh returned non-chunked content")
					}
					for _, candidate := range chunked.Chunks {
						if candidate.Offset == chunk.Offset && (chunk.Hash == "" || candidate.Hash == chunk.Hash) {
							return candidate.SignedDownloadUri, nil
						}
					}
					return "", fmt.Errorf("URL refresh returned no matching chunk at offset %d", chunk.Offset)
				})
				if err != nil {
					return err
				}
				if uint64(len(data)) != chunk.Size {
					return fmt.Errorf("download chunk at offset %d returned %d bytes, expected %d", chunk.Offset, len(data), chunk.Size)
				}
				if chunk.Hash != "" && blake3Hex(data) != chunk.Hash {
					return connect.NewError(connect.CodeDataLoss, fmt.Errorf("download chunk at offset %d failed hash validation", chunk.Offset))
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
		pinned, err := pinnedVolumeFileRequest(fileReq, resp.Msg.GetVersion())
		if err != nil {
			return nil, err
		}
		end, ok := checkedInclusiveRangeEnd(pack.Offset, pack.Size)
		if !ok {
			return nil, fmt.Errorf("packed range overflows uint64")
		}
		headers := http.Header{"Range": []string{fmt.Sprintf("bytes=%d-%d", pack.Offset, end)}}
		data, err := c.signedGetWithRetry(ctx, pack.SignedDownloadUri, headers, cfg.MaxChunkRetries, func(ctx context.Context) (string, error) {
			res, err := c.rpc.GetFile(ctx, connect.NewRequest(pinned))
			if err != nil {
				return "", err
			}
			refreshed := res.Msg.GetPacked()
			if refreshed == nil || refreshed.GetPack() == nil {
				return "", fmt.Errorf("URL refresh returned non-packed content")
			}
			return refreshed.GetPack().SignedDownloadUri, nil
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
			res, err := c.ListFiles(ctx, ListVolumeFilesParams{
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
		end := offset + length
		if end > uint64(len(c.bytes)) {
			return nil, fmt.Errorf("read_chunk out of range: %d+%d > %d", offset, length, len(c.bytes))
		}
		return append([]byte(nil), c.bytes[offset:end]...), nil
	}
	file, err := os.Open(c.localPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buf := make([]byte, length)
	_, err = file.ReadAt(buf, int64(offset))
	if err != nil && err != io.EOF {
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

func applyVolumeFileInfoMetadata(path string, info *volumev2.FileInfo) {
	if info == nil {
		return
	}
	if runtime.GOOS != "windows" && info.Mode != nil {
		_ = os.Chmod(path, os.FileMode(info.GetMode()&0o7777))
	}
	if info.UpdatedAt != nil {
		t := info.UpdatedAt.AsTime()
		_ = os.Chtimes(path, t, t)
	}
}

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

func pathFileDelta(file uploadedVolumeFile) *volumev2.PathFileDelta {
	return &volumev2.PathFileDelta{
		Path: file.path,
		Node: &volumev2.FileNode{
			Metadata: file.metadata,
			Node: &volumev2.FileNode_File{File: &volumev2.RegularFileNode{
				Content: file.contentRef,
			}},
		},
		Mode: volumev2.PathWriteMode_PATH_WRITE_MODE_UPSERT,
	}
}

func uploadedVolumeObjectRef(objectKey string, hash string, size uint64, kind volumev2.UploadedObjectKind) *volumev2.UploadedObjectReference {
	return &volumev2.UploadedObjectReference{ObjectKey: objectKey, Hash: hash, ContentSize: size, Kind: kind}
}

func dedupeVolumeObjectRefs(files []uploadedVolumeFile) []*volumev2.UploadedObjectReference {
	var refs []*volumev2.UploadedObjectReference
	for _, file := range files {
		refs = append(refs, file.uploadedObjects...)
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

func newVolumeCommitID() string {
	var raw [16]byte
	if _, err := rand.Read(raw[:]); err != nil {
		return fmt.Sprintf("commit-%d", time.Now().UnixNano())
	}
	return "commit-" + hex.EncodeToString(raw[:])
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

func checkedInclusiveRangeEnd(offset uint64, size uint64) (uint64, bool) {
	if size == 0 {
		return offset, true
	}
	end := offset + size - 1
	return end, end >= offset
}

func (cfg VolumeUploadConfig) withDefaults() VolumeUploadConfig {
	defaults := DefaultVolumeUploadConfig()
	if cfg.ChunkSize == 0 {
		cfg.ChunkSize = defaults.ChunkSize
	}
	if cfg.MaxPackBytes == 0 {
		cfg.MaxPackBytes = defaults.MaxPackBytes
	}
	if cfg.FileConcurrency == 0 {
		cfg.FileConcurrency = defaults.FileConcurrency
	}
	if cfg.ChunkConcurrency == 0 {
		cfg.ChunkConcurrency = defaults.ChunkConcurrency
	}
	if cfg.BatchSize == 0 {
		cfg.BatchSize = defaults.BatchSize
	}
	if cfg.MaxChunkRetries == 0 {
		cfg.MaxChunkRetries = defaults.MaxChunkRetries
	}
	if cfg.MaxRateLimitRetries == 0 {
		cfg.MaxRateLimitRetries = defaults.MaxRateLimitRetries
	}
	if cfg.MaxCommitRetries == 0 {
		cfg.MaxCommitRetries = defaults.MaxCommitRetries
	}
	return cfg
}

func (cfg VolumeDownloadConfig) withDefaults() VolumeDownloadConfig {
	defaults := DefaultVolumeDownloadConfig()
	if cfg.ChunkConcurrency == 0 {
		cfg.ChunkConcurrency = defaults.ChunkConcurrency
	}
	if cfg.MaxChunkRetries == 0 {
		cfg.MaxChunkRetries = defaults.MaxChunkRetries
	}
	return cfg
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

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
