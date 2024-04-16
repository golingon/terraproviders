// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_application_gateway

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

// Resource represents the Terraform resource azurerm_application_gateway.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermApplicationGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aag *Resource) Type() string {
	return "azurerm_application_gateway"
}

// LocalName returns the local name for [Resource].
func (aag *Resource) LocalName() string {
	return aag.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aag *Resource) Configuration() interface{} {
	return aag.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aag *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aag)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aag *Resource) Dependencies() terra.Dependencies {
	return aag.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aag *Resource) LifecycleManagement() *terra.Lifecycle {
	return aag.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aag *Resource) Attributes() azurermApplicationGatewayAttributes {
	return azurermApplicationGatewayAttributes{ref: terra.ReferenceResource(aag)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aag *Resource) ImportState(state io.Reader) error {
	aag.state = &azurermApplicationGatewayState{}
	if err := json.NewDecoder(state).Decode(aag.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aag.Type(), aag.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aag *Resource) State() (*azurermApplicationGatewayState, bool) {
	return aag.state, aag.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aag *Resource) StateMust() *azurermApplicationGatewayState {
	if aag.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aag.Type(), aag.LocalName()))
	}
	return aag.state
}

// Args contains the configurations for azurerm_application_gateway.
type Args struct {
	// EnableHttp2: bool, optional
	EnableHttp2 terra.BoolValue `hcl:"enable_http2,attr"`
	// FipsEnabled: bool, optional
	FipsEnabled terra.BoolValue `hcl:"fips_enabled,attr"`
	// FirewallPolicyId: string, optional
	FirewallPolicyId terra.StringValue `hcl:"firewall_policy_id,attr"`
	// ForceFirewallPolicyAssociation: bool, optional
	ForceFirewallPolicyAssociation terra.BoolValue `hcl:"force_firewall_policy_association,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Zones: set of string, optional
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr"`
	// AuthenticationCertificate: min=0
	AuthenticationCertificate []AuthenticationCertificate `hcl:"authentication_certificate,block" validate:"min=0"`
	// AutoscaleConfiguration: optional
	AutoscaleConfiguration *AutoscaleConfiguration `hcl:"autoscale_configuration,block"`
	// BackendAddressPool: min=1
	BackendAddressPool []BackendAddressPool `hcl:"backend_address_pool,block" validate:"min=1"`
	// BackendHttpSettings: min=1
	BackendHttpSettings []BackendHttpSettings `hcl:"backend_http_settings,block" validate:"min=1"`
	// CustomErrorConfiguration: min=0
	CustomErrorConfiguration []CustomErrorConfiguration `hcl:"custom_error_configuration,block" validate:"min=0"`
	// FrontendIpConfiguration: min=1
	FrontendIpConfiguration []FrontendIpConfiguration `hcl:"frontend_ip_configuration,block" validate:"min=1"`
	// FrontendPort: min=1
	FrontendPort []FrontendPort `hcl:"frontend_port,block" validate:"min=1"`
	// GatewayIpConfiguration: min=1,max=2
	GatewayIpConfiguration []GatewayIpConfiguration `hcl:"gateway_ip_configuration,block" validate:"min=1,max=2"`
	// Global: optional
	Global *Global `hcl:"global,block"`
	// HttpListener: min=1
	HttpListener []HttpListener `hcl:"http_listener,block" validate:"min=1"`
	// Identity: optional
	Identity *Identity `hcl:"identity,block"`
	// PrivateLinkConfiguration: min=0
	PrivateLinkConfiguration []PrivateLinkConfiguration `hcl:"private_link_configuration,block" validate:"min=0"`
	// Probe: min=0
	Probe []Probe `hcl:"probe,block" validate:"min=0"`
	// RedirectConfiguration: min=0
	RedirectConfiguration []RedirectConfiguration `hcl:"redirect_configuration,block" validate:"min=0"`
	// RequestRoutingRule: min=1
	RequestRoutingRule []RequestRoutingRule `hcl:"request_routing_rule,block" validate:"min=1"`
	// RewriteRuleSet: min=0
	RewriteRuleSet []RewriteRuleSet `hcl:"rewrite_rule_set,block" validate:"min=0"`
	// Sku: required
	Sku *Sku `hcl:"sku,block" validate:"required"`
	// SslCertificate: min=0
	SslCertificate []SslCertificate `hcl:"ssl_certificate,block" validate:"min=0"`
	// SslPolicy: optional
	SslPolicy *SslPolicy `hcl:"ssl_policy,block"`
	// SslProfile: min=0
	SslProfile []SslProfile `hcl:"ssl_profile,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// TrustedClientCertificate: min=0
	TrustedClientCertificate []TrustedClientCertificate `hcl:"trusted_client_certificate,block" validate:"min=0"`
	// TrustedRootCertificate: min=0
	TrustedRootCertificate []TrustedRootCertificate `hcl:"trusted_root_certificate,block" validate:"min=0"`
	// UrlPathMap: min=0
	UrlPathMap []UrlPathMap `hcl:"url_path_map,block" validate:"min=0"`
	// WafConfiguration: optional
	WafConfiguration *WafConfiguration `hcl:"waf_configuration,block"`
}

type azurermApplicationGatewayAttributes struct {
	ref terra.Reference
}

// EnableHttp2 returns a reference to field enable_http2 of azurerm_application_gateway.
func (aag azurermApplicationGatewayAttributes) EnableHttp2() terra.BoolValue {
	return terra.ReferenceAsBool(aag.ref.Append("enable_http2"))
}

// FipsEnabled returns a reference to field fips_enabled of azurerm_application_gateway.
func (aag azurermApplicationGatewayAttributes) FipsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aag.ref.Append("fips_enabled"))
}

// FirewallPolicyId returns a reference to field firewall_policy_id of azurerm_application_gateway.
func (aag azurermApplicationGatewayAttributes) FirewallPolicyId() terra.StringValue {
	return terra.ReferenceAsString(aag.ref.Append("firewall_policy_id"))
}

// ForceFirewallPolicyAssociation returns a reference to field force_firewall_policy_association of azurerm_application_gateway.
func (aag azurermApplicationGatewayAttributes) ForceFirewallPolicyAssociation() terra.BoolValue {
	return terra.ReferenceAsBool(aag.ref.Append("force_firewall_policy_association"))
}

// Id returns a reference to field id of azurerm_application_gateway.
func (aag azurermApplicationGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aag.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_application_gateway.
func (aag azurermApplicationGatewayAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aag.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_application_gateway.
func (aag azurermApplicationGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aag.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_application_gateway.
func (aag azurermApplicationGatewayAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aag.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_application_gateway.
func (aag azurermApplicationGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aag.ref.Append("tags"))
}

// Zones returns a reference to field zones of azurerm_application_gateway.
func (aag azurermApplicationGatewayAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aag.ref.Append("zones"))
}

func (aag azurermApplicationGatewayAttributes) PrivateEndpointConnection() terra.SetValue[PrivateEndpointConnectionAttributes] {
	return terra.ReferenceAsSet[PrivateEndpointConnectionAttributes](aag.ref.Append("private_endpoint_connection"))
}

func (aag azurermApplicationGatewayAttributes) AuthenticationCertificate() terra.ListValue[AuthenticationCertificateAttributes] {
	return terra.ReferenceAsList[AuthenticationCertificateAttributes](aag.ref.Append("authentication_certificate"))
}

func (aag azurermApplicationGatewayAttributes) AutoscaleConfiguration() terra.ListValue[AutoscaleConfigurationAttributes] {
	return terra.ReferenceAsList[AutoscaleConfigurationAttributes](aag.ref.Append("autoscale_configuration"))
}

func (aag azurermApplicationGatewayAttributes) BackendAddressPool() terra.SetValue[BackendAddressPoolAttributes] {
	return terra.ReferenceAsSet[BackendAddressPoolAttributes](aag.ref.Append("backend_address_pool"))
}

func (aag azurermApplicationGatewayAttributes) BackendHttpSettings() terra.SetValue[BackendHttpSettingsAttributes] {
	return terra.ReferenceAsSet[BackendHttpSettingsAttributes](aag.ref.Append("backend_http_settings"))
}

func (aag azurermApplicationGatewayAttributes) CustomErrorConfiguration() terra.ListValue[CustomErrorConfigurationAttributes] {
	return terra.ReferenceAsList[CustomErrorConfigurationAttributes](aag.ref.Append("custom_error_configuration"))
}

func (aag azurermApplicationGatewayAttributes) FrontendIpConfiguration() terra.ListValue[FrontendIpConfigurationAttributes] {
	return terra.ReferenceAsList[FrontendIpConfigurationAttributes](aag.ref.Append("frontend_ip_configuration"))
}

func (aag azurermApplicationGatewayAttributes) FrontendPort() terra.SetValue[FrontendPortAttributes] {
	return terra.ReferenceAsSet[FrontendPortAttributes](aag.ref.Append("frontend_port"))
}

func (aag azurermApplicationGatewayAttributes) GatewayIpConfiguration() terra.ListValue[GatewayIpConfigurationAttributes] {
	return terra.ReferenceAsList[GatewayIpConfigurationAttributes](aag.ref.Append("gateway_ip_configuration"))
}

func (aag azurermApplicationGatewayAttributes) Global() terra.ListValue[GlobalAttributes] {
	return terra.ReferenceAsList[GlobalAttributes](aag.ref.Append("global"))
}

func (aag azurermApplicationGatewayAttributes) HttpListener() terra.SetValue[HttpListenerAttributes] {
	return terra.ReferenceAsSet[HttpListenerAttributes](aag.ref.Append("http_listener"))
}

func (aag azurermApplicationGatewayAttributes) Identity() terra.ListValue[IdentityAttributes] {
	return terra.ReferenceAsList[IdentityAttributes](aag.ref.Append("identity"))
}

func (aag azurermApplicationGatewayAttributes) PrivateLinkConfiguration() terra.SetValue[PrivateLinkConfigurationAttributes] {
	return terra.ReferenceAsSet[PrivateLinkConfigurationAttributes](aag.ref.Append("private_link_configuration"))
}

func (aag azurermApplicationGatewayAttributes) Probe() terra.SetValue[ProbeAttributes] {
	return terra.ReferenceAsSet[ProbeAttributes](aag.ref.Append("probe"))
}

func (aag azurermApplicationGatewayAttributes) RedirectConfiguration() terra.SetValue[RedirectConfigurationAttributes] {
	return terra.ReferenceAsSet[RedirectConfigurationAttributes](aag.ref.Append("redirect_configuration"))
}

func (aag azurermApplicationGatewayAttributes) RequestRoutingRule() terra.SetValue[RequestRoutingRuleAttributes] {
	return terra.ReferenceAsSet[RequestRoutingRuleAttributes](aag.ref.Append("request_routing_rule"))
}

func (aag azurermApplicationGatewayAttributes) RewriteRuleSet() terra.ListValue[RewriteRuleSetAttributes] {
	return terra.ReferenceAsList[RewriteRuleSetAttributes](aag.ref.Append("rewrite_rule_set"))
}

func (aag azurermApplicationGatewayAttributes) Sku() terra.ListValue[SkuAttributes] {
	return terra.ReferenceAsList[SkuAttributes](aag.ref.Append("sku"))
}

func (aag azurermApplicationGatewayAttributes) SslCertificate() terra.SetValue[SslCertificateAttributes] {
	return terra.ReferenceAsSet[SslCertificateAttributes](aag.ref.Append("ssl_certificate"))
}

func (aag azurermApplicationGatewayAttributes) SslPolicy() terra.ListValue[SslPolicyAttributes] {
	return terra.ReferenceAsList[SslPolicyAttributes](aag.ref.Append("ssl_policy"))
}

func (aag azurermApplicationGatewayAttributes) SslProfile() terra.ListValue[SslProfileAttributes] {
	return terra.ReferenceAsList[SslProfileAttributes](aag.ref.Append("ssl_profile"))
}

func (aag azurermApplicationGatewayAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aag.ref.Append("timeouts"))
}

func (aag azurermApplicationGatewayAttributes) TrustedClientCertificate() terra.ListValue[TrustedClientCertificateAttributes] {
	return terra.ReferenceAsList[TrustedClientCertificateAttributes](aag.ref.Append("trusted_client_certificate"))
}

func (aag azurermApplicationGatewayAttributes) TrustedRootCertificate() terra.ListValue[TrustedRootCertificateAttributes] {
	return terra.ReferenceAsList[TrustedRootCertificateAttributes](aag.ref.Append("trusted_root_certificate"))
}

func (aag azurermApplicationGatewayAttributes) UrlPathMap() terra.ListValue[UrlPathMapAttributes] {
	return terra.ReferenceAsList[UrlPathMapAttributes](aag.ref.Append("url_path_map"))
}

func (aag azurermApplicationGatewayAttributes) WafConfiguration() terra.ListValue[WafConfigurationAttributes] {
	return terra.ReferenceAsList[WafConfigurationAttributes](aag.ref.Append("waf_configuration"))
}

type azurermApplicationGatewayState struct {
	EnableHttp2                    bool                             `json:"enable_http2"`
	FipsEnabled                    bool                             `json:"fips_enabled"`
	FirewallPolicyId               string                           `json:"firewall_policy_id"`
	ForceFirewallPolicyAssociation bool                             `json:"force_firewall_policy_association"`
	Id                             string                           `json:"id"`
	Location                       string                           `json:"location"`
	Name                           string                           `json:"name"`
	ResourceGroupName              string                           `json:"resource_group_name"`
	Tags                           map[string]string                `json:"tags"`
	Zones                          []string                         `json:"zones"`
	PrivateEndpointConnection      []PrivateEndpointConnectionState `json:"private_endpoint_connection"`
	AuthenticationCertificate      []AuthenticationCertificateState `json:"authentication_certificate"`
	AutoscaleConfiguration         []AutoscaleConfigurationState    `json:"autoscale_configuration"`
	BackendAddressPool             []BackendAddressPoolState        `json:"backend_address_pool"`
	BackendHttpSettings            []BackendHttpSettingsState       `json:"backend_http_settings"`
	CustomErrorConfiguration       []CustomErrorConfigurationState  `json:"custom_error_configuration"`
	FrontendIpConfiguration        []FrontendIpConfigurationState   `json:"frontend_ip_configuration"`
	FrontendPort                   []FrontendPortState              `json:"frontend_port"`
	GatewayIpConfiguration         []GatewayIpConfigurationState    `json:"gateway_ip_configuration"`
	Global                         []GlobalState                    `json:"global"`
	HttpListener                   []HttpListenerState              `json:"http_listener"`
	Identity                       []IdentityState                  `json:"identity"`
	PrivateLinkConfiguration       []PrivateLinkConfigurationState  `json:"private_link_configuration"`
	Probe                          []ProbeState                     `json:"probe"`
	RedirectConfiguration          []RedirectConfigurationState     `json:"redirect_configuration"`
	RequestRoutingRule             []RequestRoutingRuleState        `json:"request_routing_rule"`
	RewriteRuleSet                 []RewriteRuleSetState            `json:"rewrite_rule_set"`
	Sku                            []SkuState                       `json:"sku"`
	SslCertificate                 []SslCertificateState            `json:"ssl_certificate"`
	SslPolicy                      []SslPolicyState                 `json:"ssl_policy"`
	SslProfile                     []SslProfileState                `json:"ssl_profile"`
	Timeouts                       *TimeoutsState                   `json:"timeouts"`
	TrustedClientCertificate       []TrustedClientCertificateState  `json:"trusted_client_certificate"`
	TrustedRootCertificate         []TrustedRootCertificateState    `json:"trusted_root_certificate"`
	UrlPathMap                     []UrlPathMapState                `json:"url_path_map"`
	WafConfiguration               []WafConfigurationState          `json:"waf_configuration"`
}
