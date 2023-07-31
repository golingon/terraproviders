// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpcendpoint "github.com/golingon/terraproviders/aws/5.10.0/vpcendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcEndpoint creates a new instance of [VpcEndpoint].
func NewVpcEndpoint(name string, args VpcEndpointArgs) *VpcEndpoint {
	return &VpcEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcEndpoint)(nil)

// VpcEndpoint represents the Terraform resource aws_vpc_endpoint.
type VpcEndpoint struct {
	Name      string
	Args      VpcEndpointArgs
	state     *vpcEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcEndpoint].
func (ve *VpcEndpoint) Type() string {
	return "aws_vpc_endpoint"
}

// LocalName returns the local name for [VpcEndpoint].
func (ve *VpcEndpoint) LocalName() string {
	return ve.Name
}

// Configuration returns the configuration (args) for [VpcEndpoint].
func (ve *VpcEndpoint) Configuration() interface{} {
	return ve.Args
}

// DependOn is used for other resources to depend on [VpcEndpoint].
func (ve *VpcEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(ve)
}

// Dependencies returns the list of resources [VpcEndpoint] depends_on.
func (ve *VpcEndpoint) Dependencies() terra.Dependencies {
	return ve.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcEndpoint].
func (ve *VpcEndpoint) LifecycleManagement() *terra.Lifecycle {
	return ve.Lifecycle
}

// Attributes returns the attributes for [VpcEndpoint].
func (ve *VpcEndpoint) Attributes() vpcEndpointAttributes {
	return vpcEndpointAttributes{ref: terra.ReferenceResource(ve)}
}

// ImportState imports the given attribute values into [VpcEndpoint]'s state.
func (ve *VpcEndpoint) ImportState(av io.Reader) error {
	ve.state = &vpcEndpointState{}
	if err := json.NewDecoder(av).Decode(ve.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ve.Type(), ve.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcEndpoint] has state.
func (ve *VpcEndpoint) State() (*vpcEndpointState, bool) {
	return ve.state, ve.state != nil
}

// StateMust returns the state for [VpcEndpoint]. Panics if the state is nil.
func (ve *VpcEndpoint) StateMust() *vpcEndpointState {
	if ve.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ve.Type(), ve.LocalName()))
	}
	return ve.state
}

// VpcEndpointArgs contains the configurations for aws_vpc_endpoint.
type VpcEndpointArgs struct {
	// AutoAccept: bool, optional
	AutoAccept terra.BoolValue `hcl:"auto_accept,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpAddressType: string, optional
	IpAddressType terra.StringValue `hcl:"ip_address_type,attr"`
	// Policy: string, optional
	Policy terra.StringValue `hcl:"policy,attr"`
	// PrivateDnsEnabled: bool, optional
	PrivateDnsEnabled terra.BoolValue `hcl:"private_dns_enabled,attr"`
	// RouteTableIds: set of string, optional
	RouteTableIds terra.SetValue[terra.StringValue] `hcl:"route_table_ids,attr"`
	// SecurityGroupIds: set of string, optional
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// SubnetIds: set of string, optional
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcEndpointType: string, optional
	VpcEndpointType terra.StringValue `hcl:"vpc_endpoint_type,attr"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
	// DnsEntry: min=0
	DnsEntry []vpcendpoint.DnsEntry `hcl:"dns_entry,block" validate:"min=0"`
	// DnsOptions: optional
	DnsOptions *vpcendpoint.DnsOptions `hcl:"dns_options,block"`
	// Timeouts: optional
	Timeouts *vpcendpoint.Timeouts `hcl:"timeouts,block"`
}
type vpcEndpointAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("arn"))
}

// AutoAccept returns a reference to field auto_accept of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) AutoAccept() terra.BoolValue {
	return terra.ReferenceAsBool(ve.ref.Append("auto_accept"))
}

// CidrBlocks returns a reference to field cidr_blocks of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) CidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ve.ref.Append("cidr_blocks"))
}

// Id returns a reference to field id of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("id"))
}

// IpAddressType returns a reference to field ip_address_type of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) IpAddressType() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("ip_address_type"))
}

// NetworkInterfaceIds returns a reference to field network_interface_ids of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) NetworkInterfaceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ve.ref.Append("network_interface_ids"))
}

// OwnerId returns a reference to field owner_id of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("owner_id"))
}

// Policy returns a reference to field policy of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("policy"))
}

// PrefixListId returns a reference to field prefix_list_id of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) PrefixListId() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("prefix_list_id"))
}

// PrivateDnsEnabled returns a reference to field private_dns_enabled of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) PrivateDnsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ve.ref.Append("private_dns_enabled"))
}

// RequesterManaged returns a reference to field requester_managed of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) RequesterManaged() terra.BoolValue {
	return terra.ReferenceAsBool(ve.ref.Append("requester_managed"))
}

// RouteTableIds returns a reference to field route_table_ids of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) RouteTableIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ve.ref.Append("route_table_ids"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ve.ref.Append("security_group_ids"))
}

// ServiceName returns a reference to field service_name of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("service_name"))
}

// State returns a reference to field state of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("state"))
}

// SubnetIds returns a reference to field subnet_ids of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ve.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ve.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ve.ref.Append("tags_all"))
}

// VpcEndpointType returns a reference to field vpc_endpoint_type of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) VpcEndpointType() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("vpc_endpoint_type"))
}

// VpcId returns a reference to field vpc_id of aws_vpc_endpoint.
func (ve vpcEndpointAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("vpc_id"))
}

func (ve vpcEndpointAttributes) DnsEntry() terra.ListValue[vpcendpoint.DnsEntryAttributes] {
	return terra.ReferenceAsList[vpcendpoint.DnsEntryAttributes](ve.ref.Append("dns_entry"))
}

func (ve vpcEndpointAttributes) DnsOptions() terra.ListValue[vpcendpoint.DnsOptionsAttributes] {
	return terra.ReferenceAsList[vpcendpoint.DnsOptionsAttributes](ve.ref.Append("dns_options"))
}

func (ve vpcEndpointAttributes) Timeouts() vpcendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpcendpoint.TimeoutsAttributes](ve.ref.Append("timeouts"))
}

type vpcEndpointState struct {
	Arn                 string                        `json:"arn"`
	AutoAccept          bool                          `json:"auto_accept"`
	CidrBlocks          []string                      `json:"cidr_blocks"`
	Id                  string                        `json:"id"`
	IpAddressType       string                        `json:"ip_address_type"`
	NetworkInterfaceIds []string                      `json:"network_interface_ids"`
	OwnerId             string                        `json:"owner_id"`
	Policy              string                        `json:"policy"`
	PrefixListId        string                        `json:"prefix_list_id"`
	PrivateDnsEnabled   bool                          `json:"private_dns_enabled"`
	RequesterManaged    bool                          `json:"requester_managed"`
	RouteTableIds       []string                      `json:"route_table_ids"`
	SecurityGroupIds    []string                      `json:"security_group_ids"`
	ServiceName         string                        `json:"service_name"`
	State               string                        `json:"state"`
	SubnetIds           []string                      `json:"subnet_ids"`
	Tags                map[string]string             `json:"tags"`
	TagsAll             map[string]string             `json:"tags_all"`
	VpcEndpointType     string                        `json:"vpc_endpoint_type"`
	VpcId               string                        `json:"vpc_id"`
	DnsEntry            []vpcendpoint.DnsEntryState   `json:"dns_entry"`
	DnsOptions          []vpcendpoint.DnsOptionsState `json:"dns_options"`
	Timeouts            *vpcendpoint.TimeoutsState    `json:"timeouts"`
}
