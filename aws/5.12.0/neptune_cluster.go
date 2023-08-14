// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	neptunecluster "github.com/golingon/terraproviders/aws/5.12.0/neptunecluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNeptuneCluster creates a new instance of [NeptuneCluster].
func NewNeptuneCluster(name string, args NeptuneClusterArgs) *NeptuneCluster {
	return &NeptuneCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NeptuneCluster)(nil)

// NeptuneCluster represents the Terraform resource aws_neptune_cluster.
type NeptuneCluster struct {
	Name      string
	Args      NeptuneClusterArgs
	state     *neptuneClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NeptuneCluster].
func (nc *NeptuneCluster) Type() string {
	return "aws_neptune_cluster"
}

// LocalName returns the local name for [NeptuneCluster].
func (nc *NeptuneCluster) LocalName() string {
	return nc.Name
}

// Configuration returns the configuration (args) for [NeptuneCluster].
func (nc *NeptuneCluster) Configuration() interface{} {
	return nc.Args
}

// DependOn is used for other resources to depend on [NeptuneCluster].
func (nc *NeptuneCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(nc)
}

// Dependencies returns the list of resources [NeptuneCluster] depends_on.
func (nc *NeptuneCluster) Dependencies() terra.Dependencies {
	return nc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NeptuneCluster].
func (nc *NeptuneCluster) LifecycleManagement() *terra.Lifecycle {
	return nc.Lifecycle
}

// Attributes returns the attributes for [NeptuneCluster].
func (nc *NeptuneCluster) Attributes() neptuneClusterAttributes {
	return neptuneClusterAttributes{ref: terra.ReferenceResource(nc)}
}

// ImportState imports the given attribute values into [NeptuneCluster]'s state.
func (nc *NeptuneCluster) ImportState(av io.Reader) error {
	nc.state = &neptuneClusterState{}
	if err := json.NewDecoder(av).Decode(nc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nc.Type(), nc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NeptuneCluster] has state.
func (nc *NeptuneCluster) State() (*neptuneClusterState, bool) {
	return nc.state, nc.state != nil
}

// StateMust returns the state for [NeptuneCluster]. Panics if the state is nil.
func (nc *NeptuneCluster) StateMust() *neptuneClusterState {
	if nc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nc.Type(), nc.LocalName()))
	}
	return nc.state
}

// NeptuneClusterArgs contains the configurations for aws_neptune_cluster.
type NeptuneClusterArgs struct {
	// AllowMajorVersionUpgrade: bool, optional
	AllowMajorVersionUpgrade terra.BoolValue `hcl:"allow_major_version_upgrade,attr"`
	// ApplyImmediately: bool, optional
	ApplyImmediately terra.BoolValue `hcl:"apply_immediately,attr"`
	// AvailabilityZones: set of string, optional
	AvailabilityZones terra.SetValue[terra.StringValue] `hcl:"availability_zones,attr"`
	// BackupRetentionPeriod: number, optional
	BackupRetentionPeriod terra.NumberValue `hcl:"backup_retention_period,attr"`
	// ClusterIdentifier: string, optional
	ClusterIdentifier terra.StringValue `hcl:"cluster_identifier,attr"`
	// ClusterIdentifierPrefix: string, optional
	ClusterIdentifierPrefix terra.StringValue `hcl:"cluster_identifier_prefix,attr"`
	// CopyTagsToSnapshot: bool, optional
	CopyTagsToSnapshot terra.BoolValue `hcl:"copy_tags_to_snapshot,attr"`
	// DeletionProtection: bool, optional
	DeletionProtection terra.BoolValue `hcl:"deletion_protection,attr"`
	// EnableCloudwatchLogsExports: set of string, optional
	EnableCloudwatchLogsExports terra.SetValue[terra.StringValue] `hcl:"enable_cloudwatch_logs_exports,attr"`
	// Engine: string, optional
	Engine terra.StringValue `hcl:"engine,attr"`
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
	// KmsKeyArn: string, optional
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr"`
	// NeptuneClusterParameterGroupName: string, optional
	NeptuneClusterParameterGroupName terra.StringValue `hcl:"neptune_cluster_parameter_group_name,attr"`
	// NeptuneInstanceParameterGroupName: string, optional
	NeptuneInstanceParameterGroupName terra.StringValue `hcl:"neptune_instance_parameter_group_name,attr"`
	// NeptuneSubnetGroupName: string, optional
	NeptuneSubnetGroupName terra.StringValue `hcl:"neptune_subnet_group_name,attr"`
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
	// StorageEncrypted: bool, optional
	StorageEncrypted terra.BoolValue `hcl:"storage_encrypted,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcSecurityGroupIds: set of string, optional
	VpcSecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"vpc_security_group_ids,attr"`
	// ServerlessV2ScalingConfiguration: optional
	ServerlessV2ScalingConfiguration *neptunecluster.ServerlessV2ScalingConfiguration `hcl:"serverless_v2_scaling_configuration,block"`
	// Timeouts: optional
	Timeouts *neptunecluster.Timeouts `hcl:"timeouts,block"`
}
type neptuneClusterAttributes struct {
	ref terra.Reference
}

// AllowMajorVersionUpgrade returns a reference to field allow_major_version_upgrade of aws_neptune_cluster.
func (nc neptuneClusterAttributes) AllowMajorVersionUpgrade() terra.BoolValue {
	return terra.ReferenceAsBool(nc.ref.Append("allow_major_version_upgrade"))
}

// ApplyImmediately returns a reference to field apply_immediately of aws_neptune_cluster.
func (nc neptuneClusterAttributes) ApplyImmediately() terra.BoolValue {
	return terra.ReferenceAsBool(nc.ref.Append("apply_immediately"))
}

// Arn returns a reference to field arn of aws_neptune_cluster.
func (nc neptuneClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("arn"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_neptune_cluster.
func (nc neptuneClusterAttributes) AvailabilityZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nc.ref.Append("availability_zones"))
}

// BackupRetentionPeriod returns a reference to field backup_retention_period of aws_neptune_cluster.
func (nc neptuneClusterAttributes) BackupRetentionPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(nc.ref.Append("backup_retention_period"))
}

// ClusterIdentifier returns a reference to field cluster_identifier of aws_neptune_cluster.
func (nc neptuneClusterAttributes) ClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("cluster_identifier"))
}

// ClusterIdentifierPrefix returns a reference to field cluster_identifier_prefix of aws_neptune_cluster.
func (nc neptuneClusterAttributes) ClusterIdentifierPrefix() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("cluster_identifier_prefix"))
}

// ClusterMembers returns a reference to field cluster_members of aws_neptune_cluster.
func (nc neptuneClusterAttributes) ClusterMembers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nc.ref.Append("cluster_members"))
}

// ClusterResourceId returns a reference to field cluster_resource_id of aws_neptune_cluster.
func (nc neptuneClusterAttributes) ClusterResourceId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("cluster_resource_id"))
}

// CopyTagsToSnapshot returns a reference to field copy_tags_to_snapshot of aws_neptune_cluster.
func (nc neptuneClusterAttributes) CopyTagsToSnapshot() terra.BoolValue {
	return terra.ReferenceAsBool(nc.ref.Append("copy_tags_to_snapshot"))
}

// DeletionProtection returns a reference to field deletion_protection of aws_neptune_cluster.
func (nc neptuneClusterAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(nc.ref.Append("deletion_protection"))
}

// EnableCloudwatchLogsExports returns a reference to field enable_cloudwatch_logs_exports of aws_neptune_cluster.
func (nc neptuneClusterAttributes) EnableCloudwatchLogsExports() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nc.ref.Append("enable_cloudwatch_logs_exports"))
}

// Endpoint returns a reference to field endpoint of aws_neptune_cluster.
func (nc neptuneClusterAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("endpoint"))
}

// Engine returns a reference to field engine of aws_neptune_cluster.
func (nc neptuneClusterAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_neptune_cluster.
func (nc neptuneClusterAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("engine_version"))
}

// FinalSnapshotIdentifier returns a reference to field final_snapshot_identifier of aws_neptune_cluster.
func (nc neptuneClusterAttributes) FinalSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("final_snapshot_identifier"))
}

// GlobalClusterIdentifier returns a reference to field global_cluster_identifier of aws_neptune_cluster.
func (nc neptuneClusterAttributes) GlobalClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("global_cluster_identifier"))
}

// HostedZoneId returns a reference to field hosted_zone_id of aws_neptune_cluster.
func (nc neptuneClusterAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("hosted_zone_id"))
}

// IamDatabaseAuthenticationEnabled returns a reference to field iam_database_authentication_enabled of aws_neptune_cluster.
func (nc neptuneClusterAttributes) IamDatabaseAuthenticationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(nc.ref.Append("iam_database_authentication_enabled"))
}

// IamRoles returns a reference to field iam_roles of aws_neptune_cluster.
func (nc neptuneClusterAttributes) IamRoles() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nc.ref.Append("iam_roles"))
}

// Id returns a reference to field id of aws_neptune_cluster.
func (nc neptuneClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("id"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_neptune_cluster.
func (nc neptuneClusterAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("kms_key_arn"))
}

// NeptuneClusterParameterGroupName returns a reference to field neptune_cluster_parameter_group_name of aws_neptune_cluster.
func (nc neptuneClusterAttributes) NeptuneClusterParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("neptune_cluster_parameter_group_name"))
}

// NeptuneInstanceParameterGroupName returns a reference to field neptune_instance_parameter_group_name of aws_neptune_cluster.
func (nc neptuneClusterAttributes) NeptuneInstanceParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("neptune_instance_parameter_group_name"))
}

// NeptuneSubnetGroupName returns a reference to field neptune_subnet_group_name of aws_neptune_cluster.
func (nc neptuneClusterAttributes) NeptuneSubnetGroupName() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("neptune_subnet_group_name"))
}

// Port returns a reference to field port of aws_neptune_cluster.
func (nc neptuneClusterAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(nc.ref.Append("port"))
}

// PreferredBackupWindow returns a reference to field preferred_backup_window of aws_neptune_cluster.
func (nc neptuneClusterAttributes) PreferredBackupWindow() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("preferred_backup_window"))
}

// PreferredMaintenanceWindow returns a reference to field preferred_maintenance_window of aws_neptune_cluster.
func (nc neptuneClusterAttributes) PreferredMaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("preferred_maintenance_window"))
}

// ReaderEndpoint returns a reference to field reader_endpoint of aws_neptune_cluster.
func (nc neptuneClusterAttributes) ReaderEndpoint() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("reader_endpoint"))
}

// ReplicationSourceIdentifier returns a reference to field replication_source_identifier of aws_neptune_cluster.
func (nc neptuneClusterAttributes) ReplicationSourceIdentifier() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("replication_source_identifier"))
}

// SkipFinalSnapshot returns a reference to field skip_final_snapshot of aws_neptune_cluster.
func (nc neptuneClusterAttributes) SkipFinalSnapshot() terra.BoolValue {
	return terra.ReferenceAsBool(nc.ref.Append("skip_final_snapshot"))
}

// SnapshotIdentifier returns a reference to field snapshot_identifier of aws_neptune_cluster.
func (nc neptuneClusterAttributes) SnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("snapshot_identifier"))
}

// StorageEncrypted returns a reference to field storage_encrypted of aws_neptune_cluster.
func (nc neptuneClusterAttributes) StorageEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(nc.ref.Append("storage_encrypted"))
}

// Tags returns a reference to field tags of aws_neptune_cluster.
func (nc neptuneClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_neptune_cluster.
func (nc neptuneClusterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nc.ref.Append("tags_all"))
}

// VpcSecurityGroupIds returns a reference to field vpc_security_group_ids of aws_neptune_cluster.
func (nc neptuneClusterAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nc.ref.Append("vpc_security_group_ids"))
}

func (nc neptuneClusterAttributes) ServerlessV2ScalingConfiguration() terra.ListValue[neptunecluster.ServerlessV2ScalingConfigurationAttributes] {
	return terra.ReferenceAsList[neptunecluster.ServerlessV2ScalingConfigurationAttributes](nc.ref.Append("serverless_v2_scaling_configuration"))
}

func (nc neptuneClusterAttributes) Timeouts() neptunecluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[neptunecluster.TimeoutsAttributes](nc.ref.Append("timeouts"))
}

type neptuneClusterState struct {
	AllowMajorVersionUpgrade          bool                                                   `json:"allow_major_version_upgrade"`
	ApplyImmediately                  bool                                                   `json:"apply_immediately"`
	Arn                               string                                                 `json:"arn"`
	AvailabilityZones                 []string                                               `json:"availability_zones"`
	BackupRetentionPeriod             float64                                                `json:"backup_retention_period"`
	ClusterIdentifier                 string                                                 `json:"cluster_identifier"`
	ClusterIdentifierPrefix           string                                                 `json:"cluster_identifier_prefix"`
	ClusterMembers                    []string                                               `json:"cluster_members"`
	ClusterResourceId                 string                                                 `json:"cluster_resource_id"`
	CopyTagsToSnapshot                bool                                                   `json:"copy_tags_to_snapshot"`
	DeletionProtection                bool                                                   `json:"deletion_protection"`
	EnableCloudwatchLogsExports       []string                                               `json:"enable_cloudwatch_logs_exports"`
	Endpoint                          string                                                 `json:"endpoint"`
	Engine                            string                                                 `json:"engine"`
	EngineVersion                     string                                                 `json:"engine_version"`
	FinalSnapshotIdentifier           string                                                 `json:"final_snapshot_identifier"`
	GlobalClusterIdentifier           string                                                 `json:"global_cluster_identifier"`
	HostedZoneId                      string                                                 `json:"hosted_zone_id"`
	IamDatabaseAuthenticationEnabled  bool                                                   `json:"iam_database_authentication_enabled"`
	IamRoles                          []string                                               `json:"iam_roles"`
	Id                                string                                                 `json:"id"`
	KmsKeyArn                         string                                                 `json:"kms_key_arn"`
	NeptuneClusterParameterGroupName  string                                                 `json:"neptune_cluster_parameter_group_name"`
	NeptuneInstanceParameterGroupName string                                                 `json:"neptune_instance_parameter_group_name"`
	NeptuneSubnetGroupName            string                                                 `json:"neptune_subnet_group_name"`
	Port                              float64                                                `json:"port"`
	PreferredBackupWindow             string                                                 `json:"preferred_backup_window"`
	PreferredMaintenanceWindow        string                                                 `json:"preferred_maintenance_window"`
	ReaderEndpoint                    string                                                 `json:"reader_endpoint"`
	ReplicationSourceIdentifier       string                                                 `json:"replication_source_identifier"`
	SkipFinalSnapshot                 bool                                                   `json:"skip_final_snapshot"`
	SnapshotIdentifier                string                                                 `json:"snapshot_identifier"`
	StorageEncrypted                  bool                                                   `json:"storage_encrypted"`
	Tags                              map[string]string                                      `json:"tags"`
	TagsAll                           map[string]string                                      `json:"tags_all"`
	VpcSecurityGroupIds               []string                                               `json:"vpc_security_group_ids"`
	ServerlessV2ScalingConfiguration  []neptunecluster.ServerlessV2ScalingConfigurationState `json:"serverless_v2_scaling_configuration"`
	Timeouts                          *neptunecluster.TimeoutsState                          `json:"timeouts"`
}