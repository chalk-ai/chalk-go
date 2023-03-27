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
	initErr := initFeatures(reflect.ValueOf(&testFeatures).Elem(), "", make(map[string]bool), make(fqnToField))
	assert.Nil(t, initErr)
	assert.Nil(t, testFeatures.User.Id)
	assert.Nil(t, testFeatures.User.Name)
	assert.Nil(t, testFeatures.User.Card.Id)
	assert.Nil(t, testFeatures.User.Card.Number)
	assert.Nil(t, testFeatures.User.Address.Id)
	assert.Nil(t, testFeatures.User.Address.City)
	assert.Nil(t, testFeatures.User.FamilySize)
	assert.Nil(t, testFeatures.User.HasFamily)
	assert.Nil(t, testFeatures.User.FamilyIncome)

	assert.Nil(t, testFeatures.Card.Id)
	assert.Nil(t, testFeatures.Card.Number)

	assert.Nil(t, testFeatures.Address.Id)
	assert.Nil(t, testFeatures.Address.City)

}

func TestUnmarshal(t *testing.T) {
	userId := FeatureResult{
		Value: "1",
		Field: "go_user.id",
	}
	userName := FeatureResult{
		Value: "Jan",
		Field: "go_user.name",
	}
	userCardId := FeatureResult{
		Value: "3333-3333-3333-3333",
		Field: "go_user.card.id",
	}
	userAddressId := FeatureResult{
		Value: "2",
		Field: "go_user.address.id",
	}
	userAddressCity := FeatureResult{
		Value: "San Francisco",
		Field: "go_user.address.city",
	}
	userFamilySize := FeatureResult{
		Value: 4,
		Field: "go_user.family_size",
	}
	userHasFamily := FeatureResult{
		Value: true,
		Field: "go_user.has_family",
	}
	userFamilyIncome := FeatureResult{
		Value: float32(100000.0),
		Field: "go_user.family_income",
	}

	result := OnlineQueryResult{
		Data: []FeatureResult{
			userId,
			userName,
			userCardId,
			userAddressId,
			userAddressCity,
			userFamilySize,
			userHasFamily,
			userFamilyIncome,
		},
	}

	user := goUser{}
	err := result.UnmarshalInto(&user)
	assert.Nil(t, err)

	// In the result
	assert.Equal(t, *user.Id, userId.Value)
	assert.Equal(t, *user.Name, userName.Value)
	assert.Equal(t, *user.Card.Id, userCardId.Value)
	assert.Equal(t, *user.Address.Id, userAddressId.Value)
	assert.Equal(t, *user.Address.City, userAddressCity.Value)
	assert.Equal(t, *user.FamilySize, userFamilySize.Value)
	assert.Equal(t, *user.HasFamily, userHasFamily.Value)
	assert.Equal(t, *user.FamilyIncome, userFamilyIncome.Value)

	// Not in the result
	assert.Nil(t, user.Card.Number)
}
