// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataautomationaccount "github.com/golingon/terraproviders/azurerm/3.49.0/dataautomationaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAutomationAccount creates a new instance of [DataAutomationAccount].
func NewDataAutomationAccount(name string, args DataAutomationAccountArgs) *DataAutomationAccount {
	return &DataAutomationAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAutomationAccount)(nil)

// DataAutomationAccount represents the Terraform data resource azurerm_automation_account.
type DataAutomationAccount struct {
	Name string
	Args DataAutomationAccountArgs
}

// DataSource returns the Terraform object type for [DataAutomationAccount].
func (aa *DataAutomationAccount) DataSource() string {
	return "azurerm_automation_account"
}

// LocalName returns the local name for [DataAutomationAccount].
func (aa *DataAutomationAccount) LocalName() string {
	return aa.Name
}

// Configuration returns the configuration (args) for [DataAutomationAccount].
func (aa *DataAutomationAccount) Configuration() interface{} {
	return aa.Args
}

// Attributes returns the attributes for [DataAutomationAccount].
func (aa *DataAutomationAccount) Attributes() dataAutomationAccountAttributes {
	return dataAutomationAccountAttributes{ref: terra.ReferenceDataResource(aa)}
}

// DataAutomationAccountArgs contains the configurations for azurerm_automation_account.
type DataAutomationAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []dataautomationaccount.Identity `hcl:"identity,block" validate:"min=0"`
	// PrivateEndpointConnection: min=0
	PrivateEndpointConnection []dataautomationaccount.PrivateEndpointConnection `hcl:"private_endpoint_connection,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataautomationaccount.Timeouts `hcl:"timeouts,block"`
}
type dataAutomationAccountAttributes struct {
	ref terra.Reference
}

// Endpoint returns a reference to field endpoint of azurerm_automation_account.
func (aa dataAutomationAccountAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("endpoint"))
}

// HybridServiceUrl returns a reference to field hybrid_service_url of azurerm_automation_account.
func (aa dataAutomationAccountAttributes) HybridServiceUrl() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("hybrid_service_url"))
}

// Id returns a reference to field id of azurerm_automation_account.
func (aa dataAutomationAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_automation_account.
func (aa dataAutomationAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("name"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_automation_account.
func (aa dataAutomationAccountAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("primary_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_account.
func (aa dataAutomationAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("resource_group_name"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_automation_account.
func (aa dataAutomationAccountAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("secondary_key"))
}

func (aa dataAutomationAccountAttributes) Identity() terra.ListValue[dataautomationaccount.IdentityAttributes] {
	return terra.ReferenceAsList[dataautomationaccount.IdentityAttributes](aa.ref.Append("identity"))
}

func (aa dataAutomationAccountAttributes) PrivateEndpointConnection() terra.ListValue[dataautomationaccount.PrivateEndpointConnectionAttributes] {
	return terra.ReferenceAsList[dataautomationaccount.PrivateEndpointConnectionAttributes](aa.ref.Append("private_endpoint_connection"))
}

func (aa dataAutomationAccountAttributes) Timeouts() dataautomationaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataautomationaccount.TimeoutsAttributes](aa.ref.Append("timeouts"))
}
