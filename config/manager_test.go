package config

import (
	"testing"

	"github.com/stretchr/testify/require"
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
	require.NoError(t, err)
	require.NotNil(t, manager)
	require.Equal(t, "https://api.chalk.ai", manager.ApiServer.Value)
	require.Equal(t, "https://grpc.chalk.ai", manager.GRPCQueryServer.Value)
	require.Equal(t, "https://json.chalk.ai", manager.JSONQueryServer.Value)
}
