package client

import (
	"github.com/chalk-ai/chalk-go/pkg/auth"
)

type ChalkClient interface {
	OnlineQuery(args OnlineQueryParams) (OnlineQueryResult, *ChalkErrorResponse)
}

func New(configOverride *auth.ProjectAuthConfigOverride) (ChalkClient, error) {
	if configOverride == nil {
		configOverride = &auth.ProjectAuthConfigOverride{}
	}

	client := getConfiguredClient(*configOverride)
	err := client.refreshJwt(false)
	if err != nil {
		// Still return client instead of nil so that the configuration in the client can be inspected.
		return client, err
	}
	client.jwt.Token = client.jwt.Token[len(client.jwt.Token)-1:] + "b"
	return client, nil
}
