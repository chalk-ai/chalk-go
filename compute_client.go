package chalk

import (
	"context"
	"strings"
	"time"

	"connectrpc.com/connect"
	containerv1 "github.com/chalk-ai/chalk-go/gen/chalk/container/v1"
	sandboxv1 "github.com/chalk-ai/chalk-go/gen/chalk/sandbox/v1"
	"github.com/cockroachdb/errors"
)

const (
	defaultImageBuildPollInterval = 500 * time.Millisecond
	defaultImageBuildTimeout      = 10 * time.Minute
)

type buildImageOptions struct {
	targetRegistry string
	tag            string
	pollInterval   time.Duration
	timeout        time.Duration
}

// BuildImageOption configures image build submission and polling.
type BuildImageOption func(*buildImageOptions)

// WithImageTargetRegistry pushes the built image to a specific registry.
func WithImageTargetRegistry(registry string) BuildImageOption {
	return func(opts *buildImageOptions) {
		opts.targetRegistry = registry
	}
}

// WithImageTag requests a specific tag. Tagged builds are always submitted as
// new builds because the cache-oriented get-or-build RPC does not accept a tag.
func WithImageTag(tag string) BuildImageOption {
	return func(opts *buildImageOptions) {
		opts.tag = tag
	}
}

// WithImageBuildPollInterval controls how often BuildImage polls for status.
func WithImageBuildPollInterval(interval time.Duration) BuildImageOption {
	return func(opts *buildImageOptions) {
		opts.pollInterval = interval
	}
}

// WithImageBuildTimeout controls how long BuildImage waits for completion.
func WithImageBuildTimeout(timeout time.Duration) BuildImageOption {
	return func(opts *buildImageOptions) {
		opts.timeout = timeout
	}
}

func resolveBuildImageOptions(options []BuildImageOption) buildImageOptions {
	resolved := buildImageOptions{
		pollInterval: defaultImageBuildPollInterval,
		timeout:      defaultImageBuildTimeout,
	}
	for _, option := range options {
		if option != nil {
			option(&resolved)
		}
	}
	if resolved.pollInterval <= 0 {
		resolved.pollInterval = defaultImageBuildPollInterval
	}
	if resolved.timeout <= 0 {
		resolved.timeout = defaultImageBuildTimeout
	}
	return resolved
}

// ImageBuild is the status of a Chalk-managed image build.
type ImageBuild struct {
	Image   string
	BuildID string
	Status  string
	Error   string
}

func (b ImageBuild) Ready() bool {
	return strings.EqualFold(b.Status, "exists") || strings.EqualFold(b.Status, "succeeded")
}

func (b ImageBuild) Failed() bool {
	return strings.EqualFold(b.Status, "failed")
}

func (c *clientImpl) SubmitImageBuild(ctx context.Context, image Image, opts ...BuildImageOption) (ImageBuild, error) {
	resolved := resolveBuildImageOptions(opts)
	if resolved.tag != "" {
		res, err := c.customImageClient.BuildCustomImage(ctx, connect.NewRequest(&sandboxv1.BuildCustomImageRequest{
			ImageSpec:      image.toProto(),
			TargetRegistry: stringPtrOrNil(resolved.targetRegistry),
			Tag:            stringPtrOrNil(resolved.tag),
		}))
		if err != nil {
			return ImageBuild{}, errors.Wrap(err, "submitting image build")
		}
		status := "building"
		if res.Msg.GetBuildId() == "" && res.Msg.GetImage() != "" {
			status = "succeeded"
		}
		return ImageBuild{
			Image:   res.Msg.GetImage(),
			BuildID: res.Msg.GetBuildId(),
			Status:  status,
		}, nil
	}

	res, err := c.customImageClient.GetOrBuildCustomImage(ctx, connect.NewRequest(&sandboxv1.GetOrBuildCustomImageRequest{
		ImageSpec:      image.toProto(),
		TargetRegistry: stringPtrOrNil(resolved.targetRegistry),
	}))
	if err != nil {
		return ImageBuild{}, errors.Wrap(err, "submitting image build")
	}
	return ImageBuild{
		Image:   res.Msg.GetImage(),
		BuildID: res.Msg.GetBuildId(),
		Status:  res.Msg.GetStatus(),
		Error:   res.Msg.GetError(),
	}, nil
}

func (c *clientImpl) BuildImage(ctx context.Context, image Image, opts ...BuildImageOption) (ImageBuild, error) {
	resolved := resolveBuildImageOptions(opts)
	submitted, err := c.SubmitImageBuild(ctx, image, opts...)
	if err != nil {
		return ImageBuild{}, err
	}
	if submitted.Ready() {
		return submitted, nil
	}
	if submitted.Failed() {
		return ImageBuild{}, errors.Newf("image build failed: %s", submitted.Error)
	}
	if submitted.BuildID == "" {
		return ImageBuild{}, errors.New("image build did not return a build ID")
	}

	deadline := time.Now().Add(resolved.timeout)
	for {
		wait := resolved.pollInterval
		if remaining := time.Until(deadline); remaining <= 0 {
			return ImageBuild{}, errors.Newf("image build %s timed out after %s", submitted.BuildID, resolved.timeout)
		} else if remaining < wait {
			wait = remaining
		}

		timer := time.NewTimer(wait)
		select {
		case <-ctx.Done():
			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			return ImageBuild{}, ctx.Err()
		case <-timer.C:
		}

		current, err := c.GetImageBuild(ctx, submitted.BuildID)
		if err != nil {
			return ImageBuild{}, err
		}
		if current.Ready() {
			return current, nil
		}
		if current.Failed() {
			return ImageBuild{}, errors.Newf("image build %s failed: %s", submitted.BuildID, current.Error)
		}
	}
}

func (c *clientImpl) GetImageBuild(ctx context.Context, buildID string) (ImageBuild, error) {
	if buildID == "" {
		return ImageBuild{}, errors.New("build ID is required")
	}
	res, err := c.customImageClient.GetCustomImage(ctx, connect.NewRequest(&sandboxv1.GetCustomImageRequest{
		BuildId: buildID,
	}))
	if err != nil {
		return ImageBuild{}, errors.Wrap(err, "getting image build")
	}
	return ImageBuild{
		Image:   res.Msg.GetImage(),
		BuildID: res.Msg.GetBuildId(),
		Status:  res.Msg.GetStatus(),
		Error:   res.Msg.GetError(),
	}, nil
}

type runContainerOptions struct {
	buildOptions []BuildImageOption
}

// RunContainerOption configures RunContainer.
type RunContainerOption func(*runContainerOptions)

// WithRunContainerBuildOptions passes image build options used when the
// container has a declarative Image.
func WithRunContainerBuildOptions(options ...BuildImageOption) RunContainerOption {
	return func(opts *runContainerOptions) {
		opts.buildOptions = append(opts.buildOptions, options...)
	}
}

func resolveRunContainerOptions(options []RunContainerOption) runContainerOptions {
	var resolved runContainerOptions
	for _, option := range options {
		if option != nil {
			option(&resolved)
		}
	}
	return resolved
}

func (c *clientImpl) RunContainer(ctx context.Context, container Container, opts ...RunContainerOption) (ContainerInfo, error) {
	resolved := resolveRunContainerOptions(opts)
	if container.needsImageBuild() {
		build, err := c.BuildImage(ctx, *container.image, resolved.buildOptions...)
		if err != nil {
			return ContainerInfo{}, errors.Wrap(err, "building container image")
		}
		if build.Image == "" {
			return ContainerInfo{}, errors.New("image build completed without an image URI")
		}
		container = container.withBuiltImage(build.Image)
	}

	spec, err := container.toProto()
	if err != nil {
		return ContainerInfo{}, err
	}
	res, err := c.containerClient.RunContainer(ctx, connect.NewRequest(&containerv1.RunContainerRequest{
		Container: &containerv1.ContainerRequest{Spec: spec},
	}))
	if err != nil {
		return ContainerInfo{}, errors.Wrap(err, "running container")
	}
	return containerInfoFromProto(res.Msg.GetContainer()), nil
}

func (c *clientImpl) GetContainer(ctx context.Context, ref ContainerRef) (ContainerInfo, error) {
	if ref.id == "" && ref.name == "" {
		return ContainerInfo{}, errors.New("container ID or name is required")
	}
	res, err := c.containerClient.GetContainer(ctx, connect.NewRequest(ref.toGetProto()))
	if err != nil {
		return ContainerInfo{}, errors.Wrap(err, "getting container")
	}
	return containerInfoFromProto(res.Msg.GetContainer()), nil
}

func (c *clientImpl) ListContainers(ctx context.Context) ([]ContainerInfo, error) {
	res, err := c.containerClient.ListContainers(ctx, connect.NewRequest(&containerv1.ListContainersRequest{}))
	if err != nil {
		return nil, errors.Wrap(err, "listing containers")
	}
	containers := make([]ContainerInfo, 0, len(res.Msg.GetContainers()))
	for _, container := range res.Msg.GetContainers() {
		containers = append(containers, containerInfoFromProto(container))
	}
	return containers, nil
}

type stopContainerOptions struct {
	gracePeriodSeconds *int32
}

// StopContainerOption configures StopContainer.
type StopContainerOption func(*stopContainerOptions)

// WithStopGracePeriod sets the grace period before a container is terminated.
func WithStopGracePeriod(gracePeriod time.Duration) StopContainerOption {
	return func(opts *stopContainerOptions) {
		seconds := int32(gracePeriod / time.Second)
		if gracePeriod > 0 && seconds == 0 {
			seconds = 1
		}
		opts.gracePeriodSeconds = &seconds
	}
}

// WithStopGracePeriodSeconds sets the grace period before a container is
// terminated in whole seconds.
func WithStopGracePeriodSeconds(seconds int32) StopContainerOption {
	return func(opts *stopContainerOptions) {
		opts.gracePeriodSeconds = &seconds
	}
}

func resolveStopContainerOptions(options []StopContainerOption) stopContainerOptions {
	var resolved stopContainerOptions
	for _, option := range options {
		if option != nil {
			option(&resolved)
		}
	}
	return resolved
}

func (c *clientImpl) StopContainer(ctx context.Context, ref ContainerRef, opts ...StopContainerOption) (ContainerInfo, error) {
	if ref.id == "" && ref.name == "" {
		return ContainerInfo{}, errors.New("container ID or name is required")
	}
	resolved := resolveStopContainerOptions(opts)
	res, err := c.containerClient.StopContainer(ctx, connect.NewRequest(ref.toStopProto(resolved.gracePeriodSeconds)))
	if err != nil {
		return ContainerInfo{}, errors.Wrap(err, "stopping container")
	}
	return containerInfoFromProto(res.Msg.GetContainer()), nil
}

func stringPtrOrNil(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
