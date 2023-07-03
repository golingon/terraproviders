// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbgremlindatabase "github.com/golingon/terraproviders/azurerm/3.63.0/cosmosdbgremlindatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbGremlinDatabase creates a new instance of [CosmosdbGremlinDatabase].
func NewCosmosdbGremlinDatabase(name string, args CosmosdbGremlinDatabaseArgs) *CosmosdbGremlinDatabase {
	return &CosmosdbGremlinDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbGremlinDatabase)(nil)

// CosmosdbGremlinDatabase represents the Terraform resource azurerm_cosmosdb_gremlin_database.
type CosmosdbGremlinDatabase struct {
	Name      string
	Args      CosmosdbGremlinDatabaseArgs
	state     *cosmosdbGremlinDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbGremlinDatabase].
func (cgd *CosmosdbGremlinDatabase) Type() string {
	return "azurerm_cosmosdb_gremlin_database"
}

// LocalName returns the local name for [CosmosdbGremlinDatabase].
func (cgd *CosmosdbGremlinDatabase) LocalName() string {
	return cgd.Name
}

// Configuration returns the configuration (args) for [CosmosdbGremlinDatabase].
func (cgd *CosmosdbGremlinDatabase) Configuration() interface{} {
	return cgd.Args
}

// DependOn is used for other resources to depend on [CosmosdbGremlinDatabase].
func (cgd *CosmosdbGremlinDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(cgd)
}

// Dependencies returns the list of resources [CosmosdbGremlinDatabase] depends_on.
func (cgd *CosmosdbGremlinDatabase) Dependencies() terra.Dependencies {
	return cgd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbGremlinDatabase].
func (cgd *CosmosdbGremlinDatabase) LifecycleManagement() *terra.Lifecycle {
	return cgd.Lifecycle
}

// Attributes returns the attributes for [CosmosdbGremlinDatabase].
func (cgd *CosmosdbGremlinDatabase) Attributes() cosmosdbGremlinDatabaseAttributes {
	return cosmosdbGremlinDatabaseAttributes{ref: terra.ReferenceResource(cgd)}
}

// ImportState imports the given attribute values into [CosmosdbGremlinDatabase]'s state.
func (cgd *CosmosdbGremlinDatabase) ImportState(av io.Reader) error {
	cgd.state = &cosmosdbGremlinDatabaseState{}
	if err := json.NewDecoder(av).Decode(cgd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cgd.Type(), cgd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbGremlinDatabase] has state.
func (cgd *CosmosdbGremlinDatabase) State() (*cosmosdbGremlinDatabaseState, bool) {
	return cgd.state, cgd.state != nil
}

// StateMust returns the state for [CosmosdbGremlinDatabase]. Panics if the state is nil.
func (cgd *CosmosdbGremlinDatabase) StateMust() *cosmosdbGremlinDatabaseState {
	if cgd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cgd.Type(), cgd.LocalName()))
	}
	return cgd.state
}

// CosmosdbGremlinDatabaseArgs contains the configurations for azurerm_cosmosdb_gremlin_database.
type CosmosdbGremlinDatabaseArgs struct {
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
	AutoscaleSettings *cosmosdbgremlindatabase.AutoscaleSettings `hcl:"autoscale_settings,block"`
	// Timeouts: optional
	Timeouts *cosmosdbgremlindatabase.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbGremlinDatabaseAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_cosmosdb_gremlin_database.
func (cgd cosmosdbGremlinDatabaseAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(cgd.ref.Append("account_name"))
}

// Id returns a reference to field id of azurerm_cosmosdb_gremlin_database.
func (cgd cosmosdbGremlinDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cgd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_gremlin_database.
func (cgd cosmosdbGremlinDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cgd.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_gremlin_database.
func (cgd cosmosdbGremlinDatabaseAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cgd.ref.Append("resource_group_name"))
}

// Throughput returns a reference to field throughput of azurerm_cosmosdb_gremlin_database.
func (cgd cosmosdbGremlinDatabaseAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(cgd.ref.Append("throughput"))
}

func (cgd cosmosdbGremlinDatabaseAttributes) AutoscaleSettings() terra.ListValue[cosmosdbgremlindatabase.AutoscaleSettingsAttributes] {
	return terra.ReferenceAsList[cosmosdbgremlindatabase.AutoscaleSettingsAttributes](cgd.ref.Append("autoscale_settings"))
}

func (cgd cosmosdbGremlinDatabaseAttributes) Timeouts() cosmosdbgremlindatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbgremlindatabase.TimeoutsAttributes](cgd.ref.Append("timeouts"))
}

type cosmosdbGremlinDatabaseState struct {
	AccountName       string                                           `json:"account_name"`
	Id                string                                           `json:"id"`
	Name              string                                           `json:"name"`
	ResourceGroupName string                                           `json:"resource_group_name"`
	Throughput        float64                                          `json:"throughput"`
	AutoscaleSettings []cosmosdbgremlindatabase.AutoscaleSettingsState `json:"autoscale_settings"`
	Timeouts          *cosmosdbgremlindatabase.TimeoutsState           `json:"timeouts"`
}
