// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mysqlactivedirectoryadministrator "github.com/golingon/terraproviders/azurerm/3.63.0/mysqlactivedirectoryadministrator"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMysqlActiveDirectoryAdministrator creates a new instance of [MysqlActiveDirectoryAdministrator].
func NewMysqlActiveDirectoryAdministrator(name string, args MysqlActiveDirectoryAdministratorArgs) *MysqlActiveDirectoryAdministrator {
	return &MysqlActiveDirectoryAdministrator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MysqlActiveDirectoryAdministrator)(nil)

// MysqlActiveDirectoryAdministrator represents the Terraform resource azurerm_mysql_active_directory_administrator.
type MysqlActiveDirectoryAdministrator struct {
	Name      string
	Args      MysqlActiveDirectoryAdministratorArgs
	state     *mysqlActiveDirectoryAdministratorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MysqlActiveDirectoryAdministrator].
func (mada *MysqlActiveDirectoryAdministrator) Type() string {
	return "azurerm_mysql_active_directory_administrator"
}

// LocalName returns the local name for [MysqlActiveDirectoryAdministrator].
func (mada *MysqlActiveDirectoryAdministrator) LocalName() string {
	return mada.Name
}

// Configuration returns the configuration (args) for [MysqlActiveDirectoryAdministrator].
func (mada *MysqlActiveDirectoryAdministrator) Configuration() interface{} {
	return mada.Args
}

// DependOn is used for other resources to depend on [MysqlActiveDirectoryAdministrator].
func (mada *MysqlActiveDirectoryAdministrator) DependOn() terra.Reference {
	return terra.ReferenceResource(mada)
}

// Dependencies returns the list of resources [MysqlActiveDirectoryAdministrator] depends_on.
func (mada *MysqlActiveDirectoryAdministrator) Dependencies() terra.Dependencies {
	return mada.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MysqlActiveDirectoryAdministrator].
func (mada *MysqlActiveDirectoryAdministrator) LifecycleManagement() *terra.Lifecycle {
	return mada.Lifecycle
}

// Attributes returns the attributes for [MysqlActiveDirectoryAdministrator].
func (mada *MysqlActiveDirectoryAdministrator) Attributes() mysqlActiveDirectoryAdministratorAttributes {
	return mysqlActiveDirectoryAdministratorAttributes{ref: terra.ReferenceResource(mada)}
}

// ImportState imports the given attribute values into [MysqlActiveDirectoryAdministrator]'s state.
func (mada *MysqlActiveDirectoryAdministrator) ImportState(av io.Reader) error {
	mada.state = &mysqlActiveDirectoryAdministratorState{}
	if err := json.NewDecoder(av).Decode(mada.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mada.Type(), mada.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MysqlActiveDirectoryAdministrator] has state.
func (mada *MysqlActiveDirectoryAdministrator) State() (*mysqlActiveDirectoryAdministratorState, bool) {
	return mada.state, mada.state != nil
}

// StateMust returns the state for [MysqlActiveDirectoryAdministrator]. Panics if the state is nil.
func (mada *MysqlActiveDirectoryAdministrator) StateMust() *mysqlActiveDirectoryAdministratorState {
	if mada.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mada.Type(), mada.LocalName()))
	}
	return mada.state
}

// MysqlActiveDirectoryAdministratorArgs contains the configurations for azurerm_mysql_active_directory_administrator.
type MysqlActiveDirectoryAdministratorArgs struct {
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
	Timeouts *mysqlactivedirectoryadministrator.Timeouts `hcl:"timeouts,block"`
}
type mysqlActiveDirectoryAdministratorAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mysql_active_directory_administrator.
func (mada mysqlActiveDirectoryAdministratorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mada.ref.Append("id"))
}

// Login returns a reference to field login of azurerm_mysql_active_directory_administrator.
func (mada mysqlActiveDirectoryAdministratorAttributes) Login() terra.StringValue {
	return terra.ReferenceAsString(mada.ref.Append("login"))
}

// ObjectId returns a reference to field object_id of azurerm_mysql_active_directory_administrator.
func (mada mysqlActiveDirectoryAdministratorAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(mada.ref.Append("object_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mysql_active_directory_administrator.
func (mada mysqlActiveDirectoryAdministratorAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mada.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_mysql_active_directory_administrator.
func (mada mysqlActiveDirectoryAdministratorAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(mada.ref.Append("server_name"))
}

// TenantId returns a reference to field tenant_id of azurerm_mysql_active_directory_administrator.
func (mada mysqlActiveDirectoryAdministratorAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(mada.ref.Append("tenant_id"))
}

func (mada mysqlActiveDirectoryAdministratorAttributes) Timeouts() mysqlactivedirectoryadministrator.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mysqlactivedirectoryadministrator.TimeoutsAttributes](mada.ref.Append("timeouts"))
}

type mysqlActiveDirectoryAdministratorState struct {
	Id                string                                           `json:"id"`
	Login             string                                           `json:"login"`
	ObjectId          string                                           `json:"object_id"`
	ResourceGroupName string                                           `json:"resource_group_name"`
	ServerName        string                                           `json:"server_name"`
	TenantId          string                                           `json:"tenant_id"`
	Timeouts          *mysqlactivedirectoryadministrator.TimeoutsState `json:"timeouts"`
}
