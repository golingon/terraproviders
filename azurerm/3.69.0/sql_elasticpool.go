// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sqlelasticpool "github.com/golingon/terraproviders/azurerm/3.69.0/sqlelasticpool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSqlElasticpool creates a new instance of [SqlElasticpool].
func NewSqlElasticpool(name string, args SqlElasticpoolArgs) *SqlElasticpool {
	return &SqlElasticpool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqlElasticpool)(nil)

// SqlElasticpool represents the Terraform resource azurerm_sql_elasticpool.
type SqlElasticpool struct {
	Name      string
	Args      SqlElasticpoolArgs
	state     *sqlElasticpoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqlElasticpool].
func (se *SqlElasticpool) Type() string {
	return "azurerm_sql_elasticpool"
}

// LocalName returns the local name for [SqlElasticpool].
func (se *SqlElasticpool) LocalName() string {
	return se.Name
}

// Configuration returns the configuration (args) for [SqlElasticpool].
func (se *SqlElasticpool) Configuration() interface{} {
	return se.Args
}

// DependOn is used for other resources to depend on [SqlElasticpool].
func (se *SqlElasticpool) DependOn() terra.Reference {
	return terra.ReferenceResource(se)
}

// Dependencies returns the list of resources [SqlElasticpool] depends_on.
func (se *SqlElasticpool) Dependencies() terra.Dependencies {
	return se.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqlElasticpool].
func (se *SqlElasticpool) LifecycleManagement() *terra.Lifecycle {
	return se.Lifecycle
}

// Attributes returns the attributes for [SqlElasticpool].
func (se *SqlElasticpool) Attributes() sqlElasticpoolAttributes {
	return sqlElasticpoolAttributes{ref: terra.ReferenceResource(se)}
}

// ImportState imports the given attribute values into [SqlElasticpool]'s state.
func (se *SqlElasticpool) ImportState(av io.Reader) error {
	se.state = &sqlElasticpoolState{}
	if err := json.NewDecoder(av).Decode(se.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", se.Type(), se.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqlElasticpool] has state.
func (se *SqlElasticpool) State() (*sqlElasticpoolState, bool) {
	return se.state, se.state != nil
}

// StateMust returns the state for [SqlElasticpool]. Panics if the state is nil.
func (se *SqlElasticpool) StateMust() *sqlElasticpoolState {
	if se.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", se.Type(), se.LocalName()))
	}
	return se.state
}

// SqlElasticpoolArgs contains the configurations for azurerm_sql_elasticpool.
type SqlElasticpoolArgs struct {
	// DbDtuMax: number, optional
	DbDtuMax terra.NumberValue `hcl:"db_dtu_max,attr"`
	// DbDtuMin: number, optional
	DbDtuMin terra.NumberValue `hcl:"db_dtu_min,attr"`
	// Dtu: number, required
	Dtu terra.NumberValue `hcl:"dtu,attr" validate:"required"`
	// Edition: string, required
	Edition terra.StringValue `hcl:"edition,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PoolSize: number, optional
	PoolSize terra.NumberValue `hcl:"pool_size,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServerName: string, required
	ServerName terra.StringValue `hcl:"server_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *sqlelasticpool.Timeouts `hcl:"timeouts,block"`
}
type sqlElasticpoolAttributes struct {
	ref terra.Reference
}

// CreationDate returns a reference to field creation_date of azurerm_sql_elasticpool.
func (se sqlElasticpoolAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("creation_date"))
}

// DbDtuMax returns a reference to field db_dtu_max of azurerm_sql_elasticpool.
func (se sqlElasticpoolAttributes) DbDtuMax() terra.NumberValue {
	return terra.ReferenceAsNumber(se.ref.Append("db_dtu_max"))
}

// DbDtuMin returns a reference to field db_dtu_min of azurerm_sql_elasticpool.
func (se sqlElasticpoolAttributes) DbDtuMin() terra.NumberValue {
	return terra.ReferenceAsNumber(se.ref.Append("db_dtu_min"))
}

// Dtu returns a reference to field dtu of azurerm_sql_elasticpool.
func (se sqlElasticpoolAttributes) Dtu() terra.NumberValue {
	return terra.ReferenceAsNumber(se.ref.Append("dtu"))
}

// Edition returns a reference to field edition of azurerm_sql_elasticpool.
func (se sqlElasticpoolAttributes) Edition() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("edition"))
}

// Id returns a reference to field id of azurerm_sql_elasticpool.
func (se sqlElasticpoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_sql_elasticpool.
func (se sqlElasticpoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_sql_elasticpool.
func (se sqlElasticpoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("name"))
}

// PoolSize returns a reference to field pool_size of azurerm_sql_elasticpool.
func (se sqlElasticpoolAttributes) PoolSize() terra.NumberValue {
	return terra.ReferenceAsNumber(se.ref.Append("pool_size"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_sql_elasticpool.
func (se sqlElasticpoolAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_sql_elasticpool.
func (se sqlElasticpoolAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("server_name"))
}

// Tags returns a reference to field tags of azurerm_sql_elasticpool.
func (se sqlElasticpoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](se.ref.Append("tags"))
}

func (se sqlElasticpoolAttributes) Timeouts() sqlelasticpool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sqlelasticpool.TimeoutsAttributes](se.ref.Append("timeouts"))
}

type sqlElasticpoolState struct {
	CreationDate      string                        `json:"creation_date"`
	DbDtuMax          float64                       `json:"db_dtu_max"`
	DbDtuMin          float64                       `json:"db_dtu_min"`
	Dtu               float64                       `json:"dtu"`
	Edition           string                        `json:"edition"`
	Id                string                        `json:"id"`
	Location          string                        `json:"location"`
	Name              string                        `json:"name"`
	PoolSize          float64                       `json:"pool_size"`
	ResourceGroupName string                        `json:"resource_group_name"`
	ServerName        string                        `json:"server_name"`
	Tags              map[string]string             `json:"tags"`
	Timeouts          *sqlelasticpool.TimeoutsState `json:"timeouts"`
}
