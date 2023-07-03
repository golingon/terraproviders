// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datalogicappstandard "github.com/golingon/terraproviders/azurerm/3.63.0/datalogicappstandard"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLogicAppStandard creates a new instance of [DataLogicAppStandard].
func NewDataLogicAppStandard(name string, args DataLogicAppStandardArgs) *DataLogicAppStandard {
	return &DataLogicAppStandard{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLogicAppStandard)(nil)

// DataLogicAppStandard represents the Terraform data resource azurerm_logic_app_standard.
type DataLogicAppStandard struct {
	Name string
	Args DataLogicAppStandardArgs
}

// DataSource returns the Terraform object type for [DataLogicAppStandard].
func (las *DataLogicAppStandard) DataSource() string {
	return "azurerm_logic_app_standard"
}

// LocalName returns the local name for [DataLogicAppStandard].
func (las *DataLogicAppStandard) LocalName() string {
	return las.Name
}

// Configuration returns the configuration (args) for [DataLogicAppStandard].
func (las *DataLogicAppStandard) Configuration() interface{} {
	return las.Args
}

// Attributes returns the attributes for [DataLogicAppStandard].
func (las *DataLogicAppStandard) Attributes() dataLogicAppStandardAttributes {
	return dataLogicAppStandardAttributes{ref: terra.ReferenceDataResource(las)}
}

// DataLogicAppStandardArgs contains the configurations for azurerm_logic_app_standard.
type DataLogicAppStandardArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ConnectionString: min=0
	ConnectionString []datalogicappstandard.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// Identity: min=0
	Identity []datalogicappstandard.Identity `hcl:"identity,block" validate:"min=0"`
	// SiteCredential: min=0
	SiteCredential []datalogicappstandard.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// SiteConfig: optional
	SiteConfig *datalogicappstandard.SiteConfig `hcl:"site_config,block"`
	// Timeouts: optional
	Timeouts *datalogicappstandard.Timeouts `hcl:"timeouts,block"`
}
type dataLogicAppStandardAttributes struct {
	ref terra.Reference
}

// AppServicePlanId returns a reference to field app_service_plan_id of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) AppServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("app_service_plan_id"))
}

// AppSettings returns a reference to field app_settings of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](las.ref.Append("app_settings"))
}

// BundleVersion returns a reference to field bundle_version of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) BundleVersion() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("bundle_version"))
}

// ClientAffinityEnabled returns a reference to field client_affinity_enabled of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) ClientAffinityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(las.ref.Append("client_affinity_enabled"))
}

// ClientCertificateMode returns a reference to field client_certificate_mode of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) ClientCertificateMode() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("client_certificate_mode"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("custom_domain_verification_id"))
}

// DefaultHostname returns a reference to field default_hostname of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("default_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(las.ref.Append("enabled"))
}

// HttpsOnly returns a reference to field https_only of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(las.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("name"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("possible_outbound_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("resource_group_name"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("storage_account_access_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("storage_account_name"))
}

// StorageAccountShareName returns a reference to field storage_account_share_name of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) StorageAccountShareName() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("storage_account_share_name"))
}

// Tags returns a reference to field tags of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](las.ref.Append("tags"))
}

// UseExtensionBundle returns a reference to field use_extension_bundle of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) UseExtensionBundle() terra.BoolValue {
	return terra.ReferenceAsBool(las.ref.Append("use_extension_bundle"))
}

// Version returns a reference to field version of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("version"))
}

// VirtualNetworkSubnetId returns a reference to field virtual_network_subnet_id of azurerm_logic_app_standard.
func (las dataLogicAppStandardAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(las.ref.Append("virtual_network_subnet_id"))
}

func (las dataLogicAppStandardAttributes) ConnectionString() terra.SetValue[datalogicappstandard.ConnectionStringAttributes] {
	return terra.ReferenceAsSet[datalogicappstandard.ConnectionStringAttributes](las.ref.Append("connection_string"))
}

func (las dataLogicAppStandardAttributes) Identity() terra.ListValue[datalogicappstandard.IdentityAttributes] {
	return terra.ReferenceAsList[datalogicappstandard.IdentityAttributes](las.ref.Append("identity"))
}

func (las dataLogicAppStandardAttributes) SiteCredential() terra.ListValue[datalogicappstandard.SiteCredentialAttributes] {
	return terra.ReferenceAsList[datalogicappstandard.SiteCredentialAttributes](las.ref.Append("site_credential"))
}

func (las dataLogicAppStandardAttributes) SiteConfig() terra.ListValue[datalogicappstandard.SiteConfigAttributes] {
	return terra.ReferenceAsList[datalogicappstandard.SiteConfigAttributes](las.ref.Append("site_config"))
}

func (las dataLogicAppStandardAttributes) Timeouts() datalogicappstandard.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datalogicappstandard.TimeoutsAttributes](las.ref.Append("timeouts"))
}
