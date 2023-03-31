// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package appservicecertificateorder

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Certificates struct{}

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

type CertificatesAttributes struct {
	ref terra.Reference
}

func (c CertificatesAttributes) InternalRef() terra.Reference {
	return c.ref
}

func (c CertificatesAttributes) InternalWithRef(ref terra.Reference) CertificatesAttributes {
	return CertificatesAttributes{ref: ref}
}

func (c CertificatesAttributes) InternalTokens() hclwrite.Tokens {
	return c.ref.InternalTokens()
}

func (c CertificatesAttributes) CertificateName() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("certificate_name"))
}

func (c CertificatesAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("key_vault_id"))
}

func (c CertificatesAttributes) KeyVaultSecretName() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("key_vault_secret_name"))
}

func (c CertificatesAttributes) ProvisioningState() terra.StringValue {
	return terra.ReferenceString(c.ref.Append("provisioning_state"))
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

type CertificatesState struct {
	CertificateName    string `json:"certificate_name"`
	KeyVaultId         string `json:"key_vault_id"`
	KeyVaultSecretName string `json:"key_vault_secret_name"`
	ProvisioningState  string `json:"provisioning_state"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
