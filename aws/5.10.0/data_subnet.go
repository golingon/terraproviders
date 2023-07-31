// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datasubnet "github.com/golingon/terraproviders/aws/5.10.0/datasubnet"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSubnet creates a new instance of [DataSubnet].
func NewDataSubnet(name string, args DataSubnetArgs) *DataSubnet {
	return &DataSubnet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSubnet)(nil)

// DataSubnet represents the Terraform data resource aws_subnet.
type DataSubnet struct {
	Name string
	Args DataSubnetArgs
}

// DataSource returns the Terraform object type for [DataSubnet].
func (s *DataSubnet) DataSource() string {
	return "aws_subnet"
}

// LocalName returns the local name for [DataSubnet].
func (s *DataSubnet) LocalName() string {
	return s.Name
}

// Configuration returns the configuration (args) for [DataSubnet].
func (s *DataSubnet) Configuration() interface{} {
	return s.Args
}

// Attributes returns the attributes for [DataSubnet].
func (s *DataSubnet) Attributes() dataSubnetAttributes {
	return dataSubnetAttributes{ref: terra.ReferenceDataResource(s)}
}

// DataSubnetArgs contains the configurations for aws_subnet.
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

// Arn returns a reference to field arn of aws_subnet.
func (s dataSubnetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("arn"))
}

// AssignIpv6AddressOnCreation returns a reference to field assign_ipv6_address_on_creation of aws_subnet.
func (s dataSubnetAttributes) AssignIpv6AddressOnCreation() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("assign_ipv6_address_on_creation"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_subnet.
func (s dataSubnetAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("availability_zone"))
}

// AvailabilityZoneId returns a reference to field availability_zone_id of aws_subnet.
func (s dataSubnetAttributes) AvailabilityZoneId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("availability_zone_id"))
}

// AvailableIpAddressCount returns a reference to field available_ip_address_count of aws_subnet.
func (s dataSubnetAttributes) AvailableIpAddressCount() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("available_ip_address_count"))
}

// CidrBlock returns a reference to field cidr_block of aws_subnet.
func (s dataSubnetAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("cidr_block"))
}

// CustomerOwnedIpv4Pool returns a reference to field customer_owned_ipv4_pool of aws_subnet.
func (s dataSubnetAttributes) CustomerOwnedIpv4Pool() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("customer_owned_ipv4_pool"))
}

// DefaultForAz returns a reference to field default_for_az of aws_subnet.
func (s dataSubnetAttributes) DefaultForAz() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("default_for_az"))
}

// EnableDns64 returns a reference to field enable_dns64 of aws_subnet.
func (s dataSubnetAttributes) EnableDns64() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("enable_dns64"))
}

// EnableLniAtDeviceIndex returns a reference to field enable_lni_at_device_index of aws_subnet.
func (s dataSubnetAttributes) EnableLniAtDeviceIndex() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("enable_lni_at_device_index"))
}

// EnableResourceNameDnsARecordOnLaunch returns a reference to field enable_resource_name_dns_a_record_on_launch of aws_subnet.
func (s dataSubnetAttributes) EnableResourceNameDnsARecordOnLaunch() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("enable_resource_name_dns_a_record_on_launch"))
}

// EnableResourceNameDnsAaaaRecordOnLaunch returns a reference to field enable_resource_name_dns_aaaa_record_on_launch of aws_subnet.
func (s dataSubnetAttributes) EnableResourceNameDnsAaaaRecordOnLaunch() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("enable_resource_name_dns_aaaa_record_on_launch"))
}

// Id returns a reference to field id of aws_subnet.
func (s dataSubnetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("id"))
}

// Ipv6CidrBlock returns a reference to field ipv6_cidr_block of aws_subnet.
func (s dataSubnetAttributes) Ipv6CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("ipv6_cidr_block"))
}

// Ipv6CidrBlockAssociationId returns a reference to field ipv6_cidr_block_association_id of aws_subnet.
func (s dataSubnetAttributes) Ipv6CidrBlockAssociationId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("ipv6_cidr_block_association_id"))
}

// Ipv6Native returns a reference to field ipv6_native of aws_subnet.
func (s dataSubnetAttributes) Ipv6Native() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("ipv6_native"))
}

// MapCustomerOwnedIpOnLaunch returns a reference to field map_customer_owned_ip_on_launch of aws_subnet.
func (s dataSubnetAttributes) MapCustomerOwnedIpOnLaunch() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("map_customer_owned_ip_on_launch"))
}

// MapPublicIpOnLaunch returns a reference to field map_public_ip_on_launch of aws_subnet.
func (s dataSubnetAttributes) MapPublicIpOnLaunch() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("map_public_ip_on_launch"))
}

// OutpostArn returns a reference to field outpost_arn of aws_subnet.
func (s dataSubnetAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("outpost_arn"))
}

// OwnerId returns a reference to field owner_id of aws_subnet.
func (s dataSubnetAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("owner_id"))
}

// PrivateDnsHostnameTypeOnLaunch returns a reference to field private_dns_hostname_type_on_launch of aws_subnet.
func (s dataSubnetAttributes) PrivateDnsHostnameTypeOnLaunch() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("private_dns_hostname_type_on_launch"))
}

// State returns a reference to field state of aws_subnet.
func (s dataSubnetAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_subnet.
func (s dataSubnetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](s.ref.Append("tags"))
}

// VpcId returns a reference to field vpc_id of aws_subnet.
func (s dataSubnetAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("vpc_id"))
}

func (s dataSubnetAttributes) Filter() terra.SetValue[datasubnet.FilterAttributes] {
	return terra.ReferenceAsSet[datasubnet.FilterAttributes](s.ref.Append("filter"))
}

func (s dataSubnetAttributes) Timeouts() datasubnet.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasubnet.TimeoutsAttributes](s.ref.Append("timeouts"))
}
