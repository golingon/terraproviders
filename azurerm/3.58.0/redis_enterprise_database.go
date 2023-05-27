// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	redisenterprisedatabase "github.com/golingon/terraproviders/azurerm/3.58.0/redisenterprisedatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRedisEnterpriseDatabase creates a new instance of [RedisEnterpriseDatabase].
func NewRedisEnterpriseDatabase(name string, args RedisEnterpriseDatabaseArgs) *RedisEnterpriseDatabase {
	return &RedisEnterpriseDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedisEnterpriseDatabase)(nil)

// RedisEnterpriseDatabase represents the Terraform resource azurerm_redis_enterprise_database.
type RedisEnterpriseDatabase struct {
	Name      string
	Args      RedisEnterpriseDatabaseArgs
	state     *redisEnterpriseDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedisEnterpriseDatabase].
func (red *RedisEnterpriseDatabase) Type() string {
	return "azurerm_redis_enterprise_database"
}

// LocalName returns the local name for [RedisEnterpriseDatabase].
func (red *RedisEnterpriseDatabase) LocalName() string {
	return red.Name
}

// Configuration returns the configuration (args) for [RedisEnterpriseDatabase].
func (red *RedisEnterpriseDatabase) Configuration() interface{} {
	return red.Args
}

// DependOn is used for other resources to depend on [RedisEnterpriseDatabase].
func (red *RedisEnterpriseDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(red)
}

// Dependencies returns the list of resources [RedisEnterpriseDatabase] depends_on.
func (red *RedisEnterpriseDatabase) Dependencies() terra.Dependencies {
	return red.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedisEnterpriseDatabase].
func (red *RedisEnterpriseDatabase) LifecycleManagement() *terra.Lifecycle {
	return red.Lifecycle
}

// Attributes returns the attributes for [RedisEnterpriseDatabase].
func (red *RedisEnterpriseDatabase) Attributes() redisEnterpriseDatabaseAttributes {
	return redisEnterpriseDatabaseAttributes{ref: terra.ReferenceResource(red)}
}

// ImportState imports the given attribute values into [RedisEnterpriseDatabase]'s state.
func (red *RedisEnterpriseDatabase) ImportState(av io.Reader) error {
	red.state = &redisEnterpriseDatabaseState{}
	if err := json.NewDecoder(av).Decode(red.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", red.Type(), red.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedisEnterpriseDatabase] has state.
func (red *RedisEnterpriseDatabase) State() (*redisEnterpriseDatabaseState, bool) {
	return red.state, red.state != nil
}

// StateMust returns the state for [RedisEnterpriseDatabase]. Panics if the state is nil.
func (red *RedisEnterpriseDatabase) StateMust() *redisEnterpriseDatabaseState {
	if red.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", red.Type(), red.LocalName()))
	}
	return red.state
}

// RedisEnterpriseDatabaseArgs contains the configurations for azurerm_redis_enterprise_database.
type RedisEnterpriseDatabaseArgs struct {
	// ClientProtocol: string, optional
	ClientProtocol terra.StringValue `hcl:"client_protocol,attr"`
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
	// ClusteringPolicy: string, optional
	ClusteringPolicy terra.StringValue `hcl:"clustering_policy,attr"`
	// EvictionPolicy: string, optional
	EvictionPolicy terra.StringValue `hcl:"eviction_policy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkedDatabaseGroupNickname: string, optional
	LinkedDatabaseGroupNickname terra.StringValue `hcl:"linked_database_group_nickname,attr"`
	// LinkedDatabaseId: set of string, optional
	LinkedDatabaseId terra.SetValue[terra.StringValue] `hcl:"linked_database_id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// ResourceGroupName: string, optional
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr"`
	// Module: min=0,max=4
	Module []redisenterprisedatabase.Module `hcl:"module,block" validate:"min=0,max=4"`
	// Timeouts: optional
	Timeouts *redisenterprisedatabase.Timeouts `hcl:"timeouts,block"`
}
type redisEnterpriseDatabaseAttributes struct {
	ref terra.Reference
}

// ClientProtocol returns a reference to field client_protocol of azurerm_redis_enterprise_database.
func (red redisEnterpriseDatabaseAttributes) ClientProtocol() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("client_protocol"))
}

// ClusterId returns a reference to field cluster_id of azurerm_redis_enterprise_database.
func (red redisEnterpriseDatabaseAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("cluster_id"))
}

// ClusteringPolicy returns a reference to field clustering_policy of azurerm_redis_enterprise_database.
func (red redisEnterpriseDatabaseAttributes) ClusteringPolicy() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("clustering_policy"))
}

// EvictionPolicy returns a reference to field eviction_policy of azurerm_redis_enterprise_database.
func (red redisEnterpriseDatabaseAttributes) EvictionPolicy() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("eviction_policy"))
}

// Id returns a reference to field id of azurerm_redis_enterprise_database.
func (red redisEnterpriseDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("id"))
}

// LinkedDatabaseGroupNickname returns a reference to field linked_database_group_nickname of azurerm_redis_enterprise_database.
func (red redisEnterpriseDatabaseAttributes) LinkedDatabaseGroupNickname() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("linked_database_group_nickname"))
}

// LinkedDatabaseId returns a reference to field linked_database_id of azurerm_redis_enterprise_database.
func (red redisEnterpriseDatabaseAttributes) LinkedDatabaseId() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](red.ref.Append("linked_database_id"))
}

// Name returns a reference to field name of azurerm_redis_enterprise_database.
func (red redisEnterpriseDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("name"))
}

// Port returns a reference to field port of azurerm_redis_enterprise_database.
func (red redisEnterpriseDatabaseAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(red.ref.Append("port"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_redis_enterprise_database.
func (red redisEnterpriseDatabaseAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("primary_access_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_redis_enterprise_database.
func (red redisEnterpriseDatabaseAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_redis_enterprise_database.
func (red redisEnterpriseDatabaseAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("secondary_access_key"))
}

func (red redisEnterpriseDatabaseAttributes) Module() terra.ListValue[redisenterprisedatabase.ModuleAttributes] {
	return terra.ReferenceAsList[redisenterprisedatabase.ModuleAttributes](red.ref.Append("module"))
}

func (red redisEnterpriseDatabaseAttributes) Timeouts() redisenterprisedatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[redisenterprisedatabase.TimeoutsAttributes](red.ref.Append("timeouts"))
}

type redisEnterpriseDatabaseState struct {
	ClientProtocol              string                                 `json:"client_protocol"`
	ClusterId                   string                                 `json:"cluster_id"`
	ClusteringPolicy            string                                 `json:"clustering_policy"`
	EvictionPolicy              string                                 `json:"eviction_policy"`
	Id                          string                                 `json:"id"`
	LinkedDatabaseGroupNickname string                                 `json:"linked_database_group_nickname"`
	LinkedDatabaseId            []string                               `json:"linked_database_id"`
	Name                        string                                 `json:"name"`
	Port                        float64                                `json:"port"`
	PrimaryAccessKey            string                                 `json:"primary_access_key"`
	ResourceGroupName           string                                 `json:"resource_group_name"`
	SecondaryAccessKey          string                                 `json:"secondary_access_key"`
	Module                      []redisenterprisedatabase.ModuleState  `json:"module"`
	Timeouts                    *redisenterprisedatabase.TimeoutsState `json:"timeouts"`
}
