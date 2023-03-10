package client

import (
	"github.com/chalk-ai/chalk-go/pkg/auth"
)

type ChalkClient interface {
	OnlineQuery(args OnlineQueryParams) (OnlineQueryResult, *ChalkErrorResponse)
}

func New(configOverride *auth.ProjectAuthConfigOverride) (ChalkClient, error) {
	client := getConfiguredClient(configOverride)
	err := client.refreshJwt(false)
	if err != nil {
		return nil, err
	}
	return client, nil
}
