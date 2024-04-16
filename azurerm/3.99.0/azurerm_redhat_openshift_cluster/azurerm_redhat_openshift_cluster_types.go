// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_redhat_openshift_cluster

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ApiServerProfile struct {
	// Visibility: string, required
	Visibility terra.StringValue `hcl:"visibility,attr" validate:"required"`
}

type ClusterProfile struct {
	// Domain: string, required
	Domain terra.StringValue `hcl:"domain,attr" validate:"required"`
	// FipsEnabled: bool, optional
	FipsEnabled terra.BoolValue `hcl:"fips_enabled,attr"`
	// PullSecret: string, optional
	PullSecret terra.StringValue `hcl:"pull_secret,attr"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
}

type IngressProfile struct {
	// Visibility: string, required
	Visibility terra.StringValue `hcl:"visibility,attr" validate:"required"`
}

type MainProfile struct {
	// DiskEncryptionSetId: string, optional
	DiskEncryptionSetId terra.StringValue `hcl:"disk_encryption_set_id,attr"`
	// EncryptionAtHostEnabled: bool, optional
	EncryptionAtHostEnabled terra.BoolValue `hcl:"encryption_at_host_enabled,attr"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// VmSize: string, required
	VmSize terra.StringValue `hcl:"vm_size,attr" validate:"required"`
}

type NetworkProfile struct {
	// OutboundType: string, optional
	OutboundType terra.StringValue `hcl:"outbound_type,attr"`
	// PodCidr: string, required
	PodCidr terra.StringValue `hcl:"pod_cidr,attr" validate:"required"`
	// ServiceCidr: string, required
	ServiceCidr terra.StringValue `hcl:"service_cidr,attr" validate:"required"`
}

type ServicePrincipal struct {
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, required
	ClientSecret terra.StringValue `hcl:"client_secret,attr" validate:"required"`
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

type WorkerProfile struct {
	// DiskEncryptionSetId: string, optional
	DiskEncryptionSetId terra.StringValue `hcl:"disk_encryption_set_id,attr"`
	// DiskSizeGb: number, required
	DiskSizeGb terra.NumberValue `hcl:"disk_size_gb,attr" validate:"required"`
	// EncryptionAtHostEnabled: bool, optional
	EncryptionAtHostEnabled terra.BoolValue `hcl:"encryption_at_host_enabled,attr"`
	// NodeCount: number, required
	NodeCount terra.NumberValue `hcl:"node_count,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// VmSize: string, required
	VmSize terra.StringValue `hcl:"vm_size,attr" validate:"required"`
}

type ApiServerProfileAttributes struct {
	ref terra.Reference
}

func (asp ApiServerProfileAttributes) InternalRef() (terra.Reference, error) {
	return asp.ref, nil
}

func (asp ApiServerProfileAttributes) InternalWithRef(ref terra.Reference) ApiServerProfileAttributes {
	return ApiServerProfileAttributes{ref: ref}
}

func (asp ApiServerProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return asp.ref.InternalTokens()
}

func (asp ApiServerProfileAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("ip_address"))
}

func (asp ApiServerProfileAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("url"))
}

func (asp ApiServerProfileAttributes) Visibility() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("visibility"))
}

type ClusterProfileAttributes struct {
	ref terra.Reference
}

func (cp ClusterProfileAttributes) InternalRef() (terra.Reference, error) {
	return cp.ref, nil
}

func (cp ClusterProfileAttributes) InternalWithRef(ref terra.Reference) ClusterProfileAttributes {
	return ClusterProfileAttributes{ref: ref}
}

func (cp ClusterProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cp.ref.InternalTokens()
}

func (cp ClusterProfileAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("domain"))
}

func (cp ClusterProfileAttributes) FipsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cp.ref.Append("fips_enabled"))
}

func (cp ClusterProfileAttributes) PullSecret() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("pull_secret"))
}

func (cp ClusterProfileAttributes) ResourceGroupId() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("resource_group_id"))
}

func (cp ClusterProfileAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("version"))
}

type IngressProfileAttributes struct {
	ref terra.Reference
}

func (ip IngressProfileAttributes) InternalRef() (terra.Reference, error) {
	return ip.ref, nil
}

func (ip IngressProfileAttributes) InternalWithRef(ref terra.Reference) IngressProfileAttributes {
	return IngressProfileAttributes{ref: ref}
}

func (ip IngressProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ip.ref.InternalTokens()
}

func (ip IngressProfileAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("ip_address"))
}

func (ip IngressProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("name"))
}

func (ip IngressProfileAttributes) Visibility() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("visibility"))
}

type MainProfileAttributes struct {
	ref terra.Reference
}

func (mp MainProfileAttributes) InternalRef() (terra.Reference, error) {
	return mp.ref, nil
}

func (mp MainProfileAttributes) InternalWithRef(ref terra.Reference) MainProfileAttributes {
	return MainProfileAttributes{ref: ref}
}

func (mp MainProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mp.ref.InternalTokens()
}

func (mp MainProfileAttributes) DiskEncryptionSetId() terra.StringValue {
	return terra.ReferenceAsString(mp.ref.Append("disk_encryption_set_id"))
}

func (mp MainProfileAttributes) EncryptionAtHostEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mp.ref.Append("encryption_at_host_enabled"))
}

func (mp MainProfileAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(mp.ref.Append("subnet_id"))
}

func (mp MainProfileAttributes) VmSize() terra.StringValue {
	return terra.ReferenceAsString(mp.ref.Append("vm_size"))
}

type NetworkProfileAttributes struct {
	ref terra.Reference
}

func (np NetworkProfileAttributes) InternalRef() (terra.Reference, error) {
	return np.ref, nil
}

func (np NetworkProfileAttributes) InternalWithRef(ref terra.Reference) NetworkProfileAttributes {
	return NetworkProfileAttributes{ref: ref}
}

func (np NetworkProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return np.ref.InternalTokens()
}

func (np NetworkProfileAttributes) OutboundType() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("outbound_type"))
}

func (np NetworkProfileAttributes) PodCidr() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("pod_cidr"))
}

func (np NetworkProfileAttributes) ServiceCidr() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("service_cidr"))
}

type ServicePrincipalAttributes struct {
	ref terra.Reference
}

func (sp ServicePrincipalAttributes) InternalRef() (terra.Reference, error) {
	return sp.ref, nil
}

func (sp ServicePrincipalAttributes) InternalWithRef(ref terra.Reference) ServicePrincipalAttributes {
	return ServicePrincipalAttributes{ref: ref}
}

func (sp ServicePrincipalAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sp.ref.InternalTokens()
}

func (sp ServicePrincipalAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("client_id"))
}

func (sp ServicePrincipalAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("client_secret"))
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

type WorkerProfileAttributes struct {
	ref terra.Reference
}

func (wp WorkerProfileAttributes) InternalRef() (terra.Reference, error) {
	return wp.ref, nil
}

func (wp WorkerProfileAttributes) InternalWithRef(ref terra.Reference) WorkerProfileAttributes {
	return WorkerProfileAttributes{ref: ref}
}

func (wp WorkerProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wp.ref.InternalTokens()
}

func (wp WorkerProfileAttributes) DiskEncryptionSetId() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("disk_encryption_set_id"))
}

func (wp WorkerProfileAttributes) DiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(wp.ref.Append("disk_size_gb"))
}

func (wp WorkerProfileAttributes) EncryptionAtHostEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wp.ref.Append("encryption_at_host_enabled"))
}

func (wp WorkerProfileAttributes) NodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(wp.ref.Append("node_count"))
}

func (wp WorkerProfileAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("subnet_id"))
}

func (wp WorkerProfileAttributes) VmSize() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("vm_size"))
}

type ApiServerProfileState struct {
	IpAddress  string `json:"ip_address"`
	Url        string `json:"url"`
	Visibility string `json:"visibility"`
}

type ClusterProfileState struct {
	Domain          string `json:"domain"`
	FipsEnabled     bool   `json:"fips_enabled"`
	PullSecret      string `json:"pull_secret"`
	ResourceGroupId string `json:"resource_group_id"`
	Version         string `json:"version"`
}

type IngressProfileState struct {
	IpAddress  string `json:"ip_address"`
	Name       string `json:"name"`
	Visibility string `json:"visibility"`
}

type MainProfileState struct {
	DiskEncryptionSetId     string `json:"disk_encryption_set_id"`
	EncryptionAtHostEnabled bool   `json:"encryption_at_host_enabled"`
	SubnetId                string `json:"subnet_id"`
	VmSize                  string `json:"vm_size"`
}

type NetworkProfileState struct {
	OutboundType string `json:"outbound_type"`
	PodCidr      string `json:"pod_cidr"`
	ServiceCidr  string `json:"service_cidr"`
}

type ServicePrincipalState struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type WorkerProfileState struct {
	DiskEncryptionSetId     string  `json:"disk_encryption_set_id"`
	DiskSizeGb              float64 `json:"disk_size_gb"`
	EncryptionAtHostEnabled bool    `json:"encryption_at_host_enabled"`
	NodeCount               float64 `json:"node_count"`
	SubnetId                string  `json:"subnet_id"`
	VmSize                  string  `json:"vm_size"`
}
