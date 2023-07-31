// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	route53recoverycontrolconfigsafetyrule "github.com/golingon/terraproviders/aws/5.10.0/route53recoverycontrolconfigsafetyrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRoute53RecoverycontrolconfigSafetyRule creates a new instance of [Route53RecoverycontrolconfigSafetyRule].
func NewRoute53RecoverycontrolconfigSafetyRule(name string, args Route53RecoverycontrolconfigSafetyRuleArgs) *Route53RecoverycontrolconfigSafetyRule {
	return &Route53RecoverycontrolconfigSafetyRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53RecoverycontrolconfigSafetyRule)(nil)

// Route53RecoverycontrolconfigSafetyRule represents the Terraform resource aws_route53recoverycontrolconfig_safety_rule.
type Route53RecoverycontrolconfigSafetyRule struct {
	Name      string
	Args      Route53RecoverycontrolconfigSafetyRuleArgs
	state     *route53RecoverycontrolconfigSafetyRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53RecoverycontrolconfigSafetyRule].
func (rsr *Route53RecoverycontrolconfigSafetyRule) Type() string {
	return "aws_route53recoverycontrolconfig_safety_rule"
}

// LocalName returns the local name for [Route53RecoverycontrolconfigSafetyRule].
func (rsr *Route53RecoverycontrolconfigSafetyRule) LocalName() string {
	return rsr.Name
}

// Configuration returns the configuration (args) for [Route53RecoverycontrolconfigSafetyRule].
func (rsr *Route53RecoverycontrolconfigSafetyRule) Configuration() interface{} {
	return rsr.Args
}

// DependOn is used for other resources to depend on [Route53RecoverycontrolconfigSafetyRule].
func (rsr *Route53RecoverycontrolconfigSafetyRule) DependOn() terra.Reference {
	return terra.ReferenceResource(rsr)
}

// Dependencies returns the list of resources [Route53RecoverycontrolconfigSafetyRule] depends_on.
func (rsr *Route53RecoverycontrolconfigSafetyRule) Dependencies() terra.Dependencies {
	return rsr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53RecoverycontrolconfigSafetyRule].
func (rsr *Route53RecoverycontrolconfigSafetyRule) LifecycleManagement() *terra.Lifecycle {
	return rsr.Lifecycle
}

// Attributes returns the attributes for [Route53RecoverycontrolconfigSafetyRule].
func (rsr *Route53RecoverycontrolconfigSafetyRule) Attributes() route53RecoverycontrolconfigSafetyRuleAttributes {
	return route53RecoverycontrolconfigSafetyRuleAttributes{ref: terra.ReferenceResource(rsr)}
}

// ImportState imports the given attribute values into [Route53RecoverycontrolconfigSafetyRule]'s state.
func (rsr *Route53RecoverycontrolconfigSafetyRule) ImportState(av io.Reader) error {
	rsr.state = &route53RecoverycontrolconfigSafetyRuleState{}
	if err := json.NewDecoder(av).Decode(rsr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rsr.Type(), rsr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53RecoverycontrolconfigSafetyRule] has state.
func (rsr *Route53RecoverycontrolconfigSafetyRule) State() (*route53RecoverycontrolconfigSafetyRuleState, bool) {
	return rsr.state, rsr.state != nil
}

// StateMust returns the state for [Route53RecoverycontrolconfigSafetyRule]. Panics if the state is nil.
func (rsr *Route53RecoverycontrolconfigSafetyRule) StateMust() *route53RecoverycontrolconfigSafetyRuleState {
	if rsr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rsr.Type(), rsr.LocalName()))
	}
	return rsr.state
}

// Route53RecoverycontrolconfigSafetyRuleArgs contains the configurations for aws_route53recoverycontrolconfig_safety_rule.
type Route53RecoverycontrolconfigSafetyRuleArgs struct {
	// AssertedControls: list of string, optional
	AssertedControls terra.ListValue[terra.StringValue] `hcl:"asserted_controls,attr"`
	// ControlPanelArn: string, required
	ControlPanelArn terra.StringValue `hcl:"control_panel_arn,attr" validate:"required"`
	// GatingControls: list of string, optional
	GatingControls terra.ListValue[terra.StringValue] `hcl:"gating_controls,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TargetControls: list of string, optional
	TargetControls terra.ListValue[terra.StringValue] `hcl:"target_controls,attr"`
	// WaitPeriodMs: number, required
	WaitPeriodMs terra.NumberValue `hcl:"wait_period_ms,attr" validate:"required"`
	// RuleConfig: required
	RuleConfig *route53recoverycontrolconfigsafetyrule.RuleConfig `hcl:"rule_config,block" validate:"required"`
}
type route53RecoverycontrolconfigSafetyRuleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_route53recoverycontrolconfig_safety_rule.
func (rsr route53RecoverycontrolconfigSafetyRuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rsr.ref.Append("arn"))
}

// AssertedControls returns a reference to field asserted_controls of aws_route53recoverycontrolconfig_safety_rule.
func (rsr route53RecoverycontrolconfigSafetyRuleAttributes) AssertedControls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rsr.ref.Append("asserted_controls"))
}

// ControlPanelArn returns a reference to field control_panel_arn of aws_route53recoverycontrolconfig_safety_rule.
func (rsr route53RecoverycontrolconfigSafetyRuleAttributes) ControlPanelArn() terra.StringValue {
	return terra.ReferenceAsString(rsr.ref.Append("control_panel_arn"))
}

// GatingControls returns a reference to field gating_controls of aws_route53recoverycontrolconfig_safety_rule.
func (rsr route53RecoverycontrolconfigSafetyRuleAttributes) GatingControls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rsr.ref.Append("gating_controls"))
}

// Id returns a reference to field id of aws_route53recoverycontrolconfig_safety_rule.
func (rsr route53RecoverycontrolconfigSafetyRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rsr.ref.Append("id"))
}

// Name returns a reference to field name of aws_route53recoverycontrolconfig_safety_rule.
func (rsr route53RecoverycontrolconfigSafetyRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rsr.ref.Append("name"))
}

// Status returns a reference to field status of aws_route53recoverycontrolconfig_safety_rule.
func (rsr route53RecoverycontrolconfigSafetyRuleAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(rsr.ref.Append("status"))
}

// TargetControls returns a reference to field target_controls of aws_route53recoverycontrolconfig_safety_rule.
func (rsr route53RecoverycontrolconfigSafetyRuleAttributes) TargetControls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rsr.ref.Append("target_controls"))
}

// WaitPeriodMs returns a reference to field wait_period_ms of aws_route53recoverycontrolconfig_safety_rule.
func (rsr route53RecoverycontrolconfigSafetyRuleAttributes) WaitPeriodMs() terra.NumberValue {
	return terra.ReferenceAsNumber(rsr.ref.Append("wait_period_ms"))
}

func (rsr route53RecoverycontrolconfigSafetyRuleAttributes) RuleConfig() terra.ListValue[route53recoverycontrolconfigsafetyrule.RuleConfigAttributes] {
	return terra.ReferenceAsList[route53recoverycontrolconfigsafetyrule.RuleConfigAttributes](rsr.ref.Append("rule_config"))
}

type route53RecoverycontrolconfigSafetyRuleState struct {
	Arn              string                                                   `json:"arn"`
	AssertedControls []string                                                 `json:"asserted_controls"`
	ControlPanelArn  string                                                   `json:"control_panel_arn"`
	GatingControls   []string                                                 `json:"gating_controls"`
	Id               string                                                   `json:"id"`
	Name             string                                                   `json:"name"`
	Status           string                                                   `json:"status"`
	TargetControls   []string                                                 `json:"target_controls"`
	WaitPeriodMs     float64                                                  `json:"wait_period_ms"`
	RuleConfig       []route53recoverycontrolconfigsafetyrule.RuleConfigState `json:"rule_config"`
}
