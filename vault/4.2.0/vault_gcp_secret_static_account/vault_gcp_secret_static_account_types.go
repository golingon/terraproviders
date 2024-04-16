// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_gcp_secret_static_account

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Binding struct {
	// Resource: string, required
	Resource terra.StringValue `hcl:"resource,attr" validate:"required"`
	// Roles: set of string, required
	Roles terra.SetValue[terra.StringValue] `hcl:"roles,attr" validate:"required"`
}

type BindingAttributes struct {
	ref terra.Reference
}

func (b BindingAttributes) InternalRef() (terra.Reference, error) {
	return b.ref, nil
}

func (b BindingAttributes) InternalWithRef(ref terra.Reference) BindingAttributes {
	return BindingAttributes{ref: ref}
}

func (b BindingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return b.ref.InternalTokens()
}

func (b BindingAttributes) Resource() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("resource"))
}

func (b BindingAttributes) Roles() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](b.ref.Append("roles"))
}

type BindingState struct {
	Resource string   `json:"resource"`
	Roles    []string `json:"roles"`
}
