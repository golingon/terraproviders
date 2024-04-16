// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_connect_bot_association

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataLexBot struct {
	// LexRegion: string, optional
	LexRegion terra.StringValue `hcl:"lex_region,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type DataLexBotAttributes struct {
	ref terra.Reference
}

func (lb DataLexBotAttributes) InternalRef() (terra.Reference, error) {
	return lb.ref, nil
}

func (lb DataLexBotAttributes) InternalWithRef(ref terra.Reference) DataLexBotAttributes {
	return DataLexBotAttributes{ref: ref}
}

func (lb DataLexBotAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lb.ref.InternalTokens()
}

func (lb DataLexBotAttributes) LexRegion() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("lex_region"))
}

func (lb DataLexBotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("name"))
}

type DataLexBotState struct {
	LexRegion string `json:"lex_region"`
	Name      string `json:"name"`
}
