package config

import (
	"context"
	"testing"

	"github.com/chalk-ai/chalk-go/envfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	simpleChalkYAML = `project: test-project
environments:
  default:
    runtime: python311
    requirements: requirements.txt
`

	chalkYAMLNoRequirements = `project: test-project
environments:
  default:
    runtime: python311
`

	multiEnvChalkYAML = `project: multi-env-project
environments:
  dev:
    runtime: python311
    requirements: dev-requirements.txt
  staging:
    runtime: python311
    requirements: staging-requirements.txt
    dockerfile: Dockerfile.staging
  prod:
    runtime: python311
    requirements: prod-requirements.txt
    platform_version: "1.2.3"
`

	nestedChalkYAML = `project: nested-project
environments:
  default:
    runtime: python310
`

	parentChalkYAML = `project: parent-project
environments:
  default:
    runtime: python311
`

	yamlProjectYAML = `project: yaml-project
environments:
  default:
    runtime: python311
`

	ymlProjectYAML = `project: yml-project
environments:
  default:
    runtime: python310
`

	customRequirementsYAML = `project: test-project
environments:
  default:
    runtime: python311
    requirements: custom-requirements.txt
`

	emptyRequirementsYAML = `project: test-project
environments:
  default:
    runtime: python311
    requirements: ""
`

	multiEnvMixedRequirementsYAML = `project: test-project
environments:
  dev:
    runtime: python311
  prod:
    runtime: python311
    requirements: custom.txt
`

	invalidYAML = `project: test-project
environments:
  default:
    runtime: python311
    invalid: [unclosed array
`
)

func TestLoadProjectConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		mockEnvOpts       []envfs.MockEnvironmentOption
		expectError       bool
		errorContains     string
		expectedProject   string
		expectedLocalDir  string
		expectedFilename  string
		validateSettings  func(t *testing.T, settings *ProjectSettings)
	}{
		{
			name: "basic chalk.yaml",
			mockEnvOpts: []envfs.MockEnvironmentOption{
				envfs.WithWorkingDirectory("/project"),
				envfs.WithFileString("/project/chalk.yaml", simpleChalkYAML),
			},
			expectedProject:  "test-project",
			expectedLocalDir: "/project",
			expectedFilename: "/project/chalk.yaml",
			validateSettings: func(t *testing.T, settings *ProjectSettings) {
				require.NotNil(t, settings.Environments["default"])
				assert.Equal(t, "python311", *settings.Environments["default"].Runtime)
				assert.Equal(t, "requirements.txt", *settings.Environments["default"].Requirements)
			},
		},
		{
			name: "chalk.yml extension",
			mockEnvOpts: []envfs.MockEnvironmentOption{
				envfs.WithWorkingDirectory("/project"),
				envfs.WithFileString("/project/chalk.yml", simpleChalkYAML),
			},
			expectedProject:  "test-project",
			expectedLocalDir: "/project",
			expectedFilename: "/project/chalk.yml",
		},
		{
			name: "prefer chalk.yaml over chalk.yml",
			mockEnvOpts: []envfs.MockEnvironmentOption{
				envfs.WithWorkingDirectory("/project"),
				envfs.WithFileString("/project/chalk.yaml", yamlProjectYAML),
				envfs.WithFileString("/project/chalk.yml", ymlProjectYAML),
			},
			expectedProject:  "yaml-project",
			expectedLocalDir: "/project",
			expectedFilename: "/project/chalk.yaml",
		},
		{
			name: "find config in parent directory",
			mockEnvOpts: []envfs.MockEnvironmentOption{
				envfs.WithWorkingDirectory("/project/subdir/nested"),
				envfs.WithFileString("/project/chalk.yaml", simpleChalkYAML),
			},
			expectedProject:  "test-project",
			expectedLocalDir: "/project",
			expectedFilename: "/project/chalk.yaml",
		},
		{
			name: "no config found",
			mockEnvOpts: []envfs.MockEnvironmentOption{
				envfs.WithWorkingDirectory("/project"),
			},
			expectError:   true,
			errorContains: "no chalk.yaml or chalk.yml",
		},
		{
			name: "invalid YAML",
			mockEnvOpts: []envfs.MockEnvironmentOption{
				envfs.WithWorkingDirectory("/project"),
				envfs.WithFileString("/project/chalk.yaml", invalidYAML),
			},
			expectError: true,
		},
		{
			name: "default requirements.txt when present",
			mockEnvOpts: []envfs.MockEnvironmentOption{
				envfs.WithWorkingDirectory("/project"),
				envfs.WithFileString("/project/chalk.yaml", chalkYAMLNoRequirements),
				envfs.WithFileString("/project/requirements.txt", "numpy==1.0.0"),
			},
			expectedProject:  "test-project",
			expectedLocalDir: "/project",
			expectedFilename: "/project/chalk.yaml",
			validateSettings: func(t *testing.T, settings *ProjectSettings) {
				require.NotNil(t, settings.Environments["default"])
				require.NotNil(t, settings.Environments["default"].Requirements)
				assert.Equal(t, "requirements.txt", *settings.Environments["default"].Requirements)
			},
		},
		{
			name: "no default requirements when file absent",
			mockEnvOpts: []envfs.MockEnvironmentOption{
				envfs.WithWorkingDirectory("/project"),
				envfs.WithFileString("/project/chalk.yaml", chalkYAMLNoRequirements),
			},
			expectedProject:  "test-project",
			expectedLocalDir: "/project",
			expectedFilename: "/project/chalk.yaml",
			validateSettings: func(t *testing.T, settings *ProjectSettings) {
				require.NotNil(t, settings.Environments["default"])
				assert.Nil(t, settings.Environments["default"].Requirements)
			},
		},
		{
			name: "explicit requirements overrides default",
			mockEnvOpts: []envfs.MockEnvironmentOption{
				envfs.WithWorkingDirectory("/project"),
				envfs.WithFileString("/project/chalk.yaml", customRequirementsYAML),
				envfs.WithFileString("/project/requirements.txt", "numpy==1.0.0"),
			},
			expectedProject:  "test-project",
			expectedLocalDir: "/project",
			expectedFilename: "/project/chalk.yaml",
			validateSettings: func(t *testing.T, settings *ProjectSettings) {
				require.NotNil(t, settings.Environments["default"])
				require.NotNil(t, settings.Environments["default"].Requirements)
				assert.Equal(t, "custom-requirements.txt", *settings.Environments["default"].Requirements)
			},
		},
		{
			name: "chalkignore file detected",
			mockEnvOpts: []envfs.MockEnvironmentOption{
				envfs.WithWorkingDirectory("/project"),
				envfs.WithFileString("/project/chalk.yaml", simpleChalkYAML),
				envfs.WithFileString("/project/.chalkignore", "*.pyc\n__pycache__/"),
			},
			expectedProject:  "test-project",
			expectedLocalDir: "/project",
			expectedFilename: "/project/chalk.yaml",
			validateSettings: func(t *testing.T, settings *ProjectSettings) {
				require.NotNil(t, settings.ChalkIgnore)
				assert.Equal(t, "/project/.chalkignore", *settings.ChalkIgnore)
			},
		},
		{
			name: "no chalkignore file",
			mockEnvOpts: []envfs.MockEnvironmentOption{
				envfs.WithWorkingDirectory("/project"),
				envfs.WithFileString("/project/chalk.yaml", simpleChalkYAML),
			},
			expectedProject:  "test-project",
			expectedLocalDir: "/project",
			expectedFilename: "/project/chalk.yaml",
			validateSettings: func(t *testing.T, settings *ProjectSettings) {
				assert.Nil(t, settings.ChalkIgnore)
			},
		},
		{
			name: "multiple environments",
			mockEnvOpts: []envfs.MockEnvironmentOption{
				envfs.WithWorkingDirectory("/project"),
				envfs.WithFileString("/project/chalk.yaml", multiEnvChalkYAML),
			},
			expectedProject:  "multi-env-project",
			expectedLocalDir: "/project",
			expectedFilename: "/project/chalk.yaml",
			validateSettings: func(t *testing.T, settings *ProjectSettings) {
				assert.Len(t, settings.Environments, 3)

				// Check dev environment
				require.NotNil(t, settings.Environments["dev"])
				assert.Equal(t, "python311", *settings.Environments["dev"].Runtime)
				assert.Equal(t, "dev-requirements.txt", *settings.Environments["dev"].Requirements)

				// Check staging environment
				require.NotNil(t, settings.Environments["staging"])
				assert.Equal(t, "python311", *settings.Environments["staging"].Runtime)
				assert.Equal(t, "staging-requirements.txt", *settings.Environments["staging"].Requirements)
				assert.Equal(t, "Dockerfile.staging", *settings.Environments["staging"].Dockerfile)

				// Check prod environment
				require.NotNil(t, settings.Environments["prod"])
				assert.Equal(t, "python311", *settings.Environments["prod"].Runtime)
				assert.Equal(t, "prod-requirements.txt", *settings.Environments["prod"].Requirements)
				assert.Equal(t, "1.2.3", *settings.Environments["prod"].PlatformVersion)
			},
		},
		{
			name: "multiple envs with selective default requirements",
			mockEnvOpts: []envfs.MockEnvironmentOption{
				envfs.WithWorkingDirectory("/project"),
				envfs.WithFileString("/project/chalk.yaml", multiEnvMixedRequirementsYAML),
				envfs.WithFileString("/project/requirements.txt", "numpy==1.0.0"),
			},
			expectedProject:  "test-project",
			expectedLocalDir: "/project",
			expectedFilename: "/project/chalk.yaml",
			validateSettings: func(t *testing.T, settings *ProjectSettings) {
				// dev should get default requirements.txt
				require.NotNil(t, settings.Environments["dev"])
				require.NotNil(t, settings.Environments["dev"].Requirements)
				assert.Equal(t, "requirements.txt", *settings.Environments["dev"].Requirements)

				// prod should keep its custom requirements
				require.NotNil(t, settings.Environments["prod"])
				require.NotNil(t, settings.Environments["prod"].Requirements)
				assert.Equal(t, "custom.txt", *settings.Environments["prod"].Requirements)
			},
		},
		{
			name: "empty requirements string gets default",
			mockEnvOpts: []envfs.MockEnvironmentOption{
				envfs.WithWorkingDirectory("/project"),
				envfs.WithFileString("/project/chalk.yaml", emptyRequirementsYAML),
				envfs.WithFileString("/project/requirements.txt", "numpy==1.0.0"),
			},
			expectedProject:  "test-project",
			expectedLocalDir: "/project",
			expectedFilename: "/project/chalk.yaml",
			validateSettings: func(t *testing.T, settings *ProjectSettings) {
				require.NotNil(t, settings.Environments["default"])
				require.NotNil(t, settings.Environments["default"].Requirements)
				assert.Equal(t, "requirements.txt", *settings.Environments["default"].Requirements)
			},
		},
		{
			name: "config in root directory",
			mockEnvOpts: []envfs.MockEnvironmentOption{
				envfs.WithWorkingDirectory("/project"),
				envfs.WithFileString("/chalk.yaml", simpleChalkYAML),
			},
			expectedProject:  "test-project",
			expectedLocalDir: "/",
			expectedFilename: "/chalk.yaml",
		},
		{
			name: "stops at first config found",
			mockEnvOpts: []envfs.MockEnvironmentOption{
				envfs.WithWorkingDirectory("/project/nested"),
				envfs.WithFileString("/project/nested/chalk.yaml", nestedChalkYAML),
				envfs.WithFileString("/project/chalk.yaml", parentChalkYAML),
			},
			expectedProject:  "nested-project",
			expectedLocalDir: "/project/nested",
			expectedFilename: "/project/nested/chalk.yaml",
			validateSettings: func(t *testing.T, settings *ProjectSettings) {
				require.NotNil(t, settings.Environments["default"])
				assert.Equal(t, "python310", *settings.Environments["default"].Runtime)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := envfs.ContextWithEnvironmentGetter(
				context.Background(),
				envfs.NewMockEnvironmentGetter(tt.mockEnvOpts...),
			)

			settings, err := LoadProjectConfig(ctx)

			if tt.expectError {
				require.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
				assert.Nil(t, settings)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, settings)

			assert.Equal(t, tt.expectedProject, settings.Project)
			assert.Equal(t, tt.expectedLocalDir, settings.LocalDirectory)
			assert.Equal(t, tt.expectedFilename, settings.Filename)

			if tt.validateSettings != nil {
				tt.validateSettings(t, settings)
			}
		})
	}
}
