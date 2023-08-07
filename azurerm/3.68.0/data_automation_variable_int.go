// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataautomationvariableint "github.com/golingon/terraproviders/azurerm/3.68.0/dataautomationvariableint"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAutomationVariableInt creates a new instance of [DataAutomationVariableInt].
func NewDataAutomationVariableInt(name string, args DataAutomationVariableIntArgs) *DataAutomationVariableInt {
	return &DataAutomationVariableInt{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAutomationVariableInt)(nil)

// DataAutomationVariableInt represents the Terraform data resource azurerm_automation_variable_int.
type DataAutomationVariableInt struct {
	Name string
	Args DataAutomationVariableIntArgs
}

// DataSource returns the Terraform object type for [DataAutomationVariableInt].
func (avi *DataAutomationVariableInt) DataSource() string {
	return "azurerm_automation_variable_int"
}

// LocalName returns the local name for [DataAutomationVariableInt].
func (avi *DataAutomationVariableInt) LocalName() string {
	return avi.Name
}

// Configuration returns the configuration (args) for [DataAutomationVariableInt].
func (avi *DataAutomationVariableInt) Configuration() interface{} {
	return avi.Args
}

// Attributes returns the attributes for [DataAutomationVariableInt].
func (avi *DataAutomationVariableInt) Attributes() dataAutomationVariableIntAttributes {
	return dataAutomationVariableIntAttributes{ref: terra.ReferenceDataResource(avi)}
}

// DataAutomationVariableIntArgs contains the configurations for azurerm_automation_variable_int.
type DataAutomationVariableIntArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataautomationvariableint.Timeouts `hcl:"timeouts,block"`
}
type dataAutomationVariableIntAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_variable_int.
func (avi dataAutomationVariableIntAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(avi.ref.Append("automation_account_name"))
}

// Description returns a reference to field description of azurerm_automation_variable_int.
func (avi dataAutomationVariableIntAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(avi.ref.Append("description"))
}

// Encrypted returns a reference to field encrypted of azurerm_automation_variable_int.
func (avi dataAutomationVariableIntAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(avi.ref.Append("encrypted"))
}

// Id returns a reference to field id of azurerm_automation_variable_int.
func (avi dataAutomationVariableIntAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avi.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_variable_int.
func (avi dataAutomationVariableIntAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avi.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_variable_int.
func (avi dataAutomationVariableIntAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(avi.ref.Append("resource_group_name"))
}

// Value returns a reference to field value of azurerm_automation_variable_int.
func (avi dataAutomationVariableIntAttributes) Value() terra.NumberValue {
	return terra.ReferenceAsNumber(avi.ref.Append("value"))
}

func (avi dataAutomationVariableIntAttributes) Timeouts() dataautomationvariableint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataautomationvariableint.TimeoutsAttributes](avi.ref.Append("timeouts"))
}
