// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	linuxwebapp "github.com/golingon/terraproviders/azurerm/3.63.0/linuxwebapp"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLinuxWebApp creates a new instance of [LinuxWebApp].
func NewLinuxWebApp(name string, args LinuxWebAppArgs) *LinuxWebApp {
	return &LinuxWebApp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LinuxWebApp)(nil)

// LinuxWebApp represents the Terraform resource azurerm_linux_web_app.
type LinuxWebApp struct {
	Name      string
	Args      LinuxWebAppArgs
	state     *linuxWebAppState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LinuxWebApp].
func (lwa *LinuxWebApp) Type() string {
	return "azurerm_linux_web_app"
}

// LocalName returns the local name for [LinuxWebApp].
func (lwa *LinuxWebApp) LocalName() string {
	return lwa.Name
}

// Configuration returns the configuration (args) for [LinuxWebApp].
func (lwa *LinuxWebApp) Configuration() interface{} {
	return lwa.Args
}

// DependOn is used for other resources to depend on [LinuxWebApp].
func (lwa *LinuxWebApp) DependOn() terra.Reference {
	return terra.ReferenceResource(lwa)
}

// Dependencies returns the list of resources [LinuxWebApp] depends_on.
func (lwa *LinuxWebApp) Dependencies() terra.Dependencies {
	return lwa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LinuxWebApp].
func (lwa *LinuxWebApp) LifecycleManagement() *terra.Lifecycle {
	return lwa.Lifecycle
}

// Attributes returns the attributes for [LinuxWebApp].
func (lwa *LinuxWebApp) Attributes() linuxWebAppAttributes {
	return linuxWebAppAttributes{ref: terra.ReferenceResource(lwa)}
}

// ImportState imports the given attribute values into [LinuxWebApp]'s state.
func (lwa *LinuxWebApp) ImportState(av io.Reader) error {
	lwa.state = &linuxWebAppState{}
	if err := json.NewDecoder(av).Decode(lwa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lwa.Type(), lwa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LinuxWebApp] has state.
func (lwa *LinuxWebApp) State() (*linuxWebAppState, bool) {
	return lwa.state, lwa.state != nil
}

// StateMust returns the state for [LinuxWebApp]. Panics if the state is nil.
func (lwa *LinuxWebApp) StateMust() *linuxWebAppState {
	if lwa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lwa.Type(), lwa.LocalName()))
	}
	return lwa.state
}

// LinuxWebAppArgs contains the configurations for azurerm_linux_web_app.
type LinuxWebAppArgs struct {
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
	SiteCredential []linuxwebapp.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// AuthSettings: optional
	AuthSettings *linuxwebapp.AuthSettings `hcl:"auth_settings,block"`
	// AuthSettingsV2: optional
	AuthSettingsV2 *linuxwebapp.AuthSettingsV2 `hcl:"auth_settings_v2,block"`
	// Backup: optional
	Backup *linuxwebapp.Backup `hcl:"backup,block"`
	// ConnectionString: min=0
	ConnectionString []linuxwebapp.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: optional
	Identity *linuxwebapp.Identity `hcl:"identity,block"`
	// Logs: optional
	Logs *linuxwebapp.Logs `hcl:"logs,block"`
	// SiteConfig: required
	SiteConfig *linuxwebapp.SiteConfig `hcl:"site_config,block" validate:"required"`
	// StickySettings: optional
	StickySettings *linuxwebapp.StickySettings `hcl:"sticky_settings,block"`
	// StorageAccount: min=0
	StorageAccount []linuxwebapp.StorageAccount `hcl:"storage_account,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *linuxwebapp.Timeouts `hcl:"timeouts,block"`
}
type linuxWebAppAttributes struct {
	ref terra.Reference
}

// AppSettings returns a reference to field app_settings of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lwa.ref.Append("app_settings"))
}

// ClientAffinityEnabled returns a reference to field client_affinity_enabled of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) ClientAffinityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lwa.ref.Append("client_affinity_enabled"))
}

// ClientCertificateEnabled returns a reference to field client_certificate_enabled of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) ClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lwa.ref.Append("client_certificate_enabled"))
}

// ClientCertificateExclusionPaths returns a reference to field client_certificate_exclusion_paths of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) ClientCertificateExclusionPaths() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("client_certificate_exclusion_paths"))
}

// ClientCertificateMode returns a reference to field client_certificate_mode of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("client_certificate_mode"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("custom_domain_verification_id"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("default_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(lwa.ref.Append("enabled"))
}

// HostingEnvironmentId returns a reference to field hosting_environment_id of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) HostingEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("hosting_environment_id"))
}

// HttpsOnly returns a reference to field https_only of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(lwa.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("id"))
}

// KeyVaultReferenceIdentityId returns a reference to field key_vault_reference_identity_id of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) KeyVaultReferenceIdentityId() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("key_vault_reference_identity_id"))
}

// Kind returns a reference to field kind of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("name"))
}

// OutboundIpAddressList returns a reference to field outbound_ip_address_list of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) OutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lwa.ref.Append("outbound_ip_address_list"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddressList returns a reference to field possible_outbound_ip_address_list of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) PossibleOutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lwa.ref.Append("possible_outbound_ip_address_list"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("possible_outbound_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("resource_group_name"))
}

// ServicePlanId returns a reference to field service_plan_id of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) ServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("service_plan_id"))
}

// Tags returns a reference to field tags of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lwa.ref.Append("tags"))
}

// VirtualNetworkSubnetId returns a reference to field virtual_network_subnet_id of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("virtual_network_subnet_id"))
}

// ZipDeployFile returns a reference to field zip_deploy_file of azurerm_linux_web_app.
func (lwa linuxWebAppAttributes) ZipDeployFile() terra.StringValue {
	return terra.ReferenceAsString(lwa.ref.Append("zip_deploy_file"))
}

func (lwa linuxWebAppAttributes) SiteCredential() terra.ListValue[linuxwebapp.SiteCredentialAttributes] {
	return terra.ReferenceAsList[linuxwebapp.SiteCredentialAttributes](lwa.ref.Append("site_credential"))
}

func (lwa linuxWebAppAttributes) AuthSettings() terra.ListValue[linuxwebapp.AuthSettingsAttributes] {
	return terra.ReferenceAsList[linuxwebapp.AuthSettingsAttributes](lwa.ref.Append("auth_settings"))
}

func (lwa linuxWebAppAttributes) AuthSettingsV2() terra.ListValue[linuxwebapp.AuthSettingsV2Attributes] {
	return terra.ReferenceAsList[linuxwebapp.AuthSettingsV2Attributes](lwa.ref.Append("auth_settings_v2"))
}

func (lwa linuxWebAppAttributes) Backup() terra.ListValue[linuxwebapp.BackupAttributes] {
	return terra.ReferenceAsList[linuxwebapp.BackupAttributes](lwa.ref.Append("backup"))
}

func (lwa linuxWebAppAttributes) ConnectionString() terra.SetValue[linuxwebapp.ConnectionStringAttributes] {
	return terra.ReferenceAsSet[linuxwebapp.ConnectionStringAttributes](lwa.ref.Append("connection_string"))
}

func (lwa linuxWebAppAttributes) Identity() terra.ListValue[linuxwebapp.IdentityAttributes] {
	return terra.ReferenceAsList[linuxwebapp.IdentityAttributes](lwa.ref.Append("identity"))
}

func (lwa linuxWebAppAttributes) Logs() terra.ListValue[linuxwebapp.LogsAttributes] {
	return terra.ReferenceAsList[linuxwebapp.LogsAttributes](lwa.ref.Append("logs"))
}

func (lwa linuxWebAppAttributes) SiteConfig() terra.ListValue[linuxwebapp.SiteConfigAttributes] {
	return terra.ReferenceAsList[linuxwebapp.SiteConfigAttributes](lwa.ref.Append("site_config"))
}

func (lwa linuxWebAppAttributes) StickySettings() terra.ListValue[linuxwebapp.StickySettingsAttributes] {
	return terra.ReferenceAsList[linuxwebapp.StickySettingsAttributes](lwa.ref.Append("sticky_settings"))
}

func (lwa linuxWebAppAttributes) StorageAccount() terra.SetValue[linuxwebapp.StorageAccountAttributes] {
	return terra.ReferenceAsSet[linuxwebapp.StorageAccountAttributes](lwa.ref.Append("storage_account"))
}

func (lwa linuxWebAppAttributes) Timeouts() linuxwebapp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[linuxwebapp.TimeoutsAttributes](lwa.ref.Append("timeouts"))
}

type linuxWebAppState struct {
	AppSettings                     map[string]string                   `json:"app_settings"`
	ClientAffinityEnabled           bool                                `json:"client_affinity_enabled"`
	ClientCertificateEnabled        bool                                `json:"client_certificate_enabled"`
	ClientCertificateExclusionPaths string                              `json:"client_certificate_exclusion_paths"`
	ClientCertificateMode           string                              `json:"client_certificate_mode"`
	CustomDomainVerificationId      string                              `json:"custom_domain_verification_id"`
	DefaultHostname                 string                              `json:"default_hostname"`
	Enabled                         bool                                `json:"enabled"`
	HostingEnvironmentId            string                              `json:"hosting_environment_id"`
	HttpsOnly                       bool                                `json:"https_only"`
	Id                              string                              `json:"id"`
	KeyVaultReferenceIdentityId     string                              `json:"key_vault_reference_identity_id"`
	Kind                            string                              `json:"kind"`
	Location                        string                              `json:"location"`
	Name                            string                              `json:"name"`
	OutboundIpAddressList           []string                            `json:"outbound_ip_address_list"`
	OutboundIpAddresses             string                              `json:"outbound_ip_addresses"`
	PossibleOutboundIpAddressList   []string                            `json:"possible_outbound_ip_address_list"`
	PossibleOutboundIpAddresses     string                              `json:"possible_outbound_ip_addresses"`
	ResourceGroupName               string                              `json:"resource_group_name"`
	ServicePlanId                   string                              `json:"service_plan_id"`
	Tags                            map[string]string                   `json:"tags"`
	VirtualNetworkSubnetId          string                              `json:"virtual_network_subnet_id"`
	ZipDeployFile                   string                              `json:"zip_deploy_file"`
	SiteCredential                  []linuxwebapp.SiteCredentialState   `json:"site_credential"`
	AuthSettings                    []linuxwebapp.AuthSettingsState     `json:"auth_settings"`
	AuthSettingsV2                  []linuxwebapp.AuthSettingsV2State   `json:"auth_settings_v2"`
	Backup                          []linuxwebapp.BackupState           `json:"backup"`
	ConnectionString                []linuxwebapp.ConnectionStringState `json:"connection_string"`
	Identity                        []linuxwebapp.IdentityState         `json:"identity"`
	Logs                            []linuxwebapp.LogsState             `json:"logs"`
	SiteConfig                      []linuxwebapp.SiteConfigState       `json:"site_config"`
	StickySettings                  []linuxwebapp.StickySettingsState   `json:"sticky_settings"`
	StorageAccount                  []linuxwebapp.StorageAccountState   `json:"storage_account"`
	Timeouts                        *linuxwebapp.TimeoutsState          `json:"timeouts"`
}