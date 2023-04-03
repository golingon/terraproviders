// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbcassandrakeyspace "github.com/golingon/terraproviders/azurerm/3.49.0/cosmosdbcassandrakeyspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbCassandraKeyspace creates a new instance of [CosmosdbCassandraKeyspace].
func NewCosmosdbCassandraKeyspace(name string, args CosmosdbCassandraKeyspaceArgs) *CosmosdbCassandraKeyspace {
	return &CosmosdbCassandraKeyspace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbCassandraKeyspace)(nil)

// CosmosdbCassandraKeyspace represents the Terraform resource azurerm_cosmosdb_cassandra_keyspace.
type CosmosdbCassandraKeyspace struct {
	Name      string
	Args      CosmosdbCassandraKeyspaceArgs
	state     *cosmosdbCassandraKeyspaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbCassandraKeyspace].
func (cck *CosmosdbCassandraKeyspace) Type() string {
	return "azurerm_cosmosdb_cassandra_keyspace"
}

// LocalName returns the local name for [CosmosdbCassandraKeyspace].
func (cck *CosmosdbCassandraKeyspace) LocalName() string {
	return cck.Name
}

// Configuration returns the configuration (args) for [CosmosdbCassandraKeyspace].
func (cck *CosmosdbCassandraKeyspace) Configuration() interface{} {
	return cck.Args
}

// DependOn is used for other resources to depend on [CosmosdbCassandraKeyspace].
func (cck *CosmosdbCassandraKeyspace) DependOn() terra.Reference {
	return terra.ReferenceResource(cck)
}

// Dependencies returns the list of resources [CosmosdbCassandraKeyspace] depends_on.
func (cck *CosmosdbCassandraKeyspace) Dependencies() terra.Dependencies {
	return cck.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbCassandraKeyspace].
func (cck *CosmosdbCassandraKeyspace) LifecycleManagement() *terra.Lifecycle {
	return cck.Lifecycle
}

// Attributes returns the attributes for [CosmosdbCassandraKeyspace].
func (cck *CosmosdbCassandraKeyspace) Attributes() cosmosdbCassandraKeyspaceAttributes {
	return cosmosdbCassandraKeyspaceAttributes{ref: terra.ReferenceResource(cck)}
}

// ImportState imports the given attribute values into [CosmosdbCassandraKeyspace]'s state.
func (cck *CosmosdbCassandraKeyspace) ImportState(av io.Reader) error {
	cck.state = &cosmosdbCassandraKeyspaceState{}
	if err := json.NewDecoder(av).Decode(cck.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cck.Type(), cck.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbCassandraKeyspace] has state.
func (cck *CosmosdbCassandraKeyspace) State() (*cosmosdbCassandraKeyspaceState, bool) {
	return cck.state, cck.state != nil
}

// StateMust returns the state for [CosmosdbCassandraKeyspace]. Panics if the state is nil.
func (cck *CosmosdbCassandraKeyspace) StateMust() *cosmosdbCassandraKeyspaceState {
	if cck.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cck.Type(), cck.LocalName()))
	}
	return cck.state
}

// CosmosdbCassandraKeyspaceArgs contains the configurations for azurerm_cosmosdb_cassandra_keyspace.
type CosmosdbCassandraKeyspaceArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Throughput: number, optional
	Throughput terra.NumberValue `hcl:"throughput,attr"`
	// AutoscaleSettings: optional
	AutoscaleSettings *cosmosdbcassandrakeyspace.AutoscaleSettings `hcl:"autoscale_settings,block"`
	// Timeouts: optional
	Timeouts *cosmosdbcassandrakeyspace.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbCassandraKeyspaceAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_cosmosdb_cassandra_keyspace.
func (cck cosmosdbCassandraKeyspaceAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(cck.ref.Append("account_name"))
}

// Id returns a reference to field id of azurerm_cosmosdb_cassandra_keyspace.
func (cck cosmosdbCassandraKeyspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cck.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_cassandra_keyspace.
func (cck cosmosdbCassandraKeyspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cck.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_cassandra_keyspace.
func (cck cosmosdbCassandraKeyspaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cck.ref.Append("resource_group_name"))
}

// Throughput returns a reference to field throughput of azurerm_cosmosdb_cassandra_keyspace.
func (cck cosmosdbCassandraKeyspaceAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(cck.ref.Append("throughput"))
}

func (cck cosmosdbCassandraKeyspaceAttributes) AutoscaleSettings() terra.ListValue[cosmosdbcassandrakeyspace.AutoscaleSettingsAttributes] {
	return terra.ReferenceAsList[cosmosdbcassandrakeyspace.AutoscaleSettingsAttributes](cck.ref.Append("autoscale_settings"))
}

func (cck cosmosdbCassandraKeyspaceAttributes) Timeouts() cosmosdbcassandrakeyspace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbcassandrakeyspace.TimeoutsAttributes](cck.ref.Append("timeouts"))
}

type cosmosdbCassandraKeyspaceState struct {
	AccountName       string                                             `json:"account_name"`
	Id                string                                             `json:"id"`
	Name              string                                             `json:"name"`
	ResourceGroupName string                                             `json:"resource_group_name"`
	Throughput        float64                                            `json:"throughput"`
	AutoscaleSettings []cosmosdbcassandrakeyspace.AutoscaleSettingsState `json:"autoscale_settings"`
	Timeouts          *cosmosdbcassandrakeyspace.TimeoutsState           `json:"timeouts"`
}
