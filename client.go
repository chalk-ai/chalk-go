package chalk

import (
	"github.com/chalk-ai/chalk-go/internal/auth"
	"net/http"
)

type ChalkClient interface {
	OnlineQuery(args OnlineQueryParams) (OnlineQueryResult, *ChalkErrorResponse)
	SetLogger(logger *LeveledLogger)
	SetHTTPClient(logger *http.Client)
}

func New(configOverride *auth.ProjectAuthConfigOverride) (ChalkClient, error) {
	return getConfiguredClient(configOverride)
}
