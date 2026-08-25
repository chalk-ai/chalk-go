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
	assert.False(t, dataset.IsFinished)
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
