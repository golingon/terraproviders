// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagement "github.com/golingon/terraproviders/azurerm/3.67.0/apimanagement"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagement creates a new instance of [ApiManagement].
func NewApiManagement(name string, args ApiManagementArgs) *ApiManagement {
	return &ApiManagement{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagement)(nil)

// ApiManagement represents the Terraform resource azurerm_api_management.
type ApiManagement struct {
	Name      string
	Args      ApiManagementArgs
	state     *apiManagementState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagement].
func (am *ApiManagement) Type() string {
	return "azurerm_api_management"
}

// LocalName returns the local name for [ApiManagement].
func (am *ApiManagement) LocalName() string {
	return am.Name
}

// Configuration returns the configuration (args) for [ApiManagement].
func (am *ApiManagement) Configuration() interface{} {
	return am.Args
}

// DependOn is used for other resources to depend on [ApiManagement].
func (am *ApiManagement) DependOn() terra.Reference {
	return terra.ReferenceResource(am)
}

// Dependencies returns the list of resources [ApiManagement] depends_on.
func (am *ApiManagement) Dependencies() terra.Dependencies {
	return am.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagement].
func (am *ApiManagement) LifecycleManagement() *terra.Lifecycle {
	return am.Lifecycle
}

// Attributes returns the attributes for [ApiManagement].
func (am *ApiManagement) Attributes() apiManagementAttributes {
	return apiManagementAttributes{ref: terra.ReferenceResource(am)}
}

// ImportState imports the given attribute values into [ApiManagement]'s state.
func (am *ApiManagement) ImportState(av io.Reader) error {
	am.state = &apiManagementState{}
	if err := json.NewDecoder(av).Decode(am.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", am.Type(), am.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagement] has state.
func (am *ApiManagement) State() (*apiManagementState, bool) {
	return am.state, am.state != nil
}

// StateMust returns the state for [ApiManagement]. Panics if the state is nil.
func (am *ApiManagement) StateMust() *apiManagementState {
	if am.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", am.Type(), am.LocalName()))
	}
	return am.state
}

// ApiManagementArgs contains the configurations for azurerm_api_management.
type ApiManagementArgs struct {
	// ClientCertificateEnabled: bool, optional
	ClientCertificateEnabled terra.BoolValue `hcl:"client_certificate_enabled,attr"`
	// GatewayDisabled: bool, optional
	GatewayDisabled terra.BoolValue `hcl:"gateway_disabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MinApiVersion: string, optional
	MinApiVersion terra.StringValue `hcl:"min_api_version,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NotificationSenderEmail: string, optional
	NotificationSenderEmail terra.StringValue `hcl:"notification_sender_email,attr"`
	// PublicIpAddressId: string, optional
	PublicIpAddressId terra.StringValue `hcl:"public_ip_address_id,attr"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// PublisherEmail: string, required
	PublisherEmail terra.StringValue `hcl:"publisher_email,attr" validate:"required"`
	// PublisherName: string, required
	PublisherName terra.StringValue `hcl:"publisher_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VirtualNetworkType: string, optional
	VirtualNetworkType terra.StringValue `hcl:"virtual_network_type,attr"`
	// Zones: set of string, optional
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr"`
	// Policy: min=0
	Policy []apimanagement.Policy `hcl:"policy,block" validate:"min=0"`
	// AdditionalLocation: min=0
	AdditionalLocation []apimanagement.AdditionalLocation `hcl:"additional_location,block" validate:"min=0"`
	// Certificate: min=0,max=10
	Certificate []apimanagement.Certificate `hcl:"certificate,block" validate:"min=0,max=10"`
	// Delegation: optional
	Delegation *apimanagement.Delegation `hcl:"delegation,block"`
	// HostnameConfiguration: optional
	HostnameConfiguration *apimanagement.HostnameConfiguration `hcl:"hostname_configuration,block"`
	// Identity: optional
	Identity *apimanagement.Identity `hcl:"identity,block"`
	// Protocols: optional
	Protocols *apimanagement.Protocols `hcl:"protocols,block"`
	// Security: optional
	Security *apimanagement.Security `hcl:"security,block"`
	// SignIn: optional
	SignIn *apimanagement.SignIn `hcl:"sign_in,block"`
	// SignUp: optional
	SignUp *apimanagement.SignUp `hcl:"sign_up,block"`
	// TenantAccess: optional
	TenantAccess *apimanagement.TenantAccess `hcl:"tenant_access,block"`
	// Timeouts: optional
	Timeouts *apimanagement.Timeouts `hcl:"timeouts,block"`
	// VirtualNetworkConfiguration: optional
	VirtualNetworkConfiguration *apimanagement.VirtualNetworkConfiguration `hcl:"virtual_network_configuration,block"`
}
type apiManagementAttributes struct {
	ref terra.Reference
}

// ClientCertificateEnabled returns a reference to field client_certificate_enabled of azurerm_api_management.
func (am apiManagementAttributes) ClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(am.ref.Append("client_certificate_enabled"))
}

// DeveloperPortalUrl returns a reference to field developer_portal_url of azurerm_api_management.
func (am apiManagementAttributes) DeveloperPortalUrl() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("developer_portal_url"))
}

// GatewayDisabled returns a reference to field gateway_disabled of azurerm_api_management.
func (am apiManagementAttributes) GatewayDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(am.ref.Append("gateway_disabled"))
}

// GatewayRegionalUrl returns a reference to field gateway_regional_url of azurerm_api_management.
func (am apiManagementAttributes) GatewayRegionalUrl() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("gateway_regional_url"))
}

// GatewayUrl returns a reference to field gateway_url of azurerm_api_management.
func (am apiManagementAttributes) GatewayUrl() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("gateway_url"))
}

// Id returns a reference to field id of azurerm_api_management.
func (am apiManagementAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_api_management.
func (am apiManagementAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("location"))
}

// ManagementApiUrl returns a reference to field management_api_url of azurerm_api_management.
func (am apiManagementAttributes) ManagementApiUrl() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("management_api_url"))
}

// MinApiVersion returns a reference to field min_api_version of azurerm_api_management.
func (am apiManagementAttributes) MinApiVersion() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("min_api_version"))
}

// Name returns a reference to field name of azurerm_api_management.
func (am apiManagementAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("name"))
}

// NotificationSenderEmail returns a reference to field notification_sender_email of azurerm_api_management.
func (am apiManagementAttributes) NotificationSenderEmail() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("notification_sender_email"))
}

// PortalUrl returns a reference to field portal_url of azurerm_api_management.
func (am apiManagementAttributes) PortalUrl() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("portal_url"))
}

// PrivateIpAddresses returns a reference to field private_ip_addresses of azurerm_api_management.
func (am apiManagementAttributes) PrivateIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](am.ref.Append("private_ip_addresses"))
}

// PublicIpAddressId returns a reference to field public_ip_address_id of azurerm_api_management.
func (am apiManagementAttributes) PublicIpAddressId() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("public_ip_address_id"))
}

// PublicIpAddresses returns a reference to field public_ip_addresses of azurerm_api_management.
func (am apiManagementAttributes) PublicIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](am.ref.Append("public_ip_addresses"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_api_management.
func (am apiManagementAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(am.ref.Append("public_network_access_enabled"))
}

// PublisherEmail returns a reference to field publisher_email of azurerm_api_management.
func (am apiManagementAttributes) PublisherEmail() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("publisher_email"))
}

// PublisherName returns a reference to field publisher_name of azurerm_api_management.
func (am apiManagementAttributes) PublisherName() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("publisher_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management.
func (am apiManagementAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("resource_group_name"))
}

// ScmUrl returns a reference to field scm_url of azurerm_api_management.
func (am apiManagementAttributes) ScmUrl() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("scm_url"))
}

// SkuName returns a reference to field sku_name of azurerm_api_management.
func (am apiManagementAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_api_management.
func (am apiManagementAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](am.ref.Append("tags"))
}

// VirtualNetworkType returns a reference to field virtual_network_type of azurerm_api_management.
func (am apiManagementAttributes) VirtualNetworkType() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("virtual_network_type"))
}

// Zones returns a reference to field zones of azurerm_api_management.
func (am apiManagementAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](am.ref.Append("zones"))
}

func (am apiManagementAttributes) Policy() terra.ListValue[apimanagement.PolicyAttributes] {
	return terra.ReferenceAsList[apimanagement.PolicyAttributes](am.ref.Append("policy"))
}

func (am apiManagementAttributes) AdditionalLocation() terra.ListValue[apimanagement.AdditionalLocationAttributes] {
	return terra.ReferenceAsList[apimanagement.AdditionalLocationAttributes](am.ref.Append("additional_location"))
}

func (am apiManagementAttributes) Certificate() terra.ListValue[apimanagement.CertificateAttributes] {
	return terra.ReferenceAsList[apimanagement.CertificateAttributes](am.ref.Append("certificate"))
}

func (am apiManagementAttributes) Delegation() terra.ListValue[apimanagement.DelegationAttributes] {
	return terra.ReferenceAsList[apimanagement.DelegationAttributes](am.ref.Append("delegation"))
}

func (am apiManagementAttributes) HostnameConfiguration() terra.ListValue[apimanagement.HostnameConfigurationAttributes] {
	return terra.ReferenceAsList[apimanagement.HostnameConfigurationAttributes](am.ref.Append("hostname_configuration"))
}

func (am apiManagementAttributes) Identity() terra.ListValue[apimanagement.IdentityAttributes] {
	return terra.ReferenceAsList[apimanagement.IdentityAttributes](am.ref.Append("identity"))
}

func (am apiManagementAttributes) Protocols() terra.ListValue[apimanagement.ProtocolsAttributes] {
	return terra.ReferenceAsList[apimanagement.ProtocolsAttributes](am.ref.Append("protocols"))
}

func (am apiManagementAttributes) Security() terra.ListValue[apimanagement.SecurityAttributes] {
	return terra.ReferenceAsList[apimanagement.SecurityAttributes](am.ref.Append("security"))
}

func (am apiManagementAttributes) SignIn() terra.ListValue[apimanagement.SignInAttributes] {
	return terra.ReferenceAsList[apimanagement.SignInAttributes](am.ref.Append("sign_in"))
}

func (am apiManagementAttributes) SignUp() terra.ListValue[apimanagement.SignUpAttributes] {
	return terra.ReferenceAsList[apimanagement.SignUpAttributes](am.ref.Append("sign_up"))
}

func (am apiManagementAttributes) TenantAccess() terra.ListValue[apimanagement.TenantAccessAttributes] {
	return terra.ReferenceAsList[apimanagement.TenantAccessAttributes](am.ref.Append("tenant_access"))
}

func (am apiManagementAttributes) Timeouts() apimanagement.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagement.TimeoutsAttributes](am.ref.Append("timeouts"))
}

func (am apiManagementAttributes) VirtualNetworkConfiguration() terra.ListValue[apimanagement.VirtualNetworkConfigurationAttributes] {
	return terra.ReferenceAsList[apimanagement.VirtualNetworkConfigurationAttributes](am.ref.Append("virtual_network_configuration"))
}

type apiManagementState struct {
	ClientCertificateEnabled    bool                                             `json:"client_certificate_enabled"`
	DeveloperPortalUrl          string                                           `json:"developer_portal_url"`
	GatewayDisabled             bool                                             `json:"gateway_disabled"`
	GatewayRegionalUrl          string                                           `json:"gateway_regional_url"`
	GatewayUrl                  string                                           `json:"gateway_url"`
	Id                          string                                           `json:"id"`
	Location                    string                                           `json:"location"`
	ManagementApiUrl            string                                           `json:"management_api_url"`
	MinApiVersion               string                                           `json:"min_api_version"`
	Name                        string                                           `json:"name"`
	NotificationSenderEmail     string                                           `json:"notification_sender_email"`
	PortalUrl                   string                                           `json:"portal_url"`
	PrivateIpAddresses          []string                                         `json:"private_ip_addresses"`
	PublicIpAddressId           string                                           `json:"public_ip_address_id"`
	PublicIpAddresses           []string                                         `json:"public_ip_addresses"`
	PublicNetworkAccessEnabled  bool                                             `json:"public_network_access_enabled"`
	PublisherEmail              string                                           `json:"publisher_email"`
	PublisherName               string                                           `json:"publisher_name"`
	ResourceGroupName           string                                           `json:"resource_group_name"`
	ScmUrl                      string                                           `json:"scm_url"`
	SkuName                     string                                           `json:"sku_name"`
	Tags                        map[string]string                                `json:"tags"`
	VirtualNetworkType          string                                           `json:"virtual_network_type"`
	Zones                       []string                                         `json:"zones"`
	Policy                      []apimanagement.PolicyState                      `json:"policy"`
	AdditionalLocation          []apimanagement.AdditionalLocationState          `json:"additional_location"`
	Certificate                 []apimanagement.CertificateState                 `json:"certificate"`
	Delegation                  []apimanagement.DelegationState                  `json:"delegation"`
	HostnameConfiguration       []apimanagement.HostnameConfigurationState       `json:"hostname_configuration"`
	Identity                    []apimanagement.IdentityState                    `json:"identity"`
	Protocols                   []apimanagement.ProtocolsState                   `json:"protocols"`
	Security                    []apimanagement.SecurityState                    `json:"security"`
	SignIn                      []apimanagement.SignInState                      `json:"sign_in"`
	SignUp                      []apimanagement.SignUpState                      `json:"sign_up"`
	TenantAccess                []apimanagement.TenantAccessState                `json:"tenant_access"`
	Timeouts                    *apimanagement.TimeoutsState                     `json:"timeouts"`
	VirtualNetworkConfiguration []apimanagement.VirtualNetworkConfigurationState `json:"virtual_network_configuration"`
}
