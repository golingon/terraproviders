// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storageaccountlocaluser "github.com/golingon/terraproviders/azurerm/3.49.0/storageaccountlocaluser"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewStorageAccountLocalUser(name string, args StorageAccountLocalUserArgs) *StorageAccountLocalUser {
	return &StorageAccountLocalUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageAccountLocalUser)(nil)

type StorageAccountLocalUser struct {
	Name  string
	Args  StorageAccountLocalUserArgs
	state *storageAccountLocalUserState
}

func (salu *StorageAccountLocalUser) Type() string {
	return "azurerm_storage_account_local_user"
}

func (salu *StorageAccountLocalUser) LocalName() string {
	return salu.Name
}

func (salu *StorageAccountLocalUser) Configuration() interface{} {
	return salu.Args
}

func (salu *StorageAccountLocalUser) Attributes() storageAccountLocalUserAttributes {
	return storageAccountLocalUserAttributes{ref: terra.ReferenceResource(salu)}
}

func (salu *StorageAccountLocalUser) ImportState(av io.Reader) error {
	salu.state = &storageAccountLocalUserState{}
	if err := json.NewDecoder(av).Decode(salu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", salu.Type(), salu.LocalName(), err)
	}
	return nil
}

func (salu *StorageAccountLocalUser) State() (*storageAccountLocalUserState, bool) {
	return salu.state, salu.state != nil
}

func (salu *StorageAccountLocalUser) StateMust() *storageAccountLocalUserState {
	if salu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", salu.Type(), salu.LocalName()))
	}
	return salu.state
}

func (salu *StorageAccountLocalUser) DependOn() terra.Reference {
	return terra.ReferenceResource(salu)
}

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
	// DependsOn contains resources that StorageAccountLocalUser depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type storageAccountLocalUserAttributes struct {
	ref terra.Reference
}

func (salu storageAccountLocalUserAttributes) HomeDirectory() terra.StringValue {
	return terra.ReferenceString(salu.ref.Append("home_directory"))
}

func (salu storageAccountLocalUserAttributes) Id() terra.StringValue {
	return terra.ReferenceString(salu.ref.Append("id"))
}

func (salu storageAccountLocalUserAttributes) Name() terra.StringValue {
	return terra.ReferenceString(salu.ref.Append("name"))
}

func (salu storageAccountLocalUserAttributes) Password() terra.StringValue {
	return terra.ReferenceString(salu.ref.Append("password"))
}

func (salu storageAccountLocalUserAttributes) Sid() terra.StringValue {
	return terra.ReferenceString(salu.ref.Append("sid"))
}

func (salu storageAccountLocalUserAttributes) SshKeyEnabled() terra.BoolValue {
	return terra.ReferenceBool(salu.ref.Append("ssh_key_enabled"))
}

func (salu storageAccountLocalUserAttributes) SshPasswordEnabled() terra.BoolValue {
	return terra.ReferenceBool(salu.ref.Append("ssh_password_enabled"))
}

func (salu storageAccountLocalUserAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceString(salu.ref.Append("storage_account_id"))
}

func (salu storageAccountLocalUserAttributes) PermissionScope() terra.ListValue[storageaccountlocaluser.PermissionScopeAttributes] {
	return terra.ReferenceList[storageaccountlocaluser.PermissionScopeAttributes](salu.ref.Append("permission_scope"))
}

func (salu storageAccountLocalUserAttributes) SshAuthorizedKey() terra.ListValue[storageaccountlocaluser.SshAuthorizedKeyAttributes] {
	return terra.ReferenceList[storageaccountlocaluser.SshAuthorizedKeyAttributes](salu.ref.Append("ssh_authorized_key"))
}

func (salu storageAccountLocalUserAttributes) Timeouts() storageaccountlocaluser.TimeoutsAttributes {
	return terra.ReferenceSingle[storageaccountlocaluser.TimeoutsAttributes](salu.ref.Append("timeouts"))
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
