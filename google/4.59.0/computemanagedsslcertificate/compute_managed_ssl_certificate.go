// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package computemanagedsslcertificate

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Managed struct {
	// Domains: list of string, required
	Domains terra.ListValue[terra.StringValue] `hcl:"domains,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type ManagedAttributes struct {
	ref terra.Reference
}

func (m ManagedAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m ManagedAttributes) InternalWithRef(ref terra.Reference) ManagedAttributes {
	return ManagedAttributes{ref: ref}
}

func (m ManagedAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m ManagedAttributes) Domains() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](m.ref.Append("domains"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

type ManagedState struct {
	Domains []string `json:"domains"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
