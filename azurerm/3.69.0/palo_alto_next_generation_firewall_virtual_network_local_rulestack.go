// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	paloaltonextgenerationfirewallvirtualnetworklocalrulestack "github.com/golingon/terraproviders/azurerm/3.69.0/paloaltonextgenerationfirewallvirtualnetworklocalrulestack"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack creates a new instance of [PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack].
func NewPaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack(name string, args PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestackArgs) *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack {
	return &PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack)(nil)

// PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack represents the Terraform resource azurerm_palo_alto_next_generation_firewall_virtual_network_local_rulestack.
type PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack struct {
	Name      string
	Args      PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestackArgs
	state     *paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack].
func (pangfvnlr *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack) Type() string {
	return "azurerm_palo_alto_next_generation_firewall_virtual_network_local_rulestack"
}

// LocalName returns the local name for [PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack].
func (pangfvnlr *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack) LocalName() string {
	return pangfvnlr.Name
}

// Configuration returns the configuration (args) for [PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack].
func (pangfvnlr *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack) Configuration() interface{} {
	return pangfvnlr.Args
}

// DependOn is used for other resources to depend on [PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack].
func (pangfvnlr *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack) DependOn() terra.Reference {
	return terra.ReferenceResource(pangfvnlr)
}

// Dependencies returns the list of resources [PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack] depends_on.
func (pangfvnlr *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack) Dependencies() terra.Dependencies {
	return pangfvnlr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack].
func (pangfvnlr *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack) LifecycleManagement() *terra.Lifecycle {
	return pangfvnlr.Lifecycle
}

// Attributes returns the attributes for [PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack].
func (pangfvnlr *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack) Attributes() paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackAttributes {
	return paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackAttributes{ref: terra.ReferenceResource(pangfvnlr)}
}

// ImportState imports the given attribute values into [PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack]'s state.
func (pangfvnlr *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack) ImportState(av io.Reader) error {
	pangfvnlr.state = &paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackState{}
	if err := json.NewDecoder(av).Decode(pangfvnlr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pangfvnlr.Type(), pangfvnlr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack] has state.
func (pangfvnlr *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack) State() (*paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackState, bool) {
	return pangfvnlr.state, pangfvnlr.state != nil
}

// StateMust returns the state for [PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack]. Panics if the state is nil.
func (pangfvnlr *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack) StateMust() *paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackState {
	if pangfvnlr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pangfvnlr.Type(), pangfvnlr.LocalName()))
	}
	return pangfvnlr.state
}

// PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestackArgs contains the configurations for azurerm_palo_alto_next_generation_firewall_virtual_network_local_rulestack.
type PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestackArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RulestackId: string, required
	RulestackId terra.StringValue `hcl:"rulestack_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// DestinationNat: min=0
	DestinationNat []paloaltonextgenerationfirewallvirtualnetworklocalrulestack.DestinationNat `hcl:"destination_nat,block" validate:"min=0"`
	// DnsSettings: optional
	DnsSettings *paloaltonextgenerationfirewallvirtualnetworklocalrulestack.DnsSettings `hcl:"dns_settings,block"`
	// NetworkProfile: required
	NetworkProfile *paloaltonextgenerationfirewallvirtualnetworklocalrulestack.NetworkProfile `hcl:"network_profile,block" validate:"required"`
	// Timeouts: optional
	Timeouts *paloaltonextgenerationfirewallvirtualnetworklocalrulestack.Timeouts `hcl:"timeouts,block"`
}
type paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_palo_alto_next_generation_firewall_virtual_network_local_rulestack.
func (pangfvnlr paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pangfvnlr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_palo_alto_next_generation_firewall_virtual_network_local_rulestack.
func (pangfvnlr paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pangfvnlr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_palo_alto_next_generation_firewall_virtual_network_local_rulestack.
func (pangfvnlr paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pangfvnlr.ref.Append("resource_group_name"))
}

// RulestackId returns a reference to field rulestack_id of azurerm_palo_alto_next_generation_firewall_virtual_network_local_rulestack.
func (pangfvnlr paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackAttributes) RulestackId() terra.StringValue {
	return terra.ReferenceAsString(pangfvnlr.ref.Append("rulestack_id"))
}

// Tags returns a reference to field tags of azurerm_palo_alto_next_generation_firewall_virtual_network_local_rulestack.
func (pangfvnlr paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pangfvnlr.ref.Append("tags"))
}

func (pangfvnlr paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackAttributes) DestinationNat() terra.ListValue[paloaltonextgenerationfirewallvirtualnetworklocalrulestack.DestinationNatAttributes] {
	return terra.ReferenceAsList[paloaltonextgenerationfirewallvirtualnetworklocalrulestack.DestinationNatAttributes](pangfvnlr.ref.Append("destination_nat"))
}

func (pangfvnlr paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackAttributes) DnsSettings() terra.ListValue[paloaltonextgenerationfirewallvirtualnetworklocalrulestack.DnsSettingsAttributes] {
	return terra.ReferenceAsList[paloaltonextgenerationfirewallvirtualnetworklocalrulestack.DnsSettingsAttributes](pangfvnlr.ref.Append("dns_settings"))
}

func (pangfvnlr paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackAttributes) NetworkProfile() terra.ListValue[paloaltonextgenerationfirewallvirtualnetworklocalrulestack.NetworkProfileAttributes] {
	return terra.ReferenceAsList[paloaltonextgenerationfirewallvirtualnetworklocalrulestack.NetworkProfileAttributes](pangfvnlr.ref.Append("network_profile"))
}

func (pangfvnlr paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackAttributes) Timeouts() paloaltonextgenerationfirewallvirtualnetworklocalrulestack.TimeoutsAttributes {
	return terra.ReferenceAsSingle[paloaltonextgenerationfirewallvirtualnetworklocalrulestack.TimeoutsAttributes](pangfvnlr.ref.Append("timeouts"))
}

type paloAltoNextGenerationFirewallVirtualNetworkLocalRulestackState struct {
	Id                string                                                                           `json:"id"`
	Name              string                                                                           `json:"name"`
	ResourceGroupName string                                                                           `json:"resource_group_name"`
	RulestackId       string                                                                           `json:"rulestack_id"`
	Tags              map[string]string                                                                `json:"tags"`
	DestinationNat    []paloaltonextgenerationfirewallvirtualnetworklocalrulestack.DestinationNatState `json:"destination_nat"`
	DnsSettings       []paloaltonextgenerationfirewallvirtualnetworklocalrulestack.DnsSettingsState    `json:"dns_settings"`
	NetworkProfile    []paloaltonextgenerationfirewallvirtualnetworklocalrulestack.NetworkProfileState `json:"network_profile"`
	Timeouts          *paloaltonextgenerationfirewallvirtualnetworklocalrulestack.TimeoutsState        `json:"timeouts"`
}