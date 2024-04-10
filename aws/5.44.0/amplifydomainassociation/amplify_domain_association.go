// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package amplifydomainassociation

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type SubDomain struct {
	// BranchName: string, required
	BranchName terra.StringValue `hcl:"branch_name,attr" validate:"required"`
	// Prefix: string, required
	Prefix terra.StringValue `hcl:"prefix,attr" validate:"required"`
}

type SubDomainAttributes struct {
	ref terra.Reference
}

func (sd SubDomainAttributes) InternalRef() (terra.Reference, error) {
	return sd.ref, nil
}

func (sd SubDomainAttributes) InternalWithRef(ref terra.Reference) SubDomainAttributes {
	return SubDomainAttributes{ref: ref}
}

func (sd SubDomainAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sd.ref.InternalTokens()
}

func (sd SubDomainAttributes) BranchName() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("branch_name"))
}

func (sd SubDomainAttributes) DnsRecord() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("dns_record"))
}

func (sd SubDomainAttributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("prefix"))
}

func (sd SubDomainAttributes) Verified() terra.BoolValue {
	return terra.ReferenceAsBool(sd.ref.Append("verified"))
}

type SubDomainState struct {
	BranchName string `json:"branch_name"`
	DnsRecord  string `json:"dns_record"`
	Prefix     string `json:"prefix"`
	Verified   bool   `json:"verified"`
}
