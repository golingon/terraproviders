// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privateendpoint "github.com/golingon/terraproviders/azurerm/3.67.0/privateendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivateEndpoint creates a new instance of [PrivateEndpoint].
func NewPrivateEndpoint(name string, args PrivateEndpointArgs) *PrivateEndpoint {
	return &PrivateEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateEndpoint)(nil)

// PrivateEndpoint represents the Terraform resource azurerm_private_endpoint.
type PrivateEndpoint struct {
	Name      string
	Args      PrivateEndpointArgs
	state     *privateEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateEndpoint].
func (pe *PrivateEndpoint) Type() string {
	return "azurerm_private_endpoint"
}

// LocalName returns the local name for [PrivateEndpoint].
func (pe *PrivateEndpoint) LocalName() string {
	return pe.Name
}

// Configuration returns the configuration (args) for [PrivateEndpoint].
func (pe *PrivateEndpoint) Configuration() interface{} {
	return pe.Args
}

// DependOn is used for other resources to depend on [PrivateEndpoint].
func (pe *PrivateEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(pe)
}

// Dependencies returns the list of resources [PrivateEndpoint] depends_on.
func (pe *PrivateEndpoint) Dependencies() terra.Dependencies {
	return pe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateEndpoint].
func (pe *PrivateEndpoint) LifecycleManagement() *terra.Lifecycle {
	return pe.Lifecycle
}

// Attributes returns the attributes for [PrivateEndpoint].
func (pe *PrivateEndpoint) Attributes() privateEndpointAttributes {
	return privateEndpointAttributes{ref: terra.ReferenceResource(pe)}
}

// ImportState imports the given attribute values into [PrivateEndpoint]'s state.
func (pe *PrivateEndpoint) ImportState(av io.Reader) error {
	pe.state = &privateEndpointState{}
	if err := json.NewDecoder(av).Decode(pe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pe.Type(), pe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateEndpoint] has state.
func (pe *PrivateEndpoint) State() (*privateEndpointState, bool) {
	return pe.state, pe.state != nil
}

// StateMust returns the state for [PrivateEndpoint]. Panics if the state is nil.
func (pe *PrivateEndpoint) StateMust() *privateEndpointState {
	if pe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pe.Type(), pe.LocalName()))
	}
	return pe.state
}

// PrivateEndpointArgs contains the configurations for azurerm_private_endpoint.
type PrivateEndpointArgs struct {
	// CustomNetworkInterfaceName: string, optional
	CustomNetworkInterfaceName terra.StringValue `hcl:"custom_network_interface_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// CustomDnsConfigs: min=0
	CustomDnsConfigs []privateendpoint.CustomDnsConfigs `hcl:"custom_dns_configs,block" validate:"min=0"`
	// NetworkInterface: min=0
	NetworkInterface []privateendpoint.NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
	// PrivateDnsZoneConfigs: min=0
	PrivateDnsZoneConfigs []privateendpoint.PrivateDnsZoneConfigs `hcl:"private_dns_zone_configs,block" validate:"min=0"`
	// IpConfiguration: min=0
	IpConfiguration []privateendpoint.IpConfiguration `hcl:"ip_configuration,block" validate:"min=0"`
	// PrivateDnsZoneGroup: optional
	PrivateDnsZoneGroup *privateendpoint.PrivateDnsZoneGroup `hcl:"private_dns_zone_group,block"`
	// PrivateServiceConnection: required
	PrivateServiceConnection *privateendpoint.PrivateServiceConnection `hcl:"private_service_connection,block" validate:"required"`
	// Timeouts: optional
	Timeouts *privateendpoint.Timeouts `hcl:"timeouts,block"`
}
type privateEndpointAttributes struct {
	ref terra.Reference
}

// CustomNetworkInterfaceName returns a reference to field custom_network_interface_name of azurerm_private_endpoint.
func (pe privateEndpointAttributes) CustomNetworkInterfaceName() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("custom_network_interface_name"))
}

// Id returns a reference to field id of azurerm_private_endpoint.
func (pe privateEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_private_endpoint.
func (pe privateEndpointAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_private_endpoint.
func (pe privateEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_endpoint.
func (pe privateEndpointAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("resource_group_name"))
}

// SubnetId returns a reference to field subnet_id of azurerm_private_endpoint.
func (pe privateEndpointAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_private_endpoint.
func (pe privateEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pe.ref.Append("tags"))
}

func (pe privateEndpointAttributes) CustomDnsConfigs() terra.ListValue[privateendpoint.CustomDnsConfigsAttributes] {
	return terra.ReferenceAsList[privateendpoint.CustomDnsConfigsAttributes](pe.ref.Append("custom_dns_configs"))
}

func (pe privateEndpointAttributes) NetworkInterface() terra.ListValue[privateendpoint.NetworkInterfaceAttributes] {
	return terra.ReferenceAsList[privateendpoint.NetworkInterfaceAttributes](pe.ref.Append("network_interface"))
}

func (pe privateEndpointAttributes) PrivateDnsZoneConfigs() terra.ListValue[privateendpoint.PrivateDnsZoneConfigsAttributes] {
	return terra.ReferenceAsList[privateendpoint.PrivateDnsZoneConfigsAttributes](pe.ref.Append("private_dns_zone_configs"))
}

func (pe privateEndpointAttributes) IpConfiguration() terra.ListValue[privateendpoint.IpConfigurationAttributes] {
	return terra.ReferenceAsList[privateendpoint.IpConfigurationAttributes](pe.ref.Append("ip_configuration"))
}

func (pe privateEndpointAttributes) PrivateDnsZoneGroup() terra.ListValue[privateendpoint.PrivateDnsZoneGroupAttributes] {
	return terra.ReferenceAsList[privateendpoint.PrivateDnsZoneGroupAttributes](pe.ref.Append("private_dns_zone_group"))
}

func (pe privateEndpointAttributes) PrivateServiceConnection() terra.ListValue[privateendpoint.PrivateServiceConnectionAttributes] {
	return terra.ReferenceAsList[privateendpoint.PrivateServiceConnectionAttributes](pe.ref.Append("private_service_connection"))
}

func (pe privateEndpointAttributes) Timeouts() privateendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privateendpoint.TimeoutsAttributes](pe.ref.Append("timeouts"))
}

type privateEndpointState struct {
	CustomNetworkInterfaceName string                                          `json:"custom_network_interface_name"`
	Id                         string                                          `json:"id"`
	Location                   string                                          `json:"location"`
	Name                       string                                          `json:"name"`
	ResourceGroupName          string                                          `json:"resource_group_name"`
	SubnetId                   string                                          `json:"subnet_id"`
	Tags                       map[string]string                               `json:"tags"`
	CustomDnsConfigs           []privateendpoint.CustomDnsConfigsState         `json:"custom_dns_configs"`
	NetworkInterface           []privateendpoint.NetworkInterfaceState         `json:"network_interface"`
	PrivateDnsZoneConfigs      []privateendpoint.PrivateDnsZoneConfigsState    `json:"private_dns_zone_configs"`
	IpConfiguration            []privateendpoint.IpConfigurationState          `json:"ip_configuration"`
	PrivateDnsZoneGroup        []privateendpoint.PrivateDnsZoneGroupState      `json:"private_dns_zone_group"`
	PrivateServiceConnection   []privateendpoint.PrivateServiceConnectionState `json:"private_service_connection"`
	Timeouts                   *privateendpoint.TimeoutsState                  `json:"timeouts"`
}
