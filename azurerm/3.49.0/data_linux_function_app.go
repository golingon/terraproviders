// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datalinuxfunctionapp "github.com/golingon/terraproviders/azurerm/3.49.0/datalinuxfunctionapp"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataLinuxFunctionApp(name string, args DataLinuxFunctionAppArgs) *DataLinuxFunctionApp {
	return &DataLinuxFunctionApp{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLinuxFunctionApp)(nil)

type DataLinuxFunctionApp struct {
	Name string
	Args DataLinuxFunctionAppArgs
}

func (lfa *DataLinuxFunctionApp) DataSource() string {
	return "azurerm_linux_function_app"
}

func (lfa *DataLinuxFunctionApp) LocalName() string {
	return lfa.Name
}

func (lfa *DataLinuxFunctionApp) Configuration() interface{} {
	return lfa.Args
}

func (lfa *DataLinuxFunctionApp) Attributes() dataLinuxFunctionAppAttributes {
	return dataLinuxFunctionAppAttributes{ref: terra.ReferenceDataResource(lfa)}
}

type DataLinuxFunctionAppArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// AuthSettings: min=0
	AuthSettings []datalinuxfunctionapp.AuthSettings `hcl:"auth_settings,block" validate:"min=0"`
	// AuthSettingsV2: min=0
	AuthSettingsV2 []datalinuxfunctionapp.AuthSettingsV2 `hcl:"auth_settings_v2,block" validate:"min=0"`
	// Backup: min=0
	Backup []datalinuxfunctionapp.Backup `hcl:"backup,block" validate:"min=0"`
	// ConnectionString: min=0
	ConnectionString []datalinuxfunctionapp.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: min=0
	Identity []datalinuxfunctionapp.Identity `hcl:"identity,block" validate:"min=0"`
	// SiteConfig: min=0
	SiteConfig []datalinuxfunctionapp.SiteConfig `hcl:"site_config,block" validate:"min=0"`
	// SiteCredential: min=0
	SiteCredential []datalinuxfunctionapp.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// StickySettings: min=0
	StickySettings []datalinuxfunctionapp.StickySettings `hcl:"sticky_settings,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datalinuxfunctionapp.Timeouts `hcl:"timeouts,block"`
}
type dataLinuxFunctionAppAttributes struct {
	ref terra.Reference
}

func (lfa dataLinuxFunctionAppAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](lfa.ref.Append("app_settings"))
}

func (lfa dataLinuxFunctionAppAttributes) BuiltinLoggingEnabled() terra.BoolValue {
	return terra.ReferenceBool(lfa.ref.Append("builtin_logging_enabled"))
}

func (lfa dataLinuxFunctionAppAttributes) ClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceBool(lfa.ref.Append("client_certificate_enabled"))
}

func (lfa dataLinuxFunctionAppAttributes) ClientCertificateExclusionPaths() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("client_certificate_exclusion_paths"))
}

func (lfa dataLinuxFunctionAppAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("client_certificate_mode"))
}

func (lfa dataLinuxFunctionAppAttributes) ContentShareForceDisabled() terra.BoolValue {
	return terra.ReferenceBool(lfa.ref.Append("content_share_force_disabled"))
}

func (lfa dataLinuxFunctionAppAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("custom_domain_verification_id"))
}

func (lfa dataLinuxFunctionAppAttributes) DailyMemoryTimeQuota() terra.NumberValue {
	return terra.ReferenceNumber(lfa.ref.Append("daily_memory_time_quota"))
}

func (lfa dataLinuxFunctionAppAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("default_hostname"))
}

func (lfa dataLinuxFunctionAppAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(lfa.ref.Append("enabled"))
}

func (lfa dataLinuxFunctionAppAttributes) FunctionsExtensionVersion() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("functions_extension_version"))
}

func (lfa dataLinuxFunctionAppAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceBool(lfa.ref.Append("https_only"))
}

func (lfa dataLinuxFunctionAppAttributes) Id() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("id"))
}

func (lfa dataLinuxFunctionAppAttributes) Kind() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("kind"))
}

func (lfa dataLinuxFunctionAppAttributes) Location() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("location"))
}

func (lfa dataLinuxFunctionAppAttributes) Name() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("name"))
}

func (lfa dataLinuxFunctionAppAttributes) OutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](lfa.ref.Append("outbound_ip_address_list"))
}

func (lfa dataLinuxFunctionAppAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("outbound_ip_addresses"))
}

func (lfa dataLinuxFunctionAppAttributes) PossibleOutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](lfa.ref.Append("possible_outbound_ip_address_list"))
}

func (lfa dataLinuxFunctionAppAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("possible_outbound_ip_addresses"))
}

func (lfa dataLinuxFunctionAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("resource_group_name"))
}

func (lfa dataLinuxFunctionAppAttributes) ServicePlanId() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("service_plan_id"))
}

func (lfa dataLinuxFunctionAppAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("storage_account_access_key"))
}

func (lfa dataLinuxFunctionAppAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("storage_account_name"))
}

func (lfa dataLinuxFunctionAppAttributes) StorageKeyVaultSecretId() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("storage_key_vault_secret_id"))
}

func (lfa dataLinuxFunctionAppAttributes) StorageUsesManagedIdentity() terra.BoolValue {
	return terra.ReferenceBool(lfa.ref.Append("storage_uses_managed_identity"))
}

func (lfa dataLinuxFunctionAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](lfa.ref.Append("tags"))
}

func (lfa dataLinuxFunctionAppAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceString(lfa.ref.Append("virtual_network_subnet_id"))
}

func (lfa dataLinuxFunctionAppAttributes) AuthSettings() terra.ListValue[datalinuxfunctionapp.AuthSettingsAttributes] {
	return terra.ReferenceList[datalinuxfunctionapp.AuthSettingsAttributes](lfa.ref.Append("auth_settings"))
}

func (lfa dataLinuxFunctionAppAttributes) AuthSettingsV2() terra.ListValue[datalinuxfunctionapp.AuthSettingsV2Attributes] {
	return terra.ReferenceList[datalinuxfunctionapp.AuthSettingsV2Attributes](lfa.ref.Append("auth_settings_v2"))
}

func (lfa dataLinuxFunctionAppAttributes) Backup() terra.ListValue[datalinuxfunctionapp.BackupAttributes] {
	return terra.ReferenceList[datalinuxfunctionapp.BackupAttributes](lfa.ref.Append("backup"))
}

func (lfa dataLinuxFunctionAppAttributes) ConnectionString() terra.SetValue[datalinuxfunctionapp.ConnectionStringAttributes] {
	return terra.ReferenceSet[datalinuxfunctionapp.ConnectionStringAttributes](lfa.ref.Append("connection_string"))
}

func (lfa dataLinuxFunctionAppAttributes) Identity() terra.ListValue[datalinuxfunctionapp.IdentityAttributes] {
	return terra.ReferenceList[datalinuxfunctionapp.IdentityAttributes](lfa.ref.Append("identity"))
}

func (lfa dataLinuxFunctionAppAttributes) SiteConfig() terra.ListValue[datalinuxfunctionapp.SiteConfigAttributes] {
	return terra.ReferenceList[datalinuxfunctionapp.SiteConfigAttributes](lfa.ref.Append("site_config"))
}

func (lfa dataLinuxFunctionAppAttributes) SiteCredential() terra.ListValue[datalinuxfunctionapp.SiteCredentialAttributes] {
	return terra.ReferenceList[datalinuxfunctionapp.SiteCredentialAttributes](lfa.ref.Append("site_credential"))
}

func (lfa dataLinuxFunctionAppAttributes) StickySettings() terra.ListValue[datalinuxfunctionapp.StickySettingsAttributes] {
	return terra.ReferenceList[datalinuxfunctionapp.StickySettingsAttributes](lfa.ref.Append("sticky_settings"))
}

func (lfa dataLinuxFunctionAppAttributes) Timeouts() datalinuxfunctionapp.TimeoutsAttributes {
	return terra.ReferenceSingle[datalinuxfunctionapp.TimeoutsAttributes](lfa.ref.Append("timeouts"))
}
