// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mysqlflexibleserveractivedirectoryadministrator "github.com/golingon/terraproviders/azurerm/3.66.0/mysqlflexibleserveractivedirectoryadministrator"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMysqlFlexibleServerActiveDirectoryAdministrator creates a new instance of [MysqlFlexibleServerActiveDirectoryAdministrator].
func NewMysqlFlexibleServerActiveDirectoryAdministrator(name string, args MysqlFlexibleServerActiveDirectoryAdministratorArgs) *MysqlFlexibleServerActiveDirectoryAdministrator {
	return &MysqlFlexibleServerActiveDirectoryAdministrator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MysqlFlexibleServerActiveDirectoryAdministrator)(nil)

// MysqlFlexibleServerActiveDirectoryAdministrator represents the Terraform resource azurerm_mysql_flexible_server_active_directory_administrator.
type MysqlFlexibleServerActiveDirectoryAdministrator struct {
	Name      string
	Args      MysqlFlexibleServerActiveDirectoryAdministratorArgs
	state     *mysqlFlexibleServerActiveDirectoryAdministratorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MysqlFlexibleServerActiveDirectoryAdministrator].
func (mfsada *MysqlFlexibleServerActiveDirectoryAdministrator) Type() string {
	return "azurerm_mysql_flexible_server_active_directory_administrator"
}

// LocalName returns the local name for [MysqlFlexibleServerActiveDirectoryAdministrator].
func (mfsada *MysqlFlexibleServerActiveDirectoryAdministrator) LocalName() string {
	return mfsada.Name
}

// Configuration returns the configuration (args) for [MysqlFlexibleServerActiveDirectoryAdministrator].
func (mfsada *MysqlFlexibleServerActiveDirectoryAdministrator) Configuration() interface{} {
	return mfsada.Args
}

// DependOn is used for other resources to depend on [MysqlFlexibleServerActiveDirectoryAdministrator].
func (mfsada *MysqlFlexibleServerActiveDirectoryAdministrator) DependOn() terra.Reference {
	return terra.ReferenceResource(mfsada)
}

// Dependencies returns the list of resources [MysqlFlexibleServerActiveDirectoryAdministrator] depends_on.
func (mfsada *MysqlFlexibleServerActiveDirectoryAdministrator) Dependencies() terra.Dependencies {
	return mfsada.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MysqlFlexibleServerActiveDirectoryAdministrator].
func (mfsada *MysqlFlexibleServerActiveDirectoryAdministrator) LifecycleManagement() *terra.Lifecycle {
	return mfsada.Lifecycle
}

// Attributes returns the attributes for [MysqlFlexibleServerActiveDirectoryAdministrator].
func (mfsada *MysqlFlexibleServerActiveDirectoryAdministrator) Attributes() mysqlFlexibleServerActiveDirectoryAdministratorAttributes {
	return mysqlFlexibleServerActiveDirectoryAdministratorAttributes{ref: terra.ReferenceResource(mfsada)}
}

// ImportState imports the given attribute values into [MysqlFlexibleServerActiveDirectoryAdministrator]'s state.
func (mfsada *MysqlFlexibleServerActiveDirectoryAdministrator) ImportState(av io.Reader) error {
	mfsada.state = &mysqlFlexibleServerActiveDirectoryAdministratorState{}
	if err := json.NewDecoder(av).Decode(mfsada.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mfsada.Type(), mfsada.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MysqlFlexibleServerActiveDirectoryAdministrator] has state.
func (mfsada *MysqlFlexibleServerActiveDirectoryAdministrator) State() (*mysqlFlexibleServerActiveDirectoryAdministratorState, bool) {
	return mfsada.state, mfsada.state != nil
}

// StateMust returns the state for [MysqlFlexibleServerActiveDirectoryAdministrator]. Panics if the state is nil.
func (mfsada *MysqlFlexibleServerActiveDirectoryAdministrator) StateMust() *mysqlFlexibleServerActiveDirectoryAdministratorState {
	if mfsada.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mfsada.Type(), mfsada.LocalName()))
	}
	return mfsada.state
}

// MysqlFlexibleServerActiveDirectoryAdministratorArgs contains the configurations for azurerm_mysql_flexible_server_active_directory_administrator.
type MysqlFlexibleServerActiveDirectoryAdministratorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityId: string, required
	IdentityId terra.StringValue `hcl:"identity_id,attr" validate:"required"`
	// Login: string, required
	Login terra.StringValue `hcl:"login,attr" validate:"required"`
	// ObjectId: string, required
	ObjectId terra.StringValue `hcl:"object_id,attr" validate:"required"`
	// ServerId: string, required
	ServerId terra.StringValue `hcl:"server_id,attr" validate:"required"`
	// TenantId: string, required
	TenantId terra.StringValue `hcl:"tenant_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *mysqlflexibleserveractivedirectoryadministrator.Timeouts `hcl:"timeouts,block"`
}
type mysqlFlexibleServerActiveDirectoryAdministratorAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mysql_flexible_server_active_directory_administrator.
func (mfsada mysqlFlexibleServerActiveDirectoryAdministratorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mfsada.ref.Append("id"))
}

// IdentityId returns a reference to field identity_id of azurerm_mysql_flexible_server_active_directory_administrator.
func (mfsada mysqlFlexibleServerActiveDirectoryAdministratorAttributes) IdentityId() terra.StringValue {
	return terra.ReferenceAsString(mfsada.ref.Append("identity_id"))
}

// Login returns a reference to field login of azurerm_mysql_flexible_server_active_directory_administrator.
func (mfsada mysqlFlexibleServerActiveDirectoryAdministratorAttributes) Login() terra.StringValue {
	return terra.ReferenceAsString(mfsada.ref.Append("login"))
}

// ObjectId returns a reference to field object_id of azurerm_mysql_flexible_server_active_directory_administrator.
func (mfsada mysqlFlexibleServerActiveDirectoryAdministratorAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(mfsada.ref.Append("object_id"))
}

// ServerId returns a reference to field server_id of azurerm_mysql_flexible_server_active_directory_administrator.
func (mfsada mysqlFlexibleServerActiveDirectoryAdministratorAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(mfsada.ref.Append("server_id"))
}

// TenantId returns a reference to field tenant_id of azurerm_mysql_flexible_server_active_directory_administrator.
func (mfsada mysqlFlexibleServerActiveDirectoryAdministratorAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(mfsada.ref.Append("tenant_id"))
}

func (mfsada mysqlFlexibleServerActiveDirectoryAdministratorAttributes) Timeouts() mysqlflexibleserveractivedirectoryadministrator.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mysqlflexibleserveractivedirectoryadministrator.TimeoutsAttributes](mfsada.ref.Append("timeouts"))
}

type mysqlFlexibleServerActiveDirectoryAdministratorState struct {
	Id         string                                                         `json:"id"`
	IdentityId string                                                         `json:"identity_id"`
	Login      string                                                         `json:"login"`
	ObjectId   string                                                         `json:"object_id"`
	ServerId   string                                                         `json:"server_id"`
	TenantId   string                                                         `json:"tenant_id"`
	Timeouts   *mysqlflexibleserveractivedirectoryadministrator.TimeoutsState `json:"timeouts"`
}
