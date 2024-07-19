package integration

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"time"
)

type creditReport struct {
	Id          *string
	CreditScore *int64
}

type latLng struct {
	Lat *float64 `dataclass_field:"true"`
	Lng *float64 `dataclass_field:"true"`
}
type franchiseSet struct {
	Name            *string   `dataclass_field:"true"`
	PrimaryLocation *location `name:"primaryLocation"`
	Locations       *[]location
}

type location struct {
	Latlng      *latLng   `dataclass_field:"true"`
	Owners      *[]string `dataclass_field:"true"`
	Coordinates *[]latLng `dataclass_field:"true"`
}

type user struct {
	Id                       *int64
	FullName                 *string
	Gender                   *string
	Today                    *time.Time
	HiThere                  *int64
	TodayCached              *time.Time
	SocureScore              *float64
	CreditReportId           *string
	CreditReport             *creditReport
	FailsAlways              *int64
	FailsOnFirstTry          *int64
	CountFailedCronPods      *int64
	FavoriteColor            *string
	StsName                  *string
	SigmaFraudScore          *float64
	UnplannableFeature       *string
	LatLng                   *latLng       `dataclass:"true"`
	FranchiseSet             *franchiseSet `dataclass:"true"`
	FavoriteNumbers          *[]int64
	FavoriteColors           *[]string
	NewFavoriteNumbers       *[]int64
	NewFavoriteColors        *[]string
	NiceNewFeature           *int64
	SomethingDifferent       *string
	Newthing                 *string
	MedianTransactionAmount  *float64
	MeanTransactionAmount    *float64
	SumTransactionAmount     *float64
	CrashingFeature          *float64
	SlowFeature              *float64
	RandomNormal             *float64
	RandomUniform            *float64
	RandomDiffOnlineOffline  *float64
	RandomIncreasingOverTime *float64
	NowValue                 *time.Time
	PartiallyNoneRandom      *float64
	FullNameOptional         *string
	DeploymentId             *string
	ReversedDeploymentId     *string
	UpperDeploymentId        *string
	DefaultedFeature         *string
	VersionedFeature         *float64 `versioned:"default(2)"`
	VersionedFeatureV1       *float64 `versioned:"true"`
	VersionedFeatureV2       *float64 `versioned:"true"`
	Version1Feature          *float64 `name:"version_1_feature" versioned:"default(1)"`
	Version1FeatureV1        *float64 `name:"version_1_feature_v1" versioned:"true"`
	MaxCreditScore           *int64
	BinaryData               *[]byte
	DebugCronLogs            *string
}

var testFeatures struct {
	User *user
}

type Intercepted struct {
	Header http.Header
	Body   []byte
	URL    *url.URL
}

type InterceptorHTTPClient struct {
	Intercepted Intercepted
}

func NewInterceptorHTTPClient() *InterceptorHTTPClient {
	return &InterceptorHTTPClient{}
}

func (c *InterceptorHTTPClient) Do(req *http.Request) (*http.Response, error) {
	bodyBytes, bodyBytesErr := io.ReadAll(req.Body)
	if bodyBytesErr != nil {
		return nil, bodyBytesErr
	}
	req.Body.Close()
	c.Intercepted = Intercepted{
		Header: req.Header,
		Body:   bodyBytes,
		URL:    req.URL,
	}
	body := io.NopCloser(bytes.NewBufferString(`{"data": {"something": "exciting"}}`))
	return &http.Response{StatusCode: 200, Body: body}, nil
}

func (c *InterceptorHTTPClient) Get(url string) (*http.Response, error) {
	actualClient := &http.Client{}
	return actualClient.Get(url)
}
