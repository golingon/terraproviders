// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationconnectiontype "github.com/golingon/terraproviders/azurerm/3.68.0/automationconnectiontype"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutomationConnectionType creates a new instance of [AutomationConnectionType].
func NewAutomationConnectionType(name string, args AutomationConnectionTypeArgs) *AutomationConnectionType {
	return &AutomationConnectionType{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationConnectionType)(nil)

// AutomationConnectionType represents the Terraform resource azurerm_automation_connection_type.
type AutomationConnectionType struct {
	Name      string
	Args      AutomationConnectionTypeArgs
	state     *automationConnectionTypeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationConnectionType].
func (act *AutomationConnectionType) Type() string {
	return "azurerm_automation_connection_type"
}

// LocalName returns the local name for [AutomationConnectionType].
func (act *AutomationConnectionType) LocalName() string {
	return act.Name
}

// Configuration returns the configuration (args) for [AutomationConnectionType].
func (act *AutomationConnectionType) Configuration() interface{} {
	return act.Args
}

// DependOn is used for other resources to depend on [AutomationConnectionType].
func (act *AutomationConnectionType) DependOn() terra.Reference {
	return terra.ReferenceResource(act)
}

// Dependencies returns the list of resources [AutomationConnectionType] depends_on.
func (act *AutomationConnectionType) Dependencies() terra.Dependencies {
	return act.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationConnectionType].
func (act *AutomationConnectionType) LifecycleManagement() *terra.Lifecycle {
	return act.Lifecycle
}

// Attributes returns the attributes for [AutomationConnectionType].
func (act *AutomationConnectionType) Attributes() automationConnectionTypeAttributes {
	return automationConnectionTypeAttributes{ref: terra.ReferenceResource(act)}
}

// ImportState imports the given attribute values into [AutomationConnectionType]'s state.
func (act *AutomationConnectionType) ImportState(av io.Reader) error {
	act.state = &automationConnectionTypeState{}
	if err := json.NewDecoder(av).Decode(act.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", act.Type(), act.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationConnectionType] has state.
func (act *AutomationConnectionType) State() (*automationConnectionTypeState, bool) {
	return act.state, act.state != nil
}

// StateMust returns the state for [AutomationConnectionType]. Panics if the state is nil.
func (act *AutomationConnectionType) StateMust() *automationConnectionTypeState {
	if act.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", act.Type(), act.LocalName()))
	}
	return act.state
}

// AutomationConnectionTypeArgs contains the configurations for azurerm_automation_connection_type.
type AutomationConnectionTypeArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsGlobal: bool, optional
	IsGlobal terra.BoolValue `hcl:"is_global,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Field: min=1
	Field []automationconnectiontype.Field `hcl:"field,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *automationconnectiontype.Timeouts `hcl:"timeouts,block"`
}
type automationConnectionTypeAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_connection_type.
func (act automationConnectionTypeAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(act.ref.Append("automation_account_name"))
}

// Id returns a reference to field id of azurerm_automation_connection_type.
func (act automationConnectionTypeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(act.ref.Append("id"))
}

// IsGlobal returns a reference to field is_global of azurerm_automation_connection_type.
func (act automationConnectionTypeAttributes) IsGlobal() terra.BoolValue {
	return terra.ReferenceAsBool(act.ref.Append("is_global"))
}

// Name returns a reference to field name of azurerm_automation_connection_type.
func (act automationConnectionTypeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(act.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_connection_type.
func (act automationConnectionTypeAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(act.ref.Append("resource_group_name"))
}

func (act automationConnectionTypeAttributes) Field() terra.ListValue[automationconnectiontype.FieldAttributes] {
	return terra.ReferenceAsList[automationconnectiontype.FieldAttributes](act.ref.Append("field"))
}

func (act automationConnectionTypeAttributes) Timeouts() automationconnectiontype.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationconnectiontype.TimeoutsAttributes](act.ref.Append("timeouts"))
}

type automationConnectionTypeState struct {
	AutomationAccountName string                                  `json:"automation_account_name"`
	Id                    string                                  `json:"id"`
	IsGlobal              bool                                    `json:"is_global"`
	Name                  string                                  `json:"name"`
	ResourceGroupName     string                                  `json:"resource_group_name"`
	Field                 []automationconnectiontype.FieldState   `json:"field"`
	Timeouts              *automationconnectiontype.TimeoutsState `json:"timeouts"`
}
