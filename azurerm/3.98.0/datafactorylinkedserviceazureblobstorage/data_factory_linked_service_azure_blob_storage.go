// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datafactorylinkedserviceazureblobstorage

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type KeyVaultSasToken struct {
	// LinkedServiceName: string, required
	LinkedServiceName terra.StringValue `hcl:"linked_service_name,attr" validate:"required"`
	// SecretName: string, required
	SecretName terra.StringValue `hcl:"secret_name,attr" validate:"required"`
}

type ServicePrincipalLinkedKeyVaultKey struct {
	// LinkedServiceName: string, required
	LinkedServiceName terra.StringValue `hcl:"linked_service_name,attr" validate:"required"`
	// SecretName: string, required
	SecretName terra.StringValue `hcl:"secret_name,attr" validate:"required"`
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

type KeyVaultSasTokenAttributes struct {
	ref terra.Reference
}

func (kvst KeyVaultSasTokenAttributes) InternalRef() (terra.Reference, error) {
	return kvst.ref, nil
}

func (kvst KeyVaultSasTokenAttributes) InternalWithRef(ref terra.Reference) KeyVaultSasTokenAttributes {
	return KeyVaultSasTokenAttributes{ref: ref}
}

func (kvst KeyVaultSasTokenAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kvst.ref.InternalTokens()
}

func (kvst KeyVaultSasTokenAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceAsString(kvst.ref.Append("linked_service_name"))
}

func (kvst KeyVaultSasTokenAttributes) SecretName() terra.StringValue {
	return terra.ReferenceAsString(kvst.ref.Append("secret_name"))
}

type ServicePrincipalLinkedKeyVaultKeyAttributes struct {
	ref terra.Reference
}

func (splkvk ServicePrincipalLinkedKeyVaultKeyAttributes) InternalRef() (terra.Reference, error) {
	return splkvk.ref, nil
}

func (splkvk ServicePrincipalLinkedKeyVaultKeyAttributes) InternalWithRef(ref terra.Reference) ServicePrincipalLinkedKeyVaultKeyAttributes {
	return ServicePrincipalLinkedKeyVaultKeyAttributes{ref: ref}
}

func (splkvk ServicePrincipalLinkedKeyVaultKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return splkvk.ref.InternalTokens()
}

func (splkvk ServicePrincipalLinkedKeyVaultKeyAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceAsString(splkvk.ref.Append("linked_service_name"))
}

func (splkvk ServicePrincipalLinkedKeyVaultKeyAttributes) SecretName() terra.StringValue {
	return terra.ReferenceAsString(splkvk.ref.Append("secret_name"))
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

type KeyVaultSasTokenState struct {
	LinkedServiceName string `json:"linked_service_name"`
	SecretName        string `json:"secret_name"`
}

type ServicePrincipalLinkedKeyVaultKeyState struct {
	LinkedServiceName string `json:"linked_service_name"`
	SecretName        string `json:"secret_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
