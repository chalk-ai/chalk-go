package fixtures

import "time"

type LatLng struct {
	Lat *float64 `dataclass_field:"true"`
	Lng *float64 `dataclass_field:"true"`
}

type LatLngWithExtraField struct {
	Lat   *float64 `dataclass_field:"true"`
	Lng   *float64 `dataclass_field:"true"`
	Extra *string  `dataclass_field:"true"`
}

type FavoriteThings struct {
	Numbers *[]int64 `dataclass_field:"true"`
	Words   *[]string
}

type Possessions struct {
	Car   *string `dataclass_field:"true"`
	Yacht *string
	Plane *string
}

type Grandparent struct {
	Name *string `dataclass_field:"true"`
}

type Parent struct {
	Name *string `dataclass_field:"true"`
	Mom  *Grandparent
	Dad  *Grandparent
}

type Child struct {
	Name *string `dataclass_field:"true"`
	Mom  *Parent
	Dad  *Parent
}

type LevelOneNest struct {
	Id                *string
	ShouldAlwaysBeNil *string
	Nested            *LevelTwoNest
}

type LevelTwoNest struct {
	Id                *string
	ShouldAlwaysBeNil *string
}

type DclassWithOverrides struct {
	CamelName *string `dataclass_field:"true" name:"camelName"`
}

type HasMany struct {
	Id *string

	// The following should be kept in parity with the enumeration of
	// fields with all types in the `AllTypes` struct.
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
	Dataclass              *LatLng             `dataclass:"true"`
	DataclassList          *[]LatLng
	DataclassWithList      *FavoriteThings
	DataclassWithNils      *Possessions
	DataclassWithDataclass *Child
	DataclassWithOverrides *DclassWithOverrides
}

type AllTypes struct {
	Int                     *int64
	Float                   *float64
	String                  *string
	Bool                    *bool
	Timestamp               *time.Time
	IntList                 *[]int64
	NestedIntPointerList    *[]*[]int64
	NestedIntList           *[][]int64
	WindowedInt             map[string]*int64   `windows:"1m,5m,1h"`
	WindowedList            map[string]*[]int64 `windows:"1m"`
	Dataclass               *LatLng             `dataclass:"true"`
	DataclassWithExtraField *LatLng
	DataclassList           *[]LatLng
	DataclassWithList       *FavoriteThings
	DataclassWithNils       *Possessions
	DataclassWithDataclass  *Child
	DataclassWithOverrides  *DclassWithOverrides
	Nested                  *LevelOneNest
	HasMany                 *[]HasMany
}

var Root struct {
	AllTypes *AllTypes
}
