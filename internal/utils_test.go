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

func TestTranslateBracketFqn(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "days to seconds",
			input: "user.count[1d]",
			want:  "user.count__86400__",
		},
		{
			name:  "hours to seconds",
			input: "user.count[1h]",
			want:  "user.count__3600__",
		},
		{
			name:  "minutes to seconds",
			input: "user.count[1m]",
			want:  "user.count__60__",
		},
		{
			name:  "weeks to seconds",
			input: "user.count[1w]",
			want:  "user.count__604800__",
		},
		{
			name:  "7d and 1w map to the same seconds",
			input: "user.count[7d]",
			want:  "user.count__604800__",
		},
		{
			name:  "all bucket",
			input: "user.count[all]",
			want:  "user.count__all__",
		},
		{
			name:  "version suffix preserved",
			input: "user.count[1d]@2",
			want:  "user.count__86400__@2",
		},
		{
			name:  "all bucket with version suffix",
			input: "user.count[all]@2",
			want:  "user.count__all__@2",
		},
		{
			name:  "90m (5400s)",
			input: "user.count[90m]",
			want:  "user.count__5400__",
		},
		{
			name:  "multi-segment namespace",
			input: "user.relationship.agg[1d]",
			want:  "user.relationship.agg__86400__",
		},
		{
			name:  "non-windowed passthrough",
			input: "user.age",
			want:  "user.age",
		},
		{
			name:  "already-internal passthrough",
			input: "user.count__86400__",
			want:  "user.count__86400__",
		},
		{
			name:  "invalid duration passthrough",
			input: "user.count[banana]",
			want:  "user.count[banana]",
		},
		{
			name:  "round trips with TranslateWindowedFqn",
			input: "user.count[1w]",
			want:  "user.count__604800__",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.want, TranslateBracketFqn(tc.input))
		})
	}
}

func TestBracketFqnRoundTrip(t *testing.T) {
	t.Parallel()
	// Bracket notation -> internal -> bracket notation should be stable for
	// the canonical bracket forms emitted by TranslateWindowedFqn.
	for _, bracket := range []string{
		"user.count[1d]",
		"user.count[1h]",
		"user.count[1m]",
		"user.count[1w]",
		"user.count[all]",
		"user.count[1d]@2",
		"user.count[all]@2",
	} {
		require.Equal(t, bracket, TranslateWindowedFqn(TranslateBracketFqn(bracket)))
	}
}
