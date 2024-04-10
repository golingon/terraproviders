// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datakeyvault

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AccessPolicy struct{}

type NetworkAcls struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type AccessPolicyAttributes struct {
	ref terra.Reference
}

func (ap AccessPolicyAttributes) InternalRef() (terra.Reference, error) {
	return ap.ref, nil
}

func (ap AccessPolicyAttributes) InternalWithRef(ref terra.Reference) AccessPolicyAttributes {
	return AccessPolicyAttributes{ref: ref}
}

func (ap AccessPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ap.ref.InternalTokens()
}

func (ap AccessPolicyAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("application_id"))
}

func (ap AccessPolicyAttributes) CertificatePermissions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ap.ref.Append("certificate_permissions"))
}

func (ap AccessPolicyAttributes) KeyPermissions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ap.ref.Append("key_permissions"))
}

func (ap AccessPolicyAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("object_id"))
}

func (ap AccessPolicyAttributes) SecretPermissions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ap.ref.Append("secret_permissions"))
}

func (ap AccessPolicyAttributes) StoragePermissions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ap.ref.Append("storage_permissions"))
}

func (ap AccessPolicyAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("tenant_id"))
}

type NetworkAclsAttributes struct {
	ref terra.Reference
}

func (na NetworkAclsAttributes) InternalRef() (terra.Reference, error) {
	return na.ref, nil
}

func (na NetworkAclsAttributes) InternalWithRef(ref terra.Reference) NetworkAclsAttributes {
	return NetworkAclsAttributes{ref: ref}
}

func (na NetworkAclsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return na.ref.InternalTokens()
}

func (na NetworkAclsAttributes) Bypass() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("bypass"))
}

func (na NetworkAclsAttributes) DefaultAction() terra.StringValue {
	return terra.ReferenceAsString(na.ref.Append("default_action"))
}

func (na NetworkAclsAttributes) IpRules() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](na.ref.Append("ip_rules"))
}

func (na NetworkAclsAttributes) VirtualNetworkSubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](na.ref.Append("virtual_network_subnet_ids"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type AccessPolicyState struct {
	ApplicationId          string   `json:"application_id"`
	CertificatePermissions []string `json:"certificate_permissions"`
	KeyPermissions         []string `json:"key_permissions"`
	ObjectId               string   `json:"object_id"`
	SecretPermissions      []string `json:"secret_permissions"`
	StoragePermissions     []string `json:"storage_permissions"`
	TenantId               string   `json:"tenant_id"`
}

type NetworkAclsState struct {
	Bypass                  string   `json:"bypass"`
	DefaultAction           string   `json:"default_action"`
	IpRules                 []string `json:"ip_rules"`
	VirtualNetworkSubnetIds []string `json:"virtual_network_subnet_ids"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
