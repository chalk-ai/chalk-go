package client

import (
	"github.com/chalk-ai/chalk-go/pkg/auth"
)

// ChalkClient is the primary interface for interacting with Chalk. You can use
// it to query data, trigger resolver runs, gather offline data, and more.
type ChalkClient interface {
	OnlineQuery(args OnlineQueryParams) (OnlineQueryResult, *ChalkErrorResponse)
	TriggerResolverRun(args TriggerResolverRunParams) (TriggerResolverRunResult, *ChalkErrorResponse)
}

// New creates a ChalkClient with authentication settings configured.
// These settings can be overriden by passing in a ProjectAuthConfigOverride
// object. Otherwise, for each configuration variable, New uses its
// corresponding environment variable if it exists. The environment variables
// that New looks for are:
//
//	_CHALK_ACTIVE_ENVIRONMENT
//	_CHALK_API_SERVER
//	CHALK_CLIENT_ID
//	CHALK_CLIENT_SECRET
//
// For each config variable, if it is still not found, New will look for a
// `~/.chalk.yml` file, which is updated when you run `chalk login`.
// If a configuration for the specific project directory if found,
// that configuration will be used. Otherwise, the configuration under
// the key `default` will be used.
func New(configOverride *auth.ProjectAuthConfigOverride) (ChalkClient, error) {
	client := getConfiguredClient(configOverride)
	err := client.refreshJwt(false)
	if err != nil {
		return nil, err
	}
	return client, nil
}
