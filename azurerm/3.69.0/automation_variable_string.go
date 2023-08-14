// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationvariablestring "github.com/golingon/terraproviders/azurerm/3.69.0/automationvariablestring"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutomationVariableString creates a new instance of [AutomationVariableString].
func NewAutomationVariableString(name string, args AutomationVariableStringArgs) *AutomationVariableString {
	return &AutomationVariableString{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationVariableString)(nil)

// AutomationVariableString represents the Terraform resource azurerm_automation_variable_string.
type AutomationVariableString struct {
	Name      string
	Args      AutomationVariableStringArgs
	state     *automationVariableStringState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationVariableString].
func (avs *AutomationVariableString) Type() string {
	return "azurerm_automation_variable_string"
}

// LocalName returns the local name for [AutomationVariableString].
func (avs *AutomationVariableString) LocalName() string {
	return avs.Name
}

// Configuration returns the configuration (args) for [AutomationVariableString].
func (avs *AutomationVariableString) Configuration() interface{} {
	return avs.Args
}

// DependOn is used for other resources to depend on [AutomationVariableString].
func (avs *AutomationVariableString) DependOn() terra.Reference {
	return terra.ReferenceResource(avs)
}

// Dependencies returns the list of resources [AutomationVariableString] depends_on.
func (avs *AutomationVariableString) Dependencies() terra.Dependencies {
	return avs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationVariableString].
func (avs *AutomationVariableString) LifecycleManagement() *terra.Lifecycle {
	return avs.Lifecycle
}

// Attributes returns the attributes for [AutomationVariableString].
func (avs *AutomationVariableString) Attributes() automationVariableStringAttributes {
	return automationVariableStringAttributes{ref: terra.ReferenceResource(avs)}
}

// ImportState imports the given attribute values into [AutomationVariableString]'s state.
func (avs *AutomationVariableString) ImportState(av io.Reader) error {
	avs.state = &automationVariableStringState{}
	if err := json.NewDecoder(av).Decode(avs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avs.Type(), avs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationVariableString] has state.
func (avs *AutomationVariableString) State() (*automationVariableStringState, bool) {
	return avs.state, avs.state != nil
}

// StateMust returns the state for [AutomationVariableString]. Panics if the state is nil.
func (avs *AutomationVariableString) StateMust() *automationVariableStringState {
	if avs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avs.Type(), avs.LocalName()))
	}
	return avs.state
}

// AutomationVariableStringArgs contains the configurations for azurerm_automation_variable_string.
type AutomationVariableStringArgs struct {
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
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
	// Timeouts: optional
	Timeouts *automationvariablestring.Timeouts `hcl:"timeouts,block"`
}
type automationVariableStringAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_variable_string.
func (avs automationVariableStringAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("automation_account_name"))
}

// Description returns a reference to field description of azurerm_automation_variable_string.
func (avs automationVariableStringAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("description"))
}

// Encrypted returns a reference to field encrypted of azurerm_automation_variable_string.
func (avs automationVariableStringAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(avs.ref.Append("encrypted"))
}

// Id returns a reference to field id of azurerm_automation_variable_string.
func (avs automationVariableStringAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_variable_string.
func (avs automationVariableStringAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_variable_string.
func (avs automationVariableStringAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("resource_group_name"))
}

// Value returns a reference to field value of azurerm_automation_variable_string.
func (avs automationVariableStringAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("value"))
}

func (avs automationVariableStringAttributes) Timeouts() automationvariablestring.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationvariablestring.TimeoutsAttributes](avs.ref.Append("timeouts"))
}

type automationVariableStringState struct {
	AutomationAccountName string                                  `json:"automation_account_name"`
	Description           string                                  `json:"description"`
	Encrypted             bool                                    `json:"encrypted"`
	Id                    string                                  `json:"id"`
	Name                  string                                  `json:"name"`
	ResourceGroupName     string                                  `json:"resource_group_name"`
	Value                 string                                  `json:"value"`
	Timeouts              *automationvariablestring.TimeoutsState `json:"timeouts"`
}
