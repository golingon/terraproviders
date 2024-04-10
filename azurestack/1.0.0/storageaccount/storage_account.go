// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package storageaccount

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CustomDomain struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// UseSubdomain: bool, optional
	UseSubdomain terra.BoolValue `hcl:"use_subdomain,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type CustomDomainAttributes struct {
	ref terra.Reference
}

func (cd CustomDomainAttributes) InternalRef() (terra.Reference, error) {
	return cd.ref, nil
}

func (cd CustomDomainAttributes) InternalWithRef(ref terra.Reference) CustomDomainAttributes {
	return CustomDomainAttributes{ref: ref}
}

func (cd CustomDomainAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cd.ref.InternalTokens()
}

func (cd CustomDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("name"))
}

func (cd CustomDomainAttributes) UseSubdomain() terra.BoolValue {
	return terra.ReferenceAsBool(cd.ref.Append("use_subdomain"))
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

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type CustomDomainState struct {
	Name         string `json:"name"`
	UseSubdomain bool   `json:"use_subdomain"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
