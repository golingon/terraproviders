// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datawindowswebapp "github.com/golingon/terraproviders/azurerm/3.52.0/datawindowswebapp"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataWindowsWebApp creates a new instance of [DataWindowsWebApp].
func NewDataWindowsWebApp(name string, args DataWindowsWebAppArgs) *DataWindowsWebApp {
	return &DataWindowsWebApp{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataWindowsWebApp)(nil)

// DataWindowsWebApp represents the Terraform data resource azurerm_windows_web_app.
type DataWindowsWebApp struct {
	Name string
	Args DataWindowsWebAppArgs
}

// DataSource returns the Terraform object type for [DataWindowsWebApp].
func (wwa *DataWindowsWebApp) DataSource() string {
	return "azurerm_windows_web_app"
}

// LocalName returns the local name for [DataWindowsWebApp].
func (wwa *DataWindowsWebApp) LocalName() string {
	return wwa.Name
}

// Configuration returns the configuration (args) for [DataWindowsWebApp].
func (wwa *DataWindowsWebApp) Configuration() interface{} {
	return wwa.Args
}

// Attributes returns the attributes for [DataWindowsWebApp].
func (wwa *DataWindowsWebApp) Attributes() dataWindowsWebAppAttributes {
	return dataWindowsWebAppAttributes{ref: terra.ReferenceDataResource(wwa)}
}

// DataWindowsWebAppArgs contains the configurations for azurerm_windows_web_app.
type DataWindowsWebAppArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// AuthSettings: min=0
	AuthSettings []datawindowswebapp.AuthSettings `hcl:"auth_settings,block" validate:"min=0"`
	// AuthSettingsV2: min=0
	AuthSettingsV2 []datawindowswebapp.AuthSettingsV2 `hcl:"auth_settings_v2,block" validate:"min=0"`
	// Backup: min=0
	Backup []datawindowswebapp.Backup `hcl:"backup,block" validate:"min=0"`
	// ConnectionString: min=0
	ConnectionString []datawindowswebapp.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: min=0
	Identity []datawindowswebapp.Identity `hcl:"identity,block" validate:"min=0"`
	// Logs: min=0
	Logs []datawindowswebapp.Logs `hcl:"logs,block" validate:"min=0"`
	// SiteConfig: min=0
	SiteConfig []datawindowswebapp.SiteConfig `hcl:"site_config,block" validate:"min=0"`
	// SiteCredential: min=0
	SiteCredential []datawindowswebapp.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// StickySettings: min=0
	StickySettings []datawindowswebapp.StickySettings `hcl:"sticky_settings,block" validate:"min=0"`
	// StorageAccount: min=0
	StorageAccount []datawindowswebapp.StorageAccount `hcl:"storage_account,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datawindowswebapp.Timeouts `hcl:"timeouts,block"`
}
type dataWindowsWebAppAttributes struct {
	ref terra.Reference
}

// AppSettings returns a reference to field app_settings of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wwa.ref.Append("app_settings"))
}

// ClientAffinityEnabled returns a reference to field client_affinity_enabled of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) ClientAffinityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wwa.ref.Append("client_affinity_enabled"))
}

// ClientCertificateEnabled returns a reference to field client_certificate_enabled of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) ClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wwa.ref.Append("client_certificate_enabled"))
}

// ClientCertificateExclusionPaths returns a reference to field client_certificate_exclusion_paths of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) ClientCertificateExclusionPaths() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("client_certificate_exclusion_paths"))
}

// ClientCertificateMode returns a reference to field client_certificate_mode of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("client_certificate_mode"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("custom_domain_verification_id"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("default_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(wwa.ref.Append("enabled"))
}

// HttpsOnly returns a reference to field https_only of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(wwa.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("name"))
}

// OutboundIpAddressList returns a reference to field outbound_ip_address_list of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) OutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wwa.ref.Append("outbound_ip_address_list"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddressList returns a reference to field possible_outbound_ip_address_list of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) PossibleOutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wwa.ref.Append("possible_outbound_ip_address_list"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("possible_outbound_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("resource_group_name"))
}

// ServicePlanId returns a reference to field service_plan_id of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) ServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("service_plan_id"))
}

// Tags returns a reference to field tags of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wwa.ref.Append("tags"))
}

// VirtualNetworkSubnetId returns a reference to field virtual_network_subnet_id of azurerm_windows_web_app.
func (wwa dataWindowsWebAppAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("virtual_network_subnet_id"))
}

func (wwa dataWindowsWebAppAttributes) AuthSettings() terra.ListValue[datawindowswebapp.AuthSettingsAttributes] {
	return terra.ReferenceAsList[datawindowswebapp.AuthSettingsAttributes](wwa.ref.Append("auth_settings"))
}

func (wwa dataWindowsWebAppAttributes) AuthSettingsV2() terra.ListValue[datawindowswebapp.AuthSettingsV2Attributes] {
	return terra.ReferenceAsList[datawindowswebapp.AuthSettingsV2Attributes](wwa.ref.Append("auth_settings_v2"))
}

func (wwa dataWindowsWebAppAttributes) Backup() terra.ListValue[datawindowswebapp.BackupAttributes] {
	return terra.ReferenceAsList[datawindowswebapp.BackupAttributes](wwa.ref.Append("backup"))
}

func (wwa dataWindowsWebAppAttributes) ConnectionString() terra.SetValue[datawindowswebapp.ConnectionStringAttributes] {
	return terra.ReferenceAsSet[datawindowswebapp.ConnectionStringAttributes](wwa.ref.Append("connection_string"))
}

func (wwa dataWindowsWebAppAttributes) Identity() terra.ListValue[datawindowswebapp.IdentityAttributes] {
	return terra.ReferenceAsList[datawindowswebapp.IdentityAttributes](wwa.ref.Append("identity"))
}

func (wwa dataWindowsWebAppAttributes) Logs() terra.ListValue[datawindowswebapp.LogsAttributes] {
	return terra.ReferenceAsList[datawindowswebapp.LogsAttributes](wwa.ref.Append("logs"))
}

func (wwa dataWindowsWebAppAttributes) SiteConfig() terra.ListValue[datawindowswebapp.SiteConfigAttributes] {
	return terra.ReferenceAsList[datawindowswebapp.SiteConfigAttributes](wwa.ref.Append("site_config"))
}

func (wwa dataWindowsWebAppAttributes) SiteCredential() terra.ListValue[datawindowswebapp.SiteCredentialAttributes] {
	return terra.ReferenceAsList[datawindowswebapp.SiteCredentialAttributes](wwa.ref.Append("site_credential"))
}

func (wwa dataWindowsWebAppAttributes) StickySettings() terra.ListValue[datawindowswebapp.StickySettingsAttributes] {
	return terra.ReferenceAsList[datawindowswebapp.StickySettingsAttributes](wwa.ref.Append("sticky_settings"))
}

func (wwa dataWindowsWebAppAttributes) StorageAccount() terra.ListValue[datawindowswebapp.StorageAccountAttributes] {
	return terra.ReferenceAsList[datawindowswebapp.StorageAccountAttributes](wwa.ref.Append("storage_account"))
}

func (wwa dataWindowsWebAppAttributes) Timeouts() datawindowswebapp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datawindowswebapp.TimeoutsAttributes](wwa.ref.Append("timeouts"))
}
