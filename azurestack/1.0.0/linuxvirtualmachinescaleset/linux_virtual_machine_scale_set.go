// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package linuxvirtualmachinescaleset

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

type AutomaticInstanceRepair struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// GracePeriod: string, optional
	GracePeriod terra.StringValue `hcl:"grace_period,attr"`
}

type AutomaticOsUpgradePolicy struct {
	// DisableAutomaticRollback: bool, required
	DisableAutomaticRollback terra.BoolValue `hcl:"disable_automatic_rollback,attr" validate:"required"`
	// EnableAutomaticOsUpgrade: bool, required
	EnableAutomaticOsUpgrade terra.BoolValue `hcl:"enable_automatic_os_upgrade,attr" validate:"required"`
}

type BootDiagnostics struct {
	// StorageAccountUri: string, required
	StorageAccountUri terra.StringValue `hcl:"storage_account_uri,attr" validate:"required"`
}

type DataDisk struct {
	// Caching: string, required
	Caching terra.StringValue `hcl:"caching,attr" validate:"required"`
	// CreateOption: string, optional
	CreateOption terra.StringValue `hcl:"create_option,attr"`
	// DiskEncryptionSetId: string, optional
	DiskEncryptionSetId terra.StringValue `hcl:"disk_encryption_set_id,attr"`
	// DiskSizeGb: number, required
	DiskSizeGb terra.NumberValue `hcl:"disk_size_gb,attr" validate:"required"`
	// Lun: number, required
	Lun terra.NumberValue `hcl:"lun,attr" validate:"required"`
	// StorageAccountType: string, required
	StorageAccountType terra.StringValue `hcl:"storage_account_type,attr" validate:"required"`
	// WriteAcceleratorEnabled: bool, optional
	WriteAcceleratorEnabled terra.BoolValue `hcl:"write_accelerator_enabled,attr"`
}

type Extension struct {
	// AutoUpgradeMinorVersion: bool, optional
	AutoUpgradeMinorVersion terra.BoolValue `hcl:"auto_upgrade_minor_version,attr"`
	// AutomaticUpgradeEnabled: bool, optional
	AutomaticUpgradeEnabled terra.BoolValue `hcl:"automatic_upgrade_enabled,attr"`
	// ForceUpdateTag: string, optional
	ForceUpdateTag terra.StringValue `hcl:"force_update_tag,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProtectedSettings: string, optional
	ProtectedSettings terra.StringValue `hcl:"protected_settings,attr"`
	// ProvisionAfterExtensions: list of string, optional
	ProvisionAfterExtensions terra.ListValue[terra.StringValue] `hcl:"provision_after_extensions,attr"`
	// Publisher: string, required
	Publisher terra.StringValue `hcl:"publisher,attr" validate:"required"`
	// Settings: string, optional
	Settings terra.StringValue `hcl:"settings,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// TypeHandlerVersion: string, required
	TypeHandlerVersion terra.StringValue `hcl:"type_handler_version,attr" validate:"required"`
}

type NetworkInterface struct {
	// DnsServers: list of string, optional
	DnsServers terra.ListValue[terra.StringValue] `hcl:"dns_servers,attr"`
	// EnableIpForwarding: bool, optional
	EnableIpForwarding terra.BoolValue `hcl:"enable_ip_forwarding,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkSecurityGroupId: string, optional
	NetworkSecurityGroupId terra.StringValue `hcl:"network_security_group_id,attr"`
	// Primary: bool, optional
	Primary terra.BoolValue `hcl:"primary,attr"`
	// IpConfiguration: min=1
	IpConfiguration []IpConfiguration `hcl:"ip_configuration,block" validate:"min=1"`
}

type IpConfiguration struct {
	// LoadBalancerBackendAddressPoolIds: set of string, optional
	LoadBalancerBackendAddressPoolIds terra.SetValue[terra.StringValue] `hcl:"load_balancer_backend_address_pool_ids,attr"`
	// LoadBalancerInboundNatRulesIds: set of string, optional
	LoadBalancerInboundNatRulesIds terra.SetValue[terra.StringValue] `hcl:"load_balancer_inbound_nat_rules_ids,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Primary: bool, optional
	Primary terra.BoolValue `hcl:"primary,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
}

type OsDisk struct {
	// Caching: string, required
	Caching terra.StringValue `hcl:"caching,attr" validate:"required"`
	// DiskEncryptionSetId: string, optional
	DiskEncryptionSetId terra.StringValue `hcl:"disk_encryption_set_id,attr"`
	// DiskSizeGb: number, optional
	DiskSizeGb terra.NumberValue `hcl:"disk_size_gb,attr"`
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

type Certificate struct{}

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

type TerminateNotification struct {
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

type AutomaticInstanceRepairAttributes struct {
	ref terra.Reference
}

func (air AutomaticInstanceRepairAttributes) InternalRef() (terra.Reference, error) {
	return air.ref, nil
}

func (air AutomaticInstanceRepairAttributes) InternalWithRef(ref terra.Reference) AutomaticInstanceRepairAttributes {
	return AutomaticInstanceRepairAttributes{ref: ref}
}

func (air AutomaticInstanceRepairAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return air.ref.InternalTokens()
}

func (air AutomaticInstanceRepairAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(air.ref.Append("enabled"))
}

func (air AutomaticInstanceRepairAttributes) GracePeriod() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("grace_period"))
}

type AutomaticOsUpgradePolicyAttributes struct {
	ref terra.Reference
}

func (aoup AutomaticOsUpgradePolicyAttributes) InternalRef() (terra.Reference, error) {
	return aoup.ref, nil
}

func (aoup AutomaticOsUpgradePolicyAttributes) InternalWithRef(ref terra.Reference) AutomaticOsUpgradePolicyAttributes {
	return AutomaticOsUpgradePolicyAttributes{ref: ref}
}

func (aoup AutomaticOsUpgradePolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aoup.ref.InternalTokens()
}

func (aoup AutomaticOsUpgradePolicyAttributes) DisableAutomaticRollback() terra.BoolValue {
	return terra.ReferenceAsBool(aoup.ref.Append("disable_automatic_rollback"))
}

func (aoup AutomaticOsUpgradePolicyAttributes) EnableAutomaticOsUpgrade() terra.BoolValue {
	return terra.ReferenceAsBool(aoup.ref.Append("enable_automatic_os_upgrade"))
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

type DataDiskAttributes struct {
	ref terra.Reference
}

func (dd DataDiskAttributes) InternalRef() (terra.Reference, error) {
	return dd.ref, nil
}

func (dd DataDiskAttributes) InternalWithRef(ref terra.Reference) DataDiskAttributes {
	return DataDiskAttributes{ref: ref}
}

func (dd DataDiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dd.ref.InternalTokens()
}

func (dd DataDiskAttributes) Caching() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("caching"))
}

func (dd DataDiskAttributes) CreateOption() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("create_option"))
}

func (dd DataDiskAttributes) DiskEncryptionSetId() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("disk_encryption_set_id"))
}

func (dd DataDiskAttributes) DiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(dd.ref.Append("disk_size_gb"))
}

func (dd DataDiskAttributes) Lun() terra.NumberValue {
	return terra.ReferenceAsNumber(dd.ref.Append("lun"))
}

func (dd DataDiskAttributes) StorageAccountType() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("storage_account_type"))
}

func (dd DataDiskAttributes) WriteAcceleratorEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dd.ref.Append("write_accelerator_enabled"))
}

type ExtensionAttributes struct {
	ref terra.Reference
}

func (e ExtensionAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e ExtensionAttributes) InternalWithRef(ref terra.Reference) ExtensionAttributes {
	return ExtensionAttributes{ref: ref}
}

func (e ExtensionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e ExtensionAttributes) AutoUpgradeMinorVersion() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("auto_upgrade_minor_version"))
}

func (e ExtensionAttributes) AutomaticUpgradeEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("automatic_upgrade_enabled"))
}

func (e ExtensionAttributes) ForceUpdateTag() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("force_update_tag"))
}

func (e ExtensionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("name"))
}

func (e ExtensionAttributes) ProtectedSettings() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("protected_settings"))
}

func (e ExtensionAttributes) ProvisionAfterExtensions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](e.ref.Append("provision_after_extensions"))
}

func (e ExtensionAttributes) Publisher() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("publisher"))
}

func (e ExtensionAttributes) Settings() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("settings"))
}

func (e ExtensionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("type"))
}

func (e ExtensionAttributes) TypeHandlerVersion() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("type_handler_version"))
}

type NetworkInterfaceAttributes struct {
	ref terra.Reference
}

func (ni NetworkInterfaceAttributes) InternalRef() (terra.Reference, error) {
	return ni.ref, nil
}

func (ni NetworkInterfaceAttributes) InternalWithRef(ref terra.Reference) NetworkInterfaceAttributes {
	return NetworkInterfaceAttributes{ref: ref}
}

func (ni NetworkInterfaceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ni.ref.InternalTokens()
}

func (ni NetworkInterfaceAttributes) DnsServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ni.ref.Append("dns_servers"))
}

func (ni NetworkInterfaceAttributes) EnableIpForwarding() terra.BoolValue {
	return terra.ReferenceAsBool(ni.ref.Append("enable_ip_forwarding"))
}

func (ni NetworkInterfaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("name"))
}

func (ni NetworkInterfaceAttributes) NetworkSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("network_security_group_id"))
}

func (ni NetworkInterfaceAttributes) Primary() terra.BoolValue {
	return terra.ReferenceAsBool(ni.ref.Append("primary"))
}

func (ni NetworkInterfaceAttributes) IpConfiguration() terra.ListValue[IpConfigurationAttributes] {
	return terra.ReferenceAsList[IpConfigurationAttributes](ni.ref.Append("ip_configuration"))
}

type IpConfigurationAttributes struct {
	ref terra.Reference
}

func (ic IpConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ic.ref, nil
}

func (ic IpConfigurationAttributes) InternalWithRef(ref terra.Reference) IpConfigurationAttributes {
	return IpConfigurationAttributes{ref: ref}
}

func (ic IpConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ic.ref.InternalTokens()
}

func (ic IpConfigurationAttributes) LoadBalancerBackendAddressPoolIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ic.ref.Append("load_balancer_backend_address_pool_ids"))
}

func (ic IpConfigurationAttributes) LoadBalancerInboundNatRulesIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ic.ref.Append("load_balancer_inbound_nat_rules_ids"))
}

func (ic IpConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("name"))
}

func (ic IpConfigurationAttributes) Primary() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("primary"))
}

func (ic IpConfigurationAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("subnet_id"))
}

func (ic IpConfigurationAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("version"))
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

type TerminateNotificationAttributes struct {
	ref terra.Reference
}

func (tn TerminateNotificationAttributes) InternalRef() (terra.Reference, error) {
	return tn.ref, nil
}

func (tn TerminateNotificationAttributes) InternalWithRef(ref terra.Reference) TerminateNotificationAttributes {
	return TerminateNotificationAttributes{ref: ref}
}

func (tn TerminateNotificationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tn.ref.InternalTokens()
}

func (tn TerminateNotificationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(tn.ref.Append("enabled"))
}

func (tn TerminateNotificationAttributes) Timeout() terra.StringValue {
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

type AutomaticInstanceRepairState struct {
	Enabled     bool   `json:"enabled"`
	GracePeriod string `json:"grace_period"`
}

type AutomaticOsUpgradePolicyState struct {
	DisableAutomaticRollback bool `json:"disable_automatic_rollback"`
	EnableAutomaticOsUpgrade bool `json:"enable_automatic_os_upgrade"`
}

type BootDiagnosticsState struct {
	StorageAccountUri string `json:"storage_account_uri"`
}

type DataDiskState struct {
	Caching                 string  `json:"caching"`
	CreateOption            string  `json:"create_option"`
	DiskEncryptionSetId     string  `json:"disk_encryption_set_id"`
	DiskSizeGb              float64 `json:"disk_size_gb"`
	Lun                     float64 `json:"lun"`
	StorageAccountType      string  `json:"storage_account_type"`
	WriteAcceleratorEnabled bool    `json:"write_accelerator_enabled"`
}

type ExtensionState struct {
	AutoUpgradeMinorVersion  bool     `json:"auto_upgrade_minor_version"`
	AutomaticUpgradeEnabled  bool     `json:"automatic_upgrade_enabled"`
	ForceUpdateTag           string   `json:"force_update_tag"`
	Name                     string   `json:"name"`
	ProtectedSettings        string   `json:"protected_settings"`
	ProvisionAfterExtensions []string `json:"provision_after_extensions"`
	Publisher                string   `json:"publisher"`
	Settings                 string   `json:"settings"`
	Type                     string   `json:"type"`
	TypeHandlerVersion       string   `json:"type_handler_version"`
}

type NetworkInterfaceState struct {
	DnsServers             []string               `json:"dns_servers"`
	EnableIpForwarding     bool                   `json:"enable_ip_forwarding"`
	Name                   string                 `json:"name"`
	NetworkSecurityGroupId string                 `json:"network_security_group_id"`
	Primary                bool                   `json:"primary"`
	IpConfiguration        []IpConfigurationState `json:"ip_configuration"`
}

type IpConfigurationState struct {
	LoadBalancerBackendAddressPoolIds []string `json:"load_balancer_backend_address_pool_ids"`
	LoadBalancerInboundNatRulesIds    []string `json:"load_balancer_inbound_nat_rules_ids"`
	Name                              string   `json:"name"`
	Primary                           bool     `json:"primary"`
	SubnetId                          string   `json:"subnet_id"`
	Version                           string   `json:"version"`
}

type OsDiskState struct {
	Caching                 string                  `json:"caching"`
	DiskEncryptionSetId     string                  `json:"disk_encryption_set_id"`
	DiskSizeGb              float64                 `json:"disk_size_gb"`
	StorageAccountType      string                  `json:"storage_account_type"`
	WriteAcceleratorEnabled bool                    `json:"write_accelerator_enabled"`
	DiffDiskSettings        []DiffDiskSettingsState `json:"diff_disk_settings"`
}

type DiffDiskSettingsState struct {
	Option string `json:"option"`
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

type CertificateState struct{}

type SourceImageReferenceState struct {
	Offer     string `json:"offer"`
	Publisher string `json:"publisher"`
	Sku       string `json:"sku"`
	Version   string `json:"version"`
}

type TerminateNotificationState struct {
	Enabled bool   `json:"enabled"`
	Timeout string `json:"timeout"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
