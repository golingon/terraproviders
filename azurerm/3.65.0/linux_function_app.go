// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	linuxfunctionapp "github.com/golingon/terraproviders/azurerm/3.65.0/linuxfunctionapp"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLinuxFunctionApp creates a new instance of [LinuxFunctionApp].
func NewLinuxFunctionApp(name string, args LinuxFunctionAppArgs) *LinuxFunctionApp {
	return &LinuxFunctionApp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LinuxFunctionApp)(nil)

// LinuxFunctionApp represents the Terraform resource azurerm_linux_function_app.
type LinuxFunctionApp struct {
	Name      string
	Args      LinuxFunctionAppArgs
	state     *linuxFunctionAppState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LinuxFunctionApp].
func (lfa *LinuxFunctionApp) Type() string {
	return "azurerm_linux_function_app"
}

// LocalName returns the local name for [LinuxFunctionApp].
func (lfa *LinuxFunctionApp) LocalName() string {
	return lfa.Name
}

// Configuration returns the configuration (args) for [LinuxFunctionApp].
func (lfa *LinuxFunctionApp) Configuration() interface{} {
	return lfa.Args
}

// DependOn is used for other resources to depend on [LinuxFunctionApp].
func (lfa *LinuxFunctionApp) DependOn() terra.Reference {
	return terra.ReferenceResource(lfa)
}

// Dependencies returns the list of resources [LinuxFunctionApp] depends_on.
func (lfa *LinuxFunctionApp) Dependencies() terra.Dependencies {
	return lfa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LinuxFunctionApp].
func (lfa *LinuxFunctionApp) LifecycleManagement() *terra.Lifecycle {
	return lfa.Lifecycle
}

// Attributes returns the attributes for [LinuxFunctionApp].
func (lfa *LinuxFunctionApp) Attributes() linuxFunctionAppAttributes {
	return linuxFunctionAppAttributes{ref: terra.ReferenceResource(lfa)}
}

// ImportState imports the given attribute values into [LinuxFunctionApp]'s state.
func (lfa *LinuxFunctionApp) ImportState(av io.Reader) error {
	lfa.state = &linuxFunctionAppState{}
	if err := json.NewDecoder(av).Decode(lfa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lfa.Type(), lfa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LinuxFunctionApp] has state.
func (lfa *LinuxFunctionApp) State() (*linuxFunctionAppState, bool) {
	return lfa.state, lfa.state != nil
}

// StateMust returns the state for [LinuxFunctionApp]. Panics if the state is nil.
func (lfa *LinuxFunctionApp) StateMust() *linuxFunctionAppState {
	if lfa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lfa.Type(), lfa.LocalName()))
	}
	return lfa.state
}

// LinuxFunctionAppArgs contains the configurations for azurerm_linux_function_app.
type LinuxFunctionAppArgs struct {
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
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
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
	SiteCredential []linuxfunctionapp.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// AuthSettings: optional
	AuthSettings *linuxfunctionapp.AuthSettings `hcl:"auth_settings,block"`
	// AuthSettingsV2: optional
	AuthSettingsV2 *linuxfunctionapp.AuthSettingsV2 `hcl:"auth_settings_v2,block"`
	// Backup: optional
	Backup *linuxfunctionapp.Backup `hcl:"backup,block"`
	// ConnectionString: min=0
	ConnectionString []linuxfunctionapp.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: optional
	Identity *linuxfunctionapp.Identity `hcl:"identity,block"`
	// SiteConfig: required
	SiteConfig *linuxfunctionapp.SiteConfig `hcl:"site_config,block" validate:"required"`
	// StickySettings: optional
	StickySettings *linuxfunctionapp.StickySettings `hcl:"sticky_settings,block"`
	// StorageAccount: min=0
	StorageAccount []linuxfunctionapp.StorageAccount `hcl:"storage_account,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *linuxfunctionapp.Timeouts `hcl:"timeouts,block"`
}
type linuxFunctionAppAttributes struct {
	ref terra.Reference
}

// AppSettings returns a reference to field app_settings of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lfa.ref.Append("app_settings"))
}

// BuiltinLoggingEnabled returns a reference to field builtin_logging_enabled of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) BuiltinLoggingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfa.ref.Append("builtin_logging_enabled"))
}

// ClientCertificateEnabled returns a reference to field client_certificate_enabled of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) ClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfa.ref.Append("client_certificate_enabled"))
}

// ClientCertificateExclusionPaths returns a reference to field client_certificate_exclusion_paths of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) ClientCertificateExclusionPaths() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("client_certificate_exclusion_paths"))
}

// ClientCertificateMode returns a reference to field client_certificate_mode of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("client_certificate_mode"))
}

// ContentShareForceDisabled returns a reference to field content_share_force_disabled of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) ContentShareForceDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfa.ref.Append("content_share_force_disabled"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("custom_domain_verification_id"))
}

// DailyMemoryTimeQuota returns a reference to field daily_memory_time_quota of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) DailyMemoryTimeQuota() terra.NumberValue {
	return terra.ReferenceAsNumber(lfa.ref.Append("daily_memory_time_quota"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("default_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfa.ref.Append("enabled"))
}

// FunctionsExtensionVersion returns a reference to field functions_extension_version of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) FunctionsExtensionVersion() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("functions_extension_version"))
}

// HostingEnvironmentId returns a reference to field hosting_environment_id of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) HostingEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("hosting_environment_id"))
}

// HttpsOnly returns a reference to field https_only of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(lfa.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("id"))
}

// KeyVaultReferenceIdentityId returns a reference to field key_vault_reference_identity_id of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) KeyVaultReferenceIdentityId() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("key_vault_reference_identity_id"))
}

// Kind returns a reference to field kind of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("name"))
}

// OutboundIpAddressList returns a reference to field outbound_ip_address_list of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) OutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lfa.ref.Append("outbound_ip_address_list"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddressList returns a reference to field possible_outbound_ip_address_list of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) PossibleOutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lfa.ref.Append("possible_outbound_ip_address_list"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("possible_outbound_ip_addresses"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfa.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("resource_group_name"))
}

// ServicePlanId returns a reference to field service_plan_id of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) ServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("service_plan_id"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("storage_account_access_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("storage_account_name"))
}

// StorageKeyVaultSecretId returns a reference to field storage_key_vault_secret_id of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) StorageKeyVaultSecretId() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("storage_key_vault_secret_id"))
}

// StorageUsesManagedIdentity returns a reference to field storage_uses_managed_identity of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) StorageUsesManagedIdentity() terra.BoolValue {
	return terra.ReferenceAsBool(lfa.ref.Append("storage_uses_managed_identity"))
}

// Tags returns a reference to field tags of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lfa.ref.Append("tags"))
}

// VirtualNetworkSubnetId returns a reference to field virtual_network_subnet_id of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("virtual_network_subnet_id"))
}

// ZipDeployFile returns a reference to field zip_deploy_file of azurerm_linux_function_app.
func (lfa linuxFunctionAppAttributes) ZipDeployFile() terra.StringValue {
	return terra.ReferenceAsString(lfa.ref.Append("zip_deploy_file"))
}

func (lfa linuxFunctionAppAttributes) SiteCredential() terra.ListValue[linuxfunctionapp.SiteCredentialAttributes] {
	return terra.ReferenceAsList[linuxfunctionapp.SiteCredentialAttributes](lfa.ref.Append("site_credential"))
}

func (lfa linuxFunctionAppAttributes) AuthSettings() terra.ListValue[linuxfunctionapp.AuthSettingsAttributes] {
	return terra.ReferenceAsList[linuxfunctionapp.AuthSettingsAttributes](lfa.ref.Append("auth_settings"))
}

func (lfa linuxFunctionAppAttributes) AuthSettingsV2() terra.ListValue[linuxfunctionapp.AuthSettingsV2Attributes] {
	return terra.ReferenceAsList[linuxfunctionapp.AuthSettingsV2Attributes](lfa.ref.Append("auth_settings_v2"))
}

func (lfa linuxFunctionAppAttributes) Backup() terra.ListValue[linuxfunctionapp.BackupAttributes] {
	return terra.ReferenceAsList[linuxfunctionapp.BackupAttributes](lfa.ref.Append("backup"))
}

func (lfa linuxFunctionAppAttributes) ConnectionString() terra.SetValue[linuxfunctionapp.ConnectionStringAttributes] {
	return terra.ReferenceAsSet[linuxfunctionapp.ConnectionStringAttributes](lfa.ref.Append("connection_string"))
}

func (lfa linuxFunctionAppAttributes) Identity() terra.ListValue[linuxfunctionapp.IdentityAttributes] {
	return terra.ReferenceAsList[linuxfunctionapp.IdentityAttributes](lfa.ref.Append("identity"))
}

func (lfa linuxFunctionAppAttributes) SiteConfig() terra.ListValue[linuxfunctionapp.SiteConfigAttributes] {
	return terra.ReferenceAsList[linuxfunctionapp.SiteConfigAttributes](lfa.ref.Append("site_config"))
}

func (lfa linuxFunctionAppAttributes) StickySettings() terra.ListValue[linuxfunctionapp.StickySettingsAttributes] {
	return terra.ReferenceAsList[linuxfunctionapp.StickySettingsAttributes](lfa.ref.Append("sticky_settings"))
}

func (lfa linuxFunctionAppAttributes) StorageAccount() terra.SetValue[linuxfunctionapp.StorageAccountAttributes] {
	return terra.ReferenceAsSet[linuxfunctionapp.StorageAccountAttributes](lfa.ref.Append("storage_account"))
}

func (lfa linuxFunctionAppAttributes) Timeouts() linuxfunctionapp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[linuxfunctionapp.TimeoutsAttributes](lfa.ref.Append("timeouts"))
}

type linuxFunctionAppState struct {
	AppSettings                     map[string]string                        `json:"app_settings"`
	BuiltinLoggingEnabled           bool                                     `json:"builtin_logging_enabled"`
	ClientCertificateEnabled        bool                                     `json:"client_certificate_enabled"`
	ClientCertificateExclusionPaths string                                   `json:"client_certificate_exclusion_paths"`
	ClientCertificateMode           string                                   `json:"client_certificate_mode"`
	ContentShareForceDisabled       bool                                     `json:"content_share_force_disabled"`
	CustomDomainVerificationId      string                                   `json:"custom_domain_verification_id"`
	DailyMemoryTimeQuota            float64                                  `json:"daily_memory_time_quota"`
	DefaultHostname                 string                                   `json:"default_hostname"`
	Enabled                         bool                                     `json:"enabled"`
	FunctionsExtensionVersion       string                                   `json:"functions_extension_version"`
	HostingEnvironmentId            string                                   `json:"hosting_environment_id"`
	HttpsOnly                       bool                                     `json:"https_only"`
	Id                              string                                   `json:"id"`
	KeyVaultReferenceIdentityId     string                                   `json:"key_vault_reference_identity_id"`
	Kind                            string                                   `json:"kind"`
	Location                        string                                   `json:"location"`
	Name                            string                                   `json:"name"`
	OutboundIpAddressList           []string                                 `json:"outbound_ip_address_list"`
	OutboundIpAddresses             string                                   `json:"outbound_ip_addresses"`
	PossibleOutboundIpAddressList   []string                                 `json:"possible_outbound_ip_address_list"`
	PossibleOutboundIpAddresses     string                                   `json:"possible_outbound_ip_addresses"`
	PublicNetworkAccessEnabled      bool                                     `json:"public_network_access_enabled"`
	ResourceGroupName               string                                   `json:"resource_group_name"`
	ServicePlanId                   string                                   `json:"service_plan_id"`
	StorageAccountAccessKey         string                                   `json:"storage_account_access_key"`
	StorageAccountName              string                                   `json:"storage_account_name"`
	StorageKeyVaultSecretId         string                                   `json:"storage_key_vault_secret_id"`
	StorageUsesManagedIdentity      bool                                     `json:"storage_uses_managed_identity"`
	Tags                            map[string]string                        `json:"tags"`
	VirtualNetworkSubnetId          string                                   `json:"virtual_network_subnet_id"`
	ZipDeployFile                   string                                   `json:"zip_deploy_file"`
	SiteCredential                  []linuxfunctionapp.SiteCredentialState   `json:"site_credential"`
	AuthSettings                    []linuxfunctionapp.AuthSettingsState     `json:"auth_settings"`
	AuthSettingsV2                  []linuxfunctionapp.AuthSettingsV2State   `json:"auth_settings_v2"`
	Backup                          []linuxfunctionapp.BackupState           `json:"backup"`
	ConnectionString                []linuxfunctionapp.ConnectionStringState `json:"connection_string"`
	Identity                        []linuxfunctionapp.IdentityState         `json:"identity"`
	SiteConfig                      []linuxfunctionapp.SiteConfigState       `json:"site_config"`
	StickySettings                  []linuxfunctionapp.StickySettingsState   `json:"sticky_settings"`
	StorageAccount                  []linuxfunctionapp.StorageAccountState   `json:"storage_account"`
	Timeouts                        *linuxfunctionapp.TimeoutsState          `json:"timeouts"`
}
