// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	dataredisinstance "github.com/golingon/terraproviders/google/4.66.0/dataredisinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataRedisInstance creates a new instance of [DataRedisInstance].
func NewDataRedisInstance(name string, args DataRedisInstanceArgs) *DataRedisInstance {
	return &DataRedisInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRedisInstance)(nil)

// DataRedisInstance represents the Terraform data resource google_redis_instance.
type DataRedisInstance struct {
	Name string
	Args DataRedisInstanceArgs
}

// DataSource returns the Terraform object type for [DataRedisInstance].
func (ri *DataRedisInstance) DataSource() string {
	return "google_redis_instance"
}

// LocalName returns the local name for [DataRedisInstance].
func (ri *DataRedisInstance) LocalName() string {
	return ri.Name
}

// Configuration returns the configuration (args) for [DataRedisInstance].
func (ri *DataRedisInstance) Configuration() interface{} {
	return ri.Args
}

// Attributes returns the attributes for [DataRedisInstance].
func (ri *DataRedisInstance) Attributes() dataRedisInstanceAttributes {
	return dataRedisInstanceAttributes{ref: terra.ReferenceDataResource(ri)}
}

// DataRedisInstanceArgs contains the configurations for google_redis_instance.
type DataRedisInstanceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// MaintenancePolicy: min=0
	MaintenancePolicy []dataredisinstance.MaintenancePolicy `hcl:"maintenance_policy,block" validate:"min=0"`
	// MaintenanceSchedule: min=0
	MaintenanceSchedule []dataredisinstance.MaintenanceSchedule `hcl:"maintenance_schedule,block" validate:"min=0"`
	// Nodes: min=0
	Nodes []dataredisinstance.Nodes `hcl:"nodes,block" validate:"min=0"`
	// PersistenceConfig: min=0
	PersistenceConfig []dataredisinstance.PersistenceConfig `hcl:"persistence_config,block" validate:"min=0"`
	// ServerCaCerts: min=0
	ServerCaCerts []dataredisinstance.ServerCaCerts `hcl:"server_ca_certs,block" validate:"min=0"`
}
type dataRedisInstanceAttributes struct {
	ref terra.Reference
}

// AlternativeLocationId returns a reference to field alternative_location_id of google_redis_instance.
func (ri dataRedisInstanceAttributes) AlternativeLocationId() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("alternative_location_id"))
}

// AuthEnabled returns a reference to field auth_enabled of google_redis_instance.
func (ri dataRedisInstanceAttributes) AuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ri.ref.Append("auth_enabled"))
}

// AuthString returns a reference to field auth_string of google_redis_instance.
func (ri dataRedisInstanceAttributes) AuthString() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("auth_string"))
}

// AuthorizedNetwork returns a reference to field authorized_network of google_redis_instance.
func (ri dataRedisInstanceAttributes) AuthorizedNetwork() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("authorized_network"))
}

// ConnectMode returns a reference to field connect_mode of google_redis_instance.
func (ri dataRedisInstanceAttributes) ConnectMode() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("connect_mode"))
}

// CreateTime returns a reference to field create_time of google_redis_instance.
func (ri dataRedisInstanceAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("create_time"))
}

// CurrentLocationId returns a reference to field current_location_id of google_redis_instance.
func (ri dataRedisInstanceAttributes) CurrentLocationId() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("current_location_id"))
}

// CustomerManagedKey returns a reference to field customer_managed_key of google_redis_instance.
func (ri dataRedisInstanceAttributes) CustomerManagedKey() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("customer_managed_key"))
}

// DisplayName returns a reference to field display_name of google_redis_instance.
func (ri dataRedisInstanceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("display_name"))
}

// Host returns a reference to field host of google_redis_instance.
func (ri dataRedisInstanceAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("host"))
}

// Id returns a reference to field id of google_redis_instance.
func (ri dataRedisInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("id"))
}

// Labels returns a reference to field labels of google_redis_instance.
func (ri dataRedisInstanceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ri.ref.Append("labels"))
}

// LocationId returns a reference to field location_id of google_redis_instance.
func (ri dataRedisInstanceAttributes) LocationId() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("location_id"))
}

// MemorySizeGb returns a reference to field memory_size_gb of google_redis_instance.
func (ri dataRedisInstanceAttributes) MemorySizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(ri.ref.Append("memory_size_gb"))
}

// Name returns a reference to field name of google_redis_instance.
func (ri dataRedisInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("name"))
}

// PersistenceIamIdentity returns a reference to field persistence_iam_identity of google_redis_instance.
func (ri dataRedisInstanceAttributes) PersistenceIamIdentity() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("persistence_iam_identity"))
}

// Port returns a reference to field port of google_redis_instance.
func (ri dataRedisInstanceAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ri.ref.Append("port"))
}

// Project returns a reference to field project of google_redis_instance.
func (ri dataRedisInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("project"))
}

// ReadEndpoint returns a reference to field read_endpoint of google_redis_instance.
func (ri dataRedisInstanceAttributes) ReadEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("read_endpoint"))
}

// ReadEndpointPort returns a reference to field read_endpoint_port of google_redis_instance.
func (ri dataRedisInstanceAttributes) ReadEndpointPort() terra.NumberValue {
	return terra.ReferenceAsNumber(ri.ref.Append("read_endpoint_port"))
}

// ReadReplicasMode returns a reference to field read_replicas_mode of google_redis_instance.
func (ri dataRedisInstanceAttributes) ReadReplicasMode() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("read_replicas_mode"))
}

// RedisConfigs returns a reference to field redis_configs of google_redis_instance.
func (ri dataRedisInstanceAttributes) RedisConfigs() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ri.ref.Append("redis_configs"))
}

// RedisVersion returns a reference to field redis_version of google_redis_instance.
func (ri dataRedisInstanceAttributes) RedisVersion() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("redis_version"))
}

// Region returns a reference to field region of google_redis_instance.
func (ri dataRedisInstanceAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("region"))
}

// ReplicaCount returns a reference to field replica_count of google_redis_instance.
func (ri dataRedisInstanceAttributes) ReplicaCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ri.ref.Append("replica_count"))
}

// ReservedIpRange returns a reference to field reserved_ip_range of google_redis_instance.
func (ri dataRedisInstanceAttributes) ReservedIpRange() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("reserved_ip_range"))
}

// SecondaryIpRange returns a reference to field secondary_ip_range of google_redis_instance.
func (ri dataRedisInstanceAttributes) SecondaryIpRange() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("secondary_ip_range"))
}

// Tier returns a reference to field tier of google_redis_instance.
func (ri dataRedisInstanceAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("tier"))
}

// TransitEncryptionMode returns a reference to field transit_encryption_mode of google_redis_instance.
func (ri dataRedisInstanceAttributes) TransitEncryptionMode() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("transit_encryption_mode"))
}

func (ri dataRedisInstanceAttributes) MaintenancePolicy() terra.ListValue[dataredisinstance.MaintenancePolicyAttributes] {
	return terra.ReferenceAsList[dataredisinstance.MaintenancePolicyAttributes](ri.ref.Append("maintenance_policy"))
}

func (ri dataRedisInstanceAttributes) MaintenanceSchedule() terra.ListValue[dataredisinstance.MaintenanceScheduleAttributes] {
	return terra.ReferenceAsList[dataredisinstance.MaintenanceScheduleAttributes](ri.ref.Append("maintenance_schedule"))
}

func (ri dataRedisInstanceAttributes) Nodes() terra.ListValue[dataredisinstance.NodesAttributes] {
	return terra.ReferenceAsList[dataredisinstance.NodesAttributes](ri.ref.Append("nodes"))
}

func (ri dataRedisInstanceAttributes) PersistenceConfig() terra.ListValue[dataredisinstance.PersistenceConfigAttributes] {
	return terra.ReferenceAsList[dataredisinstance.PersistenceConfigAttributes](ri.ref.Append("persistence_config"))
}

func (ri dataRedisInstanceAttributes) ServerCaCerts() terra.ListValue[dataredisinstance.ServerCaCertsAttributes] {
	return terra.ReferenceAsList[dataredisinstance.ServerCaCertsAttributes](ri.ref.Append("server_ca_certs"))
}
