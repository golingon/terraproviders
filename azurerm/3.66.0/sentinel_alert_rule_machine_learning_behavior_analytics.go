// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentinelalertrulemachinelearningbehavioranalytics "github.com/golingon/terraproviders/azurerm/3.66.0/sentinelalertrulemachinelearningbehavioranalytics"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelAlertRuleMachineLearningBehaviorAnalytics creates a new instance of [SentinelAlertRuleMachineLearningBehaviorAnalytics].
func NewSentinelAlertRuleMachineLearningBehaviorAnalytics(name string, args SentinelAlertRuleMachineLearningBehaviorAnalyticsArgs) *SentinelAlertRuleMachineLearningBehaviorAnalytics {
	return &SentinelAlertRuleMachineLearningBehaviorAnalytics{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelAlertRuleMachineLearningBehaviorAnalytics)(nil)

// SentinelAlertRuleMachineLearningBehaviorAnalytics represents the Terraform resource azurerm_sentinel_alert_rule_machine_learning_behavior_analytics.
type SentinelAlertRuleMachineLearningBehaviorAnalytics struct {
	Name      string
	Args      SentinelAlertRuleMachineLearningBehaviorAnalyticsArgs
	state     *sentinelAlertRuleMachineLearningBehaviorAnalyticsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelAlertRuleMachineLearningBehaviorAnalytics].
func (sarmlba *SentinelAlertRuleMachineLearningBehaviorAnalytics) Type() string {
	return "azurerm_sentinel_alert_rule_machine_learning_behavior_analytics"
}

// LocalName returns the local name for [SentinelAlertRuleMachineLearningBehaviorAnalytics].
func (sarmlba *SentinelAlertRuleMachineLearningBehaviorAnalytics) LocalName() string {
	return sarmlba.Name
}

// Configuration returns the configuration (args) for [SentinelAlertRuleMachineLearningBehaviorAnalytics].
func (sarmlba *SentinelAlertRuleMachineLearningBehaviorAnalytics) Configuration() interface{} {
	return sarmlba.Args
}

// DependOn is used for other resources to depend on [SentinelAlertRuleMachineLearningBehaviorAnalytics].
func (sarmlba *SentinelAlertRuleMachineLearningBehaviorAnalytics) DependOn() terra.Reference {
	return terra.ReferenceResource(sarmlba)
}

// Dependencies returns the list of resources [SentinelAlertRuleMachineLearningBehaviorAnalytics] depends_on.
func (sarmlba *SentinelAlertRuleMachineLearningBehaviorAnalytics) Dependencies() terra.Dependencies {
	return sarmlba.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelAlertRuleMachineLearningBehaviorAnalytics].
func (sarmlba *SentinelAlertRuleMachineLearningBehaviorAnalytics) LifecycleManagement() *terra.Lifecycle {
	return sarmlba.Lifecycle
}

// Attributes returns the attributes for [SentinelAlertRuleMachineLearningBehaviorAnalytics].
func (sarmlba *SentinelAlertRuleMachineLearningBehaviorAnalytics) Attributes() sentinelAlertRuleMachineLearningBehaviorAnalyticsAttributes {
	return sentinelAlertRuleMachineLearningBehaviorAnalyticsAttributes{ref: terra.ReferenceResource(sarmlba)}
}

// ImportState imports the given attribute values into [SentinelAlertRuleMachineLearningBehaviorAnalytics]'s state.
func (sarmlba *SentinelAlertRuleMachineLearningBehaviorAnalytics) ImportState(av io.Reader) error {
	sarmlba.state = &sentinelAlertRuleMachineLearningBehaviorAnalyticsState{}
	if err := json.NewDecoder(av).Decode(sarmlba.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sarmlba.Type(), sarmlba.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelAlertRuleMachineLearningBehaviorAnalytics] has state.
func (sarmlba *SentinelAlertRuleMachineLearningBehaviorAnalytics) State() (*sentinelAlertRuleMachineLearningBehaviorAnalyticsState, bool) {
	return sarmlba.state, sarmlba.state != nil
}

// StateMust returns the state for [SentinelAlertRuleMachineLearningBehaviorAnalytics]. Panics if the state is nil.
func (sarmlba *SentinelAlertRuleMachineLearningBehaviorAnalytics) StateMust() *sentinelAlertRuleMachineLearningBehaviorAnalyticsState {
	if sarmlba.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sarmlba.Type(), sarmlba.LocalName()))
	}
	return sarmlba.state
}

// SentinelAlertRuleMachineLearningBehaviorAnalyticsArgs contains the configurations for azurerm_sentinel_alert_rule_machine_learning_behavior_analytics.
type SentinelAlertRuleMachineLearningBehaviorAnalyticsArgs struct {
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
	Timeouts *sentinelalertrulemachinelearningbehavioranalytics.Timeouts `hcl:"timeouts,block"`
}
type sentinelAlertRuleMachineLearningBehaviorAnalyticsAttributes struct {
	ref terra.Reference
}

// AlertRuleTemplateGuid returns a reference to field alert_rule_template_guid of azurerm_sentinel_alert_rule_machine_learning_behavior_analytics.
func (sarmlba sentinelAlertRuleMachineLearningBehaviorAnalyticsAttributes) AlertRuleTemplateGuid() terra.StringValue {
	return terra.ReferenceAsString(sarmlba.ref.Append("alert_rule_template_guid"))
}

// Enabled returns a reference to field enabled of azurerm_sentinel_alert_rule_machine_learning_behavior_analytics.
func (sarmlba sentinelAlertRuleMachineLearningBehaviorAnalyticsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sarmlba.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_sentinel_alert_rule_machine_learning_behavior_analytics.
func (sarmlba sentinelAlertRuleMachineLearningBehaviorAnalyticsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sarmlba.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_alert_rule_machine_learning_behavior_analytics.
func (sarmlba sentinelAlertRuleMachineLearningBehaviorAnalyticsAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sarmlba.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_alert_rule_machine_learning_behavior_analytics.
func (sarmlba sentinelAlertRuleMachineLearningBehaviorAnalyticsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sarmlba.ref.Append("name"))
}

func (sarmlba sentinelAlertRuleMachineLearningBehaviorAnalyticsAttributes) Timeouts() sentinelalertrulemachinelearningbehavioranalytics.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentinelalertrulemachinelearningbehavioranalytics.TimeoutsAttributes](sarmlba.ref.Append("timeouts"))
}

type sentinelAlertRuleMachineLearningBehaviorAnalyticsState struct {
	AlertRuleTemplateGuid   string                                                           `json:"alert_rule_template_guid"`
	Enabled                 bool                                                             `json:"enabled"`
	Id                      string                                                           `json:"id"`
	LogAnalyticsWorkspaceId string                                                           `json:"log_analytics_workspace_id"`
	Name                    string                                                           `json:"name"`
	Timeouts                *sentinelalertrulemachinelearningbehavioranalytics.TimeoutsState `json:"timeouts"`
}