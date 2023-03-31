// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package loganalyticssolution

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Plan struct {
	// Product: string, required
	Product terra.StringValue `hcl:"product,attr" validate:"required"`
	// PromotionCode: string, optional
	PromotionCode terra.StringValue `hcl:"promotion_code,attr"`
	// Publisher: string, required
	Publisher terra.StringValue `hcl:"publisher,attr" validate:"required"`
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

type PlanAttributes struct {
	ref terra.Reference
}

func (p PlanAttributes) InternalRef() terra.Reference {
	return p.ref
}

func (p PlanAttributes) InternalWithRef(ref terra.Reference) PlanAttributes {
	return PlanAttributes{ref: ref}
}

func (p PlanAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p PlanAttributes) Name() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("name"))
}

func (p PlanAttributes) Product() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("product"))
}

func (p PlanAttributes) PromotionCode() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("promotion_code"))
}

func (p PlanAttributes) Publisher() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("publisher"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type PlanState struct {
	Name          string `json:"name"`
	Product       string `json:"product"`
	PromotionCode string `json:"promotion_code"`
	Publisher     string `json:"publisher"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
