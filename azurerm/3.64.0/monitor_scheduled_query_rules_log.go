// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitorscheduledqueryruleslog "github.com/golingon/terraproviders/azurerm/3.64.0/monitorscheduledqueryruleslog"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorScheduledQueryRulesLog creates a new instance of [MonitorScheduledQueryRulesLog].
func NewMonitorScheduledQueryRulesLog(name string, args MonitorScheduledQueryRulesLogArgs) *MonitorScheduledQueryRulesLog {
	return &MonitorScheduledQueryRulesLog{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorScheduledQueryRulesLog)(nil)

// MonitorScheduledQueryRulesLog represents the Terraform resource azurerm_monitor_scheduled_query_rules_log.
type MonitorScheduledQueryRulesLog struct {
	Name      string
	Args      MonitorScheduledQueryRulesLogArgs
	state     *monitorScheduledQueryRulesLogState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorScheduledQueryRulesLog].
func (msqrl *MonitorScheduledQueryRulesLog) Type() string {
	return "azurerm_monitor_scheduled_query_rules_log"
}

// LocalName returns the local name for [MonitorScheduledQueryRulesLog].
func (msqrl *MonitorScheduledQueryRulesLog) LocalName() string {
	return msqrl.Name
}

// Configuration returns the configuration (args) for [MonitorScheduledQueryRulesLog].
func (msqrl *MonitorScheduledQueryRulesLog) Configuration() interface{} {
	return msqrl.Args
}

// DependOn is used for other resources to depend on [MonitorScheduledQueryRulesLog].
func (msqrl *MonitorScheduledQueryRulesLog) DependOn() terra.Reference {
	return terra.ReferenceResource(msqrl)
}

// Dependencies returns the list of resources [MonitorScheduledQueryRulesLog] depends_on.
func (msqrl *MonitorScheduledQueryRulesLog) Dependencies() terra.Dependencies {
	return msqrl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorScheduledQueryRulesLog].
func (msqrl *MonitorScheduledQueryRulesLog) LifecycleManagement() *terra.Lifecycle {
	return msqrl.Lifecycle
}

// Attributes returns the attributes for [MonitorScheduledQueryRulesLog].
func (msqrl *MonitorScheduledQueryRulesLog) Attributes() monitorScheduledQueryRulesLogAttributes {
	return monitorScheduledQueryRulesLogAttributes{ref: terra.ReferenceResource(msqrl)}
}

// ImportState imports the given attribute values into [MonitorScheduledQueryRulesLog]'s state.
func (msqrl *MonitorScheduledQueryRulesLog) ImportState(av io.Reader) error {
	msqrl.state = &monitorScheduledQueryRulesLogState{}
	if err := json.NewDecoder(av).Decode(msqrl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", msqrl.Type(), msqrl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorScheduledQueryRulesLog] has state.
func (msqrl *MonitorScheduledQueryRulesLog) State() (*monitorScheduledQueryRulesLogState, bool) {
	return msqrl.state, msqrl.state != nil
}

// StateMust returns the state for [MonitorScheduledQueryRulesLog]. Panics if the state is nil.
func (msqrl *MonitorScheduledQueryRulesLog) StateMust() *monitorScheduledQueryRulesLogState {
	if msqrl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", msqrl.Type(), msqrl.LocalName()))
	}
	return msqrl.state
}

// MonitorScheduledQueryRulesLogArgs contains the configurations for azurerm_monitor_scheduled_query_rules_log.
type MonitorScheduledQueryRulesLogArgs struct {
	// AuthorizedResourceIds: set of string, optional
	AuthorizedResourceIds terra.SetValue[terra.StringValue] `hcl:"authorized_resource_ids,attr"`
	// DataSourceId: string, required
	DataSourceId terra.StringValue `hcl:"data_source_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Criteria: required
	Criteria *monitorscheduledqueryruleslog.Criteria `hcl:"criteria,block" validate:"required"`
	// Timeouts: optional
	Timeouts *monitorscheduledqueryruleslog.Timeouts `hcl:"timeouts,block"`
}
type monitorScheduledQueryRulesLogAttributes struct {
	ref terra.Reference
}

// AuthorizedResourceIds returns a reference to field authorized_resource_ids of azurerm_monitor_scheduled_query_rules_log.
func (msqrl monitorScheduledQueryRulesLogAttributes) AuthorizedResourceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](msqrl.ref.Append("authorized_resource_ids"))
}

// DataSourceId returns a reference to field data_source_id of azurerm_monitor_scheduled_query_rules_log.
func (msqrl monitorScheduledQueryRulesLogAttributes) DataSourceId() terra.StringValue {
	return terra.ReferenceAsString(msqrl.ref.Append("data_source_id"))
}

// Description returns a reference to field description of azurerm_monitor_scheduled_query_rules_log.
func (msqrl monitorScheduledQueryRulesLogAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(msqrl.ref.Append("description"))
}

// Enabled returns a reference to field enabled of azurerm_monitor_scheduled_query_rules_log.
func (msqrl monitorScheduledQueryRulesLogAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(msqrl.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_monitor_scheduled_query_rules_log.
func (msqrl monitorScheduledQueryRulesLogAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(msqrl.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_monitor_scheduled_query_rules_log.
func (msqrl monitorScheduledQueryRulesLogAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(msqrl.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_monitor_scheduled_query_rules_log.
func (msqrl monitorScheduledQueryRulesLogAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(msqrl.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_scheduled_query_rules_log.
func (msqrl monitorScheduledQueryRulesLogAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(msqrl.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_monitor_scheduled_query_rules_log.
func (msqrl monitorScheduledQueryRulesLogAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](msqrl.ref.Append("tags"))
}

func (msqrl monitorScheduledQueryRulesLogAttributes) Criteria() terra.ListValue[monitorscheduledqueryruleslog.CriteriaAttributes] {
	return terra.ReferenceAsList[monitorscheduledqueryruleslog.CriteriaAttributes](msqrl.ref.Append("criteria"))
}

func (msqrl monitorScheduledQueryRulesLogAttributes) Timeouts() monitorscheduledqueryruleslog.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitorscheduledqueryruleslog.TimeoutsAttributes](msqrl.ref.Append("timeouts"))
}

type monitorScheduledQueryRulesLogState struct {
	AuthorizedResourceIds []string                                      `json:"authorized_resource_ids"`
	DataSourceId          string                                        `json:"data_source_id"`
	Description           string                                        `json:"description"`
	Enabled               bool                                          `json:"enabled"`
	Id                    string                                        `json:"id"`
	Location              string                                        `json:"location"`
	Name                  string                                        `json:"name"`
	ResourceGroupName     string                                        `json:"resource_group_name"`
	Tags                  map[string]string                             `json:"tags"`
	Criteria              []monitorscheduledqueryruleslog.CriteriaState `json:"criteria"`
	Timeouts              *monitorscheduledqueryruleslog.TimeoutsState  `json:"timeouts"`
}
