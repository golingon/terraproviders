// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	sentinelalertrulethreatintelligence "github.com/golingon/terraproviders/azurerm/3.98.0/sentinelalertrulethreatintelligence"
	"io"
)

// NewSentinelAlertRuleThreatIntelligence creates a new instance of [SentinelAlertRuleThreatIntelligence].
func NewSentinelAlertRuleThreatIntelligence(name string, args SentinelAlertRuleThreatIntelligenceArgs) *SentinelAlertRuleThreatIntelligence {
	return &SentinelAlertRuleThreatIntelligence{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelAlertRuleThreatIntelligence)(nil)

// SentinelAlertRuleThreatIntelligence represents the Terraform resource azurerm_sentinel_alert_rule_threat_intelligence.
type SentinelAlertRuleThreatIntelligence struct {
	Name      string
	Args      SentinelAlertRuleThreatIntelligenceArgs
	state     *sentinelAlertRuleThreatIntelligenceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelAlertRuleThreatIntelligence].
func (sarti *SentinelAlertRuleThreatIntelligence) Type() string {
	return "azurerm_sentinel_alert_rule_threat_intelligence"
}

// LocalName returns the local name for [SentinelAlertRuleThreatIntelligence].
func (sarti *SentinelAlertRuleThreatIntelligence) LocalName() string {
	return sarti.Name
}

// Configuration returns the configuration (args) for [SentinelAlertRuleThreatIntelligence].
func (sarti *SentinelAlertRuleThreatIntelligence) Configuration() interface{} {
	return sarti.Args
}

// DependOn is used for other resources to depend on [SentinelAlertRuleThreatIntelligence].
func (sarti *SentinelAlertRuleThreatIntelligence) DependOn() terra.Reference {
	return terra.ReferenceResource(sarti)
}

// Dependencies returns the list of resources [SentinelAlertRuleThreatIntelligence] depends_on.
func (sarti *SentinelAlertRuleThreatIntelligence) Dependencies() terra.Dependencies {
	return sarti.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelAlertRuleThreatIntelligence].
func (sarti *SentinelAlertRuleThreatIntelligence) LifecycleManagement() *terra.Lifecycle {
	return sarti.Lifecycle
}

// Attributes returns the attributes for [SentinelAlertRuleThreatIntelligence].
func (sarti *SentinelAlertRuleThreatIntelligence) Attributes() sentinelAlertRuleThreatIntelligenceAttributes {
	return sentinelAlertRuleThreatIntelligenceAttributes{ref: terra.ReferenceResource(sarti)}
}

// ImportState imports the given attribute values into [SentinelAlertRuleThreatIntelligence]'s state.
func (sarti *SentinelAlertRuleThreatIntelligence) ImportState(av io.Reader) error {
	sarti.state = &sentinelAlertRuleThreatIntelligenceState{}
	if err := json.NewDecoder(av).Decode(sarti.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sarti.Type(), sarti.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelAlertRuleThreatIntelligence] has state.
func (sarti *SentinelAlertRuleThreatIntelligence) State() (*sentinelAlertRuleThreatIntelligenceState, bool) {
	return sarti.state, sarti.state != nil
}

// StateMust returns the state for [SentinelAlertRuleThreatIntelligence]. Panics if the state is nil.
func (sarti *SentinelAlertRuleThreatIntelligence) StateMust() *sentinelAlertRuleThreatIntelligenceState {
	if sarti.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sarti.Type(), sarti.LocalName()))
	}
	return sarti.state
}

// SentinelAlertRuleThreatIntelligenceArgs contains the configurations for azurerm_sentinel_alert_rule_threat_intelligence.
type SentinelAlertRuleThreatIntelligenceArgs struct {
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
	Timeouts *sentinelalertrulethreatintelligence.Timeouts `hcl:"timeouts,block"`
}
type sentinelAlertRuleThreatIntelligenceAttributes struct {
	ref terra.Reference
}

// AlertRuleTemplateGuid returns a reference to field alert_rule_template_guid of azurerm_sentinel_alert_rule_threat_intelligence.
func (sarti sentinelAlertRuleThreatIntelligenceAttributes) AlertRuleTemplateGuid() terra.StringValue {
	return terra.ReferenceAsString(sarti.ref.Append("alert_rule_template_guid"))
}

// Enabled returns a reference to field enabled of azurerm_sentinel_alert_rule_threat_intelligence.
func (sarti sentinelAlertRuleThreatIntelligenceAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sarti.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_sentinel_alert_rule_threat_intelligence.
func (sarti sentinelAlertRuleThreatIntelligenceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sarti.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_alert_rule_threat_intelligence.
func (sarti sentinelAlertRuleThreatIntelligenceAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sarti.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_alert_rule_threat_intelligence.
func (sarti sentinelAlertRuleThreatIntelligenceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sarti.ref.Append("name"))
}

func (sarti sentinelAlertRuleThreatIntelligenceAttributes) Timeouts() sentinelalertrulethreatintelligence.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentinelalertrulethreatintelligence.TimeoutsAttributes](sarti.ref.Append("timeouts"))
}

type sentinelAlertRuleThreatIntelligenceState struct {
	AlertRuleTemplateGuid   string                                             `json:"alert_rule_template_guid"`
	Enabled                 bool                                               `json:"enabled"`
	Id                      string                                             `json:"id"`
	LogAnalyticsWorkspaceId string                                             `json:"log_analytics_workspace_id"`
	Name                    string                                             `json:"name"`
	Timeouts                *sentinelalertrulethreatintelligence.TimeoutsState `json:"timeouts"`
}
