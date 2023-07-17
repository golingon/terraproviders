// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentinelalertrulescheduled "github.com/golingon/terraproviders/azurerm/3.65.0/sentinelalertrulescheduled"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelAlertRuleScheduled creates a new instance of [SentinelAlertRuleScheduled].
func NewSentinelAlertRuleScheduled(name string, args SentinelAlertRuleScheduledArgs) *SentinelAlertRuleScheduled {
	return &SentinelAlertRuleScheduled{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelAlertRuleScheduled)(nil)

// SentinelAlertRuleScheduled represents the Terraform resource azurerm_sentinel_alert_rule_scheduled.
type SentinelAlertRuleScheduled struct {
	Name      string
	Args      SentinelAlertRuleScheduledArgs
	state     *sentinelAlertRuleScheduledState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelAlertRuleScheduled].
func (sars *SentinelAlertRuleScheduled) Type() string {
	return "azurerm_sentinel_alert_rule_scheduled"
}

// LocalName returns the local name for [SentinelAlertRuleScheduled].
func (sars *SentinelAlertRuleScheduled) LocalName() string {
	return sars.Name
}

// Configuration returns the configuration (args) for [SentinelAlertRuleScheduled].
func (sars *SentinelAlertRuleScheduled) Configuration() interface{} {
	return sars.Args
}

// DependOn is used for other resources to depend on [SentinelAlertRuleScheduled].
func (sars *SentinelAlertRuleScheduled) DependOn() terra.Reference {
	return terra.ReferenceResource(sars)
}

// Dependencies returns the list of resources [SentinelAlertRuleScheduled] depends_on.
func (sars *SentinelAlertRuleScheduled) Dependencies() terra.Dependencies {
	return sars.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelAlertRuleScheduled].
func (sars *SentinelAlertRuleScheduled) LifecycleManagement() *terra.Lifecycle {
	return sars.Lifecycle
}

// Attributes returns the attributes for [SentinelAlertRuleScheduled].
func (sars *SentinelAlertRuleScheduled) Attributes() sentinelAlertRuleScheduledAttributes {
	return sentinelAlertRuleScheduledAttributes{ref: terra.ReferenceResource(sars)}
}

// ImportState imports the given attribute values into [SentinelAlertRuleScheduled]'s state.
func (sars *SentinelAlertRuleScheduled) ImportState(av io.Reader) error {
	sars.state = &sentinelAlertRuleScheduledState{}
	if err := json.NewDecoder(av).Decode(sars.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sars.Type(), sars.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelAlertRuleScheduled] has state.
func (sars *SentinelAlertRuleScheduled) State() (*sentinelAlertRuleScheduledState, bool) {
	return sars.state, sars.state != nil
}

// StateMust returns the state for [SentinelAlertRuleScheduled]. Panics if the state is nil.
func (sars *SentinelAlertRuleScheduled) StateMust() *sentinelAlertRuleScheduledState {
	if sars.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sars.Type(), sars.LocalName()))
	}
	return sars.state
}

// SentinelAlertRuleScheduledArgs contains the configurations for azurerm_sentinel_alert_rule_scheduled.
type SentinelAlertRuleScheduledArgs struct {
	// AlertRuleTemplateGuid: string, optional
	AlertRuleTemplateGuid terra.StringValue `hcl:"alert_rule_template_guid,attr"`
	// AlertRuleTemplateVersion: string, optional
	AlertRuleTemplateVersion terra.StringValue `hcl:"alert_rule_template_version,attr"`
	// CustomDetails: map of string, optional
	CustomDetails terra.MapValue[terra.StringValue] `hcl:"custom_details,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Query: string, required
	Query terra.StringValue `hcl:"query,attr" validate:"required"`
	// QueryFrequency: string, optional
	QueryFrequency terra.StringValue `hcl:"query_frequency,attr"`
	// QueryPeriod: string, optional
	QueryPeriod terra.StringValue `hcl:"query_period,attr"`
	// Severity: string, required
	Severity terra.StringValue `hcl:"severity,attr" validate:"required"`
	// SuppressionDuration: string, optional
	SuppressionDuration terra.StringValue `hcl:"suppression_duration,attr"`
	// SuppressionEnabled: bool, optional
	SuppressionEnabled terra.BoolValue `hcl:"suppression_enabled,attr"`
	// Tactics: set of string, optional
	Tactics terra.SetValue[terra.StringValue] `hcl:"tactics,attr"`
	// Techniques: set of string, optional
	Techniques terra.SetValue[terra.StringValue] `hcl:"techniques,attr"`
	// TriggerOperator: string, optional
	TriggerOperator terra.StringValue `hcl:"trigger_operator,attr"`
	// TriggerThreshold: number, optional
	TriggerThreshold terra.NumberValue `hcl:"trigger_threshold,attr"`
	// AlertDetailsOverride: min=0
	AlertDetailsOverride []sentinelalertrulescheduled.AlertDetailsOverride `hcl:"alert_details_override,block" validate:"min=0"`
	// EntityMapping: min=0,max=5
	EntityMapping []sentinelalertrulescheduled.EntityMapping `hcl:"entity_mapping,block" validate:"min=0,max=5"`
	// EventGrouping: optional
	EventGrouping *sentinelalertrulescheduled.EventGrouping `hcl:"event_grouping,block"`
	// IncidentConfiguration: optional
	IncidentConfiguration *sentinelalertrulescheduled.IncidentConfiguration `hcl:"incident_configuration,block"`
	// SentinelEntityMapping: min=0,max=5
	SentinelEntityMapping []sentinelalertrulescheduled.SentinelEntityMapping `hcl:"sentinel_entity_mapping,block" validate:"min=0,max=5"`
	// Timeouts: optional
	Timeouts *sentinelalertrulescheduled.Timeouts `hcl:"timeouts,block"`
}
type sentinelAlertRuleScheduledAttributes struct {
	ref terra.Reference
}

// AlertRuleTemplateGuid returns a reference to field alert_rule_template_guid of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) AlertRuleTemplateGuid() terra.StringValue {
	return terra.ReferenceAsString(sars.ref.Append("alert_rule_template_guid"))
}

// AlertRuleTemplateVersion returns a reference to field alert_rule_template_version of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) AlertRuleTemplateVersion() terra.StringValue {
	return terra.ReferenceAsString(sars.ref.Append("alert_rule_template_version"))
}

// CustomDetails returns a reference to field custom_details of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) CustomDetails() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sars.ref.Append("custom_details"))
}

// Description returns a reference to field description of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sars.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(sars.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sars.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sars.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sars.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sars.ref.Append("name"))
}

// Query returns a reference to field query of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) Query() terra.StringValue {
	return terra.ReferenceAsString(sars.ref.Append("query"))
}

// QueryFrequency returns a reference to field query_frequency of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) QueryFrequency() terra.StringValue {
	return terra.ReferenceAsString(sars.ref.Append("query_frequency"))
}

// QueryPeriod returns a reference to field query_period of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) QueryPeriod() terra.StringValue {
	return terra.ReferenceAsString(sars.ref.Append("query_period"))
}

// Severity returns a reference to field severity of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) Severity() terra.StringValue {
	return terra.ReferenceAsString(sars.ref.Append("severity"))
}

// SuppressionDuration returns a reference to field suppression_duration of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) SuppressionDuration() terra.StringValue {
	return terra.ReferenceAsString(sars.ref.Append("suppression_duration"))
}

// SuppressionEnabled returns a reference to field suppression_enabled of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) SuppressionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sars.ref.Append("suppression_enabled"))
}

// Tactics returns a reference to field tactics of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) Tactics() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sars.ref.Append("tactics"))
}

// Techniques returns a reference to field techniques of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) Techniques() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sars.ref.Append("techniques"))
}

// TriggerOperator returns a reference to field trigger_operator of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) TriggerOperator() terra.StringValue {
	return terra.ReferenceAsString(sars.ref.Append("trigger_operator"))
}

// TriggerThreshold returns a reference to field trigger_threshold of azurerm_sentinel_alert_rule_scheduled.
func (sars sentinelAlertRuleScheduledAttributes) TriggerThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(sars.ref.Append("trigger_threshold"))
}

func (sars sentinelAlertRuleScheduledAttributes) AlertDetailsOverride() terra.ListValue[sentinelalertrulescheduled.AlertDetailsOverrideAttributes] {
	return terra.ReferenceAsList[sentinelalertrulescheduled.AlertDetailsOverrideAttributes](sars.ref.Append("alert_details_override"))
}

func (sars sentinelAlertRuleScheduledAttributes) EntityMapping() terra.ListValue[sentinelalertrulescheduled.EntityMappingAttributes] {
	return terra.ReferenceAsList[sentinelalertrulescheduled.EntityMappingAttributes](sars.ref.Append("entity_mapping"))
}

func (sars sentinelAlertRuleScheduledAttributes) EventGrouping() terra.ListValue[sentinelalertrulescheduled.EventGroupingAttributes] {
	return terra.ReferenceAsList[sentinelalertrulescheduled.EventGroupingAttributes](sars.ref.Append("event_grouping"))
}

func (sars sentinelAlertRuleScheduledAttributes) IncidentConfiguration() terra.ListValue[sentinelalertrulescheduled.IncidentConfigurationAttributes] {
	return terra.ReferenceAsList[sentinelalertrulescheduled.IncidentConfigurationAttributes](sars.ref.Append("incident_configuration"))
}

func (sars sentinelAlertRuleScheduledAttributes) SentinelEntityMapping() terra.ListValue[sentinelalertrulescheduled.SentinelEntityMappingAttributes] {
	return terra.ReferenceAsList[sentinelalertrulescheduled.SentinelEntityMappingAttributes](sars.ref.Append("sentinel_entity_mapping"))
}

func (sars sentinelAlertRuleScheduledAttributes) Timeouts() sentinelalertrulescheduled.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentinelalertrulescheduled.TimeoutsAttributes](sars.ref.Append("timeouts"))
}

type sentinelAlertRuleScheduledState struct {
	AlertRuleTemplateGuid    string                                                  `json:"alert_rule_template_guid"`
	AlertRuleTemplateVersion string                                                  `json:"alert_rule_template_version"`
	CustomDetails            map[string]string                                       `json:"custom_details"`
	Description              string                                                  `json:"description"`
	DisplayName              string                                                  `json:"display_name"`
	Enabled                  bool                                                    `json:"enabled"`
	Id                       string                                                  `json:"id"`
	LogAnalyticsWorkspaceId  string                                                  `json:"log_analytics_workspace_id"`
	Name                     string                                                  `json:"name"`
	Query                    string                                                  `json:"query"`
	QueryFrequency           string                                                  `json:"query_frequency"`
	QueryPeriod              string                                                  `json:"query_period"`
	Severity                 string                                                  `json:"severity"`
	SuppressionDuration      string                                                  `json:"suppression_duration"`
	SuppressionEnabled       bool                                                    `json:"suppression_enabled"`
	Tactics                  []string                                                `json:"tactics"`
	Techniques               []string                                                `json:"techniques"`
	TriggerOperator          string                                                  `json:"trigger_operator"`
	TriggerThreshold         float64                                                 `json:"trigger_threshold"`
	AlertDetailsOverride     []sentinelalertrulescheduled.AlertDetailsOverrideState  `json:"alert_details_override"`
	EntityMapping            []sentinelalertrulescheduled.EntityMappingState         `json:"entity_mapping"`
	EventGrouping            []sentinelalertrulescheduled.EventGroupingState         `json:"event_grouping"`
	IncidentConfiguration    []sentinelalertrulescheduled.IncidentConfigurationState `json:"incident_configuration"`
	SentinelEntityMapping    []sentinelalertrulescheduled.SentinelEntityMappingState `json:"sentinel_entity_mapping"`
	Timeouts                 *sentinelalertrulescheduled.TimeoutsState               `json:"timeouts"`
}
