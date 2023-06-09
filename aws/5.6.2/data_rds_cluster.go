// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datardscluster "github.com/golingon/terraproviders/aws/5.6.2/datardscluster"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataRdsCluster creates a new instance of [DataRdsCluster].
func NewDataRdsCluster(name string, args DataRdsClusterArgs) *DataRdsCluster {
	return &DataRdsCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRdsCluster)(nil)

// DataRdsCluster represents the Terraform data resource aws_rds_cluster.
type DataRdsCluster struct {
	Name string
	Args DataRdsClusterArgs
}

// DataSource returns the Terraform object type for [DataRdsCluster].
func (rc *DataRdsCluster) DataSource() string {
	return "aws_rds_cluster"
}

// LocalName returns the local name for [DataRdsCluster].
func (rc *DataRdsCluster) LocalName() string {
	return rc.Name
}

// Configuration returns the configuration (args) for [DataRdsCluster].
func (rc *DataRdsCluster) Configuration() interface{} {
	return rc.Args
}

// Attributes returns the attributes for [DataRdsCluster].
func (rc *DataRdsCluster) Attributes() dataRdsClusterAttributes {
	return dataRdsClusterAttributes{ref: terra.ReferenceDataResource(rc)}
}

// DataRdsClusterArgs contains the configurations for aws_rds_cluster.
type DataRdsClusterArgs struct {
	// ClusterIdentifier: string, required
	ClusterIdentifier terra.StringValue `hcl:"cluster_identifier,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// MasterUserSecret: min=0
	MasterUserSecret []datardscluster.MasterUserSecret `hcl:"master_user_secret,block" validate:"min=0"`
}
type dataRdsClusterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_rds_cluster.
func (rc dataRdsClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("arn"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_rds_cluster.
func (rc dataRdsClusterAttributes) AvailabilityZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rc.ref.Append("availability_zones"))
}

// BacktrackWindow returns a reference to field backtrack_window of aws_rds_cluster.
func (rc dataRdsClusterAttributes) BacktrackWindow() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("backtrack_window"))
}

// BackupRetentionPeriod returns a reference to field backup_retention_period of aws_rds_cluster.
func (rc dataRdsClusterAttributes) BackupRetentionPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("backup_retention_period"))
}

// ClusterIdentifier returns a reference to field cluster_identifier of aws_rds_cluster.
func (rc dataRdsClusterAttributes) ClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("cluster_identifier"))
}

// ClusterMembers returns a reference to field cluster_members of aws_rds_cluster.
func (rc dataRdsClusterAttributes) ClusterMembers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rc.ref.Append("cluster_members"))
}

// ClusterResourceId returns a reference to field cluster_resource_id of aws_rds_cluster.
func (rc dataRdsClusterAttributes) ClusterResourceId() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("cluster_resource_id"))
}

// DatabaseName returns a reference to field database_name of aws_rds_cluster.
func (rc dataRdsClusterAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("database_name"))
}

// DbClusterParameterGroupName returns a reference to field db_cluster_parameter_group_name of aws_rds_cluster.
func (rc dataRdsClusterAttributes) DbClusterParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("db_cluster_parameter_group_name"))
}

// DbSubnetGroupName returns a reference to field db_subnet_group_name of aws_rds_cluster.
func (rc dataRdsClusterAttributes) DbSubnetGroupName() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("db_subnet_group_name"))
}

// EnabledCloudwatchLogsExports returns a reference to field enabled_cloudwatch_logs_exports of aws_rds_cluster.
func (rc dataRdsClusterAttributes) EnabledCloudwatchLogsExports() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rc.ref.Append("enabled_cloudwatch_logs_exports"))
}

// Endpoint returns a reference to field endpoint of aws_rds_cluster.
func (rc dataRdsClusterAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("endpoint"))
}

// Engine returns a reference to field engine of aws_rds_cluster.
func (rc dataRdsClusterAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("engine"))
}

// EngineMode returns a reference to field engine_mode of aws_rds_cluster.
func (rc dataRdsClusterAttributes) EngineMode() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("engine_mode"))
}

// EngineVersion returns a reference to field engine_version of aws_rds_cluster.
func (rc dataRdsClusterAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("engine_version"))
}

// FinalSnapshotIdentifier returns a reference to field final_snapshot_identifier of aws_rds_cluster.
func (rc dataRdsClusterAttributes) FinalSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("final_snapshot_identifier"))
}

// HostedZoneId returns a reference to field hosted_zone_id of aws_rds_cluster.
func (rc dataRdsClusterAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("hosted_zone_id"))
}

// IamDatabaseAuthenticationEnabled returns a reference to field iam_database_authentication_enabled of aws_rds_cluster.
func (rc dataRdsClusterAttributes) IamDatabaseAuthenticationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("iam_database_authentication_enabled"))
}

// IamRoles returns a reference to field iam_roles of aws_rds_cluster.
func (rc dataRdsClusterAttributes) IamRoles() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rc.ref.Append("iam_roles"))
}

// Id returns a reference to field id of aws_rds_cluster.
func (rc dataRdsClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_rds_cluster.
func (rc dataRdsClusterAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("kms_key_id"))
}

// MasterUsername returns a reference to field master_username of aws_rds_cluster.
func (rc dataRdsClusterAttributes) MasterUsername() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("master_username"))
}

// NetworkType returns a reference to field network_type of aws_rds_cluster.
func (rc dataRdsClusterAttributes) NetworkType() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("network_type"))
}

// Port returns a reference to field port of aws_rds_cluster.
func (rc dataRdsClusterAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("port"))
}

// PreferredBackupWindow returns a reference to field preferred_backup_window of aws_rds_cluster.
func (rc dataRdsClusterAttributes) PreferredBackupWindow() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("preferred_backup_window"))
}

// PreferredMaintenanceWindow returns a reference to field preferred_maintenance_window of aws_rds_cluster.
func (rc dataRdsClusterAttributes) PreferredMaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("preferred_maintenance_window"))
}

// ReaderEndpoint returns a reference to field reader_endpoint of aws_rds_cluster.
func (rc dataRdsClusterAttributes) ReaderEndpoint() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("reader_endpoint"))
}

// ReplicationSourceIdentifier returns a reference to field replication_source_identifier of aws_rds_cluster.
func (rc dataRdsClusterAttributes) ReplicationSourceIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("replication_source_identifier"))
}

// StorageEncrypted returns a reference to field storage_encrypted of aws_rds_cluster.
func (rc dataRdsClusterAttributes) StorageEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("storage_encrypted"))
}

// Tags returns a reference to field tags of aws_rds_cluster.
func (rc dataRdsClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rc.ref.Append("tags"))
}

// VpcSecurityGroupIds returns a reference to field vpc_security_group_ids of aws_rds_cluster.
func (rc dataRdsClusterAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rc.ref.Append("vpc_security_group_ids"))
}

func (rc dataRdsClusterAttributes) MasterUserSecret() terra.ListValue[datardscluster.MasterUserSecretAttributes] {
	return terra.ReferenceAsList[datardscluster.MasterUserSecretAttributes](rc.ref.Append("master_user_secret"))
}
