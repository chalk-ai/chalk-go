package chalk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{
		getFile: func(context.Context, *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			return connect.NewResponse(&volumev2.GetFileResponse{
				File:    &volumev2.FileInfo{Path: "model.bin", Size: 5, Hash: blake3Hex([]byte("hello"))},
				Content: &volumev2.GetFileResponse_Data{Data: []byte("hello")},
			}), nil
		},
	})
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
	path, err := safeVolumeRelativePath("models/v1/model.bin")
	require.NoError(t, err)
	require.Equal(t, "models/v1/model.bin", path)
	_, err = safeVolumeRelativePath("../model.bin")
	require.Error(t, err)
	_, err = safeVolumeRelativePath("/tmp/model.bin")
	require.Error(t, err)
}

func TestVolumeDataPackSealParsesAndAddressesMembers(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
	builder := newDataPackBuilder()
	for i := range 30 {
		data := bytes.Repeat([]byte{byte('a' + i%26)}, 24+i)
		require.True(t, builder.fits(uint64(len(data)), 10000))
		builder.append(blake3Sum(data), data)
	}

	sealed, err := builder.seal()
	require.NoError(t, err)
	require.Equal(t, uint64(len(sealed.bytes)), builder.objectLen())
}

func TestVolumeDefaultUploadConfig(t *testing.T) {
	t.Parallel()
	cfg := DefaultVolumeUploadConfig()
	require.Greater(t, cfg.ChunkSize, uint64(0))
	require.Greater(t, cfg.MaxPackBytes, uint64(0))
	require.Greater(t, cfg.BatchSize, 0)
	require.Greater(t, cfg.FileConcurrency, 0)
}

func TestVolumeRefSelector(t *testing.T) {
	t.Parallel()
	sel := VolumeRefSelector("my-ref")
	require.Equal(t, "my-ref", sel.GetRef())
}

func TestVolumeLocalPathContentSizeAndRead(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	path := filepath.Join(dir, "data.bin")
	require.NoError(t, os.WriteFile(path, []byte("hello world"), 0o644))

	content := VolumeUploadLocalPath(path)
	size, err := content.size()
	require.NoError(t, err)
	require.Equal(t, uint64(11), size)

	chunk, err := content.readChunk(6, 5)
	require.NoError(t, err)
	require.Equal(t, []byte("world"), chunk)

	_, err = VolumeUploadLocalPath("/nonexistent/path.bin").size()
	require.Error(t, err)
}

func TestVolumeCollectLocalFiles(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	subDir := filepath.Join(dir, "sub")
	require.NoError(t, os.MkdirAll(subDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "a.txt"), []byte("a"), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(subDir, "b.txt"), []byte("b"), 0o644))

	files, err := collectVolumeLocalFiles(dir)
	require.NoError(t, err)
	require.Len(t, files, 2)
	var paths []string
	for _, f := range files {
		paths = append(paths, f.Path)
	}
	require.ElementsMatch(t, []string{"a.txt", "sub/b.txt"}, paths)
}

func TestVolumeRemoveFiles(t *testing.T) {
	t.Parallel()
	var capturedRemoves []*volumev2.PathRemoveDelta
	rpc := &fakeVolumeRPC{
		getVolume: func(_ context.Context, _ *connect.Request[volumev2.GetVolumeRequest]) (*connect.Response[volumev2.GetVolumeResponse], error) {
			return connect.NewResponse(&volumev2.GetVolumeResponse{
				Version: &volumev2.VersionInfo{VersionId: 1, SequenceNumber: 0},
			}), nil
		},
		commitVersion: func(_ context.Context, req *connect.Request[volumev2.CommitVersionRequest]) (*connect.Response[volumev2.CommitVersionResponse], error) {
			capturedRemoves = req.Msg.GetIntent().GetPathDeltas().GetRemoves()
			return connect.NewResponse(&volumev2.CommitVersionResponse{
				Status: &volumev2.CommitStatus{Result: volumev2.CommitResult_COMMIT_RESULT_COMMITTED},
			}), nil
		},
	}
	client := &volumeClientImpl{rpc: rpc, author: "test"}
	status, err := client.RemoveFiles(context.Background(), "my-vol", []VolumeRemovePath{
		{Path: "old.bin"},
		{Path: "dir/", Recursive: true},
	}, 1)
	require.NoError(t, err)
	require.Equal(t, volumev2.CommitResult_COMMIT_RESULT_COMMITTED, status.Result)
	require.Len(t, capturedRemoves, 2)
	require.Equal(t, "old.bin", capturedRemoves[0].Path)
	require.False(t, capturedRemoves[0].Recursive)
	require.Equal(t, "dir/", capturedRemoves[1].Path)
	require.True(t, capturedRemoves[1].Recursive)
}

func TestVolumeUploadFilesEmptyBatch(t *testing.T) {
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{})
	statuses, err := client.UploadFiles(context.Background(), VolumeUploadRequest{
		VolumeName: "test-vol",
		Files:      nil,
	}, nil)
	require.NoError(t, err)
	require.Empty(t, statuses)
}

func TestVolumeUploadFilesZeroSizeFile(t *testing.T) {
	t.Parallel()
	var commits []*volumev2.CommitIntent
	rpc := &fakeVolumeRPC{
		getVolume: func(_ context.Context, _ *connect.Request[volumev2.GetVolumeRequest]) (*connect.Response[volumev2.GetVolumeResponse], error) {
			return connect.NewResponse(&volumev2.GetVolumeResponse{
				Version: &volumev2.VersionInfo{VersionId: 1, SequenceNumber: 0},
			}), nil
		},
		commitVersion: func(_ context.Context, req *connect.Request[volumev2.CommitVersionRequest]) (*connect.Response[volumev2.CommitVersionResponse], error) {
			commits = append(commits, req.Msg.GetIntent())
			return connect.NewResponse(&volumev2.CommitVersionResponse{
				Status: &volumev2.CommitStatus{Result: volumev2.CommitResult_COMMIT_RESULT_COMMITTED},
			}), nil
		},
	}
	client := &volumeClientImpl{rpc: rpc, author: "test"}
	statuses, err := client.UploadFiles(context.Background(), VolumeUploadRequest{
		VolumeName: "test-vol",
		Files:      []VolumeUploadFile{{Path: "empty.bin", Content: VolumeUploadBytes(nil)}},
	}, nil)
	require.NoError(t, err)
	require.Len(t, statuses, 1)
	require.Len(t, commits, 1)
	require.Len(t, commits[0].GetPathDeltas().GetUpserts(), 1)
}

func TestVolumeUploadDirectoryPackPath(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, "a.bin"), []byte("content-a"), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "b.bin"), []byte("content-b"), 0o644))

	rpc := &fakeVolumeRPC{
		requestUploadURLs: func(_ context.Context, req *connect.Request[volumev2.RequestUploadURLsRequest]) (*connect.Response[volumev2.RequestUploadURLsResponse], error) {
			urls := make([]*volumev2.UploadURLItem, len(req.Msg.GetObjects()))
			for i, obj := range req.Msg.GetObjects() {
				urls[i] = &volumev2.UploadURLItem{ObjectKey: obj.GetObjectKey(), AlreadyExists: true}
			}
			return connect.NewResponse(&volumev2.RequestUploadURLsResponse{Urls: urls}), nil
		},
		getVolume: func(_ context.Context, _ *connect.Request[volumev2.GetVolumeRequest]) (*connect.Response[volumev2.GetVolumeResponse], error) {
			return connect.NewResponse(&volumev2.GetVolumeResponse{
				Version: &volumev2.VersionInfo{VersionId: 1, SequenceNumber: 0},
			}), nil
		},
		commitVersion: func(_ context.Context, req *connect.Request[volumev2.CommitVersionRequest]) (*connect.Response[volumev2.CommitVersionResponse], error) {
			return connect.NewResponse(&volumev2.CommitVersionResponse{
				Status: &volumev2.CommitStatus{Result: volumev2.CommitResult_COMMIT_RESULT_COMMITTED},
			}), nil
		},
	}
	client := &volumeClientImpl{rpc: rpc, author: "test"}
	statuses, err := client.UploadDirectory(context.Background(), "test-vol", dir, VolumeUploadConfig{})
	require.NoError(t, err)
	require.Len(t, statuses, 1)
	require.Equal(t, volumev2.CommitResult_COMMIT_RESULT_COMMITTED, statuses[0].Result)
}

// TestVolumeUploadFilesChunkPath forces the per-file (non-pack) upload path by setting
// MaxPackBytes=1 so no file fits in a pack, then exercises uploadOneFile with a real PUT.
func TestVolumeUploadFilesChunkPath(t *testing.T) {
	t.Parallel()
	fileData := []byte("chunk content data")
	var putBodies [][]byte
	var mu sync.Mutex
	putSrv, putURL := testHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			data, _ := io.ReadAll(r.Body)
			mu.Lock()
			putBodies = append(putBodies, data)
			mu.Unlock()
		}
		w.WriteHeader(http.StatusOK)
	})

	var commits []*volumev2.CommitIntent
	rpc := &fakeVolumeRPC{
		requestUploadURLs: func(_ context.Context, req *connect.Request[volumev2.RequestUploadURLsRequest]) (*connect.Response[volumev2.RequestUploadURLsResponse], error) {
			urls := make([]*volumev2.UploadURLItem, len(req.Msg.GetObjects()))
			for i, obj := range req.Msg.GetObjects() {
				urls[i] = &volumev2.UploadURLItem{
					ObjectKey:       obj.GetObjectKey(),
					SignedUploadUri: putURL,
					AlreadyExists:   false,
				}
			}
			return connect.NewResponse(&volumev2.RequestUploadURLsResponse{Urls: urls}), nil
		},
		getVolume: func(_ context.Context, _ *connect.Request[volumev2.GetVolumeRequest]) (*connect.Response[volumev2.GetVolumeResponse], error) {
			return connect.NewResponse(&volumev2.GetVolumeResponse{
				Version: &volumev2.VersionInfo{VersionId: 1, SequenceNumber: 0},
			}), nil
		},
		commitVersion: func(_ context.Context, req *connect.Request[volumev2.CommitVersionRequest]) (*connect.Response[volumev2.CommitVersionResponse], error) {
			commits = append(commits, req.Msg.GetIntent())
			return connect.NewResponse(&volumev2.CommitVersionResponse{
				Status: &volumev2.CommitStatus{Result: volumev2.CommitResult_COMMIT_RESULT_COMMITTED},
			}), nil
		},
	}
	client := &volumeClientImpl{httpClient: putSrv.Client(), rpc: rpc, author: "test"}
	statuses, err := client.UploadFiles(context.Background(), VolumeUploadRequest{
		VolumeName: "test-vol",
		Files:      []VolumeUploadFile{{Path: "file.bin", Content: VolumeUploadBytes(fileData)}},
		Config:     VolumeUploadConfig{MaxPackBytes: 1}, // force per-file path
	}, func(uint64) {})
	require.NoError(t, err)
	require.Len(t, statuses, 1)
	require.Equal(t, volumev2.CommitResult_COMMIT_RESULT_COMMITTED, statuses[0].Result)
	require.NotEmpty(t, putBodies)
	require.Equal(t, fileData, putBodies[0])
	require.Len(t, commits, 1)
	require.Len(t, commits[0].GetPathDeltas().GetUpserts(), 1)
}

func TestVolumeDownloadToFile(t *testing.T) {
	t.Parallel()
	content := []byte("file content data")
	client := newClientWithRPC(&fakeVolumeRPC{
		getFile: func(_ context.Context, _ *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			return connect.NewResponse(&volumev2.GetFileResponse{
				File:    &volumev2.FileInfo{Path: "model.bin", Size: uint64(len(content))},
				Content: &volumev2.GetFileResponse_Data{Data: content},
			}), nil
		},
	})
	localPath := filepath.Join(t.TempDir(), "model.bin")
	info, err := client.DownloadToFile(context.Background(), VolumeDownloadRequest{
		VolumeName: "models",
		Path:       "model.bin",
	}, localPath, nil)
	require.NoError(t, err)
	require.Equal(t, "model.bin", info.Path)
	got, err := os.ReadFile(localPath)
	require.NoError(t, err)
	require.Equal(t, content, got)
}

func TestVolumeDownloadToDirectory(t *testing.T) {
	t.Parallel()
	content := []byte("dir file content")
	var getFileCalls int
	client := newClientWithRPC(&fakeVolumeRPC{
		listFiles: func(_ context.Context, req *connect.Request[volumev2.ListFilesRequest]) (*connect.Response[volumev2.ListFilesResponse], error) {
			return connect.NewResponse(&volumev2.ListFilesResponse{
				Files: []*volumev2.FileInfo{
					{Path: "sub/model.bin", Size: uint64(len(content)), Kind: volumev2.FileKind_FILE_KIND_FILE},
				},
			}), nil
		},
		getFile: func(_ context.Context, _ *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			getFileCalls++
			return connect.NewResponse(&volumev2.GetFileResponse{
				File:    &volumev2.FileInfo{Path: "sub/model.bin", Size: uint64(len(content))},
				Content: &volumev2.GetFileResponse_Data{Data: content},
			}), nil
		},
	})
	localDir := t.TempDir()
	err := client.DownloadToDirectory(context.Background(), "models", localDir, nil, VolumeDownloadConfig{})
	require.NoError(t, err)
	require.Equal(t, 1, getFileCalls)
	got, err := os.ReadFile(filepath.Join(localDir, "sub", "model.bin"))
	require.NoError(t, err)
	require.Equal(t, content, got)
}

func TestVolumeDownloadToDirectoryWithSubdirs(t *testing.T) {
	t.Parallel()
	content := []byte("nested file")
	client := newClientWithRPC(&fakeVolumeRPC{
		listFiles: func(_ context.Context, req *connect.Request[volumev2.ListFilesRequest]) (*connect.Response[volumev2.ListFilesResponse], error) {
			if req.Msg.GetPath() == "" {
				return connect.NewResponse(&volumev2.ListFilesResponse{
					Files: []*volumev2.FileInfo{
						{Path: "data/", Kind: volumev2.FileKind_FILE_KIND_DIRECTORY},
					},
				}), nil
			}
			return connect.NewResponse(&volumev2.ListFilesResponse{
				Files: []*volumev2.FileInfo{
					{Path: "data/file.bin", Size: uint64(len(content)), Kind: volumev2.FileKind_FILE_KIND_FILE},
				},
			}), nil
		},
		getFile: func(_ context.Context, _ *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			return connect.NewResponse(&volumev2.GetFileResponse{
				File:    &volumev2.FileInfo{Path: "data/file.bin"},
				Content: &volumev2.GetFileResponse_Data{Data: content},
			}), nil
		},
	})
	localDir := t.TempDir()
	err := client.DownloadToDirectory(context.Background(), "models", localDir, nil, VolumeDownloadConfig{})
	require.NoError(t, err)
	got, err := os.ReadFile(filepath.Join(localDir, "data", "file.bin"))
	require.NoError(t, err)
	require.Equal(t, content, got)
}

func TestVolumeListFilesRecursivePagination(t *testing.T) {
	t.Parallel()
	var calls int
	client := newClientWithRPC(&fakeVolumeRPC{
		listFiles: func(_ context.Context, req *connect.Request[volumev2.ListFilesRequest]) (*connect.Response[volumev2.ListFilesResponse], error) {
			calls++
			if req.Msg.GetCursor() == "" {
				return connect.NewResponse(&volumev2.ListFilesResponse{
					Files:      []*volumev2.FileInfo{{Path: "a.bin", Kind: volumev2.FileKind_FILE_KIND_FILE}},
					NextCursor: "page2",
				}), nil
			}
			return connect.NewResponse(&volumev2.ListFilesResponse{
				Files: []*volumev2.FileInfo{{Path: "b.bin", Kind: volumev2.FileKind_FILE_KIND_FILE}},
			}), nil
		},
	})
	files, err := client.listFilesRecursive(context.Background(), "models", nil)
	require.NoError(t, err)
	require.Len(t, files, 2)
	require.Equal(t, 2, calls)
}

func TestVolumeParseDataPackHeaderErrors(t *testing.T) {
	t.Parallel()
	t.Run("too short", func(t *testing.T) {
		t.Parallel()
		_, _, err := parseDataPackHeader([]byte{0, 1, 2})
		require.Error(t, err)
	})
	t.Run("bad magic", func(t *testing.T) {
		t.Parallel()
		data := make([]byte, 16)
		copy(data[:4], []byte("BAD1"))
		_, _, err := parseDataPackHeader(data)
		require.Error(t, err)
	})
	t.Run("header length exceeds data", func(t *testing.T) {
		t.Parallel()
		// 10-byte buffer; write headerLen=1000 (0x000003E8 LE) so bodyStart > len(data).
		data := make([]byte, 10)
		copy(data[:4], cdp1Magic)
		data[4] = 0xE8
		data[5] = 0x03 // 1000 little-endian
		_, _, err := parseDataPackHeader(data)
		require.Error(t, err)
	})
}

func TestVolumeFitsOversizedFile(t *testing.T) {
	t.Parallel()
	b := newDataPackBuilder()
	require.False(t, b.fits(uint64(^uint32(0))+1, 1<<40))
}

func TestVolumeSignedRequestDefaultsMaxAttempts(t *testing.T) {
	t.Parallel()
	srv, url := testHTTPServer(t, func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK) })
	client := &volumeClientImpl{httpClient: srv.Client()}
	err := client.signedRequestWithRetry(context.Background(), http.MethodPut, url,
		nil, nil, 0, 0,
		func(context.Context) (string, error) { return url, nil })
	require.NoError(t, err)
}

func TestVolumeSignedRequestRefreshError(t *testing.T) {
	t.Parallel()
	srv, url := testHTTPServer(t, func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusForbidden) })
	client := &volumeClientImpl{httpClient: srv.Client()}
	err := client.signedRequestWithRetry(context.Background(), http.MethodPut, url,
		nil, bytes.NewReader(nil), 1, 0,
		func(context.Context) (string, error) { return "", io.ErrClosedPipe })
	require.Error(t, err)
}

func TestVolumeSignedRequestNonSuccessRetry(t *testing.T) {
	t.Parallel()
	var count int
	srv, url := testHTTPServer(t, func(w http.ResponseWriter, _ *http.Request) {
		count++
		w.WriteHeader(http.StatusInternalServerError)
	})
	client := &volumeClientImpl{httpClient: srv.Client()}
	err := client.signedRequestWithRetry(context.Background(), http.MethodPut, url,
		nil, bytes.NewReader(nil), 2, 0,
		func(context.Context) (string, error) { return url, nil })
	require.Error(t, err)
	require.Equal(t, 2, count)
}

func TestVolumeSignedRequestRateLimitExceeded(t *testing.T) {
	t.Parallel()
	srv, url := testHTTPServer(t, func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusTooManyRequests) })
	client := &volumeClientImpl{httpClient: srv.Client()}
	err := client.signedRequestWithRetry(context.Background(), http.MethodPut, url,
		nil, bytes.NewReader(nil), 1, 0,
		func(context.Context) (string, error) { return url, nil })
	require.Error(t, err)
}

func TestVolumeSignedGetDefaultsMaxAttempts(t *testing.T) {
	t.Parallel()
	srv, url := testHTTPServer(t, func(w http.ResponseWriter, _ *http.Request) { _, _ = w.Write([]byte("ok")) })
	client := &volumeClientImpl{httpClient: srv.Client()}
	data, err := client.signedGetWithRetry(context.Background(), url, nil, 0,
		func(context.Context) (string, error) { return url, nil })
	require.NoError(t, err)
	require.Equal(t, []byte("ok"), data)
}

func TestVolumeSignedGetRefreshError(t *testing.T) {
	t.Parallel()
	srv, url := testHTTPServer(t, func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusUnauthorized) })
	client := &volumeClientImpl{httpClient: srv.Client()}
	_, err := client.signedGetWithRetry(context.Background(), url, nil, 1,
		func(context.Context) (string, error) { return "", io.ErrClosedPipe })
	require.Error(t, err)
}

func TestVolumeSignedGetNonSuccessRetry(t *testing.T) {
	t.Parallel()
	var count int
	srv, url := testHTTPServer(t, func(w http.ResponseWriter, _ *http.Request) {
		count++
		w.WriteHeader(http.StatusInternalServerError)
	})
	client := &volumeClientImpl{httpClient: srv.Client()}
	_, err := client.signedGetWithRetry(context.Background(), url, nil, 2,
		func(context.Context) (string, error) { return url, nil })
	require.Error(t, err)
	require.Equal(t, 2, count)
}

func TestVolumeResolveCommitAuthorHTTPError(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))
	url := server.URL
	client := newClientWithServer(t, url)
	server.Close() // close after client creation so httpClient.Do fails

	impl := client.(*volumeClientImpl)
	require.Empty(t, impl.resolveCommitAuthor(context.Background()))
}

func TestVolumeResolveCommitAuthorNonSuccessStatus(t *testing.T) {
	t.Parallel()
	_, url := testHTTPServer(t, func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusInternalServerError) })
	client := newClientWithServer(t, url)
	require.Empty(t, client.(*volumeClientImpl).resolveCommitAuthor(context.Background()))
}

func TestVolumeResolveCommitAuthorInvalidJSON(t *testing.T) {
	t.Parallel()
	_, url := testHTTPServer(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{invalid`))
	})
	client := newClientWithServer(t, url)
	require.Empty(t, client.(*volumeClientImpl).resolveCommitAuthor(context.Background()))
}

func TestVolumeResolveCommitAuthorUsesBodyEnvID(t *testing.T) {
	t.Parallel()
	_, url := testHTTPServer(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"user":"bob","environment_id":"env-from-body"}`))
	})
	// Create a client normally, then clear envID so the body's environment_id is used.
	client := newClientWithServer(t, url)
	impl := client.(*volumeClientImpl)
	impl.envID = ""
	author := impl.resolveCommitAuthor(context.Background())
	require.Equal(t, "chalk:env-from-body:agent:bob", author)
}

func TestVolumeDownloadGetFileError(t *testing.T) {
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{
		getFile: func(_ context.Context, _ *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			return nil, fakeVolumeUnimplemented()
		},
	})
	_, _, err := client.DownloadBytes(context.Background(), VolumeDownloadRequest{
		VolumeName: "models", Path: "f.bin",
	}, nil)
	require.Error(t, err)
}

func TestVolumeDownloadPackedRangeOverflow(t *testing.T) {
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{
		getFile: func(_ context.Context, _ *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			return connect.NewResponse(&volumev2.GetFileResponse{
				File:    &volumev2.FileInfo{Path: "f.bin"},
				Version: &volumev2.VersionInfo{VersionId: 1},
				Content: &volumev2.GetFileResponse_Packed{Packed: &volumev2.PackedFileContent{
					// offset + size - 1 overflows uint64
					Pack: &volumev2.SignedPackEntryRef{Offset: ^uint64(0) - 1, Size: 3},
				}},
			}), nil
		},
	})
	_, _, err := client.DownloadBytes(context.Background(), VolumeDownloadRequest{
		VolumeName: "models", Path: "f.bin",
	}, nil)
	require.Error(t, err)
}

func TestVolumeDownloadPackedGetFailure(t *testing.T) {
	t.Parallel()
	client := newHTTPFakeClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}, func(url string) *fakeVolumeRPC {
		return &fakeVolumeRPC{getFile: func(_ context.Context, _ *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			return packedGetFileResponse(url, 0, 5, ""), nil
		}}
	})
	_, _, err := client.DownloadBytes(context.Background(), VolumeDownloadRequest{
		VolumeName: "models", Path: "f.bin",
		Config: VolumeDownloadConfig{MaxChunkRetries: 1},
	}, nil)
	require.Error(t, err)
}

func TestVolumeDownloadPackedSizeMismatch(t *testing.T) {
	t.Parallel()
	body := []byte("ab") // 2 bytes; pack declares Size: 100
	client := newHTTPFakeClient(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(body)
	}, func(url string) *fakeVolumeRPC {
		return &fakeVolumeRPC{getFile: func(_ context.Context, _ *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			return packedGetFileResponse(url, 0, 100, ""), nil
		}}
	})
	_, _, err := client.DownloadBytes(context.Background(), VolumeDownloadRequest{
		VolumeName: "models", Path: "f.bin",
	}, nil)
	require.Error(t, err)
}

func TestVolumeDownloadPackedHashMismatch(t *testing.T) {
	t.Parallel()
	body := []byte("data bytes")
	client := newHTTPFakeClient(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(body)
	}, func(url string) *fakeVolumeRPC {
		return &fakeVolumeRPC{getFile: func(_ context.Context, _ *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			return packedGetFileResponse(url, 0, uint64(len(body)), "deadbeef"), nil
		}}
	})
	_, _, err := client.DownloadBytes(context.Background(), VolumeDownloadRequest{
		VolumeName: "models", Path: "f.bin",
	}, nil)
	require.Error(t, err)
}

func TestVolumeDownloadChunkedHashMismatch(t *testing.T) {
	t.Parallel()
	body := []byte("chunk data")
	client := newHTTPFakeClient(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(body)
	}, func(url string) *fakeVolumeRPC {
		return &fakeVolumeRPC{getFile: func(_ context.Context, _ *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			return connect.NewResponse(&volumev2.GetFileResponse{
				File:    &volumev2.FileInfo{Path: "f.bin"},
				Version: &volumev2.VersionInfo{VersionId: 1},
				Content: &volumev2.GetFileResponse_Chunked{Chunked: &volumev2.ChunkedFileContent{
					Chunks: []*volumev2.SignedChunkRef{{
						SignedDownloadUri: url,
						Offset:            0,
						Size:              uint64(len(body)),
						Hash:              "deadbeef", // wrong hash
					}},
				}},
			}), nil
		}}
	})
	_, _, err := client.DownloadBytes(context.Background(), VolumeDownloadRequest{
		VolumeName: "models", Path: "f.bin",
		Config: VolumeDownloadConfig{ChunkConcurrency: 1, MaxChunkRetries: 1},
	}, nil)
	require.Error(t, err)
}

func TestVolumeRateLimitBackoff(t *testing.T) {
	t.Parallel()
	tests := []struct {
		attempt int
		want    uint64
	}{
		{0, 1000},
		{1, 2000},
		{2, 4000},
		{3, 8000},
		{4, 16000},
		{5, 30000}, // capped at 30s
		{9, 30000}, // still capped
	}
	for _, tc := range tests {
		require.Equal(t, tc.want, rateLimitBackoffMS(tc.attempt), "attempt=%d", tc.attempt)
	}
}

func TestVolumeChunkRelativeObjectKey(t *testing.T) {
	t.Parallel()
	tests := []struct{ hash, want string }{
		{"abcdef1234", "ab/abcdef1234"},
		{"x", "x"}, // len < 2: returned as-is
		{"", ""},   // empty: returned as-is
	}
	for _, tc := range tests {
		require.Equal(t, tc.want, chunkRelativeObjectKey(tc.hash), "hash=%q", tc.hash)
	}
}

func TestVolumePinnedFileRequestNilVersion(t *testing.T) {
	t.Parallel()
	_, err := pinnedVolumeFileRequest(&volumev2.GetFileRequest{}, nil)
	require.Error(t, err)
}

func TestVolumeApplyFileInfoMetadata(t *testing.T) {
	t.Parallel()
	// nil info must be a no-op (no panic)
	applyVolumeFileInfoMetadata("/nonexistent/path", nil)

	dir := t.TempDir()
	path := filepath.Join(dir, "f.bin")
	require.NoError(t, os.WriteFile(path, []byte("x"), 0o644))

	mode := uint32(0o755)
	modTime := time.Now().Add(-time.Hour).Truncate(time.Second)
	applyVolumeFileInfoMetadata(path, &volumev2.FileInfo{
		Mode:      &mode,
		UpdatedAt: timestamppb.New(modTime),
	})
	// Chmod and Chtimes results are platform-dependent; just verify the file still exists.
	_, err := os.Stat(path)
	require.NoError(t, err)
}

func TestVolumeDownloadPackedZeroSize(t *testing.T) {
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{
		getFile: func(_ context.Context, _ *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			return connect.NewResponse(&volumev2.GetFileResponse{
				File:    &volumev2.FileInfo{Path: "empty.bin"},
				Version: &volumev2.VersionInfo{VersionId: 1},
				Content: &volumev2.GetFileResponse_Packed{Packed: &volumev2.PackedFileContent{
					Pack: &volumev2.SignedPackEntryRef{Size: 0},
				}},
			}), nil
		},
	})
	got, info, err := client.DownloadBytes(context.Background(), VolumeDownloadRequest{
		VolumeName: "models",
		Path:       "empty.bin",
	}, nil)
	require.NoError(t, err)
	require.Empty(t, got)
	require.Equal(t, "empty.bin", info.Path)
}

func TestVolumeUploadPackWithPUT(t *testing.T) {
	t.Parallel()
	fileData := []byte("pack payload data")
	var putCount int
	putSrv, putURL := testHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			putCount++
		}
		w.WriteHeader(http.StatusOK)
	})

	rpc := &fakeVolumeRPC{
		requestUploadURLs: func(_ context.Context, req *connect.Request[volumev2.RequestUploadURLsRequest]) (*connect.Response[volumev2.RequestUploadURLsResponse], error) {
			urls := make([]*volumev2.UploadURLItem, len(req.Msg.GetObjects()))
			for i, obj := range req.Msg.GetObjects() {
				urls[i] = &volumev2.UploadURLItem{
					ObjectKey:       obj.GetObjectKey(),
					SignedUploadUri: putURL,
					AlreadyExists:   false,
				}
			}
			return connect.NewResponse(&volumev2.RequestUploadURLsResponse{Urls: urls}), nil
		},
		getVolume: func(_ context.Context, _ *connect.Request[volumev2.GetVolumeRequest]) (*connect.Response[volumev2.GetVolumeResponse], error) {
			return connect.NewResponse(&volumev2.GetVolumeResponse{
				Version: &volumev2.VersionInfo{VersionId: 1, SequenceNumber: 0},
			}), nil
		},
		commitVersion: func(_ context.Context, _ *connect.Request[volumev2.CommitVersionRequest]) (*connect.Response[volumev2.CommitVersionResponse], error) {
			return connect.NewResponse(&volumev2.CommitVersionResponse{
				Status: &volumev2.CommitStatus{Result: volumev2.CommitResult_COMMIT_RESULT_COMMITTED},
			}), nil
		},
	}
	client := &volumeClientImpl{httpClient: putSrv.Client(), rpc: rpc, author: "test"}
	statuses, err := client.UploadFiles(context.Background(), VolumeUploadRequest{
		VolumeName: "test-vol",
		Files:      []VolumeUploadFile{{Path: "pack.bin", Content: VolumeUploadBytes(fileData)}},
	}, nil)
	require.NoError(t, err)
	require.Len(t, statuses, 1)
	require.Equal(t, 1, putCount)
}

func TestVolumePackAndUploadMultiplePacks(t *testing.T) {
	t.Parallel()
	data1 := bytes.Repeat([]byte{0xAA}, 64)
	data2 := bytes.Repeat([]byte{0xBB}, 64)

	// Compute MaxPackBytes that fits exactly one 64-byte file but not two.
	b := newDataPackBuilder()
	b.append(blake3Sum(data1), data1)
	maxPackBytes := b.objectLen() + 1

	var uploadCount int
	rpc := &fakeVolumeRPC{
		requestUploadURLs: func(_ context.Context, req *connect.Request[volumev2.RequestUploadURLsRequest]) (*connect.Response[volumev2.RequestUploadURLsResponse], error) {
			uploadCount++
			urls := make([]*volumev2.UploadURLItem, len(req.Msg.GetObjects()))
			for i, obj := range req.Msg.GetObjects() {
				urls[i] = &volumev2.UploadURLItem{ObjectKey: obj.GetObjectKey(), AlreadyExists: true}
			}
			return connect.NewResponse(&volumev2.RequestUploadURLsResponse{Urls: urls}), nil
		},
		getVolume: func(_ context.Context, _ *connect.Request[volumev2.GetVolumeRequest]) (*connect.Response[volumev2.GetVolumeResponse], error) {
			return connect.NewResponse(&volumev2.GetVolumeResponse{
				Version: &volumev2.VersionInfo{VersionId: 1, SequenceNumber: 0},
			}), nil
		},
		commitVersion: func(_ context.Context, _ *connect.Request[volumev2.CommitVersionRequest]) (*connect.Response[volumev2.CommitVersionResponse], error) {
			return connect.NewResponse(&volumev2.CommitVersionResponse{
				Status: &volumev2.CommitStatus{Result: volumev2.CommitResult_COMMIT_RESULT_COMMITTED},
			}), nil
		},
	}
	client := &volumeClientImpl{rpc: rpc, author: "test"}
	statuses, err := client.UploadFiles(context.Background(), VolumeUploadRequest{
		VolumeName: "test-vol",
		Files: []VolumeUploadFile{
			{Path: "a.bin", Content: VolumeUploadBytes(data1)},
			{Path: "b.bin", Content: VolumeUploadBytes(data2)},
		},
		Config: VolumeUploadConfig{MaxPackBytes: maxPackBytes},
	}, nil)
	require.NoError(t, err)
	require.Len(t, statuses, 1)
	require.Equal(t, 2, uploadCount, "expected two separate pack uploads")
}

func TestVolumeResolveCommitAuthorFromAPI(t *testing.T) {
	t.Parallel()
	envID := "env-test"
	var requestCount int
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"user":"alice","environment_id":"env-test"}`))
	}))
	defer apiServer.Close()

	client, err := NewVolumeClient(context.Background(), &VolumeClientConfig{
		ApiServer:                  apiServer.URL,
		ClientId:                   "client",
		ClientSecret:               "secret",
		EnvironmentId:              envID,
		SkipEnvironmentNameMapping: true,
		JWT: &serverv1.GetTokenResponse{
			AccessToken:         "test-token",
			ExpiresAt:           timestamppb.New(time.Now().Add(time.Hour)),
			EnvironmentIdToName: map[string]string{envID: "test-env"},
		},
	})
	require.NoError(t, err)

	impl := client.(*volumeClientImpl)
	author := impl.resolveCommitAuthor(context.Background())
	require.Equal(t, "chalk:env-test:agent:alice", author)

	// Second call must use the cached value; the API must not be hit again.
	_ = impl.resolveCommitAuthor(context.Background())
	require.Equal(t, 1, requestCount)
}

func TestVolumeDownloadPackedNilVersionError(t *testing.T) {
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{
		getFile: func(_ context.Context, _ *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			return connect.NewResponse(&volumev2.GetFileResponse{
				File: &volumev2.FileInfo{Path: "f.bin"},
				// Version intentionally nil — triggers pinnedVolumeFileRequest error for packed path.
				Content: &volumev2.GetFileResponse_Packed{Packed: &volumev2.PackedFileContent{
					Pack: &volumev2.SignedPackEntryRef{SignedDownloadUri: "http://x", Offset: 0, Size: 4},
				}},
			}), nil
		},
	})
	_, _, err := client.DownloadBytes(context.Background(), VolumeDownloadRequest{
		VolumeName: "models", Path: "f.bin",
	}, nil)
	require.Error(t, err)
}

func TestVolumeCheckedRangeEndOverflow(t *testing.T) {
	t.Parallel()
	// Sanity-check the normal case.
	end, ok := checkedInclusiveRangeEnd(10, 5)
	require.True(t, ok)
	require.Equal(t, uint64(14), end)

	// Size=0 is always valid regardless of offset.
	_, ok = checkedInclusiveRangeEnd(^uint64(0), 0)
	require.True(t, ok)

	// offset + size overflows uint64.
	_, ok = checkedInclusiveRangeEnd(^uint64(0), 2)
	require.False(t, ok)
}

func TestVolumeSignedRequestNetworkError(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	url := server.URL
	server.Close() // cause connection refused on the next request

	client := &volumeClientImpl{httpClient: &http.Client{}}
	err := client.signedRequestWithRetry(
		context.Background(),
		http.MethodPut,
		url,
		nil,
		bytes.NewReader([]byte("data")),
		1, 0,
		func(context.Context) (string, error) { return url, nil },
	)
	require.Error(t, err)
}

func TestVolumeDownloadChunkedSizeMismatch(t *testing.T) {
	t.Parallel()
	body := []byte("short") // 5 bytes, but the chunk declares Size: 100
	client := newHTTPFakeClient(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(body)
	}, func(url string) *fakeVolumeRPC {
		return &fakeVolumeRPC{getFile: func(_ context.Context, _ *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			return connect.NewResponse(&volumev2.GetFileResponse{
				File:    &volumev2.FileInfo{Path: "f.bin"},
				Version: &volumev2.VersionInfo{VersionId: 1},
				Content: &volumev2.GetFileResponse_Chunked{Chunked: &volumev2.ChunkedFileContent{
					Chunks: []*volumev2.SignedChunkRef{{
						SignedDownloadUri: url,
						Offset:            0,
						Size:              100, // declared size doesn't match server response
					}},
				}},
			}), nil
		}}
	})
	_, _, err := client.DownloadBytes(context.Background(), VolumeDownloadRequest{
		VolumeName: "models",
		Path:       "f.bin",
		Config:     VolumeDownloadConfig{ChunkConcurrency: 1, MaxChunkRetries: 1},
	}, nil)
	require.Error(t, err)
}

func TestVolumeUploadOnePackEmptyURLsError(t *testing.T) {
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{
		requestUploadURLs: func(_ context.Context, _ *connect.Request[volumev2.RequestUploadURLsRequest]) (*connect.Response[volumev2.RequestUploadURLsResponse], error) {
			return connect.NewResponse(&volumev2.RequestUploadURLsResponse{Urls: nil}), nil
		},
	})
	builder := newDataPackBuilder()
	data := []byte("pack data")
	builder.append(blake3Sum(data), data)
	_, err := client.uploadOnePack(context.Background(), &volumev2.VolumeRef{Name: "vol"}, builder,
		[]packMember{{path: "f.bin", hash: blake3Sum(data)}},
		VolumeUploadConfig{}, func(uint64) {})
	require.Error(t, err)
}

func TestVolumeDownloadChunkedNilVersionError(t *testing.T) {
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{
		getFile: func(_ context.Context, _ *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			return connect.NewResponse(&volumev2.GetFileResponse{
				File: &volumev2.FileInfo{Path: "f.bin"},
				// Version intentionally nil — triggers pinnedVolumeFileRequest error
				Content: &volumev2.GetFileResponse_Chunked{Chunked: &volumev2.ChunkedFileContent{
					Chunks: []*volumev2.SignedChunkRef{{SignedDownloadUri: "http://example.com", Size: 4}},
				}},
			}), nil
		},
	})
	_, _, err := client.DownloadBytes(context.Background(), VolumeDownloadRequest{
		VolumeName: "models", Path: "f.bin",
	}, nil)
	require.Error(t, err)
}

func TestVolumeDownloadPackedNilPackError(t *testing.T) {
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{
		getFile: func(_ context.Context, _ *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
			return connect.NewResponse(&volumev2.GetFileResponse{
				File:    &volumev2.FileInfo{Path: "f.bin"},
				Version: &volumev2.VersionInfo{VersionId: 1},
				Content: &volumev2.GetFileResponse_Packed{Packed: &volumev2.PackedFileContent{Pack: nil}},
			}), nil
		},
	})
	_, _, err := client.DownloadBytes(context.Background(), VolumeDownloadRequest{
		VolumeName: "models", Path: "f.bin",
	}, nil)
	require.Error(t, err)
}

func TestVolumeSignedGetNetworkError(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	url := server.URL
	server.Close()

	client := &volumeClientImpl{httpClient: &http.Client{}}
	_, err := client.signedGetWithRetry(
		context.Background(),
		url,
		nil,
		1,
		func(context.Context) (string, error) { return url, nil },
	)
	require.Error(t, err)
}

func TestVolumeUploadOneFileMetadataError(t *testing.T) {
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{})
	file := VolumeUploadFile{
		Path:    "file.txt",
		Content: VolumeUploadLocalPath("/nonexistent-path-xyz-abc"),
	}
	_, err := client.uploadOneFile(context.Background(), &volumev2.VolumeRef{Name: "vol"}, file, 100, VolumeUploadConfig{ChunkSize: 64 * 1024}, func(uint64) {})
	require.Error(t, err)
}

func TestVolumeUploadOneFileReadChunkError(t *testing.T) {
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{})
	file := VolumeUploadFile{
		Path:     "file.txt",
		Content:  VolumeUploadBytes([]byte("hi")),
		Metadata: &volumev2.FileMetadata{},
	}
	// size=100 causes readChunk to fail since actual data is only 2 bytes
	_, err := client.uploadOneFile(context.Background(), &volumev2.VolumeRef{}, file, 100, VolumeUploadConfig{ChunkSize: 64 * 1024}, func(uint64) {})
	require.Error(t, err)
}

func TestVolumeUploadOneFileURLCountMismatch(t *testing.T) {
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{
		requestUploadURLs: func(_ context.Context, _ *connect.Request[volumev2.RequestUploadURLsRequest]) (*connect.Response[volumev2.RequestUploadURLsResponse], error) {
			return connect.NewResponse(&volumev2.RequestUploadURLsResponse{}), nil
		},
	})
	data := []byte("hello world data")
	file := VolumeUploadFile{
		Path:     "file.txt",
		Content:  VolumeUploadBytes(data),
		Metadata: &volumev2.FileMetadata{},
	}
	_, err := client.uploadOneFile(context.Background(), &volumev2.VolumeRef{}, file, uint64(len(data)), VolumeUploadConfig{ChunkSize: 64 * 1024}, func(uint64) {})
	require.Error(t, err)
}

func TestVolumeRPCWrapperErrors(t *testing.T) {
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{})
	vol := VolumeRef{Name: "vol"}
	ctx := context.Background()

	_, err := client.CreateVolume(ctx, "vol")
	require.Error(t, err)

	_, err = client.GetVolume(ctx, vol, nil)
	require.Error(t, err)

	_, err = client.ListVolumes(ctx, 10, "")
	require.Error(t, err)

	_, err = client.ListVolumeVersions(ctx, vol, 10, "")
	require.Error(t, err)

	_, err = client.ListFiles(ctx, ListVolumeFilesParams{Volume: vol})
	require.Error(t, err)

	_, err = client.GetFile(ctx, vol, "path.txt", nil)
	require.Error(t, err)
}

func TestVolumeUploadDirectoryNonExistentDir(t *testing.T) {
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{})
	_, err := client.UploadDirectory(context.Background(), "vol", "/nonexistent-dir-xyz-abc", VolumeUploadConfig{})
	require.Error(t, err)
}

func newTestCommitClient(getVolumeFn func(context.Context, *connect.Request[volumev2.GetVolumeRequest]) (*connect.Response[volumev2.GetVolumeResponse], error), commitVersionFn func(context.Context, *connect.Request[volumev2.CommitVersionRequest]) (*connect.Response[volumev2.CommitVersionResponse], error)) *volumeClientImpl {
	return newClientWithRPC(&fakeVolumeRPC{getVolume: getVolumeFn, commitVersion: commitVersionFn})
}

func fakeGetVolumeOK() func(context.Context, *connect.Request[volumev2.GetVolumeRequest]) (*connect.Response[volumev2.GetVolumeResponse], error) {
	return func(_ context.Context, _ *connect.Request[volumev2.GetVolumeRequest]) (*connect.Response[volumev2.GetVolumeResponse], error) {
		return connect.NewResponse(&volumev2.GetVolumeResponse{
			Version: &volumev2.VersionInfo{VersionId: 1},
		}), nil
	}
}

func TestVolumeCommitInlineContentError(t *testing.T) {
	t.Parallel()
	client := newTestCommitClient(fakeGetVolumeOK(), nil)
	upserts := []uploadedVolumeFile{{
		path: "file.txt",
		contentRef: &volumev2.ContentRef{
			Content: &volumev2.ContentRef_Inline{Inline: &volumev2.InlineFileContent{Data: []byte("hi")}},
		},
	}}
	_, err := client.commitPathDeltas(context.Background(), &volumev2.VolumeRef{Name: "vol"}, upserts, nil, 1)
	require.Error(t, err)
}

func TestVolumeCommitGetVolumeError(t *testing.T) {
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{
		getVolume: func(_ context.Context, _ *connect.Request[volumev2.GetVolumeRequest]) (*connect.Response[volumev2.GetVolumeResponse], error) {
			return nil, fakeVolumeUnimplemented()
		},
	})
	_, err := client.commitPathDeltas(context.Background(), &volumev2.VolumeRef{Name: "vol"}, nil, nil, 1)
	require.Error(t, err)
}

func TestVolumeCommitGetVolumeNilVersion(t *testing.T) {
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{
		getVolume: func(_ context.Context, _ *connect.Request[volumev2.GetVolumeRequest]) (*connect.Response[volumev2.GetVolumeResponse], error) {
			return connect.NewResponse(&volumev2.GetVolumeResponse{}), nil
		},
	})
	_, err := client.commitPathDeltas(context.Background(), &volumev2.VolumeRef{Name: "vol"}, nil, nil, 1)
	require.Error(t, err)
}

func TestVolumeCommitVersionError(t *testing.T) {
	t.Parallel()
	client := newTestCommitClient(fakeGetVolumeOK(), func(_ context.Context, _ *connect.Request[volumev2.CommitVersionRequest]) (*connect.Response[volumev2.CommitVersionResponse], error) {
		return nil, fakeVolumeUnimplemented()
	})
	_, err := client.commitPathDeltas(context.Background(), &volumev2.VolumeRef{Name: "vol"}, nil, nil, 1)
	require.Error(t, err)
}

func TestVolumeCommitVersionNilStatus(t *testing.T) {
	t.Parallel()
	client := newTestCommitClient(fakeGetVolumeOK(), func(_ context.Context, _ *connect.Request[volumev2.CommitVersionRequest]) (*connect.Response[volumev2.CommitVersionResponse], error) {
		return connect.NewResponse(&volumev2.CommitVersionResponse{}), nil
	})
	_, err := client.commitPathDeltas(context.Background(), &volumev2.VolumeRef{Name: "vol"}, nil, nil, 1)
	require.Error(t, err)
}

func TestVolumeCommitResultUnknown(t *testing.T) {
	t.Parallel()
	client := newTestCommitClient(fakeGetVolumeOK(), func(_ context.Context, _ *connect.Request[volumev2.CommitVersionRequest]) (*connect.Response[volumev2.CommitVersionResponse], error) {
		return connect.NewResponse(&volumev2.CommitVersionResponse{
			Status: &volumev2.CommitStatus{Result: volumev2.CommitResult(999)},
		}), nil
	})
	_, err := client.commitPathDeltas(context.Background(), &volumev2.VolumeRef{Name: "vol"}, nil, nil, 1)
	require.Error(t, err)
}

func TestVolumeCommitMaxRetriesExhausted(t *testing.T) {
	t.Parallel()
	client := newTestCommitClient(fakeGetVolumeOK(), func(_ context.Context, _ *connect.Request[volumev2.CommitVersionRequest]) (*connect.Response[volumev2.CommitVersionResponse], error) {
		return connect.NewResponse(&volumev2.CommitVersionResponse{
			Status: &volumev2.CommitStatus{Result: volumev2.CommitResult_COMMIT_RESULT_REBASE_REQUIRED},
		}), nil
	})
	_, err := client.commitPathDeltas(context.Background(), &volumev2.VolumeRef{Name: "vol"}, nil, nil, 1)
	require.Error(t, err)
}

func TestVolumeUploadOnePackRequestURLsError(t *testing.T) {
	t.Parallel()
	client := newClientWithRPC(&fakeVolumeRPC{
		requestUploadURLs: func(_ context.Context, _ *connect.Request[volumev2.RequestUploadURLsRequest]) (*connect.Response[volumev2.RequestUploadURLsResponse], error) {
			return nil, fakeVolumeUnimplemented()
		},
	})
	builder := newDataPackBuilder()
	data := []byte("upload-pack-data")
	h := blake3Sum(data)
	builder.append(h, data)
	_, err := client.uploadOnePack(context.Background(), &volumev2.VolumeRef{Name: "vol"}, builder,
		[]packMember{{path: "f.bin", hash: h}},
		VolumeUploadConfig{}, func(uint64) {})
	require.Error(t, err)
}

func TestVolumeUploadOnePackPUTFailure(t *testing.T) {
	t.Parallel()
	client := newHTTPFakeClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}, func(url string) *fakeVolumeRPC {
		return &fakeVolumeRPC{
			requestUploadURLs: func(_ context.Context, _ *connect.Request[volumev2.RequestUploadURLsRequest]) (*connect.Response[volumev2.RequestUploadURLsResponse], error) {
				return connect.NewResponse(&volumev2.RequestUploadURLsResponse{
					Urls: []*volumev2.UploadURLItem{{ObjectKey: "some-key", SignedUploadUri: url, AlreadyExists: false}},
				}), nil
			},
		}
	})
	builder := newDataPackBuilder()
	data := []byte("put-fail-data")
	h := blake3Sum(data)
	builder.append(h, data)
	_, err := client.uploadOnePack(context.Background(), &volumev2.VolumeRef{Name: "vol"}, builder,
		[]packMember{{path: "f.bin", hash: h}},
		VolumeUploadConfig{MaxChunkRetries: 1}, func(uint64) {})
	require.Error(t, err)
}

func TestVolumeCRUDPassthrough(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("CreateVolume", func(t *testing.T) {
		t.Parallel()
		var gotName string
		client := newClientWithRPC(&fakeVolumeRPC{
			createVolume: func(_ context.Context, req *connect.Request[volumev2.CreateVolumeRequest]) (*connect.Response[volumev2.CreateVolumeResponse], error) {
				gotName = req.Msg.GetName()
				return connect.NewResponse(&volumev2.CreateVolumeResponse{}), nil
			},
		})
		_, err := client.CreateVolume(ctx, "my-vol")
		require.NoError(t, err)
		require.Equal(t, "my-vol", gotName)
	})

	t.Run("GetVolume", func(t *testing.T) {
		t.Parallel()
		client := newClientWithRPC(&fakeVolumeRPC{
			getVolume: func(_ context.Context, req *connect.Request[volumev2.GetVolumeRequest]) (*connect.Response[volumev2.GetVolumeResponse], error) {
				require.Equal(t, "my-vol", req.Msg.GetVolume().GetName())
				return connect.NewResponse(&volumev2.GetVolumeResponse{
					Version: &volumev2.VersionInfo{VersionId: 5},
				}), nil
			},
		})
		resp, err := client.GetVolume(ctx, VolumeRef{Name: "my-vol"}, VolumeVersionSelector(5))
		require.NoError(t, err)
		require.Equal(t, uint64(5), resp.Version.VersionId)
	})

	t.Run("ListVolumes", func(t *testing.T) {
		t.Parallel()
		client := newClientWithRPC(&fakeVolumeRPC{
			listVolumes: func(_ context.Context, req *connect.Request[volumev2.ListVolumesRequest]) (*connect.Response[volumev2.ListVolumesResponse], error) {
				require.Equal(t, int32(10), req.Msg.GetLimit())
				require.Equal(t, "cursor1", req.Msg.GetCursor())
				return connect.NewResponse(&volumev2.ListVolumesResponse{}), nil
			},
		})
		_, err := client.ListVolumes(ctx, 10, "cursor1")
		require.NoError(t, err)
	})

	t.Run("DeleteVolume", func(t *testing.T) {
		t.Parallel()
		var gotVolume *volumev2.VolumeRef
		client := newClientWithRPC(&fakeVolumeRPC{
			deleteVolume: func(_ context.Context, req *connect.Request[volumev2.DeleteVolumeRequest]) (*connect.Response[volumev2.DeleteVolumeResponse], error) {
				gotVolume = req.Msg.GetVolume()
				return connect.NewResponse(&volumev2.DeleteVolumeResponse{}), nil
			},
		})
		err := client.DeleteVolume(ctx, VolumeRef{Name: "old-vol"})
		require.NoError(t, err)
		require.Equal(t, "old-vol", gotVolume.GetName())
	})

	t.Run("ListVolumeVersions", func(t *testing.T) {
		t.Parallel()
		client := newClientWithRPC(&fakeVolumeRPC{
			listVolumeVersions: func(_ context.Context, req *connect.Request[volumev2.ListVolumeVersionsRequest]) (*connect.Response[volumev2.ListVolumeVersionsResponse], error) {
				require.Equal(t, "my-vol", req.Msg.GetVolume().GetName())
				require.Equal(t, int32(5), req.Msg.GetLimit())
				return connect.NewResponse(&volumev2.ListVolumeVersionsResponse{}), nil
			},
		})
		_, err := client.ListVolumeVersions(ctx, VolumeRef{Name: "my-vol"}, 5, "")
		require.NoError(t, err)
	})

	t.Run("ListFiles", func(t *testing.T) {
		t.Parallel()
		client := newClientWithRPC(&fakeVolumeRPC{
			listFiles: func(_ context.Context, req *connect.Request[volumev2.ListFilesRequest]) (*connect.Response[volumev2.ListFilesResponse], error) {
				require.Equal(t, "my-vol", req.Msg.GetVolume().GetName())
				require.Equal(t, "models/", req.Msg.GetPath())
				require.True(t, req.Msg.GetRecursive())
				return connect.NewResponse(&volumev2.ListFilesResponse{
					Files: []*volumev2.FileInfo{{Path: "models/a.bin"}},
				}), nil
			},
		})
		resp, err := client.ListFiles(ctx, ListVolumeFilesParams{
			Volume:    VolumeRef{Name: "my-vol"},
			Path:      "models/",
			Recursive: true,
		})
		require.NoError(t, err)
		require.Len(t, resp.Files, 1)
		require.Equal(t, "models/a.bin", resp.Files[0].Path)
	})

	t.Run("GetFile", func(t *testing.T) {
		t.Parallel()
		client := newClientWithRPC(&fakeVolumeRPC{
			getFile: func(_ context.Context, req *connect.Request[volumev2.GetFileRequest]) (*connect.Response[volumev2.GetFileResponse], error) {
				require.Equal(t, "my-vol", req.Msg.GetVolume().GetName())
				require.Equal(t, "model.bin", req.Msg.GetPath())
				return connect.NewResponse(&volumev2.GetFileResponse{
					File: &volumev2.FileInfo{Path: "model.bin", Size: 3},
				}), nil
			},
		})
		resp, err := client.GetFile(ctx, VolumeRef{Name: "my-vol"}, "model.bin", nil)
		require.NoError(t, err)
		require.Equal(t, uint64(3), resp.File.Size)
	})
}

func newClientWithServer(t *testing.T, apiServer string) VolumeClient {
	t.Helper()
	envID := "env-test"
	client, err := NewVolumeClient(context.Background(), &VolumeClientConfig{
		ApiServer:                  apiServer,
		ClientId:                   "client",
		ClientSecret:               "secret",
		EnvironmentId:              envID,
		SkipEnvironmentNameMapping: true,
		JWT: &serverv1.GetTokenResponse{
			AccessToken:         "test-token",
			ExpiresAt:           timestamppb.New(time.Now().Add(time.Hour)),
			EnvironmentIdToName: map[string]string{envID: "test-env"},
		},
	})
	require.NoError(t, err)
	return client
}

// newFakeClient creates a volumeClientImpl backed by a fakeVolumeRPC for unit tests.
func newClientWithRPC(rpc *fakeVolumeRPC) *volumeClientImpl {
	return &volumeClientImpl{rpc: rpc, author: "test"}
}

// testHTTPServer starts an httptest.Server using handler and registers its Close
// as a test cleanup. Returns the server and its URL.
func testHTTPServer(t *testing.T, handler http.HandlerFunc) (*httptest.Server, string) {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	return srv, srv.URL
}

// newHTTPFakeClient starts an httptest.Server, registers cleanup, and returns a
// volumeClientImpl whose HTTP client points at that server. makeRPC receives the
// server URL so handlers can embed it in fake RPC responses.
func newHTTPFakeClient(t *testing.T, handler http.HandlerFunc, makeRPC func(url string) *fakeVolumeRPC) *volumeClientImpl {
	t.Helper()
	srv, url := testHTTPServer(t, handler)
	return &volumeClientImpl{rpc: makeRPC(url), httpClient: srv.Client(), author: "test"}
}

// packedGetFileResponse builds a GetFileResponse pointing at a pack at the given URL/offset/size.
func packedGetFileResponse(url string, offset, size uint64, fileHash string) *connect.Response[volumev2.GetFileResponse] {
	return connect.NewResponse(&volumev2.GetFileResponse{
		File:    &volumev2.FileInfo{Path: "f.bin", Size: size, Hash: fileHash},
		Version: &volumev2.VersionInfo{VersionId: 1},
		Content: &volumev2.GetFileResponse_Packed{Packed: &volumev2.PackedFileContent{
			Pack: &volumev2.SignedPackEntryRef{SignedDownloadUri: url, Offset: offset, Size: size},
		}},
	})
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
