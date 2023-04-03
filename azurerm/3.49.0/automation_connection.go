// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationconnection "github.com/golingon/terraproviders/azurerm/3.49.0/automationconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutomationConnection creates a new instance of [AutomationConnection].
func NewAutomationConnection(name string, args AutomationConnectionArgs) *AutomationConnection {
	return &AutomationConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationConnection)(nil)

// AutomationConnection represents the Terraform resource azurerm_automation_connection.
type AutomationConnection struct {
	Name      string
	Args      AutomationConnectionArgs
	state     *automationConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationConnection].
func (ac *AutomationConnection) Type() string {
	return "azurerm_automation_connection"
}

// LocalName returns the local name for [AutomationConnection].
func (ac *AutomationConnection) LocalName() string {
	return ac.Name
}

// Configuration returns the configuration (args) for [AutomationConnection].
func (ac *AutomationConnection) Configuration() interface{} {
	return ac.Args
}

// DependOn is used for other resources to depend on [AutomationConnection].
func (ac *AutomationConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(ac)
}

// Dependencies returns the list of resources [AutomationConnection] depends_on.
func (ac *AutomationConnection) Dependencies() terra.Dependencies {
	return ac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationConnection].
func (ac *AutomationConnection) LifecycleManagement() *terra.Lifecycle {
	return ac.Lifecycle
}

// Attributes returns the attributes for [AutomationConnection].
func (ac *AutomationConnection) Attributes() automationConnectionAttributes {
	return automationConnectionAttributes{ref: terra.ReferenceResource(ac)}
}

// ImportState imports the given attribute values into [AutomationConnection]'s state.
func (ac *AutomationConnection) ImportState(av io.Reader) error {
	ac.state = &automationConnectionState{}
	if err := json.NewDecoder(av).Decode(ac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ac.Type(), ac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationConnection] has state.
func (ac *AutomationConnection) State() (*automationConnectionState, bool) {
	return ac.state, ac.state != nil
}

// StateMust returns the state for [AutomationConnection]. Panics if the state is nil.
func (ac *AutomationConnection) StateMust() *automationConnectionState {
	if ac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ac.Type(), ac.LocalName()))
	}
	return ac.state
}

// AutomationConnectionArgs contains the configurations for azurerm_automation_connection.
type AutomationConnectionArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Values: map of string, required
	Values terra.MapValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *automationconnection.Timeouts `hcl:"timeouts,block"`
}
type automationConnectionAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_connection.
func (ac automationConnectionAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("automation_account_name"))
}

// Description returns a reference to field description of azurerm_automation_connection.
func (ac automationConnectionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_automation_connection.
func (ac automationConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_connection.
func (ac automationConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_connection.
func (ac automationConnectionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("resource_group_name"))
}

// Type returns a reference to field type of azurerm_automation_connection.
func (ac automationConnectionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("type"))
}

// Values returns a reference to field values of azurerm_automation_connection.
func (ac automationConnectionAttributes) Values() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ac.ref.Append("values"))
}

func (ac automationConnectionAttributes) Timeouts() automationconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationconnection.TimeoutsAttributes](ac.ref.Append("timeouts"))
}

type automationConnectionState struct {
	AutomationAccountName string                              `json:"automation_account_name"`
	Description           string                              `json:"description"`
	Id                    string                              `json:"id"`
	Name                  string                              `json:"name"`
	ResourceGroupName     string                              `json:"resource_group_name"`
	Type                  string                              `json:"type"`
	Values                map[string]string                   `json:"values"`
	Timeouts              *automationconnection.TimeoutsState `json:"timeouts"`
}
