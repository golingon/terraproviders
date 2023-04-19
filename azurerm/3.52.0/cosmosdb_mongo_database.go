// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbmongodatabase "github.com/golingon/terraproviders/azurerm/3.52.0/cosmosdbmongodatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbMongoDatabase creates a new instance of [CosmosdbMongoDatabase].
func NewCosmosdbMongoDatabase(name string, args CosmosdbMongoDatabaseArgs) *CosmosdbMongoDatabase {
	return &CosmosdbMongoDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbMongoDatabase)(nil)

// CosmosdbMongoDatabase represents the Terraform resource azurerm_cosmosdb_mongo_database.
type CosmosdbMongoDatabase struct {
	Name      string
	Args      CosmosdbMongoDatabaseArgs
	state     *cosmosdbMongoDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbMongoDatabase].
func (cmd *CosmosdbMongoDatabase) Type() string {
	return "azurerm_cosmosdb_mongo_database"
}

// LocalName returns the local name for [CosmosdbMongoDatabase].
func (cmd *CosmosdbMongoDatabase) LocalName() string {
	return cmd.Name
}

// Configuration returns the configuration (args) for [CosmosdbMongoDatabase].
func (cmd *CosmosdbMongoDatabase) Configuration() interface{} {
	return cmd.Args
}

// DependOn is used for other resources to depend on [CosmosdbMongoDatabase].
func (cmd *CosmosdbMongoDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(cmd)
}

// Dependencies returns the list of resources [CosmosdbMongoDatabase] depends_on.
func (cmd *CosmosdbMongoDatabase) Dependencies() terra.Dependencies {
	return cmd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbMongoDatabase].
func (cmd *CosmosdbMongoDatabase) LifecycleManagement() *terra.Lifecycle {
	return cmd.Lifecycle
}

// Attributes returns the attributes for [CosmosdbMongoDatabase].
func (cmd *CosmosdbMongoDatabase) Attributes() cosmosdbMongoDatabaseAttributes {
	return cosmosdbMongoDatabaseAttributes{ref: terra.ReferenceResource(cmd)}
}

// ImportState imports the given attribute values into [CosmosdbMongoDatabase]'s state.
func (cmd *CosmosdbMongoDatabase) ImportState(av io.Reader) error {
	cmd.state = &cosmosdbMongoDatabaseState{}
	if err := json.NewDecoder(av).Decode(cmd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmd.Type(), cmd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbMongoDatabase] has state.
func (cmd *CosmosdbMongoDatabase) State() (*cosmosdbMongoDatabaseState, bool) {
	return cmd.state, cmd.state != nil
}

// StateMust returns the state for [CosmosdbMongoDatabase]. Panics if the state is nil.
func (cmd *CosmosdbMongoDatabase) StateMust() *cosmosdbMongoDatabaseState {
	if cmd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmd.Type(), cmd.LocalName()))
	}
	return cmd.state
}

// CosmosdbMongoDatabaseArgs contains the configurations for azurerm_cosmosdb_mongo_database.
type CosmosdbMongoDatabaseArgs struct {
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
	AutoscaleSettings *cosmosdbmongodatabase.AutoscaleSettings `hcl:"autoscale_settings,block"`
	// Timeouts: optional
	Timeouts *cosmosdbmongodatabase.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbMongoDatabaseAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_cosmosdb_mongo_database.
func (cmd cosmosdbMongoDatabaseAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(cmd.ref.Append("account_name"))
}

// Id returns a reference to field id of azurerm_cosmosdb_mongo_database.
func (cmd cosmosdbMongoDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cmd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_mongo_database.
func (cmd cosmosdbMongoDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cmd.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_mongo_database.
func (cmd cosmosdbMongoDatabaseAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cmd.ref.Append("resource_group_name"))
}

// Throughput returns a reference to field throughput of azurerm_cosmosdb_mongo_database.
func (cmd cosmosdbMongoDatabaseAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(cmd.ref.Append("throughput"))
}

func (cmd cosmosdbMongoDatabaseAttributes) AutoscaleSettings() terra.ListValue[cosmosdbmongodatabase.AutoscaleSettingsAttributes] {
	return terra.ReferenceAsList[cosmosdbmongodatabase.AutoscaleSettingsAttributes](cmd.ref.Append("autoscale_settings"))
}

func (cmd cosmosdbMongoDatabaseAttributes) Timeouts() cosmosdbmongodatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbmongodatabase.TimeoutsAttributes](cmd.ref.Append("timeouts"))
}

type cosmosdbMongoDatabaseState struct {
	AccountName       string                                         `json:"account_name"`
	Id                string                                         `json:"id"`
	Name              string                                         `json:"name"`
	ResourceGroupName string                                         `json:"resource_group_name"`
	Throughput        float64                                        `json:"throughput"`
	AutoscaleSettings []cosmosdbmongodatabase.AutoscaleSettingsState `json:"autoscale_settings"`
	Timeouts          *cosmosdbmongodatabase.TimeoutsState           `json:"timeouts"`
}
