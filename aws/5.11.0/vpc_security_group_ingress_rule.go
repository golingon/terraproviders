// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcSecurityGroupIngressRule creates a new instance of [VpcSecurityGroupIngressRule].
func NewVpcSecurityGroupIngressRule(name string, args VpcSecurityGroupIngressRuleArgs) *VpcSecurityGroupIngressRule {
	return &VpcSecurityGroupIngressRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcSecurityGroupIngressRule)(nil)

// VpcSecurityGroupIngressRule represents the Terraform resource aws_vpc_security_group_ingress_rule.
type VpcSecurityGroupIngressRule struct {
	Name      string
	Args      VpcSecurityGroupIngressRuleArgs
	state     *vpcSecurityGroupIngressRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcSecurityGroupIngressRule].
func (vsgir *VpcSecurityGroupIngressRule) Type() string {
	return "aws_vpc_security_group_ingress_rule"
}

// LocalName returns the local name for [VpcSecurityGroupIngressRule].
func (vsgir *VpcSecurityGroupIngressRule) LocalName() string {
	return vsgir.Name
}

// Configuration returns the configuration (args) for [VpcSecurityGroupIngressRule].
func (vsgir *VpcSecurityGroupIngressRule) Configuration() interface{} {
	return vsgir.Args
}

// DependOn is used for other resources to depend on [VpcSecurityGroupIngressRule].
func (vsgir *VpcSecurityGroupIngressRule) DependOn() terra.Reference {
	return terra.ReferenceResource(vsgir)
}

// Dependencies returns the list of resources [VpcSecurityGroupIngressRule] depends_on.
func (vsgir *VpcSecurityGroupIngressRule) Dependencies() terra.Dependencies {
	return vsgir.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcSecurityGroupIngressRule].
func (vsgir *VpcSecurityGroupIngressRule) LifecycleManagement() *terra.Lifecycle {
	return vsgir.Lifecycle
}

// Attributes returns the attributes for [VpcSecurityGroupIngressRule].
func (vsgir *VpcSecurityGroupIngressRule) Attributes() vpcSecurityGroupIngressRuleAttributes {
	return vpcSecurityGroupIngressRuleAttributes{ref: terra.ReferenceResource(vsgir)}
}

// ImportState imports the given attribute values into [VpcSecurityGroupIngressRule]'s state.
func (vsgir *VpcSecurityGroupIngressRule) ImportState(av io.Reader) error {
	vsgir.state = &vpcSecurityGroupIngressRuleState{}
	if err := json.NewDecoder(av).Decode(vsgir.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vsgir.Type(), vsgir.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcSecurityGroupIngressRule] has state.
func (vsgir *VpcSecurityGroupIngressRule) State() (*vpcSecurityGroupIngressRuleState, bool) {
	return vsgir.state, vsgir.state != nil
}

// StateMust returns the state for [VpcSecurityGroupIngressRule]. Panics if the state is nil.
func (vsgir *VpcSecurityGroupIngressRule) StateMust() *vpcSecurityGroupIngressRuleState {
	if vsgir.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vsgir.Type(), vsgir.LocalName()))
	}
	return vsgir.state
}

// VpcSecurityGroupIngressRuleArgs contains the configurations for aws_vpc_security_group_ingress_rule.
type VpcSecurityGroupIngressRuleArgs struct {
	// CidrIpv4: string, optional
	CidrIpv4 terra.StringValue `hcl:"cidr_ipv4,attr"`
	// CidrIpv6: string, optional
	CidrIpv6 terra.StringValue `hcl:"cidr_ipv6,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FromPort: number, optional
	FromPort terra.NumberValue `hcl:"from_port,attr"`
	// IpProtocol: string, required
	IpProtocol terra.StringValue `hcl:"ip_protocol,attr" validate:"required"`
	// PrefixListId: string, optional
	PrefixListId terra.StringValue `hcl:"prefix_list_id,attr"`
	// ReferencedSecurityGroupId: string, optional
	ReferencedSecurityGroupId terra.StringValue `hcl:"referenced_security_group_id,attr"`
	// SecurityGroupId: string, required
	SecurityGroupId terra.StringValue `hcl:"security_group_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ToPort: number, optional
	ToPort terra.NumberValue `hcl:"to_port,attr"`
}
type vpcSecurityGroupIngressRuleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpc_security_group_ingress_rule.
func (vsgir vpcSecurityGroupIngressRuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(vsgir.ref.Append("arn"))
}

// CidrIpv4 returns a reference to field cidr_ipv4 of aws_vpc_security_group_ingress_rule.
func (vsgir vpcSecurityGroupIngressRuleAttributes) CidrIpv4() terra.StringValue {
	return terra.ReferenceAsString(vsgir.ref.Append("cidr_ipv4"))
}

// CidrIpv6 returns a reference to field cidr_ipv6 of aws_vpc_security_group_ingress_rule.
func (vsgir vpcSecurityGroupIngressRuleAttributes) CidrIpv6() terra.StringValue {
	return terra.ReferenceAsString(vsgir.ref.Append("cidr_ipv6"))
}

// Description returns a reference to field description of aws_vpc_security_group_ingress_rule.
func (vsgir vpcSecurityGroupIngressRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vsgir.ref.Append("description"))
}

// FromPort returns a reference to field from_port of aws_vpc_security_group_ingress_rule.
func (vsgir vpcSecurityGroupIngressRuleAttributes) FromPort() terra.NumberValue {
	return terra.ReferenceAsNumber(vsgir.ref.Append("from_port"))
}

// Id returns a reference to field id of aws_vpc_security_group_ingress_rule.
func (vsgir vpcSecurityGroupIngressRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vsgir.ref.Append("id"))
}

// IpProtocol returns a reference to field ip_protocol of aws_vpc_security_group_ingress_rule.
func (vsgir vpcSecurityGroupIngressRuleAttributes) IpProtocol() terra.StringValue {
	return terra.ReferenceAsString(vsgir.ref.Append("ip_protocol"))
}

// PrefixListId returns a reference to field prefix_list_id of aws_vpc_security_group_ingress_rule.
func (vsgir vpcSecurityGroupIngressRuleAttributes) PrefixListId() terra.StringValue {
	return terra.ReferenceAsString(vsgir.ref.Append("prefix_list_id"))
}

// ReferencedSecurityGroupId returns a reference to field referenced_security_group_id of aws_vpc_security_group_ingress_rule.
func (vsgir vpcSecurityGroupIngressRuleAttributes) ReferencedSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(vsgir.ref.Append("referenced_security_group_id"))
}

// SecurityGroupId returns a reference to field security_group_id of aws_vpc_security_group_ingress_rule.
func (vsgir vpcSecurityGroupIngressRuleAttributes) SecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(vsgir.ref.Append("security_group_id"))
}

// SecurityGroupRuleId returns a reference to field security_group_rule_id of aws_vpc_security_group_ingress_rule.
func (vsgir vpcSecurityGroupIngressRuleAttributes) SecurityGroupRuleId() terra.StringValue {
	return terra.ReferenceAsString(vsgir.ref.Append("security_group_rule_id"))
}

// Tags returns a reference to field tags of aws_vpc_security_group_ingress_rule.
func (vsgir vpcSecurityGroupIngressRuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vsgir.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpc_security_group_ingress_rule.
func (vsgir vpcSecurityGroupIngressRuleAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vsgir.ref.Append("tags_all"))
}

// ToPort returns a reference to field to_port of aws_vpc_security_group_ingress_rule.
func (vsgir vpcSecurityGroupIngressRuleAttributes) ToPort() terra.NumberValue {
	return terra.ReferenceAsNumber(vsgir.ref.Append("to_port"))
}

type vpcSecurityGroupIngressRuleState struct {
	Arn                       string            `json:"arn"`
	CidrIpv4                  string            `json:"cidr_ipv4"`
	CidrIpv6                  string            `json:"cidr_ipv6"`
	Description               string            `json:"description"`
	FromPort                  float64           `json:"from_port"`
	Id                        string            `json:"id"`
	IpProtocol                string            `json:"ip_protocol"`
	PrefixListId              string            `json:"prefix_list_id"`
	ReferencedSecurityGroupId string            `json:"referenced_security_group_id"`
	SecurityGroupId           string            `json:"security_group_id"`
	SecurityGroupRuleId       string            `json:"security_group_rule_id"`
	Tags                      map[string]string `json:"tags"`
	TagsAll                   map[string]string `json:"tags_all"`
	ToPort                    float64           `json:"to_port"`
}
