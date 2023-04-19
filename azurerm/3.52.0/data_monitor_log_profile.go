// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamonitorlogprofile "github.com/golingon/terraproviders/azurerm/3.52.0/datamonitorlogprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMonitorLogProfile creates a new instance of [DataMonitorLogProfile].
func NewDataMonitorLogProfile(name string, args DataMonitorLogProfileArgs) *DataMonitorLogProfile {
	return &DataMonitorLogProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMonitorLogProfile)(nil)

// DataMonitorLogProfile represents the Terraform data resource azurerm_monitor_log_profile.
type DataMonitorLogProfile struct {
	Name string
	Args DataMonitorLogProfileArgs
}

// DataSource returns the Terraform object type for [DataMonitorLogProfile].
func (mlp *DataMonitorLogProfile) DataSource() string {
	return "azurerm_monitor_log_profile"
}

// LocalName returns the local name for [DataMonitorLogProfile].
func (mlp *DataMonitorLogProfile) LocalName() string {
	return mlp.Name
}

// Configuration returns the configuration (args) for [DataMonitorLogProfile].
func (mlp *DataMonitorLogProfile) Configuration() interface{} {
	return mlp.Args
}

// Attributes returns the attributes for [DataMonitorLogProfile].
func (mlp *DataMonitorLogProfile) Attributes() dataMonitorLogProfileAttributes {
	return dataMonitorLogProfileAttributes{ref: terra.ReferenceDataResource(mlp)}
}

// DataMonitorLogProfileArgs contains the configurations for azurerm_monitor_log_profile.
type DataMonitorLogProfileArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RetentionPolicy: min=0
	RetentionPolicy []datamonitorlogprofile.RetentionPolicy `hcl:"retention_policy,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datamonitorlogprofile.Timeouts `hcl:"timeouts,block"`
}
type dataMonitorLogProfileAttributes struct {
	ref terra.Reference
}

// Categories returns a reference to field categories of azurerm_monitor_log_profile.
func (mlp dataMonitorLogProfileAttributes) Categories() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mlp.ref.Append("categories"))
}

// Id returns a reference to field id of azurerm_monitor_log_profile.
func (mlp dataMonitorLogProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mlp.ref.Append("id"))
}

// Locations returns a reference to field locations of azurerm_monitor_log_profile.
func (mlp dataMonitorLogProfileAttributes) Locations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mlp.ref.Append("locations"))
}

// Name returns a reference to field name of azurerm_monitor_log_profile.
func (mlp dataMonitorLogProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mlp.ref.Append("name"))
}

// ServicebusRuleId returns a reference to field servicebus_rule_id of azurerm_monitor_log_profile.
func (mlp dataMonitorLogProfileAttributes) ServicebusRuleId() terra.StringValue {
	return terra.ReferenceAsString(mlp.ref.Append("servicebus_rule_id"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_monitor_log_profile.
func (mlp dataMonitorLogProfileAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(mlp.ref.Append("storage_account_id"))
}

func (mlp dataMonitorLogProfileAttributes) RetentionPolicy() terra.ListValue[datamonitorlogprofile.RetentionPolicyAttributes] {
	return terra.ReferenceAsList[datamonitorlogprofile.RetentionPolicyAttributes](mlp.ref.Append("retention_policy"))
}

func (mlp dataMonitorLogProfileAttributes) Timeouts() datamonitorlogprofile.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamonitorlogprofile.TimeoutsAttributes](mlp.ref.Append("timeouts"))
}
