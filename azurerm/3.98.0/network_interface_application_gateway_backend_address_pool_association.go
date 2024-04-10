// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	networkinterfaceapplicationgatewaybackendaddresspoolassociation "github.com/golingon/terraproviders/azurerm/3.98.0/networkinterfaceapplicationgatewaybackendaddresspoolassociation"
	"io"
)

// NewNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation creates a new instance of [NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation].
func NewNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation(name string, args NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs) *NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation {
	return &NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation)(nil)

// NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation represents the Terraform resource azurerm_network_interface_application_gateway_backend_address_pool_association.
type NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation struct {
	Name      string
	Args      NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs
	state     *networkInterfaceApplicationGatewayBackendAddressPoolAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation].
func (niagbapa *NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation) Type() string {
	return "azurerm_network_interface_application_gateway_backend_address_pool_association"
}

// LocalName returns the local name for [NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation].
func (niagbapa *NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation) LocalName() string {
	return niagbapa.Name
}

// Configuration returns the configuration (args) for [NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation].
func (niagbapa *NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation) Configuration() interface{} {
	return niagbapa.Args
}

// DependOn is used for other resources to depend on [NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation].
func (niagbapa *NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(niagbapa)
}

// Dependencies returns the list of resources [NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation] depends_on.
func (niagbapa *NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation) Dependencies() terra.Dependencies {
	return niagbapa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation].
func (niagbapa *NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation) LifecycleManagement() *terra.Lifecycle {
	return niagbapa.Lifecycle
}

// Attributes returns the attributes for [NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation].
func (niagbapa *NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation) Attributes() networkInterfaceApplicationGatewayBackendAddressPoolAssociationAttributes {
	return networkInterfaceApplicationGatewayBackendAddressPoolAssociationAttributes{ref: terra.ReferenceResource(niagbapa)}
}

// ImportState imports the given attribute values into [NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation]'s state.
func (niagbapa *NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation) ImportState(av io.Reader) error {
	niagbapa.state = &networkInterfaceApplicationGatewayBackendAddressPoolAssociationState{}
	if err := json.NewDecoder(av).Decode(niagbapa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", niagbapa.Type(), niagbapa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation] has state.
func (niagbapa *NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation) State() (*networkInterfaceApplicationGatewayBackendAddressPoolAssociationState, bool) {
	return niagbapa.state, niagbapa.state != nil
}

// StateMust returns the state for [NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation]. Panics if the state is nil.
func (niagbapa *NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation) StateMust() *networkInterfaceApplicationGatewayBackendAddressPoolAssociationState {
	if niagbapa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", niagbapa.Type(), niagbapa.LocalName()))
	}
	return niagbapa.state
}

// NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs contains the configurations for azurerm_network_interface_application_gateway_backend_address_pool_association.
type NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs struct {
	// BackendAddressPoolId: string, required
	BackendAddressPoolId terra.StringValue `hcl:"backend_address_pool_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpConfigurationName: string, required
	IpConfigurationName terra.StringValue `hcl:"ip_configuration_name,attr" validate:"required"`
	// NetworkInterfaceId: string, required
	NetworkInterfaceId terra.StringValue `hcl:"network_interface_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkinterfaceapplicationgatewaybackendaddresspoolassociation.Timeouts `hcl:"timeouts,block"`
}
type networkInterfaceApplicationGatewayBackendAddressPoolAssociationAttributes struct {
	ref terra.Reference
}

// BackendAddressPoolId returns a reference to field backend_address_pool_id of azurerm_network_interface_application_gateway_backend_address_pool_association.
func (niagbapa networkInterfaceApplicationGatewayBackendAddressPoolAssociationAttributes) BackendAddressPoolId() terra.StringValue {
	return terra.ReferenceAsString(niagbapa.ref.Append("backend_address_pool_id"))
}

// Id returns a reference to field id of azurerm_network_interface_application_gateway_backend_address_pool_association.
func (niagbapa networkInterfaceApplicationGatewayBackendAddressPoolAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(niagbapa.ref.Append("id"))
}

// IpConfigurationName returns a reference to field ip_configuration_name of azurerm_network_interface_application_gateway_backend_address_pool_association.
func (niagbapa networkInterfaceApplicationGatewayBackendAddressPoolAssociationAttributes) IpConfigurationName() terra.StringValue {
	return terra.ReferenceAsString(niagbapa.ref.Append("ip_configuration_name"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of azurerm_network_interface_application_gateway_backend_address_pool_association.
func (niagbapa networkInterfaceApplicationGatewayBackendAddressPoolAssociationAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(niagbapa.ref.Append("network_interface_id"))
}

func (niagbapa networkInterfaceApplicationGatewayBackendAddressPoolAssociationAttributes) Timeouts() networkinterfaceapplicationgatewaybackendaddresspoolassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkinterfaceapplicationgatewaybackendaddresspoolassociation.TimeoutsAttributes](niagbapa.ref.Append("timeouts"))
}

type networkInterfaceApplicationGatewayBackendAddressPoolAssociationState struct {
	BackendAddressPoolId string                                                                         `json:"backend_address_pool_id"`
	Id                   string                                                                         `json:"id"`
	IpConfigurationName  string                                                                         `json:"ip_configuration_name"`
	NetworkInterfaceId   string                                                                         `json:"network_interface_id"`
	Timeouts             *networkinterfaceapplicationgatewaybackendaddresspoolassociation.TimeoutsState `json:"timeouts"`
}
