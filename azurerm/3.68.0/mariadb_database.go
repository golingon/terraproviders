// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mariadbdatabase "github.com/golingon/terraproviders/azurerm/3.68.0/mariadbdatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMariadbDatabase creates a new instance of [MariadbDatabase].
func NewMariadbDatabase(name string, args MariadbDatabaseArgs) *MariadbDatabase {
	return &MariadbDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MariadbDatabase)(nil)

// MariadbDatabase represents the Terraform resource azurerm_mariadb_database.
type MariadbDatabase struct {
	Name      string
	Args      MariadbDatabaseArgs
	state     *mariadbDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MariadbDatabase].
func (md *MariadbDatabase) Type() string {
	return "azurerm_mariadb_database"
}

// LocalName returns the local name for [MariadbDatabase].
func (md *MariadbDatabase) LocalName() string {
	return md.Name
}

// Configuration returns the configuration (args) for [MariadbDatabase].
func (md *MariadbDatabase) Configuration() interface{} {
	return md.Args
}

// DependOn is used for other resources to depend on [MariadbDatabase].
func (md *MariadbDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(md)
}

// Dependencies returns the list of resources [MariadbDatabase] depends_on.
func (md *MariadbDatabase) Dependencies() terra.Dependencies {
	return md.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MariadbDatabase].
func (md *MariadbDatabase) LifecycleManagement() *terra.Lifecycle {
	return md.Lifecycle
}

// Attributes returns the attributes for [MariadbDatabase].
func (md *MariadbDatabase) Attributes() mariadbDatabaseAttributes {
	return mariadbDatabaseAttributes{ref: terra.ReferenceResource(md)}
}

// ImportState imports the given attribute values into [MariadbDatabase]'s state.
func (md *MariadbDatabase) ImportState(av io.Reader) error {
	md.state = &mariadbDatabaseState{}
	if err := json.NewDecoder(av).Decode(md.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", md.Type(), md.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MariadbDatabase] has state.
func (md *MariadbDatabase) State() (*mariadbDatabaseState, bool) {
	return md.state, md.state != nil
}

// StateMust returns the state for [MariadbDatabase]. Panics if the state is nil.
func (md *MariadbDatabase) StateMust() *mariadbDatabaseState {
	if md.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", md.Type(), md.LocalName()))
	}
	return md.state
}

// MariadbDatabaseArgs contains the configurations for azurerm_mariadb_database.
type MariadbDatabaseArgs struct {
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
	Timeouts *mariadbdatabase.Timeouts `hcl:"timeouts,block"`
}
type mariadbDatabaseAttributes struct {
	ref terra.Reference
}

// Charset returns a reference to field charset of azurerm_mariadb_database.
func (md mariadbDatabaseAttributes) Charset() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("charset"))
}

// Collation returns a reference to field collation of azurerm_mariadb_database.
func (md mariadbDatabaseAttributes) Collation() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("collation"))
}

// Id returns a reference to field id of azurerm_mariadb_database.
func (md mariadbDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_mariadb_database.
func (md mariadbDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mariadb_database.
func (md mariadbDatabaseAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_mariadb_database.
func (md mariadbDatabaseAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("server_name"))
}

func (md mariadbDatabaseAttributes) Timeouts() mariadbdatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mariadbdatabase.TimeoutsAttributes](md.ref.Append("timeouts"))
}

type mariadbDatabaseState struct {
	Charset           string                         `json:"charset"`
	Collation         string                         `json:"collation"`
	Id                string                         `json:"id"`
	Name              string                         `json:"name"`
	ResourceGroupName string                         `json:"resource_group_name"`
	ServerName        string                         `json:"server_name"`
	Timeouts          *mariadbdatabase.TimeoutsState `json:"timeouts"`
}
