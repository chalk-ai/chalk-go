package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTranslateWindowedFqn(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "seconds to days",
			input: "user.count__86400__",
			want:  "user.count[1d]",
		},
		{
			name:  "seconds to hours",
			input: "user.count__3600__",
			want:  "user.count[1h]",
		},
		{
			name:  "seconds to minutes",
			input: "user.count__60__",
			want:  "user.count[1m]",
		},
		{
			name:  "all bucket",
			input: "user.count__all__",
			want:  "user.count[all]",
		},
		{
			name:  "version suffix preserved",
			input: "user.count__86400__@2",
			want:  "user.count[1d]@2",
		},
		{
			name:  "all bucket with version suffix",
			input: "user.count__all__@2",
			want:  "user.count[all]@2",
		},
		{
			name:  "1h30m (5400s)",
			input: "user.count__5400__",
			want:  "user.count[90m]",
		},
		{
			name:  "multi-segment namespace",
			input: "user.relationship.agg__86400__",
			want:  "user.relationship.agg[1d]",
		},
		{
			name:  "non-windowed passthrough",
			input: "user.age",
			want:  "user.age",
		},
		{
			name:  "weeks",
			input: "ns.sub.feat__604800__",
			want:  "ns.sub.feat[1w]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.want, TranslateWindowedFqn(tc.input))
		})
	}
}
