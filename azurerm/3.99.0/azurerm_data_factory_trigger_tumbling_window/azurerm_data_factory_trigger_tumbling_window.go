// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_data_factory_trigger_tumbling_window

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

// Resource represents the Terraform resource azurerm_data_factory_trigger_tumbling_window.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermDataFactoryTriggerTumblingWindowState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adfttw *Resource) Type() string {
	return "azurerm_data_factory_trigger_tumbling_window"
}

// LocalName returns the local name for [Resource].
func (adfttw *Resource) LocalName() string {
	return adfttw.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adfttw *Resource) Configuration() interface{} {
	return adfttw.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adfttw *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adfttw)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adfttw *Resource) Dependencies() terra.Dependencies {
	return adfttw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adfttw *Resource) LifecycleManagement() *terra.Lifecycle {
	return adfttw.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adfttw *Resource) Attributes() azurermDataFactoryTriggerTumblingWindowAttributes {
	return azurermDataFactoryTriggerTumblingWindowAttributes{ref: terra.ReferenceResource(adfttw)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adfttw *Resource) ImportState(state io.Reader) error {
	adfttw.state = &azurermDataFactoryTriggerTumblingWindowState{}
	if err := json.NewDecoder(state).Decode(adfttw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adfttw.Type(), adfttw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adfttw *Resource) State() (*azurermDataFactoryTriggerTumblingWindowState, bool) {
	return adfttw.state, adfttw.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adfttw *Resource) StateMust() *azurermDataFactoryTriggerTumblingWindowState {
	if adfttw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adfttw.Type(), adfttw.LocalName()))
	}
	return adfttw.state
}

// Args contains the configurations for azurerm_data_factory_trigger_tumbling_window.
type Args struct {
	// Activated: bool, optional
	Activated terra.BoolValue `hcl:"activated,attr"`
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Delay: string, optional
	Delay terra.StringValue `hcl:"delay,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EndTime: string, optional
	EndTime terra.StringValue `hcl:"end_time,attr"`
	// Frequency: string, required
	Frequency terra.StringValue `hcl:"frequency,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Interval: number, required
	Interval terra.NumberValue `hcl:"interval,attr" validate:"required"`
	// MaxConcurrency: number, optional
	MaxConcurrency terra.NumberValue `hcl:"max_concurrency,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StartTime: string, required
	StartTime terra.StringValue `hcl:"start_time,attr" validate:"required"`
	// Pipeline: required
	Pipeline *Pipeline `hcl:"pipeline,block" validate:"required"`
	// Retry: optional
	Retry *Retry `hcl:"retry,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// TriggerDependency: min=0
	TriggerDependency []TriggerDependency `hcl:"trigger_dependency,block" validate:"min=0"`
}

type azurermDataFactoryTriggerTumblingWindowAttributes struct {
	ref terra.Reference
}

// Activated returns a reference to field activated of azurerm_data_factory_trigger_tumbling_window.
func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) Activated() terra.BoolValue {
	return terra.ReferenceAsBool(adfttw.ref.Append("activated"))
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_trigger_tumbling_window.
func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adfttw.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_trigger_tumbling_window.
func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](adfttw.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_trigger_tumbling_window.
func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(adfttw.ref.Append("data_factory_id"))
}

// Delay returns a reference to field delay of azurerm_data_factory_trigger_tumbling_window.
func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) Delay() terra.StringValue {
	return terra.ReferenceAsString(adfttw.ref.Append("delay"))
}

// Description returns a reference to field description of azurerm_data_factory_trigger_tumbling_window.
func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(adfttw.ref.Append("description"))
}

// EndTime returns a reference to field end_time of azurerm_data_factory_trigger_tumbling_window.
func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) EndTime() terra.StringValue {
	return terra.ReferenceAsString(adfttw.ref.Append("end_time"))
}

// Frequency returns a reference to field frequency of azurerm_data_factory_trigger_tumbling_window.
func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) Frequency() terra.StringValue {
	return terra.ReferenceAsString(adfttw.ref.Append("frequency"))
}

// Id returns a reference to field id of azurerm_data_factory_trigger_tumbling_window.
func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adfttw.ref.Append("id"))
}

// Interval returns a reference to field interval of azurerm_data_factory_trigger_tumbling_window.
func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) Interval() terra.NumberValue {
	return terra.ReferenceAsNumber(adfttw.ref.Append("interval"))
}

// MaxConcurrency returns a reference to field max_concurrency of azurerm_data_factory_trigger_tumbling_window.
func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) MaxConcurrency() terra.NumberValue {
	return terra.ReferenceAsNumber(adfttw.ref.Append("max_concurrency"))
}

// Name returns a reference to field name of azurerm_data_factory_trigger_tumbling_window.
func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adfttw.ref.Append("name"))
}

// StartTime returns a reference to field start_time of azurerm_data_factory_trigger_tumbling_window.
func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(adfttw.ref.Append("start_time"))
}

func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) Pipeline() terra.ListValue[PipelineAttributes] {
	return terra.ReferenceAsList[PipelineAttributes](adfttw.ref.Append("pipeline"))
}

func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) Retry() terra.ListValue[RetryAttributes] {
	return terra.ReferenceAsList[RetryAttributes](adfttw.ref.Append("retry"))
}

func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adfttw.ref.Append("timeouts"))
}

func (adfttw azurermDataFactoryTriggerTumblingWindowAttributes) TriggerDependency() terra.SetValue[TriggerDependencyAttributes] {
	return terra.ReferenceAsSet[TriggerDependencyAttributes](adfttw.ref.Append("trigger_dependency"))
}

type azurermDataFactoryTriggerTumblingWindowState struct {
	Activated            bool                     `json:"activated"`
	AdditionalProperties map[string]string        `json:"additional_properties"`
	Annotations          []string                 `json:"annotations"`
	DataFactoryId        string                   `json:"data_factory_id"`
	Delay                string                   `json:"delay"`
	Description          string                   `json:"description"`
	EndTime              string                   `json:"end_time"`
	Frequency            string                   `json:"frequency"`
	Id                   string                   `json:"id"`
	Interval             float64                  `json:"interval"`
	MaxConcurrency       float64                  `json:"max_concurrency"`
	Name                 string                   `json:"name"`
	StartTime            string                   `json:"start_time"`
	Pipeline             []PipelineState          `json:"pipeline"`
	Retry                []RetryState             `json:"retry"`
	Timeouts             *TimeoutsState           `json:"timeouts"`
	TriggerDependency    []TriggerDependencyState `json:"trigger_dependency"`
}
