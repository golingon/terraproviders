// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iothubconsumergroup "github.com/golingon/terraproviders/azurerm/3.52.0/iothubconsumergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIothubConsumerGroup creates a new instance of [IothubConsumerGroup].
func NewIothubConsumerGroup(name string, args IothubConsumerGroupArgs) *IothubConsumerGroup {
	return &IothubConsumerGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IothubConsumerGroup)(nil)

// IothubConsumerGroup represents the Terraform resource azurerm_iothub_consumer_group.
type IothubConsumerGroup struct {
	Name      string
	Args      IothubConsumerGroupArgs
	state     *iothubConsumerGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IothubConsumerGroup].
func (icg *IothubConsumerGroup) Type() string {
	return "azurerm_iothub_consumer_group"
}

// LocalName returns the local name for [IothubConsumerGroup].
func (icg *IothubConsumerGroup) LocalName() string {
	return icg.Name
}

// Configuration returns the configuration (args) for [IothubConsumerGroup].
func (icg *IothubConsumerGroup) Configuration() interface{} {
	return icg.Args
}

// DependOn is used for other resources to depend on [IothubConsumerGroup].
func (icg *IothubConsumerGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(icg)
}

// Dependencies returns the list of resources [IothubConsumerGroup] depends_on.
func (icg *IothubConsumerGroup) Dependencies() terra.Dependencies {
	return icg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IothubConsumerGroup].
func (icg *IothubConsumerGroup) LifecycleManagement() *terra.Lifecycle {
	return icg.Lifecycle
}

// Attributes returns the attributes for [IothubConsumerGroup].
func (icg *IothubConsumerGroup) Attributes() iothubConsumerGroupAttributes {
	return iothubConsumerGroupAttributes{ref: terra.ReferenceResource(icg)}
}

// ImportState imports the given attribute values into [IothubConsumerGroup]'s state.
func (icg *IothubConsumerGroup) ImportState(av io.Reader) error {
	icg.state = &iothubConsumerGroupState{}
	if err := json.NewDecoder(av).Decode(icg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", icg.Type(), icg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IothubConsumerGroup] has state.
func (icg *IothubConsumerGroup) State() (*iothubConsumerGroupState, bool) {
	return icg.state, icg.state != nil
}

// StateMust returns the state for [IothubConsumerGroup]. Panics if the state is nil.
func (icg *IothubConsumerGroup) StateMust() *iothubConsumerGroupState {
	if icg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", icg.Type(), icg.LocalName()))
	}
	return icg.state
}

// IothubConsumerGroupArgs contains the configurations for azurerm_iothub_consumer_group.
type IothubConsumerGroupArgs struct {
	// EventhubEndpointName: string, required
	EventhubEndpointName terra.StringValue `hcl:"eventhub_endpoint_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IothubName: string, required
	IothubName terra.StringValue `hcl:"iothub_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *iothubconsumergroup.Timeouts `hcl:"timeouts,block"`
}
type iothubConsumerGroupAttributes struct {
	ref terra.Reference
}

// EventhubEndpointName returns a reference to field eventhub_endpoint_name of azurerm_iothub_consumer_group.
func (icg iothubConsumerGroupAttributes) EventhubEndpointName() terra.StringValue {
	return terra.ReferenceAsString(icg.ref.Append("eventhub_endpoint_name"))
}

// Id returns a reference to field id of azurerm_iothub_consumer_group.
func (icg iothubConsumerGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(icg.ref.Append("id"))
}

// IothubName returns a reference to field iothub_name of azurerm_iothub_consumer_group.
func (icg iothubConsumerGroupAttributes) IothubName() terra.StringValue {
	return terra.ReferenceAsString(icg.ref.Append("iothub_name"))
}

// Name returns a reference to field name of azurerm_iothub_consumer_group.
func (icg iothubConsumerGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(icg.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_iothub_consumer_group.
func (icg iothubConsumerGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(icg.ref.Append("resource_group_name"))
}

func (icg iothubConsumerGroupAttributes) Timeouts() iothubconsumergroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iothubconsumergroup.TimeoutsAttributes](icg.ref.Append("timeouts"))
}

type iothubConsumerGroupState struct {
	EventhubEndpointName string                             `json:"eventhub_endpoint_name"`
	Id                   string                             `json:"id"`
	IothubName           string                             `json:"iothub_name"`
	Name                 string                             `json:"name"`
	ResourceGroupName    string                             `json:"resource_group_name"`
	Timeouts             *iothubconsumergroup.TimeoutsState `json:"timeouts"`
}
