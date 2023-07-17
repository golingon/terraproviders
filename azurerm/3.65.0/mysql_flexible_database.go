// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mysqlflexibledatabase "github.com/golingon/terraproviders/azurerm/3.65.0/mysqlflexibledatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMysqlFlexibleDatabase creates a new instance of [MysqlFlexibleDatabase].
func NewMysqlFlexibleDatabase(name string, args MysqlFlexibleDatabaseArgs) *MysqlFlexibleDatabase {
	return &MysqlFlexibleDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MysqlFlexibleDatabase)(nil)

// MysqlFlexibleDatabase represents the Terraform resource azurerm_mysql_flexible_database.
type MysqlFlexibleDatabase struct {
	Name      string
	Args      MysqlFlexibleDatabaseArgs
	state     *mysqlFlexibleDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MysqlFlexibleDatabase].
func (mfd *MysqlFlexibleDatabase) Type() string {
	return "azurerm_mysql_flexible_database"
}

// LocalName returns the local name for [MysqlFlexibleDatabase].
func (mfd *MysqlFlexibleDatabase) LocalName() string {
	return mfd.Name
}

// Configuration returns the configuration (args) for [MysqlFlexibleDatabase].
func (mfd *MysqlFlexibleDatabase) Configuration() interface{} {
	return mfd.Args
}

// DependOn is used for other resources to depend on [MysqlFlexibleDatabase].
func (mfd *MysqlFlexibleDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(mfd)
}

// Dependencies returns the list of resources [MysqlFlexibleDatabase] depends_on.
func (mfd *MysqlFlexibleDatabase) Dependencies() terra.Dependencies {
	return mfd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MysqlFlexibleDatabase].
func (mfd *MysqlFlexibleDatabase) LifecycleManagement() *terra.Lifecycle {
	return mfd.Lifecycle
}

// Attributes returns the attributes for [MysqlFlexibleDatabase].
func (mfd *MysqlFlexibleDatabase) Attributes() mysqlFlexibleDatabaseAttributes {
	return mysqlFlexibleDatabaseAttributes{ref: terra.ReferenceResource(mfd)}
}

// ImportState imports the given attribute values into [MysqlFlexibleDatabase]'s state.
func (mfd *MysqlFlexibleDatabase) ImportState(av io.Reader) error {
	mfd.state = &mysqlFlexibleDatabaseState{}
	if err := json.NewDecoder(av).Decode(mfd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mfd.Type(), mfd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MysqlFlexibleDatabase] has state.
func (mfd *MysqlFlexibleDatabase) State() (*mysqlFlexibleDatabaseState, bool) {
	return mfd.state, mfd.state != nil
}

// StateMust returns the state for [MysqlFlexibleDatabase]. Panics if the state is nil.
func (mfd *MysqlFlexibleDatabase) StateMust() *mysqlFlexibleDatabaseState {
	if mfd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mfd.Type(), mfd.LocalName()))
	}
	return mfd.state
}

// MysqlFlexibleDatabaseArgs contains the configurations for azurerm_mysql_flexible_database.
type MysqlFlexibleDatabaseArgs struct {
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
	Timeouts *mysqlflexibledatabase.Timeouts `hcl:"timeouts,block"`
}
type mysqlFlexibleDatabaseAttributes struct {
	ref terra.Reference
}

// Charset returns a reference to field charset of azurerm_mysql_flexible_database.
func (mfd mysqlFlexibleDatabaseAttributes) Charset() terra.StringValue {
	return terra.ReferenceAsString(mfd.ref.Append("charset"))
}

// Collation returns a reference to field collation of azurerm_mysql_flexible_database.
func (mfd mysqlFlexibleDatabaseAttributes) Collation() terra.StringValue {
	return terra.ReferenceAsString(mfd.ref.Append("collation"))
}

// Id returns a reference to field id of azurerm_mysql_flexible_database.
func (mfd mysqlFlexibleDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mfd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_mysql_flexible_database.
func (mfd mysqlFlexibleDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mfd.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mysql_flexible_database.
func (mfd mysqlFlexibleDatabaseAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mfd.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_mysql_flexible_database.
func (mfd mysqlFlexibleDatabaseAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(mfd.ref.Append("server_name"))
}

func (mfd mysqlFlexibleDatabaseAttributes) Timeouts() mysqlflexibledatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mysqlflexibledatabase.TimeoutsAttributes](mfd.ref.Append("timeouts"))
}

type mysqlFlexibleDatabaseState struct {
	Charset           string                               `json:"charset"`
	Collation         string                               `json:"collation"`
	Id                string                               `json:"id"`
	Name              string                               `json:"name"`
	ResourceGroupName string                               `json:"resource_group_name"`
	ServerName        string                               `json:"server_name"`
	Timeouts          *mysqlflexibledatabase.TimeoutsState `json:"timeouts"`
}
