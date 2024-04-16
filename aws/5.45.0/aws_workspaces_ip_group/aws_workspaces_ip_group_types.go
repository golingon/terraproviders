// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_workspaces_ip_group

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Rules struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Source: string, required
	Source terra.StringValue `hcl:"source,attr" validate:"required"`
}

type RulesAttributes struct {
	ref terra.Reference
}

func (r RulesAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RulesAttributes) InternalWithRef(ref terra.Reference) RulesAttributes {
	return RulesAttributes{ref: ref}
}

func (r RulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RulesAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("description"))
}

func (r RulesAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("source"))
}

type RulesState struct {
	Description string `json:"description"`
	Source      string `json:"source"`
}
