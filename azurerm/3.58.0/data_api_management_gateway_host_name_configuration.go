// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataapimanagementgatewayhostnameconfiguration "github.com/golingon/terraproviders/azurerm/3.58.0/dataapimanagementgatewayhostnameconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataApiManagementGatewayHostNameConfiguration creates a new instance of [DataApiManagementGatewayHostNameConfiguration].
func NewDataApiManagementGatewayHostNameConfiguration(name string, args DataApiManagementGatewayHostNameConfigurationArgs) *DataApiManagementGatewayHostNameConfiguration {
	return &DataApiManagementGatewayHostNameConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApiManagementGatewayHostNameConfiguration)(nil)

// DataApiManagementGatewayHostNameConfiguration represents the Terraform data resource azurerm_api_management_gateway_host_name_configuration.
type DataApiManagementGatewayHostNameConfiguration struct {
	Name string
	Args DataApiManagementGatewayHostNameConfigurationArgs
}

// DataSource returns the Terraform object type for [DataApiManagementGatewayHostNameConfiguration].
func (amghnc *DataApiManagementGatewayHostNameConfiguration) DataSource() string {
	return "azurerm_api_management_gateway_host_name_configuration"
}

// LocalName returns the local name for [DataApiManagementGatewayHostNameConfiguration].
func (amghnc *DataApiManagementGatewayHostNameConfiguration) LocalName() string {
	return amghnc.Name
}

// Configuration returns the configuration (args) for [DataApiManagementGatewayHostNameConfiguration].
func (amghnc *DataApiManagementGatewayHostNameConfiguration) Configuration() interface{} {
	return amghnc.Args
}

// Attributes returns the attributes for [DataApiManagementGatewayHostNameConfiguration].
func (amghnc *DataApiManagementGatewayHostNameConfiguration) Attributes() dataApiManagementGatewayHostNameConfigurationAttributes {
	return dataApiManagementGatewayHostNameConfigurationAttributes{ref: terra.ReferenceDataResource(amghnc)}
}

// DataApiManagementGatewayHostNameConfigurationArgs contains the configurations for azurerm_api_management_gateway_host_name_configuration.
type DataApiManagementGatewayHostNameConfigurationArgs struct {
	// ApiManagementId: string, required
	ApiManagementId terra.StringValue `hcl:"api_management_id,attr" validate:"required"`
	// GatewayName: string, required
	GatewayName terra.StringValue `hcl:"gateway_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataapimanagementgatewayhostnameconfiguration.Timeouts `hcl:"timeouts,block"`
}
type dataApiManagementGatewayHostNameConfigurationAttributes struct {
	ref terra.Reference
}

// ApiManagementId returns a reference to field api_management_id of azurerm_api_management_gateway_host_name_configuration.
func (amghnc dataApiManagementGatewayHostNameConfigurationAttributes) ApiManagementId() terra.StringValue {
	return terra.ReferenceAsString(amghnc.ref.Append("api_management_id"))
}

// CertificateId returns a reference to field certificate_id of azurerm_api_management_gateway_host_name_configuration.
func (amghnc dataApiManagementGatewayHostNameConfigurationAttributes) CertificateId() terra.StringValue {
	return terra.ReferenceAsString(amghnc.ref.Append("certificate_id"))
}

// GatewayName returns a reference to field gateway_name of azurerm_api_management_gateway_host_name_configuration.
func (amghnc dataApiManagementGatewayHostNameConfigurationAttributes) GatewayName() terra.StringValue {
	return terra.ReferenceAsString(amghnc.ref.Append("gateway_name"))
}

// HostName returns a reference to field host_name of azurerm_api_management_gateway_host_name_configuration.
func (amghnc dataApiManagementGatewayHostNameConfigurationAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(amghnc.ref.Append("host_name"))
}

// Http2Enabled returns a reference to field http2_enabled of azurerm_api_management_gateway_host_name_configuration.
func (amghnc dataApiManagementGatewayHostNameConfigurationAttributes) Http2Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(amghnc.ref.Append("http2_enabled"))
}

// Id returns a reference to field id of azurerm_api_management_gateway_host_name_configuration.
func (amghnc dataApiManagementGatewayHostNameConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amghnc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_api_management_gateway_host_name_configuration.
func (amghnc dataApiManagementGatewayHostNameConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amghnc.ref.Append("name"))
}

// RequestClientCertificateEnabled returns a reference to field request_client_certificate_enabled of azurerm_api_management_gateway_host_name_configuration.
func (amghnc dataApiManagementGatewayHostNameConfigurationAttributes) RequestClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(amghnc.ref.Append("request_client_certificate_enabled"))
}

// Tls10Enabled returns a reference to field tls10_enabled of azurerm_api_management_gateway_host_name_configuration.
func (amghnc dataApiManagementGatewayHostNameConfigurationAttributes) Tls10Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(amghnc.ref.Append("tls10_enabled"))
}

// Tls11Enabled returns a reference to field tls11_enabled of azurerm_api_management_gateway_host_name_configuration.
func (amghnc dataApiManagementGatewayHostNameConfigurationAttributes) Tls11Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(amghnc.ref.Append("tls11_enabled"))
}

func (amghnc dataApiManagementGatewayHostNameConfigurationAttributes) Timeouts() dataapimanagementgatewayhostnameconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataapimanagementgatewayhostnameconfiguration.TimeoutsAttributes](amghnc.ref.Append("timeouts"))
}
