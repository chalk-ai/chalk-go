package proto

import (
	arrowv1 "github.com/chalk-ai/chalk-go/gen/chalk/arrow/v1"
	"github.com/cockroachdb/errors"
	"time"
)

/*


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
			expected: time.Date(2020, 1, 1, 0, 0, 1, 0, nil),
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
			expected: time.Date(2020, 1, 1, 0, 0, 0, 1_000_000, nil),
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
			expected: time.Date(2020, 1, 1, 0, 0, 0, 1000, nil),
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
			expected: time.Date(2020, 1, 1, 0, 0, 0, 1, nil),
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
			expected: time.Date(0, 1, 1, 16, 20, 0, 1, nil),
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
			expected: time.Date(0, 1, 1, 16, 20, 0, 1_000, nil),
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
			expected: time.Date(0, 1, 1, 16, 20, 0, 1_000_000, nil),
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
			expected: time.Date(0, 1, 1, 16, 20, 1, 0, nil),
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
*/

func ConvertScalarValue(value *arrowv1.ScalarValue) (any, error) {
	if value.GetValue() == nil {
		return nil, errors.New("scalar value `Value` field is nil")
	}

	switch v := value.GetValue().(type) {
	case *arrowv1.ScalarValue_BoolValue:
		return v.BoolValue, nil
	case *arrowv1.ScalarValue_Float64Value:
		return v.Float64Value, nil
	case *arrowv1.ScalarValue_Float32Value:
		return v.Float32Value, nil
	case *arrowv1.ScalarValue_Float16Value:
		return v.Float16Value, nil
	case *arrowv1.ScalarValue_Int64Value:
		return v.Int64Value, nil
	case *arrowv1.ScalarValue_Int32Value:
		return v.Int32Value, nil
	case *arrowv1.ScalarValue_Int16Value:
		return v.Int16Value, nil
	case *arrowv1.ScalarValue_Int8Value:
		return v.Int8Value, nil
	case *arrowv1.ScalarValue_Uint64Value:
		return v.Uint64Value, nil
	case *arrowv1.ScalarValue_Uint32Value:
		return v.Uint32Value, nil
	case *arrowv1.ScalarValue_Uint16Value:
		return v.Uint16Value, nil
	case *arrowv1.ScalarValue_Uint8Value:
		return v.Uint8Value, nil
	case *arrowv1.ScalarValue_LargeUtf8Value:
		return v.LargeUtf8Value, nil
	case *arrowv1.ScalarValue_TimestampValue:
		tsValue := v.TimestampValue.GetValue()
		var tz string = v.TimestampValue.GetTimezone()
		var localTime time.Time
		switch ts := tsValue.(type) {
		case *arrowv1.ScalarTimestampValue_TimeSecondValue:
			// Convert epoch seconds to time.Time
			localTime = time.Unix(ts.TimeSecondValue, 0)
		case *arrowv1.ScalarTimestampValue_TimeMillisecondValue:
			// Convert epoch milliseconds to time.Time
			localTime = time.Unix(0, ts.TimeMillisecondValue*int64(time.Millisecond))
		case *arrowv1.ScalarTimestampValue_TimeMicrosecondValue:
			// Convert epoch microseconds to time.Time
			localTime = time.Unix(0, ts.TimeMicrosecondValue*int64(time.Microsecond))
		case *arrowv1.ScalarTimestampValue_TimeNanosecondValue:
			// Convert epoch nanoseconds to time.Time
			localTime = time.Unix(0, ts.TimeNanosecondValue)
		}
		if tz != "" {
			loc, err := time.LoadLocation(tz)
			if err != nil {
				return nil, errors.Wrapf(err, "Failed to convert timezone '%s' to `time.Location`", tz)
			}
			return localTime.In(loc), nil
		}
		return localTime.In(time.UTC), nil
	case *arrowv1.ScalarValue_Date_64Value:
		// Convert date64 to time.Time
		epoch := time.Unix(0, 0)
		return epoch.AddDate(0, 0, int(v.Date_64Value)).In(time.UTC), nil
	case *arrowv1.ScalarValue_Date_32Value:
		// Convert date32 to time.Time
		epoch := time.Unix(0, 0)
		return epoch.AddDate(0, 0, int(v.Date_32Value)).In(time.UTC), nil
	case *arrowv1.ScalarValue_Time64Value:
		tsValue := v.Time64Value.GetValue()
		var datelessTime time.Time
		switch ts := tsValue.(type) {
		case *arrowv1.ScalarTime64Value_Time64NanosecondValue:
			// Convert time64 nanoseconds to time.Time
			datelessTime = time.Unix(0, ts.Time64NanosecondValue)
		case *arrowv1.ScalarTime64Value_Time64MicrosecondValue:
			// Convert time64 microseconds to time.Time
			datelessTime = time.Unix(0, ts.Time64MicrosecondValue*int64(time.Microsecond))
		}
		return datelessTime.In(time.UTC), nil
	case *arrowv1.ScalarValue_Time32Value:
		tsValue := v.Time32Value.GetValue()
		var datelessTime time.Time
		switch ts := tsValue.(type) {
		case *arrowv1.ScalarTime32Value_Time32MillisecondValue:
			// Convert time32 milliseconds to time.Time
			datelessTime = time.Unix(0, int64(ts.Time32MillisecondValue)*int64(time.Millisecond))
		case *arrowv1.ScalarTime32Value_Time32SecondValue:
			// Convert time32 seconds to time.Time
			datelessTime = time.Unix(int64(ts.Time32SecondValue), 0)
		}
		return datelessTime.In(time.UTC), nil
	case *arrowv1.ScalarValue_DurationNanosecondValue:
		// Convert duration nanoseconds to time.Duration
		return time.Duration(v.DurationNanosecondValue), nil
	case *arrowv1.ScalarValue_DurationMicrosecondValue:
		// Convert duration microseconds to time.Duration
		return time.Duration(v.DurationMicrosecondValue) * time.Microsecond, nil
	case *arrowv1.ScalarValue_DurationMillisecondValue:
		// Convert duration milliseconds to time.Duration
		return time.Duration(v.DurationMillisecondValue) * time.Millisecond, nil
	case *arrowv1.ScalarValue_DurationSecondValue:
		// Convert duration seconds to time.Duration
		return time.Duration(v.DurationSecondValue) * time.Second, nil
	case *arrowv1.ScalarValue_BinaryValue:
		return v.BinaryValue, nil
	case *arrowv1.ScalarValue_LargeBinaryValue:
		return v.LargeBinaryValue, nil
	case *arrowv1.ScalarValue_FixedSizeBinaryValue:
		return v.FixedSizeBinaryValue.Values, nil
	case *arrowv1.ScalarValue_Decimal256Value:
		return nil, errors.New("unimplemented")
	case *arrowv1.ScalarValue_Decimal128Value:
		return nil, errors.New("unimplemented")
	}

	return nil, nil
}
