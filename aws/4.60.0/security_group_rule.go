// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	securitygrouprule "github.com/golingon/terraproviders/aws/4.60.0/securitygrouprule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSecurityGroupRule creates a new instance of [SecurityGroupRule].
func NewSecurityGroupRule(name string, args SecurityGroupRuleArgs) *SecurityGroupRule {
	return &SecurityGroupRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecurityGroupRule)(nil)

// SecurityGroupRule represents the Terraform resource aws_security_group_rule.
type SecurityGroupRule struct {
	Name      string
	Args      SecurityGroupRuleArgs
	state     *securityGroupRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecurityGroupRule].
func (sgr *SecurityGroupRule) Type() string {
	return "aws_security_group_rule"
}

// LocalName returns the local name for [SecurityGroupRule].
func (sgr *SecurityGroupRule) LocalName() string {
	return sgr.Name
}

// Configuration returns the configuration (args) for [SecurityGroupRule].
func (sgr *SecurityGroupRule) Configuration() interface{} {
	return sgr.Args
}

// DependOn is used for other resources to depend on [SecurityGroupRule].
func (sgr *SecurityGroupRule) DependOn() terra.Reference {
	return terra.ReferenceResource(sgr)
}

// Dependencies returns the list of resources [SecurityGroupRule] depends_on.
func (sgr *SecurityGroupRule) Dependencies() terra.Dependencies {
	return sgr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecurityGroupRule].
func (sgr *SecurityGroupRule) LifecycleManagement() *terra.Lifecycle {
	return sgr.Lifecycle
}

// Attributes returns the attributes for [SecurityGroupRule].
func (sgr *SecurityGroupRule) Attributes() securityGroupRuleAttributes {
	return securityGroupRuleAttributes{ref: terra.ReferenceResource(sgr)}
}

// ImportState imports the given attribute values into [SecurityGroupRule]'s state.
func (sgr *SecurityGroupRule) ImportState(av io.Reader) error {
	sgr.state = &securityGroupRuleState{}
	if err := json.NewDecoder(av).Decode(sgr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sgr.Type(), sgr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecurityGroupRule] has state.
func (sgr *SecurityGroupRule) State() (*securityGroupRuleState, bool) {
	return sgr.state, sgr.state != nil
}

// StateMust returns the state for [SecurityGroupRule]. Panics if the state is nil.
func (sgr *SecurityGroupRule) StateMust() *securityGroupRuleState {
	if sgr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sgr.Type(), sgr.LocalName()))
	}
	return sgr.state
}

// SecurityGroupRuleArgs contains the configurations for aws_security_group_rule.
type SecurityGroupRuleArgs struct {
	// CidrBlocks: list of string, optional
	CidrBlocks terra.ListValue[terra.StringValue] `hcl:"cidr_blocks,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FromPort: number, required
	FromPort terra.NumberValue `hcl:"from_port,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Ipv6CidrBlocks: list of string, optional
	Ipv6CidrBlocks terra.ListValue[terra.StringValue] `hcl:"ipv6_cidr_blocks,attr"`
	// PrefixListIds: list of string, optional
	PrefixListIds terra.ListValue[terra.StringValue] `hcl:"prefix_list_ids,attr"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
	// SecurityGroupId: string, required
	SecurityGroupId terra.StringValue `hcl:"security_group_id,attr" validate:"required"`
	// Self: bool, optional
	Self terra.BoolValue `hcl:"self,attr"`
	// SourceSecurityGroupId: string, optional
	SourceSecurityGroupId terra.StringValue `hcl:"source_security_group_id,attr"`
	// ToPort: number, required
	ToPort terra.NumberValue `hcl:"to_port,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *securitygrouprule.Timeouts `hcl:"timeouts,block"`
}
type securityGroupRuleAttributes struct {
	ref terra.Reference
}

// CidrBlocks returns a reference to field cidr_blocks of aws_security_group_rule.
func (sgr securityGroupRuleAttributes) CidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sgr.ref.Append("cidr_blocks"))
}

// Description returns a reference to field description of aws_security_group_rule.
func (sgr securityGroupRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sgr.ref.Append("description"))
}

// FromPort returns a reference to field from_port of aws_security_group_rule.
func (sgr securityGroupRuleAttributes) FromPort() terra.NumberValue {
	return terra.ReferenceAsNumber(sgr.ref.Append("from_port"))
}

// Id returns a reference to field id of aws_security_group_rule.
func (sgr securityGroupRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sgr.ref.Append("id"))
}

// Ipv6CidrBlocks returns a reference to field ipv6_cidr_blocks of aws_security_group_rule.
func (sgr securityGroupRuleAttributes) Ipv6CidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sgr.ref.Append("ipv6_cidr_blocks"))
}

// PrefixListIds returns a reference to field prefix_list_ids of aws_security_group_rule.
func (sgr securityGroupRuleAttributes) PrefixListIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sgr.ref.Append("prefix_list_ids"))
}

// Protocol returns a reference to field protocol of aws_security_group_rule.
func (sgr securityGroupRuleAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(sgr.ref.Append("protocol"))
}

// SecurityGroupId returns a reference to field security_group_id of aws_security_group_rule.
func (sgr securityGroupRuleAttributes) SecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(sgr.ref.Append("security_group_id"))
}

// SecurityGroupRuleId returns a reference to field security_group_rule_id of aws_security_group_rule.
func (sgr securityGroupRuleAttributes) SecurityGroupRuleId() terra.StringValue {
	return terra.ReferenceAsString(sgr.ref.Append("security_group_rule_id"))
}

// Self returns a reference to field self of aws_security_group_rule.
func (sgr securityGroupRuleAttributes) Self() terra.BoolValue {
	return terra.ReferenceAsBool(sgr.ref.Append("self"))
}

// SourceSecurityGroupId returns a reference to field source_security_group_id of aws_security_group_rule.
func (sgr securityGroupRuleAttributes) SourceSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(sgr.ref.Append("source_security_group_id"))
}

// ToPort returns a reference to field to_port of aws_security_group_rule.
func (sgr securityGroupRuleAttributes) ToPort() terra.NumberValue {
	return terra.ReferenceAsNumber(sgr.ref.Append("to_port"))
}

// Type returns a reference to field type of aws_security_group_rule.
func (sgr securityGroupRuleAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sgr.ref.Append("type"))
}

func (sgr securityGroupRuleAttributes) Timeouts() securitygrouprule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[securitygrouprule.TimeoutsAttributes](sgr.ref.Append("timeouts"))
}

type securityGroupRuleState struct {
	CidrBlocks            []string                         `json:"cidr_blocks"`
	Description           string                           `json:"description"`
	FromPort              float64                          `json:"from_port"`
	Id                    string                           `json:"id"`
	Ipv6CidrBlocks        []string                         `json:"ipv6_cidr_blocks"`
	PrefixListIds         []string                         `json:"prefix_list_ids"`
	Protocol              string                           `json:"protocol"`
	SecurityGroupId       string                           `json:"security_group_id"`
	SecurityGroupRuleId   string                           `json:"security_group_rule_id"`
	Self                  bool                             `json:"self"`
	SourceSecurityGroupId string                           `json:"source_security_group_id"`
	ToPort                float64                          `json:"to_port"`
	Type                  string                           `json:"type"`
	Timeouts              *securitygrouprule.TimeoutsState `json:"timeouts"`
}
