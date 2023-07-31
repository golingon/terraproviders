// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbgremlingraph "github.com/golingon/terraproviders/azurerm/3.67.0/cosmosdbgremlingraph"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbGremlinGraph creates a new instance of [CosmosdbGremlinGraph].
func NewCosmosdbGremlinGraph(name string, args CosmosdbGremlinGraphArgs) *CosmosdbGremlinGraph {
	return &CosmosdbGremlinGraph{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbGremlinGraph)(nil)

// CosmosdbGremlinGraph represents the Terraform resource azurerm_cosmosdb_gremlin_graph.
type CosmosdbGremlinGraph struct {
	Name      string
	Args      CosmosdbGremlinGraphArgs
	state     *cosmosdbGremlinGraphState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbGremlinGraph].
func (cgg *CosmosdbGremlinGraph) Type() string {
	return "azurerm_cosmosdb_gremlin_graph"
}

// LocalName returns the local name for [CosmosdbGremlinGraph].
func (cgg *CosmosdbGremlinGraph) LocalName() string {
	return cgg.Name
}

// Configuration returns the configuration (args) for [CosmosdbGremlinGraph].
func (cgg *CosmosdbGremlinGraph) Configuration() interface{} {
	return cgg.Args
}

// DependOn is used for other resources to depend on [CosmosdbGremlinGraph].
func (cgg *CosmosdbGremlinGraph) DependOn() terra.Reference {
	return terra.ReferenceResource(cgg)
}

// Dependencies returns the list of resources [CosmosdbGremlinGraph] depends_on.
func (cgg *CosmosdbGremlinGraph) Dependencies() terra.Dependencies {
	return cgg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbGremlinGraph].
func (cgg *CosmosdbGremlinGraph) LifecycleManagement() *terra.Lifecycle {
	return cgg.Lifecycle
}

// Attributes returns the attributes for [CosmosdbGremlinGraph].
func (cgg *CosmosdbGremlinGraph) Attributes() cosmosdbGremlinGraphAttributes {
	return cosmosdbGremlinGraphAttributes{ref: terra.ReferenceResource(cgg)}
}

// ImportState imports the given attribute values into [CosmosdbGremlinGraph]'s state.
func (cgg *CosmosdbGremlinGraph) ImportState(av io.Reader) error {
	cgg.state = &cosmosdbGremlinGraphState{}
	if err := json.NewDecoder(av).Decode(cgg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cgg.Type(), cgg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbGremlinGraph] has state.
func (cgg *CosmosdbGremlinGraph) State() (*cosmosdbGremlinGraphState, bool) {
	return cgg.state, cgg.state != nil
}

// StateMust returns the state for [CosmosdbGremlinGraph]. Panics if the state is nil.
func (cgg *CosmosdbGremlinGraph) StateMust() *cosmosdbGremlinGraphState {
	if cgg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cgg.Type(), cgg.LocalName()))
	}
	return cgg.state
}

// CosmosdbGremlinGraphArgs contains the configurations for azurerm_cosmosdb_gremlin_graph.
type CosmosdbGremlinGraphArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// AnalyticalStorageTtl: number, optional
	AnalyticalStorageTtl terra.NumberValue `hcl:"analytical_storage_ttl,attr"`
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// DefaultTtl: number, optional
	DefaultTtl terra.NumberValue `hcl:"default_ttl,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PartitionKeyPath: string, required
	PartitionKeyPath terra.StringValue `hcl:"partition_key_path,attr" validate:"required"`
	// PartitionKeyVersion: number, optional
	PartitionKeyVersion terra.NumberValue `hcl:"partition_key_version,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Throughput: number, optional
	Throughput terra.NumberValue `hcl:"throughput,attr"`
	// AutoscaleSettings: optional
	AutoscaleSettings *cosmosdbgremlingraph.AutoscaleSettings `hcl:"autoscale_settings,block"`
	// ConflictResolutionPolicy: optional
	ConflictResolutionPolicy *cosmosdbgremlingraph.ConflictResolutionPolicy `hcl:"conflict_resolution_policy,block"`
	// IndexPolicy: optional
	IndexPolicy *cosmosdbgremlingraph.IndexPolicy `hcl:"index_policy,block"`
	// Timeouts: optional
	Timeouts *cosmosdbgremlingraph.Timeouts `hcl:"timeouts,block"`
	// UniqueKey: min=0
	UniqueKey []cosmosdbgremlingraph.UniqueKey `hcl:"unique_key,block" validate:"min=0"`
}
type cosmosdbGremlinGraphAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_cosmosdb_gremlin_graph.
func (cgg cosmosdbGremlinGraphAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(cgg.ref.Append("account_name"))
}

// AnalyticalStorageTtl returns a reference to field analytical_storage_ttl of azurerm_cosmosdb_gremlin_graph.
func (cgg cosmosdbGremlinGraphAttributes) AnalyticalStorageTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(cgg.ref.Append("analytical_storage_ttl"))
}

// DatabaseName returns a reference to field database_name of azurerm_cosmosdb_gremlin_graph.
func (cgg cosmosdbGremlinGraphAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(cgg.ref.Append("database_name"))
}

// DefaultTtl returns a reference to field default_ttl of azurerm_cosmosdb_gremlin_graph.
func (cgg cosmosdbGremlinGraphAttributes) DefaultTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(cgg.ref.Append("default_ttl"))
}

// Id returns a reference to field id of azurerm_cosmosdb_gremlin_graph.
func (cgg cosmosdbGremlinGraphAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cgg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_gremlin_graph.
func (cgg cosmosdbGremlinGraphAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cgg.ref.Append("name"))
}

// PartitionKeyPath returns a reference to field partition_key_path of azurerm_cosmosdb_gremlin_graph.
func (cgg cosmosdbGremlinGraphAttributes) PartitionKeyPath() terra.StringValue {
	return terra.ReferenceAsString(cgg.ref.Append("partition_key_path"))
}

// PartitionKeyVersion returns a reference to field partition_key_version of azurerm_cosmosdb_gremlin_graph.
func (cgg cosmosdbGremlinGraphAttributes) PartitionKeyVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(cgg.ref.Append("partition_key_version"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_gremlin_graph.
func (cgg cosmosdbGremlinGraphAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cgg.ref.Append("resource_group_name"))
}

// Throughput returns a reference to field throughput of azurerm_cosmosdb_gremlin_graph.
func (cgg cosmosdbGremlinGraphAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(cgg.ref.Append("throughput"))
}

func (cgg cosmosdbGremlinGraphAttributes) AutoscaleSettings() terra.ListValue[cosmosdbgremlingraph.AutoscaleSettingsAttributes] {
	return terra.ReferenceAsList[cosmosdbgremlingraph.AutoscaleSettingsAttributes](cgg.ref.Append("autoscale_settings"))
}

func (cgg cosmosdbGremlinGraphAttributes) ConflictResolutionPolicy() terra.ListValue[cosmosdbgremlingraph.ConflictResolutionPolicyAttributes] {
	return terra.ReferenceAsList[cosmosdbgremlingraph.ConflictResolutionPolicyAttributes](cgg.ref.Append("conflict_resolution_policy"))
}

func (cgg cosmosdbGremlinGraphAttributes) IndexPolicy() terra.ListValue[cosmosdbgremlingraph.IndexPolicyAttributes] {
	return terra.ReferenceAsList[cosmosdbgremlingraph.IndexPolicyAttributes](cgg.ref.Append("index_policy"))
}

func (cgg cosmosdbGremlinGraphAttributes) Timeouts() cosmosdbgremlingraph.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbgremlingraph.TimeoutsAttributes](cgg.ref.Append("timeouts"))
}

func (cgg cosmosdbGremlinGraphAttributes) UniqueKey() terra.SetValue[cosmosdbgremlingraph.UniqueKeyAttributes] {
	return terra.ReferenceAsSet[cosmosdbgremlingraph.UniqueKeyAttributes](cgg.ref.Append("unique_key"))
}

type cosmosdbGremlinGraphState struct {
	AccountName              string                                               `json:"account_name"`
	AnalyticalStorageTtl     float64                                              `json:"analytical_storage_ttl"`
	DatabaseName             string                                               `json:"database_name"`
	DefaultTtl               float64                                              `json:"default_ttl"`
	Id                       string                                               `json:"id"`
	Name                     string                                               `json:"name"`
	PartitionKeyPath         string                                               `json:"partition_key_path"`
	PartitionKeyVersion      float64                                              `json:"partition_key_version"`
	ResourceGroupName        string                                               `json:"resource_group_name"`
	Throughput               float64                                              `json:"throughput"`
	AutoscaleSettings        []cosmosdbgremlingraph.AutoscaleSettingsState        `json:"autoscale_settings"`
	ConflictResolutionPolicy []cosmosdbgremlingraph.ConflictResolutionPolicyState `json:"conflict_resolution_policy"`
	IndexPolicy              []cosmosdbgremlingraph.IndexPolicyState              `json:"index_policy"`
	Timeouts                 *cosmosdbgremlingraph.TimeoutsState                  `json:"timeouts"`
	UniqueKey                []cosmosdbgremlingraph.UniqueKeyState                `json:"unique_key"`
}
