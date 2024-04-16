// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_postgresql_flexible_server_active_directory_administrator

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

// Resource represents the Terraform resource azurerm_postgresql_flexible_server_active_directory_administrator.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermPostgresqlFlexibleServerActiveDirectoryAdministratorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (apfsada *Resource) Type() string {
	return "azurerm_postgresql_flexible_server_active_directory_administrator"
}

// LocalName returns the local name for [Resource].
func (apfsada *Resource) LocalName() string {
	return apfsada.Name
}

// Configuration returns the configuration (args) for [Resource].
func (apfsada *Resource) Configuration() interface{} {
	return apfsada.Args
}

// DependOn is used for other resources to depend on [Resource].
func (apfsada *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(apfsada)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (apfsada *Resource) Dependencies() terra.Dependencies {
	return apfsada.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (apfsada *Resource) LifecycleManagement() *terra.Lifecycle {
	return apfsada.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (apfsada *Resource) Attributes() azurermPostgresqlFlexibleServerActiveDirectoryAdministratorAttributes {
	return azurermPostgresqlFlexibleServerActiveDirectoryAdministratorAttributes{ref: terra.ReferenceResource(apfsada)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (apfsada *Resource) ImportState(state io.Reader) error {
	apfsada.state = &azurermPostgresqlFlexibleServerActiveDirectoryAdministratorState{}
	if err := json.NewDecoder(state).Decode(apfsada.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", apfsada.Type(), apfsada.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (apfsada *Resource) State() (*azurermPostgresqlFlexibleServerActiveDirectoryAdministratorState, bool) {
	return apfsada.state, apfsada.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (apfsada *Resource) StateMust() *azurermPostgresqlFlexibleServerActiveDirectoryAdministratorState {
	if apfsada.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", apfsada.Type(), apfsada.LocalName()))
	}
	return apfsada.state
}

// Args contains the configurations for azurerm_postgresql_flexible_server_active_directory_administrator.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermPostgresqlFlexibleServerActiveDirectoryAdministratorAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_postgresql_flexible_server_active_directory_administrator.
func (apfsada azurermPostgresqlFlexibleServerActiveDirectoryAdministratorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(apfsada.ref.Append("id"))
}

// ObjectId returns a reference to field object_id of azurerm_postgresql_flexible_server_active_directory_administrator.
func (apfsada azurermPostgresqlFlexibleServerActiveDirectoryAdministratorAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(apfsada.ref.Append("object_id"))
}

// PrincipalName returns a reference to field principal_name of azurerm_postgresql_flexible_server_active_directory_administrator.
func (apfsada azurermPostgresqlFlexibleServerActiveDirectoryAdministratorAttributes) PrincipalName() terra.StringValue {
	return terra.ReferenceAsString(apfsada.ref.Append("principal_name"))
}

// PrincipalType returns a reference to field principal_type of azurerm_postgresql_flexible_server_active_directory_administrator.
func (apfsada azurermPostgresqlFlexibleServerActiveDirectoryAdministratorAttributes) PrincipalType() terra.StringValue {
	return terra.ReferenceAsString(apfsada.ref.Append("principal_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_postgresql_flexible_server_active_directory_administrator.
func (apfsada azurermPostgresqlFlexibleServerActiveDirectoryAdministratorAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(apfsada.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_postgresql_flexible_server_active_directory_administrator.
func (apfsada azurermPostgresqlFlexibleServerActiveDirectoryAdministratorAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(apfsada.ref.Append("server_name"))
}

// TenantId returns a reference to field tenant_id of azurerm_postgresql_flexible_server_active_directory_administrator.
func (apfsada azurermPostgresqlFlexibleServerActiveDirectoryAdministratorAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(apfsada.ref.Append("tenant_id"))
}

func (apfsada azurermPostgresqlFlexibleServerActiveDirectoryAdministratorAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](apfsada.ref.Append("timeouts"))
}

type azurermPostgresqlFlexibleServerActiveDirectoryAdministratorState struct {
	Id                string         `json:"id"`
	ObjectId          string         `json:"object_id"`
	PrincipalName     string         `json:"principal_name"`
	PrincipalType     string         `json:"principal_type"`
	ResourceGroupName string         `json:"resource_group_name"`
	ServerName        string         `json:"server_name"`
	TenantId          string         `json:"tenant_id"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
