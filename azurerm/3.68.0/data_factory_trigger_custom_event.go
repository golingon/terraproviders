// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorytriggercustomevent "github.com/golingon/terraproviders/azurerm/3.68.0/datafactorytriggercustomevent"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryTriggerCustomEvent creates a new instance of [DataFactoryTriggerCustomEvent].
func NewDataFactoryTriggerCustomEvent(name string, args DataFactoryTriggerCustomEventArgs) *DataFactoryTriggerCustomEvent {
	return &DataFactoryTriggerCustomEvent{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryTriggerCustomEvent)(nil)

// DataFactoryTriggerCustomEvent represents the Terraform resource azurerm_data_factory_trigger_custom_event.
type DataFactoryTriggerCustomEvent struct {
	Name      string
	Args      DataFactoryTriggerCustomEventArgs
	state     *dataFactoryTriggerCustomEventState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryTriggerCustomEvent].
func (dftce *DataFactoryTriggerCustomEvent) Type() string {
	return "azurerm_data_factory_trigger_custom_event"
}

// LocalName returns the local name for [DataFactoryTriggerCustomEvent].
func (dftce *DataFactoryTriggerCustomEvent) LocalName() string {
	return dftce.Name
}

// Configuration returns the configuration (args) for [DataFactoryTriggerCustomEvent].
func (dftce *DataFactoryTriggerCustomEvent) Configuration() interface{} {
	return dftce.Args
}

// DependOn is used for other resources to depend on [DataFactoryTriggerCustomEvent].
func (dftce *DataFactoryTriggerCustomEvent) DependOn() terra.Reference {
	return terra.ReferenceResource(dftce)
}

// Dependencies returns the list of resources [DataFactoryTriggerCustomEvent] depends_on.
func (dftce *DataFactoryTriggerCustomEvent) Dependencies() terra.Dependencies {
	return dftce.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryTriggerCustomEvent].
func (dftce *DataFactoryTriggerCustomEvent) LifecycleManagement() *terra.Lifecycle {
	return dftce.Lifecycle
}

// Attributes returns the attributes for [DataFactoryTriggerCustomEvent].
func (dftce *DataFactoryTriggerCustomEvent) Attributes() dataFactoryTriggerCustomEventAttributes {
	return dataFactoryTriggerCustomEventAttributes{ref: terra.ReferenceResource(dftce)}
}

// ImportState imports the given attribute values into [DataFactoryTriggerCustomEvent]'s state.
func (dftce *DataFactoryTriggerCustomEvent) ImportState(av io.Reader) error {
	dftce.state = &dataFactoryTriggerCustomEventState{}
	if err := json.NewDecoder(av).Decode(dftce.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dftce.Type(), dftce.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryTriggerCustomEvent] has state.
func (dftce *DataFactoryTriggerCustomEvent) State() (*dataFactoryTriggerCustomEventState, bool) {
	return dftce.state, dftce.state != nil
}

// StateMust returns the state for [DataFactoryTriggerCustomEvent]. Panics if the state is nil.
func (dftce *DataFactoryTriggerCustomEvent) StateMust() *dataFactoryTriggerCustomEventState {
	if dftce.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dftce.Type(), dftce.LocalName()))
	}
	return dftce.state
}

// DataFactoryTriggerCustomEventArgs contains the configurations for azurerm_data_factory_trigger_custom_event.
type DataFactoryTriggerCustomEventArgs struct {
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
	Pipeline []datafactorytriggercustomevent.Pipeline `hcl:"pipeline,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *datafactorytriggercustomevent.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryTriggerCustomEventAttributes struct {
	ref terra.Reference
}

// Activated returns a reference to field activated of azurerm_data_factory_trigger_custom_event.
func (dftce dataFactoryTriggerCustomEventAttributes) Activated() terra.BoolValue {
	return terra.ReferenceAsBool(dftce.ref.Append("activated"))
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_trigger_custom_event.
func (dftce dataFactoryTriggerCustomEventAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dftce.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_trigger_custom_event.
func (dftce dataFactoryTriggerCustomEventAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dftce.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_trigger_custom_event.
func (dftce dataFactoryTriggerCustomEventAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dftce.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_trigger_custom_event.
func (dftce dataFactoryTriggerCustomEventAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dftce.ref.Append("description"))
}

// EventgridTopicId returns a reference to field eventgrid_topic_id of azurerm_data_factory_trigger_custom_event.
func (dftce dataFactoryTriggerCustomEventAttributes) EventgridTopicId() terra.StringValue {
	return terra.ReferenceAsString(dftce.ref.Append("eventgrid_topic_id"))
}

// Events returns a reference to field events of azurerm_data_factory_trigger_custom_event.
func (dftce dataFactoryTriggerCustomEventAttributes) Events() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dftce.ref.Append("events"))
}

// Id returns a reference to field id of azurerm_data_factory_trigger_custom_event.
func (dftce dataFactoryTriggerCustomEventAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dftce.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_data_factory_trigger_custom_event.
func (dftce dataFactoryTriggerCustomEventAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dftce.ref.Append("name"))
}

// SubjectBeginsWith returns a reference to field subject_begins_with of azurerm_data_factory_trigger_custom_event.
func (dftce dataFactoryTriggerCustomEventAttributes) SubjectBeginsWith() terra.StringValue {
	return terra.ReferenceAsString(dftce.ref.Append("subject_begins_with"))
}

// SubjectEndsWith returns a reference to field subject_ends_with of azurerm_data_factory_trigger_custom_event.
func (dftce dataFactoryTriggerCustomEventAttributes) SubjectEndsWith() terra.StringValue {
	return terra.ReferenceAsString(dftce.ref.Append("subject_ends_with"))
}

func (dftce dataFactoryTriggerCustomEventAttributes) Pipeline() terra.SetValue[datafactorytriggercustomevent.PipelineAttributes] {
	return terra.ReferenceAsSet[datafactorytriggercustomevent.PipelineAttributes](dftce.ref.Append("pipeline"))
}

func (dftce dataFactoryTriggerCustomEventAttributes) Timeouts() datafactorytriggercustomevent.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorytriggercustomevent.TimeoutsAttributes](dftce.ref.Append("timeouts"))
}

type dataFactoryTriggerCustomEventState struct {
	Activated            bool                                          `json:"activated"`
	AdditionalProperties map[string]string                             `json:"additional_properties"`
	Annotations          []string                                      `json:"annotations"`
	DataFactoryId        string                                        `json:"data_factory_id"`
	Description          string                                        `json:"description"`
	EventgridTopicId     string                                        `json:"eventgrid_topic_id"`
	Events               []string                                      `json:"events"`
	Id                   string                                        `json:"id"`
	Name                 string                                        `json:"name"`
	SubjectBeginsWith    string                                        `json:"subject_begins_with"`
	SubjectEndsWith      string                                        `json:"subject_ends_with"`
	Pipeline             []datafactorytriggercustomevent.PipelineState `json:"pipeline"`
	Timeouts             *datafactorytriggercustomevent.TimeoutsState  `json:"timeouts"`
}
