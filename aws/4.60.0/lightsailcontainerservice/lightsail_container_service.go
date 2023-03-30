// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package lightsailcontainerservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PrivateRegistryAccess struct {
	// EcrImagePullerRole: optional
	EcrImagePullerRole *EcrImagePullerRole `hcl:"ecr_image_puller_role,block"`
}

type EcrImagePullerRole struct {
	// IsActive: bool, optional
	IsActive terra.BoolValue `hcl:"is_active,attr"`
}

type PublicDomainNames struct {
	// Certificate: min=1
	Certificate []Certificate `hcl:"certificate,block" validate:"min=1"`
}

type Certificate struct {
	// CertificateName: string, required
	CertificateName terra.StringValue `hcl:"certificate_name,attr" validate:"required"`
	// DomainNames: list of string, required
	DomainNames terra.ListValue[terra.StringValue] `hcl:"domain_names,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type PrivateRegistryAccessAttributes struct {
	ref terra.Reference
}

func (pra PrivateRegistryAccessAttributes) InternalRef() terra.Reference {
	return pra.ref
}

func (pra PrivateRegistryAccessAttributes) InternalWithRef(ref terra.Reference) PrivateRegistryAccessAttributes {
	return PrivateRegistryAccessAttributes{ref: ref}
}

func (pra PrivateRegistryAccessAttributes) InternalTokens() hclwrite.Tokens {
	return pra.ref.InternalTokens()
}

func (pra PrivateRegistryAccessAttributes) EcrImagePullerRole() terra.ListValue[EcrImagePullerRoleAttributes] {
	return terra.ReferenceList[EcrImagePullerRoleAttributes](pra.ref.Append("ecr_image_puller_role"))
}

type EcrImagePullerRoleAttributes struct {
	ref terra.Reference
}

func (eipr EcrImagePullerRoleAttributes) InternalRef() terra.Reference {
	return eipr.ref
}

func (eipr EcrImagePullerRoleAttributes) InternalWithRef(ref terra.Reference) EcrImagePullerRoleAttributes {
	return EcrImagePullerRoleAttributes{ref: ref}
}

func (eipr EcrImagePullerRoleAttributes) InternalTokens() hclwrite.Tokens {
	return eipr.ref.InternalTokens()
}

func (eipr EcrImagePullerRoleAttributes) IsActive() terra.BoolValue {
	return terra.ReferenceBool(eipr.ref.Append("is_active"))
}

func (eipr EcrImagePullerRoleAttributes) PrincipalArn() terra.StringValue {
	return terra.ReferenceString(eipr.ref.Append("principal_arn"))
}

type PublicDomainNamesAttributes struct {
	ref terra.Reference
}

func (pdn PublicDomainNamesAttributes) InternalRef() terra.Reference {
	return pdn.ref
}

func (pdn PublicDomainNamesAttributes) InternalWithRef(ref terra.Reference) PublicDomainNamesAttributes {
	return PublicDomainNamesAttributes{ref: ref}
}

func (pdn PublicDomainNamesAttributes) InternalTokens() hclwrite.Tokens {
	return pdn.ref.InternalTokens()
}

func (pdn PublicDomainNamesAttributes) Certificate() terra.SetValue[CertificateAttributes] {
	return terra.ReferenceSet[CertificateAttributes](pdn.ref.Append("certificate"))
}

type CertificateAttributes struct {
	ref terra.Reference
}

func (c CertificateAttributes) InternalRef() terra.Reference {
	return c.ref
}

func (c CertificateAttributes) InternalWithRef(ref terra.Reference) CertificateAttributes {
	return CertificateAttributes{ref: ref}
}

func (c CertificateAttributes) InternalTokens() hclwrite.Tokens {
	return c.ref.InternalTokens()
}

func (c CertificateAttributes) CertificateName() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("certificate_name"))
}

func (c CertificateAttributes) DomainNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](c.ref.Append("domain_names"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type PrivateRegistryAccessState struct {
	EcrImagePullerRole []EcrImagePullerRoleState `json:"ecr_image_puller_role"`
}

type EcrImagePullerRoleState struct {
	IsActive     bool   `json:"is_active"`
	PrincipalArn string `json:"principal_arn"`
}

type PublicDomainNamesState struct {
	Certificate []CertificateState `json:"certificate"`
}

type CertificateState struct {
	CertificateName string   `json:"certificate_name"`
	DomainNames     []string `json:"domain_names"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
