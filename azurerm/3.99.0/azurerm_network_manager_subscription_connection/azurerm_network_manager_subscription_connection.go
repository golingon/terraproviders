// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_network_manager_subscription_connection

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

// Resource represents the Terraform resource azurerm_network_manager_subscription_connection.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermNetworkManagerSubscriptionConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (anmsc *Resource) Type() string {
	return "azurerm_network_manager_subscription_connection"
}

// LocalName returns the local name for [Resource].
func (anmsc *Resource) LocalName() string {
	return anmsc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (anmsc *Resource) Configuration() interface{} {
	return anmsc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (anmsc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(anmsc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (anmsc *Resource) Dependencies() terra.Dependencies {
	return anmsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (anmsc *Resource) LifecycleManagement() *terra.Lifecycle {
	return anmsc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (anmsc *Resource) Attributes() azurermNetworkManagerSubscriptionConnectionAttributes {
	return azurermNetworkManagerSubscriptionConnectionAttributes{ref: terra.ReferenceResource(anmsc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (anmsc *Resource) ImportState(state io.Reader) error {
	anmsc.state = &azurermNetworkManagerSubscriptionConnectionState{}
	if err := json.NewDecoder(state).Decode(anmsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", anmsc.Type(), anmsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (anmsc *Resource) State() (*azurermNetworkManagerSubscriptionConnectionState, bool) {
	return anmsc.state, anmsc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (anmsc *Resource) StateMust() *azurermNetworkManagerSubscriptionConnectionState {
	if anmsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", anmsc.Type(), anmsc.LocalName()))
	}
	return anmsc.state
}

// Args contains the configurations for azurerm_network_manager_subscription_connection.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkManagerId: string, required
	NetworkManagerId terra.StringValue `hcl:"network_manager_id,attr" validate:"required"`
	// SubscriptionId: string, required
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermNetworkManagerSubscriptionConnectionAttributes struct {
	ref terra.Reference
}

// ConnectionState returns a reference to field connection_state of azurerm_network_manager_subscription_connection.
func (anmsc azurermNetworkManagerSubscriptionConnectionAttributes) ConnectionState() terra.StringValue {
	return terra.ReferenceAsString(anmsc.ref.Append("connection_state"))
}

// Description returns a reference to field description of azurerm_network_manager_subscription_connection.
func (anmsc azurermNetworkManagerSubscriptionConnectionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(anmsc.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_network_manager_subscription_connection.
func (anmsc azurermNetworkManagerSubscriptionConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(anmsc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_network_manager_subscription_connection.
func (anmsc azurermNetworkManagerSubscriptionConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(anmsc.ref.Append("name"))
}

// NetworkManagerId returns a reference to field network_manager_id of azurerm_network_manager_subscription_connection.
func (anmsc azurermNetworkManagerSubscriptionConnectionAttributes) NetworkManagerId() terra.StringValue {
	return terra.ReferenceAsString(anmsc.ref.Append("network_manager_id"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_network_manager_subscription_connection.
func (anmsc azurermNetworkManagerSubscriptionConnectionAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(anmsc.ref.Append("subscription_id"))
}

func (anmsc azurermNetworkManagerSubscriptionConnectionAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](anmsc.ref.Append("timeouts"))
}

type azurermNetworkManagerSubscriptionConnectionState struct {
	ConnectionState  string         `json:"connection_state"`
	Description      string         `json:"description"`
	Id               string         `json:"id"`
	Name             string         `json:"name"`
	NetworkManagerId string         `json:"network_manager_id"`
	SubscriptionId   string         `json:"subscription_id"`
	Timeouts         *TimeoutsState `json:"timeouts"`
}
