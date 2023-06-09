// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataLexBotAlias creates a new instance of [DataLexBotAlias].
func NewDataLexBotAlias(name string, args DataLexBotAliasArgs) *DataLexBotAlias {
	return &DataLexBotAlias{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLexBotAlias)(nil)

// DataLexBotAlias represents the Terraform data resource aws_lex_bot_alias.
type DataLexBotAlias struct {
	Name string
	Args DataLexBotAliasArgs
}

// DataSource returns the Terraform object type for [DataLexBotAlias].
func (lba *DataLexBotAlias) DataSource() string {
	return "aws_lex_bot_alias"
}

// LocalName returns the local name for [DataLexBotAlias].
func (lba *DataLexBotAlias) LocalName() string {
	return lba.Name
}

// Configuration returns the configuration (args) for [DataLexBotAlias].
func (lba *DataLexBotAlias) Configuration() interface{} {
	return lba.Args
}

// Attributes returns the attributes for [DataLexBotAlias].
func (lba *DataLexBotAlias) Attributes() dataLexBotAliasAttributes {
	return dataLexBotAliasAttributes{ref: terra.ReferenceDataResource(lba)}
}

// DataLexBotAliasArgs contains the configurations for aws_lex_bot_alias.
type DataLexBotAliasArgs struct {
	// BotName: string, required
	BotName terra.StringValue `hcl:"bot_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type dataLexBotAliasAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lex_bot_alias.
func (lba dataLexBotAliasAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("arn"))
}

// BotName returns a reference to field bot_name of aws_lex_bot_alias.
func (lba dataLexBotAliasAttributes) BotName() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("bot_name"))
}

// BotVersion returns a reference to field bot_version of aws_lex_bot_alias.
func (lba dataLexBotAliasAttributes) BotVersion() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("bot_version"))
}

// Checksum returns a reference to field checksum of aws_lex_bot_alias.
func (lba dataLexBotAliasAttributes) Checksum() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("checksum"))
}

// CreatedDate returns a reference to field created_date of aws_lex_bot_alias.
func (lba dataLexBotAliasAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("created_date"))
}

// Description returns a reference to field description of aws_lex_bot_alias.
func (lba dataLexBotAliasAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("description"))
}

// Id returns a reference to field id of aws_lex_bot_alias.
func (lba dataLexBotAliasAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_lex_bot_alias.
func (lba dataLexBotAliasAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("last_updated_date"))
}

// Name returns a reference to field name of aws_lex_bot_alias.
func (lba dataLexBotAliasAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lba.ref.Append("name"))
}
