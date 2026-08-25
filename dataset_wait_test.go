package chalk

import (
	"context"
	"errors"
	"testing"

	assert "github.com/stretchr/testify/require"
)

// mockWaitClient implements just GetOfflineQueryStatus; any other Client
// method panics via the embedded nil interface.
type mockWaitClient struct {
	Client
	getStatus func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error)
}

func (m *mockWaitClient) GetOfflineQueryStatus(
	ctx context.Context,
	args GetOfflineQueryStatusParams,
) (GetOfflineQueryStatusResult, error) {
	return m.getStatus(ctx, args)
}

func waitTestDataset(client Client, numPartitions int, numComputers int) Dataset {
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
	// recorded failure without polling the server again.
	assert.True(t, dataset.IsFinished)
	dataset.client = &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			t.Fatal("should not poll again after a terminal failure")
			return GetOfflineQueryStatusResult{}, nil
		},
	}
	assert.Equal(t, err, dataset.Wait(context.Background()))
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
	statuses := []string{"INIT", "COMPUTE_STARTED", "COMPUTE_ENDED", "COMPLETED"}
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

func TestDatasetWaitAlreadyFinished(t *testing.T) {
	t.Parallel()
	client := &mockWaitClient{
		getStatus: func(ctx context.Context, args GetOfflineQueryStatusParams) (GetOfflineQueryStatusResult, error) {
			t.Fatal("should not poll when dataset is already finished")
			return GetOfflineQueryStatusResult{}, nil
		},
	}
	dataset := waitTestDataset(client, 1, 0)
	dataset.IsFinished = true

	assert.NoError(t, dataset.Wait(context.Background()))
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
