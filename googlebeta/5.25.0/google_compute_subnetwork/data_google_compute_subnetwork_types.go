// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_subnetwork

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataSecondaryIpRangeAttributes struct {
	ref terra.Reference
}

func (sir DataSecondaryIpRangeAttributes) InternalRef() (terra.Reference, error) {
	return sir.ref, nil
}

func (sir DataSecondaryIpRangeAttributes) InternalWithRef(ref terra.Reference) DataSecondaryIpRangeAttributes {
	return DataSecondaryIpRangeAttributes{ref: ref}
}

func (sir DataSecondaryIpRangeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sir.ref.InternalTokens()
}

func (sir DataSecondaryIpRangeAttributes) IpCidrRange() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("ip_cidr_range"))
}

func (sir DataSecondaryIpRangeAttributes) RangeName() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("range_name"))
}

type DataSecondaryIpRangeState struct {
	IpCidrRange string `json:"ip_cidr_range"`
	RangeName   string `json:"range_name"`
}
