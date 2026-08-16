package envfs

import (
	"errors"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"
)

func TestSystemEnvironmentGetterFileOperations(t *testing.T) {
	t.Parallel()

	getter := SystemEnvironmentGetter{}
	temporary, err := getter.CreateTemp(t.TempDir(), "stream-*.bin")
	if err != nil {
		t.Fatal(err)
	}
	path := temporary.Name()
	defer func() { _ = getter.Remove(path) }()

	if _, err := temporary.Write([]byte("payload")); err != nil {
		t.Fatal(err)
	}
	info, err := temporary.Stat()
	if err != nil {
		t.Fatal(err)
	}
	if info.Size() != int64(len("payload")) {
		t.Fatalf("temporary size = %d, want %d", info.Size(), len("payload"))
	}
	if err := temporary.Close(); err != nil {
		t.Fatal(err)
	}

	opened, err := getter.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = opened.Close() }()
	section := make([]byte, 3)
	if _, err := opened.ReadAt(section, 1); err != nil {
		t.Fatal(err)
	}
	if string(section) != "ayl" {
		t.Fatalf("section = %q, want %q", section, "ayl")
	}
	if _, err := opened.Seek(0, io.SeekStart); err != nil {
		t.Fatal(err)
	}
	contents, err := io.ReadAll(opened)
	if err != nil {
		t.Fatal(err)
	}
	if string(contents) != "payload" {
		t.Fatalf("contents = %q, want %q", contents, "payload")
	}
	if err := opened.Close(); err != nil {
		t.Fatal(err)
	}
	if err := getter.Remove(path); err != nil {
		t.Fatal(err)
	}
	if _, err := getter.Stat(path); !errors.Is(err, fs.ErrNotExist) {
		t.Fatalf("stat removed file error = %v, want fs.ErrNotExist", err)
	}
}

func TestMockEnvironmentGetterOpenStreamsFiles(t *testing.T) {
	t.Parallel()

	const path = "/input/data.parquet"
	getter := NewMockEnvironmentGetter(WithFileString(path, "payload"))
	opened, err := getter.Open(path)
	if err != nil {
		t.Fatal(err)
	}

	prefix := make([]byte, 3)
	if _, err := opened.Read(prefix); err != nil {
		t.Fatal(err)
	}
	if string(prefix) != "pay" {
		t.Fatalf("prefix = %q, want %q", prefix, "pay")
	}
	section := make([]byte, 3)
	if _, err := opened.ReadAt(section, 4); err != nil {
		t.Fatal(err)
	}
	if string(section) != "oad" {
		t.Fatalf("section = %q, want %q", section, "oad")
	}
	if _, err := opened.Seek(-2, io.SeekEnd); err != nil {
		t.Fatal(err)
	}
	suffix, err := io.ReadAll(opened)
	if err != nil {
		t.Fatal(err)
	}
	if string(suffix) != "ad" {
		t.Fatalf("suffix = %q, want %q", suffix, "ad")
	}
	if err := opened.Close(); err != nil {
		t.Fatal(err)
	}
	if _, err := opened.Read(prefix); !errors.Is(err, fs.ErrClosed) {
		t.Fatalf("read after close error = %v, want fs.ErrClosed", err)
	}
}

func TestMockEnvironmentGetterTempFileLifecycle(t *testing.T) {
	t.Parallel()

	temporaryDirectory := filepath.Join(string(filepath.Separator), "mock-tmp")
	getter := NewMockEnvironmentGetter(WithEnv(map[string]string{"TMPDIR": temporaryDirectory}))
	first, err := getter.CreateTemp("", "upload-*.parquet")
	if err != nil {
		t.Fatal(err)
	}
	second, err := getter.CreateTemp("", "upload-*.parquet")
	if err != nil {
		t.Fatal(err)
	}
	if first.Name() == second.Name() {
		t.Fatalf("temporary paths are not unique: %s", first.Name())
	}
	if filepath.Dir(first.Name()) != temporaryDirectory ||
		!strings.HasPrefix(filepath.Base(first.Name()), "upload-") ||
		filepath.Ext(first.Name()) != ".parquet" {
		t.Fatalf("unexpected temporary path %q", first.Name())
	}
	if err := second.Close(); err != nil {
		t.Fatal(err)
	}

	if _, err := first.Write([]byte("payload")); err != nil {
		t.Fatal(err)
	}
	info, err := first.Stat()
	if err != nil {
		t.Fatal(err)
	}
	if info.Size() != int64(len("payload")) {
		t.Fatalf("temporary size = %d, want %d", info.Size(), len("payload"))
	}
	if info.Mode().Perm() != 0o600 {
		t.Fatalf("temporary mode = %o, want 600", info.Mode().Perm())
	}
	path := first.Name()
	if err := first.Close(); err != nil {
		t.Fatal(err)
	}
	contents, err := getter.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(contents) != "payload" {
		t.Fatalf("contents = %q, want %q", contents, "payload")
	}
	if err := getter.Remove(path); err != nil {
		t.Fatal(err)
	}
	if _, err := getter.Open(path); !errors.Is(err, fs.ErrNotExist) {
		t.Fatalf("open removed file error = %v, want fs.ErrNotExist", err)
	}
}

func TestEnvironmentGetterImplementationsExposeFilesystem(t *testing.T) {
	t.Parallel()

	var _ EnvironmentGetter = SystemEnvironmentGetter{}
	var _ EnvironmentGetter = NewMockEnvironmentGetter()
	var _ EnvironmentGetter = NewDirectoryOverrideEnvironmentGetter(NewMockEnvironmentGetter(), "/workspace")
}
