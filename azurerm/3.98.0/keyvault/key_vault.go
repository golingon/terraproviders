// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package keyvault

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AccessPolicy struct {
	// ApplicationId: string, optional
	ApplicationId terra.StringValue `hcl:"application_id,attr"`
	// CertificatePermissions: list of string, optional
	CertificatePermissions terra.ListValue[terra.StringValue] `hcl:"certificate_permissions,attr"`
	// KeyPermissions: list of string, optional
	KeyPermissions terra.ListValue[terra.StringValue] `hcl:"key_permissions,attr"`
	// ObjectId: string, optional
	ObjectId terra.StringValue `hcl:"object_id,attr"`
	// SecretPermissions: list of string, optional
	SecretPermissions terra.ListValue[terra.StringValue] `hcl:"secret_permissions,attr"`
	// StoragePermissions: list of string, optional
	StoragePermissions terra.ListValue[terra.StringValue] `hcl:"storage_permissions,attr"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
}

type Contact struct {
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Phone: string, optional
	Phone terra.StringValue `hcl:"phone,attr"`
}

type NetworkAcls struct {
	// Bypass: string, required
	Bypass terra.StringValue `hcl:"bypass,attr" validate:"required"`
	// DefaultAction: string, required
	DefaultAction terra.StringValue `hcl:"default_action,attr" validate:"required"`
	// IpRules: set of string, optional
	IpRules terra.SetValue[terra.StringValue] `hcl:"ip_rules,attr"`
	// VirtualNetworkSubnetIds: set of string, optional
	VirtualNetworkSubnetIds terra.SetValue[terra.StringValue] `hcl:"virtual_network_subnet_ids,attr"`
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

type ContactAttributes struct {
	ref terra.Reference
}

func (c ContactAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ContactAttributes) InternalWithRef(ref terra.Reference) ContactAttributes {
	return ContactAttributes{ref: ref}
}

func (c ContactAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ContactAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("email"))
}

func (c ContactAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

func (c ContactAttributes) Phone() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("phone"))
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

type AccessPolicyState struct {
	ApplicationId          string   `json:"application_id"`
	CertificatePermissions []string `json:"certificate_permissions"`
	KeyPermissions         []string `json:"key_permissions"`
	ObjectId               string   `json:"object_id"`
	SecretPermissions      []string `json:"secret_permissions"`
	StoragePermissions     []string `json:"storage_permissions"`
	TenantId               string   `json:"tenant_id"`
}

type ContactState struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type NetworkAclsState struct {
	Bypass                  string   `json:"bypass"`
	DefaultAction           string   `json:"default_action"`
	IpRules                 []string `json:"ip_rules"`
	VirtualNetworkSubnetIds []string `json:"virtual_network_subnet_ids"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
