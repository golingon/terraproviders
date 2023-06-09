// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networksecurityrule "github.com/golingon/terraproviders/azurerm/3.58.0/networksecurityrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkSecurityRule creates a new instance of [NetworkSecurityRule].
func NewNetworkSecurityRule(name string, args NetworkSecurityRuleArgs) *NetworkSecurityRule {
	return &NetworkSecurityRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkSecurityRule)(nil)

// NetworkSecurityRule represents the Terraform resource azurerm_network_security_rule.
type NetworkSecurityRule struct {
	Name      string
	Args      NetworkSecurityRuleArgs
	state     *networkSecurityRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkSecurityRule].
func (nsr *NetworkSecurityRule) Type() string {
	return "azurerm_network_security_rule"
}

// LocalName returns the local name for [NetworkSecurityRule].
func (nsr *NetworkSecurityRule) LocalName() string {
	return nsr.Name
}

// Configuration returns the configuration (args) for [NetworkSecurityRule].
func (nsr *NetworkSecurityRule) Configuration() interface{} {
	return nsr.Args
}

// DependOn is used for other resources to depend on [NetworkSecurityRule].
func (nsr *NetworkSecurityRule) DependOn() terra.Reference {
	return terra.ReferenceResource(nsr)
}

// Dependencies returns the list of resources [NetworkSecurityRule] depends_on.
func (nsr *NetworkSecurityRule) Dependencies() terra.Dependencies {
	return nsr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkSecurityRule].
func (nsr *NetworkSecurityRule) LifecycleManagement() *terra.Lifecycle {
	return nsr.Lifecycle
}

// Attributes returns the attributes for [NetworkSecurityRule].
func (nsr *NetworkSecurityRule) Attributes() networkSecurityRuleAttributes {
	return networkSecurityRuleAttributes{ref: terra.ReferenceResource(nsr)}
}

// ImportState imports the given attribute values into [NetworkSecurityRule]'s state.
func (nsr *NetworkSecurityRule) ImportState(av io.Reader) error {
	nsr.state = &networkSecurityRuleState{}
	if err := json.NewDecoder(av).Decode(nsr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nsr.Type(), nsr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkSecurityRule] has state.
func (nsr *NetworkSecurityRule) State() (*networkSecurityRuleState, bool) {
	return nsr.state, nsr.state != nil
}

// StateMust returns the state for [NetworkSecurityRule]. Panics if the state is nil.
func (nsr *NetworkSecurityRule) StateMust() *networkSecurityRuleState {
	if nsr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nsr.Type(), nsr.LocalName()))
	}
	return nsr.state
}

// NetworkSecurityRuleArgs contains the configurations for azurerm_network_security_rule.
type NetworkSecurityRuleArgs struct {
	// Access: string, required
	Access terra.StringValue `hcl:"access,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DestinationAddressPrefix: string, optional
	DestinationAddressPrefix terra.StringValue `hcl:"destination_address_prefix,attr"`
	// DestinationAddressPrefixes: set of string, optional
	DestinationAddressPrefixes terra.SetValue[terra.StringValue] `hcl:"destination_address_prefixes,attr"`
	// DestinationApplicationSecurityGroupIds: set of string, optional
	DestinationApplicationSecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"destination_application_security_group_ids,attr"`
	// DestinationPortRange: string, optional
	DestinationPortRange terra.StringValue `hcl:"destination_port_range,attr"`
	// DestinationPortRanges: set of string, optional
	DestinationPortRanges terra.SetValue[terra.StringValue] `hcl:"destination_port_ranges,attr"`
	// Direction: string, required
	Direction terra.StringValue `hcl:"direction,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkSecurityGroupName: string, required
	NetworkSecurityGroupName terra.StringValue `hcl:"network_security_group_name,attr" validate:"required"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SourceAddressPrefix: string, optional
	SourceAddressPrefix terra.StringValue `hcl:"source_address_prefix,attr"`
	// SourceAddressPrefixes: set of string, optional
	SourceAddressPrefixes terra.SetValue[terra.StringValue] `hcl:"source_address_prefixes,attr"`
	// SourceApplicationSecurityGroupIds: set of string, optional
	SourceApplicationSecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"source_application_security_group_ids,attr"`
	// SourcePortRange: string, optional
	SourcePortRange terra.StringValue `hcl:"source_port_range,attr"`
	// SourcePortRanges: set of string, optional
	SourcePortRanges terra.SetValue[terra.StringValue] `hcl:"source_port_ranges,attr"`
	// Timeouts: optional
	Timeouts *networksecurityrule.Timeouts `hcl:"timeouts,block"`
}
type networkSecurityRuleAttributes struct {
	ref terra.Reference
}

// Access returns a reference to field access of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) Access() terra.StringValue {
	return terra.ReferenceAsString(nsr.ref.Append("access"))
}

// Description returns a reference to field description of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nsr.ref.Append("description"))
}

// DestinationAddressPrefix returns a reference to field destination_address_prefix of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) DestinationAddressPrefix() terra.StringValue {
	return terra.ReferenceAsString(nsr.ref.Append("destination_address_prefix"))
}

// DestinationAddressPrefixes returns a reference to field destination_address_prefixes of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) DestinationAddressPrefixes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nsr.ref.Append("destination_address_prefixes"))
}

// DestinationApplicationSecurityGroupIds returns a reference to field destination_application_security_group_ids of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) DestinationApplicationSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nsr.ref.Append("destination_application_security_group_ids"))
}

// DestinationPortRange returns a reference to field destination_port_range of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) DestinationPortRange() terra.StringValue {
	return terra.ReferenceAsString(nsr.ref.Append("destination_port_range"))
}

// DestinationPortRanges returns a reference to field destination_port_ranges of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) DestinationPortRanges() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nsr.ref.Append("destination_port_ranges"))
}

// Direction returns a reference to field direction of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(nsr.ref.Append("direction"))
}

// Id returns a reference to field id of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nsr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nsr.ref.Append("name"))
}

// NetworkSecurityGroupName returns a reference to field network_security_group_name of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) NetworkSecurityGroupName() terra.StringValue {
	return terra.ReferenceAsString(nsr.ref.Append("network_security_group_name"))
}

// Priority returns a reference to field priority of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(nsr.ref.Append("priority"))
}

// Protocol returns a reference to field protocol of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(nsr.ref.Append("protocol"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nsr.ref.Append("resource_group_name"))
}

// SourceAddressPrefix returns a reference to field source_address_prefix of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) SourceAddressPrefix() terra.StringValue {
	return terra.ReferenceAsString(nsr.ref.Append("source_address_prefix"))
}

// SourceAddressPrefixes returns a reference to field source_address_prefixes of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) SourceAddressPrefixes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nsr.ref.Append("source_address_prefixes"))
}

// SourceApplicationSecurityGroupIds returns a reference to field source_application_security_group_ids of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) SourceApplicationSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nsr.ref.Append("source_application_security_group_ids"))
}

// SourcePortRange returns a reference to field source_port_range of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) SourcePortRange() terra.StringValue {
	return terra.ReferenceAsString(nsr.ref.Append("source_port_range"))
}

// SourcePortRanges returns a reference to field source_port_ranges of azurerm_network_security_rule.
func (nsr networkSecurityRuleAttributes) SourcePortRanges() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nsr.ref.Append("source_port_ranges"))
}

func (nsr networkSecurityRuleAttributes) Timeouts() networksecurityrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networksecurityrule.TimeoutsAttributes](nsr.ref.Append("timeouts"))
}

type networkSecurityRuleState struct {
	Access                                 string                             `json:"access"`
	Description                            string                             `json:"description"`
	DestinationAddressPrefix               string                             `json:"destination_address_prefix"`
	DestinationAddressPrefixes             []string                           `json:"destination_address_prefixes"`
	DestinationApplicationSecurityGroupIds []string                           `json:"destination_application_security_group_ids"`
	DestinationPortRange                   string                             `json:"destination_port_range"`
	DestinationPortRanges                  []string                           `json:"destination_port_ranges"`
	Direction                              string                             `json:"direction"`
	Id                                     string                             `json:"id"`
	Name                                   string                             `json:"name"`
	NetworkSecurityGroupName               string                             `json:"network_security_group_name"`
	Priority                               float64                            `json:"priority"`
	Protocol                               string                             `json:"protocol"`
	ResourceGroupName                      string                             `json:"resource_group_name"`
	SourceAddressPrefix                    string                             `json:"source_address_prefix"`
	SourceAddressPrefixes                  []string                           `json:"source_address_prefixes"`
	SourceApplicationSecurityGroupIds      []string                           `json:"source_application_security_group_ids"`
	SourcePortRange                        string                             `json:"source_port_range"`
	SourcePortRanges                       []string                           `json:"source_port_ranges"`
	Timeouts                               *networksecurityrule.TimeoutsState `json:"timeouts"`
}
