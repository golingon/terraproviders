// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataelasticachereplicationgroup "github.com/golingon/terraproviders/aws/5.9.0/dataelasticachereplicationgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataElasticacheReplicationGroup creates a new instance of [DataElasticacheReplicationGroup].
func NewDataElasticacheReplicationGroup(name string, args DataElasticacheReplicationGroupArgs) *DataElasticacheReplicationGroup {
	return &DataElasticacheReplicationGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataElasticacheReplicationGroup)(nil)

// DataElasticacheReplicationGroup represents the Terraform data resource aws_elasticache_replication_group.
type DataElasticacheReplicationGroup struct {
	Name string
	Args DataElasticacheReplicationGroupArgs
}

// DataSource returns the Terraform object type for [DataElasticacheReplicationGroup].
func (erg *DataElasticacheReplicationGroup) DataSource() string {
	return "aws_elasticache_replication_group"
}

// LocalName returns the local name for [DataElasticacheReplicationGroup].
func (erg *DataElasticacheReplicationGroup) LocalName() string {
	return erg.Name
}

// Configuration returns the configuration (args) for [DataElasticacheReplicationGroup].
func (erg *DataElasticacheReplicationGroup) Configuration() interface{} {
	return erg.Args
}

// Attributes returns the attributes for [DataElasticacheReplicationGroup].
func (erg *DataElasticacheReplicationGroup) Attributes() dataElasticacheReplicationGroupAttributes {
	return dataElasticacheReplicationGroupAttributes{ref: terra.ReferenceDataResource(erg)}
}

// DataElasticacheReplicationGroupArgs contains the configurations for aws_elasticache_replication_group.
type DataElasticacheReplicationGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ReplicationGroupId: string, required
	ReplicationGroupId terra.StringValue `hcl:"replication_group_id,attr" validate:"required"`
	// LogDeliveryConfiguration: min=0
	LogDeliveryConfiguration []dataelasticachereplicationgroup.LogDeliveryConfiguration `hcl:"log_delivery_configuration,block" validate:"min=0"`
}
type dataElasticacheReplicationGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("arn"))
}

// AuthTokenEnabled returns a reference to field auth_token_enabled of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) AuthTokenEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(erg.ref.Append("auth_token_enabled"))
}

// AutomaticFailoverEnabled returns a reference to field automatic_failover_enabled of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) AutomaticFailoverEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(erg.ref.Append("automatic_failover_enabled"))
}

// ConfigurationEndpointAddress returns a reference to field configuration_endpoint_address of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) ConfigurationEndpointAddress() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("configuration_endpoint_address"))
}

// Description returns a reference to field description of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("description"))
}

// Id returns a reference to field id of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("id"))
}

// MemberClusters returns a reference to field member_clusters of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) MemberClusters() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](erg.ref.Append("member_clusters"))
}

// MultiAzEnabled returns a reference to field multi_az_enabled of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) MultiAzEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(erg.ref.Append("multi_az_enabled"))
}

// NodeType returns a reference to field node_type of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) NodeType() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("node_type"))
}

// NumCacheClusters returns a reference to field num_cache_clusters of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) NumCacheClusters() terra.NumberValue {
	return terra.ReferenceAsNumber(erg.ref.Append("num_cache_clusters"))
}

// NumNodeGroups returns a reference to field num_node_groups of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) NumNodeGroups() terra.NumberValue {
	return terra.ReferenceAsNumber(erg.ref.Append("num_node_groups"))
}

// Port returns a reference to field port of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(erg.ref.Append("port"))
}

// PrimaryEndpointAddress returns a reference to field primary_endpoint_address of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) PrimaryEndpointAddress() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("primary_endpoint_address"))
}

// ReaderEndpointAddress returns a reference to field reader_endpoint_address of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) ReaderEndpointAddress() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("reader_endpoint_address"))
}

// ReplicasPerNodeGroup returns a reference to field replicas_per_node_group of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) ReplicasPerNodeGroup() terra.NumberValue {
	return terra.ReferenceAsNumber(erg.ref.Append("replicas_per_node_group"))
}

// ReplicationGroupId returns a reference to field replication_group_id of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) ReplicationGroupId() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("replication_group_id"))
}

// SnapshotRetentionLimit returns a reference to field snapshot_retention_limit of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) SnapshotRetentionLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(erg.ref.Append("snapshot_retention_limit"))
}

// SnapshotWindow returns a reference to field snapshot_window of aws_elasticache_replication_group.
func (erg dataElasticacheReplicationGroupAttributes) SnapshotWindow() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("snapshot_window"))
}

func (erg dataElasticacheReplicationGroupAttributes) LogDeliveryConfiguration() terra.SetValue[dataelasticachereplicationgroup.LogDeliveryConfigurationAttributes] {
	return terra.ReferenceAsSet[dataelasticachereplicationgroup.LogDeliveryConfigurationAttributes](erg.ref.Append("log_delivery_configuration"))
}
