package integration

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/chalk-ai/chalk-go"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/http2"
	"net/http"
	"testing"
	"time"
)

func getParams() chalk.OnlineQueryParamsComplete {
	return chalk.OnlineQueryParams{}.
		WithInput(testFeatures.User.Id, 1).
		WithOutputs(
			testFeatures.User.Id,
			testFeatures.User.Gender,
			testFeatures.User.Today,
			testFeatures.User.NiceNewFeature,
			testFeatures.User.SocureScore,
			testFeatures.User.FavoriteNumbers,
			testFeatures.User.FavoriteColors,
			testFeatures.User.FranchiseSet,
		)
}

func testUserValues(t *testing.T, testUser *user) {
	t.Helper()
	assert.NotNil(t, testUser)
	assert.NotNil(t, testUser.Id)
	assert.Equal(t, int64(1), *testUser.Id)
	assert.NotNil(t, testUser.Gender)
	assert.Equal(t, "f", *testUser.Gender)
	assert.NotNil(t, testUser.Today)
	assert.NotNil(t, testUser.NiceNewFeature)
	assert.Equal(t, int64(9), *testUser.NiceNewFeature)
	assert.NotNil(t, testUser.SocureScore)
	assert.Equal(t, 123.0, *testUser.SocureScore)
	assert.NotNil(t, testUser.FavoriteNumbers)
	assert.Equal(t, []int64{1, 2, 3}, *testUser.FavoriteNumbers)
	assert.NotNil(t, testUser.FavoriteColors)
	assert.Equal(t, []string{"red", "green", "blue"}, *testUser.FavoriteColors)
	assert.NotNil(t, testUser.FranchiseSet)
}

// TestOnlineQueryE2E mainly tests querying real data
// from the staging server does not crash. Correctness
// is partially tested here, but is mainly tested in
// TestOnlineQueryUnmarshalNonBulkAllTypes.
func TestOnlineQueryE2E(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	for _, fixture := range []struct {
		useGrpc bool
	}{
		{useGrpc: false},
		{useGrpc: true},
	} {
		t.Run(fmt.Sprintf("grpc=%v", fixture.useGrpc), func(t *testing.T) {
			certPool, err := x509.SystemCertPool()
			if err != nil {
				t.Fatal("Failed creating a system cert pool", err)
			}
			httpClient := http.Client{
				Transport: &http2.Transport{
					TLSClientConfig: &tls.Config{
						RootCAs: certPool,
					},
				},
			}

			client, err := chalk.NewClient(context.Background(), &chalk.ClientConfig{UseGrpc: fixture.useGrpc, HTTPClient: &httpClient})
			if err != nil {
				t.Fatal("Failed creating a Chalk Client", err)
			}
			err = chalk.InitFeatures(&testFeatures)
			if err != nil {
				t.Fatal("Failed initializing features", err)
			}

			var implicitUser user
			res, queryErr := client.OnlineQuery(context.Background(), getParams(), &implicitUser)
			if queryErr != nil {
				t.Fatal("Failed querying features", queryErr)
			}

			var explicitUser user
			assert.NoError(t, res.UnmarshalInto(&explicitUser))
			testUserValues(t, &implicitUser)
			testUserValues(t, &explicitUser)
		})
	}
}

// TestNamedQueriesE2E tests that querying with a query name works.
func TestNamedQueriesE2E(t *testing.T) {
	t.Skip("CHA-5086")
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	for _, fixture := range []struct {
		useGrpc bool
	}{
		{useGrpc: false},
		{useGrpc: true},
	} {
		t.Run(fmt.Sprintf("grpc=%v", fixture.useGrpc), func(t *testing.T) {
			client, err := chalk.NewClient(context.Background(), &chalk.ClientConfig{UseGrpc: fixture.useGrpc})
			if err != nil {
				t.Fatal("Failed creating a Chalk Client", err)
			}
			err = chalk.InitFeatures(&testFeatures)
			if err != nil {
				t.Fatal("Failed initializing features", err)
			}

			var implicitUser user
			params := chalk.OnlineQueryParams{}.
				WithInput("user.id", 1).
				WithQueryName("user_socure_score")

			_, queryErr := client.OnlineQuery(context.Background(), params, &implicitUser)
			if queryErr != nil {
				t.Fatal("Failed querying features", queryErr)
			}
			assert.Equal(t, 123.0, *implicitUser.SocureScore)
		})
	}
}

// TestGRPCOnlineQueryE2E mainly tests querying real data
// from the staging server does not crash. Correctness
// is partially tested here, but is mainly tested in
// TestOnlineQueryUnmarshalNonBulkAllTypes.
//
// This test is also notably different from the E2E test
// where a gRPC client is also tested but is built on top
// of the existing REST `Client` interface.
func TestGRPCOnlineQueryE2E(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	client, err := chalk.NewGRPCClient(context.Background())
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}
	err = chalk.InitFeatures(&testFeatures)
	if err != nil {
		t.Fatal("Failed initializing features", err)
	}

	res, queryErr := client.OnlineQuery(context.Background(), getParams())
	if queryErr != nil {
		t.Fatal("Failed querying features", queryErr)
	}

	var testUser user
	assert.NoError(t, chalk.UnmarshalOnlineQueryResponse(res, &testUser))
	testUserValues(t, &testUser)
}

func TestFisQuery(t *testing.T) {
	t.Parallel()

	for _, fixture := range []struct {
		useGrpc bool
	}{
		{useGrpc: true},
	} {
		t.Run(fmt.Sprintf("grpc=%v", fixture.useGrpc), func(t *testing.T) {
			err := chalk.InitFeatures(&testFeatures)
			if err != nil {
				t.Fatal("Failed initializing features", err)
			}

			client, err := chalk.NewClient(context.Background(), &chalk.ClientConfig{UseGrpc: fixture.useGrpc,
				QueryServer: "http://localhost:6666",
				ApiServer:   "http://localhost:4002"})
			if err != nil {
				t.Fatal("Failed creating a Chalk Client", err)
			}

			var _ = map[any]any{"transaction.fi_transaction_id": "fakeee",
				"transaction.acquirer_cntry":     "USA",
				"transaction.acquirer_id":        "someId",
				"transaction.atc_crd":            "someAtcCrd",
				"transaction.atc_host":           "someAtcHost",
				"transaction.atm_network_id":     "someAtmNetworkId",
				"transaction.atm_processing_xcd": "someXcd",
				"transaction.bin":                "12345",
				"transaction.card_verif_results": "1",
				"transaction.crd_port":           "1",
				"transaction.crd_psnt_ind":       "1",
				"transaction.cryptogram_valid":   "1",
				"transaction.cus_card_typ":       "1",
				"transaction.mer_cnty_cd":        "USA",
				"transaction.mer_st":             "IL",
				"transaction.mer_st_search":      "IL",
				"transaction.mer_zip_3":          "123",
				"transaction.pad_response":       "1",
				"transaction.pos_condition_cd":   "1",
				"transaction.process_reason_cd":  "1",
				"transaction.rltm_req":           "1",
				"transaction.trans_category":     "1",
				"transaction.trn_cvv_vrfy_cd":    "1",
				"transaction.trn_pin_vrfy_cd":    "1",
				"transaction.trn_pos_ent_cd":     "1",
				"transaction.trn_typ":            "1",
				"transaction.mer_id":             "fake",
				"transaction.mer_nm":             "1",
				"transaction.customer_xid_hash":  "1",
				"transaction.sic_cd":             "1",
				"transaction.cashback_amt":       "1.0",
				"transaction.trn_amt":            "1.0",
				"transaction.usr_ind_5":          "1",
				"transaction.ext_scor1":          "1",
				"transaction.usr_dat_2":          "1",
				"transaction.trn_dt":             "2025-03-04T00:02:21.339137Z",
				"transaction.crd_exp_dt":         "2025-03-06T14:02:21.339147Z",
			}

			req := chalk.OnlineQueryParams{
				QueryNameVersion: "v1.0.0",
			}.
				WithInput(FISFeatures.Transaction.FiTransactionId, "fakeid").
				WithInput(FISFeatures.Transaction.CashbackAmt, 0.0).
				WithInput(FISFeatures.Transaction.TrnAmt, 0.0).
				WithInput(FISFeatures.Transaction.AcquirerCntry, "USA").
				WithInput(FISFeatures.Transaction.AcquirerId, "someId").
				WithInput(FISFeatures.Transaction.AtcCrd, "req.GetAtcCrd()").
				WithInput(FISFeatures.Transaction.AtcHost, "req.GetAtcHost()").
				WithInput(FISFeatures.Transaction.AtmNetworkId, "req.GetAtmNetworkId()").
				WithInput(FISFeatures.Transaction.AtmProcessingXcd, "req.GetAtmProcessingXcd()").
				WithInput(FISFeatures.Transaction.CardVerifResults, "1").
				WithInput(FISFeatures.Transaction.CrdPsntInd, "1").
				WithInput(FISFeatures.Transaction.CryptogramValid, "1").
				WithInput(FISFeatures.Transaction.CustomerXidHash, "1").
				WithInput(FISFeatures.Transaction.MerId, "fakeMerId").
				WithInput(FISFeatures.Transaction.MerNm, "1").
				WithInput(FISFeatures.Transaction.SicCd, "1").
				WithInput(FISFeatures.Transaction.TrnDt, "2025-03-04T00:02:21.339137Z").
				WithInput(FISFeatures.Transaction.CrdExpDt, "2025-03-06T14:02:21.339147Z").
				WithInput(FISFeatures.Transaction.Bin, "12345").
				WithInput(FISFeatures.Transaction.CrdPort, "1").
				WithInput(FISFeatures.Transaction.CusCardTyp, "1").
				WithInput(FISFeatures.Transaction.MerCntyCd, "USA").
				WithInput(FISFeatures.Transaction.MerSt, "IL").
				WithInput(FISFeatures.Transaction.MerStSearch, "IL").
				WithInput(FISFeatures.Transaction.MerZip3, "123").
				WithInput(FISFeatures.Transaction.PadResponse, "1").
				WithInput(FISFeatures.Transaction.PosConditionCd, "1").
				WithInput(FISFeatures.Transaction.ProcessReasonCd, "1").
				WithInput(FISFeatures.Transaction.RltmReq, "1").
				WithInput(FISFeatures.Transaction.TransCategory, "1").
				WithInput(FISFeatures.Transaction.TrnCvvVrfyCd, "1").
				WithInput(FISFeatures.Transaction.TrnPinVrfyCd, "1").
				WithInput(FISFeatures.Transaction.TrnPosEntCd, "1").
				WithInput(FISFeatures.Transaction.TrnTyp, "1").
				WithInput(FISFeatures.Transaction.UsrInd5, "1").
				WithInput(FISFeatures.Transaction.ExtScor1, "1").
				WithInput(FISFeatures.Transaction.UsrDat2, "1").
				WithQueryName("cardfraud_inference_query").
				WithQueryNameVersion("v1.0.0")

			res, err := client.OnlineQuery(context.Background(), req, nil)
			assert.NoError(t, err)
			transaction := Transaction{}
			err = res.UnmarshalInto(&transaction)
			assert.NoError(t, err)
			// features := createFeaturesFromTransaction(ctx, &transaction)
			fmt.Printf("WindowedCustomerSicFrdTxnAmtSum\t%v\n", transaction.CustomerSic.WindowedCustomerSicFrdTxnCount)
			fmt.Printf("transaction.CustomerSic.WindowedCustomerSicFrdTxnCount[\"1d\"]: ")
			if transaction.CustomerSic.WindowedCustomerSicFrdTxnCount["2d"] != nil {
				fmt.Printf("feature.CustomerSicWindowedCustomerSicFrdTransactionCount__172800__ = %d", *transaction.CustomerSic.WindowedCustomerSicFrdTxnCount["2d"])
			} else {
				fmt.Printf("feature.CustomerSicWindowedCustomerSicFrdTransactionCount__172800__ is NIL!!!")
			}
			fmt.Println()

			if transaction.CustomerSic.WindowedCustomerSicFrdTxnCount["30d"] != nil {
				fmt.Printf("feature.CustomerSicWindowedCustomerSicFrdTransactionCount__2592000__ = %d", *transaction.CustomerSic.WindowedCustomerSicFrdTxnCount["30d"])
			} else {
				fmt.Printf("feature.CustomerSicWindowedCustomerSicFrdTransactionCount__2592000__ is NIL!!!")
			}
			fmt.Println()

			if transaction.CustomerSic.WindowedCustomerSicFrdTxnCount["1w"] != nil {
				fmt.Printf("feature.CustomerSicWindowedCustomerSicFrdTransactionCount__604800__ = %d", *transaction.CustomerSic.WindowedCustomerSicFrdTxnCount["1w"])
			} else {
				fmt.Printf("feature.CustomerSicWindowedCustomerSicFrdTransactionCount__604800__ is NIL!!!")
			}
			fmt.Println()

			if transaction.CustomerSic.WindowedCustomerSicFrdTxnCount["1d"] != nil {
				fmt.Printf("feature.CustomerSicWindowedCustomerSicFrdTransactionCount__86400__ = %d", *transaction.CustomerSic.WindowedCustomerSicFrdTxnCount["1d"])
			} else {
				fmt.Printf("feature.CustomerSicWindowedCustomerSicFrdTransactionCount__86400__ is NIL!!!")

			}
			fmt.Println()
		})
	}
}

// TestOnlineQueryBulkParamsDoesNotErr tests that none
// of the feather header params causes an error when
// specified. Correctness of the thread through is
// tested in TestParamsSetInFeatherHeader. Correctness
// of the results is *not* tested here.
func TestOnlineQueryBulkParamsDoesNotErr(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	for _, fixture := range []struct {
		useGrpc bool
	}{
		{useGrpc: false},
		{useGrpc: true},
	} {
		t.Run(fmt.Sprintf("grpc=%v", fixture.useGrpc), func(t *testing.T) {
			err := chalk.InitFeatures(&testFeatures)
			if err != nil {
				t.Fatal("Failed initializing features", err)
			}

			client, err := chalk.NewClient(context.Background(), &chalk.ClientConfig{UseGrpc: fixture.useGrpc})
			if err != nil {
				t.Fatal("Failed creating a Chalk Client", err)
			}
			userIds := []int{1, 2}

			req := chalk.OnlineQueryParams{
				Tags:                 []string{"named-integration"},
				RequiredResolverTags: []string{"named-integration"},
				Now:                  []time.Time{time.Now(), time.Now()},
				StorePlanStages:      true,
				CorrelationId:        "chalk-go-int-test-correlation-id",
				QueryName:            "chalk-go-int-test-query",
				QueryNameVersion:     "1",
				Meta: map[string]string{
					"test_meta_1": "test_meta_value_1",
					"test_meta_2": "test_meta_value_2",
				},
				Explain: true,
			}.
				WithInput(testFeatures.User.Id, userIds).
				WithOutputs(testFeatures.User.FullName).
				WithStaleness(testFeatures.User.SocureScore, time.Minute*10)

			_, err = client.OnlineQueryBulk(context.Background(), req)
			assert.NoError(t, err)
		})
	}
}

// TestOnlineQueryParamsDoesNotErr tests that none
// of the feather header params causes an error when
// specified. Correctness of the thread through is
// tested in TestParamsSetInOnlineQuery. Correctness
// of the results is *not* tested here.
func TestOnlineQueryParamsDoesNotErr(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	for _, fixture := range []struct {
		useGrpc bool
	}{
		{useGrpc: false},
		{useGrpc: true},
	} {
		t.Run(fmt.Sprintf("grpc=%v", fixture.useGrpc), func(t *testing.T) {
			if fixture.useGrpc {
				t.Skip("CHA-4780")
			}
			client, err := chalk.NewClient(context.Background(), &chalk.ClientConfig{UseGrpc: fixture.useGrpc})
			if err != nil {
				t.Fatal("Failed creating a Chalk Client", err)
			}
			err = chalk.InitFeatures(&testFeatures)
			if err != nil {
				t.Fatal("Failed initializing features", err)
			}

			req := chalk.OnlineQueryParams{
				Tags:                 []string{"named-integration"},
				RequiredResolverTags: []string{"named-integration"},
				Now:                  []time.Time{time.Now()},
				StorePlanStages:      true,
				CorrelationId:        "chalk-go-int-test-correlation-id",
				QueryName:            "chalk-go-int-test-query",
				QueryNameVersion:     "1",
				Meta: map[string]string{
					"test_meta_1": "test_meta_value_1",
					"test_meta_2": "test_meta_value_2",
				},
				Explain: true,
			}.
				WithInput(testFeatures.User.Id, 1).
				WithOutputs(testFeatures.User.FullName).
				WithStaleness(testFeatures.User.SocureScore, time.Minute*10)

			_, err = client.OnlineQuery(context.Background(), req, nil)
			assert.NoError(t, err)
		})
	}
}

func TestCustomCerts(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	systemCertPool, err := x509.SystemCertPool()
	if err != nil {
		t.Fatal("Failed creating a system cert pool", err)
	}
	emptyCertPool := x509.NewCertPool()

	for _, fixture := range []struct {
		useGrpc    bool
		certPool   *x509.CertPool
		shouldFail bool
	}{
		{useGrpc: false, certPool: systemCertPool, shouldFail: false},
		{useGrpc: false, certPool: emptyCertPool, shouldFail: true},
		{useGrpc: true, certPool: systemCertPool, shouldFail: false},
		{useGrpc: true, certPool: emptyCertPool, shouldFail: true},
	} {
		t.Run(fmt.Sprintf("grpc=%v, shouldFail=%v", fixture.useGrpc, fixture.shouldFail), func(t *testing.T) {
			t.Parallel()
			httpClient := http.Client{
				Transport: &http2.Transport{
					TLSClientConfig: &tls.Config{
						RootCAs: fixture.certPool,
					},
				},
			}

			client, err := chalk.NewClient(context.Background(), &chalk.ClientConfig{UseGrpc: fixture.useGrpc, HTTPClient: &httpClient})
			if fixture.shouldFail {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
			}
			var userObj user
			_, queryErr := client.OnlineQuery(context.Background(), getParams(), &userObj)
			assert.NoError(t, queryErr)
		})
	}
}
