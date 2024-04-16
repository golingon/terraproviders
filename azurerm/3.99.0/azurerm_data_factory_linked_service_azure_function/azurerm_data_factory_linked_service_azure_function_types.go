// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_data_factory_linked_service_azure_function

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type KeyVaultKey struct {
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

type KeyVaultKeyAttributes struct {
	ref terra.Reference
}

func (kvk KeyVaultKeyAttributes) InternalRef() (terra.Reference, error) {
	return kvk.ref, nil
}

func (kvk KeyVaultKeyAttributes) InternalWithRef(ref terra.Reference) KeyVaultKeyAttributes {
	return KeyVaultKeyAttributes{ref: ref}
}

func (kvk KeyVaultKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kvk.ref.InternalTokens()
}

func (kvk KeyVaultKeyAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("linked_service_name"))
}

func (kvk KeyVaultKeyAttributes) SecretName() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("secret_name"))
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

type KeyVaultKeyState struct {
	LinkedServiceName string `json:"linked_service_name"`
	SecretName        string `json:"secret_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
