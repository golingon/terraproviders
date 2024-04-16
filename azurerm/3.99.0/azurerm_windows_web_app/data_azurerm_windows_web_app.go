// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_windows_web_app

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_windows_web_app.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (awwa *DataSource) DataSource() string {
	return "azurerm_windows_web_app"
}

// LocalName returns the local name for [DataSource].
func (awwa *DataSource) LocalName() string {
	return awwa.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (awwa *DataSource) Configuration() interface{} {
	return awwa.Args
}

// Attributes returns the attributes for [DataSource].
func (awwa *DataSource) Attributes() dataAzurermWindowsWebAppAttributes {
	return dataAzurermWindowsWebAppAttributes{ref: terra.ReferenceDataSource(awwa)}
}

// DataArgs contains the configurations for azurerm_windows_web_app.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermWindowsWebAppAttributes struct {
	ref terra.Reference
}

// AppSettings returns a reference to field app_settings of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](awwa.ref.Append("app_settings"))
}

// ClientAffinityEnabled returns a reference to field client_affinity_enabled of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) ClientAffinityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(awwa.ref.Append("client_affinity_enabled"))
}

// ClientCertificateEnabled returns a reference to field client_certificate_enabled of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) ClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(awwa.ref.Append("client_certificate_enabled"))
}

// ClientCertificateExclusionPaths returns a reference to field client_certificate_exclusion_paths of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) ClientCertificateExclusionPaths() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("client_certificate_exclusion_paths"))
}

// ClientCertificateMode returns a reference to field client_certificate_mode of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("client_certificate_mode"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("custom_domain_verification_id"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("default_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(awwa.ref.Append("enabled"))
}

// FtpPublishBasicAuthenticationEnabled returns a reference to field ftp_publish_basic_authentication_enabled of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) FtpPublishBasicAuthenticationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(awwa.ref.Append("ftp_publish_basic_authentication_enabled"))
}

// HostingEnvironmentId returns a reference to field hosting_environment_id of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) HostingEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("hosting_environment_id"))
}

// HttpsOnly returns a reference to field https_only of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(awwa.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("name"))
}

// OutboundIpAddressList returns a reference to field outbound_ip_address_list of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) OutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](awwa.ref.Append("outbound_ip_address_list"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddressList returns a reference to field possible_outbound_ip_address_list of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) PossibleOutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](awwa.ref.Append("possible_outbound_ip_address_list"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("possible_outbound_ip_addresses"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(awwa.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("resource_group_name"))
}

// ServicePlanId returns a reference to field service_plan_id of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) ServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("service_plan_id"))
}

// Tags returns a reference to field tags of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](awwa.ref.Append("tags"))
}

// VirtualNetworkSubnetId returns a reference to field virtual_network_subnet_id of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("virtual_network_subnet_id"))
}

// WebdeployPublishBasicAuthenticationEnabled returns a reference to field webdeploy_publish_basic_authentication_enabled of azurerm_windows_web_app.
func (awwa dataAzurermWindowsWebAppAttributes) WebdeployPublishBasicAuthenticationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(awwa.ref.Append("webdeploy_publish_basic_authentication_enabled"))
}

func (awwa dataAzurermWindowsWebAppAttributes) AuthSettings() terra.ListValue[DataAuthSettingsAttributes] {
	return terra.ReferenceAsList[DataAuthSettingsAttributes](awwa.ref.Append("auth_settings"))
}

func (awwa dataAzurermWindowsWebAppAttributes) AuthSettingsV2() terra.ListValue[DataAuthSettingsV2Attributes] {
	return terra.ReferenceAsList[DataAuthSettingsV2Attributes](awwa.ref.Append("auth_settings_v2"))
}

func (awwa dataAzurermWindowsWebAppAttributes) Backup() terra.ListValue[DataBackupAttributes] {
	return terra.ReferenceAsList[DataBackupAttributes](awwa.ref.Append("backup"))
}

func (awwa dataAzurermWindowsWebAppAttributes) ConnectionString() terra.SetValue[DataConnectionStringAttributes] {
	return terra.ReferenceAsSet[DataConnectionStringAttributes](awwa.ref.Append("connection_string"))
}

func (awwa dataAzurermWindowsWebAppAttributes) Identity() terra.ListValue[DataIdentityAttributes] {
	return terra.ReferenceAsList[DataIdentityAttributes](awwa.ref.Append("identity"))
}

func (awwa dataAzurermWindowsWebAppAttributes) Logs() terra.ListValue[DataLogsAttributes] {
	return terra.ReferenceAsList[DataLogsAttributes](awwa.ref.Append("logs"))
}

func (awwa dataAzurermWindowsWebAppAttributes) SiteConfig() terra.ListValue[DataSiteConfigAttributes] {
	return terra.ReferenceAsList[DataSiteConfigAttributes](awwa.ref.Append("site_config"))
}

func (awwa dataAzurermWindowsWebAppAttributes) SiteCredential() terra.ListValue[DataSiteCredentialAttributes] {
	return terra.ReferenceAsList[DataSiteCredentialAttributes](awwa.ref.Append("site_credential"))
}

func (awwa dataAzurermWindowsWebAppAttributes) StickySettings() terra.ListValue[DataStickySettingsAttributes] {
	return terra.ReferenceAsList[DataStickySettingsAttributes](awwa.ref.Append("sticky_settings"))
}

func (awwa dataAzurermWindowsWebAppAttributes) StorageAccount() terra.ListValue[DataStorageAccountAttributes] {
	return terra.ReferenceAsList[DataStorageAccountAttributes](awwa.ref.Append("storage_account"))
}

func (awwa dataAzurermWindowsWebAppAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](awwa.ref.Append("timeouts"))
}
