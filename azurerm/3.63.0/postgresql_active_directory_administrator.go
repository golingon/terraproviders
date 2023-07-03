// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	postgresqlactivedirectoryadministrator "github.com/golingon/terraproviders/azurerm/3.63.0/postgresqlactivedirectoryadministrator"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPostgresqlActiveDirectoryAdministrator creates a new instance of [PostgresqlActiveDirectoryAdministrator].
func NewPostgresqlActiveDirectoryAdministrator(name string, args PostgresqlActiveDirectoryAdministratorArgs) *PostgresqlActiveDirectoryAdministrator {
	return &PostgresqlActiveDirectoryAdministrator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PostgresqlActiveDirectoryAdministrator)(nil)

// PostgresqlActiveDirectoryAdministrator represents the Terraform resource azurerm_postgresql_active_directory_administrator.
type PostgresqlActiveDirectoryAdministrator struct {
	Name      string
	Args      PostgresqlActiveDirectoryAdministratorArgs
	state     *postgresqlActiveDirectoryAdministratorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PostgresqlActiveDirectoryAdministrator].
func (pada *PostgresqlActiveDirectoryAdministrator) Type() string {
	return "azurerm_postgresql_active_directory_administrator"
}

// LocalName returns the local name for [PostgresqlActiveDirectoryAdministrator].
func (pada *PostgresqlActiveDirectoryAdministrator) LocalName() string {
	return pada.Name
}

// Configuration returns the configuration (args) for [PostgresqlActiveDirectoryAdministrator].
func (pada *PostgresqlActiveDirectoryAdministrator) Configuration() interface{} {
	return pada.Args
}

// DependOn is used for other resources to depend on [PostgresqlActiveDirectoryAdministrator].
func (pada *PostgresqlActiveDirectoryAdministrator) DependOn() terra.Reference {
	return terra.ReferenceResource(pada)
}

// Dependencies returns the list of resources [PostgresqlActiveDirectoryAdministrator] depends_on.
func (pada *PostgresqlActiveDirectoryAdministrator) Dependencies() terra.Dependencies {
	return pada.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PostgresqlActiveDirectoryAdministrator].
func (pada *PostgresqlActiveDirectoryAdministrator) LifecycleManagement() *terra.Lifecycle {
	return pada.Lifecycle
}

// Attributes returns the attributes for [PostgresqlActiveDirectoryAdministrator].
func (pada *PostgresqlActiveDirectoryAdministrator) Attributes() postgresqlActiveDirectoryAdministratorAttributes {
	return postgresqlActiveDirectoryAdministratorAttributes{ref: terra.ReferenceResource(pada)}
}

// ImportState imports the given attribute values into [PostgresqlActiveDirectoryAdministrator]'s state.
func (pada *PostgresqlActiveDirectoryAdministrator) ImportState(av io.Reader) error {
	pada.state = &postgresqlActiveDirectoryAdministratorState{}
	if err := json.NewDecoder(av).Decode(pada.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pada.Type(), pada.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PostgresqlActiveDirectoryAdministrator] has state.
func (pada *PostgresqlActiveDirectoryAdministrator) State() (*postgresqlActiveDirectoryAdministratorState, bool) {
	return pada.state, pada.state != nil
}

// StateMust returns the state for [PostgresqlActiveDirectoryAdministrator]. Panics if the state is nil.
func (pada *PostgresqlActiveDirectoryAdministrator) StateMust() *postgresqlActiveDirectoryAdministratorState {
	if pada.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pada.Type(), pada.LocalName()))
	}
	return pada.state
}

// PostgresqlActiveDirectoryAdministratorArgs contains the configurations for azurerm_postgresql_active_directory_administrator.
type PostgresqlActiveDirectoryAdministratorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Login: string, required
	Login terra.StringValue `hcl:"login,attr" validate:"required"`
	// ObjectId: string, required
	ObjectId terra.StringValue `hcl:"object_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServerName: string, required
	ServerName terra.StringValue `hcl:"server_name,attr" validate:"required"`
	// TenantId: string, required
	TenantId terra.StringValue `hcl:"tenant_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *postgresqlactivedirectoryadministrator.Timeouts `hcl:"timeouts,block"`
}
type postgresqlActiveDirectoryAdministratorAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_postgresql_active_directory_administrator.
func (pada postgresqlActiveDirectoryAdministratorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pada.ref.Append("id"))
}

// Login returns a reference to field login of azurerm_postgresql_active_directory_administrator.
func (pada postgresqlActiveDirectoryAdministratorAttributes) Login() terra.StringValue {
	return terra.ReferenceAsString(pada.ref.Append("login"))
}

// ObjectId returns a reference to field object_id of azurerm_postgresql_active_directory_administrator.
func (pada postgresqlActiveDirectoryAdministratorAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(pada.ref.Append("object_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_postgresql_active_directory_administrator.
func (pada postgresqlActiveDirectoryAdministratorAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pada.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_postgresql_active_directory_administrator.
func (pada postgresqlActiveDirectoryAdministratorAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(pada.ref.Append("server_name"))
}

// TenantId returns a reference to field tenant_id of azurerm_postgresql_active_directory_administrator.
func (pada postgresqlActiveDirectoryAdministratorAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(pada.ref.Append("tenant_id"))
}

func (pada postgresqlActiveDirectoryAdministratorAttributes) Timeouts() postgresqlactivedirectoryadministrator.TimeoutsAttributes {
	return terra.ReferenceAsSingle[postgresqlactivedirectoryadministrator.TimeoutsAttributes](pada.ref.Append("timeouts"))
}

type postgresqlActiveDirectoryAdministratorState struct {
	Id                string                                                `json:"id"`
	Login             string                                                `json:"login"`
	ObjectId          string                                                `json:"object_id"`
	ResourceGroupName string                                                `json:"resource_group_name"`
	ServerName        string                                                `json:"server_name"`
	TenantId          string                                                `json:"tenant_id"`
	Timeouts          *postgresqlactivedirectoryadministrator.TimeoutsState `json:"timeouts"`
}
