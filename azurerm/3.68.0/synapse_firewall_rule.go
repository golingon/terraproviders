// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	synapsefirewallrule "github.com/golingon/terraproviders/azurerm/3.68.0/synapsefirewallrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSynapseFirewallRule creates a new instance of [SynapseFirewallRule].
func NewSynapseFirewallRule(name string, args SynapseFirewallRuleArgs) *SynapseFirewallRule {
	return &SynapseFirewallRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SynapseFirewallRule)(nil)

// SynapseFirewallRule represents the Terraform resource azurerm_synapse_firewall_rule.
type SynapseFirewallRule struct {
	Name      string
	Args      SynapseFirewallRuleArgs
	state     *synapseFirewallRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SynapseFirewallRule].
func (sfr *SynapseFirewallRule) Type() string {
	return "azurerm_synapse_firewall_rule"
}

// LocalName returns the local name for [SynapseFirewallRule].
func (sfr *SynapseFirewallRule) LocalName() string {
	return sfr.Name
}

// Configuration returns the configuration (args) for [SynapseFirewallRule].
func (sfr *SynapseFirewallRule) Configuration() interface{} {
	return sfr.Args
}

// DependOn is used for other resources to depend on [SynapseFirewallRule].
func (sfr *SynapseFirewallRule) DependOn() terra.Reference {
	return terra.ReferenceResource(sfr)
}

// Dependencies returns the list of resources [SynapseFirewallRule] depends_on.
func (sfr *SynapseFirewallRule) Dependencies() terra.Dependencies {
	return sfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SynapseFirewallRule].
func (sfr *SynapseFirewallRule) LifecycleManagement() *terra.Lifecycle {
	return sfr.Lifecycle
}

// Attributes returns the attributes for [SynapseFirewallRule].
func (sfr *SynapseFirewallRule) Attributes() synapseFirewallRuleAttributes {
	return synapseFirewallRuleAttributes{ref: terra.ReferenceResource(sfr)}
}

// ImportState imports the given attribute values into [SynapseFirewallRule]'s state.
func (sfr *SynapseFirewallRule) ImportState(av io.Reader) error {
	sfr.state = &synapseFirewallRuleState{}
	if err := json.NewDecoder(av).Decode(sfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sfr.Type(), sfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SynapseFirewallRule] has state.
func (sfr *SynapseFirewallRule) State() (*synapseFirewallRuleState, bool) {
	return sfr.state, sfr.state != nil
}

// StateMust returns the state for [SynapseFirewallRule]. Panics if the state is nil.
func (sfr *SynapseFirewallRule) StateMust() *synapseFirewallRuleState {
	if sfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sfr.Type(), sfr.LocalName()))
	}
	return sfr.state
}

// SynapseFirewallRuleArgs contains the configurations for azurerm_synapse_firewall_rule.
type SynapseFirewallRuleArgs struct {
	// EndIpAddress: string, required
	EndIpAddress terra.StringValue `hcl:"end_ip_address,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StartIpAddress: string, required
	StartIpAddress terra.StringValue `hcl:"start_ip_address,attr" validate:"required"`
	// SynapseWorkspaceId: string, required
	SynapseWorkspaceId terra.StringValue `hcl:"synapse_workspace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *synapsefirewallrule.Timeouts `hcl:"timeouts,block"`
}
type synapseFirewallRuleAttributes struct {
	ref terra.Reference
}

// EndIpAddress returns a reference to field end_ip_address of azurerm_synapse_firewall_rule.
func (sfr synapseFirewallRuleAttributes) EndIpAddress() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("end_ip_address"))
}

// Id returns a reference to field id of azurerm_synapse_firewall_rule.
func (sfr synapseFirewallRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_synapse_firewall_rule.
func (sfr synapseFirewallRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("name"))
}

// StartIpAddress returns a reference to field start_ip_address of azurerm_synapse_firewall_rule.
func (sfr synapseFirewallRuleAttributes) StartIpAddress() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("start_ip_address"))
}

// SynapseWorkspaceId returns a reference to field synapse_workspace_id of azurerm_synapse_firewall_rule.
func (sfr synapseFirewallRuleAttributes) SynapseWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("synapse_workspace_id"))
}

func (sfr synapseFirewallRuleAttributes) Timeouts() synapsefirewallrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[synapsefirewallrule.TimeoutsAttributes](sfr.ref.Append("timeouts"))
}

type synapseFirewallRuleState struct {
	EndIpAddress       string                             `json:"end_ip_address"`
	Id                 string                             `json:"id"`
	Name               string                             `json:"name"`
	StartIpAddress     string                             `json:"start_ip_address"`
	SynapseWorkspaceId string                             `json:"synapse_workspace_id"`
	Timeouts           *synapsefirewallrule.TimeoutsState `json:"timeouts"`
}
