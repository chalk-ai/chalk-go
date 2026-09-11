package chalk

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/chalk-ai/chalk-go/auth"
	"github.com/cockroachdb/errors"
)

const authRefreshHeadroom = time.Minute

func setAuthHeaders(header http.Header, snapshot *auth.AuthSnapshot) {
	header.Set("x-chalk-env-id", snapshot.EnvironmentID)
	header.Set("Authorization", fmt.Sprintf("Bearer %s", snapshot.Token.AccessToken))
}

// sendAuthenticatedUnary adds the current authentication snapshot and retries
// once when the server rejects it. RefreshAuthAfterRejection ensures concurrent
// failures do not each refresh or discard a snapshot another request installed.
func sendAuthenticatedUnary(
	ctx context.Context,
	req connect.AnyRequest,
	next connect.UnaryFunc,
	tokenManager *auth.Manager,
) (connect.AnyResponse, error) {
	snapshot, err := tokenManager.GetAuth(ctx, time.Now().Add(authRefreshHeadroom))
	if err != nil {
		return nil, errors.Wrap(err, "refreshing authentication")
	}
	setAuthHeaders(req.Header(), snapshot)

	res, err := next(ctx, req)
	if err == nil || connect.CodeOf(err) != connect.CodeUnauthenticated {
		return res, err
	}

	refreshed, refreshErr := tokenManager.RefreshAuthAfterRejection(
		ctx,
		snapshot,
		time.Now().Add(authRefreshHeadroom),
	)
	if refreshErr != nil {
		return res, errors.Wrap(refreshErr, "refreshing authentication after rejection")
	}
	setAuthHeaders(req.Header(), refreshed)
	return next(ctx, req)
}
