// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	vpnserverconfiguration "github.com/golingon/terraproviders/azurerm/3.67.0/vpnserverconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpnServerConfiguration creates a new instance of [VpnServerConfiguration].
func NewVpnServerConfiguration(name string, args VpnServerConfigurationArgs) *VpnServerConfiguration {
	return &VpnServerConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpnServerConfiguration)(nil)

// VpnServerConfiguration represents the Terraform resource azurerm_vpn_server_configuration.
type VpnServerConfiguration struct {
	Name      string
	Args      VpnServerConfigurationArgs
	state     *vpnServerConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpnServerConfiguration].
func (vsc *VpnServerConfiguration) Type() string {
	return "azurerm_vpn_server_configuration"
}

// LocalName returns the local name for [VpnServerConfiguration].
func (vsc *VpnServerConfiguration) LocalName() string {
	return vsc.Name
}

// Configuration returns the configuration (args) for [VpnServerConfiguration].
func (vsc *VpnServerConfiguration) Configuration() interface{} {
	return vsc.Args
}

// DependOn is used for other resources to depend on [VpnServerConfiguration].
func (vsc *VpnServerConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(vsc)
}

// Dependencies returns the list of resources [VpnServerConfiguration] depends_on.
func (vsc *VpnServerConfiguration) Dependencies() terra.Dependencies {
	return vsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpnServerConfiguration].
func (vsc *VpnServerConfiguration) LifecycleManagement() *terra.Lifecycle {
	return vsc.Lifecycle
}

// Attributes returns the attributes for [VpnServerConfiguration].
func (vsc *VpnServerConfiguration) Attributes() vpnServerConfigurationAttributes {
	return vpnServerConfigurationAttributes{ref: terra.ReferenceResource(vsc)}
}

// ImportState imports the given attribute values into [VpnServerConfiguration]'s state.
func (vsc *VpnServerConfiguration) ImportState(av io.Reader) error {
	vsc.state = &vpnServerConfigurationState{}
	if err := json.NewDecoder(av).Decode(vsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vsc.Type(), vsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpnServerConfiguration] has state.
func (vsc *VpnServerConfiguration) State() (*vpnServerConfigurationState, bool) {
	return vsc.state, vsc.state != nil
}

// StateMust returns the state for [VpnServerConfiguration]. Panics if the state is nil.
func (vsc *VpnServerConfiguration) StateMust() *vpnServerConfigurationState {
	if vsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vsc.Type(), vsc.LocalName()))
	}
	return vsc.state
}

// VpnServerConfigurationArgs contains the configurations for azurerm_vpn_server_configuration.
type VpnServerConfigurationArgs struct {
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
	// VpnAuthenticationTypes: list of string, required
	VpnAuthenticationTypes terra.ListValue[terra.StringValue] `hcl:"vpn_authentication_types,attr" validate:"required"`
	// VpnProtocols: set of string, optional
	VpnProtocols terra.SetValue[terra.StringValue] `hcl:"vpn_protocols,attr"`
	// AzureActiveDirectoryAuthentication: min=0
	AzureActiveDirectoryAuthentication []vpnserverconfiguration.AzureActiveDirectoryAuthentication `hcl:"azure_active_directory_authentication,block" validate:"min=0"`
	// ClientRevokedCertificate: min=0
	ClientRevokedCertificate []vpnserverconfiguration.ClientRevokedCertificate `hcl:"client_revoked_certificate,block" validate:"min=0"`
	// ClientRootCertificate: min=0
	ClientRootCertificate []vpnserverconfiguration.ClientRootCertificate `hcl:"client_root_certificate,block" validate:"min=0"`
	// IpsecPolicy: optional
	IpsecPolicy *vpnserverconfiguration.IpsecPolicy `hcl:"ipsec_policy,block"`
	// Radius: optional
	Radius *vpnserverconfiguration.Radius `hcl:"radius,block"`
	// Timeouts: optional
	Timeouts *vpnserverconfiguration.Timeouts `hcl:"timeouts,block"`
}
type vpnServerConfigurationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_vpn_server_configuration.
func (vsc vpnServerConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vsc.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_vpn_server_configuration.
func (vsc vpnServerConfigurationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vsc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_vpn_server_configuration.
func (vsc vpnServerConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vsc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_vpn_server_configuration.
func (vsc vpnServerConfigurationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vsc.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_vpn_server_configuration.
func (vsc vpnServerConfigurationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vsc.ref.Append("tags"))
}

// VpnAuthenticationTypes returns a reference to field vpn_authentication_types of azurerm_vpn_server_configuration.
func (vsc vpnServerConfigurationAttributes) VpnAuthenticationTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vsc.ref.Append("vpn_authentication_types"))
}

// VpnProtocols returns a reference to field vpn_protocols of azurerm_vpn_server_configuration.
func (vsc vpnServerConfigurationAttributes) VpnProtocols() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vsc.ref.Append("vpn_protocols"))
}

func (vsc vpnServerConfigurationAttributes) AzureActiveDirectoryAuthentication() terra.ListValue[vpnserverconfiguration.AzureActiveDirectoryAuthenticationAttributes] {
	return terra.ReferenceAsList[vpnserverconfiguration.AzureActiveDirectoryAuthenticationAttributes](vsc.ref.Append("azure_active_directory_authentication"))
}

func (vsc vpnServerConfigurationAttributes) ClientRevokedCertificate() terra.SetValue[vpnserverconfiguration.ClientRevokedCertificateAttributes] {
	return terra.ReferenceAsSet[vpnserverconfiguration.ClientRevokedCertificateAttributes](vsc.ref.Append("client_revoked_certificate"))
}

func (vsc vpnServerConfigurationAttributes) ClientRootCertificate() terra.SetValue[vpnserverconfiguration.ClientRootCertificateAttributes] {
	return terra.ReferenceAsSet[vpnserverconfiguration.ClientRootCertificateAttributes](vsc.ref.Append("client_root_certificate"))
}

func (vsc vpnServerConfigurationAttributes) IpsecPolicy() terra.ListValue[vpnserverconfiguration.IpsecPolicyAttributes] {
	return terra.ReferenceAsList[vpnserverconfiguration.IpsecPolicyAttributes](vsc.ref.Append("ipsec_policy"))
}

func (vsc vpnServerConfigurationAttributes) Radius() terra.ListValue[vpnserverconfiguration.RadiusAttributes] {
	return terra.ReferenceAsList[vpnserverconfiguration.RadiusAttributes](vsc.ref.Append("radius"))
}

func (vsc vpnServerConfigurationAttributes) Timeouts() vpnserverconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpnserverconfiguration.TimeoutsAttributes](vsc.ref.Append("timeouts"))
}

type vpnServerConfigurationState struct {
	Id                                 string                                                           `json:"id"`
	Location                           string                                                           `json:"location"`
	Name                               string                                                           `json:"name"`
	ResourceGroupName                  string                                                           `json:"resource_group_name"`
	Tags                               map[string]string                                                `json:"tags"`
	VpnAuthenticationTypes             []string                                                         `json:"vpn_authentication_types"`
	VpnProtocols                       []string                                                         `json:"vpn_protocols"`
	AzureActiveDirectoryAuthentication []vpnserverconfiguration.AzureActiveDirectoryAuthenticationState `json:"azure_active_directory_authentication"`
	ClientRevokedCertificate           []vpnserverconfiguration.ClientRevokedCertificateState           `json:"client_revoked_certificate"`
	ClientRootCertificate              []vpnserverconfiguration.ClientRootCertificateState              `json:"client_root_certificate"`
	IpsecPolicy                        []vpnserverconfiguration.IpsecPolicyState                        `json:"ipsec_policy"`
	Radius                             []vpnserverconfiguration.RadiusState                             `json:"radius"`
	Timeouts                           *vpnserverconfiguration.TimeoutsState                            `json:"timeouts"`
}
