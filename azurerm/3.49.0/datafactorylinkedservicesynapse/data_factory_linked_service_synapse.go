// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datafactorylinkedservicesynapse

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

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

type KeyVaultPasswordAttributes struct {
	ref terra.Reference
}

func (kvp KeyVaultPasswordAttributes) InternalRef() terra.Reference {
	return kvp.ref
}

func (kvp KeyVaultPasswordAttributes) InternalWithRef(ref terra.Reference) KeyVaultPasswordAttributes {
	return KeyVaultPasswordAttributes{ref: ref}
}

func (kvp KeyVaultPasswordAttributes) InternalTokens() hclwrite.Tokens {
	return kvp.ref.InternalTokens()
}

func (kvp KeyVaultPasswordAttributes) LinkedServiceName() terra.StringValue {
	return terra.ReferenceString(kvp.ref.Append("linked_service_name"))
}

func (kvp KeyVaultPasswordAttributes) SecretName() terra.StringValue {
	return terra.ReferenceString(kvp.ref.Append("secret_name"))
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
