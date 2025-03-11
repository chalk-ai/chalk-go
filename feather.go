package chalk

import (
	"encoding/json"
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/colls"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"time"
)

func (r OnlineQueryBulkResult) Release() {
	r.ScalarsTable.Release()
	for _, table := range r.GroupsTables {
		table.Release()
	}
}

type SerializationOptions struct {
	ClientConfigBranchId string
	Allocator            memory.Allocator
}

func (p OnlineQueryParamsComplete) ToBytes(options ...*SerializationOptions) ([]byte, error) {
	branchId := p.underlying.BranchId
	allocator := memory.DefaultAllocator
	if len(options) > 1 {
		return nil, fmt.Errorf("expected 1 SerializationOptions, got %d", len(options))
	} else if len(options) == 1 {
		if branchId == nil || *branchId == "" && options[0].ClientConfigBranchId != "" {
			branchId = &options[0].ClientConfigBranchId
		}
		if options[0].Allocator != nil {
			allocator = options[0].Allocator
		}
	}

	convertedStaleness := make(map[string]string)
	for k, v := range p.underlying.staleness {
		convertedStaleness[k] = internal.FormatBucketDuration(int(v.Seconds()))
	}

	outputs := p.underlying.outputs
	if outputs == nil {
		// `outputs` is a non-optional field
		outputs = []string{}
	}

	return internal.CreateOnlineQueryBulkBody(
		p.underlying.inputs,
		internal.FeatherRequestHeader{
			Outputs:     outputs,
			Explain:     p.underlying.Explain,
			IncludeMeta: p.underlying.IncludeMeta || p.underlying.Explain,
			BranchId:    branchId,
			Context: &internal.OnlineQueryContext{
				Environment:          ptr.PtrOrNil(p.underlying.EnvironmentId),
				Tags:                 p.underlying.Tags,
				RequiredResolverTags: p.underlying.RequiredResolverTags,
			},
			StorePlanStages:  p.underlying.StorePlanStages,
			CorrelationId:    ptr.PtrOrNil(p.underlying.CorrelationId),
			QueryName:        ptr.PtrOrNil(p.underlying.QueryName),
			QueryNameVersion: ptr.PtrOrNil(p.underlying.QueryNameVersion),
			QueryContext:     p.underlying.QueryContext.ToMap(),
			Meta:             p.underlying.Meta,
			Staleness:        convertedStaleness,
			Now: colls.Map(p.underlying.Now, func(val time.Time) string {
				return val.Format(internal.NowTimeFormat)
			}),
			PlannerOptions: p.underlying.PlannerOptions,
		},
		allocator,
	)
}

func (r *OnlineQueryBulkResponse) Unmarshal(body []byte) error {
	res, err := internal.ChalkUnmarshal(body)
	if err != nil {
		return fmt.Errorf("failed to unmarshal bytes: %w", err)
	}
	resultMap := map[QueryName]onlineQueryResultFeather{}

	queryNameToBytesInBytes, ok := res["query_results_bytes"]
	if !ok {
		return fmt.Errorf("malformed bulk online query response - missing 'query_results_bytes' attribute")
	}

	queryNameToBytesInBytesCast, ok := queryNameToBytesInBytes.([]byte)
	if !ok {
		return fmt.Errorf("failed to convert 'query_results_bytes' value to a byte array")
	}
	queryNameToBytesInMap, err := internal.ChalkUnmarshal(queryNameToBytesInBytesCast)
	if err != nil {
		return fmt.Errorf("failed to unmarshal 'query_results_bytes' value: %w", err)
	}

	allocator := r.allocator
	if allocator == nil {
		allocator = memory.DefaultAllocator
	}

	for queryName, queryResultBytes := range queryNameToBytesInMap {
		queryResultBytesCast, ok := queryResultBytes.([]byte)
		if !ok {
			return fmt.Errorf("failed to cast bytes to byte array for query name: %s", queryName)
		}
		resultFeather := onlineQueryResultFeather{}

		err := resultFeather.Unmarshal(queryResultBytesCast, allocator)
		if err != nil {
			return fmt.Errorf("failed to unmarshal result bytes for query name '%s': %w", queryName, err)
		}
		resultMap[queryName] = resultFeather
	}

	r.QueryResults = resultMap
	return nil
}

func (r *onlineQueryResultFeather) Unmarshal(body []byte, allocator memory.Allocator) error {
	res, err := internal.ChalkUnmarshal(body)
	if err != nil {
		return fmt.Errorf("failed to unmarshal bytes: %w", err)
	}

	hasData, ok := res["has_data"]
	if !ok {
		return fmt.Errorf("missing attribute 'has_data'")
	}
	hasDataBool, ok := hasData.(bool)
	if !ok {
		return fmt.Errorf("cannot cast attribute 'has_data' to bool")
	}

	var table arrow.Table
	var groupsTables map[string]arrow.Table
	if hasDataBool {
		scalarDataBytes, ok := res["scalar_data"]
		if !ok {
			return fmt.Errorf("missing attribute 'scalar_data'")
		}
		scalarDataBytesCast, ok := scalarDataBytes.([]byte)
		if !ok {
			return fmt.Errorf("failed to cast scalar data bytes to bytes array")
		}
		table, err = internal.ConvertBytesToTable(scalarDataBytesCast, allocator)
		if err != nil {
			return fmt.Errorf("failed to convert scalar data bytes to an Arrow Table: %w", err)
		}

		groupsDataBytes, ok := res["groups_data"]
		if !ok {
			return fmt.Errorf("missing attribute 'groups_data'")
		}
		groupsDataBytesCast, ok := groupsDataBytes.([]byte)
		if !ok {
			return fmt.Errorf("failed to cast groups data bytes to bytes array")
		}
		groupsDataMap, err := internal.ChalkUnmarshal(groupsDataBytesCast)
		if err != nil {
			return fmt.Errorf("failed to unmarshal 'groups_data' value: %w", err)
		}

		groupsTables = map[string]arrow.Table{}
		for k, v := range groupsDataMap {
			vBytes, ok := v.([]byte)
			if !ok {
				return fmt.Errorf("failed to cast data for has-many feature '%s': %w", k, err)
			}
			vTable, err := internal.ConvertBytesToTable(vBytes, allocator)
			if err != nil {
				return fmt.Errorf("failed to convert bytes for has-many feature '%s' to Arrow table batch: %w", k, err)
			}
			groupsTables[k] = vTable
		}
	}

	r.HasData = hasDataBool
	r.ScalarData = table
	r.GroupsData = groupsTables

	errorsAny, ok := res["errors"]
	if !ok {
		return fmt.Errorf("missing attribute 'errors'")
	}

	var errorsAnyArr []any
	if errorsAny == nil {
		errorsAnyArr = []any{}
	} else {
		errorsAnyArr, ok = errorsAny.([]any)
		if !ok {
			return fmt.Errorf("cannot cast attribute 'errors' to an array")
		}
	}

	if !ok {
		return fmt.Errorf("cannot cast attribute 'errors' to an array")
	}
	errorsStrArr := make([]string, len(errorsAnyArr))
	for i, errorAny := range errorsAnyArr {
		errorStr, ok := errorAny.(string)
		if !ok {
			return fmt.Errorf("cannot cast error to string")
		}
		errorsStrArr[i] = errorStr
	}

	errorsArr := make([]ServerError, len(errorsStrArr))
	for i, errorStr := range errorsStrArr {
		errorsArr[i] = ServerError{}
		err = json.Unmarshal([]byte(errorStr), &errorsArr[i])
		if err != nil {
			return fmt.Errorf("failed to unmarshal error string: %w", err)
		}
	}
	r.Errors = errorsArr

	metaString, ok := res["meta"]
	if !ok {
		return fmt.Errorf("missing attribute 'meta'")
	}
	if metaString == nil {

	} else {
		metaStringCast, ok := metaString.(string)
		if !ok {
			return fmt.Errorf("'meta' attribute cannot be casted to string")
		}
		meta := QueryMeta{}
		err = json.Unmarshal([]byte(metaStringCast), &meta)
		if err != nil {
			return fmt.Errorf("failed to unmarshal query 'meta' attribute: %w", err)
		}
		r.Meta = &meta
	}

	return nil
}
