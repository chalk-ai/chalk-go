package chalk

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	sandboxv1 "github.com/chalk-ai/chalk-go/gen/chalk/sandbox/v1"
	"github.com/cockroachdb/errors"
)

type imageStepKind int

const (
	imageStepRunCommands imageStepKind = iota + 1
	imageStepPipInstall
	imageStepUvPipInstall
	imageStepAddFile
	imageStepDockerfileCommands
)

type imageStep struct {
	kind        imageStepKind
	commands    []string
	packages    []string
	destination string
	content     []byte
	mode        *uint32
}

// Image is a declarative, immutable container image definition.
//
// Builder methods return a new Image, so a base image can be safely reused:
//
//	base := chalk.DebianSlimImage().PipInstall("requests")
//	api := base.PipInstall("flask").Workdir("/app")
//	worker := base.PipInstall("celery").Workdir("/worker")
type Image struct {
	baseImage  string
	steps      []imageStep
	entrypoint []string
	cmd        []string
	workdir    *string
	env        map[string]string
}

// BaseImage starts an image from an arbitrary image reference, such as
// "python:3.14-slim-trixie" or "ghcr.io/acme/service:latest".
func BaseImage(image string) Image {
	return Image{baseImage: image}
}

// PythonSlimImage starts from python:<version>-slim-trixie.
// If no version is supplied, Python 3.14 is used.
func PythonSlimImage(version ...string) Image {
	pythonVersion := "3.14"
	if len(version) > 0 && version[0] != "" {
		pythonVersion = version[0]
	}
	return BaseImage(fmt.Sprintf("python:%s-slim-trixie", pythonVersion))
}

// DebianSlimImage is a convenience alias for PythonSlimImage.
func DebianSlimImage(version ...string) Image {
	return PythonSlimImage(version...)
}

// ImageFromDockerfile creates an Image from a Dockerfile and leaves it
// chainable. The last FROM line becomes the base image; remaining non-empty,
// non-comment lines are sent as Dockerfile commands.
func ImageFromDockerfile(path string) (Image, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Image{}, errors.Wrap(err, "reading Dockerfile")
	}

	base := "scratch"
	var commands []string
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(strings.ToUpper(line), "FROM ") {
			base = strings.TrimSpace(line[len("FROM "):])
			continue
		}
		commands = append(commands, line)
	}

	image := BaseImage(base)
	if len(commands) > 0 {
		image = image.DockerfileCommands(commands...)
	}
	return image, nil
}

func (i Image) clone() Image {
	next := Image{
		baseImage:  i.baseImage,
		steps:      make([]imageStep, len(i.steps)),
		entrypoint: append([]string(nil), i.entrypoint...),
		cmd:        append([]string(nil), i.cmd...),
		env:        cloneStringMap(i.env),
	}
	if i.workdir != nil {
		workdir := *i.workdir
		next.workdir = &workdir
	}
	for idx, step := range i.steps {
		next.steps[idx] = step.clone()
	}
	return next
}

func (s imageStep) clone() imageStep {
	next := s
	next.commands = append([]string(nil), s.commands...)
	next.packages = append([]string(nil), s.packages...)
	next.content = append([]byte(nil), s.content...)
	if s.mode != nil {
		mode := *s.mode
		next.mode = &mode
	}
	return next
}

// RunCommands appends shell commands to execute during the image build. Each
// command becomes a separate build instruction.
func (i Image) RunCommands(commands ...string) Image {
	next := i.clone()
	next.steps = append(next.steps, imageStep{
		kind:     imageStepRunCommands,
		commands: append([]string(nil), commands...),
	})
	return next
}

// PipInstall appends a pip install step.
func (i Image) PipInstall(packages ...string) Image {
	next := i.clone()
	next.steps = append(next.steps, imageStep{
		kind:     imageStepPipInstall,
		packages: append([]string(nil), packages...),
	})
	return next
}

// PipInstallFromRequirements reads a requirements.txt-style file and appends
// a pip install step for its non-empty, non-comment lines.
func (i Image) PipInstallFromRequirements(path string) (Image, error) {
	packages, err := readRequirements(path)
	if err != nil {
		return Image{}, err
	}
	return i.PipInstall(packages...), nil
}

// UvPipInstall appends a uv pip install step. The base image must include uv.
func (i Image) UvPipInstall(packages ...string) Image {
	next := i.clone()
	next.steps = append(next.steps, imageStep{
		kind:     imageStepUvPipInstall,
		packages: append([]string(nil), packages...),
	})
	return next
}

// UvPipInstallFromRequirements reads a requirements.txt-style file and appends
// a uv pip install step for its non-empty, non-comment lines.
func (i Image) UvPipInstallFromRequirements(path string) (Image, error) {
	packages, err := readRequirements(path)
	if err != nil {
		return Image{}, err
	}
	return i.UvPipInstall(packages...), nil
}

type addFileOptions struct {
	mode *uint32
}

// FileOption configures files added to an Image.
type FileOption func(*addFileOptions)

// FileMode sets the POSIX file mode for a file added to an Image.
func FileMode(mode uint32) FileOption {
	return func(opts *addFileOptions) {
		opts.mode = &mode
	}
}

func applyFileOptions(options []FileOption) addFileOptions {
	var opts addFileOptions
	for _, option := range options {
		if option != nil {
			option(&opts)
		}
	}
	return opts
}

// AddFile adds in-memory file content to the image at destination.
func (i Image) AddFile(destination string, content []byte, options ...FileOption) Image {
	opts := applyFileOptions(options)
	next := i.clone()
	next.steps = append(next.steps, imageStep{
		kind:        imageStepAddFile,
		destination: destination,
		content:     append([]byte(nil), content...),
		mode:        opts.mode,
	})
	return next
}

// AddLocalFile reads a local file and adds its content to the image at
// destination.
func (i Image) AddLocalFile(path string, destination string, options ...FileOption) (Image, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return Image{}, errors.Wrap(err, "reading local file")
	}
	return i.AddFile(destination, content, options...), nil
}

// AddLocalDir walks a local directory and adds each regular file under
// destination. The .git directory and .chalkignore files are skipped.
func (i Image) AddLocalDir(path string, destination string) (Image, error) {
	var files []string
	err := filepath.WalkDir(path, func(filePath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() && entry.Name() == ".git" {
			return filepath.SkipDir
		}
		if entry.IsDir() {
			return nil
		}
		if entry.Name() == ".chalkignore" {
			return nil
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		files = append(files, filePath)
		return nil
	})
	if err != nil {
		return Image{}, errors.Wrap(err, "walking local directory")
	}

	sort.Strings(files)
	next := i.clone()
	for _, filePath := range files {
		rel, err := filepath.Rel(path, filePath)
		if err != nil {
			return Image{}, errors.Wrap(err, "resolving relative file path")
		}
		content, err := os.ReadFile(filePath)
		if err != nil {
			return Image{}, errors.Wrap(err, "reading local directory file")
		}
		target := filepath.ToSlash(filepath.Join(destination, rel))
		next = next.AddFile(target, content)
	}
	return next, nil
}

// DockerfileCommands appends raw Dockerfile instructions to the image build.
func (i Image) DockerfileCommands(commands ...string) Image {
	next := i.clone()
	next.steps = append(next.steps, imageStep{
		kind:     imageStepDockerfileCommands,
		commands: append([]string(nil), commands...),
	})
	return next
}

// Env applies environment variables baked into the image.
func (i Image) Env(vars map[string]string) Image {
	next := i.clone()
	if next.env == nil {
		next.env = map[string]string{}
	}
	for key, value := range vars {
		next.env[key] = value
	}
	return next
}

// EnvVar sets one environment variable baked into the image.
func (i Image) EnvVar(key string, value string) Image {
	return i.Env(map[string]string{key: value})
}

// Workdir sets the final image working directory.
func (i Image) Workdir(path string) Image {
	next := i.clone()
	next.workdir = &path
	return next
}

// Entrypoint sets the image entrypoint in exec form.
func (i Image) Entrypoint(args ...string) Image {
	next := i.clone()
	next.entrypoint = append([]string(nil), args...)
	return next
}

// Cmd sets the image CMD in exec form.
func (i Image) Cmd(args ...string) Image {
	next := i.clone()
	next.cmd = append([]string(nil), args...)
	return next
}

// Base returns the base image reference.
func (i Image) Base() string {
	return i.baseImage
}

// String returns a concise description of the image definition.
func (i Image) String() string {
	parts := []string{fmt.Sprintf("Image(base=%q", i.baseImage)}
	if len(i.steps) > 0 {
		parts = append(parts, fmt.Sprintf(", steps=%d", len(i.steps)))
	}
	if i.workdir != nil {
		parts = append(parts, fmt.Sprintf(", workdir=%q", *i.workdir))
	}
	if len(i.env) > 0 {
		parts = append(parts, fmt.Sprintf(", env=%d", len(i.env)))
	}
	parts = append(parts, ")")
	return strings.Join(parts, "")
}

func (i Image) toProto() *sandboxv1.ImageSpec {
	spec := &sandboxv1.ImageSpec{
		BaseImage:  i.baseImage,
		Steps:      make([]*sandboxv1.BuildStep, 0, len(i.steps)),
		Entrypoint: append([]string(nil), i.entrypoint...),
		Cmd:        append([]string(nil), i.cmd...),
		Env:        cloneStringMap(i.env),
	}
	if i.workdir != nil {
		workdir := *i.workdir
		spec.Workdir = &workdir
	}
	for _, step := range i.steps {
		spec.Steps = append(spec.Steps, step.toProto())
	}
	return spec
}

func (s imageStep) toProto() *sandboxv1.BuildStep {
	switch s.kind {
	case imageStepRunCommands:
		return &sandboxv1.BuildStep{
			Step: &sandboxv1.BuildStep_RunCommands{
				RunCommands: &sandboxv1.RunCommandsStep{Commands: append([]string(nil), s.commands...)},
			},
		}
	case imageStepPipInstall:
		return &sandboxv1.BuildStep{
			Step: &sandboxv1.BuildStep_PipInstall{
				PipInstall: &sandboxv1.PipInstallStep{Packages: append([]string(nil), s.packages...)},
			},
		}
	case imageStepUvPipInstall:
		return &sandboxv1.BuildStep{
			Step: &sandboxv1.BuildStep_UvPipInstall{
				UvPipInstall: &sandboxv1.UvPipInstallStep{Packages: append([]string(nil), s.packages...)},
			},
		}
	case imageStepAddFile:
		addFile := &sandboxv1.AddFileStep{
			Destination: s.destination,
			Content:     append([]byte(nil), s.content...),
		}
		if s.mode != nil {
			mode := *s.mode
			addFile.Mode = &mode
		}
		return &sandboxv1.BuildStep{
			Step: &sandboxv1.BuildStep_AddFile{AddFile: addFile},
		}
	case imageStepDockerfileCommands:
		return &sandboxv1.BuildStep{
			Step: &sandboxv1.BuildStep_DockerfileCommands{
				DockerfileCommands: &sandboxv1.DockerfileCommandsStep{Commands: append([]string(nil), s.commands...)},
			},
		}
	default:
		return &sandboxv1.BuildStep{}
	}
}

func readRequirements(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "opening requirements file")
	}
	defer file.Close()

	var packages []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		packages = append(packages, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "reading requirements file")
	}
	return packages, nil
}

func cloneStringMap(input map[string]string) map[string]string {
	if input == nil {
		return nil
	}
	output := make(map[string]string, len(input))
	for key, value := range input {
		output[key] = value
	}
	return output
}
