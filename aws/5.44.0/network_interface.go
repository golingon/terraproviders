// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	networkinterface "github.com/golingon/terraproviders/aws/5.44.0/networkinterface"
	"io"
)

// NewNetworkInterface creates a new instance of [NetworkInterface].
func NewNetworkInterface(name string, args NetworkInterfaceArgs) *NetworkInterface {
	return &NetworkInterface{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkInterface)(nil)

// NetworkInterface represents the Terraform resource aws_network_interface.
type NetworkInterface struct {
	Name      string
	Args      NetworkInterfaceArgs
	state     *networkInterfaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkInterface].
func (ni *NetworkInterface) Type() string {
	return "aws_network_interface"
}

// LocalName returns the local name for [NetworkInterface].
func (ni *NetworkInterface) LocalName() string {
	return ni.Name
}

// Configuration returns the configuration (args) for [NetworkInterface].
func (ni *NetworkInterface) Configuration() interface{} {
	return ni.Args
}

// DependOn is used for other resources to depend on [NetworkInterface].
func (ni *NetworkInterface) DependOn() terra.Reference {
	return terra.ReferenceResource(ni)
}

// Dependencies returns the list of resources [NetworkInterface] depends_on.
func (ni *NetworkInterface) Dependencies() terra.Dependencies {
	return ni.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkInterface].
func (ni *NetworkInterface) LifecycleManagement() *terra.Lifecycle {
	return ni.Lifecycle
}

// Attributes returns the attributes for [NetworkInterface].
func (ni *NetworkInterface) Attributes() networkInterfaceAttributes {
	return networkInterfaceAttributes{ref: terra.ReferenceResource(ni)}
}

// ImportState imports the given attribute values into [NetworkInterface]'s state.
func (ni *NetworkInterface) ImportState(av io.Reader) error {
	ni.state = &networkInterfaceState{}
	if err := json.NewDecoder(av).Decode(ni.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ni.Type(), ni.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkInterface] has state.
func (ni *NetworkInterface) State() (*networkInterfaceState, bool) {
	return ni.state, ni.state != nil
}

// StateMust returns the state for [NetworkInterface]. Panics if the state is nil.
func (ni *NetworkInterface) StateMust() *networkInterfaceState {
	if ni.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ni.Type(), ni.LocalName()))
	}
	return ni.state
}

// NetworkInterfaceArgs contains the configurations for aws_network_interface.
type NetworkInterfaceArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InterfaceType: string, optional
	InterfaceType terra.StringValue `hcl:"interface_type,attr"`
	// Ipv4PrefixCount: number, optional
	Ipv4PrefixCount terra.NumberValue `hcl:"ipv4_prefix_count,attr"`
	// Ipv4Prefixes: set of string, optional
	Ipv4Prefixes terra.SetValue[terra.StringValue] `hcl:"ipv4_prefixes,attr"`
	// Ipv6AddressCount: number, optional
	Ipv6AddressCount terra.NumberValue `hcl:"ipv6_address_count,attr"`
	// Ipv6AddressList: list of string, optional
	Ipv6AddressList terra.ListValue[terra.StringValue] `hcl:"ipv6_address_list,attr"`
	// Ipv6AddressListEnabled: bool, optional
	Ipv6AddressListEnabled terra.BoolValue `hcl:"ipv6_address_list_enabled,attr"`
	// Ipv6Addresses: set of string, optional
	Ipv6Addresses terra.SetValue[terra.StringValue] `hcl:"ipv6_addresses,attr"`
	// Ipv6PrefixCount: number, optional
	Ipv6PrefixCount terra.NumberValue `hcl:"ipv6_prefix_count,attr"`
	// Ipv6Prefixes: set of string, optional
	Ipv6Prefixes terra.SetValue[terra.StringValue] `hcl:"ipv6_prefixes,attr"`
	// PrivateIp: string, optional
	PrivateIp terra.StringValue `hcl:"private_ip,attr"`
	// PrivateIpList: list of string, optional
	PrivateIpList terra.ListValue[terra.StringValue] `hcl:"private_ip_list,attr"`
	// PrivateIpListEnabled: bool, optional
	PrivateIpListEnabled terra.BoolValue `hcl:"private_ip_list_enabled,attr"`
	// PrivateIps: set of string, optional
	PrivateIps terra.SetValue[terra.StringValue] `hcl:"private_ips,attr"`
	// PrivateIpsCount: number, optional
	PrivateIpsCount terra.NumberValue `hcl:"private_ips_count,attr"`
	// SecurityGroups: set of string, optional
	SecurityGroups terra.SetValue[terra.StringValue] `hcl:"security_groups,attr"`
	// SourceDestCheck: bool, optional
	SourceDestCheck terra.BoolValue `hcl:"source_dest_check,attr"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Attachment: min=0
	Attachment []networkinterface.Attachment `hcl:"attachment,block" validate:"min=0"`
}
type networkInterfaceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_network_interface.
func (ni networkInterfaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("arn"))
}

// Description returns a reference to field description of aws_network_interface.
func (ni networkInterfaceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("description"))
}

// Id returns a reference to field id of aws_network_interface.
func (ni networkInterfaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("id"))
}

// InterfaceType returns a reference to field interface_type of aws_network_interface.
func (ni networkInterfaceAttributes) InterfaceType() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("interface_type"))
}

// Ipv4PrefixCount returns a reference to field ipv4_prefix_count of aws_network_interface.
func (ni networkInterfaceAttributes) Ipv4PrefixCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ni.ref.Append("ipv4_prefix_count"))
}

// Ipv4Prefixes returns a reference to field ipv4_prefixes of aws_network_interface.
func (ni networkInterfaceAttributes) Ipv4Prefixes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ni.ref.Append("ipv4_prefixes"))
}

// Ipv6AddressCount returns a reference to field ipv6_address_count of aws_network_interface.
func (ni networkInterfaceAttributes) Ipv6AddressCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ni.ref.Append("ipv6_address_count"))
}

// Ipv6AddressList returns a reference to field ipv6_address_list of aws_network_interface.
func (ni networkInterfaceAttributes) Ipv6AddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ni.ref.Append("ipv6_address_list"))
}

// Ipv6AddressListEnabled returns a reference to field ipv6_address_list_enabled of aws_network_interface.
func (ni networkInterfaceAttributes) Ipv6AddressListEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ni.ref.Append("ipv6_address_list_enabled"))
}

// Ipv6Addresses returns a reference to field ipv6_addresses of aws_network_interface.
func (ni networkInterfaceAttributes) Ipv6Addresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ni.ref.Append("ipv6_addresses"))
}

// Ipv6PrefixCount returns a reference to field ipv6_prefix_count of aws_network_interface.
func (ni networkInterfaceAttributes) Ipv6PrefixCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ni.ref.Append("ipv6_prefix_count"))
}

// Ipv6Prefixes returns a reference to field ipv6_prefixes of aws_network_interface.
func (ni networkInterfaceAttributes) Ipv6Prefixes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ni.ref.Append("ipv6_prefixes"))
}

// MacAddress returns a reference to field mac_address of aws_network_interface.
func (ni networkInterfaceAttributes) MacAddress() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("mac_address"))
}

// OutpostArn returns a reference to field outpost_arn of aws_network_interface.
func (ni networkInterfaceAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("outpost_arn"))
}

// OwnerId returns a reference to field owner_id of aws_network_interface.
func (ni networkInterfaceAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("owner_id"))
}

// PrivateDnsName returns a reference to field private_dns_name of aws_network_interface.
func (ni networkInterfaceAttributes) PrivateDnsName() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("private_dns_name"))
}

// PrivateIp returns a reference to field private_ip of aws_network_interface.
func (ni networkInterfaceAttributes) PrivateIp() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("private_ip"))
}

// PrivateIpList returns a reference to field private_ip_list of aws_network_interface.
func (ni networkInterfaceAttributes) PrivateIpList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ni.ref.Append("private_ip_list"))
}

// PrivateIpListEnabled returns a reference to field private_ip_list_enabled of aws_network_interface.
func (ni networkInterfaceAttributes) PrivateIpListEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ni.ref.Append("private_ip_list_enabled"))
}

// PrivateIps returns a reference to field private_ips of aws_network_interface.
func (ni networkInterfaceAttributes) PrivateIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ni.ref.Append("private_ips"))
}

// PrivateIpsCount returns a reference to field private_ips_count of aws_network_interface.
func (ni networkInterfaceAttributes) PrivateIpsCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ni.ref.Append("private_ips_count"))
}

// SecurityGroups returns a reference to field security_groups of aws_network_interface.
func (ni networkInterfaceAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ni.ref.Append("security_groups"))
}

// SourceDestCheck returns a reference to field source_dest_check of aws_network_interface.
func (ni networkInterfaceAttributes) SourceDestCheck() terra.BoolValue {
	return terra.ReferenceAsBool(ni.ref.Append("source_dest_check"))
}

// SubnetId returns a reference to field subnet_id of aws_network_interface.
func (ni networkInterfaceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of aws_network_interface.
func (ni networkInterfaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ni.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_network_interface.
func (ni networkInterfaceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ni.ref.Append("tags_all"))
}

func (ni networkInterfaceAttributes) Attachment() terra.SetValue[networkinterface.AttachmentAttributes] {
	return terra.ReferenceAsSet[networkinterface.AttachmentAttributes](ni.ref.Append("attachment"))
}

type networkInterfaceState struct {
	Arn                    string                             `json:"arn"`
	Description            string                             `json:"description"`
	Id                     string                             `json:"id"`
	InterfaceType          string                             `json:"interface_type"`
	Ipv4PrefixCount        float64                            `json:"ipv4_prefix_count"`
	Ipv4Prefixes           []string                           `json:"ipv4_prefixes"`
	Ipv6AddressCount       float64                            `json:"ipv6_address_count"`
	Ipv6AddressList        []string                           `json:"ipv6_address_list"`
	Ipv6AddressListEnabled bool                               `json:"ipv6_address_list_enabled"`
	Ipv6Addresses          []string                           `json:"ipv6_addresses"`
	Ipv6PrefixCount        float64                            `json:"ipv6_prefix_count"`
	Ipv6Prefixes           []string                           `json:"ipv6_prefixes"`
	MacAddress             string                             `json:"mac_address"`
	OutpostArn             string                             `json:"outpost_arn"`
	OwnerId                string                             `json:"owner_id"`
	PrivateDnsName         string                             `json:"private_dns_name"`
	PrivateIp              string                             `json:"private_ip"`
	PrivateIpList          []string                           `json:"private_ip_list"`
	PrivateIpListEnabled   bool                               `json:"private_ip_list_enabled"`
	PrivateIps             []string                           `json:"private_ips"`
	PrivateIpsCount        float64                            `json:"private_ips_count"`
	SecurityGroups         []string                           `json:"security_groups"`
	SourceDestCheck        bool                               `json:"source_dest_check"`
	SubnetId               string                             `json:"subnet_id"`
	Tags                   map[string]string                  `json:"tags"`
	TagsAll                map[string]string                  `json:"tags_all"`
	Attachment             []networkinterface.AttachmentState `json:"attachment"`
}
