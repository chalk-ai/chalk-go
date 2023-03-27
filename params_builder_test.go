package chalk

import "time"

type testLatLng struct {
	Lat *float64 `dataclass_field:"true"`
	Lng *float64 `dataclass_field:"true"`
}

type anotherFeature struct {
	Id *string
}

type allTypes struct {
	Int          *int64
	Float        *float64
	String       *string
	Bool         *bool
	Timestamp    *time.Time
	IntList      *[]int64
	WindowedInt  map[string]*int64   `windows:"1m,5m,1h"`
	WindowedList map[string]*[]int64 `windows:"1m"`
	Dataclass    *testLatLng         `dataclass:"true"`
	Nested       *anotherFeature
}

// ptf is short for "params test features"
var ptf struct {
	AllTypes *allTypes
}
