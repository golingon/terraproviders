// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentinelalertrulemssecurityincident "github.com/golingon/terraproviders/azurerm/3.65.0/sentinelalertrulemssecurityincident"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelAlertRuleMsSecurityIncident creates a new instance of [SentinelAlertRuleMsSecurityIncident].
func NewSentinelAlertRuleMsSecurityIncident(name string, args SentinelAlertRuleMsSecurityIncidentArgs) *SentinelAlertRuleMsSecurityIncident {
	return &SentinelAlertRuleMsSecurityIncident{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelAlertRuleMsSecurityIncident)(nil)

// SentinelAlertRuleMsSecurityIncident represents the Terraform resource azurerm_sentinel_alert_rule_ms_security_incident.
type SentinelAlertRuleMsSecurityIncident struct {
	Name      string
	Args      SentinelAlertRuleMsSecurityIncidentArgs
	state     *sentinelAlertRuleMsSecurityIncidentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelAlertRuleMsSecurityIncident].
func (sarmsi *SentinelAlertRuleMsSecurityIncident) Type() string {
	return "azurerm_sentinel_alert_rule_ms_security_incident"
}

// LocalName returns the local name for [SentinelAlertRuleMsSecurityIncident].
func (sarmsi *SentinelAlertRuleMsSecurityIncident) LocalName() string {
	return sarmsi.Name
}

// Configuration returns the configuration (args) for [SentinelAlertRuleMsSecurityIncident].
func (sarmsi *SentinelAlertRuleMsSecurityIncident) Configuration() interface{} {
	return sarmsi.Args
}

// DependOn is used for other resources to depend on [SentinelAlertRuleMsSecurityIncident].
func (sarmsi *SentinelAlertRuleMsSecurityIncident) DependOn() terra.Reference {
	return terra.ReferenceResource(sarmsi)
}

// Dependencies returns the list of resources [SentinelAlertRuleMsSecurityIncident] depends_on.
func (sarmsi *SentinelAlertRuleMsSecurityIncident) Dependencies() terra.Dependencies {
	return sarmsi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelAlertRuleMsSecurityIncident].
func (sarmsi *SentinelAlertRuleMsSecurityIncident) LifecycleManagement() *terra.Lifecycle {
	return sarmsi.Lifecycle
}

// Attributes returns the attributes for [SentinelAlertRuleMsSecurityIncident].
func (sarmsi *SentinelAlertRuleMsSecurityIncident) Attributes() sentinelAlertRuleMsSecurityIncidentAttributes {
	return sentinelAlertRuleMsSecurityIncidentAttributes{ref: terra.ReferenceResource(sarmsi)}
}

// ImportState imports the given attribute values into [SentinelAlertRuleMsSecurityIncident]'s state.
func (sarmsi *SentinelAlertRuleMsSecurityIncident) ImportState(av io.Reader) error {
	sarmsi.state = &sentinelAlertRuleMsSecurityIncidentState{}
	if err := json.NewDecoder(av).Decode(sarmsi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sarmsi.Type(), sarmsi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelAlertRuleMsSecurityIncident] has state.
func (sarmsi *SentinelAlertRuleMsSecurityIncident) State() (*sentinelAlertRuleMsSecurityIncidentState, bool) {
	return sarmsi.state, sarmsi.state != nil
}

// StateMust returns the state for [SentinelAlertRuleMsSecurityIncident]. Panics if the state is nil.
func (sarmsi *SentinelAlertRuleMsSecurityIncident) StateMust() *sentinelAlertRuleMsSecurityIncidentState {
	if sarmsi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sarmsi.Type(), sarmsi.LocalName()))
	}
	return sarmsi.state
}

// SentinelAlertRuleMsSecurityIncidentArgs contains the configurations for azurerm_sentinel_alert_rule_ms_security_incident.
type SentinelAlertRuleMsSecurityIncidentArgs struct {
	// AlertRuleTemplateGuid: string, optional
	AlertRuleTemplateGuid terra.StringValue `hcl:"alert_rule_template_guid,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// DisplayNameExcludeFilter: set of string, optional
	DisplayNameExcludeFilter terra.SetValue[terra.StringValue] `hcl:"display_name_exclude_filter,attr"`
	// DisplayNameFilter: set of string, optional
	DisplayNameFilter terra.SetValue[terra.StringValue] `hcl:"display_name_filter,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProductFilter: string, required
	ProductFilter terra.StringValue `hcl:"product_filter,attr" validate:"required"`
	// SeverityFilter: set of string, required
	SeverityFilter terra.SetValue[terra.StringValue] `hcl:"severity_filter,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *sentinelalertrulemssecurityincident.Timeouts `hcl:"timeouts,block"`
}
type sentinelAlertRuleMsSecurityIncidentAttributes struct {
	ref terra.Reference
}

// AlertRuleTemplateGuid returns a reference to field alert_rule_template_guid of azurerm_sentinel_alert_rule_ms_security_incident.
func (sarmsi sentinelAlertRuleMsSecurityIncidentAttributes) AlertRuleTemplateGuid() terra.StringValue {
	return terra.ReferenceAsString(sarmsi.ref.Append("alert_rule_template_guid"))
}

// Description returns a reference to field description of azurerm_sentinel_alert_rule_ms_security_incident.
func (sarmsi sentinelAlertRuleMsSecurityIncidentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sarmsi.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_sentinel_alert_rule_ms_security_incident.
func (sarmsi sentinelAlertRuleMsSecurityIncidentAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(sarmsi.ref.Append("display_name"))
}

// DisplayNameExcludeFilter returns a reference to field display_name_exclude_filter of azurerm_sentinel_alert_rule_ms_security_incident.
func (sarmsi sentinelAlertRuleMsSecurityIncidentAttributes) DisplayNameExcludeFilter() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sarmsi.ref.Append("display_name_exclude_filter"))
}

// DisplayNameFilter returns a reference to field display_name_filter of azurerm_sentinel_alert_rule_ms_security_incident.
func (sarmsi sentinelAlertRuleMsSecurityIncidentAttributes) DisplayNameFilter() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sarmsi.ref.Append("display_name_filter"))
}

// Enabled returns a reference to field enabled of azurerm_sentinel_alert_rule_ms_security_incident.
func (sarmsi sentinelAlertRuleMsSecurityIncidentAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sarmsi.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_sentinel_alert_rule_ms_security_incident.
func (sarmsi sentinelAlertRuleMsSecurityIncidentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sarmsi.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_alert_rule_ms_security_incident.
func (sarmsi sentinelAlertRuleMsSecurityIncidentAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sarmsi.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_alert_rule_ms_security_incident.
func (sarmsi sentinelAlertRuleMsSecurityIncidentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sarmsi.ref.Append("name"))
}

// ProductFilter returns a reference to field product_filter of azurerm_sentinel_alert_rule_ms_security_incident.
func (sarmsi sentinelAlertRuleMsSecurityIncidentAttributes) ProductFilter() terra.StringValue {
	return terra.ReferenceAsString(sarmsi.ref.Append("product_filter"))
}

// SeverityFilter returns a reference to field severity_filter of azurerm_sentinel_alert_rule_ms_security_incident.
func (sarmsi sentinelAlertRuleMsSecurityIncidentAttributes) SeverityFilter() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sarmsi.ref.Append("severity_filter"))
}

func (sarmsi sentinelAlertRuleMsSecurityIncidentAttributes) Timeouts() sentinelalertrulemssecurityincident.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentinelalertrulemssecurityincident.TimeoutsAttributes](sarmsi.ref.Append("timeouts"))
}

type sentinelAlertRuleMsSecurityIncidentState struct {
	AlertRuleTemplateGuid    string                                             `json:"alert_rule_template_guid"`
	Description              string                                             `json:"description"`
	DisplayName              string                                             `json:"display_name"`
	DisplayNameExcludeFilter []string                                           `json:"display_name_exclude_filter"`
	DisplayNameFilter        []string                                           `json:"display_name_filter"`
	Enabled                  bool                                               `json:"enabled"`
	Id                       string                                             `json:"id"`
	LogAnalyticsWorkspaceId  string                                             `json:"log_analytics_workspace_id"`
	Name                     string                                             `json:"name"`
	ProductFilter            string                                             `json:"product_filter"`
	SeverityFilter           []string                                           `json:"severity_filter"`
	Timeouts                 *sentinelalertrulemssecurityincident.TimeoutsState `json:"timeouts"`
}