// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorydataflow "github.com/golingon/terraproviders/azurerm/3.49.0/datafactorydataflow"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDataFactoryDataFlow(name string, args DataFactoryDataFlowArgs) *DataFactoryDataFlow {
	return &DataFactoryDataFlow{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryDataFlow)(nil)

type DataFactoryDataFlow struct {
	Name  string
	Args  DataFactoryDataFlowArgs
	state *dataFactoryDataFlowState
}

func (dfdf *DataFactoryDataFlow) Type() string {
	return "azurerm_data_factory_data_flow"
}

func (dfdf *DataFactoryDataFlow) LocalName() string {
	return dfdf.Name
}

func (dfdf *DataFactoryDataFlow) Configuration() interface{} {
	return dfdf.Args
}

func (dfdf *DataFactoryDataFlow) Attributes() dataFactoryDataFlowAttributes {
	return dataFactoryDataFlowAttributes{ref: terra.ReferenceResource(dfdf)}
}

func (dfdf *DataFactoryDataFlow) ImportState(av io.Reader) error {
	dfdf.state = &dataFactoryDataFlowState{}
	if err := json.NewDecoder(av).Decode(dfdf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfdf.Type(), dfdf.LocalName(), err)
	}
	return nil
}

func (dfdf *DataFactoryDataFlow) State() (*dataFactoryDataFlowState, bool) {
	return dfdf.state, dfdf.state != nil
}

func (dfdf *DataFactoryDataFlow) StateMust() *dataFactoryDataFlowState {
	if dfdf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfdf.Type(), dfdf.LocalName()))
	}
	return dfdf.state
}

func (dfdf *DataFactoryDataFlow) DependOn() terra.Reference {
	return terra.ReferenceResource(dfdf)
}

type DataFactoryDataFlowArgs struct {
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Folder: string, optional
	Folder terra.StringValue `hcl:"folder,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Script: string, optional
	Script terra.StringValue `hcl:"script,attr"`
	// ScriptLines: list of string, optional
	ScriptLines terra.ListValue[terra.StringValue] `hcl:"script_lines,attr"`
	// Sink: min=1
	Sink []datafactorydataflow.Sink `hcl:"sink,block" validate:"min=1"`
	// Source: min=1
	Source []datafactorydataflow.Source `hcl:"source,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *datafactorydataflow.Timeouts `hcl:"timeouts,block"`
	// Transformation: min=0
	Transformation []datafactorydataflow.Transformation `hcl:"transformation,block" validate:"min=0"`
	// DependsOn contains resources that DataFactoryDataFlow depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dataFactoryDataFlowAttributes struct {
	ref terra.Reference
}

func (dfdf dataFactoryDataFlowAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](dfdf.ref.Append("annotations"))
}

func (dfdf dataFactoryDataFlowAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceString(dfdf.ref.Append("data_factory_id"))
}

func (dfdf dataFactoryDataFlowAttributes) Description() terra.StringValue {
	return terra.ReferenceString(dfdf.ref.Append("description"))
}

func (dfdf dataFactoryDataFlowAttributes) Folder() terra.StringValue {
	return terra.ReferenceString(dfdf.ref.Append("folder"))
}

func (dfdf dataFactoryDataFlowAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dfdf.ref.Append("id"))
}

func (dfdf dataFactoryDataFlowAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dfdf.ref.Append("name"))
}

func (dfdf dataFactoryDataFlowAttributes) Script() terra.StringValue {
	return terra.ReferenceString(dfdf.ref.Append("script"))
}

func (dfdf dataFactoryDataFlowAttributes) ScriptLines() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](dfdf.ref.Append("script_lines"))
}

func (dfdf dataFactoryDataFlowAttributes) Sink() terra.ListValue[datafactorydataflow.SinkAttributes] {
	return terra.ReferenceList[datafactorydataflow.SinkAttributes](dfdf.ref.Append("sink"))
}

func (dfdf dataFactoryDataFlowAttributes) Source() terra.ListValue[datafactorydataflow.SourceAttributes] {
	return terra.ReferenceList[datafactorydataflow.SourceAttributes](dfdf.ref.Append("source"))
}

func (dfdf dataFactoryDataFlowAttributes) Timeouts() datafactorydataflow.TimeoutsAttributes {
	return terra.ReferenceSingle[datafactorydataflow.TimeoutsAttributes](dfdf.ref.Append("timeouts"))
}

func (dfdf dataFactoryDataFlowAttributes) Transformation() terra.ListValue[datafactorydataflow.TransformationAttributes] {
	return terra.ReferenceList[datafactorydataflow.TransformationAttributes](dfdf.ref.Append("transformation"))
}

type dataFactoryDataFlowState struct {
	Annotations    []string                                  `json:"annotations"`
	DataFactoryId  string                                    `json:"data_factory_id"`
	Description    string                                    `json:"description"`
	Folder         string                                    `json:"folder"`
	Id             string                                    `json:"id"`
	Name           string                                    `json:"name"`
	Script         string                                    `json:"script"`
	ScriptLines    []string                                  `json:"script_lines"`
	Sink           []datafactorydataflow.SinkState           `json:"sink"`
	Source         []datafactorydataflow.SourceState         `json:"source"`
	Timeouts       *datafactorydataflow.TimeoutsState        `json:"timeouts"`
	Transformation []datafactorydataflow.TransformationState `json:"transformation"`
}
