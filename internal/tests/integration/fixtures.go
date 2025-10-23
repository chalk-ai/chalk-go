package integration

import (
	"context"
	chalk "github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/envfs"
	"sync"
)

var restClient chalk.Client
var grpcClient chalk.GRPCClient

type testFeaturesType struct {
	AllTypes            *allTypes
	Cached              *cached
	Crashing            *crashing
	CrashingHasManyRoot *crashingHasManyRoot
	NQFeatures          *nqFeatures `name:"nq_features"`
	Optionals           *optionals
}

var initTestFeaturesOnce sync.Once
var testFeaturesSingleton testFeaturesType
var initTestFeaturesError error

func GetTestFeatures() (testFeaturesType, error) {
	initTestFeaturesOnce.Do(func() {
		initTestFeaturesError = chalk.InitFeatures(&testFeaturesSingleton)
	})
	return testFeaturesSingleton, initTestFeaturesError
}

func init() {
	if _, err := GetTestFeatures(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	if envfs.EnvironmentGetterFromContext(ctx).Getenv("INTEGRATION_TESTER") == "" {
		return
	}

	client, err := chalk.NewClient(ctx)
	if err != nil {
		panic(err)
	}
	restClient = client

	clientGrpc, err := chalk.NewGRPCClient(ctx)
	if err != nil {
		panic(err)
	}
	grpcClient = clientGrpc
}

type allTypes struct {
	Id      *int64
	StrFeat *string
	IntFeat *int64
	HasMany *[]hasManyFeature `has_many:"id,all_types_id"`
}

type cached struct {
	Id                   *string
	RandomUploadedNumber *float64
}

type crashing struct {
	Id   *int64
	Name *string
}

type crashingHasMany struct {
	Id     *string
	Amount *float64
	RootId *string
}

type crashingHasManyRoot struct {
	Id              *string
	Name            *string
	CrashingHasMany *[]crashingHasMany
}

type hasManyFeature struct {
	Id         *string
	Name       *string
	AllTypesId *int64
}

type nqFeatures struct {
	Id   *int64
	Name *string
}

type optionals struct {
	Id   *int64
	Name *string
}
