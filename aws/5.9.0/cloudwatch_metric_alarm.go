// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudwatchmetricalarm "github.com/golingon/terraproviders/aws/5.9.0/cloudwatchmetricalarm"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudwatchMetricAlarm creates a new instance of [CloudwatchMetricAlarm].
func NewCloudwatchMetricAlarm(name string, args CloudwatchMetricAlarmArgs) *CloudwatchMetricAlarm {
	return &CloudwatchMetricAlarm{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudwatchMetricAlarm)(nil)

// CloudwatchMetricAlarm represents the Terraform resource aws_cloudwatch_metric_alarm.
type CloudwatchMetricAlarm struct {
	Name      string
	Args      CloudwatchMetricAlarmArgs
	state     *cloudwatchMetricAlarmState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudwatchMetricAlarm].
func (cma *CloudwatchMetricAlarm) Type() string {
	return "aws_cloudwatch_metric_alarm"
}

// LocalName returns the local name for [CloudwatchMetricAlarm].
func (cma *CloudwatchMetricAlarm) LocalName() string {
	return cma.Name
}

// Configuration returns the configuration (args) for [CloudwatchMetricAlarm].
func (cma *CloudwatchMetricAlarm) Configuration() interface{} {
	return cma.Args
}

// DependOn is used for other resources to depend on [CloudwatchMetricAlarm].
func (cma *CloudwatchMetricAlarm) DependOn() terra.Reference {
	return terra.ReferenceResource(cma)
}

// Dependencies returns the list of resources [CloudwatchMetricAlarm] depends_on.
func (cma *CloudwatchMetricAlarm) Dependencies() terra.Dependencies {
	return cma.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudwatchMetricAlarm].
func (cma *CloudwatchMetricAlarm) LifecycleManagement() *terra.Lifecycle {
	return cma.Lifecycle
}

// Attributes returns the attributes for [CloudwatchMetricAlarm].
func (cma *CloudwatchMetricAlarm) Attributes() cloudwatchMetricAlarmAttributes {
	return cloudwatchMetricAlarmAttributes{ref: terra.ReferenceResource(cma)}
}

// ImportState imports the given attribute values into [CloudwatchMetricAlarm]'s state.
func (cma *CloudwatchMetricAlarm) ImportState(av io.Reader) error {
	cma.state = &cloudwatchMetricAlarmState{}
	if err := json.NewDecoder(av).Decode(cma.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cma.Type(), cma.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudwatchMetricAlarm] has state.
func (cma *CloudwatchMetricAlarm) State() (*cloudwatchMetricAlarmState, bool) {
	return cma.state, cma.state != nil
}

// StateMust returns the state for [CloudwatchMetricAlarm]. Panics if the state is nil.
func (cma *CloudwatchMetricAlarm) StateMust() *cloudwatchMetricAlarmState {
	if cma.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cma.Type(), cma.LocalName()))
	}
	return cma.state
}

// CloudwatchMetricAlarmArgs contains the configurations for aws_cloudwatch_metric_alarm.
type CloudwatchMetricAlarmArgs struct {
	// ActionsEnabled: bool, optional
	ActionsEnabled terra.BoolValue `hcl:"actions_enabled,attr"`
	// AlarmActions: set of string, optional
	AlarmActions terra.SetValue[terra.StringValue] `hcl:"alarm_actions,attr"`
	// AlarmDescription: string, optional
	AlarmDescription terra.StringValue `hcl:"alarm_description,attr"`
	// AlarmName: string, required
	AlarmName terra.StringValue `hcl:"alarm_name,attr" validate:"required"`
	// ComparisonOperator: string, required
	ComparisonOperator terra.StringValue `hcl:"comparison_operator,attr" validate:"required"`
	// DatapointsToAlarm: number, optional
	DatapointsToAlarm terra.NumberValue `hcl:"datapoints_to_alarm,attr"`
	// Dimensions: map of string, optional
	Dimensions terra.MapValue[terra.StringValue] `hcl:"dimensions,attr"`
	// EvaluateLowSampleCountPercentiles: string, optional
	EvaluateLowSampleCountPercentiles terra.StringValue `hcl:"evaluate_low_sample_count_percentiles,attr"`
	// EvaluationPeriods: number, required
	EvaluationPeriods terra.NumberValue `hcl:"evaluation_periods,attr" validate:"required"`
	// ExtendedStatistic: string, optional
	ExtendedStatistic terra.StringValue `hcl:"extended_statistic,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InsufficientDataActions: set of string, optional
	InsufficientDataActions terra.SetValue[terra.StringValue] `hcl:"insufficient_data_actions,attr"`
	// MetricName: string, optional
	MetricName terra.StringValue `hcl:"metric_name,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// OkActions: set of string, optional
	OkActions terra.SetValue[terra.StringValue] `hcl:"ok_actions,attr"`
	// Period: number, optional
	Period terra.NumberValue `hcl:"period,attr"`
	// Statistic: string, optional
	Statistic terra.StringValue `hcl:"statistic,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Threshold: number, optional
	Threshold terra.NumberValue `hcl:"threshold,attr"`
	// ThresholdMetricId: string, optional
	ThresholdMetricId terra.StringValue `hcl:"threshold_metric_id,attr"`
	// TreatMissingData: string, optional
	TreatMissingData terra.StringValue `hcl:"treat_missing_data,attr"`
	// Unit: string, optional
	Unit terra.StringValue `hcl:"unit,attr"`
	// MetricQuery: min=0
	MetricQuery []cloudwatchmetricalarm.MetricQuery `hcl:"metric_query,block" validate:"min=0"`
}
type cloudwatchMetricAlarmAttributes struct {
	ref terra.Reference
}

// ActionsEnabled returns a reference to field actions_enabled of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) ActionsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cma.ref.Append("actions_enabled"))
}

// AlarmActions returns a reference to field alarm_actions of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) AlarmActions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cma.ref.Append("alarm_actions"))
}

// AlarmDescription returns a reference to field alarm_description of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) AlarmDescription() terra.StringValue {
	return terra.ReferenceAsString(cma.ref.Append("alarm_description"))
}

// AlarmName returns a reference to field alarm_name of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) AlarmName() terra.StringValue {
	return terra.ReferenceAsString(cma.ref.Append("alarm_name"))
}

// Arn returns a reference to field arn of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cma.ref.Append("arn"))
}

// ComparisonOperator returns a reference to field comparison_operator of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) ComparisonOperator() terra.StringValue {
	return terra.ReferenceAsString(cma.ref.Append("comparison_operator"))
}

// DatapointsToAlarm returns a reference to field datapoints_to_alarm of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) DatapointsToAlarm() terra.NumberValue {
	return terra.ReferenceAsNumber(cma.ref.Append("datapoints_to_alarm"))
}

// Dimensions returns a reference to field dimensions of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) Dimensions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cma.ref.Append("dimensions"))
}

// EvaluateLowSampleCountPercentiles returns a reference to field evaluate_low_sample_count_percentiles of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) EvaluateLowSampleCountPercentiles() terra.StringValue {
	return terra.ReferenceAsString(cma.ref.Append("evaluate_low_sample_count_percentiles"))
}

// EvaluationPeriods returns a reference to field evaluation_periods of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) EvaluationPeriods() terra.NumberValue {
	return terra.ReferenceAsNumber(cma.ref.Append("evaluation_periods"))
}

// ExtendedStatistic returns a reference to field extended_statistic of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) ExtendedStatistic() terra.StringValue {
	return terra.ReferenceAsString(cma.ref.Append("extended_statistic"))
}

// Id returns a reference to field id of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cma.ref.Append("id"))
}

// InsufficientDataActions returns a reference to field insufficient_data_actions of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) InsufficientDataActions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cma.ref.Append("insufficient_data_actions"))
}

// MetricName returns a reference to field metric_name of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) MetricName() terra.StringValue {
	return terra.ReferenceAsString(cma.ref.Append("metric_name"))
}

// Namespace returns a reference to field namespace of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(cma.ref.Append("namespace"))
}

// OkActions returns a reference to field ok_actions of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) OkActions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cma.ref.Append("ok_actions"))
}

// Period returns a reference to field period of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) Period() terra.NumberValue {
	return terra.ReferenceAsNumber(cma.ref.Append("period"))
}

// Statistic returns a reference to field statistic of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) Statistic() terra.StringValue {
	return terra.ReferenceAsString(cma.ref.Append("statistic"))
}

// Tags returns a reference to field tags of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cma.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cma.ref.Append("tags_all"))
}

// Threshold returns a reference to field threshold of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) Threshold() terra.NumberValue {
	return terra.ReferenceAsNumber(cma.ref.Append("threshold"))
}

// ThresholdMetricId returns a reference to field threshold_metric_id of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) ThresholdMetricId() terra.StringValue {
	return terra.ReferenceAsString(cma.ref.Append("threshold_metric_id"))
}

// TreatMissingData returns a reference to field treat_missing_data of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) TreatMissingData() terra.StringValue {
	return terra.ReferenceAsString(cma.ref.Append("treat_missing_data"))
}

// Unit returns a reference to field unit of aws_cloudwatch_metric_alarm.
func (cma cloudwatchMetricAlarmAttributes) Unit() terra.StringValue {
	return terra.ReferenceAsString(cma.ref.Append("unit"))
}

func (cma cloudwatchMetricAlarmAttributes) MetricQuery() terra.SetValue[cloudwatchmetricalarm.MetricQueryAttributes] {
	return terra.ReferenceAsSet[cloudwatchmetricalarm.MetricQueryAttributes](cma.ref.Append("metric_query"))
}

type cloudwatchMetricAlarmState struct {
	ActionsEnabled                    bool                                     `json:"actions_enabled"`
	AlarmActions                      []string                                 `json:"alarm_actions"`
	AlarmDescription                  string                                   `json:"alarm_description"`
	AlarmName                         string                                   `json:"alarm_name"`
	Arn                               string                                   `json:"arn"`
	ComparisonOperator                string                                   `json:"comparison_operator"`
	DatapointsToAlarm                 float64                                  `json:"datapoints_to_alarm"`
	Dimensions                        map[string]string                        `json:"dimensions"`
	EvaluateLowSampleCountPercentiles string                                   `json:"evaluate_low_sample_count_percentiles"`
	EvaluationPeriods                 float64                                  `json:"evaluation_periods"`
	ExtendedStatistic                 string                                   `json:"extended_statistic"`
	Id                                string                                   `json:"id"`
	InsufficientDataActions           []string                                 `json:"insufficient_data_actions"`
	MetricName                        string                                   `json:"metric_name"`
	Namespace                         string                                   `json:"namespace"`
	OkActions                         []string                                 `json:"ok_actions"`
	Period                            float64                                  `json:"period"`
	Statistic                         string                                   `json:"statistic"`
	Tags                              map[string]string                        `json:"tags"`
	TagsAll                           map[string]string                        `json:"tags_all"`
	Threshold                         float64                                  `json:"threshold"`
	ThresholdMetricId                 string                                   `json:"threshold_metric_id"`
	TreatMissingData                  string                                   `json:"treat_missing_data"`
	Unit                              string                                   `json:"unit"`
	MetricQuery                       []cloudwatchmetricalarm.MetricQueryState `json:"metric_query"`
}
