// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package resourcedeploymentscriptazurecli

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Container struct {
	// ContainerGroupName: string, optional
	ContainerGroupName terra.StringValue `hcl:"container_group_name,attr"`
}

type EnvironmentVariable struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SecureValue: string, optional
	SecureValue terra.StringValue `hcl:"secure_value,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type Identity struct {
	// IdentityIds: set of string, required
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type StorageAccount struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
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

type ContainerAttributes struct {
	ref terra.Reference
}

func (c ContainerAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ContainerAttributes) InternalWithRef(ref terra.Reference) ContainerAttributes {
	return ContainerAttributes{ref: ref}
}

func (c ContainerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ContainerAttributes) ContainerGroupName() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("container_group_name"))
}

type EnvironmentVariableAttributes struct {
	ref terra.Reference
}

func (ev EnvironmentVariableAttributes) InternalRef() (terra.Reference, error) {
	return ev.ref, nil
}

func (ev EnvironmentVariableAttributes) InternalWithRef(ref terra.Reference) EnvironmentVariableAttributes {
	return EnvironmentVariableAttributes{ref: ref}
}

func (ev EnvironmentVariableAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ev.ref.InternalTokens()
}

func (ev EnvironmentVariableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("name"))
}

func (ev EnvironmentVariableAttributes) SecureValue() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("secure_value"))
}

func (ev EnvironmentVariableAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("value"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type StorageAccountAttributes struct {
	ref terra.Reference
}

func (sa StorageAccountAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa StorageAccountAttributes) InternalWithRef(ref terra.Reference) StorageAccountAttributes {
	return StorageAccountAttributes{ref: ref}
}

func (sa StorageAccountAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa StorageAccountAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("key"))
}

func (sa StorageAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("name"))
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

type ContainerState struct {
	ContainerGroupName string `json:"container_group_name"`
}

type EnvironmentVariableState struct {
	Name        string `json:"name"`
	SecureValue string `json:"secure_value"`
	Value       string `json:"value"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	Type        string   `json:"type"`
}

type StorageAccountState struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
