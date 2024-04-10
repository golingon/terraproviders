// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	dataautomationvariabledatetime "github.com/golingon/terraproviders/azurerm/3.98.0/dataautomationvariabledatetime"
)

// NewDataAutomationVariableDatetime creates a new instance of [DataAutomationVariableDatetime].
func NewDataAutomationVariableDatetime(name string, args DataAutomationVariableDatetimeArgs) *DataAutomationVariableDatetime {
	return &DataAutomationVariableDatetime{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAutomationVariableDatetime)(nil)

// DataAutomationVariableDatetime represents the Terraform data resource azurerm_automation_variable_datetime.
type DataAutomationVariableDatetime struct {
	Name string
	Args DataAutomationVariableDatetimeArgs
}

// DataSource returns the Terraform object type for [DataAutomationVariableDatetime].
func (avd *DataAutomationVariableDatetime) DataSource() string {
	return "azurerm_automation_variable_datetime"
}

// LocalName returns the local name for [DataAutomationVariableDatetime].
func (avd *DataAutomationVariableDatetime) LocalName() string {
	return avd.Name
}

// Configuration returns the configuration (args) for [DataAutomationVariableDatetime].
func (avd *DataAutomationVariableDatetime) Configuration() interface{} {
	return avd.Args
}

// Attributes returns the attributes for [DataAutomationVariableDatetime].
func (avd *DataAutomationVariableDatetime) Attributes() dataAutomationVariableDatetimeAttributes {
	return dataAutomationVariableDatetimeAttributes{ref: terra.ReferenceDataResource(avd)}
}

// DataAutomationVariableDatetimeArgs contains the configurations for azurerm_automation_variable_datetime.
type DataAutomationVariableDatetimeArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataautomationvariabledatetime.Timeouts `hcl:"timeouts,block"`
}
type dataAutomationVariableDatetimeAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_variable_datetime.
func (avd dataAutomationVariableDatetimeAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(avd.ref.Append("automation_account_name"))
}

// Description returns a reference to field description of azurerm_automation_variable_datetime.
func (avd dataAutomationVariableDatetimeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(avd.ref.Append("description"))
}

// Encrypted returns a reference to field encrypted of azurerm_automation_variable_datetime.
func (avd dataAutomationVariableDatetimeAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(avd.ref.Append("encrypted"))
}

// Id returns a reference to field id of azurerm_automation_variable_datetime.
func (avd dataAutomationVariableDatetimeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_variable_datetime.
func (avd dataAutomationVariableDatetimeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avd.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_variable_datetime.
func (avd dataAutomationVariableDatetimeAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(avd.ref.Append("resource_group_name"))
}

// Value returns a reference to field value of azurerm_automation_variable_datetime.
func (avd dataAutomationVariableDatetimeAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(avd.ref.Append("value"))
}

func (avd dataAutomationVariableDatetimeAttributes) Timeouts() dataautomationvariabledatetime.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataautomationvariabledatetime.TimeoutsAttributes](avd.ref.Append("timeouts"))
}
