// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	sentinelautomationrule "github.com/golingon/terraproviders/azurerm/3.98.0/sentinelautomationrule"
	"io"
)

// NewSentinelAutomationRule creates a new instance of [SentinelAutomationRule].
func NewSentinelAutomationRule(name string, args SentinelAutomationRuleArgs) *SentinelAutomationRule {
	return &SentinelAutomationRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelAutomationRule)(nil)

// SentinelAutomationRule represents the Terraform resource azurerm_sentinel_automation_rule.
type SentinelAutomationRule struct {
	Name      string
	Args      SentinelAutomationRuleArgs
	state     *sentinelAutomationRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelAutomationRule].
func (sar *SentinelAutomationRule) Type() string {
	return "azurerm_sentinel_automation_rule"
}

// LocalName returns the local name for [SentinelAutomationRule].
func (sar *SentinelAutomationRule) LocalName() string {
	return sar.Name
}

// Configuration returns the configuration (args) for [SentinelAutomationRule].
func (sar *SentinelAutomationRule) Configuration() interface{} {
	return sar.Args
}

// DependOn is used for other resources to depend on [SentinelAutomationRule].
func (sar *SentinelAutomationRule) DependOn() terra.Reference {
	return terra.ReferenceResource(sar)
}

// Dependencies returns the list of resources [SentinelAutomationRule] depends_on.
func (sar *SentinelAutomationRule) Dependencies() terra.Dependencies {
	return sar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelAutomationRule].
func (sar *SentinelAutomationRule) LifecycleManagement() *terra.Lifecycle {
	return sar.Lifecycle
}

// Attributes returns the attributes for [SentinelAutomationRule].
func (sar *SentinelAutomationRule) Attributes() sentinelAutomationRuleAttributes {
	return sentinelAutomationRuleAttributes{ref: terra.ReferenceResource(sar)}
}

// ImportState imports the given attribute values into [SentinelAutomationRule]'s state.
func (sar *SentinelAutomationRule) ImportState(av io.Reader) error {
	sar.state = &sentinelAutomationRuleState{}
	if err := json.NewDecoder(av).Decode(sar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sar.Type(), sar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelAutomationRule] has state.
func (sar *SentinelAutomationRule) State() (*sentinelAutomationRuleState, bool) {
	return sar.state, sar.state != nil
}

// StateMust returns the state for [SentinelAutomationRule]. Panics if the state is nil.
func (sar *SentinelAutomationRule) StateMust() *sentinelAutomationRuleState {
	if sar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sar.Type(), sar.LocalName()))
	}
	return sar.state
}

// SentinelAutomationRuleArgs contains the configurations for azurerm_sentinel_automation_rule.
type SentinelAutomationRuleArgs struct {
	// ConditionJson: string, optional
	ConditionJson terra.StringValue `hcl:"condition_json,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Expiration: string, optional
	Expiration terra.StringValue `hcl:"expiration,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Order: number, required
	Order terra.NumberValue `hcl:"order,attr" validate:"required"`
	// TriggersOn: string, optional
	TriggersOn terra.StringValue `hcl:"triggers_on,attr"`
	// TriggersWhen: string, optional
	TriggersWhen terra.StringValue `hcl:"triggers_when,attr"`
	// ActionIncident: min=0
	ActionIncident []sentinelautomationrule.ActionIncident `hcl:"action_incident,block" validate:"min=0"`
	// ActionPlaybook: min=0
	ActionPlaybook []sentinelautomationrule.ActionPlaybook `hcl:"action_playbook,block" validate:"min=0"`
	// Condition: min=0
	Condition []sentinelautomationrule.Condition `hcl:"condition,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *sentinelautomationrule.Timeouts `hcl:"timeouts,block"`
}
type sentinelAutomationRuleAttributes struct {
	ref terra.Reference
}

// ConditionJson returns a reference to field condition_json of azurerm_sentinel_automation_rule.
func (sar sentinelAutomationRuleAttributes) ConditionJson() terra.StringValue {
	return terra.ReferenceAsString(sar.ref.Append("condition_json"))
}

// DisplayName returns a reference to field display_name of azurerm_sentinel_automation_rule.
func (sar sentinelAutomationRuleAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(sar.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of azurerm_sentinel_automation_rule.
func (sar sentinelAutomationRuleAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sar.ref.Append("enabled"))
}

// Expiration returns a reference to field expiration of azurerm_sentinel_automation_rule.
func (sar sentinelAutomationRuleAttributes) Expiration() terra.StringValue {
	return terra.ReferenceAsString(sar.ref.Append("expiration"))
}

// Id returns a reference to field id of azurerm_sentinel_automation_rule.
func (sar sentinelAutomationRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sar.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_automation_rule.
func (sar sentinelAutomationRuleAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sar.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_automation_rule.
func (sar sentinelAutomationRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sar.ref.Append("name"))
}

// Order returns a reference to field order of azurerm_sentinel_automation_rule.
func (sar sentinelAutomationRuleAttributes) Order() terra.NumberValue {
	return terra.ReferenceAsNumber(sar.ref.Append("order"))
}

// TriggersOn returns a reference to field triggers_on of azurerm_sentinel_automation_rule.
func (sar sentinelAutomationRuleAttributes) TriggersOn() terra.StringValue {
	return terra.ReferenceAsString(sar.ref.Append("triggers_on"))
}

// TriggersWhen returns a reference to field triggers_when of azurerm_sentinel_automation_rule.
func (sar sentinelAutomationRuleAttributes) TriggersWhen() terra.StringValue {
	return terra.ReferenceAsString(sar.ref.Append("triggers_when"))
}

func (sar sentinelAutomationRuleAttributes) ActionIncident() terra.ListValue[sentinelautomationrule.ActionIncidentAttributes] {
	return terra.ReferenceAsList[sentinelautomationrule.ActionIncidentAttributes](sar.ref.Append("action_incident"))
}

func (sar sentinelAutomationRuleAttributes) ActionPlaybook() terra.ListValue[sentinelautomationrule.ActionPlaybookAttributes] {
	return terra.ReferenceAsList[sentinelautomationrule.ActionPlaybookAttributes](sar.ref.Append("action_playbook"))
}

func (sar sentinelAutomationRuleAttributes) Condition() terra.ListValue[sentinelautomationrule.ConditionAttributes] {
	return terra.ReferenceAsList[sentinelautomationrule.ConditionAttributes](sar.ref.Append("condition"))
}

func (sar sentinelAutomationRuleAttributes) Timeouts() sentinelautomationrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentinelautomationrule.TimeoutsAttributes](sar.ref.Append("timeouts"))
}

type sentinelAutomationRuleState struct {
	ConditionJson           string                                       `json:"condition_json"`
	DisplayName             string                                       `json:"display_name"`
	Enabled                 bool                                         `json:"enabled"`
	Expiration              string                                       `json:"expiration"`
	Id                      string                                       `json:"id"`
	LogAnalyticsWorkspaceId string                                       `json:"log_analytics_workspace_id"`
	Name                    string                                       `json:"name"`
	Order                   float64                                      `json:"order"`
	TriggersOn              string                                       `json:"triggers_on"`
	TriggersWhen            string                                       `json:"triggers_when"`
	ActionIncident          []sentinelautomationrule.ActionIncidentState `json:"action_incident"`
	ActionPlaybook          []sentinelautomationrule.ActionPlaybookState `json:"action_playbook"`
	Condition               []sentinelautomationrule.ConditionState      `json:"condition"`
	Timeouts                *sentinelautomationrule.TimeoutsState        `json:"timeouts"`
}
