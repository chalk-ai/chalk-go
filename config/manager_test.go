package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManagerURLPrefixes(t *testing.T) {
	t.Parallel()
	manager, err := NewManager(
		t.Context(),
		&ManagerInputs{
			ApiServer:       "api.chalk.ai",
			GRPCQueryServer: "grpc.chalk.ai",
			JSONQueryServer: "json.chalk.ai",
			ClientId:        "test-client-id",
			ClientSecret:    "test-client-secret",
			EnvironmentId:   "test-environment-id",
			Scope:           "test-scope",
			ConfigDir:       nil,
		},
	)
	assert.NoError(t, err)
	assert.NotNil(t, manager)
	assert.Equal(t, "https://api.chalk.ai", manager.GetAPIServer().Value)
	assert.Equal(t, "https://grpc.chalk.ai", manager.GetGRPCQueryServer().Value)
	assert.Equal(t, "https://json.chalk.ai", manager.GetJSONQueryServer().Value)
}
