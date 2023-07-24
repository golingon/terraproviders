// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqlservertransparentdataencryption "github.com/golingon/terraproviders/azurerm/3.66.0/mssqlservertransparentdataencryption"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlServerTransparentDataEncryption creates a new instance of [MssqlServerTransparentDataEncryption].
func NewMssqlServerTransparentDataEncryption(name string, args MssqlServerTransparentDataEncryptionArgs) *MssqlServerTransparentDataEncryption {
	return &MssqlServerTransparentDataEncryption{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlServerTransparentDataEncryption)(nil)

// MssqlServerTransparentDataEncryption represents the Terraform resource azurerm_mssql_server_transparent_data_encryption.
type MssqlServerTransparentDataEncryption struct {
	Name      string
	Args      MssqlServerTransparentDataEncryptionArgs
	state     *mssqlServerTransparentDataEncryptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlServerTransparentDataEncryption].
func (mstde *MssqlServerTransparentDataEncryption) Type() string {
	return "azurerm_mssql_server_transparent_data_encryption"
}

// LocalName returns the local name for [MssqlServerTransparentDataEncryption].
func (mstde *MssqlServerTransparentDataEncryption) LocalName() string {
	return mstde.Name
}

// Configuration returns the configuration (args) for [MssqlServerTransparentDataEncryption].
func (mstde *MssqlServerTransparentDataEncryption) Configuration() interface{} {
	return mstde.Args
}

// DependOn is used for other resources to depend on [MssqlServerTransparentDataEncryption].
func (mstde *MssqlServerTransparentDataEncryption) DependOn() terra.Reference {
	return terra.ReferenceResource(mstde)
}

// Dependencies returns the list of resources [MssqlServerTransparentDataEncryption] depends_on.
func (mstde *MssqlServerTransparentDataEncryption) Dependencies() terra.Dependencies {
	return mstde.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlServerTransparentDataEncryption].
func (mstde *MssqlServerTransparentDataEncryption) LifecycleManagement() *terra.Lifecycle {
	return mstde.Lifecycle
}

// Attributes returns the attributes for [MssqlServerTransparentDataEncryption].
func (mstde *MssqlServerTransparentDataEncryption) Attributes() mssqlServerTransparentDataEncryptionAttributes {
	return mssqlServerTransparentDataEncryptionAttributes{ref: terra.ReferenceResource(mstde)}
}

// ImportState imports the given attribute values into [MssqlServerTransparentDataEncryption]'s state.
func (mstde *MssqlServerTransparentDataEncryption) ImportState(av io.Reader) error {
	mstde.state = &mssqlServerTransparentDataEncryptionState{}
	if err := json.NewDecoder(av).Decode(mstde.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mstde.Type(), mstde.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlServerTransparentDataEncryption] has state.
func (mstde *MssqlServerTransparentDataEncryption) State() (*mssqlServerTransparentDataEncryptionState, bool) {
	return mstde.state, mstde.state != nil
}

// StateMust returns the state for [MssqlServerTransparentDataEncryption]. Panics if the state is nil.
func (mstde *MssqlServerTransparentDataEncryption) StateMust() *mssqlServerTransparentDataEncryptionState {
	if mstde.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mstde.Type(), mstde.LocalName()))
	}
	return mstde.state
}

// MssqlServerTransparentDataEncryptionArgs contains the configurations for azurerm_mssql_server_transparent_data_encryption.
type MssqlServerTransparentDataEncryptionArgs struct {
	// AutoRotationEnabled: bool, optional
	AutoRotationEnabled terra.BoolValue `hcl:"auto_rotation_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultKeyId: string, optional
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr"`
	// ServerId: string, required
	ServerId terra.StringValue `hcl:"server_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *mssqlservertransparentdataencryption.Timeouts `hcl:"timeouts,block"`
}
type mssqlServerTransparentDataEncryptionAttributes struct {
	ref terra.Reference
}

// AutoRotationEnabled returns a reference to field auto_rotation_enabled of azurerm_mssql_server_transparent_data_encryption.
func (mstde mssqlServerTransparentDataEncryptionAttributes) AutoRotationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mstde.ref.Append("auto_rotation_enabled"))
}

// Id returns a reference to field id of azurerm_mssql_server_transparent_data_encryption.
func (mstde mssqlServerTransparentDataEncryptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mstde.ref.Append("id"))
}

// KeyVaultKeyId returns a reference to field key_vault_key_id of azurerm_mssql_server_transparent_data_encryption.
func (mstde mssqlServerTransparentDataEncryptionAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(mstde.ref.Append("key_vault_key_id"))
}

// ServerId returns a reference to field server_id of azurerm_mssql_server_transparent_data_encryption.
func (mstde mssqlServerTransparentDataEncryptionAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(mstde.ref.Append("server_id"))
}

func (mstde mssqlServerTransparentDataEncryptionAttributes) Timeouts() mssqlservertransparentdataencryption.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqlservertransparentdataencryption.TimeoutsAttributes](mstde.ref.Append("timeouts"))
}

type mssqlServerTransparentDataEncryptionState struct {
	AutoRotationEnabled bool                                                `json:"auto_rotation_enabled"`
	Id                  string                                              `json:"id"`
	KeyVaultKeyId       string                                              `json:"key_vault_key_id"`
	ServerId            string                                              `json:"server_id"`
	Timeouts            *mssqlservertransparentdataencryption.TimeoutsState `json:"timeouts"`
}
