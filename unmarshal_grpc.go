package chalk

import (
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
)

type FeatureResultGRPC struct {
	Field string
	Value any
	Meta  *FeatureMeta
}

type FeatureMeta struct {
	Pkey        any
	ResolverFqn string
	SourceType  string
	SourceId    string
}

func GetRow(response *commonv1.OnlineQueryBulkResponse, rowIndex int) ([]FeatureResultGRPC, error) {
	if len(response.GetScalarsData()) == 0 {
		return nil, errors.New("results table empty, either the query has errors or the data is malformed")
	}

	scalarsTable, err := internal.ConvertBytesToTable(response.GetScalarsData())
	if err != nil {
		return nil, errors.Wrap(err, "converting scalars data to table")
	}

	results := make([]FeatureResultGRPC, 0, scalarsTable.NumCols())
	// Need to obtain time.Time values as string because structpb.NewValue does not support time.Time.
	rows, meta, err := internal.ExtractFeaturesFromTable(scalarsTable, true)
	if err != nil {
		return nil, errors.Wrap(err, "extracting features from scalars table")
	}

	var rowMeta map[string]internal.FeatureMeta
	if len(meta) != len(rows) {
		return nil, errors.New("metadata length does not match rows length")
	}
	rowMeta = meta[rowIndex]
	for fqn, value := range rows[rowIndex] {
		featureRes := FeatureResultGRPC{
			Field: fqn,
			Value: value,
		}
		if rowMeta != nil {
			internalMeta, ok := rowMeta[fqn]
			if !ok {
				// Features such as has-many features do not have a metadata column.
				continue
			}
			publicMeta := &FeatureMeta{
				Pkey: internalMeta.Pkey,
			}
			if internalMeta.ResolverFqn != nil {
				publicMeta.ResolverFqn = *internalMeta.ResolverFqn
			}
			if internalMeta.SourceType != nil {
				publicMeta.SourceType = *internalMeta.SourceType
			}
			if internalMeta.SourceId != nil {
				publicMeta.SourceId = *internalMeta.SourceId
			}
		}
		results = append(results, featureRes)
	}
	return results, nil
}

func UnmarshalOnlineQueryResponse(response *commonv1.OnlineQueryResponse, resultHolder any) error {
	if err := validateOnlineQueryResultHolder(resultHolder); err != nil {
		return err
	}
	fqnToValue := map[Fqn]any{}
	for _, featureResult := range response.GetData().GetResults() {
		convertedValue, err := convertIfHasManyMap(featureResult.Value.AsInterface())
		if err != nil {
			return errors.Wrapf(err, "converting has-many value for feature '%s'", featureResult.Field)
		}
		fqnToValue[featureResult.Field] = convertedValue
	}
	return UnmarshalInto(resultHolder, fqnToValue, nil)
}

func UnmarshalOnlineQueryBulkResponse(response *commonv1.OnlineQueryBulkResponse, resultHolders any) error {
	scalars, err := internal.ConvertBytesToTable(response.GetScalarsData())
	if err != nil {
		return errors.Wrap(err, "deserializing scalars table")
	}
	return unmarshalTableInto(scalars, resultHolders)
}
