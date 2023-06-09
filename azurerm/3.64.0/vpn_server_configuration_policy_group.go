// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	vpnserverconfigurationpolicygroup "github.com/golingon/terraproviders/azurerm/3.64.0/vpnserverconfigurationpolicygroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpnServerConfigurationPolicyGroup creates a new instance of [VpnServerConfigurationPolicyGroup].
func NewVpnServerConfigurationPolicyGroup(name string, args VpnServerConfigurationPolicyGroupArgs) *VpnServerConfigurationPolicyGroup {
	return &VpnServerConfigurationPolicyGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpnServerConfigurationPolicyGroup)(nil)

// VpnServerConfigurationPolicyGroup represents the Terraform resource azurerm_vpn_server_configuration_policy_group.
type VpnServerConfigurationPolicyGroup struct {
	Name      string
	Args      VpnServerConfigurationPolicyGroupArgs
	state     *vpnServerConfigurationPolicyGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpnServerConfigurationPolicyGroup].
func (vscpg *VpnServerConfigurationPolicyGroup) Type() string {
	return "azurerm_vpn_server_configuration_policy_group"
}

// LocalName returns the local name for [VpnServerConfigurationPolicyGroup].
func (vscpg *VpnServerConfigurationPolicyGroup) LocalName() string {
	return vscpg.Name
}

// Configuration returns the configuration (args) for [VpnServerConfigurationPolicyGroup].
func (vscpg *VpnServerConfigurationPolicyGroup) Configuration() interface{} {
	return vscpg.Args
}

// DependOn is used for other resources to depend on [VpnServerConfigurationPolicyGroup].
func (vscpg *VpnServerConfigurationPolicyGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(vscpg)
}

// Dependencies returns the list of resources [VpnServerConfigurationPolicyGroup] depends_on.
func (vscpg *VpnServerConfigurationPolicyGroup) Dependencies() terra.Dependencies {
	return vscpg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpnServerConfigurationPolicyGroup].
func (vscpg *VpnServerConfigurationPolicyGroup) LifecycleManagement() *terra.Lifecycle {
	return vscpg.Lifecycle
}

// Attributes returns the attributes for [VpnServerConfigurationPolicyGroup].
func (vscpg *VpnServerConfigurationPolicyGroup) Attributes() vpnServerConfigurationPolicyGroupAttributes {
	return vpnServerConfigurationPolicyGroupAttributes{ref: terra.ReferenceResource(vscpg)}
}

// ImportState imports the given attribute values into [VpnServerConfigurationPolicyGroup]'s state.
func (vscpg *VpnServerConfigurationPolicyGroup) ImportState(av io.Reader) error {
	vscpg.state = &vpnServerConfigurationPolicyGroupState{}
	if err := json.NewDecoder(av).Decode(vscpg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vscpg.Type(), vscpg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpnServerConfigurationPolicyGroup] has state.
func (vscpg *VpnServerConfigurationPolicyGroup) State() (*vpnServerConfigurationPolicyGroupState, bool) {
	return vscpg.state, vscpg.state != nil
}

// StateMust returns the state for [VpnServerConfigurationPolicyGroup]. Panics if the state is nil.
func (vscpg *VpnServerConfigurationPolicyGroup) StateMust() *vpnServerConfigurationPolicyGroupState {
	if vscpg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vscpg.Type(), vscpg.LocalName()))
	}
	return vscpg.state
}

// VpnServerConfigurationPolicyGroupArgs contains the configurations for azurerm_vpn_server_configuration_policy_group.
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
}
type vpnServerConfigurationPolicyGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_vpn_server_configuration_policy_group.
func (vscpg vpnServerConfigurationPolicyGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vscpg.ref.Append("id"))
}

// IsDefault returns a reference to field is_default of azurerm_vpn_server_configuration_policy_group.
func (vscpg vpnServerConfigurationPolicyGroupAttributes) IsDefault() terra.BoolValue {
	return terra.ReferenceAsBool(vscpg.ref.Append("is_default"))
}

// Name returns a reference to field name of azurerm_vpn_server_configuration_policy_group.
func (vscpg vpnServerConfigurationPolicyGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vscpg.ref.Append("name"))
}

// Priority returns a reference to field priority of azurerm_vpn_server_configuration_policy_group.
func (vscpg vpnServerConfigurationPolicyGroupAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(vscpg.ref.Append("priority"))
}

// VpnServerConfigurationId returns a reference to field vpn_server_configuration_id of azurerm_vpn_server_configuration_policy_group.
func (vscpg vpnServerConfigurationPolicyGroupAttributes) VpnServerConfigurationId() terra.StringValue {
	return terra.ReferenceAsString(vscpg.ref.Append("vpn_server_configuration_id"))
}

func (vscpg vpnServerConfigurationPolicyGroupAttributes) Policy() terra.SetValue[vpnserverconfigurationpolicygroup.PolicyAttributes] {
	return terra.ReferenceAsSet[vpnserverconfigurationpolicygroup.PolicyAttributes](vscpg.ref.Append("policy"))
}

func (vscpg vpnServerConfigurationPolicyGroupAttributes) Timeouts() vpnserverconfigurationpolicygroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpnserverconfigurationpolicygroup.TimeoutsAttributes](vscpg.ref.Append("timeouts"))
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
