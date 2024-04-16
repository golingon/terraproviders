// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_dev_test_windows_virtual_machine

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type GalleryImageReference struct {
	// Offer: string, required
	Offer terra.StringValue `hcl:"offer,attr" validate:"required"`
	// Publisher: string, required
	Publisher terra.StringValue `hcl:"publisher,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
}

type InboundNatRule struct {
	// BackendPort: number, required
	BackendPort terra.NumberValue `hcl:"backend_port,attr" validate:"required"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
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

type GalleryImageReferenceAttributes struct {
	ref terra.Reference
}

func (gir GalleryImageReferenceAttributes) InternalRef() (terra.Reference, error) {
	return gir.ref, nil
}

func (gir GalleryImageReferenceAttributes) InternalWithRef(ref terra.Reference) GalleryImageReferenceAttributes {
	return GalleryImageReferenceAttributes{ref: ref}
}

func (gir GalleryImageReferenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gir.ref.InternalTokens()
}

func (gir GalleryImageReferenceAttributes) Offer() terra.StringValue {
	return terra.ReferenceAsString(gir.ref.Append("offer"))
}

func (gir GalleryImageReferenceAttributes) Publisher() terra.StringValue {
	return terra.ReferenceAsString(gir.ref.Append("publisher"))
}

func (gir GalleryImageReferenceAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(gir.ref.Append("sku"))
}

func (gir GalleryImageReferenceAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(gir.ref.Append("version"))
}

type InboundNatRuleAttributes struct {
	ref terra.Reference
}

func (inr InboundNatRuleAttributes) InternalRef() (terra.Reference, error) {
	return inr.ref, nil
}

func (inr InboundNatRuleAttributes) InternalWithRef(ref terra.Reference) InboundNatRuleAttributes {
	return InboundNatRuleAttributes{ref: ref}
}

func (inr InboundNatRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return inr.ref.InternalTokens()
}

func (inr InboundNatRuleAttributes) BackendPort() terra.NumberValue {
	return terra.ReferenceAsNumber(inr.ref.Append("backend_port"))
}

func (inr InboundNatRuleAttributes) FrontendPort() terra.NumberValue {
	return terra.ReferenceAsNumber(inr.ref.Append("frontend_port"))
}

func (inr InboundNatRuleAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(inr.ref.Append("protocol"))
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

type GalleryImageReferenceState struct {
	Offer     string `json:"offer"`
	Publisher string `json:"publisher"`
	Sku       string `json:"sku"`
	Version   string `json:"version"`
}

type InboundNatRuleState struct {
	BackendPort  float64 `json:"backend_port"`
	FrontendPort float64 `json:"frontend_port"`
	Protocol     string  `json:"protocol"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
