// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkmanageradminrule "github.com/golingon/terraproviders/azurerm/3.64.0/networkmanageradminrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkManagerAdminRule creates a new instance of [NetworkManagerAdminRule].
func NewNetworkManagerAdminRule(name string, args NetworkManagerAdminRuleArgs) *NetworkManagerAdminRule {
	return &NetworkManagerAdminRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkManagerAdminRule)(nil)

// NetworkManagerAdminRule represents the Terraform resource azurerm_network_manager_admin_rule.
type NetworkManagerAdminRule struct {
	Name      string
	Args      NetworkManagerAdminRuleArgs
	state     *networkManagerAdminRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkManagerAdminRule].
func (nmar *NetworkManagerAdminRule) Type() string {
	return "azurerm_network_manager_admin_rule"
}

// LocalName returns the local name for [NetworkManagerAdminRule].
func (nmar *NetworkManagerAdminRule) LocalName() string {
	return nmar.Name
}

// Configuration returns the configuration (args) for [NetworkManagerAdminRule].
func (nmar *NetworkManagerAdminRule) Configuration() interface{} {
	return nmar.Args
}

// DependOn is used for other resources to depend on [NetworkManagerAdminRule].
func (nmar *NetworkManagerAdminRule) DependOn() terra.Reference {
	return terra.ReferenceResource(nmar)
}

// Dependencies returns the list of resources [NetworkManagerAdminRule] depends_on.
func (nmar *NetworkManagerAdminRule) Dependencies() terra.Dependencies {
	return nmar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkManagerAdminRule].
func (nmar *NetworkManagerAdminRule) LifecycleManagement() *terra.Lifecycle {
	return nmar.Lifecycle
}

// Attributes returns the attributes for [NetworkManagerAdminRule].
func (nmar *NetworkManagerAdminRule) Attributes() networkManagerAdminRuleAttributes {
	return networkManagerAdminRuleAttributes{ref: terra.ReferenceResource(nmar)}
}

// ImportState imports the given attribute values into [NetworkManagerAdminRule]'s state.
func (nmar *NetworkManagerAdminRule) ImportState(av io.Reader) error {
	nmar.state = &networkManagerAdminRuleState{}
	if err := json.NewDecoder(av).Decode(nmar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nmar.Type(), nmar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkManagerAdminRule] has state.
func (nmar *NetworkManagerAdminRule) State() (*networkManagerAdminRuleState, bool) {
	return nmar.state, nmar.state != nil
}

// StateMust returns the state for [NetworkManagerAdminRule]. Panics if the state is nil.
func (nmar *NetworkManagerAdminRule) StateMust() *networkManagerAdminRuleState {
	if nmar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nmar.Type(), nmar.LocalName()))
	}
	return nmar.state
}

// NetworkManagerAdminRuleArgs contains the configurations for azurerm_network_manager_admin_rule.
type NetworkManagerAdminRuleArgs struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// AdminRuleCollectionId: string, required
	AdminRuleCollectionId terra.StringValue `hcl:"admin_rule_collection_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DestinationPortRanges: list of string, optional
	DestinationPortRanges terra.ListValue[terra.StringValue] `hcl:"destination_port_ranges,attr"`
	// Direction: string, required
	Direction terra.StringValue `hcl:"direction,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
	// SourcePortRanges: list of string, optional
	SourcePortRanges terra.ListValue[terra.StringValue] `hcl:"source_port_ranges,attr"`
	// Destination: min=0
	Destination []networkmanageradminrule.Destination `hcl:"destination,block" validate:"min=0"`
	// Source: min=0
	Source []networkmanageradminrule.Source `hcl:"source,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *networkmanageradminrule.Timeouts `hcl:"timeouts,block"`
}
type networkManagerAdminRuleAttributes struct {
	ref terra.Reference
}

// Action returns a reference to field action of azurerm_network_manager_admin_rule.
func (nmar networkManagerAdminRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(nmar.ref.Append("action"))
}

// AdminRuleCollectionId returns a reference to field admin_rule_collection_id of azurerm_network_manager_admin_rule.
func (nmar networkManagerAdminRuleAttributes) AdminRuleCollectionId() terra.StringValue {
	return terra.ReferenceAsString(nmar.ref.Append("admin_rule_collection_id"))
}

// Description returns a reference to field description of azurerm_network_manager_admin_rule.
func (nmar networkManagerAdminRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nmar.ref.Append("description"))
}

// DestinationPortRanges returns a reference to field destination_port_ranges of azurerm_network_manager_admin_rule.
func (nmar networkManagerAdminRuleAttributes) DestinationPortRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nmar.ref.Append("destination_port_ranges"))
}

// Direction returns a reference to field direction of azurerm_network_manager_admin_rule.
func (nmar networkManagerAdminRuleAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(nmar.ref.Append("direction"))
}

// Id returns a reference to field id of azurerm_network_manager_admin_rule.
func (nmar networkManagerAdminRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nmar.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_network_manager_admin_rule.
func (nmar networkManagerAdminRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nmar.ref.Append("name"))
}

// Priority returns a reference to field priority of azurerm_network_manager_admin_rule.
func (nmar networkManagerAdminRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(nmar.ref.Append("priority"))
}

// Protocol returns a reference to field protocol of azurerm_network_manager_admin_rule.
func (nmar networkManagerAdminRuleAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(nmar.ref.Append("protocol"))
}

// SourcePortRanges returns a reference to field source_port_ranges of azurerm_network_manager_admin_rule.
func (nmar networkManagerAdminRuleAttributes) SourcePortRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nmar.ref.Append("source_port_ranges"))
}

func (nmar networkManagerAdminRuleAttributes) Destination() terra.ListValue[networkmanageradminrule.DestinationAttributes] {
	return terra.ReferenceAsList[networkmanageradminrule.DestinationAttributes](nmar.ref.Append("destination"))
}

func (nmar networkManagerAdminRuleAttributes) Source() terra.ListValue[networkmanageradminrule.SourceAttributes] {
	return terra.ReferenceAsList[networkmanageradminrule.SourceAttributes](nmar.ref.Append("source"))
}

func (nmar networkManagerAdminRuleAttributes) Timeouts() networkmanageradminrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanageradminrule.TimeoutsAttributes](nmar.ref.Append("timeouts"))
}

type networkManagerAdminRuleState struct {
	Action                string                                     `json:"action"`
	AdminRuleCollectionId string                                     `json:"admin_rule_collection_id"`
	Description           string                                     `json:"description"`
	DestinationPortRanges []string                                   `json:"destination_port_ranges"`
	Direction             string                                     `json:"direction"`
	Id                    string                                     `json:"id"`
	Name                  string                                     `json:"name"`
	Priority              float64                                    `json:"priority"`
	Protocol              string                                     `json:"protocol"`
	SourcePortRanges      []string                                   `json:"source_port_ranges"`
	Destination           []networkmanageradminrule.DestinationState `json:"destination"`
	Source                []networkmanageradminrule.SourceState      `json:"source"`
	Timeouts              *networkmanageradminrule.TimeoutsState     `json:"timeouts"`
}
