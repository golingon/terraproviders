// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datamemorydbcluster "github.com/golingon/terraproviders/aws/5.12.0/datamemorydbcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMemorydbCluster creates a new instance of [DataMemorydbCluster].
func NewDataMemorydbCluster(name string, args DataMemorydbClusterArgs) *DataMemorydbCluster {
	return &DataMemorydbCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMemorydbCluster)(nil)

// DataMemorydbCluster represents the Terraform data resource aws_memorydb_cluster.
type DataMemorydbCluster struct {
	Name string
	Args DataMemorydbClusterArgs
}

// DataSource returns the Terraform object type for [DataMemorydbCluster].
func (mc *DataMemorydbCluster) DataSource() string {
	return "aws_memorydb_cluster"
}

// LocalName returns the local name for [DataMemorydbCluster].
func (mc *DataMemorydbCluster) LocalName() string {
	return mc.Name
}

// Configuration returns the configuration (args) for [DataMemorydbCluster].
func (mc *DataMemorydbCluster) Configuration() interface{} {
	return mc.Args
}

// Attributes returns the attributes for [DataMemorydbCluster].
func (mc *DataMemorydbCluster) Attributes() dataMemorydbClusterAttributes {
	return dataMemorydbClusterAttributes{ref: terra.ReferenceDataResource(mc)}
}

// DataMemorydbClusterArgs contains the configurations for aws_memorydb_cluster.
type DataMemorydbClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ClusterEndpoint: min=0
	ClusterEndpoint []datamemorydbcluster.ClusterEndpoint `hcl:"cluster_endpoint,block" validate:"min=0"`
	// Shards: min=0
	Shards []datamemorydbcluster.Shards `hcl:"shards,block" validate:"min=0"`
}
type dataMemorydbClusterAttributes struct {
	ref terra.Reference
}

// AclName returns a reference to field acl_name of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) AclName() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("acl_name"))
}

// Arn returns a reference to field arn of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("arn"))
}

// AutoMinorVersionUpgrade returns a reference to field auto_minor_version_upgrade of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) AutoMinorVersionUpgrade() terra.BoolValue {
	return terra.ReferenceAsBool(mc.ref.Append("auto_minor_version_upgrade"))
}

// DataTiering returns a reference to field data_tiering of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) DataTiering() terra.BoolValue {
	return terra.ReferenceAsBool(mc.ref.Append("data_tiering"))
}

// Description returns a reference to field description of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("description"))
}

// EnginePatchVersion returns a reference to field engine_patch_version of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) EnginePatchVersion() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("engine_patch_version"))
}

// EngineVersion returns a reference to field engine_version of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("engine_version"))
}

// FinalSnapshotName returns a reference to field final_snapshot_name of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) FinalSnapshotName() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("final_snapshot_name"))
}

// Id returns a reference to field id of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("id"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("kms_key_arn"))
}

// MaintenanceWindow returns a reference to field maintenance_window of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) MaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("maintenance_window"))
}

// Name returns a reference to field name of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("name"))
}

// NodeType returns a reference to field node_type of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) NodeType() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("node_type"))
}

// NumReplicasPerShard returns a reference to field num_replicas_per_shard of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) NumReplicasPerShard() terra.NumberValue {
	return terra.ReferenceAsNumber(mc.ref.Append("num_replicas_per_shard"))
}

// NumShards returns a reference to field num_shards of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) NumShards() terra.NumberValue {
	return terra.ReferenceAsNumber(mc.ref.Append("num_shards"))
}

// ParameterGroupName returns a reference to field parameter_group_name of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) ParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("parameter_group_name"))
}

// Port returns a reference to field port of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(mc.ref.Append("port"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mc.ref.Append("security_group_ids"))
}

// SnapshotRetentionLimit returns a reference to field snapshot_retention_limit of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) SnapshotRetentionLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(mc.ref.Append("snapshot_retention_limit"))
}

// SnapshotWindow returns a reference to field snapshot_window of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) SnapshotWindow() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("snapshot_window"))
}

// SnsTopicArn returns a reference to field sns_topic_arn of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) SnsTopicArn() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("sns_topic_arn"))
}

// SubnetGroupName returns a reference to field subnet_group_name of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) SubnetGroupName() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("subnet_group_name"))
}

// Tags returns a reference to field tags of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mc.ref.Append("tags"))
}

// TlsEnabled returns a reference to field tls_enabled of aws_memorydb_cluster.
func (mc dataMemorydbClusterAttributes) TlsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mc.ref.Append("tls_enabled"))
}

func (mc dataMemorydbClusterAttributes) ClusterEndpoint() terra.ListValue[datamemorydbcluster.ClusterEndpointAttributes] {
	return terra.ReferenceAsList[datamemorydbcluster.ClusterEndpointAttributes](mc.ref.Append("cluster_endpoint"))
}

func (mc dataMemorydbClusterAttributes) Shards() terra.SetValue[datamemorydbcluster.ShardsAttributes] {
	return terra.ReferenceAsSet[datamemorydbcluster.ShardsAttributes](mc.ref.Append("shards"))
}
