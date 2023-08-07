// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	billingaccountcostmanagementexport "github.com/golingon/terraproviders/azurerm/3.68.0/billingaccountcostmanagementexport"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBillingAccountCostManagementExport creates a new instance of [BillingAccountCostManagementExport].
func NewBillingAccountCostManagementExport(name string, args BillingAccountCostManagementExportArgs) *BillingAccountCostManagementExport {
	return &BillingAccountCostManagementExport{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BillingAccountCostManagementExport)(nil)

// BillingAccountCostManagementExport represents the Terraform resource azurerm_billing_account_cost_management_export.
type BillingAccountCostManagementExport struct {
	Name      string
	Args      BillingAccountCostManagementExportArgs
	state     *billingAccountCostManagementExportState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BillingAccountCostManagementExport].
func (bacme *BillingAccountCostManagementExport) Type() string {
	return "azurerm_billing_account_cost_management_export"
}

// LocalName returns the local name for [BillingAccountCostManagementExport].
func (bacme *BillingAccountCostManagementExport) LocalName() string {
	return bacme.Name
}

// Configuration returns the configuration (args) for [BillingAccountCostManagementExport].
func (bacme *BillingAccountCostManagementExport) Configuration() interface{} {
	return bacme.Args
}

// DependOn is used for other resources to depend on [BillingAccountCostManagementExport].
func (bacme *BillingAccountCostManagementExport) DependOn() terra.Reference {
	return terra.ReferenceResource(bacme)
}

// Dependencies returns the list of resources [BillingAccountCostManagementExport] depends_on.
func (bacme *BillingAccountCostManagementExport) Dependencies() terra.Dependencies {
	return bacme.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BillingAccountCostManagementExport].
func (bacme *BillingAccountCostManagementExport) LifecycleManagement() *terra.Lifecycle {
	return bacme.Lifecycle
}

// Attributes returns the attributes for [BillingAccountCostManagementExport].
func (bacme *BillingAccountCostManagementExport) Attributes() billingAccountCostManagementExportAttributes {
	return billingAccountCostManagementExportAttributes{ref: terra.ReferenceResource(bacme)}
}

// ImportState imports the given attribute values into [BillingAccountCostManagementExport]'s state.
func (bacme *BillingAccountCostManagementExport) ImportState(av io.Reader) error {
	bacme.state = &billingAccountCostManagementExportState{}
	if err := json.NewDecoder(av).Decode(bacme.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bacme.Type(), bacme.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BillingAccountCostManagementExport] has state.
func (bacme *BillingAccountCostManagementExport) State() (*billingAccountCostManagementExportState, bool) {
	return bacme.state, bacme.state != nil
}

// StateMust returns the state for [BillingAccountCostManagementExport]. Panics if the state is nil.
func (bacme *BillingAccountCostManagementExport) StateMust() *billingAccountCostManagementExportState {
	if bacme.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bacme.Type(), bacme.LocalName()))
	}
	return bacme.state
}

// BillingAccountCostManagementExportArgs contains the configurations for azurerm_billing_account_cost_management_export.
type BillingAccountCostManagementExportArgs struct {
	// Active: bool, optional
	Active terra.BoolValue `hcl:"active,attr"`
	// BillingAccountId: string, required
	BillingAccountId terra.StringValue `hcl:"billing_account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecurrencePeriodEndDate: string, required
	RecurrencePeriodEndDate terra.StringValue `hcl:"recurrence_period_end_date,attr" validate:"required"`
	// RecurrencePeriodStartDate: string, required
	RecurrencePeriodStartDate terra.StringValue `hcl:"recurrence_period_start_date,attr" validate:"required"`
	// RecurrenceType: string, required
	RecurrenceType terra.StringValue `hcl:"recurrence_type,attr" validate:"required"`
	// ExportDataOptions: required
	ExportDataOptions *billingaccountcostmanagementexport.ExportDataOptions `hcl:"export_data_options,block" validate:"required"`
	// ExportDataStorageLocation: required
	ExportDataStorageLocation *billingaccountcostmanagementexport.ExportDataStorageLocation `hcl:"export_data_storage_location,block" validate:"required"`
	// Timeouts: optional
	Timeouts *billingaccountcostmanagementexport.Timeouts `hcl:"timeouts,block"`
}
type billingAccountCostManagementExportAttributes struct {
	ref terra.Reference
}

// Active returns a reference to field active of azurerm_billing_account_cost_management_export.
func (bacme billingAccountCostManagementExportAttributes) Active() terra.BoolValue {
	return terra.ReferenceAsBool(bacme.ref.Append("active"))
}

// BillingAccountId returns a reference to field billing_account_id of azurerm_billing_account_cost_management_export.
func (bacme billingAccountCostManagementExportAttributes) BillingAccountId() terra.StringValue {
	return terra.ReferenceAsString(bacme.ref.Append("billing_account_id"))
}

// Id returns a reference to field id of azurerm_billing_account_cost_management_export.
func (bacme billingAccountCostManagementExportAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bacme.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_billing_account_cost_management_export.
func (bacme billingAccountCostManagementExportAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bacme.ref.Append("name"))
}

// RecurrencePeriodEndDate returns a reference to field recurrence_period_end_date of azurerm_billing_account_cost_management_export.
func (bacme billingAccountCostManagementExportAttributes) RecurrencePeriodEndDate() terra.StringValue {
	return terra.ReferenceAsString(bacme.ref.Append("recurrence_period_end_date"))
}

// RecurrencePeriodStartDate returns a reference to field recurrence_period_start_date of azurerm_billing_account_cost_management_export.
func (bacme billingAccountCostManagementExportAttributes) RecurrencePeriodStartDate() terra.StringValue {
	return terra.ReferenceAsString(bacme.ref.Append("recurrence_period_start_date"))
}

// RecurrenceType returns a reference to field recurrence_type of azurerm_billing_account_cost_management_export.
func (bacme billingAccountCostManagementExportAttributes) RecurrenceType() terra.StringValue {
	return terra.ReferenceAsString(bacme.ref.Append("recurrence_type"))
}

func (bacme billingAccountCostManagementExportAttributes) ExportDataOptions() terra.ListValue[billingaccountcostmanagementexport.ExportDataOptionsAttributes] {
	return terra.ReferenceAsList[billingaccountcostmanagementexport.ExportDataOptionsAttributes](bacme.ref.Append("export_data_options"))
}

func (bacme billingAccountCostManagementExportAttributes) ExportDataStorageLocation() terra.ListValue[billingaccountcostmanagementexport.ExportDataStorageLocationAttributes] {
	return terra.ReferenceAsList[billingaccountcostmanagementexport.ExportDataStorageLocationAttributes](bacme.ref.Append("export_data_storage_location"))
}

func (bacme billingAccountCostManagementExportAttributes) Timeouts() billingaccountcostmanagementexport.TimeoutsAttributes {
	return terra.ReferenceAsSingle[billingaccountcostmanagementexport.TimeoutsAttributes](bacme.ref.Append("timeouts"))
}

type billingAccountCostManagementExportState struct {
	Active                    bool                                                                `json:"active"`
	BillingAccountId          string                                                              `json:"billing_account_id"`
	Id                        string                                                              `json:"id"`
	Name                      string                                                              `json:"name"`
	RecurrencePeriodEndDate   string                                                              `json:"recurrence_period_end_date"`
	RecurrencePeriodStartDate string                                                              `json:"recurrence_period_start_date"`
	RecurrenceType            string                                                              `json:"recurrence_type"`
	ExportDataOptions         []billingaccountcostmanagementexport.ExportDataOptionsState         `json:"export_data_options"`
	ExportDataStorageLocation []billingaccountcostmanagementexport.ExportDataStorageLocationState `json:"export_data_storage_location"`
	Timeouts                  *billingaccountcostmanagementexport.TimeoutsState                   `json:"timeouts"`
}
