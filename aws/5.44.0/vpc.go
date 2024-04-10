// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewVpc creates a new instance of [Vpc].
func NewVpc(name string, args VpcArgs) *Vpc {
	return &Vpc{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Vpc)(nil)

// Vpc represents the Terraform resource aws_vpc.
type Vpc struct {
	Name      string
	Args      VpcArgs
	state     *vpcState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Vpc].
func (v *Vpc) Type() string {
	return "aws_vpc"
}

// LocalName returns the local name for [Vpc].
func (v *Vpc) LocalName() string {
	return v.Name
}

// Configuration returns the configuration (args) for [Vpc].
func (v *Vpc) Configuration() interface{} {
	return v.Args
}

// DependOn is used for other resources to depend on [Vpc].
func (v *Vpc) DependOn() terra.Reference {
	return terra.ReferenceResource(v)
}

// Dependencies returns the list of resources [Vpc] depends_on.
func (v *Vpc) Dependencies() terra.Dependencies {
	return v.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Vpc].
func (v *Vpc) LifecycleManagement() *terra.Lifecycle {
	return v.Lifecycle
}

// Attributes returns the attributes for [Vpc].
func (v *Vpc) Attributes() vpcAttributes {
	return vpcAttributes{ref: terra.ReferenceResource(v)}
}

// ImportState imports the given attribute values into [Vpc]'s state.
func (v *Vpc) ImportState(av io.Reader) error {
	v.state = &vpcState{}
	if err := json.NewDecoder(av).Decode(v.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", v.Type(), v.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Vpc] has state.
func (v *Vpc) State() (*vpcState, bool) {
	return v.state, v.state != nil
}

// StateMust returns the state for [Vpc]. Panics if the state is nil.
func (v *Vpc) StateMust() *vpcState {
	if v.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", v.Type(), v.LocalName()))
	}
	return v.state
}

// VpcArgs contains the configurations for aws_vpc.
type VpcArgs struct {
	// AssignGeneratedIpv6CidrBlock: bool, optional
	AssignGeneratedIpv6CidrBlock terra.BoolValue `hcl:"assign_generated_ipv6_cidr_block,attr"`
	// CidrBlock: string, optional
	CidrBlock terra.StringValue `hcl:"cidr_block,attr"`
	// EnableDnsHostnames: bool, optional
	EnableDnsHostnames terra.BoolValue `hcl:"enable_dns_hostnames,attr"`
	// EnableDnsSupport: bool, optional
	EnableDnsSupport terra.BoolValue `hcl:"enable_dns_support,attr"`
	// EnableNetworkAddressUsageMetrics: bool, optional
	EnableNetworkAddressUsageMetrics terra.BoolValue `hcl:"enable_network_address_usage_metrics,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceTenancy: string, optional
	InstanceTenancy terra.StringValue `hcl:"instance_tenancy,attr"`
	// Ipv4IpamPoolId: string, optional
	Ipv4IpamPoolId terra.StringValue `hcl:"ipv4_ipam_pool_id,attr"`
	// Ipv4NetmaskLength: number, optional
	Ipv4NetmaskLength terra.NumberValue `hcl:"ipv4_netmask_length,attr"`
	// Ipv6CidrBlock: string, optional
	Ipv6CidrBlock terra.StringValue `hcl:"ipv6_cidr_block,attr"`
	// Ipv6CidrBlockNetworkBorderGroup: string, optional
	Ipv6CidrBlockNetworkBorderGroup terra.StringValue `hcl:"ipv6_cidr_block_network_border_group,attr"`
	// Ipv6IpamPoolId: string, optional
	Ipv6IpamPoolId terra.StringValue `hcl:"ipv6_ipam_pool_id,attr"`
	// Ipv6NetmaskLength: number, optional
	Ipv6NetmaskLength terra.NumberValue `hcl:"ipv6_netmask_length,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type vpcAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpc.
func (v vpcAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("arn"))
}

// AssignGeneratedIpv6CidrBlock returns a reference to field assign_generated_ipv6_cidr_block of aws_vpc.
func (v vpcAttributes) AssignGeneratedIpv6CidrBlock() terra.BoolValue {
	return terra.ReferenceAsBool(v.ref.Append("assign_generated_ipv6_cidr_block"))
}

// CidrBlock returns a reference to field cidr_block of aws_vpc.
func (v vpcAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("cidr_block"))
}

// DefaultNetworkAclId returns a reference to field default_network_acl_id of aws_vpc.
func (v vpcAttributes) DefaultNetworkAclId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("default_network_acl_id"))
}

// DefaultRouteTableId returns a reference to field default_route_table_id of aws_vpc.
func (v vpcAttributes) DefaultRouteTableId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("default_route_table_id"))
}

// DefaultSecurityGroupId returns a reference to field default_security_group_id of aws_vpc.
func (v vpcAttributes) DefaultSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("default_security_group_id"))
}

// DhcpOptionsId returns a reference to field dhcp_options_id of aws_vpc.
func (v vpcAttributes) DhcpOptionsId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("dhcp_options_id"))
}

// EnableDnsHostnames returns a reference to field enable_dns_hostnames of aws_vpc.
func (v vpcAttributes) EnableDnsHostnames() terra.BoolValue {
	return terra.ReferenceAsBool(v.ref.Append("enable_dns_hostnames"))
}

// EnableDnsSupport returns a reference to field enable_dns_support of aws_vpc.
func (v vpcAttributes) EnableDnsSupport() terra.BoolValue {
	return terra.ReferenceAsBool(v.ref.Append("enable_dns_support"))
}

// EnableNetworkAddressUsageMetrics returns a reference to field enable_network_address_usage_metrics of aws_vpc.
func (v vpcAttributes) EnableNetworkAddressUsageMetrics() terra.BoolValue {
	return terra.ReferenceAsBool(v.ref.Append("enable_network_address_usage_metrics"))
}

// Id returns a reference to field id of aws_vpc.
func (v vpcAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("id"))
}

// InstanceTenancy returns a reference to field instance_tenancy of aws_vpc.
func (v vpcAttributes) InstanceTenancy() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("instance_tenancy"))
}

// Ipv4IpamPoolId returns a reference to field ipv4_ipam_pool_id of aws_vpc.
func (v vpcAttributes) Ipv4IpamPoolId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("ipv4_ipam_pool_id"))
}

// Ipv4NetmaskLength returns a reference to field ipv4_netmask_length of aws_vpc.
func (v vpcAttributes) Ipv4NetmaskLength() terra.NumberValue {
	return terra.ReferenceAsNumber(v.ref.Append("ipv4_netmask_length"))
}

// Ipv6AssociationId returns a reference to field ipv6_association_id of aws_vpc.
func (v vpcAttributes) Ipv6AssociationId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("ipv6_association_id"))
}

// Ipv6CidrBlock returns a reference to field ipv6_cidr_block of aws_vpc.
func (v vpcAttributes) Ipv6CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("ipv6_cidr_block"))
}

// Ipv6CidrBlockNetworkBorderGroup returns a reference to field ipv6_cidr_block_network_border_group of aws_vpc.
func (v vpcAttributes) Ipv6CidrBlockNetworkBorderGroup() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("ipv6_cidr_block_network_border_group"))
}

// Ipv6IpamPoolId returns a reference to field ipv6_ipam_pool_id of aws_vpc.
func (v vpcAttributes) Ipv6IpamPoolId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("ipv6_ipam_pool_id"))
}

// Ipv6NetmaskLength returns a reference to field ipv6_netmask_length of aws_vpc.
func (v vpcAttributes) Ipv6NetmaskLength() terra.NumberValue {
	return terra.ReferenceAsNumber(v.ref.Append("ipv6_netmask_length"))
}

// MainRouteTableId returns a reference to field main_route_table_id of aws_vpc.
func (v vpcAttributes) MainRouteTableId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("main_route_table_id"))
}

// OwnerId returns a reference to field owner_id of aws_vpc.
func (v vpcAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("owner_id"))
}

// Tags returns a reference to field tags of aws_vpc.
func (v vpcAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](v.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpc.
func (v vpcAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](v.ref.Append("tags_all"))
}

type vpcState struct {
	Arn                              string            `json:"arn"`
	AssignGeneratedIpv6CidrBlock     bool              `json:"assign_generated_ipv6_cidr_block"`
	CidrBlock                        string            `json:"cidr_block"`
	DefaultNetworkAclId              string            `json:"default_network_acl_id"`
	DefaultRouteTableId              string            `json:"default_route_table_id"`
	DefaultSecurityGroupId           string            `json:"default_security_group_id"`
	DhcpOptionsId                    string            `json:"dhcp_options_id"`
	EnableDnsHostnames               bool              `json:"enable_dns_hostnames"`
	EnableDnsSupport                 bool              `json:"enable_dns_support"`
	EnableNetworkAddressUsageMetrics bool              `json:"enable_network_address_usage_metrics"`
	Id                               string            `json:"id"`
	InstanceTenancy                  string            `json:"instance_tenancy"`
	Ipv4IpamPoolId                   string            `json:"ipv4_ipam_pool_id"`
	Ipv4NetmaskLength                float64           `json:"ipv4_netmask_length"`
	Ipv6AssociationId                string            `json:"ipv6_association_id"`
	Ipv6CidrBlock                    string            `json:"ipv6_cidr_block"`
	Ipv6CidrBlockNetworkBorderGroup  string            `json:"ipv6_cidr_block_network_border_group"`
	Ipv6IpamPoolId                   string            `json:"ipv6_ipam_pool_id"`
	Ipv6NetmaskLength                float64           `json:"ipv6_netmask_length"`
	MainRouteTableId                 string            `json:"main_route_table_id"`
	OwnerId                          string            `json:"owner_id"`
	Tags                             map[string]string `json:"tags"`
	TagsAll                          map[string]string `json:"tags_all"`
}
