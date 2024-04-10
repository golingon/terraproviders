// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datasharedimage

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Identifier struct{}

type PurchasePlan struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type IdentifierAttributes struct {
	ref terra.Reference
}

func (i IdentifierAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentifierAttributes) InternalWithRef(ref terra.Reference) IdentifierAttributes {
	return IdentifierAttributes{ref: ref}
}

func (i IdentifierAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentifierAttributes) Offer() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("offer"))
}

func (i IdentifierAttributes) Publisher() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("publisher"))
}

func (i IdentifierAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("sku"))
}

type PurchasePlanAttributes struct {
	ref terra.Reference
}

func (pp PurchasePlanAttributes) InternalRef() (terra.Reference, error) {
	return pp.ref, nil
}

func (pp PurchasePlanAttributes) InternalWithRef(ref terra.Reference) PurchasePlanAttributes {
	return PurchasePlanAttributes{ref: ref}
}

func (pp PurchasePlanAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pp.ref.InternalTokens()
}

func (pp PurchasePlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pp.ref.Append("name"))
}

func (pp PurchasePlanAttributes) Product() terra.StringValue {
	return terra.ReferenceAsString(pp.ref.Append("product"))
}

func (pp PurchasePlanAttributes) Publisher() terra.StringValue {
	return terra.ReferenceAsString(pp.ref.Append("publisher"))
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

type IdentifierState struct {
	Offer     string `json:"offer"`
	Publisher string `json:"publisher"`
	Sku       string `json:"sku"`
}

type PurchasePlanState struct {
	Name      string `json:"name"`
	Product   string `json:"product"`
	Publisher string `json:"publisher"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
