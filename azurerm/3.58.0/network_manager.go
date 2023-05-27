// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkmanager "github.com/golingon/terraproviders/azurerm/3.58.0/networkmanager"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkManager creates a new instance of [NetworkManager].
func NewNetworkManager(name string, args NetworkManagerArgs) *NetworkManager {
	return &NetworkManager{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkManager)(nil)

// NetworkManager represents the Terraform resource azurerm_network_manager.
type NetworkManager struct {
	Name      string
	Args      NetworkManagerArgs
	state     *networkManagerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkManager].
func (nm *NetworkManager) Type() string {
	return "azurerm_network_manager"
}

// LocalName returns the local name for [NetworkManager].
func (nm *NetworkManager) LocalName() string {
	return nm.Name
}

// Configuration returns the configuration (args) for [NetworkManager].
func (nm *NetworkManager) Configuration() interface{} {
	return nm.Args
}

// DependOn is used for other resources to depend on [NetworkManager].
func (nm *NetworkManager) DependOn() terra.Reference {
	return terra.ReferenceResource(nm)
}

// Dependencies returns the list of resources [NetworkManager] depends_on.
func (nm *NetworkManager) Dependencies() terra.Dependencies {
	return nm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkManager].
func (nm *NetworkManager) LifecycleManagement() *terra.Lifecycle {
	return nm.Lifecycle
}

// Attributes returns the attributes for [NetworkManager].
func (nm *NetworkManager) Attributes() networkManagerAttributes {
	return networkManagerAttributes{ref: terra.ReferenceResource(nm)}
}

// ImportState imports the given attribute values into [NetworkManager]'s state.
func (nm *NetworkManager) ImportState(av io.Reader) error {
	nm.state = &networkManagerState{}
	if err := json.NewDecoder(av).Decode(nm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nm.Type(), nm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkManager] has state.
func (nm *NetworkManager) State() (*networkManagerState, bool) {
	return nm.state, nm.state != nil
}

// StateMust returns the state for [NetworkManager]. Panics if the state is nil.
func (nm *NetworkManager) StateMust() *networkManagerState {
	if nm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nm.Type(), nm.LocalName()))
	}
	return nm.state
}

// NetworkManagerArgs contains the configurations for azurerm_network_manager.
type NetworkManagerArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ScopeAccesses: list of string, required
	ScopeAccesses terra.ListValue[terra.StringValue] `hcl:"scope_accesses,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// CrossTenantScopes: min=0
	CrossTenantScopes []networkmanager.CrossTenantScopes `hcl:"cross_tenant_scopes,block" validate:"min=0"`
	// Scope: required
	Scope *networkmanager.Scope `hcl:"scope,block" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanager.Timeouts `hcl:"timeouts,block"`
}
type networkManagerAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_network_manager.
func (nm networkManagerAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nm.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_network_manager.
func (nm networkManagerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nm.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_network_manager.
func (nm networkManagerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nm.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_network_manager.
func (nm networkManagerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nm.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_network_manager.
func (nm networkManagerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nm.ref.Append("resource_group_name"))
}

// ScopeAccesses returns a reference to field scope_accesses of azurerm_network_manager.
func (nm networkManagerAttributes) ScopeAccesses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nm.ref.Append("scope_accesses"))
}

// Tags returns a reference to field tags of azurerm_network_manager.
func (nm networkManagerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nm.ref.Append("tags"))
}

func (nm networkManagerAttributes) CrossTenantScopes() terra.ListValue[networkmanager.CrossTenantScopesAttributes] {
	return terra.ReferenceAsList[networkmanager.CrossTenantScopesAttributes](nm.ref.Append("cross_tenant_scopes"))
}

func (nm networkManagerAttributes) Scope() terra.ListValue[networkmanager.ScopeAttributes] {
	return terra.ReferenceAsList[networkmanager.ScopeAttributes](nm.ref.Append("scope"))
}

func (nm networkManagerAttributes) Timeouts() networkmanager.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanager.TimeoutsAttributes](nm.ref.Append("timeouts"))
}

type networkManagerState struct {
	Description       string                                  `json:"description"`
	Id                string                                  `json:"id"`
	Location          string                                  `json:"location"`
	Name              string                                  `json:"name"`
	ResourceGroupName string                                  `json:"resource_group_name"`
	ScopeAccesses     []string                                `json:"scope_accesses"`
	Tags              map[string]string                       `json:"tags"`
	CrossTenantScopes []networkmanager.CrossTenantScopesState `json:"cross_tenant_scopes"`
	Scope             []networkmanager.ScopeState             `json:"scope"`
	Timeouts          *networkmanager.TimeoutsState           `json:"timeouts"`
}
