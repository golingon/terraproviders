// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataLexBotAlias(name string, args DataLexBotAliasArgs) *DataLexBotAlias {
	return &DataLexBotAlias{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLexBotAlias)(nil)

type DataLexBotAlias struct {
	Name string
	Args DataLexBotAliasArgs
}

func (lba *DataLexBotAlias) DataSource() string {
	return "aws_lex_bot_alias"
}

func (lba *DataLexBotAlias) LocalName() string {
	return lba.Name
}

func (lba *DataLexBotAlias) Configuration() interface{} {
	return lba.Args
}

func (lba *DataLexBotAlias) Attributes() dataLexBotAliasAttributes {
	return dataLexBotAliasAttributes{ref: terra.ReferenceDataResource(lba)}
}

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

func (lba dataLexBotAliasAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(lba.ref.Append("arn"))
}

func (lba dataLexBotAliasAttributes) BotName() terra.StringValue {
	return terra.ReferenceString(lba.ref.Append("bot_name"))
}

func (lba dataLexBotAliasAttributes) BotVersion() terra.StringValue {
	return terra.ReferenceString(lba.ref.Append("bot_version"))
}

func (lba dataLexBotAliasAttributes) Checksum() terra.StringValue {
	return terra.ReferenceString(lba.ref.Append("checksum"))
}

func (lba dataLexBotAliasAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceString(lba.ref.Append("created_date"))
}

func (lba dataLexBotAliasAttributes) Description() terra.StringValue {
	return terra.ReferenceString(lba.ref.Append("description"))
}

func (lba dataLexBotAliasAttributes) Id() terra.StringValue {
	return terra.ReferenceString(lba.ref.Append("id"))
}

func (lba dataLexBotAliasAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceString(lba.ref.Append("last_updated_date"))
}

func (lba dataLexBotAliasAttributes) Name() terra.StringValue {
	return terra.ReferenceString(lba.ref.Append("name"))
}
