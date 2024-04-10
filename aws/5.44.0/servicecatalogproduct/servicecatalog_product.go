// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package servicecatalogproduct

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ProvisioningArtifactParameters struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisableTemplateValidation: bool, optional
	DisableTemplateValidation terra.BoolValue `hcl:"disable_template_validation,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// TemplatePhysicalId: string, optional
	TemplatePhysicalId terra.StringValue `hcl:"template_physical_id,attr"`
	// TemplateUrl: string, optional
	TemplateUrl terra.StringValue `hcl:"template_url,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
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

type ProvisioningArtifactParametersAttributes struct {
	ref terra.Reference
}

func (pap ProvisioningArtifactParametersAttributes) InternalRef() (terra.Reference, error) {
	return pap.ref, nil
}

func (pap ProvisioningArtifactParametersAttributes) InternalWithRef(ref terra.Reference) ProvisioningArtifactParametersAttributes {
	return ProvisioningArtifactParametersAttributes{ref: ref}
}

func (pap ProvisioningArtifactParametersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pap.ref.InternalTokens()
}

func (pap ProvisioningArtifactParametersAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(pap.ref.Append("description"))
}

func (pap ProvisioningArtifactParametersAttributes) DisableTemplateValidation() terra.BoolValue {
	return terra.ReferenceAsBool(pap.ref.Append("disable_template_validation"))
}

func (pap ProvisioningArtifactParametersAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pap.ref.Append("name"))
}

func (pap ProvisioningArtifactParametersAttributes) TemplatePhysicalId() terra.StringValue {
	return terra.ReferenceAsString(pap.ref.Append("template_physical_id"))
}

func (pap ProvisioningArtifactParametersAttributes) TemplateUrl() terra.StringValue {
	return terra.ReferenceAsString(pap.ref.Append("template_url"))
}

func (pap ProvisioningArtifactParametersAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(pap.ref.Append("type"))
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

type ProvisioningArtifactParametersState struct {
	Description               string `json:"description"`
	DisableTemplateValidation bool   `json:"disable_template_validation"`
	Name                      string `json:"name"`
	TemplatePhysicalId        string `json:"template_physical_id"`
	TemplateUrl               string `json:"template_url"`
	Type                      string `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
