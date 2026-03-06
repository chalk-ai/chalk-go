package chalk

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/memory"
)

type TimeBound struct {
	Time  *time.Time
	Delta *float64
}

var timeBoundFormats = []string{
	time.RFC3339Nano,
	time.RFC3339,
	"2006-01-02T15:04:05.999999999",
	"2006-01-02T15:04:05",
}

func (t *TimeBound) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	if len(data) > 0 && data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		for _, layout := range timeBoundFormats {
			if ts, err := time.Parse(layout, s); err == nil {
				t.Time = &ts
				return nil
			}
		}
		return &time.ParseError{Value: s, Message: ": unrecognized timestamp format"}
	}
	var delta float64
	if err := json.Unmarshal(data, &delta); err != nil {
		return err
	}
	t.Delta = &delta
	return nil
}

func (t TimeBound) MarshalJSON() ([]byte, error) {
	if t.Delta != nil {
		return json.Marshal(*t.Delta)
	}
	if t.Time != nil {
		return json.Marshal(*t.Time)
	}
	return []byte("null"), nil
}

type onlineQueryResponseSerialized struct {
	Data   []featureResultSerialized `json:"data"`
	Errors []chalkErrorSerialized    `json:"errors"`
	Meta   *QueryMeta                `json:"meta"`
}

type onlineQueryResultFeather struct {
	HasData    bool
	ScalarData arrow.Table
	GroupsData map[Fqn]arrow.Table
	Errors     ServerErrors
	Meta       *QueryMeta
}

type QueryName = string
type Fqn = string

type OnlineQueryBulkResponse struct {
	QueryResults map[QueryName]onlineQueryResultFeather

	allocator memory.Allocator
}

type featureResultSerialized struct {
	Field     string                 `json:"field"`
	Value     any                    `json:"value"`
	Pkey      any                    `json:"pkey"`
	Timestamp string                 `json:"ts"`
	Meta      *FeatureResolutionMeta `json:"meta"`
	Error     *chalkErrorSerialized  `json:"error"`
}

type DatasetSampleFilter struct {
	LowerBound *TimeBound `json:"lower_bound"`
	UpperBound *TimeBound `json:"upper_bound"`
	MaxSamples *int       `json:"max_samples"`
}

type DatasetFilter struct {
	SampleFilters DatasetSampleFilter `json:"sample_filters"`
	MaxCacheAge   *float64            `json:"max_cache_age_secs"`
}

type chalkHttpException struct {
	Detail any     `json:"detail"`
	Trace  *string `json:"trace"`
}

type sendRequestParams struct {
	Body                  any
	Method                string
	URL                   string
	Response              any
	ResponseHeaders       *http.Header
	Versioned             bool
	Branch                *string
	ResourceGroupOverride *string
	IsEngineRequest       bool
}

type chalkErrorSerialized struct {
	Code      string             `json:"code"`
	Category  string             `json:"category"`
	Message   string             `json:"message"`
	Exception *ResolverException `json:"exception"`
	Feature   string             `json:"feature"`
	Resolver  string             `json:"resolver"`
}

func UnmarshalDatasetResponse(data []byte) (Dataset, error) {
	var d Dataset
	err := json.Unmarshal(data, &d)
	return d, err
}
