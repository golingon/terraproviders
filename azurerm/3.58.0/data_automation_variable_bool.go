// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataautomationvariablebool "github.com/golingon/terraproviders/azurerm/3.58.0/dataautomationvariablebool"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAutomationVariableBool creates a new instance of [DataAutomationVariableBool].
func NewDataAutomationVariableBool(name string, args DataAutomationVariableBoolArgs) *DataAutomationVariableBool {
	return &DataAutomationVariableBool{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAutomationVariableBool)(nil)

// DataAutomationVariableBool represents the Terraform data resource azurerm_automation_variable_bool.
type DataAutomationVariableBool struct {
	Name string
	Args DataAutomationVariableBoolArgs
}

// DataSource returns the Terraform object type for [DataAutomationVariableBool].
func (avb *DataAutomationVariableBool) DataSource() string {
	return "azurerm_automation_variable_bool"
}

// LocalName returns the local name for [DataAutomationVariableBool].
func (avb *DataAutomationVariableBool) LocalName() string {
	return avb.Name
}

// Configuration returns the configuration (args) for [DataAutomationVariableBool].
func (avb *DataAutomationVariableBool) Configuration() interface{} {
	return avb.Args
}

// Attributes returns the attributes for [DataAutomationVariableBool].
func (avb *DataAutomationVariableBool) Attributes() dataAutomationVariableBoolAttributes {
	return dataAutomationVariableBoolAttributes{ref: terra.ReferenceDataResource(avb)}
}

// DataAutomationVariableBoolArgs contains the configurations for azurerm_automation_variable_bool.
type DataAutomationVariableBoolArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataautomationvariablebool.Timeouts `hcl:"timeouts,block"`
}
type dataAutomationVariableBoolAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_variable_bool.
func (avb dataAutomationVariableBoolAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(avb.ref.Append("automation_account_name"))
}

// Description returns a reference to field description of azurerm_automation_variable_bool.
func (avb dataAutomationVariableBoolAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(avb.ref.Append("description"))
}

// Encrypted returns a reference to field encrypted of azurerm_automation_variable_bool.
func (avb dataAutomationVariableBoolAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(avb.ref.Append("encrypted"))
}

// Id returns a reference to field id of azurerm_automation_variable_bool.
func (avb dataAutomationVariableBoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avb.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_variable_bool.
func (avb dataAutomationVariableBoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avb.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_variable_bool.
func (avb dataAutomationVariableBoolAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(avb.ref.Append("resource_group_name"))
}

// Value returns a reference to field value of azurerm_automation_variable_bool.
func (avb dataAutomationVariableBoolAttributes) Value() terra.BoolValue {
	return terra.ReferenceAsBool(avb.ref.Append("value"))
}

func (avb dataAutomationVariableBoolAttributes) Timeouts() dataautomationvariablebool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataautomationvariablebool.TimeoutsAttributes](avb.ref.Append("timeouts"))
}
