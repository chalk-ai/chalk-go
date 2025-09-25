package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManagerURLPrefixes(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		Name            string
		APIServer       string
		GRPCQueryServer string
		JSONQueryServer string

		ExpectedAPIServer       string
		ExpectedGRPCQueryServer string
		ExpectedJSONQueryServer string
	}{
		{
			Name: "no prefixes",

			APIServer:       "api.chalk.ai",
			GRPCQueryServer: "grpc.chalk.ai",
			JSONQueryServer: "json.chalk.ai",

			ExpectedAPIServer:       "https://api.chalk.ai",
			ExpectedGRPCQueryServer: "https://grpc.chalk.ai",
			ExpectedJSONQueryServer: "https://json.chalk.ai",
		},
		{
			Name: "https scheme",

			APIServer:       "https://api.chalk.ai",
			GRPCQueryServer: "https://grpc.chalk.ai",
			JSONQueryServer: "https://json.chalk.ai",

			ExpectedAPIServer:       "https://api.chalk.ai",
			ExpectedGRPCQueryServer: "https://grpc.chalk.ai",
			ExpectedJSONQueryServer: "https://json.chalk.ai",
		},
		{
			Name: "http scheme",

			APIServer:       "http://api.chalk.ai",
			GRPCQueryServer: "http://grpc.chalk.ai",
			JSONQueryServer: "http://json.chalk.ai",

			ExpectedAPIServer:       "http://api.chalk.ai",
			ExpectedGRPCQueryServer: "http://grpc.chalk.ai",
			ExpectedJSONQueryServer: "http://json.chalk.ai",
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			manager, err := NewManager(t.Context(), &ManagerInputs{
				APIServer:       tc.APIServer,
				GRPCQueryServer: tc.GRPCQueryServer,
				JSONQueryServer: tc.JSONQueryServer,
				ClientId:        "test-client-id",
				ClientSecret:    "test-client-secret",
			})
			assert.NoError(t, err)
			assert.NotNil(t, manager)
			assert.Equal(t, tc.ExpectedAPIServer, manager.GetAPIServer().Value)
			assert.Equal(t, tc.ExpectedGRPCQueryServer, manager.GetGRPCQueryServer().Value)
			assert.Equal(t, tc.ExpectedJSONQueryServer, manager.GetJSONQueryServer().Value)
		})
	}
}

func TestSetters(t *testing.T) {
	t.Parallel()
	manager, err := NewManager(t.Context(), &ManagerInputs{
		APIServer:       "api.chalk.ai",
		GRPCQueryServer: "grpc.chalk.ai",
		JSONQueryServer: "json.chalk.ai",
		ClientId:        "test-client-id",
		ClientSecret:    "test-abc",
	})

	assert.NoError(t, err)

	manager.SetAPIServer(NewFromToken("chalk.myendpoint.ai", "wow nice"))
	assert.Equal(t, "https://chalk.myendpoint.ai", manager.GetAPIServer().Value)
	manager.SetAPIServer(NewFromToken("https://chalk.myendpoint.ai", "wow"))
	assert.Equal(t, "https://chalk.myendpoint.ai", manager.GetAPIServer().Value)

	manager.SetGRPCQueryServer(NewFromToken("grpc.chalk.ai", "wow nice"))
	assert.Equal(t, "https://grpc.chalk.ai", manager.GetGRPCQueryServer().Value)
	manager.SetGRPCQueryServer(NewFromToken("http://grpc.chalk.ai", "wow"))
	assert.Equal(t, "http://grpc.chalk.ai", manager.GetGRPCQueryServer().Value)

	manager.SetJSONQueryServer(NewFromToken("json.chalk.ai", "wow nice"))
	assert.Equal(t, "https://json.chalk.ai", manager.GetJSONQueryServer().Value)
	manager.SetJSONQueryServer(NewFromToken("http://json.chalk.ai", "wow"))
	assert.Equal(t, "http://json.chalk.ai", manager.GetJSONQueryServer().Value)
}
