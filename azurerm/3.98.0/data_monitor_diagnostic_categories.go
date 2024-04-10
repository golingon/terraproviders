// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	datamonitordiagnosticcategories "github.com/golingon/terraproviders/azurerm/3.98.0/datamonitordiagnosticcategories"
)

// NewDataMonitorDiagnosticCategories creates a new instance of [DataMonitorDiagnosticCategories].
func NewDataMonitorDiagnosticCategories(name string, args DataMonitorDiagnosticCategoriesArgs) *DataMonitorDiagnosticCategories {
	return &DataMonitorDiagnosticCategories{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMonitorDiagnosticCategories)(nil)

// DataMonitorDiagnosticCategories represents the Terraform data resource azurerm_monitor_diagnostic_categories.
type DataMonitorDiagnosticCategories struct {
	Name string
	Args DataMonitorDiagnosticCategoriesArgs
}

// DataSource returns the Terraform object type for [DataMonitorDiagnosticCategories].
func (mdc *DataMonitorDiagnosticCategories) DataSource() string {
	return "azurerm_monitor_diagnostic_categories"
}

// LocalName returns the local name for [DataMonitorDiagnosticCategories].
func (mdc *DataMonitorDiagnosticCategories) LocalName() string {
	return mdc.Name
}

// Configuration returns the configuration (args) for [DataMonitorDiagnosticCategories].
func (mdc *DataMonitorDiagnosticCategories) Configuration() interface{} {
	return mdc.Args
}

// Attributes returns the attributes for [DataMonitorDiagnosticCategories].
func (mdc *DataMonitorDiagnosticCategories) Attributes() dataMonitorDiagnosticCategoriesAttributes {
	return dataMonitorDiagnosticCategoriesAttributes{ref: terra.ReferenceDataResource(mdc)}
}

// DataMonitorDiagnosticCategoriesArgs contains the configurations for azurerm_monitor_diagnostic_categories.
type DataMonitorDiagnosticCategoriesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceId: string, required
	ResourceId terra.StringValue `hcl:"resource_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datamonitordiagnosticcategories.Timeouts `hcl:"timeouts,block"`
}
type dataMonitorDiagnosticCategoriesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_monitor_diagnostic_categories.
func (mdc dataMonitorDiagnosticCategoriesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mdc.ref.Append("id"))
}

// LogCategoryGroups returns a reference to field log_category_groups of azurerm_monitor_diagnostic_categories.
func (mdc dataMonitorDiagnosticCategoriesAttributes) LogCategoryGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mdc.ref.Append("log_category_groups"))
}

// LogCategoryTypes returns a reference to field log_category_types of azurerm_monitor_diagnostic_categories.
func (mdc dataMonitorDiagnosticCategoriesAttributes) LogCategoryTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mdc.ref.Append("log_category_types"))
}

// Logs returns a reference to field logs of azurerm_monitor_diagnostic_categories.
func (mdc dataMonitorDiagnosticCategoriesAttributes) Logs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mdc.ref.Append("logs"))
}

// Metrics returns a reference to field metrics of azurerm_monitor_diagnostic_categories.
func (mdc dataMonitorDiagnosticCategoriesAttributes) Metrics() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mdc.ref.Append("metrics"))
}

// ResourceId returns a reference to field resource_id of azurerm_monitor_diagnostic_categories.
func (mdc dataMonitorDiagnosticCategoriesAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(mdc.ref.Append("resource_id"))
}

func (mdc dataMonitorDiagnosticCategoriesAttributes) Timeouts() datamonitordiagnosticcategories.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamonitordiagnosticcategories.TimeoutsAttributes](mdc.ref.Append("timeouts"))
}
