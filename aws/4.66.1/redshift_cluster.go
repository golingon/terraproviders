// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	redshiftcluster "github.com/golingon/terraproviders/aws/4.66.1/redshiftcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRedshiftCluster creates a new instance of [RedshiftCluster].
func NewRedshiftCluster(name string, args RedshiftClusterArgs) *RedshiftCluster {
	return &RedshiftCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedshiftCluster)(nil)

// RedshiftCluster represents the Terraform resource aws_redshift_cluster.
type RedshiftCluster struct {
	Name      string
	Args      RedshiftClusterArgs
	state     *redshiftClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedshiftCluster].
func (rc *RedshiftCluster) Type() string {
	return "aws_redshift_cluster"
}

// LocalName returns the local name for [RedshiftCluster].
func (rc *RedshiftCluster) LocalName() string {
	return rc.Name
}

// Configuration returns the configuration (args) for [RedshiftCluster].
func (rc *RedshiftCluster) Configuration() interface{} {
	return rc.Args
}

// DependOn is used for other resources to depend on [RedshiftCluster].
func (rc *RedshiftCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(rc)
}

// Dependencies returns the list of resources [RedshiftCluster] depends_on.
func (rc *RedshiftCluster) Dependencies() terra.Dependencies {
	return rc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedshiftCluster].
func (rc *RedshiftCluster) LifecycleManagement() *terra.Lifecycle {
	return rc.Lifecycle
}

// Attributes returns the attributes for [RedshiftCluster].
func (rc *RedshiftCluster) Attributes() redshiftClusterAttributes {
	return redshiftClusterAttributes{ref: terra.ReferenceResource(rc)}
}

// ImportState imports the given attribute values into [RedshiftCluster]'s state.
func (rc *RedshiftCluster) ImportState(av io.Reader) error {
	rc.state = &redshiftClusterState{}
	if err := json.NewDecoder(av).Decode(rc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rc.Type(), rc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedshiftCluster] has state.
func (rc *RedshiftCluster) State() (*redshiftClusterState, bool) {
	return rc.state, rc.state != nil
}

// StateMust returns the state for [RedshiftCluster]. Panics if the state is nil.
func (rc *RedshiftCluster) StateMust() *redshiftClusterState {
	if rc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rc.Type(), rc.LocalName()))
	}
	return rc.state
}

// RedshiftClusterArgs contains the configurations for aws_redshift_cluster.
type RedshiftClusterArgs struct {
	// AllowVersionUpgrade: bool, optional
	AllowVersionUpgrade terra.BoolValue `hcl:"allow_version_upgrade,attr"`
	// ApplyImmediately: bool, optional
	ApplyImmediately terra.BoolValue `hcl:"apply_immediately,attr"`
	// AquaConfigurationStatus: string, optional
	AquaConfigurationStatus terra.StringValue `hcl:"aqua_configuration_status,attr"`
	// AutomatedSnapshotRetentionPeriod: number, optional
	AutomatedSnapshotRetentionPeriod terra.NumberValue `hcl:"automated_snapshot_retention_period,attr"`
	// AvailabilityZone: string, optional
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr"`
	// AvailabilityZoneRelocationEnabled: bool, optional
	AvailabilityZoneRelocationEnabled terra.BoolValue `hcl:"availability_zone_relocation_enabled,attr"`
	// ClusterIdentifier: string, required
	ClusterIdentifier terra.StringValue `hcl:"cluster_identifier,attr" validate:"required"`
	// ClusterParameterGroupName: string, optional
	ClusterParameterGroupName terra.StringValue `hcl:"cluster_parameter_group_name,attr"`
	// ClusterPublicKey: string, optional
	ClusterPublicKey terra.StringValue `hcl:"cluster_public_key,attr"`
	// ClusterRevisionNumber: string, optional
	ClusterRevisionNumber terra.StringValue `hcl:"cluster_revision_number,attr"`
	// ClusterSecurityGroups: set of string, optional
	ClusterSecurityGroups terra.SetValue[terra.StringValue] `hcl:"cluster_security_groups,attr"`
	// ClusterSubnetGroupName: string, optional
	ClusterSubnetGroupName terra.StringValue `hcl:"cluster_subnet_group_name,attr"`
	// ClusterType: string, optional
	ClusterType terra.StringValue `hcl:"cluster_type,attr"`
	// ClusterVersion: string, optional
	ClusterVersion terra.StringValue `hcl:"cluster_version,attr"`
	// DatabaseName: string, optional
	DatabaseName terra.StringValue `hcl:"database_name,attr"`
	// DefaultIamRoleArn: string, optional
	DefaultIamRoleArn terra.StringValue `hcl:"default_iam_role_arn,attr"`
	// ElasticIp: string, optional
	ElasticIp terra.StringValue `hcl:"elastic_ip,attr"`
	// Encrypted: bool, optional
	Encrypted terra.BoolValue `hcl:"encrypted,attr"`
	// Endpoint: string, optional
	Endpoint terra.StringValue `hcl:"endpoint,attr"`
	// EnhancedVpcRouting: bool, optional
	EnhancedVpcRouting terra.BoolValue `hcl:"enhanced_vpc_routing,attr"`
	// FinalSnapshotIdentifier: string, optional
	FinalSnapshotIdentifier terra.StringValue `hcl:"final_snapshot_identifier,attr"`
	// IamRoles: set of string, optional
	IamRoles terra.SetValue[terra.StringValue] `hcl:"iam_roles,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// MaintenanceTrackName: string, optional
	MaintenanceTrackName terra.StringValue `hcl:"maintenance_track_name,attr"`
	// ManualSnapshotRetentionPeriod: number, optional
	ManualSnapshotRetentionPeriod terra.NumberValue `hcl:"manual_snapshot_retention_period,attr"`
	// MasterPassword: string, optional
	MasterPassword terra.StringValue `hcl:"master_password,attr"`
	// MasterUsername: string, optional
	MasterUsername terra.StringValue `hcl:"master_username,attr"`
	// NodeType: string, required
	NodeType terra.StringValue `hcl:"node_type,attr" validate:"required"`
	// NumberOfNodes: number, optional
	NumberOfNodes terra.NumberValue `hcl:"number_of_nodes,attr"`
	// OwnerAccount: string, optional
	OwnerAccount terra.StringValue `hcl:"owner_account,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// PreferredMaintenanceWindow: string, optional
	PreferredMaintenanceWindow terra.StringValue `hcl:"preferred_maintenance_window,attr"`
	// PubliclyAccessible: bool, optional
	PubliclyAccessible terra.BoolValue `hcl:"publicly_accessible,attr"`
	// SkipFinalSnapshot: bool, optional
	SkipFinalSnapshot terra.BoolValue `hcl:"skip_final_snapshot,attr"`
	// SnapshotClusterIdentifier: string, optional
	SnapshotClusterIdentifier terra.StringValue `hcl:"snapshot_cluster_identifier,attr"`
	// SnapshotIdentifier: string, optional
	SnapshotIdentifier terra.StringValue `hcl:"snapshot_identifier,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcSecurityGroupIds: set of string, optional
	VpcSecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"vpc_security_group_ids,attr"`
	// ClusterNodes: min=0
	ClusterNodes []redshiftcluster.ClusterNodes `hcl:"cluster_nodes,block" validate:"min=0"`
	// Logging: optional
	Logging *redshiftcluster.Logging `hcl:"logging,block"`
	// SnapshotCopy: optional
	SnapshotCopy *redshiftcluster.SnapshotCopy `hcl:"snapshot_copy,block"`
	// Timeouts: optional
	Timeouts *redshiftcluster.Timeouts `hcl:"timeouts,block"`
}
type redshiftClusterAttributes struct {
	ref terra.Reference
}

// AllowVersionUpgrade returns a reference to field allow_version_upgrade of aws_redshift_cluster.
func (rc redshiftClusterAttributes) AllowVersionUpgrade() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("allow_version_upgrade"))
}

// ApplyImmediately returns a reference to field apply_immediately of aws_redshift_cluster.
func (rc redshiftClusterAttributes) ApplyImmediately() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("apply_immediately"))
}

// AquaConfigurationStatus returns a reference to field aqua_configuration_status of aws_redshift_cluster.
func (rc redshiftClusterAttributes) AquaConfigurationStatus() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("aqua_configuration_status"))
}

// Arn returns a reference to field arn of aws_redshift_cluster.
func (rc redshiftClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("arn"))
}

// AutomatedSnapshotRetentionPeriod returns a reference to field automated_snapshot_retention_period of aws_redshift_cluster.
func (rc redshiftClusterAttributes) AutomatedSnapshotRetentionPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("automated_snapshot_retention_period"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_redshift_cluster.
func (rc redshiftClusterAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("availability_zone"))
}

// AvailabilityZoneRelocationEnabled returns a reference to field availability_zone_relocation_enabled of aws_redshift_cluster.
func (rc redshiftClusterAttributes) AvailabilityZoneRelocationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("availability_zone_relocation_enabled"))
}

// ClusterIdentifier returns a reference to field cluster_identifier of aws_redshift_cluster.
func (rc redshiftClusterAttributes) ClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("cluster_identifier"))
}

// ClusterParameterGroupName returns a reference to field cluster_parameter_group_name of aws_redshift_cluster.
func (rc redshiftClusterAttributes) ClusterParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("cluster_parameter_group_name"))
}

// ClusterPublicKey returns a reference to field cluster_public_key of aws_redshift_cluster.
func (rc redshiftClusterAttributes) ClusterPublicKey() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("cluster_public_key"))
}

// ClusterRevisionNumber returns a reference to field cluster_revision_number of aws_redshift_cluster.
func (rc redshiftClusterAttributes) ClusterRevisionNumber() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("cluster_revision_number"))
}

// ClusterSecurityGroups returns a reference to field cluster_security_groups of aws_redshift_cluster.
func (rc redshiftClusterAttributes) ClusterSecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rc.ref.Append("cluster_security_groups"))
}

// ClusterSubnetGroupName returns a reference to field cluster_subnet_group_name of aws_redshift_cluster.
func (rc redshiftClusterAttributes) ClusterSubnetGroupName() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("cluster_subnet_group_name"))
}

// ClusterType returns a reference to field cluster_type of aws_redshift_cluster.
func (rc redshiftClusterAttributes) ClusterType() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("cluster_type"))
}

// ClusterVersion returns a reference to field cluster_version of aws_redshift_cluster.
func (rc redshiftClusterAttributes) ClusterVersion() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("cluster_version"))
}

// DatabaseName returns a reference to field database_name of aws_redshift_cluster.
func (rc redshiftClusterAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("database_name"))
}

// DefaultIamRoleArn returns a reference to field default_iam_role_arn of aws_redshift_cluster.
func (rc redshiftClusterAttributes) DefaultIamRoleArn() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("default_iam_role_arn"))
}

// DnsName returns a reference to field dns_name of aws_redshift_cluster.
func (rc redshiftClusterAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("dns_name"))
}

// ElasticIp returns a reference to field elastic_ip of aws_redshift_cluster.
func (rc redshiftClusterAttributes) ElasticIp() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("elastic_ip"))
}

// Encrypted returns a reference to field encrypted of aws_redshift_cluster.
func (rc redshiftClusterAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("encrypted"))
}

// Endpoint returns a reference to field endpoint of aws_redshift_cluster.
func (rc redshiftClusterAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("endpoint"))
}

// EnhancedVpcRouting returns a reference to field enhanced_vpc_routing of aws_redshift_cluster.
func (rc redshiftClusterAttributes) EnhancedVpcRouting() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("enhanced_vpc_routing"))
}

// FinalSnapshotIdentifier returns a reference to field final_snapshot_identifier of aws_redshift_cluster.
func (rc redshiftClusterAttributes) FinalSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("final_snapshot_identifier"))
}

// IamRoles returns a reference to field iam_roles of aws_redshift_cluster.
func (rc redshiftClusterAttributes) IamRoles() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rc.ref.Append("iam_roles"))
}

// Id returns a reference to field id of aws_redshift_cluster.
func (rc redshiftClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_redshift_cluster.
func (rc redshiftClusterAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("kms_key_id"))
}

// MaintenanceTrackName returns a reference to field maintenance_track_name of aws_redshift_cluster.
func (rc redshiftClusterAttributes) MaintenanceTrackName() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("maintenance_track_name"))
}

// ManualSnapshotRetentionPeriod returns a reference to field manual_snapshot_retention_period of aws_redshift_cluster.
func (rc redshiftClusterAttributes) ManualSnapshotRetentionPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("manual_snapshot_retention_period"))
}

// MasterPassword returns a reference to field master_password of aws_redshift_cluster.
func (rc redshiftClusterAttributes) MasterPassword() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("master_password"))
}

// MasterUsername returns a reference to field master_username of aws_redshift_cluster.
func (rc redshiftClusterAttributes) MasterUsername() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("master_username"))
}

// NodeType returns a reference to field node_type of aws_redshift_cluster.
func (rc redshiftClusterAttributes) NodeType() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("node_type"))
}

// NumberOfNodes returns a reference to field number_of_nodes of aws_redshift_cluster.
func (rc redshiftClusterAttributes) NumberOfNodes() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("number_of_nodes"))
}

// OwnerAccount returns a reference to field owner_account of aws_redshift_cluster.
func (rc redshiftClusterAttributes) OwnerAccount() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("owner_account"))
}

// Port returns a reference to field port of aws_redshift_cluster.
func (rc redshiftClusterAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("port"))
}

// PreferredMaintenanceWindow returns a reference to field preferred_maintenance_window of aws_redshift_cluster.
func (rc redshiftClusterAttributes) PreferredMaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("preferred_maintenance_window"))
}

// PubliclyAccessible returns a reference to field publicly_accessible of aws_redshift_cluster.
func (rc redshiftClusterAttributes) PubliclyAccessible() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("publicly_accessible"))
}

// SkipFinalSnapshot returns a reference to field skip_final_snapshot of aws_redshift_cluster.
func (rc redshiftClusterAttributes) SkipFinalSnapshot() terra.BoolValue {
	return terra.ReferenceAsBool(rc.ref.Append("skip_final_snapshot"))
}

// SnapshotClusterIdentifier returns a reference to field snapshot_cluster_identifier of aws_redshift_cluster.
func (rc redshiftClusterAttributes) SnapshotClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("snapshot_cluster_identifier"))
}

// SnapshotIdentifier returns a reference to field snapshot_identifier of aws_redshift_cluster.
func (rc redshiftClusterAttributes) SnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("snapshot_identifier"))
}

// Tags returns a reference to field tags of aws_redshift_cluster.
func (rc redshiftClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_redshift_cluster.
func (rc redshiftClusterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rc.ref.Append("tags_all"))
}

// VpcSecurityGroupIds returns a reference to field vpc_security_group_ids of aws_redshift_cluster.
func (rc redshiftClusterAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rc.ref.Append("vpc_security_group_ids"))
}

func (rc redshiftClusterAttributes) ClusterNodes() terra.ListValue[redshiftcluster.ClusterNodesAttributes] {
	return terra.ReferenceAsList[redshiftcluster.ClusterNodesAttributes](rc.ref.Append("cluster_nodes"))
}

func (rc redshiftClusterAttributes) Logging() terra.ListValue[redshiftcluster.LoggingAttributes] {
	return terra.ReferenceAsList[redshiftcluster.LoggingAttributes](rc.ref.Append("logging"))
}

func (rc redshiftClusterAttributes) SnapshotCopy() terra.ListValue[redshiftcluster.SnapshotCopyAttributes] {
	return terra.ReferenceAsList[redshiftcluster.SnapshotCopyAttributes](rc.ref.Append("snapshot_copy"))
}

func (rc redshiftClusterAttributes) Timeouts() redshiftcluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[redshiftcluster.TimeoutsAttributes](rc.ref.Append("timeouts"))
}

type redshiftClusterState struct {
	AllowVersionUpgrade               bool                                `json:"allow_version_upgrade"`
	ApplyImmediately                  bool                                `json:"apply_immediately"`
	AquaConfigurationStatus           string                              `json:"aqua_configuration_status"`
	Arn                               string                              `json:"arn"`
	AutomatedSnapshotRetentionPeriod  float64                             `json:"automated_snapshot_retention_period"`
	AvailabilityZone                  string                              `json:"availability_zone"`
	AvailabilityZoneRelocationEnabled bool                                `json:"availability_zone_relocation_enabled"`
	ClusterIdentifier                 string                              `json:"cluster_identifier"`
	ClusterParameterGroupName         string                              `json:"cluster_parameter_group_name"`
	ClusterPublicKey                  string                              `json:"cluster_public_key"`
	ClusterRevisionNumber             string                              `json:"cluster_revision_number"`
	ClusterSecurityGroups             []string                            `json:"cluster_security_groups"`
	ClusterSubnetGroupName            string                              `json:"cluster_subnet_group_name"`
	ClusterType                       string                              `json:"cluster_type"`
	ClusterVersion                    string                              `json:"cluster_version"`
	DatabaseName                      string                              `json:"database_name"`
	DefaultIamRoleArn                 string                              `json:"default_iam_role_arn"`
	DnsName                           string                              `json:"dns_name"`
	ElasticIp                         string                              `json:"elastic_ip"`
	Encrypted                         bool                                `json:"encrypted"`
	Endpoint                          string                              `json:"endpoint"`
	EnhancedVpcRouting                bool                                `json:"enhanced_vpc_routing"`
	FinalSnapshotIdentifier           string                              `json:"final_snapshot_identifier"`
	IamRoles                          []string                            `json:"iam_roles"`
	Id                                string                              `json:"id"`
	KmsKeyId                          string                              `json:"kms_key_id"`
	MaintenanceTrackName              string                              `json:"maintenance_track_name"`
	ManualSnapshotRetentionPeriod     float64                             `json:"manual_snapshot_retention_period"`
	MasterPassword                    string                              `json:"master_password"`
	MasterUsername                    string                              `json:"master_username"`
	NodeType                          string                              `json:"node_type"`
	NumberOfNodes                     float64                             `json:"number_of_nodes"`
	OwnerAccount                      string                              `json:"owner_account"`
	Port                              float64                             `json:"port"`
	PreferredMaintenanceWindow        string                              `json:"preferred_maintenance_window"`
	PubliclyAccessible                bool                                `json:"publicly_accessible"`
	SkipFinalSnapshot                 bool                                `json:"skip_final_snapshot"`
	SnapshotClusterIdentifier         string                              `json:"snapshot_cluster_identifier"`
	SnapshotIdentifier                string                              `json:"snapshot_identifier"`
	Tags                              map[string]string                   `json:"tags"`
	TagsAll                           map[string]string                   `json:"tags_all"`
	VpcSecurityGroupIds               []string                            `json:"vpc_security_group_ids"`
	ClusterNodes                      []redshiftcluster.ClusterNodesState `json:"cluster_nodes"`
	Logging                           []redshiftcluster.LoggingState      `json:"logging"`
	SnapshotCopy                      []redshiftcluster.SnapshotCopyState `json:"snapshot_copy"`
	Timeouts                          *redshiftcluster.TimeoutsState      `json:"timeouts"`
}
