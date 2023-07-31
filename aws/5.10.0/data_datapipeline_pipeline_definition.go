// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datadatapipelinepipelinedefinition "github.com/golingon/terraproviders/aws/5.10.0/datadatapipelinepipelinedefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDatapipelinePipelineDefinition creates a new instance of [DataDatapipelinePipelineDefinition].
func NewDataDatapipelinePipelineDefinition(name string, args DataDatapipelinePipelineDefinitionArgs) *DataDatapipelinePipelineDefinition {
	return &DataDatapipelinePipelineDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDatapipelinePipelineDefinition)(nil)

// DataDatapipelinePipelineDefinition represents the Terraform data resource aws_datapipeline_pipeline_definition.
type DataDatapipelinePipelineDefinition struct {
	Name string
	Args DataDatapipelinePipelineDefinitionArgs
}

// DataSource returns the Terraform object type for [DataDatapipelinePipelineDefinition].
func (dpd *DataDatapipelinePipelineDefinition) DataSource() string {
	return "aws_datapipeline_pipeline_definition"
}

// LocalName returns the local name for [DataDatapipelinePipelineDefinition].
func (dpd *DataDatapipelinePipelineDefinition) LocalName() string {
	return dpd.Name
}

// Configuration returns the configuration (args) for [DataDatapipelinePipelineDefinition].
func (dpd *DataDatapipelinePipelineDefinition) Configuration() interface{} {
	return dpd.Args
}

// Attributes returns the attributes for [DataDatapipelinePipelineDefinition].
func (dpd *DataDatapipelinePipelineDefinition) Attributes() dataDatapipelinePipelineDefinitionAttributes {
	return dataDatapipelinePipelineDefinitionAttributes{ref: terra.ReferenceDataResource(dpd)}
}

// DataDatapipelinePipelineDefinitionArgs contains the configurations for aws_datapipeline_pipeline_definition.
type DataDatapipelinePipelineDefinitionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PipelineId: string, required
	PipelineId terra.StringValue `hcl:"pipeline_id,attr" validate:"required"`
	// ParameterObject: min=0
	ParameterObject []datadatapipelinepipelinedefinition.ParameterObject `hcl:"parameter_object,block" validate:"min=0"`
	// PipelineObject: min=0
	PipelineObject []datadatapipelinepipelinedefinition.PipelineObject `hcl:"pipeline_object,block" validate:"min=0"`
	// ParameterValue: min=0
	ParameterValue []datadatapipelinepipelinedefinition.ParameterValue `hcl:"parameter_value,block" validate:"min=0"`
}
type dataDatapipelinePipelineDefinitionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_datapipeline_pipeline_definition.
func (dpd dataDatapipelinePipelineDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpd.ref.Append("id"))
}

// PipelineId returns a reference to field pipeline_id of aws_datapipeline_pipeline_definition.
func (dpd dataDatapipelinePipelineDefinitionAttributes) PipelineId() terra.StringValue {
	return terra.ReferenceAsString(dpd.ref.Append("pipeline_id"))
}

func (dpd dataDatapipelinePipelineDefinitionAttributes) ParameterObject() terra.SetValue[datadatapipelinepipelinedefinition.ParameterObjectAttributes] {
	return terra.ReferenceAsSet[datadatapipelinepipelinedefinition.ParameterObjectAttributes](dpd.ref.Append("parameter_object"))
}

func (dpd dataDatapipelinePipelineDefinitionAttributes) PipelineObject() terra.SetValue[datadatapipelinepipelinedefinition.PipelineObjectAttributes] {
	return terra.ReferenceAsSet[datadatapipelinepipelinedefinition.PipelineObjectAttributes](dpd.ref.Append("pipeline_object"))
}

func (dpd dataDatapipelinePipelineDefinitionAttributes) ParameterValue() terra.SetValue[datadatapipelinepipelinedefinition.ParameterValueAttributes] {
	return terra.ReferenceAsSet[datadatapipelinepipelinedefinition.ParameterValueAttributes](dpd.ref.Append("parameter_value"))
}
