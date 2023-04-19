// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sqlactivedirectoryadministrator "github.com/golingon/terraproviders/azurerm/3.52.0/sqlactivedirectoryadministrator"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSqlActiveDirectoryAdministrator creates a new instance of [SqlActiveDirectoryAdministrator].
func NewSqlActiveDirectoryAdministrator(name string, args SqlActiveDirectoryAdministratorArgs) *SqlActiveDirectoryAdministrator {
	return &SqlActiveDirectoryAdministrator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqlActiveDirectoryAdministrator)(nil)

// SqlActiveDirectoryAdministrator represents the Terraform resource azurerm_sql_active_directory_administrator.
type SqlActiveDirectoryAdministrator struct {
	Name      string
	Args      SqlActiveDirectoryAdministratorArgs
	state     *sqlActiveDirectoryAdministratorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqlActiveDirectoryAdministrator].
func (sada *SqlActiveDirectoryAdministrator) Type() string {
	return "azurerm_sql_active_directory_administrator"
}

// LocalName returns the local name for [SqlActiveDirectoryAdministrator].
func (sada *SqlActiveDirectoryAdministrator) LocalName() string {
	return sada.Name
}

// Configuration returns the configuration (args) for [SqlActiveDirectoryAdministrator].
func (sada *SqlActiveDirectoryAdministrator) Configuration() interface{} {
	return sada.Args
}

// DependOn is used for other resources to depend on [SqlActiveDirectoryAdministrator].
func (sada *SqlActiveDirectoryAdministrator) DependOn() terra.Reference {
	return terra.ReferenceResource(sada)
}

// Dependencies returns the list of resources [SqlActiveDirectoryAdministrator] depends_on.
func (sada *SqlActiveDirectoryAdministrator) Dependencies() terra.Dependencies {
	return sada.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqlActiveDirectoryAdministrator].
func (sada *SqlActiveDirectoryAdministrator) LifecycleManagement() *terra.Lifecycle {
	return sada.Lifecycle
}

// Attributes returns the attributes for [SqlActiveDirectoryAdministrator].
func (sada *SqlActiveDirectoryAdministrator) Attributes() sqlActiveDirectoryAdministratorAttributes {
	return sqlActiveDirectoryAdministratorAttributes{ref: terra.ReferenceResource(sada)}
}

// ImportState imports the given attribute values into [SqlActiveDirectoryAdministrator]'s state.
func (sada *SqlActiveDirectoryAdministrator) ImportState(av io.Reader) error {
	sada.state = &sqlActiveDirectoryAdministratorState{}
	if err := json.NewDecoder(av).Decode(sada.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sada.Type(), sada.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqlActiveDirectoryAdministrator] has state.
func (sada *SqlActiveDirectoryAdministrator) State() (*sqlActiveDirectoryAdministratorState, bool) {
	return sada.state, sada.state != nil
}

// StateMust returns the state for [SqlActiveDirectoryAdministrator]. Panics if the state is nil.
func (sada *SqlActiveDirectoryAdministrator) StateMust() *sqlActiveDirectoryAdministratorState {
	if sada.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sada.Type(), sada.LocalName()))
	}
	return sada.state
}

// SqlActiveDirectoryAdministratorArgs contains the configurations for azurerm_sql_active_directory_administrator.
type SqlActiveDirectoryAdministratorArgs struct {
	// AzureadAuthenticationOnly: bool, optional
	AzureadAuthenticationOnly terra.BoolValue `hcl:"azuread_authentication_only,attr"`
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
	Timeouts *sqlactivedirectoryadministrator.Timeouts `hcl:"timeouts,block"`
}
type sqlActiveDirectoryAdministratorAttributes struct {
	ref terra.Reference
}

// AzureadAuthenticationOnly returns a reference to field azuread_authentication_only of azurerm_sql_active_directory_administrator.
func (sada sqlActiveDirectoryAdministratorAttributes) AzureadAuthenticationOnly() terra.BoolValue {
	return terra.ReferenceAsBool(sada.ref.Append("azuread_authentication_only"))
}

// Id returns a reference to field id of azurerm_sql_active_directory_administrator.
func (sada sqlActiveDirectoryAdministratorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sada.ref.Append("id"))
}

// Login returns a reference to field login of azurerm_sql_active_directory_administrator.
func (sada sqlActiveDirectoryAdministratorAttributes) Login() terra.StringValue {
	return terra.ReferenceAsString(sada.ref.Append("login"))
}

// ObjectId returns a reference to field object_id of azurerm_sql_active_directory_administrator.
func (sada sqlActiveDirectoryAdministratorAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(sada.ref.Append("object_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_sql_active_directory_administrator.
func (sada sqlActiveDirectoryAdministratorAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sada.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_sql_active_directory_administrator.
func (sada sqlActiveDirectoryAdministratorAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(sada.ref.Append("server_name"))
}

// TenantId returns a reference to field tenant_id of azurerm_sql_active_directory_administrator.
func (sada sqlActiveDirectoryAdministratorAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(sada.ref.Append("tenant_id"))
}

func (sada sqlActiveDirectoryAdministratorAttributes) Timeouts() sqlactivedirectoryadministrator.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sqlactivedirectoryadministrator.TimeoutsAttributes](sada.ref.Append("timeouts"))
}

type sqlActiveDirectoryAdministratorState struct {
	AzureadAuthenticationOnly bool                                           `json:"azuread_authentication_only"`
	Id                        string                                         `json:"id"`
	Login                     string                                         `json:"login"`
	ObjectId                  string                                         `json:"object_id"`
	ResourceGroupName         string                                         `json:"resource_group_name"`
	ServerName                string                                         `json:"server_name"`
	TenantId                  string                                         `json:"tenant_id"`
	Timeouts                  *sqlactivedirectoryadministrator.TimeoutsState `json:"timeouts"`
}