// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	elasticacheglobalreplicationgroup "github.com/golingon/terraproviders/aws/4.63.0/elasticacheglobalreplicationgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewElasticacheGlobalReplicationGroup creates a new instance of [ElasticacheGlobalReplicationGroup].
func NewElasticacheGlobalReplicationGroup(name string, args ElasticacheGlobalReplicationGroupArgs) *ElasticacheGlobalReplicationGroup {
	return &ElasticacheGlobalReplicationGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElasticacheGlobalReplicationGroup)(nil)

// ElasticacheGlobalReplicationGroup represents the Terraform resource aws_elasticache_global_replication_group.
type ElasticacheGlobalReplicationGroup struct {
	Name      string
	Args      ElasticacheGlobalReplicationGroupArgs
	state     *elasticacheGlobalReplicationGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ElasticacheGlobalReplicationGroup].
func (egrg *ElasticacheGlobalReplicationGroup) Type() string {
	return "aws_elasticache_global_replication_group"
}

// LocalName returns the local name for [ElasticacheGlobalReplicationGroup].
func (egrg *ElasticacheGlobalReplicationGroup) LocalName() string {
	return egrg.Name
}

// Configuration returns the configuration (args) for [ElasticacheGlobalReplicationGroup].
func (egrg *ElasticacheGlobalReplicationGroup) Configuration() interface{} {
	return egrg.Args
}

// DependOn is used for other resources to depend on [ElasticacheGlobalReplicationGroup].
func (egrg *ElasticacheGlobalReplicationGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(egrg)
}

// Dependencies returns the list of resources [ElasticacheGlobalReplicationGroup] depends_on.
func (egrg *ElasticacheGlobalReplicationGroup) Dependencies() terra.Dependencies {
	return egrg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ElasticacheGlobalReplicationGroup].
func (egrg *ElasticacheGlobalReplicationGroup) LifecycleManagement() *terra.Lifecycle {
	return egrg.Lifecycle
}

// Attributes returns the attributes for [ElasticacheGlobalReplicationGroup].
func (egrg *ElasticacheGlobalReplicationGroup) Attributes() elasticacheGlobalReplicationGroupAttributes {
	return elasticacheGlobalReplicationGroupAttributes{ref: terra.ReferenceResource(egrg)}
}

// ImportState imports the given attribute values into [ElasticacheGlobalReplicationGroup]'s state.
func (egrg *ElasticacheGlobalReplicationGroup) ImportState(av io.Reader) error {
	egrg.state = &elasticacheGlobalReplicationGroupState{}
	if err := json.NewDecoder(av).Decode(egrg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", egrg.Type(), egrg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ElasticacheGlobalReplicationGroup] has state.
func (egrg *ElasticacheGlobalReplicationGroup) State() (*elasticacheGlobalReplicationGroupState, bool) {
	return egrg.state, egrg.state != nil
}

// StateMust returns the state for [ElasticacheGlobalReplicationGroup]. Panics if the state is nil.
func (egrg *ElasticacheGlobalReplicationGroup) StateMust() *elasticacheGlobalReplicationGroupState {
	if egrg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", egrg.Type(), egrg.LocalName()))
	}
	return egrg.state
}

// ElasticacheGlobalReplicationGroupArgs contains the configurations for aws_elasticache_global_replication_group.
type ElasticacheGlobalReplicationGroupArgs struct {
	// AutomaticFailoverEnabled: bool, optional
	AutomaticFailoverEnabled terra.BoolValue `hcl:"automatic_failover_enabled,attr"`
	// CacheNodeType: string, optional
	CacheNodeType terra.StringValue `hcl:"cache_node_type,attr"`
	// EngineVersion: string, optional
	EngineVersion terra.StringValue `hcl:"engine_version,attr"`
	// GlobalReplicationGroupDescription: string, optional
	GlobalReplicationGroupDescription terra.StringValue `hcl:"global_replication_group_description,attr"`
	// GlobalReplicationGroupIdSuffix: string, required
	GlobalReplicationGroupIdSuffix terra.StringValue `hcl:"global_replication_group_id_suffix,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NumNodeGroups: number, optional
	NumNodeGroups terra.NumberValue `hcl:"num_node_groups,attr"`
	// ParameterGroupName: string, optional
	ParameterGroupName terra.StringValue `hcl:"parameter_group_name,attr"`
	// PrimaryReplicationGroupId: string, required
	PrimaryReplicationGroupId terra.StringValue `hcl:"primary_replication_group_id,attr" validate:"required"`
	// GlobalNodeGroups: min=0
	GlobalNodeGroups []elasticacheglobalreplicationgroup.GlobalNodeGroups `hcl:"global_node_groups,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *elasticacheglobalreplicationgroup.Timeouts `hcl:"timeouts,block"`
}
type elasticacheGlobalReplicationGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(egrg.ref.Append("arn"))
}

// AtRestEncryptionEnabled returns a reference to field at_rest_encryption_enabled of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) AtRestEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(egrg.ref.Append("at_rest_encryption_enabled"))
}

// AuthTokenEnabled returns a reference to field auth_token_enabled of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) AuthTokenEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(egrg.ref.Append("auth_token_enabled"))
}

// AutomaticFailoverEnabled returns a reference to field automatic_failover_enabled of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) AutomaticFailoverEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(egrg.ref.Append("automatic_failover_enabled"))
}

// CacheNodeType returns a reference to field cache_node_type of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) CacheNodeType() terra.StringValue {
	return terra.ReferenceAsString(egrg.ref.Append("cache_node_type"))
}

// ClusterEnabled returns a reference to field cluster_enabled of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) ClusterEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(egrg.ref.Append("cluster_enabled"))
}

// Engine returns a reference to field engine of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(egrg.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(egrg.ref.Append("engine_version"))
}

// EngineVersionActual returns a reference to field engine_version_actual of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) EngineVersionActual() terra.StringValue {
	return terra.ReferenceAsString(egrg.ref.Append("engine_version_actual"))
}

// GlobalReplicationGroupDescription returns a reference to field global_replication_group_description of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) GlobalReplicationGroupDescription() terra.StringValue {
	return terra.ReferenceAsString(egrg.ref.Append("global_replication_group_description"))
}

// GlobalReplicationGroupId returns a reference to field global_replication_group_id of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) GlobalReplicationGroupId() terra.StringValue {
	return terra.ReferenceAsString(egrg.ref.Append("global_replication_group_id"))
}

// GlobalReplicationGroupIdSuffix returns a reference to field global_replication_group_id_suffix of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) GlobalReplicationGroupIdSuffix() terra.StringValue {
	return terra.ReferenceAsString(egrg.ref.Append("global_replication_group_id_suffix"))
}

// Id returns a reference to field id of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(egrg.ref.Append("id"))
}

// NumNodeGroups returns a reference to field num_node_groups of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) NumNodeGroups() terra.NumberValue {
	return terra.ReferenceAsNumber(egrg.ref.Append("num_node_groups"))
}

// ParameterGroupName returns a reference to field parameter_group_name of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) ParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(egrg.ref.Append("parameter_group_name"))
}

// PrimaryReplicationGroupId returns a reference to field primary_replication_group_id of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) PrimaryReplicationGroupId() terra.StringValue {
	return terra.ReferenceAsString(egrg.ref.Append("primary_replication_group_id"))
}

// TransitEncryptionEnabled returns a reference to field transit_encryption_enabled of aws_elasticache_global_replication_group.
func (egrg elasticacheGlobalReplicationGroupAttributes) TransitEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(egrg.ref.Append("transit_encryption_enabled"))
}

func (egrg elasticacheGlobalReplicationGroupAttributes) GlobalNodeGroups() terra.SetValue[elasticacheglobalreplicationgroup.GlobalNodeGroupsAttributes] {
	return terra.ReferenceAsSet[elasticacheglobalreplicationgroup.GlobalNodeGroupsAttributes](egrg.ref.Append("global_node_groups"))
}

func (egrg elasticacheGlobalReplicationGroupAttributes) Timeouts() elasticacheglobalreplicationgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[elasticacheglobalreplicationgroup.TimeoutsAttributes](egrg.ref.Append("timeouts"))
}

type elasticacheGlobalReplicationGroupState struct {
	Arn                               string                                                    `json:"arn"`
	AtRestEncryptionEnabled           bool                                                      `json:"at_rest_encryption_enabled"`
	AuthTokenEnabled                  bool                                                      `json:"auth_token_enabled"`
	AutomaticFailoverEnabled          bool                                                      `json:"automatic_failover_enabled"`
	CacheNodeType                     string                                                    `json:"cache_node_type"`
	ClusterEnabled                    bool                                                      `json:"cluster_enabled"`
	Engine                            string                                                    `json:"engine"`
	EngineVersion                     string                                                    `json:"engine_version"`
	EngineVersionActual               string                                                    `json:"engine_version_actual"`
	GlobalReplicationGroupDescription string                                                    `json:"global_replication_group_description"`
	GlobalReplicationGroupId          string                                                    `json:"global_replication_group_id"`
	GlobalReplicationGroupIdSuffix    string                                                    `json:"global_replication_group_id_suffix"`
	Id                                string                                                    `json:"id"`
	NumNodeGroups                     float64                                                   `json:"num_node_groups"`
	ParameterGroupName                string                                                    `json:"parameter_group_name"`
	PrimaryReplicationGroupId         string                                                    `json:"primary_replication_group_id"`
	TransitEncryptionEnabled          bool                                                      `json:"transit_encryption_enabled"`
	GlobalNodeGroups                  []elasticacheglobalreplicationgroup.GlobalNodeGroupsState `json:"global_node_groups"`
	Timeouts                          *elasticacheglobalreplicationgroup.TimeoutsState          `json:"timeouts"`
}
