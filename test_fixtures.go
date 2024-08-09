package chalk

import "time"

type testLatLng struct {
	Lat *float64 `dataclass_field:"true" json:"lat"`
	Lng *float64 `dataclass_field:"true" json:"lng"`
}

type favoriteThings struct {
	Numbers *[]int64  `dataclass_field:"true" json:"numbers"`
	Words   *[]string `json:"words"`
}

type possessions struct {
	Car   *string `dataclass_field:"true" json:"car"`
	Yacht *string `json:"yacht"`
	Plane *string `json:"plane"`
}

type grandparent struct {
	Name *string `dataclass_field:"true" json:"name"`
}

type parent struct {
	Name *string      `dataclass_field:"true" json:"name"`
	Mom  *grandparent `json:"mom"`
	Dad  *grandparent `json:"dad"`
}

type child struct {
	Name *string `dataclass_field:"true" json:"name"`
	Mom  *parent `json:"mom"`
	Dad  *parent `json:"dad"`
}

type levelOneNest struct {
	Id                *string
	ShouldAlwaysBeNil *string
	Nested            *levelTwoNest
}

type levelTwoNest struct {
	Id                *string
	ShouldAlwaysBeNil *string
}

type dclassWithOverrides struct {
	CamelName *string `dataclass_field:"true" name:"camelName" json:"camelName"`
}

type hasMany struct {
	Id *string

	// The following should be kept in parity with the enumeration of
	// fields with all types in the `allTypes` struct.
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
	DataclassWithOverrides *dclassWithOverrides
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
	DataclassWithOverrides *dclassWithOverrides
	Nested                 *levelOneNest
	HasMany                *[]hasMany
}

var testRootFeatures struct {
	AllTypes *allTypes
}
