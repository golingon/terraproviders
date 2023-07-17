// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkmanagerscopeconnection "github.com/golingon/terraproviders/azurerm/3.65.0/networkmanagerscopeconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkManagerScopeConnection creates a new instance of [NetworkManagerScopeConnection].
func NewNetworkManagerScopeConnection(name string, args NetworkManagerScopeConnectionArgs) *NetworkManagerScopeConnection {
	return &NetworkManagerScopeConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkManagerScopeConnection)(nil)

// NetworkManagerScopeConnection represents the Terraform resource azurerm_network_manager_scope_connection.
type NetworkManagerScopeConnection struct {
	Name      string
	Args      NetworkManagerScopeConnectionArgs
	state     *networkManagerScopeConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkManagerScopeConnection].
func (nmsc *NetworkManagerScopeConnection) Type() string {
	return "azurerm_network_manager_scope_connection"
}

// LocalName returns the local name for [NetworkManagerScopeConnection].
func (nmsc *NetworkManagerScopeConnection) LocalName() string {
	return nmsc.Name
}

// Configuration returns the configuration (args) for [NetworkManagerScopeConnection].
func (nmsc *NetworkManagerScopeConnection) Configuration() interface{} {
	return nmsc.Args
}

// DependOn is used for other resources to depend on [NetworkManagerScopeConnection].
func (nmsc *NetworkManagerScopeConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(nmsc)
}

// Dependencies returns the list of resources [NetworkManagerScopeConnection] depends_on.
func (nmsc *NetworkManagerScopeConnection) Dependencies() terra.Dependencies {
	return nmsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkManagerScopeConnection].
func (nmsc *NetworkManagerScopeConnection) LifecycleManagement() *terra.Lifecycle {
	return nmsc.Lifecycle
}

// Attributes returns the attributes for [NetworkManagerScopeConnection].
func (nmsc *NetworkManagerScopeConnection) Attributes() networkManagerScopeConnectionAttributes {
	return networkManagerScopeConnectionAttributes{ref: terra.ReferenceResource(nmsc)}
}

// ImportState imports the given attribute values into [NetworkManagerScopeConnection]'s state.
func (nmsc *NetworkManagerScopeConnection) ImportState(av io.Reader) error {
	nmsc.state = &networkManagerScopeConnectionState{}
	if err := json.NewDecoder(av).Decode(nmsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nmsc.Type(), nmsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkManagerScopeConnection] has state.
func (nmsc *NetworkManagerScopeConnection) State() (*networkManagerScopeConnectionState, bool) {
	return nmsc.state, nmsc.state != nil
}

// StateMust returns the state for [NetworkManagerScopeConnection]. Panics if the state is nil.
func (nmsc *NetworkManagerScopeConnection) StateMust() *networkManagerScopeConnectionState {
	if nmsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nmsc.Type(), nmsc.LocalName()))
	}
	return nmsc.state
}

// NetworkManagerScopeConnectionArgs contains the configurations for azurerm_network_manager_scope_connection.
type NetworkManagerScopeConnectionArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkManagerId: string, required
	NetworkManagerId terra.StringValue `hcl:"network_manager_id,attr" validate:"required"`
	// TargetScopeId: string, required
	TargetScopeId terra.StringValue `hcl:"target_scope_id,attr" validate:"required"`
	// TenantId: string, required
	TenantId terra.StringValue `hcl:"tenant_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanagerscopeconnection.Timeouts `hcl:"timeouts,block"`
}
type networkManagerScopeConnectionAttributes struct {
	ref terra.Reference
}

// ConnectionState returns a reference to field connection_state of azurerm_network_manager_scope_connection.
func (nmsc networkManagerScopeConnectionAttributes) ConnectionState() terra.StringValue {
	return terra.ReferenceAsString(nmsc.ref.Append("connection_state"))
}

// Description returns a reference to field description of azurerm_network_manager_scope_connection.
func (nmsc networkManagerScopeConnectionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nmsc.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_network_manager_scope_connection.
func (nmsc networkManagerScopeConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nmsc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_network_manager_scope_connection.
func (nmsc networkManagerScopeConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nmsc.ref.Append("name"))
}

// NetworkManagerId returns a reference to field network_manager_id of azurerm_network_manager_scope_connection.
func (nmsc networkManagerScopeConnectionAttributes) NetworkManagerId() terra.StringValue {
	return terra.ReferenceAsString(nmsc.ref.Append("network_manager_id"))
}

// TargetScopeId returns a reference to field target_scope_id of azurerm_network_manager_scope_connection.
func (nmsc networkManagerScopeConnectionAttributes) TargetScopeId() terra.StringValue {
	return terra.ReferenceAsString(nmsc.ref.Append("target_scope_id"))
}

// TenantId returns a reference to field tenant_id of azurerm_network_manager_scope_connection.
func (nmsc networkManagerScopeConnectionAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(nmsc.ref.Append("tenant_id"))
}

func (nmsc networkManagerScopeConnectionAttributes) Timeouts() networkmanagerscopeconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagerscopeconnection.TimeoutsAttributes](nmsc.ref.Append("timeouts"))
}

type networkManagerScopeConnectionState struct {
	ConnectionState  string                                       `json:"connection_state"`
	Description      string                                       `json:"description"`
	Id               string                                       `json:"id"`
	Name             string                                       `json:"name"`
	NetworkManagerId string                                       `json:"network_manager_id"`
	TargetScopeId    string                                       `json:"target_scope_id"`
	TenantId         string                                       `json:"tenant_id"`
	Timeouts         *networkmanagerscopeconnection.TimeoutsState `json:"timeouts"`
}
