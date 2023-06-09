// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationvariablebool "github.com/golingon/terraproviders/azurerm/3.63.0/automationvariablebool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutomationVariableBool creates a new instance of [AutomationVariableBool].
func NewAutomationVariableBool(name string, args AutomationVariableBoolArgs) *AutomationVariableBool {
	return &AutomationVariableBool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationVariableBool)(nil)

// AutomationVariableBool represents the Terraform resource azurerm_automation_variable_bool.
type AutomationVariableBool struct {
	Name      string
	Args      AutomationVariableBoolArgs
	state     *automationVariableBoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationVariableBool].
func (avb *AutomationVariableBool) Type() string {
	return "azurerm_automation_variable_bool"
}

// LocalName returns the local name for [AutomationVariableBool].
func (avb *AutomationVariableBool) LocalName() string {
	return avb.Name
}

// Configuration returns the configuration (args) for [AutomationVariableBool].
func (avb *AutomationVariableBool) Configuration() interface{} {
	return avb.Args
}

// DependOn is used for other resources to depend on [AutomationVariableBool].
func (avb *AutomationVariableBool) DependOn() terra.Reference {
	return terra.ReferenceResource(avb)
}

// Dependencies returns the list of resources [AutomationVariableBool] depends_on.
func (avb *AutomationVariableBool) Dependencies() terra.Dependencies {
	return avb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationVariableBool].
func (avb *AutomationVariableBool) LifecycleManagement() *terra.Lifecycle {
	return avb.Lifecycle
}

// Attributes returns the attributes for [AutomationVariableBool].
func (avb *AutomationVariableBool) Attributes() automationVariableBoolAttributes {
	return automationVariableBoolAttributes{ref: terra.ReferenceResource(avb)}
}

// ImportState imports the given attribute values into [AutomationVariableBool]'s state.
func (avb *AutomationVariableBool) ImportState(av io.Reader) error {
	avb.state = &automationVariableBoolState{}
	if err := json.NewDecoder(av).Decode(avb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avb.Type(), avb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationVariableBool] has state.
func (avb *AutomationVariableBool) State() (*automationVariableBoolState, bool) {
	return avb.state, avb.state != nil
}

// StateMust returns the state for [AutomationVariableBool]. Panics if the state is nil.
func (avb *AutomationVariableBool) StateMust() *automationVariableBoolState {
	if avb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avb.Type(), avb.LocalName()))
	}
	return avb.state
}

// AutomationVariableBoolArgs contains the configurations for azurerm_automation_variable_bool.
type AutomationVariableBoolArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Encrypted: bool, optional
	Encrypted terra.BoolValue `hcl:"encrypted,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Value: bool, optional
	Value terra.BoolValue `hcl:"value,attr"`
	// Timeouts: optional
	Timeouts *automationvariablebool.Timeouts `hcl:"timeouts,block"`
}
type automationVariableBoolAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_variable_bool.
func (avb automationVariableBoolAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(avb.ref.Append("automation_account_name"))
}

// Description returns a reference to field description of azurerm_automation_variable_bool.
func (avb automationVariableBoolAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(avb.ref.Append("description"))
}

// Encrypted returns a reference to field encrypted of azurerm_automation_variable_bool.
func (avb automationVariableBoolAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(avb.ref.Append("encrypted"))
}

// Id returns a reference to field id of azurerm_automation_variable_bool.
func (avb automationVariableBoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avb.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_variable_bool.
func (avb automationVariableBoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avb.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_variable_bool.
func (avb automationVariableBoolAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(avb.ref.Append("resource_group_name"))
}

// Value returns a reference to field value of azurerm_automation_variable_bool.
func (avb automationVariableBoolAttributes) Value() terra.BoolValue {
	return terra.ReferenceAsBool(avb.ref.Append("value"))
}

func (avb automationVariableBoolAttributes) Timeouts() automationvariablebool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationvariablebool.TimeoutsAttributes](avb.ref.Append("timeouts"))
}

type automationVariableBoolState struct {
	AutomationAccountName string                                `json:"automation_account_name"`
	Description           string                                `json:"description"`
	Encrypted             bool                                  `json:"encrypted"`
	Id                    string                                `json:"id"`
	Name                  string                                `json:"name"`
	ResourceGroupName     string                                `json:"resource_group_name"`
	Value                 bool                                  `json:"value"`
	Timeouts              *automationvariablebool.TimeoutsState `json:"timeouts"`
}
