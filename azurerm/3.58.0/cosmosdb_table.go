// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbtable "github.com/golingon/terraproviders/azurerm/3.58.0/cosmosdbtable"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbTable creates a new instance of [CosmosdbTable].
func NewCosmosdbTable(name string, args CosmosdbTableArgs) *CosmosdbTable {
	return &CosmosdbTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbTable)(nil)

// CosmosdbTable represents the Terraform resource azurerm_cosmosdb_table.
type CosmosdbTable struct {
	Name      string
	Args      CosmosdbTableArgs
	state     *cosmosdbTableState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbTable].
func (ct *CosmosdbTable) Type() string {
	return "azurerm_cosmosdb_table"
}

// LocalName returns the local name for [CosmosdbTable].
func (ct *CosmosdbTable) LocalName() string {
	return ct.Name
}

// Configuration returns the configuration (args) for [CosmosdbTable].
func (ct *CosmosdbTable) Configuration() interface{} {
	return ct.Args
}

// DependOn is used for other resources to depend on [CosmosdbTable].
func (ct *CosmosdbTable) DependOn() terra.Reference {
	return terra.ReferenceResource(ct)
}

// Dependencies returns the list of resources [CosmosdbTable] depends_on.
func (ct *CosmosdbTable) Dependencies() terra.Dependencies {
	return ct.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbTable].
func (ct *CosmosdbTable) LifecycleManagement() *terra.Lifecycle {
	return ct.Lifecycle
}

// Attributes returns the attributes for [CosmosdbTable].
func (ct *CosmosdbTable) Attributes() cosmosdbTableAttributes {
	return cosmosdbTableAttributes{ref: terra.ReferenceResource(ct)}
}

// ImportState imports the given attribute values into [CosmosdbTable]'s state.
func (ct *CosmosdbTable) ImportState(av io.Reader) error {
	ct.state = &cosmosdbTableState{}
	if err := json.NewDecoder(av).Decode(ct.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ct.Type(), ct.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbTable] has state.
func (ct *CosmosdbTable) State() (*cosmosdbTableState, bool) {
	return ct.state, ct.state != nil
}

// StateMust returns the state for [CosmosdbTable]. Panics if the state is nil.
func (ct *CosmosdbTable) StateMust() *cosmosdbTableState {
	if ct.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ct.Type(), ct.LocalName()))
	}
	return ct.state
}

// CosmosdbTableArgs contains the configurations for azurerm_cosmosdb_table.
type CosmosdbTableArgs struct {
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
	AutoscaleSettings *cosmosdbtable.AutoscaleSettings `hcl:"autoscale_settings,block"`
	// Timeouts: optional
	Timeouts *cosmosdbtable.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbTableAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_cosmosdb_table.
func (ct cosmosdbTableAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("account_name"))
}

// Id returns a reference to field id of azurerm_cosmosdb_table.
func (ct cosmosdbTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_table.
func (ct cosmosdbTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_table.
func (ct cosmosdbTableAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ct.ref.Append("resource_group_name"))
}

// Throughput returns a reference to field throughput of azurerm_cosmosdb_table.
func (ct cosmosdbTableAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(ct.ref.Append("throughput"))
}

func (ct cosmosdbTableAttributes) AutoscaleSettings() terra.ListValue[cosmosdbtable.AutoscaleSettingsAttributes] {
	return terra.ReferenceAsList[cosmosdbtable.AutoscaleSettingsAttributes](ct.ref.Append("autoscale_settings"))
}

func (ct cosmosdbTableAttributes) Timeouts() cosmosdbtable.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbtable.TimeoutsAttributes](ct.ref.Append("timeouts"))
}

type cosmosdbTableState struct {
	AccountName       string                                 `json:"account_name"`
	Id                string                                 `json:"id"`
	Name              string                                 `json:"name"`
	ResourceGroupName string                                 `json:"resource_group_name"`
	Throughput        float64                                `json:"throughput"`
	AutoscaleSettings []cosmosdbtable.AutoscaleSettingsState `json:"autoscale_settings"`
	Timeouts          *cosmosdbtable.TimeoutsState           `json:"timeouts"`
}
