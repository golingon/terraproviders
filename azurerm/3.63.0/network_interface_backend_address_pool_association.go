// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkinterfacebackendaddresspoolassociation "github.com/golingon/terraproviders/azurerm/3.63.0/networkinterfacebackendaddresspoolassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkInterfaceBackendAddressPoolAssociation creates a new instance of [NetworkInterfaceBackendAddressPoolAssociation].
func NewNetworkInterfaceBackendAddressPoolAssociation(name string, args NetworkInterfaceBackendAddressPoolAssociationArgs) *NetworkInterfaceBackendAddressPoolAssociation {
	return &NetworkInterfaceBackendAddressPoolAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkInterfaceBackendAddressPoolAssociation)(nil)

// NetworkInterfaceBackendAddressPoolAssociation represents the Terraform resource azurerm_network_interface_backend_address_pool_association.
type NetworkInterfaceBackendAddressPoolAssociation struct {
	Name      string
	Args      NetworkInterfaceBackendAddressPoolAssociationArgs
	state     *networkInterfaceBackendAddressPoolAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkInterfaceBackendAddressPoolAssociation].
func (nibapa *NetworkInterfaceBackendAddressPoolAssociation) Type() string {
	return "azurerm_network_interface_backend_address_pool_association"
}

// LocalName returns the local name for [NetworkInterfaceBackendAddressPoolAssociation].
func (nibapa *NetworkInterfaceBackendAddressPoolAssociation) LocalName() string {
	return nibapa.Name
}

// Configuration returns the configuration (args) for [NetworkInterfaceBackendAddressPoolAssociation].
func (nibapa *NetworkInterfaceBackendAddressPoolAssociation) Configuration() interface{} {
	return nibapa.Args
}

// DependOn is used for other resources to depend on [NetworkInterfaceBackendAddressPoolAssociation].
func (nibapa *NetworkInterfaceBackendAddressPoolAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(nibapa)
}

// Dependencies returns the list of resources [NetworkInterfaceBackendAddressPoolAssociation] depends_on.
func (nibapa *NetworkInterfaceBackendAddressPoolAssociation) Dependencies() terra.Dependencies {
	return nibapa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkInterfaceBackendAddressPoolAssociation].
func (nibapa *NetworkInterfaceBackendAddressPoolAssociation) LifecycleManagement() *terra.Lifecycle {
	return nibapa.Lifecycle
}

// Attributes returns the attributes for [NetworkInterfaceBackendAddressPoolAssociation].
func (nibapa *NetworkInterfaceBackendAddressPoolAssociation) Attributes() networkInterfaceBackendAddressPoolAssociationAttributes {
	return networkInterfaceBackendAddressPoolAssociationAttributes{ref: terra.ReferenceResource(nibapa)}
}

// ImportState imports the given attribute values into [NetworkInterfaceBackendAddressPoolAssociation]'s state.
func (nibapa *NetworkInterfaceBackendAddressPoolAssociation) ImportState(av io.Reader) error {
	nibapa.state = &networkInterfaceBackendAddressPoolAssociationState{}
	if err := json.NewDecoder(av).Decode(nibapa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nibapa.Type(), nibapa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkInterfaceBackendAddressPoolAssociation] has state.
func (nibapa *NetworkInterfaceBackendAddressPoolAssociation) State() (*networkInterfaceBackendAddressPoolAssociationState, bool) {
	return nibapa.state, nibapa.state != nil
}

// StateMust returns the state for [NetworkInterfaceBackendAddressPoolAssociation]. Panics if the state is nil.
func (nibapa *NetworkInterfaceBackendAddressPoolAssociation) StateMust() *networkInterfaceBackendAddressPoolAssociationState {
	if nibapa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nibapa.Type(), nibapa.LocalName()))
	}
	return nibapa.state
}

// NetworkInterfaceBackendAddressPoolAssociationArgs contains the configurations for azurerm_network_interface_backend_address_pool_association.
type NetworkInterfaceBackendAddressPoolAssociationArgs struct {
	// BackendAddressPoolId: string, required
	BackendAddressPoolId terra.StringValue `hcl:"backend_address_pool_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpConfigurationName: string, required
	IpConfigurationName terra.StringValue `hcl:"ip_configuration_name,attr" validate:"required"`
	// NetworkInterfaceId: string, required
	NetworkInterfaceId terra.StringValue `hcl:"network_interface_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkinterfacebackendaddresspoolassociation.Timeouts `hcl:"timeouts,block"`
}
type networkInterfaceBackendAddressPoolAssociationAttributes struct {
	ref terra.Reference
}

// BackendAddressPoolId returns a reference to field backend_address_pool_id of azurerm_network_interface_backend_address_pool_association.
func (nibapa networkInterfaceBackendAddressPoolAssociationAttributes) BackendAddressPoolId() terra.StringValue {
	return terra.ReferenceAsString(nibapa.ref.Append("backend_address_pool_id"))
}

// Id returns a reference to field id of azurerm_network_interface_backend_address_pool_association.
func (nibapa networkInterfaceBackendAddressPoolAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nibapa.ref.Append("id"))
}

// IpConfigurationName returns a reference to field ip_configuration_name of azurerm_network_interface_backend_address_pool_association.
func (nibapa networkInterfaceBackendAddressPoolAssociationAttributes) IpConfigurationName() terra.StringValue {
	return terra.ReferenceAsString(nibapa.ref.Append("ip_configuration_name"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of azurerm_network_interface_backend_address_pool_association.
func (nibapa networkInterfaceBackendAddressPoolAssociationAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(nibapa.ref.Append("network_interface_id"))
}

func (nibapa networkInterfaceBackendAddressPoolAssociationAttributes) Timeouts() networkinterfacebackendaddresspoolassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkinterfacebackendaddresspoolassociation.TimeoutsAttributes](nibapa.ref.Append("timeouts"))
}

type networkInterfaceBackendAddressPoolAssociationState struct {
	BackendAddressPoolId string                                                       `json:"backend_address_pool_id"`
	Id                   string                                                       `json:"id"`
	IpConfigurationName  string                                                       `json:"ip_configuration_name"`
	NetworkInterfaceId   string                                                       `json:"network_interface_id"`
	Timeouts             *networkinterfacebackendaddresspoolassociation.TimeoutsState `json:"timeouts"`
}
