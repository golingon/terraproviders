// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package mssqlelasticpool

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PerDatabaseSettings struct {
	// MaxCapacity: number, required
	MaxCapacity terra.NumberValue `hcl:"max_capacity,attr" validate:"required"`
	// MinCapacity: number, required
	MinCapacity terra.NumberValue `hcl:"min_capacity,attr" validate:"required"`
}

type Sku struct {
	// Capacity: number, required
	Capacity terra.NumberValue `hcl:"capacity,attr" validate:"required"`
	// Family: string, optional
	Family terra.StringValue `hcl:"family,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tier: string, required
	Tier terra.StringValue `hcl:"tier,attr" validate:"required"`
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

type PerDatabaseSettingsAttributes struct {
	ref terra.Reference
}

func (pds PerDatabaseSettingsAttributes) InternalRef() (terra.Reference, error) {
	return pds.ref, nil
}

func (pds PerDatabaseSettingsAttributes) InternalWithRef(ref terra.Reference) PerDatabaseSettingsAttributes {
	return PerDatabaseSettingsAttributes{ref: ref}
}

func (pds PerDatabaseSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pds.ref.InternalTokens()
}

func (pds PerDatabaseSettingsAttributes) MaxCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(pds.ref.Append("max_capacity"))
}

func (pds PerDatabaseSettingsAttributes) MinCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(pds.ref.Append("min_capacity"))
}

type SkuAttributes struct {
	ref terra.Reference
}

func (s SkuAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SkuAttributes) InternalWithRef(ref terra.Reference) SkuAttributes {
	return SkuAttributes{ref: ref}
}

func (s SkuAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SkuAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("capacity"))
}

func (s SkuAttributes) Family() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("family"))
}

func (s SkuAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SkuAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("tier"))
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

type PerDatabaseSettingsState struct {
	MaxCapacity float64 `json:"max_capacity"`
	MinCapacity float64 `json:"min_capacity"`
}

type SkuState struct {
	Capacity float64 `json:"capacity"`
	Family   string  `json:"family"`
	Name     string  `json:"name"`
	Tier     string  `json:"tier"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}