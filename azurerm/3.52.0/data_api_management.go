// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataapimanagement "github.com/golingon/terraproviders/azurerm/3.52.0/dataapimanagement"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataApiManagement creates a new instance of [DataApiManagement].
func NewDataApiManagement(name string, args DataApiManagementArgs) *DataApiManagement {
	return &DataApiManagement{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApiManagement)(nil)

// DataApiManagement represents the Terraform data resource azurerm_api_management.
type DataApiManagement struct {
	Name string
	Args DataApiManagementArgs
}

// DataSource returns the Terraform object type for [DataApiManagement].
func (am *DataApiManagement) DataSource() string {
	return "azurerm_api_management"
}

// LocalName returns the local name for [DataApiManagement].
func (am *DataApiManagement) LocalName() string {
	return am.Name
}

// Configuration returns the configuration (args) for [DataApiManagement].
func (am *DataApiManagement) Configuration() interface{} {
	return am.Args
}

// Attributes returns the attributes for [DataApiManagement].
func (am *DataApiManagement) Attributes() dataApiManagementAttributes {
	return dataApiManagementAttributes{ref: terra.ReferenceDataResource(am)}
}

// DataApiManagementArgs contains the configurations for azurerm_api_management.
type DataApiManagementArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// AdditionalLocation: min=0
	AdditionalLocation []dataapimanagement.AdditionalLocation `hcl:"additional_location,block" validate:"min=0"`
	// HostnameConfiguration: min=0
	HostnameConfiguration []dataapimanagement.HostnameConfiguration `hcl:"hostname_configuration,block" validate:"min=0"`
	// Identity: min=0
	Identity []dataapimanagement.Identity `hcl:"identity,block" validate:"min=0"`
	// TenantAccess: min=0
	TenantAccess []dataapimanagement.TenantAccess `hcl:"tenant_access,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataapimanagement.Timeouts `hcl:"timeouts,block"`
}
type dataApiManagementAttributes struct {
	ref terra.Reference
}

// DeveloperPortalUrl returns a reference to field developer_portal_url of azurerm_api_management.
func (am dataApiManagementAttributes) DeveloperPortalUrl() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("developer_portal_url"))
}

// GatewayRegionalUrl returns a reference to field gateway_regional_url of azurerm_api_management.
func (am dataApiManagementAttributes) GatewayRegionalUrl() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("gateway_regional_url"))
}

// GatewayUrl returns a reference to field gateway_url of azurerm_api_management.
func (am dataApiManagementAttributes) GatewayUrl() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("gateway_url"))
}

// Id returns a reference to field id of azurerm_api_management.
func (am dataApiManagementAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_api_management.
func (am dataApiManagementAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("location"))
}

// ManagementApiUrl returns a reference to field management_api_url of azurerm_api_management.
func (am dataApiManagementAttributes) ManagementApiUrl() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("management_api_url"))
}

// Name returns a reference to field name of azurerm_api_management.
func (am dataApiManagementAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("name"))
}

// NotificationSenderEmail returns a reference to field notification_sender_email of azurerm_api_management.
func (am dataApiManagementAttributes) NotificationSenderEmail() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("notification_sender_email"))
}

// PortalUrl returns a reference to field portal_url of azurerm_api_management.
func (am dataApiManagementAttributes) PortalUrl() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("portal_url"))
}

// PrivateIpAddresses returns a reference to field private_ip_addresses of azurerm_api_management.
func (am dataApiManagementAttributes) PrivateIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](am.ref.Append("private_ip_addresses"))
}

// PublicIpAddressId returns a reference to field public_ip_address_id of azurerm_api_management.
func (am dataApiManagementAttributes) PublicIpAddressId() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("public_ip_address_id"))
}

// PublicIpAddresses returns a reference to field public_ip_addresses of azurerm_api_management.
func (am dataApiManagementAttributes) PublicIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](am.ref.Append("public_ip_addresses"))
}

// PublisherEmail returns a reference to field publisher_email of azurerm_api_management.
func (am dataApiManagementAttributes) PublisherEmail() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("publisher_email"))
}

// PublisherName returns a reference to field publisher_name of azurerm_api_management.
func (am dataApiManagementAttributes) PublisherName() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("publisher_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management.
func (am dataApiManagementAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("resource_group_name"))
}

// ScmUrl returns a reference to field scm_url of azurerm_api_management.
func (am dataApiManagementAttributes) ScmUrl() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("scm_url"))
}

// SkuName returns a reference to field sku_name of azurerm_api_management.
func (am dataApiManagementAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_api_management.
func (am dataApiManagementAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](am.ref.Append("tags"))
}

func (am dataApiManagementAttributes) AdditionalLocation() terra.ListValue[dataapimanagement.AdditionalLocationAttributes] {
	return terra.ReferenceAsList[dataapimanagement.AdditionalLocationAttributes](am.ref.Append("additional_location"))
}

func (am dataApiManagementAttributes) HostnameConfiguration() terra.ListValue[dataapimanagement.HostnameConfigurationAttributes] {
	return terra.ReferenceAsList[dataapimanagement.HostnameConfigurationAttributes](am.ref.Append("hostname_configuration"))
}

func (am dataApiManagementAttributes) Identity() terra.ListValue[dataapimanagement.IdentityAttributes] {
	return terra.ReferenceAsList[dataapimanagement.IdentityAttributes](am.ref.Append("identity"))
}

func (am dataApiManagementAttributes) TenantAccess() terra.ListValue[dataapimanagement.TenantAccessAttributes] {
	return terra.ReferenceAsList[dataapimanagement.TenantAccessAttributes](am.ref.Append("tenant_access"))
}

func (am dataApiManagementAttributes) Timeouts() dataapimanagement.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataapimanagement.TimeoutsAttributes](am.ref.Append("timeouts"))
}
