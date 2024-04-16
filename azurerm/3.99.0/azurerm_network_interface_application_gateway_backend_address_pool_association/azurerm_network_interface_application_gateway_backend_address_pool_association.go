// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_network_interface_application_gateway_backend_address_pool_association

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

// Resource represents the Terraform resource azurerm_network_interface_application_gateway_backend_address_pool_association.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aniagbapa *Resource) Type() string {
	return "azurerm_network_interface_application_gateway_backend_address_pool_association"
}

// LocalName returns the local name for [Resource].
func (aniagbapa *Resource) LocalName() string {
	return aniagbapa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aniagbapa *Resource) Configuration() interface{} {
	return aniagbapa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aniagbapa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aniagbapa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aniagbapa *Resource) Dependencies() terra.Dependencies {
	return aniagbapa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aniagbapa *Resource) LifecycleManagement() *terra.Lifecycle {
	return aniagbapa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aniagbapa *Resource) Attributes() azurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationAttributes {
	return azurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationAttributes{ref: terra.ReferenceResource(aniagbapa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aniagbapa *Resource) ImportState(state io.Reader) error {
	aniagbapa.state = &azurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationState{}
	if err := json.NewDecoder(state).Decode(aniagbapa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aniagbapa.Type(), aniagbapa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aniagbapa *Resource) State() (*azurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationState, bool) {
	return aniagbapa.state, aniagbapa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aniagbapa *Resource) StateMust() *azurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationState {
	if aniagbapa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aniagbapa.Type(), aniagbapa.LocalName()))
	}
	return aniagbapa.state
}

// Args contains the configurations for azurerm_network_interface_application_gateway_backend_address_pool_association.
type Args struct {
	// BackendAddressPoolId: string, required
	BackendAddressPoolId terra.StringValue `hcl:"backend_address_pool_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpConfigurationName: string, required
	IpConfigurationName terra.StringValue `hcl:"ip_configuration_name,attr" validate:"required"`
	// NetworkInterfaceId: string, required
	NetworkInterfaceId terra.StringValue `hcl:"network_interface_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationAttributes struct {
	ref terra.Reference
}

// BackendAddressPoolId returns a reference to field backend_address_pool_id of azurerm_network_interface_application_gateway_backend_address_pool_association.
func (aniagbapa azurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationAttributes) BackendAddressPoolId() terra.StringValue {
	return terra.ReferenceAsString(aniagbapa.ref.Append("backend_address_pool_id"))
}

// Id returns a reference to field id of azurerm_network_interface_application_gateway_backend_address_pool_association.
func (aniagbapa azurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aniagbapa.ref.Append("id"))
}

// IpConfigurationName returns a reference to field ip_configuration_name of azurerm_network_interface_application_gateway_backend_address_pool_association.
func (aniagbapa azurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationAttributes) IpConfigurationName() terra.StringValue {
	return terra.ReferenceAsString(aniagbapa.ref.Append("ip_configuration_name"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of azurerm_network_interface_application_gateway_backend_address_pool_association.
func (aniagbapa azurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(aniagbapa.ref.Append("network_interface_id"))
}

func (aniagbapa azurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aniagbapa.ref.Append("timeouts"))
}

type azurermNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationState struct {
	BackendAddressPoolId string         `json:"backend_address_pool_id"`
	Id                   string         `json:"id"`
	IpConfigurationName  string         `json:"ip_configuration_name"`
	NetworkInterfaceId   string         `json:"network_interface_id"`
	Timeouts             *TimeoutsState `json:"timeouts"`
}
