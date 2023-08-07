// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataLexBot creates a new instance of [DataLexBot].
func NewDataLexBot(name string, args DataLexBotArgs) *DataLexBot {
	return &DataLexBot{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLexBot)(nil)

// DataLexBot represents the Terraform data resource aws_lex_bot.
type DataLexBot struct {
	Name string
	Args DataLexBotArgs
}

// DataSource returns the Terraform object type for [DataLexBot].
func (lb *DataLexBot) DataSource() string {
	return "aws_lex_bot"
}

// LocalName returns the local name for [DataLexBot].
func (lb *DataLexBot) LocalName() string {
	return lb.Name
}

// Configuration returns the configuration (args) for [DataLexBot].
func (lb *DataLexBot) Configuration() interface{} {
	return lb.Args
}

// Attributes returns the attributes for [DataLexBot].
func (lb *DataLexBot) Attributes() dataLexBotAttributes {
	return dataLexBotAttributes{ref: terra.ReferenceDataResource(lb)}
}

// DataLexBotArgs contains the configurations for aws_lex_bot.
type DataLexBotArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
}
type dataLexBotAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lex_bot.
func (lb dataLexBotAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("arn"))
}

// Checksum returns a reference to field checksum of aws_lex_bot.
func (lb dataLexBotAttributes) Checksum() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("checksum"))
}

// ChildDirected returns a reference to field child_directed of aws_lex_bot.
func (lb dataLexBotAttributes) ChildDirected() terra.BoolValue {
	return terra.ReferenceAsBool(lb.ref.Append("child_directed"))
}

// CreatedDate returns a reference to field created_date of aws_lex_bot.
func (lb dataLexBotAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("created_date"))
}

// Description returns a reference to field description of aws_lex_bot.
func (lb dataLexBotAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("description"))
}

// DetectSentiment returns a reference to field detect_sentiment of aws_lex_bot.
func (lb dataLexBotAttributes) DetectSentiment() terra.BoolValue {
	return terra.ReferenceAsBool(lb.ref.Append("detect_sentiment"))
}

// EnableModelImprovements returns a reference to field enable_model_improvements of aws_lex_bot.
func (lb dataLexBotAttributes) EnableModelImprovements() terra.BoolValue {
	return terra.ReferenceAsBool(lb.ref.Append("enable_model_improvements"))
}

// FailureReason returns a reference to field failure_reason of aws_lex_bot.
func (lb dataLexBotAttributes) FailureReason() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("failure_reason"))
}

// Id returns a reference to field id of aws_lex_bot.
func (lb dataLexBotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("id"))
}

// IdleSessionTtlInSeconds returns a reference to field idle_session_ttl_in_seconds of aws_lex_bot.
func (lb dataLexBotAttributes) IdleSessionTtlInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(lb.ref.Append("idle_session_ttl_in_seconds"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_lex_bot.
func (lb dataLexBotAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("last_updated_date"))
}

// Locale returns a reference to field locale of aws_lex_bot.
func (lb dataLexBotAttributes) Locale() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("locale"))
}

// Name returns a reference to field name of aws_lex_bot.
func (lb dataLexBotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("name"))
}

// NluIntentConfidenceThreshold returns a reference to field nlu_intent_confidence_threshold of aws_lex_bot.
func (lb dataLexBotAttributes) NluIntentConfidenceThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(lb.ref.Append("nlu_intent_confidence_threshold"))
}

// Status returns a reference to field status of aws_lex_bot.
func (lb dataLexBotAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("status"))
}

// Version returns a reference to field version of aws_lex_bot.
func (lb dataLexBotAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("version"))
}

// VoiceId returns a reference to field voice_id of aws_lex_bot.
func (lb dataLexBotAttributes) VoiceId() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("voice_id"))
}
