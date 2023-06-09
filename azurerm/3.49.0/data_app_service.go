// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataappservice "github.com/golingon/terraproviders/azurerm/3.49.0/dataappservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAppService creates a new instance of [DataAppService].
func NewDataAppService(name string, args DataAppServiceArgs) *DataAppService {
	return &DataAppService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppService)(nil)

// DataAppService represents the Terraform data resource azurerm_app_service.
type DataAppService struct {
	Name string
	Args DataAppServiceArgs
}

// DataSource returns the Terraform object type for [DataAppService].
func (as *DataAppService) DataSource() string {
	return "azurerm_app_service"
}

// LocalName returns the local name for [DataAppService].
func (as *DataAppService) LocalName() string {
	return as.Name
}

// Configuration returns the configuration (args) for [DataAppService].
func (as *DataAppService) Configuration() interface{} {
	return as.Args
}

// Attributes returns the attributes for [DataAppService].
func (as *DataAppService) Attributes() dataAppServiceAttributes {
	return dataAppServiceAttributes{ref: terra.ReferenceDataResource(as)}
}

// DataAppServiceArgs contains the configurations for azurerm_app_service.
type DataAppServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ConnectionString: min=0
	ConnectionString []dataappservice.ConnectionString `hcl:"connection_string,block" validate:"min=0"`
	// SiteConfig: min=0
	SiteConfig []dataappservice.SiteConfig `hcl:"site_config,block" validate:"min=0"`
	// SiteCredential: min=0
	SiteCredential []dataappservice.SiteCredential `hcl:"site_credential,block" validate:"min=0"`
	// SourceControl: min=0
	SourceControl []dataappservice.SourceControl `hcl:"source_control,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataappservice.Timeouts `hcl:"timeouts,block"`
}
type dataAppServiceAttributes struct {
	ref terra.Reference
}

// AppServicePlanId returns a reference to field app_service_plan_id of azurerm_app_service.
func (as dataAppServiceAttributes) AppServicePlanId() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("app_service_plan_id"))
}

// AppSettings returns a reference to field app_settings of azurerm_app_service.
func (as dataAppServiceAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](as.ref.Append("app_settings"))
}

// ClientAffinityEnabled returns a reference to field client_affinity_enabled of azurerm_app_service.
func (as dataAppServiceAttributes) ClientAffinityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(as.ref.Append("client_affinity_enabled"))
}

// ClientCertEnabled returns a reference to field client_cert_enabled of azurerm_app_service.
func (as dataAppServiceAttributes) ClientCertEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(as.ref.Append("client_cert_enabled"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_app_service.
func (as dataAppServiceAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("custom_domain_verification_id"))
}

// DefaultSiteHostname returns a reference to field default_site_hostname of azurerm_app_service.
func (as dataAppServiceAttributes) DefaultSiteHostname() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("default_site_hostname"))
}

// Enabled returns a reference to field enabled of azurerm_app_service.
func (as dataAppServiceAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(as.ref.Append("enabled"))
}

// HttpsOnly returns a reference to field https_only of azurerm_app_service.
func (as dataAppServiceAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(as.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_app_service.
func (as dataAppServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_app_service.
func (as dataAppServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_app_service.
func (as dataAppServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("name"))
}

// OutboundIpAddressList returns a reference to field outbound_ip_address_list of azurerm_app_service.
func (as dataAppServiceAttributes) OutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](as.ref.Append("outbound_ip_address_list"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_app_service.
func (as dataAppServiceAttributes) OutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("outbound_ip_addresses"))
}

// PossibleOutboundIpAddressList returns a reference to field possible_outbound_ip_address_list of azurerm_app_service.
func (as dataAppServiceAttributes) PossibleOutboundIpAddressList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](as.ref.Append("possible_outbound_ip_address_list"))
}

// PossibleOutboundIpAddresses returns a reference to field possible_outbound_ip_addresses of azurerm_app_service.
func (as dataAppServiceAttributes) PossibleOutboundIpAddresses() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("possible_outbound_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_app_service.
func (as dataAppServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_app_service.
func (as dataAppServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](as.ref.Append("tags"))
}

func (as dataAppServiceAttributes) ConnectionString() terra.ListValue[dataappservice.ConnectionStringAttributes] {
	return terra.ReferenceAsList[dataappservice.ConnectionStringAttributes](as.ref.Append("connection_string"))
}

func (as dataAppServiceAttributes) SiteConfig() terra.ListValue[dataappservice.SiteConfigAttributes] {
	return terra.ReferenceAsList[dataappservice.SiteConfigAttributes](as.ref.Append("site_config"))
}

func (as dataAppServiceAttributes) SiteCredential() terra.ListValue[dataappservice.SiteCredentialAttributes] {
	return terra.ReferenceAsList[dataappservice.SiteCredentialAttributes](as.ref.Append("site_credential"))
}

func (as dataAppServiceAttributes) SourceControl() terra.ListValue[dataappservice.SourceControlAttributes] {
	return terra.ReferenceAsList[dataappservice.SourceControlAttributes](as.ref.Append("source_control"))
}

func (as dataAppServiceAttributes) Timeouts() dataappservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataappservice.TimeoutsAttributes](as.ref.Append("timeouts"))
}
