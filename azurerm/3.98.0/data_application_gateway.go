// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	dataapplicationgateway "github.com/golingon/terraproviders/azurerm/3.98.0/dataapplicationgateway"
)

// NewDataApplicationGateway creates a new instance of [DataApplicationGateway].
func NewDataApplicationGateway(name string, args DataApplicationGatewayArgs) *DataApplicationGateway {
	return &DataApplicationGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApplicationGateway)(nil)

// DataApplicationGateway represents the Terraform data resource azurerm_application_gateway.
type DataApplicationGateway struct {
	Name string
	Args DataApplicationGatewayArgs
}

// DataSource returns the Terraform object type for [DataApplicationGateway].
func (ag *DataApplicationGateway) DataSource() string {
	return "azurerm_application_gateway"
}

// LocalName returns the local name for [DataApplicationGateway].
func (ag *DataApplicationGateway) LocalName() string {
	return ag.Name
}

// Configuration returns the configuration (args) for [DataApplicationGateway].
func (ag *DataApplicationGateway) Configuration() interface{} {
	return ag.Args
}

// Attributes returns the attributes for [DataApplicationGateway].
func (ag *DataApplicationGateway) Attributes() dataApplicationGatewayAttributes {
	return dataApplicationGatewayAttributes{ref: terra.ReferenceDataResource(ag)}
}

// DataApplicationGatewayArgs contains the configurations for azurerm_application_gateway.
type DataApplicationGatewayArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// AuthenticationCertificate: min=0
	AuthenticationCertificate []dataapplicationgateway.AuthenticationCertificate `hcl:"authentication_certificate,block" validate:"min=0"`
	// AutoscaleConfiguration: min=0
	AutoscaleConfiguration []dataapplicationgateway.AutoscaleConfiguration `hcl:"autoscale_configuration,block" validate:"min=0"`
	// BackendAddressPool: min=0
	BackendAddressPool []dataapplicationgateway.BackendAddressPool `hcl:"backend_address_pool,block" validate:"min=0"`
	// BackendHttpSettings: min=0
	BackendHttpSettings []dataapplicationgateway.BackendHttpSettings `hcl:"backend_http_settings,block" validate:"min=0"`
	// CustomErrorConfiguration: min=0
	CustomErrorConfiguration []dataapplicationgateway.CustomErrorConfiguration `hcl:"custom_error_configuration,block" validate:"min=0"`
	// FrontendIpConfiguration: min=0
	FrontendIpConfiguration []dataapplicationgateway.FrontendIpConfiguration `hcl:"frontend_ip_configuration,block" validate:"min=0"`
	// FrontendPort: min=0
	FrontendPort []dataapplicationgateway.FrontendPort `hcl:"frontend_port,block" validate:"min=0"`
	// GatewayIpConfiguration: min=0
	GatewayIpConfiguration []dataapplicationgateway.GatewayIpConfiguration `hcl:"gateway_ip_configuration,block" validate:"min=0"`
	// Global: min=0
	Global []dataapplicationgateway.Global `hcl:"global,block" validate:"min=0"`
	// HttpListener: min=0
	HttpListener []dataapplicationgateway.HttpListener `hcl:"http_listener,block" validate:"min=0"`
	// Identity: min=0
	Identity []dataapplicationgateway.Identity `hcl:"identity,block" validate:"min=0"`
	// PrivateEndpointConnection: min=0
	PrivateEndpointConnection []dataapplicationgateway.PrivateEndpointConnection `hcl:"private_endpoint_connection,block" validate:"min=0"`
	// PrivateLinkConfiguration: min=0
	PrivateLinkConfiguration []dataapplicationgateway.PrivateLinkConfiguration `hcl:"private_link_configuration,block" validate:"min=0"`
	// Probe: min=0
	Probe []dataapplicationgateway.Probe `hcl:"probe,block" validate:"min=0"`
	// RedirectConfiguration: min=0
	RedirectConfiguration []dataapplicationgateway.RedirectConfiguration `hcl:"redirect_configuration,block" validate:"min=0"`
	// RequestRoutingRule: min=0
	RequestRoutingRule []dataapplicationgateway.RequestRoutingRule `hcl:"request_routing_rule,block" validate:"min=0"`
	// RewriteRuleSet: min=0
	RewriteRuleSet []dataapplicationgateway.RewriteRuleSet `hcl:"rewrite_rule_set,block" validate:"min=0"`
	// Sku: min=0
	Sku []dataapplicationgateway.Sku `hcl:"sku,block" validate:"min=0"`
	// SslCertificate: min=0
	SslCertificate []dataapplicationgateway.SslCertificate `hcl:"ssl_certificate,block" validate:"min=0"`
	// SslPolicy: min=0
	SslPolicy []dataapplicationgateway.SslPolicy `hcl:"ssl_policy,block" validate:"min=0"`
	// SslProfile: min=0
	SslProfile []dataapplicationgateway.SslProfile `hcl:"ssl_profile,block" validate:"min=0"`
	// TrustedClientCertificate: min=0
	TrustedClientCertificate []dataapplicationgateway.TrustedClientCertificate `hcl:"trusted_client_certificate,block" validate:"min=0"`
	// TrustedRootCertificate: min=0
	TrustedRootCertificate []dataapplicationgateway.TrustedRootCertificate `hcl:"trusted_root_certificate,block" validate:"min=0"`
	// UrlPathMap: min=0
	UrlPathMap []dataapplicationgateway.UrlPathMap `hcl:"url_path_map,block" validate:"min=0"`
	// WafConfiguration: min=0
	WafConfiguration []dataapplicationgateway.WafConfiguration `hcl:"waf_configuration,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataapplicationgateway.Timeouts `hcl:"timeouts,block"`
}
type dataApplicationGatewayAttributes struct {
	ref terra.Reference
}

// FipsEnabled returns a reference to field fips_enabled of azurerm_application_gateway.
func (ag dataApplicationGatewayAttributes) FipsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ag.ref.Append("fips_enabled"))
}

// FirewallPolicyId returns a reference to field firewall_policy_id of azurerm_application_gateway.
func (ag dataApplicationGatewayAttributes) FirewallPolicyId() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("firewall_policy_id"))
}

// ForceFirewallPolicyAssociation returns a reference to field force_firewall_policy_association of azurerm_application_gateway.
func (ag dataApplicationGatewayAttributes) ForceFirewallPolicyAssociation() terra.BoolValue {
	return terra.ReferenceAsBool(ag.ref.Append("force_firewall_policy_association"))
}

// Http2Enabled returns a reference to field http2_enabled of azurerm_application_gateway.
func (ag dataApplicationGatewayAttributes) Http2Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ag.ref.Append("http2_enabled"))
}

// Id returns a reference to field id of azurerm_application_gateway.
func (ag dataApplicationGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_application_gateway.
func (ag dataApplicationGatewayAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_application_gateway.
func (ag dataApplicationGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_application_gateway.
func (ag dataApplicationGatewayAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_application_gateway.
func (ag dataApplicationGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ag.ref.Append("tags"))
}

// Zones returns a reference to field zones of azurerm_application_gateway.
func (ag dataApplicationGatewayAttributes) Zones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ag.ref.Append("zones"))
}

func (ag dataApplicationGatewayAttributes) AuthenticationCertificate() terra.ListValue[dataapplicationgateway.AuthenticationCertificateAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.AuthenticationCertificateAttributes](ag.ref.Append("authentication_certificate"))
}

func (ag dataApplicationGatewayAttributes) AutoscaleConfiguration() terra.ListValue[dataapplicationgateway.AutoscaleConfigurationAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.AutoscaleConfigurationAttributes](ag.ref.Append("autoscale_configuration"))
}

func (ag dataApplicationGatewayAttributes) BackendAddressPool() terra.ListValue[dataapplicationgateway.BackendAddressPoolAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.BackendAddressPoolAttributes](ag.ref.Append("backend_address_pool"))
}

func (ag dataApplicationGatewayAttributes) BackendHttpSettings() terra.ListValue[dataapplicationgateway.BackendHttpSettingsAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.BackendHttpSettingsAttributes](ag.ref.Append("backend_http_settings"))
}

func (ag dataApplicationGatewayAttributes) CustomErrorConfiguration() terra.ListValue[dataapplicationgateway.CustomErrorConfigurationAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.CustomErrorConfigurationAttributes](ag.ref.Append("custom_error_configuration"))
}

func (ag dataApplicationGatewayAttributes) FrontendIpConfiguration() terra.ListValue[dataapplicationgateway.FrontendIpConfigurationAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.FrontendIpConfigurationAttributes](ag.ref.Append("frontend_ip_configuration"))
}

func (ag dataApplicationGatewayAttributes) FrontendPort() terra.ListValue[dataapplicationgateway.FrontendPortAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.FrontendPortAttributes](ag.ref.Append("frontend_port"))
}

func (ag dataApplicationGatewayAttributes) GatewayIpConfiguration() terra.ListValue[dataapplicationgateway.GatewayIpConfigurationAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.GatewayIpConfigurationAttributes](ag.ref.Append("gateway_ip_configuration"))
}

func (ag dataApplicationGatewayAttributes) Global() terra.ListValue[dataapplicationgateway.GlobalAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.GlobalAttributes](ag.ref.Append("global"))
}

func (ag dataApplicationGatewayAttributes) HttpListener() terra.ListValue[dataapplicationgateway.HttpListenerAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.HttpListenerAttributes](ag.ref.Append("http_listener"))
}

func (ag dataApplicationGatewayAttributes) Identity() terra.ListValue[dataapplicationgateway.IdentityAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.IdentityAttributes](ag.ref.Append("identity"))
}

func (ag dataApplicationGatewayAttributes) PrivateEndpointConnection() terra.ListValue[dataapplicationgateway.PrivateEndpointConnectionAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.PrivateEndpointConnectionAttributes](ag.ref.Append("private_endpoint_connection"))
}

func (ag dataApplicationGatewayAttributes) PrivateLinkConfiguration() terra.ListValue[dataapplicationgateway.PrivateLinkConfigurationAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.PrivateLinkConfigurationAttributes](ag.ref.Append("private_link_configuration"))
}

func (ag dataApplicationGatewayAttributes) Probe() terra.ListValue[dataapplicationgateway.ProbeAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.ProbeAttributes](ag.ref.Append("probe"))
}

func (ag dataApplicationGatewayAttributes) RedirectConfiguration() terra.ListValue[dataapplicationgateway.RedirectConfigurationAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.RedirectConfigurationAttributes](ag.ref.Append("redirect_configuration"))
}

func (ag dataApplicationGatewayAttributes) RequestRoutingRule() terra.ListValue[dataapplicationgateway.RequestRoutingRuleAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.RequestRoutingRuleAttributes](ag.ref.Append("request_routing_rule"))
}

func (ag dataApplicationGatewayAttributes) RewriteRuleSet() terra.ListValue[dataapplicationgateway.RewriteRuleSetAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.RewriteRuleSetAttributes](ag.ref.Append("rewrite_rule_set"))
}

func (ag dataApplicationGatewayAttributes) Sku() terra.ListValue[dataapplicationgateway.SkuAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.SkuAttributes](ag.ref.Append("sku"))
}

func (ag dataApplicationGatewayAttributes) SslCertificate() terra.ListValue[dataapplicationgateway.SslCertificateAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.SslCertificateAttributes](ag.ref.Append("ssl_certificate"))
}

func (ag dataApplicationGatewayAttributes) SslPolicy() terra.ListValue[dataapplicationgateway.SslPolicyAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.SslPolicyAttributes](ag.ref.Append("ssl_policy"))
}

func (ag dataApplicationGatewayAttributes) SslProfile() terra.ListValue[dataapplicationgateway.SslProfileAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.SslProfileAttributes](ag.ref.Append("ssl_profile"))
}

func (ag dataApplicationGatewayAttributes) TrustedClientCertificate() terra.ListValue[dataapplicationgateway.TrustedClientCertificateAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.TrustedClientCertificateAttributes](ag.ref.Append("trusted_client_certificate"))
}

func (ag dataApplicationGatewayAttributes) TrustedRootCertificate() terra.ListValue[dataapplicationgateway.TrustedRootCertificateAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.TrustedRootCertificateAttributes](ag.ref.Append("trusted_root_certificate"))
}

func (ag dataApplicationGatewayAttributes) UrlPathMap() terra.ListValue[dataapplicationgateway.UrlPathMapAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.UrlPathMapAttributes](ag.ref.Append("url_path_map"))
}

func (ag dataApplicationGatewayAttributes) WafConfiguration() terra.ListValue[dataapplicationgateway.WafConfigurationAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.WafConfigurationAttributes](ag.ref.Append("waf_configuration"))
}

func (ag dataApplicationGatewayAttributes) Timeouts() dataapplicationgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataapplicationgateway.TimeoutsAttributes](ag.ref.Append("timeouts"))
}
