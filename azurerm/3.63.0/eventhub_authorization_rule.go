// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	eventhubauthorizationrule "github.com/golingon/terraproviders/azurerm/3.63.0/eventhubauthorizationrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEventhubAuthorizationRule creates a new instance of [EventhubAuthorizationRule].
func NewEventhubAuthorizationRule(name string, args EventhubAuthorizationRuleArgs) *EventhubAuthorizationRule {
	return &EventhubAuthorizationRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EventhubAuthorizationRule)(nil)

// EventhubAuthorizationRule represents the Terraform resource azurerm_eventhub_authorization_rule.
type EventhubAuthorizationRule struct {
	Name      string
	Args      EventhubAuthorizationRuleArgs
	state     *eventhubAuthorizationRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EventhubAuthorizationRule].
func (ear *EventhubAuthorizationRule) Type() string {
	return "azurerm_eventhub_authorization_rule"
}

// LocalName returns the local name for [EventhubAuthorizationRule].
func (ear *EventhubAuthorizationRule) LocalName() string {
	return ear.Name
}

// Configuration returns the configuration (args) for [EventhubAuthorizationRule].
func (ear *EventhubAuthorizationRule) Configuration() interface{} {
	return ear.Args
}

// DependOn is used for other resources to depend on [EventhubAuthorizationRule].
func (ear *EventhubAuthorizationRule) DependOn() terra.Reference {
	return terra.ReferenceResource(ear)
}

// Dependencies returns the list of resources [EventhubAuthorizationRule] depends_on.
func (ear *EventhubAuthorizationRule) Dependencies() terra.Dependencies {
	return ear.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EventhubAuthorizationRule].
func (ear *EventhubAuthorizationRule) LifecycleManagement() *terra.Lifecycle {
	return ear.Lifecycle
}

// Attributes returns the attributes for [EventhubAuthorizationRule].
func (ear *EventhubAuthorizationRule) Attributes() eventhubAuthorizationRuleAttributes {
	return eventhubAuthorizationRuleAttributes{ref: terra.ReferenceResource(ear)}
}

// ImportState imports the given attribute values into [EventhubAuthorizationRule]'s state.
func (ear *EventhubAuthorizationRule) ImportState(av io.Reader) error {
	ear.state = &eventhubAuthorizationRuleState{}
	if err := json.NewDecoder(av).Decode(ear.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ear.Type(), ear.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EventhubAuthorizationRule] has state.
func (ear *EventhubAuthorizationRule) State() (*eventhubAuthorizationRuleState, bool) {
	return ear.state, ear.state != nil
}

// StateMust returns the state for [EventhubAuthorizationRule]. Panics if the state is nil.
func (ear *EventhubAuthorizationRule) StateMust() *eventhubAuthorizationRuleState {
	if ear.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ear.Type(), ear.LocalName()))
	}
	return ear.state
}

// EventhubAuthorizationRuleArgs contains the configurations for azurerm_eventhub_authorization_rule.
type EventhubAuthorizationRuleArgs struct {
	// EventhubName: string, required
	EventhubName terra.StringValue `hcl:"eventhub_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Listen: bool, optional
	Listen terra.BoolValue `hcl:"listen,attr"`
	// Manage: bool, optional
	Manage terra.BoolValue `hcl:"manage,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceName: string, required
	NamespaceName terra.StringValue `hcl:"namespace_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Send: bool, optional
	Send terra.BoolValue `hcl:"send,attr"`
	// Timeouts: optional
	Timeouts *eventhubauthorizationrule.Timeouts `hcl:"timeouts,block"`
}
type eventhubAuthorizationRuleAttributes struct {
	ref terra.Reference
}

// EventhubName returns a reference to field eventhub_name of azurerm_eventhub_authorization_rule.
func (ear eventhubAuthorizationRuleAttributes) EventhubName() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("eventhub_name"))
}

// Id returns a reference to field id of azurerm_eventhub_authorization_rule.
func (ear eventhubAuthorizationRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("id"))
}

// Listen returns a reference to field listen of azurerm_eventhub_authorization_rule.
func (ear eventhubAuthorizationRuleAttributes) Listen() terra.BoolValue {
	return terra.ReferenceAsBool(ear.ref.Append("listen"))
}

// Manage returns a reference to field manage of azurerm_eventhub_authorization_rule.
func (ear eventhubAuthorizationRuleAttributes) Manage() terra.BoolValue {
	return terra.ReferenceAsBool(ear.ref.Append("manage"))
}

// Name returns a reference to field name of azurerm_eventhub_authorization_rule.
func (ear eventhubAuthorizationRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("name"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_eventhub_authorization_rule.
func (ear eventhubAuthorizationRuleAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("namespace_name"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_eventhub_authorization_rule.
func (ear eventhubAuthorizationRuleAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("primary_connection_string"))
}

// PrimaryConnectionStringAlias returns a reference to field primary_connection_string_alias of azurerm_eventhub_authorization_rule.
func (ear eventhubAuthorizationRuleAttributes) PrimaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("primary_connection_string_alias"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_eventhub_authorization_rule.
func (ear eventhubAuthorizationRuleAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("primary_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventhub_authorization_rule.
func (ear eventhubAuthorizationRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("resource_group_name"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_eventhub_authorization_rule.
func (ear eventhubAuthorizationRuleAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("secondary_connection_string"))
}

// SecondaryConnectionStringAlias returns a reference to field secondary_connection_string_alias of azurerm_eventhub_authorization_rule.
func (ear eventhubAuthorizationRuleAttributes) SecondaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("secondary_connection_string_alias"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_eventhub_authorization_rule.
func (ear eventhubAuthorizationRuleAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("secondary_key"))
}

// Send returns a reference to field send of azurerm_eventhub_authorization_rule.
func (ear eventhubAuthorizationRuleAttributes) Send() terra.BoolValue {
	return terra.ReferenceAsBool(ear.ref.Append("send"))
}

func (ear eventhubAuthorizationRuleAttributes) Timeouts() eventhubauthorizationrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eventhubauthorizationrule.TimeoutsAttributes](ear.ref.Append("timeouts"))
}

type eventhubAuthorizationRuleState struct {
	EventhubName                   string                                   `json:"eventhub_name"`
	Id                             string                                   `json:"id"`
	Listen                         bool                                     `json:"listen"`
	Manage                         bool                                     `json:"manage"`
	Name                           string                                   `json:"name"`
	NamespaceName                  string                                   `json:"namespace_name"`
	PrimaryConnectionString        string                                   `json:"primary_connection_string"`
	PrimaryConnectionStringAlias   string                                   `json:"primary_connection_string_alias"`
	PrimaryKey                     string                                   `json:"primary_key"`
	ResourceGroupName              string                                   `json:"resource_group_name"`
	SecondaryConnectionString      string                                   `json:"secondary_connection_string"`
	SecondaryConnectionStringAlias string                                   `json:"secondary_connection_string_alias"`
	SecondaryKey                   string                                   `json:"secondary_key"`
	Send                           bool                                     `json:"send"`
	Timeouts                       *eventhubauthorizationrule.TimeoutsState `json:"timeouts"`
}
