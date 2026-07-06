package chalk

import (
	"testing"

	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	assert "github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/structpb"
)

// TestQueryMetaFromProtoMissedPlanCache verifies that the plan-cache participation flag is
// read from the free-form additional_metadata map, and that absent or non-bool values are
// reported as nil (server did not report it).
func TestQueryMetaFromProtoMissedPlanCache(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		name     string
		meta     *commonv1.OnlineQueryMetadata
		expected *bool
	}{
		{
			name: "missed",
			meta: &commonv1.OnlineQueryMetadata{
				AdditionalMetadata: map[string]*structpb.Value{
					"missed_plan_cache": structpb.NewBoolValue(true),
				},
			},
			expected: new(true),
		},
		{
			name: "served from cache",
			meta: &commonv1.OnlineQueryMetadata{
				AdditionalMetadata: map[string]*structpb.Value{
					"missed_plan_cache": structpb.NewBoolValue(false),
				},
			},
			expected: new(false),
		},
		{
			name:     "not reported",
			meta:     &commonv1.OnlineQueryMetadata{},
			expected: nil,
		},
		{
			name: "non-bool value",
			meta: &commonv1.OnlineQueryMetadata{
				AdditionalMetadata: map[string]*structpb.Value{
					"missed_plan_cache": structpb.NewStringValue("yes"),
				},
			},
			expected: nil,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			converted := queryMetaFromProto(tt.meta)
			assert.NotNil(t, converted)
			if tt.expected == nil {
				assert.Nil(t, converted.MissedPlanCache)
			} else {
				assert.NotNil(t, converted.MissedPlanCache)
				assert.Equal(t, *tt.expected, *converted.MissedPlanCache)
			}
		})
	}
}
