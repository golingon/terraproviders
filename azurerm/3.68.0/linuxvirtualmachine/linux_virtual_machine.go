// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package linuxvirtualmachine

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AdditionalCapabilities struct {
	// UltraSsdEnabled: bool, optional
	UltraSsdEnabled terra.BoolValue `hcl:"ultra_ssd_enabled,attr"`
}

type AdminSshKey struct {
	// PublicKey: string, required
	PublicKey terra.StringValue `hcl:"public_key,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type BootDiagnostics struct {
	// StorageAccountUri: string, optional
	StorageAccountUri terra.StringValue `hcl:"storage_account_uri,attr"`
}

type GalleryApplication struct {
	// ConfigurationBlobUri: string, optional
	ConfigurationBlobUri terra.StringValue `hcl:"configuration_blob_uri,attr"`
	// Order: number, optional
	Order terra.NumberValue `hcl:"order,attr"`
	// Tag: string, optional
	Tag terra.StringValue `hcl:"tag,attr"`
	// VersionId: string, required
	VersionId terra.StringValue `hcl:"version_id,attr" validate:"required"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type OsDisk struct {
	// Caching: string, required
	Caching terra.StringValue `hcl:"caching,attr" validate:"required"`
	// DiskEncryptionSetId: string, optional
	DiskEncryptionSetId terra.StringValue `hcl:"disk_encryption_set_id,attr"`
	// DiskSizeGb: number, optional
	DiskSizeGb terra.NumberValue `hcl:"disk_size_gb,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// SecureVmDiskEncryptionSetId: string, optional
	SecureVmDiskEncryptionSetId terra.StringValue `hcl:"secure_vm_disk_encryption_set_id,attr"`
	// SecurityEncryptionType: string, optional
	SecurityEncryptionType terra.StringValue `hcl:"security_encryption_type,attr"`
	// StorageAccountType: string, required
	StorageAccountType terra.StringValue `hcl:"storage_account_type,attr" validate:"required"`
	// WriteAcceleratorEnabled: bool, optional
	WriteAcceleratorEnabled terra.BoolValue `hcl:"write_accelerator_enabled,attr"`
	// DiffDiskSettings: optional
	DiffDiskSettings *DiffDiskSettings `hcl:"diff_disk_settings,block"`
}

type DiffDiskSettings struct {
	// Option: string, required
	Option terra.StringValue `hcl:"option,attr" validate:"required"`
	// Placement: string, optional
	Placement terra.StringValue `hcl:"placement,attr"`
}

type Plan struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Product: string, required
	Product terra.StringValue `hcl:"product,attr" validate:"required"`
	// Publisher: string, required
	Publisher terra.StringValue `hcl:"publisher,attr" validate:"required"`
}

type Secret struct {
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// Certificate: min=1
	Certificate []Certificate `hcl:"certificate,block" validate:"min=1"`
}

type Certificate struct {
	// Url: string, required
	Url terra.StringValue `hcl:"url,attr" validate:"required"`
}

type SourceImageReference struct {
	// Offer: string, required
	Offer terra.StringValue `hcl:"offer,attr" validate:"required"`
	// Publisher: string, required
	Publisher terra.StringValue `hcl:"publisher,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
}

type TerminationNotification struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Timeout: string, optional
	Timeout terra.StringValue `hcl:"timeout,attr"`
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

type AdditionalCapabilitiesAttributes struct {
	ref terra.Reference
}

func (ac AdditionalCapabilitiesAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac AdditionalCapabilitiesAttributes) InternalWithRef(ref terra.Reference) AdditionalCapabilitiesAttributes {
	return AdditionalCapabilitiesAttributes{ref: ref}
}

func (ac AdditionalCapabilitiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac AdditionalCapabilitiesAttributes) UltraSsdEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ac.ref.Append("ultra_ssd_enabled"))
}

type AdminSshKeyAttributes struct {
	ref terra.Reference
}

func (ask AdminSshKeyAttributes) InternalRef() (terra.Reference, error) {
	return ask.ref, nil
}

func (ask AdminSshKeyAttributes) InternalWithRef(ref terra.Reference) AdminSshKeyAttributes {
	return AdminSshKeyAttributes{ref: ref}
}

func (ask AdminSshKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ask.ref.InternalTokens()
}

func (ask AdminSshKeyAttributes) PublicKey() terra.StringValue {
	return terra.ReferenceAsString(ask.ref.Append("public_key"))
}

func (ask AdminSshKeyAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(ask.ref.Append("username"))
}

type BootDiagnosticsAttributes struct {
	ref terra.Reference
}

func (bd BootDiagnosticsAttributes) InternalRef() (terra.Reference, error) {
	return bd.ref, nil
}

func (bd BootDiagnosticsAttributes) InternalWithRef(ref terra.Reference) BootDiagnosticsAttributes {
	return BootDiagnosticsAttributes{ref: ref}
}

func (bd BootDiagnosticsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bd.ref.InternalTokens()
}

func (bd BootDiagnosticsAttributes) StorageAccountUri() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("storage_account_uri"))
}

type GalleryApplicationAttributes struct {
	ref terra.Reference
}

func (ga GalleryApplicationAttributes) InternalRef() (terra.Reference, error) {
	return ga.ref, nil
}

func (ga GalleryApplicationAttributes) InternalWithRef(ref terra.Reference) GalleryApplicationAttributes {
	return GalleryApplicationAttributes{ref: ref}
}

func (ga GalleryApplicationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ga.ref.InternalTokens()
}

func (ga GalleryApplicationAttributes) ConfigurationBlobUri() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("configuration_blob_uri"))
}

func (ga GalleryApplicationAttributes) Order() terra.NumberValue {
	return terra.ReferenceAsNumber(ga.ref.Append("order"))
}

func (ga GalleryApplicationAttributes) Tag() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("tag"))
}

func (ga GalleryApplicationAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("version_id"))
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

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type OsDiskAttributes struct {
	ref terra.Reference
}

func (od OsDiskAttributes) InternalRef() (terra.Reference, error) {
	return od.ref, nil
}

func (od OsDiskAttributes) InternalWithRef(ref terra.Reference) OsDiskAttributes {
	return OsDiskAttributes{ref: ref}
}

func (od OsDiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return od.ref.InternalTokens()
}

func (od OsDiskAttributes) Caching() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("caching"))
}

func (od OsDiskAttributes) DiskEncryptionSetId() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("disk_encryption_set_id"))
}

func (od OsDiskAttributes) DiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(od.ref.Append("disk_size_gb"))
}

func (od OsDiskAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("name"))
}

func (od OsDiskAttributes) SecureVmDiskEncryptionSetId() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("secure_vm_disk_encryption_set_id"))
}

func (od OsDiskAttributes) SecurityEncryptionType() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("security_encryption_type"))
}

func (od OsDiskAttributes) StorageAccountType() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("storage_account_type"))
}

func (od OsDiskAttributes) WriteAcceleratorEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(od.ref.Append("write_accelerator_enabled"))
}

func (od OsDiskAttributes) DiffDiskSettings() terra.ListValue[DiffDiskSettingsAttributes] {
	return terra.ReferenceAsList[DiffDiskSettingsAttributes](od.ref.Append("diff_disk_settings"))
}

type DiffDiskSettingsAttributes struct {
	ref terra.Reference
}

func (dds DiffDiskSettingsAttributes) InternalRef() (terra.Reference, error) {
	return dds.ref, nil
}

func (dds DiffDiskSettingsAttributes) InternalWithRef(ref terra.Reference) DiffDiskSettingsAttributes {
	return DiffDiskSettingsAttributes{ref: ref}
}

func (dds DiffDiskSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dds.ref.InternalTokens()
}

func (dds DiffDiskSettingsAttributes) Option() terra.StringValue {
	return terra.ReferenceAsString(dds.ref.Append("option"))
}

func (dds DiffDiskSettingsAttributes) Placement() terra.StringValue {
	return terra.ReferenceAsString(dds.ref.Append("placement"))
}

type PlanAttributes struct {
	ref terra.Reference
}

func (p PlanAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PlanAttributes) InternalWithRef(ref terra.Reference) PlanAttributes {
	return PlanAttributes{ref: ref}
}

func (p PlanAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

func (p PlanAttributes) Product() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("product"))
}

func (p PlanAttributes) Publisher() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("publisher"))
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

func (s SecretAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("key_vault_id"))
}

func (s SecretAttributes) Certificate() terra.SetValue[CertificateAttributes] {
	return terra.ReferenceAsSet[CertificateAttributes](s.ref.Append("certificate"))
}

type CertificateAttributes struct {
	ref terra.Reference
}

func (c CertificateAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CertificateAttributes) InternalWithRef(ref terra.Reference) CertificateAttributes {
	return CertificateAttributes{ref: ref}
}

func (c CertificateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CertificateAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("url"))
}

type SourceImageReferenceAttributes struct {
	ref terra.Reference
}

func (sir SourceImageReferenceAttributes) InternalRef() (terra.Reference, error) {
	return sir.ref, nil
}

func (sir SourceImageReferenceAttributes) InternalWithRef(ref terra.Reference) SourceImageReferenceAttributes {
	return SourceImageReferenceAttributes{ref: ref}
}

func (sir SourceImageReferenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sir.ref.InternalTokens()
}

func (sir SourceImageReferenceAttributes) Offer() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("offer"))
}

func (sir SourceImageReferenceAttributes) Publisher() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("publisher"))
}

func (sir SourceImageReferenceAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("sku"))
}

func (sir SourceImageReferenceAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("version"))
}

type TerminationNotificationAttributes struct {
	ref terra.Reference
}

func (tn TerminationNotificationAttributes) InternalRef() (terra.Reference, error) {
	return tn.ref, nil
}

func (tn TerminationNotificationAttributes) InternalWithRef(ref terra.Reference) TerminationNotificationAttributes {
	return TerminationNotificationAttributes{ref: ref}
}

func (tn TerminationNotificationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tn.ref.InternalTokens()
}

func (tn TerminationNotificationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(tn.ref.Append("enabled"))
}

func (tn TerminationNotificationAttributes) Timeout() terra.StringValue {
	return terra.ReferenceAsString(tn.ref.Append("timeout"))
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

type AdditionalCapabilitiesState struct {
	UltraSsdEnabled bool `json:"ultra_ssd_enabled"`
}

type AdminSshKeyState struct {
	PublicKey string `json:"public_key"`
	Username  string `json:"username"`
}

type BootDiagnosticsState struct {
	StorageAccountUri string `json:"storage_account_uri"`
}

type GalleryApplicationState struct {
	ConfigurationBlobUri string  `json:"configuration_blob_uri"`
	Order                float64 `json:"order"`
	Tag                  string  `json:"tag"`
	VersionId            string  `json:"version_id"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type OsDiskState struct {
	Caching                     string                  `json:"caching"`
	DiskEncryptionSetId         string                  `json:"disk_encryption_set_id"`
	DiskSizeGb                  float64                 `json:"disk_size_gb"`
	Name                        string                  `json:"name"`
	SecureVmDiskEncryptionSetId string                  `json:"secure_vm_disk_encryption_set_id"`
	SecurityEncryptionType      string                  `json:"security_encryption_type"`
	StorageAccountType          string                  `json:"storage_account_type"`
	WriteAcceleratorEnabled     bool                    `json:"write_accelerator_enabled"`
	DiffDiskSettings            []DiffDiskSettingsState `json:"diff_disk_settings"`
}

type DiffDiskSettingsState struct {
	Option    string `json:"option"`
	Placement string `json:"placement"`
}

type PlanState struct {
	Name      string `json:"name"`
	Product   string `json:"product"`
	Publisher string `json:"publisher"`
}

type SecretState struct {
	KeyVaultId  string             `json:"key_vault_id"`
	Certificate []CertificateState `json:"certificate"`
}

type CertificateState struct {
	Url string `json:"url"`
}

type SourceImageReferenceState struct {
	Offer     string `json:"offer"`
	Publisher string `json:"publisher"`
	Sku       string `json:"sku"`
	Version   string `json:"version"`
}

type TerminationNotificationState struct {
	Enabled bool   `json:"enabled"`
	Timeout string `json:"timeout"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
