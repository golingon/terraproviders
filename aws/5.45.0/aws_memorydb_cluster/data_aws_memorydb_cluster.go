// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_memorydb_cluster

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_memorydb_cluster.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (amc *DataSource) DataSource() string {
	return "aws_memorydb_cluster"
}

// LocalName returns the local name for [DataSource].
func (amc *DataSource) LocalName() string {
	return amc.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (amc *DataSource) Configuration() interface{} {
	return amc.Args
}

// Attributes returns the attributes for [DataSource].
func (amc *DataSource) Attributes() dataAwsMemorydbClusterAttributes {
	return dataAwsMemorydbClusterAttributes{ref: terra.ReferenceDataSource(amc)}
}

// DataArgs contains the configurations for aws_memorydb_cluster.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}

type dataAwsMemorydbClusterAttributes struct {
	ref terra.Reference
}

// AclName returns a reference to field acl_name of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) AclName() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("acl_name"))
}

// Arn returns a reference to field arn of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("arn"))
}

// AutoMinorVersionUpgrade returns a reference to field auto_minor_version_upgrade of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) AutoMinorVersionUpgrade() terra.BoolValue {
	return terra.ReferenceAsBool(amc.ref.Append("auto_minor_version_upgrade"))
}

// DataTiering returns a reference to field data_tiering of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) DataTiering() terra.BoolValue {
	return terra.ReferenceAsBool(amc.ref.Append("data_tiering"))
}

// Description returns a reference to field description of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("description"))
}

// EnginePatchVersion returns a reference to field engine_patch_version of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) EnginePatchVersion() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("engine_patch_version"))
}

// EngineVersion returns a reference to field engine_version of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("engine_version"))
}

// FinalSnapshotName returns a reference to field final_snapshot_name of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) FinalSnapshotName() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("final_snapshot_name"))
}

// Id returns a reference to field id of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("id"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("kms_key_arn"))
}

// MaintenanceWindow returns a reference to field maintenance_window of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) MaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("maintenance_window"))
}

// Name returns a reference to field name of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("name"))
}

// NodeType returns a reference to field node_type of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) NodeType() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("node_type"))
}

// NumReplicasPerShard returns a reference to field num_replicas_per_shard of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) NumReplicasPerShard() terra.NumberValue {
	return terra.ReferenceAsNumber(amc.ref.Append("num_replicas_per_shard"))
}

// NumShards returns a reference to field num_shards of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) NumShards() terra.NumberValue {
	return terra.ReferenceAsNumber(amc.ref.Append("num_shards"))
}

// ParameterGroupName returns a reference to field parameter_group_name of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) ParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("parameter_group_name"))
}

// Port returns a reference to field port of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(amc.ref.Append("port"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](amc.ref.Append("security_group_ids"))
}

// SnapshotRetentionLimit returns a reference to field snapshot_retention_limit of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) SnapshotRetentionLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(amc.ref.Append("snapshot_retention_limit"))
}

// SnapshotWindow returns a reference to field snapshot_window of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) SnapshotWindow() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("snapshot_window"))
}

// SnsTopicArn returns a reference to field sns_topic_arn of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) SnsTopicArn() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("sns_topic_arn"))
}

// SubnetGroupName returns a reference to field subnet_group_name of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) SubnetGroupName() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("subnet_group_name"))
}

// Tags returns a reference to field tags of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](amc.ref.Append("tags"))
}

// TlsEnabled returns a reference to field tls_enabled of aws_memorydb_cluster.
func (amc dataAwsMemorydbClusterAttributes) TlsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(amc.ref.Append("tls_enabled"))
}

func (amc dataAwsMemorydbClusterAttributes) ClusterEndpoint() terra.ListValue[DataClusterEndpointAttributes] {
	return terra.ReferenceAsList[DataClusterEndpointAttributes](amc.ref.Append("cluster_endpoint"))
}

func (amc dataAwsMemorydbClusterAttributes) Shards() terra.SetValue[DataShardsAttributes] {
	return terra.ReferenceAsSet[DataShardsAttributes](amc.ref.Append("shards"))
}
