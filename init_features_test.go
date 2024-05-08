package chalk_test

import (
	"github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"testing"
)

var initFeaturesErr error

type goLatLng struct {
	Lat *float64 `dataclass_field:"true"`
	Lng *float64 `dataclass_field:"true"`
}

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

	CustomUnderscoresL90DP90DAustralia *int `name:"custom_underscores_l90d_p90d_australia"`

	// Versioned features
	Grade   *int `versioned:"default(2)"`
	GradeV1 *int `versioned:"true"`
	GradeV2 *int `versioned:"true"`

	// Windowed features
	AvgSpend map[string]*float64 `windows:"1m,5m,1h"`

	// Dataclass features
	LatLng *goLatLng
}

var testFeatures struct {
	User    *goUser
	Card    *goCard
	Address *goAddress
}

func init() {
	initFeaturesErr = chalk.InitFeatures(&testFeatures)
}

func TestInitFeatures(t *testing.T) {
	assert.Nil(t, initFeaturesErr)

	userId, err := chalk.UnwrapFeature(testFeatures.User.Id)
	assert.Nil(t, err)
	assert.Equal(t, "user.id", userId.Fqn)

	userName, err := chalk.UnwrapFeature(testFeatures.User.Name)
	assert.Nil(t, err)
	assert.Equal(t, "user.name", userName.Fqn)

	userCardId, err := chalk.UnwrapFeature(testFeatures.User.Card.Id)
	assert.Nil(t, err)
	assert.Equal(t, "user.card.id", userCardId.Fqn)

	userCardNumber, err := chalk.UnwrapFeature(testFeatures.User.Card.Number)
	assert.Nil(t, err)
	assert.Equal(t, "user.card.number", userCardNumber.Fqn)

	userAddressId, err := chalk.UnwrapFeature(testFeatures.User.Address.Id)
	assert.Nil(t, err)
	assert.Equal(t, "user.address.id", userAddressId.Fqn)

	userAddressCity, err := chalk.UnwrapFeature(testFeatures.User.Address.City)
	assert.Nil(t, err)
	assert.Equal(t, "user.address.city", userAddressCity.Fqn)

	userFamilySize, err := chalk.UnwrapFeature(testFeatures.User.FamilySize)
	assert.Nil(t, err)
	assert.Equal(t, "user.family_size", userFamilySize.Fqn)

	userHasFamily, err := chalk.UnwrapFeature(testFeatures.User.HasFamily)
	assert.Nil(t, err)
	assert.Equal(t, "user.has_family", userHasFamily.Fqn)

	userFamilyIncome, err := chalk.UnwrapFeature(testFeatures.User.FamilyIncome)
	assert.Nil(t, err)
	assert.Equal(t, "user.family_income", userFamilyIncome.Fqn)

	cardId, err := chalk.UnwrapFeature(testFeatures.Card.Id)
	assert.Nil(t, err)
	assert.Equal(t, "card.id", cardId.Fqn)

	cardNumber, err := chalk.UnwrapFeature(testFeatures.Card.Number)
	assert.Nil(t, err)
	assert.Equal(t, "card.number", cardNumber.Fqn)

	addressId, err := chalk.UnwrapFeature(testFeatures.Address.Id)
	assert.Nil(t, err)
	assert.Equal(t, "address.id", addressId.Fqn)

	addressCity, err := chalk.UnwrapFeature(testFeatures.Address.City)
	assert.Nil(t, err)
	assert.Equal(t, "address.city", addressCity.Fqn)

	customUnderscore, err := chalk.UnwrapFeature(testFeatures.User.CustomUnderscoresL90DP90DAustralia)
	assert.Nil(t, err)
	assert.Equal(t, "user.custom_underscores_l90d_p90d_australia", customUnderscore.Fqn)

	grade, err := chalk.UnwrapFeature(testFeatures.User.Grade)
	assert.Nil(t, err)
	assert.Equal(t, "user.grade@2", grade.Fqn)

	gradeV1, err := chalk.UnwrapFeature(testFeatures.User.GradeV1)
	assert.Nil(t, err)
	assert.Equal(t, "user.grade", gradeV1.Fqn)

	gradeV2, err := chalk.UnwrapFeature(testFeatures.User.GradeV2)
	assert.Nil(t, err)
	assert.Equal(t, "user.grade@2", gradeV2.Fqn)

	avgSpend1m, err := chalk.UnwrapFeature(testFeatures.User.AvgSpend["1m"])
	assert.Nil(t, err)
	assert.Equal(t, "user.avg_spend__60__", avgSpend1m.Fqn)

	avgSpend5m, err := chalk.UnwrapFeature(testFeatures.User.AvgSpend["5m"])
	assert.Nil(t, err)
	assert.Equal(t, "user.avg_spend__300__", avgSpend5m.Fqn)

	avgSpend1h, err := chalk.UnwrapFeature(testFeatures.User.AvgSpend["1h"])
	assert.Nil(t, err)
	assert.Equal(t, "user.avg_spend__3600__", avgSpend1h.Fqn)

	latLngLat, err := chalk.UnwrapFeature(testFeatures.User.LatLng.Lat)
	assert.Nil(t, err)
	assert.Equal(t, "user.lat_lng", latLngLat.Fqn)

	latLngLng, err := chalk.UnwrapFeature(testFeatures.User.LatLng.Lng)
	assert.Nil(t, err)
	assert.Equal(t, "user.lat_lng", latLngLng.Fqn)
}
