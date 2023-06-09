// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataiamrole

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RoleLastUsed struct{}

type RoleLastUsedAttributes struct {
	ref terra.Reference
}

func (rlu RoleLastUsedAttributes) InternalRef() (terra.Reference, error) {
	return rlu.ref, nil
}

func (rlu RoleLastUsedAttributes) InternalWithRef(ref terra.Reference) RoleLastUsedAttributes {
	return RoleLastUsedAttributes{ref: ref}
}

func (rlu RoleLastUsedAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rlu.ref.InternalTokens()
}

func (rlu RoleLastUsedAttributes) LastUsedDate() terra.StringValue {
	return terra.ReferenceAsString(rlu.ref.Append("last_used_date"))
}

func (rlu RoleLastUsedAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(rlu.ref.Append("region"))
}

type RoleLastUsedState struct {
	LastUsedDate string `json:"last_used_date"`
	Region       string `json:"region"`
}
