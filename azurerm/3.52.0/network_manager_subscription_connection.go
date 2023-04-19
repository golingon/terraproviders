// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkmanagersubscriptionconnection "github.com/golingon/terraproviders/azurerm/3.52.0/networkmanagersubscriptionconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkManagerSubscriptionConnection creates a new instance of [NetworkManagerSubscriptionConnection].
func NewNetworkManagerSubscriptionConnection(name string, args NetworkManagerSubscriptionConnectionArgs) *NetworkManagerSubscriptionConnection {
	return &NetworkManagerSubscriptionConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkManagerSubscriptionConnection)(nil)

// NetworkManagerSubscriptionConnection represents the Terraform resource azurerm_network_manager_subscription_connection.
type NetworkManagerSubscriptionConnection struct {
	Name      string
	Args      NetworkManagerSubscriptionConnectionArgs
	state     *networkManagerSubscriptionConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkManagerSubscriptionConnection].
func (nmsc *NetworkManagerSubscriptionConnection) Type() string {
	return "azurerm_network_manager_subscription_connection"
}

// LocalName returns the local name for [NetworkManagerSubscriptionConnection].
func (nmsc *NetworkManagerSubscriptionConnection) LocalName() string {
	return nmsc.Name
}

// Configuration returns the configuration (args) for [NetworkManagerSubscriptionConnection].
func (nmsc *NetworkManagerSubscriptionConnection) Configuration() interface{} {
	return nmsc.Args
}

// DependOn is used for other resources to depend on [NetworkManagerSubscriptionConnection].
func (nmsc *NetworkManagerSubscriptionConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(nmsc)
}

// Dependencies returns the list of resources [NetworkManagerSubscriptionConnection] depends_on.
func (nmsc *NetworkManagerSubscriptionConnection) Dependencies() terra.Dependencies {
	return nmsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkManagerSubscriptionConnection].
func (nmsc *NetworkManagerSubscriptionConnection) LifecycleManagement() *terra.Lifecycle {
	return nmsc.Lifecycle
}

// Attributes returns the attributes for [NetworkManagerSubscriptionConnection].
func (nmsc *NetworkManagerSubscriptionConnection) Attributes() networkManagerSubscriptionConnectionAttributes {
	return networkManagerSubscriptionConnectionAttributes{ref: terra.ReferenceResource(nmsc)}
}

// ImportState imports the given attribute values into [NetworkManagerSubscriptionConnection]'s state.
func (nmsc *NetworkManagerSubscriptionConnection) ImportState(av io.Reader) error {
	nmsc.state = &networkManagerSubscriptionConnectionState{}
	if err := json.NewDecoder(av).Decode(nmsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nmsc.Type(), nmsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkManagerSubscriptionConnection] has state.
func (nmsc *NetworkManagerSubscriptionConnection) State() (*networkManagerSubscriptionConnectionState, bool) {
	return nmsc.state, nmsc.state != nil
}

// StateMust returns the state for [NetworkManagerSubscriptionConnection]. Panics if the state is nil.
func (nmsc *NetworkManagerSubscriptionConnection) StateMust() *networkManagerSubscriptionConnectionState {
	if nmsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nmsc.Type(), nmsc.LocalName()))
	}
	return nmsc.state
}

// NetworkManagerSubscriptionConnectionArgs contains the configurations for azurerm_network_manager_subscription_connection.
type NetworkManagerSubscriptionConnectionArgs struct {
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
	Timeouts *networkmanagersubscriptionconnection.Timeouts `hcl:"timeouts,block"`
}
type networkManagerSubscriptionConnectionAttributes struct {
	ref terra.Reference
}

// ConnectionState returns a reference to field connection_state of azurerm_network_manager_subscription_connection.
func (nmsc networkManagerSubscriptionConnectionAttributes) ConnectionState() terra.StringValue {
	return terra.ReferenceAsString(nmsc.ref.Append("connection_state"))
}

// Description returns a reference to field description of azurerm_network_manager_subscription_connection.
func (nmsc networkManagerSubscriptionConnectionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nmsc.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_network_manager_subscription_connection.
func (nmsc networkManagerSubscriptionConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nmsc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_network_manager_subscription_connection.
func (nmsc networkManagerSubscriptionConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nmsc.ref.Append("name"))
}

// NetworkManagerId returns a reference to field network_manager_id of azurerm_network_manager_subscription_connection.
func (nmsc networkManagerSubscriptionConnectionAttributes) NetworkManagerId() terra.StringValue {
	return terra.ReferenceAsString(nmsc.ref.Append("network_manager_id"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_network_manager_subscription_connection.
func (nmsc networkManagerSubscriptionConnectionAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(nmsc.ref.Append("subscription_id"))
}

func (nmsc networkManagerSubscriptionConnectionAttributes) Timeouts() networkmanagersubscriptionconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagersubscriptionconnection.TimeoutsAttributes](nmsc.ref.Append("timeouts"))
}

type networkManagerSubscriptionConnectionState struct {
	ConnectionState  string                                              `json:"connection_state"`
	Description      string                                              `json:"description"`
	Id               string                                              `json:"id"`
	Name             string                                              `json:"name"`
	NetworkManagerId string                                              `json:"network_manager_id"`
	SubscriptionId   string                                              `json:"subscription_id"`
	Timeouts         *networkmanagersubscriptionconnection.TimeoutsState `json:"timeouts"`
}
