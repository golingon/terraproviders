// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSesActiveReceiptRuleSet creates a new instance of [SesActiveReceiptRuleSet].
func NewSesActiveReceiptRuleSet(name string, args SesActiveReceiptRuleSetArgs) *SesActiveReceiptRuleSet {
	return &SesActiveReceiptRuleSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SesActiveReceiptRuleSet)(nil)

// SesActiveReceiptRuleSet represents the Terraform resource aws_ses_active_receipt_rule_set.
type SesActiveReceiptRuleSet struct {
	Name      string
	Args      SesActiveReceiptRuleSetArgs
	state     *sesActiveReceiptRuleSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SesActiveReceiptRuleSet].
func (sarrs *SesActiveReceiptRuleSet) Type() string {
	return "aws_ses_active_receipt_rule_set"
}

// LocalName returns the local name for [SesActiveReceiptRuleSet].
func (sarrs *SesActiveReceiptRuleSet) LocalName() string {
	return sarrs.Name
}

// Configuration returns the configuration (args) for [SesActiveReceiptRuleSet].
func (sarrs *SesActiveReceiptRuleSet) Configuration() interface{} {
	return sarrs.Args
}

// DependOn is used for other resources to depend on [SesActiveReceiptRuleSet].
func (sarrs *SesActiveReceiptRuleSet) DependOn() terra.Reference {
	return terra.ReferenceResource(sarrs)
}

// Dependencies returns the list of resources [SesActiveReceiptRuleSet] depends_on.
func (sarrs *SesActiveReceiptRuleSet) Dependencies() terra.Dependencies {
	return sarrs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SesActiveReceiptRuleSet].
func (sarrs *SesActiveReceiptRuleSet) LifecycleManagement() *terra.Lifecycle {
	return sarrs.Lifecycle
}

// Attributes returns the attributes for [SesActiveReceiptRuleSet].
func (sarrs *SesActiveReceiptRuleSet) Attributes() sesActiveReceiptRuleSetAttributes {
	return sesActiveReceiptRuleSetAttributes{ref: terra.ReferenceResource(sarrs)}
}

// ImportState imports the given attribute values into [SesActiveReceiptRuleSet]'s state.
func (sarrs *SesActiveReceiptRuleSet) ImportState(av io.Reader) error {
	sarrs.state = &sesActiveReceiptRuleSetState{}
	if err := json.NewDecoder(av).Decode(sarrs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sarrs.Type(), sarrs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SesActiveReceiptRuleSet] has state.
func (sarrs *SesActiveReceiptRuleSet) State() (*sesActiveReceiptRuleSetState, bool) {
	return sarrs.state, sarrs.state != nil
}

// StateMust returns the state for [SesActiveReceiptRuleSet]. Panics if the state is nil.
func (sarrs *SesActiveReceiptRuleSet) StateMust() *sesActiveReceiptRuleSetState {
	if sarrs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sarrs.Type(), sarrs.LocalName()))
	}
	return sarrs.state
}

// SesActiveReceiptRuleSetArgs contains the configurations for aws_ses_active_receipt_rule_set.
type SesActiveReceiptRuleSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RuleSetName: string, required
	RuleSetName terra.StringValue `hcl:"rule_set_name,attr" validate:"required"`
}
type sesActiveReceiptRuleSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ses_active_receipt_rule_set.
func (sarrs sesActiveReceiptRuleSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sarrs.ref.Append("arn"))
}

// Id returns a reference to field id of aws_ses_active_receipt_rule_set.
func (sarrs sesActiveReceiptRuleSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sarrs.ref.Append("id"))
}

// RuleSetName returns a reference to field rule_set_name of aws_ses_active_receipt_rule_set.
func (sarrs sesActiveReceiptRuleSetAttributes) RuleSetName() terra.StringValue {
	return terra.ReferenceAsString(sarrs.ref.Append("rule_set_name"))
}

type sesActiveReceiptRuleSetState struct {
	Arn         string `json:"arn"`
	Id          string `json:"id"`
	RuleSetName string `json:"rule_set_name"`
}
