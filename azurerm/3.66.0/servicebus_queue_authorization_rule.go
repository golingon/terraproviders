// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	servicebusqueueauthorizationrule "github.com/golingon/terraproviders/azurerm/3.66.0/servicebusqueueauthorizationrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicebusQueueAuthorizationRule creates a new instance of [ServicebusQueueAuthorizationRule].
func NewServicebusQueueAuthorizationRule(name string, args ServicebusQueueAuthorizationRuleArgs) *ServicebusQueueAuthorizationRule {
	return &ServicebusQueueAuthorizationRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicebusQueueAuthorizationRule)(nil)

// ServicebusQueueAuthorizationRule represents the Terraform resource azurerm_servicebus_queue_authorization_rule.
type ServicebusQueueAuthorizationRule struct {
	Name      string
	Args      ServicebusQueueAuthorizationRuleArgs
	state     *servicebusQueueAuthorizationRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicebusQueueAuthorizationRule].
func (sqar *ServicebusQueueAuthorizationRule) Type() string {
	return "azurerm_servicebus_queue_authorization_rule"
}

// LocalName returns the local name for [ServicebusQueueAuthorizationRule].
func (sqar *ServicebusQueueAuthorizationRule) LocalName() string {
	return sqar.Name
}

// Configuration returns the configuration (args) for [ServicebusQueueAuthorizationRule].
func (sqar *ServicebusQueueAuthorizationRule) Configuration() interface{} {
	return sqar.Args
}

// DependOn is used for other resources to depend on [ServicebusQueueAuthorizationRule].
func (sqar *ServicebusQueueAuthorizationRule) DependOn() terra.Reference {
	return terra.ReferenceResource(sqar)
}

// Dependencies returns the list of resources [ServicebusQueueAuthorizationRule] depends_on.
func (sqar *ServicebusQueueAuthorizationRule) Dependencies() terra.Dependencies {
	return sqar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicebusQueueAuthorizationRule].
func (sqar *ServicebusQueueAuthorizationRule) LifecycleManagement() *terra.Lifecycle {
	return sqar.Lifecycle
}

// Attributes returns the attributes for [ServicebusQueueAuthorizationRule].
func (sqar *ServicebusQueueAuthorizationRule) Attributes() servicebusQueueAuthorizationRuleAttributes {
	return servicebusQueueAuthorizationRuleAttributes{ref: terra.ReferenceResource(sqar)}
}

// ImportState imports the given attribute values into [ServicebusQueueAuthorizationRule]'s state.
func (sqar *ServicebusQueueAuthorizationRule) ImportState(av io.Reader) error {
	sqar.state = &servicebusQueueAuthorizationRuleState{}
	if err := json.NewDecoder(av).Decode(sqar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sqar.Type(), sqar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicebusQueueAuthorizationRule] has state.
func (sqar *ServicebusQueueAuthorizationRule) State() (*servicebusQueueAuthorizationRuleState, bool) {
	return sqar.state, sqar.state != nil
}

// StateMust returns the state for [ServicebusQueueAuthorizationRule]. Panics if the state is nil.
func (sqar *ServicebusQueueAuthorizationRule) StateMust() *servicebusQueueAuthorizationRuleState {
	if sqar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sqar.Type(), sqar.LocalName()))
	}
	return sqar.state
}

// ServicebusQueueAuthorizationRuleArgs contains the configurations for azurerm_servicebus_queue_authorization_rule.
type ServicebusQueueAuthorizationRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Listen: bool, optional
	Listen terra.BoolValue `hcl:"listen,attr"`
	// Manage: bool, optional
	Manage terra.BoolValue `hcl:"manage,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// QueueId: string, required
	QueueId terra.StringValue `hcl:"queue_id,attr" validate:"required"`
	// Send: bool, optional
	Send terra.BoolValue `hcl:"send,attr"`
	// Timeouts: optional
	Timeouts *servicebusqueueauthorizationrule.Timeouts `hcl:"timeouts,block"`
}
type servicebusQueueAuthorizationRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_servicebus_queue_authorization_rule.
func (sqar servicebusQueueAuthorizationRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("id"))
}

// Listen returns a reference to field listen of azurerm_servicebus_queue_authorization_rule.
func (sqar servicebusQueueAuthorizationRuleAttributes) Listen() terra.BoolValue {
	return terra.ReferenceAsBool(sqar.ref.Append("listen"))
}

// Manage returns a reference to field manage of azurerm_servicebus_queue_authorization_rule.
func (sqar servicebusQueueAuthorizationRuleAttributes) Manage() terra.BoolValue {
	return terra.ReferenceAsBool(sqar.ref.Append("manage"))
}

// Name returns a reference to field name of azurerm_servicebus_queue_authorization_rule.
func (sqar servicebusQueueAuthorizationRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("name"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_servicebus_queue_authorization_rule.
func (sqar servicebusQueueAuthorizationRuleAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("primary_connection_string"))
}

// PrimaryConnectionStringAlias returns a reference to field primary_connection_string_alias of azurerm_servicebus_queue_authorization_rule.
func (sqar servicebusQueueAuthorizationRuleAttributes) PrimaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("primary_connection_string_alias"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_servicebus_queue_authorization_rule.
func (sqar servicebusQueueAuthorizationRuleAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("primary_key"))
}

// QueueId returns a reference to field queue_id of azurerm_servicebus_queue_authorization_rule.
func (sqar servicebusQueueAuthorizationRuleAttributes) QueueId() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("queue_id"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_servicebus_queue_authorization_rule.
func (sqar servicebusQueueAuthorizationRuleAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("secondary_connection_string"))
}

// SecondaryConnectionStringAlias returns a reference to field secondary_connection_string_alias of azurerm_servicebus_queue_authorization_rule.
func (sqar servicebusQueueAuthorizationRuleAttributes) SecondaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("secondary_connection_string_alias"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_servicebus_queue_authorization_rule.
func (sqar servicebusQueueAuthorizationRuleAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("secondary_key"))
}

// Send returns a reference to field send of azurerm_servicebus_queue_authorization_rule.
func (sqar servicebusQueueAuthorizationRuleAttributes) Send() terra.BoolValue {
	return terra.ReferenceAsBool(sqar.ref.Append("send"))
}

func (sqar servicebusQueueAuthorizationRuleAttributes) Timeouts() servicebusqueueauthorizationrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicebusqueueauthorizationrule.TimeoutsAttributes](sqar.ref.Append("timeouts"))
}

type servicebusQueueAuthorizationRuleState struct {
	Id                             string                                          `json:"id"`
	Listen                         bool                                            `json:"listen"`
	Manage                         bool                                            `json:"manage"`
	Name                           string                                          `json:"name"`
	PrimaryConnectionString        string                                          `json:"primary_connection_string"`
	PrimaryConnectionStringAlias   string                                          `json:"primary_connection_string_alias"`
	PrimaryKey                     string                                          `json:"primary_key"`
	QueueId                        string                                          `json:"queue_id"`
	SecondaryConnectionString      string                                          `json:"secondary_connection_string"`
	SecondaryConnectionStringAlias string                                          `json:"secondary_connection_string_alias"`
	SecondaryKey                   string                                          `json:"secondary_key"`
	Send                           bool                                            `json:"send"`
	Timeouts                       *servicebusqueueauthorizationrule.TimeoutsState `json:"timeouts"`
}
