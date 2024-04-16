// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_billing_mpa_account_scope

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_billing_mpa_account_scope.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (abmas *DataSource) DataSource() string {
	return "azurerm_billing_mpa_account_scope"
}

// LocalName returns the local name for [DataSource].
func (abmas *DataSource) LocalName() string {
	return abmas.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (abmas *DataSource) Configuration() interface{} {
	return abmas.Args
}

// Attributes returns the attributes for [DataSource].
func (abmas *DataSource) Attributes() dataAzurermBillingMpaAccountScopeAttributes {
	return dataAzurermBillingMpaAccountScopeAttributes{ref: terra.ReferenceDataSource(abmas)}
}

// DataArgs contains the configurations for azurerm_billing_mpa_account_scope.
type DataArgs struct {
	// BillingAccountName: string, required
	BillingAccountName terra.StringValue `hcl:"billing_account_name,attr" validate:"required"`
	// CustomerName: string, required
	CustomerName terra.StringValue `hcl:"customer_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermBillingMpaAccountScopeAttributes struct {
	ref terra.Reference
}

// BillingAccountName returns a reference to field billing_account_name of azurerm_billing_mpa_account_scope.
func (abmas dataAzurermBillingMpaAccountScopeAttributes) BillingAccountName() terra.StringValue {
	return terra.ReferenceAsString(abmas.ref.Append("billing_account_name"))
}

// CustomerName returns a reference to field customer_name of azurerm_billing_mpa_account_scope.
func (abmas dataAzurermBillingMpaAccountScopeAttributes) CustomerName() terra.StringValue {
	return terra.ReferenceAsString(abmas.ref.Append("customer_name"))
}

// Id returns a reference to field id of azurerm_billing_mpa_account_scope.
func (abmas dataAzurermBillingMpaAccountScopeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(abmas.ref.Append("id"))
}

func (abmas dataAzurermBillingMpaAccountScopeAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](abmas.ref.Append("timeouts"))
}
