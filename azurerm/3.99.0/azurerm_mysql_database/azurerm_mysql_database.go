// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_mysql_database

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

// Resource represents the Terraform resource azurerm_mysql_database.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermMysqlDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (amd *Resource) Type() string {
	return "azurerm_mysql_database"
}

// LocalName returns the local name for [Resource].
func (amd *Resource) LocalName() string {
	return amd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (amd *Resource) Configuration() interface{} {
	return amd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (amd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(amd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (amd *Resource) Dependencies() terra.Dependencies {
	return amd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (amd *Resource) LifecycleManagement() *terra.Lifecycle {
	return amd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (amd *Resource) Attributes() azurermMysqlDatabaseAttributes {
	return azurermMysqlDatabaseAttributes{ref: terra.ReferenceResource(amd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (amd *Resource) ImportState(state io.Reader) error {
	amd.state = &azurermMysqlDatabaseState{}
	if err := json.NewDecoder(state).Decode(amd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amd.Type(), amd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (amd *Resource) State() (*azurermMysqlDatabaseState, bool) {
	return amd.state, amd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (amd *Resource) StateMust() *azurermMysqlDatabaseState {
	if amd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amd.Type(), amd.LocalName()))
	}
	return amd.state
}

// Args contains the configurations for azurerm_mysql_database.
type Args struct {
	// Charset: string, required
	Charset terra.StringValue `hcl:"charset,attr" validate:"required"`
	// Collation: string, required
	Collation terra.StringValue `hcl:"collation,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServerName: string, required
	ServerName terra.StringValue `hcl:"server_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermMysqlDatabaseAttributes struct {
	ref terra.Reference
}

// Charset returns a reference to field charset of azurerm_mysql_database.
func (amd azurermMysqlDatabaseAttributes) Charset() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("charset"))
}

// Collation returns a reference to field collation of azurerm_mysql_database.
func (amd azurermMysqlDatabaseAttributes) Collation() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("collation"))
}

// Id returns a reference to field id of azurerm_mysql_database.
func (amd azurermMysqlDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_mysql_database.
func (amd azurermMysqlDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mysql_database.
func (amd azurermMysqlDatabaseAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_mysql_database.
func (amd azurermMysqlDatabaseAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("server_name"))
}

func (amd azurermMysqlDatabaseAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](amd.ref.Append("timeouts"))
}

type azurermMysqlDatabaseState struct {
	Charset           string         `json:"charset"`
	Collation         string         `json:"collation"`
	Id                string         `json:"id"`
	Name              string         `json:"name"`
	ResourceGroupName string         `json:"resource_group_name"`
	ServerName        string         `json:"server_name"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
