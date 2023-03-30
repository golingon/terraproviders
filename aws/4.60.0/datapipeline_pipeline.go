// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDatapipelinePipeline(name string, args DatapipelinePipelineArgs) *DatapipelinePipeline {
	return &DatapipelinePipeline{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatapipelinePipeline)(nil)

type DatapipelinePipeline struct {
	Name  string
	Args  DatapipelinePipelineArgs
	state *datapipelinePipelineState
}

func (dp *DatapipelinePipeline) Type() string {
	return "aws_datapipeline_pipeline"
}

func (dp *DatapipelinePipeline) LocalName() string {
	return dp.Name
}

func (dp *DatapipelinePipeline) Configuration() interface{} {
	return dp.Args
}

func (dp *DatapipelinePipeline) Attributes() datapipelinePipelineAttributes {
	return datapipelinePipelineAttributes{ref: terra.ReferenceResource(dp)}
}

func (dp *DatapipelinePipeline) ImportState(av io.Reader) error {
	dp.state = &datapipelinePipelineState{}
	if err := json.NewDecoder(av).Decode(dp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dp.Type(), dp.LocalName(), err)
	}
	return nil
}

func (dp *DatapipelinePipeline) State() (*datapipelinePipelineState, bool) {
	return dp.state, dp.state != nil
}

func (dp *DatapipelinePipeline) StateMust() *datapipelinePipelineState {
	if dp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dp.Type(), dp.LocalName()))
	}
	return dp.state
}

func (dp *DatapipelinePipeline) DependOn() terra.Reference {
	return terra.ReferenceResource(dp)
}

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
	// DependsOn contains resources that DatapipelinePipeline depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type datapipelinePipelineAttributes struct {
	ref terra.Reference
}

func (dp datapipelinePipelineAttributes) Description() terra.StringValue {
	return terra.ReferenceString(dp.ref.Append("description"))
}

func (dp datapipelinePipelineAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dp.ref.Append("id"))
}

func (dp datapipelinePipelineAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dp.ref.Append("name"))
}

func (dp datapipelinePipelineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dp.ref.Append("tags"))
}

func (dp datapipelinePipelineAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dp.ref.Append("tags_all"))
}

type datapipelinePipelineState struct {
	Description string            `json:"description"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Tags        map[string]string `json:"tags"`
	TagsAll     map[string]string `json:"tags_all"`
}
