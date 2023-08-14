// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqlmanagedinstanceactivedirectoryadministrator "github.com/golingon/terraproviders/azurerm/3.69.0/mssqlmanagedinstanceactivedirectoryadministrator"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlManagedInstanceActiveDirectoryAdministrator creates a new instance of [MssqlManagedInstanceActiveDirectoryAdministrator].
func NewMssqlManagedInstanceActiveDirectoryAdministrator(name string, args MssqlManagedInstanceActiveDirectoryAdministratorArgs) *MssqlManagedInstanceActiveDirectoryAdministrator {
	return &MssqlManagedInstanceActiveDirectoryAdministrator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlManagedInstanceActiveDirectoryAdministrator)(nil)

// MssqlManagedInstanceActiveDirectoryAdministrator represents the Terraform resource azurerm_mssql_managed_instance_active_directory_administrator.
type MssqlManagedInstanceActiveDirectoryAdministrator struct {
	Name      string
	Args      MssqlManagedInstanceActiveDirectoryAdministratorArgs
	state     *mssqlManagedInstanceActiveDirectoryAdministratorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlManagedInstanceActiveDirectoryAdministrator].
func (mmiada *MssqlManagedInstanceActiveDirectoryAdministrator) Type() string {
	return "azurerm_mssql_managed_instance_active_directory_administrator"
}

// LocalName returns the local name for [MssqlManagedInstanceActiveDirectoryAdministrator].
func (mmiada *MssqlManagedInstanceActiveDirectoryAdministrator) LocalName() string {
	return mmiada.Name
}

// Configuration returns the configuration (args) for [MssqlManagedInstanceActiveDirectoryAdministrator].
func (mmiada *MssqlManagedInstanceActiveDirectoryAdministrator) Configuration() interface{} {
	return mmiada.Args
}

// DependOn is used for other resources to depend on [MssqlManagedInstanceActiveDirectoryAdministrator].
func (mmiada *MssqlManagedInstanceActiveDirectoryAdministrator) DependOn() terra.Reference {
	return terra.ReferenceResource(mmiada)
}

// Dependencies returns the list of resources [MssqlManagedInstanceActiveDirectoryAdministrator] depends_on.
func (mmiada *MssqlManagedInstanceActiveDirectoryAdministrator) Dependencies() terra.Dependencies {
	return mmiada.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlManagedInstanceActiveDirectoryAdministrator].
func (mmiada *MssqlManagedInstanceActiveDirectoryAdministrator) LifecycleManagement() *terra.Lifecycle {
	return mmiada.Lifecycle
}

// Attributes returns the attributes for [MssqlManagedInstanceActiveDirectoryAdministrator].
func (mmiada *MssqlManagedInstanceActiveDirectoryAdministrator) Attributes() mssqlManagedInstanceActiveDirectoryAdministratorAttributes {
	return mssqlManagedInstanceActiveDirectoryAdministratorAttributes{ref: terra.ReferenceResource(mmiada)}
}

// ImportState imports the given attribute values into [MssqlManagedInstanceActiveDirectoryAdministrator]'s state.
func (mmiada *MssqlManagedInstanceActiveDirectoryAdministrator) ImportState(av io.Reader) error {
	mmiada.state = &mssqlManagedInstanceActiveDirectoryAdministratorState{}
	if err := json.NewDecoder(av).Decode(mmiada.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mmiada.Type(), mmiada.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlManagedInstanceActiveDirectoryAdministrator] has state.
func (mmiada *MssqlManagedInstanceActiveDirectoryAdministrator) State() (*mssqlManagedInstanceActiveDirectoryAdministratorState, bool) {
	return mmiada.state, mmiada.state != nil
}

// StateMust returns the state for [MssqlManagedInstanceActiveDirectoryAdministrator]. Panics if the state is nil.
func (mmiada *MssqlManagedInstanceActiveDirectoryAdministrator) StateMust() *mssqlManagedInstanceActiveDirectoryAdministratorState {
	if mmiada.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mmiada.Type(), mmiada.LocalName()))
	}
	return mmiada.state
}

// MssqlManagedInstanceActiveDirectoryAdministratorArgs contains the configurations for azurerm_mssql_managed_instance_active_directory_administrator.
type MssqlManagedInstanceActiveDirectoryAdministratorArgs struct {
	// AzureadAuthenticationOnly: bool, optional
	AzureadAuthenticationOnly terra.BoolValue `hcl:"azuread_authentication_only,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LoginUsername: string, required
	LoginUsername terra.StringValue `hcl:"login_username,attr" validate:"required"`
	// ManagedInstanceId: string, required
	ManagedInstanceId terra.StringValue `hcl:"managed_instance_id,attr" validate:"required"`
	// ObjectId: string, required
	ObjectId terra.StringValue `hcl:"object_id,attr" validate:"required"`
	// TenantId: string, required
	TenantId terra.StringValue `hcl:"tenant_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *mssqlmanagedinstanceactivedirectoryadministrator.Timeouts `hcl:"timeouts,block"`
}
type mssqlManagedInstanceActiveDirectoryAdministratorAttributes struct {
	ref terra.Reference
}

// AzureadAuthenticationOnly returns a reference to field azuread_authentication_only of azurerm_mssql_managed_instance_active_directory_administrator.
func (mmiada mssqlManagedInstanceActiveDirectoryAdministratorAttributes) AzureadAuthenticationOnly() terra.BoolValue {
	return terra.ReferenceAsBool(mmiada.ref.Append("azuread_authentication_only"))
}

// Id returns a reference to field id of azurerm_mssql_managed_instance_active_directory_administrator.
func (mmiada mssqlManagedInstanceActiveDirectoryAdministratorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mmiada.ref.Append("id"))
}

// LoginUsername returns a reference to field login_username of azurerm_mssql_managed_instance_active_directory_administrator.
func (mmiada mssqlManagedInstanceActiveDirectoryAdministratorAttributes) LoginUsername() terra.StringValue {
	return terra.ReferenceAsString(mmiada.ref.Append("login_username"))
}

// ManagedInstanceId returns a reference to field managed_instance_id of azurerm_mssql_managed_instance_active_directory_administrator.
func (mmiada mssqlManagedInstanceActiveDirectoryAdministratorAttributes) ManagedInstanceId() terra.StringValue {
	return terra.ReferenceAsString(mmiada.ref.Append("managed_instance_id"))
}

// ObjectId returns a reference to field object_id of azurerm_mssql_managed_instance_active_directory_administrator.
func (mmiada mssqlManagedInstanceActiveDirectoryAdministratorAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(mmiada.ref.Append("object_id"))
}

// TenantId returns a reference to field tenant_id of azurerm_mssql_managed_instance_active_directory_administrator.
func (mmiada mssqlManagedInstanceActiveDirectoryAdministratorAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(mmiada.ref.Append("tenant_id"))
}

func (mmiada mssqlManagedInstanceActiveDirectoryAdministratorAttributes) Timeouts() mssqlmanagedinstanceactivedirectoryadministrator.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqlmanagedinstanceactivedirectoryadministrator.TimeoutsAttributes](mmiada.ref.Append("timeouts"))
}

type mssqlManagedInstanceActiveDirectoryAdministratorState struct {
	AzureadAuthenticationOnly bool                                                            `json:"azuread_authentication_only"`
	Id                        string                                                          `json:"id"`
	LoginUsername             string                                                          `json:"login_username"`
	ManagedInstanceId         string                                                          `json:"managed_instance_id"`
	ObjectId                  string                                                          `json:"object_id"`
	TenantId                  string                                                          `json:"tenant_id"`
	Timeouts                  *mssqlmanagedinstanceactivedirectoryadministrator.TimeoutsState `json:"timeouts"`
}
