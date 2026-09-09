package chalk

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	"github.com/stretchr/testify/require"
)

type aggregateBackfillService struct {
	planRequests   int
	createRequests []*aggregatev1.CreateAggregateBackfillJobRequest
}

func (h *aggregateBackfillService) PlanAggregateBackfill(
	_ context.Context,
	_ *connect.Request[aggregatev1.PlanAggregateBackfillRequest],
) (*connect.Response[aggregatev1.PlanAggregateBackfillResponse], error) {
	h.planRequests++
	return connect.NewResponse(&aggregatev1.PlanAggregateBackfillResponse{
		AggregateBackfillId: "aggregate-backfill-id",
		Backfills: []*aggregatev1.AggregateBackfillWithCostEstimate{{
			Backfill: &aggregatev1.AggregateBackfill{
				Resolver:        "resolver",
				DatetimeFeature: "events.timestamp",
				Series: []*aggregatev1.AggregateTimeSeries{{
					Rules: []*aggregatev1.AggregateTimeSeriesRule{{DependentFeatures: []string{"feature.count"}}},
				}},
			},
		}},
	}), nil
}

func TestTriggerAggregateBackfillRequiresBoundsForOfflineStorage(t *testing.T) {
	service := &aggregateBackfillService{}
	client := &clientImpl{aggregateClient: service}
	storeOffline := true

	result, err := client.TriggerAggregateBackfill(context.Background(), TriggerAggregateBackfillParams{
		Features:     []string{"feature.count"},
		StoreOffline: &storeOffline,
	})

	require.Nil(t, result)
	require.EqualError(t, err, "lower and upper bounds are required when StoreOffline is true")
	require.Zero(t, service.planRequests)
	require.Empty(t, service.createRequests)
}

func (h *aggregateBackfillService) CreateAggregateBackfillJob(
	_ context.Context,
	req *connect.Request[aggregatev1.CreateAggregateBackfillJobRequest],
) (*connect.Response[aggregatev1.CreateAggregateBackfillJobResponse], error) {
	h.createRequests = append(h.createRequests, req.Msg)
	return connect.NewResponse(&aggregatev1.CreateAggregateBackfillJobResponse{JobId: "job-id"}), nil
}

func (h *aggregateBackfillService) GetAggregates(context.Context, *connect.Request[aggregatev1.GetAggregatesRequest]) (*connect.Response[aggregatev1.GetAggregatesResponse], error) {
	return nil, nil
}

func (h *aggregateBackfillService) GetAggregateBackfillJobs(context.Context, *connect.Request[aggregatev1.GetAggregateBackfillJobsRequest]) (*connect.Response[aggregatev1.GetAggregateBackfillJobsResponse], error) {
	return nil, nil
}

func (h *aggregateBackfillService) GetAggregateBackfillJob(context.Context, *connect.Request[aggregatev1.GetAggregateBackfillJobRequest]) (*connect.Response[aggregatev1.GetAggregateBackfillJobResponse], error) {
	return nil, nil
}

func (h *aggregateBackfillService) GetCronAggregateBackfill(context.Context, *connect.Request[aggregatev1.GetCronAggregateBackfillRequest]) (*connect.Response[aggregatev1.GetCronAggregateBackfillResponse], error) {
	return nil, nil
}

func (h *aggregateBackfillService) CreateAggregateBackfillV2(context.Context, *connect.Request[aggregatev1.CreateAggregateBackfillV2Request]) (*connect.Response[aggregatev1.CreateAggregateBackfillV2Response], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}

func (h *aggregateBackfillService) GetActiveCronAggregateBackfills(context.Context, *connect.Request[aggregatev1.GetActiveCronAggregateBackfillsRequest]) (*connect.Response[aggregatev1.GetActiveCronAggregateBackfillsResponse], error) {
	return nil, nil
}

func TestTriggerAggregateBackfillPlanOnly(t *testing.T) {
	service := &aggregateBackfillService{}
	client := &clientImpl{aggregateClient: service}

	result, err := client.TriggerAggregateBackfill(context.Background(), TriggerAggregateBackfillParams{
		Features: []string{"feature.count"},
		PlanOnly: true,
	})

	require.NoError(t, err)
	require.Equal(t, "aggregate-backfill-id", result.Plan.AggregateBackfillId)
	require.Empty(t, result.Jobs)
	require.Empty(t, service.createRequests)
}

func TestTriggerAggregateBackfillSubmitsPlannedJobs(t *testing.T) {
	service := &aggregateBackfillService{}
	client := &clientImpl{aggregateClient: service}
	allowEmptyTiles := true

	result, err := client.TriggerAggregateBackfill(context.Background(), TriggerAggregateBackfillParams{
		Features:        []string{"feature.count"},
		QueryTags:       []string{"nightly"},
		AllowEmptyTiles: &allowEmptyTiles,
	})

	require.NoError(t, err)
	require.Len(t, result.Jobs, 1)
	require.Len(t, service.createRequests, 1)
	require.Equal(t, []string{"feature.count"}, service.createRequests[0].Features)
	require.Equal(t, "aggregate-backfill-id", service.createRequests[0].GetAggregateBackfillId())
	require.True(t, service.createRequests[0].GetAllowEmptyTiles())
}
