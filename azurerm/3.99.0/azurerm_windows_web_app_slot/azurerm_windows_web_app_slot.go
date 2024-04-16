// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_windows_web_app_slot

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_windows_web_app_slot.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermWindowsWebAppSlotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (awwas *Resource) Type() string {
	return "azurerm_windows_web_app_slot"
}

// LocalName returns the local name for [Resource].
func (awwas *Resource) LocalName() string {
	return awwas.Name
}

// Configuration returns the configuration (args) for [Resource].
func (awwas *Resource) Configuration() interface{} {
	return awwas.Args
}

// DependOn is used for other resources to depend on [Resource].
func (awwas *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(awwas)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (awwas *Resource) Dependencies() terra.Dependencies {
	return awwas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (awwas *Resource) LifecycleManagement() *terra.Lifecycle {
	return awwas.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (awwas *Resource) Attributes() azurermWindowsWebAppSlotAttributes {
	return azurermWindowsWebAppSlotAttributes{ref: terra.ReferenceResource(awwas)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (awwas *Resource) ImportState(state io.Reader) error {
	awwas.state = &azurermWindowsWebAppSlotState{}
	if err := json.NewDecoder(state).Decode(awwas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", awwas.Type(), awwas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (awwas *Resource) State() (*azurermWindowsWebAppSlotState, bool) {
	return awwas.state, awwas.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (awwas *Resource) StateMust() *azurermWindowsWebAppSlotState {
	if awwas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", awwas.Type(), awwas.LocalName()))
	}
	return awwas.state
}

// Args contains the configurations for azurerm_windows_web_app_slot.
type Args struct {
	// AppServiceId: string, required
	AppServiceId terra.StringValue `hcl:"app_service_id,attr" validate:"required"`
	// AppSettings: map of string, optional
	AppSettings terra.MapValue[terra.StringValue] `hcl:"app_settings,attr"`
	// ClientAffinityEnabled: bool, optional
	ClientAffinityEnabled terra.BoolValue `hcl:"client_affinity_enabled,attr"`
	// ClientCertificateEnabled: bool, optional
	ClientCertificateEnabled terra.BoolValue `hcl:"client_certificate_enabled,attr"`
	// ClientCertificateExclusionPaths: string, optional
	ClientCertificateExclusionPaths terra.StringValue `hcl:"client_certificate_exclusion_paths,attr"`
	// ClientCertificateMode: string, optional
	ClientCertificateMode terra.StringValue `hcl:"client_certificate_mode,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// FtpPublishBasicAuthenticationEnabled: bool, optional
	FtpPublishBasicAuthenticationEnabled terra.BoolValue `hcl:"ftp_publish_basic_authentication_enabled,attr"`
	// HttpsOnly: bool, optional
	HttpsOnly terra.BoolValue `hcl:"https_only,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultReferenceIdentityId: string, optional
	KeyVaultReferenceIdentityId terra.StringValue `hcl:"key_vault_reference_identity_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ServicePlanId: string, optional
	ServicePlanId terra.StringValue `hcl:"service_plan_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VirtualNetworkSubnetId: string, optional
	VirtualNetworkSubnetId terra.StringValue `hcl:"virtual_network_subnet_id,attr"`
	// WebdeployPublishBasicAuthenticationEnabled: bool, optional
	WebdeployPublishBasicAuthenticationEnabled terra.BoolValue `hcl:"webdeploy_publish_basic_authentication_enabled,attr"`
	// ZipDeployFile: string, optional
	ZipDeployFile terra.StringValue `hcl:"zip_deploy_file,attr"`
	// AuthSettings: optional
	AuthSettings *AuthSettings `hcl:"auth_settings,block"`
	// AuthSettingsV2: optional
	AuthSettingsV2 *AuthSettingsV2 `hcl:"auth_settings_v2,block"`
	// Backup: optional
	Backup *Backup `hcl:"backup,block"`
	// ConnectionString: min=0
	ConnectionString []ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: optional
	Identity *Identity `hcl:"identity,block"`
	// Logs: optional
	Logs *Logs `hcl:"logs,block"`
	// SiteConfig: required
	SiteConfig *SiteConfig `hcl:"site_config,block" validate:"required"`
	// StorageAccount: min=0
	StorageAccount []StorageAccount `hcl:"storage_account,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermWindowsWebAppSlotAttributes struct {
	ref terra.Reference
}

// AppServiceId returns a reference to field app_service_id of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) AppServiceId() terra.StringValue {
	return terra.ReferenceAsString(awwas.ref.Append("app_service_id"))
}

// AppSettings returns a reference to field app_settings of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](awwas.ref.Append("app_settings"))
}

// ClientAffinityEnabled returns a reference to field client_affinity_enabled of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) ClientAffinityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(awwas.ref.Append("client_affinity_enabled"))
}

// ClientCertificateEnabled returns a reference to field client_certificate_enabled of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) ClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(awwas.ref.Append("client_certificate_enabled"))
}

// ClientCertificateExclusionPaths returns a reference to field client_certificate_exclusion_paths of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) ClientCertificateExclusionPaths() terra.StringValue {
	return terra.ReferenceAsString(awwas.ref.Append("client_certificate_exclusion_paths"))
}

// ClientCertificateMode returns a reference to field client_certificate_mode of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceAsString(awwas.ref.Append("client_certificate_mode"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(awwas.ref.Append("custom_domain_verification_id"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(awwas.ref.Append("default_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(awwas.ref.Append("enabled"))
}

// FtpPublishBasicAuthenticationEnabled returns a reference to field ftp_publish_basic_authentication_enabled of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) FtpPublishBasicAuthenticationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(awwas.ref.Append("ftp_publish_basic_authentication_enabled"))
}

// HostingEnvironmentId returns a reference to field hosting_environment_id of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) HostingEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(awwas.ref.Append("hosting_environment_id"))
}

// HttpsOnly returns a reference to field https_only of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(awwas.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(awwas.ref.Append("id"))
}

// KeyVaultReferenceIdentityId returns a reference to field key_vault_reference_identity_id of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) KeyVaultReferenceIdentityId() terra.StringValue {
	return terra.ReferenceAsString(awwas.ref.Append("key_vault_reference_identity_id"))
}

// Kind returns a reference to field kind of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(awwas.ref.Append("kind"))
}

// Name returns a reference to field name of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(awwas.ref.Append("name"))
}

// OutboundIpAddressList returns a reference to field outbound_ip_address_list of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) OutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](awwas.ref.Append("outbound_ip_address_list"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(awwas.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddressList returns a reference to field possible_outbound_ip_address_list of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) PossibleOutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](awwas.ref.Append("possible_outbound_ip_address_list"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(awwas.ref.Append("possible_outbound_ip_addresses"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(awwas.ref.Append("public_network_access_enabled"))
}

// ServicePlanId returns a reference to field service_plan_id of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) ServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(awwas.ref.Append("service_plan_id"))
}

// Tags returns a reference to field tags of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](awwas.ref.Append("tags"))
}

// VirtualNetworkSubnetId returns a reference to field virtual_network_subnet_id of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(awwas.ref.Append("virtual_network_subnet_id"))
}

// WebdeployPublishBasicAuthenticationEnabled returns a reference to field webdeploy_publish_basic_authentication_enabled of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) WebdeployPublishBasicAuthenticationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(awwas.ref.Append("webdeploy_publish_basic_authentication_enabled"))
}

// ZipDeployFile returns a reference to field zip_deploy_file of azurerm_windows_web_app_slot.
func (awwas azurermWindowsWebAppSlotAttributes) ZipDeployFile() terra.StringValue {
	return terra.ReferenceAsString(awwas.ref.Append("zip_deploy_file"))
}

func (awwas azurermWindowsWebAppSlotAttributes) SiteCredential() terra.ListValue[SiteCredentialAttributes] {
	return terra.ReferenceAsList[SiteCredentialAttributes](awwas.ref.Append("site_credential"))
}

func (awwas azurermWindowsWebAppSlotAttributes) AuthSettings() terra.ListValue[AuthSettingsAttributes] {
	return terra.ReferenceAsList[AuthSettingsAttributes](awwas.ref.Append("auth_settings"))
}

func (awwas azurermWindowsWebAppSlotAttributes) AuthSettingsV2() terra.ListValue[AuthSettingsV2Attributes] {
	return terra.ReferenceAsList[AuthSettingsV2Attributes](awwas.ref.Append("auth_settings_v2"))
}

func (awwas azurermWindowsWebAppSlotAttributes) Backup() terra.ListValue[BackupAttributes] {
	return terra.ReferenceAsList[BackupAttributes](awwas.ref.Append("backup"))
}

func (awwas azurermWindowsWebAppSlotAttributes) ConnectionString() terra.SetValue[ConnectionStringAttributes] {
	return terra.ReferenceAsSet[ConnectionStringAttributes](awwas.ref.Append("connection_string"))
}

func (awwas azurermWindowsWebAppSlotAttributes) Identity() terra.ListValue[IdentityAttributes] {
	return terra.ReferenceAsList[IdentityAttributes](awwas.ref.Append("identity"))
}

func (awwas azurermWindowsWebAppSlotAttributes) Logs() terra.ListValue[LogsAttributes] {
	return terra.ReferenceAsList[LogsAttributes](awwas.ref.Append("logs"))
}

func (awwas azurermWindowsWebAppSlotAttributes) SiteConfig() terra.ListValue[SiteConfigAttributes] {
	return terra.ReferenceAsList[SiteConfigAttributes](awwas.ref.Append("site_config"))
}

func (awwas azurermWindowsWebAppSlotAttributes) StorageAccount() terra.SetValue[StorageAccountAttributes] {
	return terra.ReferenceAsSet[StorageAccountAttributes](awwas.ref.Append("storage_account"))
}

func (awwas azurermWindowsWebAppSlotAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](awwas.ref.Append("timeouts"))
}

type azurermWindowsWebAppSlotState struct {
	AppServiceId                               string                  `json:"app_service_id"`
	AppSettings                                map[string]string       `json:"app_settings"`
	ClientAffinityEnabled                      bool                    `json:"client_affinity_enabled"`
	ClientCertificateEnabled                   bool                    `json:"client_certificate_enabled"`
	ClientCertificateExclusionPaths            string                  `json:"client_certificate_exclusion_paths"`
	ClientCertificateMode                      string                  `json:"client_certificate_mode"`
	CustomDomainVerificationId                 string                  `json:"custom_domain_verification_id"`
	DefaultHostname                            string                  `json:"default_hostname"`
	Enabled                                    bool                    `json:"enabled"`
	FtpPublishBasicAuthenticationEnabled       bool                    `json:"ftp_publish_basic_authentication_enabled"`
	HostingEnvironmentId                       string                  `json:"hosting_environment_id"`
	HttpsOnly                                  bool                    `json:"https_only"`
	Id                                         string                  `json:"id"`
	KeyVaultReferenceIdentityId                string                  `json:"key_vault_reference_identity_id"`
	Kind                                       string                  `json:"kind"`
	Name                                       string                  `json:"name"`
	OutboundIpAddressList                      []string                `json:"outbound_ip_address_list"`
	OutboundIpAddresses                        string                  `json:"outbound_ip_addresses"`
	PossibleOutboundIpAddressList              []string                `json:"possible_outbound_ip_address_list"`
	PossibleOutboundIpAddresses                string                  `json:"possible_outbound_ip_addresses"`
	PublicNetworkAccessEnabled                 bool                    `json:"public_network_access_enabled"`
	ServicePlanId                              string                  `json:"service_plan_id"`
	Tags                                       map[string]string       `json:"tags"`
	VirtualNetworkSubnetId                     string                  `json:"virtual_network_subnet_id"`
	WebdeployPublishBasicAuthenticationEnabled bool                    `json:"webdeploy_publish_basic_authentication_enabled"`
	ZipDeployFile                              string                  `json:"zip_deploy_file"`
	SiteCredential                             []SiteCredentialState   `json:"site_credential"`
	AuthSettings                               []AuthSettingsState     `json:"auth_settings"`
	AuthSettingsV2                             []AuthSettingsV2State   `json:"auth_settings_v2"`
	Backup                                     []BackupState           `json:"backup"`
	ConnectionString                           []ConnectionStringState `json:"connection_string"`
	Identity                                   []IdentityState         `json:"identity"`
	Logs                                       []LogsState             `json:"logs"`
	SiteConfig                                 []SiteConfigState       `json:"site_config"`
	StorageAccount                             []StorageAccountState   `json:"storage_account"`
	Timeouts                                   *TimeoutsState          `json:"timeouts"`
}
