package fixtures

import "time"

type IntFeatures struct {
	Int1  *int64
	Int2  *int64
	Int3  *int64
	Int4  *int64
	Int5  *int64
	Int6  *int64
	Int7  *int64
	Int8  *int64
	Int9  *int64
	Int10 *int64
	Int11 *int64
	Int12 *int64
	Int13 *int64
	Int14 *int64
	Int15 *int64
	Int16 *int64
	Int17 *int64
	Int18 *int64
	Int19 *int64
	Int20 *int64
	Int21 *int64
	Int22 *int64
	Int23 *int64
	Int24 *int64
	Int25 *int64
	Int26 *int64
	Int27 *int64
	Int28 *int64
	Int29 *int64
	Int30 *int64
	Int31 *int64
	Int32 *int64
	Int33 *int64
	Int34 *int64
	Int35 *int64
	Int36 *int64
	Int37 *int64
	Int38 *int64
	Int39 *int64
	Int40 *int64
}

type FloatFeatures struct {
	Float1  *float64
	Float2  *float64
	Float3  *float64
	Float4  *float64
	Float5  *float64
	Float6  *float64
	Float7  *float64
	Float8  *float64
	Float9  *float64
	Float10 *float64
	Float11 *float64
	Float12 *float64
	Float13 *float64
	Float14 *float64
	Float15 *float64
	Float16 *float64
	Float17 *float64
	Float18 *float64
	Float19 *float64
	Float20 *float64
	Float21 *float64
	Float22 *float64
	Float23 *float64
	Float24 *float64
	Float25 *float64
	Float26 *float64
	Float27 *float64
	Float28 *float64
	Float29 *float64
	Float30 *float64
	Float31 *float64
	Float32 *float64
	Float33 *float64
	Float34 *float64
	Float35 *float64
	Float36 *float64
	Float37 *float64
	Float38 *float64
	Float39 *float64
	Float40 *float64
}

type StringFeatures struct {
	String1  *string
	String2  *string
	String3  *string
	String4  *string
	String5  *string
	String6  *string
	String7  *string
	String8  *string
	String9  *string
	String10 *string
	String11 *string
	String12 *string
	String13 *string
	String14 *string
	String15 *string
	String16 *string
	String17 *string
	String18 *string
	String19 *string
	String20 *string
	String21 *string
	String22 *string
	String23 *string
	String24 *string
	String25 *string
	String26 *string
	String27 *string
	String28 *string
	String29 *string
	String30 *string
	String31 *string
	String32 *string
	String33 *string
	String34 *string
	String35 *string
	String36 *string
	String37 *string
	String38 *string
	String39 *string
	String40 *string
}

type BoolFeatures struct {
	Bool1  *bool
	Bool2  *bool
	Bool3  *bool
	Bool4  *bool
	Bool5  *bool
	Bool6  *bool
	Bool7  *bool
	Bool8  *bool
	Bool9  *bool
	Bool10 *bool
	Bool11 *bool
	Bool12 *bool
	Bool13 *bool
	Bool14 *bool
	Bool15 *bool
	Bool16 *bool
	Bool17 *bool
	Bool18 *bool
	Bool19 *bool
	Bool20 *bool
	Bool21 *bool
	Bool22 *bool
	Bool23 *bool
	Bool24 *bool
	Bool25 *bool
	Bool26 *bool
	Bool27 *bool
	Bool28 *bool
	Bool29 *bool
	Bool30 *bool
	Bool31 *bool
	Bool32 *bool
	Bool33 *bool
	Bool34 *bool
	Bool35 *bool
	Bool36 *bool
	Bool37 *bool
	Bool38 *bool
	Bool39 *bool
	Bool40 *bool
}

type TimestampFeatures struct {
	Timestamp1  *time.Time
	Timestamp2  *time.Time
	Timestamp3  *time.Time
	Timestamp4  *time.Time
	Timestamp5  *time.Time
	Timestamp6  *time.Time
	Timestamp7  *time.Time
	Timestamp8  *time.Time
	Timestamp9  *time.Time
	Timestamp10 *time.Time
	Timestamp11 *time.Time
	Timestamp12 *time.Time
	Timestamp13 *time.Time
	Timestamp14 *time.Time
	Timestamp15 *time.Time
	Timestamp16 *time.Time
	Timestamp17 *time.Time
	Timestamp18 *time.Time
	Timestamp19 *time.Time
	Timestamp20 *time.Time
	Timestamp21 *time.Time
	Timestamp22 *time.Time
	Timestamp23 *time.Time
	Timestamp24 *time.Time
	Timestamp25 *time.Time
	Timestamp26 *time.Time
	Timestamp27 *time.Time
	Timestamp28 *time.Time
	Timestamp29 *time.Time
	Timestamp30 *time.Time
	Timestamp31 *time.Time
	Timestamp32 *time.Time
	Timestamp33 *time.Time
	Timestamp34 *time.Time
	Timestamp35 *time.Time
	Timestamp36 *time.Time
	Timestamp37 *time.Time
	Timestamp38 *time.Time
	Timestamp39 *time.Time
	Timestamp40 *time.Time
}

type WindowedIntFeatures struct {
	WindowedInt1  map[string]*int64 `windows:"1m,5m,1h"`
	WindowedInt2  map[string]*int64 `windows:"1m,5m,1h"`
	WindowedInt3  map[string]*int64 `windows:"1m,5m,1h"`
	WindowedInt4  map[string]*int64 `windows:"1m,5m,1h"`
	WindowedInt5  map[string]*int64 `windows:"1m,5m,1h"`
	WindowedInt6  map[string]*int64 `windows:"1m,5m,1h"`
	WindowedInt7  map[string]*int64 `windows:"1m,5m,1h"`
	WindowedInt8  map[string]*int64 `windows:"1m,5m,1h"`
	WindowedInt9  map[string]*int64 `windows:"1m,5m,1h"`
	WindowedInt10 map[string]*int64 `windows:"1m,5m,1h"`
	WindowedInt11 map[string]*int64 `windows:"1m,5m,1h"`
	WindowedInt12 map[string]*int64 `windows:"1m,5m,1h"`
	WindowedInt13 map[string]*int64 `windows:"1m,5m,1h"`
}

type WindowedBoolFeatures struct {
	WindowedBool1  map[string]*bool `windows:"1m,5m,1h"`
	WindowedBool2  map[string]*bool `windows:"1m,5m,1h"`
	WindowedBool3  map[string]*bool `windows:"1m,5m,1h"`
	WindowedBool4  map[string]*bool `windows:"1m,5m,1h"`
	WindowedBool5  map[string]*bool `windows:"1m,5m,1h"`
	WindowedBool6  map[string]*bool `windows:"1m,5m,1h"`
	WindowedBool7  map[string]*bool `windows:"1m,5m,1h"`
	WindowedBool8  map[string]*bool `windows:"1m,5m,1h"`
	WindowedBool9  map[string]*bool `windows:"1m,5m,1h"`
	WindowedBool10 map[string]*bool `windows:"1m,5m,1h"`
	WindowedBool11 map[string]*bool `windows:"1m,5m,1h"`
	WindowedBool12 map[string]*bool `windows:"1m,5m,1h"`
	WindowedBool13 map[string]*bool `windows:"1m,5m,1h"`
}

type WindowedFloatFeatures struct {
	WindowedFloat1  map[string]*float64 `windows:"1m,5m,1h"`
	WindowedFloat2  map[string]*float64 `windows:"1m,5m,1h"`
	WindowedFloat3  map[string]*float64 `windows:"1m,5m,1h"`
	WindowedFloat4  map[string]*float64 `windows:"1m,5m,1h"`
	WindowedFloat5  map[string]*float64 `windows:"1m,5m,1h"`
	WindowedFloat6  map[string]*float64 `windows:"1m,5m,1h"`
	WindowedFloat7  map[string]*float64 `windows:"1m,5m,1h"`
	WindowedFloat8  map[string]*float64 `windows:"1m,5m,1h"`
	WindowedFloat9  map[string]*float64 `windows:"1m,5m,1h"`
	WindowedFloat10 map[string]*float64 `windows:"1m,5m,1h"`
	WindowedFloat11 map[string]*float64 `windows:"1m,5m,1h"`
	WindowedFloat12 map[string]*float64 `windows:"1m,5m,1h"`
	WindowedFloat13 map[string]*float64 `windows:"1m,5m,1h"`
}

type WindowedStringFeatures struct {
	WindowedString1  map[string]*string `windows:"1m,5m,1h"`
	WindowedString2  map[string]*string `windows:"1m,5m,1h"`
	WindowedString3  map[string]*string `windows:"1m,5m,1h"`
	WindowedString4  map[string]*string `windows:"1m,5m,1h"`
	WindowedString5  map[string]*string `windows:"1m,5m,1h"`
	WindowedString6  map[string]*string `windows:"1m,5m,1h"`
	WindowedString7  map[string]*string `windows:"1m,5m,1h"`
	WindowedString8  map[string]*string `windows:"1m,5m,1h"`
	WindowedString9  map[string]*string `windows:"1m,5m,1h"`
	WindowedString10 map[string]*string `windows:"1m,5m,1h"`
	WindowedString11 map[string]*string `windows:"1m,5m,1h"`
	WindowedString12 map[string]*string `windows:"1m,5m,1h"`
	WindowedString13 map[string]*string `windows:"1m,5m,1h"`
}

type WindowedTimestampFeatures struct {
	WindowedTimestamp1  map[string]*time.Time `windows:"1m,5m,1h"`
	WindowedTimestamp2  map[string]*time.Time `windows:"1m,5m,1h"`
	WindowedTimestamp3  map[string]*time.Time `windows:"1m,5m,1h"`
	WindowedTimestamp4  map[string]*time.Time `windows:"1m,5m,1h"`
	WindowedTimestamp5  map[string]*time.Time `windows:"1m,5m,1h"`
	WindowedTimestamp6  map[string]*time.Time `windows:"1m,5m,1h"`
	WindowedTimestamp7  map[string]*time.Time `windows:"1m,5m,1h"`
	WindowedTimestamp8  map[string]*time.Time `windows:"1m,5m,1h"`
	WindowedTimestamp9  map[string]*time.Time `windows:"1m,5m,1h"`
	WindowedTimestamp10 map[string]*time.Time `windows:"1m,5m,1h"`
	WindowedTimestamp11 map[string]*time.Time `windows:"1m,5m,1h"`
	WindowedTimestamp12 map[string]*time.Time `windows:"1m,5m,1h"`
	WindowedTimestamp13 map[string]*time.Time `windows:"1m,5m,1h"`
}
