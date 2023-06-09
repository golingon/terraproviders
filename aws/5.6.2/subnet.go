// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	subnet "github.com/golingon/terraproviders/aws/5.6.2/subnet"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSubnet creates a new instance of [Subnet].
func NewSubnet(name string, args SubnetArgs) *Subnet {
	return &Subnet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Subnet)(nil)

// Subnet represents the Terraform resource aws_subnet.
type Subnet struct {
	Name      string
	Args      SubnetArgs
	state     *subnetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Subnet].
func (s *Subnet) Type() string {
	return "aws_subnet"
}

// LocalName returns the local name for [Subnet].
func (s *Subnet) LocalName() string {
	return s.Name
}

// Configuration returns the configuration (args) for [Subnet].
func (s *Subnet) Configuration() interface{} {
	return s.Args
}

// DependOn is used for other resources to depend on [Subnet].
func (s *Subnet) DependOn() terra.Reference {
	return terra.ReferenceResource(s)
}

// Dependencies returns the list of resources [Subnet] depends_on.
func (s *Subnet) Dependencies() terra.Dependencies {
	return s.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Subnet].
func (s *Subnet) LifecycleManagement() *terra.Lifecycle {
	return s.Lifecycle
}

// Attributes returns the attributes for [Subnet].
func (s *Subnet) Attributes() subnetAttributes {
	return subnetAttributes{ref: terra.ReferenceResource(s)}
}

// ImportState imports the given attribute values into [Subnet]'s state.
func (s *Subnet) ImportState(av io.Reader) error {
	s.state = &subnetState{}
	if err := json.NewDecoder(av).Decode(s.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", s.Type(), s.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Subnet] has state.
func (s *Subnet) State() (*subnetState, bool) {
	return s.state, s.state != nil
}

// StateMust returns the state for [Subnet]. Panics if the state is nil.
func (s *Subnet) StateMust() *subnetState {
	if s.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", s.Type(), s.LocalName()))
	}
	return s.state
}

// SubnetArgs contains the configurations for aws_subnet.
type SubnetArgs struct {
	// AssignIpv6AddressOnCreation: bool, optional
	AssignIpv6AddressOnCreation terra.BoolValue `hcl:"assign_ipv6_address_on_creation,attr"`
	// AvailabilityZone: string, optional
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr"`
	// AvailabilityZoneId: string, optional
	AvailabilityZoneId terra.StringValue `hcl:"availability_zone_id,attr"`
	// CidrBlock: string, optional
	CidrBlock terra.StringValue `hcl:"cidr_block,attr"`
	// CustomerOwnedIpv4Pool: string, optional
	CustomerOwnedIpv4Pool terra.StringValue `hcl:"customer_owned_ipv4_pool,attr"`
	// EnableDns64: bool, optional
	EnableDns64 terra.BoolValue `hcl:"enable_dns64,attr"`
	// EnableLniAtDeviceIndex: number, optional
	EnableLniAtDeviceIndex terra.NumberValue `hcl:"enable_lni_at_device_index,attr"`
	// EnableResourceNameDnsARecordOnLaunch: bool, optional
	EnableResourceNameDnsARecordOnLaunch terra.BoolValue `hcl:"enable_resource_name_dns_a_record_on_launch,attr"`
	// EnableResourceNameDnsAaaaRecordOnLaunch: bool, optional
	EnableResourceNameDnsAaaaRecordOnLaunch terra.BoolValue `hcl:"enable_resource_name_dns_aaaa_record_on_launch,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Ipv6CidrBlock: string, optional
	Ipv6CidrBlock terra.StringValue `hcl:"ipv6_cidr_block,attr"`
	// Ipv6Native: bool, optional
	Ipv6Native terra.BoolValue `hcl:"ipv6_native,attr"`
	// MapCustomerOwnedIpOnLaunch: bool, optional
	MapCustomerOwnedIpOnLaunch terra.BoolValue `hcl:"map_customer_owned_ip_on_launch,attr"`
	// MapPublicIpOnLaunch: bool, optional
	MapPublicIpOnLaunch terra.BoolValue `hcl:"map_public_ip_on_launch,attr"`
	// OutpostArn: string, optional
	OutpostArn terra.StringValue `hcl:"outpost_arn,attr"`
	// PrivateDnsHostnameTypeOnLaunch: string, optional
	PrivateDnsHostnameTypeOnLaunch terra.StringValue `hcl:"private_dns_hostname_type_on_launch,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *subnet.Timeouts `hcl:"timeouts,block"`
}
type subnetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_subnet.
func (s subnetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("arn"))
}

// AssignIpv6AddressOnCreation returns a reference to field assign_ipv6_address_on_creation of aws_subnet.
func (s subnetAttributes) AssignIpv6AddressOnCreation() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("assign_ipv6_address_on_creation"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_subnet.
func (s subnetAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("availability_zone"))
}

// AvailabilityZoneId returns a reference to field availability_zone_id of aws_subnet.
func (s subnetAttributes) AvailabilityZoneId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("availability_zone_id"))
}

// CidrBlock returns a reference to field cidr_block of aws_subnet.
func (s subnetAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("cidr_block"))
}

// CustomerOwnedIpv4Pool returns a reference to field customer_owned_ipv4_pool of aws_subnet.
func (s subnetAttributes) CustomerOwnedIpv4Pool() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("customer_owned_ipv4_pool"))
}

// EnableDns64 returns a reference to field enable_dns64 of aws_subnet.
func (s subnetAttributes) EnableDns64() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("enable_dns64"))
}

// EnableLniAtDeviceIndex returns a reference to field enable_lni_at_device_index of aws_subnet.
func (s subnetAttributes) EnableLniAtDeviceIndex() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("enable_lni_at_device_index"))
}

// EnableResourceNameDnsARecordOnLaunch returns a reference to field enable_resource_name_dns_a_record_on_launch of aws_subnet.
func (s subnetAttributes) EnableResourceNameDnsARecordOnLaunch() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("enable_resource_name_dns_a_record_on_launch"))
}

// EnableResourceNameDnsAaaaRecordOnLaunch returns a reference to field enable_resource_name_dns_aaaa_record_on_launch of aws_subnet.
func (s subnetAttributes) EnableResourceNameDnsAaaaRecordOnLaunch() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("enable_resource_name_dns_aaaa_record_on_launch"))
}

// Id returns a reference to field id of aws_subnet.
func (s subnetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("id"))
}

// Ipv6CidrBlock returns a reference to field ipv6_cidr_block of aws_subnet.
func (s subnetAttributes) Ipv6CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("ipv6_cidr_block"))
}

// Ipv6CidrBlockAssociationId returns a reference to field ipv6_cidr_block_association_id of aws_subnet.
func (s subnetAttributes) Ipv6CidrBlockAssociationId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("ipv6_cidr_block_association_id"))
}

// Ipv6Native returns a reference to field ipv6_native of aws_subnet.
func (s subnetAttributes) Ipv6Native() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("ipv6_native"))
}

// MapCustomerOwnedIpOnLaunch returns a reference to field map_customer_owned_ip_on_launch of aws_subnet.
func (s subnetAttributes) MapCustomerOwnedIpOnLaunch() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("map_customer_owned_ip_on_launch"))
}

// MapPublicIpOnLaunch returns a reference to field map_public_ip_on_launch of aws_subnet.
func (s subnetAttributes) MapPublicIpOnLaunch() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("map_public_ip_on_launch"))
}

// OutpostArn returns a reference to field outpost_arn of aws_subnet.
func (s subnetAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("outpost_arn"))
}

// OwnerId returns a reference to field owner_id of aws_subnet.
func (s subnetAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("owner_id"))
}

// PrivateDnsHostnameTypeOnLaunch returns a reference to field private_dns_hostname_type_on_launch of aws_subnet.
func (s subnetAttributes) PrivateDnsHostnameTypeOnLaunch() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("private_dns_hostname_type_on_launch"))
}

// Tags returns a reference to field tags of aws_subnet.
func (s subnetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](s.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_subnet.
func (s subnetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](s.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_subnet.
func (s subnetAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("vpc_id"))
}

func (s subnetAttributes) Timeouts() subnet.TimeoutsAttributes {
	return terra.ReferenceAsSingle[subnet.TimeoutsAttributes](s.ref.Append("timeouts"))
}

type subnetState struct {
	Arn                                     string                `json:"arn"`
	AssignIpv6AddressOnCreation             bool                  `json:"assign_ipv6_address_on_creation"`
	AvailabilityZone                        string                `json:"availability_zone"`
	AvailabilityZoneId                      string                `json:"availability_zone_id"`
	CidrBlock                               string                `json:"cidr_block"`
	CustomerOwnedIpv4Pool                   string                `json:"customer_owned_ipv4_pool"`
	EnableDns64                             bool                  `json:"enable_dns64"`
	EnableLniAtDeviceIndex                  float64               `json:"enable_lni_at_device_index"`
	EnableResourceNameDnsARecordOnLaunch    bool                  `json:"enable_resource_name_dns_a_record_on_launch"`
	EnableResourceNameDnsAaaaRecordOnLaunch bool                  `json:"enable_resource_name_dns_aaaa_record_on_launch"`
	Id                                      string                `json:"id"`
	Ipv6CidrBlock                           string                `json:"ipv6_cidr_block"`
	Ipv6CidrBlockAssociationId              string                `json:"ipv6_cidr_block_association_id"`
	Ipv6Native                              bool                  `json:"ipv6_native"`
	MapCustomerOwnedIpOnLaunch              bool                  `json:"map_customer_owned_ip_on_launch"`
	MapPublicIpOnLaunch                     bool                  `json:"map_public_ip_on_launch"`
	OutpostArn                              string                `json:"outpost_arn"`
	OwnerId                                 string                `json:"owner_id"`
	PrivateDnsHostnameTypeOnLaunch          string                `json:"private_dns_hostname_type_on_launch"`
	Tags                                    map[string]string     `json:"tags"`
	TagsAll                                 map[string]string     `json:"tags_all"`
	VpcId                                   string                `json:"vpc_id"`
	Timeouts                                *subnet.TimeoutsState `json:"timeouts"`
}
