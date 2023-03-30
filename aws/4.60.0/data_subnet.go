// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datasubnet "github.com/golingon/terraproviders/aws/4.60.0/datasubnet"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataSubnet(name string, args DataSubnetArgs) *DataSubnet {
	return &DataSubnet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSubnet)(nil)

type DataSubnet struct {
	Name string
	Args DataSubnetArgs
}

func (s *DataSubnet) DataSource() string {
	return "aws_subnet"
}

func (s *DataSubnet) LocalName() string {
	return s.Name
}

func (s *DataSubnet) Configuration() interface{} {
	return s.Args
}

func (s *DataSubnet) Attributes() dataSubnetAttributes {
	return dataSubnetAttributes{ref: terra.ReferenceDataResource(s)}
}

type DataSubnetArgs struct {
	// AvailabilityZone: string, optional
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr"`
	// AvailabilityZoneId: string, optional
	AvailabilityZoneId terra.StringValue `hcl:"availability_zone_id,attr"`
	// CidrBlock: string, optional
	CidrBlock terra.StringValue `hcl:"cidr_block,attr"`
	// DefaultForAz: bool, optional
	DefaultForAz terra.BoolValue `hcl:"default_for_az,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Ipv6CidrBlock: string, optional
	Ipv6CidrBlock terra.StringValue `hcl:"ipv6_cidr_block,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VpcId: string, optional
	VpcId terra.StringValue `hcl:"vpc_id,attr"`
	// Filter: min=0
	Filter []datasubnet.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datasubnet.Timeouts `hcl:"timeouts,block"`
}
type dataSubnetAttributes struct {
	ref terra.Reference
}

func (s dataSubnetAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("arn"))
}

func (s dataSubnetAttributes) AssignIpv6AddressOnCreation() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("assign_ipv6_address_on_creation"))
}

func (s dataSubnetAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("availability_zone"))
}

func (s dataSubnetAttributes) AvailabilityZoneId() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("availability_zone_id"))
}

func (s dataSubnetAttributes) AvailableIpAddressCount() terra.NumberValue {
	return terra.ReferenceNumber(s.ref.Append("available_ip_address_count"))
}

func (s dataSubnetAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("cidr_block"))
}

func (s dataSubnetAttributes) CustomerOwnedIpv4Pool() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("customer_owned_ipv4_pool"))
}

func (s dataSubnetAttributes) DefaultForAz() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("default_for_az"))
}

func (s dataSubnetAttributes) EnableDns64() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("enable_dns64"))
}

func (s dataSubnetAttributes) EnableResourceNameDnsARecordOnLaunch() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("enable_resource_name_dns_a_record_on_launch"))
}

func (s dataSubnetAttributes) EnableResourceNameDnsAaaaRecordOnLaunch() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("enable_resource_name_dns_aaaa_record_on_launch"))
}

func (s dataSubnetAttributes) Id() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("id"))
}

func (s dataSubnetAttributes) Ipv6CidrBlock() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("ipv6_cidr_block"))
}

func (s dataSubnetAttributes) Ipv6CidrBlockAssociationId() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("ipv6_cidr_block_association_id"))
}

func (s dataSubnetAttributes) Ipv6Native() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("ipv6_native"))
}

func (s dataSubnetAttributes) MapCustomerOwnedIpOnLaunch() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("map_customer_owned_ip_on_launch"))
}

func (s dataSubnetAttributes) MapPublicIpOnLaunch() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("map_public_ip_on_launch"))
}

func (s dataSubnetAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("outpost_arn"))
}

func (s dataSubnetAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("owner_id"))
}

func (s dataSubnetAttributes) PrivateDnsHostnameTypeOnLaunch() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("private_dns_hostname_type_on_launch"))
}

func (s dataSubnetAttributes) State() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("state"))
}

func (s dataSubnetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](s.ref.Append("tags"))
}

func (s dataSubnetAttributes) VpcId() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("vpc_id"))
}

func (s dataSubnetAttributes) Filter() terra.SetValue[datasubnet.FilterAttributes] {
	return terra.ReferenceSet[datasubnet.FilterAttributes](s.ref.Append("filter"))
}

func (s dataSubnetAttributes) Timeouts() datasubnet.TimeoutsAttributes {
	return terra.ReferenceSingle[datasubnet.TimeoutsAttributes](s.ref.Append("timeouts"))
}
