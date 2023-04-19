// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqlmanagedinstancetransparentdataencryption "github.com/golingon/terraproviders/azurerm/3.52.0/mssqlmanagedinstancetransparentdataencryption"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlManagedInstanceTransparentDataEncryption creates a new instance of [MssqlManagedInstanceTransparentDataEncryption].
func NewMssqlManagedInstanceTransparentDataEncryption(name string, args MssqlManagedInstanceTransparentDataEncryptionArgs) *MssqlManagedInstanceTransparentDataEncryption {
	return &MssqlManagedInstanceTransparentDataEncryption{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlManagedInstanceTransparentDataEncryption)(nil)

// MssqlManagedInstanceTransparentDataEncryption represents the Terraform resource azurerm_mssql_managed_instance_transparent_data_encryption.
type MssqlManagedInstanceTransparentDataEncryption struct {
	Name      string
	Args      MssqlManagedInstanceTransparentDataEncryptionArgs
	state     *mssqlManagedInstanceTransparentDataEncryptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlManagedInstanceTransparentDataEncryption].
func (mmitde *MssqlManagedInstanceTransparentDataEncryption) Type() string {
	return "azurerm_mssql_managed_instance_transparent_data_encryption"
}

// LocalName returns the local name for [MssqlManagedInstanceTransparentDataEncryption].
func (mmitde *MssqlManagedInstanceTransparentDataEncryption) LocalName() string {
	return mmitde.Name
}

// Configuration returns the configuration (args) for [MssqlManagedInstanceTransparentDataEncryption].
func (mmitde *MssqlManagedInstanceTransparentDataEncryption) Configuration() interface{} {
	return mmitde.Args
}

// DependOn is used for other resources to depend on [MssqlManagedInstanceTransparentDataEncryption].
func (mmitde *MssqlManagedInstanceTransparentDataEncryption) DependOn() terra.Reference {
	return terra.ReferenceResource(mmitde)
}

// Dependencies returns the list of resources [MssqlManagedInstanceTransparentDataEncryption] depends_on.
func (mmitde *MssqlManagedInstanceTransparentDataEncryption) Dependencies() terra.Dependencies {
	return mmitde.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlManagedInstanceTransparentDataEncryption].
func (mmitde *MssqlManagedInstanceTransparentDataEncryption) LifecycleManagement() *terra.Lifecycle {
	return mmitde.Lifecycle
}

// Attributes returns the attributes for [MssqlManagedInstanceTransparentDataEncryption].
func (mmitde *MssqlManagedInstanceTransparentDataEncryption) Attributes() mssqlManagedInstanceTransparentDataEncryptionAttributes {
	return mssqlManagedInstanceTransparentDataEncryptionAttributes{ref: terra.ReferenceResource(mmitde)}
}

// ImportState imports the given attribute values into [MssqlManagedInstanceTransparentDataEncryption]'s state.
func (mmitde *MssqlManagedInstanceTransparentDataEncryption) ImportState(av io.Reader) error {
	mmitde.state = &mssqlManagedInstanceTransparentDataEncryptionState{}
	if err := json.NewDecoder(av).Decode(mmitde.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mmitde.Type(), mmitde.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlManagedInstanceTransparentDataEncryption] has state.
func (mmitde *MssqlManagedInstanceTransparentDataEncryption) State() (*mssqlManagedInstanceTransparentDataEncryptionState, bool) {
	return mmitde.state, mmitde.state != nil
}

// StateMust returns the state for [MssqlManagedInstanceTransparentDataEncryption]. Panics if the state is nil.
func (mmitde *MssqlManagedInstanceTransparentDataEncryption) StateMust() *mssqlManagedInstanceTransparentDataEncryptionState {
	if mmitde.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mmitde.Type(), mmitde.LocalName()))
	}
	return mmitde.state
}

// MssqlManagedInstanceTransparentDataEncryptionArgs contains the configurations for azurerm_mssql_managed_instance_transparent_data_encryption.
type MssqlManagedInstanceTransparentDataEncryptionArgs struct {
	// AutoRotationEnabled: bool, optional
	AutoRotationEnabled terra.BoolValue `hcl:"auto_rotation_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultKeyId: string, optional
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr"`
	// ManagedInstanceId: string, required
	ManagedInstanceId terra.StringValue `hcl:"managed_instance_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *mssqlmanagedinstancetransparentdataencryption.Timeouts `hcl:"timeouts,block"`
}
type mssqlManagedInstanceTransparentDataEncryptionAttributes struct {
	ref terra.Reference
}

// AutoRotationEnabled returns a reference to field auto_rotation_enabled of azurerm_mssql_managed_instance_transparent_data_encryption.
func (mmitde mssqlManagedInstanceTransparentDataEncryptionAttributes) AutoRotationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mmitde.ref.Append("auto_rotation_enabled"))
}

// Id returns a reference to field id of azurerm_mssql_managed_instance_transparent_data_encryption.
func (mmitde mssqlManagedInstanceTransparentDataEncryptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mmitde.ref.Append("id"))
}

// KeyVaultKeyId returns a reference to field key_vault_key_id of azurerm_mssql_managed_instance_transparent_data_encryption.
func (mmitde mssqlManagedInstanceTransparentDataEncryptionAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(mmitde.ref.Append("key_vault_key_id"))
}

// ManagedInstanceId returns a reference to field managed_instance_id of azurerm_mssql_managed_instance_transparent_data_encryption.
func (mmitde mssqlManagedInstanceTransparentDataEncryptionAttributes) ManagedInstanceId() terra.StringValue {
	return terra.ReferenceAsString(mmitde.ref.Append("managed_instance_id"))
}

func (mmitde mssqlManagedInstanceTransparentDataEncryptionAttributes) Timeouts() mssqlmanagedinstancetransparentdataencryption.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqlmanagedinstancetransparentdataencryption.TimeoutsAttributes](mmitde.ref.Append("timeouts"))
}

type mssqlManagedInstanceTransparentDataEncryptionState struct {
	AutoRotationEnabled bool                                                         `json:"auto_rotation_enabled"`
	Id                  string                                                       `json:"id"`
	KeyVaultKeyId       string                                                       `json:"key_vault_key_id"`
	ManagedInstanceId   string                                                       `json:"managed_instance_id"`
	Timeouts            *mssqlmanagedinstancetransparentdataencryption.TimeoutsState `json:"timeouts"`
}
