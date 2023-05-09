// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	codepipelinecustomactiontype "github.com/golingon/terraproviders/aws/4.66.1/codepipelinecustomactiontype"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCodepipelineCustomActionType creates a new instance of [CodepipelineCustomActionType].
func NewCodepipelineCustomActionType(name string, args CodepipelineCustomActionTypeArgs) *CodepipelineCustomActionType {
	return &CodepipelineCustomActionType{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodepipelineCustomActionType)(nil)

// CodepipelineCustomActionType represents the Terraform resource aws_codepipeline_custom_action_type.
type CodepipelineCustomActionType struct {
	Name      string
	Args      CodepipelineCustomActionTypeArgs
	state     *codepipelineCustomActionTypeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CodepipelineCustomActionType].
func (ccat *CodepipelineCustomActionType) Type() string {
	return "aws_codepipeline_custom_action_type"
}

// LocalName returns the local name for [CodepipelineCustomActionType].
func (ccat *CodepipelineCustomActionType) LocalName() string {
	return ccat.Name
}

// Configuration returns the configuration (args) for [CodepipelineCustomActionType].
func (ccat *CodepipelineCustomActionType) Configuration() interface{} {
	return ccat.Args
}

// DependOn is used for other resources to depend on [CodepipelineCustomActionType].
func (ccat *CodepipelineCustomActionType) DependOn() terra.Reference {
	return terra.ReferenceResource(ccat)
}

// Dependencies returns the list of resources [CodepipelineCustomActionType] depends_on.
func (ccat *CodepipelineCustomActionType) Dependencies() terra.Dependencies {
	return ccat.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CodepipelineCustomActionType].
func (ccat *CodepipelineCustomActionType) LifecycleManagement() *terra.Lifecycle {
	return ccat.Lifecycle
}

// Attributes returns the attributes for [CodepipelineCustomActionType].
func (ccat *CodepipelineCustomActionType) Attributes() codepipelineCustomActionTypeAttributes {
	return codepipelineCustomActionTypeAttributes{ref: terra.ReferenceResource(ccat)}
}

// ImportState imports the given attribute values into [CodepipelineCustomActionType]'s state.
func (ccat *CodepipelineCustomActionType) ImportState(av io.Reader) error {
	ccat.state = &codepipelineCustomActionTypeState{}
	if err := json.NewDecoder(av).Decode(ccat.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ccat.Type(), ccat.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CodepipelineCustomActionType] has state.
func (ccat *CodepipelineCustomActionType) State() (*codepipelineCustomActionTypeState, bool) {
	return ccat.state, ccat.state != nil
}

// StateMust returns the state for [CodepipelineCustomActionType]. Panics if the state is nil.
func (ccat *CodepipelineCustomActionType) StateMust() *codepipelineCustomActionTypeState {
	if ccat.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ccat.Type(), ccat.LocalName()))
	}
	return ccat.state
}

// CodepipelineCustomActionTypeArgs contains the configurations for aws_codepipeline_custom_action_type.
type CodepipelineCustomActionTypeArgs struct {
	// Category: string, required
	Category terra.StringValue `hcl:"category,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProviderName: string, required
	ProviderName terra.StringValue `hcl:"provider_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// ConfigurationProperty: min=0,max=10
	ConfigurationProperty []codepipelinecustomactiontype.ConfigurationProperty `hcl:"configuration_property,block" validate:"min=0,max=10"`
	// InputArtifactDetails: required
	InputArtifactDetails *codepipelinecustomactiontype.InputArtifactDetails `hcl:"input_artifact_details,block" validate:"required"`
	// OutputArtifactDetails: required
	OutputArtifactDetails *codepipelinecustomactiontype.OutputArtifactDetails `hcl:"output_artifact_details,block" validate:"required"`
	// Settings: optional
	Settings *codepipelinecustomactiontype.Settings `hcl:"settings,block"`
}
type codepipelineCustomActionTypeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_codepipeline_custom_action_type.
func (ccat codepipelineCustomActionTypeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ccat.ref.Append("arn"))
}

// Category returns a reference to field category of aws_codepipeline_custom_action_type.
func (ccat codepipelineCustomActionTypeAttributes) Category() terra.StringValue {
	return terra.ReferenceAsString(ccat.ref.Append("category"))
}

// Id returns a reference to field id of aws_codepipeline_custom_action_type.
func (ccat codepipelineCustomActionTypeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccat.ref.Append("id"))
}

// Owner returns a reference to field owner of aws_codepipeline_custom_action_type.
func (ccat codepipelineCustomActionTypeAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(ccat.ref.Append("owner"))
}

// ProviderName returns a reference to field provider_name of aws_codepipeline_custom_action_type.
func (ccat codepipelineCustomActionTypeAttributes) ProviderName() terra.StringValue {
	return terra.ReferenceAsString(ccat.ref.Append("provider_name"))
}

// Tags returns a reference to field tags of aws_codepipeline_custom_action_type.
func (ccat codepipelineCustomActionTypeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ccat.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_codepipeline_custom_action_type.
func (ccat codepipelineCustomActionTypeAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ccat.ref.Append("tags_all"))
}

// Version returns a reference to field version of aws_codepipeline_custom_action_type.
func (ccat codepipelineCustomActionTypeAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ccat.ref.Append("version"))
}

func (ccat codepipelineCustomActionTypeAttributes) ConfigurationProperty() terra.ListValue[codepipelinecustomactiontype.ConfigurationPropertyAttributes] {
	return terra.ReferenceAsList[codepipelinecustomactiontype.ConfigurationPropertyAttributes](ccat.ref.Append("configuration_property"))
}

func (ccat codepipelineCustomActionTypeAttributes) InputArtifactDetails() terra.ListValue[codepipelinecustomactiontype.InputArtifactDetailsAttributes] {
	return terra.ReferenceAsList[codepipelinecustomactiontype.InputArtifactDetailsAttributes](ccat.ref.Append("input_artifact_details"))
}

func (ccat codepipelineCustomActionTypeAttributes) OutputArtifactDetails() terra.ListValue[codepipelinecustomactiontype.OutputArtifactDetailsAttributes] {
	return terra.ReferenceAsList[codepipelinecustomactiontype.OutputArtifactDetailsAttributes](ccat.ref.Append("output_artifact_details"))
}

func (ccat codepipelineCustomActionTypeAttributes) Settings() terra.ListValue[codepipelinecustomactiontype.SettingsAttributes] {
	return terra.ReferenceAsList[codepipelinecustomactiontype.SettingsAttributes](ccat.ref.Append("settings"))
}

type codepipelineCustomActionTypeState struct {
	Arn                   string                                                    `json:"arn"`
	Category              string                                                    `json:"category"`
	Id                    string                                                    `json:"id"`
	Owner                 string                                                    `json:"owner"`
	ProviderName          string                                                    `json:"provider_name"`
	Tags                  map[string]string                                         `json:"tags"`
	TagsAll               map[string]string                                         `json:"tags_all"`
	Version               string                                                    `json:"version"`
	ConfigurationProperty []codepipelinecustomactiontype.ConfigurationPropertyState `json:"configuration_property"`
	InputArtifactDetails  []codepipelinecustomactiontype.InputArtifactDetailsState  `json:"input_artifact_details"`
	OutputArtifactDetails []codepipelinecustomactiontype.OutputArtifactDetailsState `json:"output_artifact_details"`
	Settings              []codepipelinecustomactiontype.SettingsState              `json:"settings"`
}
