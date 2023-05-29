// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcDhcpOptions creates a new instance of [VpcDhcpOptions].
func NewVpcDhcpOptions(name string, args VpcDhcpOptionsArgs) *VpcDhcpOptions {
	return &VpcDhcpOptions{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcDhcpOptions)(nil)

// VpcDhcpOptions represents the Terraform resource aws_vpc_dhcp_options.
type VpcDhcpOptions struct {
	Name      string
	Args      VpcDhcpOptionsArgs
	state     *vpcDhcpOptionsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcDhcpOptions].
func (vdo *VpcDhcpOptions) Type() string {
	return "aws_vpc_dhcp_options"
}

// LocalName returns the local name for [VpcDhcpOptions].
func (vdo *VpcDhcpOptions) LocalName() string {
	return vdo.Name
}

// Configuration returns the configuration (args) for [VpcDhcpOptions].
func (vdo *VpcDhcpOptions) Configuration() interface{} {
	return vdo.Args
}

// DependOn is used for other resources to depend on [VpcDhcpOptions].
func (vdo *VpcDhcpOptions) DependOn() terra.Reference {
	return terra.ReferenceResource(vdo)
}

// Dependencies returns the list of resources [VpcDhcpOptions] depends_on.
func (vdo *VpcDhcpOptions) Dependencies() terra.Dependencies {
	return vdo.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcDhcpOptions].
func (vdo *VpcDhcpOptions) LifecycleManagement() *terra.Lifecycle {
	return vdo.Lifecycle
}

// Attributes returns the attributes for [VpcDhcpOptions].
func (vdo *VpcDhcpOptions) Attributes() vpcDhcpOptionsAttributes {
	return vpcDhcpOptionsAttributes{ref: terra.ReferenceResource(vdo)}
}

// ImportState imports the given attribute values into [VpcDhcpOptions]'s state.
func (vdo *VpcDhcpOptions) ImportState(av io.Reader) error {
	vdo.state = &vpcDhcpOptionsState{}
	if err := json.NewDecoder(av).Decode(vdo.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vdo.Type(), vdo.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcDhcpOptions] has state.
func (vdo *VpcDhcpOptions) State() (*vpcDhcpOptionsState, bool) {
	return vdo.state, vdo.state != nil
}

// StateMust returns the state for [VpcDhcpOptions]. Panics if the state is nil.
func (vdo *VpcDhcpOptions) StateMust() *vpcDhcpOptionsState {
	if vdo.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vdo.Type(), vdo.LocalName()))
	}
	return vdo.state
}

// VpcDhcpOptionsArgs contains the configurations for aws_vpc_dhcp_options.
type VpcDhcpOptionsArgs struct {
	// DomainName: string, optional
	DomainName terra.StringValue `hcl:"domain_name,attr"`
	// DomainNameServers: list of string, optional
	DomainNameServers terra.ListValue[terra.StringValue] `hcl:"domain_name_servers,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NetbiosNameServers: list of string, optional
	NetbiosNameServers terra.ListValue[terra.StringValue] `hcl:"netbios_name_servers,attr"`
	// NetbiosNodeType: string, optional
	NetbiosNodeType terra.StringValue `hcl:"netbios_node_type,attr"`
	// NtpServers: list of string, optional
	NtpServers terra.ListValue[terra.StringValue] `hcl:"ntp_servers,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type vpcDhcpOptionsAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpc_dhcp_options.
func (vdo vpcDhcpOptionsAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(vdo.ref.Append("arn"))
}

// DomainName returns a reference to field domain_name of aws_vpc_dhcp_options.
func (vdo vpcDhcpOptionsAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(vdo.ref.Append("domain_name"))
}

// DomainNameServers returns a reference to field domain_name_servers of aws_vpc_dhcp_options.
func (vdo vpcDhcpOptionsAttributes) DomainNameServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vdo.ref.Append("domain_name_servers"))
}

// Id returns a reference to field id of aws_vpc_dhcp_options.
func (vdo vpcDhcpOptionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vdo.ref.Append("id"))
}

// NetbiosNameServers returns a reference to field netbios_name_servers of aws_vpc_dhcp_options.
func (vdo vpcDhcpOptionsAttributes) NetbiosNameServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vdo.ref.Append("netbios_name_servers"))
}

// NetbiosNodeType returns a reference to field netbios_node_type of aws_vpc_dhcp_options.
func (vdo vpcDhcpOptionsAttributes) NetbiosNodeType() terra.StringValue {
	return terra.ReferenceAsString(vdo.ref.Append("netbios_node_type"))
}

// NtpServers returns a reference to field ntp_servers of aws_vpc_dhcp_options.
func (vdo vpcDhcpOptionsAttributes) NtpServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vdo.ref.Append("ntp_servers"))
}

// OwnerId returns a reference to field owner_id of aws_vpc_dhcp_options.
func (vdo vpcDhcpOptionsAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(vdo.ref.Append("owner_id"))
}

// Tags returns a reference to field tags of aws_vpc_dhcp_options.
func (vdo vpcDhcpOptionsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vdo.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpc_dhcp_options.
func (vdo vpcDhcpOptionsAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vdo.ref.Append("tags_all"))
}

type vpcDhcpOptionsState struct {
	Arn                string            `json:"arn"`
	DomainName         string            `json:"domain_name"`
	DomainNameServers  []string          `json:"domain_name_servers"`
	Id                 string            `json:"id"`
	NetbiosNameServers []string          `json:"netbios_name_servers"`
	NetbiosNodeType    string            `json:"netbios_node_type"`
	NtpServers         []string          `json:"ntp_servers"`
	OwnerId            string            `json:"owner_id"`
	Tags               map[string]string `json:"tags"`
	TagsAll            map[string]string `json:"tags_all"`
}