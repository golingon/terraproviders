// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatapipelinePipeline creates a new instance of [DatapipelinePipeline].
func NewDatapipelinePipeline(name string, args DatapipelinePipelineArgs) *DatapipelinePipeline {
	return &DatapipelinePipeline{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatapipelinePipeline)(nil)

// DatapipelinePipeline represents the Terraform resource aws_datapipeline_pipeline.
type DatapipelinePipeline struct {
	Name      string
	Args      DatapipelinePipelineArgs
	state     *datapipelinePipelineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatapipelinePipeline].
func (dp *DatapipelinePipeline) Type() string {
	return "aws_datapipeline_pipeline"
}

// LocalName returns the local name for [DatapipelinePipeline].
func (dp *DatapipelinePipeline) LocalName() string {
	return dp.Name
}

// Configuration returns the configuration (args) for [DatapipelinePipeline].
func (dp *DatapipelinePipeline) Configuration() interface{} {
	return dp.Args
}

// DependOn is used for other resources to depend on [DatapipelinePipeline].
func (dp *DatapipelinePipeline) DependOn() terra.Reference {
	return terra.ReferenceResource(dp)
}

// Dependencies returns the list of resources [DatapipelinePipeline] depends_on.
func (dp *DatapipelinePipeline) Dependencies() terra.Dependencies {
	return dp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatapipelinePipeline].
func (dp *DatapipelinePipeline) LifecycleManagement() *terra.Lifecycle {
	return dp.Lifecycle
}

// Attributes returns the attributes for [DatapipelinePipeline].
func (dp *DatapipelinePipeline) Attributes() datapipelinePipelineAttributes {
	return datapipelinePipelineAttributes{ref: terra.ReferenceResource(dp)}
}

// ImportState imports the given attribute values into [DatapipelinePipeline]'s state.
func (dp *DatapipelinePipeline) ImportState(av io.Reader) error {
	dp.state = &datapipelinePipelineState{}
	if err := json.NewDecoder(av).Decode(dp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dp.Type(), dp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatapipelinePipeline] has state.
func (dp *DatapipelinePipeline) State() (*datapipelinePipelineState, bool) {
	return dp.state, dp.state != nil
}

// StateMust returns the state for [DatapipelinePipeline]. Panics if the state is nil.
func (dp *DatapipelinePipeline) StateMust() *datapipelinePipelineState {
	if dp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dp.Type(), dp.LocalName()))
	}
	return dp.state
}

// DatapipelinePipelineArgs contains the configurations for aws_datapipeline_pipeline.
type DatapipelinePipelineArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type datapipelinePipelineAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of aws_datapipeline_pipeline.
func (dp datapipelinePipelineAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("description"))
}

// Id returns a reference to field id of aws_datapipeline_pipeline.
func (dp datapipelinePipelineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("id"))
}

// Name returns a reference to field name of aws_datapipeline_pipeline.
func (dp datapipelinePipelineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_datapipeline_pipeline.
func (dp datapipelinePipelineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_datapipeline_pipeline.
func (dp datapipelinePipelineAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dp.ref.Append("tags_all"))
}

type datapipelinePipelineState struct {
	Description string            `json:"description"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Tags        map[string]string `json:"tags"`
	TagsAll     map[string]string `json:"tags_all"`
}