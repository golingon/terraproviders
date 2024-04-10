// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	cosmosdbsqltrigger "github.com/golingon/terraproviders/azurerm/3.98.0/cosmosdbsqltrigger"
	"io"
)

// NewCosmosdbSqlTrigger creates a new instance of [CosmosdbSqlTrigger].
func NewCosmosdbSqlTrigger(name string, args CosmosdbSqlTriggerArgs) *CosmosdbSqlTrigger {
	return &CosmosdbSqlTrigger{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbSqlTrigger)(nil)

// CosmosdbSqlTrigger represents the Terraform resource azurerm_cosmosdb_sql_trigger.
type CosmosdbSqlTrigger struct {
	Name      string
	Args      CosmosdbSqlTriggerArgs
	state     *cosmosdbSqlTriggerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbSqlTrigger].
func (cst *CosmosdbSqlTrigger) Type() string {
	return "azurerm_cosmosdb_sql_trigger"
}

// LocalName returns the local name for [CosmosdbSqlTrigger].
func (cst *CosmosdbSqlTrigger) LocalName() string {
	return cst.Name
}

// Configuration returns the configuration (args) for [CosmosdbSqlTrigger].
func (cst *CosmosdbSqlTrigger) Configuration() interface{} {
	return cst.Args
}

// DependOn is used for other resources to depend on [CosmosdbSqlTrigger].
func (cst *CosmosdbSqlTrigger) DependOn() terra.Reference {
	return terra.ReferenceResource(cst)
}

// Dependencies returns the list of resources [CosmosdbSqlTrigger] depends_on.
func (cst *CosmosdbSqlTrigger) Dependencies() terra.Dependencies {
	return cst.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbSqlTrigger].
func (cst *CosmosdbSqlTrigger) LifecycleManagement() *terra.Lifecycle {
	return cst.Lifecycle
}

// Attributes returns the attributes for [CosmosdbSqlTrigger].
func (cst *CosmosdbSqlTrigger) Attributes() cosmosdbSqlTriggerAttributes {
	return cosmosdbSqlTriggerAttributes{ref: terra.ReferenceResource(cst)}
}

// ImportState imports the given attribute values into [CosmosdbSqlTrigger]'s state.
func (cst *CosmosdbSqlTrigger) ImportState(av io.Reader) error {
	cst.state = &cosmosdbSqlTriggerState{}
	if err := json.NewDecoder(av).Decode(cst.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cst.Type(), cst.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbSqlTrigger] has state.
func (cst *CosmosdbSqlTrigger) State() (*cosmosdbSqlTriggerState, bool) {
	return cst.state, cst.state != nil
}

// StateMust returns the state for [CosmosdbSqlTrigger]. Panics if the state is nil.
func (cst *CosmosdbSqlTrigger) StateMust() *cosmosdbSqlTriggerState {
	if cst.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cst.Type(), cst.LocalName()))
	}
	return cst.state
}

// CosmosdbSqlTriggerArgs contains the configurations for azurerm_cosmosdb_sql_trigger.
type CosmosdbSqlTriggerArgs struct {
	// Body: string, required
	Body terra.StringValue `hcl:"body,attr" validate:"required"`
	// ContainerId: string, required
	ContainerId terra.StringValue `hcl:"container_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Operation: string, required
	Operation terra.StringValue `hcl:"operation,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *cosmosdbsqltrigger.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbSqlTriggerAttributes struct {
	ref terra.Reference
}

// Body returns a reference to field body of azurerm_cosmosdb_sql_trigger.
func (cst cosmosdbSqlTriggerAttributes) Body() terra.StringValue {
	return terra.ReferenceAsString(cst.ref.Append("body"))
}

// ContainerId returns a reference to field container_id of azurerm_cosmosdb_sql_trigger.
func (cst cosmosdbSqlTriggerAttributes) ContainerId() terra.StringValue {
	return terra.ReferenceAsString(cst.ref.Append("container_id"))
}

// Id returns a reference to field id of azurerm_cosmosdb_sql_trigger.
func (cst cosmosdbSqlTriggerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cst.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_sql_trigger.
func (cst cosmosdbSqlTriggerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cst.ref.Append("name"))
}

// Operation returns a reference to field operation of azurerm_cosmosdb_sql_trigger.
func (cst cosmosdbSqlTriggerAttributes) Operation() terra.StringValue {
	return terra.ReferenceAsString(cst.ref.Append("operation"))
}

// Type returns a reference to field type of azurerm_cosmosdb_sql_trigger.
func (cst cosmosdbSqlTriggerAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(cst.ref.Append("type"))
}

func (cst cosmosdbSqlTriggerAttributes) Timeouts() cosmosdbsqltrigger.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbsqltrigger.TimeoutsAttributes](cst.ref.Append("timeouts"))
}

type cosmosdbSqlTriggerState struct {
	Body        string                            `json:"body"`
	ContainerId string                            `json:"container_id"`
	Id          string                            `json:"id"`
	Name        string                            `json:"name"`
	Operation   string                            `json:"operation"`
	Type        string                            `json:"type"`
	Timeouts    *cosmosdbsqltrigger.TimeoutsState `json:"timeouts"`
}
