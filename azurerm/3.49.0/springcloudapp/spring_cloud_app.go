// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package springcloudapp

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CustomPersistentDisk struct {
	// MountOptions: set of string, optional
	MountOptions terra.SetValue[terra.StringValue] `hcl:"mount_options,attr"`
	// MountPath: string, required
	MountPath terra.StringValue `hcl:"mount_path,attr" validate:"required"`
	// ReadOnlyEnabled: bool, optional
	ReadOnlyEnabled terra.BoolValue `hcl:"read_only_enabled,attr"`
	// ShareName: string, required
	ShareName terra.StringValue `hcl:"share_name,attr" validate:"required"`
	// StorageName: string, required
	StorageName terra.StringValue `hcl:"storage_name,attr" validate:"required"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type IngressSettings struct {
	// BackendProtocol: string, optional
	BackendProtocol terra.StringValue `hcl:"backend_protocol,attr"`
	// ReadTimeoutInSeconds: number, optional
	ReadTimeoutInSeconds terra.NumberValue `hcl:"read_timeout_in_seconds,attr"`
	// SendTimeoutInSeconds: number, optional
	SendTimeoutInSeconds terra.NumberValue `hcl:"send_timeout_in_seconds,attr"`
	// SessionAffinity: string, optional
	SessionAffinity terra.StringValue `hcl:"session_affinity,attr"`
	// SessionCookieMaxAge: number, optional
	SessionCookieMaxAge terra.NumberValue `hcl:"session_cookie_max_age,attr"`
}

type PersistentDisk struct {
	// MountPath: string, optional
	MountPath terra.StringValue `hcl:"mount_path,attr"`
	// SizeInGb: number, required
	SizeInGb terra.NumberValue `hcl:"size_in_gb,attr" validate:"required"`
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

type CustomPersistentDiskAttributes struct {
	ref terra.Reference
}

func (cpd CustomPersistentDiskAttributes) InternalRef() terra.Reference {
	return cpd.ref
}

func (cpd CustomPersistentDiskAttributes) InternalWithRef(ref terra.Reference) CustomPersistentDiskAttributes {
	return CustomPersistentDiskAttributes{ref: ref}
}

func (cpd CustomPersistentDiskAttributes) InternalTokens() hclwrite.Tokens {
	return cpd.ref.InternalTokens()
}

func (cpd CustomPersistentDiskAttributes) MountOptions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cpd.ref.Append("mount_options"))
}

func (cpd CustomPersistentDiskAttributes) MountPath() terra.StringValue {
	return terra.ReferenceString(cpd.ref.Append("mount_path"))
}

func (cpd CustomPersistentDiskAttributes) ReadOnlyEnabled() terra.BoolValue {
	return terra.ReferenceBool(cpd.ref.Append("read_only_enabled"))
}

func (cpd CustomPersistentDiskAttributes) ShareName() terra.StringValue {
	return terra.ReferenceString(cpd.ref.Append("share_name"))
}

func (cpd CustomPersistentDiskAttributes) StorageName() terra.StringValue {
	return terra.ReferenceString(cpd.ref.Append("storage_name"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() terra.Reference {
	return i.ref
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() hclwrite.Tokens {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceString(i.ref.Append("type"))
}

type IngressSettingsAttributes struct {
	ref terra.Reference
}

func (is IngressSettingsAttributes) InternalRef() terra.Reference {
	return is.ref
}

func (is IngressSettingsAttributes) InternalWithRef(ref terra.Reference) IngressSettingsAttributes {
	return IngressSettingsAttributes{ref: ref}
}

func (is IngressSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return is.ref.InternalTokens()
}

func (is IngressSettingsAttributes) BackendProtocol() terra.StringValue {
	return terra.ReferenceString(is.ref.Append("backend_protocol"))
}

func (is IngressSettingsAttributes) ReadTimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceNumber(is.ref.Append("read_timeout_in_seconds"))
}

func (is IngressSettingsAttributes) SendTimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceNumber(is.ref.Append("send_timeout_in_seconds"))
}

func (is IngressSettingsAttributes) SessionAffinity() terra.StringValue {
	return terra.ReferenceString(is.ref.Append("session_affinity"))
}

func (is IngressSettingsAttributes) SessionCookieMaxAge() terra.NumberValue {
	return terra.ReferenceNumber(is.ref.Append("session_cookie_max_age"))
}

type PersistentDiskAttributes struct {
	ref terra.Reference
}

func (pd PersistentDiskAttributes) InternalRef() terra.Reference {
	return pd.ref
}

func (pd PersistentDiskAttributes) InternalWithRef(ref terra.Reference) PersistentDiskAttributes {
	return PersistentDiskAttributes{ref: ref}
}

func (pd PersistentDiskAttributes) InternalTokens() hclwrite.Tokens {
	return pd.ref.InternalTokens()
}

func (pd PersistentDiskAttributes) MountPath() terra.StringValue {
	return terra.ReferenceString(pd.ref.Append("mount_path"))
}

func (pd PersistentDiskAttributes) SizeInGb() terra.NumberValue {
	return terra.ReferenceNumber(pd.ref.Append("size_in_gb"))
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

type CustomPersistentDiskState struct {
	MountOptions    []string `json:"mount_options"`
	MountPath       string   `json:"mount_path"`
	ReadOnlyEnabled bool     `json:"read_only_enabled"`
	ShareName       string   `json:"share_name"`
	StorageName     string   `json:"storage_name"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type IngressSettingsState struct {
	BackendProtocol      string  `json:"backend_protocol"`
	ReadTimeoutInSeconds float64 `json:"read_timeout_in_seconds"`
	SendTimeoutInSeconds float64 `json:"send_timeout_in_seconds"`
	SessionAffinity      string  `json:"session_affinity"`
	SessionCookieMaxAge  float64 `json:"session_cookie_max_age"`
}

type PersistentDiskState struct {
	MountPath string  `json:"mount_path"`
	SizeInGb  float64 `json:"size_in_gb"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
