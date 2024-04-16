// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_cosmosdb_sql_trigger

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_cosmosdb_sql_trigger.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermCosmosdbSqlTriggerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acst *Resource) Type() string {
	return "azurerm_cosmosdb_sql_trigger"
}

// LocalName returns the local name for [Resource].
func (acst *Resource) LocalName() string {
	return acst.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acst *Resource) Configuration() interface{} {
	return acst.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acst *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acst)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acst *Resource) Dependencies() terra.Dependencies {
	return acst.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acst *Resource) LifecycleManagement() *terra.Lifecycle {
	return acst.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acst *Resource) Attributes() azurermCosmosdbSqlTriggerAttributes {
	return azurermCosmosdbSqlTriggerAttributes{ref: terra.ReferenceResource(acst)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acst *Resource) ImportState(state io.Reader) error {
	acst.state = &azurermCosmosdbSqlTriggerState{}
	if err := json.NewDecoder(state).Decode(acst.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acst.Type(), acst.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acst *Resource) State() (*azurermCosmosdbSqlTriggerState, bool) {
	return acst.state, acst.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acst *Resource) StateMust() *azurermCosmosdbSqlTriggerState {
	if acst.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acst.Type(), acst.LocalName()))
	}
	return acst.state
}

// Args contains the configurations for azurerm_cosmosdb_sql_trigger.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermCosmosdbSqlTriggerAttributes struct {
	ref terra.Reference
}

// Body returns a reference to field body of azurerm_cosmosdb_sql_trigger.
func (acst azurermCosmosdbSqlTriggerAttributes) Body() terra.StringValue {
	return terra.ReferenceAsString(acst.ref.Append("body"))
}

// ContainerId returns a reference to field container_id of azurerm_cosmosdb_sql_trigger.
func (acst azurermCosmosdbSqlTriggerAttributes) ContainerId() terra.StringValue {
	return terra.ReferenceAsString(acst.ref.Append("container_id"))
}

// Id returns a reference to field id of azurerm_cosmosdb_sql_trigger.
func (acst azurermCosmosdbSqlTriggerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acst.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_sql_trigger.
func (acst azurermCosmosdbSqlTriggerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acst.ref.Append("name"))
}

// Operation returns a reference to field operation of azurerm_cosmosdb_sql_trigger.
func (acst azurermCosmosdbSqlTriggerAttributes) Operation() terra.StringValue {
	return terra.ReferenceAsString(acst.ref.Append("operation"))
}

// Type returns a reference to field type of azurerm_cosmosdb_sql_trigger.
func (acst azurermCosmosdbSqlTriggerAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(acst.ref.Append("type"))
}

func (acst azurermCosmosdbSqlTriggerAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](acst.ref.Append("timeouts"))
}

type azurermCosmosdbSqlTriggerState struct {
	Body        string         `json:"body"`
	ContainerId string         `json:"container_id"`
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	Operation   string         `json:"operation"`
	Type        string         `json:"type"`
	Timeouts    *TimeoutsState `json:"timeouts"`
}
