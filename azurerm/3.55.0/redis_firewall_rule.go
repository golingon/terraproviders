// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	redisfirewallrule "github.com/golingon/terraproviders/azurerm/3.55.0/redisfirewallrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRedisFirewallRule creates a new instance of [RedisFirewallRule].
func NewRedisFirewallRule(name string, args RedisFirewallRuleArgs) *RedisFirewallRule {
	return &RedisFirewallRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedisFirewallRule)(nil)

// RedisFirewallRule represents the Terraform resource azurerm_redis_firewall_rule.
type RedisFirewallRule struct {
	Name      string
	Args      RedisFirewallRuleArgs
	state     *redisFirewallRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedisFirewallRule].
func (rfr *RedisFirewallRule) Type() string {
	return "azurerm_redis_firewall_rule"
}

// LocalName returns the local name for [RedisFirewallRule].
func (rfr *RedisFirewallRule) LocalName() string {
	return rfr.Name
}

// Configuration returns the configuration (args) for [RedisFirewallRule].
func (rfr *RedisFirewallRule) Configuration() interface{} {
	return rfr.Args
}

// DependOn is used for other resources to depend on [RedisFirewallRule].
func (rfr *RedisFirewallRule) DependOn() terra.Reference {
	return terra.ReferenceResource(rfr)
}

// Dependencies returns the list of resources [RedisFirewallRule] depends_on.
func (rfr *RedisFirewallRule) Dependencies() terra.Dependencies {
	return rfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedisFirewallRule].
func (rfr *RedisFirewallRule) LifecycleManagement() *terra.Lifecycle {
	return rfr.Lifecycle
}

// Attributes returns the attributes for [RedisFirewallRule].
func (rfr *RedisFirewallRule) Attributes() redisFirewallRuleAttributes {
	return redisFirewallRuleAttributes{ref: terra.ReferenceResource(rfr)}
}

// ImportState imports the given attribute values into [RedisFirewallRule]'s state.
func (rfr *RedisFirewallRule) ImportState(av io.Reader) error {
	rfr.state = &redisFirewallRuleState{}
	if err := json.NewDecoder(av).Decode(rfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rfr.Type(), rfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedisFirewallRule] has state.
func (rfr *RedisFirewallRule) State() (*redisFirewallRuleState, bool) {
	return rfr.state, rfr.state != nil
}

// StateMust returns the state for [RedisFirewallRule]. Panics if the state is nil.
func (rfr *RedisFirewallRule) StateMust() *redisFirewallRuleState {
	if rfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rfr.Type(), rfr.LocalName()))
	}
	return rfr.state
}

// RedisFirewallRuleArgs contains the configurations for azurerm_redis_firewall_rule.
type RedisFirewallRuleArgs struct {
	// EndIp: string, required
	EndIp terra.StringValue `hcl:"end_ip,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RedisCacheName: string, required
	RedisCacheName terra.StringValue `hcl:"redis_cache_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StartIp: string, required
	StartIp terra.StringValue `hcl:"start_ip,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *redisfirewallrule.Timeouts `hcl:"timeouts,block"`
}
type redisFirewallRuleAttributes struct {
	ref terra.Reference
}

// EndIp returns a reference to field end_ip of azurerm_redis_firewall_rule.
func (rfr redisFirewallRuleAttributes) EndIp() terra.StringValue {
	return terra.ReferenceAsString(rfr.ref.Append("end_ip"))
}

// Id returns a reference to field id of azurerm_redis_firewall_rule.
func (rfr redisFirewallRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rfr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_redis_firewall_rule.
func (rfr redisFirewallRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rfr.ref.Append("name"))
}

// RedisCacheName returns a reference to field redis_cache_name of azurerm_redis_firewall_rule.
func (rfr redisFirewallRuleAttributes) RedisCacheName() terra.StringValue {
	return terra.ReferenceAsString(rfr.ref.Append("redis_cache_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_redis_firewall_rule.
func (rfr redisFirewallRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(rfr.ref.Append("resource_group_name"))
}

// StartIp returns a reference to field start_ip of azurerm_redis_firewall_rule.
func (rfr redisFirewallRuleAttributes) StartIp() terra.StringValue {
	return terra.ReferenceAsString(rfr.ref.Append("start_ip"))
}

func (rfr redisFirewallRuleAttributes) Timeouts() redisfirewallrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[redisfirewallrule.TimeoutsAttributes](rfr.ref.Append("timeouts"))
}

type redisFirewallRuleState struct {
	EndIp             string                           `json:"end_ip"`
	Id                string                           `json:"id"`
	Name              string                           `json:"name"`
	RedisCacheName    string                           `json:"redis_cache_name"`
	ResourceGroupName string                           `json:"resource_group_name"`
	StartIp           string                           `json:"start_ip"`
	Timeouts          *redisfirewallrule.TimeoutsState `json:"timeouts"`
}
