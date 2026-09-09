package chalk

import (
	"context"
	"slices"
	"strings"
	"time"

	"connectrpc.com/connect"
	serverv1 "github.com/chalk-ai/chalk-go/gen/chalk/server/v1"
	"github.com/cockroachdb/errors"
)

const integrationKindPrefix = "INTEGRATION_KIND_"

// IntegrationParams describes a data source to create or update. It mirrors the
// YAML accepted by the `chalk integration apply` CLI command.
type IntegrationParams struct {
	// Name identifies the integration within the environment. Required.
	Name string

	// Kind is the data source type, such as "postgresql", "snowflake", or
	// "bigquery". The fully qualified proto name ("INTEGRATION_KIND_POSTGRESQL")
	// is also accepted. Required.
	Kind string

	// Variables holds the connection settings and credentials, keyed by the
	// variable name the data source expects, e.g. "PGHOST" or "PGPASSWORD".
	// Values are stored as managed secrets in the environment's secret store.
	Variables map[string]string
}

// Integration is a Chalk data source: a named, environment-scoped set of
// credentials that resolvers use to reach an external system.
type Integration struct {
	ID   string
	Name string

	// Kind is the lowercase data source type, e.g. "postgresql". It round-trips
	// with IntegrationParams.Kind.
	Kind string

	EnvironmentID string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// ApplyIntegrationResult is the outcome of Client.ApplyIntegration.
type ApplyIntegrationResult struct {
	Integration Integration

	// Created is true when no integration with the given name existed and one
	// was created, and false when an existing integration was updated.
	Created bool
}

func (c *clientImpl) CreateIntegration(ctx context.Context, params IntegrationParams) (Integration, error) {
	kind, err := params.resolveKind()
	if err != nil {
		return Integration{}, err
	}

	res, err := c.integrationsClient.InsertIntegration(ctx, connect.NewRequest(&serverv1.InsertIntegrationRequest{
		Name:            params.Name,
		IntegrationKind: kind,
		Config:          integrationConfigFromVariables(params.Variables),
	}))
	if err != nil {
		return Integration{}, errors.Wrap(err, "creating integration")
	}
	return integrationFromProto(res.Msg.GetIntegration()), nil
}

func (c *clientImpl) ApplyIntegration(ctx context.Context, params IntegrationParams) (ApplyIntegrationResult, error) {
	kind, err := params.resolveKind()
	if err != nil {
		return ApplyIntegrationResult{}, err
	}

	existing, err := c.integrationByName(ctx, params.Name)
	if err != nil {
		return ApplyIntegrationResult{}, err
	}
	if existing == nil {
		created, err := c.CreateIntegration(ctx, params)
		if err != nil {
			return ApplyIntegrationResult{}, err
		}
		return ApplyIntegrationResult{Integration: created, Created: true}, nil
	}

	// An update cannot change an integration's kind, so applying a different
	// kind over an existing name would silently keep the stored one.
	if existing.GetKind() != kind {
		return ApplyIntegrationResult{}, errors.Newf(
			"integration '%s' already exists with kind '%s'; delete it before applying kind '%s'",
			params.Name,
			integrationKindNameFromProtoName(existing.GetKind().String()),
			integrationKindNameFromProtoName(kind.String()),
		)
	}

	updated, err := c.integrationsClient.UpdateIntegration(ctx, connect.NewRequest(&serverv1.UpdateIntegrationRequest{
		Name:          params.Name,
		IntegrationId: existing.GetId(),
		Config:        integrationConfigFromVariables(params.Variables),
	}))
	if err != nil {
		return ApplyIntegrationResult{}, errors.Wrapf(err, "updating integration '%s'", params.Name)
	}
	return ApplyIntegrationResult{Integration: integrationFromProto(updated.Msg.GetIntegration())}, nil
}

func (c *clientImpl) DeleteIntegrationByName(ctx context.Context, name string) error {
	if name == "" {
		return errors.New("integration name is required")
	}

	existing, err := c.integrationByName(ctx, name)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.Newf("integration '%s' not found", name)
	}

	_, err = c.integrationsClient.DeleteIntegration(ctx, connect.NewRequest(&serverv1.DeleteIntegrationRequest{
		Id: existing.GetId(),
	}))
	return errors.Wrapf(err, "deleting integration '%s'", name)
}

// integrationByName returns nil without an error when the active environment has
// no integration with the given name.
func (c *clientImpl) integrationByName(ctx context.Context, name string) (*serverv1.Integration, error) {
	res, err := c.integrationsClient.GetIntegrationByName(ctx, connect.NewRequest(&serverv1.GetIntegrationByNameRequest{
		IntegrationName: name,
	}))
	if err != nil {
		return nil, errors.Wrapf(err, "looking up integration '%s'", name)
	}
	return res.Msg.GetIntegration(), nil
}

// resolveKind checks the required fields and translates Kind, which accepts
// either a bare kind ("postgresql") or the fully qualified proto name
// ("INTEGRATION_KIND_POSTGRESQL").
func (p IntegrationParams) resolveKind() (serverv1.IntegrationKind, error) {
	unspecified := serverv1.IntegrationKind_INTEGRATION_KIND_UNSPECIFIED
	if p.Name == "" {
		return unspecified, errors.New("integration name is required")
	}
	if p.Kind == "" {
		return unspecified, errors.New("integration kind is required")
	}
	if value, ok := serverv1.IntegrationKind_value[p.Kind]; ok {
		return serverv1.IntegrationKind(value), nil
	}
	prefixed := integrationKindPrefix + strings.ToUpper(p.Kind)
	if value, ok := serverv1.IntegrationKind_value[prefixed]; ok {
		return serverv1.IntegrationKind(value), nil
	}
	return unspecified, errors.Newf(
		"unknown integration kind '%s'; valid kinds: %s",
		p.Kind,
		strings.Join(validIntegrationKinds(), ", "),
	)
}

func validIntegrationKinds() []string {
	kinds := make([]string, 0, len(serverv1.IntegrationKind_value))
	for name := range serverv1.IntegrationKind_value {
		if name == integrationKindPrefix+"UNSPECIFIED" {
			continue
		}
		kinds = append(kinds, integrationKindNameFromProtoName(name))
	}
	slices.Sort(kinds)
	return kinds
}

func integrationKindNameFromProtoName(name string) string {
	return strings.ToLower(strings.TrimPrefix(name, integrationKindPrefix))
}

func integrationConfigFromVariables(variables map[string]string) map[string]*serverv1.IntegrationConfigValue {
	if len(variables) == 0 {
		return nil
	}
	config := make(map[string]*serverv1.IntegrationConfigValue, len(variables))
	for name, value := range variables {
		config[name] = &serverv1.IntegrationConfigValue{
			Value: &serverv1.IntegrationConfigValue_Literal{Literal: value},
		}
	}
	return config
}

func integrationFromProto(proto *serverv1.Integration) Integration {
	integration := Integration{
		ID:            proto.GetId(),
		Name:          proto.GetName(),
		Kind:          integrationKindNameFromProtoName(proto.GetKind().String()),
		EnvironmentID: proto.GetEnvironmentId(),
	}
	if proto.GetCreatedAt() != nil {
		integration.CreatedAt = proto.GetCreatedAt().AsTime()
	}
	if proto.GetUpdatedAt() != nil {
		integration.UpdatedAt = proto.GetUpdatedAt().AsTime()
	}
	return integration
}
