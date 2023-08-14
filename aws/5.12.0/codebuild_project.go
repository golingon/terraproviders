// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	codebuildproject "github.com/golingon/terraproviders/aws/5.12.0/codebuildproject"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCodebuildProject creates a new instance of [CodebuildProject].
func NewCodebuildProject(name string, args CodebuildProjectArgs) *CodebuildProject {
	return &CodebuildProject{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodebuildProject)(nil)

// CodebuildProject represents the Terraform resource aws_codebuild_project.
type CodebuildProject struct {
	Name      string
	Args      CodebuildProjectArgs
	state     *codebuildProjectState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CodebuildProject].
func (cp *CodebuildProject) Type() string {
	return "aws_codebuild_project"
}

// LocalName returns the local name for [CodebuildProject].
func (cp *CodebuildProject) LocalName() string {
	return cp.Name
}

// Configuration returns the configuration (args) for [CodebuildProject].
func (cp *CodebuildProject) Configuration() interface{} {
	return cp.Args
}

// DependOn is used for other resources to depend on [CodebuildProject].
func (cp *CodebuildProject) DependOn() terra.Reference {
	return terra.ReferenceResource(cp)
}

// Dependencies returns the list of resources [CodebuildProject] depends_on.
func (cp *CodebuildProject) Dependencies() terra.Dependencies {
	return cp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CodebuildProject].
func (cp *CodebuildProject) LifecycleManagement() *terra.Lifecycle {
	return cp.Lifecycle
}

// Attributes returns the attributes for [CodebuildProject].
func (cp *CodebuildProject) Attributes() codebuildProjectAttributes {
	return codebuildProjectAttributes{ref: terra.ReferenceResource(cp)}
}

// ImportState imports the given attribute values into [CodebuildProject]'s state.
func (cp *CodebuildProject) ImportState(av io.Reader) error {
	cp.state = &codebuildProjectState{}
	if err := json.NewDecoder(av).Decode(cp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cp.Type(), cp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CodebuildProject] has state.
func (cp *CodebuildProject) State() (*codebuildProjectState, bool) {
	return cp.state, cp.state != nil
}

// StateMust returns the state for [CodebuildProject]. Panics if the state is nil.
func (cp *CodebuildProject) StateMust() *codebuildProjectState {
	if cp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cp.Type(), cp.LocalName()))
	}
	return cp.state
}

// CodebuildProjectArgs contains the configurations for aws_codebuild_project.
type CodebuildProjectArgs struct {
	// BadgeEnabled: bool, optional
	BadgeEnabled terra.BoolValue `hcl:"badge_enabled,attr"`
	// BuildTimeout: number, optional
	BuildTimeout terra.NumberValue `hcl:"build_timeout,attr"`
	// ConcurrentBuildLimit: number, optional
	ConcurrentBuildLimit terra.NumberValue `hcl:"concurrent_build_limit,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EncryptionKey: string, optional
	EncryptionKey terra.StringValue `hcl:"encryption_key,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProjectVisibility: string, optional
	ProjectVisibility terra.StringValue `hcl:"project_visibility,attr"`
	// QueuedTimeout: number, optional
	QueuedTimeout terra.NumberValue `hcl:"queued_timeout,attr"`
	// ResourceAccessRole: string, optional
	ResourceAccessRole terra.StringValue `hcl:"resource_access_role,attr"`
	// ServiceRole: string, required
	ServiceRole terra.StringValue `hcl:"service_role,attr" validate:"required"`
	// SourceVersion: string, optional
	SourceVersion terra.StringValue `hcl:"source_version,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Artifacts: required
	Artifacts *codebuildproject.Artifacts `hcl:"artifacts,block" validate:"required"`
	// BuildBatchConfig: optional
	BuildBatchConfig *codebuildproject.BuildBatchConfig `hcl:"build_batch_config,block"`
	// Cache: optional
	Cache *codebuildproject.Cache `hcl:"cache,block"`
	// Environment: required
	Environment *codebuildproject.Environment `hcl:"environment,block" validate:"required"`
	// FileSystemLocations: min=0
	FileSystemLocations []codebuildproject.FileSystemLocations `hcl:"file_system_locations,block" validate:"min=0"`
	// LogsConfig: optional
	LogsConfig *codebuildproject.LogsConfig `hcl:"logs_config,block"`
	// SecondaryArtifacts: min=0,max=12
	SecondaryArtifacts []codebuildproject.SecondaryArtifacts `hcl:"secondary_artifacts,block" validate:"min=0,max=12"`
	// SecondarySourceVersion: min=0,max=12
	SecondarySourceVersion []codebuildproject.SecondarySourceVersion `hcl:"secondary_source_version,block" validate:"min=0,max=12"`
	// SecondarySources: min=0,max=12
	SecondarySources []codebuildproject.SecondarySources `hcl:"secondary_sources,block" validate:"min=0,max=12"`
	// Source: required
	Source *codebuildproject.Source `hcl:"source,block" validate:"required"`
	// VpcConfig: optional
	VpcConfig *codebuildproject.VpcConfig `hcl:"vpc_config,block"`
}
type codebuildProjectAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_codebuild_project.
func (cp codebuildProjectAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("arn"))
}

// BadgeEnabled returns a reference to field badge_enabled of aws_codebuild_project.
func (cp codebuildProjectAttributes) BadgeEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cp.ref.Append("badge_enabled"))
}

// BadgeUrl returns a reference to field badge_url of aws_codebuild_project.
func (cp codebuildProjectAttributes) BadgeUrl() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("badge_url"))
}

// BuildTimeout returns a reference to field build_timeout of aws_codebuild_project.
func (cp codebuildProjectAttributes) BuildTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(cp.ref.Append("build_timeout"))
}

// ConcurrentBuildLimit returns a reference to field concurrent_build_limit of aws_codebuild_project.
func (cp codebuildProjectAttributes) ConcurrentBuildLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(cp.ref.Append("concurrent_build_limit"))
}

// Description returns a reference to field description of aws_codebuild_project.
func (cp codebuildProjectAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("description"))
}

// EncryptionKey returns a reference to field encryption_key of aws_codebuild_project.
func (cp codebuildProjectAttributes) EncryptionKey() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("encryption_key"))
}

// Id returns a reference to field id of aws_codebuild_project.
func (cp codebuildProjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("id"))
}

// Name returns a reference to field name of aws_codebuild_project.
func (cp codebuildProjectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("name"))
}

// ProjectVisibility returns a reference to field project_visibility of aws_codebuild_project.
func (cp codebuildProjectAttributes) ProjectVisibility() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("project_visibility"))
}

// PublicProjectAlias returns a reference to field public_project_alias of aws_codebuild_project.
func (cp codebuildProjectAttributes) PublicProjectAlias() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("public_project_alias"))
}

// QueuedTimeout returns a reference to field queued_timeout of aws_codebuild_project.
func (cp codebuildProjectAttributes) QueuedTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(cp.ref.Append("queued_timeout"))
}

// ResourceAccessRole returns a reference to field resource_access_role of aws_codebuild_project.
func (cp codebuildProjectAttributes) ResourceAccessRole() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("resource_access_role"))
}

// ServiceRole returns a reference to field service_role of aws_codebuild_project.
func (cp codebuildProjectAttributes) ServiceRole() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("service_role"))
}

// SourceVersion returns a reference to field source_version of aws_codebuild_project.
func (cp codebuildProjectAttributes) SourceVersion() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("source_version"))
}

// Tags returns a reference to field tags of aws_codebuild_project.
func (cp codebuildProjectAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_codebuild_project.
func (cp codebuildProjectAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cp.ref.Append("tags_all"))
}

func (cp codebuildProjectAttributes) Artifacts() terra.ListValue[codebuildproject.ArtifactsAttributes] {
	return terra.ReferenceAsList[codebuildproject.ArtifactsAttributes](cp.ref.Append("artifacts"))
}

func (cp codebuildProjectAttributes) BuildBatchConfig() terra.ListValue[codebuildproject.BuildBatchConfigAttributes] {
	return terra.ReferenceAsList[codebuildproject.BuildBatchConfigAttributes](cp.ref.Append("build_batch_config"))
}

func (cp codebuildProjectAttributes) Cache() terra.ListValue[codebuildproject.CacheAttributes] {
	return terra.ReferenceAsList[codebuildproject.CacheAttributes](cp.ref.Append("cache"))
}

func (cp codebuildProjectAttributes) Environment() terra.ListValue[codebuildproject.EnvironmentAttributes] {
	return terra.ReferenceAsList[codebuildproject.EnvironmentAttributes](cp.ref.Append("environment"))
}

func (cp codebuildProjectAttributes) FileSystemLocations() terra.SetValue[codebuildproject.FileSystemLocationsAttributes] {
	return terra.ReferenceAsSet[codebuildproject.FileSystemLocationsAttributes](cp.ref.Append("file_system_locations"))
}

func (cp codebuildProjectAttributes) LogsConfig() terra.ListValue[codebuildproject.LogsConfigAttributes] {
	return terra.ReferenceAsList[codebuildproject.LogsConfigAttributes](cp.ref.Append("logs_config"))
}

func (cp codebuildProjectAttributes) SecondaryArtifacts() terra.SetValue[codebuildproject.SecondaryArtifactsAttributes] {
	return terra.ReferenceAsSet[codebuildproject.SecondaryArtifactsAttributes](cp.ref.Append("secondary_artifacts"))
}

func (cp codebuildProjectAttributes) SecondarySourceVersion() terra.SetValue[codebuildproject.SecondarySourceVersionAttributes] {
	return terra.ReferenceAsSet[codebuildproject.SecondarySourceVersionAttributes](cp.ref.Append("secondary_source_version"))
}

func (cp codebuildProjectAttributes) SecondarySources() terra.SetValue[codebuildproject.SecondarySourcesAttributes] {
	return terra.ReferenceAsSet[codebuildproject.SecondarySourcesAttributes](cp.ref.Append("secondary_sources"))
}

func (cp codebuildProjectAttributes) Source() terra.ListValue[codebuildproject.SourceAttributes] {
	return terra.ReferenceAsList[codebuildproject.SourceAttributes](cp.ref.Append("source"))
}

func (cp codebuildProjectAttributes) VpcConfig() terra.ListValue[codebuildproject.VpcConfigAttributes] {
	return terra.ReferenceAsList[codebuildproject.VpcConfigAttributes](cp.ref.Append("vpc_config"))
}

type codebuildProjectState struct {
	Arn                    string                                         `json:"arn"`
	BadgeEnabled           bool                                           `json:"badge_enabled"`
	BadgeUrl               string                                         `json:"badge_url"`
	BuildTimeout           float64                                        `json:"build_timeout"`
	ConcurrentBuildLimit   float64                                        `json:"concurrent_build_limit"`
	Description            string                                         `json:"description"`
	EncryptionKey          string                                         `json:"encryption_key"`
	Id                     string                                         `json:"id"`
	Name                   string                                         `json:"name"`
	ProjectVisibility      string                                         `json:"project_visibility"`
	PublicProjectAlias     string                                         `json:"public_project_alias"`
	QueuedTimeout          float64                                        `json:"queued_timeout"`
	ResourceAccessRole     string                                         `json:"resource_access_role"`
	ServiceRole            string                                         `json:"service_role"`
	SourceVersion          string                                         `json:"source_version"`
	Tags                   map[string]string                              `json:"tags"`
	TagsAll                map[string]string                              `json:"tags_all"`
	Artifacts              []codebuildproject.ArtifactsState              `json:"artifacts"`
	BuildBatchConfig       []codebuildproject.BuildBatchConfigState       `json:"build_batch_config"`
	Cache                  []codebuildproject.CacheState                  `json:"cache"`
	Environment            []codebuildproject.EnvironmentState            `json:"environment"`
	FileSystemLocations    []codebuildproject.FileSystemLocationsState    `json:"file_system_locations"`
	LogsConfig             []codebuildproject.LogsConfigState             `json:"logs_config"`
	SecondaryArtifacts     []codebuildproject.SecondaryArtifactsState     `json:"secondary_artifacts"`
	SecondarySourceVersion []codebuildproject.SecondarySourceVersionState `json:"secondary_source_version"`
	SecondarySources       []codebuildproject.SecondarySourcesState       `json:"secondary_sources"`
	Source                 []codebuildproject.SourceState                 `json:"source"`
	VpcConfig              []codebuildproject.VpcConfigState              `json:"vpc_config"`
}
