// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	rbinrule "github.com/golingon/terraproviders/aws/5.44.0/rbinrule"
	"io"
)

// NewRbinRule creates a new instance of [RbinRule].
func NewRbinRule(name string, args RbinRuleArgs) *RbinRule {
	return &RbinRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RbinRule)(nil)

// RbinRule represents the Terraform resource aws_rbin_rule.
type RbinRule struct {
	Name      string
	Args      RbinRuleArgs
	state     *rbinRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RbinRule].
func (rr *RbinRule) Type() string {
	return "aws_rbin_rule"
}

// LocalName returns the local name for [RbinRule].
func (rr *RbinRule) LocalName() string {
	return rr.Name
}

// Configuration returns the configuration (args) for [RbinRule].
func (rr *RbinRule) Configuration() interface{} {
	return rr.Args
}

// DependOn is used for other resources to depend on [RbinRule].
func (rr *RbinRule) DependOn() terra.Reference {
	return terra.ReferenceResource(rr)
}

// Dependencies returns the list of resources [RbinRule] depends_on.
func (rr *RbinRule) Dependencies() terra.Dependencies {
	return rr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RbinRule].
func (rr *RbinRule) LifecycleManagement() *terra.Lifecycle {
	return rr.Lifecycle
}

// Attributes returns the attributes for [RbinRule].
func (rr *RbinRule) Attributes() rbinRuleAttributes {
	return rbinRuleAttributes{ref: terra.ReferenceResource(rr)}
}

// ImportState imports the given attribute values into [RbinRule]'s state.
func (rr *RbinRule) ImportState(av io.Reader) error {
	rr.state = &rbinRuleState{}
	if err := json.NewDecoder(av).Decode(rr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rr.Type(), rr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RbinRule] has state.
func (rr *RbinRule) State() (*rbinRuleState, bool) {
	return rr.state, rr.state != nil
}

// StateMust returns the state for [RbinRule]. Panics if the state is nil.
func (rr *RbinRule) StateMust() *rbinRuleState {
	if rr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rr.Type(), rr.LocalName()))
	}
	return rr.state
}

// RbinRuleArgs contains the configurations for aws_rbin_rule.
type RbinRuleArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ResourceType: string, required
	ResourceType terra.StringValue `hcl:"resource_type,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// LockConfiguration: optional
	LockConfiguration *rbinrule.LockConfiguration `hcl:"lock_configuration,block"`
	// ResourceTags: min=0,max=50
	ResourceTags []rbinrule.ResourceTags `hcl:"resource_tags,block" validate:"min=0,max=50"`
	// RetentionPeriod: required
	RetentionPeriod *rbinrule.RetentionPeriod `hcl:"retention_period,block" validate:"required"`
	// Timeouts: optional
	Timeouts *rbinrule.Timeouts `hcl:"timeouts,block"`
}
type rbinRuleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_rbin_rule.
func (rr rbinRuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("arn"))
}

// Description returns a reference to field description of aws_rbin_rule.
func (rr rbinRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("description"))
}

// Id returns a reference to field id of aws_rbin_rule.
func (rr rbinRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("id"))
}

// LockEndTime returns a reference to field lock_end_time of aws_rbin_rule.
func (rr rbinRuleAttributes) LockEndTime() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("lock_end_time"))
}

// LockState returns a reference to field lock_state of aws_rbin_rule.
func (rr rbinRuleAttributes) LockState() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("lock_state"))
}

// ResourceType returns a reference to field resource_type of aws_rbin_rule.
func (rr rbinRuleAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("resource_type"))
}

// Status returns a reference to field status of aws_rbin_rule.
func (rr rbinRuleAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_rbin_rule.
func (rr rbinRuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_rbin_rule.
func (rr rbinRuleAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rr.ref.Append("tags_all"))
}

func (rr rbinRuleAttributes) LockConfiguration() terra.ListValue[rbinrule.LockConfigurationAttributes] {
	return terra.ReferenceAsList[rbinrule.LockConfigurationAttributes](rr.ref.Append("lock_configuration"))
}

func (rr rbinRuleAttributes) ResourceTags() terra.SetValue[rbinrule.ResourceTagsAttributes] {
	return terra.ReferenceAsSet[rbinrule.ResourceTagsAttributes](rr.ref.Append("resource_tags"))
}

func (rr rbinRuleAttributes) RetentionPeriod() terra.ListValue[rbinrule.RetentionPeriodAttributes] {
	return terra.ReferenceAsList[rbinrule.RetentionPeriodAttributes](rr.ref.Append("retention_period"))
}

func (rr rbinRuleAttributes) Timeouts() rbinrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[rbinrule.TimeoutsAttributes](rr.ref.Append("timeouts"))
}

type rbinRuleState struct {
	Arn               string                            `json:"arn"`
	Description       string                            `json:"description"`
	Id                string                            `json:"id"`
	LockEndTime       string                            `json:"lock_end_time"`
	LockState         string                            `json:"lock_state"`
	ResourceType      string                            `json:"resource_type"`
	Status            string                            `json:"status"`
	Tags              map[string]string                 `json:"tags"`
	TagsAll           map[string]string                 `json:"tags_all"`
	LockConfiguration []rbinrule.LockConfigurationState `json:"lock_configuration"`
	ResourceTags      []rbinrule.ResourceTagsState      `json:"resource_tags"`
	RetentionPeriod   []rbinrule.RetentionPeriodState   `json:"retention_period"`
	Timeouts          *rbinrule.TimeoutsState           `json:"timeouts"`
}
