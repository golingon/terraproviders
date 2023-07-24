// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	elasticachereplicationgroup "github.com/golingon/terraproviders/aws/5.9.0/elasticachereplicationgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewElasticacheReplicationGroup creates a new instance of [ElasticacheReplicationGroup].
func NewElasticacheReplicationGroup(name string, args ElasticacheReplicationGroupArgs) *ElasticacheReplicationGroup {
	return &ElasticacheReplicationGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElasticacheReplicationGroup)(nil)

// ElasticacheReplicationGroup represents the Terraform resource aws_elasticache_replication_group.
type ElasticacheReplicationGroup struct {
	Name      string
	Args      ElasticacheReplicationGroupArgs
	state     *elasticacheReplicationGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ElasticacheReplicationGroup].
func (erg *ElasticacheReplicationGroup) Type() string {
	return "aws_elasticache_replication_group"
}

// LocalName returns the local name for [ElasticacheReplicationGroup].
func (erg *ElasticacheReplicationGroup) LocalName() string {
	return erg.Name
}

// Configuration returns the configuration (args) for [ElasticacheReplicationGroup].
func (erg *ElasticacheReplicationGroup) Configuration() interface{} {
	return erg.Args
}

// DependOn is used for other resources to depend on [ElasticacheReplicationGroup].
func (erg *ElasticacheReplicationGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(erg)
}

// Dependencies returns the list of resources [ElasticacheReplicationGroup] depends_on.
func (erg *ElasticacheReplicationGroup) Dependencies() terra.Dependencies {
	return erg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ElasticacheReplicationGroup].
func (erg *ElasticacheReplicationGroup) LifecycleManagement() *terra.Lifecycle {
	return erg.Lifecycle
}

// Attributes returns the attributes for [ElasticacheReplicationGroup].
func (erg *ElasticacheReplicationGroup) Attributes() elasticacheReplicationGroupAttributes {
	return elasticacheReplicationGroupAttributes{ref: terra.ReferenceResource(erg)}
}

// ImportState imports the given attribute values into [ElasticacheReplicationGroup]'s state.
func (erg *ElasticacheReplicationGroup) ImportState(av io.Reader) error {
	erg.state = &elasticacheReplicationGroupState{}
	if err := json.NewDecoder(av).Decode(erg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", erg.Type(), erg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ElasticacheReplicationGroup] has state.
func (erg *ElasticacheReplicationGroup) State() (*elasticacheReplicationGroupState, bool) {
	return erg.state, erg.state != nil
}

// StateMust returns the state for [ElasticacheReplicationGroup]. Panics if the state is nil.
func (erg *ElasticacheReplicationGroup) StateMust() *elasticacheReplicationGroupState {
	if erg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", erg.Type(), erg.LocalName()))
	}
	return erg.state
}

// ElasticacheReplicationGroupArgs contains the configurations for aws_elasticache_replication_group.
type ElasticacheReplicationGroupArgs struct {
	// ApplyImmediately: bool, optional
	ApplyImmediately terra.BoolValue `hcl:"apply_immediately,attr"`
	// AtRestEncryptionEnabled: bool, optional
	AtRestEncryptionEnabled terra.BoolValue `hcl:"at_rest_encryption_enabled,attr"`
	// AuthToken: string, optional
	AuthToken terra.StringValue `hcl:"auth_token,attr"`
	// AutoMinorVersionUpgrade: string, optional
	AutoMinorVersionUpgrade terra.StringValue `hcl:"auto_minor_version_upgrade,attr"`
	// AutomaticFailoverEnabled: bool, optional
	AutomaticFailoverEnabled terra.BoolValue `hcl:"automatic_failover_enabled,attr"`
	// DataTieringEnabled: bool, optional
	DataTieringEnabled terra.BoolValue `hcl:"data_tiering_enabled,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Engine: string, optional
	Engine terra.StringValue `hcl:"engine,attr"`
	// EngineVersion: string, optional
	EngineVersion terra.StringValue `hcl:"engine_version,attr"`
	// FinalSnapshotIdentifier: string, optional
	FinalSnapshotIdentifier terra.StringValue `hcl:"final_snapshot_identifier,attr"`
	// GlobalReplicationGroupId: string, optional
	GlobalReplicationGroupId terra.StringValue `hcl:"global_replication_group_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// MaintenanceWindow: string, optional
	MaintenanceWindow terra.StringValue `hcl:"maintenance_window,attr"`
	// MultiAzEnabled: bool, optional
	MultiAzEnabled terra.BoolValue `hcl:"multi_az_enabled,attr"`
	// NodeType: string, optional
	NodeType terra.StringValue `hcl:"node_type,attr"`
	// NotificationTopicArn: string, optional
	NotificationTopicArn terra.StringValue `hcl:"notification_topic_arn,attr"`
	// NumCacheClusters: number, optional
	NumCacheClusters terra.NumberValue `hcl:"num_cache_clusters,attr"`
	// NumNodeGroups: number, optional
	NumNodeGroups terra.NumberValue `hcl:"num_node_groups,attr"`
	// ParameterGroupName: string, optional
	ParameterGroupName terra.StringValue `hcl:"parameter_group_name,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// PreferredCacheClusterAzs: list of string, optional
	PreferredCacheClusterAzs terra.ListValue[terra.StringValue] `hcl:"preferred_cache_cluster_azs,attr"`
	// ReplicasPerNodeGroup: number, optional
	ReplicasPerNodeGroup terra.NumberValue `hcl:"replicas_per_node_group,attr"`
	// ReplicationGroupId: string, required
	ReplicationGroupId terra.StringValue `hcl:"replication_group_id,attr" validate:"required"`
	// SecurityGroupIds: set of string, optional
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// SecurityGroupNames: set of string, optional
	SecurityGroupNames terra.SetValue[terra.StringValue] `hcl:"security_group_names,attr"`
	// SnapshotArns: set of string, optional
	SnapshotArns terra.SetValue[terra.StringValue] `hcl:"snapshot_arns,attr"`
	// SnapshotName: string, optional
	SnapshotName terra.StringValue `hcl:"snapshot_name,attr"`
	// SnapshotRetentionLimit: number, optional
	SnapshotRetentionLimit terra.NumberValue `hcl:"snapshot_retention_limit,attr"`
	// SnapshotWindow: string, optional
	SnapshotWindow terra.StringValue `hcl:"snapshot_window,attr"`
	// SubnetGroupName: string, optional
	SubnetGroupName terra.StringValue `hcl:"subnet_group_name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TransitEncryptionEnabled: bool, optional
	TransitEncryptionEnabled terra.BoolValue `hcl:"transit_encryption_enabled,attr"`
	// UserGroupIds: set of string, optional
	UserGroupIds terra.SetValue[terra.StringValue] `hcl:"user_group_ids,attr"`
	// LogDeliveryConfiguration: min=0,max=2
	LogDeliveryConfiguration []elasticachereplicationgroup.LogDeliveryConfiguration `hcl:"log_delivery_configuration,block" validate:"min=0,max=2"`
	// Timeouts: optional
	Timeouts *elasticachereplicationgroup.Timeouts `hcl:"timeouts,block"`
}
type elasticacheReplicationGroupAttributes struct {
	ref terra.Reference
}

// ApplyImmediately returns a reference to field apply_immediately of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) ApplyImmediately() terra.BoolValue {
	return terra.ReferenceAsBool(erg.ref.Append("apply_immediately"))
}

// Arn returns a reference to field arn of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("arn"))
}

// AtRestEncryptionEnabled returns a reference to field at_rest_encryption_enabled of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) AtRestEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(erg.ref.Append("at_rest_encryption_enabled"))
}

// AuthToken returns a reference to field auth_token of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) AuthToken() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("auth_token"))
}

// AutoMinorVersionUpgrade returns a reference to field auto_minor_version_upgrade of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) AutoMinorVersionUpgrade() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("auto_minor_version_upgrade"))
}

// AutomaticFailoverEnabled returns a reference to field automatic_failover_enabled of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) AutomaticFailoverEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(erg.ref.Append("automatic_failover_enabled"))
}

// ClusterEnabled returns a reference to field cluster_enabled of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) ClusterEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(erg.ref.Append("cluster_enabled"))
}

// ConfigurationEndpointAddress returns a reference to field configuration_endpoint_address of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) ConfigurationEndpointAddress() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("configuration_endpoint_address"))
}

// DataTieringEnabled returns a reference to field data_tiering_enabled of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) DataTieringEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(erg.ref.Append("data_tiering_enabled"))
}

// Description returns a reference to field description of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("description"))
}

// Engine returns a reference to field engine of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("engine_version"))
}

// EngineVersionActual returns a reference to field engine_version_actual of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) EngineVersionActual() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("engine_version_actual"))
}

// FinalSnapshotIdentifier returns a reference to field final_snapshot_identifier of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) FinalSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("final_snapshot_identifier"))
}

// GlobalReplicationGroupId returns a reference to field global_replication_group_id of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) GlobalReplicationGroupId() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("global_replication_group_id"))
}

// Id returns a reference to field id of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("kms_key_id"))
}

// MaintenanceWindow returns a reference to field maintenance_window of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) MaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("maintenance_window"))
}

// MemberClusters returns a reference to field member_clusters of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) MemberClusters() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](erg.ref.Append("member_clusters"))
}

// MultiAzEnabled returns a reference to field multi_az_enabled of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) MultiAzEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(erg.ref.Append("multi_az_enabled"))
}

// NodeType returns a reference to field node_type of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) NodeType() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("node_type"))
}

// NotificationTopicArn returns a reference to field notification_topic_arn of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) NotificationTopicArn() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("notification_topic_arn"))
}

// NumCacheClusters returns a reference to field num_cache_clusters of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) NumCacheClusters() terra.NumberValue {
	return terra.ReferenceAsNumber(erg.ref.Append("num_cache_clusters"))
}

// NumNodeGroups returns a reference to field num_node_groups of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) NumNodeGroups() terra.NumberValue {
	return terra.ReferenceAsNumber(erg.ref.Append("num_node_groups"))
}

// ParameterGroupName returns a reference to field parameter_group_name of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) ParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("parameter_group_name"))
}

// Port returns a reference to field port of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(erg.ref.Append("port"))
}

// PreferredCacheClusterAzs returns a reference to field preferred_cache_cluster_azs of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) PreferredCacheClusterAzs() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](erg.ref.Append("preferred_cache_cluster_azs"))
}

// PrimaryEndpointAddress returns a reference to field primary_endpoint_address of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) PrimaryEndpointAddress() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("primary_endpoint_address"))
}

// ReaderEndpointAddress returns a reference to field reader_endpoint_address of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) ReaderEndpointAddress() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("reader_endpoint_address"))
}

// ReplicasPerNodeGroup returns a reference to field replicas_per_node_group of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) ReplicasPerNodeGroup() terra.NumberValue {
	return terra.ReferenceAsNumber(erg.ref.Append("replicas_per_node_group"))
}

// ReplicationGroupId returns a reference to field replication_group_id of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) ReplicationGroupId() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("replication_group_id"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](erg.ref.Append("security_group_ids"))
}

// SecurityGroupNames returns a reference to field security_group_names of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) SecurityGroupNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](erg.ref.Append("security_group_names"))
}

// SnapshotArns returns a reference to field snapshot_arns of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) SnapshotArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](erg.ref.Append("snapshot_arns"))
}

// SnapshotName returns a reference to field snapshot_name of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) SnapshotName() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("snapshot_name"))
}

// SnapshotRetentionLimit returns a reference to field snapshot_retention_limit of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) SnapshotRetentionLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(erg.ref.Append("snapshot_retention_limit"))
}

// SnapshotWindow returns a reference to field snapshot_window of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) SnapshotWindow() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("snapshot_window"))
}

// SubnetGroupName returns a reference to field subnet_group_name of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) SubnetGroupName() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("subnet_group_name"))
}

// Tags returns a reference to field tags of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](erg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](erg.ref.Append("tags_all"))
}

// TransitEncryptionEnabled returns a reference to field transit_encryption_enabled of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) TransitEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(erg.ref.Append("transit_encryption_enabled"))
}

// UserGroupIds returns a reference to field user_group_ids of aws_elasticache_replication_group.
func (erg elasticacheReplicationGroupAttributes) UserGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](erg.ref.Append("user_group_ids"))
}

func (erg elasticacheReplicationGroupAttributes) LogDeliveryConfiguration() terra.SetValue[elasticachereplicationgroup.LogDeliveryConfigurationAttributes] {
	return terra.ReferenceAsSet[elasticachereplicationgroup.LogDeliveryConfigurationAttributes](erg.ref.Append("log_delivery_configuration"))
}

func (erg elasticacheReplicationGroupAttributes) Timeouts() elasticachereplicationgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[elasticachereplicationgroup.TimeoutsAttributes](erg.ref.Append("timeouts"))
}

type elasticacheReplicationGroupState struct {
	ApplyImmediately             bool                                                        `json:"apply_immediately"`
	Arn                          string                                                      `json:"arn"`
	AtRestEncryptionEnabled      bool                                                        `json:"at_rest_encryption_enabled"`
	AuthToken                    string                                                      `json:"auth_token"`
	AutoMinorVersionUpgrade      string                                                      `json:"auto_minor_version_upgrade"`
	AutomaticFailoverEnabled     bool                                                        `json:"automatic_failover_enabled"`
	ClusterEnabled               bool                                                        `json:"cluster_enabled"`
	ConfigurationEndpointAddress string                                                      `json:"configuration_endpoint_address"`
	DataTieringEnabled           bool                                                        `json:"data_tiering_enabled"`
	Description                  string                                                      `json:"description"`
	Engine                       string                                                      `json:"engine"`
	EngineVersion                string                                                      `json:"engine_version"`
	EngineVersionActual          string                                                      `json:"engine_version_actual"`
	FinalSnapshotIdentifier      string                                                      `json:"final_snapshot_identifier"`
	GlobalReplicationGroupId     string                                                      `json:"global_replication_group_id"`
	Id                           string                                                      `json:"id"`
	KmsKeyId                     string                                                      `json:"kms_key_id"`
	MaintenanceWindow            string                                                      `json:"maintenance_window"`
	MemberClusters               []string                                                    `json:"member_clusters"`
	MultiAzEnabled               bool                                                        `json:"multi_az_enabled"`
	NodeType                     string                                                      `json:"node_type"`
	NotificationTopicArn         string                                                      `json:"notification_topic_arn"`
	NumCacheClusters             float64                                                     `json:"num_cache_clusters"`
	NumNodeGroups                float64                                                     `json:"num_node_groups"`
	ParameterGroupName           string                                                      `json:"parameter_group_name"`
	Port                         float64                                                     `json:"port"`
	PreferredCacheClusterAzs     []string                                                    `json:"preferred_cache_cluster_azs"`
	PrimaryEndpointAddress       string                                                      `json:"primary_endpoint_address"`
	ReaderEndpointAddress        string                                                      `json:"reader_endpoint_address"`
	ReplicasPerNodeGroup         float64                                                     `json:"replicas_per_node_group"`
	ReplicationGroupId           string                                                      `json:"replication_group_id"`
	SecurityGroupIds             []string                                                    `json:"security_group_ids"`
	SecurityGroupNames           []string                                                    `json:"security_group_names"`
	SnapshotArns                 []string                                                    `json:"snapshot_arns"`
	SnapshotName                 string                                                      `json:"snapshot_name"`
	SnapshotRetentionLimit       float64                                                     `json:"snapshot_retention_limit"`
	SnapshotWindow               string                                                      `json:"snapshot_window"`
	SubnetGroupName              string                                                      `json:"subnet_group_name"`
	Tags                         map[string]string                                           `json:"tags"`
	TagsAll                      map[string]string                                           `json:"tags_all"`
	TransitEncryptionEnabled     bool                                                        `json:"transit_encryption_enabled"`
	UserGroupIds                 []string                                                    `json:"user_group_ids"`
	LogDeliveryConfiguration     []elasticachereplicationgroup.LogDeliveryConfigurationState `json:"log_delivery_configuration"`
	Timeouts                     *elasticachereplicationgroup.TimeoutsState                  `json:"timeouts"`
}
