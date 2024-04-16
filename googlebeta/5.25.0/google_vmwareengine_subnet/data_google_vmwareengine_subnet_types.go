// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_vmwareengine_subnet

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataDhcpAddressRangesAttributes struct {
	ref terra.Reference
}

func (dar DataDhcpAddressRangesAttributes) InternalRef() (terra.Reference, error) {
	return dar.ref, nil
}

func (dar DataDhcpAddressRangesAttributes) InternalWithRef(ref terra.Reference) DataDhcpAddressRangesAttributes {
	return DataDhcpAddressRangesAttributes{ref: ref}
}

func (dar DataDhcpAddressRangesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dar.ref.InternalTokens()
}

func (dar DataDhcpAddressRangesAttributes) FirstAddress() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("first_address"))
}

func (dar DataDhcpAddressRangesAttributes) LastAddress() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("last_address"))
}

type DataDhcpAddressRangesState struct {
	FirstAddress string `json:"first_address"`
	LastAddress  string `json:"last_address"`
}
