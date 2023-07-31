// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	recoveryservicesvault "github.com/golingon/terraproviders/azurerm/3.67.0/recoveryservicesvault"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRecoveryServicesVault creates a new instance of [RecoveryServicesVault].
func NewRecoveryServicesVault(name string, args RecoveryServicesVaultArgs) *RecoveryServicesVault {
	return &RecoveryServicesVault{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RecoveryServicesVault)(nil)

// RecoveryServicesVault represents the Terraform resource azurerm_recovery_services_vault.
type RecoveryServicesVault struct {
	Name      string
	Args      RecoveryServicesVaultArgs
	state     *recoveryServicesVaultState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RecoveryServicesVault].
func (rsv *RecoveryServicesVault) Type() string {
	return "azurerm_recovery_services_vault"
}

// LocalName returns the local name for [RecoveryServicesVault].
func (rsv *RecoveryServicesVault) LocalName() string {
	return rsv.Name
}

// Configuration returns the configuration (args) for [RecoveryServicesVault].
func (rsv *RecoveryServicesVault) Configuration() interface{} {
	return rsv.Args
}

// DependOn is used for other resources to depend on [RecoveryServicesVault].
func (rsv *RecoveryServicesVault) DependOn() terra.Reference {
	return terra.ReferenceResource(rsv)
}

// Dependencies returns the list of resources [RecoveryServicesVault] depends_on.
func (rsv *RecoveryServicesVault) Dependencies() terra.Dependencies {
	return rsv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RecoveryServicesVault].
func (rsv *RecoveryServicesVault) LifecycleManagement() *terra.Lifecycle {
	return rsv.Lifecycle
}

// Attributes returns the attributes for [RecoveryServicesVault].
func (rsv *RecoveryServicesVault) Attributes() recoveryServicesVaultAttributes {
	return recoveryServicesVaultAttributes{ref: terra.ReferenceResource(rsv)}
}

// ImportState imports the given attribute values into [RecoveryServicesVault]'s state.
func (rsv *RecoveryServicesVault) ImportState(av io.Reader) error {
	rsv.state = &recoveryServicesVaultState{}
	if err := json.NewDecoder(av).Decode(rsv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rsv.Type(), rsv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RecoveryServicesVault] has state.
func (rsv *RecoveryServicesVault) State() (*recoveryServicesVaultState, bool) {
	return rsv.state, rsv.state != nil
}

// StateMust returns the state for [RecoveryServicesVault]. Panics if the state is nil.
func (rsv *RecoveryServicesVault) StateMust() *recoveryServicesVaultState {
	if rsv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rsv.Type(), rsv.LocalName()))
	}
	return rsv.state
}

// RecoveryServicesVaultArgs contains the configurations for azurerm_recovery_services_vault.
type RecoveryServicesVaultArgs struct {
	// ClassicVmwareReplicationEnabled: bool, optional
	ClassicVmwareReplicationEnabled terra.BoolValue `hcl:"classic_vmware_replication_enabled,attr"`
	// CrossRegionRestoreEnabled: bool, optional
	CrossRegionRestoreEnabled terra.BoolValue `hcl:"cross_region_restore_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Immutability: string, optional
	Immutability terra.StringValue `hcl:"immutability,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// SoftDeleteEnabled: bool, optional
	SoftDeleteEnabled terra.BoolValue `hcl:"soft_delete_enabled,attr"`
	// StorageModeType: string, optional
	StorageModeType terra.StringValue `hcl:"storage_mode_type,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Encryption: optional
	Encryption *recoveryservicesvault.Encryption `hcl:"encryption,block"`
	// Identity: optional
	Identity *recoveryservicesvault.Identity `hcl:"identity,block"`
	// Monitoring: optional
	Monitoring *recoveryservicesvault.Monitoring `hcl:"monitoring,block"`
	// Timeouts: optional
	Timeouts *recoveryservicesvault.Timeouts `hcl:"timeouts,block"`
}
type recoveryServicesVaultAttributes struct {
	ref terra.Reference
}

// ClassicVmwareReplicationEnabled returns a reference to field classic_vmware_replication_enabled of azurerm_recovery_services_vault.
func (rsv recoveryServicesVaultAttributes) ClassicVmwareReplicationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(rsv.ref.Append("classic_vmware_replication_enabled"))
}

// CrossRegionRestoreEnabled returns a reference to field cross_region_restore_enabled of azurerm_recovery_services_vault.
func (rsv recoveryServicesVaultAttributes) CrossRegionRestoreEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(rsv.ref.Append("cross_region_restore_enabled"))
}

// Id returns a reference to field id of azurerm_recovery_services_vault.
func (rsv recoveryServicesVaultAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rsv.ref.Append("id"))
}

// Immutability returns a reference to field immutability of azurerm_recovery_services_vault.
func (rsv recoveryServicesVaultAttributes) Immutability() terra.StringValue {
	return terra.ReferenceAsString(rsv.ref.Append("immutability"))
}

// Location returns a reference to field location of azurerm_recovery_services_vault.
func (rsv recoveryServicesVaultAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(rsv.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_recovery_services_vault.
func (rsv recoveryServicesVaultAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rsv.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_recovery_services_vault.
func (rsv recoveryServicesVaultAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(rsv.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_recovery_services_vault.
func (rsv recoveryServicesVaultAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(rsv.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_recovery_services_vault.
func (rsv recoveryServicesVaultAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(rsv.ref.Append("sku"))
}

// SoftDeleteEnabled returns a reference to field soft_delete_enabled of azurerm_recovery_services_vault.
func (rsv recoveryServicesVaultAttributes) SoftDeleteEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(rsv.ref.Append("soft_delete_enabled"))
}

// StorageModeType returns a reference to field storage_mode_type of azurerm_recovery_services_vault.
func (rsv recoveryServicesVaultAttributes) StorageModeType() terra.StringValue {
	return terra.ReferenceAsString(rsv.ref.Append("storage_mode_type"))
}

// Tags returns a reference to field tags of azurerm_recovery_services_vault.
func (rsv recoveryServicesVaultAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rsv.ref.Append("tags"))
}

func (rsv recoveryServicesVaultAttributes) Encryption() terra.ListValue[recoveryservicesvault.EncryptionAttributes] {
	return terra.ReferenceAsList[recoveryservicesvault.EncryptionAttributes](rsv.ref.Append("encryption"))
}

func (rsv recoveryServicesVaultAttributes) Identity() terra.ListValue[recoveryservicesvault.IdentityAttributes] {
	return terra.ReferenceAsList[recoveryservicesvault.IdentityAttributes](rsv.ref.Append("identity"))
}

func (rsv recoveryServicesVaultAttributes) Monitoring() terra.ListValue[recoveryservicesvault.MonitoringAttributes] {
	return terra.ReferenceAsList[recoveryservicesvault.MonitoringAttributes](rsv.ref.Append("monitoring"))
}

func (rsv recoveryServicesVaultAttributes) Timeouts() recoveryservicesvault.TimeoutsAttributes {
	return terra.ReferenceAsSingle[recoveryservicesvault.TimeoutsAttributes](rsv.ref.Append("timeouts"))
}

type recoveryServicesVaultState struct {
	ClassicVmwareReplicationEnabled bool                                    `json:"classic_vmware_replication_enabled"`
	CrossRegionRestoreEnabled       bool                                    `json:"cross_region_restore_enabled"`
	Id                              string                                  `json:"id"`
	Immutability                    string                                  `json:"immutability"`
	Location                        string                                  `json:"location"`
	Name                            string                                  `json:"name"`
	PublicNetworkAccessEnabled      bool                                    `json:"public_network_access_enabled"`
	ResourceGroupName               string                                  `json:"resource_group_name"`
	Sku                             string                                  `json:"sku"`
	SoftDeleteEnabled               bool                                    `json:"soft_delete_enabled"`
	StorageModeType                 string                                  `json:"storage_mode_type"`
	Tags                            map[string]string                       `json:"tags"`
	Encryption                      []recoveryservicesvault.EncryptionState `json:"encryption"`
	Identity                        []recoveryservicesvault.IdentityState   `json:"identity"`
	Monitoring                      []recoveryservicesvault.MonitoringState `json:"monitoring"`
	Timeouts                        *recoveryservicesvault.TimeoutsState    `json:"timeouts"`
}
