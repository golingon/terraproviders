// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storageaccountlocaluser "github.com/golingon/terraproviders/azurerm/3.63.0/storageaccountlocaluser"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageAccountLocalUser creates a new instance of [StorageAccountLocalUser].
func NewStorageAccountLocalUser(name string, args StorageAccountLocalUserArgs) *StorageAccountLocalUser {
	return &StorageAccountLocalUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageAccountLocalUser)(nil)

// StorageAccountLocalUser represents the Terraform resource azurerm_storage_account_local_user.
type StorageAccountLocalUser struct {
	Name      string
	Args      StorageAccountLocalUserArgs
	state     *storageAccountLocalUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageAccountLocalUser].
func (salu *StorageAccountLocalUser) Type() string {
	return "azurerm_storage_account_local_user"
}

// LocalName returns the local name for [StorageAccountLocalUser].
func (salu *StorageAccountLocalUser) LocalName() string {
	return salu.Name
}

// Configuration returns the configuration (args) for [StorageAccountLocalUser].
func (salu *StorageAccountLocalUser) Configuration() interface{} {
	return salu.Args
}

// DependOn is used for other resources to depend on [StorageAccountLocalUser].
func (salu *StorageAccountLocalUser) DependOn() terra.Reference {
	return terra.ReferenceResource(salu)
}

// Dependencies returns the list of resources [StorageAccountLocalUser] depends_on.
func (salu *StorageAccountLocalUser) Dependencies() terra.Dependencies {
	return salu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageAccountLocalUser].
func (salu *StorageAccountLocalUser) LifecycleManagement() *terra.Lifecycle {
	return salu.Lifecycle
}

// Attributes returns the attributes for [StorageAccountLocalUser].
func (salu *StorageAccountLocalUser) Attributes() storageAccountLocalUserAttributes {
	return storageAccountLocalUserAttributes{ref: terra.ReferenceResource(salu)}
}

// ImportState imports the given attribute values into [StorageAccountLocalUser]'s state.
func (salu *StorageAccountLocalUser) ImportState(av io.Reader) error {
	salu.state = &storageAccountLocalUserState{}
	if err := json.NewDecoder(av).Decode(salu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", salu.Type(), salu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageAccountLocalUser] has state.
func (salu *StorageAccountLocalUser) State() (*storageAccountLocalUserState, bool) {
	return salu.state, salu.state != nil
}

// StateMust returns the state for [StorageAccountLocalUser]. Panics if the state is nil.
func (salu *StorageAccountLocalUser) StateMust() *storageAccountLocalUserState {
	if salu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", salu.Type(), salu.LocalName()))
	}
	return salu.state
}

// StorageAccountLocalUserArgs contains the configurations for azurerm_storage_account_local_user.
type StorageAccountLocalUserArgs struct {
	// HomeDirectory: string, optional
	HomeDirectory terra.StringValue `hcl:"home_directory,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SshKeyEnabled: bool, optional
	SshKeyEnabled terra.BoolValue `hcl:"ssh_key_enabled,attr"`
	// SshPasswordEnabled: bool, optional
	SshPasswordEnabled terra.BoolValue `hcl:"ssh_password_enabled,attr"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// PermissionScope: min=0
	PermissionScope []storageaccountlocaluser.PermissionScope `hcl:"permission_scope,block" validate:"min=0"`
	// SshAuthorizedKey: min=0
	SshAuthorizedKey []storageaccountlocaluser.SshAuthorizedKey `hcl:"ssh_authorized_key,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *storageaccountlocaluser.Timeouts `hcl:"timeouts,block"`
}
type storageAccountLocalUserAttributes struct {
	ref terra.Reference
}

// HomeDirectory returns a reference to field home_directory of azurerm_storage_account_local_user.
func (salu storageAccountLocalUserAttributes) HomeDirectory() terra.StringValue {
	return terra.ReferenceAsString(salu.ref.Append("home_directory"))
}

// Id returns a reference to field id of azurerm_storage_account_local_user.
func (salu storageAccountLocalUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(salu.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_storage_account_local_user.
func (salu storageAccountLocalUserAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(salu.ref.Append("name"))
}

// Password returns a reference to field password of azurerm_storage_account_local_user.
func (salu storageAccountLocalUserAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(salu.ref.Append("password"))
}

// Sid returns a reference to field sid of azurerm_storage_account_local_user.
func (salu storageAccountLocalUserAttributes) Sid() terra.StringValue {
	return terra.ReferenceAsString(salu.ref.Append("sid"))
}

// SshKeyEnabled returns a reference to field ssh_key_enabled of azurerm_storage_account_local_user.
func (salu storageAccountLocalUserAttributes) SshKeyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(salu.ref.Append("ssh_key_enabled"))
}

// SshPasswordEnabled returns a reference to field ssh_password_enabled of azurerm_storage_account_local_user.
func (salu storageAccountLocalUserAttributes) SshPasswordEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(salu.ref.Append("ssh_password_enabled"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_storage_account_local_user.
func (salu storageAccountLocalUserAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(salu.ref.Append("storage_account_id"))
}

func (salu storageAccountLocalUserAttributes) PermissionScope() terra.ListValue[storageaccountlocaluser.PermissionScopeAttributes] {
	return terra.ReferenceAsList[storageaccountlocaluser.PermissionScopeAttributes](salu.ref.Append("permission_scope"))
}

func (salu storageAccountLocalUserAttributes) SshAuthorizedKey() terra.ListValue[storageaccountlocaluser.SshAuthorizedKeyAttributes] {
	return terra.ReferenceAsList[storageaccountlocaluser.SshAuthorizedKeyAttributes](salu.ref.Append("ssh_authorized_key"))
}

func (salu storageAccountLocalUserAttributes) Timeouts() storageaccountlocaluser.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storageaccountlocaluser.TimeoutsAttributes](salu.ref.Append("timeouts"))
}

type storageAccountLocalUserState struct {
	HomeDirectory      string                                          `json:"home_directory"`
	Id                 string                                          `json:"id"`
	Name               string                                          `json:"name"`
	Password           string                                          `json:"password"`
	Sid                string                                          `json:"sid"`
	SshKeyEnabled      bool                                            `json:"ssh_key_enabled"`
	SshPasswordEnabled bool                                            `json:"ssh_password_enabled"`
	StorageAccountId   string                                          `json:"storage_account_id"`
	PermissionScope    []storageaccountlocaluser.PermissionScopeState  `json:"permission_scope"`
	SshAuthorizedKey   []storageaccountlocaluser.SshAuthorizedKeyState `json:"ssh_authorized_key"`
	Timeouts           *storageaccountlocaluser.TimeoutsState          `json:"timeouts"`
}
