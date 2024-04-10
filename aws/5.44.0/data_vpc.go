// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"github.com/golingon/lingon/pkg/terra"
	datavpc "github.com/golingon/terraproviders/aws/5.44.0/datavpc"
)

// NewDataVpc creates a new instance of [DataVpc].
func NewDataVpc(name string, args DataVpcArgs) *DataVpc {
	return &DataVpc{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpc)(nil)

// DataVpc represents the Terraform data resource aws_vpc.
type DataVpc struct {
	Name string
	Args DataVpcArgs
}

// DataSource returns the Terraform object type for [DataVpc].
func (v *DataVpc) DataSource() string {
	return "aws_vpc"
}

// LocalName returns the local name for [DataVpc].
func (v *DataVpc) LocalName() string {
	return v.Name
}

// Configuration returns the configuration (args) for [DataVpc].
func (v *DataVpc) Configuration() interface{} {
	return v.Args
}

// Attributes returns the attributes for [DataVpc].
func (v *DataVpc) Attributes() dataVpcAttributes {
	return dataVpcAttributes{ref: terra.ReferenceDataResource(v)}
}

// DataVpcArgs contains the configurations for aws_vpc.
type DataVpcArgs struct {
	// CidrBlock: string, optional
	CidrBlock terra.StringValue `hcl:"cidr_block,attr"`
	// Default: bool, optional
	Default terra.BoolValue `hcl:"default,attr"`
	// DhcpOptionsId: string, optional
	DhcpOptionsId terra.StringValue `hcl:"dhcp_options_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// CidrBlockAssociations: min=0
	CidrBlockAssociations []datavpc.CidrBlockAssociations `hcl:"cidr_block_associations,block" validate:"min=0"`
	// Filter: min=0
	Filter []datavpc.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datavpc.Timeouts `hcl:"timeouts,block"`
}
type dataVpcAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpc.
func (v dataVpcAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("arn"))
}

// CidrBlock returns a reference to field cidr_block of aws_vpc.
func (v dataVpcAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("cidr_block"))
}

// Default returns a reference to field default of aws_vpc.
func (v dataVpcAttributes) Default() terra.BoolValue {
	return terra.ReferenceAsBool(v.ref.Append("default"))
}

// DhcpOptionsId returns a reference to field dhcp_options_id of aws_vpc.
func (v dataVpcAttributes) DhcpOptionsId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("dhcp_options_id"))
}

// EnableDnsHostnames returns a reference to field enable_dns_hostnames of aws_vpc.
func (v dataVpcAttributes) EnableDnsHostnames() terra.BoolValue {
	return terra.ReferenceAsBool(v.ref.Append("enable_dns_hostnames"))
}

// EnableDnsSupport returns a reference to field enable_dns_support of aws_vpc.
func (v dataVpcAttributes) EnableDnsSupport() terra.BoolValue {
	return terra.ReferenceAsBool(v.ref.Append("enable_dns_support"))
}

// EnableNetworkAddressUsageMetrics returns a reference to field enable_network_address_usage_metrics of aws_vpc.
func (v dataVpcAttributes) EnableNetworkAddressUsageMetrics() terra.BoolValue {
	return terra.ReferenceAsBool(v.ref.Append("enable_network_address_usage_metrics"))
}

// Id returns a reference to field id of aws_vpc.
func (v dataVpcAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("id"))
}

// InstanceTenancy returns a reference to field instance_tenancy of aws_vpc.
func (v dataVpcAttributes) InstanceTenancy() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("instance_tenancy"))
}

// Ipv6AssociationId returns a reference to field ipv6_association_id of aws_vpc.
func (v dataVpcAttributes) Ipv6AssociationId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("ipv6_association_id"))
}

// Ipv6CidrBlock returns a reference to field ipv6_cidr_block of aws_vpc.
func (v dataVpcAttributes) Ipv6CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("ipv6_cidr_block"))
}

// MainRouteTableId returns a reference to field main_route_table_id of aws_vpc.
func (v dataVpcAttributes) MainRouteTableId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("main_route_table_id"))
}

// OwnerId returns a reference to field owner_id of aws_vpc.
func (v dataVpcAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("owner_id"))
}

// State returns a reference to field state of aws_vpc.
func (v dataVpcAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_vpc.
func (v dataVpcAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](v.ref.Append("tags"))
}

func (v dataVpcAttributes) CidrBlockAssociations() terra.ListValue[datavpc.CidrBlockAssociationsAttributes] {
	return terra.ReferenceAsList[datavpc.CidrBlockAssociationsAttributes](v.ref.Append("cidr_block_associations"))
}

func (v dataVpcAttributes) Filter() terra.SetValue[datavpc.FilterAttributes] {
	return terra.ReferenceAsSet[datavpc.FilterAttributes](v.ref.Append("filter"))
}

func (v dataVpcAttributes) Timeouts() datavpc.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavpc.TimeoutsAttributes](v.ref.Append("timeouts"))
}
