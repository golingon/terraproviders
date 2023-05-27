// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	windowswebapp "github.com/golingon/terraproviders/azurerm/3.58.0/windowswebapp"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWindowsWebApp creates a new instance of [WindowsWebApp].
func NewWindowsWebApp(name string, args WindowsWebAppArgs) *WindowsWebApp {
	return &WindowsWebApp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WindowsWebApp)(nil)

// WindowsWebApp represents the Terraform resource azurerm_windows_web_app.
type WindowsWebApp struct {
	Name      string
	Args      WindowsWebAppArgs
	state     *windowsWebAppState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WindowsWebApp].
func (wwa *WindowsWebApp) Type() string {
	return "azurerm_windows_web_app"
}

// LocalName returns the local name for [WindowsWebApp].
func (wwa *WindowsWebApp) LocalName() string {
	return wwa.Name
}

// Configuration returns the configuration (args) for [WindowsWebApp].
func (wwa *WindowsWebApp) Configuration() interface{} {
	return wwa.Args
}

// DependOn is used for other resources to depend on [WindowsWebApp].
func (wwa *WindowsWebApp) DependOn() terra.Reference {
	return terra.ReferenceResource(wwa)
}

// Dependencies returns the list of resources [WindowsWebApp] depends_on.
func (wwa *WindowsWebApp) Dependencies() terra.Dependencies {
	return wwa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WindowsWebApp].
func (wwa *WindowsWebApp) LifecycleManagement() *terra.Lifecycle {
	return wwa.Lifecycle
}

// Attributes returns the attributes for [WindowsWebApp].
func (wwa *WindowsWebApp) Attributes() windowsWebAppAttributes {
	return windowsWebAppAttributes{ref: terra.ReferenceResource(wwa)}
}

// ImportState imports the given attribute values into [WindowsWebApp]'s state.
func (wwa *WindowsWebApp) ImportState(av io.Reader) error {
	wwa.state = &windowsWebAppState{}
	if err := json.NewDecoder(av).Decode(wwa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wwa.Type(), wwa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WindowsWebApp] has state.
func (wwa *WindowsWebApp) State() (*windowsWebAppState, bool) {
	return wwa.state, wwa.state != nil
}

// StateMust returns the state for [WindowsWebApp]. Panics if the state is nil.
func (wwa *WindowsWebApp) StateMust() *windowsWebAppState {
	if wwa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wwa.Type(), wwa.LocalName()))
	}
	return wwa.state
}

// WindowsWebAppArgs contains the configurations for azurerm_windows_web_app.
type WindowsWebAppArgs struct {
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
	// HttpsOnly: bool, optional
	HttpsOnly terra.BoolValue `hcl:"https_only,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultReferenceIdentityId: string, optional
	KeyVaultReferenceIdentityId terra.StringValue `hcl:"key_vault_reference_identity_id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServicePlanId: string, required
	ServicePlanId terra.StringValue `hcl:"service_plan_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VirtualNetworkSubnetId: string, optional
	VirtualNetworkSubnetId terra.StringValue `hcl:"virtual_network_subnet_id,attr"`
	// ZipDeployFile: string, optional
	ZipDeployFile terra.StringValue `hcl:"zip_deploy_file,attr"`
	// SiteCredential: min=0
	SiteCredential []windowswebapp.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// AuthSettings: optional
	AuthSettings *windowswebapp.AuthSettings `hcl:"auth_settings,block"`
	// AuthSettingsV2: optional
	AuthSettingsV2 *windowswebapp.AuthSettingsV2 `hcl:"auth_settings_v2,block"`
	// Backup: optional
	Backup *windowswebapp.Backup `hcl:"backup,block"`
	// ConnectionString: min=0
	ConnectionString []windowswebapp.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: optional
	Identity *windowswebapp.Identity `hcl:"identity,block"`
	// Logs: optional
	Logs *windowswebapp.Logs `hcl:"logs,block"`
	// SiteConfig: required
	SiteConfig *windowswebapp.SiteConfig `hcl:"site_config,block" validate:"required"`
	// StickySettings: optional
	StickySettings *windowswebapp.StickySettings `hcl:"sticky_settings,block"`
	// StorageAccount: min=0
	StorageAccount []windowswebapp.StorageAccount `hcl:"storage_account,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *windowswebapp.Timeouts `hcl:"timeouts,block"`
}
type windowsWebAppAttributes struct {
	ref terra.Reference
}

// AppSettings returns a reference to field app_settings of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wwa.ref.Append("app_settings"))
}

// ClientAffinityEnabled returns a reference to field client_affinity_enabled of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) ClientAffinityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wwa.ref.Append("client_affinity_enabled"))
}

// ClientCertificateEnabled returns a reference to field client_certificate_enabled of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) ClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wwa.ref.Append("client_certificate_enabled"))
}

// ClientCertificateExclusionPaths returns a reference to field client_certificate_exclusion_paths of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) ClientCertificateExclusionPaths() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("client_certificate_exclusion_paths"))
}

// ClientCertificateMode returns a reference to field client_certificate_mode of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("client_certificate_mode"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("custom_domain_verification_id"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("default_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(wwa.ref.Append("enabled"))
}

// HostingEnvironmentId returns a reference to field hosting_environment_id of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) HostingEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("hosting_environment_id"))
}

// HttpsOnly returns a reference to field https_only of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(wwa.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("id"))
}

// KeyVaultReferenceIdentityId returns a reference to field key_vault_reference_identity_id of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) KeyVaultReferenceIdentityId() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("key_vault_reference_identity_id"))
}

// Kind returns a reference to field kind of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("name"))
}

// OutboundIpAddressList returns a reference to field outbound_ip_address_list of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) OutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wwa.ref.Append("outbound_ip_address_list"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddressList returns a reference to field possible_outbound_ip_address_list of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) PossibleOutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wwa.ref.Append("possible_outbound_ip_address_list"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("possible_outbound_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("resource_group_name"))
}

// ServicePlanId returns a reference to field service_plan_id of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) ServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("service_plan_id"))
}

// Tags returns a reference to field tags of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wwa.ref.Append("tags"))
}

// VirtualNetworkSubnetId returns a reference to field virtual_network_subnet_id of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("virtual_network_subnet_id"))
}

// ZipDeployFile returns a reference to field zip_deploy_file of azurerm_windows_web_app.
func (wwa windowsWebAppAttributes) ZipDeployFile() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("zip_deploy_file"))
}

func (wwa windowsWebAppAttributes) SiteCredential() terra.ListValue[windowswebapp.SiteCredentialAttributes] {
	return terra.ReferenceAsList[windowswebapp.SiteCredentialAttributes](wwa.ref.Append("site_credential"))
}

func (wwa windowsWebAppAttributes) AuthSettings() terra.ListValue[windowswebapp.AuthSettingsAttributes] {
	return terra.ReferenceAsList[windowswebapp.AuthSettingsAttributes](wwa.ref.Append("auth_settings"))
}

func (wwa windowsWebAppAttributes) AuthSettingsV2() terra.ListValue[windowswebapp.AuthSettingsV2Attributes] {
	return terra.ReferenceAsList[windowswebapp.AuthSettingsV2Attributes](wwa.ref.Append("auth_settings_v2"))
}

func (wwa windowsWebAppAttributes) Backup() terra.ListValue[windowswebapp.BackupAttributes] {
	return terra.ReferenceAsList[windowswebapp.BackupAttributes](wwa.ref.Append("backup"))
}

func (wwa windowsWebAppAttributes) ConnectionString() terra.SetValue[windowswebapp.ConnectionStringAttributes] {
	return terra.ReferenceAsSet[windowswebapp.ConnectionStringAttributes](wwa.ref.Append("connection_string"))
}

func (wwa windowsWebAppAttributes) Identity() terra.ListValue[windowswebapp.IdentityAttributes] {
	return terra.ReferenceAsList[windowswebapp.IdentityAttributes](wwa.ref.Append("identity"))
}

func (wwa windowsWebAppAttributes) Logs() terra.ListValue[windowswebapp.LogsAttributes] {
	return terra.ReferenceAsList[windowswebapp.LogsAttributes](wwa.ref.Append("logs"))
}

func (wwa windowsWebAppAttributes) SiteConfig() terra.ListValue[windowswebapp.SiteConfigAttributes] {
	return terra.ReferenceAsList[windowswebapp.SiteConfigAttributes](wwa.ref.Append("site_config"))
}

func (wwa windowsWebAppAttributes) StickySettings() terra.ListValue[windowswebapp.StickySettingsAttributes] {
	return terra.ReferenceAsList[windowswebapp.StickySettingsAttributes](wwa.ref.Append("sticky_settings"))
}

func (wwa windowsWebAppAttributes) StorageAccount() terra.SetValue[windowswebapp.StorageAccountAttributes] {
	return terra.ReferenceAsSet[windowswebapp.StorageAccountAttributes](wwa.ref.Append("storage_account"))
}

func (wwa windowsWebAppAttributes) Timeouts() windowswebapp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[windowswebapp.TimeoutsAttributes](wwa.ref.Append("timeouts"))
}

type windowsWebAppState struct {
	AppSettings                     map[string]string                     `json:"app_settings"`
	ClientAffinityEnabled           bool                                  `json:"client_affinity_enabled"`
	ClientCertificateEnabled        bool                                  `json:"client_certificate_enabled"`
	ClientCertificateExclusionPaths string                                `json:"client_certificate_exclusion_paths"`
	ClientCertificateMode           string                                `json:"client_certificate_mode"`
	CustomDomainVerificationId      string                                `json:"custom_domain_verification_id"`
	DefaultHostname                 string                                `json:"default_hostname"`
	Enabled                         bool                                  `json:"enabled"`
	HostingEnvironmentId            string                                `json:"hosting_environment_id"`
	HttpsOnly                       bool                                  `json:"https_only"`
	Id                              string                                `json:"id"`
	KeyVaultReferenceIdentityId     string                                `json:"key_vault_reference_identity_id"`
	Kind                            string                                `json:"kind"`
	Location                        string                                `json:"location"`
	Name                            string                                `json:"name"`
	OutboundIpAddressList           []string                              `json:"outbound_ip_address_list"`
	OutboundIpAddresses             string                                `json:"outbound_ip_addresses"`
	PossibleOutboundIpAddressList   []string                              `json:"possible_outbound_ip_address_list"`
	PossibleOutboundIpAddresses     string                                `json:"possible_outbound_ip_addresses"`
	ResourceGroupName               string                                `json:"resource_group_name"`
	ServicePlanId                   string                                `json:"service_plan_id"`
	Tags                            map[string]string                     `json:"tags"`
	VirtualNetworkSubnetId          string                                `json:"virtual_network_subnet_id"`
	ZipDeployFile                   string                                `json:"zip_deploy_file"`
	SiteCredential                  []windowswebapp.SiteCredentialState   `json:"site_credential"`
	AuthSettings                    []windowswebapp.AuthSettingsState     `json:"auth_settings"`
	AuthSettingsV2                  []windowswebapp.AuthSettingsV2State   `json:"auth_settings_v2"`
	Backup                          []windowswebapp.BackupState           `json:"backup"`
	ConnectionString                []windowswebapp.ConnectionStringState `json:"connection_string"`
	Identity                        []windowswebapp.IdentityState         `json:"identity"`
	Logs                            []windowswebapp.LogsState             `json:"logs"`
	SiteConfig                      []windowswebapp.SiteConfigState       `json:"site_config"`
	StickySettings                  []windowswebapp.StickySettingsState   `json:"sticky_settings"`
	StorageAccount                  []windowswebapp.StorageAccountState   `json:"storage_account"`
	Timeouts                        *windowswebapp.TimeoutsState          `json:"timeouts"`
}
