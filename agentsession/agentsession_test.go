package agentsession

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/chalk-ai/chalk-go/envfs"
	"github.com/stretchr/testify/assert"
)

func ctxWithEnv(t *testing.T, env map[string]string) context.Context {
	t.Helper()
	return envfs.ContextWithEnvironmentGetter(t.Context(), envfs.NewMockEnvironmentGetter(envfs.WithEnv(env)))
}

func TestDetect(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name              string
		env               map[string]string
		expectedSessionId string
		expectedAgent     string
	}{
		{
			name: "claude code",
			env: map[string]string{
				"CLAUDE_CODE_SESSION_ID": "3ac46c84-a182-43b3-a619-8f86349b222c",
				"AI_AGENT":               "claude-code_2-1-222_agent",
			},
			expectedSessionId: "3ac46c84-a182-43b3-a619-8f86349b222c",
			expectedAgent:     "claude-code_2-1-222_agent",
		},
		{
			name:              "no agent",
			env:               map[string]string{},
			expectedSessionId: "",
			expectedAgent:     "",
		},
		{
			name: "codex",
			env: map[string]string{
				"CODEX_THREAD_ID": "019fea01-5967-7c31-b12a-de0e9cc19702",
			},
			expectedSessionId: "019fea01-5967-7c31-b12a-de0e9cc19702",
			expectedAgent:     "codex",
		},
		{
			name: "chalk overrides win over the harness",
			env: map[string]string{
				"CHALK_AGENT_SESSION_ID": "explicit-session",
				"CHALK_AGENT":            "explicit-agent",
				"CLAUDE_CODE_SESSION_ID": "3ac46c84-a182-43b3-a619-8f86349b222c",
				"AI_AGENT":               "claude-code_2-1-222_agent",
			},
			expectedSessionId: "explicit-session",
			expectedAgent:     "explicit-agent",
		},
		{
			name: "unrecognized harness supplies values directly",
			env: map[string]string{
				"CHALK_AGENT_SESSION_ID": "some-session",
				"CHALK_AGENT":            "some-agent",
			},
			expectedSessionId: "some-session",
			expectedAgent:     "some-agent",
		},
		{
			name: "session without a known agent name",
			env: map[string]string{
				"CLAUDE_CODE_SESSION_ID": "session-only",
			},
			expectedSessionId: "session-only",
			expectedAgent:     "claude-code",
		},
		{
			name: "explicitly disabled",
			env: map[string]string{
				"CHALK_AGENT_ATTRIBUTION": "0",
				"CLAUDE_CODE_SESSION_ID":  "3ac46c84-a182-43b3-a619-8f86349b222c",
				"AI_AGENT":                "claude-code_2-1-222_agent",
			},
			expectedSessionId: "",
			expectedAgent:     "",
		},
		{
			name: "disabled by word",
			env: map[string]string{
				"CHALK_AGENT_ATTRIBUTION": "False",
				"CLAUDE_CODE_SESSION_ID":  "3ac46c84-a182-43b3-a619-8f86349b222c",
			},
			expectedSessionId: "",
			expectedAgent:     "",
		},
		{
			name: "an unrelated value leaves attribution on",
			env: map[string]string{
				"CHALK_AGENT_ATTRIBUTION": "1",
				"CLAUDE_CODE_SESSION_ID":  "3ac46c84-a182-43b3-a619-8f86349b222c",
			},
			expectedSessionId: "3ac46c84-a182-43b3-a619-8f86349b222c",
			expectedAgent:     "claude-code",
		},
		{
			name: "surrounding whitespace is trimmed",
			env: map[string]string{
				"CLAUDE_CODE_SESSION_ID": "  padded-session  ",
			},
			expectedSessionId: "padded-session",
			expectedAgent:     "claude-code",
		},
		{
			name: "a value with a newline is dropped, not sanitized",
			env: map[string]string{
				"CLAUDE_CODE_SESSION_ID": "abc\r\nX-Evil: 1",
			},
			expectedSessionId: "",
			expectedAgent:     "",
		},
		{
			name: "a non-ascii value is dropped",
			env: map[string]string{
				"CHALK_AGENT": "claude-codé",
			},
			expectedSessionId: "",
			expectedAgent:     "",
		},
		{
			name: "an over-long value is dropped rather than truncated",
			env: map[string]string{
				"CLAUDE_CODE_SESSION_ID": strings.Repeat("a", maxValueLen+1),
			},
			expectedSessionId: "",
			expectedAgent:     "",
		},
		{
			name: "an invalid override falls through to the harness",
			env: map[string]string{
				"CHALK_AGENT_SESSION_ID": "bad\nvalue",
				"CLAUDE_CODE_SESSION_ID": "good-session",
			},
			expectedSessionId: "good-session",
			expectedAgent:     "claude-code",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			attribution := Detect(ctxWithEnv(t, tc.env))
			assert.Equal(t, tc.expectedSessionId, attribution.SessionId)
			assert.Equal(t, tc.expectedAgent, attribution.Agent)
			assert.Equal(t, tc.expectedSessionId == "" && tc.expectedAgent == "", attribution.IsZero())
		})
	}
}

func TestSetHeadersOmitsUnknownFields(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name        string
		attribution Attribution
		expected    map[string]string
	}{
		{
			name:        "both known",
			attribution: Attribution{SessionId: "session", Agent: "agent"},
			expected:    map[string]string{SessionIdHeader: "session", AgentHeader: "agent"},
		},
		{
			name:        "session only",
			attribution: Attribution{SessionId: "session"},
			expected:    map[string]string{SessionIdHeader: "session"},
		},
		{
			name:        "agent only",
			attribution: Attribution{Agent: "agent"},
			expected:    map[string]string{AgentHeader: "agent"},
		},
		{
			name:        "zero value sets nothing",
			attribution: Attribution{},
			expected:    map[string]string{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			header := http.Header{}
			tc.attribution.SetHeaders(header)

			assert.Equal(t, len(tc.expected), len(header))
			for name, value := range tc.expected {
				assert.Equal(t, value, header.Get(name))
			}
		})
	}
}

func TestUserAgentAppendsAgentComment(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name        string
		attribution Attribution
		expected    string
	}{
		{
			name:        "agent",
			attribution: Attribution{SessionId: "session", Agent: "claude-code_2-1-222_agent"},
			expected:    "base/1.0 (agent=claude-code_2-1-222_agent)",
		},
		{
			name:        "no agent",
			attribution: Attribution{SessionId: "session"},
			expected:    "base/1.0",
		},
		{
			name:        "an agent that would break out of the comment is left off",
			attribution: Attribution{Agent: "evil) (x"},
			expected:    "base/1.0",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expected, tc.attribution.UserAgent("base/1.0"))
		})
	}
}
