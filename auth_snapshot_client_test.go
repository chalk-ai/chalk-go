package chalk

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/chalk-ai/chalk-go/auth"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type authCaptureTransport struct {
	request *http.Request
}

func (t *authCaptureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.request = req
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
		Body:       io.NopCloser(bytes.NewBufferString(`{}`)),
		Request:    req,
	}, nil
}

type authRetryTransport struct {
	statuses []int
	headers  []http.Header
}

func (t *authRetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.headers = append(t.headers, req.Header.Clone())
	status := t.statuses[min(len(t.headers)-1, len(t.statuses)-1)]
	return &http.Response{
		StatusCode: status,
		Header:     make(http.Header),
		Body:       io.NopCloser(bytes.NewBufferString("{}")),
		Request:    req,
	}, nil
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func staleAuthToken() *serverv1.GetTokenResponse {
	return &serverv1.GetTokenResponse{
		AccessToken: "token-initial",
		ExpiresAt:   timestamppb.New(time.Now().Add(30 * time.Second)),
	}
}

func freshAuthToken() *serverv1.GetTokenResponse {
	return &serverv1.GetTokenResponse{
		AccessToken: "token-initial",
		ExpiresAt:   timestamppb.New(time.Now().Add(time.Hour)),
	}
}

func rotatedAuthProvider(context.Context) (*auth.AuthSnapshot, error) {
	return &auth.AuthSnapshot{
		Token: &serverv1.GetTokenResponse{
			AccessToken: "token-rotated",
			ExpiresAt:   timestamppb.New(time.Now().Add(time.Hour)),
		},
		EnvironmentID: "env-rotated",
	}, nil
}

func TestHTTPClientSendsRotatedAuthSnapshot(t *testing.T) {
	capture := &authCaptureTransport{}
	client, err := NewClient(t.Context(), &ClientConfig{
		ApiServer:                  "https://api.chalk.ai",
		QueryServer:                "https://engine.chalk.ai",
		EnvironmentId:              "env-initial",
		JWT:                        staleAuthToken(),
		AuthProvider:               rotatedAuthProvider,
		HTTPClient:                 &http.Client{Transport: capture},
		SkipEnvironmentNameMapping: true,
		SkipEngineMapping:          true,
	})
	require.NoError(t, err)

	impl := client.(*clientImpl)
	var response map[string]any
	require.NoError(t, impl.sendRequest(t.Context(), &sendRequestParams{
		Method:   http.MethodGet,
		URL:      "https://api.chalk.ai/test",
		Response: &response,
	}))
	require.Equal(t, "Bearer token-rotated", capture.request.Header.Get("Authorization"))
	require.Equal(t, "env-rotated", capture.request.Header.Get("X-Chalk-Env-Id"))
}

func TestHTTPClientRefreshesRejectedFreshAuth(t *testing.T) {
	for _, status := range []int{http.StatusUnauthorized, http.StatusForbidden} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			transport := &authRetryTransport{statuses: []int{status, http.StatusOK}}
			providerCalls := 0
			invalidatorCalls := 0
			client, err := NewClient(t.Context(), &ClientConfig{
				ApiServer:     "https://api.chalk.ai",
				QueryServer:   "https://engine.chalk.ai",
				EnvironmentId: "env-initial",
				JWT:           freshAuthToken(),
				AuthProvider: func(ctx context.Context) (*auth.AuthSnapshot, error) {
					providerCalls++
					return rotatedAuthProvider(ctx)
				},
				AuthProviderInvalidator: func(context.Context) {
					invalidatorCalls++
				},
				HTTPClient:                 &http.Client{Transport: transport},
				SkipEnvironmentNameMapping: true,
				SkipEngineMapping:          true,
			})
			require.NoError(t, err)

			var response map[string]any
			require.NoError(t, client.(*clientImpl).sendRequest(t.Context(), &sendRequestParams{
				Method:   http.MethodGet,
				URL:      "https://api.chalk.ai/test",
				Response: &response,
			}))

			require.Len(t, transport.headers, 2)
			require.Equal(t, "Bearer token-initial", transport.headers[0].Get("Authorization"))
			require.Equal(t, "env-initial", transport.headers[0].Get("X-Chalk-Env-Id"))
			require.Equal(t, "Bearer token-rotated", transport.headers[1].Get("Authorization"))
			require.Equal(t, "env-rotated", transport.headers[1].Get("X-Chalk-Env-Id"))
			require.Equal(t, 1, providerCalls)
			require.Equal(t, 1, invalidatorCalls)
		})
	}
}

func TestHTTPClientRetriesAuthenticationRejectionOnlyOnce(t *testing.T) {
	transport := &authRetryTransport{statuses: []int{
		http.StatusUnauthorized,
		http.StatusUnauthorized,
		http.StatusOK,
	}}
	client, err := NewClient(t.Context(), &ClientConfig{
		ApiServer:                  "https://api.chalk.ai",
		QueryServer:                "https://engine.chalk.ai",
		EnvironmentId:              "env-initial",
		JWT:                        freshAuthToken(),
		AuthProvider:               rotatedAuthProvider,
		HTTPClient:                 &http.Client{Transport: transport},
		SkipEnvironmentNameMapping: true,
		SkipEngineMapping:          true,
	})
	require.NoError(t, err)

	var response map[string]any
	err = client.(*clientImpl).sendRequest(t.Context(), &sendRequestParams{
		Method:   http.MethodGet,
		URL:      "https://api.chalk.ai/test",
		Response: &response,
	})

	require.Error(t, err)
	require.Len(t, transport.headers, 2)
}

func TestGRPCClientSendsRotatedAuthSnapshot(t *testing.T) {
	client, err := NewGRPCClient(t.Context(), &GRPCClientConfig{
		ApiServer:                  "https://api.chalk.ai",
		QueryServer:                "https://engine.chalk.ai",
		EnvironmentId:              "env-initial",
		JWT:                        staleAuthToken(),
		AuthProvider:               rotatedAuthProvider,
		SkipEnvironmentNameMapping: true,
		SkipEngineMapping:          true,
	})
	require.NoError(t, err)

	var captured http.Header
	next := func(_ context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		captured = req.Header().Clone()
		return nil, nil
	}
	req := connect.NewRequest(&serverv1.GetGraphRequest{})
	_, err = client.(*grpcClientImpl).engineInterceptor(next)(t.Context(), req)
	require.NoError(t, err)
	require.Equal(t, "Bearer token-rotated", captured.Get("Authorization"))
	require.Equal(t, "env-rotated", captured.Get("x-chalk-env-id"))
}

func TestGRPCClientRefreshesRejectedFreshAuth(t *testing.T) {
	providerCalls := 0
	invalidatorCalls := 0
	client, err := NewGRPCClient(t.Context(), &GRPCClientConfig{
		ApiServer:     "https://api.chalk.ai",
		QueryServer:   "https://engine.chalk.ai",
		EnvironmentId: "env-initial",
		JWT:           freshAuthToken(),
		AuthProvider: func(ctx context.Context) (*auth.AuthSnapshot, error) {
			providerCalls++
			return rotatedAuthProvider(ctx)
		},
		AuthProviderInvalidator: func(context.Context) {
			invalidatorCalls++
		},
		SkipEnvironmentNameMapping: true,
		SkipEngineMapping:          true,
	})
	require.NoError(t, err)

	var headers []http.Header
	next := func(_ context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		headers = append(headers, req.Header().Clone())
		if len(headers) == 1 {
			return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("rejected"))
		}
		return nil, nil
	}
	req := connect.NewRequest(&serverv1.GetGraphRequest{})
	_, err = client.(*grpcClientImpl).engineInterceptor(next)(t.Context(), req)

	require.NoError(t, err)
	require.Len(t, headers, 2)
	require.Equal(t, "Bearer token-initial", headers[0].Get("Authorization"))
	require.Equal(t, "env-initial", headers[0].Get("x-chalk-env-id"))
	require.Equal(t, "Bearer token-rotated", headers[1].Get("Authorization"))
	require.Equal(t, "env-rotated", headers[1].Get("x-chalk-env-id"))
	require.Equal(t, 1, providerCalls)
	require.Equal(t, 1, invalidatorCalls)
}

func TestGRPCClientAuthenticationRetryBounds(t *testing.T) {
	testCases := []struct {
		name          string
		code          connect.Code
		expectedCalls int
		expectedFresh int
	}{
		{
			name:          "unauthenticated retries once",
			code:          connect.CodeUnauthenticated,
			expectedCalls: 2,
			expectedFresh: 1,
		},
		{
			name:          "permission denied is not an authentication retry",
			code:          connect.CodePermissionDenied,
			expectedCalls: 1,
			expectedFresh: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			providerCalls := 0
			client, err := NewGRPCClient(t.Context(), &GRPCClientConfig{
				ApiServer:     "https://api.chalk.ai",
				QueryServer:   "https://engine.chalk.ai",
				EnvironmentId: "env-initial",
				JWT:           freshAuthToken(),
				AuthProvider: func(ctx context.Context) (*auth.AuthSnapshot, error) {
					providerCalls++
					return rotatedAuthProvider(ctx)
				},
				SkipEnvironmentNameMapping: true,
				SkipEngineMapping:          true,
			})
			require.NoError(t, err)

			calls := 0
			next := func(context.Context, connect.AnyRequest) (connect.AnyResponse, error) {
				calls++
				return nil, connect.NewError(tc.code, errors.New("rejected"))
			}
			req := connect.NewRequest(&serverv1.GetGraphRequest{})
			_, err = client.(*grpcClientImpl).engineInterceptor(next)(t.Context(), req)

			require.Error(t, err)
			require.Equal(t, tc.expectedCalls, calls)
			require.Equal(t, tc.expectedFresh, providerCalls)
		})
	}
}

func TestVolumeClientSendsRotatedAuthSnapshot(t *testing.T) {
	client, err := NewVolumeClient(t.Context(), &VolumeClientConfig{
		ApiServer:                  "https://api.chalk.ai",
		EnvironmentId:              "env-initial",
		JWT:                        staleAuthToken(),
		AuthProvider:               rotatedAuthProvider,
		SkipEnvironmentNameMapping: true,
		SkipEngineMapping:          true,
	})
	require.NoError(t, err)

	header := http.Header{}
	_, err = client.(*volumeClientImpl).addAuthHeaders(t.Context(), header)
	require.NoError(t, err)
	require.Equal(t, "Bearer token-rotated", header.Get("Authorization"))
	require.Equal(t, "env-rotated", header.Get("x-chalk-env-id"))
}

func TestVolumeClientRefreshesRejectedFreshAuth(t *testing.T) {
	providerCalls := 0
	invalidatorCalls := 0
	client, err := NewVolumeClient(t.Context(), &VolumeClientConfig{
		ApiServer:     "https://api.chalk.ai",
		EnvironmentId: "env-initial",
		JWT:           freshAuthToken(),
		AuthProvider: func(ctx context.Context) (*auth.AuthSnapshot, error) {
			providerCalls++
			return rotatedAuthProvider(ctx)
		},
		AuthProviderInvalidator: func(context.Context) {
			invalidatorCalls++
		},
		SkipEnvironmentNameMapping: true,
		SkipEngineMapping:          true,
	})
	require.NoError(t, err)

	var headers []http.Header
	next := func(_ context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		headers = append(headers, req.Header().Clone())
		if len(headers) == 1 {
			return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("rejected"))
		}
		return nil, nil
	}
	req := connect.NewRequest(&serverv1.GetGraphRequest{})
	_, err = client.(*volumeClientImpl).authInterceptor()(next)(t.Context(), req)

	require.NoError(t, err)
	require.Len(t, headers, 2)
	require.Equal(t, "Bearer token-initial", headers[0].Get("Authorization"))
	require.Equal(t, "env-initial", headers[0].Get("x-chalk-env-id"))
	require.Equal(t, "Bearer token-rotated", headers[1].Get("Authorization"))
	require.Equal(t, "env-rotated", headers[1].Get("x-chalk-env-id"))
	require.Equal(t, 1, providerCalls)
	require.Equal(t, 1, invalidatorCalls)
}

func TestVolumeHTTPRefreshesRejectedFreshAuth(t *testing.T) {
	var headers []http.Header
	transport := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		headers = append(headers, req.Header.Clone())
		status := http.StatusUnauthorized
		body := "{}"
		if len(headers) == 2 {
			status = http.StatusOK
			body = "{\"user\":\"test-user\"}"
		}
		return &http.Response{
			StatusCode: status,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString(body)),
			Request:    req,
		}, nil
	})
	invalidatorCalls := 0
	client, err := NewVolumeClient(t.Context(), &VolumeClientConfig{
		ApiServer:     "https://api.chalk.ai",
		EnvironmentId: "env-initial",
		JWT:           freshAuthToken(),
		AuthProvider:  rotatedAuthProvider,
		AuthProviderInvalidator: func(context.Context) {
			invalidatorCalls++
		},
		HTTPClient:                 &http.Client{Transport: transport},
		SkipEnvironmentNameMapping: true,
		SkipEngineMapping:          true,
	})
	require.NoError(t, err)

	author := client.(*volumeClientImpl).resolveCommitAuthor(t.Context())

	require.Equal(t, "chalk:env-rotated:agent:test-user", author)
	require.Len(t, headers, 2)
	require.Equal(t, "Bearer token-initial", headers[0].Get("Authorization"))
	require.Equal(t, "Bearer token-rotated", headers[1].Get("Authorization"))
	require.Equal(t, 1, invalidatorCalls)
}
