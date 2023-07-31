// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datanamedlocation

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Country struct{}

type Ip struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type CountryAttributes struct {
	ref terra.Reference
}

func (c CountryAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CountryAttributes) InternalWithRef(ref terra.Reference) CountryAttributes {
	return CountryAttributes{ref: ref}
}

func (c CountryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CountryAttributes) CountriesAndRegions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("countries_and_regions"))
}

func (c CountryAttributes) IncludeUnknownCountriesAndRegions() terra.BoolValue {
	return terra.ReferenceAsBool(c.ref.Append("include_unknown_countries_and_regions"))
}

type IpAttributes struct {
	ref terra.Reference
}

func (i IpAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IpAttributes) InternalWithRef(ref terra.Reference) IpAttributes {
	return IpAttributes{ref: ref}
}

func (i IpAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IpAttributes) IpRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("ip_ranges"))
}

func (i IpAttributes) Trusted() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("trusted"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type CountryState struct {
	CountriesAndRegions               []string `json:"countries_and_regions"`
	IncludeUnknownCountriesAndRegions bool     `json:"include_unknown_countries_and_regions"`
}

type IpState struct {
	IpRanges []string `json:"ip_ranges"`
	Trusted  bool     `json:"trusted"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
