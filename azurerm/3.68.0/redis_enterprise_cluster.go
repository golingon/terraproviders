// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	redisenterprisecluster "github.com/golingon/terraproviders/azurerm/3.68.0/redisenterprisecluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRedisEnterpriseCluster creates a new instance of [RedisEnterpriseCluster].
func NewRedisEnterpriseCluster(name string, args RedisEnterpriseClusterArgs) *RedisEnterpriseCluster {
	return &RedisEnterpriseCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedisEnterpriseCluster)(nil)

// RedisEnterpriseCluster represents the Terraform resource azurerm_redis_enterprise_cluster.
type RedisEnterpriseCluster struct {
	Name      string
	Args      RedisEnterpriseClusterArgs
	state     *redisEnterpriseClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedisEnterpriseCluster].
func (rec *RedisEnterpriseCluster) Type() string {
	return "azurerm_redis_enterprise_cluster"
}

// LocalName returns the local name for [RedisEnterpriseCluster].
func (rec *RedisEnterpriseCluster) LocalName() string {
	return rec.Name
}

// Configuration returns the configuration (args) for [RedisEnterpriseCluster].
func (rec *RedisEnterpriseCluster) Configuration() interface{} {
	return rec.Args
}

// DependOn is used for other resources to depend on [RedisEnterpriseCluster].
func (rec *RedisEnterpriseCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(rec)
}

// Dependencies returns the list of resources [RedisEnterpriseCluster] depends_on.
func (rec *RedisEnterpriseCluster) Dependencies() terra.Dependencies {
	return rec.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedisEnterpriseCluster].
func (rec *RedisEnterpriseCluster) LifecycleManagement() *terra.Lifecycle {
	return rec.Lifecycle
}

// Attributes returns the attributes for [RedisEnterpriseCluster].
func (rec *RedisEnterpriseCluster) Attributes() redisEnterpriseClusterAttributes {
	return redisEnterpriseClusterAttributes{ref: terra.ReferenceResource(rec)}
}

// ImportState imports the given attribute values into [RedisEnterpriseCluster]'s state.
func (rec *RedisEnterpriseCluster) ImportState(av io.Reader) error {
	rec.state = &redisEnterpriseClusterState{}
	if err := json.NewDecoder(av).Decode(rec.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rec.Type(), rec.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedisEnterpriseCluster] has state.
func (rec *RedisEnterpriseCluster) State() (*redisEnterpriseClusterState, bool) {
	return rec.state, rec.state != nil
}

// StateMust returns the state for [RedisEnterpriseCluster]. Panics if the state is nil.
func (rec *RedisEnterpriseCluster) StateMust() *redisEnterpriseClusterState {
	if rec.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rec.Type(), rec.LocalName()))
	}
	return rec.state
}

// RedisEnterpriseClusterArgs contains the configurations for azurerm_redis_enterprise_cluster.
type RedisEnterpriseClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MinimumTlsVersion: string, optional
	MinimumTlsVersion terra.StringValue `hcl:"minimum_tls_version,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Zones: set of string, optional
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr"`
	// Timeouts: optional
	Timeouts *redisenterprisecluster.Timeouts `hcl:"timeouts,block"`
}
type redisEnterpriseClusterAttributes struct {
	ref terra.Reference
}

// Hostname returns a reference to field hostname of azurerm_redis_enterprise_cluster.
func (rec redisEnterpriseClusterAttributes) Hostname() terra.StringValue {
	return terra.ReferenceAsString(rec.ref.Append("hostname"))
}

// Id returns a reference to field id of azurerm_redis_enterprise_cluster.
func (rec redisEnterpriseClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rec.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_redis_enterprise_cluster.
func (rec redisEnterpriseClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(rec.ref.Append("location"))
}

// MinimumTlsVersion returns a reference to field minimum_tls_version of azurerm_redis_enterprise_cluster.
func (rec redisEnterpriseClusterAttributes) MinimumTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(rec.ref.Append("minimum_tls_version"))
}

// Name returns a reference to field name of azurerm_redis_enterprise_cluster.
func (rec redisEnterpriseClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rec.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_redis_enterprise_cluster.
func (rec redisEnterpriseClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(rec.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_redis_enterprise_cluster.
func (rec redisEnterpriseClusterAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(rec.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_redis_enterprise_cluster.
func (rec redisEnterpriseClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rec.ref.Append("tags"))
}

// Zones returns a reference to field zones of azurerm_redis_enterprise_cluster.
func (rec redisEnterpriseClusterAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rec.ref.Append("zones"))
}

func (rec redisEnterpriseClusterAttributes) Timeouts() redisenterprisecluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[redisenterprisecluster.TimeoutsAttributes](rec.ref.Append("timeouts"))
}

type redisEnterpriseClusterState struct {
	Hostname          string                                `json:"hostname"`
	Id                string                                `json:"id"`
	Location          string                                `json:"location"`
	MinimumTlsVersion string                                `json:"minimum_tls_version"`
	Name              string                                `json:"name"`
	ResourceGroupName string                                `json:"resource_group_name"`
	SkuName           string                                `json:"sku_name"`
	Tags              map[string]string                     `json:"tags"`
	Zones             []string                              `json:"zones"`
	Timeouts          *redisenterprisecluster.TimeoutsState `json:"timeouts"`
}
