// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sqlmanageddatabase "github.com/golingon/terraproviders/azurerm/3.58.0/sqlmanageddatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSqlManagedDatabase creates a new instance of [SqlManagedDatabase].
func NewSqlManagedDatabase(name string, args SqlManagedDatabaseArgs) *SqlManagedDatabase {
	return &SqlManagedDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqlManagedDatabase)(nil)

// SqlManagedDatabase represents the Terraform resource azurerm_sql_managed_database.
type SqlManagedDatabase struct {
	Name      string
	Args      SqlManagedDatabaseArgs
	state     *sqlManagedDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqlManagedDatabase].
func (smd *SqlManagedDatabase) Type() string {
	return "azurerm_sql_managed_database"
}

// LocalName returns the local name for [SqlManagedDatabase].
func (smd *SqlManagedDatabase) LocalName() string {
	return smd.Name
}

// Configuration returns the configuration (args) for [SqlManagedDatabase].
func (smd *SqlManagedDatabase) Configuration() interface{} {
	return smd.Args
}

// DependOn is used for other resources to depend on [SqlManagedDatabase].
func (smd *SqlManagedDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(smd)
}

// Dependencies returns the list of resources [SqlManagedDatabase] depends_on.
func (smd *SqlManagedDatabase) Dependencies() terra.Dependencies {
	return smd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqlManagedDatabase].
func (smd *SqlManagedDatabase) LifecycleManagement() *terra.Lifecycle {
	return smd.Lifecycle
}

// Attributes returns the attributes for [SqlManagedDatabase].
func (smd *SqlManagedDatabase) Attributes() sqlManagedDatabaseAttributes {
	return sqlManagedDatabaseAttributes{ref: terra.ReferenceResource(smd)}
}

// ImportState imports the given attribute values into [SqlManagedDatabase]'s state.
func (smd *SqlManagedDatabase) ImportState(av io.Reader) error {
	smd.state = &sqlManagedDatabaseState{}
	if err := json.NewDecoder(av).Decode(smd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smd.Type(), smd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqlManagedDatabase] has state.
func (smd *SqlManagedDatabase) State() (*sqlManagedDatabaseState, bool) {
	return smd.state, smd.state != nil
}

// StateMust returns the state for [SqlManagedDatabase]. Panics if the state is nil.
func (smd *SqlManagedDatabase) StateMust() *sqlManagedDatabaseState {
	if smd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smd.Type(), smd.LocalName()))
	}
	return smd.state
}

// SqlManagedDatabaseArgs contains the configurations for azurerm_sql_managed_database.
type SqlManagedDatabaseArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SqlManagedInstanceId: string, required
	SqlManagedInstanceId terra.StringValue `hcl:"sql_managed_instance_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *sqlmanageddatabase.Timeouts `hcl:"timeouts,block"`
}
type sqlManagedDatabaseAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sql_managed_database.
func (smd sqlManagedDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smd.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_sql_managed_database.
func (smd sqlManagedDatabaseAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(smd.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_sql_managed_database.
func (smd sqlManagedDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(smd.ref.Append("name"))
}

// SqlManagedInstanceId returns a reference to field sql_managed_instance_id of azurerm_sql_managed_database.
func (smd sqlManagedDatabaseAttributes) SqlManagedInstanceId() terra.StringValue {
	return terra.ReferenceAsString(smd.ref.Append("sql_managed_instance_id"))
}

func (smd sqlManagedDatabaseAttributes) Timeouts() sqlmanageddatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sqlmanageddatabase.TimeoutsAttributes](smd.ref.Append("timeouts"))
}

type sqlManagedDatabaseState struct {
	Id                   string                            `json:"id"`
	Location             string                            `json:"location"`
	Name                 string                            `json:"name"`
	SqlManagedInstanceId string                            `json:"sql_managed_instance_id"`
	Timeouts             *sqlmanageddatabase.TimeoutsState `json:"timeouts"`
}
