// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	rdscluster "github.com/golingon/terraproviders/aws/5.44.0/rdscluster"
	"io"
)

// NewRdsCluster creates a new instance of [RdsCluster].
func NewRdsCluster(name string, args RdsClusterArgs) *RdsCluster {
	return &RdsCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RdsCluster)(nil)

// RdsCluster represents the Terraform resource aws_rds_cluster.
type RdsCluster struct {
	Name      string
	Args      RdsClusterArgs
	state     *rdsClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RdsCluster].
func (rc *RdsCluster) Type() string {
	return "aws_rds_cluster"
}

// LocalName returns the local name for [RdsCluster].
func (rc *RdsCluster) LocalName() string {
	return rc.Name
}

// Configuration returns the configuration (args) for [RdsCluster].
func (rc *RdsCluster) Configuration() interface{} {
	return rc.Args
}

// DependOn is used for other resources to depend on [RdsCluster].
func (rc *RdsCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(rc)
}

// Dependencies returns the list of resources [RdsCluster] depends_on.
func (rc *RdsCluster) Dependencies() terra.Dependencies {
	return rc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RdsCluster].
func (rc *RdsCluster) LifecycleManagement() *terra.Lifecycle {
	return rc.Lifecycle
}

// Attributes returns the attributes for [RdsCluster].
func (rc *RdsCluster) Attributes() rdsClusterAttributes {
	return rdsClusterAttributes{ref: terra.ReferenceResource(rc)}
}

// ImportState imports the given attribute values into [RdsCluster]'s state.
func (rc *RdsCluster) ImportState(av io.Reader) error {
	rc.state = &rdsClusterState{}
	if err := json.NewDecoder(av).Decode(rc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rc.Type(), rc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RdsCluster] has state.
func (rc *RdsCluster) State() (*rdsClusterState, bool) {
	return rc.state, rc.state != nil
}

// StateMust returns the state for [RdsCluster]. Panics if the state is nil.
func (rc *RdsCluster) StateMust() *rdsClusterState {
	if rc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rc.Type(), rc.LocalName()))
	}
	return rc.state
}

// RdsClusterArgs contains the configurations for aws_rds_cluster.
type RdsClusterArgs struct {
	// AllocatedStorage: number, optional
	AllocatedStorage terra.NumberValue `hcl:"allocated_storage,attr"`
	// AllowMajorVersionUpgrade: bool, optional
	AllowMajorVersionUpgrade terra.BoolValue `hcl:"allow_major_version_upgrade,attr"`
	// ApplyImmediately: bool, optional
	ApplyImmediately terra.BoolValue `hcl:"apply_immediately,attr"`
	// AvailabilityZones: set of string, optional
	AvailabilityZones terra.SetValue[terra.StringValue] `hcl:"availability_zones,attr"`
	// BacktrackWindow: number, optional
	BacktrackWindow terra.NumberValue `hcl:"backtrack_window,attr"`
	// BackupRetentionPeriod: number, optional
	BackupRetentionPeriod terra.NumberValue `hcl:"backup_retention_period,attr"`
	// ClusterIdentifier: string, optional
	ClusterIdentifier terra.StringValue `hcl:"cluster_identifier,attr"`
	// ClusterIdentifierPrefix: string, optional
	ClusterIdentifierPrefix terra.StringValue `hcl:"cluster_identifier_prefix,attr"`
	// ClusterMembers: set of string, optional
	ClusterMembers terra.SetValue[terra.StringValue] `hcl:"cluster_members,attr"`
	// CopyTagsToSnapshot: bool, optional
	CopyTagsToSnapshot terra.BoolValue `hcl:"copy_tags_to_snapshot,attr"`
	// DatabaseName: string, optional
	DatabaseName terra.StringValue `hcl:"database_name,attr"`
	// DbClusterInstanceClass: string, optional
	DbClusterInstanceClass terra.StringValue `hcl:"db_cluster_instance_class,attr"`
	// DbClusterParameterGroupName: string, optional
	DbClusterParameterGroupName terra.StringValue `hcl:"db_cluster_parameter_group_name,attr"`
	// DbInstanceParameterGroupName: string, optional
	DbInstanceParameterGroupName terra.StringValue `hcl:"db_instance_parameter_group_name,attr"`
	// DbSubnetGroupName: string, optional
	DbSubnetGroupName terra.StringValue `hcl:"db_subnet_group_name,attr"`
	// DbSystemId: string, optional
	DbSystemId terra.StringValue `hcl:"db_system_id,attr"`
	// DeleteAutomatedBackups: bool, optional
	DeleteAutomatedBackups terra.BoolValue `hcl:"delete_automated_backups,attr"`
	// DeletionProtection: bool, optional
	DeletionProtection terra.BoolValue `hcl:"deletion_protection,attr"`
	// Domain: string, optional
	Domain terra.StringValue `hcl:"domain,attr"`
	// DomainIamRoleName: string, optional
	DomainIamRoleName terra.StringValue `hcl:"domain_iam_role_name,attr"`
	// EnableGlobalWriteForwarding: bool, optional
	EnableGlobalWriteForwarding terra.BoolValue `hcl:"enable_global_write_forwarding,attr"`
	// EnableHttpEndpoint: bool, optional
	EnableHttpEndpoint terra.BoolValue `hcl:"enable_http_endpoint,attr"`
	// EnableLocalWriteForwarding: bool, optional
	EnableLocalWriteForwarding terra.BoolValue `hcl:"enable_local_write_forwarding,attr"`
	// EnabledCloudwatchLogsExports: set of string, optional
	EnabledCloudwatchLogsExports terra.SetValue[terra.StringValue] `hcl:"enabled_cloudwatch_logs_exports,attr"`
	// Engine: string, required
	Engine terra.StringValue `hcl:"engine,attr" validate:"required"`
	// EngineMode: string, optional
	EngineMode terra.StringValue `hcl:"engine_mode,attr"`
	// EngineVersion: string, optional
	EngineVersion terra.StringValue `hcl:"engine_version,attr"`
	// FinalSnapshotIdentifier: string, optional
	FinalSnapshotIdentifier terra.StringValue `hcl:"final_snapshot_identifier,attr"`
	// GlobalClusterIdentifier: string, optional
	GlobalClusterIdentifier terra.StringValue `hcl:"global_cluster_identifier,attr"`
	// IamDatabaseAuthenticationEnabled: bool, optional
	IamDatabaseAuthenticationEnabled terra.BoolValue `hcl:"iam_database_authentication_enabled,attr"`
	// IamRoles: set of string, optional
	IamRoles terra.SetValue[terra.StringValue] `hcl:"iam_roles,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Iops: number, optional
	Iops terra.NumberValue `hcl:"iops,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// ManageMasterUserPassword: bool, optional
	ManageMasterUserPassword terra.BoolValue `hcl:"manage_master_user_password,attr"`
	// MasterPassword: string, optional
	MasterPassword terra.StringValue `hcl:"master_password,attr"`
	// MasterUserSecretKmsKeyId: string, optional
	MasterUserSecretKmsKeyId terra.StringValue `hcl:"master_user_secret_kms_key_id,attr"`
	// MasterUsername: string, optional
	MasterUsername terra.StringValue `hcl:"master_username,attr"`
	// NetworkType: string, optional
	NetworkType terra.StringValue `hcl:"network_type,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// PreferredBackupWindow: string, optional
	PreferredBackupWindow terra.StringValue `hcl:"preferred_backup_window,attr"`
	// PreferredMaintenanceWindow: string, optional
	PreferredMaintenanceWindow terra.StringValue `hcl:"preferred_maintenance_window,attr"`
	// ReplicationSourceIdentifier: string, optional
	ReplicationSourceIdentifier terra.StringValue `hcl:"replication_source_identifier,attr"`
	// SkipFinalSnapshot: bool, optional
	SkipFinalSnapshot terra.BoolValue `hcl:"skip_final_snapshot,attr"`
	// SnapshotIdentifier: string, optional
	SnapshotIdentifier terra.StringValue `hcl:"snapshot_identifier,attr"`
	// SourceRegion: string, optional
	SourceRegion terra.StringValue `hcl:"source_region,attr"`
	// StorageEncrypted: bool, optional
	StorageEncrypted terra.BoolValue `hcl:"storage_encrypted,attr"`
	// StorageType: string, optional
	StorageType terra.StringValue `hcl:"storage_type,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcSecurityGroupIds: set of string, optional
	VpcSecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"vpc_security_group_ids,attr"`
	// MasterUserSecret: min=0
	MasterUserSecret []rdscluster.MasterUserSecret `hcl:"master_user_secret,block" validate:"min=0"`
	// RestoreToPointInTime: optional
	RestoreToPointInTime *rdscluster.RestoreToPointInTime `hcl:"restore_to_point_in_time,block"`
	// S3Import: optional
	S3Import *rdscluster.S3Import `hcl:"s3_import,block"`
	// ScalingConfiguration: optional
	ScalingConfiguration *rdscluster.ScalingConfiguration `hcl:"scaling_configuration,block"`
	// Serverlessv2ScalingConfiguration: optional
	Serverlessv2ScalingConfiguration *rdscluster.Serverlessv2ScalingConfiguration `hcl:"serverlessv2_scaling_configuration,block"`
	// Timeouts: optional
	Timeouts *rdscluster.Timeouts `hcl:"timeouts,block"`
}
type rdsClusterAttributes struct {
	ref terra.Reference
}

// AllocatedStorage returns a reference to field allocated_storage of aws_rds_cluster.
func (rc rdsClusterAttributes) AllocatedStorage() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("allocated_storage"))
}

// AllowMajorVersionUpgrade returns a reference to field allow_major_version_upgrade of aws_rds_cluster.
func (rc rdsClusterAttributes) AllowMajorVersionUpgrade() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("allow_major_version_upgrade"))
}

// ApplyImmediately returns a reference to field apply_immediately of aws_rds_cluster.
func (rc rdsClusterAttributes) ApplyImmediately() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("apply_immediately"))
}

// Arn returns a reference to field arn of aws_rds_cluster.
func (rc rdsClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("arn"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_rds_cluster.
func (rc rdsClusterAttributes) AvailabilityZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rc.ref.Append("availability_zones"))
}

// BacktrackWindow returns a reference to field backtrack_window of aws_rds_cluster.
func (rc rdsClusterAttributes) BacktrackWindow() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("backtrack_window"))
}

// BackupRetentionPeriod returns a reference to field backup_retention_period of aws_rds_cluster.
func (rc rdsClusterAttributes) BackupRetentionPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("backup_retention_period"))
}

// ClusterIdentifier returns a reference to field cluster_identifier of aws_rds_cluster.
func (rc rdsClusterAttributes) ClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("cluster_identifier"))
}

// ClusterIdentifierPrefix returns a reference to field cluster_identifier_prefix of aws_rds_cluster.
func (rc rdsClusterAttributes) ClusterIdentifierPrefix() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("cluster_identifier_prefix"))
}

// ClusterMembers returns a reference to field cluster_members of aws_rds_cluster.
func (rc rdsClusterAttributes) ClusterMembers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rc.ref.Append("cluster_members"))
}

// ClusterResourceId returns a reference to field cluster_resource_id of aws_rds_cluster.
func (rc rdsClusterAttributes) ClusterResourceId() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("cluster_resource_id"))
}

// CopyTagsToSnapshot returns a reference to field copy_tags_to_snapshot of aws_rds_cluster.
func (rc rdsClusterAttributes) CopyTagsToSnapshot() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("copy_tags_to_snapshot"))
}

// DatabaseName returns a reference to field database_name of aws_rds_cluster.
func (rc rdsClusterAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("database_name"))
}

// DbClusterInstanceClass returns a reference to field db_cluster_instance_class of aws_rds_cluster.
func (rc rdsClusterAttributes) DbClusterInstanceClass() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("db_cluster_instance_class"))
}

// DbClusterParameterGroupName returns a reference to field db_cluster_parameter_group_name of aws_rds_cluster.
func (rc rdsClusterAttributes) DbClusterParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("db_cluster_parameter_group_name"))
}

// DbInstanceParameterGroupName returns a reference to field db_instance_parameter_group_name of aws_rds_cluster.
func (rc rdsClusterAttributes) DbInstanceParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("db_instance_parameter_group_name"))
}

// DbSubnetGroupName returns a reference to field db_subnet_group_name of aws_rds_cluster.
func (rc rdsClusterAttributes) DbSubnetGroupName() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("db_subnet_group_name"))
}

// DbSystemId returns a reference to field db_system_id of aws_rds_cluster.
func (rc rdsClusterAttributes) DbSystemId() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("db_system_id"))
}

// DeleteAutomatedBackups returns a reference to field delete_automated_backups of aws_rds_cluster.
func (rc rdsClusterAttributes) DeleteAutomatedBackups() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("delete_automated_backups"))
}

// DeletionProtection returns a reference to field deletion_protection of aws_rds_cluster.
func (rc rdsClusterAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("deletion_protection"))
}

// Domain returns a reference to field domain of aws_rds_cluster.
func (rc rdsClusterAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("domain"))
}

// DomainIamRoleName returns a reference to field domain_iam_role_name of aws_rds_cluster.
func (rc rdsClusterAttributes) DomainIamRoleName() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("domain_iam_role_name"))
}

// EnableGlobalWriteForwarding returns a reference to field enable_global_write_forwarding of aws_rds_cluster.
func (rc rdsClusterAttributes) EnableGlobalWriteForwarding() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("enable_global_write_forwarding"))
}

// EnableHttpEndpoint returns a reference to field enable_http_endpoint of aws_rds_cluster.
func (rc rdsClusterAttributes) EnableHttpEndpoint() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("enable_http_endpoint"))
}

// EnableLocalWriteForwarding returns a reference to field enable_local_write_forwarding of aws_rds_cluster.
func (rc rdsClusterAttributes) EnableLocalWriteForwarding() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("enable_local_write_forwarding"))
}

// EnabledCloudwatchLogsExports returns a reference to field enabled_cloudwatch_logs_exports of aws_rds_cluster.
func (rc rdsClusterAttributes) EnabledCloudwatchLogsExports() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rc.ref.Append("enabled_cloudwatch_logs_exports"))
}

// Endpoint returns a reference to field endpoint of aws_rds_cluster.
func (rc rdsClusterAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("endpoint"))
}

// Engine returns a reference to field engine of aws_rds_cluster.
func (rc rdsClusterAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("engine"))
}

// EngineMode returns a reference to field engine_mode of aws_rds_cluster.
func (rc rdsClusterAttributes) EngineMode() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("engine_mode"))
}

// EngineVersion returns a reference to field engine_version of aws_rds_cluster.
func (rc rdsClusterAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("engine_version"))
}

// EngineVersionActual returns a reference to field engine_version_actual of aws_rds_cluster.
func (rc rdsClusterAttributes) EngineVersionActual() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("engine_version_actual"))
}

// FinalSnapshotIdentifier returns a reference to field final_snapshot_identifier of aws_rds_cluster.
func (rc rdsClusterAttributes) FinalSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("final_snapshot_identifier"))
}

// GlobalClusterIdentifier returns a reference to field global_cluster_identifier of aws_rds_cluster.
func (rc rdsClusterAttributes) GlobalClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("global_cluster_identifier"))
}

// HostedZoneId returns a reference to field hosted_zone_id of aws_rds_cluster.
func (rc rdsClusterAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("hosted_zone_id"))
}

// IamDatabaseAuthenticationEnabled returns a reference to field iam_database_authentication_enabled of aws_rds_cluster.
func (rc rdsClusterAttributes) IamDatabaseAuthenticationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("iam_database_authentication_enabled"))
}

// IamRoles returns a reference to field iam_roles of aws_rds_cluster.
func (rc rdsClusterAttributes) IamRoles() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rc.ref.Append("iam_roles"))
}

// Id returns a reference to field id of aws_rds_cluster.
func (rc rdsClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("id"))
}

// Iops returns a reference to field iops of aws_rds_cluster.
func (rc rdsClusterAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("iops"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_rds_cluster.
func (rc rdsClusterAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("kms_key_id"))
}

// ManageMasterUserPassword returns a reference to field manage_master_user_password of aws_rds_cluster.
func (rc rdsClusterAttributes) ManageMasterUserPassword() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("manage_master_user_password"))
}

// MasterPassword returns a reference to field master_password of aws_rds_cluster.
func (rc rdsClusterAttributes) MasterPassword() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("master_password"))
}

// MasterUserSecretKmsKeyId returns a reference to field master_user_secret_kms_key_id of aws_rds_cluster.
func (rc rdsClusterAttributes) MasterUserSecretKmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("master_user_secret_kms_key_id"))
}

// MasterUsername returns a reference to field master_username of aws_rds_cluster.
func (rc rdsClusterAttributes) MasterUsername() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("master_username"))
}

// NetworkType returns a reference to field network_type of aws_rds_cluster.
func (rc rdsClusterAttributes) NetworkType() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("network_type"))
}

// Port returns a reference to field port of aws_rds_cluster.
func (rc rdsClusterAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("port"))
}

// PreferredBackupWindow returns a reference to field preferred_backup_window of aws_rds_cluster.
func (rc rdsClusterAttributes) PreferredBackupWindow() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("preferred_backup_window"))
}

// PreferredMaintenanceWindow returns a reference to field preferred_maintenance_window of aws_rds_cluster.
func (rc rdsClusterAttributes) PreferredMaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("preferred_maintenance_window"))
}

// ReaderEndpoint returns a reference to field reader_endpoint of aws_rds_cluster.
func (rc rdsClusterAttributes) ReaderEndpoint() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("reader_endpoint"))
}

// ReplicationSourceIdentifier returns a reference to field replication_source_identifier of aws_rds_cluster.
func (rc rdsClusterAttributes) ReplicationSourceIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("replication_source_identifier"))
}

// SkipFinalSnapshot returns a reference to field skip_final_snapshot of aws_rds_cluster.
func (rc rdsClusterAttributes) SkipFinalSnapshot() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("skip_final_snapshot"))
}

// SnapshotIdentifier returns a reference to field snapshot_identifier of aws_rds_cluster.
func (rc rdsClusterAttributes) SnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("snapshot_identifier"))
}

// SourceRegion returns a reference to field source_region of aws_rds_cluster.
func (rc rdsClusterAttributes) SourceRegion() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("source_region"))
}

// StorageEncrypted returns a reference to field storage_encrypted of aws_rds_cluster.
func (rc rdsClusterAttributes) StorageEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("storage_encrypted"))
}

// StorageType returns a reference to field storage_type of aws_rds_cluster.
func (rc rdsClusterAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("storage_type"))
}

// Tags returns a reference to field tags of aws_rds_cluster.
func (rc rdsClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_rds_cluster.
func (rc rdsClusterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rc.ref.Append("tags_all"))
}

// VpcSecurityGroupIds returns a reference to field vpc_security_group_ids of aws_rds_cluster.
func (rc rdsClusterAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rc.ref.Append("vpc_security_group_ids"))
}

func (rc rdsClusterAttributes) MasterUserSecret() terra.ListValue[rdscluster.MasterUserSecretAttributes] {
	return terra.ReferenceAsList[rdscluster.MasterUserSecretAttributes](rc.ref.Append("master_user_secret"))
}

func (rc rdsClusterAttributes) RestoreToPointInTime() terra.ListValue[rdscluster.RestoreToPointInTimeAttributes] {
	return terra.ReferenceAsList[rdscluster.RestoreToPointInTimeAttributes](rc.ref.Append("restore_to_point_in_time"))
}

func (rc rdsClusterAttributes) S3Import() terra.ListValue[rdscluster.S3ImportAttributes] {
	return terra.ReferenceAsList[rdscluster.S3ImportAttributes](rc.ref.Append("s3_import"))
}

func (rc rdsClusterAttributes) ScalingConfiguration() terra.ListValue[rdscluster.ScalingConfigurationAttributes] {
	return terra.ReferenceAsList[rdscluster.ScalingConfigurationAttributes](rc.ref.Append("scaling_configuration"))
}

func (rc rdsClusterAttributes) Serverlessv2ScalingConfiguration() terra.ListValue[rdscluster.Serverlessv2ScalingConfigurationAttributes] {
	return terra.ReferenceAsList[rdscluster.Serverlessv2ScalingConfigurationAttributes](rc.ref.Append("serverlessv2_scaling_configuration"))
}

func (rc rdsClusterAttributes) Timeouts() rdscluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[rdscluster.TimeoutsAttributes](rc.ref.Append("timeouts"))
}

type rdsClusterState struct {
	AllocatedStorage                 float64                                            `json:"allocated_storage"`
	AllowMajorVersionUpgrade         bool                                               `json:"allow_major_version_upgrade"`
	ApplyImmediately                 bool                                               `json:"apply_immediately"`
	Arn                              string                                             `json:"arn"`
	AvailabilityZones                []string                                           `json:"availability_zones"`
	BacktrackWindow                  float64                                            `json:"backtrack_window"`
	BackupRetentionPeriod            float64                                            `json:"backup_retention_period"`
	ClusterIdentifier                string                                             `json:"cluster_identifier"`
	ClusterIdentifierPrefix          string                                             `json:"cluster_identifier_prefix"`
	ClusterMembers                   []string                                           `json:"cluster_members"`
	ClusterResourceId                string                                             `json:"cluster_resource_id"`
	CopyTagsToSnapshot               bool                                               `json:"copy_tags_to_snapshot"`
	DatabaseName                     string                                             `json:"database_name"`
	DbClusterInstanceClass           string                                             `json:"db_cluster_instance_class"`
	DbClusterParameterGroupName      string                                             `json:"db_cluster_parameter_group_name"`
	DbInstanceParameterGroupName     string                                             `json:"db_instance_parameter_group_name"`
	DbSubnetGroupName                string                                             `json:"db_subnet_group_name"`
	DbSystemId                       string                                             `json:"db_system_id"`
	DeleteAutomatedBackups           bool                                               `json:"delete_automated_backups"`
	DeletionProtection               bool                                               `json:"deletion_protection"`
	Domain                           string                                             `json:"domain"`
	DomainIamRoleName                string                                             `json:"domain_iam_role_name"`
	EnableGlobalWriteForwarding      bool                                               `json:"enable_global_write_forwarding"`
	EnableHttpEndpoint               bool                                               `json:"enable_http_endpoint"`
	EnableLocalWriteForwarding       bool                                               `json:"enable_local_write_forwarding"`
	EnabledCloudwatchLogsExports     []string                                           `json:"enabled_cloudwatch_logs_exports"`
	Endpoint                         string                                             `json:"endpoint"`
	Engine                           string                                             `json:"engine"`
	EngineMode                       string                                             `json:"engine_mode"`
	EngineVersion                    string                                             `json:"engine_version"`
	EngineVersionActual              string                                             `json:"engine_version_actual"`
	FinalSnapshotIdentifier          string                                             `json:"final_snapshot_identifier"`
	GlobalClusterIdentifier          string                                             `json:"global_cluster_identifier"`
	HostedZoneId                     string                                             `json:"hosted_zone_id"`
	IamDatabaseAuthenticationEnabled bool                                               `json:"iam_database_authentication_enabled"`
	IamRoles                         []string                                           `json:"iam_roles"`
	Id                               string                                             `json:"id"`
	Iops                             float64                                            `json:"iops"`
	KmsKeyId                         string                                             `json:"kms_key_id"`
	ManageMasterUserPassword         bool                                               `json:"manage_master_user_password"`
	MasterPassword                   string                                             `json:"master_password"`
	MasterUserSecretKmsKeyId         string                                             `json:"master_user_secret_kms_key_id"`
	MasterUsername                   string                                             `json:"master_username"`
	NetworkType                      string                                             `json:"network_type"`
	Port                             float64                                            `json:"port"`
	PreferredBackupWindow            string                                             `json:"preferred_backup_window"`
	PreferredMaintenanceWindow       string                                             `json:"preferred_maintenance_window"`
	ReaderEndpoint                   string                                             `json:"reader_endpoint"`
	ReplicationSourceIdentifier      string                                             `json:"replication_source_identifier"`
	SkipFinalSnapshot                bool                                               `json:"skip_final_snapshot"`
	SnapshotIdentifier               string                                             `json:"snapshot_identifier"`
	SourceRegion                     string                                             `json:"source_region"`
	StorageEncrypted                 bool                                               `json:"storage_encrypted"`
	StorageType                      string                                             `json:"storage_type"`
	Tags                             map[string]string                                  `json:"tags"`
	TagsAll                          map[string]string                                  `json:"tags_all"`
	VpcSecurityGroupIds              []string                                           `json:"vpc_security_group_ids"`
	MasterUserSecret                 []rdscluster.MasterUserSecretState                 `json:"master_user_secret"`
	RestoreToPointInTime             []rdscluster.RestoreToPointInTimeState             `json:"restore_to_point_in_time"`
	S3Import                         []rdscluster.S3ImportState                         `json:"s3_import"`
	ScalingConfiguration             []rdscluster.ScalingConfigurationState             `json:"scaling_configuration"`
	Serverlessv2ScalingConfiguration []rdscluster.Serverlessv2ScalingConfigurationState `json:"serverlessv2_scaling_configuration"`
	Timeouts                         *rdscluster.TimeoutsState                          `json:"timeouts"`
}
