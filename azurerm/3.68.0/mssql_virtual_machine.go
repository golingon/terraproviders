// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqlvirtualmachine "github.com/golingon/terraproviders/azurerm/3.68.0/mssqlvirtualmachine"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlVirtualMachine creates a new instance of [MssqlVirtualMachine].
func NewMssqlVirtualMachine(name string, args MssqlVirtualMachineArgs) *MssqlVirtualMachine {
	return &MssqlVirtualMachine{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlVirtualMachine)(nil)

// MssqlVirtualMachine represents the Terraform resource azurerm_mssql_virtual_machine.
type MssqlVirtualMachine struct {
	Name      string
	Args      MssqlVirtualMachineArgs
	state     *mssqlVirtualMachineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlVirtualMachine].
func (mvm *MssqlVirtualMachine) Type() string {
	return "azurerm_mssql_virtual_machine"
}

// LocalName returns the local name for [MssqlVirtualMachine].
func (mvm *MssqlVirtualMachine) LocalName() string {
	return mvm.Name
}

// Configuration returns the configuration (args) for [MssqlVirtualMachine].
func (mvm *MssqlVirtualMachine) Configuration() interface{} {
	return mvm.Args
}

// DependOn is used for other resources to depend on [MssqlVirtualMachine].
func (mvm *MssqlVirtualMachine) DependOn() terra.Reference {
	return terra.ReferenceResource(mvm)
}

// Dependencies returns the list of resources [MssqlVirtualMachine] depends_on.
func (mvm *MssqlVirtualMachine) Dependencies() terra.Dependencies {
	return mvm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlVirtualMachine].
func (mvm *MssqlVirtualMachine) LifecycleManagement() *terra.Lifecycle {
	return mvm.Lifecycle
}

// Attributes returns the attributes for [MssqlVirtualMachine].
func (mvm *MssqlVirtualMachine) Attributes() mssqlVirtualMachineAttributes {
	return mssqlVirtualMachineAttributes{ref: terra.ReferenceResource(mvm)}
}

// ImportState imports the given attribute values into [MssqlVirtualMachine]'s state.
func (mvm *MssqlVirtualMachine) ImportState(av io.Reader) error {
	mvm.state = &mssqlVirtualMachineState{}
	if err := json.NewDecoder(av).Decode(mvm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mvm.Type(), mvm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlVirtualMachine] has state.
func (mvm *MssqlVirtualMachine) State() (*mssqlVirtualMachineState, bool) {
	return mvm.state, mvm.state != nil
}

// StateMust returns the state for [MssqlVirtualMachine]. Panics if the state is nil.
func (mvm *MssqlVirtualMachine) StateMust() *mssqlVirtualMachineState {
	if mvm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mvm.Type(), mvm.LocalName()))
	}
	return mvm.state
}

// MssqlVirtualMachineArgs contains the configurations for azurerm_mssql_virtual_machine.
type MssqlVirtualMachineArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RServicesEnabled: bool, optional
	RServicesEnabled terra.BoolValue `hcl:"r_services_enabled,attr"`
	// SqlConnectivityPort: number, optional
	SqlConnectivityPort terra.NumberValue `hcl:"sql_connectivity_port,attr"`
	// SqlConnectivityType: string, optional
	SqlConnectivityType terra.StringValue `hcl:"sql_connectivity_type,attr"`
	// SqlConnectivityUpdatePassword: string, optional
	SqlConnectivityUpdatePassword terra.StringValue `hcl:"sql_connectivity_update_password,attr"`
	// SqlConnectivityUpdateUsername: string, optional
	SqlConnectivityUpdateUsername terra.StringValue `hcl:"sql_connectivity_update_username,attr"`
	// SqlLicenseType: string, optional
	SqlLicenseType terra.StringValue `hcl:"sql_license_type,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VirtualMachineId: string, required
	VirtualMachineId terra.StringValue `hcl:"virtual_machine_id,attr" validate:"required"`
	// Assessment: optional
	Assessment *mssqlvirtualmachine.Assessment `hcl:"assessment,block"`
	// AutoBackup: optional
	AutoBackup *mssqlvirtualmachine.AutoBackup `hcl:"auto_backup,block"`
	// AutoPatching: optional
	AutoPatching *mssqlvirtualmachine.AutoPatching `hcl:"auto_patching,block"`
	// KeyVaultCredential: optional
	KeyVaultCredential *mssqlvirtualmachine.KeyVaultCredential `hcl:"key_vault_credential,block"`
	// SqlInstance: optional
	SqlInstance *mssqlvirtualmachine.SqlInstance `hcl:"sql_instance,block"`
	// StorageConfiguration: optional
	StorageConfiguration *mssqlvirtualmachine.StorageConfiguration `hcl:"storage_configuration,block"`
	// Timeouts: optional
	Timeouts *mssqlvirtualmachine.Timeouts `hcl:"timeouts,block"`
}
type mssqlVirtualMachineAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mssql_virtual_machine.
func (mvm mssqlVirtualMachineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mvm.ref.Append("id"))
}

// RServicesEnabled returns a reference to field r_services_enabled of azurerm_mssql_virtual_machine.
func (mvm mssqlVirtualMachineAttributes) RServicesEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mvm.ref.Append("r_services_enabled"))
}

// SqlConnectivityPort returns a reference to field sql_connectivity_port of azurerm_mssql_virtual_machine.
func (mvm mssqlVirtualMachineAttributes) SqlConnectivityPort() terra.NumberValue {
	return terra.ReferenceAsNumber(mvm.ref.Append("sql_connectivity_port"))
}

// SqlConnectivityType returns a reference to field sql_connectivity_type of azurerm_mssql_virtual_machine.
func (mvm mssqlVirtualMachineAttributes) SqlConnectivityType() terra.StringValue {
	return terra.ReferenceAsString(mvm.ref.Append("sql_connectivity_type"))
}

// SqlConnectivityUpdatePassword returns a reference to field sql_connectivity_update_password of azurerm_mssql_virtual_machine.
func (mvm mssqlVirtualMachineAttributes) SqlConnectivityUpdatePassword() terra.StringValue {
	return terra.ReferenceAsString(mvm.ref.Append("sql_connectivity_update_password"))
}

// SqlConnectivityUpdateUsername returns a reference to field sql_connectivity_update_username of azurerm_mssql_virtual_machine.
func (mvm mssqlVirtualMachineAttributes) SqlConnectivityUpdateUsername() terra.StringValue {
	return terra.ReferenceAsString(mvm.ref.Append("sql_connectivity_update_username"))
}

// SqlLicenseType returns a reference to field sql_license_type of azurerm_mssql_virtual_machine.
func (mvm mssqlVirtualMachineAttributes) SqlLicenseType() terra.StringValue {
	return terra.ReferenceAsString(mvm.ref.Append("sql_license_type"))
}

// Tags returns a reference to field tags of azurerm_mssql_virtual_machine.
func (mvm mssqlVirtualMachineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mvm.ref.Append("tags"))
}

// VirtualMachineId returns a reference to field virtual_machine_id of azurerm_mssql_virtual_machine.
func (mvm mssqlVirtualMachineAttributes) VirtualMachineId() terra.StringValue {
	return terra.ReferenceAsString(mvm.ref.Append("virtual_machine_id"))
}

func (mvm mssqlVirtualMachineAttributes) Assessment() terra.ListValue[mssqlvirtualmachine.AssessmentAttributes] {
	return terra.ReferenceAsList[mssqlvirtualmachine.AssessmentAttributes](mvm.ref.Append("assessment"))
}

func (mvm mssqlVirtualMachineAttributes) AutoBackup() terra.ListValue[mssqlvirtualmachine.AutoBackupAttributes] {
	return terra.ReferenceAsList[mssqlvirtualmachine.AutoBackupAttributes](mvm.ref.Append("auto_backup"))
}

func (mvm mssqlVirtualMachineAttributes) AutoPatching() terra.ListValue[mssqlvirtualmachine.AutoPatchingAttributes] {
	return terra.ReferenceAsList[mssqlvirtualmachine.AutoPatchingAttributes](mvm.ref.Append("auto_patching"))
}

func (mvm mssqlVirtualMachineAttributes) KeyVaultCredential() terra.ListValue[mssqlvirtualmachine.KeyVaultCredentialAttributes] {
	return terra.ReferenceAsList[mssqlvirtualmachine.KeyVaultCredentialAttributes](mvm.ref.Append("key_vault_credential"))
}

func (mvm mssqlVirtualMachineAttributes) SqlInstance() terra.ListValue[mssqlvirtualmachine.SqlInstanceAttributes] {
	return terra.ReferenceAsList[mssqlvirtualmachine.SqlInstanceAttributes](mvm.ref.Append("sql_instance"))
}

func (mvm mssqlVirtualMachineAttributes) StorageConfiguration() terra.ListValue[mssqlvirtualmachine.StorageConfigurationAttributes] {
	return terra.ReferenceAsList[mssqlvirtualmachine.StorageConfigurationAttributes](mvm.ref.Append("storage_configuration"))
}

func (mvm mssqlVirtualMachineAttributes) Timeouts() mssqlvirtualmachine.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqlvirtualmachine.TimeoutsAttributes](mvm.ref.Append("timeouts"))
}

type mssqlVirtualMachineState struct {
	Id                            string                                          `json:"id"`
	RServicesEnabled              bool                                            `json:"r_services_enabled"`
	SqlConnectivityPort           float64                                         `json:"sql_connectivity_port"`
	SqlConnectivityType           string                                          `json:"sql_connectivity_type"`
	SqlConnectivityUpdatePassword string                                          `json:"sql_connectivity_update_password"`
	SqlConnectivityUpdateUsername string                                          `json:"sql_connectivity_update_username"`
	SqlLicenseType                string                                          `json:"sql_license_type"`
	Tags                          map[string]string                               `json:"tags"`
	VirtualMachineId              string                                          `json:"virtual_machine_id"`
	Assessment                    []mssqlvirtualmachine.AssessmentState           `json:"assessment"`
	AutoBackup                    []mssqlvirtualmachine.AutoBackupState           `json:"auto_backup"`
	AutoPatching                  []mssqlvirtualmachine.AutoPatchingState         `json:"auto_patching"`
	KeyVaultCredential            []mssqlvirtualmachine.KeyVaultCredentialState   `json:"key_vault_credential"`
	SqlInstance                   []mssqlvirtualmachine.SqlInstanceState          `json:"sql_instance"`
	StorageConfiguration          []mssqlvirtualmachine.StorageConfigurationState `json:"storage_configuration"`
	Timeouts                      *mssqlvirtualmachine.TimeoutsState              `json:"timeouts"`
}
