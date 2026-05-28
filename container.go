package chalk

import (
	"time"

	containerv1 "github.com/chalk-ai/chalk-go/gen/chalk/container/v1"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"github.com/cockroachdb/errors"
	"google.golang.org/protobuf/types/known/durationpb"
)

type VolumeKind string

const (
	VolumeKindEmptyDir     VolumeKind = "empty_dir"
	VolumeKindSharedMemory VolumeKind = "shared_memory"
	VolumeKindChalkFS      VolumeKind = "chalkfs"
)

// VolumeMount describes storage mounted into a container.
type VolumeMount struct {
	name      string
	mountPath string
	kind      VolumeKind
	sizeLimit string
}

// EmptyDirVolume mounts scratch storage into a container.
func EmptyDirVolume(name string, mountPath string, sizeLimit string) VolumeMount {
	return VolumeMount{name: name, mountPath: mountPath, kind: VolumeKindEmptyDir, sizeLimit: sizeLimit}
}

// SharedMemoryVolume mounts shared memory storage into a container.
func SharedMemoryVolume(name string, mountPath string, sizeLimit string) VolumeMount {
	return VolumeMount{name: name, mountPath: mountPath, kind: VolumeKindSharedMemory, sizeLimit: sizeLimit}
}

// ChalkFSVolume mounts an existing ChalkFS volume into a container.
func ChalkFSVolume(name string, mountPath string) VolumeMount {
	return VolumeMount{name: name, mountPath: mountPath, kind: VolumeKindChalkFS}
}

func (v VolumeMount) Name() string {
	return v.name
}

func (v VolumeMount) MountPath() string {
	return v.mountPath
}

func (v VolumeMount) Kind() VolumeKind {
	return v.kind
}

func (v VolumeMount) SizeLimit() string {
	return v.sizeLimit
}

func (v VolumeMount) toProto() *containerv1.VolumeMount {
	return &containerv1.VolumeMount{
		Name:      v.name,
		MountPath: v.mountPath,
		Type:      string(v.kind),
		SizeLimit: ptr.OrNil(v.sizeLimit),
	}
}

type ContainerProtocol string

const (
	ContainerProtocolHTTP ContainerProtocol = "http"
	ContainerProtocolGRPC ContainerProtocol = "grpc"
)

type ContainerRouting string

const (
	ContainerRoutingPublic  ContainerRouting = "PUBLIC"
	ContainerRoutingPrivate ContainerRouting = "PRIVATE"
)

type ContainerAuthentication string

const (
	ContainerAuthenticationUnauthenticated ContainerAuthentication = "UNAUTHENTICATED"
	ContainerAuthenticationAuthenticated   ContainerAuthentication = "AUTHENTICATED"
)

type ResourceLimits struct {
	CPU    string
	Memory string
	GPU    string
}

func (r ResourceLimits) isZero() bool {
	return r.CPU == "" && r.Memory == "" && r.GPU == ""
}

func (r ResourceLimits) toProto() *containerv1.ResourceLimits {
	if r.isZero() {
		return nil
	}
	return &containerv1.ResourceLimits{
		Cpu:    ptr.OrNil(r.CPU),
		Memory: ptr.OrNil(r.Memory),
		Gpu:    ptr.OrNil(r.GPU),
	}
}

// Container is an immutable specification for a managed container.
//
// Use NewContainer for declarative images that should be built by Chalk, or
// NewContainerFromImageRef when you already have an image URI.
type Container struct {
	name           string
	image          *Image
	imageRef       string
	entrypoint     []string
	tags           map[string]string
	port           *int32
	lifetime       *time.Duration
	resources      ResourceLimits
	enableSSH      *bool
	env            map[string]string
	volumes        []VolumeMount
	protocol       *ContainerProtocol
	routing        *ContainerRouting
	authentication *ContainerAuthentication
}

// NewContainer creates a container that runs a declarative Image. RunContainer
// builds the image before submitting the container request.
func NewContainer(name string, image Image) Container {
	imageCopy := image.clone()
	return Container{name: name, image: &imageCopy}
}

// NewContainerFromImageRef creates a container from an existing image URI.
func NewContainerFromImageRef(name string, imageRef string) Container {
	return Container{name: name, imageRef: imageRef}
}

func (c Container) clone() Container {
	next := c
	if c.image != nil {
		image := c.image.clone()
		next.image = &image
	}
	next.entrypoint = append([]string(nil), c.entrypoint...)
	next.tags = cloneStringMap(c.tags)
	next.env = cloneStringMap(c.env)
	next.volumes = append([]VolumeMount(nil), c.volumes...)
	if c.port != nil {
		port := *c.port
		next.port = &port
	}
	if c.lifetime != nil {
		lifetime := *c.lifetime
		next.lifetime = &lifetime
	}
	if c.enableSSH != nil {
		enableSSH := *c.enableSSH
		next.enableSSH = &enableSSH
	}
	if c.protocol != nil {
		protocol := *c.protocol
		next.protocol = &protocol
	}
	if c.routing != nil {
		routing := *c.routing
		next.routing = &routing
	}
	if c.authentication != nil {
		authentication := *c.authentication
		next.authentication = &authentication
	}
	return next
}

func (c Container) Name() string {
	return c.name
}

func (c Container) ImageRef() string {
	return c.imageRef
}

func (c Container) WithName(name string) Container {
	next := c.clone()
	next.name = name
	return next
}

func (c Container) WithImage(image Image) Container {
	next := c.clone()
	imageCopy := image.clone()
	next.image = &imageCopy
	next.imageRef = ""
	return next
}

func (c Container) WithImageRef(imageRef string) Container {
	next := c.clone()
	next.image = nil
	next.imageRef = imageRef
	return next
}

func (c Container) WithEntrypoint(args ...string) Container {
	next := c.clone()
	next.entrypoint = append([]string(nil), args...)
	return next
}

func (c Container) WithTag(key string, value string) Container {
	return c.WithTags(map[string]string{key: value})
}

func (c Container) WithTags(tags map[string]string) Container {
	next := c.clone()
	if next.tags == nil {
		next.tags = map[string]string{}
	}
	for key, value := range tags {
		next.tags[key] = value
	}
	return next
}

func (c Container) WithPort(port int32) Container {
	next := c.clone()
	next.port = &port
	return next
}

func (c Container) WithLifetime(lifetime time.Duration) Container {
	next := c.clone()
	next.lifetime = &lifetime
	return next
}

func (c Container) WithResources(resources ResourceLimits) Container {
	next := c.clone()
	next.resources = resources
	return next
}

func (c Container) WithCPU(cpu string) Container {
	next := c.clone()
	next.resources.CPU = cpu
	return next
}

func (c Container) WithMemory(memory string) Container {
	next := c.clone()
	next.resources.Memory = memory
	return next
}

func (c Container) WithGPU(gpu string) Container {
	next := c.clone()
	next.resources.GPU = gpu
	return next
}

func (c Container) WithSSH(enabled bool) Container {
	next := c.clone()
	next.enableSSH = &enabled
	return next
}

func (c Container) WithEnv(key string, value string) Container {
	return c.WithEnvVars(map[string]string{key: value})
}

func (c Container) WithEnvVars(env map[string]string) Container {
	next := c.clone()
	if next.env == nil {
		next.env = map[string]string{}
	}
	for key, value := range env {
		next.env[key] = value
	}
	return next
}

func (c Container) WithVolume(volume VolumeMount) Container {
	next := c.clone()
	next.volumes = append(next.volumes, volume)
	return next
}

func (c Container) WithProtocol(protocol ContainerProtocol) Container {
	next := c.clone()
	next.protocol = &protocol
	return next
}

func (c Container) WithRouting(routing ContainerRouting) Container {
	next := c.clone()
	next.routing = &routing
	return next
}

func (c Container) WithAuthentication(authentication ContainerAuthentication) Container {
	next := c.clone()
	next.authentication = &authentication
	return next
}

func (c Container) needsImageBuild() bool {
	return c.image != nil
}

func (c Container) withBuiltImage(imageRef string) Container {
	next := c.clone()
	next.image = nil
	next.imageRef = imageRef
	return next
}

func (c Container) toProto() (*containerv1.ChalkContainerSpec, error) {
	if c.image != nil {
		return nil, errors.New("container image has not been built")
	}
	if c.imageRef == "" {
		return nil, errors.New("container image is required")
	}

	spec := &containerv1.ChalkContainerSpec{
		Name:       c.name,
		Image:      c.imageRef,
		Entrypoint: append([]string(nil), c.entrypoint...),
		Tags:       cloneStringMap(c.tags),
		Resources:  c.resources.toProto(),
		EnableSsh:  c.enableSSH,
		EnvVars:    cloneStringMap(c.env),
		Volumes:    make([]*containerv1.VolumeMount, 0, len(c.volumes)),
	}
	for _, volume := range c.volumes {
		spec.Volumes = append(spec.Volumes, volume.toProto())
	}
	if c.port != nil {
		port := *c.port
		spec.Port = &port
	}
	if c.lifetime != nil {
		spec.Lifetime = durationpb.New(*c.lifetime)
	}
	if c.protocol != nil {
		protocol := string(*c.protocol)
		spec.Protocol = &protocol
	}
	if c.routing != nil {
		routing := string(*c.routing)
		spec.Routing = &routing
	}
	if c.authentication != nil {
		authentication := string(*c.authentication)
		spec.Authentication = &authentication
	}
	return spec, nil
}

type ContainerRef struct {
	id   string
	name string
}

func ContainerID(id string) ContainerRef {
	return ContainerRef{id: id}
}

func ContainerName(name string) ContainerRef {
	return ContainerRef{name: name}
}

func (r ContainerRef) toGetProto() *containerv1.GetContainerRequest {
	return &containerv1.GetContainerRequest{
		Id:   ptr.OrNil(r.id),
		Name: ptr.OrNil(r.name),
	}
}

func (r ContainerRef) toStopProto(gracePeriodSeconds *int32) *containerv1.StopContainerRequest {
	return &containerv1.StopContainerRequest{
		Id:                 ptr.OrNil(r.id),
		Name:               ptr.OrNil(r.name),
		GracePeriodSeconds: gracePeriodSeconds,
	}
}

type HealthCheck struct {
	Healthy    bool
	StatusCode *int32
	Error      string
}

type ContainerInfo struct {
	ID            string
	Name          string
	Status        string
	StatusMessage string
	Spec          Container
	CreatedAt     *time.Time
	StoppedAt     *time.Time
	PodName       string
	WebURL        string
	SSHPrivateKey string
	SSHUsername   string
	SSHHost       string
	SSHPort       *int32
	HealthCheck   *HealthCheck
}

func containerInfoFromProto(proto *containerv1.ContainerResponse) ContainerInfo {
	if proto == nil {
		return ContainerInfo{}
	}
	info := ContainerInfo{
		ID:            proto.GetId(),
		Name:          proto.GetName(),
		Status:        proto.GetStatus(),
		StatusMessage: proto.GetStatusMessage(),
		Spec:          containerFromProto(proto.GetSpec()),
		PodName:       proto.GetPodName(),
		WebURL:        proto.GetWebUrl(),
		SSHPrivateKey: proto.GetSshPrivateKey(),
		SSHUsername:   proto.GetSshUsername(),
		SSHHost:       proto.GetSshHost(),
		SSHPort:       proto.SshPort,
	}
	if proto.GetCreatedAt() != nil {
		createdAt := proto.GetCreatedAt().AsTime()
		info.CreatedAt = &createdAt
	}
	if proto.GetStoppedAt() != nil {
		stoppedAt := proto.GetStoppedAt().AsTime()
		info.StoppedAt = &stoppedAt
	}
	if proto.GetHealthCheck() != nil {
		info.HealthCheck = &HealthCheck{
			Healthy:    proto.GetHealthCheck().GetHealthy(),
			StatusCode: proto.GetHealthCheck().StatusCode,
			Error:      proto.GetHealthCheck().GetError(),
		}
	}
	return info
}

func containerFromProto(proto *containerv1.ChalkContainerSpec) Container {
	if proto == nil {
		return Container{}
	}
	container := NewContainerFromImageRef(proto.GetName(), proto.GetImage()).
		WithEntrypoint(proto.GetEntrypoint()...).
		WithTags(proto.GetTags()).
		WithEnvVars(proto.GetEnvVars())

	if proto.Port != nil {
		container = container.WithPort(proto.GetPort())
	}
	if proto.Lifetime != nil {
		container = container.WithLifetime(proto.GetLifetime().AsDuration())
	}
	if proto.Resources != nil {
		container = container.WithResources(ResourceLimits{
			CPU:    proto.GetResources().GetCpu(),
			Memory: proto.GetResources().GetMemory(),
			GPU:    proto.GetResources().GetGpu(),
		})
	}
	if proto.EnableSsh != nil {
		container = container.WithSSH(proto.GetEnableSsh())
	}
	for _, volume := range proto.GetVolumes() {
		container = container.WithVolume(VolumeMount{
			name:      volume.GetName(),
			mountPath: volume.GetMountPath(),
			kind:      VolumeKind(volume.GetType()),
			sizeLimit: volume.GetSizeLimit(),
		})
	}
	if proto.Protocol != nil {
		container = container.WithProtocol(ContainerProtocol(proto.GetProtocol()))
	}
	if proto.Routing != nil {
		container = container.WithRouting(ContainerRouting(proto.GetRouting()))
	}
	if proto.Authentication != nil {
		container = container.WithAuthentication(ContainerAuthentication(proto.GetAuthentication()))
	}
	return container
}
