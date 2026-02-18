package downloader

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type DownloadRequest struct {
	Id      string
	Url     string
	Retries int
}

type ProgressUpdate struct {
	Id            string
	ContentLength int64
	BytesWritten  int64
	Status        string
}

func (p ProgressUpdate) Percentage() float64 {
	if p.ContentLength <= 0 {
		return 0
	}
	percentage := float64(p.BytesWritten) / float64(p.ContentLength) * 100
	if percentage > 100 {
		percentage = 100
	}
	return percentage
}

type Downloader struct {
	outputFolder   string
	maxParallelism int
	progressChan   chan<- ProgressUpdate
	client         *http.Client
}

func New(outputFolder string, maxParallelism int, progressChan chan<- ProgressUpdate, client *http.Client) *Downloader {
	return &Downloader{
		outputFolder:   outputFolder,
		maxParallelism: maxParallelism,
		progressChan:   progressChan,
		client:         client,
	}
}

func (d *Downloader) Download(ctx context.Context, requests []DownloadRequest) error {
	sem := make(chan struct{}, d.maxParallelism)
	var wg sync.WaitGroup

	for _, req := range requests {
		wg.Add(1)
		go func(req DownloadRequest) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			d.downloadFile(ctx, req)
		}(req)
	}

	wg.Wait()
	return nil
}

func (d *Downloader) downloadFile(ctx context.Context, req DownloadRequest) {
	for attempt := 0; attempt <= req.Retries; attempt++ {
		if attempt > 0 {
			d.sendProgressUpdate(req.Id, 0, 0, fmt.Sprintf("Retrying (%d/%d)", attempt, req.Retries))
		}

		if err := d.attemptDownload(ctx, req); err != nil {
			if attempt < req.Retries {
				_ = d.deleteFile(req.Id)
				continue
			}
			d.sendProgressUpdate(req.Id, 0, 0, fmt.Sprintf("Failed: %v", err))
			return
		}

		d.sendProgressUpdate(req.Id, 1, 1, "Completed")
		return
	}
}

func (d *Downloader) attemptDownload(ctx context.Context, req DownloadRequest) error {
	d.sendProgressUpdate(req.Id, 0, 0, "Starting download")

	httpReq, err := http.NewRequestWithContext(ctx, "GET", req.Url, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	resp, err := d.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	filename := d.extractFilename(resp, req.Url)
	file, err := os.Create(filepath.Join(d.outputFolder, filename))
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer file.Close()

	contentLength := resp.ContentLength
	if contentLength <= 0 {
		contentLength = 1
	}

	var written int64
	buf := make([]byte, 8*1024) // Smaller buffer for more frequent updates
	var lastUpdateBytes int64

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		n, err := resp.Body.Read(buf)
		if n > 0 {
			if _, writeErr := file.Write(buf[:n]); writeErr != nil {
				return fmt.Errorf("writing to file: %w", writeErr)
			}
			written += int64(n)

			// Only send progress update if we've written at least 64KB since last update
			if written-lastUpdateBytes >= 64*1024 || written == int64(n) {
				d.sendProgressUpdate(req.Id, contentLength, written, "Downloading")
				lastUpdateBytes = written
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("reading response: %w", err)
		}
	}

	return nil
}

func (d *Downloader) extractFilename(resp *http.Response, url string) string {
	contentDisposition := resp.Header.Get("Content-Disposition")
	if contentDisposition != "" {
		for part := range strings.SplitSeq(contentDisposition, ";") {
			part = strings.TrimSpace(part)
			if filename, ok := strings.CutPrefix(part, "filename="); ok {
				filename = strings.Trim(filename, `"`)
				if filename != "" {
					return filename
				}
			}
		}
	}

	parts := strings.Split(url, "/")
	if len(parts) > 0 {
		filename := parts[len(parts)-1]
		if filename != "" {
			return filename
		}
	}

	return "download"
}

func (d *Downloader) deleteFile(id string) error {
	matches, err := filepath.Glob(filepath.Join(d.outputFolder, "*"))
	if err != nil {
		return err
	}

	for _, match := range matches {
		if strings.Contains(match, id) {
			if err := os.Remove(match); err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *Downloader) sendProgressUpdate(id string, contentLength, bytesWritten int64, status string) {
	if d.progressChan != nil {
		select {
		case d.progressChan <- ProgressUpdate{
			Id:            id,
			ContentLength: contentLength,
			BytesWritten:  bytesWritten,
			Status:        status,
		}:
		default:
		}
	}
}
