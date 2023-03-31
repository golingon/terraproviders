// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datavpcdhcpoptions "github.com/golingon/terraproviders/aws/4.60.0/datavpcdhcpoptions"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVpcDhcpOptions creates a new instance of [DataVpcDhcpOptions].
func NewDataVpcDhcpOptions(name string, args DataVpcDhcpOptionsArgs) *DataVpcDhcpOptions {
	return &DataVpcDhcpOptions{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpcDhcpOptions)(nil)

// DataVpcDhcpOptions represents the Terraform data resource aws_vpc_dhcp_options.
type DataVpcDhcpOptions struct {
	Name string
	Args DataVpcDhcpOptionsArgs
}

// DataSource returns the Terraform object type for [DataVpcDhcpOptions].
func (vdo *DataVpcDhcpOptions) DataSource() string {
	return "aws_vpc_dhcp_options"
}

// LocalName returns the local name for [DataVpcDhcpOptions].
func (vdo *DataVpcDhcpOptions) LocalName() string {
	return vdo.Name
}

// Configuration returns the configuration (args) for [DataVpcDhcpOptions].
func (vdo *DataVpcDhcpOptions) Configuration() interface{} {
	return vdo.Args
}

// Attributes returns the attributes for [DataVpcDhcpOptions].
func (vdo *DataVpcDhcpOptions) Attributes() dataVpcDhcpOptionsAttributes {
	return dataVpcDhcpOptionsAttributes{ref: terra.ReferenceDataResource(vdo)}
}

// DataVpcDhcpOptionsArgs contains the configurations for aws_vpc_dhcp_options.
type DataVpcDhcpOptionsArgs struct {
	// DhcpOptionsId: string, optional
	DhcpOptionsId terra.StringValue `hcl:"dhcp_options_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []datavpcdhcpoptions.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datavpcdhcpoptions.Timeouts `hcl:"timeouts,block"`
}
type dataVpcDhcpOptionsAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpc_dhcp_options.
func (vdo dataVpcDhcpOptionsAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(vdo.ref.Append("arn"))
}

// DhcpOptionsId returns a reference to field dhcp_options_id of aws_vpc_dhcp_options.
func (vdo dataVpcDhcpOptionsAttributes) DhcpOptionsId() terra.StringValue {
	return terra.ReferenceAsString(vdo.ref.Append("dhcp_options_id"))
}

// DomainName returns a reference to field domain_name of aws_vpc_dhcp_options.
func (vdo dataVpcDhcpOptionsAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(vdo.ref.Append("domain_name"))
}

// DomainNameServers returns a reference to field domain_name_servers of aws_vpc_dhcp_options.
func (vdo dataVpcDhcpOptionsAttributes) DomainNameServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vdo.ref.Append("domain_name_servers"))
}

// Id returns a reference to field id of aws_vpc_dhcp_options.
func (vdo dataVpcDhcpOptionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vdo.ref.Append("id"))
}

// NetbiosNameServers returns a reference to field netbios_name_servers of aws_vpc_dhcp_options.
func (vdo dataVpcDhcpOptionsAttributes) NetbiosNameServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vdo.ref.Append("netbios_name_servers"))
}

// NetbiosNodeType returns a reference to field netbios_node_type of aws_vpc_dhcp_options.
func (vdo dataVpcDhcpOptionsAttributes) NetbiosNodeType() terra.StringValue {
	return terra.ReferenceAsString(vdo.ref.Append("netbios_node_type"))
}

// NtpServers returns a reference to field ntp_servers of aws_vpc_dhcp_options.
func (vdo dataVpcDhcpOptionsAttributes) NtpServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vdo.ref.Append("ntp_servers"))
}

// OwnerId returns a reference to field owner_id of aws_vpc_dhcp_options.
func (vdo dataVpcDhcpOptionsAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(vdo.ref.Append("owner_id"))
}

// Tags returns a reference to field tags of aws_vpc_dhcp_options.
func (vdo dataVpcDhcpOptionsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vdo.ref.Append("tags"))
}

func (vdo dataVpcDhcpOptionsAttributes) Filter() terra.SetValue[datavpcdhcpoptions.FilterAttributes] {
	return terra.ReferenceAsSet[datavpcdhcpoptions.FilterAttributes](vdo.ref.Append("filter"))
}

func (vdo dataVpcDhcpOptionsAttributes) Timeouts() datavpcdhcpoptions.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavpcdhcpoptions.TimeoutsAttributes](vdo.ref.Append("timeouts"))
}
