// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_data_factory_trigger_custom_event

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

// Resource represents the Terraform resource azurerm_data_factory_trigger_custom_event.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermDataFactoryTriggerCustomEventState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adftce *Resource) Type() string {
	return "azurerm_data_factory_trigger_custom_event"
}

// LocalName returns the local name for [Resource].
func (adftce *Resource) LocalName() string {
	return adftce.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adftce *Resource) Configuration() interface{} {
	return adftce.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adftce *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adftce)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adftce *Resource) Dependencies() terra.Dependencies {
	return adftce.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adftce *Resource) LifecycleManagement() *terra.Lifecycle {
	return adftce.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adftce *Resource) Attributes() azurermDataFactoryTriggerCustomEventAttributes {
	return azurermDataFactoryTriggerCustomEventAttributes{ref: terra.ReferenceResource(adftce)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adftce *Resource) ImportState(state io.Reader) error {
	adftce.state = &azurermDataFactoryTriggerCustomEventState{}
	if err := json.NewDecoder(state).Decode(adftce.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adftce.Type(), adftce.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adftce *Resource) State() (*azurermDataFactoryTriggerCustomEventState, bool) {
	return adftce.state, adftce.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adftce *Resource) StateMust() *azurermDataFactoryTriggerCustomEventState {
	if adftce.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adftce.Type(), adftce.LocalName()))
	}
	return adftce.state
}

// Args contains the configurations for azurerm_data_factory_trigger_custom_event.
type Args struct {
	// Activated: bool, optional
	Activated terra.BoolValue `hcl:"activated,attr"`
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EventgridTopicId: string, required
	EventgridTopicId terra.StringValue `hcl:"eventgrid_topic_id,attr" validate:"required"`
	// Events: set of string, required
	Events terra.SetValue[terra.StringValue] `hcl:"events,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SubjectBeginsWith: string, optional
	SubjectBeginsWith terra.StringValue `hcl:"subject_begins_with,attr"`
	// SubjectEndsWith: string, optional
	SubjectEndsWith terra.StringValue `hcl:"subject_ends_with,attr"`
	// Pipeline: min=1
	Pipeline []Pipeline `hcl:"pipeline,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermDataFactoryTriggerCustomEventAttributes struct {
	ref terra.Reference
}

// Activated returns a reference to field activated of azurerm_data_factory_trigger_custom_event.
func (adftce azurermDataFactoryTriggerCustomEventAttributes) Activated() terra.BoolValue {
	return terra.ReferenceAsBool(adftce.ref.Append("activated"))
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_trigger_custom_event.
func (adftce azurermDataFactoryTriggerCustomEventAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adftce.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_trigger_custom_event.
func (adftce azurermDataFactoryTriggerCustomEventAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](adftce.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_trigger_custom_event.
func (adftce azurermDataFactoryTriggerCustomEventAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(adftce.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_trigger_custom_event.
func (adftce azurermDataFactoryTriggerCustomEventAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(adftce.ref.Append("description"))
}

// EventgridTopicId returns a reference to field eventgrid_topic_id of azurerm_data_factory_trigger_custom_event.
func (adftce azurermDataFactoryTriggerCustomEventAttributes) EventgridTopicId() terra.StringValue {
	return terra.ReferenceAsString(adftce.ref.Append("eventgrid_topic_id"))
}

// Events returns a reference to field events of azurerm_data_factory_trigger_custom_event.
func (adftce azurermDataFactoryTriggerCustomEventAttributes) Events() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](adftce.ref.Append("events"))
}

// Id returns a reference to field id of azurerm_data_factory_trigger_custom_event.
func (adftce azurermDataFactoryTriggerCustomEventAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adftce.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_data_factory_trigger_custom_event.
func (adftce azurermDataFactoryTriggerCustomEventAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adftce.ref.Append("name"))
}

// SubjectBeginsWith returns a reference to field subject_begins_with of azurerm_data_factory_trigger_custom_event.
func (adftce azurermDataFactoryTriggerCustomEventAttributes) SubjectBeginsWith() terra.StringValue {
	return terra.ReferenceAsString(adftce.ref.Append("subject_begins_with"))
}

// SubjectEndsWith returns a reference to field subject_ends_with of azurerm_data_factory_trigger_custom_event.
func (adftce azurermDataFactoryTriggerCustomEventAttributes) SubjectEndsWith() terra.StringValue {
	return terra.ReferenceAsString(adftce.ref.Append("subject_ends_with"))
}

func (adftce azurermDataFactoryTriggerCustomEventAttributes) Pipeline() terra.SetValue[PipelineAttributes] {
	return terra.ReferenceAsSet[PipelineAttributes](adftce.ref.Append("pipeline"))
}

func (adftce azurermDataFactoryTriggerCustomEventAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adftce.ref.Append("timeouts"))
}

type azurermDataFactoryTriggerCustomEventState struct {
	Activated            bool              `json:"activated"`
	AdditionalProperties map[string]string `json:"additional_properties"`
	Annotations          []string          `json:"annotations"`
	DataFactoryId        string            `json:"data_factory_id"`
	Description          string            `json:"description"`
	EventgridTopicId     string            `json:"eventgrid_topic_id"`
	Events               []string          `json:"events"`
	Id                   string            `json:"id"`
	Name                 string            `json:"name"`
	SubjectBeginsWith    string            `json:"subject_begins_with"`
	SubjectEndsWith      string            `json:"subject_ends_with"`
	Pipeline             []PipelineState   `json:"pipeline"`
	Timeouts             *TimeoutsState    `json:"timeouts"`
}
