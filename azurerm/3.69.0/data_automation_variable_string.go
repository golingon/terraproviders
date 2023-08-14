// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataautomationvariablestring "github.com/golingon/terraproviders/azurerm/3.69.0/dataautomationvariablestring"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAutomationVariableString creates a new instance of [DataAutomationVariableString].
func NewDataAutomationVariableString(name string, args DataAutomationVariableStringArgs) *DataAutomationVariableString {
	return &DataAutomationVariableString{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAutomationVariableString)(nil)

// DataAutomationVariableString represents the Terraform data resource azurerm_automation_variable_string.
type DataAutomationVariableString struct {
	Name string
	Args DataAutomationVariableStringArgs
}

// DataSource returns the Terraform object type for [DataAutomationVariableString].
func (avs *DataAutomationVariableString) DataSource() string {
	return "azurerm_automation_variable_string"
}

// LocalName returns the local name for [DataAutomationVariableString].
func (avs *DataAutomationVariableString) LocalName() string {
	return avs.Name
}

// Configuration returns the configuration (args) for [DataAutomationVariableString].
func (avs *DataAutomationVariableString) Configuration() interface{} {
	return avs.Args
}

// Attributes returns the attributes for [DataAutomationVariableString].
func (avs *DataAutomationVariableString) Attributes() dataAutomationVariableStringAttributes {
	return dataAutomationVariableStringAttributes{ref: terra.ReferenceDataResource(avs)}
}

// DataAutomationVariableStringArgs contains the configurations for azurerm_automation_variable_string.
type DataAutomationVariableStringArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataautomationvariablestring.Timeouts `hcl:"timeouts,block"`
}
type dataAutomationVariableStringAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_variable_string.
func (avs dataAutomationVariableStringAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("automation_account_name"))
}

// Description returns a reference to field description of azurerm_automation_variable_string.
func (avs dataAutomationVariableStringAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("description"))
}

// Encrypted returns a reference to field encrypted of azurerm_automation_variable_string.
func (avs dataAutomationVariableStringAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(avs.ref.Append("encrypted"))
}

// Id returns a reference to field id of azurerm_automation_variable_string.
func (avs dataAutomationVariableStringAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_variable_string.
func (avs dataAutomationVariableStringAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_variable_string.
func (avs dataAutomationVariableStringAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("resource_group_name"))
}

// Value returns a reference to field value of azurerm_automation_variable_string.
func (avs dataAutomationVariableStringAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("value"))
}

func (avs dataAutomationVariableStringAttributes) Timeouts() dataautomationvariablestring.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataautomationvariablestring.TimeoutsAttributes](avs.ref.Append("timeouts"))
}
