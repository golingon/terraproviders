// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_sentinel_alert_rule_threat_intelligence

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_sentinel_alert_rule_threat_intelligence.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermSentinelAlertRuleThreatIntelligenceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asarti *Resource) Type() string {
	return "azurerm_sentinel_alert_rule_threat_intelligence"
}

// LocalName returns the local name for [Resource].
func (asarti *Resource) LocalName() string {
	return asarti.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asarti *Resource) Configuration() interface{} {
	return asarti.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asarti *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asarti)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asarti *Resource) Dependencies() terra.Dependencies {
	return asarti.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asarti *Resource) LifecycleManagement() *terra.Lifecycle {
	return asarti.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asarti *Resource) Attributes() azurermSentinelAlertRuleThreatIntelligenceAttributes {
	return azurermSentinelAlertRuleThreatIntelligenceAttributes{ref: terra.ReferenceResource(asarti)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asarti *Resource) ImportState(state io.Reader) error {
	asarti.state = &azurermSentinelAlertRuleThreatIntelligenceState{}
	if err := json.NewDecoder(state).Decode(asarti.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asarti.Type(), asarti.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asarti *Resource) State() (*azurermSentinelAlertRuleThreatIntelligenceState, bool) {
	return asarti.state, asarti.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asarti *Resource) StateMust() *azurermSentinelAlertRuleThreatIntelligenceState {
	if asarti.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asarti.Type(), asarti.LocalName()))
	}
	return asarti.state
}

// Args contains the configurations for azurerm_sentinel_alert_rule_threat_intelligence.
type Args struct {
	// AlertRuleTemplateGuid: string, required
	AlertRuleTemplateGuid terra.StringValue `hcl:"alert_rule_template_guid,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermSentinelAlertRuleThreatIntelligenceAttributes struct {
	ref terra.Reference
}

// AlertRuleTemplateGuid returns a reference to field alert_rule_template_guid of azurerm_sentinel_alert_rule_threat_intelligence.
func (asarti azurermSentinelAlertRuleThreatIntelligenceAttributes) AlertRuleTemplateGuid() terra.StringValue {
	return terra.ReferenceAsString(asarti.ref.Append("alert_rule_template_guid"))
}

// Enabled returns a reference to field enabled of azurerm_sentinel_alert_rule_threat_intelligence.
func (asarti azurermSentinelAlertRuleThreatIntelligenceAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(asarti.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_sentinel_alert_rule_threat_intelligence.
func (asarti azurermSentinelAlertRuleThreatIntelligenceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asarti.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_alert_rule_threat_intelligence.
func (asarti azurermSentinelAlertRuleThreatIntelligenceAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(asarti.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_alert_rule_threat_intelligence.
func (asarti azurermSentinelAlertRuleThreatIntelligenceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asarti.ref.Append("name"))
}

func (asarti azurermSentinelAlertRuleThreatIntelligenceAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](asarti.ref.Append("timeouts"))
}

type azurermSentinelAlertRuleThreatIntelligenceState struct {
	AlertRuleTemplateGuid   string         `json:"alert_rule_template_guid"`
	Enabled                 bool           `json:"enabled"`
	Id                      string         `json:"id"`
	LogAnalyticsWorkspaceId string         `json:"log_analytics_workspace_id"`
	Name                    string         `json:"name"`
	Timeouts                *TimeoutsState `json:"timeouts"`
}
