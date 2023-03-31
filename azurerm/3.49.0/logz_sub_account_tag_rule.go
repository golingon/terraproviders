// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	logzsubaccounttagrule "github.com/golingon/terraproviders/azurerm/3.49.0/logzsubaccounttagrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewLogzSubAccountTagRule(name string, args LogzSubAccountTagRuleArgs) *LogzSubAccountTagRule {
	return &LogzSubAccountTagRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogzSubAccountTagRule)(nil)

type LogzSubAccountTagRule struct {
	Name  string
	Args  LogzSubAccountTagRuleArgs
	state *logzSubAccountTagRuleState
}

func (lsatr *LogzSubAccountTagRule) Type() string {
	return "azurerm_logz_sub_account_tag_rule"
}

func (lsatr *LogzSubAccountTagRule) LocalName() string {
	return lsatr.Name
}

func (lsatr *LogzSubAccountTagRule) Configuration() interface{} {
	return lsatr.Args
}

func (lsatr *LogzSubAccountTagRule) Attributes() logzSubAccountTagRuleAttributes {
	return logzSubAccountTagRuleAttributes{ref: terra.ReferenceResource(lsatr)}
}

func (lsatr *LogzSubAccountTagRule) ImportState(av io.Reader) error {
	lsatr.state = &logzSubAccountTagRuleState{}
	if err := json.NewDecoder(av).Decode(lsatr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lsatr.Type(), lsatr.LocalName(), err)
	}
	return nil
}

func (lsatr *LogzSubAccountTagRule) State() (*logzSubAccountTagRuleState, bool) {
	return lsatr.state, lsatr.state != nil
}

func (lsatr *LogzSubAccountTagRule) StateMust() *logzSubAccountTagRuleState {
	if lsatr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lsatr.Type(), lsatr.LocalName()))
	}
	return lsatr.state
}

func (lsatr *LogzSubAccountTagRule) DependOn() terra.Reference {
	return terra.ReferenceResource(lsatr)
}

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
	// DependsOn contains resources that LogzSubAccountTagRule depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type logzSubAccountTagRuleAttributes struct {
	ref terra.Reference
}

func (lsatr logzSubAccountTagRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceString(lsatr.ref.Append("id"))
}

func (lsatr logzSubAccountTagRuleAttributes) LogzSubAccountId() terra.StringValue {
	return terra.ReferenceString(lsatr.ref.Append("logz_sub_account_id"))
}

func (lsatr logzSubAccountTagRuleAttributes) SendAadLogs() terra.BoolValue {
	return terra.ReferenceBool(lsatr.ref.Append("send_aad_logs"))
}

func (lsatr logzSubAccountTagRuleAttributes) SendActivityLogs() terra.BoolValue {
	return terra.ReferenceBool(lsatr.ref.Append("send_activity_logs"))
}

func (lsatr logzSubAccountTagRuleAttributes) SendSubscriptionLogs() terra.BoolValue {
	return terra.ReferenceBool(lsatr.ref.Append("send_subscription_logs"))
}

func (lsatr logzSubAccountTagRuleAttributes) TagFilter() terra.ListValue[logzsubaccounttagrule.TagFilterAttributes] {
	return terra.ReferenceList[logzsubaccounttagrule.TagFilterAttributes](lsatr.ref.Append("tag_filter"))
}

func (lsatr logzSubAccountTagRuleAttributes) Timeouts() logzsubaccounttagrule.TimeoutsAttributes {
	return terra.ReferenceSingle[logzsubaccounttagrule.TimeoutsAttributes](lsatr.ref.Append("timeouts"))
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
