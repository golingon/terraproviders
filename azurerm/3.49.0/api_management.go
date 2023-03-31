// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagement "github.com/golingon/terraproviders/azurerm/3.49.0/apimanagement"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewApiManagement(name string, args ApiManagementArgs) *ApiManagement {
	return &ApiManagement{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagement)(nil)

type ApiManagement struct {
	Name  string
	Args  ApiManagementArgs
	state *apiManagementState
}

func (am *ApiManagement) Type() string {
	return "azurerm_api_management"
}

func (am *ApiManagement) LocalName() string {
	return am.Name
}

func (am *ApiManagement) Configuration() interface{} {
	return am.Args
}

func (am *ApiManagement) Attributes() apiManagementAttributes {
	return apiManagementAttributes{ref: terra.ReferenceResource(am)}
}

func (am *ApiManagement) ImportState(av io.Reader) error {
	am.state = &apiManagementState{}
	if err := json.NewDecoder(av).Decode(am.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", am.Type(), am.LocalName(), err)
	}
	return nil
}

func (am *ApiManagement) State() (*apiManagementState, bool) {
	return am.state, am.state != nil
}

func (am *ApiManagement) StateMust() *apiManagementState {
	if am.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", am.Type(), am.LocalName()))
	}
	return am.state
}

func (am *ApiManagement) DependOn() terra.Reference {
	return terra.ReferenceResource(am)
}

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
	// DependsOn contains resources that ApiManagement depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type apiManagementAttributes struct {
	ref terra.Reference
}

func (am apiManagementAttributes) ClientCertificateEnabled() terra.BoolValue {
	return terra.ReferenceBool(am.ref.Append("client_certificate_enabled"))
}

func (am apiManagementAttributes) DeveloperPortalUrl() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("developer_portal_url"))
}

func (am apiManagementAttributes) GatewayDisabled() terra.BoolValue {
	return terra.ReferenceBool(am.ref.Append("gateway_disabled"))
}

func (am apiManagementAttributes) GatewayRegionalUrl() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("gateway_regional_url"))
}

func (am apiManagementAttributes) GatewayUrl() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("gateway_url"))
}

func (am apiManagementAttributes) Id() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("id"))
}

func (am apiManagementAttributes) Location() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("location"))
}

func (am apiManagementAttributes) ManagementApiUrl() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("management_api_url"))
}

func (am apiManagementAttributes) MinApiVersion() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("min_api_version"))
}

func (am apiManagementAttributes) Name() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("name"))
}

func (am apiManagementAttributes) NotificationSenderEmail() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("notification_sender_email"))
}

func (am apiManagementAttributes) PortalUrl() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("portal_url"))
}

func (am apiManagementAttributes) PrivateIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](am.ref.Append("private_ip_addresses"))
}

func (am apiManagementAttributes) PublicIpAddressId() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("public_ip_address_id"))
}

func (am apiManagementAttributes) PublicIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](am.ref.Append("public_ip_addresses"))
}

func (am apiManagementAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceBool(am.ref.Append("public_network_access_enabled"))
}

func (am apiManagementAttributes) PublisherEmail() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("publisher_email"))
}

func (am apiManagementAttributes) PublisherName() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("publisher_name"))
}

func (am apiManagementAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("resource_group_name"))
}

func (am apiManagementAttributes) ScmUrl() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("scm_url"))
}

func (am apiManagementAttributes) SkuName() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("sku_name"))
}

func (am apiManagementAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](am.ref.Append("tags"))
}

func (am apiManagementAttributes) VirtualNetworkType() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("virtual_network_type"))
}

func (am apiManagementAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](am.ref.Append("zones"))
}

func (am apiManagementAttributes) Policy() terra.ListValue[apimanagement.PolicyAttributes] {
	return terra.ReferenceList[apimanagement.PolicyAttributes](am.ref.Append("policy"))
}

func (am apiManagementAttributes) AdditionalLocation() terra.ListValue[apimanagement.AdditionalLocationAttributes] {
	return terra.ReferenceList[apimanagement.AdditionalLocationAttributes](am.ref.Append("additional_location"))
}

func (am apiManagementAttributes) Certificate() terra.ListValue[apimanagement.CertificateAttributes] {
	return terra.ReferenceList[apimanagement.CertificateAttributes](am.ref.Append("certificate"))
}

func (am apiManagementAttributes) Delegation() terra.ListValue[apimanagement.DelegationAttributes] {
	return terra.ReferenceList[apimanagement.DelegationAttributes](am.ref.Append("delegation"))
}

func (am apiManagementAttributes) HostnameConfiguration() terra.ListValue[apimanagement.HostnameConfigurationAttributes] {
	return terra.ReferenceList[apimanagement.HostnameConfigurationAttributes](am.ref.Append("hostname_configuration"))
}

func (am apiManagementAttributes) Identity() terra.ListValue[apimanagement.IdentityAttributes] {
	return terra.ReferenceList[apimanagement.IdentityAttributes](am.ref.Append("identity"))
}

func (am apiManagementAttributes) Protocols() terra.ListValue[apimanagement.ProtocolsAttributes] {
	return terra.ReferenceList[apimanagement.ProtocolsAttributes](am.ref.Append("protocols"))
}

func (am apiManagementAttributes) Security() terra.ListValue[apimanagement.SecurityAttributes] {
	return terra.ReferenceList[apimanagement.SecurityAttributes](am.ref.Append("security"))
}

func (am apiManagementAttributes) SignIn() terra.ListValue[apimanagement.SignInAttributes] {
	return terra.ReferenceList[apimanagement.SignInAttributes](am.ref.Append("sign_in"))
}

func (am apiManagementAttributes) SignUp() terra.ListValue[apimanagement.SignUpAttributes] {
	return terra.ReferenceList[apimanagement.SignUpAttributes](am.ref.Append("sign_up"))
}

func (am apiManagementAttributes) TenantAccess() terra.ListValue[apimanagement.TenantAccessAttributes] {
	return terra.ReferenceList[apimanagement.TenantAccessAttributes](am.ref.Append("tenant_access"))
}

func (am apiManagementAttributes) Timeouts() apimanagement.TimeoutsAttributes {
	return terra.ReferenceSingle[apimanagement.TimeoutsAttributes](am.ref.Append("timeouts"))
}

func (am apiManagementAttributes) VirtualNetworkConfiguration() terra.ListValue[apimanagement.VirtualNetworkConfigurationAttributes] {
	return terra.ReferenceList[apimanagement.VirtualNetworkConfigurationAttributes](am.ref.Append("virtual_network_configuration"))
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
