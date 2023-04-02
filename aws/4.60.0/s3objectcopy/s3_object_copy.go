// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package s3objectcopy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Grant struct {
	// Email: string, optional
	Email terra.StringValue `hcl:"email,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Permissions: set of string, required
	Permissions terra.SetValue[terra.StringValue] `hcl:"permissions,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Uri: string, optional
	Uri terra.StringValue `hcl:"uri,attr"`
}

type GrantAttributes struct {
	ref terra.Reference
}

func (g GrantAttributes) InternalRef() (terra.Reference, error) {
	return g.ref, nil
}

func (g GrantAttributes) InternalWithRef(ref terra.Reference) GrantAttributes {
	return GrantAttributes{ref: ref}
}

func (g GrantAttributes) InternalTokens() hclwrite.Tokens {
	return g.ref.InternalTokens()
}

func (g GrantAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("email"))
}

func (g GrantAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("id"))
}

func (g GrantAttributes) Permissions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](g.ref.Append("permissions"))
}

func (g GrantAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("type"))
}

func (g GrantAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("uri"))
}

type GrantState struct {
	Email       string   `json:"email"`
	Id          string   `json:"id"`
	Permissions []string `json:"permissions"`
	Type        string   `json:"type"`
	Uri         string   `json:"uri"`
}
