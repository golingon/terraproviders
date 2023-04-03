// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	logzsubaccounttagrule "github.com/golingon/terraproviders/azurerm/3.49.0/logzsubaccounttagrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogzSubAccountTagRule creates a new instance of [LogzSubAccountTagRule].
func NewLogzSubAccountTagRule(name string, args LogzSubAccountTagRuleArgs) *LogzSubAccountTagRule {
	return &LogzSubAccountTagRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogzSubAccountTagRule)(nil)

// LogzSubAccountTagRule represents the Terraform resource azurerm_logz_sub_account_tag_rule.
type LogzSubAccountTagRule struct {
	Name      string
	Args      LogzSubAccountTagRuleArgs
	state     *logzSubAccountTagRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogzSubAccountTagRule].
func (lsatr *LogzSubAccountTagRule) Type() string {
	return "azurerm_logz_sub_account_tag_rule"
}

// LocalName returns the local name for [LogzSubAccountTagRule].
func (lsatr *LogzSubAccountTagRule) LocalName() string {
	return lsatr.Name
}

// Configuration returns the configuration (args) for [LogzSubAccountTagRule].
func (lsatr *LogzSubAccountTagRule) Configuration() interface{} {
	return lsatr.Args
}

// DependOn is used for other resources to depend on [LogzSubAccountTagRule].
func (lsatr *LogzSubAccountTagRule) DependOn() terra.Reference {
	return terra.ReferenceResource(lsatr)
}

// Dependencies returns the list of resources [LogzSubAccountTagRule] depends_on.
func (lsatr *LogzSubAccountTagRule) Dependencies() terra.Dependencies {
	return lsatr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogzSubAccountTagRule].
func (lsatr *LogzSubAccountTagRule) LifecycleManagement() *terra.Lifecycle {
	return lsatr.Lifecycle
}

// Attributes returns the attributes for [LogzSubAccountTagRule].
func (lsatr *LogzSubAccountTagRule) Attributes() logzSubAccountTagRuleAttributes {
	return logzSubAccountTagRuleAttributes{ref: terra.ReferenceResource(lsatr)}
}

// ImportState imports the given attribute values into [LogzSubAccountTagRule]'s state.
func (lsatr *LogzSubAccountTagRule) ImportState(av io.Reader) error {
	lsatr.state = &logzSubAccountTagRuleState{}
	if err := json.NewDecoder(av).Decode(lsatr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lsatr.Type(), lsatr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogzSubAccountTagRule] has state.
func (lsatr *LogzSubAccountTagRule) State() (*logzSubAccountTagRuleState, bool) {
	return lsatr.state, lsatr.state != nil
}

// StateMust returns the state for [LogzSubAccountTagRule]. Panics if the state is nil.
func (lsatr *LogzSubAccountTagRule) StateMust() *logzSubAccountTagRuleState {
	if lsatr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lsatr.Type(), lsatr.LocalName()))
	}
	return lsatr.state
}

// LogzSubAccountTagRuleArgs contains the configurations for azurerm_logz_sub_account_tag_rule.
type LogzSubAccountTagRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogzSubAccountId: string, required
	LogzSubAccountId terra.StringValue `hcl:"logz_sub_account_id,attr" validate:"required"`
	// SendAadLogs: bool, optional
	SendAadLogs terra.BoolValue `hcl:"send_aad_logs,attr"`
	// SendActivityLogs: bool, optional
	SendActivityLogs terra.BoolValue `hcl:"send_activity_logs,attr"`
	// SendSubscriptionLogs: bool, optional
	SendSubscriptionLogs terra.BoolValue `hcl:"send_subscription_logs,attr"`
	// TagFilter: min=0,max=10
	TagFilter []logzsubaccounttagrule.TagFilter `hcl:"tag_filter,block" validate:"min=0,max=10"`
	// Timeouts: optional
	Timeouts *logzsubaccounttagrule.Timeouts `hcl:"timeouts,block"`
}
type logzSubAccountTagRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_logz_sub_account_tag_rule.
func (lsatr logzSubAccountTagRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lsatr.ref.Append("id"))
}

// LogzSubAccountId returns a reference to field logz_sub_account_id of azurerm_logz_sub_account_tag_rule.
func (lsatr logzSubAccountTagRuleAttributes) LogzSubAccountId() terra.StringValue {
	return terra.ReferenceAsString(lsatr.ref.Append("logz_sub_account_id"))
}

// SendAadLogs returns a reference to field send_aad_logs of azurerm_logz_sub_account_tag_rule.
func (lsatr logzSubAccountTagRuleAttributes) SendAadLogs() terra.BoolValue {
	return terra.ReferenceAsBool(lsatr.ref.Append("send_aad_logs"))
}

// SendActivityLogs returns a reference to field send_activity_logs of azurerm_logz_sub_account_tag_rule.
func (lsatr logzSubAccountTagRuleAttributes) SendActivityLogs() terra.BoolValue {
	return terra.ReferenceAsBool(lsatr.ref.Append("send_activity_logs"))
}

// SendSubscriptionLogs returns a reference to field send_subscription_logs of azurerm_logz_sub_account_tag_rule.
func (lsatr logzSubAccountTagRuleAttributes) SendSubscriptionLogs() terra.BoolValue {
	return terra.ReferenceAsBool(lsatr.ref.Append("send_subscription_logs"))
}

func (lsatr logzSubAccountTagRuleAttributes) TagFilter() terra.ListValue[logzsubaccounttagrule.TagFilterAttributes] {
	return terra.ReferenceAsList[logzsubaccounttagrule.TagFilterAttributes](lsatr.ref.Append("tag_filter"))
}

func (lsatr logzSubAccountTagRuleAttributes) Timeouts() logzsubaccounttagrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[logzsubaccounttagrule.TimeoutsAttributes](lsatr.ref.Append("timeouts"))
}

type logzSubAccountTagRuleState struct {
	Id                   string                                 `json:"id"`
	LogzSubAccountId     string                                 `json:"logz_sub_account_id"`
	SendAadLogs          bool                                   `json:"send_aad_logs"`
	SendActivityLogs     bool                                   `json:"send_activity_logs"`
	SendSubscriptionLogs bool                                   `json:"send_subscription_logs"`
	TagFilter            []logzsubaccounttagrule.TagFilterState `json:"tag_filter"`
	Timeouts             *logzsubaccounttagrule.TimeoutsState   `json:"timeouts"`
}
