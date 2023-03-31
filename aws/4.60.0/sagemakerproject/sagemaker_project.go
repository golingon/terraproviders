// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sagemakerproject

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ServiceCatalogProvisioningDetails struct {
	// PathId: string, optional
	PathId terra.StringValue `hcl:"path_id,attr"`
	// ProductId: string, required
	ProductId terra.StringValue `hcl:"product_id,attr" validate:"required"`
	// ProvisioningArtifactId: string, optional
	ProvisioningArtifactId terra.StringValue `hcl:"provisioning_artifact_id,attr"`
	// ProvisioningParameter: min=0
	ProvisioningParameter []ProvisioningParameter `hcl:"provisioning_parameter,block" validate:"min=0"`
}

type ProvisioningParameter struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type ServiceCatalogProvisioningDetailsAttributes struct {
	ref terra.Reference
}

func (scpd ServiceCatalogProvisioningDetailsAttributes) InternalRef() terra.Reference {
	return scpd.ref
}

func (scpd ServiceCatalogProvisioningDetailsAttributes) InternalWithRef(ref terra.Reference) ServiceCatalogProvisioningDetailsAttributes {
	return ServiceCatalogProvisioningDetailsAttributes{ref: ref}
}

func (scpd ServiceCatalogProvisioningDetailsAttributes) InternalTokens() hclwrite.Tokens {
	return scpd.ref.InternalTokens()
}

func (scpd ServiceCatalogProvisioningDetailsAttributes) PathId() terra.StringValue {
	return terra.ReferenceAsString(scpd.ref.Append("path_id"))
}

func (scpd ServiceCatalogProvisioningDetailsAttributes) ProductId() terra.StringValue {
	return terra.ReferenceAsString(scpd.ref.Append("product_id"))
}

func (scpd ServiceCatalogProvisioningDetailsAttributes) ProvisioningArtifactId() terra.StringValue {
	return terra.ReferenceAsString(scpd.ref.Append("provisioning_artifact_id"))
}

func (scpd ServiceCatalogProvisioningDetailsAttributes) ProvisioningParameter() terra.ListValue[ProvisioningParameterAttributes] {
	return terra.ReferenceAsList[ProvisioningParameterAttributes](scpd.ref.Append("provisioning_parameter"))
}

type ProvisioningParameterAttributes struct {
	ref terra.Reference
}

func (pp ProvisioningParameterAttributes) InternalRef() terra.Reference {
	return pp.ref
}

func (pp ProvisioningParameterAttributes) InternalWithRef(ref terra.Reference) ProvisioningParameterAttributes {
	return ProvisioningParameterAttributes{ref: ref}
}

func (pp ProvisioningParameterAttributes) InternalTokens() hclwrite.Tokens {
	return pp.ref.InternalTokens()
}

func (pp ProvisioningParameterAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(pp.ref.Append("key"))
}

func (pp ProvisioningParameterAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(pp.ref.Append("value"))
}

type ServiceCatalogProvisioningDetailsState struct {
	PathId                 string                       `json:"path_id"`
	ProductId              string                       `json:"product_id"`
	ProvisioningArtifactId string                       `json:"provisioning_artifact_id"`
	ProvisioningParameter  []ProvisioningParameterState `json:"provisioning_parameter"`
}

type ProvisioningParameterState struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}