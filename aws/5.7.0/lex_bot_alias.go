// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lexbotalias "github.com/golingon/terraproviders/aws/5.7.0/lexbotalias"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLexBotAlias creates a new instance of [LexBotAlias].
func NewLexBotAlias(name string, args LexBotAliasArgs) *LexBotAlias {
	return &LexBotAlias{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LexBotAlias)(nil)

// LexBotAlias represents the Terraform resource aws_lex_bot_alias.
type LexBotAlias struct {
	Name      string
	Args      LexBotAliasArgs
	state     *lexBotAliasState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LexBotAlias].
func (lba *LexBotAlias) Type() string {
	return "aws_lex_bot_alias"
}

// LocalName returns the local name for [LexBotAlias].
func (lba *LexBotAlias) LocalName() string {
	return lba.Name
}

// Configuration returns the configuration (args) for [LexBotAlias].
func (lba *LexBotAlias) Configuration() interface{} {
	return lba.Args
}

// DependOn is used for other resources to depend on [LexBotAlias].
func (lba *LexBotAlias) DependOn() terra.Reference {
	return terra.ReferenceResource(lba)
}

// Dependencies returns the list of resources [LexBotAlias] depends_on.
func (lba *LexBotAlias) Dependencies() terra.Dependencies {
	return lba.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LexBotAlias].
func (lba *LexBotAlias) LifecycleManagement() *terra.Lifecycle {
	return lba.Lifecycle
}

// Attributes returns the attributes for [LexBotAlias].
func (lba *LexBotAlias) Attributes() lexBotAliasAttributes {
	return lexBotAliasAttributes{ref: terra.ReferenceResource(lba)}
}

// ImportState imports the given attribute values into [LexBotAlias]'s state.
func (lba *LexBotAlias) ImportState(av io.Reader) error {
	lba.state = &lexBotAliasState{}
	if err := json.NewDecoder(av).Decode(lba.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lba.Type(), lba.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LexBotAlias] has state.
func (lba *LexBotAlias) State() (*lexBotAliasState, bool) {
	return lba.state, lba.state != nil
}

// StateMust returns the state for [LexBotAlias]. Panics if the state is nil.
func (lba *LexBotAlias) StateMust() *lexBotAliasState {
	if lba.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lba.Type(), lba.LocalName()))
	}
	return lba.state
}

// LexBotAliasArgs contains the configurations for aws_lex_bot_alias.
type LexBotAliasArgs struct {
	// BotName: string, required
	BotName terra.StringValue `hcl:"bot_name,attr" validate:"required"`
	// BotVersion: string, required
	BotVersion terra.StringValue `hcl:"bot_version,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ConversationLogs: optional
	ConversationLogs *lexbotalias.ConversationLogs `hcl:"conversation_logs,block"`
	// Timeouts: optional
	Timeouts *lexbotalias.Timeouts `hcl:"timeouts,block"`
}
type lexBotAliasAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lex_bot_alias.
func (lba lexBotAliasAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("arn"))
}

// BotName returns a reference to field bot_name of aws_lex_bot_alias.
func (lba lexBotAliasAttributes) BotName() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("bot_name"))
}

// BotVersion returns a reference to field bot_version of aws_lex_bot_alias.
func (lba lexBotAliasAttributes) BotVersion() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("bot_version"))
}

// Checksum returns a reference to field checksum of aws_lex_bot_alias.
func (lba lexBotAliasAttributes) Checksum() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("checksum"))
}

// CreatedDate returns a reference to field created_date of aws_lex_bot_alias.
func (lba lexBotAliasAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("created_date"))
}

// Description returns a reference to field description of aws_lex_bot_alias.
func (lba lexBotAliasAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("description"))
}

// Id returns a reference to field id of aws_lex_bot_alias.
func (lba lexBotAliasAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_lex_bot_alias.
func (lba lexBotAliasAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("last_updated_date"))
}

// Name returns a reference to field name of aws_lex_bot_alias.
func (lba lexBotAliasAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("name"))
}

func (lba lexBotAliasAttributes) ConversationLogs() terra.ListValue[lexbotalias.ConversationLogsAttributes] {
	return terra.ReferenceAsList[lexbotalias.ConversationLogsAttributes](lba.ref.Append("conversation_logs"))
}

func (lba lexBotAliasAttributes) Timeouts() lexbotalias.TimeoutsAttributes {
	return terra.ReferenceAsSingle[lexbotalias.TimeoutsAttributes](lba.ref.Append("timeouts"))
}

type lexBotAliasState struct {
	Arn              string                              `json:"arn"`
	BotName          string                              `json:"bot_name"`
	BotVersion       string                              `json:"bot_version"`
	Checksum         string                              `json:"checksum"`
	CreatedDate      string                              `json:"created_date"`
	Description      string                              `json:"description"`
	Id               string                              `json:"id"`
	LastUpdatedDate  string                              `json:"last_updated_date"`
	Name             string                              `json:"name"`
	ConversationLogs []lexbotalias.ConversationLogsState `json:"conversation_logs"`
	Timeouts         *lexbotalias.TimeoutsState          `json:"timeouts"`
}
