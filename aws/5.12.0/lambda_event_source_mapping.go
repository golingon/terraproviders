// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lambdaeventsourcemapping "github.com/golingon/terraproviders/aws/5.12.0/lambdaeventsourcemapping"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLambdaEventSourceMapping creates a new instance of [LambdaEventSourceMapping].
func NewLambdaEventSourceMapping(name string, args LambdaEventSourceMappingArgs) *LambdaEventSourceMapping {
	return &LambdaEventSourceMapping{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LambdaEventSourceMapping)(nil)

// LambdaEventSourceMapping represents the Terraform resource aws_lambda_event_source_mapping.
type LambdaEventSourceMapping struct {
	Name      string
	Args      LambdaEventSourceMappingArgs
	state     *lambdaEventSourceMappingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LambdaEventSourceMapping].
func (lesm *LambdaEventSourceMapping) Type() string {
	return "aws_lambda_event_source_mapping"
}

// LocalName returns the local name for [LambdaEventSourceMapping].
func (lesm *LambdaEventSourceMapping) LocalName() string {
	return lesm.Name
}

// Configuration returns the configuration (args) for [LambdaEventSourceMapping].
func (lesm *LambdaEventSourceMapping) Configuration() interface{} {
	return lesm.Args
}

// DependOn is used for other resources to depend on [LambdaEventSourceMapping].
func (lesm *LambdaEventSourceMapping) DependOn() terra.Reference {
	return terra.ReferenceResource(lesm)
}

// Dependencies returns the list of resources [LambdaEventSourceMapping] depends_on.
func (lesm *LambdaEventSourceMapping) Dependencies() terra.Dependencies {
	return lesm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LambdaEventSourceMapping].
func (lesm *LambdaEventSourceMapping) LifecycleManagement() *terra.Lifecycle {
	return lesm.Lifecycle
}

// Attributes returns the attributes for [LambdaEventSourceMapping].
func (lesm *LambdaEventSourceMapping) Attributes() lambdaEventSourceMappingAttributes {
	return lambdaEventSourceMappingAttributes{ref: terra.ReferenceResource(lesm)}
}

// ImportState imports the given attribute values into [LambdaEventSourceMapping]'s state.
func (lesm *LambdaEventSourceMapping) ImportState(av io.Reader) error {
	lesm.state = &lambdaEventSourceMappingState{}
	if err := json.NewDecoder(av).Decode(lesm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lesm.Type(), lesm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LambdaEventSourceMapping] has state.
func (lesm *LambdaEventSourceMapping) State() (*lambdaEventSourceMappingState, bool) {
	return lesm.state, lesm.state != nil
}

// StateMust returns the state for [LambdaEventSourceMapping]. Panics if the state is nil.
func (lesm *LambdaEventSourceMapping) StateMust() *lambdaEventSourceMappingState {
	if lesm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lesm.Type(), lesm.LocalName()))
	}
	return lesm.state
}

// LambdaEventSourceMappingArgs contains the configurations for aws_lambda_event_source_mapping.
type LambdaEventSourceMappingArgs struct {
	// BatchSize: number, optional
	BatchSize terra.NumberValue `hcl:"batch_size,attr"`
	// BisectBatchOnFunctionError: bool, optional
	BisectBatchOnFunctionError terra.BoolValue `hcl:"bisect_batch_on_function_error,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// EventSourceArn: string, optional
	EventSourceArn terra.StringValue `hcl:"event_source_arn,attr"`
	// FunctionName: string, required
	FunctionName terra.StringValue `hcl:"function_name,attr" validate:"required"`
	// FunctionResponseTypes: set of string, optional
	FunctionResponseTypes terra.SetValue[terra.StringValue] `hcl:"function_response_types,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaximumBatchingWindowInSeconds: number, optional
	MaximumBatchingWindowInSeconds terra.NumberValue `hcl:"maximum_batching_window_in_seconds,attr"`
	// MaximumRecordAgeInSeconds: number, optional
	MaximumRecordAgeInSeconds terra.NumberValue `hcl:"maximum_record_age_in_seconds,attr"`
	// MaximumRetryAttempts: number, optional
	MaximumRetryAttempts terra.NumberValue `hcl:"maximum_retry_attempts,attr"`
	// ParallelizationFactor: number, optional
	ParallelizationFactor terra.NumberValue `hcl:"parallelization_factor,attr"`
	// Queues: list of string, optional
	Queues terra.ListValue[terra.StringValue] `hcl:"queues,attr"`
	// StartingPosition: string, optional
	StartingPosition terra.StringValue `hcl:"starting_position,attr"`
	// StartingPositionTimestamp: string, optional
	StartingPositionTimestamp terra.StringValue `hcl:"starting_position_timestamp,attr"`
	// Topics: set of string, optional
	Topics terra.SetValue[terra.StringValue] `hcl:"topics,attr"`
	// TumblingWindowInSeconds: number, optional
	TumblingWindowInSeconds terra.NumberValue `hcl:"tumbling_window_in_seconds,attr"`
	// AmazonManagedKafkaEventSourceConfig: optional
	AmazonManagedKafkaEventSourceConfig *lambdaeventsourcemapping.AmazonManagedKafkaEventSourceConfig `hcl:"amazon_managed_kafka_event_source_config,block"`
	// DestinationConfig: optional
	DestinationConfig *lambdaeventsourcemapping.DestinationConfig `hcl:"destination_config,block"`
	// DocumentDbEventSourceConfig: optional
	DocumentDbEventSourceConfig *lambdaeventsourcemapping.DocumentDbEventSourceConfig `hcl:"document_db_event_source_config,block"`
	// FilterCriteria: optional
	FilterCriteria *lambdaeventsourcemapping.FilterCriteria `hcl:"filter_criteria,block"`
	// ScalingConfig: optional
	ScalingConfig *lambdaeventsourcemapping.ScalingConfig `hcl:"scaling_config,block"`
	// SelfManagedEventSource: optional
	SelfManagedEventSource *lambdaeventsourcemapping.SelfManagedEventSource `hcl:"self_managed_event_source,block"`
	// SelfManagedKafkaEventSourceConfig: optional
	SelfManagedKafkaEventSourceConfig *lambdaeventsourcemapping.SelfManagedKafkaEventSourceConfig `hcl:"self_managed_kafka_event_source_config,block"`
	// SourceAccessConfiguration: min=0,max=22
	SourceAccessConfiguration []lambdaeventsourcemapping.SourceAccessConfiguration `hcl:"source_access_configuration,block" validate:"min=0,max=22"`
}
type lambdaEventSourceMappingAttributes struct {
	ref terra.Reference
}

// BatchSize returns a reference to field batch_size of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) BatchSize() terra.NumberValue {
	return terra.ReferenceAsNumber(lesm.ref.Append("batch_size"))
}

// BisectBatchOnFunctionError returns a reference to field bisect_batch_on_function_error of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) BisectBatchOnFunctionError() terra.BoolValue {
	return terra.ReferenceAsBool(lesm.ref.Append("bisect_batch_on_function_error"))
}

// Enabled returns a reference to field enabled of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(lesm.ref.Append("enabled"))
}

// EventSourceArn returns a reference to field event_source_arn of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) EventSourceArn() terra.StringValue {
	return terra.ReferenceAsString(lesm.ref.Append("event_source_arn"))
}

// FunctionArn returns a reference to field function_arn of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) FunctionArn() terra.StringValue {
	return terra.ReferenceAsString(lesm.ref.Append("function_arn"))
}

// FunctionName returns a reference to field function_name of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) FunctionName() terra.StringValue {
	return terra.ReferenceAsString(lesm.ref.Append("function_name"))
}

// FunctionResponseTypes returns a reference to field function_response_types of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) FunctionResponseTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](lesm.ref.Append("function_response_types"))
}

// Id returns a reference to field id of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lesm.ref.Append("id"))
}

// LastModified returns a reference to field last_modified of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) LastModified() terra.StringValue {
	return terra.ReferenceAsString(lesm.ref.Append("last_modified"))
}

// LastProcessingResult returns a reference to field last_processing_result of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) LastProcessingResult() terra.StringValue {
	return terra.ReferenceAsString(lesm.ref.Append("last_processing_result"))
}

// MaximumBatchingWindowInSeconds returns a reference to field maximum_batching_window_in_seconds of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) MaximumBatchingWindowInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(lesm.ref.Append("maximum_batching_window_in_seconds"))
}

// MaximumRecordAgeInSeconds returns a reference to field maximum_record_age_in_seconds of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) MaximumRecordAgeInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(lesm.ref.Append("maximum_record_age_in_seconds"))
}

// MaximumRetryAttempts returns a reference to field maximum_retry_attempts of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) MaximumRetryAttempts() terra.NumberValue {
	return terra.ReferenceAsNumber(lesm.ref.Append("maximum_retry_attempts"))
}

// ParallelizationFactor returns a reference to field parallelization_factor of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) ParallelizationFactor() terra.NumberValue {
	return terra.ReferenceAsNumber(lesm.ref.Append("parallelization_factor"))
}

// Queues returns a reference to field queues of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) Queues() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lesm.ref.Append("queues"))
}

// StartingPosition returns a reference to field starting_position of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) StartingPosition() terra.StringValue {
	return terra.ReferenceAsString(lesm.ref.Append("starting_position"))
}

// StartingPositionTimestamp returns a reference to field starting_position_timestamp of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) StartingPositionTimestamp() terra.StringValue {
	return terra.ReferenceAsString(lesm.ref.Append("starting_position_timestamp"))
}

// State returns a reference to field state of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(lesm.ref.Append("state"))
}

// StateTransitionReason returns a reference to field state_transition_reason of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) StateTransitionReason() terra.StringValue {
	return terra.ReferenceAsString(lesm.ref.Append("state_transition_reason"))
}

// Topics returns a reference to field topics of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) Topics() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](lesm.ref.Append("topics"))
}

// TumblingWindowInSeconds returns a reference to field tumbling_window_in_seconds of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) TumblingWindowInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(lesm.ref.Append("tumbling_window_in_seconds"))
}

// Uuid returns a reference to field uuid of aws_lambda_event_source_mapping.
func (lesm lambdaEventSourceMappingAttributes) Uuid() terra.StringValue {
	return terra.ReferenceAsString(lesm.ref.Append("uuid"))
}

func (lesm lambdaEventSourceMappingAttributes) AmazonManagedKafkaEventSourceConfig() terra.ListValue[lambdaeventsourcemapping.AmazonManagedKafkaEventSourceConfigAttributes] {
	return terra.ReferenceAsList[lambdaeventsourcemapping.AmazonManagedKafkaEventSourceConfigAttributes](lesm.ref.Append("amazon_managed_kafka_event_source_config"))
}

func (lesm lambdaEventSourceMappingAttributes) DestinationConfig() terra.ListValue[lambdaeventsourcemapping.DestinationConfigAttributes] {
	return terra.ReferenceAsList[lambdaeventsourcemapping.DestinationConfigAttributes](lesm.ref.Append("destination_config"))
}

func (lesm lambdaEventSourceMappingAttributes) DocumentDbEventSourceConfig() terra.ListValue[lambdaeventsourcemapping.DocumentDbEventSourceConfigAttributes] {
	return terra.ReferenceAsList[lambdaeventsourcemapping.DocumentDbEventSourceConfigAttributes](lesm.ref.Append("document_db_event_source_config"))
}

func (lesm lambdaEventSourceMappingAttributes) FilterCriteria() terra.ListValue[lambdaeventsourcemapping.FilterCriteriaAttributes] {
	return terra.ReferenceAsList[lambdaeventsourcemapping.FilterCriteriaAttributes](lesm.ref.Append("filter_criteria"))
}

func (lesm lambdaEventSourceMappingAttributes) ScalingConfig() terra.ListValue[lambdaeventsourcemapping.ScalingConfigAttributes] {
	return terra.ReferenceAsList[lambdaeventsourcemapping.ScalingConfigAttributes](lesm.ref.Append("scaling_config"))
}

func (lesm lambdaEventSourceMappingAttributes) SelfManagedEventSource() terra.ListValue[lambdaeventsourcemapping.SelfManagedEventSourceAttributes] {
	return terra.ReferenceAsList[lambdaeventsourcemapping.SelfManagedEventSourceAttributes](lesm.ref.Append("self_managed_event_source"))
}

func (lesm lambdaEventSourceMappingAttributes) SelfManagedKafkaEventSourceConfig() terra.ListValue[lambdaeventsourcemapping.SelfManagedKafkaEventSourceConfigAttributes] {
	return terra.ReferenceAsList[lambdaeventsourcemapping.SelfManagedKafkaEventSourceConfigAttributes](lesm.ref.Append("self_managed_kafka_event_source_config"))
}

func (lesm lambdaEventSourceMappingAttributes) SourceAccessConfiguration() terra.SetValue[lambdaeventsourcemapping.SourceAccessConfigurationAttributes] {
	return terra.ReferenceAsSet[lambdaeventsourcemapping.SourceAccessConfigurationAttributes](lesm.ref.Append("source_access_configuration"))
}

type lambdaEventSourceMappingState struct {
	BatchSize                           float64                                                             `json:"batch_size"`
	BisectBatchOnFunctionError          bool                                                                `json:"bisect_batch_on_function_error"`
	Enabled                             bool                                                                `json:"enabled"`
	EventSourceArn                      string                                                              `json:"event_source_arn"`
	FunctionArn                         string                                                              `json:"function_arn"`
	FunctionName                        string                                                              `json:"function_name"`
	FunctionResponseTypes               []string                                                            `json:"function_response_types"`
	Id                                  string                                                              `json:"id"`
	LastModified                        string                                                              `json:"last_modified"`
	LastProcessingResult                string                                                              `json:"last_processing_result"`
	MaximumBatchingWindowInSeconds      float64                                                             `json:"maximum_batching_window_in_seconds"`
	MaximumRecordAgeInSeconds           float64                                                             `json:"maximum_record_age_in_seconds"`
	MaximumRetryAttempts                float64                                                             `json:"maximum_retry_attempts"`
	ParallelizationFactor               float64                                                             `json:"parallelization_factor"`
	Queues                              []string                                                            `json:"queues"`
	StartingPosition                    string                                                              `json:"starting_position"`
	StartingPositionTimestamp           string                                                              `json:"starting_position_timestamp"`
	State                               string                                                              `json:"state"`
	StateTransitionReason               string                                                              `json:"state_transition_reason"`
	Topics                              []string                                                            `json:"topics"`
	TumblingWindowInSeconds             float64                                                             `json:"tumbling_window_in_seconds"`
	Uuid                                string                                                              `json:"uuid"`
	AmazonManagedKafkaEventSourceConfig []lambdaeventsourcemapping.AmazonManagedKafkaEventSourceConfigState `json:"amazon_managed_kafka_event_source_config"`
	DestinationConfig                   []lambdaeventsourcemapping.DestinationConfigState                   `json:"destination_config"`
	DocumentDbEventSourceConfig         []lambdaeventsourcemapping.DocumentDbEventSourceConfigState         `json:"document_db_event_source_config"`
	FilterCriteria                      []lambdaeventsourcemapping.FilterCriteriaState                      `json:"filter_criteria"`
	ScalingConfig                       []lambdaeventsourcemapping.ScalingConfigState                       `json:"scaling_config"`
	SelfManagedEventSource              []lambdaeventsourcemapping.SelfManagedEventSourceState              `json:"self_managed_event_source"`
	SelfManagedKafkaEventSourceConfig   []lambdaeventsourcemapping.SelfManagedKafkaEventSourceConfigState   `json:"self_managed_kafka_event_source_config"`
	SourceAccessConfiguration           []lambdaeventsourcemapping.SourceAccessConfigurationState           `json:"source_access_configuration"`
}
