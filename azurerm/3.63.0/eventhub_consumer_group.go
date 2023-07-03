// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	eventhubconsumergroup "github.com/golingon/terraproviders/azurerm/3.63.0/eventhubconsumergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEventhubConsumerGroup creates a new instance of [EventhubConsumerGroup].
func NewEventhubConsumerGroup(name string, args EventhubConsumerGroupArgs) *EventhubConsumerGroup {
	return &EventhubConsumerGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EventhubConsumerGroup)(nil)

// EventhubConsumerGroup represents the Terraform resource azurerm_eventhub_consumer_group.
type EventhubConsumerGroup struct {
	Name      string
	Args      EventhubConsumerGroupArgs
	state     *eventhubConsumerGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EventhubConsumerGroup].
func (ecg *EventhubConsumerGroup) Type() string {
	return "azurerm_eventhub_consumer_group"
}

// LocalName returns the local name for [EventhubConsumerGroup].
func (ecg *EventhubConsumerGroup) LocalName() string {
	return ecg.Name
}

// Configuration returns the configuration (args) for [EventhubConsumerGroup].
func (ecg *EventhubConsumerGroup) Configuration() interface{} {
	return ecg.Args
}

// DependOn is used for other resources to depend on [EventhubConsumerGroup].
func (ecg *EventhubConsumerGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(ecg)
}

// Dependencies returns the list of resources [EventhubConsumerGroup] depends_on.
func (ecg *EventhubConsumerGroup) Dependencies() terra.Dependencies {
	return ecg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EventhubConsumerGroup].
func (ecg *EventhubConsumerGroup) LifecycleManagement() *terra.Lifecycle {
	return ecg.Lifecycle
}

// Attributes returns the attributes for [EventhubConsumerGroup].
func (ecg *EventhubConsumerGroup) Attributes() eventhubConsumerGroupAttributes {
	return eventhubConsumerGroupAttributes{ref: terra.ReferenceResource(ecg)}
}

// ImportState imports the given attribute values into [EventhubConsumerGroup]'s state.
func (ecg *EventhubConsumerGroup) ImportState(av io.Reader) error {
	ecg.state = &eventhubConsumerGroupState{}
	if err := json.NewDecoder(av).Decode(ecg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ecg.Type(), ecg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EventhubConsumerGroup] has state.
func (ecg *EventhubConsumerGroup) State() (*eventhubConsumerGroupState, bool) {
	return ecg.state, ecg.state != nil
}

// StateMust returns the state for [EventhubConsumerGroup]. Panics if the state is nil.
func (ecg *EventhubConsumerGroup) StateMust() *eventhubConsumerGroupState {
	if ecg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ecg.Type(), ecg.LocalName()))
	}
	return ecg.state
}

// EventhubConsumerGroupArgs contains the configurations for azurerm_eventhub_consumer_group.
type EventhubConsumerGroupArgs struct {
	// EventhubName: string, required
	EventhubName terra.StringValue `hcl:"eventhub_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceName: string, required
	NamespaceName terra.StringValue `hcl:"namespace_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// UserMetadata: string, optional
	UserMetadata terra.StringValue `hcl:"user_metadata,attr"`
	// Timeouts: optional
	Timeouts *eventhubconsumergroup.Timeouts `hcl:"timeouts,block"`
}
type eventhubConsumerGroupAttributes struct {
	ref terra.Reference
}

// EventhubName returns a reference to field eventhub_name of azurerm_eventhub_consumer_group.
func (ecg eventhubConsumerGroupAttributes) EventhubName() terra.StringValue {
	return terra.ReferenceAsString(ecg.ref.Append("eventhub_name"))
}

// Id returns a reference to field id of azurerm_eventhub_consumer_group.
func (ecg eventhubConsumerGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ecg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_eventhub_consumer_group.
func (ecg eventhubConsumerGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ecg.ref.Append("name"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_eventhub_consumer_group.
func (ecg eventhubConsumerGroupAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(ecg.ref.Append("namespace_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventhub_consumer_group.
func (ecg eventhubConsumerGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ecg.ref.Append("resource_group_name"))
}

// UserMetadata returns a reference to field user_metadata of azurerm_eventhub_consumer_group.
func (ecg eventhubConsumerGroupAttributes) UserMetadata() terra.StringValue {
	return terra.ReferenceAsString(ecg.ref.Append("user_metadata"))
}

func (ecg eventhubConsumerGroupAttributes) Timeouts() eventhubconsumergroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eventhubconsumergroup.TimeoutsAttributes](ecg.ref.Append("timeouts"))
}

type eventhubConsumerGroupState struct {
	EventhubName      string                               `json:"eventhub_name"`
	Id                string                               `json:"id"`
	Name              string                               `json:"name"`
	NamespaceName     string                               `json:"namespace_name"`
	ResourceGroupName string                               `json:"resource_group_name"`
	UserMetadata      string                               `json:"user_metadata"`
	Timeouts          *eventhubconsumergroup.TimeoutsState `json:"timeouts"`
}