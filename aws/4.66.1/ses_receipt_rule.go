// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sesreceiptrule "github.com/golingon/terraproviders/aws/4.66.1/sesreceiptrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSesReceiptRule creates a new instance of [SesReceiptRule].
func NewSesReceiptRule(name string, args SesReceiptRuleArgs) *SesReceiptRule {
	return &SesReceiptRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SesReceiptRule)(nil)

// SesReceiptRule represents the Terraform resource aws_ses_receipt_rule.
type SesReceiptRule struct {
	Name      string
	Args      SesReceiptRuleArgs
	state     *sesReceiptRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SesReceiptRule].
func (srr *SesReceiptRule) Type() string {
	return "aws_ses_receipt_rule"
}

// LocalName returns the local name for [SesReceiptRule].
func (srr *SesReceiptRule) LocalName() string {
	return srr.Name
}

// Configuration returns the configuration (args) for [SesReceiptRule].
func (srr *SesReceiptRule) Configuration() interface{} {
	return srr.Args
}

// DependOn is used for other resources to depend on [SesReceiptRule].
func (srr *SesReceiptRule) DependOn() terra.Reference {
	return terra.ReferenceResource(srr)
}

// Dependencies returns the list of resources [SesReceiptRule] depends_on.
func (srr *SesReceiptRule) Dependencies() terra.Dependencies {
	return srr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SesReceiptRule].
func (srr *SesReceiptRule) LifecycleManagement() *terra.Lifecycle {
	return srr.Lifecycle
}

// Attributes returns the attributes for [SesReceiptRule].
func (srr *SesReceiptRule) Attributes() sesReceiptRuleAttributes {
	return sesReceiptRuleAttributes{ref: terra.ReferenceResource(srr)}
}

// ImportState imports the given attribute values into [SesReceiptRule]'s state.
func (srr *SesReceiptRule) ImportState(av io.Reader) error {
	srr.state = &sesReceiptRuleState{}
	if err := json.NewDecoder(av).Decode(srr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srr.Type(), srr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SesReceiptRule] has state.
func (srr *SesReceiptRule) State() (*sesReceiptRuleState, bool) {
	return srr.state, srr.state != nil
}

// StateMust returns the state for [SesReceiptRule]. Panics if the state is nil.
func (srr *SesReceiptRule) StateMust() *sesReceiptRuleState {
	if srr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srr.Type(), srr.LocalName()))
	}
	return srr.state
}

// SesReceiptRuleArgs contains the configurations for aws_ses_receipt_rule.
type SesReceiptRuleArgs struct {
	// After: string, optional
	After terra.StringValue `hcl:"after,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Recipients: set of string, optional
	Recipients terra.SetValue[terra.StringValue] `hcl:"recipients,attr"`
	// RuleSetName: string, required
	RuleSetName terra.StringValue `hcl:"rule_set_name,attr" validate:"required"`
	// ScanEnabled: bool, optional
	ScanEnabled terra.BoolValue `hcl:"scan_enabled,attr"`
	// TlsPolicy: string, optional
	TlsPolicy terra.StringValue `hcl:"tls_policy,attr"`
	// AddHeaderAction: min=0
	AddHeaderAction []sesreceiptrule.AddHeaderAction `hcl:"add_header_action,block" validate:"min=0"`
	// BounceAction: min=0
	BounceAction []sesreceiptrule.BounceAction `hcl:"bounce_action,block" validate:"min=0"`
	// LambdaAction: min=0
	LambdaAction []sesreceiptrule.LambdaAction `hcl:"lambda_action,block" validate:"min=0"`
	// S3Action: min=0
	S3Action []sesreceiptrule.S3Action `hcl:"s3_action,block" validate:"min=0"`
	// SnsAction: min=0
	SnsAction []sesreceiptrule.SnsAction `hcl:"sns_action,block" validate:"min=0"`
	// StopAction: min=0
	StopAction []sesreceiptrule.StopAction `hcl:"stop_action,block" validate:"min=0"`
	// WorkmailAction: min=0
	WorkmailAction []sesreceiptrule.WorkmailAction `hcl:"workmail_action,block" validate:"min=0"`
}
type sesReceiptRuleAttributes struct {
	ref terra.Reference
}

// After returns a reference to field after of aws_ses_receipt_rule.
func (srr sesReceiptRuleAttributes) After() terra.StringValue {
	return terra.ReferenceAsString(srr.ref.Append("after"))
}

// Arn returns a reference to field arn of aws_ses_receipt_rule.
func (srr sesReceiptRuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(srr.ref.Append("arn"))
}

// Enabled returns a reference to field enabled of aws_ses_receipt_rule.
func (srr sesReceiptRuleAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(srr.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_ses_receipt_rule.
func (srr sesReceiptRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srr.ref.Append("id"))
}

// Name returns a reference to field name of aws_ses_receipt_rule.
func (srr sesReceiptRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srr.ref.Append("name"))
}

// Recipients returns a reference to field recipients of aws_ses_receipt_rule.
func (srr sesReceiptRuleAttributes) Recipients() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](srr.ref.Append("recipients"))
}

// RuleSetName returns a reference to field rule_set_name of aws_ses_receipt_rule.
func (srr sesReceiptRuleAttributes) RuleSetName() terra.StringValue {
	return terra.ReferenceAsString(srr.ref.Append("rule_set_name"))
}

// ScanEnabled returns a reference to field scan_enabled of aws_ses_receipt_rule.
func (srr sesReceiptRuleAttributes) ScanEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(srr.ref.Append("scan_enabled"))
}

// TlsPolicy returns a reference to field tls_policy of aws_ses_receipt_rule.
func (srr sesReceiptRuleAttributes) TlsPolicy() terra.StringValue {
	return terra.ReferenceAsString(srr.ref.Append("tls_policy"))
}

func (srr sesReceiptRuleAttributes) AddHeaderAction() terra.SetValue[sesreceiptrule.AddHeaderActionAttributes] {
	return terra.ReferenceAsSet[sesreceiptrule.AddHeaderActionAttributes](srr.ref.Append("add_header_action"))
}

func (srr sesReceiptRuleAttributes) BounceAction() terra.SetValue[sesreceiptrule.BounceActionAttributes] {
	return terra.ReferenceAsSet[sesreceiptrule.BounceActionAttributes](srr.ref.Append("bounce_action"))
}

func (srr sesReceiptRuleAttributes) LambdaAction() terra.SetValue[sesreceiptrule.LambdaActionAttributes] {
	return terra.ReferenceAsSet[sesreceiptrule.LambdaActionAttributes](srr.ref.Append("lambda_action"))
}

func (srr sesReceiptRuleAttributes) S3Action() terra.SetValue[sesreceiptrule.S3ActionAttributes] {
	return terra.ReferenceAsSet[sesreceiptrule.S3ActionAttributes](srr.ref.Append("s3_action"))
}

func (srr sesReceiptRuleAttributes) SnsAction() terra.SetValue[sesreceiptrule.SnsActionAttributes] {
	return terra.ReferenceAsSet[sesreceiptrule.SnsActionAttributes](srr.ref.Append("sns_action"))
}

func (srr sesReceiptRuleAttributes) StopAction() terra.SetValue[sesreceiptrule.StopActionAttributes] {
	return terra.ReferenceAsSet[sesreceiptrule.StopActionAttributes](srr.ref.Append("stop_action"))
}

func (srr sesReceiptRuleAttributes) WorkmailAction() terra.SetValue[sesreceiptrule.WorkmailActionAttributes] {
	return terra.ReferenceAsSet[sesreceiptrule.WorkmailActionAttributes](srr.ref.Append("workmail_action"))
}

type sesReceiptRuleState struct {
	After           string                                `json:"after"`
	Arn             string                                `json:"arn"`
	Enabled         bool                                  `json:"enabled"`
	Id              string                                `json:"id"`
	Name            string                                `json:"name"`
	Recipients      []string                              `json:"recipients"`
	RuleSetName     string                                `json:"rule_set_name"`
	ScanEnabled     bool                                  `json:"scan_enabled"`
	TlsPolicy       string                                `json:"tls_policy"`
	AddHeaderAction []sesreceiptrule.AddHeaderActionState `json:"add_header_action"`
	BounceAction    []sesreceiptrule.BounceActionState    `json:"bounce_action"`
	LambdaAction    []sesreceiptrule.LambdaActionState    `json:"lambda_action"`
	S3Action        []sesreceiptrule.S3ActionState        `json:"s3_action"`
	SnsAction       []sesreceiptrule.SnsActionState       `json:"sns_action"`
	StopAction      []sesreceiptrule.StopActionState      `json:"stop_action"`
	WorkmailAction  []sesreceiptrule.WorkmailActionState  `json:"workmail_action"`
}
