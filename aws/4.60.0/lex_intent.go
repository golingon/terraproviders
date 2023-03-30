// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lexintent "github.com/golingon/terraproviders/aws/4.60.0/lexintent"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewLexIntent(name string, args LexIntentArgs) *LexIntent {
	return &LexIntent{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LexIntent)(nil)

type LexIntent struct {
	Name  string
	Args  LexIntentArgs
	state *lexIntentState
}

func (li *LexIntent) Type() string {
	return "aws_lex_intent"
}

func (li *LexIntent) LocalName() string {
	return li.Name
}

func (li *LexIntent) Configuration() interface{} {
	return li.Args
}

func (li *LexIntent) Attributes() lexIntentAttributes {
	return lexIntentAttributes{ref: terra.ReferenceResource(li)}
}

func (li *LexIntent) ImportState(av io.Reader) error {
	li.state = &lexIntentState{}
	if err := json.NewDecoder(av).Decode(li.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", li.Type(), li.LocalName(), err)
	}
	return nil
}

func (li *LexIntent) State() (*lexIntentState, bool) {
	return li.state, li.state != nil
}

func (li *LexIntent) StateMust() *lexIntentState {
	if li.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", li.Type(), li.LocalName()))
	}
	return li.state
}

func (li *LexIntent) DependOn() terra.Reference {
	return terra.ReferenceResource(li)
}

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
	// DependsOn contains resources that LexIntent depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type lexIntentAttributes struct {
	ref terra.Reference
}

func (li lexIntentAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(li.ref.Append("arn"))
}

func (li lexIntentAttributes) Checksum() terra.StringValue {
	return terra.ReferenceString(li.ref.Append("checksum"))
}

func (li lexIntentAttributes) CreateVersion() terra.BoolValue {
	return terra.ReferenceBool(li.ref.Append("create_version"))
}

func (li lexIntentAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceString(li.ref.Append("created_date"))
}

func (li lexIntentAttributes) Description() terra.StringValue {
	return terra.ReferenceString(li.ref.Append("description"))
}

func (li lexIntentAttributes) Id() terra.StringValue {
	return terra.ReferenceString(li.ref.Append("id"))
}

func (li lexIntentAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceString(li.ref.Append("last_updated_date"))
}

func (li lexIntentAttributes) Name() terra.StringValue {
	return terra.ReferenceString(li.ref.Append("name"))
}

func (li lexIntentAttributes) ParentIntentSignature() terra.StringValue {
	return terra.ReferenceString(li.ref.Append("parent_intent_signature"))
}

func (li lexIntentAttributes) SampleUtterances() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](li.ref.Append("sample_utterances"))
}

func (li lexIntentAttributes) Version() terra.StringValue {
	return terra.ReferenceString(li.ref.Append("version"))
}

func (li lexIntentAttributes) ConclusionStatement() terra.ListValue[lexintent.ConclusionStatementAttributes] {
	return terra.ReferenceList[lexintent.ConclusionStatementAttributes](li.ref.Append("conclusion_statement"))
}

func (li lexIntentAttributes) ConfirmationPrompt() terra.ListValue[lexintent.ConfirmationPromptAttributes] {
	return terra.ReferenceList[lexintent.ConfirmationPromptAttributes](li.ref.Append("confirmation_prompt"))
}

func (li lexIntentAttributes) DialogCodeHook() terra.ListValue[lexintent.DialogCodeHookAttributes] {
	return terra.ReferenceList[lexintent.DialogCodeHookAttributes](li.ref.Append("dialog_code_hook"))
}

func (li lexIntentAttributes) FollowUpPrompt() terra.ListValue[lexintent.FollowUpPromptAttributes] {
	return terra.ReferenceList[lexintent.FollowUpPromptAttributes](li.ref.Append("follow_up_prompt"))
}

func (li lexIntentAttributes) FulfillmentActivity() terra.ListValue[lexintent.FulfillmentActivityAttributes] {
	return terra.ReferenceList[lexintent.FulfillmentActivityAttributes](li.ref.Append("fulfillment_activity"))
}

func (li lexIntentAttributes) RejectionStatement() terra.ListValue[lexintent.RejectionStatementAttributes] {
	return terra.ReferenceList[lexintent.RejectionStatementAttributes](li.ref.Append("rejection_statement"))
}

func (li lexIntentAttributes) Slot() terra.SetValue[lexintent.SlotAttributes] {
	return terra.ReferenceSet[lexintent.SlotAttributes](li.ref.Append("slot"))
}

func (li lexIntentAttributes) Timeouts() lexintent.TimeoutsAttributes {
	return terra.ReferenceSingle[lexintent.TimeoutsAttributes](li.ref.Append("timeouts"))
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
