// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbcassandratable "github.com/golingon/terraproviders/azurerm/3.49.0/cosmosdbcassandratable"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbCassandraTable creates a new instance of [CosmosdbCassandraTable].
func NewCosmosdbCassandraTable(name string, args CosmosdbCassandraTableArgs) *CosmosdbCassandraTable {
	return &CosmosdbCassandraTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbCassandraTable)(nil)

// CosmosdbCassandraTable represents the Terraform resource azurerm_cosmosdb_cassandra_table.
type CosmosdbCassandraTable struct {
	Name      string
	Args      CosmosdbCassandraTableArgs
	state     *cosmosdbCassandraTableState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbCassandraTable].
func (cct *CosmosdbCassandraTable) Type() string {
	return "azurerm_cosmosdb_cassandra_table"
}

// LocalName returns the local name for [CosmosdbCassandraTable].
func (cct *CosmosdbCassandraTable) LocalName() string {
	return cct.Name
}

// Configuration returns the configuration (args) for [CosmosdbCassandraTable].
func (cct *CosmosdbCassandraTable) Configuration() interface{} {
	return cct.Args
}

// DependOn is used for other resources to depend on [CosmosdbCassandraTable].
func (cct *CosmosdbCassandraTable) DependOn() terra.Reference {
	return terra.ReferenceResource(cct)
}

// Dependencies returns the list of resources [CosmosdbCassandraTable] depends_on.
func (cct *CosmosdbCassandraTable) Dependencies() terra.Dependencies {
	return cct.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbCassandraTable].
func (cct *CosmosdbCassandraTable) LifecycleManagement() *terra.Lifecycle {
	return cct.Lifecycle
}

// Attributes returns the attributes for [CosmosdbCassandraTable].
func (cct *CosmosdbCassandraTable) Attributes() cosmosdbCassandraTableAttributes {
	return cosmosdbCassandraTableAttributes{ref: terra.ReferenceResource(cct)}
}

// ImportState imports the given attribute values into [CosmosdbCassandraTable]'s state.
func (cct *CosmosdbCassandraTable) ImportState(av io.Reader) error {
	cct.state = &cosmosdbCassandraTableState{}
	if err := json.NewDecoder(av).Decode(cct.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cct.Type(), cct.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbCassandraTable] has state.
func (cct *CosmosdbCassandraTable) State() (*cosmosdbCassandraTableState, bool) {
	return cct.state, cct.state != nil
}

// StateMust returns the state for [CosmosdbCassandraTable]. Panics if the state is nil.
func (cct *CosmosdbCassandraTable) StateMust() *cosmosdbCassandraTableState {
	if cct.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cct.Type(), cct.LocalName()))
	}
	return cct.state
}

// CosmosdbCassandraTableArgs contains the configurations for azurerm_cosmosdb_cassandra_table.
type CosmosdbCassandraTableArgs struct {
	// AnalyticalStorageTtl: number, optional
	AnalyticalStorageTtl terra.NumberValue `hcl:"analytical_storage_ttl,attr"`
	// CassandraKeyspaceId: string, required
	CassandraKeyspaceId terra.StringValue `hcl:"cassandra_keyspace_id,attr" validate:"required"`
	// DefaultTtl: number, optional
	DefaultTtl terra.NumberValue `hcl:"default_ttl,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Throughput: number, optional
	Throughput terra.NumberValue `hcl:"throughput,attr"`
	// AutoscaleSettings: optional
	AutoscaleSettings *cosmosdbcassandratable.AutoscaleSettings `hcl:"autoscale_settings,block"`
	// Schema: required
	Schema *cosmosdbcassandratable.Schema `hcl:"schema,block" validate:"required"`
	// Timeouts: optional
	Timeouts *cosmosdbcassandratable.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbCassandraTableAttributes struct {
	ref terra.Reference
}

// AnalyticalStorageTtl returns a reference to field analytical_storage_ttl of azurerm_cosmosdb_cassandra_table.
func (cct cosmosdbCassandraTableAttributes) AnalyticalStorageTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(cct.ref.Append("analytical_storage_ttl"))
}

// CassandraKeyspaceId returns a reference to field cassandra_keyspace_id of azurerm_cosmosdb_cassandra_table.
func (cct cosmosdbCassandraTableAttributes) CassandraKeyspaceId() terra.StringValue {
	return terra.ReferenceAsString(cct.ref.Append("cassandra_keyspace_id"))
}

// DefaultTtl returns a reference to field default_ttl of azurerm_cosmosdb_cassandra_table.
func (cct cosmosdbCassandraTableAttributes) DefaultTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(cct.ref.Append("default_ttl"))
}

// Id returns a reference to field id of azurerm_cosmosdb_cassandra_table.
func (cct cosmosdbCassandraTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cct.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_cassandra_table.
func (cct cosmosdbCassandraTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cct.ref.Append("name"))
}

// Throughput returns a reference to field throughput of azurerm_cosmosdb_cassandra_table.
func (cct cosmosdbCassandraTableAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(cct.ref.Append("throughput"))
}

func (cct cosmosdbCassandraTableAttributes) AutoscaleSettings() terra.ListValue[cosmosdbcassandratable.AutoscaleSettingsAttributes] {
	return terra.ReferenceAsList[cosmosdbcassandratable.AutoscaleSettingsAttributes](cct.ref.Append("autoscale_settings"))
}

func (cct cosmosdbCassandraTableAttributes) Schema() terra.ListValue[cosmosdbcassandratable.SchemaAttributes] {
	return terra.ReferenceAsList[cosmosdbcassandratable.SchemaAttributes](cct.ref.Append("schema"))
}

func (cct cosmosdbCassandraTableAttributes) Timeouts() cosmosdbcassandratable.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbcassandratable.TimeoutsAttributes](cct.ref.Append("timeouts"))
}

type cosmosdbCassandraTableState struct {
	AnalyticalStorageTtl float64                                         `json:"analytical_storage_ttl"`
	CassandraKeyspaceId  string                                          `json:"cassandra_keyspace_id"`
	DefaultTtl           float64                                         `json:"default_ttl"`
	Id                   string                                          `json:"id"`
	Name                 string                                          `json:"name"`
	Throughput           float64                                         `json:"throughput"`
	AutoscaleSettings    []cosmosdbcassandratable.AutoscaleSettingsState `json:"autoscale_settings"`
	Schema               []cosmosdbcassandratable.SchemaState            `json:"schema"`
	Timeouts             *cosmosdbcassandratable.TimeoutsState           `json:"timeouts"`
}
