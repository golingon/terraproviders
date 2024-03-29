// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	linuxfunctionappslot "github.com/golingon/terraproviders/azurerm/3.74.0/linuxfunctionappslot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLinuxFunctionAppSlot creates a new instance of [LinuxFunctionAppSlot].
func NewLinuxFunctionAppSlot(name string, args LinuxFunctionAppSlotArgs) *LinuxFunctionAppSlot {
	return &LinuxFunctionAppSlot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LinuxFunctionAppSlot)(nil)

// LinuxFunctionAppSlot represents the Terraform resource azurerm_linux_function_app_slot.
type LinuxFunctionAppSlot struct {
	Name      string
	Args      LinuxFunctionAppSlotArgs
	state     *linuxFunctionAppSlotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LinuxFunctionAppSlot].
func (lfas *LinuxFunctionAppSlot) Type() string {
	return "azurerm_linux_function_app_slot"
}

// LocalName returns the local name for [LinuxFunctionAppSlot].
func (lfas *LinuxFunctionAppSlot) LocalName() string {
	return lfas.Name
}

// Configuration returns the configuration (args) for [LinuxFunctionAppSlot].
func (lfas *LinuxFunctionAppSlot) Configuration() interface{} {
	return lfas.Args
}

// DependOn is used for other resources to depend on [LinuxFunctionAppSlot].
func (lfas *LinuxFunctionAppSlot) DependOn() terra.Reference {
	return terra.ReferenceResource(lfas)
}

// Dependencies returns the list of resources [LinuxFunctionAppSlot] depends_on.
func (lfas *LinuxFunctionAppSlot) Dependencies() terra.Dependencies {
	return lfas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LinuxFunctionAppSlot].
func (lfas *LinuxFunctionAppSlot) LifecycleManagement() *terra.Lifecycle {
	return lfas.Lifecycle
}

// Attributes returns the attributes for [LinuxFunctionAppSlot].
func (lfas *LinuxFunctionAppSlot) Attributes() linuxFunctionAppSlotAttributes {
	return linuxFunctionAppSlotAttributes{ref: terra.ReferenceResource(lfas)}
}

// ImportState imports the given attribute values into [LinuxFunctionAppSlot]'s state.
func (lfas *LinuxFunctionAppSlot) ImportState(av io.Reader) error {
	lfas.state = &linuxFunctionAppSlotState{}
	if err := json.NewDecoder(av).Decode(lfas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lfas.Type(), lfas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LinuxFunctionAppSlot] has state.
func (lfas *LinuxFunctionAppSlot) State() (*linuxFunctionAppSlotState, bool) {
	return lfas.state, lfas.state != nil
}

// StateMust returns the state for [LinuxFunctionAppSlot]. Panics if the state is nil.
func (lfas *LinuxFunctionAppSlot) StateMust() *linuxFunctionAppSlotState {
	if lfas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lfas.Type(), lfas.LocalName()))
	}
	return lfas.state
}

// LinuxFunctionAppSlotArgs contains the configurations for azurerm_linux_function_app_slot.
type LinuxFunctionAppSlotArgs struct {
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
	// FunctionAppId: string, required
	FunctionAppId terra.StringValue `hcl:"function_app_id,attr" validate:"required"`
	// FunctionsExtensionVersion: string, optional
	FunctionsExtensionVersion terra.StringValue `hcl:"functions_extension_version,attr"`
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
	// SiteCredential: min=0
	SiteCredential []linuxfunctionappslot.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// AuthSettings: optional
	AuthSettings *linuxfunctionappslot.AuthSettings `hcl:"auth_settings,block"`
	// AuthSettingsV2: optional
	AuthSettingsV2 *linuxfunctionappslot.AuthSettingsV2 `hcl:"auth_settings_v2,block"`
	// Backup: optional
	Backup *linuxfunctionappslot.Backup `hcl:"backup,block"`
	// ConnectionString: min=0
	ConnectionString []linuxfunctionappslot.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: optional
	Identity *linuxfunctionappslot.Identity `hcl:"identity,block"`
	// SiteConfig: required
	SiteConfig *linuxfunctionappslot.SiteConfig `hcl:"site_config,block" validate:"required"`
	// StorageAccount: min=0
	StorageAccount []linuxfunctionappslot.StorageAccount `hcl:"storage_account,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *linuxfunctionappslot.Timeouts `hcl:"timeouts,block"`
}
type linuxFunctionAppSlotAttributes struct {
	ref terra.Reference
}

// AppSettings returns a reference to field app_settings of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lfas.ref.Append("app_settings"))
}

// BuiltinLoggingEnabled returns a reference to field builtin_logging_enabled of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) BuiltinLoggingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfas.ref.Append("builtin_logging_enabled"))
}

// ClientCertificateEnabled returns a reference to field client_certificate_enabled of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) ClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfas.ref.Append("client_certificate_enabled"))
}

// ClientCertificateExclusionPaths returns a reference to field client_certificate_exclusion_paths of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) ClientCertificateExclusionPaths() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("client_certificate_exclusion_paths"))
}

// ClientCertificateMode returns a reference to field client_certificate_mode of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("client_certificate_mode"))
}

// ContentShareForceDisabled returns a reference to field content_share_force_disabled of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) ContentShareForceDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfas.ref.Append("content_share_force_disabled"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("custom_domain_verification_id"))
}

// DailyMemoryTimeQuota returns a reference to field daily_memory_time_quota of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) DailyMemoryTimeQuota() terra.NumberValue {
	return terra.ReferenceAsNumber(lfas.ref.Append("daily_memory_time_quota"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("default_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfas.ref.Append("enabled"))
}

// FunctionAppId returns a reference to field function_app_id of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) FunctionAppId() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("function_app_id"))
}

// FunctionsExtensionVersion returns a reference to field functions_extension_version of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) FunctionsExtensionVersion() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("functions_extension_version"))
}

// HostingEnvironmentId returns a reference to field hosting_environment_id of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) HostingEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("hosting_environment_id"))
}

// HttpsOnly returns a reference to field https_only of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(lfas.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("id"))
}

// KeyVaultReferenceIdentityId returns a reference to field key_vault_reference_identity_id of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) KeyVaultReferenceIdentityId() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("key_vault_reference_identity_id"))
}

// Kind returns a reference to field kind of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("kind"))
}

// Name returns a reference to field name of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("name"))
}

// OutboundIpAddressList returns a reference to field outbound_ip_address_list of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) OutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lfas.ref.Append("outbound_ip_address_list"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddressList returns a reference to field possible_outbound_ip_address_list of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) PossibleOutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lfas.ref.Append("possible_outbound_ip_address_list"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("possible_outbound_ip_addresses"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lfas.ref.Append("public_network_access_enabled"))
}

// ServicePlanId returns a reference to field service_plan_id of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) ServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("service_plan_id"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("storage_account_access_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("storage_account_name"))
}

// StorageKeyVaultSecretId returns a reference to field storage_key_vault_secret_id of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) StorageKeyVaultSecretId() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("storage_key_vault_secret_id"))
}

// StorageUsesManagedIdentity returns a reference to field storage_uses_managed_identity of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) StorageUsesManagedIdentity() terra.BoolValue {
	return terra.ReferenceAsBool(lfas.ref.Append("storage_uses_managed_identity"))
}

// Tags returns a reference to field tags of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lfas.ref.Append("tags"))
}

// VirtualNetworkSubnetId returns a reference to field virtual_network_subnet_id of azurerm_linux_function_app_slot.
func (lfas linuxFunctionAppSlotAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(lfas.ref.Append("virtual_network_subnet_id"))
}

func (lfas linuxFunctionAppSlotAttributes) SiteCredential() terra.ListValue[linuxfunctionappslot.SiteCredentialAttributes] {
	return terra.ReferenceAsList[linuxfunctionappslot.SiteCredentialAttributes](lfas.ref.Append("site_credential"))
}

func (lfas linuxFunctionAppSlotAttributes) AuthSettings() terra.ListValue[linuxfunctionappslot.AuthSettingsAttributes] {
	return terra.ReferenceAsList[linuxfunctionappslot.AuthSettingsAttributes](lfas.ref.Append("auth_settings"))
}

func (lfas linuxFunctionAppSlotAttributes) AuthSettingsV2() terra.ListValue[linuxfunctionappslot.AuthSettingsV2Attributes] {
	return terra.ReferenceAsList[linuxfunctionappslot.AuthSettingsV2Attributes](lfas.ref.Append("auth_settings_v2"))
}

func (lfas linuxFunctionAppSlotAttributes) Backup() terra.ListValue[linuxfunctionappslot.BackupAttributes] {
	return terra.ReferenceAsList[linuxfunctionappslot.BackupAttributes](lfas.ref.Append("backup"))
}

func (lfas linuxFunctionAppSlotAttributes) ConnectionString() terra.SetValue[linuxfunctionappslot.ConnectionStringAttributes] {
	return terra.ReferenceAsSet[linuxfunctionappslot.ConnectionStringAttributes](lfas.ref.Append("connection_string"))
}

func (lfas linuxFunctionAppSlotAttributes) Identity() terra.ListValue[linuxfunctionappslot.IdentityAttributes] {
	return terra.ReferenceAsList[linuxfunctionappslot.IdentityAttributes](lfas.ref.Append("identity"))
}

func (lfas linuxFunctionAppSlotAttributes) SiteConfig() terra.ListValue[linuxfunctionappslot.SiteConfigAttributes] {
	return terra.ReferenceAsList[linuxfunctionappslot.SiteConfigAttributes](lfas.ref.Append("site_config"))
}

func (lfas linuxFunctionAppSlotAttributes) StorageAccount() terra.SetValue[linuxfunctionappslot.StorageAccountAttributes] {
	return terra.ReferenceAsSet[linuxfunctionappslot.StorageAccountAttributes](lfas.ref.Append("storage_account"))
}

func (lfas linuxFunctionAppSlotAttributes) Timeouts() linuxfunctionappslot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[linuxfunctionappslot.TimeoutsAttributes](lfas.ref.Append("timeouts"))
}

type linuxFunctionAppSlotState struct {
	AppSettings                     map[string]string                            `json:"app_settings"`
	BuiltinLoggingEnabled           bool                                         `json:"builtin_logging_enabled"`
	ClientCertificateEnabled        bool                                         `json:"client_certificate_enabled"`
	ClientCertificateExclusionPaths string                                       `json:"client_certificate_exclusion_paths"`
	ClientCertificateMode           string                                       `json:"client_certificate_mode"`
	ContentShareForceDisabled       bool                                         `json:"content_share_force_disabled"`
	CustomDomainVerificationId      string                                       `json:"custom_domain_verification_id"`
	DailyMemoryTimeQuota            float64                                      `json:"daily_memory_time_quota"`
	DefaultHostname                 string                                       `json:"default_hostname"`
	Enabled                         bool                                         `json:"enabled"`
	FunctionAppId                   string                                       `json:"function_app_id"`
	FunctionsExtensionVersion       string                                       `json:"functions_extension_version"`
	HostingEnvironmentId            string                                       `json:"hosting_environment_id"`
	HttpsOnly                       bool                                         `json:"https_only"`
	Id                              string                                       `json:"id"`
	KeyVaultReferenceIdentityId     string                                       `json:"key_vault_reference_identity_id"`
	Kind                            string                                       `json:"kind"`
	Name                            string                                       `json:"name"`
	OutboundIpAddressList           []string                                     `json:"outbound_ip_address_list"`
	OutboundIpAddresses             string                                       `json:"outbound_ip_addresses"`
	PossibleOutboundIpAddressList   []string                                     `json:"possible_outbound_ip_address_list"`
	PossibleOutboundIpAddresses     string                                       `json:"possible_outbound_ip_addresses"`
	PublicNetworkAccessEnabled      bool                                         `json:"public_network_access_enabled"`
	ServicePlanId                   string                                       `json:"service_plan_id"`
	StorageAccountAccessKey         string                                       `json:"storage_account_access_key"`
	StorageAccountName              string                                       `json:"storage_account_name"`
	StorageKeyVaultSecretId         string                                       `json:"storage_key_vault_secret_id"`
	StorageUsesManagedIdentity      bool                                         `json:"storage_uses_managed_identity"`
	Tags                            map[string]string                            `json:"tags"`
	VirtualNetworkSubnetId          string                                       `json:"virtual_network_subnet_id"`
	SiteCredential                  []linuxfunctionappslot.SiteCredentialState   `json:"site_credential"`
	AuthSettings                    []linuxfunctionappslot.AuthSettingsState     `json:"auth_settings"`
	AuthSettingsV2                  []linuxfunctionappslot.AuthSettingsV2State   `json:"auth_settings_v2"`
	Backup                          []linuxfunctionappslot.BackupState           `json:"backup"`
	ConnectionString                []linuxfunctionappslot.ConnectionStringState `json:"connection_string"`
	Identity                        []linuxfunctionappslot.IdentityState         `json:"identity"`
	SiteConfig                      []linuxfunctionappslot.SiteConfigState       `json:"site_config"`
	StorageAccount                  []linuxfunctionappslot.StorageAccountState   `json:"storage_account"`
	Timeouts                        *linuxfunctionappslot.TimeoutsState          `json:"timeouts"`
}
