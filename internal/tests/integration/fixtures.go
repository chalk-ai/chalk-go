package integration

import (
	"sync"
	"testing"

	chalk "github.com/chalk-ai/chalk-go"
	"github.com/stretchr/testify/assert"
)

var initOnce sync.Once

func newRestClient(t *testing.T) chalk.Client {
	initOnce.Do(func() {
		assert.NoError(t, chalk.InitFeatures(&testFeatures))
	})
	client, err := chalk.NewClient(t.Context())
	assert.NoError(t, err)
	return client
}

func newGRPCClient(t *testing.T) chalk.GRPCClient {
	initOnce.Do(func() {
		assert.NoError(t, chalk.InitFeatures(&testFeatures))
	})
	client, err := chalk.NewGRPCClient(t.Context())
	assert.NoError(t, err)
	return client
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
