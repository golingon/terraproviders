// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lexintent "github.com/golingon/terraproviders/aws/5.0.1/lexintent"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLexIntent creates a new instance of [LexIntent].
func NewLexIntent(name string, args LexIntentArgs) *LexIntent {
	return &LexIntent{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LexIntent)(nil)

// LexIntent represents the Terraform resource aws_lex_intent.
type LexIntent struct {
	Name      string
	Args      LexIntentArgs
	state     *lexIntentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LexIntent].
func (li *LexIntent) Type() string {
	return "aws_lex_intent"
}

// LocalName returns the local name for [LexIntent].
func (li *LexIntent) LocalName() string {
	return li.Name
}

// Configuration returns the configuration (args) for [LexIntent].
func (li *LexIntent) Configuration() interface{} {
	return li.Args
}

// DependOn is used for other resources to depend on [LexIntent].
func (li *LexIntent) DependOn() terra.Reference {
	return terra.ReferenceResource(li)
}

// Dependencies returns the list of resources [LexIntent] depends_on.
func (li *LexIntent) Dependencies() terra.Dependencies {
	return li.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LexIntent].
func (li *LexIntent) LifecycleManagement() *terra.Lifecycle {
	return li.Lifecycle
}

// Attributes returns the attributes for [LexIntent].
func (li *LexIntent) Attributes() lexIntentAttributes {
	return lexIntentAttributes{ref: terra.ReferenceResource(li)}
}

// ImportState imports the given attribute values into [LexIntent]'s state.
func (li *LexIntent) ImportState(av io.Reader) error {
	li.state = &lexIntentState{}
	if err := json.NewDecoder(av).Decode(li.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", li.Type(), li.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LexIntent] has state.
func (li *LexIntent) State() (*lexIntentState, bool) {
	return li.state, li.state != nil
}

// StateMust returns the state for [LexIntent]. Panics if the state is nil.
func (li *LexIntent) StateMust() *lexIntentState {
	if li.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", li.Type(), li.LocalName()))
	}
	return li.state
}

// LexIntentArgs contains the configurations for aws_lex_intent.
type LexIntentArgs struct {
	// CreateVersion: bool, optional
	CreateVersion terra.BoolValue `hcl:"create_version,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParentIntentSignature: string, optional
	ParentIntentSignature terra.StringValue `hcl:"parent_intent_signature,attr"`
	// SampleUtterances: set of string, optional
	SampleUtterances terra.SetValue[terra.StringValue] `hcl:"sample_utterances,attr"`
	// ConclusionStatement: optional
	ConclusionStatement *lexintent.ConclusionStatement `hcl:"conclusion_statement,block"`
	// ConfirmationPrompt: optional
	ConfirmationPrompt *lexintent.ConfirmationPrompt `hcl:"confirmation_prompt,block"`
	// DialogCodeHook: optional
	DialogCodeHook *lexintent.DialogCodeHook `hcl:"dialog_code_hook,block"`
	// FollowUpPrompt: optional
	FollowUpPrompt *lexintent.FollowUpPrompt `hcl:"follow_up_prompt,block"`
	// FulfillmentActivity: required
	FulfillmentActivity *lexintent.FulfillmentActivity `hcl:"fulfillment_activity,block" validate:"required"`
	// RejectionStatement: optional
	RejectionStatement *lexintent.RejectionStatement `hcl:"rejection_statement,block"`
	// Slot: min=0,max=100
	Slot []lexintent.Slot `hcl:"slot,block" validate:"min=0,max=100"`
	// Timeouts: optional
	Timeouts *lexintent.Timeouts `hcl:"timeouts,block"`
}
type lexIntentAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lex_intent.
func (li lexIntentAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(li.ref.Append("arn"))
}

// Checksum returns a reference to field checksum of aws_lex_intent.
func (li lexIntentAttributes) Checksum() terra.StringValue {
	return terra.ReferenceAsString(li.ref.Append("checksum"))
}

// CreateVersion returns a reference to field create_version of aws_lex_intent.
func (li lexIntentAttributes) CreateVersion() terra.BoolValue {
	return terra.ReferenceAsBool(li.ref.Append("create_version"))
}

// CreatedDate returns a reference to field created_date of aws_lex_intent.
func (li lexIntentAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(li.ref.Append("created_date"))
}

// Description returns a reference to field description of aws_lex_intent.
func (li lexIntentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(li.ref.Append("description"))
}

// Id returns a reference to field id of aws_lex_intent.
func (li lexIntentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(li.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_lex_intent.
func (li lexIntentAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(li.ref.Append("last_updated_date"))
}

// Name returns a reference to field name of aws_lex_intent.
func (li lexIntentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(li.ref.Append("name"))
}

// ParentIntentSignature returns a reference to field parent_intent_signature of aws_lex_intent.
func (li lexIntentAttributes) ParentIntentSignature() terra.StringValue {
	return terra.ReferenceAsString(li.ref.Append("parent_intent_signature"))
}

// SampleUtterances returns a reference to field sample_utterances of aws_lex_intent.
func (li lexIntentAttributes) SampleUtterances() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](li.ref.Append("sample_utterances"))
}

// Version returns a reference to field version of aws_lex_intent.
func (li lexIntentAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(li.ref.Append("version"))
}

func (li lexIntentAttributes) ConclusionStatement() terra.ListValue[lexintent.ConclusionStatementAttributes] {
	return terra.ReferenceAsList[lexintent.ConclusionStatementAttributes](li.ref.Append("conclusion_statement"))
}

func (li lexIntentAttributes) ConfirmationPrompt() terra.ListValue[lexintent.ConfirmationPromptAttributes] {
	return terra.ReferenceAsList[lexintent.ConfirmationPromptAttributes](li.ref.Append("confirmation_prompt"))
}

func (li lexIntentAttributes) DialogCodeHook() terra.ListValue[lexintent.DialogCodeHookAttributes] {
	return terra.ReferenceAsList[lexintent.DialogCodeHookAttributes](li.ref.Append("dialog_code_hook"))
}

func (li lexIntentAttributes) FollowUpPrompt() terra.ListValue[lexintent.FollowUpPromptAttributes] {
	return terra.ReferenceAsList[lexintent.FollowUpPromptAttributes](li.ref.Append("follow_up_prompt"))
}

func (li lexIntentAttributes) FulfillmentActivity() terra.ListValue[lexintent.FulfillmentActivityAttributes] {
	return terra.ReferenceAsList[lexintent.FulfillmentActivityAttributes](li.ref.Append("fulfillment_activity"))
}

func (li lexIntentAttributes) RejectionStatement() terra.ListValue[lexintent.RejectionStatementAttributes] {
	return terra.ReferenceAsList[lexintent.RejectionStatementAttributes](li.ref.Append("rejection_statement"))
}

func (li lexIntentAttributes) Slot() terra.SetValue[lexintent.SlotAttributes] {
	return terra.ReferenceAsSet[lexintent.SlotAttributes](li.ref.Append("slot"))
}

func (li lexIntentAttributes) Timeouts() lexintent.TimeoutsAttributes {
	return terra.ReferenceAsSingle[lexintent.TimeoutsAttributes](li.ref.Append("timeouts"))
}

type lexIntentState struct {
	Arn                   string                               `json:"arn"`
	Checksum              string                               `json:"checksum"`
	CreateVersion         bool                                 `json:"create_version"`
	CreatedDate           string                               `json:"created_date"`
	Description           string                               `json:"description"`
	Id                    string                               `json:"id"`
	LastUpdatedDate       string                               `json:"last_updated_date"`
	Name                  string                               `json:"name"`
	ParentIntentSignature string                               `json:"parent_intent_signature"`
	SampleUtterances      []string                             `json:"sample_utterances"`
	Version               string                               `json:"version"`
	ConclusionStatement   []lexintent.ConclusionStatementState `json:"conclusion_statement"`
	ConfirmationPrompt    []lexintent.ConfirmationPromptState  `json:"confirmation_prompt"`
	DialogCodeHook        []lexintent.DialogCodeHookState      `json:"dialog_code_hook"`
	FollowUpPrompt        []lexintent.FollowUpPromptState      `json:"follow_up_prompt"`
	FulfillmentActivity   []lexintent.FulfillmentActivityState `json:"fulfillment_activity"`
	RejectionStatement    []lexintent.RejectionStatementState  `json:"rejection_statement"`
	Slot                  []lexintent.SlotState                `json:"slot"`
	Timeouts              *lexintent.TimeoutsState             `json:"timeouts"`
}
