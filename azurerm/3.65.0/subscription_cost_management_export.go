// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	subscriptioncostmanagementexport "github.com/golingon/terraproviders/azurerm/3.65.0/subscriptioncostmanagementexport"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSubscriptionCostManagementExport creates a new instance of [SubscriptionCostManagementExport].
func NewSubscriptionCostManagementExport(name string, args SubscriptionCostManagementExportArgs) *SubscriptionCostManagementExport {
	return &SubscriptionCostManagementExport{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SubscriptionCostManagementExport)(nil)

// SubscriptionCostManagementExport represents the Terraform resource azurerm_subscription_cost_management_export.
type SubscriptionCostManagementExport struct {
	Name      string
	Args      SubscriptionCostManagementExportArgs
	state     *subscriptionCostManagementExportState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SubscriptionCostManagementExport].
func (scme *SubscriptionCostManagementExport) Type() string {
	return "azurerm_subscription_cost_management_export"
}

// LocalName returns the local name for [SubscriptionCostManagementExport].
func (scme *SubscriptionCostManagementExport) LocalName() string {
	return scme.Name
}

// Configuration returns the configuration (args) for [SubscriptionCostManagementExport].
func (scme *SubscriptionCostManagementExport) Configuration() interface{} {
	return scme.Args
}

// DependOn is used for other resources to depend on [SubscriptionCostManagementExport].
func (scme *SubscriptionCostManagementExport) DependOn() terra.Reference {
	return terra.ReferenceResource(scme)
}

// Dependencies returns the list of resources [SubscriptionCostManagementExport] depends_on.
func (scme *SubscriptionCostManagementExport) Dependencies() terra.Dependencies {
	return scme.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SubscriptionCostManagementExport].
func (scme *SubscriptionCostManagementExport) LifecycleManagement() *terra.Lifecycle {
	return scme.Lifecycle
}

// Attributes returns the attributes for [SubscriptionCostManagementExport].
func (scme *SubscriptionCostManagementExport) Attributes() subscriptionCostManagementExportAttributes {
	return subscriptionCostManagementExportAttributes{ref: terra.ReferenceResource(scme)}
}

// ImportState imports the given attribute values into [SubscriptionCostManagementExport]'s state.
func (scme *SubscriptionCostManagementExport) ImportState(av io.Reader) error {
	scme.state = &subscriptionCostManagementExportState{}
	if err := json.NewDecoder(av).Decode(scme.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scme.Type(), scme.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SubscriptionCostManagementExport] has state.
func (scme *SubscriptionCostManagementExport) State() (*subscriptionCostManagementExportState, bool) {
	return scme.state, scme.state != nil
}

// StateMust returns the state for [SubscriptionCostManagementExport]. Panics if the state is nil.
func (scme *SubscriptionCostManagementExport) StateMust() *subscriptionCostManagementExportState {
	if scme.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scme.Type(), scme.LocalName()))
	}
	return scme.state
}

// SubscriptionCostManagementExportArgs contains the configurations for azurerm_subscription_cost_management_export.
type SubscriptionCostManagementExportArgs struct {
	// Active: bool, optional
	Active terra.BoolValue `hcl:"active,attr"`
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
	// SubscriptionId: string, required
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr" validate:"required"`
	// ExportDataOptions: required
	ExportDataOptions *subscriptioncostmanagementexport.ExportDataOptions `hcl:"export_data_options,block" validate:"required"`
	// ExportDataStorageLocation: required
	ExportDataStorageLocation *subscriptioncostmanagementexport.ExportDataStorageLocation `hcl:"export_data_storage_location,block" validate:"required"`
	// Timeouts: optional
	Timeouts *subscriptioncostmanagementexport.Timeouts `hcl:"timeouts,block"`
}
type subscriptionCostManagementExportAttributes struct {
	ref terra.Reference
}

// Active returns a reference to field active of azurerm_subscription_cost_management_export.
func (scme subscriptionCostManagementExportAttributes) Active() terra.BoolValue {
	return terra.ReferenceAsBool(scme.ref.Append("active"))
}

// Id returns a reference to field id of azurerm_subscription_cost_management_export.
func (scme subscriptionCostManagementExportAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scme.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_subscription_cost_management_export.
func (scme subscriptionCostManagementExportAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scme.ref.Append("name"))
}

// RecurrencePeriodEndDate returns a reference to field recurrence_period_end_date of azurerm_subscription_cost_management_export.
func (scme subscriptionCostManagementExportAttributes) RecurrencePeriodEndDate() terra.StringValue {
	return terra.ReferenceAsString(scme.ref.Append("recurrence_period_end_date"))
}

// RecurrencePeriodStartDate returns a reference to field recurrence_period_start_date of azurerm_subscription_cost_management_export.
func (scme subscriptionCostManagementExportAttributes) RecurrencePeriodStartDate() terra.StringValue {
	return terra.ReferenceAsString(scme.ref.Append("recurrence_period_start_date"))
}

// RecurrenceType returns a reference to field recurrence_type of azurerm_subscription_cost_management_export.
func (scme subscriptionCostManagementExportAttributes) RecurrenceType() terra.StringValue {
	return terra.ReferenceAsString(scme.ref.Append("recurrence_type"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_subscription_cost_management_export.
func (scme subscriptionCostManagementExportAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(scme.ref.Append("subscription_id"))
}

func (scme subscriptionCostManagementExportAttributes) ExportDataOptions() terra.ListValue[subscriptioncostmanagementexport.ExportDataOptionsAttributes] {
	return terra.ReferenceAsList[subscriptioncostmanagementexport.ExportDataOptionsAttributes](scme.ref.Append("export_data_options"))
}

func (scme subscriptionCostManagementExportAttributes) ExportDataStorageLocation() terra.ListValue[subscriptioncostmanagementexport.ExportDataStorageLocationAttributes] {
	return terra.ReferenceAsList[subscriptioncostmanagementexport.ExportDataStorageLocationAttributes](scme.ref.Append("export_data_storage_location"))
}

func (scme subscriptionCostManagementExportAttributes) Timeouts() subscriptioncostmanagementexport.TimeoutsAttributes {
	return terra.ReferenceAsSingle[subscriptioncostmanagementexport.TimeoutsAttributes](scme.ref.Append("timeouts"))
}

type subscriptionCostManagementExportState struct {
	Active                    bool                                                              `json:"active"`
	Id                        string                                                            `json:"id"`
	Name                      string                                                            `json:"name"`
	RecurrencePeriodEndDate   string                                                            `json:"recurrence_period_end_date"`
	RecurrencePeriodStartDate string                                                            `json:"recurrence_period_start_date"`
	RecurrenceType            string                                                            `json:"recurrence_type"`
	SubscriptionId            string                                                            `json:"subscription_id"`
	ExportDataOptions         []subscriptioncostmanagementexport.ExportDataOptionsState         `json:"export_data_options"`
	ExportDataStorageLocation []subscriptioncostmanagementexport.ExportDataStorageLocationState `json:"export_data_storage_location"`
	Timeouts                  *subscriptioncostmanagementexport.TimeoutsState                   `json:"timeouts"`
}
