// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	redislinkedserver "github.com/golingon/terraproviders/azurerm/3.52.0/redislinkedserver"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRedisLinkedServer creates a new instance of [RedisLinkedServer].
func NewRedisLinkedServer(name string, args RedisLinkedServerArgs) *RedisLinkedServer {
	return &RedisLinkedServer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedisLinkedServer)(nil)

// RedisLinkedServer represents the Terraform resource azurerm_redis_linked_server.
type RedisLinkedServer struct {
	Name      string
	Args      RedisLinkedServerArgs
	state     *redisLinkedServerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedisLinkedServer].
func (rls *RedisLinkedServer) Type() string {
	return "azurerm_redis_linked_server"
}

// LocalName returns the local name for [RedisLinkedServer].
func (rls *RedisLinkedServer) LocalName() string {
	return rls.Name
}

// Configuration returns the configuration (args) for [RedisLinkedServer].
func (rls *RedisLinkedServer) Configuration() interface{} {
	return rls.Args
}

// DependOn is used for other resources to depend on [RedisLinkedServer].
func (rls *RedisLinkedServer) DependOn() terra.Reference {
	return terra.ReferenceResource(rls)
}

// Dependencies returns the list of resources [RedisLinkedServer] depends_on.
func (rls *RedisLinkedServer) Dependencies() terra.Dependencies {
	return rls.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedisLinkedServer].
func (rls *RedisLinkedServer) LifecycleManagement() *terra.Lifecycle {
	return rls.Lifecycle
}

// Attributes returns the attributes for [RedisLinkedServer].
func (rls *RedisLinkedServer) Attributes() redisLinkedServerAttributes {
	return redisLinkedServerAttributes{ref: terra.ReferenceResource(rls)}
}

// ImportState imports the given attribute values into [RedisLinkedServer]'s state.
func (rls *RedisLinkedServer) ImportState(av io.Reader) error {
	rls.state = &redisLinkedServerState{}
	if err := json.NewDecoder(av).Decode(rls.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rls.Type(), rls.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedisLinkedServer] has state.
func (rls *RedisLinkedServer) State() (*redisLinkedServerState, bool) {
	return rls.state, rls.state != nil
}

// StateMust returns the state for [RedisLinkedServer]. Panics if the state is nil.
func (rls *RedisLinkedServer) StateMust() *redisLinkedServerState {
	if rls.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rls.Type(), rls.LocalName()))
	}
	return rls.state
}

// RedisLinkedServerArgs contains the configurations for azurerm_redis_linked_server.
type RedisLinkedServerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkedRedisCacheId: string, required
	LinkedRedisCacheId terra.StringValue `hcl:"linked_redis_cache_id,attr" validate:"required"`
	// LinkedRedisCacheLocation: string, required
	LinkedRedisCacheLocation terra.StringValue `hcl:"linked_redis_cache_location,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServerRole: string, required
	ServerRole terra.StringValue `hcl:"server_role,attr" validate:"required"`
	// TargetRedisCacheName: string, required
	TargetRedisCacheName terra.StringValue `hcl:"target_redis_cache_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *redislinkedserver.Timeouts `hcl:"timeouts,block"`
}
type redisLinkedServerAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_redis_linked_server.
func (rls redisLinkedServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rls.ref.Append("id"))
}

// LinkedRedisCacheId returns a reference to field linked_redis_cache_id of azurerm_redis_linked_server.
func (rls redisLinkedServerAttributes) LinkedRedisCacheId() terra.StringValue {
	return terra.ReferenceAsString(rls.ref.Append("linked_redis_cache_id"))
}

// LinkedRedisCacheLocation returns a reference to field linked_redis_cache_location of azurerm_redis_linked_server.
func (rls redisLinkedServerAttributes) LinkedRedisCacheLocation() terra.StringValue {
	return terra.ReferenceAsString(rls.ref.Append("linked_redis_cache_location"))
}

// Name returns a reference to field name of azurerm_redis_linked_server.
func (rls redisLinkedServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rls.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_redis_linked_server.
func (rls redisLinkedServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(rls.ref.Append("resource_group_name"))
}

// ServerRole returns a reference to field server_role of azurerm_redis_linked_server.
func (rls redisLinkedServerAttributes) ServerRole() terra.StringValue {
	return terra.ReferenceAsString(rls.ref.Append("server_role"))
}

// TargetRedisCacheName returns a reference to field target_redis_cache_name of azurerm_redis_linked_server.
func (rls redisLinkedServerAttributes) TargetRedisCacheName() terra.StringValue {
	return terra.ReferenceAsString(rls.ref.Append("target_redis_cache_name"))
}

func (rls redisLinkedServerAttributes) Timeouts() redislinkedserver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[redislinkedserver.TimeoutsAttributes](rls.ref.Append("timeouts"))
}

type redisLinkedServerState struct {
	Id                       string                           `json:"id"`
	LinkedRedisCacheId       string                           `json:"linked_redis_cache_id"`
	LinkedRedisCacheLocation string                           `json:"linked_redis_cache_location"`
	Name                     string                           `json:"name"`
	ResourceGroupName        string                           `json:"resource_group_name"`
	ServerRole               string                           `json:"server_role"`
	TargetRedisCacheName     string                           `json:"target_redis_cache_name"`
	Timeouts                 *redislinkedserver.TimeoutsState `json:"timeouts"`
}
