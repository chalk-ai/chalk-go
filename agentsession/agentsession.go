// Package agentsession derives coding-agent provenance -- which agent harness (Claude Code,
// Codex, ...) is driving this process, and the session the work belongs to -- from the
// environment variables those harnesses export to their subprocesses.
//
// A coding agent runs Chalk tooling as a child process, so the harness's own environment is
// already visible here; nothing has to be threaded through flags or config. Detection is
// best-effort and silent: an absent or unrecognized harness yields no attribution, the client
// sends nothing extra, and the request is indistinguishable from a human-driven one.
//
// The server does no detection of its own -- it records exactly what the client sends. This is
// therefore a client-asserted hint suitable for correlation and display, never for authorization.
//
// The Chalk CLI's pkg/agentsession applies the same rules; keep the two in step.
package agentsession

import (
	"context"
	"net/http"
	"strings"
	"sync"

	"github.com/chalk-ai/chalk-go/envfs"
)

// Headers carrying attribution to the Chalk API, following the X-Chalk-* convention.
const (
	SessionIdHeader = "X-Chalk-Agent-Session-Id"
	AgentHeader     = "X-Chalk-Agent"
)

// Environment variables Chalk defines itself.
const (
	// EnvEnabled disables attribution entirely when set to a false-y value, for callers who do
	// not want agent identity leaving the machine.
	EnvEnabled = "CHALK_AGENT_ATTRIBUTION"
	// EnvSessionId and EnvAgent let an unrecognized harness supply the values directly. They
	// also let a caller group several harness invocations under one session id.
	EnvSessionId = "CHALK_AGENT_SESSION_ID"
	EnvAgent     = "CHALK_AGENT"
)

// maxValueLen bounds what we are willing to put on the wire. Real session ids are uuid-sized;
// anything substantially longer is a misconfigured variable rather than an id.
const maxValueLen = 200

// sessionIdVars are checked in order, so an explicit CHALK_AGENT_SESSION_ID always wins over a
// harness-supplied one.
var sessionIdVars = []string{
	EnvSessionId,
	"CLAUDE_CODE_SESSION_ID", // Claude Code
	"CODEX_THREAD_ID",        // Codex CLI and desktop app
}

var agentVars = []string{
	EnvAgent,
	// Claude Code exports AI_AGENT ("claude-code_<version>_agent") specifically so that
	// subprocesses can attribute their traffic, so treat it as the cross-vendor spelling rather
	// than adding a per-harness entry for every agent that adopts it.
	"AI_AGENT",
}

// Attribution identifies the agent session driving this CLI invocation. The zero value means
// "not agent-driven, or not recognized" and contributes no headers.
type Attribution struct {
	// SessionId groups every request made during a single agent session.
	SessionId string
	// Agent names the harness, e.g. "claude-code_2-1-222_agent".
	Agent string
}

// IsZero reports whether nothing at all was detected.
func (a Attribution) IsZero() bool {
	return a.SessionId == "" && a.Agent == ""
}

// SetHeaders stamps whichever fields are known onto h. An unknown field is left off entirely
// rather than sent empty, so the server can read "header present" as "value known".
func (a Attribution) SetHeaders(header http.Header) {
	if a.SessionId != "" {
		header.Set(SessionIdHeader, a.SessionId)
	}
	if a.Agent != "" {
		header.Set(AgentHeader, a.Agent)
	}
}

// UserAgent appends the agent to base as a User-Agent comment, e.g.
// "chalk-go/v1.3.13 (agent=claude-code_2-1-222_agent)", so request logs that see only the
// User-Agent still show which harness made the call. The session id stays in SessionIdHeader:
// it is unique per session, and a User-Agent should stay low-cardinality.
func (a Attribution) UserAgent(base string) string {
	// Parentheses or a backslash would end or escape the comment early (RFC 9110 section 5.6.5).
	if a.Agent == "" || strings.ContainsAny(a.Agent, `()\`) {
		return base
	}
	return base + " (agent=" + a.Agent + ")"
}

var (
	detectOnce sync.Once
	detected   Attribution
)

// Detected is Detect memoized for the life of the process. A process cannot change its
// own environment, so re-reading it on every RPC would be pure overhead. Tests should call
// Detect directly, which reads the context's environment getter each time.
func Detected(ctx context.Context) Attribution {
	detectOnce.Do(func() {
		detected = Detect(ctx)
	})
	return detected
}

// Detect resolves attribution from the environment getter on ctx.
func Detect(ctx context.Context) Attribution {
	getter := envfs.EnvironmentGetterFromContext(ctx)
	if !enabled(getter) {
		return Attribution{}
	}
	agent := firstHeaderSafe(getter, agentVars)
	if agent == "" {
		// Some harnesses export a stable session id but no general agent-name
		// variable. Infer only from their documented, harness-specific ids.
		if _, ok := headerSafe(getter.Getenv("CLAUDE_CODE_SESSION_ID")); ok {
			agent = "claude-code"
		} else if _, ok := headerSafe(getter.Getenv("CODEX_THREAD_ID")); ok {
			agent = "codex"
		}
	}
	return Attribution{
		SessionId: firstHeaderSafe(getter, sessionIdVars),
		Agent:     agent,
	}
}

// enabled defaults to true: attribution is opt-out, because the common case is an agent that
// set nothing beyond what its harness already exports.
func enabled(getter envfs.EnvironmentGetter) bool {
	switch strings.ToLower(strings.TrimSpace(getter.Getenv(EnvEnabled))) {
	case "0", "false", "no", "off":
		return false
	default:
		return true
	}
}

func firstHeaderSafe(getter envfs.EnvironmentGetter, keys []string) string {
	for _, key := range keys {
		if value, ok := headerSafe(getter.Getenv(key)); ok {
			return value
		}
	}
	return ""
}

// headerSafe accepts a value only if it can be sent verbatim. An implausible value is dropped
// rather than repaired: a truncated or rewritten session id still looks like an id but
// correlates to nothing, which is worse than sending no attribution at all. Restricting to
// printable ASCII also rules out the CR/LF that would make this a header-injection vector.
func headerSafe(value string) (string, bool) {
	value = strings.TrimSpace(value)
	if value == "" || len(value) > maxValueLen {
		return "", false
	}
	for i := 0; i < len(value); i++ {
		if value[i] < 0x20 || value[i] > 0x7E {
			return "", false
		}
	}
	return value, true
}
