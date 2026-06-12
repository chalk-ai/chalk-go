package chalk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"connectrpc.com/connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	volumev2 "github.com/chalk-ai/chalk-go/gen/chalk/volume/v2"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestVolumeAuthHeadersUseConfiguredToken(t *testing.T) {
	envID := "env-123"
	client, err := NewVolumeClient(context.Background(), &VolumeClientConfig{
		ApiServer:                  "https://api.chalk.ai",
		ClientId:                   "client",
		ClientSecret:               "secret",
		EnvironmentId:              envID,
		SkipEnvironmentNameMapping: true,
		JWT: &serverv1.GetTokenResponse{
			AccessToken:         "token-abc",
			ExpiresAt:           timestamppb.New(time.Now().Add(time.Hour)),
			EnvironmentIdToName: map[string]string{envID: "test"},
		},
	})
	require.NoError(t, err)

	header := http.Header{}
	require.NoError(t, client.(*volumeClientImpl).addAuthHeaders(context.Background(), header))
	require.Equal(t, "Bearer token-abc", header.Get("Authorization"))
	require.Equal(t, envID, header.Get("x-chalk-env-id"))
	require.Equal(t, "go-api", header.Get("x-chalk-server"))
	require.NotEmpty(t, header.Get("User-Agent"))
}

func TestVolumeUploadContentCopiesBytes(t *testing.T) {
	src := []byte("abcdef")
	content := VolumeUploadBytes(src)
	src[0] = 'z'

	size, err := content.size()
	require.NoError(t, err)
	require.Equal(t, uint64(6), size)

	chunk, err := content.readChunk(1, 3)
	require.NoError(t, err)
	require.Equal(t, []byte("bcd"), chunk)

	_, err = content.readChunk(4, 10)
	require.Error(t, err)
}

func TestVolumeSlicesCoverRange(t *testing.T) {
	require.Equal(t, []volumeSlice{
		{offset: 0, size: 10},
		{offset: 10, size: 10},
		{offset: 20, size: 5},
	}, computeVolumeSlices(25, 10))
	require.Empty(t, computeVolumeSlices(0, 10))
	require.Equal(t, []volumeSlice{
		{offset: 0, size: 1},
		{offset: 1, size: 1},
		{offset: 2, size: 1},
	}, computeVolumeSlices(3, 0))
}

func TestVolumeCommitRetriesRebaseAndDedupesRefs(t *testing.T) {
	rpc := &fakeVolumeRPC{
		getVolume: func(context.Context, *connect.Request[volumev2.GetVolumeRequest]) (*connect.Response[volumev2.GetVolumeResponse], error) {
			return connect.NewResponse(&volumev2.GetVolumeResponse{
				Version: &volumev2.VersionInfo{VersionId: 7, SequenceNumber: 11},
			}), nil
		},
	}
	rpc.commitVersion = func(_ context.Context, req *connect.Request[volumev2.CommitVersionRequest]) (*connect.Response[volumev2.CommitVersionResponse], error) {
		rpc.commits = append(rpc.commits, req.Msg.GetIntent())
		result := volumev2.CommitResult_COMMIT_RESULT_REBASE_REQUIRED
		if len(rpc.commits) == 2 {
			result = volumev2.CommitResult_COMMIT_RESULT_COMMITTED
		}
		return connect.NewResponse(&volumev2.CommitVersionResponse{
			Status: &volumev2.CommitStatus{CommitId: req.Msg.GetIntent().GetCommitId(), Result: result},
		}), nil
	}

	object := uploadedVolumeObjectRef("aa/object", "hash", 12, volumev2.UploadedObjectKind_UPLOADED_OBJECT_KIND_CHUNK)
	client := &volumeClientImpl{rpc: rpc, author: "chalk:env:agent:test"}
	status, err := client.commitPathDeltas(context.Background(), &volumev2.VolumeRef{Name: "models"}, []uploadedVolumeFile{
		{path: "a.bin", contentRef: emptyVolumeContentRef(), uploadedObjects: []*volumev2.UploadedObjectReference{object}},
		{path: "b.bin", contentRef: emptyVolumeContentRef(), uploadedObjects: []*volumev2.UploadedObjectReference{object}},
	}, nil, 2)

	require.NoError(t, err)
	require.Equal(t, volumev2.CommitResult_COMMIT_RESULT_COMMITTED, status.Result)
	require.Len(t, rpc.commits, 2)
	require.NotEmpty(t, rpc.commits[0].CommitId)
	require.NotEqual(t, rpc.commits[0].CommitId, rpc.commits[1].CommitId)
	require.Equal(t, uint64(7), rpc.commits[1].GetBaseVersionId())
	require.Equal(t, uint64(11), rpc.commits[1].GetBaseSequenceNumber())
	require.Equal(t, "chalk:env:agent:test", rpc.commits[1].Author)
	require.Len(t, rpc.commits[1].UploadedObjectReferences, 1)
	require.Len(t, rpc.commits[1].GetPathDeltas().Upserts, 2)
}

func TestVolumeDownloadInlineBytes(t *testing.T) {
	client := &volumeClientImpl{rpc: &fakeVolumeRPC{
		getFile: func(context.Context, *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			return connect.NewResponse(&volumev2.GetFileResponse{
				File:    &volumev2.FileInfo{Path: "model.bin", Size: 5, Hash: blake3Hex([]byte("hello"))},
				Content: &volumev2.GetFileResponse_Data{Data: []byte("hello")},
			}), nil
		},
	}}
	var progress uint64
	got, info, err := client.DownloadBytes(context.Background(), VolumeDownloadRequest{
		VolumeName: "models",
		Path:       "model.bin",
	}, func(n uint64) { progress += n })

	require.NoError(t, err)
	require.Equal(t, []byte("hello"), got)
	require.Equal(t, "model.bin", info.Path)
	require.Equal(t, uint64(5), progress)
}

func TestVolumeDownloadChunkedRefreshesURL(t *testing.T) {
	body := []byte("chunk-data")
	var objectHits int
	objectServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		objectHits++
		if objectHits == 1 {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		_, _ = w.Write(body)
	}))
	defer objectServer.Close()

	var getFileCalls int
	client := &volumeClientImpl{httpClient: objectServer.Client(), rpc: &fakeVolumeRPC{
		getFile: func(context.Context, *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			getFileCalls++
			return connect.NewResponse(&volumev2.GetFileResponse{
				File:    &volumev2.FileInfo{Path: "chunk.bin", Size: uint64(len(body))},
				Version: &volumev2.VersionInfo{VersionId: 42},
				Content: &volumev2.GetFileResponse_Chunked{Chunked: &volumev2.ChunkedFileContent{Chunks: []*volumev2.SignedChunkRef{{
					SignedDownloadUri: objectServer.URL,
					Offset:            0,
					Size:              uint64(len(body)),
					Hash:              blake3Hex(body),
				}}}},
			}), nil
		},
	}}

	got, _, err := client.DownloadBytes(context.Background(), VolumeDownloadRequest{
		VolumeName: "models",
		Path:       "chunk.bin",
		Config:     VolumeDownloadConfig{ChunkConcurrency: 1, MaxChunkRetries: 2},
	}, nil)

	require.NoError(t, err)
	require.Equal(t, body, got)
	require.Equal(t, 2, objectHits)
	require.Equal(t, 2, getFileCalls)
}

func TestVolumeDownloadPackedUsesRange(t *testing.T) {
	body := []byte("pack member")
	var gotRange string
	objectServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotRange = r.Header.Get("Range")
		_, _ = w.Write(body)
	}))
	defer objectServer.Close()

	client := &volumeClientImpl{httpClient: objectServer.Client(), rpc: &fakeVolumeRPC{
		getFile: func(context.Context, *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			return connect.NewResponse(&volumev2.GetFileResponse{
				File:    &volumev2.FileInfo{Path: "packed.bin", Size: uint64(len(body)), Hash: blake3Hex(body)},
				Version: &volumev2.VersionInfo{VersionId: 9},
				Content: &volumev2.GetFileResponse_Packed{Packed: &volumev2.PackedFileContent{Pack: &volumev2.SignedPackEntryRef{
					SignedDownloadUri: objectServer.URL,
					Offset:            7,
					Size:              uint64(len(body)),
				}}},
			}), nil
		},
	}}

	got, _, err := client.DownloadBytes(context.Background(), VolumeDownloadRequest{
		VolumeName: "models",
		Path:       "packed.bin",
	}, nil)

	require.NoError(t, err)
	require.Equal(t, body, got)
	require.Equal(t, "bytes=7-17", gotRange)
}

func TestVolumeSignedPutRefreshesURLAndRewindsBody(t *testing.T) {
	var bodies [][]byte
	var mu sync.Mutex
	objectServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		mu.Lock()
		bodies = append(bodies, data)
		attempt := len(bodies)
		mu.Unlock()
		if attempt == 1 {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer objectServer.Close()

	refreshes := 0
	client := &volumeClientImpl{httpClient: objectServer.Client()}
	err := client.signedRequestWithRetry(
		context.Background(),
		http.MethodPut,
		objectServer.URL,
		octetStreamHeaders(),
		bytes.NewReader([]byte("payload")),
		2,
		0,
		func(context.Context) (string, error) {
			refreshes++
			return objectServer.URL, nil
		},
	)

	require.NoError(t, err)
	require.Equal(t, 1, refreshes)
	require.Equal(t, [][]byte{[]byte("payload"), []byte("payload")}, bodies)
}

func TestVolumeSafeRelativePath(t *testing.T) {
	path, err := safeVolumeRelativePath("models/v1/model.bin")
	require.NoError(t, err)
	require.Equal(t, "models/v1/model.bin", path)
	_, err = safeVolumeRelativePath("../model.bin")
	require.Error(t, err)
	_, err = safeVolumeRelativePath("/tmp/model.bin")
	require.Error(t, err)
}

func TestVolumeDataPackSealParsesAndAddressesMembers(t *testing.T) {
	files := [][]byte{
		[]byte("alpha payload"),
		[]byte("bravo payload longer"),
		[]byte("charlie"),
	}
	builder := newDataPackBuilder()
	hashes := make([][32]byte, 0, len(files))
	for _, file := range files {
		hash := blake3Sum(file)
		hashes = append(hashes, hash)
		builder.append(hash, file)
	}

	sealed, err := builder.seal()
	require.NoError(t, err)
	require.Equal(t, []byte("CDP1"), sealed.bytes[:4])
	require.Equal(t, blake3Sum(sealed.bytes), sealed.chunkID)

	entries, bodyStart, err := parseDataPackHeader(sealed.bytes)
	require.NoError(t, err)
	require.Equal(t, sealed.entries, entries)

	for i, hash := range hashes {
		entry, ok := sealed.entryFor(hash)
		require.True(t, ok)
		require.GreaterOrEqual(t, int(entry.offset), bodyStart)
		got := sealed.bytes[entry.offset : entry.offset+uint64(entry.length)]
		require.Equal(t, files[i], got)
	}
}

func TestVolumeDataPackDedupesIdenticalContent(t *testing.T) {
	data := []byte("same bytes")
	hash := blake3Sum(data)
	builder := newDataPackBuilder()
	builder.append(hash, data)
	builder.append(hash, data)

	sealed, err := builder.seal()
	require.NoError(t, err)
	require.Len(t, sealed.entries, 1)
}

func TestVolumeDataPackSealIsDeterministic(t *testing.T) {
	build := func() []byte {
		builder := newDataPackBuilder()
		for _, data := range [][]byte{[]byte("alpha"), []byte("charlie")} {
			builder.append(blake3Sum(data), data)
		}
		sealed, err := builder.seal()
		require.NoError(t, err)
		return sealed.bytes
	}

	require.True(t, bytes.Equal(build(), build()))
}

func TestVolumeDataPackFitsMatchesSealedLength(t *testing.T) {
	builder := newDataPackBuilder()
	for i := 0; i < 30; i++ {
		data := bytes.Repeat([]byte{byte('a' + i%26)}, 24+i)
		require.True(t, builder.fits(uint64(len(data)), 10000))
		builder.append(blake3Sum(data), data)
	}

	sealed, err := builder.seal()
	require.NoError(t, err)
	require.Equal(t, uint64(len(sealed.bytes)), builder.objectLen())
}

type fakeVolumeRPC struct {
	createVolume       func(context.Context, *connect.Request[volumev2.CreateVolumeRequest]) (*connect.Response[volumev2.CreateVolumeResponse], error)
	getVolume          func(context.Context, *connect.Request[volumev2.GetVolumeRequest]) (*connect.Response[volumev2.GetVolumeResponse], error)
	listVolumes        func(context.Context, *connect.Request[volumev2.ListVolumesRequest]) (*connect.Response[volumev2.ListVolumesResponse], error)
	deleteVolume       func(context.Context, *connect.Request[volumev2.DeleteVolumeRequest]) (*connect.Response[volumev2.DeleteVolumeResponse], error)
	listVolumeVersions func(context.Context, *connect.Request[volumev2.ListVolumeVersionsRequest]) (*connect.Response[volumev2.ListVolumeVersionsResponse], error)
	commitVersion      func(context.Context, *connect.Request[volumev2.CommitVersionRequest]) (*connect.Response[volumev2.CommitVersionResponse], error)
	getCommitStatus    func(context.Context, *connect.Request[volumev2.GetCommitStatusRequest]) (*connect.Response[volumev2.GetCommitStatusResponse], error)
	allocateInodeRange func(context.Context, *connect.Request[volumev2.AllocateInodeRangeRequest]) (*connect.Response[volumev2.AllocateInodeRangeResponse], error)
	requestUploadURLs  func(context.Context, *connect.Request[volumev2.RequestUploadURLsRequest]) (*connect.Response[volumev2.RequestUploadURLsResponse], error)
	listFiles          func(context.Context, *connect.Request[volumev2.ListFilesRequest]) (*connect.Response[volumev2.ListFilesResponse], error)
	getFile            func(context.Context, *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error)
	commits            []*volumev2.CommitIntent
}

func (f *fakeVolumeRPC) CreateVolume(ctx context.Context, req *connect.Request[volumev2.CreateVolumeRequest]) (*connect.Response[volumev2.CreateVolumeResponse], error) {
	if f.createVolume != nil {
		return f.createVolume(ctx, req)
	}
	return nil, fakeVolumeUnimplemented()
}

func (f *fakeVolumeRPC) GetVolume(ctx context.Context, req *connect.Request[volumev2.GetVolumeRequest]) (*connect.Response[volumev2.GetVolumeResponse], error) {
	if f.getVolume != nil {
		return f.getVolume(ctx, req)
	}
	return nil, fakeVolumeUnimplemented()
}

func (f *fakeVolumeRPC) ListVolumes(ctx context.Context, req *connect.Request[volumev2.ListVolumesRequest]) (*connect.Response[volumev2.ListVolumesResponse], error) {
	if f.listVolumes != nil {
		return f.listVolumes(ctx, req)
	}
	return nil, fakeVolumeUnimplemented()
}

func (f *fakeVolumeRPC) DeleteVolume(ctx context.Context, req *connect.Request[volumev2.DeleteVolumeRequest]) (*connect.Response[volumev2.DeleteVolumeResponse], error) {
	if f.deleteVolume != nil {
		return f.deleteVolume(ctx, req)
	}
	return nil, fakeVolumeUnimplemented()
}

func (f *fakeVolumeRPC) ListVolumeVersions(ctx context.Context, req *connect.Request[volumev2.ListVolumeVersionsRequest]) (*connect.Response[volumev2.ListVolumeVersionsResponse], error) {
	if f.listVolumeVersions != nil {
		return f.listVolumeVersions(ctx, req)
	}
	return nil, fakeVolumeUnimplemented()
}

func (f *fakeVolumeRPC) CommitVersion(ctx context.Context, req *connect.Request[volumev2.CommitVersionRequest]) (*connect.Response[volumev2.CommitVersionResponse], error) {
	if f.commitVersion != nil {
		return f.commitVersion(ctx, req)
	}
	return nil, fakeVolumeUnimplemented()
}

func (f *fakeVolumeRPC) GetCommitStatus(ctx context.Context, req *connect.Request[volumev2.GetCommitStatusRequest]) (*connect.Response[volumev2.GetCommitStatusResponse], error) {
	if f.getCommitStatus != nil {
		return f.getCommitStatus(ctx, req)
	}
	return nil, fakeVolumeUnimplemented()
}

func (f *fakeVolumeRPC) AllocateInodeRange(ctx context.Context, req *connect.Request[volumev2.AllocateInodeRangeRequest]) (*connect.Response[volumev2.AllocateInodeRangeResponse], error) {
	if f.allocateInodeRange != nil {
		return f.allocateInodeRange(ctx, req)
	}
	return nil, fakeVolumeUnimplemented()
}

func (f *fakeVolumeRPC) RequestUploadURLs(ctx context.Context, req *connect.Request[volumev2.RequestUploadURLsRequest]) (*connect.Response[volumev2.RequestUploadURLsResponse], error) {
	if f.requestUploadURLs != nil {
		return f.requestUploadURLs(ctx, req)
	}
	return nil, fakeVolumeUnimplemented()
}

func (f *fakeVolumeRPC) ListFiles(ctx context.Context, req *connect.Request[volumev2.ListFilesRequest]) (*connect.Response[volumev2.ListFilesResponse], error) {
	if f.listFiles != nil {
		return f.listFiles(ctx, req)
	}
	return nil, fakeVolumeUnimplemented()
}

func (f *fakeVolumeRPC) GetFile(ctx context.Context, req *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
	if f.getFile != nil {
		return f.getFile(ctx, req)
	}
	return nil, fakeVolumeUnimplemented()
}

func fakeVolumeUnimplemented() error {
	return connect.NewError(connect.CodeUnimplemented, io.ErrClosedPipe)
}
