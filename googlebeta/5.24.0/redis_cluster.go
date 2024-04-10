// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	rediscluster "github.com/golingon/terraproviders/googlebeta/5.24.0/rediscluster"
	"io"
)

// NewRedisCluster creates a new instance of [RedisCluster].
func NewRedisCluster(name string, args RedisClusterArgs) *RedisCluster {
	return &RedisCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedisCluster)(nil)

// RedisCluster represents the Terraform resource google_redis_cluster.
type RedisCluster struct {
	Name      string
	Args      RedisClusterArgs
	state     *redisClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedisCluster].
func (rc *RedisCluster) Type() string {
	return "google_redis_cluster"
}

// LocalName returns the local name for [RedisCluster].
func (rc *RedisCluster) LocalName() string {
	return rc.Name
}

// Configuration returns the configuration (args) for [RedisCluster].
func (rc *RedisCluster) Configuration() interface{} {
	return rc.Args
}

// DependOn is used for other resources to depend on [RedisCluster].
func (rc *RedisCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(rc)
}

// Dependencies returns the list of resources [RedisCluster] depends_on.
func (rc *RedisCluster) Dependencies() terra.Dependencies {
	return rc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedisCluster].
func (rc *RedisCluster) LifecycleManagement() *terra.Lifecycle {
	return rc.Lifecycle
}

// Attributes returns the attributes for [RedisCluster].
func (rc *RedisCluster) Attributes() redisClusterAttributes {
	return redisClusterAttributes{ref: terra.ReferenceResource(rc)}
}

// ImportState imports the given attribute values into [RedisCluster]'s state.
func (rc *RedisCluster) ImportState(av io.Reader) error {
	rc.state = &redisClusterState{}
	if err := json.NewDecoder(av).Decode(rc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rc.Type(), rc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedisCluster] has state.
func (rc *RedisCluster) State() (*redisClusterState, bool) {
	return rc.state, rc.state != nil
}

// StateMust returns the state for [RedisCluster]. Panics if the state is nil.
func (rc *RedisCluster) StateMust() *redisClusterState {
	if rc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rc.Type(), rc.LocalName()))
	}
	return rc.state
}

// RedisClusterArgs contains the configurations for google_redis_cluster.
type RedisClusterArgs struct {
	// AuthorizationMode: string, optional
	AuthorizationMode terra.StringValue `hcl:"authorization_mode,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NodeType: string, optional
	NodeType terra.StringValue `hcl:"node_type,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// ReplicaCount: number, optional
	ReplicaCount terra.NumberValue `hcl:"replica_count,attr"`
	// ShardCount: number, required
	ShardCount terra.NumberValue `hcl:"shard_count,attr" validate:"required"`
	// TransitEncryptionMode: string, optional
	TransitEncryptionMode terra.StringValue `hcl:"transit_encryption_mode,attr"`
	// DiscoveryEndpoints: min=0
	DiscoveryEndpoints []rediscluster.DiscoveryEndpoints `hcl:"discovery_endpoints,block" validate:"min=0"`
	// PscConnections: min=0
	PscConnections []rediscluster.PscConnections `hcl:"psc_connections,block" validate:"min=0"`
	// StateInfo: min=0
	StateInfo []rediscluster.StateInfo `hcl:"state_info,block" validate:"min=0"`
	// PscConfigs: min=1
	PscConfigs []rediscluster.PscConfigs `hcl:"psc_configs,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *rediscluster.Timeouts `hcl:"timeouts,block"`
}
type redisClusterAttributes struct {
	ref terra.Reference
}

// AuthorizationMode returns a reference to field authorization_mode of google_redis_cluster.
func (rc redisClusterAttributes) AuthorizationMode() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("authorization_mode"))
}

// CreateTime returns a reference to field create_time of google_redis_cluster.
func (rc redisClusterAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("create_time"))
}

// Id returns a reference to field id of google_redis_cluster.
func (rc redisClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("id"))
}

// Name returns a reference to field name of google_redis_cluster.
func (rc redisClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("name"))
}

// NodeType returns a reference to field node_type of google_redis_cluster.
func (rc redisClusterAttributes) NodeType() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("node_type"))
}

// PreciseSizeGb returns a reference to field precise_size_gb of google_redis_cluster.
func (rc redisClusterAttributes) PreciseSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("precise_size_gb"))
}

// Project returns a reference to field project of google_redis_cluster.
func (rc redisClusterAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("project"))
}

// Region returns a reference to field region of google_redis_cluster.
func (rc redisClusterAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("region"))
}

// ReplicaCount returns a reference to field replica_count of google_redis_cluster.
func (rc redisClusterAttributes) ReplicaCount() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("replica_count"))
}

// ShardCount returns a reference to field shard_count of google_redis_cluster.
func (rc redisClusterAttributes) ShardCount() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("shard_count"))
}

// SizeGb returns a reference to field size_gb of google_redis_cluster.
func (rc redisClusterAttributes) SizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("size_gb"))
}

// State returns a reference to field state of google_redis_cluster.
func (rc redisClusterAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("state"))
}

// TransitEncryptionMode returns a reference to field transit_encryption_mode of google_redis_cluster.
func (rc redisClusterAttributes) TransitEncryptionMode() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("transit_encryption_mode"))
}

// Uid returns a reference to field uid of google_redis_cluster.
func (rc redisClusterAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("uid"))
}

func (rc redisClusterAttributes) DiscoveryEndpoints() terra.ListValue[rediscluster.DiscoveryEndpointsAttributes] {
	return terra.ReferenceAsList[rediscluster.DiscoveryEndpointsAttributes](rc.ref.Append("discovery_endpoints"))
}

func (rc redisClusterAttributes) PscConnections() terra.ListValue[rediscluster.PscConnectionsAttributes] {
	return terra.ReferenceAsList[rediscluster.PscConnectionsAttributes](rc.ref.Append("psc_connections"))
}

func (rc redisClusterAttributes) StateInfo() terra.ListValue[rediscluster.StateInfoAttributes] {
	return terra.ReferenceAsList[rediscluster.StateInfoAttributes](rc.ref.Append("state_info"))
}

func (rc redisClusterAttributes) PscConfigs() terra.ListValue[rediscluster.PscConfigsAttributes] {
	return terra.ReferenceAsList[rediscluster.PscConfigsAttributes](rc.ref.Append("psc_configs"))
}

func (rc redisClusterAttributes) Timeouts() rediscluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[rediscluster.TimeoutsAttributes](rc.ref.Append("timeouts"))
}

type redisClusterState struct {
	AuthorizationMode     string                                 `json:"authorization_mode"`
	CreateTime            string                                 `json:"create_time"`
	Id                    string                                 `json:"id"`
	Name                  string                                 `json:"name"`
	NodeType              string                                 `json:"node_type"`
	PreciseSizeGb         float64                                `json:"precise_size_gb"`
	Project               string                                 `json:"project"`
	Region                string                                 `json:"region"`
	ReplicaCount          float64                                `json:"replica_count"`
	ShardCount            float64                                `json:"shard_count"`
	SizeGb                float64                                `json:"size_gb"`
	State                 string                                 `json:"state"`
	TransitEncryptionMode string                                 `json:"transit_encryption_mode"`
	Uid                   string                                 `json:"uid"`
	DiscoveryEndpoints    []rediscluster.DiscoveryEndpointsState `json:"discovery_endpoints"`
	PscConnections        []rediscluster.PscConnectionsState     `json:"psc_connections"`
	StateInfo             []rediscluster.StateInfoState          `json:"state_info"`
	PscConfigs            []rediscluster.PscConfigsState         `json:"psc_configs"`
	Timeouts              *rediscluster.TimeoutsState            `json:"timeouts"`
}
