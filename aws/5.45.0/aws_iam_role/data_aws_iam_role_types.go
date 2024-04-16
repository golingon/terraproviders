// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_iam_role

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataRoleLastUsedAttributes struct {
	ref terra.Reference
}

func (rlu DataRoleLastUsedAttributes) InternalRef() (terra.Reference, error) {
	return rlu.ref, nil
}

func (rlu DataRoleLastUsedAttributes) InternalWithRef(ref terra.Reference) DataRoleLastUsedAttributes {
	return DataRoleLastUsedAttributes{ref: ref}
}

func (rlu DataRoleLastUsedAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rlu.ref.InternalTokens()
}

func (rlu DataRoleLastUsedAttributes) LastUsedDate() terra.StringValue {
	return terra.ReferenceAsString(rlu.ref.Append("last_used_date"))
}

func (rlu DataRoleLastUsedAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(rlu.ref.Append("region"))
}

type DataRoleLastUsedState struct {
	LastUsedDate string `json:"last_used_date"`
	Region       string `json:"region"`
}
