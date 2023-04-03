// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	eventhubnamespaceauthorizationrule "github.com/golingon/terraproviders/azurerm/3.49.0/eventhubnamespaceauthorizationrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEventhubNamespaceAuthorizationRule creates a new instance of [EventhubNamespaceAuthorizationRule].
func NewEventhubNamespaceAuthorizationRule(name string, args EventhubNamespaceAuthorizationRuleArgs) *EventhubNamespaceAuthorizationRule {
	return &EventhubNamespaceAuthorizationRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EventhubNamespaceAuthorizationRule)(nil)

// EventhubNamespaceAuthorizationRule represents the Terraform resource azurerm_eventhub_namespace_authorization_rule.
type EventhubNamespaceAuthorizationRule struct {
	Name      string
	Args      EventhubNamespaceAuthorizationRuleArgs
	state     *eventhubNamespaceAuthorizationRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EventhubNamespaceAuthorizationRule].
func (enar *EventhubNamespaceAuthorizationRule) Type() string {
	return "azurerm_eventhub_namespace_authorization_rule"
}

// LocalName returns the local name for [EventhubNamespaceAuthorizationRule].
func (enar *EventhubNamespaceAuthorizationRule) LocalName() string {
	return enar.Name
}

// Configuration returns the configuration (args) for [EventhubNamespaceAuthorizationRule].
func (enar *EventhubNamespaceAuthorizationRule) Configuration() interface{} {
	return enar.Args
}

// DependOn is used for other resources to depend on [EventhubNamespaceAuthorizationRule].
func (enar *EventhubNamespaceAuthorizationRule) DependOn() terra.Reference {
	return terra.ReferenceResource(enar)
}

// Dependencies returns the list of resources [EventhubNamespaceAuthorizationRule] depends_on.
func (enar *EventhubNamespaceAuthorizationRule) Dependencies() terra.Dependencies {
	return enar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EventhubNamespaceAuthorizationRule].
func (enar *EventhubNamespaceAuthorizationRule) LifecycleManagement() *terra.Lifecycle {
	return enar.Lifecycle
}

// Attributes returns the attributes for [EventhubNamespaceAuthorizationRule].
func (enar *EventhubNamespaceAuthorizationRule) Attributes() eventhubNamespaceAuthorizationRuleAttributes {
	return eventhubNamespaceAuthorizationRuleAttributes{ref: terra.ReferenceResource(enar)}
}

// ImportState imports the given attribute values into [EventhubNamespaceAuthorizationRule]'s state.
func (enar *EventhubNamespaceAuthorizationRule) ImportState(av io.Reader) error {
	enar.state = &eventhubNamespaceAuthorizationRuleState{}
	if err := json.NewDecoder(av).Decode(enar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", enar.Type(), enar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EventhubNamespaceAuthorizationRule] has state.
func (enar *EventhubNamespaceAuthorizationRule) State() (*eventhubNamespaceAuthorizationRuleState, bool) {
	return enar.state, enar.state != nil
}

// StateMust returns the state for [EventhubNamespaceAuthorizationRule]. Panics if the state is nil.
func (enar *EventhubNamespaceAuthorizationRule) StateMust() *eventhubNamespaceAuthorizationRuleState {
	if enar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", enar.Type(), enar.LocalName()))
	}
	return enar.state
}

// EventhubNamespaceAuthorizationRuleArgs contains the configurations for azurerm_eventhub_namespace_authorization_rule.
type EventhubNamespaceAuthorizationRuleArgs struct {
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
	Timeouts *eventhubnamespaceauthorizationrule.Timeouts `hcl:"timeouts,block"`
}
type eventhubNamespaceAuthorizationRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_eventhub_namespace_authorization_rule.
func (enar eventhubNamespaceAuthorizationRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("id"))
}

// Listen returns a reference to field listen of azurerm_eventhub_namespace_authorization_rule.
func (enar eventhubNamespaceAuthorizationRuleAttributes) Listen() terra.BoolValue {
	return terra.ReferenceAsBool(enar.ref.Append("listen"))
}

// Manage returns a reference to field manage of azurerm_eventhub_namespace_authorization_rule.
func (enar eventhubNamespaceAuthorizationRuleAttributes) Manage() terra.BoolValue {
	return terra.ReferenceAsBool(enar.ref.Append("manage"))
}

// Name returns a reference to field name of azurerm_eventhub_namespace_authorization_rule.
func (enar eventhubNamespaceAuthorizationRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("name"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_eventhub_namespace_authorization_rule.
func (enar eventhubNamespaceAuthorizationRuleAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("namespace_name"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_eventhub_namespace_authorization_rule.
func (enar eventhubNamespaceAuthorizationRuleAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("primary_connection_string"))
}

// PrimaryConnectionStringAlias returns a reference to field primary_connection_string_alias of azurerm_eventhub_namespace_authorization_rule.
func (enar eventhubNamespaceAuthorizationRuleAttributes) PrimaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("primary_connection_string_alias"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_eventhub_namespace_authorization_rule.
func (enar eventhubNamespaceAuthorizationRuleAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("primary_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventhub_namespace_authorization_rule.
func (enar eventhubNamespaceAuthorizationRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("resource_group_name"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_eventhub_namespace_authorization_rule.
func (enar eventhubNamespaceAuthorizationRuleAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("secondary_connection_string"))
}

// SecondaryConnectionStringAlias returns a reference to field secondary_connection_string_alias of azurerm_eventhub_namespace_authorization_rule.
func (enar eventhubNamespaceAuthorizationRuleAttributes) SecondaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("secondary_connection_string_alias"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_eventhub_namespace_authorization_rule.
func (enar eventhubNamespaceAuthorizationRuleAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("secondary_key"))
}

// Send returns a reference to field send of azurerm_eventhub_namespace_authorization_rule.
func (enar eventhubNamespaceAuthorizationRuleAttributes) Send() terra.BoolValue {
	return terra.ReferenceAsBool(enar.ref.Append("send"))
}

func (enar eventhubNamespaceAuthorizationRuleAttributes) Timeouts() eventhubnamespaceauthorizationrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eventhubnamespaceauthorizationrule.TimeoutsAttributes](enar.ref.Append("timeouts"))
}

type eventhubNamespaceAuthorizationRuleState struct {
	Id                             string                                            `json:"id"`
	Listen                         bool                                              `json:"listen"`
	Manage                         bool                                              `json:"manage"`
	Name                           string                                            `json:"name"`
	NamespaceName                  string                                            `json:"namespace_name"`
	PrimaryConnectionString        string                                            `json:"primary_connection_string"`
	PrimaryConnectionStringAlias   string                                            `json:"primary_connection_string_alias"`
	PrimaryKey                     string                                            `json:"primary_key"`
	ResourceGroupName              string                                            `json:"resource_group_name"`
	SecondaryConnectionString      string                                            `json:"secondary_connection_string"`
	SecondaryConnectionStringAlias string                                            `json:"secondary_connection_string_alias"`
	SecondaryKey                   string                                            `json:"secondary_key"`
	Send                           bool                                              `json:"send"`
	Timeouts                       *eventhubnamespaceauthorizationrule.TimeoutsState `json:"timeouts"`
}
