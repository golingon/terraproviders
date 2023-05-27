// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	iottopicrule "github.com/golingon/terraproviders/aws/5.0.1/iottopicrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIotTopicRule creates a new instance of [IotTopicRule].
func NewIotTopicRule(name string, args IotTopicRuleArgs) *IotTopicRule {
	return &IotTopicRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotTopicRule)(nil)

// IotTopicRule represents the Terraform resource aws_iot_topic_rule.
type IotTopicRule struct {
	Name      string
	Args      IotTopicRuleArgs
	state     *iotTopicRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotTopicRule].
func (itr *IotTopicRule) Type() string {
	return "aws_iot_topic_rule"
}

// LocalName returns the local name for [IotTopicRule].
func (itr *IotTopicRule) LocalName() string {
	return itr.Name
}

// Configuration returns the configuration (args) for [IotTopicRule].
func (itr *IotTopicRule) Configuration() interface{} {
	return itr.Args
}

// DependOn is used for other resources to depend on [IotTopicRule].
func (itr *IotTopicRule) DependOn() terra.Reference {
	return terra.ReferenceResource(itr)
}

// Dependencies returns the list of resources [IotTopicRule] depends_on.
func (itr *IotTopicRule) Dependencies() terra.Dependencies {
	return itr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotTopicRule].
func (itr *IotTopicRule) LifecycleManagement() *terra.Lifecycle {
	return itr.Lifecycle
}

// Attributes returns the attributes for [IotTopicRule].
func (itr *IotTopicRule) Attributes() iotTopicRuleAttributes {
	return iotTopicRuleAttributes{ref: terra.ReferenceResource(itr)}
}

// ImportState imports the given attribute values into [IotTopicRule]'s state.
func (itr *IotTopicRule) ImportState(av io.Reader) error {
	itr.state = &iotTopicRuleState{}
	if err := json.NewDecoder(av).Decode(itr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", itr.Type(), itr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotTopicRule] has state.
func (itr *IotTopicRule) State() (*iotTopicRuleState, bool) {
	return itr.state, itr.state != nil
}

// StateMust returns the state for [IotTopicRule]. Panics if the state is nil.
func (itr *IotTopicRule) StateMust() *iotTopicRuleState {
	if itr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", itr.Type(), itr.LocalName()))
	}
	return itr.state
}

// IotTopicRuleArgs contains the configurations for aws_iot_topic_rule.
type IotTopicRuleArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Sql: string, required
	Sql terra.StringValue `hcl:"sql,attr" validate:"required"`
	// SqlVersion: string, required
	SqlVersion terra.StringValue `hcl:"sql_version,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// CloudwatchAlarm: min=0
	CloudwatchAlarm []iottopicrule.CloudwatchAlarm `hcl:"cloudwatch_alarm,block" validate:"min=0"`
	// CloudwatchLogs: min=0
	CloudwatchLogs []iottopicrule.CloudwatchLogs `hcl:"cloudwatch_logs,block" validate:"min=0"`
	// CloudwatchMetric: min=0
	CloudwatchMetric []iottopicrule.CloudwatchMetric `hcl:"cloudwatch_metric,block" validate:"min=0"`
	// Dynamodb: min=0
	Dynamodb []iottopicrule.Dynamodb `hcl:"dynamodb,block" validate:"min=0"`
	// Dynamodbv2: min=0
	Dynamodbv2 []iottopicrule.Dynamodbv2 `hcl:"dynamodbv2,block" validate:"min=0"`
	// Elasticsearch: min=0
	Elasticsearch []iottopicrule.Elasticsearch `hcl:"elasticsearch,block" validate:"min=0"`
	// ErrorAction: optional
	ErrorAction *iottopicrule.ErrorAction `hcl:"error_action,block"`
	// Firehose: min=0
	Firehose []iottopicrule.Firehose `hcl:"firehose,block" validate:"min=0"`
	// Http: min=0
	Http []iottopicrule.Http `hcl:"http,block" validate:"min=0"`
	// IotAnalytics: min=0
	IotAnalytics []iottopicrule.IotAnalytics `hcl:"iot_analytics,block" validate:"min=0"`
	// IotEvents: min=0
	IotEvents []iottopicrule.IotEvents `hcl:"iot_events,block" validate:"min=0"`
	// Kafka: min=0
	Kafka []iottopicrule.Kafka `hcl:"kafka,block" validate:"min=0"`
	// Kinesis: min=0
	Kinesis []iottopicrule.Kinesis `hcl:"kinesis,block" validate:"min=0"`
	// Lambda: min=0
	Lambda []iottopicrule.Lambda `hcl:"lambda,block" validate:"min=0"`
	// Republish: min=0
	Republish []iottopicrule.Republish `hcl:"republish,block" validate:"min=0"`
	// S3: min=0
	S3 []iottopicrule.S3 `hcl:"s3,block" validate:"min=0"`
	// Sns: min=0
	Sns []iottopicrule.Sns `hcl:"sns,block" validate:"min=0"`
	// Sqs: min=0
	Sqs []iottopicrule.Sqs `hcl:"sqs,block" validate:"min=0"`
	// StepFunctions: min=0
	StepFunctions []iottopicrule.StepFunctions `hcl:"step_functions,block" validate:"min=0"`
	// Timestream: min=0
	Timestream []iottopicrule.Timestream `hcl:"timestream,block" validate:"min=0"`
}
type iotTopicRuleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iot_topic_rule.
func (itr iotTopicRuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(itr.ref.Append("arn"))
}

// Description returns a reference to field description of aws_iot_topic_rule.
func (itr iotTopicRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(itr.ref.Append("description"))
}

// Enabled returns a reference to field enabled of aws_iot_topic_rule.
func (itr iotTopicRuleAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(itr.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_iot_topic_rule.
func (itr iotTopicRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(itr.ref.Append("id"))
}

// Name returns a reference to field name of aws_iot_topic_rule.
func (itr iotTopicRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(itr.ref.Append("name"))
}

// Sql returns a reference to field sql of aws_iot_topic_rule.
func (itr iotTopicRuleAttributes) Sql() terra.StringValue {
	return terra.ReferenceAsString(itr.ref.Append("sql"))
}

// SqlVersion returns a reference to field sql_version of aws_iot_topic_rule.
func (itr iotTopicRuleAttributes) SqlVersion() terra.StringValue {
	return terra.ReferenceAsString(itr.ref.Append("sql_version"))
}

// Tags returns a reference to field tags of aws_iot_topic_rule.
func (itr iotTopicRuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](itr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_iot_topic_rule.
func (itr iotTopicRuleAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](itr.ref.Append("tags_all"))
}

func (itr iotTopicRuleAttributes) CloudwatchAlarm() terra.SetValue[iottopicrule.CloudwatchAlarmAttributes] {
	return terra.ReferenceAsSet[iottopicrule.CloudwatchAlarmAttributes](itr.ref.Append("cloudwatch_alarm"))
}

func (itr iotTopicRuleAttributes) CloudwatchLogs() terra.SetValue[iottopicrule.CloudwatchLogsAttributes] {
	return terra.ReferenceAsSet[iottopicrule.CloudwatchLogsAttributes](itr.ref.Append("cloudwatch_logs"))
}

func (itr iotTopicRuleAttributes) CloudwatchMetric() terra.SetValue[iottopicrule.CloudwatchMetricAttributes] {
	return terra.ReferenceAsSet[iottopicrule.CloudwatchMetricAttributes](itr.ref.Append("cloudwatch_metric"))
}

func (itr iotTopicRuleAttributes) Dynamodb() terra.SetValue[iottopicrule.DynamodbAttributes] {
	return terra.ReferenceAsSet[iottopicrule.DynamodbAttributes](itr.ref.Append("dynamodb"))
}

func (itr iotTopicRuleAttributes) Dynamodbv2() terra.SetValue[iottopicrule.Dynamodbv2Attributes] {
	return terra.ReferenceAsSet[iottopicrule.Dynamodbv2Attributes](itr.ref.Append("dynamodbv2"))
}

func (itr iotTopicRuleAttributes) Elasticsearch() terra.SetValue[iottopicrule.ElasticsearchAttributes] {
	return terra.ReferenceAsSet[iottopicrule.ElasticsearchAttributes](itr.ref.Append("elasticsearch"))
}

func (itr iotTopicRuleAttributes) ErrorAction() terra.ListValue[iottopicrule.ErrorActionAttributes] {
	return terra.ReferenceAsList[iottopicrule.ErrorActionAttributes](itr.ref.Append("error_action"))
}

func (itr iotTopicRuleAttributes) Firehose() terra.SetValue[iottopicrule.FirehoseAttributes] {
	return terra.ReferenceAsSet[iottopicrule.FirehoseAttributes](itr.ref.Append("firehose"))
}

func (itr iotTopicRuleAttributes) Http() terra.SetValue[iottopicrule.HttpAttributes] {
	return terra.ReferenceAsSet[iottopicrule.HttpAttributes](itr.ref.Append("http"))
}

func (itr iotTopicRuleAttributes) IotAnalytics() terra.SetValue[iottopicrule.IotAnalyticsAttributes] {
	return terra.ReferenceAsSet[iottopicrule.IotAnalyticsAttributes](itr.ref.Append("iot_analytics"))
}

func (itr iotTopicRuleAttributes) IotEvents() terra.SetValue[iottopicrule.IotEventsAttributes] {
	return terra.ReferenceAsSet[iottopicrule.IotEventsAttributes](itr.ref.Append("iot_events"))
}

func (itr iotTopicRuleAttributes) Kafka() terra.SetValue[iottopicrule.KafkaAttributes] {
	return terra.ReferenceAsSet[iottopicrule.KafkaAttributes](itr.ref.Append("kafka"))
}

func (itr iotTopicRuleAttributes) Kinesis() terra.SetValue[iottopicrule.KinesisAttributes] {
	return terra.ReferenceAsSet[iottopicrule.KinesisAttributes](itr.ref.Append("kinesis"))
}

func (itr iotTopicRuleAttributes) Lambda() terra.SetValue[iottopicrule.LambdaAttributes] {
	return terra.ReferenceAsSet[iottopicrule.LambdaAttributes](itr.ref.Append("lambda"))
}

func (itr iotTopicRuleAttributes) Republish() terra.SetValue[iottopicrule.RepublishAttributes] {
	return terra.ReferenceAsSet[iottopicrule.RepublishAttributes](itr.ref.Append("republish"))
}

func (itr iotTopicRuleAttributes) S3() terra.SetValue[iottopicrule.S3Attributes] {
	return terra.ReferenceAsSet[iottopicrule.S3Attributes](itr.ref.Append("s3"))
}

func (itr iotTopicRuleAttributes) Sns() terra.SetValue[iottopicrule.SnsAttributes] {
	return terra.ReferenceAsSet[iottopicrule.SnsAttributes](itr.ref.Append("sns"))
}

func (itr iotTopicRuleAttributes) Sqs() terra.SetValue[iottopicrule.SqsAttributes] {
	return terra.ReferenceAsSet[iottopicrule.SqsAttributes](itr.ref.Append("sqs"))
}

func (itr iotTopicRuleAttributes) StepFunctions() terra.SetValue[iottopicrule.StepFunctionsAttributes] {
	return terra.ReferenceAsSet[iottopicrule.StepFunctionsAttributes](itr.ref.Append("step_functions"))
}

func (itr iotTopicRuleAttributes) Timestream() terra.SetValue[iottopicrule.TimestreamAttributes] {
	return terra.ReferenceAsSet[iottopicrule.TimestreamAttributes](itr.ref.Append("timestream"))
}

type iotTopicRuleState struct {
	Arn              string                               `json:"arn"`
	Description      string                               `json:"description"`
	Enabled          bool                                 `json:"enabled"`
	Id               string                               `json:"id"`
	Name             string                               `json:"name"`
	Sql              string                               `json:"sql"`
	SqlVersion       string                               `json:"sql_version"`
	Tags             map[string]string                    `json:"tags"`
	TagsAll          map[string]string                    `json:"tags_all"`
	CloudwatchAlarm  []iottopicrule.CloudwatchAlarmState  `json:"cloudwatch_alarm"`
	CloudwatchLogs   []iottopicrule.CloudwatchLogsState   `json:"cloudwatch_logs"`
	CloudwatchMetric []iottopicrule.CloudwatchMetricState `json:"cloudwatch_metric"`
	Dynamodb         []iottopicrule.DynamodbState         `json:"dynamodb"`
	Dynamodbv2       []iottopicrule.Dynamodbv2State       `json:"dynamodbv2"`
	Elasticsearch    []iottopicrule.ElasticsearchState    `json:"elasticsearch"`
	ErrorAction      []iottopicrule.ErrorActionState      `json:"error_action"`
	Firehose         []iottopicrule.FirehoseState         `json:"firehose"`
	Http             []iottopicrule.HttpState             `json:"http"`
	IotAnalytics     []iottopicrule.IotAnalyticsState     `json:"iot_analytics"`
	IotEvents        []iottopicrule.IotEventsState        `json:"iot_events"`
	Kafka            []iottopicrule.KafkaState            `json:"kafka"`
	Kinesis          []iottopicrule.KinesisState          `json:"kinesis"`
	Lambda           []iottopicrule.LambdaState           `json:"lambda"`
	Republish        []iottopicrule.RepublishState        `json:"republish"`
	S3               []iottopicrule.S3State               `json:"s3"`
	Sns              []iottopicrule.SnsState              `json:"sns"`
	Sqs              []iottopicrule.SqsState              `json:"sqs"`
	StepFunctions    []iottopicrule.StepFunctionsState    `json:"step_functions"`
	Timestream       []iottopicrule.TimestreamState       `json:"timestream"`
}
