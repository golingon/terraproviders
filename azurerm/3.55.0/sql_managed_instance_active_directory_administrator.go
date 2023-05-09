// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sqlmanagedinstanceactivedirectoryadministrator "github.com/golingon/terraproviders/azurerm/3.55.0/sqlmanagedinstanceactivedirectoryadministrator"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSqlManagedInstanceActiveDirectoryAdministrator creates a new instance of [SqlManagedInstanceActiveDirectoryAdministrator].
func NewSqlManagedInstanceActiveDirectoryAdministrator(name string, args SqlManagedInstanceActiveDirectoryAdministratorArgs) *SqlManagedInstanceActiveDirectoryAdministrator {
	return &SqlManagedInstanceActiveDirectoryAdministrator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqlManagedInstanceActiveDirectoryAdministrator)(nil)

// SqlManagedInstanceActiveDirectoryAdministrator represents the Terraform resource azurerm_sql_managed_instance_active_directory_administrator.
type SqlManagedInstanceActiveDirectoryAdministrator struct {
	Name      string
	Args      SqlManagedInstanceActiveDirectoryAdministratorArgs
	state     *sqlManagedInstanceActiveDirectoryAdministratorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqlManagedInstanceActiveDirectoryAdministrator].
func (smiada *SqlManagedInstanceActiveDirectoryAdministrator) Type() string {
	return "azurerm_sql_managed_instance_active_directory_administrator"
}

// LocalName returns the local name for [SqlManagedInstanceActiveDirectoryAdministrator].
func (smiada *SqlManagedInstanceActiveDirectoryAdministrator) LocalName() string {
	return smiada.Name
}

// Configuration returns the configuration (args) for [SqlManagedInstanceActiveDirectoryAdministrator].
func (smiada *SqlManagedInstanceActiveDirectoryAdministrator) Configuration() interface{} {
	return smiada.Args
}

// DependOn is used for other resources to depend on [SqlManagedInstanceActiveDirectoryAdministrator].
func (smiada *SqlManagedInstanceActiveDirectoryAdministrator) DependOn() terra.Reference {
	return terra.ReferenceResource(smiada)
}

// Dependencies returns the list of resources [SqlManagedInstanceActiveDirectoryAdministrator] depends_on.
func (smiada *SqlManagedInstanceActiveDirectoryAdministrator) Dependencies() terra.Dependencies {
	return smiada.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqlManagedInstanceActiveDirectoryAdministrator].
func (smiada *SqlManagedInstanceActiveDirectoryAdministrator) LifecycleManagement() *terra.Lifecycle {
	return smiada.Lifecycle
}

// Attributes returns the attributes for [SqlManagedInstanceActiveDirectoryAdministrator].
func (smiada *SqlManagedInstanceActiveDirectoryAdministrator) Attributes() sqlManagedInstanceActiveDirectoryAdministratorAttributes {
	return sqlManagedInstanceActiveDirectoryAdministratorAttributes{ref: terra.ReferenceResource(smiada)}
}

// ImportState imports the given attribute values into [SqlManagedInstanceActiveDirectoryAdministrator]'s state.
func (smiada *SqlManagedInstanceActiveDirectoryAdministrator) ImportState(av io.Reader) error {
	smiada.state = &sqlManagedInstanceActiveDirectoryAdministratorState{}
	if err := json.NewDecoder(av).Decode(smiada.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smiada.Type(), smiada.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqlManagedInstanceActiveDirectoryAdministrator] has state.
func (smiada *SqlManagedInstanceActiveDirectoryAdministrator) State() (*sqlManagedInstanceActiveDirectoryAdministratorState, bool) {
	return smiada.state, smiada.state != nil
}

// StateMust returns the state for [SqlManagedInstanceActiveDirectoryAdministrator]. Panics if the state is nil.
func (smiada *SqlManagedInstanceActiveDirectoryAdministrator) StateMust() *sqlManagedInstanceActiveDirectoryAdministratorState {
	if smiada.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smiada.Type(), smiada.LocalName()))
	}
	return smiada.state
}

// SqlManagedInstanceActiveDirectoryAdministratorArgs contains the configurations for azurerm_sql_managed_instance_active_directory_administrator.
type SqlManagedInstanceActiveDirectoryAdministratorArgs struct {
	// AzureadAuthenticationOnly: bool, optional
	AzureadAuthenticationOnly terra.BoolValue `hcl:"azuread_authentication_only,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Login: string, required
	Login terra.StringValue `hcl:"login,attr" validate:"required"`
	// ManagedInstanceName: string, required
	ManagedInstanceName terra.StringValue `hcl:"managed_instance_name,attr" validate:"required"`
	// ObjectId: string, required
	ObjectId terra.StringValue `hcl:"object_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// TenantId: string, required
	TenantId terra.StringValue `hcl:"tenant_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *sqlmanagedinstanceactivedirectoryadministrator.Timeouts `hcl:"timeouts,block"`
}
type sqlManagedInstanceActiveDirectoryAdministratorAttributes struct {
	ref terra.Reference
}

// AzureadAuthenticationOnly returns a reference to field azuread_authentication_only of azurerm_sql_managed_instance_active_directory_administrator.
func (smiada sqlManagedInstanceActiveDirectoryAdministratorAttributes) AzureadAuthenticationOnly() terra.BoolValue {
	return terra.ReferenceAsBool(smiada.ref.Append("azuread_authentication_only"))
}

// Id returns a reference to field id of azurerm_sql_managed_instance_active_directory_administrator.
func (smiada sqlManagedInstanceActiveDirectoryAdministratorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smiada.ref.Append("id"))
}

// Login returns a reference to field login of azurerm_sql_managed_instance_active_directory_administrator.
func (smiada sqlManagedInstanceActiveDirectoryAdministratorAttributes) Login() terra.StringValue {
	return terra.ReferenceAsString(smiada.ref.Append("login"))
}

// ManagedInstanceName returns a reference to field managed_instance_name of azurerm_sql_managed_instance_active_directory_administrator.
func (smiada sqlManagedInstanceActiveDirectoryAdministratorAttributes) ManagedInstanceName() terra.StringValue {
	return terra.ReferenceAsString(smiada.ref.Append("managed_instance_name"))
}

// ObjectId returns a reference to field object_id of azurerm_sql_managed_instance_active_directory_administrator.
func (smiada sqlManagedInstanceActiveDirectoryAdministratorAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(smiada.ref.Append("object_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_sql_managed_instance_active_directory_administrator.
func (smiada sqlManagedInstanceActiveDirectoryAdministratorAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(smiada.ref.Append("resource_group_name"))
}

// TenantId returns a reference to field tenant_id of azurerm_sql_managed_instance_active_directory_administrator.
func (smiada sqlManagedInstanceActiveDirectoryAdministratorAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(smiada.ref.Append("tenant_id"))
}

func (smiada sqlManagedInstanceActiveDirectoryAdministratorAttributes) Timeouts() sqlmanagedinstanceactivedirectoryadministrator.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sqlmanagedinstanceactivedirectoryadministrator.TimeoutsAttributes](smiada.ref.Append("timeouts"))
}

type sqlManagedInstanceActiveDirectoryAdministratorState struct {
	AzureadAuthenticationOnly bool                                                          `json:"azuread_authentication_only"`
	Id                        string                                                        `json:"id"`
	Login                     string                                                        `json:"login"`
	ManagedInstanceName       string                                                        `json:"managed_instance_name"`
	ObjectId                  string                                                        `json:"object_id"`
	ResourceGroupName         string                                                        `json:"resource_group_name"`
	TenantId                  string                                                        `json:"tenant_id"`
	Timeouts                  *sqlmanagedinstanceactivedirectoryadministrator.TimeoutsState `json:"timeouts"`
}
