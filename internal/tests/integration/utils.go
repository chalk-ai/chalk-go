package integration

import (
	"os"
	"testing"
)

func SkipIfNotIntegrationTester(t *testing.T) {
	if os.Getenv("INTEGRATION_TESTER") != "true" {
		//t.Skip()
	}
}
