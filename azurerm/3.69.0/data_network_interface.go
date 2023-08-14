// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datanetworkinterface "github.com/golingon/terraproviders/azurerm/3.69.0/datanetworkinterface"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataNetworkInterface creates a new instance of [DataNetworkInterface].
func NewDataNetworkInterface(name string, args DataNetworkInterfaceArgs) *DataNetworkInterface {
	return &DataNetworkInterface{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkInterface)(nil)

// DataNetworkInterface represents the Terraform data resource azurerm_network_interface.
type DataNetworkInterface struct {
	Name string
	Args DataNetworkInterfaceArgs
}

// DataSource returns the Terraform object type for [DataNetworkInterface].
func (ni *DataNetworkInterface) DataSource() string {
	return "azurerm_network_interface"
}

// LocalName returns the local name for [DataNetworkInterface].
func (ni *DataNetworkInterface) LocalName() string {
	return ni.Name
}

// Configuration returns the configuration (args) for [DataNetworkInterface].
func (ni *DataNetworkInterface) Configuration() interface{} {
	return ni.Args
}

// Attributes returns the attributes for [DataNetworkInterface].
func (ni *DataNetworkInterface) Attributes() dataNetworkInterfaceAttributes {
	return dataNetworkInterfaceAttributes{ref: terra.ReferenceDataResource(ni)}
}

// DataNetworkInterfaceArgs contains the configurations for azurerm_network_interface.
type DataNetworkInterfaceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// IpConfiguration: min=0
	IpConfiguration []datanetworkinterface.IpConfiguration `hcl:"ip_configuration,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datanetworkinterface.Timeouts `hcl:"timeouts,block"`
}
type dataNetworkInterfaceAttributes struct {
	ref terra.Reference
}

// AppliedDnsServers returns a reference to field applied_dns_servers of azurerm_network_interface.
func (ni dataNetworkInterfaceAttributes) AppliedDnsServers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ni.ref.Append("applied_dns_servers"))
}

// DnsServers returns a reference to field dns_servers of azurerm_network_interface.
func (ni dataNetworkInterfaceAttributes) DnsServers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ni.ref.Append("dns_servers"))
}

// EnableAcceleratedNetworking returns a reference to field enable_accelerated_networking of azurerm_network_interface.
func (ni dataNetworkInterfaceAttributes) EnableAcceleratedNetworking() terra.BoolValue {
	return terra.ReferenceAsBool(ni.ref.Append("enable_accelerated_networking"))
}

// EnableIpForwarding returns a reference to field enable_ip_forwarding of azurerm_network_interface.
func (ni dataNetworkInterfaceAttributes) EnableIpForwarding() terra.BoolValue {
	return terra.ReferenceAsBool(ni.ref.Append("enable_ip_forwarding"))
}

// Id returns a reference to field id of azurerm_network_interface.
func (ni dataNetworkInterfaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("id"))
}

// InternalDnsNameLabel returns a reference to field internal_dns_name_label of azurerm_network_interface.
func (ni dataNetworkInterfaceAttributes) InternalDnsNameLabel() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("internal_dns_name_label"))
}

// Location returns a reference to field location of azurerm_network_interface.
func (ni dataNetworkInterfaceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("location"))
}

// MacAddress returns a reference to field mac_address of azurerm_network_interface.
func (ni dataNetworkInterfaceAttributes) MacAddress() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("mac_address"))
}

// Name returns a reference to field name of azurerm_network_interface.
func (ni dataNetworkInterfaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("name"))
}

// NetworkSecurityGroupId returns a reference to field network_security_group_id of azurerm_network_interface.
func (ni dataNetworkInterfaceAttributes) NetworkSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("network_security_group_id"))
}

// PrivateIpAddress returns a reference to field private_ip_address of azurerm_network_interface.
func (ni dataNetworkInterfaceAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("private_ip_address"))
}

// PrivateIpAddresses returns a reference to field private_ip_addresses of azurerm_network_interface.
func (ni dataNetworkInterfaceAttributes) PrivateIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ni.ref.Append("private_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_network_interface.
func (ni dataNetworkInterfaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_network_interface.
func (ni dataNetworkInterfaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ni.ref.Append("tags"))
}

// VirtualMachineId returns a reference to field virtual_machine_id of azurerm_network_interface.
func (ni dataNetworkInterfaceAttributes) VirtualMachineId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("virtual_machine_id"))
}

func (ni dataNetworkInterfaceAttributes) IpConfiguration() terra.ListValue[datanetworkinterface.IpConfigurationAttributes] {
	return terra.ReferenceAsList[datanetworkinterface.IpConfigurationAttributes](ni.ref.Append("ip_configuration"))
}

func (ni dataNetworkInterfaceAttributes) Timeouts() datanetworkinterface.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datanetworkinterface.TimeoutsAttributes](ni.ref.Append("timeouts"))
}
