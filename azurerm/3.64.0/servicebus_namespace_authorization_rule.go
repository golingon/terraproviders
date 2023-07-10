// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	servicebusnamespaceauthorizationrule "github.com/golingon/terraproviders/azurerm/3.64.0/servicebusnamespaceauthorizationrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicebusNamespaceAuthorizationRule creates a new instance of [ServicebusNamespaceAuthorizationRule].
func NewServicebusNamespaceAuthorizationRule(name string, args ServicebusNamespaceAuthorizationRuleArgs) *ServicebusNamespaceAuthorizationRule {
	return &ServicebusNamespaceAuthorizationRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicebusNamespaceAuthorizationRule)(nil)

// ServicebusNamespaceAuthorizationRule represents the Terraform resource azurerm_servicebus_namespace_authorization_rule.
type ServicebusNamespaceAuthorizationRule struct {
	Name      string
	Args      ServicebusNamespaceAuthorizationRuleArgs
	state     *servicebusNamespaceAuthorizationRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicebusNamespaceAuthorizationRule].
func (snar *ServicebusNamespaceAuthorizationRule) Type() string {
	return "azurerm_servicebus_namespace_authorization_rule"
}

// LocalName returns the local name for [ServicebusNamespaceAuthorizationRule].
func (snar *ServicebusNamespaceAuthorizationRule) LocalName() string {
	return snar.Name
}

// Configuration returns the configuration (args) for [ServicebusNamespaceAuthorizationRule].
func (snar *ServicebusNamespaceAuthorizationRule) Configuration() interface{} {
	return snar.Args
}

// DependOn is used for other resources to depend on [ServicebusNamespaceAuthorizationRule].
func (snar *ServicebusNamespaceAuthorizationRule) DependOn() terra.Reference {
	return terra.ReferenceResource(snar)
}

// Dependencies returns the list of resources [ServicebusNamespaceAuthorizationRule] depends_on.
func (snar *ServicebusNamespaceAuthorizationRule) Dependencies() terra.Dependencies {
	return snar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicebusNamespaceAuthorizationRule].
func (snar *ServicebusNamespaceAuthorizationRule) LifecycleManagement() *terra.Lifecycle {
	return snar.Lifecycle
}

// Attributes returns the attributes for [ServicebusNamespaceAuthorizationRule].
func (snar *ServicebusNamespaceAuthorizationRule) Attributes() servicebusNamespaceAuthorizationRuleAttributes {
	return servicebusNamespaceAuthorizationRuleAttributes{ref: terra.ReferenceResource(snar)}
}

// ImportState imports the given attribute values into [ServicebusNamespaceAuthorizationRule]'s state.
func (snar *ServicebusNamespaceAuthorizationRule) ImportState(av io.Reader) error {
	snar.state = &servicebusNamespaceAuthorizationRuleState{}
	if err := json.NewDecoder(av).Decode(snar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", snar.Type(), snar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicebusNamespaceAuthorizationRule] has state.
func (snar *ServicebusNamespaceAuthorizationRule) State() (*servicebusNamespaceAuthorizationRuleState, bool) {
	return snar.state, snar.state != nil
}

// StateMust returns the state for [ServicebusNamespaceAuthorizationRule]. Panics if the state is nil.
func (snar *ServicebusNamespaceAuthorizationRule) StateMust() *servicebusNamespaceAuthorizationRuleState {
	if snar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", snar.Type(), snar.LocalName()))
	}
	return snar.state
}

// ServicebusNamespaceAuthorizationRuleArgs contains the configurations for azurerm_servicebus_namespace_authorization_rule.
type ServicebusNamespaceAuthorizationRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Listen: bool, optional
	Listen terra.BoolValue `hcl:"listen,attr"`
	// Manage: bool, optional
	Manage terra.BoolValue `hcl:"manage,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceId: string, required
	NamespaceId terra.StringValue `hcl:"namespace_id,attr" validate:"required"`
	// Send: bool, optional
	Send terra.BoolValue `hcl:"send,attr"`
	// Timeouts: optional
	Timeouts *servicebusnamespaceauthorizationrule.Timeouts `hcl:"timeouts,block"`
}
type servicebusNamespaceAuthorizationRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_servicebus_namespace_authorization_rule.
func (snar servicebusNamespaceAuthorizationRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("id"))
}

// Listen returns a reference to field listen of azurerm_servicebus_namespace_authorization_rule.
func (snar servicebusNamespaceAuthorizationRuleAttributes) Listen() terra.BoolValue {
	return terra.ReferenceAsBool(snar.ref.Append("listen"))
}

// Manage returns a reference to field manage of azurerm_servicebus_namespace_authorization_rule.
func (snar servicebusNamespaceAuthorizationRuleAttributes) Manage() terra.BoolValue {
	return terra.ReferenceAsBool(snar.ref.Append("manage"))
}

// Name returns a reference to field name of azurerm_servicebus_namespace_authorization_rule.
func (snar servicebusNamespaceAuthorizationRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("name"))
}

// NamespaceId returns a reference to field namespace_id of azurerm_servicebus_namespace_authorization_rule.
func (snar servicebusNamespaceAuthorizationRuleAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("namespace_id"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_servicebus_namespace_authorization_rule.
func (snar servicebusNamespaceAuthorizationRuleAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("primary_connection_string"))
}

// PrimaryConnectionStringAlias returns a reference to field primary_connection_string_alias of azurerm_servicebus_namespace_authorization_rule.
func (snar servicebusNamespaceAuthorizationRuleAttributes) PrimaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("primary_connection_string_alias"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_servicebus_namespace_authorization_rule.
func (snar servicebusNamespaceAuthorizationRuleAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("primary_key"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_servicebus_namespace_authorization_rule.
func (snar servicebusNamespaceAuthorizationRuleAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("secondary_connection_string"))
}

// SecondaryConnectionStringAlias returns a reference to field secondary_connection_string_alias of azurerm_servicebus_namespace_authorization_rule.
func (snar servicebusNamespaceAuthorizationRuleAttributes) SecondaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("secondary_connection_string_alias"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_servicebus_namespace_authorization_rule.
func (snar servicebusNamespaceAuthorizationRuleAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("secondary_key"))
}

// Send returns a reference to field send of azurerm_servicebus_namespace_authorization_rule.
func (snar servicebusNamespaceAuthorizationRuleAttributes) Send() terra.BoolValue {
	return terra.ReferenceAsBool(snar.ref.Append("send"))
}

func (snar servicebusNamespaceAuthorizationRuleAttributes) Timeouts() servicebusnamespaceauthorizationrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicebusnamespaceauthorizationrule.TimeoutsAttributes](snar.ref.Append("timeouts"))
}

type servicebusNamespaceAuthorizationRuleState struct {
	Id                             string                                              `json:"id"`
	Listen                         bool                                                `json:"listen"`
	Manage                         bool                                                `json:"manage"`
	Name                           string                                              `json:"name"`
	NamespaceId                    string                                              `json:"namespace_id"`
	PrimaryConnectionString        string                                              `json:"primary_connection_string"`
	PrimaryConnectionStringAlias   string                                              `json:"primary_connection_string_alias"`
	PrimaryKey                     string                                              `json:"primary_key"`
	SecondaryConnectionString      string                                              `json:"secondary_connection_string"`
	SecondaryConnectionStringAlias string                                              `json:"secondary_connection_string_alias"`
	SecondaryKey                   string                                              `json:"secondary_key"`
	Send                           bool                                                `json:"send"`
	Timeouts                       *servicebusnamespaceauthorizationrule.TimeoutsState `json:"timeouts"`
}
