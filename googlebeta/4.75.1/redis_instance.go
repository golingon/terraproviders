// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	redisinstance "github.com/golingon/terraproviders/googlebeta/4.75.1/redisinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRedisInstance creates a new instance of [RedisInstance].
func NewRedisInstance(name string, args RedisInstanceArgs) *RedisInstance {
	return &RedisInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedisInstance)(nil)

// RedisInstance represents the Terraform resource google_redis_instance.
type RedisInstance struct {
	Name      string
	Args      RedisInstanceArgs
	state     *redisInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedisInstance].
func (ri *RedisInstance) Type() string {
	return "google_redis_instance"
}

// LocalName returns the local name for [RedisInstance].
func (ri *RedisInstance) LocalName() string {
	return ri.Name
}

// Configuration returns the configuration (args) for [RedisInstance].
func (ri *RedisInstance) Configuration() interface{} {
	return ri.Args
}

// DependOn is used for other resources to depend on [RedisInstance].
func (ri *RedisInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(ri)
}

// Dependencies returns the list of resources [RedisInstance] depends_on.
func (ri *RedisInstance) Dependencies() terra.Dependencies {
	return ri.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedisInstance].
func (ri *RedisInstance) LifecycleManagement() *terra.Lifecycle {
	return ri.Lifecycle
}

// Attributes returns the attributes for [RedisInstance].
func (ri *RedisInstance) Attributes() redisInstanceAttributes {
	return redisInstanceAttributes{ref: terra.ReferenceResource(ri)}
}

// ImportState imports the given attribute values into [RedisInstance]'s state.
func (ri *RedisInstance) ImportState(av io.Reader) error {
	ri.state = &redisInstanceState{}
	if err := json.NewDecoder(av).Decode(ri.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ri.Type(), ri.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedisInstance] has state.
func (ri *RedisInstance) State() (*redisInstanceState, bool) {
	return ri.state, ri.state != nil
}

// StateMust returns the state for [RedisInstance]. Panics if the state is nil.
func (ri *RedisInstance) StateMust() *redisInstanceState {
	if ri.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ri.Type(), ri.LocalName()))
	}
	return ri.state
}

// RedisInstanceArgs contains the configurations for google_redis_instance.
type RedisInstanceArgs struct {
	// AlternativeLocationId: string, optional
	AlternativeLocationId terra.StringValue `hcl:"alternative_location_id,attr"`
	// AuthEnabled: bool, optional
	AuthEnabled terra.BoolValue `hcl:"auth_enabled,attr"`
	// AuthorizedNetwork: string, optional
	AuthorizedNetwork terra.StringValue `hcl:"authorized_network,attr"`
	// ConnectMode: string, optional
	ConnectMode terra.StringValue `hcl:"connect_mode,attr"`
	// CustomerManagedKey: string, optional
	CustomerManagedKey terra.StringValue `hcl:"customer_managed_key,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// LocationId: string, optional
	LocationId terra.StringValue `hcl:"location_id,attr"`
	// MemorySizeGb: number, required
	MemorySizeGb terra.NumberValue `hcl:"memory_size_gb,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ReadReplicasMode: string, optional
	ReadReplicasMode terra.StringValue `hcl:"read_replicas_mode,attr"`
	// RedisConfigs: map of string, optional
	RedisConfigs terra.MapValue[terra.StringValue] `hcl:"redis_configs,attr"`
	// RedisVersion: string, optional
	RedisVersion terra.StringValue `hcl:"redis_version,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// ReplicaCount: number, optional
	ReplicaCount terra.NumberValue `hcl:"replica_count,attr"`
	// ReservedIpRange: string, optional
	ReservedIpRange terra.StringValue `hcl:"reserved_ip_range,attr"`
	// SecondaryIpRange: string, optional
	SecondaryIpRange terra.StringValue `hcl:"secondary_ip_range,attr"`
	// Tier: string, optional
	Tier terra.StringValue `hcl:"tier,attr"`
	// TransitEncryptionMode: string, optional
	TransitEncryptionMode terra.StringValue `hcl:"transit_encryption_mode,attr"`
	// MaintenanceSchedule: min=0
	MaintenanceSchedule []redisinstance.MaintenanceSchedule `hcl:"maintenance_schedule,block" validate:"min=0"`
	// Nodes: min=0
	Nodes []redisinstance.Nodes `hcl:"nodes,block" validate:"min=0"`
	// ServerCaCerts: min=0
	ServerCaCerts []redisinstance.ServerCaCerts `hcl:"server_ca_certs,block" validate:"min=0"`
	// MaintenancePolicy: optional
	MaintenancePolicy *redisinstance.MaintenancePolicy `hcl:"maintenance_policy,block"`
	// PersistenceConfig: optional
	PersistenceConfig *redisinstance.PersistenceConfig `hcl:"persistence_config,block"`
	// Timeouts: optional
	Timeouts *redisinstance.Timeouts `hcl:"timeouts,block"`
}
type redisInstanceAttributes struct {
	ref terra.Reference
}

// AlternativeLocationId returns a reference to field alternative_location_id of google_redis_instance.
func (ri redisInstanceAttributes) AlternativeLocationId() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("alternative_location_id"))
}

// AuthEnabled returns a reference to field auth_enabled of google_redis_instance.
func (ri redisInstanceAttributes) AuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ri.ref.Append("auth_enabled"))
}

// AuthString returns a reference to field auth_string of google_redis_instance.
func (ri redisInstanceAttributes) AuthString() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("auth_string"))
}

// AuthorizedNetwork returns a reference to field authorized_network of google_redis_instance.
func (ri redisInstanceAttributes) AuthorizedNetwork() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("authorized_network"))
}

// ConnectMode returns a reference to field connect_mode of google_redis_instance.
func (ri redisInstanceAttributes) ConnectMode() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("connect_mode"))
}

// CreateTime returns a reference to field create_time of google_redis_instance.
func (ri redisInstanceAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("create_time"))
}

// CurrentLocationId returns a reference to field current_location_id of google_redis_instance.
func (ri redisInstanceAttributes) CurrentLocationId() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("current_location_id"))
}

// CustomerManagedKey returns a reference to field customer_managed_key of google_redis_instance.
func (ri redisInstanceAttributes) CustomerManagedKey() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("customer_managed_key"))
}

// DisplayName returns a reference to field display_name of google_redis_instance.
func (ri redisInstanceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("display_name"))
}

// Host returns a reference to field host of google_redis_instance.
func (ri redisInstanceAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("host"))
}

// Id returns a reference to field id of google_redis_instance.
func (ri redisInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("id"))
}

// Labels returns a reference to field labels of google_redis_instance.
func (ri redisInstanceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ri.ref.Append("labels"))
}

// LocationId returns a reference to field location_id of google_redis_instance.
func (ri redisInstanceAttributes) LocationId() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("location_id"))
}

// MemorySizeGb returns a reference to field memory_size_gb of google_redis_instance.
func (ri redisInstanceAttributes) MemorySizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(ri.ref.Append("memory_size_gb"))
}

// Name returns a reference to field name of google_redis_instance.
func (ri redisInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("name"))
}

// PersistenceIamIdentity returns a reference to field persistence_iam_identity of google_redis_instance.
func (ri redisInstanceAttributes) PersistenceIamIdentity() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("persistence_iam_identity"))
}

// Port returns a reference to field port of google_redis_instance.
func (ri redisInstanceAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ri.ref.Append("port"))
}

// Project returns a reference to field project of google_redis_instance.
func (ri redisInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("project"))
}

// ReadEndpoint returns a reference to field read_endpoint of google_redis_instance.
func (ri redisInstanceAttributes) ReadEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("read_endpoint"))
}

// ReadEndpointPort returns a reference to field read_endpoint_port of google_redis_instance.
func (ri redisInstanceAttributes) ReadEndpointPort() terra.NumberValue {
	return terra.ReferenceAsNumber(ri.ref.Append("read_endpoint_port"))
}

// ReadReplicasMode returns a reference to field read_replicas_mode of google_redis_instance.
func (ri redisInstanceAttributes) ReadReplicasMode() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("read_replicas_mode"))
}

// RedisConfigs returns a reference to field redis_configs of google_redis_instance.
func (ri redisInstanceAttributes) RedisConfigs() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ri.ref.Append("redis_configs"))
}

// RedisVersion returns a reference to field redis_version of google_redis_instance.
func (ri redisInstanceAttributes) RedisVersion() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("redis_version"))
}

// Region returns a reference to field region of google_redis_instance.
func (ri redisInstanceAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("region"))
}

// ReplicaCount returns a reference to field replica_count of google_redis_instance.
func (ri redisInstanceAttributes) ReplicaCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ri.ref.Append("replica_count"))
}

// ReservedIpRange returns a reference to field reserved_ip_range of google_redis_instance.
func (ri redisInstanceAttributes) ReservedIpRange() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("reserved_ip_range"))
}

// SecondaryIpRange returns a reference to field secondary_ip_range of google_redis_instance.
func (ri redisInstanceAttributes) SecondaryIpRange() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("secondary_ip_range"))
}

// Tier returns a reference to field tier of google_redis_instance.
func (ri redisInstanceAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("tier"))
}

// TransitEncryptionMode returns a reference to field transit_encryption_mode of google_redis_instance.
func (ri redisInstanceAttributes) TransitEncryptionMode() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("transit_encryption_mode"))
}

func (ri redisInstanceAttributes) MaintenanceSchedule() terra.ListValue[redisinstance.MaintenanceScheduleAttributes] {
	return terra.ReferenceAsList[redisinstance.MaintenanceScheduleAttributes](ri.ref.Append("maintenance_schedule"))
}

func (ri redisInstanceAttributes) Nodes() terra.ListValue[redisinstance.NodesAttributes] {
	return terra.ReferenceAsList[redisinstance.NodesAttributes](ri.ref.Append("nodes"))
}

func (ri redisInstanceAttributes) ServerCaCerts() terra.ListValue[redisinstance.ServerCaCertsAttributes] {
	return terra.ReferenceAsList[redisinstance.ServerCaCertsAttributes](ri.ref.Append("server_ca_certs"))
}

func (ri redisInstanceAttributes) MaintenancePolicy() terra.ListValue[redisinstance.MaintenancePolicyAttributes] {
	return terra.ReferenceAsList[redisinstance.MaintenancePolicyAttributes](ri.ref.Append("maintenance_policy"))
}

func (ri redisInstanceAttributes) PersistenceConfig() terra.ListValue[redisinstance.PersistenceConfigAttributes] {
	return terra.ReferenceAsList[redisinstance.PersistenceConfigAttributes](ri.ref.Append("persistence_config"))
}

func (ri redisInstanceAttributes) Timeouts() redisinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[redisinstance.TimeoutsAttributes](ri.ref.Append("timeouts"))
}

type redisInstanceState struct {
	AlternativeLocationId  string                                   `json:"alternative_location_id"`
	AuthEnabled            bool                                     `json:"auth_enabled"`
	AuthString             string                                   `json:"auth_string"`
	AuthorizedNetwork      string                                   `json:"authorized_network"`
	ConnectMode            string                                   `json:"connect_mode"`
	CreateTime             string                                   `json:"create_time"`
	CurrentLocationId      string                                   `json:"current_location_id"`
	CustomerManagedKey     string                                   `json:"customer_managed_key"`
	DisplayName            string                                   `json:"display_name"`
	Host                   string                                   `json:"host"`
	Id                     string                                   `json:"id"`
	Labels                 map[string]string                        `json:"labels"`
	LocationId             string                                   `json:"location_id"`
	MemorySizeGb           float64                                  `json:"memory_size_gb"`
	Name                   string                                   `json:"name"`
	PersistenceIamIdentity string                                   `json:"persistence_iam_identity"`
	Port                   float64                                  `json:"port"`
	Project                string                                   `json:"project"`
	ReadEndpoint           string                                   `json:"read_endpoint"`
	ReadEndpointPort       float64                                  `json:"read_endpoint_port"`
	ReadReplicasMode       string                                   `json:"read_replicas_mode"`
	RedisConfigs           map[string]string                        `json:"redis_configs"`
	RedisVersion           string                                   `json:"redis_version"`
	Region                 string                                   `json:"region"`
	ReplicaCount           float64                                  `json:"replica_count"`
	ReservedIpRange        string                                   `json:"reserved_ip_range"`
	SecondaryIpRange       string                                   `json:"secondary_ip_range"`
	Tier                   string                                   `json:"tier"`
	TransitEncryptionMode  string                                   `json:"transit_encryption_mode"`
	MaintenanceSchedule    []redisinstance.MaintenanceScheduleState `json:"maintenance_schedule"`
	Nodes                  []redisinstance.NodesState               `json:"nodes"`
	ServerCaCerts          []redisinstance.ServerCaCertsState       `json:"server_ca_certs"`
	MaintenancePolicy      []redisinstance.MaintenancePolicyState   `json:"maintenance_policy"`
	PersistenceConfig      []redisinstance.PersistenceConfigState   `json:"persistence_config"`
	Timeouts               *redisinstance.TimeoutsState             `json:"timeouts"`
}
