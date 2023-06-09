// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package opensearchserverlesssecurityconfig

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SamlOptions struct {
	// GroupAttribute: string, optional
	GroupAttribute terra.StringValue `hcl:"group_attribute,attr"`
	// Metadata: string, required
	Metadata terra.StringValue `hcl:"metadata,attr" validate:"required"`
	// SessionTimeout: number, optional
	SessionTimeout terra.NumberValue `hcl:"session_timeout,attr"`
	// UserAttribute: string, optional
	UserAttribute terra.StringValue `hcl:"user_attribute,attr"`
}

type SamlOptionsAttributes struct {
	ref terra.Reference
}

func (so SamlOptionsAttributes) InternalRef() (terra.Reference, error) {
	return so.ref, nil
}

func (so SamlOptionsAttributes) InternalWithRef(ref terra.Reference) SamlOptionsAttributes {
	return SamlOptionsAttributes{ref: ref}
}

func (so SamlOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return so.ref.InternalTokens()
}

func (so SamlOptionsAttributes) GroupAttribute() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("group_attribute"))
}

func (so SamlOptionsAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("metadata"))
}

func (so SamlOptionsAttributes) SessionTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(so.ref.Append("session_timeout"))
}

func (so SamlOptionsAttributes) UserAttribute() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("user_attribute"))
}

type SamlOptionsState struct {
	GroupAttribute string  `json:"group_attribute"`
	Metadata       string  `json:"metadata"`
	SessionTimeout float64 `json:"session_timeout"`
	UserAttribute  string  `json:"user_attribute"`
}
