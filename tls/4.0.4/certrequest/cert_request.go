// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package certrequest

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Subject struct {
	// CommonName: string, optional
	CommonName terra.StringValue `hcl:"common_name,attr"`
	// Country: string, optional
	Country terra.StringValue `hcl:"country,attr"`
	// Locality: string, optional
	Locality terra.StringValue `hcl:"locality,attr"`
	// Organization: string, optional
	Organization terra.StringValue `hcl:"organization,attr"`
	// OrganizationalUnit: string, optional
	OrganizationalUnit terra.StringValue `hcl:"organizational_unit,attr"`
	// PostalCode: string, optional
	PostalCode terra.StringValue `hcl:"postal_code,attr"`
	// Province: string, optional
	Province terra.StringValue `hcl:"province,attr"`
	// SerialNumber: string, optional
	SerialNumber terra.StringValue `hcl:"serial_number,attr"`
	// StreetAddress: list of string, optional
	StreetAddress terra.ListValue[terra.StringValue] `hcl:"street_address,attr"`
}

type SubjectAttributes struct {
	ref terra.Reference
}

func (s SubjectAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SubjectAttributes) InternalWithRef(ref terra.Reference) SubjectAttributes {
	return SubjectAttributes{ref: ref}
}

func (s SubjectAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SubjectAttributes) CommonName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("common_name"))
}

func (s SubjectAttributes) Country() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("country"))
}

func (s SubjectAttributes) Locality() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("locality"))
}

func (s SubjectAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("organization"))
}

func (s SubjectAttributes) OrganizationalUnit() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("organizational_unit"))
}

func (s SubjectAttributes) PostalCode() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("postal_code"))
}

func (s SubjectAttributes) Province() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("province"))
}

func (s SubjectAttributes) SerialNumber() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("serial_number"))
}

func (s SubjectAttributes) StreetAddress() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("street_address"))
}

type SubjectState struct {
	CommonName         string   `json:"common_name"`
	Country            string   `json:"country"`
	Locality           string   `json:"locality"`
	Organization       string   `json:"organization"`
	OrganizationalUnit string   `json:"organizational_unit"`
	PostalCode         string   `json:"postal_code"`
	Province           string   `json:"province"`
	SerialNumber       string   `json:"serial_number"`
	StreetAddress      []string `json:"street_address"`
}
