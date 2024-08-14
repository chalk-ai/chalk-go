package chalk

import (
	"connectrpc.com/connect"
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	aggregatev1 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
	"github.com/samber/lo"
	"net/http"
	"reflect"
	"strings"
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
	config *configManager

	branch     string
	logger     LeveledLogger
	httpClient *http.Client

	authClient  serverv1connect.AuthServiceClient
	queryClient enginev1connect.QueryServiceClient
}

func newClientGrpc(cfg ClientConfig) (*clientGrpc, error) {
	config, err := getConfigManager(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error getting resolved config")
	}
	var logger LeveledLogger
	if cfg.Logger == nil {
		logger = DefaultLeveledLogger
	}
	client := &clientGrpc{
		branch:     cfg.Branch,
		httpClient: http.DefaultClient,
		logger:     logger,
		config:     config,
	}
	if err := client.init(); err != nil {
		return nil, errors.Wrap(err, "error initializing gRPC service clients")
	}

	return client, nil
}

func (c *clientGrpc) init() error {
	authClient, err := c.NewAuthClient()
	if err != nil {
		return errors.Wrap(err, "error creating auth client")
	}
	c.authClient = authClient

	c.config.getToken = c.getToken
	// Necessary to get GRPC engines URL
	if err := c.config.refresh(false); err != nil {
		return errors.Wrap(err, "error fetching initial config")
	}

	queryClient, err := c.NewQueryClient()
	if err != nil {
		return errors.Wrap(err, "error creating query client")
	}
	c.queryClient = queryClient

	return nil
}

func withChalkInterceptors(serverType string, interceptors ...connect.Interceptor) connect.Option {
	return connect.WithInterceptors(
		append(
			interceptors,
			headerInterceptor(map[string]string{
				headerKeyServerType: serverType,
			}),
		)...,
	)
}

func (c *clientGrpc) NewAuthClient() (serverv1connect.AuthServiceClient, error) {
	return serverv1connect.NewAuthServiceClient(
		c.httpClient,
		c.config.apiServer.Value,
		withChalkInterceptors(
			serverTypeApi,
			headerInterceptor(map[string]string{
				headerKeyServerType: serverTypeApi,
			}),
		),
	), nil
}

func (c *clientGrpc) NewQueryClient() (enginev1connect.QueryServiceClient, error) {
	endpoint, ok := c.config.engines[c.config.environmentId.Value]
	if !ok {
		c.logger.Errorf(
			"query endpoint falling back to api server - no engine found for environment '%s' - engine map keys: '%v'",
			c.config.environmentId.Value,
			&c.config.engines,
		)
		endpoint = c.config.apiServer.Value
	}

	return enginev1connect.NewQueryServiceClient(
		c.httpClient,
		ensureHTTPSPrefix(endpoint),
		withChalkInterceptors(
			serverTypeEngine,
			c.tokenInterceptor(),
			headerInterceptor(map[string]string{
				headerKeyDeploymentType: "engine-grpc",
			}),
		),
		connect.WithGRPC(),
	), nil
}

func ensureHTTPSPrefix(inputURL string) string {
	if strings.HasPrefix(inputURL, "https://") || strings.HasPrefix(inputURL, "http://") {
		return inputURL
	}
	return "https://" + inputURL
}

func headerInterceptor(headers map[string]string) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			for k, v := range headers {
				req.Header().Set(k, v)
			}
			return next(ctx, req)
		}
	}
}

func (c *clientGrpc) GetToken() (*TokenResult, error) {
	getTokenResult, err := c.getToken()
	if err != nil {
		return nil, getErrorResponse(err)
	}
	return &TokenResult{
		AccessToken:        getTokenResult.AccessToken,
		PrimaryEnvironment: getTokenResult.PrimaryEnvironment,
		ValidUntil:         getTokenResult.ValidUntil,
		Engines:            getTokenResult.Engines,
	}, nil
}

func (c *clientGrpc) getToken() (*getTokenResult, error) {
	c.logger.Debugf("Getting new token via gRPC")
	authRequest := connect.NewRequest(
		&serverv1.GetTokenRequest{
			ClientId:     c.config.clientId.Value,
			ClientSecret: c.config.clientSecret.Value,
			GrantType:    "client_credentials",
		},
	)
	token, err := c.authClient.GetToken(context.Background(), authRequest)
	if err != nil {
		c.logger.Debugf("Failed to get a new token: %s", err.Error())
		return nil, err
	}

	expiresAt := token.Msg.GetExpiresAt()
	if token.Msg.GetExpiresAt() == nil {
		return nil, errors.New("token has no expiration date")
	}

	return &getTokenResult{
		ValidUntil:         expiresAt.AsTime(),
		AccessToken:        token.Msg.GetAccessToken(),
		PrimaryEnvironment: token.Msg.GetPrimaryEnvironment(),
		Engines:            token.Msg.GetGrpcEngines(),
	}, nil
}

func (c *clientGrpc) tokenInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if err := c.config.refresh(false); err != nil {
				return nil, errors.Wrap(err, "error refreshing config")
			}
			req.Header().Set(headerKeyEnvironmentId, c.config.environmentId.Value)
			req.Header().Set("Authorization", fmt.Sprintf("Bearer %s", c.config.jwt.Token))
			return next(ctx, req)
		}
	}
}

func (c *clientGrpc) getHasManyJson(columns []string, values [][]any) (string, error) {
	result := struct {
		Columns []string `json:"columns"`
		Values  [][]any  `json:"values"`
	}{
		Columns: columns,
		Values:  values,
	}
	res, err := json.Marshal(result)
	return string(res), err
}

func (c *clientGrpc) onlineQueryBulk(args OnlineQueryParamsComplete) (OnlineQueryBulkResult, error) {
	paramsProto, err := convertOnlineQueryParamsToProto(&args.underlying)
	if err != nil {
		return OnlineQueryBulkResult{}, errors.Wrap(err, "error converting online query params to proto")
	}
	req := connect.NewRequest(paramsProto)

	res, err := c.queryClient.OnlineQueryBulk(context.Background(), req)
	if err != nil {
		return OnlineQueryBulkResult{}, wrapClientError(err, "error executing online query")
	}

	if len(res.Msg.Errors) > 0 {
		convertedErrs, err := serverErrorsFromProto(res.Msg.Errors)
		if err != nil {
			return OnlineQueryBulkResult{}, wrapClientError(err, "error converting server errors")
		}
		return OnlineQueryBulkResult{}, newServerError(convertedErrs)
	}

	scalars, err := internal.ConvertBytesToTable(res.Msg.GetScalarsData())
	if err != nil {
		return OnlineQueryBulkResult{}, wrapClientError(err, "error deserializing scalars table")
	}

	groups := make(map[string]arrow.Table)
	for k, v := range res.Msg.GetGroupsData() {
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

	metaRaw := res.Msg.GetResponseMeta()

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
	bulkInputs := make(map[string]any)
	for k, singleValue := range args.underlying.inputs {
		slice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(singleValue)), 1, 1)
		slice.Index(0).Set(reflect.ValueOf(singleValue))
		bulkInputs[k] = slice.Interface()
	}
	args.underlying.inputs = bulkInputs

	bulkRes, err := c.onlineQueryBulk(args)
	if err != nil {
		// intentionally don't wrap, original error is good enough
		return OnlineQueryResult{}, err
	}

	if resultHolder != nil {
		// Create a pointer to an empty slice and unmarshal into that
		var holderType = reflect.TypeOf(resultHolder)
		if holderType.Kind() != reflect.Ptr {
			return OnlineQueryResult{}, errors.Newf(
				"result holder must be a pointer, found %s",
				holderType.Kind(),
			)
		}
		structType := holderType.Elem()
		sliceType := reflect.SliceOf(structType)
		ptrToSlice := reflect.New(sliceType)
		if err := bulkRes.UnmarshalInto(ptrToSlice.Interface()); err != nil {
			return OnlineQueryResult{}, errors.Wrap(err, "error unmarshalling result into result holder struct")
		}

		if ptrToSlice.Elem().Len() != 1 {
			return OnlineQueryResult{}, errors.Newf(
				"expected 1 element in the intermediate slice after unmarshalling, got %d",
				ptrToSlice.Elem().Len(),
			)
		}

		// Point the result holder to the first element of the slice
		var holderValue = reflect.ValueOf(resultHolder)
		holderValue.Elem().Set(ptrToSlice.Elem().Index(0))
	}

	rows, err := internal.ExtractFeaturesFromTable(bulkRes.ScalarsTable)
	if err != nil {
		return OnlineQueryResult{}, errors.Wrap(err, "error extracting features from scalars table")
	}

	features := make(map[string]FeatureResult)
	if len(rows) != 1 {
		return OnlineQueryResult{}, errors.Newf(
			"expected 1 row from scalars table, got %d",
			len(rows),
		)
	}
	for fqn, value := range rows[0] {
		features[fqn] = FeatureResult{
			Field: fqn,
			Value: value,
		}
	}

	for fqn, table := range bulkRes.GroupsTables {
		if table.NumRows() == 0 {
			columns := lo.Map(
				table.Schema().Fields(),
				func(f arrow.Field, _ int) string {
					return f.Name
				},
			)
			values := make([][]any, len(columns))
			jsonRepr, err := c.getHasManyJson(columns, values)
			if err != nil {
				return OnlineQueryResult{}, errors.Wrapf(
					err,
					"error creating JSON representation for 0-row has-many table for feature '%s'",
					fqn,
				)
			}
			features[fqn] = FeatureResult{
				Field: fqn,
				Value: jsonRepr,
			}
			continue
		}

		rowsHm, err := internal.ExtractFeaturesFromTable(table)
		if err != nil {
			return OnlineQueryResult{}, errors.Wrapf(
				err,
				"error extracting features from has-many table for feature '%s'",
				fqn,
			)
		}

		colNames := lo.Keys(rowsHm[0])
		colValues := make([][]any, 0, len(rowsHm))
		for _, col := range colNames {
			colValues = append(
				colValues,
				lo.Map(rowsHm, func(row map[string]any, _ int) any {
					return row[col]
				}),
			)
		}
		jsonRepr, err := c.getHasManyJson(colNames, colValues)
		if err != nil {
			return OnlineQueryResult{}, errors.Wrapf(
				err,
				"error creating JSON representation for has-many table for feature '%s'",
				fqn,
			)
		}

		features[fqn] = FeatureResult{
			Field: fqn,
			Value: jsonRepr,
		}
	}

	return OnlineQueryResult{
		Data:     lo.Values(features),
		Meta:     bulkRes.Meta,
		features: features,
	}, nil
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
		return nil, wrapClientError(err, "error making upload features request")
	}

	return res.Msg, err
}

func (c *clientGrpc) PlanAggregateBackfill(
	ctx context.Context,
	req *aggregatev1.PlanAggregateBackfillRequest,
) (*aggregatev1.PlanAggregateBackfillResponse, error) {
	res, err := c.queryClient.PlanAggregateBackfill(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, wrapClientError(err, "error making upload features request")
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
