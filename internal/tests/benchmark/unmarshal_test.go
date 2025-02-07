package benchmark

import (
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal/tests/benchmark/fixtures"
	"testing"
)

func BenchmarkMultiNsDeserFromSingleQueryResult(t *testing.B) {
	data := []chalk.FeatureResult{
		{
			Field: "int_features.int1",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int2",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int3",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int4",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int5",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int6",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int7",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int8",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int9",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int10",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int11",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int12",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int13",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int14",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int15",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int16",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int17",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int18",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int19",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int20",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int21",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int22",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int23",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int24",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int25",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int26",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int27",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int28",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int29",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int30",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int31",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int32",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int33",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int34",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int35",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int36",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int37",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int38",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int39",
			Value: float64(122.0),
		},
		{
			Field: "int_features.int40",
			Value: float64(122.0),
		},
		{
			Field: "float_features.float1",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float2",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float3",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float4",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float5",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float6",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float7",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float8",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float9",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float10",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float11",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float12",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float13",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float14",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float15",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float16",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float17",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float18",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float19",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float20",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float21",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float22",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float23",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float24",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float25",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float26",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float27",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float28",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float29",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float30",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float31",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float32",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float33",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float34",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float35",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float36",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float37",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float38",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float39",
			Value: float64(1.234),
		},
		{
			Field: "float_features.float40",
			Value: float64(1.234),
		},
		{
			Field: "bool_features.bool1",
			Value: true,
		},
		{
			Field: "bool_features.bool2",
			Value: true,
		},
		{
			Field: "bool_features.bool3",
			Value: true,
		},
		{
			Field: "bool_features.bool4",
			Value: true,
		},
		{
			Field: "bool_features.bool5",
			Value: true,
		},
		{
			Field: "bool_features.bool6",
			Value: true,
		},
		{
			Field: "bool_features.bool7",
			Value: true,
		},
		{
			Field: "bool_features.bool8",
			Value: true,
		},
		{
			Field: "bool_features.bool9",
			Value: true,
		},
		{
			Field: "bool_features.bool10",
			Value: true,
		},
		{
			Field: "bool_features.bool11",
			Value: true,
		},
		{
			Field: "bool_features.bool12",
			Value: true,
		},
		{
			Field: "bool_features.bool13",
			Value: true,
		},
		{
			Field: "bool_features.bool14",
			Value: true,
		},
		{
			Field: "bool_features.bool15",
			Value: true,
		},
		{
			Field: "bool_features.bool16",
			Value: true,
		},
		{
			Field: "bool_features.bool17",
			Value: true,
		},
		{
			Field: "bool_features.bool18",
			Value: true,
		},
		{
			Field: "bool_features.bool19",
			Value: true,
		},
		{
			Field: "bool_features.bool20",
			Value: true,
		},
		{
			Field: "bool_features.bool21",
			Value: true,
		},
		{
			Field: "bool_features.bool22",
			Value: true,
		},
		{
			Field: "bool_features.bool23",
			Value: true,
		},
		{
			Field: "bool_features.bool24",
			Value: true,
		},
		{
			Field: "bool_features.bool25",
			Value: true,
		},
		{
			Field: "bool_features.bool26",
			Value: true,
		},
		{
			Field: "bool_features.bool27",
			Value: true,
		},
		{
			Field: "bool_features.bool28",
			Value: true,
		},
		{
			Field: "bool_features.bool29",
			Value: true,
		},
		{
			Field: "bool_features.bool30",
			Value: true,
		},
		{
			Field: "bool_features.bool31",
			Value: true,
		},
		{
			Field: "bool_features.bool32",
			Value: true,
		},
		{
			Field: "bool_features.bool33",
			Value: true,
		},
		{
			Field: "bool_features.bool34",
			Value: true,
		},
		{
			Field: "bool_features.bool35",
			Value: true,
		},
		{
			Field: "bool_features.bool36",
			Value: true,
		},
		{
			Field: "bool_features.bool37",
			Value: true,
		},
		{
			Field: "bool_features.bool38",
			Value: true,
		},
		{
			Field: "bool_features.bool39",
			Value: true,
		},
		{
			Field: "bool_features.bool40",
			Value: true,
		},
		{
			Field: "string_features.string1",
			Value: "string_val",
		},
		{
			Field: "string_features.string2",
			Value: "string_val",
		},
		{
			Field: "string_features.string3",
			Value: "string_val",
		},
		{
			Field: "string_features.string4",
			Value: "string_val",
		},
		{
			Field: "string_features.string5",
			Value: "string_val",
		},
		{
			Field: "string_features.string6",
			Value: "string_val",
		},
		{
			Field: "string_features.string7",
			Value: "string_val",
		},
		{
			Field: "string_features.string8",
			Value: "string_val",
		},
		{
			Field: "string_features.string9",
			Value: "string_val",
		},
		{
			Field: "string_features.string10",
			Value: "string_val",
		},
		{
			Field: "string_features.string11",
			Value: "string_val",
		},
		{
			Field: "string_features.string12",
			Value: "string_val",
		},
		{
			Field: "string_features.string13",
			Value: "string_val",
		},
		{
			Field: "string_features.string14",
			Value: "string_val",
		},
		{
			Field: "string_features.string15",
			Value: "string_val",
		},
		{
			Field: "string_features.string16",
			Value: "string_val",
		},
		{
			Field: "string_features.string17",
			Value: "string_val",
		},
		{
			Field: "string_features.string18",
			Value: "string_val",
		},
		{
			Field: "string_features.string19",
			Value: "string_val",
		},
		{
			Field: "string_features.string20",
			Value: "string_val",
		},
		{
			Field: "string_features.string21",
			Value: "string_val",
		},
		{
			Field: "string_features.string22",
			Value: "string_val",
		},
		{
			Field: "string_features.string23",
			Value: "string_val",
		},
		{
			Field: "string_features.string24",
			Value: "string_val",
		},
		{
			Field: "string_features.string25",
			Value: "string_val",
		},
		{
			Field: "string_features.string26",
			Value: "string_val",
		},
		{
			Field: "string_features.string27",
			Value: "string_val",
		},
		{
			Field: "string_features.string28",
			Value: "string_val",
		},
		{
			Field: "string_features.string29",
			Value: "string_val",
		},
		{
			Field: "string_features.string30",
			Value: "string_val",
		},
		{
			Field: "string_features.string31",
			Value: "string_val",
		},
		{
			Field: "string_features.string32",
			Value: "string_val",
		},
		{
			Field: "string_features.string33",
			Value: "string_val",
		},
		{
			Field: "string_features.string34",
			Value: "string_val",
		},
		{
			Field: "string_features.string35",
			Value: "string_val",
		},
		{
			Field: "string_features.string36",
			Value: "string_val",
		},
		{
			Field: "string_features.string37",
			Value: "string_val",
		},
		{
			Field: "string_features.string38",
			Value: "string_val",
		},
		{
			Field: "string_features.string39",
			Value: "string_val",
		},
		{
			Field: "string_features.string40",
			Value: "string_val",
		},
		{
			Field: "timestamp_features.timestamp1",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp2",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp3",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp4",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp5",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp6",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp7",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp8",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp9",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp10",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp11",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp12",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp13",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp14",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp15",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp16",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp17",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp18",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp19",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp20",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp21",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp22",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp23",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp24",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp25",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp26",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp27",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp28",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp29",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp30",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp31",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp32",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp33",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp34",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp35",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp36",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp37",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp38",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp39",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "timestamp_features.timestamp40",
			Value: "2024-05-09T22:29:00Z",
		},
	}
	res := chalk.OnlineQueryResult{
		Data: data,
	}
	for i := 0; i < t.N; i++ {
		res.UnmarshalInto(&fixtures.IntFeatures{})
		res.UnmarshalInto(&fixtures.FloatFeatures{})
		res.UnmarshalInto(&fixtures.BoolFeatures{})
		res.UnmarshalInto(&fixtures.StringFeatures{})
		res.UnmarshalInto(&fixtures.TimestampFeatures{})
	}

}

func TestBenchmarkMultiNsDeserFromMultiQueryResult(t *testing.T) {

}
