// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataDatapipelinePipeline creates a new instance of [DataDatapipelinePipeline].
func NewDataDatapipelinePipeline(name string, args DataDatapipelinePipelineArgs) *DataDatapipelinePipeline {
	return &DataDatapipelinePipeline{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDatapipelinePipeline)(nil)

// DataDatapipelinePipeline represents the Terraform data resource aws_datapipeline_pipeline.
type DataDatapipelinePipeline struct {
	Name string
	Args DataDatapipelinePipelineArgs
}

// DataSource returns the Terraform object type for [DataDatapipelinePipeline].
func (dp *DataDatapipelinePipeline) DataSource() string {
	return "aws_datapipeline_pipeline"
}

// LocalName returns the local name for [DataDatapipelinePipeline].
func (dp *DataDatapipelinePipeline) LocalName() string {
	return dp.Name
}

// Configuration returns the configuration (args) for [DataDatapipelinePipeline].
func (dp *DataDatapipelinePipeline) Configuration() interface{} {
	return dp.Args
}

// Attributes returns the attributes for [DataDatapipelinePipeline].
func (dp *DataDatapipelinePipeline) Attributes() dataDatapipelinePipelineAttributes {
	return dataDatapipelinePipelineAttributes{ref: terra.ReferenceDataResource(dp)}
}

// DataDatapipelinePipelineArgs contains the configurations for aws_datapipeline_pipeline.
type DataDatapipelinePipelineArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PipelineId: string, required
	PipelineId terra.StringValue `hcl:"pipeline_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataDatapipelinePipelineAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of aws_datapipeline_pipeline.
func (dp dataDatapipelinePipelineAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("description"))
}

// Id returns a reference to field id of aws_datapipeline_pipeline.
func (dp dataDatapipelinePipelineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("id"))
}

// Name returns a reference to field name of aws_datapipeline_pipeline.
func (dp dataDatapipelinePipelineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("name"))
}

// PipelineId returns a reference to field pipeline_id of aws_datapipeline_pipeline.
func (dp dataDatapipelinePipelineAttributes) PipelineId() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("pipeline_id"))
}

// Tags returns a reference to field tags of aws_datapipeline_pipeline.
func (dp dataDatapipelinePipelineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dp.ref.Append("tags"))
}