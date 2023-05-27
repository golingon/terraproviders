// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	windowsfunctionapp "github.com/golingon/terraproviders/azurerm/3.58.0/windowsfunctionapp"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWindowsFunctionApp creates a new instance of [WindowsFunctionApp].
func NewWindowsFunctionApp(name string, args WindowsFunctionAppArgs) *WindowsFunctionApp {
	return &WindowsFunctionApp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WindowsFunctionApp)(nil)

// WindowsFunctionApp represents the Terraform resource azurerm_windows_function_app.
type WindowsFunctionApp struct {
	Name      string
	Args      WindowsFunctionAppArgs
	state     *windowsFunctionAppState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WindowsFunctionApp].
func (wfa *WindowsFunctionApp) Type() string {
	return "azurerm_windows_function_app"
}

// LocalName returns the local name for [WindowsFunctionApp].
func (wfa *WindowsFunctionApp) LocalName() string {
	return wfa.Name
}

// Configuration returns the configuration (args) for [WindowsFunctionApp].
func (wfa *WindowsFunctionApp) Configuration() interface{} {
	return wfa.Args
}

// DependOn is used for other resources to depend on [WindowsFunctionApp].
func (wfa *WindowsFunctionApp) DependOn() terra.Reference {
	return terra.ReferenceResource(wfa)
}

// Dependencies returns the list of resources [WindowsFunctionApp] depends_on.
func (wfa *WindowsFunctionApp) Dependencies() terra.Dependencies {
	return wfa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WindowsFunctionApp].
func (wfa *WindowsFunctionApp) LifecycleManagement() *terra.Lifecycle {
	return wfa.Lifecycle
}

// Attributes returns the attributes for [WindowsFunctionApp].
func (wfa *WindowsFunctionApp) Attributes() windowsFunctionAppAttributes {
	return windowsFunctionAppAttributes{ref: terra.ReferenceResource(wfa)}
}

// ImportState imports the given attribute values into [WindowsFunctionApp]'s state.
func (wfa *WindowsFunctionApp) ImportState(av io.Reader) error {
	wfa.state = &windowsFunctionAppState{}
	if err := json.NewDecoder(av).Decode(wfa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wfa.Type(), wfa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WindowsFunctionApp] has state.
func (wfa *WindowsFunctionApp) State() (*windowsFunctionAppState, bool) {
	return wfa.state, wfa.state != nil
}

// StateMust returns the state for [WindowsFunctionApp]. Panics if the state is nil.
func (wfa *WindowsFunctionApp) StateMust() *windowsFunctionAppState {
	if wfa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wfa.Type(), wfa.LocalName()))
	}
	return wfa.state
}

// WindowsFunctionAppArgs contains the configurations for azurerm_windows_function_app.
type WindowsFunctionAppArgs struct {
	// AppSettings: map of string, optional
	AppSettings terra.MapValue[terra.StringValue] `hcl:"app_settings,attr"`
	// BuiltinLoggingEnabled: bool, optional
	BuiltinLoggingEnabled terra.BoolValue `hcl:"builtin_logging_enabled,attr"`
	// ClientCertificateEnabled: bool, optional
	ClientCertificateEnabled terra.BoolValue `hcl:"client_certificate_enabled,attr"`
	// ClientCertificateExclusionPaths: string, optional
	ClientCertificateExclusionPaths terra.StringValue `hcl:"client_certificate_exclusion_paths,attr"`
	// ClientCertificateMode: string, optional
	ClientCertificateMode terra.StringValue `hcl:"client_certificate_mode,attr"`
	// ContentShareForceDisabled: bool, optional
	ContentShareForceDisabled terra.BoolValue `hcl:"content_share_force_disabled,attr"`
	// DailyMemoryTimeQuota: number, optional
	DailyMemoryTimeQuota terra.NumberValue `hcl:"daily_memory_time_quota,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// FunctionsExtensionVersion: string, optional
	FunctionsExtensionVersion terra.StringValue `hcl:"functions_extension_version,attr"`
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
	// StorageAccountAccessKey: string, optional
	StorageAccountAccessKey terra.StringValue `hcl:"storage_account_access_key,attr"`
	// StorageAccountName: string, optional
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr"`
	// StorageKeyVaultSecretId: string, optional
	StorageKeyVaultSecretId terra.StringValue `hcl:"storage_key_vault_secret_id,attr"`
	// StorageUsesManagedIdentity: bool, optional
	StorageUsesManagedIdentity terra.BoolValue `hcl:"storage_uses_managed_identity,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VirtualNetworkSubnetId: string, optional
	VirtualNetworkSubnetId terra.StringValue `hcl:"virtual_network_subnet_id,attr"`
	// ZipDeployFile: string, optional
	ZipDeployFile terra.StringValue `hcl:"zip_deploy_file,attr"`
	// SiteCredential: min=0
	SiteCredential []windowsfunctionapp.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// AuthSettings: optional
	AuthSettings *windowsfunctionapp.AuthSettings `hcl:"auth_settings,block"`
	// AuthSettingsV2: optional
	AuthSettingsV2 *windowsfunctionapp.AuthSettingsV2 `hcl:"auth_settings_v2,block"`
	// Backup: optional
	Backup *windowsfunctionapp.Backup `hcl:"backup,block"`
	// ConnectionString: min=0
	ConnectionString []windowsfunctionapp.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: optional
	Identity *windowsfunctionapp.Identity `hcl:"identity,block"`
	// SiteConfig: required
	SiteConfig *windowsfunctionapp.SiteConfig `hcl:"site_config,block" validate:"required"`
	// StickySettings: optional
	StickySettings *windowsfunctionapp.StickySettings `hcl:"sticky_settings,block"`
	// StorageAccount: min=0
	StorageAccount []windowsfunctionapp.StorageAccount `hcl:"storage_account,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *windowsfunctionapp.Timeouts `hcl:"timeouts,block"`
}
type windowsFunctionAppAttributes struct {
	ref terra.Reference
}

// AppSettings returns a reference to field app_settings of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wfa.ref.Append("app_settings"))
}

// BuiltinLoggingEnabled returns a reference to field builtin_logging_enabled of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) BuiltinLoggingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wfa.ref.Append("builtin_logging_enabled"))
}

// ClientCertificateEnabled returns a reference to field client_certificate_enabled of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) ClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wfa.ref.Append("client_certificate_enabled"))
}

// ClientCertificateExclusionPaths returns a reference to field client_certificate_exclusion_paths of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) ClientCertificateExclusionPaths() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("client_certificate_exclusion_paths"))
}

// ClientCertificateMode returns a reference to field client_certificate_mode of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("client_certificate_mode"))
}

// ContentShareForceDisabled returns a reference to field content_share_force_disabled of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) ContentShareForceDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(wfa.ref.Append("content_share_force_disabled"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("custom_domain_verification_id"))
}

// DailyMemoryTimeQuota returns a reference to field daily_memory_time_quota of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) DailyMemoryTimeQuota() terra.NumberValue {
	return terra.ReferenceAsNumber(wfa.ref.Append("daily_memory_time_quota"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("default_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(wfa.ref.Append("enabled"))
}

// FunctionsExtensionVersion returns a reference to field functions_extension_version of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) FunctionsExtensionVersion() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("functions_extension_version"))
}

// HostingEnvironmentId returns a reference to field hosting_environment_id of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) HostingEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("hosting_environment_id"))
}

// HttpsOnly returns a reference to field https_only of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(wfa.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("id"))
}

// KeyVaultReferenceIdentityId returns a reference to field key_vault_reference_identity_id of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) KeyVaultReferenceIdentityId() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("key_vault_reference_identity_id"))
}

// Kind returns a reference to field kind of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("name"))
}

// OutboundIpAddressList returns a reference to field outbound_ip_address_list of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) OutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wfa.ref.Append("outbound_ip_address_list"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddressList returns a reference to field possible_outbound_ip_address_list of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) PossibleOutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wfa.ref.Append("possible_outbound_ip_address_list"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("possible_outbound_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("resource_group_name"))
}

// ServicePlanId returns a reference to field service_plan_id of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) ServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("service_plan_id"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("storage_account_access_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("storage_account_name"))
}

// StorageKeyVaultSecretId returns a reference to field storage_key_vault_secret_id of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) StorageKeyVaultSecretId() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("storage_key_vault_secret_id"))
}

// StorageUsesManagedIdentity returns a reference to field storage_uses_managed_identity of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) StorageUsesManagedIdentity() terra.BoolValue {
	return terra.ReferenceAsBool(wfa.ref.Append("storage_uses_managed_identity"))
}

// Tags returns a reference to field tags of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wfa.ref.Append("tags"))
}

// VirtualNetworkSubnetId returns a reference to field virtual_network_subnet_id of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("virtual_network_subnet_id"))
}

// ZipDeployFile returns a reference to field zip_deploy_file of azurerm_windows_function_app.
func (wfa windowsFunctionAppAttributes) ZipDeployFile() terra.StringValue {
	return terra.ReferenceAsString(wfa.ref.Append("zip_deploy_file"))
}

func (wfa windowsFunctionAppAttributes) SiteCredential() terra.ListValue[windowsfunctionapp.SiteCredentialAttributes] {
	return terra.ReferenceAsList[windowsfunctionapp.SiteCredentialAttributes](wfa.ref.Append("site_credential"))
}

func (wfa windowsFunctionAppAttributes) AuthSettings() terra.ListValue[windowsfunctionapp.AuthSettingsAttributes] {
	return terra.ReferenceAsList[windowsfunctionapp.AuthSettingsAttributes](wfa.ref.Append("auth_settings"))
}

func (wfa windowsFunctionAppAttributes) AuthSettingsV2() terra.ListValue[windowsfunctionapp.AuthSettingsV2Attributes] {
	return terra.ReferenceAsList[windowsfunctionapp.AuthSettingsV2Attributes](wfa.ref.Append("auth_settings_v2"))
}

func (wfa windowsFunctionAppAttributes) Backup() terra.ListValue[windowsfunctionapp.BackupAttributes] {
	return terra.ReferenceAsList[windowsfunctionapp.BackupAttributes](wfa.ref.Append("backup"))
}

func (wfa windowsFunctionAppAttributes) ConnectionString() terra.SetValue[windowsfunctionapp.ConnectionStringAttributes] {
	return terra.ReferenceAsSet[windowsfunctionapp.ConnectionStringAttributes](wfa.ref.Append("connection_string"))
}

func (wfa windowsFunctionAppAttributes) Identity() terra.ListValue[windowsfunctionapp.IdentityAttributes] {
	return terra.ReferenceAsList[windowsfunctionapp.IdentityAttributes](wfa.ref.Append("identity"))
}

func (wfa windowsFunctionAppAttributes) SiteConfig() terra.ListValue[windowsfunctionapp.SiteConfigAttributes] {
	return terra.ReferenceAsList[windowsfunctionapp.SiteConfigAttributes](wfa.ref.Append("site_config"))
}

func (wfa windowsFunctionAppAttributes) StickySettings() terra.ListValue[windowsfunctionapp.StickySettingsAttributes] {
	return terra.ReferenceAsList[windowsfunctionapp.StickySettingsAttributes](wfa.ref.Append("sticky_settings"))
}

func (wfa windowsFunctionAppAttributes) StorageAccount() terra.SetValue[windowsfunctionapp.StorageAccountAttributes] {
	return terra.ReferenceAsSet[windowsfunctionapp.StorageAccountAttributes](wfa.ref.Append("storage_account"))
}

func (wfa windowsFunctionAppAttributes) Timeouts() windowsfunctionapp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[windowsfunctionapp.TimeoutsAttributes](wfa.ref.Append("timeouts"))
}

type windowsFunctionAppState struct {
	AppSettings                     map[string]string                          `json:"app_settings"`
	BuiltinLoggingEnabled           bool                                       `json:"builtin_logging_enabled"`
	ClientCertificateEnabled        bool                                       `json:"client_certificate_enabled"`
	ClientCertificateExclusionPaths string                                     `json:"client_certificate_exclusion_paths"`
	ClientCertificateMode           string                                     `json:"client_certificate_mode"`
	ContentShareForceDisabled       bool                                       `json:"content_share_force_disabled"`
	CustomDomainVerificationId      string                                     `json:"custom_domain_verification_id"`
	DailyMemoryTimeQuota            float64                                    `json:"daily_memory_time_quota"`
	DefaultHostname                 string                                     `json:"default_hostname"`
	Enabled                         bool                                       `json:"enabled"`
	FunctionsExtensionVersion       string                                     `json:"functions_extension_version"`
	HostingEnvironmentId            string                                     `json:"hosting_environment_id"`
	HttpsOnly                       bool                                       `json:"https_only"`
	Id                              string                                     `json:"id"`
	KeyVaultReferenceIdentityId     string                                     `json:"key_vault_reference_identity_id"`
	Kind                            string                                     `json:"kind"`
	Location                        string                                     `json:"location"`
	Name                            string                                     `json:"name"`
	OutboundIpAddressList           []string                                   `json:"outbound_ip_address_list"`
	OutboundIpAddresses             string                                     `json:"outbound_ip_addresses"`
	PossibleOutboundIpAddressList   []string                                   `json:"possible_outbound_ip_address_list"`
	PossibleOutboundIpAddresses     string                                     `json:"possible_outbound_ip_addresses"`
	ResourceGroupName               string                                     `json:"resource_group_name"`
	ServicePlanId                   string                                     `json:"service_plan_id"`
	StorageAccountAccessKey         string                                     `json:"storage_account_access_key"`
	StorageAccountName              string                                     `json:"storage_account_name"`
	StorageKeyVaultSecretId         string                                     `json:"storage_key_vault_secret_id"`
	StorageUsesManagedIdentity      bool                                       `json:"storage_uses_managed_identity"`
	Tags                            map[string]string                          `json:"tags"`
	VirtualNetworkSubnetId          string                                     `json:"virtual_network_subnet_id"`
	ZipDeployFile                   string                                     `json:"zip_deploy_file"`
	SiteCredential                  []windowsfunctionapp.SiteCredentialState   `json:"site_credential"`
	AuthSettings                    []windowsfunctionapp.AuthSettingsState     `json:"auth_settings"`
	AuthSettingsV2                  []windowsfunctionapp.AuthSettingsV2State   `json:"auth_settings_v2"`
	Backup                          []windowsfunctionapp.BackupState           `json:"backup"`
	ConnectionString                []windowsfunctionapp.ConnectionStringState `json:"connection_string"`
	Identity                        []windowsfunctionapp.IdentityState         `json:"identity"`
	SiteConfig                      []windowsfunctionapp.SiteConfigState       `json:"site_config"`
	StickySettings                  []windowsfunctionapp.StickySettingsState   `json:"sticky_settings"`
	StorageAccount                  []windowsfunctionapp.StorageAccountState   `json:"storage_account"`
	Timeouts                        *windowsfunctionapp.TimeoutsState          `json:"timeouts"`
}
