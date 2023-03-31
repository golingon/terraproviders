// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbmongocollection "github.com/golingon/terraproviders/azurerm/3.49.0/cosmosdbmongocollection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewCosmosdbMongoCollection(name string, args CosmosdbMongoCollectionArgs) *CosmosdbMongoCollection {
	return &CosmosdbMongoCollection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbMongoCollection)(nil)

type CosmosdbMongoCollection struct {
	Name  string
	Args  CosmosdbMongoCollectionArgs
	state *cosmosdbMongoCollectionState
}

func (cmc *CosmosdbMongoCollection) Type() string {
	return "azurerm_cosmosdb_mongo_collection"
}

func (cmc *CosmosdbMongoCollection) LocalName() string {
	return cmc.Name
}

func (cmc *CosmosdbMongoCollection) Configuration() interface{} {
	return cmc.Args
}

func (cmc *CosmosdbMongoCollection) Attributes() cosmosdbMongoCollectionAttributes {
	return cosmosdbMongoCollectionAttributes{ref: terra.ReferenceResource(cmc)}
}

func (cmc *CosmosdbMongoCollection) ImportState(av io.Reader) error {
	cmc.state = &cosmosdbMongoCollectionState{}
	if err := json.NewDecoder(av).Decode(cmc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmc.Type(), cmc.LocalName(), err)
	}
	return nil
}

func (cmc *CosmosdbMongoCollection) State() (*cosmosdbMongoCollectionState, bool) {
	return cmc.state, cmc.state != nil
}

func (cmc *CosmosdbMongoCollection) StateMust() *cosmosdbMongoCollectionState {
	if cmc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmc.Type(), cmc.LocalName()))
	}
	return cmc.state
}

func (cmc *CosmosdbMongoCollection) DependOn() terra.Reference {
	return terra.ReferenceResource(cmc)
}

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
	// DependsOn contains resources that CosmosdbMongoCollection depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type cosmosdbMongoCollectionAttributes struct {
	ref terra.Reference
}

func (cmc cosmosdbMongoCollectionAttributes) AccountName() terra.StringValue {
	return terra.ReferenceString(cmc.ref.Append("account_name"))
}

func (cmc cosmosdbMongoCollectionAttributes) AnalyticalStorageTtl() terra.NumberValue {
	return terra.ReferenceNumber(cmc.ref.Append("analytical_storage_ttl"))
}

func (cmc cosmosdbMongoCollectionAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceString(cmc.ref.Append("database_name"))
}

func (cmc cosmosdbMongoCollectionAttributes) DefaultTtlSeconds() terra.NumberValue {
	return terra.ReferenceNumber(cmc.ref.Append("default_ttl_seconds"))
}

func (cmc cosmosdbMongoCollectionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cmc.ref.Append("id"))
}

func (cmc cosmosdbMongoCollectionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cmc.ref.Append("name"))
}

func (cmc cosmosdbMongoCollectionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(cmc.ref.Append("resource_group_name"))
}

func (cmc cosmosdbMongoCollectionAttributes) ShardKey() terra.StringValue {
	return terra.ReferenceString(cmc.ref.Append("shard_key"))
}

func (cmc cosmosdbMongoCollectionAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceNumber(cmc.ref.Append("throughput"))
}

func (cmc cosmosdbMongoCollectionAttributes) SystemIndexes() terra.ListValue[cosmosdbmongocollection.SystemIndexesAttributes] {
	return terra.ReferenceList[cosmosdbmongocollection.SystemIndexesAttributes](cmc.ref.Append("system_indexes"))
}

func (cmc cosmosdbMongoCollectionAttributes) AutoscaleSettings() terra.ListValue[cosmosdbmongocollection.AutoscaleSettingsAttributes] {
	return terra.ReferenceList[cosmosdbmongocollection.AutoscaleSettingsAttributes](cmc.ref.Append("autoscale_settings"))
}

func (cmc cosmosdbMongoCollectionAttributes) Index() terra.SetValue[cosmosdbmongocollection.IndexAttributes] {
	return terra.ReferenceSet[cosmosdbmongocollection.IndexAttributes](cmc.ref.Append("index"))
}

func (cmc cosmosdbMongoCollectionAttributes) Timeouts() cosmosdbmongocollection.TimeoutsAttributes {
	return terra.ReferenceSingle[cosmosdbmongocollection.TimeoutsAttributes](cmc.ref.Append("timeouts"))
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
