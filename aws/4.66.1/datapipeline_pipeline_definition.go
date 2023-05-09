// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	datapipelinepipelinedefinition "github.com/golingon/terraproviders/aws/4.66.1/datapipelinepipelinedefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatapipelinePipelineDefinition creates a new instance of [DatapipelinePipelineDefinition].
func NewDatapipelinePipelineDefinition(name string, args DatapipelinePipelineDefinitionArgs) *DatapipelinePipelineDefinition {
	return &DatapipelinePipelineDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatapipelinePipelineDefinition)(nil)

// DatapipelinePipelineDefinition represents the Terraform resource aws_datapipeline_pipeline_definition.
type DatapipelinePipelineDefinition struct {
	Name      string
	Args      DatapipelinePipelineDefinitionArgs
	state     *datapipelinePipelineDefinitionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatapipelinePipelineDefinition].
func (dpd *DatapipelinePipelineDefinition) Type() string {
	return "aws_datapipeline_pipeline_definition"
}

// LocalName returns the local name for [DatapipelinePipelineDefinition].
func (dpd *DatapipelinePipelineDefinition) LocalName() string {
	return dpd.Name
}

// Configuration returns the configuration (args) for [DatapipelinePipelineDefinition].
func (dpd *DatapipelinePipelineDefinition) Configuration() interface{} {
	return dpd.Args
}

// DependOn is used for other resources to depend on [DatapipelinePipelineDefinition].
func (dpd *DatapipelinePipelineDefinition) DependOn() terra.Reference {
	return terra.ReferenceResource(dpd)
}

// Dependencies returns the list of resources [DatapipelinePipelineDefinition] depends_on.
func (dpd *DatapipelinePipelineDefinition) Dependencies() terra.Dependencies {
	return dpd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatapipelinePipelineDefinition].
func (dpd *DatapipelinePipelineDefinition) LifecycleManagement() *terra.Lifecycle {
	return dpd.Lifecycle
}

// Attributes returns the attributes for [DatapipelinePipelineDefinition].
func (dpd *DatapipelinePipelineDefinition) Attributes() datapipelinePipelineDefinitionAttributes {
	return datapipelinePipelineDefinitionAttributes{ref: terra.ReferenceResource(dpd)}
}

// ImportState imports the given attribute values into [DatapipelinePipelineDefinition]'s state.
func (dpd *DatapipelinePipelineDefinition) ImportState(av io.Reader) error {
	dpd.state = &datapipelinePipelineDefinitionState{}
	if err := json.NewDecoder(av).Decode(dpd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpd.Type(), dpd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatapipelinePipelineDefinition] has state.
func (dpd *DatapipelinePipelineDefinition) State() (*datapipelinePipelineDefinitionState, bool) {
	return dpd.state, dpd.state != nil
}

// StateMust returns the state for [DatapipelinePipelineDefinition]. Panics if the state is nil.
func (dpd *DatapipelinePipelineDefinition) StateMust() *datapipelinePipelineDefinitionState {
	if dpd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpd.Type(), dpd.LocalName()))
	}
	return dpd.state
}

// DatapipelinePipelineDefinitionArgs contains the configurations for aws_datapipeline_pipeline_definition.
type DatapipelinePipelineDefinitionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PipelineId: string, required
	PipelineId terra.StringValue `hcl:"pipeline_id,attr" validate:"required"`
	// ParameterObject: min=0
	ParameterObject []datapipelinepipelinedefinition.ParameterObject `hcl:"parameter_object,block" validate:"min=0"`
	// ParameterValue: min=0
	ParameterValue []datapipelinepipelinedefinition.ParameterValue `hcl:"parameter_value,block" validate:"min=0"`
	// PipelineObject: min=1
	PipelineObject []datapipelinepipelinedefinition.PipelineObject `hcl:"pipeline_object,block" validate:"min=1"`
}
type datapipelinePipelineDefinitionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_datapipeline_pipeline_definition.
func (dpd datapipelinePipelineDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpd.ref.Append("id"))
}

// PipelineId returns a reference to field pipeline_id of aws_datapipeline_pipeline_definition.
func (dpd datapipelinePipelineDefinitionAttributes) PipelineId() terra.StringValue {
	return terra.ReferenceAsString(dpd.ref.Append("pipeline_id"))
}

func (dpd datapipelinePipelineDefinitionAttributes) ParameterObject() terra.SetValue[datapipelinepipelinedefinition.ParameterObjectAttributes] {
	return terra.ReferenceAsSet[datapipelinepipelinedefinition.ParameterObjectAttributes](dpd.ref.Append("parameter_object"))
}

func (dpd datapipelinePipelineDefinitionAttributes) ParameterValue() terra.SetValue[datapipelinepipelinedefinition.ParameterValueAttributes] {
	return terra.ReferenceAsSet[datapipelinepipelinedefinition.ParameterValueAttributes](dpd.ref.Append("parameter_value"))
}

func (dpd datapipelinePipelineDefinitionAttributes) PipelineObject() terra.SetValue[datapipelinepipelinedefinition.PipelineObjectAttributes] {
	return terra.ReferenceAsSet[datapipelinepipelinedefinition.PipelineObjectAttributes](dpd.ref.Append("pipeline_object"))
}

type datapipelinePipelineDefinitionState struct {
	Id              string                                                `json:"id"`
	PipelineId      string                                                `json:"pipeline_id"`
	ParameterObject []datapipelinepipelinedefinition.ParameterObjectState `json:"parameter_object"`
	ParameterValue  []datapipelinepipelinedefinition.ParameterValueState  `json:"parameter_value"`
	PipelineObject  []datapipelinepipelinedefinition.PipelineObjectState  `json:"pipeline_object"`
}
