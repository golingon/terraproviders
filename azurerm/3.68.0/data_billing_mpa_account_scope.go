// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	databillingmpaaccountscope "github.com/golingon/terraproviders/azurerm/3.68.0/databillingmpaaccountscope"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataBillingMpaAccountScope creates a new instance of [DataBillingMpaAccountScope].
func NewDataBillingMpaAccountScope(name string, args DataBillingMpaAccountScopeArgs) *DataBillingMpaAccountScope {
	return &DataBillingMpaAccountScope{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBillingMpaAccountScope)(nil)

// DataBillingMpaAccountScope represents the Terraform data resource azurerm_billing_mpa_account_scope.
type DataBillingMpaAccountScope struct {
	Name string
	Args DataBillingMpaAccountScopeArgs
}

// DataSource returns the Terraform object type for [DataBillingMpaAccountScope].
func (bmas *DataBillingMpaAccountScope) DataSource() string {
	return "azurerm_billing_mpa_account_scope"
}

// LocalName returns the local name for [DataBillingMpaAccountScope].
func (bmas *DataBillingMpaAccountScope) LocalName() string {
	return bmas.Name
}

// Configuration returns the configuration (args) for [DataBillingMpaAccountScope].
func (bmas *DataBillingMpaAccountScope) Configuration() interface{} {
	return bmas.Args
}

// Attributes returns the attributes for [DataBillingMpaAccountScope].
func (bmas *DataBillingMpaAccountScope) Attributes() dataBillingMpaAccountScopeAttributes {
	return dataBillingMpaAccountScopeAttributes{ref: terra.ReferenceDataResource(bmas)}
}

// DataBillingMpaAccountScopeArgs contains the configurations for azurerm_billing_mpa_account_scope.
type DataBillingMpaAccountScopeArgs struct {
	// BillingAccountName: string, required
	BillingAccountName terra.StringValue `hcl:"billing_account_name,attr" validate:"required"`
	// CustomerName: string, required
	CustomerName terra.StringValue `hcl:"customer_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *databillingmpaaccountscope.Timeouts `hcl:"timeouts,block"`
}
type dataBillingMpaAccountScopeAttributes struct {
	ref terra.Reference
}

// BillingAccountName returns a reference to field billing_account_name of azurerm_billing_mpa_account_scope.
func (bmas dataBillingMpaAccountScopeAttributes) BillingAccountName() terra.StringValue {
	return terra.ReferenceAsString(bmas.ref.Append("billing_account_name"))
}

// CustomerName returns a reference to field customer_name of azurerm_billing_mpa_account_scope.
func (bmas dataBillingMpaAccountScopeAttributes) CustomerName() terra.StringValue {
	return terra.ReferenceAsString(bmas.ref.Append("customer_name"))
}

// Id returns a reference to field id of azurerm_billing_mpa_account_scope.
func (bmas dataBillingMpaAccountScopeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bmas.ref.Append("id"))
}

func (bmas dataBillingMpaAccountScopeAttributes) Timeouts() databillingmpaaccountscope.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databillingmpaaccountscope.TimeoutsAttributes](bmas.ref.Append("timeouts"))
}
