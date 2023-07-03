// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementgatewayhostnameconfiguration "github.com/golingon/terraproviders/azurerm/3.63.0/apimanagementgatewayhostnameconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementGatewayHostNameConfiguration creates a new instance of [ApiManagementGatewayHostNameConfiguration].
func NewApiManagementGatewayHostNameConfiguration(name string, args ApiManagementGatewayHostNameConfigurationArgs) *ApiManagementGatewayHostNameConfiguration {
	return &ApiManagementGatewayHostNameConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementGatewayHostNameConfiguration)(nil)

// ApiManagementGatewayHostNameConfiguration represents the Terraform resource azurerm_api_management_gateway_host_name_configuration.
type ApiManagementGatewayHostNameConfiguration struct {
	Name      string
	Args      ApiManagementGatewayHostNameConfigurationArgs
	state     *apiManagementGatewayHostNameConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementGatewayHostNameConfiguration].
func (amghnc *ApiManagementGatewayHostNameConfiguration) Type() string {
	return "azurerm_api_management_gateway_host_name_configuration"
}

// LocalName returns the local name for [ApiManagementGatewayHostNameConfiguration].
func (amghnc *ApiManagementGatewayHostNameConfiguration) LocalName() string {
	return amghnc.Name
}

// Configuration returns the configuration (args) for [ApiManagementGatewayHostNameConfiguration].
func (amghnc *ApiManagementGatewayHostNameConfiguration) Configuration() interface{} {
	return amghnc.Args
}

// DependOn is used for other resources to depend on [ApiManagementGatewayHostNameConfiguration].
func (amghnc *ApiManagementGatewayHostNameConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(amghnc)
}

// Dependencies returns the list of resources [ApiManagementGatewayHostNameConfiguration] depends_on.
func (amghnc *ApiManagementGatewayHostNameConfiguration) Dependencies() terra.Dependencies {
	return amghnc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementGatewayHostNameConfiguration].
func (amghnc *ApiManagementGatewayHostNameConfiguration) LifecycleManagement() *terra.Lifecycle {
	return amghnc.Lifecycle
}

// Attributes returns the attributes for [ApiManagementGatewayHostNameConfiguration].
func (amghnc *ApiManagementGatewayHostNameConfiguration) Attributes() apiManagementGatewayHostNameConfigurationAttributes {
	return apiManagementGatewayHostNameConfigurationAttributes{ref: terra.ReferenceResource(amghnc)}
}

// ImportState imports the given attribute values into [ApiManagementGatewayHostNameConfiguration]'s state.
func (amghnc *ApiManagementGatewayHostNameConfiguration) ImportState(av io.Reader) error {
	amghnc.state = &apiManagementGatewayHostNameConfigurationState{}
	if err := json.NewDecoder(av).Decode(amghnc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amghnc.Type(), amghnc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementGatewayHostNameConfiguration] has state.
func (amghnc *ApiManagementGatewayHostNameConfiguration) State() (*apiManagementGatewayHostNameConfigurationState, bool) {
	return amghnc.state, amghnc.state != nil
}

// StateMust returns the state for [ApiManagementGatewayHostNameConfiguration]. Panics if the state is nil.
func (amghnc *ApiManagementGatewayHostNameConfiguration) StateMust() *apiManagementGatewayHostNameConfigurationState {
	if amghnc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amghnc.Type(), amghnc.LocalName()))
	}
	return amghnc.state
}

// ApiManagementGatewayHostNameConfigurationArgs contains the configurations for azurerm_api_management_gateway_host_name_configuration.
type ApiManagementGatewayHostNameConfigurationArgs struct {
	// ApiManagementId: string, required
	ApiManagementId terra.StringValue `hcl:"api_management_id,attr" validate:"required"`
	// CertificateId: string, required
	CertificateId terra.StringValue `hcl:"certificate_id,attr" validate:"required"`
	// GatewayName: string, required
	GatewayName terra.StringValue `hcl:"gateway_name,attr" validate:"required"`
	// HostName: string, required
	HostName terra.StringValue `hcl:"host_name,attr" validate:"required"`
	// Http2Enabled: bool, optional
	Http2Enabled terra.BoolValue `hcl:"http2_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RequestClientCertificateEnabled: bool, optional
	RequestClientCertificateEnabled terra.BoolValue `hcl:"request_client_certificate_enabled,attr"`
	// Tls10Enabled: bool, optional
	Tls10Enabled terra.BoolValue `hcl:"tls10_enabled,attr"`
	// Tls11Enabled: bool, optional
	Tls11Enabled terra.BoolValue `hcl:"tls11_enabled,attr"`
	// Timeouts: optional
	Timeouts *apimanagementgatewayhostnameconfiguration.Timeouts `hcl:"timeouts,block"`
}
type apiManagementGatewayHostNameConfigurationAttributes struct {
	ref terra.Reference
}

// ApiManagementId returns a reference to field api_management_id of azurerm_api_management_gateway_host_name_configuration.
func (amghnc apiManagementGatewayHostNameConfigurationAttributes) ApiManagementId() terra.StringValue {
	return terra.ReferenceAsString(amghnc.ref.Append("api_management_id"))
}

// CertificateId returns a reference to field certificate_id of azurerm_api_management_gateway_host_name_configuration.
func (amghnc apiManagementGatewayHostNameConfigurationAttributes) CertificateId() terra.StringValue {
	return terra.ReferenceAsString(amghnc.ref.Append("certificate_id"))
}

// GatewayName returns a reference to field gateway_name of azurerm_api_management_gateway_host_name_configuration.
func (amghnc apiManagementGatewayHostNameConfigurationAttributes) GatewayName() terra.StringValue {
	return terra.ReferenceAsString(amghnc.ref.Append("gateway_name"))
}

// HostName returns a reference to field host_name of azurerm_api_management_gateway_host_name_configuration.
func (amghnc apiManagementGatewayHostNameConfigurationAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(amghnc.ref.Append("host_name"))
}

// Http2Enabled returns a reference to field http2_enabled of azurerm_api_management_gateway_host_name_configuration.
func (amghnc apiManagementGatewayHostNameConfigurationAttributes) Http2Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(amghnc.ref.Append("http2_enabled"))
}

// Id returns a reference to field id of azurerm_api_management_gateway_host_name_configuration.
func (amghnc apiManagementGatewayHostNameConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amghnc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_api_management_gateway_host_name_configuration.
func (amghnc apiManagementGatewayHostNameConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amghnc.ref.Append("name"))
}

// RequestClientCertificateEnabled returns a reference to field request_client_certificate_enabled of azurerm_api_management_gateway_host_name_configuration.
func (amghnc apiManagementGatewayHostNameConfigurationAttributes) RequestClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(amghnc.ref.Append("request_client_certificate_enabled"))
}

// Tls10Enabled returns a reference to field tls10_enabled of azurerm_api_management_gateway_host_name_configuration.
func (amghnc apiManagementGatewayHostNameConfigurationAttributes) Tls10Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(amghnc.ref.Append("tls10_enabled"))
}

// Tls11Enabled returns a reference to field tls11_enabled of azurerm_api_management_gateway_host_name_configuration.
func (amghnc apiManagementGatewayHostNameConfigurationAttributes) Tls11Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(amghnc.ref.Append("tls11_enabled"))
}

func (amghnc apiManagementGatewayHostNameConfigurationAttributes) Timeouts() apimanagementgatewayhostnameconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementgatewayhostnameconfiguration.TimeoutsAttributes](amghnc.ref.Append("timeouts"))
}

type apiManagementGatewayHostNameConfigurationState struct {
	ApiManagementId                 string                                                   `json:"api_management_id"`
	CertificateId                   string                                                   `json:"certificate_id"`
	GatewayName                     string                                                   `json:"gateway_name"`
	HostName                        string                                                   `json:"host_name"`
	Http2Enabled                    bool                                                     `json:"http2_enabled"`
	Id                              string                                                   `json:"id"`
	Name                            string                                                   `json:"name"`
	RequestClientCertificateEnabled bool                                                     `json:"request_client_certificate_enabled"`
	Tls10Enabled                    bool                                                     `json:"tls10_enabled"`
	Tls11Enabled                    bool                                                     `json:"tls11_enabled"`
	Timeouts                        *apimanagementgatewayhostnameconfiguration.TimeoutsState `json:"timeouts"`
}
