package chalk

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/apache/arrow-go/v18/arrow/array"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/gen/chalk/engine/v1/enginev1connect"
	serverv1connect "github.com/chalk-ai/chalk-go/gen/chalk/server/v1/serverv1connect"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/tests/fixtures"
	assert "github.com/stretchr/testify/require"
)

const roundTripRows = 1000

// bulkEchoHandler plays the engine: it decodes the Arrow request exactly as
// received over the wire and answers with an Arrow scalars_data response
// derived from it.
type bulkEchoHandler struct {
	enginev1connect.UnimplementedQueryServiceHandler
	mu     sync.Mutex
	gotIds []string
}

func (h *bulkEchoHandler) OnlineQueryBulk(
	_ context.Context,
	req *connect.Request[commonv1.OnlineQueryBulkRequest],
) (*connect.Response[commonv1.OnlineQueryBulkResponse], error) {
	table, err := internal.ConvertBytesToTable(req.Msg.GetInputsFeather(), fixtures.TestAllocator)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	if table == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("nil table"))
	}
	idCol := table.Column(0).Data().Chunk(0).(*array.LargeString)
	ids := make([]string, idCol.Len())
	scores := make([]float64, idCol.Len())
	counts := make([]*int64, idCol.Len())
	for i := range ids {
		ids[i] = idCol.Value(i)
		scores[i] = float64(i) / 4
		if i%50 == 0 {
			v := int64(i)
			counts[i] = &v
		}
	}
	h.mu.Lock()
	h.gotIds = ids
	h.mu.Unlock()
	scalars, err := internal.InputsToArrowBytes(map[string]any{
		"round_trip_row.id":    ids,
		"round_trip_row.score": scores,
		"round_trip_row.count": counts,
	}, fixtures.TestAllocator)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&commonv1.OnlineQueryBulkResponse{ScalarsData: scalars}), nil
}

type roundTripRow struct {
	Id    *string
	Score *float64
	Count *int64
}

// The full gRPC path: build request, encode over the connect protocol to a
// server that decodes it, return an Arrow response, then read it back via both
// UnmarshalInto and GetRow.
func TestGRPCOnlineQueryBulkRoundTrip(t *testing.T) {
	handler := &bulkEchoHandler{}
	mux := http.NewServeMux()
	path, h := enginev1connect.NewQueryServiceHandler(handler)
	mux.Handle(path, h)
	authPath, authHandler := serverv1connect.NewAuthServiceHandler(&minimalAuthHandler{})
	mux.Handle(authPath, authHandler)
	graphPath, graphHandler := serverv1connect.NewGraphServiceHandler(&minimalGraphHandler{})
	mux.Handle(graphPath, graphHandler)
	server := startUnencryptedHTTP2Server(t, mux)
	client := newTestGRPCClient(t, server.URL, server.URL, "")

	ids := make([]string, roundTripRows)
	for i := range ids {
		ids[i] = fmt.Sprintf("id-%d", i)
	}
	params := OnlineQueryParamsComplete{}.
		WithInput("round_trip_row.id", ids).
		WithOutputs("round_trip_row.score", "round_trip_row.count")
	res, err := client.OnlineQueryBulk(context.Background(), params)
	assert.NoError(t, err)
	assert.Equal(t, ids, handler.gotIds)

	var rows []roundTripRow
	assert.NoError(t, res.UnmarshalInto(&rows))
	assert.Len(t, rows, roundTripRows)
	for i, r := range rows {
		assert.Equal(t, ids[i], *r.Id)
		assert.Equal(t, float64(i)/4, *r.Score)
		if i%50 == 0 {
			assert.Equal(t, int64(i), *r.Count)
		} else {
			assert.Nil(t, r.Count)
		}
	}

	for _, i := range []int{0, 1, 50, 999} {
		row, err := res.GetRow(i)
		assert.NoError(t, err)
		assert.Equal(t, float64(i)/4, row.Features["round_trip_row.score"].Value)
		assert.Equal(t, ids[i], row.Features["round_trip_row.id"].Value)
	}
	_, err = res.GetRow(roundTripRows)
	assert.Error(t, err)

	// GetTable hands back an independent table each time.
	t1, err := res.GetTable()
	assert.NoError(t, err)
	t1.Release()
	t2, err := res.GetTable()
	assert.NoError(t, err)
	assert.Equal(t, int64(roundTripRows), t2.NumRows())
	t2.Release()
}

// Corrupt Arrow bytes may make the decoder panic. A recovered panic must not
// leave the result's cache lock held, which would hang every later call.
func TestGRPCBulkResultCorruptBytesDoNotDeadlock(t *testing.T) {
	good, err := internal.InputsToArrowBytes(map[string]any{"round_trip_row.score": []float64{1, 2, 3, 4}}, fixtures.TestAllocator)
	assert.NoError(t, err)
	call := func(res *GRPCOnlineQueryBulkResult) (panicked bool) {
		defer func() { panicked = recover() != nil }()
		_, _ = res.GetRow(1)
		return false
	}
	for seed := 0; seed < 200; seed++ {
		bad := append([]byte(nil), good...)
		bad[len(bad)/2+seed%7] ^= 0xFF
		bad[len(bad)-12-seed%5] ^= 0x5A
		res, err := NewGRPCOnlineQueryBulkResult(&commonv1.OnlineQueryBulkResponse{ScalarsData: bad})
		assert.NoError(t, err)
		done := make(chan struct{})
		go func() {
			defer close(done)
			call(res)
			call(res) // would block forever if the first call leaked the lock
		}()
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			t.Fatalf("deadlock on corrupt input (seed %d)", seed)
		}
	}
}
