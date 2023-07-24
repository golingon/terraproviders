// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datalinuxwebapp "github.com/golingon/terraproviders/azurerm/3.66.0/datalinuxwebapp"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLinuxWebApp creates a new instance of [DataLinuxWebApp].
func NewDataLinuxWebApp(name string, args DataLinuxWebAppArgs) *DataLinuxWebApp {
	return &DataLinuxWebApp{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLinuxWebApp)(nil)

// DataLinuxWebApp represents the Terraform data resource azurerm_linux_web_app.
type DataLinuxWebApp struct {
	Name string
	Args DataLinuxWebAppArgs
}

// DataSource returns the Terraform object type for [DataLinuxWebApp].
func (lwa *DataLinuxWebApp) DataSource() string {
	return "azurerm_linux_web_app"
}

// LocalName returns the local name for [DataLinuxWebApp].
func (lwa *DataLinuxWebApp) LocalName() string {
	return lwa.Name
}

// Configuration returns the configuration (args) for [DataLinuxWebApp].
func (lwa *DataLinuxWebApp) Configuration() interface{} {
	return lwa.Args
}

// Attributes returns the attributes for [DataLinuxWebApp].
func (lwa *DataLinuxWebApp) Attributes() dataLinuxWebAppAttributes {
	return dataLinuxWebAppAttributes{ref: terra.ReferenceDataResource(lwa)}
}

// DataLinuxWebAppArgs contains the configurations for azurerm_linux_web_app.
type DataLinuxWebAppArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// AuthSettings: min=0
	AuthSettings []datalinuxwebapp.AuthSettings `hcl:"auth_settings,block" validate:"min=0"`
	// AuthSettingsV2: min=0
	AuthSettingsV2 []datalinuxwebapp.AuthSettingsV2 `hcl:"auth_settings_v2,block" validate:"min=0"`
	// Backup: min=0
	Backup []datalinuxwebapp.Backup `hcl:"backup,block" validate:"min=0"`
	// ConnectionString: min=0
	ConnectionString []datalinuxwebapp.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: min=0
	Identity []datalinuxwebapp.Identity `hcl:"identity,block" validate:"min=0"`
	// Logs: min=0
	Logs []datalinuxwebapp.Logs `hcl:"logs,block" validate:"min=0"`
	// SiteConfig: min=0
	SiteConfig []datalinuxwebapp.SiteConfig `hcl:"site_config,block" validate:"min=0"`
	// SiteCredential: min=0
	SiteCredential []datalinuxwebapp.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// StickySettings: min=0
	StickySettings []datalinuxwebapp.StickySettings `hcl:"sticky_settings,block" validate:"min=0"`
	// StorageAccount: min=0
	StorageAccount []datalinuxwebapp.StorageAccount `hcl:"storage_account,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datalinuxwebapp.Timeouts `hcl:"timeouts,block"`
}
type dataLinuxWebAppAttributes struct {
	ref terra.Reference
}

// AppMetadata returns a reference to field app_metadata of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) AppMetadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lwa.ref.Append("app_metadata"))
}

// AppSettings returns a reference to field app_settings of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lwa.ref.Append("app_settings"))
}

// Availability returns a reference to field availability of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) Availability() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("availability"))
}

// ClientAffinityEnabled returns a reference to field client_affinity_enabled of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) ClientAffinityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lwa.ref.Append("client_affinity_enabled"))
}

// ClientCertificateEnabled returns a reference to field client_certificate_enabled of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) ClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lwa.ref.Append("client_certificate_enabled"))
}

// ClientCertificateExclusionPaths returns a reference to field client_certificate_exclusion_paths of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) ClientCertificateExclusionPaths() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("client_certificate_exclusion_paths"))
}

// ClientCertificateMode returns a reference to field client_certificate_mode of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("client_certificate_mode"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("custom_domain_verification_id"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("default_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(lwa.ref.Append("enabled"))
}

// HostingEnvironmentId returns a reference to field hosting_environment_id of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) HostingEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("hosting_environment_id"))
}

// HttpsOnly returns a reference to field https_only of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(lwa.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("id"))
}

// KeyVaultReferenceIdentityId returns a reference to field key_vault_reference_identity_id of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) KeyVaultReferenceIdentityId() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("key_vault_reference_identity_id"))
}

// Kind returns a reference to field kind of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("name"))
}

// OutboundIpAddressList returns a reference to field outbound_ip_address_list of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) OutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lwa.ref.Append("outbound_ip_address_list"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddressList returns a reference to field possible_outbound_ip_address_list of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) PossibleOutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lwa.ref.Append("possible_outbound_ip_address_list"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("possible_outbound_ip_addresses"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lwa.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("resource_group_name"))
}

// ServicePlanId returns a reference to field service_plan_id of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) ServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("service_plan_id"))
}

// Tags returns a reference to field tags of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lwa.ref.Append("tags"))
}

// Usage returns a reference to field usage of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) Usage() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("usage"))
}

// VirtualNetworkSubnetId returns a reference to field virtual_network_subnet_id of azurerm_linux_web_app.
func (lwa dataLinuxWebAppAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("virtual_network_subnet_id"))
}

func (lwa dataLinuxWebAppAttributes) AuthSettings() terra.ListValue[datalinuxwebapp.AuthSettingsAttributes] {
	return terra.ReferenceAsList[datalinuxwebapp.AuthSettingsAttributes](lwa.ref.Append("auth_settings"))
}

func (lwa dataLinuxWebAppAttributes) AuthSettingsV2() terra.ListValue[datalinuxwebapp.AuthSettingsV2Attributes] {
	return terra.ReferenceAsList[datalinuxwebapp.AuthSettingsV2Attributes](lwa.ref.Append("auth_settings_v2"))
}

func (lwa dataLinuxWebAppAttributes) Backup() terra.ListValue[datalinuxwebapp.BackupAttributes] {
	return terra.ReferenceAsList[datalinuxwebapp.BackupAttributes](lwa.ref.Append("backup"))
}

func (lwa dataLinuxWebAppAttributes) ConnectionString() terra.SetValue[datalinuxwebapp.ConnectionStringAttributes] {
	return terra.ReferenceAsSet[datalinuxwebapp.ConnectionStringAttributes](lwa.ref.Append("connection_string"))
}

func (lwa dataLinuxWebAppAttributes) Identity() terra.ListValue[datalinuxwebapp.IdentityAttributes] {
	return terra.ReferenceAsList[datalinuxwebapp.IdentityAttributes](lwa.ref.Append("identity"))
}

func (lwa dataLinuxWebAppAttributes) Logs() terra.ListValue[datalinuxwebapp.LogsAttributes] {
	return terra.ReferenceAsList[datalinuxwebapp.LogsAttributes](lwa.ref.Append("logs"))
}

func (lwa dataLinuxWebAppAttributes) SiteConfig() terra.ListValue[datalinuxwebapp.SiteConfigAttributes] {
	return terra.ReferenceAsList[datalinuxwebapp.SiteConfigAttributes](lwa.ref.Append("site_config"))
}

func (lwa dataLinuxWebAppAttributes) SiteCredential() terra.ListValue[datalinuxwebapp.SiteCredentialAttributes] {
	return terra.ReferenceAsList[datalinuxwebapp.SiteCredentialAttributes](lwa.ref.Append("site_credential"))
}

func (lwa dataLinuxWebAppAttributes) StickySettings() terra.ListValue[datalinuxwebapp.StickySettingsAttributes] {
	return terra.ReferenceAsList[datalinuxwebapp.StickySettingsAttributes](lwa.ref.Append("sticky_settings"))
}

func (lwa dataLinuxWebAppAttributes) StorageAccount() terra.ListValue[datalinuxwebapp.StorageAccountAttributes] {
	return terra.ReferenceAsList[datalinuxwebapp.StorageAccountAttributes](lwa.ref.Append("storage_account"))
}

func (lwa dataLinuxWebAppAttributes) Timeouts() datalinuxwebapp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datalinuxwebapp.TimeoutsAttributes](lwa.ref.Append("timeouts"))
}
