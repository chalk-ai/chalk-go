package chalk

import (
	assert "github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type goAddress struct {
	Id   *string
	City *string
}

type goCard struct {
	Id     *string
	Number *int
}

type goUser struct {
	Id           *string
	Name         *string
	Card         *goCard
	Address      *goAddress
	FamilySize   *int
	HasFamily    *bool
	FamilyIncome *float32
}

var testFeatures struct {
	User    *goUser
	Card    *goCard
	Address *goAddress
}

func TestInitFeaturesToNil(t *testing.T) {
	initFeatures(reflect.ValueOf(&testFeatures).Elem(), "", make(map[string]bool), make(fqnToField))
	assert.Equal(t, (*string)(nil), testFeatures.User.Id)
	assert.Equal(t, (*string)(nil), testFeatures.User.Name)
	assert.Equal(t, (*string)(nil), testFeatures.User.Card.Id)
	assert.Equal(t, (*int)(nil), testFeatures.User.Card.Number)
	assert.Equal(t, (*string)(nil), testFeatures.User.Address.Id)
	assert.Equal(t, (*string)(nil), testFeatures.User.Address.City)
	assert.Equal(t, (*int)(nil), testFeatures.User.FamilySize)
	assert.Equal(t, (*bool)(nil), testFeatures.User.HasFamily)
	assert.Equal(t, (*float32)(nil), testFeatures.User.FamilyIncome)

	assert.Equal(t, (*string)(nil), testFeatures.Card.Id)
	assert.Equal(t, (*int)(nil), testFeatures.Card.Number)

	assert.Equal(t, (*string)(nil), testFeatures.Address.Id)
	assert.Equal(t, (*string)(nil), testFeatures.Address.City)
}
