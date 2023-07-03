// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamonitorscheduledqueryrulesalert "github.com/golingon/terraproviders/azurerm/3.63.0/datamonitorscheduledqueryrulesalert"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMonitorScheduledQueryRulesAlert creates a new instance of [DataMonitorScheduledQueryRulesAlert].
func NewDataMonitorScheduledQueryRulesAlert(name string, args DataMonitorScheduledQueryRulesAlertArgs) *DataMonitorScheduledQueryRulesAlert {
	return &DataMonitorScheduledQueryRulesAlert{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMonitorScheduledQueryRulesAlert)(nil)

// DataMonitorScheduledQueryRulesAlert represents the Terraform data resource azurerm_monitor_scheduled_query_rules_alert.
type DataMonitorScheduledQueryRulesAlert struct {
	Name string
	Args DataMonitorScheduledQueryRulesAlertArgs
}

// DataSource returns the Terraform object type for [DataMonitorScheduledQueryRulesAlert].
func (msqra *DataMonitorScheduledQueryRulesAlert) DataSource() string {
	return "azurerm_monitor_scheduled_query_rules_alert"
}

// LocalName returns the local name for [DataMonitorScheduledQueryRulesAlert].
func (msqra *DataMonitorScheduledQueryRulesAlert) LocalName() string {
	return msqra.Name
}

// Configuration returns the configuration (args) for [DataMonitorScheduledQueryRulesAlert].
func (msqra *DataMonitorScheduledQueryRulesAlert) Configuration() interface{} {
	return msqra.Args
}

// Attributes returns the attributes for [DataMonitorScheduledQueryRulesAlert].
func (msqra *DataMonitorScheduledQueryRulesAlert) Attributes() dataMonitorScheduledQueryRulesAlertAttributes {
	return dataMonitorScheduledQueryRulesAlertAttributes{ref: terra.ReferenceDataResource(msqra)}
}

// DataMonitorScheduledQueryRulesAlertArgs contains the configurations for azurerm_monitor_scheduled_query_rules_alert.
type DataMonitorScheduledQueryRulesAlertArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Action: min=0
	Action []datamonitorscheduledqueryrulesalert.Action `hcl:"action,block" validate:"min=0"`
	// Trigger: min=0
	Trigger []datamonitorscheduledqueryrulesalert.Trigger `hcl:"trigger,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datamonitorscheduledqueryrulesalert.Timeouts `hcl:"timeouts,block"`
}
type dataMonitorScheduledQueryRulesAlertAttributes struct {
	ref terra.Reference
}

// AuthorizedResourceIds returns a reference to field authorized_resource_ids of azurerm_monitor_scheduled_query_rules_alert.
func (msqra dataMonitorScheduledQueryRulesAlertAttributes) AuthorizedResourceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](msqra.ref.Append("authorized_resource_ids"))
}

// DataSourceId returns a reference to field data_source_id of azurerm_monitor_scheduled_query_rules_alert.
func (msqra dataMonitorScheduledQueryRulesAlertAttributes) DataSourceId() terra.StringValue {
	return terra.ReferenceAsString(msqra.ref.Append("data_source_id"))
}

// Description returns a reference to field description of azurerm_monitor_scheduled_query_rules_alert.
func (msqra dataMonitorScheduledQueryRulesAlertAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(msqra.ref.Append("description"))
}

// Enabled returns a reference to field enabled of azurerm_monitor_scheduled_query_rules_alert.
func (msqra dataMonitorScheduledQueryRulesAlertAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(msqra.ref.Append("enabled"))
}

// Frequency returns a reference to field frequency of azurerm_monitor_scheduled_query_rules_alert.
func (msqra dataMonitorScheduledQueryRulesAlertAttributes) Frequency() terra.NumberValue {
	return terra.ReferenceAsNumber(msqra.ref.Append("frequency"))
}

// Id returns a reference to field id of azurerm_monitor_scheduled_query_rules_alert.
func (msqra dataMonitorScheduledQueryRulesAlertAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(msqra.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_monitor_scheduled_query_rules_alert.
func (msqra dataMonitorScheduledQueryRulesAlertAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(msqra.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_monitor_scheduled_query_rules_alert.
func (msqra dataMonitorScheduledQueryRulesAlertAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(msqra.ref.Append("name"))
}

// Query returns a reference to field query of azurerm_monitor_scheduled_query_rules_alert.
func (msqra dataMonitorScheduledQueryRulesAlertAttributes) Query() terra.StringValue {
	return terra.ReferenceAsString(msqra.ref.Append("query"))
}

// QueryType returns a reference to field query_type of azurerm_monitor_scheduled_query_rules_alert.
func (msqra dataMonitorScheduledQueryRulesAlertAttributes) QueryType() terra.StringValue {
	return terra.ReferenceAsString(msqra.ref.Append("query_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_scheduled_query_rules_alert.
func (msqra dataMonitorScheduledQueryRulesAlertAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(msqra.ref.Append("resource_group_name"))
}

// Severity returns a reference to field severity of azurerm_monitor_scheduled_query_rules_alert.
func (msqra dataMonitorScheduledQueryRulesAlertAttributes) Severity() terra.NumberValue {
	return terra.ReferenceAsNumber(msqra.ref.Append("severity"))
}

// Tags returns a reference to field tags of azurerm_monitor_scheduled_query_rules_alert.
func (msqra dataMonitorScheduledQueryRulesAlertAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](msqra.ref.Append("tags"))
}

// Throttling returns a reference to field throttling of azurerm_monitor_scheduled_query_rules_alert.
func (msqra dataMonitorScheduledQueryRulesAlertAttributes) Throttling() terra.NumberValue {
	return terra.ReferenceAsNumber(msqra.ref.Append("throttling"))
}

// TimeWindow returns a reference to field time_window of azurerm_monitor_scheduled_query_rules_alert.
func (msqra dataMonitorScheduledQueryRulesAlertAttributes) TimeWindow() terra.NumberValue {
	return terra.ReferenceAsNumber(msqra.ref.Append("time_window"))
}

func (msqra dataMonitorScheduledQueryRulesAlertAttributes) Action() terra.SetValue[datamonitorscheduledqueryrulesalert.ActionAttributes] {
	return terra.ReferenceAsSet[datamonitorscheduledqueryrulesalert.ActionAttributes](msqra.ref.Append("action"))
}

func (msqra dataMonitorScheduledQueryRulesAlertAttributes) Trigger() terra.SetValue[datamonitorscheduledqueryrulesalert.TriggerAttributes] {
	return terra.ReferenceAsSet[datamonitorscheduledqueryrulesalert.TriggerAttributes](msqra.ref.Append("trigger"))
}

func (msqra dataMonitorScheduledQueryRulesAlertAttributes) Timeouts() datamonitorscheduledqueryrulesalert.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamonitorscheduledqueryrulesalert.TimeoutsAttributes](msqra.ref.Append("timeouts"))
}
