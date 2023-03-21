package chalk_test

import (
	"github.com/chalk-ai/chalk-go"
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
	assert.Equal(t, "user.id", chalk.UnwrapFeature(testFeatures.User.Id).Fqn)
	assert.Equal(t, "user.name", chalk.UnwrapFeature(testFeatures.User.Name).Fqn)
	assert.Equal(t, "user.card.id", chalk.UnwrapFeature(testFeatures.User.Card.Id).Fqn)
	assert.Equal(t, "user.card.number", chalk.UnwrapFeature(testFeatures.User.Card.Number).Fqn)
	assert.Equal(t, "user.address.id", chalk.UnwrapFeature(testFeatures.User.Address.Id).Fqn)
	assert.Equal(t, "user.address.city", chalk.UnwrapFeature(testFeatures.User.Address.City).Fqn)
	assert.Equal(t, "user.family_size", chalk.UnwrapFeature(testFeatures.User.FamilySize).Fqn)
	assert.Equal(t, "user.has_family", chalk.UnwrapFeature(testFeatures.User.HasFamily).Fqn)
	assert.Equal(t, "user.family_income", chalk.UnwrapFeature(testFeatures.User.FamilyIncome).Fqn)

	assert.Equal(t, "card.id", chalk.UnwrapFeature(testFeatures.Card.Id).Fqn)
	assert.Equal(t, "card.number", chalk.UnwrapFeature(testFeatures.Card.Number).Fqn)

	assert.Equal(t, "address.id", chalk.UnwrapFeature(testFeatures.Address.Id).Fqn)
	assert.Equal(t, "address.city", chalk.UnwrapFeature(testFeatures.Address.City).Fqn)
}
