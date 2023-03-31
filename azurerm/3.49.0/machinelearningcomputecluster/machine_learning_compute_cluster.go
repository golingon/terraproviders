// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package machinelearningcomputecluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type ScaleSettings struct {
	// MaxNodeCount: number, required
	MaxNodeCount terra.NumberValue `hcl:"max_node_count,attr" validate:"required"`
	// MinNodeCount: number, required
	MinNodeCount terra.NumberValue `hcl:"min_node_count,attr" validate:"required"`
	// ScaleDownNodesAfterIdleDuration: string, required
	ScaleDownNodesAfterIdleDuration terra.StringValue `hcl:"scale_down_nodes_after_idle_duration,attr" validate:"required"`
}

type Ssh struct {
	// AdminPassword: string, optional
	AdminPassword terra.StringValue `hcl:"admin_password,attr"`
	// AdminUsername: string, required
	AdminUsername terra.StringValue `hcl:"admin_username,attr" validate:"required"`
	// KeyValue: string, optional
	KeyValue terra.StringValue `hcl:"key_value,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
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

type ScaleSettingsAttributes struct {
	ref terra.Reference
}

func (ss ScaleSettingsAttributes) InternalRef() terra.Reference {
	return ss.ref
}

func (ss ScaleSettingsAttributes) InternalWithRef(ref terra.Reference) ScaleSettingsAttributes {
	return ScaleSettingsAttributes{ref: ref}
}

func (ss ScaleSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return ss.ref.InternalTokens()
}

func (ss ScaleSettingsAttributes) MaxNodeCount() terra.NumberValue {
	return terra.ReferenceNumber(ss.ref.Append("max_node_count"))
}

func (ss ScaleSettingsAttributes) MinNodeCount() terra.NumberValue {
	return terra.ReferenceNumber(ss.ref.Append("min_node_count"))
}

func (ss ScaleSettingsAttributes) ScaleDownNodesAfterIdleDuration() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("scale_down_nodes_after_idle_duration"))
}

type SshAttributes struct {
	ref terra.Reference
}

func (s SshAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s SshAttributes) InternalWithRef(ref terra.Reference) SshAttributes {
	return SshAttributes{ref: ref}
}

func (s SshAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s SshAttributes) AdminPassword() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("admin_password"))
}

func (s SshAttributes) AdminUsername() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("admin_username"))
}

func (s SshAttributes) KeyValue() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("key_value"))
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

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type ScaleSettingsState struct {
	MaxNodeCount                    float64 `json:"max_node_count"`
	MinNodeCount                    float64 `json:"min_node_count"`
	ScaleDownNodesAfterIdleDuration string  `json:"scale_down_nodes_after_idle_duration"`
}

type SshState struct {
	AdminPassword string `json:"admin_password"`
	AdminUsername string `json:"admin_username"`
	KeyValue      string `json:"key_value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
}
