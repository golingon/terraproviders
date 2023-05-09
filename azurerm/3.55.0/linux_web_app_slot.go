// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	linuxwebappslot "github.com/golingon/terraproviders/azurerm/3.55.0/linuxwebappslot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLinuxWebAppSlot creates a new instance of [LinuxWebAppSlot].
func NewLinuxWebAppSlot(name string, args LinuxWebAppSlotArgs) *LinuxWebAppSlot {
	return &LinuxWebAppSlot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LinuxWebAppSlot)(nil)

// LinuxWebAppSlot represents the Terraform resource azurerm_linux_web_app_slot.
type LinuxWebAppSlot struct {
	Name      string
	Args      LinuxWebAppSlotArgs
	state     *linuxWebAppSlotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LinuxWebAppSlot].
func (lwas *LinuxWebAppSlot) Type() string {
	return "azurerm_linux_web_app_slot"
}

// LocalName returns the local name for [LinuxWebAppSlot].
func (lwas *LinuxWebAppSlot) LocalName() string {
	return lwas.Name
}

// Configuration returns the configuration (args) for [LinuxWebAppSlot].
func (lwas *LinuxWebAppSlot) Configuration() interface{} {
	return lwas.Args
}

// DependOn is used for other resources to depend on [LinuxWebAppSlot].
func (lwas *LinuxWebAppSlot) DependOn() terra.Reference {
	return terra.ReferenceResource(lwas)
}

// Dependencies returns the list of resources [LinuxWebAppSlot] depends_on.
func (lwas *LinuxWebAppSlot) Dependencies() terra.Dependencies {
	return lwas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LinuxWebAppSlot].
func (lwas *LinuxWebAppSlot) LifecycleManagement() *terra.Lifecycle {
	return lwas.Lifecycle
}

// Attributes returns the attributes for [LinuxWebAppSlot].
func (lwas *LinuxWebAppSlot) Attributes() linuxWebAppSlotAttributes {
	return linuxWebAppSlotAttributes{ref: terra.ReferenceResource(lwas)}
}

// ImportState imports the given attribute values into [LinuxWebAppSlot]'s state.
func (lwas *LinuxWebAppSlot) ImportState(av io.Reader) error {
	lwas.state = &linuxWebAppSlotState{}
	if err := json.NewDecoder(av).Decode(lwas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lwas.Type(), lwas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LinuxWebAppSlot] has state.
func (lwas *LinuxWebAppSlot) State() (*linuxWebAppSlotState, bool) {
	return lwas.state, lwas.state != nil
}

// StateMust returns the state for [LinuxWebAppSlot]. Panics if the state is nil.
func (lwas *LinuxWebAppSlot) StateMust() *linuxWebAppSlotState {
	if lwas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lwas.Type(), lwas.LocalName()))
	}
	return lwas.state
}

// LinuxWebAppSlotArgs contains the configurations for azurerm_linux_web_app_slot.
type LinuxWebAppSlotArgs struct {
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
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VirtualNetworkSubnetId: string, optional
	VirtualNetworkSubnetId terra.StringValue `hcl:"virtual_network_subnet_id,attr"`
	// ZipDeployFile: string, optional
	ZipDeployFile terra.StringValue `hcl:"zip_deploy_file,attr"`
	// SiteCredential: min=0
	SiteCredential []linuxwebappslot.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// AuthSettings: optional
	AuthSettings *linuxwebappslot.AuthSettings `hcl:"auth_settings,block"`
	// AuthSettingsV2: optional
	AuthSettingsV2 *linuxwebappslot.AuthSettingsV2 `hcl:"auth_settings_v2,block"`
	// Backup: optional
	Backup *linuxwebappslot.Backup `hcl:"backup,block"`
	// ConnectionString: min=0
	ConnectionString []linuxwebappslot.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: optional
	Identity *linuxwebappslot.Identity `hcl:"identity,block"`
	// Logs: optional
	Logs *linuxwebappslot.Logs `hcl:"logs,block"`
	// SiteConfig: required
	SiteConfig *linuxwebappslot.SiteConfig `hcl:"site_config,block" validate:"required"`
	// StorageAccount: min=0
	StorageAccount []linuxwebappslot.StorageAccount `hcl:"storage_account,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *linuxwebappslot.Timeouts `hcl:"timeouts,block"`
}
type linuxWebAppSlotAttributes struct {
	ref terra.Reference
}

// AppMetadata returns a reference to field app_metadata of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) AppMetadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lwas.ref.Append("app_metadata"))
}

// AppServiceId returns a reference to field app_service_id of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) AppServiceId() terra.StringValue {
	return terra.ReferenceAsString(lwas.ref.Append("app_service_id"))
}

// AppSettings returns a reference to field app_settings of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lwas.ref.Append("app_settings"))
}

// ClientAffinityEnabled returns a reference to field client_affinity_enabled of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) ClientAffinityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lwas.ref.Append("client_affinity_enabled"))
}

// ClientCertificateEnabled returns a reference to field client_certificate_enabled of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) ClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lwas.ref.Append("client_certificate_enabled"))
}

// ClientCertificateExclusionPaths returns a reference to field client_certificate_exclusion_paths of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) ClientCertificateExclusionPaths() terra.StringValue {
	return terra.ReferenceAsString(lwas.ref.Append("client_certificate_exclusion_paths"))
}

// ClientCertificateMode returns a reference to field client_certificate_mode of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceAsString(lwas.ref.Append("client_certificate_mode"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(lwas.ref.Append("custom_domain_verification_id"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(lwas.ref.Append("default_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(lwas.ref.Append("enabled"))
}

// HostingEnvironmentId returns a reference to field hosting_environment_id of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) HostingEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(lwas.ref.Append("hosting_environment_id"))
}

// HttpsOnly returns a reference to field https_only of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(lwas.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lwas.ref.Append("id"))
}

// KeyVaultReferenceIdentityId returns a reference to field key_vault_reference_identity_id of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) KeyVaultReferenceIdentityId() terra.StringValue {
	return terra.ReferenceAsString(lwas.ref.Append("key_vault_reference_identity_id"))
}

// Kind returns a reference to field kind of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(lwas.ref.Append("kind"))
}

// Name returns a reference to field name of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lwas.ref.Append("name"))
}

// OutboundIpAddressList returns a reference to field outbound_ip_address_list of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) OutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lwas.ref.Append("outbound_ip_address_list"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(lwas.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddressList returns a reference to field possible_outbound_ip_address_list of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) PossibleOutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lwas.ref.Append("possible_outbound_ip_address_list"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(lwas.ref.Append("possible_outbound_ip_addresses"))
}

// ServicePlanId returns a reference to field service_plan_id of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) ServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(lwas.ref.Append("service_plan_id"))
}

// Tags returns a reference to field tags of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lwas.ref.Append("tags"))
}

// VirtualNetworkSubnetId returns a reference to field virtual_network_subnet_id of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(lwas.ref.Append("virtual_network_subnet_id"))
}

// ZipDeployFile returns a reference to field zip_deploy_file of azurerm_linux_web_app_slot.
func (lwas linuxWebAppSlotAttributes) ZipDeployFile() terra.StringValue {
	return terra.ReferenceAsString(lwas.ref.Append("zip_deploy_file"))
}

func (lwas linuxWebAppSlotAttributes) SiteCredential() terra.ListValue[linuxwebappslot.SiteCredentialAttributes] {
	return terra.ReferenceAsList[linuxwebappslot.SiteCredentialAttributes](lwas.ref.Append("site_credential"))
}

func (lwas linuxWebAppSlotAttributes) AuthSettings() terra.ListValue[linuxwebappslot.AuthSettingsAttributes] {
	return terra.ReferenceAsList[linuxwebappslot.AuthSettingsAttributes](lwas.ref.Append("auth_settings"))
}

func (lwas linuxWebAppSlotAttributes) AuthSettingsV2() terra.ListValue[linuxwebappslot.AuthSettingsV2Attributes] {
	return terra.ReferenceAsList[linuxwebappslot.AuthSettingsV2Attributes](lwas.ref.Append("auth_settings_v2"))
}

func (lwas linuxWebAppSlotAttributes) Backup() terra.ListValue[linuxwebappslot.BackupAttributes] {
	return terra.ReferenceAsList[linuxwebappslot.BackupAttributes](lwas.ref.Append("backup"))
}

func (lwas linuxWebAppSlotAttributes) ConnectionString() terra.SetValue[linuxwebappslot.ConnectionStringAttributes] {
	return terra.ReferenceAsSet[linuxwebappslot.ConnectionStringAttributes](lwas.ref.Append("connection_string"))
}

func (lwas linuxWebAppSlotAttributes) Identity() terra.ListValue[linuxwebappslot.IdentityAttributes] {
	return terra.ReferenceAsList[linuxwebappslot.IdentityAttributes](lwas.ref.Append("identity"))
}

func (lwas linuxWebAppSlotAttributes) Logs() terra.ListValue[linuxwebappslot.LogsAttributes] {
	return terra.ReferenceAsList[linuxwebappslot.LogsAttributes](lwas.ref.Append("logs"))
}

func (lwas linuxWebAppSlotAttributes) SiteConfig() terra.ListValue[linuxwebappslot.SiteConfigAttributes] {
	return terra.ReferenceAsList[linuxwebappslot.SiteConfigAttributes](lwas.ref.Append("site_config"))
}

func (lwas linuxWebAppSlotAttributes) StorageAccount() terra.SetValue[linuxwebappslot.StorageAccountAttributes] {
	return terra.ReferenceAsSet[linuxwebappslot.StorageAccountAttributes](lwas.ref.Append("storage_account"))
}

func (lwas linuxWebAppSlotAttributes) Timeouts() linuxwebappslot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[linuxwebappslot.TimeoutsAttributes](lwas.ref.Append("timeouts"))
}

type linuxWebAppSlotState struct {
	AppMetadata                     map[string]string                       `json:"app_metadata"`
	AppServiceId                    string                                  `json:"app_service_id"`
	AppSettings                     map[string]string                       `json:"app_settings"`
	ClientAffinityEnabled           bool                                    `json:"client_affinity_enabled"`
	ClientCertificateEnabled        bool                                    `json:"client_certificate_enabled"`
	ClientCertificateExclusionPaths string                                  `json:"client_certificate_exclusion_paths"`
	ClientCertificateMode           string                                  `json:"client_certificate_mode"`
	CustomDomainVerificationId      string                                  `json:"custom_domain_verification_id"`
	DefaultHostname                 string                                  `json:"default_hostname"`
	Enabled                         bool                                    `json:"enabled"`
	HostingEnvironmentId            string                                  `json:"hosting_environment_id"`
	HttpsOnly                       bool                                    `json:"https_only"`
	Id                              string                                  `json:"id"`
	KeyVaultReferenceIdentityId     string                                  `json:"key_vault_reference_identity_id"`
	Kind                            string                                  `json:"kind"`
	Name                            string                                  `json:"name"`
	OutboundIpAddressList           []string                                `json:"outbound_ip_address_list"`
	OutboundIpAddresses             string                                  `json:"outbound_ip_addresses"`
	PossibleOutboundIpAddressList   []string                                `json:"possible_outbound_ip_address_list"`
	PossibleOutboundIpAddresses     string                                  `json:"possible_outbound_ip_addresses"`
	ServicePlanId                   string                                  `json:"service_plan_id"`
	Tags                            map[string]string                       `json:"tags"`
	VirtualNetworkSubnetId          string                                  `json:"virtual_network_subnet_id"`
	ZipDeployFile                   string                                  `json:"zip_deploy_file"`
	SiteCredential                  []linuxwebappslot.SiteCredentialState   `json:"site_credential"`
	AuthSettings                    []linuxwebappslot.AuthSettingsState     `json:"auth_settings"`
	AuthSettingsV2                  []linuxwebappslot.AuthSettingsV2State   `json:"auth_settings_v2"`
	Backup                          []linuxwebappslot.BackupState           `json:"backup"`
	ConnectionString                []linuxwebappslot.ConnectionStringState `json:"connection_string"`
	Identity                        []linuxwebappslot.IdentityState         `json:"identity"`
	Logs                            []linuxwebappslot.LogsState             `json:"logs"`
	SiteConfig                      []linuxwebappslot.SiteConfigState       `json:"site_config"`
	StorageAccount                  []linuxwebappslot.StorageAccountState   `json:"storage_account"`
	Timeouts                        *linuxwebappslot.TimeoutsState          `json:"timeouts"`
}
