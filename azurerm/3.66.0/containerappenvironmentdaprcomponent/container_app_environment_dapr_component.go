// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package containerappenvironmentdaprcomponent

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Metadata struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SecretName: string, optional
	SecretName terra.StringValue `hcl:"secret_name,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type Secret struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
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

type MetadataAttributes struct {
	ref terra.Reference
}

func (m MetadataAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MetadataAttributes) InternalWithRef(ref terra.Reference) MetadataAttributes {
	return MetadataAttributes{ref: ref}
}

func (m MetadataAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MetadataAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("name"))
}

func (m MetadataAttributes) SecretName() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("secret_name"))
}

func (m MetadataAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("value"))
}

type SecretAttributes struct {
	ref terra.Reference
}

func (s SecretAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SecretAttributes) InternalWithRef(ref terra.Reference) SecretAttributes {
	return SecretAttributes{ref: ref}
}

func (s SecretAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SecretAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SecretAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("value"))
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

type MetadataState struct {
	Name       string `json:"name"`
	SecretName string `json:"secret_name"`
	Value      string `json:"value"`
}

type SecretState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
