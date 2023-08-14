// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datawindowsfunctionapp "github.com/golingon/terraproviders/azurerm/3.66.0/datawindowsfunctionapp"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataWindowsFunctionApp creates a new instance of [DataWindowsFunctionApp].
func NewDataWindowsFunctionApp(name string, args DataWindowsFunctionAppArgs) *DataWindowsFunctionApp {
	return &DataWindowsFunctionApp{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataWindowsFunctionApp)(nil)

// DataWindowsFunctionApp represents the Terraform data resource azurerm_windows_function_app.
type DataWindowsFunctionApp struct {
	Name string
	Args DataWindowsFunctionAppArgs
}

// DataSource returns the Terraform object type for [DataWindowsFunctionApp].
func (wfa *DataWindowsFunctionApp) DataSource() string {
	return "azurerm_windows_function_app"
}

// LocalName returns the local name for [DataWindowsFunctionApp].
func (wfa *DataWindowsFunctionApp) LocalName() string {
	return wfa.Name
}

// Configuration returns the configuration (args) for [DataWindowsFunctionApp].
func (wfa *DataWindowsFunctionApp) Configuration() interface{} {
	return wfa.Args
}

// Attributes returns the attributes for [DataWindowsFunctionApp].
func (wfa *DataWindowsFunctionApp) Attributes() dataWindowsFunctionAppAttributes {
	return dataWindowsFunctionAppAttributes{ref: terra.ReferenceDataResource(wfa)}
}

// DataWindowsFunctionAppArgs contains the configurations for azurerm_windows_function_app.
type DataWindowsFunctionAppArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// AuthSettings: min=0
	AuthSettings []datawindowsfunctionapp.AuthSettings `hcl:"auth_settings,block" validate:"min=0"`
	// AuthSettingsV2: min=0
	AuthSettingsV2 []datawindowsfunctionapp.AuthSettingsV2 `hcl:"auth_settings_v2,block" validate:"min=0"`
	// Backup: min=0
	Backup []datawindowsfunctionapp.Backup `hcl:"backup,block" validate:"min=0"`
	// ConnectionString: min=0
	ConnectionString []datawindowsfunctionapp.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: min=0
	Identity []datawindowsfunctionapp.Identity `hcl:"identity,block" validate:"min=0"`
	// SiteConfig: min=0
	SiteConfig []datawindowsfunctionapp.SiteConfig `hcl:"site_config,block" validate:"min=0"`
	// SiteCredential: min=0
	SiteCredential []datawindowsfunctionapp.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// StickySettings: min=0
	StickySettings []datawindowsfunctionapp.StickySettings `hcl:"sticky_settings,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datawindowsfunctionapp.Timeouts `hcl:"timeouts,block"`
}
type dataWindowsFunctionAppAttributes struct {
	ref terra.Reference
}

// AppSettings returns a reference to field app_settings of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wfa.ref.Append("app_settings"))
}

// BuiltinLoggingEnabled returns a reference to field builtin_logging_enabled of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) BuiltinLoggingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wfa.ref.Append("builtin_logging_enabled"))
}

// ClientCertificateEnabled returns a reference to field client_certificate_enabled of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) ClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wfa.ref.Append("client_certificate_enabled"))
}

// ClientCertificateExclusionPaths returns a reference to field client_certificate_exclusion_paths of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) ClientCertificateExclusionPaths() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("client_certificate_exclusion_paths"))
}

// ClientCertificateMode returns a reference to field client_certificate_mode of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("client_certificate_mode"))
}

// ContentShareForceDisabled returns a reference to field content_share_force_disabled of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) ContentShareForceDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(wfa.ref.Append("content_share_force_disabled"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("custom_domain_verification_id"))
}

// DailyMemoryTimeQuota returns a reference to field daily_memory_time_quota of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) DailyMemoryTimeQuota() terra.NumberValue {
	return terra.ReferenceAsNumber(wfa.ref.Append("daily_memory_time_quota"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("default_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(wfa.ref.Append("enabled"))
}

// FunctionsExtensionVersion returns a reference to field functions_extension_version of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) FunctionsExtensionVersion() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("functions_extension_version"))
}

// HostingEnvironmentId returns a reference to field hosting_environment_id of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) HostingEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("hosting_environment_id"))
}

// HttpsOnly returns a reference to field https_only of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(wfa.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("name"))
}

// OutboundIpAddressList returns a reference to field outbound_ip_address_list of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) OutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wfa.ref.Append("outbound_ip_address_list"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddressList returns a reference to field possible_outbound_ip_address_list of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) PossibleOutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wfa.ref.Append("possible_outbound_ip_address_list"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("possible_outbound_ip_addresses"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wfa.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("resource_group_name"))
}

// ServicePlanId returns a reference to field service_plan_id of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) ServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("service_plan_id"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("storage_account_access_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("storage_account_name"))
}

// StorageKeyVaultSecretId returns a reference to field storage_key_vault_secret_id of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) StorageKeyVaultSecretId() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("storage_key_vault_secret_id"))
}

// StorageUsesManagedIdentity returns a reference to field storage_uses_managed_identity of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) StorageUsesManagedIdentity() terra.BoolValue {
	return terra.ReferenceAsBool(wfa.ref.Append("storage_uses_managed_identity"))
}

// Tags returns a reference to field tags of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wfa.ref.Append("tags"))
}

// VirtualNetworkSubnetId returns a reference to field virtual_network_subnet_id of azurerm_windows_function_app.
func (wfa dataWindowsFunctionAppAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("virtual_network_subnet_id"))
}

func (wfa dataWindowsFunctionAppAttributes) AuthSettings() terra.ListValue[datawindowsfunctionapp.AuthSettingsAttributes] {
	return terra.ReferenceAsList[datawindowsfunctionapp.AuthSettingsAttributes](wfa.ref.Append("auth_settings"))
}

func (wfa dataWindowsFunctionAppAttributes) AuthSettingsV2() terra.ListValue[datawindowsfunctionapp.AuthSettingsV2Attributes] {
	return terra.ReferenceAsList[datawindowsfunctionapp.AuthSettingsV2Attributes](wfa.ref.Append("auth_settings_v2"))
}

func (wfa dataWindowsFunctionAppAttributes) Backup() terra.ListValue[datawindowsfunctionapp.BackupAttributes] {
	return terra.ReferenceAsList[datawindowsfunctionapp.BackupAttributes](wfa.ref.Append("backup"))
}

func (wfa dataWindowsFunctionAppAttributes) ConnectionString() terra.SetValue[datawindowsfunctionapp.ConnectionStringAttributes] {
	return terra.ReferenceAsSet[datawindowsfunctionapp.ConnectionStringAttributes](wfa.ref.Append("connection_string"))
}

func (wfa dataWindowsFunctionAppAttributes) Identity() terra.ListValue[datawindowsfunctionapp.IdentityAttributes] {
	return terra.ReferenceAsList[datawindowsfunctionapp.IdentityAttributes](wfa.ref.Append("identity"))
}

func (wfa dataWindowsFunctionAppAttributes) SiteConfig() terra.ListValue[datawindowsfunctionapp.SiteConfigAttributes] {
	return terra.ReferenceAsList[datawindowsfunctionapp.SiteConfigAttributes](wfa.ref.Append("site_config"))
}

func (wfa dataWindowsFunctionAppAttributes) SiteCredential() terra.ListValue[datawindowsfunctionapp.SiteCredentialAttributes] {
	return terra.ReferenceAsList[datawindowsfunctionapp.SiteCredentialAttributes](wfa.ref.Append("site_credential"))
}

func (wfa dataWindowsFunctionAppAttributes) StickySettings() terra.ListValue[datawindowsfunctionapp.StickySettingsAttributes] {
	return terra.ReferenceAsList[datawindowsfunctionapp.StickySettingsAttributes](wfa.ref.Append("sticky_settings"))
}

func (wfa dataWindowsFunctionAppAttributes) Timeouts() datawindowsfunctionapp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datawindowsfunctionapp.TimeoutsAttributes](wfa.ref.Append("timeouts"))
}