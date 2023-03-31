// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	vpnserverconfigurationpolicygroup "github.com/golingon/terraproviders/azurerm/3.49.0/vpnserverconfigurationpolicygroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewVpnServerConfigurationPolicyGroup(name string, args VpnServerConfigurationPolicyGroupArgs) *VpnServerConfigurationPolicyGroup {
	return &VpnServerConfigurationPolicyGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpnServerConfigurationPolicyGroup)(nil)

type VpnServerConfigurationPolicyGroup struct {
	Name  string
	Args  VpnServerConfigurationPolicyGroupArgs
	state *vpnServerConfigurationPolicyGroupState
}

func (vscpg *VpnServerConfigurationPolicyGroup) Type() string {
	return "azurerm_vpn_server_configuration_policy_group"
}

func (vscpg *VpnServerConfigurationPolicyGroup) LocalName() string {
	return vscpg.Name
}

func (vscpg *VpnServerConfigurationPolicyGroup) Configuration() interface{} {
	return vscpg.Args
}

func (vscpg *VpnServerConfigurationPolicyGroup) Attributes() vpnServerConfigurationPolicyGroupAttributes {
	return vpnServerConfigurationPolicyGroupAttributes{ref: terra.ReferenceResource(vscpg)}
}

func (vscpg *VpnServerConfigurationPolicyGroup) ImportState(av io.Reader) error {
	vscpg.state = &vpnServerConfigurationPolicyGroupState{}
	if err := json.NewDecoder(av).Decode(vscpg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vscpg.Type(), vscpg.LocalName(), err)
	}
	return nil
}

func (vscpg *VpnServerConfigurationPolicyGroup) State() (*vpnServerConfigurationPolicyGroupState, bool) {
	return vscpg.state, vscpg.state != nil
}

func (vscpg *VpnServerConfigurationPolicyGroup) StateMust() *vpnServerConfigurationPolicyGroupState {
	if vscpg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vscpg.Type(), vscpg.LocalName()))
	}
	return vscpg.state
}

func (vscpg *VpnServerConfigurationPolicyGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(vscpg)
}

type VpnServerConfigurationPolicyGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsDefault: bool, optional
	IsDefault terra.BoolValue `hcl:"is_default,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// VpnServerConfigurationId: string, required
	VpnServerConfigurationId terra.StringValue `hcl:"vpn_server_configuration_id,attr" validate:"required"`
	// Policy: min=1
	Policy []vpnserverconfigurationpolicygroup.Policy `hcl:"policy,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *vpnserverconfigurationpolicygroup.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that VpnServerConfigurationPolicyGroup depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type vpnServerConfigurationPolicyGroupAttributes struct {
	ref terra.Reference
}

func (vscpg vpnServerConfigurationPolicyGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(vscpg.ref.Append("id"))
}

func (vscpg vpnServerConfigurationPolicyGroupAttributes) IsDefault() terra.BoolValue {
	return terra.ReferenceBool(vscpg.ref.Append("is_default"))
}

func (vscpg vpnServerConfigurationPolicyGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(vscpg.ref.Append("name"))
}

func (vscpg vpnServerConfigurationPolicyGroupAttributes) Priority() terra.NumberValue {
	return terra.ReferenceNumber(vscpg.ref.Append("priority"))
}

func (vscpg vpnServerConfigurationPolicyGroupAttributes) VpnServerConfigurationId() terra.StringValue {
	return terra.ReferenceString(vscpg.ref.Append("vpn_server_configuration_id"))
}

func (vscpg vpnServerConfigurationPolicyGroupAttributes) Policy() terra.SetValue[vpnserverconfigurationpolicygroup.PolicyAttributes] {
	return terra.ReferenceSet[vpnserverconfigurationpolicygroup.PolicyAttributes](vscpg.ref.Append("policy"))
}

func (vscpg vpnServerConfigurationPolicyGroupAttributes) Timeouts() vpnserverconfigurationpolicygroup.TimeoutsAttributes {
	return terra.ReferenceSingle[vpnserverconfigurationpolicygroup.TimeoutsAttributes](vscpg.ref.Append("timeouts"))
}

type vpnServerConfigurationPolicyGroupState struct {
	Id                       string                                           `json:"id"`
	IsDefault                bool                                             `json:"is_default"`
	Name                     string                                           `json:"name"`
	Priority                 float64                                          `json:"priority"`
	VpnServerConfigurationId string                                           `json:"vpn_server_configuration_id"`
	Policy                   []vpnserverconfigurationpolicygroup.PolicyState  `json:"policy"`
	Timeouts                 *vpnserverconfigurationpolicygroup.TimeoutsState `json:"timeouts"`
}
