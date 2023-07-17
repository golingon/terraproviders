// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datafactorylinkedservicesqlserver

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type KeyVaultConnectionString struct {
	// LinkedServiceName: string, required
	LinkedServiceName terra.StringValue `hcl:"linked_service_name,attr" validate:"required"`
	// SecretName: string, required
	SecretName terra.StringValue `hcl:"secret_name,attr" validate:"required"`
}

type KeyVaultPassword struct {
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

type KeyVaultConnectionStringAttributes struct {
	ref terra.Reference
}

func (kvcs KeyVaultConnectionStringAttributes) InternalRef() (terra.Reference, error) {
	return kvcs.ref, nil
}

func (kvcs KeyVaultConnectionStringAttributes) InternalWithRef(ref terra.Reference) KeyVaultConnectionStringAttributes {
	return KeyVaultConnectionStringAttributes{ref: ref}
}

func (kvcs KeyVaultConnectionStringAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kvcs.ref.InternalTokens()
}

func (kvcs KeyVaultConnectionStringAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceAsString(kvcs.ref.Append("linked_service_name"))
}

func (kvcs KeyVaultConnectionStringAttributes) SecretName() terra.StringValue {
	return terra.ReferenceAsString(kvcs.ref.Append("secret_name"))
}

type KeyVaultPasswordAttributes struct {
	ref terra.Reference
}

func (kvp KeyVaultPasswordAttributes) InternalRef() (terra.Reference, error) {
	return kvp.ref, nil
}

func (kvp KeyVaultPasswordAttributes) InternalWithRef(ref terra.Reference) KeyVaultPasswordAttributes {
	return KeyVaultPasswordAttributes{ref: ref}
}

func (kvp KeyVaultPasswordAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kvp.ref.InternalTokens()
}

func (kvp KeyVaultPasswordAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceAsString(kvp.ref.Append("linked_service_name"))
}

func (kvp KeyVaultPasswordAttributes) SecretName() terra.StringValue {
	return terra.ReferenceAsString(kvp.ref.Append("secret_name"))
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

type KeyVaultConnectionStringState struct {
	LinkedServiceName string `json:"linked_service_name"`
	SecretName        string `json:"secret_name"`
}

type KeyVaultPasswordState struct {
	LinkedServiceName string `json:"linked_service_name"`
	SecretName        string `json:"secret_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}