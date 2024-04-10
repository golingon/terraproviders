// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	wafv2rulegroup "github.com/golingon/terraproviders/aws/5.44.0/wafv2rulegroup"
	"io"
)

// NewWafv2RuleGroup creates a new instance of [Wafv2RuleGroup].
func NewWafv2RuleGroup(name string, args Wafv2RuleGroupArgs) *Wafv2RuleGroup {
	return &Wafv2RuleGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Wafv2RuleGroup)(nil)

// Wafv2RuleGroup represents the Terraform resource aws_wafv2_rule_group.
type Wafv2RuleGroup struct {
	Name      string
	Args      Wafv2RuleGroupArgs
	state     *wafv2RuleGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Wafv2RuleGroup].
func (wrg *Wafv2RuleGroup) Type() string {
	return "aws_wafv2_rule_group"
}

// LocalName returns the local name for [Wafv2RuleGroup].
func (wrg *Wafv2RuleGroup) LocalName() string {
	return wrg.Name
}

// Configuration returns the configuration (args) for [Wafv2RuleGroup].
func (wrg *Wafv2RuleGroup) Configuration() interface{} {
	return wrg.Args
}

// DependOn is used for other resources to depend on [Wafv2RuleGroup].
func (wrg *Wafv2RuleGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(wrg)
}

// Dependencies returns the list of resources [Wafv2RuleGroup] depends_on.
func (wrg *Wafv2RuleGroup) Dependencies() terra.Dependencies {
	return wrg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Wafv2RuleGroup].
func (wrg *Wafv2RuleGroup) LifecycleManagement() *terra.Lifecycle {
	return wrg.Lifecycle
}

// Attributes returns the attributes for [Wafv2RuleGroup].
func (wrg *Wafv2RuleGroup) Attributes() wafv2RuleGroupAttributes {
	return wafv2RuleGroupAttributes{ref: terra.ReferenceResource(wrg)}
}

// ImportState imports the given attribute values into [Wafv2RuleGroup]'s state.
func (wrg *Wafv2RuleGroup) ImportState(av io.Reader) error {
	wrg.state = &wafv2RuleGroupState{}
	if err := json.NewDecoder(av).Decode(wrg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wrg.Type(), wrg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Wafv2RuleGroup] has state.
func (wrg *Wafv2RuleGroup) State() (*wafv2RuleGroupState, bool) {
	return wrg.state, wrg.state != nil
}

// StateMust returns the state for [Wafv2RuleGroup]. Panics if the state is nil.
func (wrg *Wafv2RuleGroup) StateMust() *wafv2RuleGroupState {
	if wrg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wrg.Type(), wrg.LocalName()))
	}
	return wrg.state
}

// Wafv2RuleGroupArgs contains the configurations for aws_wafv2_rule_group.
type Wafv2RuleGroupArgs struct {
	// Capacity: number, required
	Capacity terra.NumberValue `hcl:"capacity,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// CustomResponseBody: min=0
	CustomResponseBody []wafv2rulegroup.CustomResponseBody `hcl:"custom_response_body,block" validate:"min=0"`
	// Rule: min=0
	Rule []wafv2rulegroup.Rule `hcl:"rule,block" validate:"min=0"`
	// VisibilityConfig: required
	VisibilityConfig *wafv2rulegroup.VisibilityConfig `hcl:"visibility_config,block" validate:"required"`
}
type wafv2RuleGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_wafv2_rule_group.
func (wrg wafv2RuleGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("arn"))
}

// Capacity returns a reference to field capacity of aws_wafv2_rule_group.
func (wrg wafv2RuleGroupAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(wrg.ref.Append("capacity"))
}

// Description returns a reference to field description of aws_wafv2_rule_group.
func (wrg wafv2RuleGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("description"))
}

// Id returns a reference to field id of aws_wafv2_rule_group.
func (wrg wafv2RuleGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("id"))
}

// LockToken returns a reference to field lock_token of aws_wafv2_rule_group.
func (wrg wafv2RuleGroupAttributes) LockToken() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("lock_token"))
}

// Name returns a reference to field name of aws_wafv2_rule_group.
func (wrg wafv2RuleGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_wafv2_rule_group.
func (wrg wafv2RuleGroupAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("name_prefix"))
}

// Scope returns a reference to field scope of aws_wafv2_rule_group.
func (wrg wafv2RuleGroupAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("scope"))
}

// Tags returns a reference to field tags of aws_wafv2_rule_group.
func (wrg wafv2RuleGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wrg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_wafv2_rule_group.
func (wrg wafv2RuleGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wrg.ref.Append("tags_all"))
}

func (wrg wafv2RuleGroupAttributes) CustomResponseBody() terra.SetValue[wafv2rulegroup.CustomResponseBodyAttributes] {
	return terra.ReferenceAsSet[wafv2rulegroup.CustomResponseBodyAttributes](wrg.ref.Append("custom_response_body"))
}

func (wrg wafv2RuleGroupAttributes) Rule() terra.SetValue[wafv2rulegroup.RuleAttributes] {
	return terra.ReferenceAsSet[wafv2rulegroup.RuleAttributes](wrg.ref.Append("rule"))
}

func (wrg wafv2RuleGroupAttributes) VisibilityConfig() terra.ListValue[wafv2rulegroup.VisibilityConfigAttributes] {
	return terra.ReferenceAsList[wafv2rulegroup.VisibilityConfigAttributes](wrg.ref.Append("visibility_config"))
}

type wafv2RuleGroupState struct {
	Arn                string                                   `json:"arn"`
	Capacity           float64                                  `json:"capacity"`
	Description        string                                   `json:"description"`
	Id                 string                                   `json:"id"`
	LockToken          string                                   `json:"lock_token"`
	Name               string                                   `json:"name"`
	NamePrefix         string                                   `json:"name_prefix"`
	Scope              string                                   `json:"scope"`
	Tags               map[string]string                        `json:"tags"`
	TagsAll            map[string]string                        `json:"tags_all"`
	CustomResponseBody []wafv2rulegroup.CustomResponseBodyState `json:"custom_response_body"`
	Rule               []wafv2rulegroup.RuleState               `json:"rule"`
	VisibilityConfig   []wafv2rulegroup.VisibilityConfigState   `json:"visibility_config"`
}
