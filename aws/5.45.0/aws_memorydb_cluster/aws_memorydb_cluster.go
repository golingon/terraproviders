// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_memorydb_cluster

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

// Resource represents the Terraform resource aws_memorydb_cluster.
type Resource struct {
	Name      string
	Args      Args
	state     *awsMemorydbClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (amc *Resource) Type() string {
	return "aws_memorydb_cluster"
}

// LocalName returns the local name for [Resource].
func (amc *Resource) LocalName() string {
	return amc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (amc *Resource) Configuration() interface{} {
	return amc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (amc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(amc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (amc *Resource) Dependencies() terra.Dependencies {
	return amc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (amc *Resource) LifecycleManagement() *terra.Lifecycle {
	return amc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (amc *Resource) Attributes() awsMemorydbClusterAttributes {
	return awsMemorydbClusterAttributes{ref: terra.ReferenceResource(amc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (amc *Resource) ImportState(state io.Reader) error {
	amc.state = &awsMemorydbClusterState{}
	if err := json.NewDecoder(state).Decode(amc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amc.Type(), amc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (amc *Resource) State() (*awsMemorydbClusterState, bool) {
	return amc.state, amc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (amc *Resource) StateMust() *awsMemorydbClusterState {
	if amc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amc.Type(), amc.LocalName()))
	}
	return amc.state
}

// Args contains the configurations for aws_memorydb_cluster.
type Args struct {
	// AclName: string, required
	AclName terra.StringValue `hcl:"acl_name,attr" validate:"required"`
	// AutoMinorVersionUpgrade: bool, optional
	AutoMinorVersionUpgrade terra.BoolValue `hcl:"auto_minor_version_upgrade,attr"`
	// DataTiering: bool, optional
	DataTiering terra.BoolValue `hcl:"data_tiering,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EngineVersion: string, optional
	EngineVersion terra.StringValue `hcl:"engine_version,attr"`
	// FinalSnapshotName: string, optional
	FinalSnapshotName terra.StringValue `hcl:"final_snapshot_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyArn: string, optional
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr"`
	// MaintenanceWindow: string, optional
	MaintenanceWindow terra.StringValue `hcl:"maintenance_window,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// NodeType: string, required
	NodeType terra.StringValue `hcl:"node_type,attr" validate:"required"`
	// NumReplicasPerShard: number, optional
	NumReplicasPerShard terra.NumberValue `hcl:"num_replicas_per_shard,attr"`
	// NumShards: number, optional
	NumShards terra.NumberValue `hcl:"num_shards,attr"`
	// ParameterGroupName: string, optional
	ParameterGroupName terra.StringValue `hcl:"parameter_group_name,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// SecurityGroupIds: set of string, optional
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// SnapshotArns: list of string, optional
	SnapshotArns terra.ListValue[terra.StringValue] `hcl:"snapshot_arns,attr"`
	// SnapshotName: string, optional
	SnapshotName terra.StringValue `hcl:"snapshot_name,attr"`
	// SnapshotRetentionLimit: number, optional
	SnapshotRetentionLimit terra.NumberValue `hcl:"snapshot_retention_limit,attr"`
	// SnapshotWindow: string, optional
	SnapshotWindow terra.StringValue `hcl:"snapshot_window,attr"`
	// SnsTopicArn: string, optional
	SnsTopicArn terra.StringValue `hcl:"sns_topic_arn,attr"`
	// SubnetGroupName: string, optional
	SubnetGroupName terra.StringValue `hcl:"subnet_group_name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TlsEnabled: bool, optional
	TlsEnabled terra.BoolValue `hcl:"tls_enabled,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsMemorydbClusterAttributes struct {
	ref terra.Reference
}

// AclName returns a reference to field acl_name of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) AclName() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("acl_name"))
}

// Arn returns a reference to field arn of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("arn"))
}

// AutoMinorVersionUpgrade returns a reference to field auto_minor_version_upgrade of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) AutoMinorVersionUpgrade() terra.BoolValue {
	return terra.ReferenceAsBool(amc.ref.Append("auto_minor_version_upgrade"))
}

// DataTiering returns a reference to field data_tiering of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) DataTiering() terra.BoolValue {
	return terra.ReferenceAsBool(amc.ref.Append("data_tiering"))
}

// Description returns a reference to field description of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("description"))
}

// EnginePatchVersion returns a reference to field engine_patch_version of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) EnginePatchVersion() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("engine_patch_version"))
}

// EngineVersion returns a reference to field engine_version of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("engine_version"))
}

// FinalSnapshotName returns a reference to field final_snapshot_name of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) FinalSnapshotName() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("final_snapshot_name"))
}

// Id returns a reference to field id of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("id"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("kms_key_arn"))
}

// MaintenanceWindow returns a reference to field maintenance_window of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) MaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("maintenance_window"))
}

// Name returns a reference to field name of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("name_prefix"))
}

// NodeType returns a reference to field node_type of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) NodeType() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("node_type"))
}

// NumReplicasPerShard returns a reference to field num_replicas_per_shard of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) NumReplicasPerShard() terra.NumberValue {
	return terra.ReferenceAsNumber(amc.ref.Append("num_replicas_per_shard"))
}

// NumShards returns a reference to field num_shards of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) NumShards() terra.NumberValue {
	return terra.ReferenceAsNumber(amc.ref.Append("num_shards"))
}

// ParameterGroupName returns a reference to field parameter_group_name of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) ParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("parameter_group_name"))
}

// Port returns a reference to field port of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(amc.ref.Append("port"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](amc.ref.Append("security_group_ids"))
}

// SnapshotArns returns a reference to field snapshot_arns of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) SnapshotArns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](amc.ref.Append("snapshot_arns"))
}

// SnapshotName returns a reference to field snapshot_name of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) SnapshotName() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("snapshot_name"))
}

// SnapshotRetentionLimit returns a reference to field snapshot_retention_limit of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) SnapshotRetentionLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(amc.ref.Append("snapshot_retention_limit"))
}

// SnapshotWindow returns a reference to field snapshot_window of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) SnapshotWindow() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("snapshot_window"))
}

// SnsTopicArn returns a reference to field sns_topic_arn of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) SnsTopicArn() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("sns_topic_arn"))
}

// SubnetGroupName returns a reference to field subnet_group_name of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) SubnetGroupName() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("subnet_group_name"))
}

// Tags returns a reference to field tags of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](amc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](amc.ref.Append("tags_all"))
}

// TlsEnabled returns a reference to field tls_enabled of aws_memorydb_cluster.
func (amc awsMemorydbClusterAttributes) TlsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(amc.ref.Append("tls_enabled"))
}

func (amc awsMemorydbClusterAttributes) ClusterEndpoint() terra.ListValue[ClusterEndpointAttributes] {
	return terra.ReferenceAsList[ClusterEndpointAttributes](amc.ref.Append("cluster_endpoint"))
}

func (amc awsMemorydbClusterAttributes) Shards() terra.SetValue[ShardsAttributes] {
	return terra.ReferenceAsSet[ShardsAttributes](amc.ref.Append("shards"))
}

func (amc awsMemorydbClusterAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](amc.ref.Append("timeouts"))
}

type awsMemorydbClusterState struct {
	AclName                 string                 `json:"acl_name"`
	Arn                     string                 `json:"arn"`
	AutoMinorVersionUpgrade bool                   `json:"auto_minor_version_upgrade"`
	DataTiering             bool                   `json:"data_tiering"`
	Description             string                 `json:"description"`
	EnginePatchVersion      string                 `json:"engine_patch_version"`
	EngineVersion           string                 `json:"engine_version"`
	FinalSnapshotName       string                 `json:"final_snapshot_name"`
	Id                      string                 `json:"id"`
	KmsKeyArn               string                 `json:"kms_key_arn"`
	MaintenanceWindow       string                 `json:"maintenance_window"`
	Name                    string                 `json:"name"`
	NamePrefix              string                 `json:"name_prefix"`
	NodeType                string                 `json:"node_type"`
	NumReplicasPerShard     float64                `json:"num_replicas_per_shard"`
	NumShards               float64                `json:"num_shards"`
	ParameterGroupName      string                 `json:"parameter_group_name"`
	Port                    float64                `json:"port"`
	SecurityGroupIds        []string               `json:"security_group_ids"`
	SnapshotArns            []string               `json:"snapshot_arns"`
	SnapshotName            string                 `json:"snapshot_name"`
	SnapshotRetentionLimit  float64                `json:"snapshot_retention_limit"`
	SnapshotWindow          string                 `json:"snapshot_window"`
	SnsTopicArn             string                 `json:"sns_topic_arn"`
	SubnetGroupName         string                 `json:"subnet_group_name"`
	Tags                    map[string]string      `json:"tags"`
	TagsAll                 map[string]string      `json:"tags_all"`
	TlsEnabled              bool                   `json:"tls_enabled"`
	ClusterEndpoint         []ClusterEndpointState `json:"cluster_endpoint"`
	Shards                  []ShardsState          `json:"shards"`
	Timeouts                *TimeoutsState         `json:"timeouts"`
}
