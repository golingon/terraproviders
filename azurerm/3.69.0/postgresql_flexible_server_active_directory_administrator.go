// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	postgresqlflexibleserveractivedirectoryadministrator "github.com/golingon/terraproviders/azurerm/3.69.0/postgresqlflexibleserveractivedirectoryadministrator"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPostgresqlFlexibleServerActiveDirectoryAdministrator creates a new instance of [PostgresqlFlexibleServerActiveDirectoryAdministrator].
func NewPostgresqlFlexibleServerActiveDirectoryAdministrator(name string, args PostgresqlFlexibleServerActiveDirectoryAdministratorArgs) *PostgresqlFlexibleServerActiveDirectoryAdministrator {
	return &PostgresqlFlexibleServerActiveDirectoryAdministrator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PostgresqlFlexibleServerActiveDirectoryAdministrator)(nil)

// PostgresqlFlexibleServerActiveDirectoryAdministrator represents the Terraform resource azurerm_postgresql_flexible_server_active_directory_administrator.
type PostgresqlFlexibleServerActiveDirectoryAdministrator struct {
	Name      string
	Args      PostgresqlFlexibleServerActiveDirectoryAdministratorArgs
	state     *postgresqlFlexibleServerActiveDirectoryAdministratorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PostgresqlFlexibleServerActiveDirectoryAdministrator].
func (pfsada *PostgresqlFlexibleServerActiveDirectoryAdministrator) Type() string {
	return "azurerm_postgresql_flexible_server_active_directory_administrator"
}

// LocalName returns the local name for [PostgresqlFlexibleServerActiveDirectoryAdministrator].
func (pfsada *PostgresqlFlexibleServerActiveDirectoryAdministrator) LocalName() string {
	return pfsada.Name
}

// Configuration returns the configuration (args) for [PostgresqlFlexibleServerActiveDirectoryAdministrator].
func (pfsada *PostgresqlFlexibleServerActiveDirectoryAdministrator) Configuration() interface{} {
	return pfsada.Args
}

// DependOn is used for other resources to depend on [PostgresqlFlexibleServerActiveDirectoryAdministrator].
func (pfsada *PostgresqlFlexibleServerActiveDirectoryAdministrator) DependOn() terra.Reference {
	return terra.ReferenceResource(pfsada)
}

// Dependencies returns the list of resources [PostgresqlFlexibleServerActiveDirectoryAdministrator] depends_on.
func (pfsada *PostgresqlFlexibleServerActiveDirectoryAdministrator) Dependencies() terra.Dependencies {
	return pfsada.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PostgresqlFlexibleServerActiveDirectoryAdministrator].
func (pfsada *PostgresqlFlexibleServerActiveDirectoryAdministrator) LifecycleManagement() *terra.Lifecycle {
	return pfsada.Lifecycle
}

// Attributes returns the attributes for [PostgresqlFlexibleServerActiveDirectoryAdministrator].
func (pfsada *PostgresqlFlexibleServerActiveDirectoryAdministrator) Attributes() postgresqlFlexibleServerActiveDirectoryAdministratorAttributes {
	return postgresqlFlexibleServerActiveDirectoryAdministratorAttributes{ref: terra.ReferenceResource(pfsada)}
}

// ImportState imports the given attribute values into [PostgresqlFlexibleServerActiveDirectoryAdministrator]'s state.
func (pfsada *PostgresqlFlexibleServerActiveDirectoryAdministrator) ImportState(av io.Reader) error {
	pfsada.state = &postgresqlFlexibleServerActiveDirectoryAdministratorState{}
	if err := json.NewDecoder(av).Decode(pfsada.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pfsada.Type(), pfsada.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PostgresqlFlexibleServerActiveDirectoryAdministrator] has state.
func (pfsada *PostgresqlFlexibleServerActiveDirectoryAdministrator) State() (*postgresqlFlexibleServerActiveDirectoryAdministratorState, bool) {
	return pfsada.state, pfsada.state != nil
}

// StateMust returns the state for [PostgresqlFlexibleServerActiveDirectoryAdministrator]. Panics if the state is nil.
func (pfsada *PostgresqlFlexibleServerActiveDirectoryAdministrator) StateMust() *postgresqlFlexibleServerActiveDirectoryAdministratorState {
	if pfsada.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pfsada.Type(), pfsada.LocalName()))
	}
	return pfsada.state
}

// PostgresqlFlexibleServerActiveDirectoryAdministratorArgs contains the configurations for azurerm_postgresql_flexible_server_active_directory_administrator.
type PostgresqlFlexibleServerActiveDirectoryAdministratorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ObjectId: string, required
	ObjectId terra.StringValue `hcl:"object_id,attr" validate:"required"`
	// PrincipalName: string, required
	PrincipalName terra.StringValue `hcl:"principal_name,attr" validate:"required"`
	// PrincipalType: string, required
	PrincipalType terra.StringValue `hcl:"principal_type,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServerName: string, required
	ServerName terra.StringValue `hcl:"server_name,attr" validate:"required"`
	// TenantId: string, required
	TenantId terra.StringValue `hcl:"tenant_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *postgresqlflexibleserveractivedirectoryadministrator.Timeouts `hcl:"timeouts,block"`
}
type postgresqlFlexibleServerActiveDirectoryAdministratorAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_postgresql_flexible_server_active_directory_administrator.
func (pfsada postgresqlFlexibleServerActiveDirectoryAdministratorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pfsada.ref.Append("id"))
}

// ObjectId returns a reference to field object_id of azurerm_postgresql_flexible_server_active_directory_administrator.
func (pfsada postgresqlFlexibleServerActiveDirectoryAdministratorAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(pfsada.ref.Append("object_id"))
}

// PrincipalName returns a reference to field principal_name of azurerm_postgresql_flexible_server_active_directory_administrator.
func (pfsada postgresqlFlexibleServerActiveDirectoryAdministratorAttributes) PrincipalName() terra.StringValue {
	return terra.ReferenceAsString(pfsada.ref.Append("principal_name"))
}

// PrincipalType returns a reference to field principal_type of azurerm_postgresql_flexible_server_active_directory_administrator.
func (pfsada postgresqlFlexibleServerActiveDirectoryAdministratorAttributes) PrincipalType() terra.StringValue {
	return terra.ReferenceAsString(pfsada.ref.Append("principal_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_postgresql_flexible_server_active_directory_administrator.
func (pfsada postgresqlFlexibleServerActiveDirectoryAdministratorAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pfsada.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_postgresql_flexible_server_active_directory_administrator.
func (pfsada postgresqlFlexibleServerActiveDirectoryAdministratorAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(pfsada.ref.Append("server_name"))
}

// TenantId returns a reference to field tenant_id of azurerm_postgresql_flexible_server_active_directory_administrator.
func (pfsada postgresqlFlexibleServerActiveDirectoryAdministratorAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(pfsada.ref.Append("tenant_id"))
}

func (pfsada postgresqlFlexibleServerActiveDirectoryAdministratorAttributes) Timeouts() postgresqlflexibleserveractivedirectoryadministrator.TimeoutsAttributes {
	return terra.ReferenceAsSingle[postgresqlflexibleserveractivedirectoryadministrator.TimeoutsAttributes](pfsada.ref.Append("timeouts"))
}

type postgresqlFlexibleServerActiveDirectoryAdministratorState struct {
	Id                string                                                              `json:"id"`
	ObjectId          string                                                              `json:"object_id"`
	PrincipalName     string                                                              `json:"principal_name"`
	PrincipalType     string                                                              `json:"principal_type"`
	ResourceGroupName string                                                              `json:"resource_group_name"`
	ServerName        string                                                              `json:"server_name"`
	TenantId          string                                                              `json:"tenant_id"`
	Timeouts          *postgresqlflexibleserveractivedirectoryadministrator.TimeoutsState `json:"timeouts"`
}
