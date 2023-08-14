// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakerpipeline "github.com/golingon/terraproviders/aws/5.12.0/sagemakerpipeline"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSagemakerPipeline creates a new instance of [SagemakerPipeline].
func NewSagemakerPipeline(name string, args SagemakerPipelineArgs) *SagemakerPipeline {
	return &SagemakerPipeline{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerPipeline)(nil)

// SagemakerPipeline represents the Terraform resource aws_sagemaker_pipeline.
type SagemakerPipeline struct {
	Name      string
	Args      SagemakerPipelineArgs
	state     *sagemakerPipelineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerPipeline].
func (sp *SagemakerPipeline) Type() string {
	return "aws_sagemaker_pipeline"
}

// LocalName returns the local name for [SagemakerPipeline].
func (sp *SagemakerPipeline) LocalName() string {
	return sp.Name
}

// Configuration returns the configuration (args) for [SagemakerPipeline].
func (sp *SagemakerPipeline) Configuration() interface{} {
	return sp.Args
}

// DependOn is used for other resources to depend on [SagemakerPipeline].
func (sp *SagemakerPipeline) DependOn() terra.Reference {
	return terra.ReferenceResource(sp)
}

// Dependencies returns the list of resources [SagemakerPipeline] depends_on.
func (sp *SagemakerPipeline) Dependencies() terra.Dependencies {
	return sp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerPipeline].
func (sp *SagemakerPipeline) LifecycleManagement() *terra.Lifecycle {
	return sp.Lifecycle
}

// Attributes returns the attributes for [SagemakerPipeline].
func (sp *SagemakerPipeline) Attributes() sagemakerPipelineAttributes {
	return sagemakerPipelineAttributes{ref: terra.ReferenceResource(sp)}
}

// ImportState imports the given attribute values into [SagemakerPipeline]'s state.
func (sp *SagemakerPipeline) ImportState(av io.Reader) error {
	sp.state = &sagemakerPipelineState{}
	if err := json.NewDecoder(av).Decode(sp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sp.Type(), sp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerPipeline] has state.
func (sp *SagemakerPipeline) State() (*sagemakerPipelineState, bool) {
	return sp.state, sp.state != nil
}

// StateMust returns the state for [SagemakerPipeline]. Panics if the state is nil.
func (sp *SagemakerPipeline) StateMust() *sagemakerPipelineState {
	if sp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sp.Type(), sp.LocalName()))
	}
	return sp.state
}

// SagemakerPipelineArgs contains the configurations for aws_sagemaker_pipeline.
type SagemakerPipelineArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PipelineDefinition: string, optional
	PipelineDefinition terra.StringValue `hcl:"pipeline_definition,attr"`
	// PipelineDescription: string, optional
	PipelineDescription terra.StringValue `hcl:"pipeline_description,attr"`
	// PipelineDisplayName: string, required
	PipelineDisplayName terra.StringValue `hcl:"pipeline_display_name,attr" validate:"required"`
	// PipelineName: string, required
	PipelineName terra.StringValue `hcl:"pipeline_name,attr" validate:"required"`
	// RoleArn: string, optional
	RoleArn terra.StringValue `hcl:"role_arn,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ParallelismConfiguration: optional
	ParallelismConfiguration *sagemakerpipeline.ParallelismConfiguration `hcl:"parallelism_configuration,block"`
	// PipelineDefinitionS3Location: optional
	PipelineDefinitionS3Location *sagemakerpipeline.PipelineDefinitionS3Location `hcl:"pipeline_definition_s3_location,block"`
}
type sagemakerPipelineAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sagemaker_pipeline.
func (sp sagemakerPipelineAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("arn"))
}

// Id returns a reference to field id of aws_sagemaker_pipeline.
func (sp sagemakerPipelineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("id"))
}

// PipelineDefinition returns a reference to field pipeline_definition of aws_sagemaker_pipeline.
func (sp sagemakerPipelineAttributes) PipelineDefinition() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("pipeline_definition"))
}

// PipelineDescription returns a reference to field pipeline_description of aws_sagemaker_pipeline.
func (sp sagemakerPipelineAttributes) PipelineDescription() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("pipeline_description"))
}

// PipelineDisplayName returns a reference to field pipeline_display_name of aws_sagemaker_pipeline.
func (sp sagemakerPipelineAttributes) PipelineDisplayName() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("pipeline_display_name"))
}

// PipelineName returns a reference to field pipeline_name of aws_sagemaker_pipeline.
func (sp sagemakerPipelineAttributes) PipelineName() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("pipeline_name"))
}

// RoleArn returns a reference to field role_arn of aws_sagemaker_pipeline.
func (sp sagemakerPipelineAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("role_arn"))
}

// Tags returns a reference to field tags of aws_sagemaker_pipeline.
func (sp sagemakerPipelineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_pipeline.
func (sp sagemakerPipelineAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sp.ref.Append("tags_all"))
}

func (sp sagemakerPipelineAttributes) ParallelismConfiguration() terra.ListValue[sagemakerpipeline.ParallelismConfigurationAttributes] {
	return terra.ReferenceAsList[sagemakerpipeline.ParallelismConfigurationAttributes](sp.ref.Append("parallelism_configuration"))
}

func (sp sagemakerPipelineAttributes) PipelineDefinitionS3Location() terra.ListValue[sagemakerpipeline.PipelineDefinitionS3LocationAttributes] {
	return terra.ReferenceAsList[sagemakerpipeline.PipelineDefinitionS3LocationAttributes](sp.ref.Append("pipeline_definition_s3_location"))
}

type sagemakerPipelineState struct {
	Arn                          string                                                `json:"arn"`
	Id                           string                                                `json:"id"`
	PipelineDefinition           string                                                `json:"pipeline_definition"`
	PipelineDescription          string                                                `json:"pipeline_description"`
	PipelineDisplayName          string                                                `json:"pipeline_display_name"`
	PipelineName                 string                                                `json:"pipeline_name"`
	RoleArn                      string                                                `json:"role_arn"`
	Tags                         map[string]string                                     `json:"tags"`
	TagsAll                      map[string]string                                     `json:"tags_all"`
	ParallelismConfiguration     []sagemakerpipeline.ParallelismConfigurationState     `json:"parallelism_configuration"`
	PipelineDefinitionS3Location []sagemakerpipeline.PipelineDefinitionS3LocationState `json:"pipeline_definition_s3_location"`
}
