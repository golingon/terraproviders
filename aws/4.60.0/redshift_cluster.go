// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	redshiftcluster "github.com/golingon/terraproviders/aws/4.60.0/redshiftcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewRedshiftCluster(name string, args RedshiftClusterArgs) *RedshiftCluster {
	return &RedshiftCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedshiftCluster)(nil)

type RedshiftCluster struct {
	Name  string
	Args  RedshiftClusterArgs
	state *redshiftClusterState
}

func (rc *RedshiftCluster) Type() string {
	return "aws_redshift_cluster"
}

func (rc *RedshiftCluster) LocalName() string {
	return rc.Name
}

func (rc *RedshiftCluster) Configuration() interface{} {
	return rc.Args
}

func (rc *RedshiftCluster) Attributes() redshiftClusterAttributes {
	return redshiftClusterAttributes{ref: terra.ReferenceResource(rc)}
}

func (rc *RedshiftCluster) ImportState(av io.Reader) error {
	rc.state = &redshiftClusterState{}
	if err := json.NewDecoder(av).Decode(rc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rc.Type(), rc.LocalName(), err)
	}
	return nil
}

func (rc *RedshiftCluster) State() (*redshiftClusterState, bool) {
	return rc.state, rc.state != nil
}

func (rc *RedshiftCluster) StateMust() *redshiftClusterState {
	if rc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rc.Type(), rc.LocalName()))
	}
	return rc.state
}

func (rc *RedshiftCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(rc)
}

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
	// DependsOn contains resources that RedshiftCluster depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type redshiftClusterAttributes struct {
	ref terra.Reference
}

func (rc redshiftClusterAttributes) AllowVersionUpgrade() terra.BoolValue {
	return terra.ReferenceBool(rc.ref.Append("allow_version_upgrade"))
}

func (rc redshiftClusterAttributes) ApplyImmediately() terra.BoolValue {
	return terra.ReferenceBool(rc.ref.Append("apply_immediately"))
}

func (rc redshiftClusterAttributes) AquaConfigurationStatus() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("aqua_configuration_status"))
}

func (rc redshiftClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("arn"))
}

func (rc redshiftClusterAttributes) AutomatedSnapshotRetentionPeriod() terra.NumberValue {
	return terra.ReferenceNumber(rc.ref.Append("automated_snapshot_retention_period"))
}

func (rc redshiftClusterAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("availability_zone"))
}

func (rc redshiftClusterAttributes) AvailabilityZoneRelocationEnabled() terra.BoolValue {
	return terra.ReferenceBool(rc.ref.Append("availability_zone_relocation_enabled"))
}

func (rc redshiftClusterAttributes) ClusterIdentifier() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("cluster_identifier"))
}

func (rc redshiftClusterAttributes) ClusterParameterGroupName() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("cluster_parameter_group_name"))
}

func (rc redshiftClusterAttributes) ClusterPublicKey() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("cluster_public_key"))
}

func (rc redshiftClusterAttributes) ClusterRevisionNumber() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("cluster_revision_number"))
}

func (rc redshiftClusterAttributes) ClusterSecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](rc.ref.Append("cluster_security_groups"))
}

func (rc redshiftClusterAttributes) ClusterSubnetGroupName() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("cluster_subnet_group_name"))
}

func (rc redshiftClusterAttributes) ClusterType() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("cluster_type"))
}

func (rc redshiftClusterAttributes) ClusterVersion() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("cluster_version"))
}

func (rc redshiftClusterAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("database_name"))
}

func (rc redshiftClusterAttributes) DefaultIamRoleArn() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("default_iam_role_arn"))
}

func (rc redshiftClusterAttributes) DnsName() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("dns_name"))
}

func (rc redshiftClusterAttributes) ElasticIp() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("elastic_ip"))
}

func (rc redshiftClusterAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceBool(rc.ref.Append("encrypted"))
}

func (rc redshiftClusterAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("endpoint"))
}

func (rc redshiftClusterAttributes) EnhancedVpcRouting() terra.BoolValue {
	return terra.ReferenceBool(rc.ref.Append("enhanced_vpc_routing"))
}

func (rc redshiftClusterAttributes) FinalSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("final_snapshot_identifier"))
}

func (rc redshiftClusterAttributes) IamRoles() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](rc.ref.Append("iam_roles"))
}

func (rc redshiftClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("id"))
}

func (rc redshiftClusterAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("kms_key_id"))
}

func (rc redshiftClusterAttributes) MaintenanceTrackName() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("maintenance_track_name"))
}

func (rc redshiftClusterAttributes) ManualSnapshotRetentionPeriod() terra.NumberValue {
	return terra.ReferenceNumber(rc.ref.Append("manual_snapshot_retention_period"))
}

func (rc redshiftClusterAttributes) MasterPassword() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("master_password"))
}

func (rc redshiftClusterAttributes) MasterUsername() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("master_username"))
}

func (rc redshiftClusterAttributes) NodeType() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("node_type"))
}

func (rc redshiftClusterAttributes) NumberOfNodes() terra.NumberValue {
	return terra.ReferenceNumber(rc.ref.Append("number_of_nodes"))
}

func (rc redshiftClusterAttributes) OwnerAccount() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("owner_account"))
}

func (rc redshiftClusterAttributes) Port() terra.NumberValue {
	return terra.ReferenceNumber(rc.ref.Append("port"))
}

func (rc redshiftClusterAttributes) PreferredMaintenanceWindow() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("preferred_maintenance_window"))
}

func (rc redshiftClusterAttributes) PubliclyAccessible() terra.BoolValue {
	return terra.ReferenceBool(rc.ref.Append("publicly_accessible"))
}

func (rc redshiftClusterAttributes) SkipFinalSnapshot() terra.BoolValue {
	return terra.ReferenceBool(rc.ref.Append("skip_final_snapshot"))
}

func (rc redshiftClusterAttributes) SnapshotClusterIdentifier() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("snapshot_cluster_identifier"))
}

func (rc redshiftClusterAttributes) SnapshotIdentifier() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("snapshot_identifier"))
}

func (rc redshiftClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](rc.ref.Append("tags"))
}

func (rc redshiftClusterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](rc.ref.Append("tags_all"))
}

func (rc redshiftClusterAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](rc.ref.Append("vpc_security_group_ids"))
}

func (rc redshiftClusterAttributes) ClusterNodes() terra.ListValue[redshiftcluster.ClusterNodesAttributes] {
	return terra.ReferenceList[redshiftcluster.ClusterNodesAttributes](rc.ref.Append("cluster_nodes"))
}

func (rc redshiftClusterAttributes) Logging() terra.ListValue[redshiftcluster.LoggingAttributes] {
	return terra.ReferenceList[redshiftcluster.LoggingAttributes](rc.ref.Append("logging"))
}

func (rc redshiftClusterAttributes) SnapshotCopy() terra.ListValue[redshiftcluster.SnapshotCopyAttributes] {
	return terra.ReferenceList[redshiftcluster.SnapshotCopyAttributes](rc.ref.Append("snapshot_copy"))
}

func (rc redshiftClusterAttributes) Timeouts() redshiftcluster.TimeoutsAttributes {
	return terra.ReferenceSingle[redshiftcluster.TimeoutsAttributes](rc.ref.Append("timeouts"))
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
