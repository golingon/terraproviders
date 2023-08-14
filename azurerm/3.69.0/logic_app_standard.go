// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	logicappstandard "github.com/golingon/terraproviders/azurerm/3.69.0/logicappstandard"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogicAppStandard creates a new instance of [LogicAppStandard].
func NewLogicAppStandard(name string, args LogicAppStandardArgs) *LogicAppStandard {
	return &LogicAppStandard{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogicAppStandard)(nil)

// LogicAppStandard represents the Terraform resource azurerm_logic_app_standard.
type LogicAppStandard struct {
	Name      string
	Args      LogicAppStandardArgs
	state     *logicAppStandardState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogicAppStandard].
func (las *LogicAppStandard) Type() string {
	return "azurerm_logic_app_standard"
}

// LocalName returns the local name for [LogicAppStandard].
func (las *LogicAppStandard) LocalName() string {
	return las.Name
}

// Configuration returns the configuration (args) for [LogicAppStandard].
func (las *LogicAppStandard) Configuration() interface{} {
	return las.Args
}

// DependOn is used for other resources to depend on [LogicAppStandard].
func (las *LogicAppStandard) DependOn() terra.Reference {
	return terra.ReferenceResource(las)
}

// Dependencies returns the list of resources [LogicAppStandard] depends_on.
func (las *LogicAppStandard) Dependencies() terra.Dependencies {
	return las.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogicAppStandard].
func (las *LogicAppStandard) LifecycleManagement() *terra.Lifecycle {
	return las.Lifecycle
}

// Attributes returns the attributes for [LogicAppStandard].
func (las *LogicAppStandard) Attributes() logicAppStandardAttributes {
	return logicAppStandardAttributes{ref: terra.ReferenceResource(las)}
}

// ImportState imports the given attribute values into [LogicAppStandard]'s state.
func (las *LogicAppStandard) ImportState(av io.Reader) error {
	las.state = &logicAppStandardState{}
	if err := json.NewDecoder(av).Decode(las.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", las.Type(), las.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogicAppStandard] has state.
func (las *LogicAppStandard) State() (*logicAppStandardState, bool) {
	return las.state, las.state != nil
}

// StateMust returns the state for [LogicAppStandard]. Panics if the state is nil.
func (las *LogicAppStandard) StateMust() *logicAppStandardState {
	if las.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", las.Type(), las.LocalName()))
	}
	return las.state
}

// LogicAppStandardArgs contains the configurations for azurerm_logic_app_standard.
type LogicAppStandardArgs struct {
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
	// SiteCredential: min=0
	SiteCredential []logicappstandard.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// ConnectionString: min=0
	ConnectionString []logicappstandard.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: optional
	Identity *logicappstandard.Identity `hcl:"identity,block"`
	// SiteConfig: optional
	SiteConfig *logicappstandard.SiteConfig `hcl:"site_config,block"`
	// Timeouts: optional
	Timeouts *logicappstandard.Timeouts `hcl:"timeouts,block"`
}
type logicAppStandardAttributes struct {
	ref terra.Reference
}

// AppServicePlanId returns a reference to field app_service_plan_id of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) AppServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("app_service_plan_id"))
}

// AppSettings returns a reference to field app_settings of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](las.ref.Append("app_settings"))
}

// BundleVersion returns a reference to field bundle_version of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) BundleVersion() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("bundle_version"))
}

// ClientAffinityEnabled returns a reference to field client_affinity_enabled of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) ClientAffinityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(las.ref.Append("client_affinity_enabled"))
}

// ClientCertificateMode returns a reference to field client_certificate_mode of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("client_certificate_mode"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("custom_domain_verification_id"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("default_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(las.ref.Append("enabled"))
}

// HttpsOnly returns a reference to field https_only of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(las.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("name"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("possible_outbound_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("resource_group_name"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("storage_account_access_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("storage_account_name"))
}

// StorageAccountShareName returns a reference to field storage_account_share_name of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) StorageAccountShareName() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("storage_account_share_name"))
}

// Tags returns a reference to field tags of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](las.ref.Append("tags"))
}

// UseExtensionBundle returns a reference to field use_extension_bundle of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) UseExtensionBundle() terra.BoolValue {
	return terra.ReferenceAsBool(las.ref.Append("use_extension_bundle"))
}

// Version returns a reference to field version of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("version"))
}

// VirtualNetworkSubnetId returns a reference to field virtual_network_subnet_id of azurerm_logic_app_standard.
func (las logicAppStandardAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("virtual_network_subnet_id"))
}

func (las logicAppStandardAttributes) SiteCredential() terra.ListValue[logicappstandard.SiteCredentialAttributes] {
	return terra.ReferenceAsList[logicappstandard.SiteCredentialAttributes](las.ref.Append("site_credential"))
}

func (las logicAppStandardAttributes) ConnectionString() terra.SetValue[logicappstandard.ConnectionStringAttributes] {
	return terra.ReferenceAsSet[logicappstandard.ConnectionStringAttributes](las.ref.Append("connection_string"))
}

func (las logicAppStandardAttributes) Identity() terra.ListValue[logicappstandard.IdentityAttributes] {
	return terra.ReferenceAsList[logicappstandard.IdentityAttributes](las.ref.Append("identity"))
}

func (las logicAppStandardAttributes) SiteConfig() terra.ListValue[logicappstandard.SiteConfigAttributes] {
	return terra.ReferenceAsList[logicappstandard.SiteConfigAttributes](las.ref.Append("site_config"))
}

func (las logicAppStandardAttributes) Timeouts() logicappstandard.TimeoutsAttributes {
	return terra.ReferenceAsSingle[logicappstandard.TimeoutsAttributes](las.ref.Append("timeouts"))
}

type logicAppStandardState struct {
	AppServicePlanId            string                                   `json:"app_service_plan_id"`
	AppSettings                 map[string]string                        `json:"app_settings"`
	BundleVersion               string                                   `json:"bundle_version"`
	ClientAffinityEnabled       bool                                     `json:"client_affinity_enabled"`
	ClientCertificateMode       string                                   `json:"client_certificate_mode"`
	CustomDomainVerificationId  string                                   `json:"custom_domain_verification_id"`
	DefaultHostname             string                                   `json:"default_hostname"`
	Enabled                     bool                                     `json:"enabled"`
	HttpsOnly                   bool                                     `json:"https_only"`
	Id                          string                                   `json:"id"`
	Kind                        string                                   `json:"kind"`
	Location                    string                                   `json:"location"`
	Name                        string                                   `json:"name"`
	OutboundIpAddresses         string                                   `json:"outbound_ip_addresses"`
	PossibleOutboundIpAddresses string                                   `json:"possible_outbound_ip_addresses"`
	ResourceGroupName           string                                   `json:"resource_group_name"`
	StorageAccountAccessKey     string                                   `json:"storage_account_access_key"`
	StorageAccountName          string                                   `json:"storage_account_name"`
	StorageAccountShareName     string                                   `json:"storage_account_share_name"`
	Tags                        map[string]string                        `json:"tags"`
	UseExtensionBundle          bool                                     `json:"use_extension_bundle"`
	Version                     string                                   `json:"version"`
	VirtualNetworkSubnetId      string                                   `json:"virtual_network_subnet_id"`
	SiteCredential              []logicappstandard.SiteCredentialState   `json:"site_credential"`
	ConnectionString            []logicappstandard.ConnectionStringState `json:"connection_string"`
	Identity                    []logicappstandard.IdentityState         `json:"identity"`
	SiteConfig                  []logicappstandard.SiteConfigState       `json:"site_config"`
	Timeouts                    *logicappstandard.TimeoutsState          `json:"timeouts"`
}
