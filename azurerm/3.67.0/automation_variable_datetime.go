// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationvariabledatetime "github.com/golingon/terraproviders/azurerm/3.67.0/automationvariabledatetime"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutomationVariableDatetime creates a new instance of [AutomationVariableDatetime].
func NewAutomationVariableDatetime(name string, args AutomationVariableDatetimeArgs) *AutomationVariableDatetime {
	return &AutomationVariableDatetime{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationVariableDatetime)(nil)

// AutomationVariableDatetime represents the Terraform resource azurerm_automation_variable_datetime.
type AutomationVariableDatetime struct {
	Name      string
	Args      AutomationVariableDatetimeArgs
	state     *automationVariableDatetimeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationVariableDatetime].
func (avd *AutomationVariableDatetime) Type() string {
	return "azurerm_automation_variable_datetime"
}

// LocalName returns the local name for [AutomationVariableDatetime].
func (avd *AutomationVariableDatetime) LocalName() string {
	return avd.Name
}

// Configuration returns the configuration (args) for [AutomationVariableDatetime].
func (avd *AutomationVariableDatetime) Configuration() interface{} {
	return avd.Args
}

// DependOn is used for other resources to depend on [AutomationVariableDatetime].
func (avd *AutomationVariableDatetime) DependOn() terra.Reference {
	return terra.ReferenceResource(avd)
}

// Dependencies returns the list of resources [AutomationVariableDatetime] depends_on.
func (avd *AutomationVariableDatetime) Dependencies() terra.Dependencies {
	return avd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationVariableDatetime].
func (avd *AutomationVariableDatetime) LifecycleManagement() *terra.Lifecycle {
	return avd.Lifecycle
}

// Attributes returns the attributes for [AutomationVariableDatetime].
func (avd *AutomationVariableDatetime) Attributes() automationVariableDatetimeAttributes {
	return automationVariableDatetimeAttributes{ref: terra.ReferenceResource(avd)}
}

// ImportState imports the given attribute values into [AutomationVariableDatetime]'s state.
func (avd *AutomationVariableDatetime) ImportState(av io.Reader) error {
	avd.state = &automationVariableDatetimeState{}
	if err := json.NewDecoder(av).Decode(avd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avd.Type(), avd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationVariableDatetime] has state.
func (avd *AutomationVariableDatetime) State() (*automationVariableDatetimeState, bool) {
	return avd.state, avd.state != nil
}

// StateMust returns the state for [AutomationVariableDatetime]. Panics if the state is nil.
func (avd *AutomationVariableDatetime) StateMust() *automationVariableDatetimeState {
	if avd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avd.Type(), avd.LocalName()))
	}
	return avd.state
}

// AutomationVariableDatetimeArgs contains the configurations for azurerm_automation_variable_datetime.
type AutomationVariableDatetimeArgs struct {
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
	Timeouts *automationvariabledatetime.Timeouts `hcl:"timeouts,block"`
}
type automationVariableDatetimeAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_variable_datetime.
func (avd automationVariableDatetimeAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(avd.ref.Append("automation_account_name"))
}

// Description returns a reference to field description of azurerm_automation_variable_datetime.
func (avd automationVariableDatetimeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(avd.ref.Append("description"))
}

// Encrypted returns a reference to field encrypted of azurerm_automation_variable_datetime.
func (avd automationVariableDatetimeAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(avd.ref.Append("encrypted"))
}

// Id returns a reference to field id of azurerm_automation_variable_datetime.
func (avd automationVariableDatetimeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_variable_datetime.
func (avd automationVariableDatetimeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avd.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_variable_datetime.
func (avd automationVariableDatetimeAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(avd.ref.Append("resource_group_name"))
}

// Value returns a reference to field value of azurerm_automation_variable_datetime.
func (avd automationVariableDatetimeAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(avd.ref.Append("value"))
}

func (avd automationVariableDatetimeAttributes) Timeouts() automationvariabledatetime.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationvariabledatetime.TimeoutsAttributes](avd.ref.Append("timeouts"))
}

type automationVariableDatetimeState struct {
	AutomationAccountName string                                    `json:"automation_account_name"`
	Description           string                                    `json:"description"`
	Encrypted             bool                                      `json:"encrypted"`
	Id                    string                                    `json:"id"`
	Name                  string                                    `json:"name"`
	ResourceGroupName     string                                    `json:"resource_group_name"`
	Value                 string                                    `json:"value"`
	Timeouts              *automationvariabledatetime.TimeoutsState `json:"timeouts"`
}
