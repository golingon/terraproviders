// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_neptune_cluster

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

// Resource represents the Terraform resource aws_neptune_cluster.
type Resource struct {
	Name      string
	Args      Args
	state     *awsNeptuneClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (anc *Resource) Type() string {
	return "aws_neptune_cluster"
}

// LocalName returns the local name for [Resource].
func (anc *Resource) LocalName() string {
	return anc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (anc *Resource) Configuration() interface{} {
	return anc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (anc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(anc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (anc *Resource) Dependencies() terra.Dependencies {
	return anc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (anc *Resource) LifecycleManagement() *terra.Lifecycle {
	return anc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (anc *Resource) Attributes() awsNeptuneClusterAttributes {
	return awsNeptuneClusterAttributes{ref: terra.ReferenceResource(anc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (anc *Resource) ImportState(state io.Reader) error {
	anc.state = &awsNeptuneClusterState{}
	if err := json.NewDecoder(state).Decode(anc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", anc.Type(), anc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (anc *Resource) State() (*awsNeptuneClusterState, bool) {
	return anc.state, anc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (anc *Resource) StateMust() *awsNeptuneClusterState {
	if anc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", anc.Type(), anc.LocalName()))
	}
	return anc.state
}

// Args contains the configurations for aws_neptune_cluster.
type Args struct {
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
	// StorageType: string, optional
	StorageType terra.StringValue `hcl:"storage_type,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcSecurityGroupIds: set of string, optional
	VpcSecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"vpc_security_group_ids,attr"`
	// ServerlessV2ScalingConfiguration: optional
	ServerlessV2ScalingConfiguration *ServerlessV2ScalingConfiguration `hcl:"serverless_v2_scaling_configuration,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsNeptuneClusterAttributes struct {
	ref terra.Reference
}

// AllowMajorVersionUpgrade returns a reference to field allow_major_version_upgrade of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) AllowMajorVersionUpgrade() terra.BoolValue {
	return terra.ReferenceAsBool(anc.ref.Append("allow_major_version_upgrade"))
}

// ApplyImmediately returns a reference to field apply_immediately of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) ApplyImmediately() terra.BoolValue {
	return terra.ReferenceAsBool(anc.ref.Append("apply_immediately"))
}

// Arn returns a reference to field arn of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("arn"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) AvailabilityZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](anc.ref.Append("availability_zones"))
}

// BackupRetentionPeriod returns a reference to field backup_retention_period of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) BackupRetentionPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(anc.ref.Append("backup_retention_period"))
}

// ClusterIdentifier returns a reference to field cluster_identifier of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) ClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("cluster_identifier"))
}

// ClusterIdentifierPrefix returns a reference to field cluster_identifier_prefix of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) ClusterIdentifierPrefix() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("cluster_identifier_prefix"))
}

// ClusterMembers returns a reference to field cluster_members of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) ClusterMembers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](anc.ref.Append("cluster_members"))
}

// ClusterResourceId returns a reference to field cluster_resource_id of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) ClusterResourceId() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("cluster_resource_id"))
}

// CopyTagsToSnapshot returns a reference to field copy_tags_to_snapshot of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) CopyTagsToSnapshot() terra.BoolValue {
	return terra.ReferenceAsBool(anc.ref.Append("copy_tags_to_snapshot"))
}

// DeletionProtection returns a reference to field deletion_protection of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(anc.ref.Append("deletion_protection"))
}

// EnableCloudwatchLogsExports returns a reference to field enable_cloudwatch_logs_exports of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) EnableCloudwatchLogsExports() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](anc.ref.Append("enable_cloudwatch_logs_exports"))
}

// Endpoint returns a reference to field endpoint of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("endpoint"))
}

// Engine returns a reference to field engine of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("engine_version"))
}

// FinalSnapshotIdentifier returns a reference to field final_snapshot_identifier of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) FinalSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("final_snapshot_identifier"))
}

// GlobalClusterIdentifier returns a reference to field global_cluster_identifier of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) GlobalClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("global_cluster_identifier"))
}

// HostedZoneId returns a reference to field hosted_zone_id of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("hosted_zone_id"))
}

// IamDatabaseAuthenticationEnabled returns a reference to field iam_database_authentication_enabled of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) IamDatabaseAuthenticationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(anc.ref.Append("iam_database_authentication_enabled"))
}

// IamRoles returns a reference to field iam_roles of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) IamRoles() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](anc.ref.Append("iam_roles"))
}

// Id returns a reference to field id of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("id"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("kms_key_arn"))
}

// NeptuneClusterParameterGroupName returns a reference to field neptune_cluster_parameter_group_name of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) NeptuneClusterParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("neptune_cluster_parameter_group_name"))
}

// NeptuneInstanceParameterGroupName returns a reference to field neptune_instance_parameter_group_name of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) NeptuneInstanceParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("neptune_instance_parameter_group_name"))
}

// NeptuneSubnetGroupName returns a reference to field neptune_subnet_group_name of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) NeptuneSubnetGroupName() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("neptune_subnet_group_name"))
}

// Port returns a reference to field port of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(anc.ref.Append("port"))
}

// PreferredBackupWindow returns a reference to field preferred_backup_window of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) PreferredBackupWindow() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("preferred_backup_window"))
}

// PreferredMaintenanceWindow returns a reference to field preferred_maintenance_window of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) PreferredMaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("preferred_maintenance_window"))
}

// ReaderEndpoint returns a reference to field reader_endpoint of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) ReaderEndpoint() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("reader_endpoint"))
}

// ReplicationSourceIdentifier returns a reference to field replication_source_identifier of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) ReplicationSourceIdentifier() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("replication_source_identifier"))
}

// SkipFinalSnapshot returns a reference to field skip_final_snapshot of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) SkipFinalSnapshot() terra.BoolValue {
	return terra.ReferenceAsBool(anc.ref.Append("skip_final_snapshot"))
}

// SnapshotIdentifier returns a reference to field snapshot_identifier of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) SnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("snapshot_identifier"))
}

// StorageEncrypted returns a reference to field storage_encrypted of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) StorageEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(anc.ref.Append("storage_encrypted"))
}

// StorageType returns a reference to field storage_type of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("storage_type"))
}

// Tags returns a reference to field tags of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](anc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](anc.ref.Append("tags_all"))
}

// VpcSecurityGroupIds returns a reference to field vpc_security_group_ids of aws_neptune_cluster.
func (anc awsNeptuneClusterAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](anc.ref.Append("vpc_security_group_ids"))
}

func (anc awsNeptuneClusterAttributes) ServerlessV2ScalingConfiguration() terra.ListValue[ServerlessV2ScalingConfigurationAttributes] {
	return terra.ReferenceAsList[ServerlessV2ScalingConfigurationAttributes](anc.ref.Append("serverless_v2_scaling_configuration"))
}

func (anc awsNeptuneClusterAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](anc.ref.Append("timeouts"))
}

type awsNeptuneClusterState struct {
	AllowMajorVersionUpgrade          bool                                    `json:"allow_major_version_upgrade"`
	ApplyImmediately                  bool                                    `json:"apply_immediately"`
	Arn                               string                                  `json:"arn"`
	AvailabilityZones                 []string                                `json:"availability_zones"`
	BackupRetentionPeriod             float64                                 `json:"backup_retention_period"`
	ClusterIdentifier                 string                                  `json:"cluster_identifier"`
	ClusterIdentifierPrefix           string                                  `json:"cluster_identifier_prefix"`
	ClusterMembers                    []string                                `json:"cluster_members"`
	ClusterResourceId                 string                                  `json:"cluster_resource_id"`
	CopyTagsToSnapshot                bool                                    `json:"copy_tags_to_snapshot"`
	DeletionProtection                bool                                    `json:"deletion_protection"`
	EnableCloudwatchLogsExports       []string                                `json:"enable_cloudwatch_logs_exports"`
	Endpoint                          string                                  `json:"endpoint"`
	Engine                            string                                  `json:"engine"`
	EngineVersion                     string                                  `json:"engine_version"`
	FinalSnapshotIdentifier           string                                  `json:"final_snapshot_identifier"`
	GlobalClusterIdentifier           string                                  `json:"global_cluster_identifier"`
	HostedZoneId                      string                                  `json:"hosted_zone_id"`
	IamDatabaseAuthenticationEnabled  bool                                    `json:"iam_database_authentication_enabled"`
	IamRoles                          []string                                `json:"iam_roles"`
	Id                                string                                  `json:"id"`
	KmsKeyArn                         string                                  `json:"kms_key_arn"`
	NeptuneClusterParameterGroupName  string                                  `json:"neptune_cluster_parameter_group_name"`
	NeptuneInstanceParameterGroupName string                                  `json:"neptune_instance_parameter_group_name"`
	NeptuneSubnetGroupName            string                                  `json:"neptune_subnet_group_name"`
	Port                              float64                                 `json:"port"`
	PreferredBackupWindow             string                                  `json:"preferred_backup_window"`
	PreferredMaintenanceWindow        string                                  `json:"preferred_maintenance_window"`
	ReaderEndpoint                    string                                  `json:"reader_endpoint"`
	ReplicationSourceIdentifier       string                                  `json:"replication_source_identifier"`
	SkipFinalSnapshot                 bool                                    `json:"skip_final_snapshot"`
	SnapshotIdentifier                string                                  `json:"snapshot_identifier"`
	StorageEncrypted                  bool                                    `json:"storage_encrypted"`
	StorageType                       string                                  `json:"storage_type"`
	Tags                              map[string]string                       `json:"tags"`
	TagsAll                           map[string]string                       `json:"tags_all"`
	VpcSecurityGroupIds               []string                                `json:"vpc_security_group_ids"`
	ServerlessV2ScalingConfiguration  []ServerlessV2ScalingConfigurationState `json:"serverless_v2_scaling_configuration"`
	Timeouts                          *TimeoutsState                          `json:"timeouts"`
}
