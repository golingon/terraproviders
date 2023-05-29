// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitorscheduledqueryrulesalert "github.com/golingon/terraproviders/azurerm/3.55.0/monitorscheduledqueryrulesalert"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorScheduledQueryRulesAlert creates a new instance of [MonitorScheduledQueryRulesAlert].
func NewMonitorScheduledQueryRulesAlert(name string, args MonitorScheduledQueryRulesAlertArgs) *MonitorScheduledQueryRulesAlert {
	return &MonitorScheduledQueryRulesAlert{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorScheduledQueryRulesAlert)(nil)

// MonitorScheduledQueryRulesAlert represents the Terraform resource azurerm_monitor_scheduled_query_rules_alert.
type MonitorScheduledQueryRulesAlert struct {
	Name      string
	Args      MonitorScheduledQueryRulesAlertArgs
	state     *monitorScheduledQueryRulesAlertState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorScheduledQueryRulesAlert].
func (msqra *MonitorScheduledQueryRulesAlert) Type() string {
	return "azurerm_monitor_scheduled_query_rules_alert"
}

// LocalName returns the local name for [MonitorScheduledQueryRulesAlert].
func (msqra *MonitorScheduledQueryRulesAlert) LocalName() string {
	return msqra.Name
}

// Configuration returns the configuration (args) for [MonitorScheduledQueryRulesAlert].
func (msqra *MonitorScheduledQueryRulesAlert) Configuration() interface{} {
	return msqra.Args
}

// DependOn is used for other resources to depend on [MonitorScheduledQueryRulesAlert].
func (msqra *MonitorScheduledQueryRulesAlert) DependOn() terra.Reference {
	return terra.ReferenceResource(msqra)
}

// Dependencies returns the list of resources [MonitorScheduledQueryRulesAlert] depends_on.
func (msqra *MonitorScheduledQueryRulesAlert) Dependencies() terra.Dependencies {
	return msqra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorScheduledQueryRulesAlert].
func (msqra *MonitorScheduledQueryRulesAlert) LifecycleManagement() *terra.Lifecycle {
	return msqra.Lifecycle
}

// Attributes returns the attributes for [MonitorScheduledQueryRulesAlert].
func (msqra *MonitorScheduledQueryRulesAlert) Attributes() monitorScheduledQueryRulesAlertAttributes {
	return monitorScheduledQueryRulesAlertAttributes{ref: terra.ReferenceResource(msqra)}
}

// ImportState imports the given attribute values into [MonitorScheduledQueryRulesAlert]'s state.
func (msqra *MonitorScheduledQueryRulesAlert) ImportState(av io.Reader) error {
	msqra.state = &monitorScheduledQueryRulesAlertState{}
	if err := json.NewDecoder(av).Decode(msqra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", msqra.Type(), msqra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorScheduledQueryRulesAlert] has state.
func (msqra *MonitorScheduledQueryRulesAlert) State() (*monitorScheduledQueryRulesAlertState, bool) {
	return msqra.state, msqra.state != nil
}

// StateMust returns the state for [MonitorScheduledQueryRulesAlert]. Panics if the state is nil.
func (msqra *MonitorScheduledQueryRulesAlert) StateMust() *monitorScheduledQueryRulesAlertState {
	if msqra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", msqra.Type(), msqra.LocalName()))
	}
	return msqra.state
}

// MonitorScheduledQueryRulesAlertArgs contains the configurations for azurerm_monitor_scheduled_query_rules_alert.
type MonitorScheduledQueryRulesAlertArgs struct {
	// AuthorizedResourceIds: set of string, optional
	AuthorizedResourceIds terra.SetValue[terra.StringValue] `hcl:"authorized_resource_ids,attr"`
	// AutoMitigationEnabled: bool, optional
	AutoMitigationEnabled terra.BoolValue `hcl:"auto_mitigation_enabled,attr"`
	// DataSourceId: string, required
	DataSourceId terra.StringValue `hcl:"data_source_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Frequency: number, required
	Frequency terra.NumberValue `hcl:"frequency,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Query: string, required
	Query terra.StringValue `hcl:"query,attr" validate:"required"`
	// QueryType: string, optional
	QueryType terra.StringValue `hcl:"query_type,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Severity: number, optional
	Severity terra.NumberValue `hcl:"severity,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Throttling: number, optional
	Throttling terra.NumberValue `hcl:"throttling,attr"`
	// TimeWindow: number, required
	TimeWindow terra.NumberValue `hcl:"time_window,attr" validate:"required"`
	// Action: required
	Action *monitorscheduledqueryrulesalert.Action `hcl:"action,block" validate:"required"`
	// Timeouts: optional
	Timeouts *monitorscheduledqueryrulesalert.Timeouts `hcl:"timeouts,block"`
	// Trigger: required
	Trigger *monitorscheduledqueryrulesalert.Trigger `hcl:"trigger,block" validate:"required"`
}
type monitorScheduledQueryRulesAlertAttributes struct {
	ref terra.Reference
}

// AuthorizedResourceIds returns a reference to field authorized_resource_ids of azurerm_monitor_scheduled_query_rules_alert.
func (msqra monitorScheduledQueryRulesAlertAttributes) AuthorizedResourceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](msqra.ref.Append("authorized_resource_ids"))
}

// AutoMitigationEnabled returns a reference to field auto_mitigation_enabled of azurerm_monitor_scheduled_query_rules_alert.
func (msqra monitorScheduledQueryRulesAlertAttributes) AutoMitigationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(msqra.ref.Append("auto_mitigation_enabled"))
}

// DataSourceId returns a reference to field data_source_id of azurerm_monitor_scheduled_query_rules_alert.
func (msqra monitorScheduledQueryRulesAlertAttributes) DataSourceId() terra.StringValue {
	return terra.ReferenceAsString(msqra.ref.Append("data_source_id"))
}

// Description returns a reference to field description of azurerm_monitor_scheduled_query_rules_alert.
func (msqra monitorScheduledQueryRulesAlertAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(msqra.ref.Append("description"))
}

// Enabled returns a reference to field enabled of azurerm_monitor_scheduled_query_rules_alert.
func (msqra monitorScheduledQueryRulesAlertAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(msqra.ref.Append("enabled"))
}

// Frequency returns a reference to field frequency of azurerm_monitor_scheduled_query_rules_alert.
func (msqra monitorScheduledQueryRulesAlertAttributes) Frequency() terra.NumberValue {
	return terra.ReferenceAsNumber(msqra.ref.Append("frequency"))
}

// Id returns a reference to field id of azurerm_monitor_scheduled_query_rules_alert.
func (msqra monitorScheduledQueryRulesAlertAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(msqra.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_monitor_scheduled_query_rules_alert.
func (msqra monitorScheduledQueryRulesAlertAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(msqra.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_monitor_scheduled_query_rules_alert.
func (msqra monitorScheduledQueryRulesAlertAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(msqra.ref.Append("name"))
}

// Query returns a reference to field query of azurerm_monitor_scheduled_query_rules_alert.
func (msqra monitorScheduledQueryRulesAlertAttributes) Query() terra.StringValue {
	return terra.ReferenceAsString(msqra.ref.Append("query"))
}

// QueryType returns a reference to field query_type of azurerm_monitor_scheduled_query_rules_alert.
func (msqra monitorScheduledQueryRulesAlertAttributes) QueryType() terra.StringValue {
	return terra.ReferenceAsString(msqra.ref.Append("query_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_scheduled_query_rules_alert.
func (msqra monitorScheduledQueryRulesAlertAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(msqra.ref.Append("resource_group_name"))
}

// Severity returns a reference to field severity of azurerm_monitor_scheduled_query_rules_alert.
func (msqra monitorScheduledQueryRulesAlertAttributes) Severity() terra.NumberValue {
	return terra.ReferenceAsNumber(msqra.ref.Append("severity"))
}

// Tags returns a reference to field tags of azurerm_monitor_scheduled_query_rules_alert.
func (msqra monitorScheduledQueryRulesAlertAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](msqra.ref.Append("tags"))
}

// Throttling returns a reference to field throttling of azurerm_monitor_scheduled_query_rules_alert.
func (msqra monitorScheduledQueryRulesAlertAttributes) Throttling() terra.NumberValue {
	return terra.ReferenceAsNumber(msqra.ref.Append("throttling"))
}

// TimeWindow returns a reference to field time_window of azurerm_monitor_scheduled_query_rules_alert.
func (msqra monitorScheduledQueryRulesAlertAttributes) TimeWindow() terra.NumberValue {
	return terra.ReferenceAsNumber(msqra.ref.Append("time_window"))
}

func (msqra monitorScheduledQueryRulesAlertAttributes) Action() terra.ListValue[monitorscheduledqueryrulesalert.ActionAttributes] {
	return terra.ReferenceAsList[monitorscheduledqueryrulesalert.ActionAttributes](msqra.ref.Append("action"))
}

func (msqra monitorScheduledQueryRulesAlertAttributes) Timeouts() monitorscheduledqueryrulesalert.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitorscheduledqueryrulesalert.TimeoutsAttributes](msqra.ref.Append("timeouts"))
}

func (msqra monitorScheduledQueryRulesAlertAttributes) Trigger() terra.ListValue[monitorscheduledqueryrulesalert.TriggerAttributes] {
	return terra.ReferenceAsList[monitorscheduledqueryrulesalert.TriggerAttributes](msqra.ref.Append("trigger"))
}

type monitorScheduledQueryRulesAlertState struct {
	AuthorizedResourceIds []string                                       `json:"authorized_resource_ids"`
	AutoMitigationEnabled bool                                           `json:"auto_mitigation_enabled"`
	DataSourceId          string                                         `json:"data_source_id"`
	Description           string                                         `json:"description"`
	Enabled               bool                                           `json:"enabled"`
	Frequency             float64                                        `json:"frequency"`
	Id                    string                                         `json:"id"`
	Location              string                                         `json:"location"`
	Name                  string                                         `json:"name"`
	Query                 string                                         `json:"query"`
	QueryType             string                                         `json:"query_type"`
	ResourceGroupName     string                                         `json:"resource_group_name"`
	Severity              float64                                        `json:"severity"`
	Tags                  map[string]string                              `json:"tags"`
	Throttling            float64                                        `json:"throttling"`
	TimeWindow            float64                                        `json:"time_window"`
	Action                []monitorscheduledqueryrulesalert.ActionState  `json:"action"`
	Timeouts              *monitorscheduledqueryrulesalert.TimeoutsState `json:"timeouts"`
	Trigger               []monitorscheduledqueryrulesalert.TriggerState `json:"trigger"`
}