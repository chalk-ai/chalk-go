package chalk

import (
	"connectrpc.com/connect"
	"context"
	"github.com/apache/arrow/go/v16/arrow"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
	"github.com/samber/lo"
	"time"
)

var (
	headerKeyDeploymentType = "x-chalk-deployment-type"
	headerKeyEnvironmentId  = "x-chalk-env-id"
	headerKeyServerType     = "x-chalk-server"
	headerKeyTraceId        = "x-chalk-trace-id"

	serverTypeApi    = "go-api"
	serverTypeEngine = "engine"
)

type clientGrpc struct {
	Client
	underlying GRPCClient
}

func newClientGrpc(cfg ClientConfig) (*clientGrpc, error) {
	nativeClient, err := newGrpcClient(GRPCClientConfig{
		ApiServer:     cfg.ApiServer,
		ClientId:      cfg.ClientId,
		ClientSecret:  cfg.ClientSecret,
		EnvironmentId: cfg.EnvironmentId,
		Logger:        cfg.Logger,
		Branch:        cfg.Branch,
		QueryServer:   cfg.QueryServer,
		HTTPClient:    cfg.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return &clientGrpc{underlying: nativeClient}, nil
}

func (c *clientGrpc) GetToken() (*TokenResult, error) {
	return c.underlying.GetToken()
}

func (c *clientGrpc) onlineQueryBulk(args OnlineQueryParamsComplete) (OnlineQueryBulkResult, error) {
	res, err := c.underlying.OnlineQueryBulk(context.Background(), args)
	if err != nil {
		return OnlineQueryBulkResult{}, wrapClientError(err, "error executing online query")
	}

	if len(res.Errors) > 0 {
		convertedErrs, err := serverErrorsFromProto(res.Errors)
		if err != nil {
			return OnlineQueryBulkResult{}, wrapClientError(err, "error converting server errors")
		}
		return OnlineQueryBulkResult{}, newServerError(convertedErrs)
	}

	scalars, err := internal.ConvertBytesToTable(res.GetScalarsData())
	if err != nil {
		return OnlineQueryBulkResult{}, wrapClientError(err, "error deserializing scalars table")
	}

	groups := make(map[string]arrow.Table)
	for k, v := range res.GetGroupsData() {
		g, err := internal.ConvertBytesToTable(v)
		if err != nil {
			return OnlineQueryBulkResult{}, errors.Wrapf(
				err,
				"error deserializing has-many table for feature '%s'",
				k,
			)
		}
		groups[k] = g
	}

	metaRaw := res.GetResponseMeta()

	var executionDuration float64
	if metaRaw.ExecutionDuration != nil {
		executionDuration = metaRaw.ExecutionDuration.AsDuration().Seconds()
	}

	var queryTimestamp *time.Time
	if metaRaw.QueryTimestamp != nil {
		queryTimestamp = lo.ToPtr(metaRaw.QueryTimestamp.AsTime())
	}

	return OnlineQueryBulkResult{
		ScalarsTable: scalars,
		GroupsTables: groups,
		Meta: &QueryMeta{
			ExecutionDurationS: executionDuration,
			DeploymentId:       metaRaw.DeploymentId,
			EnvironmentId:      metaRaw.EnvironmentId,
			EnvironmentName:    metaRaw.EnvironmentName,
			QueryId:            metaRaw.QueryId,
			QueryTimestamp:     queryTimestamp,
			QueryHash:          metaRaw.QueryHash,
		},
	}, nil
}

func (c *clientGrpc) OnlineQuery(args OnlineQueryParamsComplete, resultHolder any) (OnlineQueryResult, error) {
	c.underlying.OnlineQuery(context.Background(), args)

	//bulkInputs := make(map[string]any)
	//for k, singleValue := range args.underlying.inputs {
	//	slice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(singleValue)), 1, 1)
	//	slice.Index(0).Set(reflect.ValueOf(singleValue))
	//	bulkInputs[k] = slice.Interface()
	//}
	//args.underlying.inputs = bulkInputs
	//
	//bulkRes, err := c.onlineQueryBulk(args)
	//if err != nil {
	//	// intentionally don't wrap, original error is good enough
	//	return OnlineQueryResult{}, err
	//}
	//
	//if resultHolder != nil {
	//	// Create a pointer to an empty slice and unmarshal into that
	//	var holderType = reflect.TypeOf(resultHolder)
	//	if holderType.Kind() != reflect.Ptr {
	//		return OnlineQueryResult{}, errors.Newf(
	//			"result holder must be a pointer, found %s",
	//			holderType.Kind(),
	//		)
	//	}
	//	structType := holderType.Elem()
	//	sliceType := reflect.SliceOf(structType)
	//	ptrToSlice := reflect.New(sliceType)
	//	if err := bulkRes.UnmarshalInto(ptrToSlice.Interface()); err != nil {
	//		return OnlineQueryResult{}, errors.Wrap(err, "error unmarshalling result into result holder struct")
	//	}
	//
	//	if ptrToSlice.Elem().Len() != 1 {
	//		return OnlineQueryResult{}, errors.Newf(
	//			"expected 1 element in the intermediate slice after unmarshalling, got %d",
	//			ptrToSlice.Elem().Len(),
	//		)
	//	}
	//
	//	// Point the result holder to the first element of the slice
	//	var holderValue = reflect.ValueOf(resultHolder)
	//	holderValue.Elem().Set(ptrToSlice.Elem().Index(0))
	//}
	//
	//rows, err := internal.ExtractFeaturesFromTable(bulkRes.ScalarsTable)
	//if err != nil {
	//	return OnlineQueryResult{}, errors.Wrap(err, "error extracting features from scalars table")
	//}
	//
	//features := make(map[string]FeatureResult)
	//if len(rows) != 1 {
	//	return OnlineQueryResult{}, errors.Newf(
	//		"expected 1 row from scalars table, got %d",
	//		len(rows),
	//	)
	//}
	//for fqn, value := range rows[0] {
	//	features[fqn] = FeatureResult{
	//		Field: fqn,
	//		Value: value,
	//	}
	//}
	//
	//for fqn, table := range bulkRes.GroupsTables {
	//	rowsHm, err := internal.ExtractFeaturesFromTable(table)
	//	if err != nil {
	//		return OnlineQueryResult{}, errors.Wrapf(
	//			err,
	//			"error extracting features from has-many table for feature '%s'",
	//			fqn,
	//		)
	//	}
	//
	//	colNames := lo.Map(
	//		table.Schema().Fields(),
	//		func(f arrow.Field, _ int) string {
	//			return f.Name
	//		},
	//	)
	//	colValues := make([][]any, 0, len(rowsHm))
	//	for _, col := range colNames {
	//		colValues = append(
	//			colValues,
	//			lo.Map(rowsHm, func(row map[string]any, _ int) any {
	//				return row[col]
	//			}),
	//		)
	//	}
	//	if err != nil {
	//		return OnlineQueryResult{}, errors.Wrapf(
	//			err,
	//			"error creating JSON representation for has-many table for feature '%s'",
	//			fqn,
	//		)
	//	}
	//
	//	features[fqn] = FeatureResult{
	//		Field: fqn,
	//		Value: map[string]any{
	//			"columns": colNames,
	//			"values":  colValues,
	//		},
	//	}
	//}
	//
	//return OnlineQueryResult{
	//	Data:     lo.Values(features),
	//	Meta:     bulkRes.Meta,
	//	features: features,
	//}, nil
}

func (c *clientGrpc) OnlineQueryBulk(args OnlineQueryParamsComplete) (OnlineQueryBulkResult, error) {
	return c.onlineQueryBulk(args)
}

func (c *clientGrpc) UpdateAggregates(args UpdateAggregatesParams) (UpdateAggregatesResult, error) {
	inputsConverted, err := getConvertedInputsMap(args.Inputs)
	if err != nil {
		return UpdateAggregatesResult{}, wrapClientError(err, "error converting inputs map")
	}
	inputsFeather, err := internal.InputsToArrowBytes(inputsConverted)
	if err != nil {
		return UpdateAggregatesResult{}, wrapClientError(err, "error serializing inputs as feather")
	}

	req := connect.NewRequest(&commonv1.UploadFeaturesBulkRequest{
		InputsFeather: inputsFeather,
		BodyType:      commonv1.FeatherBodyType_FEATHER_BODY_TYPE_TABLE,
	})

	ctx := context.Background()
	if args.Context != nil {
		ctx = args.Context
	}

	res, err := c.queryClient.UploadFeaturesBulk(ctx, req)
	if err != nil {
		return UpdateAggregatesResult{}, wrapClientError(err, "error making upload features request")
	}

	if len(res.Msg.Errors) > 0 {
		convertedErrs, err := serverErrorsFromProto(res.Msg.Errors)
		if err != nil {
			return UpdateAggregatesResult{}, errors.Wrap(err, "error converting server errors")
		}
		return UpdateAggregatesResult{}, newServerError(convertedErrs)
	}
	return UpdateAggregatesResult{
		res.Trailer().Get(headerKeyTraceId),
	}, nil
}

func (c *clientGrpc) GetAggregates(ctx context.Context, features []string) (*aggregatev1.GetAggregatesResponse, error) {
	req := connect.NewRequest(&aggregatev1.GetAggregatesRequest{
		ForFeatures: features,
	})
	res, err := c.queryClient.GetAggregates(ctx, req)
	if err != nil {
		return nil, wrapClientError(err, "fetching aggregates")
	}

	return res.Msg, err
}

func (c *clientGrpc) PlanAggregateBackfill(
	ctx context.Context,
	req *aggregatev1.PlanAggregateBackfillRequest,
) (*aggregatev1.PlanAggregateBackfillResponse, error) {
	res, err := c.queryClient.PlanAggregateBackfill(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, wrapClientError(err, "planning aggregate backfill")
	}
	return res.Msg, err
}

func (c *clientGrpc) OfflineQuery(args OfflineQueryParamsComplete) (Dataset, error) {
	return Dataset{}, errors.New("not implemented")
}

func (c *clientGrpc) TriggerResolverRun(args TriggerResolverRunParams) (TriggerResolverRunResult, error) {
	return TriggerResolverRunResult{}, errors.New("not implemented")
}

func (c *clientGrpc) GetRunStatus(args GetRunStatusParams) (GetRunStatusResult, error) {
	return GetRunStatusResult{}, errors.New("not implemented")
}

func (c *clientGrpc) UploadFeatures(args UploadFeaturesParams) (UploadFeaturesResult, error) {
	return UploadFeaturesResult{}, errors.New("not implemented")
}
