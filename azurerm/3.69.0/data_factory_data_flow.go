// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorydataflow "github.com/golingon/terraproviders/azurerm/3.69.0/datafactorydataflow"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryDataFlow creates a new instance of [DataFactoryDataFlow].
func NewDataFactoryDataFlow(name string, args DataFactoryDataFlowArgs) *DataFactoryDataFlow {
	return &DataFactoryDataFlow{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryDataFlow)(nil)

// DataFactoryDataFlow represents the Terraform resource azurerm_data_factory_data_flow.
type DataFactoryDataFlow struct {
	Name      string
	Args      DataFactoryDataFlowArgs
	state     *dataFactoryDataFlowState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryDataFlow].
func (dfdf *DataFactoryDataFlow) Type() string {
	return "azurerm_data_factory_data_flow"
}

// LocalName returns the local name for [DataFactoryDataFlow].
func (dfdf *DataFactoryDataFlow) LocalName() string {
	return dfdf.Name
}

// Configuration returns the configuration (args) for [DataFactoryDataFlow].
func (dfdf *DataFactoryDataFlow) Configuration() interface{} {
	return dfdf.Args
}

// DependOn is used for other resources to depend on [DataFactoryDataFlow].
func (dfdf *DataFactoryDataFlow) DependOn() terra.Reference {
	return terra.ReferenceResource(dfdf)
}

// Dependencies returns the list of resources [DataFactoryDataFlow] depends_on.
func (dfdf *DataFactoryDataFlow) Dependencies() terra.Dependencies {
	return dfdf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryDataFlow].
func (dfdf *DataFactoryDataFlow) LifecycleManagement() *terra.Lifecycle {
	return dfdf.Lifecycle
}

// Attributes returns the attributes for [DataFactoryDataFlow].
func (dfdf *DataFactoryDataFlow) Attributes() dataFactoryDataFlowAttributes {
	return dataFactoryDataFlowAttributes{ref: terra.ReferenceResource(dfdf)}
}

// ImportState imports the given attribute values into [DataFactoryDataFlow]'s state.
func (dfdf *DataFactoryDataFlow) ImportState(av io.Reader) error {
	dfdf.state = &dataFactoryDataFlowState{}
	if err := json.NewDecoder(av).Decode(dfdf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfdf.Type(), dfdf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryDataFlow] has state.
func (dfdf *DataFactoryDataFlow) State() (*dataFactoryDataFlowState, bool) {
	return dfdf.state, dfdf.state != nil
}

// StateMust returns the state for [DataFactoryDataFlow]. Panics if the state is nil.
func (dfdf *DataFactoryDataFlow) StateMust() *dataFactoryDataFlowState {
	if dfdf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfdf.Type(), dfdf.LocalName()))
	}
	return dfdf.state
}

// DataFactoryDataFlowArgs contains the configurations for azurerm_data_factory_data_flow.
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
}
type dataFactoryDataFlowAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of azurerm_data_factory_data_flow.
func (dfdf dataFactoryDataFlowAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dfdf.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_data_flow.
func (dfdf dataFactoryDataFlowAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dfdf.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_data_flow.
func (dfdf dataFactoryDataFlowAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dfdf.ref.Append("description"))
}

// Folder returns a reference to field folder of azurerm_data_factory_data_flow.
func (dfdf dataFactoryDataFlowAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(dfdf.ref.Append("folder"))
}

// Id returns a reference to field id of azurerm_data_factory_data_flow.
func (dfdf dataFactoryDataFlowAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfdf.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_data_factory_data_flow.
func (dfdf dataFactoryDataFlowAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfdf.ref.Append("name"))
}

// Script returns a reference to field script of azurerm_data_factory_data_flow.
func (dfdf dataFactoryDataFlowAttributes) Script() terra.StringValue {
	return terra.ReferenceAsString(dfdf.ref.Append("script"))
}

// ScriptLines returns a reference to field script_lines of azurerm_data_factory_data_flow.
func (dfdf dataFactoryDataFlowAttributes) ScriptLines() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dfdf.ref.Append("script_lines"))
}

func (dfdf dataFactoryDataFlowAttributes) Sink() terra.ListValue[datafactorydataflow.SinkAttributes] {
	return terra.ReferenceAsList[datafactorydataflow.SinkAttributes](dfdf.ref.Append("sink"))
}

func (dfdf dataFactoryDataFlowAttributes) Source() terra.ListValue[datafactorydataflow.SourceAttributes] {
	return terra.ReferenceAsList[datafactorydataflow.SourceAttributes](dfdf.ref.Append("source"))
}

func (dfdf dataFactoryDataFlowAttributes) Timeouts() datafactorydataflow.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorydataflow.TimeoutsAttributes](dfdf.ref.Append("timeouts"))
}

func (dfdf dataFactoryDataFlowAttributes) Transformation() terra.ListValue[datafactorydataflow.TransformationAttributes] {
	return terra.ReferenceAsList[datafactorydataflow.TransformationAttributes](dfdf.ref.Append("transformation"))
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
