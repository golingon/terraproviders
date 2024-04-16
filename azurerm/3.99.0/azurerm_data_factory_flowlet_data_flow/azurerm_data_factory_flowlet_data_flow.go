// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_data_factory_flowlet_data_flow

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_data_factory_flowlet_data_flow.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermDataFactoryFlowletDataFlowState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adffdf *Resource) Type() string {
	return "azurerm_data_factory_flowlet_data_flow"
}

// LocalName returns the local name for [Resource].
func (adffdf *Resource) LocalName() string {
	return adffdf.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adffdf *Resource) Configuration() interface{} {
	return adffdf.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adffdf *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adffdf)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adffdf *Resource) Dependencies() terra.Dependencies {
	return adffdf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adffdf *Resource) LifecycleManagement() *terra.Lifecycle {
	return adffdf.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adffdf *Resource) Attributes() azurermDataFactoryFlowletDataFlowAttributes {
	return azurermDataFactoryFlowletDataFlowAttributes{ref: terra.ReferenceResource(adffdf)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adffdf *Resource) ImportState(state io.Reader) error {
	adffdf.state = &azurermDataFactoryFlowletDataFlowState{}
	if err := json.NewDecoder(state).Decode(adffdf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adffdf.Type(), adffdf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adffdf *Resource) State() (*azurermDataFactoryFlowletDataFlowState, bool) {
	return adffdf.state, adffdf.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adffdf *Resource) StateMust() *azurermDataFactoryFlowletDataFlowState {
	if adffdf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adffdf.Type(), adffdf.LocalName()))
	}
	return adffdf.state
}

// Args contains the configurations for azurerm_data_factory_flowlet_data_flow.
type Args struct {
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
	// Sink: min=0
	Sink []Sink `hcl:"sink,block" validate:"min=0"`
	// Source: min=0
	Source []Source `hcl:"source,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// Transformation: min=0
	Transformation []Transformation `hcl:"transformation,block" validate:"min=0"`
}

type azurermDataFactoryFlowletDataFlowAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of azurerm_data_factory_flowlet_data_flow.
func (adffdf azurermDataFactoryFlowletDataFlowAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](adffdf.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_flowlet_data_flow.
func (adffdf azurermDataFactoryFlowletDataFlowAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(adffdf.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_flowlet_data_flow.
func (adffdf azurermDataFactoryFlowletDataFlowAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(adffdf.ref.Append("description"))
}

// Folder returns a reference to field folder of azurerm_data_factory_flowlet_data_flow.
func (adffdf azurermDataFactoryFlowletDataFlowAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(adffdf.ref.Append("folder"))
}

// Id returns a reference to field id of azurerm_data_factory_flowlet_data_flow.
func (adffdf azurermDataFactoryFlowletDataFlowAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adffdf.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_data_factory_flowlet_data_flow.
func (adffdf azurermDataFactoryFlowletDataFlowAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adffdf.ref.Append("name"))
}

// Script returns a reference to field script of azurerm_data_factory_flowlet_data_flow.
func (adffdf azurermDataFactoryFlowletDataFlowAttributes) Script() terra.StringValue {
	return terra.ReferenceAsString(adffdf.ref.Append("script"))
}

// ScriptLines returns a reference to field script_lines of azurerm_data_factory_flowlet_data_flow.
func (adffdf azurermDataFactoryFlowletDataFlowAttributes) ScriptLines() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](adffdf.ref.Append("script_lines"))
}

func (adffdf azurermDataFactoryFlowletDataFlowAttributes) Sink() terra.ListValue[SinkAttributes] {
	return terra.ReferenceAsList[SinkAttributes](adffdf.ref.Append("sink"))
}

func (adffdf azurermDataFactoryFlowletDataFlowAttributes) Source() terra.ListValue[SourceAttributes] {
	return terra.ReferenceAsList[SourceAttributes](adffdf.ref.Append("source"))
}

func (adffdf azurermDataFactoryFlowletDataFlowAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adffdf.ref.Append("timeouts"))
}

func (adffdf azurermDataFactoryFlowletDataFlowAttributes) Transformation() terra.ListValue[TransformationAttributes] {
	return terra.ReferenceAsList[TransformationAttributes](adffdf.ref.Append("transformation"))
}

type azurermDataFactoryFlowletDataFlowState struct {
	Annotations    []string              `json:"annotations"`
	DataFactoryId  string                `json:"data_factory_id"`
	Description    string                `json:"description"`
	Folder         string                `json:"folder"`
	Id             string                `json:"id"`
	Name           string                `json:"name"`
	Script         string                `json:"script"`
	ScriptLines    []string              `json:"script_lines"`
	Sink           []SinkState           `json:"sink"`
	Source         []SourceState         `json:"source"`
	Timeouts       *TimeoutsState        `json:"timeouts"`
	Transformation []TransformationState `json:"transformation"`
}
