// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datavpcendpoint "github.com/golingon/terraproviders/aws/4.60.0/datavpcendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVpcEndpoint creates a new instance of [DataVpcEndpoint].
func NewDataVpcEndpoint(name string, args DataVpcEndpointArgs) *DataVpcEndpoint {
	return &DataVpcEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpcEndpoint)(nil)

// DataVpcEndpoint represents the Terraform data resource aws_vpc_endpoint.
type DataVpcEndpoint struct {
	Name string
	Args DataVpcEndpointArgs
}

// DataSource returns the Terraform object type for [DataVpcEndpoint].
func (ve *DataVpcEndpoint) DataSource() string {
	return "aws_vpc_endpoint"
}

// LocalName returns the local name for [DataVpcEndpoint].
func (ve *DataVpcEndpoint) LocalName() string {
	return ve.Name
}

// Configuration returns the configuration (args) for [DataVpcEndpoint].
func (ve *DataVpcEndpoint) Configuration() interface{} {
	return ve.Args
}

// Attributes returns the attributes for [DataVpcEndpoint].
func (ve *DataVpcEndpoint) Attributes() dataVpcEndpointAttributes {
	return dataVpcEndpointAttributes{ref: terra.ReferenceDataResource(ve)}
}

// DataVpcEndpointArgs contains the configurations for aws_vpc_endpoint.
type DataVpcEndpointArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServiceName: string, optional
	ServiceName terra.StringValue `hcl:"service_name,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VpcId: string, optional
	VpcId terra.StringValue `hcl:"vpc_id,attr"`
	// DnsEntry: min=0
	DnsEntry []datavpcendpoint.DnsEntry `hcl:"dns_entry,block" validate:"min=0"`
	// DnsOptions: min=0
	DnsOptions []datavpcendpoint.DnsOptions `hcl:"dns_options,block" validate:"min=0"`
	// Filter: min=0
	Filter []datavpcendpoint.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datavpcendpoint.Timeouts `hcl:"timeouts,block"`
}
type dataVpcEndpointAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("arn"))
}

// CidrBlocks returns a reference to field cidr_blocks of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) CidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ve.ref.Append("cidr_blocks"))
}

// Id returns a reference to field id of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("id"))
}

// IpAddressType returns a reference to field ip_address_type of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) IpAddressType() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("ip_address_type"))
}

// NetworkInterfaceIds returns a reference to field network_interface_ids of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) NetworkInterfaceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ve.ref.Append("network_interface_ids"))
}

// OwnerId returns a reference to field owner_id of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("owner_id"))
}

// Policy returns a reference to field policy of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("policy"))
}

// PrefixListId returns a reference to field prefix_list_id of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) PrefixListId() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("prefix_list_id"))
}

// PrivateDnsEnabled returns a reference to field private_dns_enabled of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) PrivateDnsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ve.ref.Append("private_dns_enabled"))
}

// RequesterManaged returns a reference to field requester_managed of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) RequesterManaged() terra.BoolValue {
	return terra.ReferenceAsBool(ve.ref.Append("requester_managed"))
}

// RouteTableIds returns a reference to field route_table_ids of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) RouteTableIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ve.ref.Append("route_table_ids"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ve.ref.Append("security_group_ids"))
}

// ServiceName returns a reference to field service_name of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("service_name"))
}

// State returns a reference to field state of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("state"))
}

// SubnetIds returns a reference to field subnet_ids of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ve.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ve.ref.Append("tags"))
}

// VpcEndpointType returns a reference to field vpc_endpoint_type of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) VpcEndpointType() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("vpc_endpoint_type"))
}

// VpcId returns a reference to field vpc_id of aws_vpc_endpoint.
func (ve dataVpcEndpointAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("vpc_id"))
}

func (ve dataVpcEndpointAttributes) DnsEntry() terra.ListValue[datavpcendpoint.DnsEntryAttributes] {
	return terra.ReferenceAsList[datavpcendpoint.DnsEntryAttributes](ve.ref.Append("dns_entry"))
}

func (ve dataVpcEndpointAttributes) DnsOptions() terra.ListValue[datavpcendpoint.DnsOptionsAttributes] {
	return terra.ReferenceAsList[datavpcendpoint.DnsOptionsAttributes](ve.ref.Append("dns_options"))
}

func (ve dataVpcEndpointAttributes) Filter() terra.SetValue[datavpcendpoint.FilterAttributes] {
	return terra.ReferenceAsSet[datavpcendpoint.FilterAttributes](ve.ref.Append("filter"))
}

func (ve dataVpcEndpointAttributes) Timeouts() datavpcendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavpcendpoint.TimeoutsAttributes](ve.ref.Append("timeouts"))
}
