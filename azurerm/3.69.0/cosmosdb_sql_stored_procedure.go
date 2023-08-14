// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbsqlstoredprocedure "github.com/golingon/terraproviders/azurerm/3.69.0/cosmosdbsqlstoredprocedure"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbSqlStoredProcedure creates a new instance of [CosmosdbSqlStoredProcedure].
func NewCosmosdbSqlStoredProcedure(name string, args CosmosdbSqlStoredProcedureArgs) *CosmosdbSqlStoredProcedure {
	return &CosmosdbSqlStoredProcedure{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbSqlStoredProcedure)(nil)

// CosmosdbSqlStoredProcedure represents the Terraform resource azurerm_cosmosdb_sql_stored_procedure.
type CosmosdbSqlStoredProcedure struct {
	Name      string
	Args      CosmosdbSqlStoredProcedureArgs
	state     *cosmosdbSqlStoredProcedureState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbSqlStoredProcedure].
func (cssp *CosmosdbSqlStoredProcedure) Type() string {
	return "azurerm_cosmosdb_sql_stored_procedure"
}

// LocalName returns the local name for [CosmosdbSqlStoredProcedure].
func (cssp *CosmosdbSqlStoredProcedure) LocalName() string {
	return cssp.Name
}

// Configuration returns the configuration (args) for [CosmosdbSqlStoredProcedure].
func (cssp *CosmosdbSqlStoredProcedure) Configuration() interface{} {
	return cssp.Args
}

// DependOn is used for other resources to depend on [CosmosdbSqlStoredProcedure].
func (cssp *CosmosdbSqlStoredProcedure) DependOn() terra.Reference {
	return terra.ReferenceResource(cssp)
}

// Dependencies returns the list of resources [CosmosdbSqlStoredProcedure] depends_on.
func (cssp *CosmosdbSqlStoredProcedure) Dependencies() terra.Dependencies {
	return cssp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbSqlStoredProcedure].
func (cssp *CosmosdbSqlStoredProcedure) LifecycleManagement() *terra.Lifecycle {
	return cssp.Lifecycle
}

// Attributes returns the attributes for [CosmosdbSqlStoredProcedure].
func (cssp *CosmosdbSqlStoredProcedure) Attributes() cosmosdbSqlStoredProcedureAttributes {
	return cosmosdbSqlStoredProcedureAttributes{ref: terra.ReferenceResource(cssp)}
}

// ImportState imports the given attribute values into [CosmosdbSqlStoredProcedure]'s state.
func (cssp *CosmosdbSqlStoredProcedure) ImportState(av io.Reader) error {
	cssp.state = &cosmosdbSqlStoredProcedureState{}
	if err := json.NewDecoder(av).Decode(cssp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cssp.Type(), cssp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbSqlStoredProcedure] has state.
func (cssp *CosmosdbSqlStoredProcedure) State() (*cosmosdbSqlStoredProcedureState, bool) {
	return cssp.state, cssp.state != nil
}

// StateMust returns the state for [CosmosdbSqlStoredProcedure]. Panics if the state is nil.
func (cssp *CosmosdbSqlStoredProcedure) StateMust() *cosmosdbSqlStoredProcedureState {
	if cssp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cssp.Type(), cssp.LocalName()))
	}
	return cssp.state
}

// CosmosdbSqlStoredProcedureArgs contains the configurations for azurerm_cosmosdb_sql_stored_procedure.
type CosmosdbSqlStoredProcedureArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// Body: string, required
	Body terra.StringValue `hcl:"body,attr" validate:"required"`
	// ContainerName: string, required
	ContainerName terra.StringValue `hcl:"container_name,attr" validate:"required"`
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *cosmosdbsqlstoredprocedure.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbSqlStoredProcedureAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_cosmosdb_sql_stored_procedure.
func (cssp cosmosdbSqlStoredProcedureAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(cssp.ref.Append("account_name"))
}

// Body returns a reference to field body of azurerm_cosmosdb_sql_stored_procedure.
func (cssp cosmosdbSqlStoredProcedureAttributes) Body() terra.StringValue {
	return terra.ReferenceAsString(cssp.ref.Append("body"))
}

// ContainerName returns a reference to field container_name of azurerm_cosmosdb_sql_stored_procedure.
func (cssp cosmosdbSqlStoredProcedureAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(cssp.ref.Append("container_name"))
}

// DatabaseName returns a reference to field database_name of azurerm_cosmosdb_sql_stored_procedure.
func (cssp cosmosdbSqlStoredProcedureAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(cssp.ref.Append("database_name"))
}

// Id returns a reference to field id of azurerm_cosmosdb_sql_stored_procedure.
func (cssp cosmosdbSqlStoredProcedureAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cssp.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_sql_stored_procedure.
func (cssp cosmosdbSqlStoredProcedureAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cssp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_sql_stored_procedure.
func (cssp cosmosdbSqlStoredProcedureAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cssp.ref.Append("resource_group_name"))
}

func (cssp cosmosdbSqlStoredProcedureAttributes) Timeouts() cosmosdbsqlstoredprocedure.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbsqlstoredprocedure.TimeoutsAttributes](cssp.ref.Append("timeouts"))
}

type cosmosdbSqlStoredProcedureState struct {
	AccountName       string                                    `json:"account_name"`
	Body              string                                    `json:"body"`
	ContainerName     string                                    `json:"container_name"`
	DatabaseName      string                                    `json:"database_name"`
	Id                string                                    `json:"id"`
	Name              string                                    `json:"name"`
	ResourceGroupName string                                    `json:"resource_group_name"`
	Timeouts          *cosmosdbsqlstoredprocedure.TimeoutsState `json:"timeouts"`
}