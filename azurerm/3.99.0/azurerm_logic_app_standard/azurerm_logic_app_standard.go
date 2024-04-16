// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_logic_app_standard

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

// Resource represents the Terraform resource azurerm_logic_app_standard.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermLogicAppStandardState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (alas *Resource) Type() string {
	return "azurerm_logic_app_standard"
}

// LocalName returns the local name for [Resource].
func (alas *Resource) LocalName() string {
	return alas.Name
}

// Configuration returns the configuration (args) for [Resource].
func (alas *Resource) Configuration() interface{} {
	return alas.Args
}

// DependOn is used for other resources to depend on [Resource].
func (alas *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(alas)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (alas *Resource) Dependencies() terra.Dependencies {
	return alas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (alas *Resource) LifecycleManagement() *terra.Lifecycle {
	return alas.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (alas *Resource) Attributes() azurermLogicAppStandardAttributes {
	return azurermLogicAppStandardAttributes{ref: terra.ReferenceResource(alas)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (alas *Resource) ImportState(state io.Reader) error {
	alas.state = &azurermLogicAppStandardState{}
	if err := json.NewDecoder(state).Decode(alas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", alas.Type(), alas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (alas *Resource) State() (*azurermLogicAppStandardState, bool) {
	return alas.state, alas.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (alas *Resource) StateMust() *azurermLogicAppStandardState {
	if alas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", alas.Type(), alas.LocalName()))
	}
	return alas.state
}

// Args contains the configurations for azurerm_logic_app_standard.
type Args struct {
	// AppServicePlanId: string, required
	AppServicePlanId terra.StringValue `hcl:"app_service_plan_id,attr" validate:"required"`
	// AppSettings: map of string, optional
	AppSettings terra.MapValue[terra.StringValue] `hcl:"app_settings,attr"`
	// BundleVersion: string, optional
	BundleVersion terra.StringValue `hcl:"bundle_version,attr"`
	// ClientAffinityEnabled: bool, optional
	ClientAffinityEnabled terra.BoolValue `hcl:"client_affinity_enabled,attr"`
	// ClientCertificateMode: string, optional
	ClientCertificateMode terra.StringValue `hcl:"client_certificate_mode,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// HttpsOnly: bool, optional
	HttpsOnly terra.BoolValue `hcl:"https_only,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StorageAccountAccessKey: string, required
	StorageAccountAccessKey terra.StringValue `hcl:"storage_account_access_key,attr" validate:"required"`
	// StorageAccountName: string, required
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr" validate:"required"`
	// StorageAccountShareName: string, optional
	StorageAccountShareName terra.StringValue `hcl:"storage_account_share_name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// UseExtensionBundle: bool, optional
	UseExtensionBundle terra.BoolValue `hcl:"use_extension_bundle,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// VirtualNetworkSubnetId: string, optional
	VirtualNetworkSubnetId terra.StringValue `hcl:"virtual_network_subnet_id,attr"`
	// ConnectionString: min=0
	ConnectionString []ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: optional
	Identity *Identity `hcl:"identity,block"`
	// SiteConfig: optional
	SiteConfig *SiteConfig `hcl:"site_config,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermLogicAppStandardAttributes struct {
	ref terra.Reference
}

// AppServicePlanId returns a reference to field app_service_plan_id of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) AppServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("app_service_plan_id"))
}

// AppSettings returns a reference to field app_settings of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](alas.ref.Append("app_settings"))
}

// BundleVersion returns a reference to field bundle_version of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) BundleVersion() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("bundle_version"))
}

// ClientAffinityEnabled returns a reference to field client_affinity_enabled of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) ClientAffinityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(alas.ref.Append("client_affinity_enabled"))
}

// ClientCertificateMode returns a reference to field client_certificate_mode of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("client_certificate_mode"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("custom_domain_verification_id"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("default_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(alas.ref.Append("enabled"))
}

// HttpsOnly returns a reference to field https_only of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(alas.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("name"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("possible_outbound_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("resource_group_name"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("storage_account_access_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("storage_account_name"))
}

// StorageAccountShareName returns a reference to field storage_account_share_name of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) StorageAccountShareName() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("storage_account_share_name"))
}

// Tags returns a reference to field tags of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](alas.ref.Append("tags"))
}

// UseExtensionBundle returns a reference to field use_extension_bundle of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) UseExtensionBundle() terra.BoolValue {
	return terra.ReferenceAsBool(alas.ref.Append("use_extension_bundle"))
}

// Version returns a reference to field version of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("version"))
}

// VirtualNetworkSubnetId returns a reference to field virtual_network_subnet_id of azurerm_logic_app_standard.
func (alas azurermLogicAppStandardAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(alas.ref.Append("virtual_network_subnet_id"))
}

func (alas azurermLogicAppStandardAttributes) SiteCredential() terra.ListValue[SiteCredentialAttributes] {
	return terra.ReferenceAsList[SiteCredentialAttributes](alas.ref.Append("site_credential"))
}

func (alas azurermLogicAppStandardAttributes) ConnectionString() terra.SetValue[ConnectionStringAttributes] {
	return terra.ReferenceAsSet[ConnectionStringAttributes](alas.ref.Append("connection_string"))
}

func (alas azurermLogicAppStandardAttributes) Identity() terra.ListValue[IdentityAttributes] {
	return terra.ReferenceAsList[IdentityAttributes](alas.ref.Append("identity"))
}

func (alas azurermLogicAppStandardAttributes) SiteConfig() terra.ListValue[SiteConfigAttributes] {
	return terra.ReferenceAsList[SiteConfigAttributes](alas.ref.Append("site_config"))
}

func (alas azurermLogicAppStandardAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](alas.ref.Append("timeouts"))
}

type azurermLogicAppStandardState struct {
	AppServicePlanId            string                  `json:"app_service_plan_id"`
	AppSettings                 map[string]string       `json:"app_settings"`
	BundleVersion               string                  `json:"bundle_version"`
	ClientAffinityEnabled       bool                    `json:"client_affinity_enabled"`
	ClientCertificateMode       string                  `json:"client_certificate_mode"`
	CustomDomainVerificationId  string                  `json:"custom_domain_verification_id"`
	DefaultHostname             string                  `json:"default_hostname"`
	Enabled                     bool                    `json:"enabled"`
	HttpsOnly                   bool                    `json:"https_only"`
	Id                          string                  `json:"id"`
	Kind                        string                  `json:"kind"`
	Location                    string                  `json:"location"`
	Name                        string                  `json:"name"`
	OutboundIpAddresses         string                  `json:"outbound_ip_addresses"`
	PossibleOutboundIpAddresses string                  `json:"possible_outbound_ip_addresses"`
	ResourceGroupName           string                  `json:"resource_group_name"`
	StorageAccountAccessKey     string                  `json:"storage_account_access_key"`
	StorageAccountName          string                  `json:"storage_account_name"`
	StorageAccountShareName     string                  `json:"storage_account_share_name"`
	Tags                        map[string]string       `json:"tags"`
	UseExtensionBundle          bool                    `json:"use_extension_bundle"`
	Version                     string                  `json:"version"`
	VirtualNetworkSubnetId      string                  `json:"virtual_network_subnet_id"`
	SiteCredential              []SiteCredentialState   `json:"site_credential"`
	ConnectionString            []ConnectionStringState `json:"connection_string"`
	Identity                    []IdentityState         `json:"identity"`
	SiteConfig                  []SiteConfigState       `json:"site_config"`
	Timeouts                    *TimeoutsState          `json:"timeouts"`
}
