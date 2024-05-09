package chalk

import (
	assert "github.com/stretchr/testify/require"
	"testing"
	"time"
)

type unmarshalLatLng struct {
	Lat *float64 `dataclass_field:"true"`
	Lng *float64 `dataclass_field:"true"`
}

type unmarshalUser struct {
	Id *string

	// Versioned features
	Grade   *int `versioned:"default(2)"`
	GradeV1 *int `versioned:"true"`
	GradeV2 *int `versioned:"true"`

	// Windowed features
	AvgSpend map[string]*float64 `windows:"1m,5m,1h"`

	// Dataclass features
	LatLng *unmarshalLatLng
}

func TestUnmarshalVersionedFeatures(t *testing.T) {
	data := []FeatureResult{
		{
			Field:     "unmarshal_user.grade",
			Value:     1,
			Pkey:      "khjdsfjhdksjfh",
			Timestamp: time.Time{},
			Meta:      nil,
			Error:     nil,
		},
		{
			Field:     "unmarshal_user.grade@2",
			Value:     2,
			Pkey:      "kjhsdfkjkdjfk",
			Timestamp: time.Time{},
			Meta:      nil,
			Error:     nil,
		},
	}
	result := OnlineQueryResult{
		Data:            data,
		Meta:            nil,
		features:        nil,
		expectedOutputs: nil,
	}
	user := unmarshalUser{}
	unmarshalErr := result.UnmarshalInto(&user)
	if unmarshalErr != nil {
		t.Fatal(unmarshalErr)
	}
	assert.Nil(t, unmarshalErr)
	assert.Equal(t, 2, *user.Grade)
	assert.Equal(t, 1, *user.GradeV1)
	assert.Equal(t, 2, *user.GradeV2)
}

func TestUnmarshalWindowedFeatures(t *testing.T) {
	data := []FeatureResult{
		{
			Field:     "unmarshal_user.avg_spend__60__",
			Value:     60.0,
			Pkey:      "khjdsfjhdksjfh",
			Timestamp: time.Time{},
			Meta:      nil,
			Error:     nil,
		},
		{
			Field:     "unmarshal_user.avg_spend__300__",
			Value:     300.0,
			Pkey:      "kjhsdfkjkdjfk",
			Timestamp: time.Time{},
			Meta:      nil,
			Error:     nil,
		},
		{
			Field:     "unmarshal_user.avg_spend__3600__",
			Value:     3600.0,
			Pkey:      "kjhsdfkjkdjfk",
			Timestamp: time.Time{},
			Meta:      nil,
			Error:     nil,
		},
	}
	result := OnlineQueryResult{
		Data:            data,
		Meta:            nil,
		features:        nil,
		expectedOutputs: nil,
	}
	user := unmarshalUser{}
	unmarshalErr := result.UnmarshalInto(&user)
	if unmarshalErr != nil {
		t.Fatal(unmarshalErr)
	}
	assert.Nil(t, unmarshalErr)
	assert.Equal(t, 60.0, *user.AvgSpend["1m"])
	assert.Equal(t, 300.0, *user.AvgSpend["5m"])
	assert.Equal(t, 3600.0, *user.AvgSpend["1h"])
}

func TestUnmarshalDataclassFeatures(t *testing.T) {
	data := []FeatureResult{
		{
			Field:     "unmarshal_user.lat_lng",
			Value:     []any{1.0, 2.0},
			Pkey:      "khjdsfjhdksjfh",
			Timestamp: time.Time{},
			Meta:      nil,
			Error:     nil,
		},
	}
	result := OnlineQueryResult{
		Data:            data,
		Meta:            nil,
		features:        nil,
		expectedOutputs: nil,
	}
	user := unmarshalUser{}
	unmarshalErr := result.UnmarshalInto(&user)
	if unmarshalErr != nil {
		t.Fatal(unmarshalErr)
	}
	assert.Nil(t, unmarshalErr)
	assert.Equal(t, 1.0, *user.LatLng.Lat)
	assert.Equal(t, 2.0, *user.LatLng.Lng)
}
