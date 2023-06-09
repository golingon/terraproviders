// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkAclRule creates a new instance of [NetworkAclRule].
func NewNetworkAclRule(name string, args NetworkAclRuleArgs) *NetworkAclRule {
	return &NetworkAclRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkAclRule)(nil)

// NetworkAclRule represents the Terraform resource aws_network_acl_rule.
type NetworkAclRule struct {
	Name      string
	Args      NetworkAclRuleArgs
	state     *networkAclRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkAclRule].
func (nar *NetworkAclRule) Type() string {
	return "aws_network_acl_rule"
}

// LocalName returns the local name for [NetworkAclRule].
func (nar *NetworkAclRule) LocalName() string {
	return nar.Name
}

// Configuration returns the configuration (args) for [NetworkAclRule].
func (nar *NetworkAclRule) Configuration() interface{} {
	return nar.Args
}

// DependOn is used for other resources to depend on [NetworkAclRule].
func (nar *NetworkAclRule) DependOn() terra.Reference {
	return terra.ReferenceResource(nar)
}

// Dependencies returns the list of resources [NetworkAclRule] depends_on.
func (nar *NetworkAclRule) Dependencies() terra.Dependencies {
	return nar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkAclRule].
func (nar *NetworkAclRule) LifecycleManagement() *terra.Lifecycle {
	return nar.Lifecycle
}

// Attributes returns the attributes for [NetworkAclRule].
func (nar *NetworkAclRule) Attributes() networkAclRuleAttributes {
	return networkAclRuleAttributes{ref: terra.ReferenceResource(nar)}
}

// ImportState imports the given attribute values into [NetworkAclRule]'s state.
func (nar *NetworkAclRule) ImportState(av io.Reader) error {
	nar.state = &networkAclRuleState{}
	if err := json.NewDecoder(av).Decode(nar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nar.Type(), nar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkAclRule] has state.
func (nar *NetworkAclRule) State() (*networkAclRuleState, bool) {
	return nar.state, nar.state != nil
}

// StateMust returns the state for [NetworkAclRule]. Panics if the state is nil.
func (nar *NetworkAclRule) StateMust() *networkAclRuleState {
	if nar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nar.Type(), nar.LocalName()))
	}
	return nar.state
}

// NetworkAclRuleArgs contains the configurations for aws_network_acl_rule.
type NetworkAclRuleArgs struct {
	// CidrBlock: string, optional
	CidrBlock terra.StringValue `hcl:"cidr_block,attr"`
	// Egress: bool, optional
	Egress terra.BoolValue `hcl:"egress,attr"`
	// FromPort: number, optional
	FromPort terra.NumberValue `hcl:"from_port,attr"`
	// IcmpCode: number, optional
	IcmpCode terra.NumberValue `hcl:"icmp_code,attr"`
	// IcmpType: number, optional
	IcmpType terra.NumberValue `hcl:"icmp_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Ipv6CidrBlock: string, optional
	Ipv6CidrBlock terra.StringValue `hcl:"ipv6_cidr_block,attr"`
	// NetworkAclId: string, required
	NetworkAclId terra.StringValue `hcl:"network_acl_id,attr" validate:"required"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
	// RuleAction: string, required
	RuleAction terra.StringValue `hcl:"rule_action,attr" validate:"required"`
	// RuleNumber: number, required
	RuleNumber terra.NumberValue `hcl:"rule_number,attr" validate:"required"`
	// ToPort: number, optional
	ToPort terra.NumberValue `hcl:"to_port,attr"`
}
type networkAclRuleAttributes struct {
	ref terra.Reference
}

// CidrBlock returns a reference to field cidr_block of aws_network_acl_rule.
func (nar networkAclRuleAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(nar.ref.Append("cidr_block"))
}

// Egress returns a reference to field egress of aws_network_acl_rule.
func (nar networkAclRuleAttributes) Egress() terra.BoolValue {
	return terra.ReferenceAsBool(nar.ref.Append("egress"))
}

// FromPort returns a reference to field from_port of aws_network_acl_rule.
func (nar networkAclRuleAttributes) FromPort() terra.NumberValue {
	return terra.ReferenceAsNumber(nar.ref.Append("from_port"))
}

// IcmpCode returns a reference to field icmp_code of aws_network_acl_rule.
func (nar networkAclRuleAttributes) IcmpCode() terra.NumberValue {
	return terra.ReferenceAsNumber(nar.ref.Append("icmp_code"))
}

// IcmpType returns a reference to field icmp_type of aws_network_acl_rule.
func (nar networkAclRuleAttributes) IcmpType() terra.NumberValue {
	return terra.ReferenceAsNumber(nar.ref.Append("icmp_type"))
}

// Id returns a reference to field id of aws_network_acl_rule.
func (nar networkAclRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nar.ref.Append("id"))
}

// Ipv6CidrBlock returns a reference to field ipv6_cidr_block of aws_network_acl_rule.
func (nar networkAclRuleAttributes) Ipv6CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(nar.ref.Append("ipv6_cidr_block"))
}

// NetworkAclId returns a reference to field network_acl_id of aws_network_acl_rule.
func (nar networkAclRuleAttributes) NetworkAclId() terra.StringValue {
	return terra.ReferenceAsString(nar.ref.Append("network_acl_id"))
}

// Protocol returns a reference to field protocol of aws_network_acl_rule.
func (nar networkAclRuleAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(nar.ref.Append("protocol"))
}

// RuleAction returns a reference to field rule_action of aws_network_acl_rule.
func (nar networkAclRuleAttributes) RuleAction() terra.StringValue {
	return terra.ReferenceAsString(nar.ref.Append("rule_action"))
}

// RuleNumber returns a reference to field rule_number of aws_network_acl_rule.
func (nar networkAclRuleAttributes) RuleNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(nar.ref.Append("rule_number"))
}

// ToPort returns a reference to field to_port of aws_network_acl_rule.
func (nar networkAclRuleAttributes) ToPort() terra.NumberValue {
	return terra.ReferenceAsNumber(nar.ref.Append("to_port"))
}

type networkAclRuleState struct {
	CidrBlock     string  `json:"cidr_block"`
	Egress        bool    `json:"egress"`
	FromPort      float64 `json:"from_port"`
	IcmpCode      float64 `json:"icmp_code"`
	IcmpType      float64 `json:"icmp_type"`
	Id            string  `json:"id"`
	Ipv6CidrBlock string  `json:"ipv6_cidr_block"`
	NetworkAclId  string  `json:"network_acl_id"`
	Protocol      string  `json:"protocol"`
	RuleAction    string  `json:"rule_action"`
	RuleNumber    float64 `json:"rule_number"`
	ToPort        float64 `json:"to_port"`
}
