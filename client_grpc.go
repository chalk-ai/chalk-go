package chalk

import (
	"context"
	"github.com/apache/arrow/go/v16/arrow"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
	"time"
)

type clientGrpc struct {
	Client
	underlying GRPCClient
}

func newClientGrpc(ctx context.Context, cfg ClientConfig) (*clientGrpc, error) {
	nativeClient, err := newGrpcClient(
		ctx,
		&GRPCClientConfig{
			ApiServer:     cfg.ApiServer,
			ClientId:      cfg.ClientId,
			ClientSecret:  cfg.ClientSecret,
			EnvironmentId: cfg.EnvironmentId,
			Logger:        cfg.Logger,
			Branch:        cfg.Branch,
			QueryServer:   cfg.QueryServer,
			HTTPClient:    cfg.HTTPClient,
			DeploymentTag: cfg.DeploymentTag,
			ResourceGroup: cfg.ResourceGroup,
			Timeout:       cfg.Timeout,
		},
	)
	if err != nil {
		return nil, err
	}
	return &clientGrpc{underlying: nativeClient}, nil
}

func (c *clientGrpc) GetToken(ctx context.Context) (*TokenResult, error) {
	return c.underlying.GetToken(ctx)
}

func (c *clientGrpc) onlineQueryBulk(ctx context.Context, args OnlineQueryParamsComplete) (OnlineQueryBulkResult, error) {
	res, err := c.underlying.OnlineQueryBulk(ctx, args)
	if err != nil {
		return OnlineQueryBulkResult{}, errors.Wrapf(err, "executing online query")
	}

	if len(res.Errors) > 0 {
		convertedErrs, err := serverErrorsFromProto(res.Errors)
		if err != nil {
			return OnlineQueryBulkResult{}, errors.Wrapf(err, "converting server errors")
		}
		return OnlineQueryBulkResult{}, convertedErrs
	}

	scalars, err := internal.ConvertBytesToTable(res.GetScalarsData())
	if err != nil {
		return OnlineQueryBulkResult{}, errors.Wrapf(err, "deserializing scalars table")
	}

	groups := make(map[string]arrow.Table)
	for k, v := range res.GetGroupsData() {
		g, err := internal.ConvertBytesToTable(v)
		if err != nil {
			return OnlineQueryBulkResult{}, errors.Wrapf(
				err,
				"deserializing has-many table for feature '%s'",
				k,
			)
		}
		groups[k] = g
	}

	return OnlineQueryBulkResult{
		ScalarsTable: scalars,
		GroupsTables: groups,
		Meta:         queryMetaFromProto(res.GetResponseMeta()),
	}, nil
}

func (c *clientGrpc) OnlineQuery(ctx context.Context, args OnlineQueryParamsComplete, resultHolder any) (OnlineQueryResult, error) {
	res, err := c.underlying.OnlineQuery(ctx, args)
	if err != nil {
		return OnlineQueryResult{}, err
	}

	if len(res.GetErrors()) > 0 {
		convertedErrs, err := serverErrorsFromProto(res.GetErrors())
		if err != nil {
			return OnlineQueryResult{}, errors.Wrapf(err, "converting server errors")
		}
		return OnlineQueryResult{}, convertedErrs
	}

	if resultHolder != nil {
		if err := UnmarshalOnlineQueryResponse(res, resultHolder); err != nil {
			return OnlineQueryResult{}, err
		}
	}

	featureResults := make([]FeatureResult, 0)
	for _, r := range res.GetData().GetResults() {
		var value any
		if r.GetValue() != nil {
			value = r.GetValue().AsInterface()
		}
		var pkey any
		if r.GetPkey() != nil {
			pkey = r.GetPkey().AsInterface()
		}
		var timestamp time.Time
		if r.GetTs() != nil {
			timestamp = r.GetTs().AsTime()
		}
		serverErr, err := serverErrorFromProto(r.GetError())
		if err != nil {
			return OnlineQueryResult{}, errors.Wrap(err, "converting server error")
		}
		var featureMeta *FeatureResolutionMeta
		if r.GetMeta() != nil {
			metaRaw := r.GetMeta()
			featureMeta = &FeatureResolutionMeta{
				ChosenResolverFqn: metaRaw.GetChosenResolverFqn(),
				CacheHit:          metaRaw.GetCacheHit(),
				PrimitiveType:     metaRaw.GetPrimitiveType(),
				Version:           int(metaRaw.GetVersion()),
			}
		}

		featureResults = append(featureResults, FeatureResult{
			Field:     r.GetField(),
			Value:     value,
			Pkey:      pkey,
			Meta:      featureMeta,
			Error:     serverErr,
			Timestamp: &timestamp,
		})
	}

	features := make(map[string]FeatureResult)
	for _, result := range featureResults {
		features[result.Field] = result
	}

	return OnlineQueryResult{
		Data:     featureResults,
		Meta:     queryMetaFromProto(res.GetResponseMeta()),
		features: features,
	}, nil

}

func (c *clientGrpc) OnlineQueryBulk(ctx context.Context, args OnlineQueryParamsComplete) (OnlineQueryBulkResult, error) {
	return c.onlineQueryBulk(ctx, args)
}

func (c *clientGrpc) UpdateAggregates(ctx context.Context, args UpdateAggregatesParams) (UpdateAggregatesResult, error) {
	res, err := c.underlying.UpdateAggregates(ctx, args)
	if err != nil {
		return UpdateAggregatesResult{}, errors.Wrapf(err, "executing update aggregates")
	}

	if len(res.Errors) > 0 {
		convertedErrs, err := serverErrorsFromProto(res.Errors)
		if err != nil {
			return UpdateAggregatesResult{}, errors.Wrapf(err, "converting server errors")
		}
		return UpdateAggregatesResult{}, convertedErrs
	}

	return UpdateAggregatesResult{
		// When we made requests directly in this client we were able
		// to get the trace ID from the resonse trailing metadata. We
		// lost the ability to get the trace ID since we now depend
		// on methods on the native client.
		// TraceId: "",
	}, nil
}

func (c *clientGrpc) GetAggregates(ctx context.Context, features []string) (*aggregatev1.GetAggregatesResponse, error) {
	return c.underlying.GetAggregates(ctx, features)
}

func (c *clientGrpc) PlanAggregateBackfill(
	ctx context.Context,
	req *aggregatev1.PlanAggregateBackfillRequest,
) (*aggregatev1.PlanAggregateBackfillResponse, error) {
	return c.underlying.PlanAggregateBackfill(ctx, req)
}

func (c *clientGrpc) OfflineQuery(ctx context.Context, args OfflineQueryParamsComplete) (Dataset, error) {
	return Dataset{}, errors.New("not implemented")
}

func (c *clientGrpc) TriggerResolverRun(ctx context.Context, args TriggerResolverRunParams) (TriggerResolverRunResult, error) {
	return TriggerResolverRunResult{}, errors.New("not implemented")
}

func (c *clientGrpc) GetRunStatus(ctx context.Context, args GetRunStatusParams) (GetRunStatusResult, error) {
	return GetRunStatusResult{}, errors.New("not implemented")
}

func (c *clientGrpc) UploadFeatures(ctx context.Context, args UploadFeaturesParams) (UploadFeaturesResult, error) {
	return UploadFeaturesResult{}, errors.New("not implemented")
}
