package chalk

import "time"

type testLatLng struct {
	Lat *float64 `dataclass_field:"true"`
	Lng *float64 `dataclass_field:"true"`
}

type favoriteThings struct {
	Numbers *[]int64 `dataclass_field:"true"`
	Words   *[]string
}

type possessions struct {
	Car   *string `dataclass_field:"true"`
	Yacht *string
	Plane *string
}

type grandparent struct {
	Name *string `dataclass_field:"true"`
}

type parent struct {
	Name *string `dataclass_field:"true"`
	Mom  *grandparent
	Dad  *grandparent
}

type child struct {
	Name *string `dataclass_field:"true"`
	Mom  *parent
	Dad  *parent
}

type anotherFeature struct {
	Id *string
}

type allTypes struct {
	Int                    *int64
	Float                  *float64
	String                 *string
	Bool                   *bool
	Timestamp              *time.Time
	IntList                *[]int64
	NestedIntPointerList   *[]*[]int64
	NestedIntList          *[][]int64
	WindowedInt            map[string]*int64   `windows:"1m,5m,1h"`
	WindowedList           map[string]*[]int64 `windows:"1m"`
	Dataclass              *testLatLng         `dataclass:"true"`
	DataclassList          *[]testLatLng
	DataclassWithList      *favoriteThings
	DataclassWithNils      *possessions
	DataclassWithDataclass *child
	Nested                 *anotherFeature
}

var testRootFeatures struct {
	AllTypes *allTypes
}
