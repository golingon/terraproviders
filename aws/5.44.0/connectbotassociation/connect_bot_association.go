// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package connectbotassociation

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type LexBot struct {
	// LexRegion: string, optional
	LexRegion terra.StringValue `hcl:"lex_region,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type LexBotAttributes struct {
	ref terra.Reference
}

func (lb LexBotAttributes) InternalRef() (terra.Reference, error) {
	return lb.ref, nil
}

func (lb LexBotAttributes) InternalWithRef(ref terra.Reference) LexBotAttributes {
	return LexBotAttributes{ref: ref}
}

func (lb LexBotAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lb.ref.InternalTokens()
}

func (lb LexBotAttributes) LexRegion() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("lex_region"))
}

func (lb LexBotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("name"))
}

type LexBotState struct {
	LexRegion string `json:"lex_region"`
	Name      string `json:"name"`
}
