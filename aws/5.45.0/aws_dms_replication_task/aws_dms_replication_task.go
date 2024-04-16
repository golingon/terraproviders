// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_dms_replication_task

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

// Resource represents the Terraform resource aws_dms_replication_task.
type Resource struct {
	Name      string
	Args      Args
	state     *awsDmsReplicationTaskState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adrt *Resource) Type() string {
	return "aws_dms_replication_task"
}

// LocalName returns the local name for [Resource].
func (adrt *Resource) LocalName() string {
	return adrt.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adrt *Resource) Configuration() interface{} {
	return adrt.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adrt *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adrt)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adrt *Resource) Dependencies() terra.Dependencies {
	return adrt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adrt *Resource) LifecycleManagement() *terra.Lifecycle {
	return adrt.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adrt *Resource) Attributes() awsDmsReplicationTaskAttributes {
	return awsDmsReplicationTaskAttributes{ref: terra.ReferenceResource(adrt)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adrt *Resource) ImportState(state io.Reader) error {
	adrt.state = &awsDmsReplicationTaskState{}
	if err := json.NewDecoder(state).Decode(adrt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adrt.Type(), adrt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adrt *Resource) State() (*awsDmsReplicationTaskState, bool) {
	return adrt.state, adrt.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adrt *Resource) StateMust() *awsDmsReplicationTaskState {
	if adrt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adrt.Type(), adrt.LocalName()))
	}
	return adrt.state
}

// Args contains the configurations for aws_dms_replication_task.
type Args struct {
	// CdcStartPosition: string, optional
	CdcStartPosition terra.StringValue `hcl:"cdc_start_position,attr"`
	// CdcStartTime: string, optional
	CdcStartTime terra.StringValue `hcl:"cdc_start_time,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MigrationType: string, required
	MigrationType terra.StringValue `hcl:"migration_type,attr" validate:"required"`
	// ReplicationInstanceArn: string, required
	ReplicationInstanceArn terra.StringValue `hcl:"replication_instance_arn,attr" validate:"required"`
	// ReplicationTaskId: string, required
	ReplicationTaskId terra.StringValue `hcl:"replication_task_id,attr" validate:"required"`
	// ReplicationTaskSettings: string, optional
	ReplicationTaskSettings terra.StringValue `hcl:"replication_task_settings,attr"`
	// SourceEndpointArn: string, required
	SourceEndpointArn terra.StringValue `hcl:"source_endpoint_arn,attr" validate:"required"`
	// StartReplicationTask: bool, optional
	StartReplicationTask terra.BoolValue `hcl:"start_replication_task,attr"`
	// TableMappings: string, required
	TableMappings terra.StringValue `hcl:"table_mappings,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TargetEndpointArn: string, required
	TargetEndpointArn terra.StringValue `hcl:"target_endpoint_arn,attr" validate:"required"`
}

type awsDmsReplicationTaskAttributes struct {
	ref terra.Reference
}

// CdcStartPosition returns a reference to field cdc_start_position of aws_dms_replication_task.
func (adrt awsDmsReplicationTaskAttributes) CdcStartPosition() terra.StringValue {
	return terra.ReferenceAsString(adrt.ref.Append("cdc_start_position"))
}

// CdcStartTime returns a reference to field cdc_start_time of aws_dms_replication_task.
func (adrt awsDmsReplicationTaskAttributes) CdcStartTime() terra.StringValue {
	return terra.ReferenceAsString(adrt.ref.Append("cdc_start_time"))
}

// Id returns a reference to field id of aws_dms_replication_task.
func (adrt awsDmsReplicationTaskAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adrt.ref.Append("id"))
}

// MigrationType returns a reference to field migration_type of aws_dms_replication_task.
func (adrt awsDmsReplicationTaskAttributes) MigrationType() terra.StringValue {
	return terra.ReferenceAsString(adrt.ref.Append("migration_type"))
}

// ReplicationInstanceArn returns a reference to field replication_instance_arn of aws_dms_replication_task.
func (adrt awsDmsReplicationTaskAttributes) ReplicationInstanceArn() terra.StringValue {
	return terra.ReferenceAsString(adrt.ref.Append("replication_instance_arn"))
}

// ReplicationTaskArn returns a reference to field replication_task_arn of aws_dms_replication_task.
func (adrt awsDmsReplicationTaskAttributes) ReplicationTaskArn() terra.StringValue {
	return terra.ReferenceAsString(adrt.ref.Append("replication_task_arn"))
}

// ReplicationTaskId returns a reference to field replication_task_id of aws_dms_replication_task.
func (adrt awsDmsReplicationTaskAttributes) ReplicationTaskId() terra.StringValue {
	return terra.ReferenceAsString(adrt.ref.Append("replication_task_id"))
}

// ReplicationTaskSettings returns a reference to field replication_task_settings of aws_dms_replication_task.
func (adrt awsDmsReplicationTaskAttributes) ReplicationTaskSettings() terra.StringValue {
	return terra.ReferenceAsString(adrt.ref.Append("replication_task_settings"))
}

// SourceEndpointArn returns a reference to field source_endpoint_arn of aws_dms_replication_task.
func (adrt awsDmsReplicationTaskAttributes) SourceEndpointArn() terra.StringValue {
	return terra.ReferenceAsString(adrt.ref.Append("source_endpoint_arn"))
}

// StartReplicationTask returns a reference to field start_replication_task of aws_dms_replication_task.
func (adrt awsDmsReplicationTaskAttributes) StartReplicationTask() terra.BoolValue {
	return terra.ReferenceAsBool(adrt.ref.Append("start_replication_task"))
}

// Status returns a reference to field status of aws_dms_replication_task.
func (adrt awsDmsReplicationTaskAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(adrt.ref.Append("status"))
}

// TableMappings returns a reference to field table_mappings of aws_dms_replication_task.
func (adrt awsDmsReplicationTaskAttributes) TableMappings() terra.StringValue {
	return terra.ReferenceAsString(adrt.ref.Append("table_mappings"))
}

// Tags returns a reference to field tags of aws_dms_replication_task.
func (adrt awsDmsReplicationTaskAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adrt.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_dms_replication_task.
func (adrt awsDmsReplicationTaskAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adrt.ref.Append("tags_all"))
}

// TargetEndpointArn returns a reference to field target_endpoint_arn of aws_dms_replication_task.
func (adrt awsDmsReplicationTaskAttributes) TargetEndpointArn() terra.StringValue {
	return terra.ReferenceAsString(adrt.ref.Append("target_endpoint_arn"))
}

type awsDmsReplicationTaskState struct {
	CdcStartPosition        string            `json:"cdc_start_position"`
	CdcStartTime            string            `json:"cdc_start_time"`
	Id                      string            `json:"id"`
	MigrationType           string            `json:"migration_type"`
	ReplicationInstanceArn  string            `json:"replication_instance_arn"`
	ReplicationTaskArn      string            `json:"replication_task_arn"`
	ReplicationTaskId       string            `json:"replication_task_id"`
	ReplicationTaskSettings string            `json:"replication_task_settings"`
	SourceEndpointArn       string            `json:"source_endpoint_arn"`
	StartReplicationTask    bool              `json:"start_replication_task"`
	Status                  string            `json:"status"`
	TableMappings           string            `json:"table_mappings"`
	Tags                    map[string]string `json:"tags"`
	TagsAll                 map[string]string `json:"tags_all"`
	TargetEndpointArn       string            `json:"target_endpoint_arn"`
}
