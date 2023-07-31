// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mysqldatabase "github.com/golingon/terraproviders/azurerm/3.67.0/mysqldatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMysqlDatabase creates a new instance of [MysqlDatabase].
func NewMysqlDatabase(name string, args MysqlDatabaseArgs) *MysqlDatabase {
	return &MysqlDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MysqlDatabase)(nil)

// MysqlDatabase represents the Terraform resource azurerm_mysql_database.
type MysqlDatabase struct {
	Name      string
	Args      MysqlDatabaseArgs
	state     *mysqlDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MysqlDatabase].
func (md *MysqlDatabase) Type() string {
	return "azurerm_mysql_database"
}

// LocalName returns the local name for [MysqlDatabase].
func (md *MysqlDatabase) LocalName() string {
	return md.Name
}

// Configuration returns the configuration (args) for [MysqlDatabase].
func (md *MysqlDatabase) Configuration() interface{} {
	return md.Args
}

// DependOn is used for other resources to depend on [MysqlDatabase].
func (md *MysqlDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(md)
}

// Dependencies returns the list of resources [MysqlDatabase] depends_on.
func (md *MysqlDatabase) Dependencies() terra.Dependencies {
	return md.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MysqlDatabase].
func (md *MysqlDatabase) LifecycleManagement() *terra.Lifecycle {
	return md.Lifecycle
}

// Attributes returns the attributes for [MysqlDatabase].
func (md *MysqlDatabase) Attributes() mysqlDatabaseAttributes {
	return mysqlDatabaseAttributes{ref: terra.ReferenceResource(md)}
}

// ImportState imports the given attribute values into [MysqlDatabase]'s state.
func (md *MysqlDatabase) ImportState(av io.Reader) error {
	md.state = &mysqlDatabaseState{}
	if err := json.NewDecoder(av).Decode(md.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", md.Type(), md.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MysqlDatabase] has state.
func (md *MysqlDatabase) State() (*mysqlDatabaseState, bool) {
	return md.state, md.state != nil
}

// StateMust returns the state for [MysqlDatabase]. Panics if the state is nil.
func (md *MysqlDatabase) StateMust() *mysqlDatabaseState {
	if md.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", md.Type(), md.LocalName()))
	}
	return md.state
}

// MysqlDatabaseArgs contains the configurations for azurerm_mysql_database.
type MysqlDatabaseArgs struct {
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
	Timeouts *mysqldatabase.Timeouts `hcl:"timeouts,block"`
}
type mysqlDatabaseAttributes struct {
	ref terra.Reference
}

// Charset returns a reference to field charset of azurerm_mysql_database.
func (md mysqlDatabaseAttributes) Charset() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("charset"))
}

// Collation returns a reference to field collation of azurerm_mysql_database.
func (md mysqlDatabaseAttributes) Collation() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("collation"))
}

// Id returns a reference to field id of azurerm_mysql_database.
func (md mysqlDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_mysql_database.
func (md mysqlDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mysql_database.
func (md mysqlDatabaseAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_mysql_database.
func (md mysqlDatabaseAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("server_name"))
}

func (md mysqlDatabaseAttributes) Timeouts() mysqldatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mysqldatabase.TimeoutsAttributes](md.ref.Append("timeouts"))
}

type mysqlDatabaseState struct {
	Charset           string                       `json:"charset"`
	Collation         string                       `json:"collation"`
	Id                string                       `json:"id"`
	Name              string                       `json:"name"`
	ResourceGroupName string                       `json:"resource_group_name"`
	ServerName        string                       `json:"server_name"`
	Timeouts          *mysqldatabase.TimeoutsState `json:"timeouts"`
}
