// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_mysql_server_key

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

// Resource represents the Terraform resource azurerm_mysql_server_key.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermMysqlServerKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (amsk *Resource) Type() string {
	return "azurerm_mysql_server_key"
}

// LocalName returns the local name for [Resource].
func (amsk *Resource) LocalName() string {
	return amsk.Name
}

// Configuration returns the configuration (args) for [Resource].
func (amsk *Resource) Configuration() interface{} {
	return amsk.Args
}

// DependOn is used for other resources to depend on [Resource].
func (amsk *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(amsk)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (amsk *Resource) Dependencies() terra.Dependencies {
	return amsk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (amsk *Resource) LifecycleManagement() *terra.Lifecycle {
	return amsk.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (amsk *Resource) Attributes() azurermMysqlServerKeyAttributes {
	return azurermMysqlServerKeyAttributes{ref: terra.ReferenceResource(amsk)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (amsk *Resource) ImportState(state io.Reader) error {
	amsk.state = &azurermMysqlServerKeyState{}
	if err := json.NewDecoder(state).Decode(amsk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amsk.Type(), amsk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (amsk *Resource) State() (*azurermMysqlServerKeyState, bool) {
	return amsk.state, amsk.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (amsk *Resource) StateMust() *azurermMysqlServerKeyState {
	if amsk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amsk.Type(), amsk.LocalName()))
	}
	return amsk.state
}

// Args contains the configurations for azurerm_mysql_server_key.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultKeyId: string, required
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr" validate:"required"`
	// ServerId: string, required
	ServerId terra.StringValue `hcl:"server_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermMysqlServerKeyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mysql_server_key.
func (amsk azurermMysqlServerKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amsk.ref.Append("id"))
}

// KeyVaultKeyId returns a reference to field key_vault_key_id of azurerm_mysql_server_key.
func (amsk azurermMysqlServerKeyAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(amsk.ref.Append("key_vault_key_id"))
}

// ServerId returns a reference to field server_id of azurerm_mysql_server_key.
func (amsk azurermMysqlServerKeyAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(amsk.ref.Append("server_id"))
}

func (amsk azurermMysqlServerKeyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](amsk.ref.Append("timeouts"))
}

type azurermMysqlServerKeyState struct {
	Id            string         `json:"id"`
	KeyVaultKeyId string         `json:"key_vault_key_id"`
	ServerId      string         `json:"server_id"`
	Timeouts      *TimeoutsState `json:"timeouts"`
}
