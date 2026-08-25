package chalk

import (
	"context"
	"errors"
	"testing"

	assert "github.com/stretchr/testify/require"
)

// mockWaitClient implements datasetClient. Unset hooks fall back to benign
// defaults so each test only has to describe the behavior it cares about.
type mockWaitClient struct {
	getStatus    func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error)
	getJobStatus func(ctx context.Context, revisionId string) (GetOfflineQueryJobResponse, error)
	getDataset   func(ctx context.Context, revisionId string) (Dataset, error)
	getUrls      func(ctx context.Context, revisionId string) ([]string, error)
	saveUrl      func(url string, directory string) error

	jobStatusCalls int
	refreshCalls   int
	urlCalls       int
}

func (m *mockWaitClient) GetOfflineQueryStatus(
	ctx context.Context,
	args GetOfflineQueryStatusParams,
) (GetOfflineQueryStatusResult, error) {
	return m.getStatus(ctx, args)
}

func (m *mockWaitClient) getOfflineQueryJobStatus(
	ctx context.Context,
	revisionId string,
) (GetOfflineQueryJobResponse, error) {
	m.jobStatusCalls++
	if m.getJobStatus == nil {
		return GetOfflineQueryJobResponse{IsFinished: true}, nil
	}
	return m.getJobStatus(ctx, revisionId)
}

func (m *mockWaitClient) getDatasetByRevisionId(ctx context.Context, revisionId string) (Dataset, error) {
	m.refreshCalls++
	if m.getDataset == nil {
		// The minimum a refresh has to return: the same revision, now marked
		// successful, as the server would once offline_query.status is COMPLETED.
		return Dataset{
			IsFinished: true,
			Revisions: []DatasetRevision{{
				RevisionId: revisionId,
				Status:     QueryStatusSuccessful,
				OutputUris: "s3://bucket/" + revisionId,
			}},
		}, nil
	}
	return m.getDataset(ctx, revisionId)
}

func (m *mockWaitClient) getDatasetUrls(ctx context.Context, revisionId string) ([]string, error) {
	m.urlCalls++
	if m.getUrls == nil {
		return []string{"https://example.com/" + revisionId + "/shard_000.parquet"}, nil
	}
	return m.getUrls(ctx, revisionId)
}

func (m *mockWaitClient) saveUrlToDirectory(url string, directory string) error {
	if m.saveUrl == nil {
		return nil
	}
	return m.saveUrl(url, directory)
}

func completedStatus(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
	return GetOfflineQueryStatusResult{Report: BatchReport{Status: "COMPLETED"}}, nil
}

func waitTestDataset(client datasetClient, numPartitions int, numComputers int) Dataset {
	return Dataset{
		client: client,
		Revisions: []DatasetRevision{
			{
				RevisionId:    "rev-1",
				NumPartitions: numPartitions,
				NumComputers:  numComputers,
			},
		},
	}
}

func TestDatasetWaitPollsAllShards(t *testing.T) {
	t.Parallel()
	var polledShards []int
	client := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			polledShards = append(polledShards, args.ComputerId)
			return GetOfflineQueryStatusResult{Report: BatchReport{Status: "COMPLETED"}}, nil
		},
	}
	dataset := waitTestDataset(client, 3, 0)

	assert.NoError(t, dataset.Wait(context.Background()))
	assert.True(t, dataset.IsFinished)
	assert.Equal(t, []int{0, 1, 2}, polledShards)
}

func TestDatasetWaitLegacyNumComputersFallback(t *testing.T) {
	t.Parallel()
	var polledShards []int
	client := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			polledShards = append(polledShards, args.ComputerId)
			return GetOfflineQueryStatusResult{Report: BatchReport{Status: "COMPLETED"}}, nil
		},
	}
	dataset := waitTestDataset(client, 0, 2)

	assert.NoError(t, dataset.Wait(context.Background()))
	assert.Equal(t, []int{0, 1}, polledShards)
}

func TestDatasetWaitShardFailure(t *testing.T) {
	t.Parallel()
	client := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			if args.ComputerId == 0 {
				return GetOfflineQueryStatusResult{Report: BatchReport{Status: "COMPLETED"}}, nil
			}
			return GetOfflineQueryStatusResult{
				Report: BatchReport{
					Status: "FAILED",
					AllErrors: []ServerError{
						{Message: "resolver exploded"},
						{Message: "upstream failed"},
					},
				},
			}, nil
		},
	}
	dataset := waitTestDataset(client, 2, 0)

	err := dataset.Wait(context.Background())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "offline query failed")
	assert.Contains(t, err.Error(), "resolver exploded")
	assert.Contains(t, err.Error(), "upstream failed")

	// Failure is terminal: IsFinished is set, and a repeated Wait reports the
	// recorded failure without polling the server again. The memo is recorded on
	// the revision, so swap the client there too -- otherwise this asserts
	// nothing, since Dataset.Wait only falls back to its own client when the
	// revision has none.
	assert.True(t, dataset.IsFinished)
	panicking := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			t.Fatal("should not poll again after a terminal failure")
			return GetOfflineQueryStatusResult{}, nil
		},
	}
	dataset.client = panicking
	dataset.Revisions[0].client = panicking
	assert.Equal(t, err, dataset.Wait(context.Background()))
	assert.Equal(t, 0, panicking.refreshCalls)
	assert.Equal(t, 0, panicking.jobStatusCalls)
}

func TestDatasetWaitToleratesTransientPollErrors(t *testing.T) {
	t.Parallel()
	calls := 0
	client := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			calls++
			if calls == 1 {
				return GetOfflineQueryStatusResult{}, errors.New("transient network error")
			}
			return GetOfflineQueryStatusResult{Report: BatchReport{Status: "COMPLETED"}}, nil
		},
	}
	dataset := waitTestDataset(client, 1, 0)

	assert.NoError(t, dataset.Wait(context.Background()))
	assert.Equal(t, 2, calls)
}

func TestDatasetWaitNonTerminalStatusesKeepPolling(t *testing.T) {
	t.Parallel()
	// SOME_FUTURE_STATUS stands in for a status introduced by a newer server:
	// anything that is neither COMPLETED nor FAILED keeps polling, matching
	// the Python client.
	statuses := []string{"INIT", "COMPUTE_STARTED", "COMPUTE_ENDED", "SOME_FUTURE_STATUS", "COMPLETED"}
	calls := 0
	client := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			status := statuses[calls]
			calls++
			return GetOfflineQueryStatusResult{Report: BatchReport{Status: status}}, nil
		},
	}
	dataset := waitTestDataset(client, 1, 0)

	assert.NoError(t, dataset.Wait(context.Background()))
	assert.Equal(t, len(statuses), calls)
}

func TestDatasetWaitContextCancellation(t *testing.T) {
	t.Parallel()
	client := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			return GetOfflineQueryStatusResult{Report: BatchReport{Status: "COMPUTE_STARTED"}}, nil
		},
	}
	dataset := waitTestDataset(client, 1, 0)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	assert.ErrorIs(t, dataset.Wait(ctx), context.Canceled)
}

func TestDatasetWaitSkipsAlreadySuccessfulRevision(t *testing.T) {
	t.Parallel()
	// The Python client seeds DatasetRevision._hydrated from the revision's own
	// status, so a revision the server already called SUCCESSFUL needs no work.
	client := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			t.Fatal("should not poll a revision the server already reported successful")
			return GetOfflineQueryStatusResult{}, nil
		},
	}
	dataset := waitTestDataset(client, 1, 0)
	dataset.Revisions[0].Status = QueryStatusSuccessful

	assert.NoError(t, dataset.Wait(context.Background()))
	assert.Equal(t, 0, client.refreshCalls)
}

func TestDatasetWaitDoesNotTrustDatasetLevelIsFinished(t *testing.T) {
	t.Parallel()
	// is_finished off the wire means "reached a terminal state", which is also
	// true of a failed or cancelled query, so it cannot decide whether waiting
	// is still needed. Both fields derive from OfflineQuerySQL.status server
	// side: COMPLETED yields is_finished=true AND status=SUCCESSFUL, while
	// FAILED and CANCELED yield is_finished=true with status=ERROR/CANCELLED.
	// Keying off is_finished reported those as successes.
	client := &mockWaitClient{getStatus: completedStatus}
	dataset := waitTestDataset(client, 2, 0)
	dataset.IsFinished = true
	dataset.Revisions[0].Status = QueryStatusError

	assert.NoError(t, dataset.Wait(context.Background()))
	assert.Equal(t, 1, client.refreshCalls)
	assert.Equal(t, 1, client.jobStatusCalls)
}

func TestDatasetWaitTimesOutWhenShardNeverReports(t *testing.T) {
	t.Parallel()
	// The status route answers 200 with a null report for a shard that has not
	// published one yet, which decodes to a zero-value BatchReport rather than
	// an error. Wait must not treat that as a received report, or it would
	// refresh the report deadline on every poll and never terminate.
	polls := 0
	client := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			polls++
			if args.ComputerId == 0 {
				return GetOfflineQueryStatusResult{Report: BatchReport{Status: "COMPLETED"}}, nil
			}
			return GetOfflineQueryStatusResult{}, nil
		},
	}
	dataset := waitTestDataset(client, 2, 0)
	dataset.waitReportTimeout = 2 * datasetWaitPollInterval

	err := dataset.Wait(context.Background())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "timed out waiting")
	assert.Contains(t, err.Error(), "shard 2 of 2")
	assert.Contains(t, err.Error(), "never reported a status")
	assert.False(t, dataset.IsFinished)
	// Bounded, not spinning forever: one poll for shard 0 plus a handful for
	// shard 1 before the deadline passes.
	assert.Less(t, polls, 10)
}

func TestDatasetWaitTimesOutOnPersistentPollError(t *testing.T) {
	t.Parallel()
	// A poll error that never clears must surface the underlying cause rather
	// than being swallowed. errors.Wrapf returns nil for a nil cause, so the
	// timeout error has to be built without a cause when polling never failed
	// and with one when it did.
	client := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			return GetOfflineQueryStatusResult{}, errors.New("connection refused")
		},
	}
	dataset := waitTestDataset(client, 1, 0)
	dataset.waitReportTimeout = 2 * datasetWaitPollInterval

	err := dataset.Wait(context.Background())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "timed out waiting")
	assert.Contains(t, err.Error(), "connection refused")
	assert.False(t, dataset.IsFinished)
}

func TestDatasetWaitReportDeadlineResetsOnNonTerminalReports(t *testing.T) {
	t.Parallel()
	// The deadline bounds the gap between reports, not the total wait, so a
	// shard that keeps reporting a non-terminal status must not time out even
	// though it runs well past the report timeout.
	polls := 0
	client := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			polls++
			if polls < 6 {
				return GetOfflineQueryStatusResult{Report: BatchReport{Status: "COMPUTE_STARTED"}}, nil
			}
			return GetOfflineQueryStatusResult{Report: BatchReport{Status: "COMPLETED"}}, nil
		},
	}
	dataset := waitTestDataset(client, 1, 0)
	dataset.waitReportTimeout = 2 * datasetWaitPollInterval

	assert.NoError(t, dataset.Wait(context.Background()))
	assert.True(t, dataset.IsFinished)
	assert.Equal(t, 6, polls)
}

func TestDatasetWaitNullReportBeforeShardStarts(t *testing.T) {
	t.Parallel()
	// The common healthy case: a shard's report row does not exist for the
	// first few polls, then appears and completes. Wait should ride through the
	// null reports without erroring.
	polls := 0
	client := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			polls++
			if polls < 3 {
				return GetOfflineQueryStatusResult{}, nil
			}
			return GetOfflineQueryStatusResult{Report: BatchReport{Status: "COMPLETED"}}, nil
		},
	}
	dataset := waitTestDataset(client, 1, 0)
	dataset.waitReportTimeout = 10 * datasetWaitPollInterval

	assert.NoError(t, dataset.Wait(context.Background()))
	assert.True(t, dataset.IsFinished)
	assert.Equal(t, 3, polls)
}

func TestDatasetWaitFailsOnRevisionLevelJobErrors(t *testing.T) {
	t.Parallel()
	// Every shard report can say COMPLETED while the revision-level job status
	// still carries errors. The Python client checks this at the end of
	// DatasetRevision.wait(); without it a failed query looks like a success.
	client := &mockWaitClient{
		getStatus: completedStatus,
		getJobStatus: func(ctx context.Context, revisionId string) (GetOfflineQueryJobResponse, error) {
			return GetOfflineQueryJobResponse{
				IsFinished: true,
				Errors:     []ServerError{{Message: "offline job failed"}},
			}, nil
		},
	}
	dataset := waitTestDataset(client, 2, 0)

	err := dataset.Wait(context.Background())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "offline query failed")
	assert.Contains(t, err.Error(), "offline job failed")
	// Terminal, so it is recorded and replayed like a shard failure.
	assert.True(t, dataset.IsFinished)
	assert.Equal(t, err, dataset.Wait(context.Background()))
	assert.Equal(t, 1, client.jobStatusCalls)
}

func TestDatasetWaitRefreshesRevisionMetadata(t *testing.T) {
	t.Parallel()
	// Until the refresh, every revision field is the submission-time value:
	// OutputUris is empty and NumBytes is unset.
	numBytes := 4096
	client := &mockWaitClient{
		getStatus: completedStatus,
		getDataset: func(ctx context.Context, revisionId string) (Dataset, error) {
			return Dataset{
				IsFinished: true,
				Revisions: []DatasetRevision{
					{RevisionId: "some-other-revision", Status: QueryStatusSuccessful},
					{
						RevisionId:    revisionId,
						Status:        QueryStatusSuccessful,
						OutputUris:    "s3://bucket/out",
						NumPartitions: 2,
						NumBytes:      &numBytes,
					},
				},
			}, nil
		},
	}
	dataset := waitTestDataset(client, 2, 0)
	assert.Empty(t, dataset.Revisions[0].OutputUris)

	assert.NoError(t, dataset.Wait(context.Background()))
	assert.Equal(t, "s3://bucket/out", dataset.Revisions[0].OutputUris)
	assert.Equal(t, QueryStatusSuccessful, dataset.Revisions[0].Status)
	assert.NotNil(t, dataset.Revisions[0].NumBytes)
	assert.Equal(t, 4096, *dataset.Revisions[0].NumBytes)
	// The refresh must not drop the client, or later calls on the revision fail.
	assert.NotNil(t, dataset.Revisions[0].client)
}

func TestDatasetWaitErrorsWhenRefreshOmitsRevision(t *testing.T) {
	t.Parallel()
	client := &mockWaitClient{
		getStatus: completedStatus,
		getDataset: func(ctx context.Context, revisionId string) (Dataset, error) {
			return Dataset{Revisions: []DatasetRevision{{RevisionId: "a-different-revision"}}}, nil
		},
	}
	dataset := waitTestDataset(client, 1, 0)

	err := dataset.Wait(context.Background())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "did not return revision")
	// Not a query failure, so not memoized as terminal.
	assert.False(t, dataset.IsFinished)
}

func TestDatasetWaitIsIdempotentAfterSuccess(t *testing.T) {
	t.Parallel()
	// The refresh marks the revision successful, so a second Wait short-circuits
	// on the revision's own status instead of re-polling every shard.
	polls := 0
	client := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			polls++
			return completedStatus(ctx, args)
		},
	}
	dataset := waitTestDataset(client, 2, 0)

	assert.NoError(t, dataset.Wait(context.Background()))
	assert.Equal(t, 2, polls)
	assert.NoError(t, dataset.Wait(context.Background()))
	assert.Equal(t, 2, polls)
	assert.Equal(t, 1, client.refreshCalls)
}

func TestDatasetDownloadUrisWaitsForAllShards(t *testing.T) {
	t.Parallel()
	// The reported failure mode: a multi-shard query is treated as complete when
	// the first shard finishes and the download then comes back empty. Every
	// data-access path has to wait first, as the Python client does by routing
	// them all through DatasetRevision._hydrate.
	var polledShards []int
	client := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			polledShards = append(polledShards, args.ComputerId)
			return completedStatus(ctx, args)
		},
	}
	dataset := waitTestDataset(client, 3, 0)

	urls, err := dataset.DownloadUris(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, []string{"https://example.com/rev-1/shard_000.parquet"}, urls)
	assert.Equal(t, []int{0, 1, 2}, polledShards)
	assert.Equal(t, 1, client.urlCalls)
}

func TestDatasetDownloadUrisFailsWhenAShardFails(t *testing.T) {
	t.Parallel()
	client := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			return GetOfflineQueryStatusResult{
				Report: BatchReport{Status: "FAILED", AllErrors: []ServerError{{Message: "shard died"}}},
			}, nil
		},
	}
	dataset := waitTestDataset(client, 2, 0)

	_, err := dataset.DownloadUris(context.Background())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "shard died")
	// Never asked for urls for a query that failed.
	assert.Equal(t, 0, client.urlCalls)
}

func TestDatasetRevisionDownloadDataWaitsForAllShards(t *testing.T) {
	t.Parallel()
	var polledShards []int
	saved := 0
	client := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			polledShards = append(polledShards, args.ComputerId)
			return completedStatus(ctx, args)
		},
		saveUrl: func(url string, directory string) error {
			saved++
			return nil
		},
	}
	revision := DatasetRevision{RevisionId: "rev-1", NumPartitions: 2, client: client}

	assert.NoError(t, revision.DownloadData(context.Background(), t.TempDir()))
	assert.Equal(t, []int{0, 1}, polledShards)
	assert.Equal(t, 1, saved)
}

func TestDatasetRevisionWaitRequiresClient(t *testing.T) {
	t.Parallel()
	revision := DatasetRevision{RevisionId: "rev-1"}
	assert.ErrorContains(t, revision.Wait(context.Background()), "client is not initialized")
}
