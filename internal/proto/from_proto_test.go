package proto

import (
	"fmt"
	arrowv1 "github.com/chalk-ai/chalk-go/gen/chalk/arrow/v1"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"
)

func loadLocationOrPanic(timezone string) *time.Location {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		panic(fmt.Sprintf("failed to load location %s: %v", timezone, err))
	}
	return location
}

func convertDecimalToProtoOrPanic(value *big.Float) *arrowv1.DecimalValue {
	decimalValue, err := convertDecimalToProto(value)
	if err != nil {
		panic(err)
	}
	return decimalValue
}

func TestConvertScalarValue(t *testing.T) {
	for _, fixture := range []struct {
		name     string
		value    *arrowv1.ScalarValue
		expected any
	}{
		{
			name: "bool",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_BoolValue{
					BoolValue: true,
				},
			},
			expected: true,
		},
		{
			name: "float64",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Float64Value{
					Float64Value: 1.0,
				},
			},
			expected: 1.0,
		},
		{
			name: "float32",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Float32Value{
					Float32Value: 1.0,
				},
			},
			expected: float32(1.0),
		},
		{
			name: "float16",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Float16Value{
					Float16Value: 1.0,
				},
			},
			expected: float32(1.0),
		},
		{
			name: "int64",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Int64Value{
					Int64Value: 1,
				},
			},
			expected: int64(1),
		},
		{
			name: "int32",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Int32Value{
					Int32Value: 1,
				},
			},
			expected: int32(1),
		},
		{
			name: "int16",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Int16Value{
					Int16Value: 1,
				},
			},
			expected: int16(1),
		},
		{
			name: "int8",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Int8Value{
					Int8Value: 1,
				},
			},
			expected: int8(1),
		},
		{
			name: "uint64",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Uint64Value{
					Uint64Value: 1,
				},
			},
			expected: uint64(1),
		},
		{
			name: "uint32",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Uint32Value{
					Uint32Value: 1,
				},
			},
			expected: uint32(1),
		},
		{
			name: "uint16",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Uint16Value{
					Uint16Value: 1,
				},
			},
			expected: uint16(1),
		},
		{
			name: "uint8",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Uint8Value{
					Uint8Value: 1,
				},
			},
			expected: uint8(1),
		},
		{
			name: "utf8",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_LargeUtf8Value{
					LargeUtf8Value: "foo",
				},
			},
			expected: "foo",
		},
		{
			name: "timestamp",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_TimestampValue{
					TimestampValue: &arrowv1.ScalarTimestampValue{
						Value: &arrowv1.ScalarTimestampValue_TimeSecondValue{
							TimeSecondValue: 1577836801,
						},
						Timezone: "",
					},
				},
			},
			expected: time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
		},
		{
			name: "timestamp[ms,no-tz]",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_TimestampValue{
					TimestampValue: &arrowv1.ScalarTimestampValue{
						Value: &arrowv1.ScalarTimestampValue_TimeMillisecondValue{
							TimeMillisecondValue: 1577836800_001,
						},
						Timezone: "",
					},
				},
			},
			expected: time.Date(2020, 1, 1, 0, 0, 0, 1_000_000, time.UTC),
		},
		{
			name: "timestamp[us,no-tz]",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_TimestampValue{
					TimestampValue: &arrowv1.ScalarTimestampValue{
						Value: &arrowv1.ScalarTimestampValue_TimeMicrosecondValue{
							TimeMicrosecondValue: 1577836800_000_001,
						},
						Timezone: "",
					},
				},
			},
			expected: time.Date(2020, 1, 1, 0, 0, 0, 1000, time.UTC),
		},
		{
			name: "timestamp[ns,no-tz]",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_TimestampValue{
					TimestampValue: &arrowv1.ScalarTimestampValue{
						Value: &arrowv1.ScalarTimestampValue_TimeNanosecondValue{
							TimeNanosecondValue: 1577836800_000_000_001,
						},
						Timezone: "",
					},
				},
			},
			expected: time.Date(2020, 1, 1, 0, 0, 0, 1, time.UTC),
		},
		{
			name: "timestamp[ns,new-york]",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_TimestampValue{
					TimestampValue: &arrowv1.ScalarTimestampValue{
						Value: &arrowv1.ScalarTimestampValue_TimeSecondValue{
							TimeSecondValue: 1577836800,
						},
						Timezone: loadLocationOrPanic("America/New_York").String(),
					},
				},
			},
			expected: time.Date(2020, 1, 1, 0, 0, 0, 0, loadLocationOrPanic("America/New_York")),
		},
		{
			name: "date64",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Date_64Value{
					Date_64Value: 18262,
				},
			},
			expected: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "date32",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Date_32Value{
					Date_32Value: 18262,
				},
			},
			expected: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "time64[ns]",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Time64Value{
					Time64Value: &arrowv1.ScalarTime64Value{
						Value: &arrowv1.ScalarTime64Value_Time64NanosecondValue{
							Time64NanosecondValue: 58800_000_000_001,
						},
					},
				},
			},
			expected: time.Date(0, 1, 1, 16, 20, 0, 1, time.UTC),
		},
		{
			name: "time64[us]",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Time64Value{
					Time64Value: &arrowv1.ScalarTime64Value{
						Value: &arrowv1.ScalarTime64Value_Time64MicrosecondValue{
							Time64MicrosecondValue: 58800_000_001,
						},
					},
				},
			},
			expected: time.Date(0, 1, 1, 16, 20, 0, 1_000, time.UTC),
		},
		{
			name: "time64[ms]",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Time32Value{
					Time32Value: &arrowv1.ScalarTime32Value{
						Value: &arrowv1.ScalarTime32Value_Time32MillisecondValue{
							Time32MillisecondValue: 58800_001,
						},
					},
				},
			},
			expected: time.Date(0, 1, 1, 16, 20, 0, 1_000_000, time.UTC),
		},
		{
			name: "time32[s]",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Time32Value{
					Time32Value: &arrowv1.ScalarTime32Value{
						Value: &arrowv1.ScalarTime32Value_Time32SecondValue{
							Time32SecondValue: 58801,
						},
					},
				},
			},
			expected: time.Date(0, 1, 1, 16, 20, 1, 0, time.UTC),
		},
		{
			name: "duration[ns]",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_DurationNanosecondValue{
					DurationNanosecondValue: 58800_000_000_001,
				},
			},
			expected: time.Nanosecond * 58800_000_000_001,
		},
		{
			name: "duration[us]",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_DurationMicrosecondValue{
					DurationMicrosecondValue: 58800_000_001,
				},
			},
			expected: time.Microsecond * 58800_000_001,
		},
		{
			name: "duration[ms]",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_DurationMillisecondValue{
					DurationMillisecondValue: 58800_001,
				},
			},
			expected: time.Millisecond * 58800_001,
		},
		{
			name: "duration[s]",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_DurationSecondValue{
					DurationSecondValue: 58800,
				},
			},
			expected: time.Second * 58800,
		},
		{
			name: "binary",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_BinaryValue{
					BinaryValue: []byte{1, 2, 3},
				},
			},
			expected: []byte{1, 2, 3},
		},
		{
			name: "large-binary",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_LargeBinaryValue{
					LargeBinaryValue: []byte{1, 2, 3},
				},
			},
			expected: []byte{1, 2, 3},
		},
		{
			name: "fixed-size-binary",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_FixedSizeBinaryValue{
					FixedSizeBinaryValue: &arrowv1.ScalarFixedSizeBinary{
						Values: []byte{1, 2, 3},
						Length: 3,
					},
				},
			},
			expected: []byte{1, 2, 3},
		},
		{
			name: "decimal256",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Decimal256Value{
					Decimal256Value: convertDecimalToProtoOrPanic(big.NewFloat(123.45)),
				},
			},
			expected: big.NewFloat(123.45),
		},
		{
			name: "decimal128",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Decimal256Value{
					Decimal256Value: convertDecimalToProtoOrPanic(big.NewFloat(12345)),
				},
			},
			expected: big.NewFloat(12345),
		},
		{
			name: "decimal256",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Decimal256Value{
					Decimal256Value: convertDecimalToProtoOrPanic(big.NewFloat(0.12345)),
				},
			},
			expected: big.NewFloat(0.12345),
		},
		{
			name: "decimal256",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Decimal256Value{
					Decimal256Value: convertDecimalToProtoOrPanic(big.NewFloat(-0.0010)),
				},
			},
			expected: big.NewFloat(-0.0010),
		},
		{
			name: "decimal128",
			value: &arrowv1.ScalarValue{
				Value: &arrowv1.ScalarValue_Decimal128Value{
					Decimal128Value: convertDecimalToProtoOrPanic(big.NewFloat(123.45)),
				},
			},
			expected: big.NewFloat(123.45),
		},
	} {
		t.Run(fixture.name, func(t *testing.T) {
			t.Parallel()
			actual, err := ConvertScalarValue(fixture.value)
			assert.NoError(t, err)
			assert.Equal(t, fixture.expected, actual)
		})
	}

}
