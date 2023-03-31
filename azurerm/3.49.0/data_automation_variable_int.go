// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataautomationvariableint "github.com/golingon/terraproviders/azurerm/3.49.0/dataautomationvariableint"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataAutomationVariableInt(name string, args DataAutomationVariableIntArgs) *DataAutomationVariableInt {
	return &DataAutomationVariableInt{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAutomationVariableInt)(nil)

type DataAutomationVariableInt struct {
	Name string
	Args DataAutomationVariableIntArgs
}

func (avi *DataAutomationVariableInt) DataSource() string {
	return "azurerm_automation_variable_int"
}

func (avi *DataAutomationVariableInt) LocalName() string {
	return avi.Name
}

func (avi *DataAutomationVariableInt) Configuration() interface{} {
	return avi.Args
}

func (avi *DataAutomationVariableInt) Attributes() dataAutomationVariableIntAttributes {
	return dataAutomationVariableIntAttributes{ref: terra.ReferenceDataResource(avi)}
}

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

func (avi dataAutomationVariableIntAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceString(avi.ref.Append("automation_account_name"))
}

func (avi dataAutomationVariableIntAttributes) Description() terra.StringValue {
	return terra.ReferenceString(avi.ref.Append("description"))
}

func (avi dataAutomationVariableIntAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceBool(avi.ref.Append("encrypted"))
}

func (avi dataAutomationVariableIntAttributes) Id() terra.StringValue {
	return terra.ReferenceString(avi.ref.Append("id"))
}

func (avi dataAutomationVariableIntAttributes) Name() terra.StringValue {
	return terra.ReferenceString(avi.ref.Append("name"))
}

func (avi dataAutomationVariableIntAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(avi.ref.Append("resource_group_name"))
}

func (avi dataAutomationVariableIntAttributes) Value() terra.NumberValue {
	return terra.ReferenceNumber(avi.ref.Append("value"))
}

func (avi dataAutomationVariableIntAttributes) Timeouts() dataautomationvariableint.TimeoutsAttributes {
	return terra.ReferenceSingle[dataautomationvariableint.TimeoutsAttributes](avi.ref.Append("timeouts"))
}
