// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	rdsexporttask "github.com/golingon/terraproviders/aws/5.10.0/rdsexporttask"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRdsExportTask creates a new instance of [RdsExportTask].
func NewRdsExportTask(name string, args RdsExportTaskArgs) *RdsExportTask {
	return &RdsExportTask{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RdsExportTask)(nil)

// RdsExportTask represents the Terraform resource aws_rds_export_task.
type RdsExportTask struct {
	Name      string
	Args      RdsExportTaskArgs
	state     *rdsExportTaskState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RdsExportTask].
func (ret *RdsExportTask) Type() string {
	return "aws_rds_export_task"
}

// LocalName returns the local name for [RdsExportTask].
func (ret *RdsExportTask) LocalName() string {
	return ret.Name
}

// Configuration returns the configuration (args) for [RdsExportTask].
func (ret *RdsExportTask) Configuration() interface{} {
	return ret.Args
}

// DependOn is used for other resources to depend on [RdsExportTask].
func (ret *RdsExportTask) DependOn() terra.Reference {
	return terra.ReferenceResource(ret)
}

// Dependencies returns the list of resources [RdsExportTask] depends_on.
func (ret *RdsExportTask) Dependencies() terra.Dependencies {
	return ret.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RdsExportTask].
func (ret *RdsExportTask) LifecycleManagement() *terra.Lifecycle {
	return ret.Lifecycle
}

// Attributes returns the attributes for [RdsExportTask].
func (ret *RdsExportTask) Attributes() rdsExportTaskAttributes {
	return rdsExportTaskAttributes{ref: terra.ReferenceResource(ret)}
}

// ImportState imports the given attribute values into [RdsExportTask]'s state.
func (ret *RdsExportTask) ImportState(av io.Reader) error {
	ret.state = &rdsExportTaskState{}
	if err := json.NewDecoder(av).Decode(ret.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ret.Type(), ret.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RdsExportTask] has state.
func (ret *RdsExportTask) State() (*rdsExportTaskState, bool) {
	return ret.state, ret.state != nil
}

// StateMust returns the state for [RdsExportTask]. Panics if the state is nil.
func (ret *RdsExportTask) StateMust() *rdsExportTaskState {
	if ret.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ret.Type(), ret.LocalName()))
	}
	return ret.state
}

// RdsExportTaskArgs contains the configurations for aws_rds_export_task.
type RdsExportTaskArgs struct {
	// ExportOnly: list of string, optional
	ExportOnly terra.ListValue[terra.StringValue] `hcl:"export_only,attr"`
	// ExportTaskIdentifier: string, required
	ExportTaskIdentifier terra.StringValue `hcl:"export_task_identifier,attr" validate:"required"`
	// IamRoleArn: string, required
	IamRoleArn terra.StringValue `hcl:"iam_role_arn,attr" validate:"required"`
	// KmsKeyId: string, required
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr" validate:"required"`
	// S3BucketName: string, required
	S3BucketName terra.StringValue `hcl:"s3_bucket_name,attr" validate:"required"`
	// S3Prefix: string, optional
	S3Prefix terra.StringValue `hcl:"s3_prefix,attr"`
	// SourceArn: string, required
	SourceArn terra.StringValue `hcl:"source_arn,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *rdsexporttask.Timeouts `hcl:"timeouts,block"`
}
type rdsExportTaskAttributes struct {
	ref terra.Reference
}

// ExportOnly returns a reference to field export_only of aws_rds_export_task.
func (ret rdsExportTaskAttributes) ExportOnly() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ret.ref.Append("export_only"))
}

// ExportTaskIdentifier returns a reference to field export_task_identifier of aws_rds_export_task.
func (ret rdsExportTaskAttributes) ExportTaskIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ret.ref.Append("export_task_identifier"))
}

// FailureCause returns a reference to field failure_cause of aws_rds_export_task.
func (ret rdsExportTaskAttributes) FailureCause() terra.StringValue {
	return terra.ReferenceAsString(ret.ref.Append("failure_cause"))
}

// IamRoleArn returns a reference to field iam_role_arn of aws_rds_export_task.
func (ret rdsExportTaskAttributes) IamRoleArn() terra.StringValue {
	return terra.ReferenceAsString(ret.ref.Append("iam_role_arn"))
}

// Id returns a reference to field id of aws_rds_export_task.
func (ret rdsExportTaskAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ret.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_rds_export_task.
func (ret rdsExportTaskAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ret.ref.Append("kms_key_id"))
}

// PercentProgress returns a reference to field percent_progress of aws_rds_export_task.
func (ret rdsExportTaskAttributes) PercentProgress() terra.NumberValue {
	return terra.ReferenceAsNumber(ret.ref.Append("percent_progress"))
}

// S3BucketName returns a reference to field s3_bucket_name of aws_rds_export_task.
func (ret rdsExportTaskAttributes) S3BucketName() terra.StringValue {
	return terra.ReferenceAsString(ret.ref.Append("s3_bucket_name"))
}

// S3Prefix returns a reference to field s3_prefix of aws_rds_export_task.
func (ret rdsExportTaskAttributes) S3Prefix() terra.StringValue {
	return terra.ReferenceAsString(ret.ref.Append("s3_prefix"))
}

// SnapshotTime returns a reference to field snapshot_time of aws_rds_export_task.
func (ret rdsExportTaskAttributes) SnapshotTime() terra.StringValue {
	return terra.ReferenceAsString(ret.ref.Append("snapshot_time"))
}

// SourceArn returns a reference to field source_arn of aws_rds_export_task.
func (ret rdsExportTaskAttributes) SourceArn() terra.StringValue {
	return terra.ReferenceAsString(ret.ref.Append("source_arn"))
}

// SourceType returns a reference to field source_type of aws_rds_export_task.
func (ret rdsExportTaskAttributes) SourceType() terra.StringValue {
	return terra.ReferenceAsString(ret.ref.Append("source_type"))
}

// Status returns a reference to field status of aws_rds_export_task.
func (ret rdsExportTaskAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ret.ref.Append("status"))
}

// TaskEndTime returns a reference to field task_end_time of aws_rds_export_task.
func (ret rdsExportTaskAttributes) TaskEndTime() terra.StringValue {
	return terra.ReferenceAsString(ret.ref.Append("task_end_time"))
}

// TaskStartTime returns a reference to field task_start_time of aws_rds_export_task.
func (ret rdsExportTaskAttributes) TaskStartTime() terra.StringValue {
	return terra.ReferenceAsString(ret.ref.Append("task_start_time"))
}

// WarningMessage returns a reference to field warning_message of aws_rds_export_task.
func (ret rdsExportTaskAttributes) WarningMessage() terra.StringValue {
	return terra.ReferenceAsString(ret.ref.Append("warning_message"))
}

func (ret rdsExportTaskAttributes) Timeouts() rdsexporttask.TimeoutsAttributes {
	return terra.ReferenceAsSingle[rdsexporttask.TimeoutsAttributes](ret.ref.Append("timeouts"))
}

type rdsExportTaskState struct {
	ExportOnly           []string                     `json:"export_only"`
	ExportTaskIdentifier string                       `json:"export_task_identifier"`
	FailureCause         string                       `json:"failure_cause"`
	IamRoleArn           string                       `json:"iam_role_arn"`
	Id                   string                       `json:"id"`
	KmsKeyId             string                       `json:"kms_key_id"`
	PercentProgress      float64                      `json:"percent_progress"`
	S3BucketName         string                       `json:"s3_bucket_name"`
	S3Prefix             string                       `json:"s3_prefix"`
	SnapshotTime         string                       `json:"snapshot_time"`
	SourceArn            string                       `json:"source_arn"`
	SourceType           string                       `json:"source_type"`
	Status               string                       `json:"status"`
	TaskEndTime          string                       `json:"task_end_time"`
	TaskStartTime        string                       `json:"task_start_time"`
	WarningMessage       string                       `json:"warning_message"`
	Timeouts             *rdsexporttask.TimeoutsState `json:"timeouts"`
}
