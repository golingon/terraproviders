// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datalinuxfunctionapp "github.com/golingon/terraproviders/azurerm/3.64.0/datalinuxfunctionapp"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLinuxFunctionApp creates a new instance of [DataLinuxFunctionApp].
func NewDataLinuxFunctionApp(name string, args DataLinuxFunctionAppArgs) *DataLinuxFunctionApp {
	return &DataLinuxFunctionApp{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLinuxFunctionApp)(nil)

// DataLinuxFunctionApp represents the Terraform data resource azurerm_linux_function_app.
type DataLinuxFunctionApp struct {
	Name string
	Args DataLinuxFunctionAppArgs
}

// DataSource returns the Terraform object type for [DataLinuxFunctionApp].
func (lfa *DataLinuxFunctionApp) DataSource() string {
	return "azurerm_linux_function_app"
}

// LocalName returns the local name for [DataLinuxFunctionApp].
func (lfa *DataLinuxFunctionApp) LocalName() string {
	return lfa.Name
}

// Configuration returns the configuration (args) for [DataLinuxFunctionApp].
func (lfa *DataLinuxFunctionApp) Configuration() interface{} {
	return lfa.Args
}

// Attributes returns the attributes for [DataLinuxFunctionApp].
func (lfa *DataLinuxFunctionApp) Attributes() dataLinuxFunctionAppAttributes {
	return dataLinuxFunctionAppAttributes{ref: terra.ReferenceDataResource(lfa)}
}

// DataLinuxFunctionAppArgs contains the configurations for azurerm_linux_function_app.
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

// AppSettings returns a reference to field app_settings of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lfa.ref.Append("app_settings"))
}

// Availability returns a reference to field availability of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) Availability() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("availability"))
}

// BuiltinLoggingEnabled returns a reference to field builtin_logging_enabled of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) BuiltinLoggingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfa.ref.Append("builtin_logging_enabled"))
}

// ClientCertificateEnabled returns a reference to field client_certificate_enabled of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) ClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfa.ref.Append("client_certificate_enabled"))
}

// ClientCertificateExclusionPaths returns a reference to field client_certificate_exclusion_paths of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) ClientCertificateExclusionPaths() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("client_certificate_exclusion_paths"))
}

// ClientCertificateMode returns a reference to field client_certificate_mode of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("client_certificate_mode"))
}

// ContentShareForceDisabled returns a reference to field content_share_force_disabled of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) ContentShareForceDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfa.ref.Append("content_share_force_disabled"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("custom_domain_verification_id"))
}

// DailyMemoryTimeQuota returns a reference to field daily_memory_time_quota of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) DailyMemoryTimeQuota() terra.NumberValue {
	return terra.ReferenceAsNumber(lfa.ref.Append("daily_memory_time_quota"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("default_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfa.ref.Append("enabled"))
}

// FunctionsExtensionVersion returns a reference to field functions_extension_version of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) FunctionsExtensionVersion() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("functions_extension_version"))
}

// HostingEnvironmentId returns a reference to field hosting_environment_id of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) HostingEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("hosting_environment_id"))
}

// HttpsOnly returns a reference to field https_only of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(lfa.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("name"))
}

// OutboundIpAddressList returns a reference to field outbound_ip_address_list of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) OutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lfa.ref.Append("outbound_ip_address_list"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddressList returns a reference to field possible_outbound_ip_address_list of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) PossibleOutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lfa.ref.Append("possible_outbound_ip_address_list"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("possible_outbound_ip_addresses"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfa.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("resource_group_name"))
}

// ServicePlanId returns a reference to field service_plan_id of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) ServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("service_plan_id"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("storage_account_access_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("storage_account_name"))
}

// StorageKeyVaultSecretId returns a reference to field storage_key_vault_secret_id of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) StorageKeyVaultSecretId() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("storage_key_vault_secret_id"))
}

// StorageUsesManagedIdentity returns a reference to field storage_uses_managed_identity of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) StorageUsesManagedIdentity() terra.BoolValue {
	return terra.ReferenceAsBool(lfa.ref.Append("storage_uses_managed_identity"))
}

// Tags returns a reference to field tags of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lfa.ref.Append("tags"))
}

// Usage returns a reference to field usage of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) Usage() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("usage"))
}

// VirtualNetworkSubnetId returns a reference to field virtual_network_subnet_id of azurerm_linux_function_app.
func (lfa dataLinuxFunctionAppAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("virtual_network_subnet_id"))
}

func (lfa dataLinuxFunctionAppAttributes) AuthSettings() terra.ListValue[datalinuxfunctionapp.AuthSettingsAttributes] {
	return terra.ReferenceAsList[datalinuxfunctionapp.AuthSettingsAttributes](lfa.ref.Append("auth_settings"))
}

func (lfa dataLinuxFunctionAppAttributes) AuthSettingsV2() terra.ListValue[datalinuxfunctionapp.AuthSettingsV2Attributes] {
	return terra.ReferenceAsList[datalinuxfunctionapp.AuthSettingsV2Attributes](lfa.ref.Append("auth_settings_v2"))
}

func (lfa dataLinuxFunctionAppAttributes) Backup() terra.ListValue[datalinuxfunctionapp.BackupAttributes] {
	return terra.ReferenceAsList[datalinuxfunctionapp.BackupAttributes](lfa.ref.Append("backup"))
}

func (lfa dataLinuxFunctionAppAttributes) ConnectionString() terra.SetValue[datalinuxfunctionapp.ConnectionStringAttributes] {
	return terra.ReferenceAsSet[datalinuxfunctionapp.ConnectionStringAttributes](lfa.ref.Append("connection_string"))
}

func (lfa dataLinuxFunctionAppAttributes) Identity() terra.ListValue[datalinuxfunctionapp.IdentityAttributes] {
	return terra.ReferenceAsList[datalinuxfunctionapp.IdentityAttributes](lfa.ref.Append("identity"))
}

func (lfa dataLinuxFunctionAppAttributes) SiteConfig() terra.ListValue[datalinuxfunctionapp.SiteConfigAttributes] {
	return terra.ReferenceAsList[datalinuxfunctionapp.SiteConfigAttributes](lfa.ref.Append("site_config"))
}

func (lfa dataLinuxFunctionAppAttributes) SiteCredential() terra.ListValue[datalinuxfunctionapp.SiteCredentialAttributes] {
	return terra.ReferenceAsList[datalinuxfunctionapp.SiteCredentialAttributes](lfa.ref.Append("site_credential"))
}

func (lfa dataLinuxFunctionAppAttributes) StickySettings() terra.ListValue[datalinuxfunctionapp.StickySettingsAttributes] {
	return terra.ReferenceAsList[datalinuxfunctionapp.StickySettingsAttributes](lfa.ref.Append("sticky_settings"))
}

func (lfa dataLinuxFunctionAppAttributes) Timeouts() datalinuxfunctionapp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datalinuxfunctionapp.TimeoutsAttributes](lfa.ref.Append("timeouts"))
}
