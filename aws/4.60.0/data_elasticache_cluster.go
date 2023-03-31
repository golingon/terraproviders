// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataelasticachecluster "github.com/golingon/terraproviders/aws/4.60.0/dataelasticachecluster"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataElasticacheCluster creates a new instance of [DataElasticacheCluster].
func NewDataElasticacheCluster(name string, args DataElasticacheClusterArgs) *DataElasticacheCluster {
	return &DataElasticacheCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataElasticacheCluster)(nil)

// DataElasticacheCluster represents the Terraform data resource aws_elasticache_cluster.
type DataElasticacheCluster struct {
	Name string
	Args DataElasticacheClusterArgs
}

// DataSource returns the Terraform object type for [DataElasticacheCluster].
func (ec *DataElasticacheCluster) DataSource() string {
	return "aws_elasticache_cluster"
}

// LocalName returns the local name for [DataElasticacheCluster].
func (ec *DataElasticacheCluster) LocalName() string {
	return ec.Name
}

// Configuration returns the configuration (args) for [DataElasticacheCluster].
func (ec *DataElasticacheCluster) Configuration() interface{} {
	return ec.Args
}

// Attributes returns the attributes for [DataElasticacheCluster].
func (ec *DataElasticacheCluster) Attributes() dataElasticacheClusterAttributes {
	return dataElasticacheClusterAttributes{ref: terra.ReferenceDataResource(ec)}
}

// DataElasticacheClusterArgs contains the configurations for aws_elasticache_cluster.
type DataElasticacheClusterArgs struct {
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// CacheNodes: min=0
	CacheNodes []dataelasticachecluster.CacheNodes `hcl:"cache_nodes,block" validate:"min=0"`
	// LogDeliveryConfiguration: min=0
	LogDeliveryConfiguration []dataelasticachecluster.LogDeliveryConfiguration `hcl:"log_delivery_configuration,block" validate:"min=0"`
}
type dataElasticacheClusterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("arn"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("availability_zone"))
}

// ClusterAddress returns a reference to field cluster_address of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) ClusterAddress() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("cluster_address"))
}

// ClusterId returns a reference to field cluster_id of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("cluster_id"))
}

// ConfigurationEndpoint returns a reference to field configuration_endpoint of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) ConfigurationEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("configuration_endpoint"))
}

// Engine returns a reference to field engine of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("id"))
}

// IpDiscovery returns a reference to field ip_discovery of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) IpDiscovery() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("ip_discovery"))
}

// MaintenanceWindow returns a reference to field maintenance_window of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) MaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("maintenance_window"))
}

// NetworkType returns a reference to field network_type of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) NetworkType() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("network_type"))
}

// NodeType returns a reference to field node_type of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) NodeType() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("node_type"))
}

// NotificationTopicArn returns a reference to field notification_topic_arn of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) NotificationTopicArn() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("notification_topic_arn"))
}

// NumCacheNodes returns a reference to field num_cache_nodes of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) NumCacheNodes() terra.NumberValue {
	return terra.ReferenceAsNumber(ec.ref.Append("num_cache_nodes"))
}

// ParameterGroupName returns a reference to field parameter_group_name of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) ParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("parameter_group_name"))
}

// Port returns a reference to field port of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ec.ref.Append("port"))
}

// PreferredOutpostArn returns a reference to field preferred_outpost_arn of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) PreferredOutpostArn() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("preferred_outpost_arn"))
}

// ReplicationGroupId returns a reference to field replication_group_id of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) ReplicationGroupId() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("replication_group_id"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ec.ref.Append("security_group_ids"))
}

// SecurityGroupNames returns a reference to field security_group_names of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) SecurityGroupNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ec.ref.Append("security_group_names"))
}

// SnapshotRetentionLimit returns a reference to field snapshot_retention_limit of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) SnapshotRetentionLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(ec.ref.Append("snapshot_retention_limit"))
}

// SnapshotWindow returns a reference to field snapshot_window of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) SnapshotWindow() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("snapshot_window"))
}

// SubnetGroupName returns a reference to field subnet_group_name of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) SubnetGroupName() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("subnet_group_name"))
}

// Tags returns a reference to field tags of aws_elasticache_cluster.
func (ec dataElasticacheClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ec.ref.Append("tags"))
}

func (ec dataElasticacheClusterAttributes) CacheNodes() terra.ListValue[dataelasticachecluster.CacheNodesAttributes] {
	return terra.ReferenceAsList[dataelasticachecluster.CacheNodesAttributes](ec.ref.Append("cache_nodes"))
}

func (ec dataElasticacheClusterAttributes) LogDeliveryConfiguration() terra.SetValue[dataelasticachecluster.LogDeliveryConfigurationAttributes] {
	return terra.ReferenceAsSet[dataelasticachecluster.LogDeliveryConfigurationAttributes](ec.ref.Append("log_delivery_configuration"))
}
