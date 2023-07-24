// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	defaultsubnet "github.com/golingon/terraproviders/aws/5.9.0/defaultsubnet"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDefaultSubnet creates a new instance of [DefaultSubnet].
func NewDefaultSubnet(name string, args DefaultSubnetArgs) *DefaultSubnet {
	return &DefaultSubnet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DefaultSubnet)(nil)

// DefaultSubnet represents the Terraform resource aws_default_subnet.
type DefaultSubnet struct {
	Name      string
	Args      DefaultSubnetArgs
	state     *defaultSubnetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DefaultSubnet].
func (ds *DefaultSubnet) Type() string {
	return "aws_default_subnet"
}

// LocalName returns the local name for [DefaultSubnet].
func (ds *DefaultSubnet) LocalName() string {
	return ds.Name
}

// Configuration returns the configuration (args) for [DefaultSubnet].
func (ds *DefaultSubnet) Configuration() interface{} {
	return ds.Args
}

// DependOn is used for other resources to depend on [DefaultSubnet].
func (ds *DefaultSubnet) DependOn() terra.Reference {
	return terra.ReferenceResource(ds)
}

// Dependencies returns the list of resources [DefaultSubnet] depends_on.
func (ds *DefaultSubnet) Dependencies() terra.Dependencies {
	return ds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DefaultSubnet].
func (ds *DefaultSubnet) LifecycleManagement() *terra.Lifecycle {
	return ds.Lifecycle
}

// Attributes returns the attributes for [DefaultSubnet].
func (ds *DefaultSubnet) Attributes() defaultSubnetAttributes {
	return defaultSubnetAttributes{ref: terra.ReferenceResource(ds)}
}

// ImportState imports the given attribute values into [DefaultSubnet]'s state.
func (ds *DefaultSubnet) ImportState(av io.Reader) error {
	ds.state = &defaultSubnetState{}
	if err := json.NewDecoder(av).Decode(ds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ds.Type(), ds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DefaultSubnet] has state.
func (ds *DefaultSubnet) State() (*defaultSubnetState, bool) {
	return ds.state, ds.state != nil
}

// StateMust returns the state for [DefaultSubnet]. Panics if the state is nil.
func (ds *DefaultSubnet) StateMust() *defaultSubnetState {
	if ds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ds.Type(), ds.LocalName()))
	}
	return ds.state
}

// DefaultSubnetArgs contains the configurations for aws_default_subnet.
type DefaultSubnetArgs struct {
	// AssignIpv6AddressOnCreation: bool, optional
	AssignIpv6AddressOnCreation terra.BoolValue `hcl:"assign_ipv6_address_on_creation,attr"`
	// AvailabilityZone: string, required
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr" validate:"required"`
	// CustomerOwnedIpv4Pool: string, optional
	CustomerOwnedIpv4Pool terra.StringValue `hcl:"customer_owned_ipv4_pool,attr"`
	// EnableDns64: bool, optional
	EnableDns64 terra.BoolValue `hcl:"enable_dns64,attr"`
	// EnableResourceNameDnsARecordOnLaunch: bool, optional
	EnableResourceNameDnsARecordOnLaunch terra.BoolValue `hcl:"enable_resource_name_dns_a_record_on_launch,attr"`
	// EnableResourceNameDnsAaaaRecordOnLaunch: bool, optional
	EnableResourceNameDnsAaaaRecordOnLaunch terra.BoolValue `hcl:"enable_resource_name_dns_aaaa_record_on_launch,attr"`
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
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
	// PrivateDnsHostnameTypeOnLaunch: string, optional
	PrivateDnsHostnameTypeOnLaunch terra.StringValue `hcl:"private_dns_hostname_type_on_launch,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *defaultsubnet.Timeouts `hcl:"timeouts,block"`
}
type defaultSubnetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_default_subnet.
func (ds defaultSubnetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("arn"))
}

// AssignIpv6AddressOnCreation returns a reference to field assign_ipv6_address_on_creation of aws_default_subnet.
func (ds defaultSubnetAttributes) AssignIpv6AddressOnCreation() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("assign_ipv6_address_on_creation"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_default_subnet.
func (ds defaultSubnetAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("availability_zone"))
}

// AvailabilityZoneId returns a reference to field availability_zone_id of aws_default_subnet.
func (ds defaultSubnetAttributes) AvailabilityZoneId() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("availability_zone_id"))
}

// CidrBlock returns a reference to field cidr_block of aws_default_subnet.
func (ds defaultSubnetAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("cidr_block"))
}

// CustomerOwnedIpv4Pool returns a reference to field customer_owned_ipv4_pool of aws_default_subnet.
func (ds defaultSubnetAttributes) CustomerOwnedIpv4Pool() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("customer_owned_ipv4_pool"))
}

// EnableDns64 returns a reference to field enable_dns64 of aws_default_subnet.
func (ds defaultSubnetAttributes) EnableDns64() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("enable_dns64"))
}

// EnableLniAtDeviceIndex returns a reference to field enable_lni_at_device_index of aws_default_subnet.
func (ds defaultSubnetAttributes) EnableLniAtDeviceIndex() terra.NumberValue {
	return terra.ReferenceAsNumber(ds.ref.Append("enable_lni_at_device_index"))
}

// EnableResourceNameDnsARecordOnLaunch returns a reference to field enable_resource_name_dns_a_record_on_launch of aws_default_subnet.
func (ds defaultSubnetAttributes) EnableResourceNameDnsARecordOnLaunch() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("enable_resource_name_dns_a_record_on_launch"))
}

// EnableResourceNameDnsAaaaRecordOnLaunch returns a reference to field enable_resource_name_dns_aaaa_record_on_launch of aws_default_subnet.
func (ds defaultSubnetAttributes) EnableResourceNameDnsAaaaRecordOnLaunch() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("enable_resource_name_dns_aaaa_record_on_launch"))
}

// ExistingDefaultSubnet returns a reference to field existing_default_subnet of aws_default_subnet.
func (ds defaultSubnetAttributes) ExistingDefaultSubnet() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("existing_default_subnet"))
}

// ForceDestroy returns a reference to field force_destroy of aws_default_subnet.
func (ds defaultSubnetAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("force_destroy"))
}

// Id returns a reference to field id of aws_default_subnet.
func (ds defaultSubnetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("id"))
}

// Ipv6CidrBlock returns a reference to field ipv6_cidr_block of aws_default_subnet.
func (ds defaultSubnetAttributes) Ipv6CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("ipv6_cidr_block"))
}

// Ipv6CidrBlockAssociationId returns a reference to field ipv6_cidr_block_association_id of aws_default_subnet.
func (ds defaultSubnetAttributes) Ipv6CidrBlockAssociationId() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("ipv6_cidr_block_association_id"))
}

// Ipv6Native returns a reference to field ipv6_native of aws_default_subnet.
func (ds defaultSubnetAttributes) Ipv6Native() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("ipv6_native"))
}

// MapCustomerOwnedIpOnLaunch returns a reference to field map_customer_owned_ip_on_launch of aws_default_subnet.
func (ds defaultSubnetAttributes) MapCustomerOwnedIpOnLaunch() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("map_customer_owned_ip_on_launch"))
}

// MapPublicIpOnLaunch returns a reference to field map_public_ip_on_launch of aws_default_subnet.
func (ds defaultSubnetAttributes) MapPublicIpOnLaunch() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("map_public_ip_on_launch"))
}

// OutpostArn returns a reference to field outpost_arn of aws_default_subnet.
func (ds defaultSubnetAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("outpost_arn"))
}

// OwnerId returns a reference to field owner_id of aws_default_subnet.
func (ds defaultSubnetAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("owner_id"))
}

// PrivateDnsHostnameTypeOnLaunch returns a reference to field private_dns_hostname_type_on_launch of aws_default_subnet.
func (ds defaultSubnetAttributes) PrivateDnsHostnameTypeOnLaunch() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("private_dns_hostname_type_on_launch"))
}

// Tags returns a reference to field tags of aws_default_subnet.
func (ds defaultSubnetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ds.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_default_subnet.
func (ds defaultSubnetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ds.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_default_subnet.
func (ds defaultSubnetAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("vpc_id"))
}

func (ds defaultSubnetAttributes) Timeouts() defaultsubnet.TimeoutsAttributes {
	return terra.ReferenceAsSingle[defaultsubnet.TimeoutsAttributes](ds.ref.Append("timeouts"))
}

type defaultSubnetState struct {
	Arn                                     string                       `json:"arn"`
	AssignIpv6AddressOnCreation             bool                         `json:"assign_ipv6_address_on_creation"`
	AvailabilityZone                        string                       `json:"availability_zone"`
	AvailabilityZoneId                      string                       `json:"availability_zone_id"`
	CidrBlock                               string                       `json:"cidr_block"`
	CustomerOwnedIpv4Pool                   string                       `json:"customer_owned_ipv4_pool"`
	EnableDns64                             bool                         `json:"enable_dns64"`
	EnableLniAtDeviceIndex                  float64                      `json:"enable_lni_at_device_index"`
	EnableResourceNameDnsARecordOnLaunch    bool                         `json:"enable_resource_name_dns_a_record_on_launch"`
	EnableResourceNameDnsAaaaRecordOnLaunch bool                         `json:"enable_resource_name_dns_aaaa_record_on_launch"`
	ExistingDefaultSubnet                   bool                         `json:"existing_default_subnet"`
	ForceDestroy                            bool                         `json:"force_destroy"`
	Id                                      string                       `json:"id"`
	Ipv6CidrBlock                           string                       `json:"ipv6_cidr_block"`
	Ipv6CidrBlockAssociationId              string                       `json:"ipv6_cidr_block_association_id"`
	Ipv6Native                              bool                         `json:"ipv6_native"`
	MapCustomerOwnedIpOnLaunch              bool                         `json:"map_customer_owned_ip_on_launch"`
	MapPublicIpOnLaunch                     bool                         `json:"map_public_ip_on_launch"`
	OutpostArn                              string                       `json:"outpost_arn"`
	OwnerId                                 string                       `json:"owner_id"`
	PrivateDnsHostnameTypeOnLaunch          string                       `json:"private_dns_hostname_type_on_launch"`
	Tags                                    map[string]string            `json:"tags"`
	TagsAll                                 map[string]string            `json:"tags_all"`
	VpcId                                   string                       `json:"vpc_id"`
	Timeouts                                *defaultsubnet.TimeoutsState `json:"timeouts"`
}
