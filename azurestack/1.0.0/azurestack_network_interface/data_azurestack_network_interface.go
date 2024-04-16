// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurestack_network_interface

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurestack_network_interface.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (ani *DataSource) DataSource() string {
	return "azurestack_network_interface"
}

// LocalName returns the local name for [DataSource].
func (ani *DataSource) LocalName() string {
	return ani.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (ani *DataSource) Configuration() interface{} {
	return ani.Args
}

// Attributes returns the attributes for [DataSource].
func (ani *DataSource) Attributes() dataAzurestackNetworkInterfaceAttributes {
	return dataAzurestackNetworkInterfaceAttributes{ref: terra.ReferenceDataSource(ani)}
}

// DataArgs contains the configurations for azurestack_network_interface.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurestackNetworkInterfaceAttributes struct {
	ref terra.Reference
}

// AppliedDnsServers returns a reference to field applied_dns_servers of azurestack_network_interface.
func (ani dataAzurestackNetworkInterfaceAttributes) AppliedDnsServers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ani.ref.Append("applied_dns_servers"))
}

// DnsServers returns a reference to field dns_servers of azurestack_network_interface.
func (ani dataAzurestackNetworkInterfaceAttributes) DnsServers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ani.ref.Append("dns_servers"))
}

// EnableIpForwarding returns a reference to field enable_ip_forwarding of azurestack_network_interface.
func (ani dataAzurestackNetworkInterfaceAttributes) EnableIpForwarding() terra.BoolValue {
	return terra.ReferenceAsBool(ani.ref.Append("enable_ip_forwarding"))
}

// Id returns a reference to field id of azurestack_network_interface.
func (ani dataAzurestackNetworkInterfaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ani.ref.Append("id"))
}

// InternalDnsNameLabel returns a reference to field internal_dns_name_label of azurestack_network_interface.
func (ani dataAzurestackNetworkInterfaceAttributes) InternalDnsNameLabel() terra.StringValue {
	return terra.ReferenceAsString(ani.ref.Append("internal_dns_name_label"))
}

// Location returns a reference to field location of azurestack_network_interface.
func (ani dataAzurestackNetworkInterfaceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ani.ref.Append("location"))
}

// MacAddress returns a reference to field mac_address of azurestack_network_interface.
func (ani dataAzurestackNetworkInterfaceAttributes) MacAddress() terra.StringValue {
	return terra.ReferenceAsString(ani.ref.Append("mac_address"))
}

// Name returns a reference to field name of azurestack_network_interface.
func (ani dataAzurestackNetworkInterfaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ani.ref.Append("name"))
}

// NetworkSecurityGroupId returns a reference to field network_security_group_id of azurestack_network_interface.
func (ani dataAzurestackNetworkInterfaceAttributes) NetworkSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(ani.ref.Append("network_security_group_id"))
}

// PrivateIpAddress returns a reference to field private_ip_address of azurestack_network_interface.
func (ani dataAzurestackNetworkInterfaceAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceAsString(ani.ref.Append("private_ip_address"))
}

// PrivateIpAddresses returns a reference to field private_ip_addresses of azurestack_network_interface.
func (ani dataAzurestackNetworkInterfaceAttributes) PrivateIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ani.ref.Append("private_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_network_interface.
func (ani dataAzurestackNetworkInterfaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ani.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurestack_network_interface.
func (ani dataAzurestackNetworkInterfaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ani.ref.Append("tags"))
}

// VirtualMachineId returns a reference to field virtual_machine_id of azurestack_network_interface.
func (ani dataAzurestackNetworkInterfaceAttributes) VirtualMachineId() terra.StringValue {
	return terra.ReferenceAsString(ani.ref.Append("virtual_machine_id"))
}

func (ani dataAzurestackNetworkInterfaceAttributes) IpConfiguration() terra.ListValue[DataIpConfigurationAttributes] {
	return terra.ReferenceAsList[DataIpConfigurationAttributes](ani.ref.Append("ip_configuration"))
}

func (ani dataAzurestackNetworkInterfaceAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](ani.ref.Append("timeouts"))
}
