// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_network_manager_admin_rule

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

// Resource represents the Terraform resource azurerm_network_manager_admin_rule.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermNetworkManagerAdminRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (anmar *Resource) Type() string {
	return "azurerm_network_manager_admin_rule"
}

// LocalName returns the local name for [Resource].
func (anmar *Resource) LocalName() string {
	return anmar.Name
}

// Configuration returns the configuration (args) for [Resource].
func (anmar *Resource) Configuration() interface{} {
	return anmar.Args
}

// DependOn is used for other resources to depend on [Resource].
func (anmar *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(anmar)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (anmar *Resource) Dependencies() terra.Dependencies {
	return anmar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (anmar *Resource) LifecycleManagement() *terra.Lifecycle {
	return anmar.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (anmar *Resource) Attributes() azurermNetworkManagerAdminRuleAttributes {
	return azurermNetworkManagerAdminRuleAttributes{ref: terra.ReferenceResource(anmar)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (anmar *Resource) ImportState(state io.Reader) error {
	anmar.state = &azurermNetworkManagerAdminRuleState{}
	if err := json.NewDecoder(state).Decode(anmar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", anmar.Type(), anmar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (anmar *Resource) State() (*azurermNetworkManagerAdminRuleState, bool) {
	return anmar.state, anmar.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (anmar *Resource) StateMust() *azurermNetworkManagerAdminRuleState {
	if anmar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", anmar.Type(), anmar.LocalName()))
	}
	return anmar.state
}

// Args contains the configurations for azurerm_network_manager_admin_rule.
type Args struct {
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
	Destination []Destination `hcl:"destination,block" validate:"min=0"`
	// Source: min=0
	Source []Source `hcl:"source,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermNetworkManagerAdminRuleAttributes struct {
	ref terra.Reference
}

// Action returns a reference to field action of azurerm_network_manager_admin_rule.
func (anmar azurermNetworkManagerAdminRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(anmar.ref.Append("action"))
}

// AdminRuleCollectionId returns a reference to field admin_rule_collection_id of azurerm_network_manager_admin_rule.
func (anmar azurermNetworkManagerAdminRuleAttributes) AdminRuleCollectionId() terra.StringValue {
	return terra.ReferenceAsString(anmar.ref.Append("admin_rule_collection_id"))
}

// Description returns a reference to field description of azurerm_network_manager_admin_rule.
func (anmar azurermNetworkManagerAdminRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(anmar.ref.Append("description"))
}

// DestinationPortRanges returns a reference to field destination_port_ranges of azurerm_network_manager_admin_rule.
func (anmar azurermNetworkManagerAdminRuleAttributes) DestinationPortRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](anmar.ref.Append("destination_port_ranges"))
}

// Direction returns a reference to field direction of azurerm_network_manager_admin_rule.
func (anmar azurermNetworkManagerAdminRuleAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(anmar.ref.Append("direction"))
}

// Id returns a reference to field id of azurerm_network_manager_admin_rule.
func (anmar azurermNetworkManagerAdminRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(anmar.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_network_manager_admin_rule.
func (anmar azurermNetworkManagerAdminRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(anmar.ref.Append("name"))
}

// Priority returns a reference to field priority of azurerm_network_manager_admin_rule.
func (anmar azurermNetworkManagerAdminRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(anmar.ref.Append("priority"))
}

// Protocol returns a reference to field protocol of azurerm_network_manager_admin_rule.
func (anmar azurermNetworkManagerAdminRuleAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(anmar.ref.Append("protocol"))
}

// SourcePortRanges returns a reference to field source_port_ranges of azurerm_network_manager_admin_rule.
func (anmar azurermNetworkManagerAdminRuleAttributes) SourcePortRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](anmar.ref.Append("source_port_ranges"))
}

func (anmar azurermNetworkManagerAdminRuleAttributes) Destination() terra.ListValue[DestinationAttributes] {
	return terra.ReferenceAsList[DestinationAttributes](anmar.ref.Append("destination"))
}

func (anmar azurermNetworkManagerAdminRuleAttributes) Source() terra.ListValue[SourceAttributes] {
	return terra.ReferenceAsList[SourceAttributes](anmar.ref.Append("source"))
}

func (anmar azurermNetworkManagerAdminRuleAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](anmar.ref.Append("timeouts"))
}

type azurermNetworkManagerAdminRuleState struct {
	Action                string             `json:"action"`
	AdminRuleCollectionId string             `json:"admin_rule_collection_id"`
	Description           string             `json:"description"`
	DestinationPortRanges []string           `json:"destination_port_ranges"`
	Direction             string             `json:"direction"`
	Id                    string             `json:"id"`
	Name                  string             `json:"name"`
	Priority              float64            `json:"priority"`
	Protocol              string             `json:"protocol"`
	SourcePortRanges      []string           `json:"source_port_ranges"`
	Destination           []DestinationState `json:"destination"`
	Source                []SourceState      `json:"source"`
	Timeouts              *TimeoutsState     `json:"timeouts"`
}
