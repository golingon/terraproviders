// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mysqlserverkey "github.com/golingon/terraproviders/azurerm/3.64.0/mysqlserverkey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMysqlServerKey creates a new instance of [MysqlServerKey].
func NewMysqlServerKey(name string, args MysqlServerKeyArgs) *MysqlServerKey {
	return &MysqlServerKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MysqlServerKey)(nil)

// MysqlServerKey represents the Terraform resource azurerm_mysql_server_key.
type MysqlServerKey struct {
	Name      string
	Args      MysqlServerKeyArgs
	state     *mysqlServerKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MysqlServerKey].
func (msk *MysqlServerKey) Type() string {
	return "azurerm_mysql_server_key"
}

// LocalName returns the local name for [MysqlServerKey].
func (msk *MysqlServerKey) LocalName() string {
	return msk.Name
}

// Configuration returns the configuration (args) for [MysqlServerKey].
func (msk *MysqlServerKey) Configuration() interface{} {
	return msk.Args
}

// DependOn is used for other resources to depend on [MysqlServerKey].
func (msk *MysqlServerKey) DependOn() terra.Reference {
	return terra.ReferenceResource(msk)
}

// Dependencies returns the list of resources [MysqlServerKey] depends_on.
func (msk *MysqlServerKey) Dependencies() terra.Dependencies {
	return msk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MysqlServerKey].
func (msk *MysqlServerKey) LifecycleManagement() *terra.Lifecycle {
	return msk.Lifecycle
}

// Attributes returns the attributes for [MysqlServerKey].
func (msk *MysqlServerKey) Attributes() mysqlServerKeyAttributes {
	return mysqlServerKeyAttributes{ref: terra.ReferenceResource(msk)}
}

// ImportState imports the given attribute values into [MysqlServerKey]'s state.
func (msk *MysqlServerKey) ImportState(av io.Reader) error {
	msk.state = &mysqlServerKeyState{}
	if err := json.NewDecoder(av).Decode(msk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", msk.Type(), msk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MysqlServerKey] has state.
func (msk *MysqlServerKey) State() (*mysqlServerKeyState, bool) {
	return msk.state, msk.state != nil
}

// StateMust returns the state for [MysqlServerKey]. Panics if the state is nil.
func (msk *MysqlServerKey) StateMust() *mysqlServerKeyState {
	if msk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", msk.Type(), msk.LocalName()))
	}
	return msk.state
}

// MysqlServerKeyArgs contains the configurations for azurerm_mysql_server_key.
type MysqlServerKeyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultKeyId: string, required
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr" validate:"required"`
	// ServerId: string, required
	ServerId terra.StringValue `hcl:"server_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *mysqlserverkey.Timeouts `hcl:"timeouts,block"`
}
type mysqlServerKeyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mysql_server_key.
func (msk mysqlServerKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(msk.ref.Append("id"))
}

// KeyVaultKeyId returns a reference to field key_vault_key_id of azurerm_mysql_server_key.
func (msk mysqlServerKeyAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(msk.ref.Append("key_vault_key_id"))
}

// ServerId returns a reference to field server_id of azurerm_mysql_server_key.
func (msk mysqlServerKeyAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(msk.ref.Append("server_id"))
}

func (msk mysqlServerKeyAttributes) Timeouts() mysqlserverkey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mysqlserverkey.TimeoutsAttributes](msk.ref.Append("timeouts"))
}

type mysqlServerKeyState struct {
	Id            string                        `json:"id"`
	KeyVaultKeyId string                        `json:"key_vault_key_id"`
	ServerId      string                        `json:"server_id"`
	Timeouts      *mysqlserverkey.TimeoutsState `json:"timeouts"`
}
