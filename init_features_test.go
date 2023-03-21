package chalk_test

import (
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal"
	assert "github.com/stretchr/testify/require"
	"testing"
)

type goAddress struct {
	Id   *string
	City *string
}

type goCard struct {
	Id     *string
	Number *string
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

func init() {
	chalk.InitFeatures(&testFeatures)
}

func TestInitFeatures(t *testing.T) {
	// Make sure each one can be unwrapped into a feature.
	assert.Equal(t, "user.id", internal.UnwrapFeature(testFeatures.User.Id).Fqn)
	assert.Equal(t, "user.name", internal.UnwrapFeature(testFeatures.User.Name).Fqn)
	assert.Equal(t, "user.card.id", internal.UnwrapFeature(testFeatures.User.Card.Id).Fqn)
	assert.Equal(t, "user.card.number", internal.UnwrapFeature(testFeatures.User.Card.Number).Fqn)
	assert.Equal(t, "user.address.id", internal.UnwrapFeature(testFeatures.User.Address.Id).Fqn)
	assert.Equal(t, "user.address.city", internal.UnwrapFeature(testFeatures.User.Address.City).Fqn)
	assert.Equal(t, "user.family_size", internal.UnwrapFeature(testFeatures.User.FamilySize).Fqn)
	assert.Equal(t, "user.has_family", internal.UnwrapFeature(testFeatures.User.HasFamily).Fqn)
	assert.Equal(t, "user.family_income", internal.UnwrapFeature(testFeatures.User.FamilyIncome).Fqn)

	assert.Equal(t, "card.id", internal.UnwrapFeature(testFeatures.Card.Id).Fqn)
	assert.Equal(t, "card.number", internal.UnwrapFeature(testFeatures.Card.Number).Fqn)

	assert.Equal(t, "address.id", internal.UnwrapFeature(testFeatures.Address.Id).Fqn)
	assert.Equal(t, "address.city", internal.UnwrapFeature(testFeatures.Address.City).Fqn)
}
