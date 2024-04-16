// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ce_anomaly_subscription

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

// Resource represents the Terraform resource aws_ce_anomaly_subscription.
type Resource struct {
	Name      string
	Args      Args
	state     *awsCeAnomalySubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acas *Resource) Type() string {
	return "aws_ce_anomaly_subscription"
}

// LocalName returns the local name for [Resource].
func (acas *Resource) LocalName() string {
	return acas.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acas *Resource) Configuration() interface{} {
	return acas.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acas *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acas)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acas *Resource) Dependencies() terra.Dependencies {
	return acas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acas *Resource) LifecycleManagement() *terra.Lifecycle {
	return acas.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acas *Resource) Attributes() awsCeAnomalySubscriptionAttributes {
	return awsCeAnomalySubscriptionAttributes{ref: terra.ReferenceResource(acas)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acas *Resource) ImportState(state io.Reader) error {
	acas.state = &awsCeAnomalySubscriptionState{}
	if err := json.NewDecoder(state).Decode(acas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acas.Type(), acas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acas *Resource) State() (*awsCeAnomalySubscriptionState, bool) {
	return acas.state, acas.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acas *Resource) StateMust() *awsCeAnomalySubscriptionState {
	if acas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acas.Type(), acas.LocalName()))
	}
	return acas.state
}

// Args contains the configurations for aws_ce_anomaly_subscription.
type Args struct {
	// AccountId: string, optional
	AccountId terra.StringValue `hcl:"account_id,attr"`
	// Frequency: string, required
	Frequency terra.StringValue `hcl:"frequency,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MonitorArnList: list of string, required
	MonitorArnList terra.ListValue[terra.StringValue] `hcl:"monitor_arn_list,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Subscriber: min=1
	Subscriber []Subscriber `hcl:"subscriber,block" validate:"min=1"`
	// ThresholdExpression: optional
	ThresholdExpression *ThresholdExpression `hcl:"threshold_expression,block"`
}

type awsCeAnomalySubscriptionAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_ce_anomaly_subscription.
func (acas awsCeAnomalySubscriptionAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(acas.ref.Append("account_id"))
}

// Arn returns a reference to field arn of aws_ce_anomaly_subscription.
func (acas awsCeAnomalySubscriptionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(acas.ref.Append("arn"))
}

// Frequency returns a reference to field frequency of aws_ce_anomaly_subscription.
func (acas awsCeAnomalySubscriptionAttributes) Frequency() terra.StringValue {
	return terra.ReferenceAsString(acas.ref.Append("frequency"))
}

// Id returns a reference to field id of aws_ce_anomaly_subscription.
func (acas awsCeAnomalySubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acas.ref.Append("id"))
}

// MonitorArnList returns a reference to field monitor_arn_list of aws_ce_anomaly_subscription.
func (acas awsCeAnomalySubscriptionAttributes) MonitorArnList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](acas.ref.Append("monitor_arn_list"))
}

// Name returns a reference to field name of aws_ce_anomaly_subscription.
func (acas awsCeAnomalySubscriptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acas.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_ce_anomaly_subscription.
func (acas awsCeAnomalySubscriptionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acas.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ce_anomaly_subscription.
func (acas awsCeAnomalySubscriptionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acas.ref.Append("tags_all"))
}

func (acas awsCeAnomalySubscriptionAttributes) Subscriber() terra.SetValue[SubscriberAttributes] {
	return terra.ReferenceAsSet[SubscriberAttributes](acas.ref.Append("subscriber"))
}

func (acas awsCeAnomalySubscriptionAttributes) ThresholdExpression() terra.ListValue[ThresholdExpressionAttributes] {
	return terra.ReferenceAsList[ThresholdExpressionAttributes](acas.ref.Append("threshold_expression"))
}

type awsCeAnomalySubscriptionState struct {
	AccountId           string                     `json:"account_id"`
	Arn                 string                     `json:"arn"`
	Frequency           string                     `json:"frequency"`
	Id                  string                     `json:"id"`
	MonitorArnList      []string                   `json:"monitor_arn_list"`
	Name                string                     `json:"name"`
	Tags                map[string]string          `json:"tags"`
	TagsAll             map[string]string          `json:"tags_all"`
	Subscriber          []SubscriberState          `json:"subscriber"`
	ThresholdExpression []ThresholdExpressionState `json:"threshold_expression"`
}
