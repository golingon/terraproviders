// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	syntheticscanary "github.com/golingon/terraproviders/aws/5.12.0/syntheticscanary"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSyntheticsCanary creates a new instance of [SyntheticsCanary].
func NewSyntheticsCanary(name string, args SyntheticsCanaryArgs) *SyntheticsCanary {
	return &SyntheticsCanary{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SyntheticsCanary)(nil)

// SyntheticsCanary represents the Terraform resource aws_synthetics_canary.
type SyntheticsCanary struct {
	Name      string
	Args      SyntheticsCanaryArgs
	state     *syntheticsCanaryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SyntheticsCanary].
func (sc *SyntheticsCanary) Type() string {
	return "aws_synthetics_canary"
}

// LocalName returns the local name for [SyntheticsCanary].
func (sc *SyntheticsCanary) LocalName() string {
	return sc.Name
}

// Configuration returns the configuration (args) for [SyntheticsCanary].
func (sc *SyntheticsCanary) Configuration() interface{} {
	return sc.Args
}

// DependOn is used for other resources to depend on [SyntheticsCanary].
func (sc *SyntheticsCanary) DependOn() terra.Reference {
	return terra.ReferenceResource(sc)
}

// Dependencies returns the list of resources [SyntheticsCanary] depends_on.
func (sc *SyntheticsCanary) Dependencies() terra.Dependencies {
	return sc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SyntheticsCanary].
func (sc *SyntheticsCanary) LifecycleManagement() *terra.Lifecycle {
	return sc.Lifecycle
}

// Attributes returns the attributes for [SyntheticsCanary].
func (sc *SyntheticsCanary) Attributes() syntheticsCanaryAttributes {
	return syntheticsCanaryAttributes{ref: terra.ReferenceResource(sc)}
}

// ImportState imports the given attribute values into [SyntheticsCanary]'s state.
func (sc *SyntheticsCanary) ImportState(av io.Reader) error {
	sc.state = &syntheticsCanaryState{}
	if err := json.NewDecoder(av).Decode(sc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sc.Type(), sc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SyntheticsCanary] has state.
func (sc *SyntheticsCanary) State() (*syntheticsCanaryState, bool) {
	return sc.state, sc.state != nil
}

// StateMust returns the state for [SyntheticsCanary]. Panics if the state is nil.
func (sc *SyntheticsCanary) StateMust() *syntheticsCanaryState {
	if sc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sc.Type(), sc.LocalName()))
	}
	return sc.state
}

// SyntheticsCanaryArgs contains the configurations for aws_synthetics_canary.
type SyntheticsCanaryArgs struct {
	// ArtifactS3Location: string, required
	ArtifactS3Location terra.StringValue `hcl:"artifact_s3_location,attr" validate:"required"`
	// DeleteLambda: bool, optional
	DeleteLambda terra.BoolValue `hcl:"delete_lambda,attr"`
	// ExecutionRoleArn: string, required
	ExecutionRoleArn terra.StringValue `hcl:"execution_role_arn,attr" validate:"required"`
	// FailureRetentionPeriod: number, optional
	FailureRetentionPeriod terra.NumberValue `hcl:"failure_retention_period,attr"`
	// Handler: string, required
	Handler terra.StringValue `hcl:"handler,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RuntimeVersion: string, required
	RuntimeVersion terra.StringValue `hcl:"runtime_version,attr" validate:"required"`
	// S3Bucket: string, optional
	S3Bucket terra.StringValue `hcl:"s3_bucket,attr"`
	// S3Key: string, optional
	S3Key terra.StringValue `hcl:"s3_key,attr"`
	// S3Version: string, optional
	S3Version terra.StringValue `hcl:"s3_version,attr"`
	// StartCanary: bool, optional
	StartCanary terra.BoolValue `hcl:"start_canary,attr"`
	// SuccessRetentionPeriod: number, optional
	SuccessRetentionPeriod terra.NumberValue `hcl:"success_retention_period,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ZipFile: string, optional
	ZipFile terra.StringValue `hcl:"zip_file,attr"`
	// Timeline: min=0
	Timeline []syntheticscanary.Timeline `hcl:"timeline,block" validate:"min=0"`
	// ArtifactConfig: optional
	ArtifactConfig *syntheticscanary.ArtifactConfig `hcl:"artifact_config,block"`
	// RunConfig: optional
	RunConfig *syntheticscanary.RunConfig `hcl:"run_config,block"`
	// Schedule: required
	Schedule *syntheticscanary.Schedule `hcl:"schedule,block" validate:"required"`
	// VpcConfig: optional
	VpcConfig *syntheticscanary.VpcConfig `hcl:"vpc_config,block"`
}
type syntheticsCanaryAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("arn"))
}

// ArtifactS3Location returns a reference to field artifact_s3_location of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) ArtifactS3Location() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("artifact_s3_location"))
}

// DeleteLambda returns a reference to field delete_lambda of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) DeleteLambda() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("delete_lambda"))
}

// EngineArn returns a reference to field engine_arn of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) EngineArn() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("engine_arn"))
}

// ExecutionRoleArn returns a reference to field execution_role_arn of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) ExecutionRoleArn() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("execution_role_arn"))
}

// FailureRetentionPeriod returns a reference to field failure_retention_period of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) FailureRetentionPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("failure_retention_period"))
}

// Handler returns a reference to field handler of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) Handler() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("handler"))
}

// Id returns a reference to field id of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("id"))
}

// Name returns a reference to field name of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("name"))
}

// RuntimeVersion returns a reference to field runtime_version of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) RuntimeVersion() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("runtime_version"))
}

// S3Bucket returns a reference to field s3_bucket of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) S3Bucket() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("s3_bucket"))
}

// S3Key returns a reference to field s3_key of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) S3Key() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("s3_key"))
}

// S3Version returns a reference to field s3_version of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) S3Version() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("s3_version"))
}

// SourceLocationArn returns a reference to field source_location_arn of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) SourceLocationArn() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("source_location_arn"))
}

// StartCanary returns a reference to field start_canary of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) StartCanary() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("start_canary"))
}

// Status returns a reference to field status of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("status"))
}

// SuccessRetentionPeriod returns a reference to field success_retention_period of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) SuccessRetentionPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("success_retention_period"))
}

// Tags returns a reference to field tags of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sc.ref.Append("tags_all"))
}

// ZipFile returns a reference to field zip_file of aws_synthetics_canary.
func (sc syntheticsCanaryAttributes) ZipFile() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("zip_file"))
}

func (sc syntheticsCanaryAttributes) Timeline() terra.ListValue[syntheticscanary.TimelineAttributes] {
	return terra.ReferenceAsList[syntheticscanary.TimelineAttributes](sc.ref.Append("timeline"))
}

func (sc syntheticsCanaryAttributes) ArtifactConfig() terra.ListValue[syntheticscanary.ArtifactConfigAttributes] {
	return terra.ReferenceAsList[syntheticscanary.ArtifactConfigAttributes](sc.ref.Append("artifact_config"))
}

func (sc syntheticsCanaryAttributes) RunConfig() terra.ListValue[syntheticscanary.RunConfigAttributes] {
	return terra.ReferenceAsList[syntheticscanary.RunConfigAttributes](sc.ref.Append("run_config"))
}

func (sc syntheticsCanaryAttributes) Schedule() terra.ListValue[syntheticscanary.ScheduleAttributes] {
	return terra.ReferenceAsList[syntheticscanary.ScheduleAttributes](sc.ref.Append("schedule"))
}

func (sc syntheticsCanaryAttributes) VpcConfig() terra.ListValue[syntheticscanary.VpcConfigAttributes] {
	return terra.ReferenceAsList[syntheticscanary.VpcConfigAttributes](sc.ref.Append("vpc_config"))
}

type syntheticsCanaryState struct {
	Arn                    string                                 `json:"arn"`
	ArtifactS3Location     string                                 `json:"artifact_s3_location"`
	DeleteLambda           bool                                   `json:"delete_lambda"`
	EngineArn              string                                 `json:"engine_arn"`
	ExecutionRoleArn       string                                 `json:"execution_role_arn"`
	FailureRetentionPeriod float64                                `json:"failure_retention_period"`
	Handler                string                                 `json:"handler"`
	Id                     string                                 `json:"id"`
	Name                   string                                 `json:"name"`
	RuntimeVersion         string                                 `json:"runtime_version"`
	S3Bucket               string                                 `json:"s3_bucket"`
	S3Key                  string                                 `json:"s3_key"`
	S3Version              string                                 `json:"s3_version"`
	SourceLocationArn      string                                 `json:"source_location_arn"`
	StartCanary            bool                                   `json:"start_canary"`
	Status                 string                                 `json:"status"`
	SuccessRetentionPeriod float64                                `json:"success_retention_period"`
	Tags                   map[string]string                      `json:"tags"`
	TagsAll                map[string]string                      `json:"tags_all"`
	ZipFile                string                                 `json:"zip_file"`
	Timeline               []syntheticscanary.TimelineState       `json:"timeline"`
	ArtifactConfig         []syntheticscanary.ArtifactConfigState `json:"artifact_config"`
	RunConfig              []syntheticscanary.RunConfigState      `json:"run_config"`
	Schedule               []syntheticscanary.ScheduleState       `json:"schedule"`
	VpcConfig              []syntheticscanary.VpcConfigState      `json:"vpc_config"`
}
