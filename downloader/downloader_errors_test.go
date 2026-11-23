package downloader

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
)

type errorTestTransport func(*http.Request) (*http.Response, error)

func (f errorTestTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// TestDownloadReturnsFailures verifies outcomes through the public download API.
func TestDownloadReturnsFailures(t *testing.T) {
	for _, tc := range []struct {
		name      string
		failures  int32
		retries   int
		wantError bool
		wantCalls int32
	}{
		{"success", 0, 0, false, 1},
		{"failure without retries", 1, 0, true, 1},
		{"retry exhaustion", 3, 2, true, 3},
		{"retry recovery", 2, 2, false, 3},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var calls atomic.Int32
			progress := make(chan ProgressUpdate, 20)
			client := &http.Client{Transport: errorTestTransport(func(r *http.Request) (*http.Response, error) {
				status := http.StatusOK
				if calls.Add(1) <= tc.failures {
					status = http.StatusServiceUnavailable
				}
				return &http.Response{StatusCode: status, Status: http.StatusText(status), Header: make(http.Header), Body: io.NopCloser(strings.NewReader("synthetic")), Request: r}, nil
			})}
			err := New(t.TempDir(), 1, progress, client).Download(context.Background(), []DownloadRequest{{Id: "file-one", Url: "https://example.test/file-one", Retries: tc.retries}})
			if (err != nil) != tc.wantError {
				t.Fatalf("Download() error = %v, wantError %v", err, tc.wantError)
			}
			if calls.Load() != tc.wantCalls {
				t.Fatalf("requests = %d, want %d", calls.Load(), tc.wantCalls)
			}
			close(progress)
			terminal := ""
			for update := range progress {
				terminal = update.Status
			}
			if tc.wantError && (!strings.HasPrefix(terminal, "Failed:") || !strings.Contains(err.Error(), "file-one")) {
				t.Fatalf("failure lost: err=%v, progress=%q", err, terminal)
			}
			if !tc.wantError && terminal != "Completed" {
				t.Fatalf("terminal progress = %q", terminal)
			}
		})
	}
}

// TestDownloadPreservesTransportErrors verifies concurrent failures remain inspectable.
func TestDownloadPreservesTransportErrors(t *testing.T) {
	first := errors.New("first transport failure")
	second := errors.New("second transport failure")
	var successes atomic.Int32
	client := &http.Client{Transport: errorTestTransport(func(r *http.Request) (*http.Response, error) {
		switch r.URL.Path {
		case "/first":
			return nil, first
		case "/second":
			return nil, second
		}
		successes.Add(1)
		return &http.Response{StatusCode: http.StatusOK, Header: make(http.Header), Body: io.NopCloser(strings.NewReader("synthetic")), Request: r}, nil
	})}
	requests := []DownloadRequest{{Id: "first", Url: "https://example.test/first"}, {Id: "second", Url: "https://example.test/second"}, {Id: "good", Url: "https://example.test/good"}}
	err := New(t.TempDir(), 2, nil, client).Download(context.Background(), requests)
	if !errors.Is(err, first) || !errors.Is(err, second) {
		t.Fatalf("Download() lost causes: %v", err)
	}
	if successes.Load() != 1 {
		t.Fatal("successful request was not completed")
	}
}
