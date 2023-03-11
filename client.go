package chalk

import (
	"net/http"
)

// ChalkClient is the primary interface for interacting with Chalk. You can use
// it to query data, trigger resolver runs, gather offline data, and more.
type ChalkClient interface {
	OnlineQuery(args OnlineQueryParams) (OnlineQueryResult, *ChalkErrorResponse)
}

type ClientConfig struct {
	ClientId      string
	ClientSecret  string
	ApiServer     string
	EnvironmentId string
	Logger        *LeveledLogger
	Client        *http.Client
}

// New creates a ChalkClient with authentication settings configured.
// These settings can be overriden by passing in a ClientConfig
// object. Otherwise, for each configuration variable, New uses its
// corresponding environment variable if it exists. The environment variables
// that New looks for are:
//
//	CHALK_ACTIVE_ENVIRONMENT
//	CHALK_API_SERVER
//	CHALK_CLIENT_ID
//	CHALK_CLIENT_SECRET
//
// For each config variable, if it is still not found, New will look for a
// `~/.chalk.yml` file, which is updated when you run `chalk login`.
// If a configuration for the specific project directory if found,
// that configuration will be used. Otherwise, the configuration under
// the key `default` will be used.
func New(config *ClientConfig) (ChalkClient, error) {
	c := getConfiguredClient(config)
	err := c.refreshJwt(false)
	if config.Logger != nil {
		c.logger = config.Logger
	}
	if err != nil {
		return nil, err
	}
	return c, nil
}
