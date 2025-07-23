package downloader

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestDownloader(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "downloader_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	progressChan := make(chan ProgressUpdate, 100)
	client := &http.Client{Timeout: 30 * time.Second}
	downloader := New(tmpDir, 3, progressChan, client)

	requests := []DownloadRequest{
		{Id: "test1", Url: "https://httpbin.org/bytes/1024", Retries: 2},
		{Id: "test2", Url: "https://httpbin.org/bytes/2048", Retries: 2},
	}

	go func() {
		for update := range progressChan {
			fmt.Printf("Progress: %s - %.1f%% - %s\n", update.Id, update.Percentage(), update.Status)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = downloader.Download(ctx, requests)
	if err != nil {
		t.Fatal(err)
	}

	close(progressChan)

	files, err := filepath.Glob(filepath.Join(tmpDir, "*"))
	if err != nil {
		t.Fatal(err)
	}

	if len(files) != 2 {
		t.Errorf("Expected 2 files, got %d", len(files))
	}
}