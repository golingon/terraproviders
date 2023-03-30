// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	elasticachereplicationgroup "github.com/golingon/terraproviders/aws/4.60.0/elasticachereplicationgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewElasticacheReplicationGroup(name string, args ElasticacheReplicationGroupArgs) *ElasticacheReplicationGroup {
	return &ElasticacheReplicationGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElasticacheReplicationGroup)(nil)

type ElasticacheReplicationGroup struct {
	Name  string
	Args  ElasticacheReplicationGroupArgs
	state *elasticacheReplicationGroupState
}

func (erg *ElasticacheReplicationGroup) Type() string {
	return "aws_elasticache_replication_group"
}

func (erg *ElasticacheReplicationGroup) LocalName() string {
	return erg.Name
}

func (erg *ElasticacheReplicationGroup) Configuration() interface{} {
	return erg.Args
}

func (erg *ElasticacheReplicationGroup) Attributes() elasticacheReplicationGroupAttributes {
	return elasticacheReplicationGroupAttributes{ref: terra.ReferenceResource(erg)}
}

func (erg *ElasticacheReplicationGroup) ImportState(av io.Reader) error {
	erg.state = &elasticacheReplicationGroupState{}
	if err := json.NewDecoder(av).Decode(erg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", erg.Type(), erg.LocalName(), err)
	}
	return nil
}

func (erg *ElasticacheReplicationGroup) State() (*elasticacheReplicationGroupState, bool) {
	return erg.state, erg.state != nil
}

func (erg *ElasticacheReplicationGroup) StateMust() *elasticacheReplicationGroupState {
	if erg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", erg.Type(), erg.LocalName()))
	}
	return erg.state
}

func (erg *ElasticacheReplicationGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(erg)
}

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
	// AvailabilityZones: set of string, optional
	AvailabilityZones terra.SetValue[terra.StringValue] `hcl:"availability_zones,attr"`
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
	// NumberCacheClusters: number, optional
	NumberCacheClusters terra.NumberValue `hcl:"number_cache_clusters,attr"`
	// ParameterGroupName: string, optional
	ParameterGroupName terra.StringValue `hcl:"parameter_group_name,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// PreferredCacheClusterAzs: list of string, optional
	PreferredCacheClusterAzs terra.ListValue[terra.StringValue] `hcl:"preferred_cache_cluster_azs,attr"`
	// ReplicasPerNodeGroup: number, optional
	ReplicasPerNodeGroup terra.NumberValue `hcl:"replicas_per_node_group,attr"`
	// ReplicationGroupDescription: string, optional
	ReplicationGroupDescription terra.StringValue `hcl:"replication_group_description,attr"`
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
	// ClusterMode: optional
	ClusterMode *elasticachereplicationgroup.ClusterMode `hcl:"cluster_mode,block"`
	// LogDeliveryConfiguration: min=0,max=2
	LogDeliveryConfiguration []elasticachereplicationgroup.LogDeliveryConfiguration `hcl:"log_delivery_configuration,block" validate:"min=0,max=2"`
	// Timeouts: optional
	Timeouts *elasticachereplicationgroup.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ElasticacheReplicationGroup depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type elasticacheReplicationGroupAttributes struct {
	ref terra.Reference
}

func (erg elasticacheReplicationGroupAttributes) ApplyImmediately() terra.BoolValue {
	return terra.ReferenceBool(erg.ref.Append("apply_immediately"))
}

func (erg elasticacheReplicationGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("arn"))
}

func (erg elasticacheReplicationGroupAttributes) AtRestEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceBool(erg.ref.Append("at_rest_encryption_enabled"))
}

func (erg elasticacheReplicationGroupAttributes) AuthToken() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("auth_token"))
}

func (erg elasticacheReplicationGroupAttributes) AutoMinorVersionUpgrade() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("auto_minor_version_upgrade"))
}

func (erg elasticacheReplicationGroupAttributes) AutomaticFailoverEnabled() terra.BoolValue {
	return terra.ReferenceBool(erg.ref.Append("automatic_failover_enabled"))
}

func (erg elasticacheReplicationGroupAttributes) AvailabilityZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](erg.ref.Append("availability_zones"))
}

func (erg elasticacheReplicationGroupAttributes) ClusterEnabled() terra.BoolValue {
	return terra.ReferenceBool(erg.ref.Append("cluster_enabled"))
}

func (erg elasticacheReplicationGroupAttributes) ConfigurationEndpointAddress() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("configuration_endpoint_address"))
}

func (erg elasticacheReplicationGroupAttributes) DataTieringEnabled() terra.BoolValue {
	return terra.ReferenceBool(erg.ref.Append("data_tiering_enabled"))
}

func (erg elasticacheReplicationGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("description"))
}

func (erg elasticacheReplicationGroupAttributes) Engine() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("engine"))
}

func (erg elasticacheReplicationGroupAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("engine_version"))
}

func (erg elasticacheReplicationGroupAttributes) EngineVersionActual() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("engine_version_actual"))
}

func (erg elasticacheReplicationGroupAttributes) FinalSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("final_snapshot_identifier"))
}

func (erg elasticacheReplicationGroupAttributes) GlobalReplicationGroupId() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("global_replication_group_id"))
}

func (erg elasticacheReplicationGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("id"))
}

func (erg elasticacheReplicationGroupAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("kms_key_id"))
}

func (erg elasticacheReplicationGroupAttributes) MaintenanceWindow() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("maintenance_window"))
}

func (erg elasticacheReplicationGroupAttributes) MemberClusters() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](erg.ref.Append("member_clusters"))
}

func (erg elasticacheReplicationGroupAttributes) MultiAzEnabled() terra.BoolValue {
	return terra.ReferenceBool(erg.ref.Append("multi_az_enabled"))
}

func (erg elasticacheReplicationGroupAttributes) NodeType() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("node_type"))
}

func (erg elasticacheReplicationGroupAttributes) NotificationTopicArn() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("notification_topic_arn"))
}

func (erg elasticacheReplicationGroupAttributes) NumCacheClusters() terra.NumberValue {
	return terra.ReferenceNumber(erg.ref.Append("num_cache_clusters"))
}

func (erg elasticacheReplicationGroupAttributes) NumNodeGroups() terra.NumberValue {
	return terra.ReferenceNumber(erg.ref.Append("num_node_groups"))
}

func (erg elasticacheReplicationGroupAttributes) NumberCacheClusters() terra.NumberValue {
	return terra.ReferenceNumber(erg.ref.Append("number_cache_clusters"))
}

func (erg elasticacheReplicationGroupAttributes) ParameterGroupName() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("parameter_group_name"))
}

func (erg elasticacheReplicationGroupAttributes) Port() terra.NumberValue {
	return terra.ReferenceNumber(erg.ref.Append("port"))
}

func (erg elasticacheReplicationGroupAttributes) PreferredCacheClusterAzs() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](erg.ref.Append("preferred_cache_cluster_azs"))
}

func (erg elasticacheReplicationGroupAttributes) PrimaryEndpointAddress() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("primary_endpoint_address"))
}

func (erg elasticacheReplicationGroupAttributes) ReaderEndpointAddress() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("reader_endpoint_address"))
}

func (erg elasticacheReplicationGroupAttributes) ReplicasPerNodeGroup() terra.NumberValue {
	return terra.ReferenceNumber(erg.ref.Append("replicas_per_node_group"))
}

func (erg elasticacheReplicationGroupAttributes) ReplicationGroupDescription() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("replication_group_description"))
}

func (erg elasticacheReplicationGroupAttributes) ReplicationGroupId() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("replication_group_id"))
}

func (erg elasticacheReplicationGroupAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](erg.ref.Append("security_group_ids"))
}

func (erg elasticacheReplicationGroupAttributes) SecurityGroupNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](erg.ref.Append("security_group_names"))
}

func (erg elasticacheReplicationGroupAttributes) SnapshotArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](erg.ref.Append("snapshot_arns"))
}

func (erg elasticacheReplicationGroupAttributes) SnapshotName() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("snapshot_name"))
}

func (erg elasticacheReplicationGroupAttributes) SnapshotRetentionLimit() terra.NumberValue {
	return terra.ReferenceNumber(erg.ref.Append("snapshot_retention_limit"))
}

func (erg elasticacheReplicationGroupAttributes) SnapshotWindow() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("snapshot_window"))
}

func (erg elasticacheReplicationGroupAttributes) SubnetGroupName() terra.StringValue {
	return terra.ReferenceString(erg.ref.Append("subnet_group_name"))
}

func (erg elasticacheReplicationGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](erg.ref.Append("tags"))
}

func (erg elasticacheReplicationGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](erg.ref.Append("tags_all"))
}

func (erg elasticacheReplicationGroupAttributes) TransitEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceBool(erg.ref.Append("transit_encryption_enabled"))
}

func (erg elasticacheReplicationGroupAttributes) UserGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](erg.ref.Append("user_group_ids"))
}

func (erg elasticacheReplicationGroupAttributes) ClusterMode() terra.ListValue[elasticachereplicationgroup.ClusterModeAttributes] {
	return terra.ReferenceList[elasticachereplicationgroup.ClusterModeAttributes](erg.ref.Append("cluster_mode"))
}

func (erg elasticacheReplicationGroupAttributes) LogDeliveryConfiguration() terra.SetValue[elasticachereplicationgroup.LogDeliveryConfigurationAttributes] {
	return terra.ReferenceSet[elasticachereplicationgroup.LogDeliveryConfigurationAttributes](erg.ref.Append("log_delivery_configuration"))
}

func (erg elasticacheReplicationGroupAttributes) Timeouts() elasticachereplicationgroup.TimeoutsAttributes {
	return terra.ReferenceSingle[elasticachereplicationgroup.TimeoutsAttributes](erg.ref.Append("timeouts"))
}

type elasticacheReplicationGroupState struct {
	ApplyImmediately             bool                                                        `json:"apply_immediately"`
	Arn                          string                                                      `json:"arn"`
	AtRestEncryptionEnabled      bool                                                        `json:"at_rest_encryption_enabled"`
	AuthToken                    string                                                      `json:"auth_token"`
	AutoMinorVersionUpgrade      string                                                      `json:"auto_minor_version_upgrade"`
	AutomaticFailoverEnabled     bool                                                        `json:"automatic_failover_enabled"`
	AvailabilityZones            []string                                                    `json:"availability_zones"`
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
	NumberCacheClusters          float64                                                     `json:"number_cache_clusters"`
	ParameterGroupName           string                                                      `json:"parameter_group_name"`
	Port                         float64                                                     `json:"port"`
	PreferredCacheClusterAzs     []string                                                    `json:"preferred_cache_cluster_azs"`
	PrimaryEndpointAddress       string                                                      `json:"primary_endpoint_address"`
	ReaderEndpointAddress        string                                                      `json:"reader_endpoint_address"`
	ReplicasPerNodeGroup         float64                                                     `json:"replicas_per_node_group"`
	ReplicationGroupDescription  string                                                      `json:"replication_group_description"`
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
	ClusterMode                  []elasticachereplicationgroup.ClusterModeState              `json:"cluster_mode"`
	LogDeliveryConfiguration     []elasticachereplicationgroup.LogDeliveryConfigurationState `json:"log_delivery_configuration"`
	Timeouts                     *elasticachereplicationgroup.TimeoutsState                  `json:"timeouts"`
}
