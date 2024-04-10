// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package kubernetesfluxconfiguration

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type BlobStorage struct {
	// AccountKey: string, optional
	AccountKey terra.StringValue `hcl:"account_key,attr"`
	// ContainerId: string, required
	ContainerId terra.StringValue `hcl:"container_id,attr" validate:"required"`
	// LocalAuthReference: string, optional
	LocalAuthReference terra.StringValue `hcl:"local_auth_reference,attr"`
	// SasToken: string, optional
	SasToken terra.StringValue `hcl:"sas_token,attr"`
	// SyncIntervalInSeconds: number, optional
	SyncIntervalInSeconds terra.NumberValue `hcl:"sync_interval_in_seconds,attr"`
	// TimeoutInSeconds: number, optional
	TimeoutInSeconds terra.NumberValue `hcl:"timeout_in_seconds,attr"`
	// ManagedIdentity: optional
	ManagedIdentity *ManagedIdentity `hcl:"managed_identity,block"`
	// ServicePrincipal: optional
	ServicePrincipal *ServicePrincipal `hcl:"service_principal,block"`
}

type ManagedIdentity struct {
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
}

type ServicePrincipal struct {
	// ClientCertificateBase64: string, optional
	ClientCertificateBase64 terra.StringValue `hcl:"client_certificate_base64,attr"`
	// ClientCertificatePassword: string, optional
	ClientCertificatePassword terra.StringValue `hcl:"client_certificate_password,attr"`
	// ClientCertificateSendChain: bool, optional
	ClientCertificateSendChain terra.BoolValue `hcl:"client_certificate_send_chain,attr"`
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, optional
	ClientSecret terra.StringValue `hcl:"client_secret,attr"`
	// TenantId: string, required
	TenantId terra.StringValue `hcl:"tenant_id,attr" validate:"required"`
}

type Bucket struct {
	// AccessKey: string, optional
	AccessKey terra.StringValue `hcl:"access_key,attr"`
	// BucketName: string, required
	BucketName terra.StringValue `hcl:"bucket_name,attr" validate:"required"`
	// LocalAuthReference: string, optional
	LocalAuthReference terra.StringValue `hcl:"local_auth_reference,attr"`
	// SecretKeyBase64: string, optional
	SecretKeyBase64 terra.StringValue `hcl:"secret_key_base64,attr"`
	// SyncIntervalInSeconds: number, optional
	SyncIntervalInSeconds terra.NumberValue `hcl:"sync_interval_in_seconds,attr"`
	// TimeoutInSeconds: number, optional
	TimeoutInSeconds terra.NumberValue `hcl:"timeout_in_seconds,attr"`
	// TlsEnabled: bool, optional
	TlsEnabled terra.BoolValue `hcl:"tls_enabled,attr"`
	// Url: string, required
	Url terra.StringValue `hcl:"url,attr" validate:"required"`
}

type GitRepository struct {
	// HttpsCaCertBase64: string, optional
	HttpsCaCertBase64 terra.StringValue `hcl:"https_ca_cert_base64,attr"`
	// HttpsKeyBase64: string, optional
	HttpsKeyBase64 terra.StringValue `hcl:"https_key_base64,attr"`
	// HttpsUser: string, optional
	HttpsUser terra.StringValue `hcl:"https_user,attr"`
	// LocalAuthReference: string, optional
	LocalAuthReference terra.StringValue `hcl:"local_auth_reference,attr"`
	// ReferenceType: string, required
	ReferenceType terra.StringValue `hcl:"reference_type,attr" validate:"required"`
	// ReferenceValue: string, required
	ReferenceValue terra.StringValue `hcl:"reference_value,attr" validate:"required"`
	// SshKnownHostsBase64: string, optional
	SshKnownHostsBase64 terra.StringValue `hcl:"ssh_known_hosts_base64,attr"`
	// SshPrivateKeyBase64: string, optional
	SshPrivateKeyBase64 terra.StringValue `hcl:"ssh_private_key_base64,attr"`
	// SyncIntervalInSeconds: number, optional
	SyncIntervalInSeconds terra.NumberValue `hcl:"sync_interval_in_seconds,attr"`
	// TimeoutInSeconds: number, optional
	TimeoutInSeconds terra.NumberValue `hcl:"timeout_in_seconds,attr"`
	// Url: string, required
	Url terra.StringValue `hcl:"url,attr" validate:"required"`
}

type Kustomizations struct {
	// DependsOn: list of string, optional
	DependsOn terra.ListValue[terra.StringValue] `hcl:"depends_on,attr"`
	// GarbageCollectionEnabled: bool, optional
	GarbageCollectionEnabled terra.BoolValue `hcl:"garbage_collection_enabled,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// RecreatingEnabled: bool, optional
	RecreatingEnabled terra.BoolValue `hcl:"recreating_enabled,attr"`
	// RetryIntervalInSeconds: number, optional
	RetryIntervalInSeconds terra.NumberValue `hcl:"retry_interval_in_seconds,attr"`
	// SyncIntervalInSeconds: number, optional
	SyncIntervalInSeconds terra.NumberValue `hcl:"sync_interval_in_seconds,attr"`
	// TimeoutInSeconds: number, optional
	TimeoutInSeconds terra.NumberValue `hcl:"timeout_in_seconds,attr"`
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

type BlobStorageAttributes struct {
	ref terra.Reference
}

func (bs BlobStorageAttributes) InternalRef() (terra.Reference, error) {
	return bs.ref, nil
}

func (bs BlobStorageAttributes) InternalWithRef(ref terra.Reference) BlobStorageAttributes {
	return BlobStorageAttributes{ref: ref}
}

func (bs BlobStorageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bs.ref.InternalTokens()
}

func (bs BlobStorageAttributes) AccountKey() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("account_key"))
}

func (bs BlobStorageAttributes) ContainerId() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("container_id"))
}

func (bs BlobStorageAttributes) LocalAuthReference() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("local_auth_reference"))
}

func (bs BlobStorageAttributes) SasToken() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("sas_token"))
}

func (bs BlobStorageAttributes) SyncIntervalInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(bs.ref.Append("sync_interval_in_seconds"))
}

func (bs BlobStorageAttributes) TimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(bs.ref.Append("timeout_in_seconds"))
}

func (bs BlobStorageAttributes) ManagedIdentity() terra.ListValue[ManagedIdentityAttributes] {
	return terra.ReferenceAsList[ManagedIdentityAttributes](bs.ref.Append("managed_identity"))
}

func (bs BlobStorageAttributes) ServicePrincipal() terra.ListValue[ServicePrincipalAttributes] {
	return terra.ReferenceAsList[ServicePrincipalAttributes](bs.ref.Append("service_principal"))
}

type ManagedIdentityAttributes struct {
	ref terra.Reference
}

func (mi ManagedIdentityAttributes) InternalRef() (terra.Reference, error) {
	return mi.ref, nil
}

func (mi ManagedIdentityAttributes) InternalWithRef(ref terra.Reference) ManagedIdentityAttributes {
	return ManagedIdentityAttributes{ref: ref}
}

func (mi ManagedIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mi.ref.InternalTokens()
}

func (mi ManagedIdentityAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(mi.ref.Append("client_id"))
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

func (sp ServicePrincipalAttributes) ClientCertificateBase64() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("client_certificate_base64"))
}

func (sp ServicePrincipalAttributes) ClientCertificatePassword() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("client_certificate_password"))
}

func (sp ServicePrincipalAttributes) ClientCertificateSendChain() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("client_certificate_send_chain"))
}

func (sp ServicePrincipalAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("client_id"))
}

func (sp ServicePrincipalAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("client_secret"))
}

func (sp ServicePrincipalAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("tenant_id"))
}

type BucketAttributes struct {
	ref terra.Reference
}

func (b BucketAttributes) InternalRef() (terra.Reference, error) {
	return b.ref, nil
}

func (b BucketAttributes) InternalWithRef(ref terra.Reference) BucketAttributes {
	return BucketAttributes{ref: ref}
}

func (b BucketAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return b.ref.InternalTokens()
}

func (b BucketAttributes) AccessKey() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("access_key"))
}

func (b BucketAttributes) BucketName() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("bucket_name"))
}

func (b BucketAttributes) LocalAuthReference() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("local_auth_reference"))
}

func (b BucketAttributes) SecretKeyBase64() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("secret_key_base64"))
}

func (b BucketAttributes) SyncIntervalInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("sync_interval_in_seconds"))
}

func (b BucketAttributes) TimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("timeout_in_seconds"))
}

func (b BucketAttributes) TlsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(b.ref.Append("tls_enabled"))
}

func (b BucketAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("url"))
}

type GitRepositoryAttributes struct {
	ref terra.Reference
}

func (gr GitRepositoryAttributes) InternalRef() (terra.Reference, error) {
	return gr.ref, nil
}

func (gr GitRepositoryAttributes) InternalWithRef(ref terra.Reference) GitRepositoryAttributes {
	return GitRepositoryAttributes{ref: ref}
}

func (gr GitRepositoryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gr.ref.InternalTokens()
}

func (gr GitRepositoryAttributes) HttpsCaCertBase64() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("https_ca_cert_base64"))
}

func (gr GitRepositoryAttributes) HttpsKeyBase64() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("https_key_base64"))
}

func (gr GitRepositoryAttributes) HttpsUser() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("https_user"))
}

func (gr GitRepositoryAttributes) LocalAuthReference() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("local_auth_reference"))
}

func (gr GitRepositoryAttributes) ReferenceType() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("reference_type"))
}

func (gr GitRepositoryAttributes) ReferenceValue() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("reference_value"))
}

func (gr GitRepositoryAttributes) SshKnownHostsBase64() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("ssh_known_hosts_base64"))
}

func (gr GitRepositoryAttributes) SshPrivateKeyBase64() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("ssh_private_key_base64"))
}

func (gr GitRepositoryAttributes) SyncIntervalInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(gr.ref.Append("sync_interval_in_seconds"))
}

func (gr GitRepositoryAttributes) TimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(gr.ref.Append("timeout_in_seconds"))
}

func (gr GitRepositoryAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("url"))
}

type KustomizationsAttributes struct {
	ref terra.Reference
}

func (k KustomizationsAttributes) InternalRef() (terra.Reference, error) {
	return k.ref, nil
}

func (k KustomizationsAttributes) InternalWithRef(ref terra.Reference) KustomizationsAttributes {
	return KustomizationsAttributes{ref: ref}
}

func (k KustomizationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return k.ref.InternalTokens()
}

func (k KustomizationsAttributes) DependsOn() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](k.ref.Append("depends_on"))
}

func (k KustomizationsAttributes) GarbageCollectionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(k.ref.Append("garbage_collection_enabled"))
}

func (k KustomizationsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(k.ref.Append("name"))
}

func (k KustomizationsAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(k.ref.Append("path"))
}

func (k KustomizationsAttributes) RecreatingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(k.ref.Append("recreating_enabled"))
}

func (k KustomizationsAttributes) RetryIntervalInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(k.ref.Append("retry_interval_in_seconds"))
}

func (k KustomizationsAttributes) SyncIntervalInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(k.ref.Append("sync_interval_in_seconds"))
}

func (k KustomizationsAttributes) TimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(k.ref.Append("timeout_in_seconds"))
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

type BlobStorageState struct {
	AccountKey            string                  `json:"account_key"`
	ContainerId           string                  `json:"container_id"`
	LocalAuthReference    string                  `json:"local_auth_reference"`
	SasToken              string                  `json:"sas_token"`
	SyncIntervalInSeconds float64                 `json:"sync_interval_in_seconds"`
	TimeoutInSeconds      float64                 `json:"timeout_in_seconds"`
	ManagedIdentity       []ManagedIdentityState  `json:"managed_identity"`
	ServicePrincipal      []ServicePrincipalState `json:"service_principal"`
}

type ManagedIdentityState struct {
	ClientId string `json:"client_id"`
}

type ServicePrincipalState struct {
	ClientCertificateBase64    string `json:"client_certificate_base64"`
	ClientCertificatePassword  string `json:"client_certificate_password"`
	ClientCertificateSendChain bool   `json:"client_certificate_send_chain"`
	ClientId                   string `json:"client_id"`
	ClientSecret               string `json:"client_secret"`
	TenantId                   string `json:"tenant_id"`
}

type BucketState struct {
	AccessKey             string  `json:"access_key"`
	BucketName            string  `json:"bucket_name"`
	LocalAuthReference    string  `json:"local_auth_reference"`
	SecretKeyBase64       string  `json:"secret_key_base64"`
	SyncIntervalInSeconds float64 `json:"sync_interval_in_seconds"`
	TimeoutInSeconds      float64 `json:"timeout_in_seconds"`
	TlsEnabled            bool    `json:"tls_enabled"`
	Url                   string  `json:"url"`
}

type GitRepositoryState struct {
	HttpsCaCertBase64     string  `json:"https_ca_cert_base64"`
	HttpsKeyBase64        string  `json:"https_key_base64"`
	HttpsUser             string  `json:"https_user"`
	LocalAuthReference    string  `json:"local_auth_reference"`
	ReferenceType         string  `json:"reference_type"`
	ReferenceValue        string  `json:"reference_value"`
	SshKnownHostsBase64   string  `json:"ssh_known_hosts_base64"`
	SshPrivateKeyBase64   string  `json:"ssh_private_key_base64"`
	SyncIntervalInSeconds float64 `json:"sync_interval_in_seconds"`
	TimeoutInSeconds      float64 `json:"timeout_in_seconds"`
	Url                   string  `json:"url"`
}

type KustomizationsState struct {
	DependsOn                []string `json:"depends_on"`
	GarbageCollectionEnabled bool     `json:"garbage_collection_enabled"`
	Name                     string   `json:"name"`
	Path                     string   `json:"path"`
	RecreatingEnabled        bool     `json:"recreating_enabled"`
	RetryIntervalInSeconds   float64  `json:"retry_interval_in_seconds"`
	SyncIntervalInSeconds    float64  `json:"sync_interval_in_seconds"`
	TimeoutInSeconds         float64  `json:"timeout_in_seconds"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
