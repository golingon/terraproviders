// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_sql_managed_instance_failover_group

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ReadWriteEndpointFailoverPolicy struct {
	// GraceMinutes: number, optional
	GraceMinutes terra.NumberValue `hcl:"grace_minutes,attr"`
	// Mode: string, required
	Mode terra.StringValue `hcl:"mode,attr" validate:"required"`
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

type PartnerRegionAttributes struct {
	ref terra.Reference
}

func (pr PartnerRegionAttributes) InternalRef() (terra.Reference, error) {
	return pr.ref, nil
}

func (pr PartnerRegionAttributes) InternalWithRef(ref terra.Reference) PartnerRegionAttributes {
	return PartnerRegionAttributes{ref: ref}
}

func (pr PartnerRegionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pr.ref.InternalTokens()
}

func (pr PartnerRegionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pr.ref.Append("location"))
}

func (pr PartnerRegionAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(pr.ref.Append("role"))
}

type ReadWriteEndpointFailoverPolicyAttributes struct {
	ref terra.Reference
}

func (rwefp ReadWriteEndpointFailoverPolicyAttributes) InternalRef() (terra.Reference, error) {
	return rwefp.ref, nil
}

func (rwefp ReadWriteEndpointFailoverPolicyAttributes) InternalWithRef(ref terra.Reference) ReadWriteEndpointFailoverPolicyAttributes {
	return ReadWriteEndpointFailoverPolicyAttributes{ref: ref}
}

func (rwefp ReadWriteEndpointFailoverPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rwefp.ref.InternalTokens()
}

func (rwefp ReadWriteEndpointFailoverPolicyAttributes) GraceMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(rwefp.ref.Append("grace_minutes"))
}

func (rwefp ReadWriteEndpointFailoverPolicyAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(rwefp.ref.Append("mode"))
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

type PartnerRegionState struct {
	Location string `json:"location"`
	Role     string `json:"role"`
}

type ReadWriteEndpointFailoverPolicyState struct {
	GraceMinutes float64 `json:"grace_minutes"`
	Mode         string  `json:"mode"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
