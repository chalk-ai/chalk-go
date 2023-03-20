package chalk_test

import (
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal"
	assert "github.com/stretchr/testify/require"
	"testing"
)

type GoAddress struct {
	Id   *string
	City *string
}

type GoCard struct {
	Id     *string
	Number *string
}

type GoUser struct {
	Id      *string
	Name    *string
	Card    *GoCard
	Address *GoAddress
}

var TestFeatures struct {
	User    *GoUser
	Card    *GoCard
	Address *GoAddress
}

func init() {
	chalk.InitFeatures(&TestFeatures)
}

func TestInitFeatures(t *testing.T) {
	// Make sure each one can be unwrapped into a feature.
	assert.Equal(t, internal.UnwrapFeature(TestFeatures.User.Id).Fqn, "user.id")
	assert.Equal(t, internal.UnwrapFeature(TestFeatures.User.Name).Fqn, "user.name")
	assert.Equal(t, internal.UnwrapFeature(TestFeatures.User.Card.Id).Fqn, "user.card.id")
	assert.Equal(t, internal.UnwrapFeature(TestFeatures.User.Card.Number).Fqn, "user.card.number")
	assert.Equal(t, internal.UnwrapFeature(TestFeatures.User.Address.Id).Fqn, "user.address.id")
	assert.Equal(t, internal.UnwrapFeature(TestFeatures.User.Address.City).Fqn, "user.address.city")

	assert.Equal(t, internal.UnwrapFeature(TestFeatures.Card.Id).Fqn, "card.id")
	assert.Equal(t, internal.UnwrapFeature(TestFeatures.Card.Number).Fqn, "card.number")

	assert.Equal(t, internal.UnwrapFeature(TestFeatures.Address.Id).Fqn, "address.id")
	assert.Equal(t, internal.UnwrapFeature(TestFeatures.Address.City).Fqn, "address.city")
}
