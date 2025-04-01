package integration

import (
	"context"
	chalk "github.com/chalk-ai/chalk-go"
	"os"
)

var restClient chalk.Client
var grpcClient chalk.GRPCClient

func init() {
	if err := chalk.InitFeatures(&testFeatures); err != nil {
		panic(err)
	}

	if os.Getenv("INTEGRATION_TESTER") == "" {
		return
	}

	ctx := context.Background()

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

var testFeatures struct {
	AllTypes            *allTypes
	Cached              *cached
	Crashing            *crashing
	CrashingHasManyRoot *crashingHasManyRoot
	NQFeatures          *nqFeatures `name:"nq_features"`
	Optionals           *optionals
}
