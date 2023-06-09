// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	databillingenrollmentaccountscope "github.com/golingon/terraproviders/azurerm/3.55.0/databillingenrollmentaccountscope"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataBillingEnrollmentAccountScope creates a new instance of [DataBillingEnrollmentAccountScope].
func NewDataBillingEnrollmentAccountScope(name string, args DataBillingEnrollmentAccountScopeArgs) *DataBillingEnrollmentAccountScope {
	return &DataBillingEnrollmentAccountScope{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBillingEnrollmentAccountScope)(nil)

// DataBillingEnrollmentAccountScope represents the Terraform data resource azurerm_billing_enrollment_account_scope.
type DataBillingEnrollmentAccountScope struct {
	Name string
	Args DataBillingEnrollmentAccountScopeArgs
}

// DataSource returns the Terraform object type for [DataBillingEnrollmentAccountScope].
func (beas *DataBillingEnrollmentAccountScope) DataSource() string {
	return "azurerm_billing_enrollment_account_scope"
}

// LocalName returns the local name for [DataBillingEnrollmentAccountScope].
func (beas *DataBillingEnrollmentAccountScope) LocalName() string {
	return beas.Name
}

// Configuration returns the configuration (args) for [DataBillingEnrollmentAccountScope].
func (beas *DataBillingEnrollmentAccountScope) Configuration() interface{} {
	return beas.Args
}

// Attributes returns the attributes for [DataBillingEnrollmentAccountScope].
func (beas *DataBillingEnrollmentAccountScope) Attributes() dataBillingEnrollmentAccountScopeAttributes {
	return dataBillingEnrollmentAccountScopeAttributes{ref: terra.ReferenceDataResource(beas)}
}

// DataBillingEnrollmentAccountScopeArgs contains the configurations for azurerm_billing_enrollment_account_scope.
type DataBillingEnrollmentAccountScopeArgs struct {
	// BillingAccountName: string, required
	BillingAccountName terra.StringValue `hcl:"billing_account_name,attr" validate:"required"`
	// EnrollmentAccountName: string, required
	EnrollmentAccountName terra.StringValue `hcl:"enrollment_account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *databillingenrollmentaccountscope.Timeouts `hcl:"timeouts,block"`
}
type dataBillingEnrollmentAccountScopeAttributes struct {
	ref terra.Reference
}

// BillingAccountName returns a reference to field billing_account_name of azurerm_billing_enrollment_account_scope.
func (beas dataBillingEnrollmentAccountScopeAttributes) BillingAccountName() terra.StringValue {
	return terra.ReferenceAsString(beas.ref.Append("billing_account_name"))
}

// EnrollmentAccountName returns a reference to field enrollment_account_name of azurerm_billing_enrollment_account_scope.
func (beas dataBillingEnrollmentAccountScopeAttributes) EnrollmentAccountName() terra.StringValue {
	return terra.ReferenceAsString(beas.ref.Append("enrollment_account_name"))
}

// Id returns a reference to field id of azurerm_billing_enrollment_account_scope.
func (beas dataBillingEnrollmentAccountScopeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(beas.ref.Append("id"))
}

func (beas dataBillingEnrollmentAccountScopeAttributes) Timeouts() databillingenrollmentaccountscope.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databillingenrollmentaccountscope.TimeoutsAttributes](beas.ref.Append("timeouts"))
}
