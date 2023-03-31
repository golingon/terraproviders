// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ec2managedprefixlist

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Entry struct {
	// Cidr: string, required
	Cidr terra.StringValue `hcl:"cidr,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
}

type EntryAttributes struct {
	ref terra.Reference
}

func (e EntryAttributes) InternalRef() terra.Reference {
	return e.ref
}

func (e EntryAttributes) InternalWithRef(ref terra.Reference) EntryAttributes {
	return EntryAttributes{ref: ref}
}

func (e EntryAttributes) InternalTokens() hclwrite.Tokens {
	return e.ref.InternalTokens()
}

func (e EntryAttributes) Cidr() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("cidr"))
}

func (e EntryAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("description"))
}

type EntryState struct {
	Cidr        string `json:"cidr"`
	Description string `json:"description"`
}