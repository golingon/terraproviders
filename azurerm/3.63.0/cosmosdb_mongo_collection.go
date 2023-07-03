// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbmongocollection "github.com/golingon/terraproviders/azurerm/3.63.0/cosmosdbmongocollection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbMongoCollection creates a new instance of [CosmosdbMongoCollection].
func NewCosmosdbMongoCollection(name string, args CosmosdbMongoCollectionArgs) *CosmosdbMongoCollection {
	return &CosmosdbMongoCollection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbMongoCollection)(nil)

// CosmosdbMongoCollection represents the Terraform resource azurerm_cosmosdb_mongo_collection.
type CosmosdbMongoCollection struct {
	Name      string
	Args      CosmosdbMongoCollectionArgs
	state     *cosmosdbMongoCollectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbMongoCollection].
func (cmc *CosmosdbMongoCollection) Type() string {
	return "azurerm_cosmosdb_mongo_collection"
}

// LocalName returns the local name for [CosmosdbMongoCollection].
func (cmc *CosmosdbMongoCollection) LocalName() string {
	return cmc.Name
}

// Configuration returns the configuration (args) for [CosmosdbMongoCollection].
func (cmc *CosmosdbMongoCollection) Configuration() interface{} {
	return cmc.Args
}

// DependOn is used for other resources to depend on [CosmosdbMongoCollection].
func (cmc *CosmosdbMongoCollection) DependOn() terra.Reference {
	return terra.ReferenceResource(cmc)
}

// Dependencies returns the list of resources [CosmosdbMongoCollection] depends_on.
func (cmc *CosmosdbMongoCollection) Dependencies() terra.Dependencies {
	return cmc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbMongoCollection].
func (cmc *CosmosdbMongoCollection) LifecycleManagement() *terra.Lifecycle {
	return cmc.Lifecycle
}

// Attributes returns the attributes for [CosmosdbMongoCollection].
func (cmc *CosmosdbMongoCollection) Attributes() cosmosdbMongoCollectionAttributes {
	return cosmosdbMongoCollectionAttributes{ref: terra.ReferenceResource(cmc)}
}

// ImportState imports the given attribute values into [CosmosdbMongoCollection]'s state.
func (cmc *CosmosdbMongoCollection) ImportState(av io.Reader) error {
	cmc.state = &cosmosdbMongoCollectionState{}
	if err := json.NewDecoder(av).Decode(cmc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmc.Type(), cmc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbMongoCollection] has state.
func (cmc *CosmosdbMongoCollection) State() (*cosmosdbMongoCollectionState, bool) {
	return cmc.state, cmc.state != nil
}

// StateMust returns the state for [CosmosdbMongoCollection]. Panics if the state is nil.
func (cmc *CosmosdbMongoCollection) StateMust() *cosmosdbMongoCollectionState {
	if cmc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmc.Type(), cmc.LocalName()))
	}
	return cmc.state
}

// CosmosdbMongoCollectionArgs contains the configurations for azurerm_cosmosdb_mongo_collection.
type CosmosdbMongoCollectionArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// AnalyticalStorageTtl: number, optional
	AnalyticalStorageTtl terra.NumberValue `hcl:"analytical_storage_ttl,attr"`
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// DefaultTtlSeconds: number, optional
	DefaultTtlSeconds terra.NumberValue `hcl:"default_ttl_seconds,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ShardKey: string, optional
	ShardKey terra.StringValue `hcl:"shard_key,attr"`
	// Throughput: number, optional
	Throughput terra.NumberValue `hcl:"throughput,attr"`
	// SystemIndexes: min=0
	SystemIndexes []cosmosdbmongocollection.SystemIndexes `hcl:"system_indexes,block" validate:"min=0"`
	// AutoscaleSettings: optional
	AutoscaleSettings *cosmosdbmongocollection.AutoscaleSettings `hcl:"autoscale_settings,block"`
	// Index: min=0
	Index []cosmosdbmongocollection.Index `hcl:"index,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *cosmosdbmongocollection.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbMongoCollectionAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_cosmosdb_mongo_collection.
func (cmc cosmosdbMongoCollectionAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(cmc.ref.Append("account_name"))
}

// AnalyticalStorageTtl returns a reference to field analytical_storage_ttl of azurerm_cosmosdb_mongo_collection.
func (cmc cosmosdbMongoCollectionAttributes) AnalyticalStorageTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(cmc.ref.Append("analytical_storage_ttl"))
}

// DatabaseName returns a reference to field database_name of azurerm_cosmosdb_mongo_collection.
func (cmc cosmosdbMongoCollectionAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(cmc.ref.Append("database_name"))
}

// DefaultTtlSeconds returns a reference to field default_ttl_seconds of azurerm_cosmosdb_mongo_collection.
func (cmc cosmosdbMongoCollectionAttributes) DefaultTtlSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(cmc.ref.Append("default_ttl_seconds"))
}

// Id returns a reference to field id of azurerm_cosmosdb_mongo_collection.
func (cmc cosmosdbMongoCollectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cmc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_mongo_collection.
func (cmc cosmosdbMongoCollectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cmc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_mongo_collection.
func (cmc cosmosdbMongoCollectionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cmc.ref.Append("resource_group_name"))
}

// ShardKey returns a reference to field shard_key of azurerm_cosmosdb_mongo_collection.
func (cmc cosmosdbMongoCollectionAttributes) ShardKey() terra.StringValue {
	return terra.ReferenceAsString(cmc.ref.Append("shard_key"))
}

// Throughput returns a reference to field throughput of azurerm_cosmosdb_mongo_collection.
func (cmc cosmosdbMongoCollectionAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(cmc.ref.Append("throughput"))
}

func (cmc cosmosdbMongoCollectionAttributes) SystemIndexes() terra.ListValue[cosmosdbmongocollection.SystemIndexesAttributes] {
	return terra.ReferenceAsList[cosmosdbmongocollection.SystemIndexesAttributes](cmc.ref.Append("system_indexes"))
}

func (cmc cosmosdbMongoCollectionAttributes) AutoscaleSettings() terra.ListValue[cosmosdbmongocollection.AutoscaleSettingsAttributes] {
	return terra.ReferenceAsList[cosmosdbmongocollection.AutoscaleSettingsAttributes](cmc.ref.Append("autoscale_settings"))
}

func (cmc cosmosdbMongoCollectionAttributes) Index() terra.SetValue[cosmosdbmongocollection.IndexAttributes] {
	return terra.ReferenceAsSet[cosmosdbmongocollection.IndexAttributes](cmc.ref.Append("index"))
}

func (cmc cosmosdbMongoCollectionAttributes) Timeouts() cosmosdbmongocollection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbmongocollection.TimeoutsAttributes](cmc.ref.Append("timeouts"))
}

type cosmosdbMongoCollectionState struct {
	AccountName          string                                           `json:"account_name"`
	AnalyticalStorageTtl float64                                          `json:"analytical_storage_ttl"`
	DatabaseName         string                                           `json:"database_name"`
	DefaultTtlSeconds    float64                                          `json:"default_ttl_seconds"`
	Id                   string                                           `json:"id"`
	Name                 string                                           `json:"name"`
	ResourceGroupName    string                                           `json:"resource_group_name"`
	ShardKey             string                                           `json:"shard_key"`
	Throughput           float64                                          `json:"throughput"`
	SystemIndexes        []cosmosdbmongocollection.SystemIndexesState     `json:"system_indexes"`
	AutoscaleSettings    []cosmosdbmongocollection.AutoscaleSettingsState `json:"autoscale_settings"`
	Index                []cosmosdbmongocollection.IndexState             `json:"index"`
	Timeouts             *cosmosdbmongocollection.TimeoutsState           `json:"timeouts"`
}
