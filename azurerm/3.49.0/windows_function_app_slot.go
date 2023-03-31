// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	windowsfunctionappslot "github.com/golingon/terraproviders/azurerm/3.49.0/windowsfunctionappslot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewWindowsFunctionAppSlot(name string, args WindowsFunctionAppSlotArgs) *WindowsFunctionAppSlot {
	return &WindowsFunctionAppSlot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WindowsFunctionAppSlot)(nil)

type WindowsFunctionAppSlot struct {
	Name  string
	Args  WindowsFunctionAppSlotArgs
	state *windowsFunctionAppSlotState
}

func (wfas *WindowsFunctionAppSlot) Type() string {
	return "azurerm_windows_function_app_slot"
}

func (wfas *WindowsFunctionAppSlot) LocalName() string {
	return wfas.Name
}

func (wfas *WindowsFunctionAppSlot) Configuration() interface{} {
	return wfas.Args
}

func (wfas *WindowsFunctionAppSlot) Attributes() windowsFunctionAppSlotAttributes {
	return windowsFunctionAppSlotAttributes{ref: terra.ReferenceResource(wfas)}
}

func (wfas *WindowsFunctionAppSlot) ImportState(av io.Reader) error {
	wfas.state = &windowsFunctionAppSlotState{}
	if err := json.NewDecoder(av).Decode(wfas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wfas.Type(), wfas.LocalName(), err)
	}
	return nil
}

func (wfas *WindowsFunctionAppSlot) State() (*windowsFunctionAppSlotState, bool) {
	return wfas.state, wfas.state != nil
}

func (wfas *WindowsFunctionAppSlot) StateMust() *windowsFunctionAppSlotState {
	if wfas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wfas.Type(), wfas.LocalName()))
	}
	return wfas.state
}

func (wfas *WindowsFunctionAppSlot) DependOn() terra.Reference {
	return terra.ReferenceResource(wfas)
}

type WindowsFunctionAppSlotArgs struct {
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
	SiteCredential []windowsfunctionappslot.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// AuthSettings: optional
	AuthSettings *windowsfunctionappslot.AuthSettings `hcl:"auth_settings,block"`
	// AuthSettingsV2: optional
	AuthSettingsV2 *windowsfunctionappslot.AuthSettingsV2 `hcl:"auth_settings_v2,block"`
	// Backup: optional
	Backup *windowsfunctionappslot.Backup `hcl:"backup,block"`
	// ConnectionString: min=0
	ConnectionString []windowsfunctionappslot.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: optional
	Identity *windowsfunctionappslot.Identity `hcl:"identity,block"`
	// SiteConfig: required
	SiteConfig *windowsfunctionappslot.SiteConfig `hcl:"site_config,block" validate:"required"`
	// StorageAccount: min=0
	StorageAccount []windowsfunctionappslot.StorageAccount `hcl:"storage_account,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *windowsfunctionappslot.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that WindowsFunctionAppSlot depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type windowsFunctionAppSlotAttributes struct {
	ref terra.Reference
}

func (wfas windowsFunctionAppSlotAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](wfas.ref.Append("app_settings"))
}

func (wfas windowsFunctionAppSlotAttributes) BuiltinLoggingEnabled() terra.BoolValue {
	return terra.ReferenceBool(wfas.ref.Append("builtin_logging_enabled"))
}

func (wfas windowsFunctionAppSlotAttributes) ClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceBool(wfas.ref.Append("client_certificate_enabled"))
}

func (wfas windowsFunctionAppSlotAttributes) ClientCertificateExclusionPaths() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("client_certificate_exclusion_paths"))
}

func (wfas windowsFunctionAppSlotAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("client_certificate_mode"))
}

func (wfas windowsFunctionAppSlotAttributes) ContentShareForceDisabled() terra.BoolValue {
	return terra.ReferenceBool(wfas.ref.Append("content_share_force_disabled"))
}

func (wfas windowsFunctionAppSlotAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("custom_domain_verification_id"))
}

func (wfas windowsFunctionAppSlotAttributes) DailyMemoryTimeQuota() terra.NumberValue {
	return terra.ReferenceNumber(wfas.ref.Append("daily_memory_time_quota"))
}

func (wfas windowsFunctionAppSlotAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("default_hostname"))
}

func (wfas windowsFunctionAppSlotAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(wfas.ref.Append("enabled"))
}

func (wfas windowsFunctionAppSlotAttributes) FunctionAppId() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("function_app_id"))
}

func (wfas windowsFunctionAppSlotAttributes) FunctionsExtensionVersion() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("functions_extension_version"))
}

func (wfas windowsFunctionAppSlotAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceBool(wfas.ref.Append("https_only"))
}

func (wfas windowsFunctionAppSlotAttributes) Id() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("id"))
}

func (wfas windowsFunctionAppSlotAttributes) KeyVaultReferenceIdentityId() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("key_vault_reference_identity_id"))
}

func (wfas windowsFunctionAppSlotAttributes) Kind() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("kind"))
}

func (wfas windowsFunctionAppSlotAttributes) Name() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("name"))
}

func (wfas windowsFunctionAppSlotAttributes) OutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](wfas.ref.Append("outbound_ip_address_list"))
}

func (wfas windowsFunctionAppSlotAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("outbound_ip_addresses"))
}

func (wfas windowsFunctionAppSlotAttributes) PossibleOutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](wfas.ref.Append("possible_outbound_ip_address_list"))
}

func (wfas windowsFunctionAppSlotAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("possible_outbound_ip_addresses"))
}

func (wfas windowsFunctionAppSlotAttributes) ServicePlanId() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("service_plan_id"))
}

func (wfas windowsFunctionAppSlotAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("storage_account_access_key"))
}

func (wfas windowsFunctionAppSlotAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("storage_account_name"))
}

func (wfas windowsFunctionAppSlotAttributes) StorageKeyVaultSecretId() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("storage_key_vault_secret_id"))
}

func (wfas windowsFunctionAppSlotAttributes) StorageUsesManagedIdentity() terra.BoolValue {
	return terra.ReferenceBool(wfas.ref.Append("storage_uses_managed_identity"))
}

func (wfas windowsFunctionAppSlotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](wfas.ref.Append("tags"))
}

func (wfas windowsFunctionAppSlotAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceString(wfas.ref.Append("virtual_network_subnet_id"))
}

func (wfas windowsFunctionAppSlotAttributes) SiteCredential() terra.ListValue[windowsfunctionappslot.SiteCredentialAttributes] {
	return terra.ReferenceList[windowsfunctionappslot.SiteCredentialAttributes](wfas.ref.Append("site_credential"))
}

func (wfas windowsFunctionAppSlotAttributes) AuthSettings() terra.ListValue[windowsfunctionappslot.AuthSettingsAttributes] {
	return terra.ReferenceList[windowsfunctionappslot.AuthSettingsAttributes](wfas.ref.Append("auth_settings"))
}

func (wfas windowsFunctionAppSlotAttributes) AuthSettingsV2() terra.ListValue[windowsfunctionappslot.AuthSettingsV2Attributes] {
	return terra.ReferenceList[windowsfunctionappslot.AuthSettingsV2Attributes](wfas.ref.Append("auth_settings_v2"))
}

func (wfas windowsFunctionAppSlotAttributes) Backup() terra.ListValue[windowsfunctionappslot.BackupAttributes] {
	return terra.ReferenceList[windowsfunctionappslot.BackupAttributes](wfas.ref.Append("backup"))
}

func (wfas windowsFunctionAppSlotAttributes) ConnectionString() terra.SetValue[windowsfunctionappslot.ConnectionStringAttributes] {
	return terra.ReferenceSet[windowsfunctionappslot.ConnectionStringAttributes](wfas.ref.Append("connection_string"))
}

func (wfas windowsFunctionAppSlotAttributes) Identity() terra.ListValue[windowsfunctionappslot.IdentityAttributes] {
	return terra.ReferenceList[windowsfunctionappslot.IdentityAttributes](wfas.ref.Append("identity"))
}

func (wfas windowsFunctionAppSlotAttributes) SiteConfig() terra.ListValue[windowsfunctionappslot.SiteConfigAttributes] {
	return terra.ReferenceList[windowsfunctionappslot.SiteConfigAttributes](wfas.ref.Append("site_config"))
}

func (wfas windowsFunctionAppSlotAttributes) StorageAccount() terra.SetValue[windowsfunctionappslot.StorageAccountAttributes] {
	return terra.ReferenceSet[windowsfunctionappslot.StorageAccountAttributes](wfas.ref.Append("storage_account"))
}

func (wfas windowsFunctionAppSlotAttributes) Timeouts() windowsfunctionappslot.TimeoutsAttributes {
	return terra.ReferenceSingle[windowsfunctionappslot.TimeoutsAttributes](wfas.ref.Append("timeouts"))
}

type windowsFunctionAppSlotState struct {
	AppSettings                     map[string]string                              `json:"app_settings"`
	BuiltinLoggingEnabled           bool                                           `json:"builtin_logging_enabled"`
	ClientCertificateEnabled        bool                                           `json:"client_certificate_enabled"`
	ClientCertificateExclusionPaths string                                         `json:"client_certificate_exclusion_paths"`
	ClientCertificateMode           string                                         `json:"client_certificate_mode"`
	ContentShareForceDisabled       bool                                           `json:"content_share_force_disabled"`
	CustomDomainVerificationId      string                                         `json:"custom_domain_verification_id"`
	DailyMemoryTimeQuota            float64                                        `json:"daily_memory_time_quota"`
	DefaultHostname                 string                                         `json:"default_hostname"`
	Enabled                         bool                                           `json:"enabled"`
	FunctionAppId                   string                                         `json:"function_app_id"`
	FunctionsExtensionVersion       string                                         `json:"functions_extension_version"`
	HttpsOnly                       bool                                           `json:"https_only"`
	Id                              string                                         `json:"id"`
	KeyVaultReferenceIdentityId     string                                         `json:"key_vault_reference_identity_id"`
	Kind                            string                                         `json:"kind"`
	Name                            string                                         `json:"name"`
	OutboundIpAddressList           []string                                       `json:"outbound_ip_address_list"`
	OutboundIpAddresses             string                                         `json:"outbound_ip_addresses"`
	PossibleOutboundIpAddressList   []string                                       `json:"possible_outbound_ip_address_list"`
	PossibleOutboundIpAddresses     string                                         `json:"possible_outbound_ip_addresses"`
	ServicePlanId                   string                                         `json:"service_plan_id"`
	StorageAccountAccessKey         string                                         `json:"storage_account_access_key"`
	StorageAccountName              string                                         `json:"storage_account_name"`
	StorageKeyVaultSecretId         string                                         `json:"storage_key_vault_secret_id"`
	StorageUsesManagedIdentity      bool                                           `json:"storage_uses_managed_identity"`
	Tags                            map[string]string                              `json:"tags"`
	VirtualNetworkSubnetId          string                                         `json:"virtual_network_subnet_id"`
	SiteCredential                  []windowsfunctionappslot.SiteCredentialState   `json:"site_credential"`
	AuthSettings                    []windowsfunctionappslot.AuthSettingsState     `json:"auth_settings"`
	AuthSettingsV2                  []windowsfunctionappslot.AuthSettingsV2State   `json:"auth_settings_v2"`
	Backup                          []windowsfunctionappslot.BackupState           `json:"backup"`
	ConnectionString                []windowsfunctionappslot.ConnectionStringState `json:"connection_string"`
	Identity                        []windowsfunctionappslot.IdentityState         `json:"identity"`
	SiteConfig                      []windowsfunctionappslot.SiteConfigState       `json:"site_config"`
	StorageAccount                  []windowsfunctionappslot.StorageAccountState   `json:"storage_account"`
	Timeouts                        *windowsfunctionappslot.TimeoutsState          `json:"timeouts"`
}
