// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualmachineextension "github.com/golingon/terraproviders/azurerm/3.69.0/virtualmachineextension"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualMachineExtension creates a new instance of [VirtualMachineExtension].
func NewVirtualMachineExtension(name string, args VirtualMachineExtensionArgs) *VirtualMachineExtension {
	return &VirtualMachineExtension{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualMachineExtension)(nil)

// VirtualMachineExtension represents the Terraform resource azurerm_virtual_machine_extension.
type VirtualMachineExtension struct {
	Name      string
	Args      VirtualMachineExtensionArgs
	state     *virtualMachineExtensionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualMachineExtension].
func (vme *VirtualMachineExtension) Type() string {
	return "azurerm_virtual_machine_extension"
}

// LocalName returns the local name for [VirtualMachineExtension].
func (vme *VirtualMachineExtension) LocalName() string {
	return vme.Name
}

// Configuration returns the configuration (args) for [VirtualMachineExtension].
func (vme *VirtualMachineExtension) Configuration() interface{} {
	return vme.Args
}

// DependOn is used for other resources to depend on [VirtualMachineExtension].
func (vme *VirtualMachineExtension) DependOn() terra.Reference {
	return terra.ReferenceResource(vme)
}

// Dependencies returns the list of resources [VirtualMachineExtension] depends_on.
func (vme *VirtualMachineExtension) Dependencies() terra.Dependencies {
	return vme.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualMachineExtension].
func (vme *VirtualMachineExtension) LifecycleManagement() *terra.Lifecycle {
	return vme.Lifecycle
}

// Attributes returns the attributes for [VirtualMachineExtension].
func (vme *VirtualMachineExtension) Attributes() virtualMachineExtensionAttributes {
	return virtualMachineExtensionAttributes{ref: terra.ReferenceResource(vme)}
}

// ImportState imports the given attribute values into [VirtualMachineExtension]'s state.
func (vme *VirtualMachineExtension) ImportState(av io.Reader) error {
	vme.state = &virtualMachineExtensionState{}
	if err := json.NewDecoder(av).Decode(vme.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vme.Type(), vme.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualMachineExtension] has state.
func (vme *VirtualMachineExtension) State() (*virtualMachineExtensionState, bool) {
	return vme.state, vme.state != nil
}

// StateMust returns the state for [VirtualMachineExtension]. Panics if the state is nil.
func (vme *VirtualMachineExtension) StateMust() *virtualMachineExtensionState {
	if vme.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vme.Type(), vme.LocalName()))
	}
	return vme.state
}

// VirtualMachineExtensionArgs contains the configurations for azurerm_virtual_machine_extension.
type VirtualMachineExtensionArgs struct {
	// AutoUpgradeMinorVersion: bool, optional
	AutoUpgradeMinorVersion terra.BoolValue `hcl:"auto_upgrade_minor_version,attr"`
	// AutomaticUpgradeEnabled: bool, optional
	AutomaticUpgradeEnabled terra.BoolValue `hcl:"automatic_upgrade_enabled,attr"`
	// FailureSuppressionEnabled: bool, optional
	FailureSuppressionEnabled terra.BoolValue `hcl:"failure_suppression_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProtectedSettings: string, optional
	ProtectedSettings terra.StringValue `hcl:"protected_settings,attr"`
	// Publisher: string, required
	Publisher terra.StringValue `hcl:"publisher,attr" validate:"required"`
	// Settings: string, optional
	Settings terra.StringValue `hcl:"settings,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// TypeHandlerVersion: string, required
	TypeHandlerVersion terra.StringValue `hcl:"type_handler_version,attr" validate:"required"`
	// VirtualMachineId: string, required
	VirtualMachineId terra.StringValue `hcl:"virtual_machine_id,attr" validate:"required"`
	// ProtectedSettingsFromKeyVault: optional
	ProtectedSettingsFromKeyVault *virtualmachineextension.ProtectedSettingsFromKeyVault `hcl:"protected_settings_from_key_vault,block"`
	// Timeouts: optional
	Timeouts *virtualmachineextension.Timeouts `hcl:"timeouts,block"`
}
type virtualMachineExtensionAttributes struct {
	ref terra.Reference
}

// AutoUpgradeMinorVersion returns a reference to field auto_upgrade_minor_version of azurerm_virtual_machine_extension.
func (vme virtualMachineExtensionAttributes) AutoUpgradeMinorVersion() terra.BoolValue {
	return terra.ReferenceAsBool(vme.ref.Append("auto_upgrade_minor_version"))
}

// AutomaticUpgradeEnabled returns a reference to field automatic_upgrade_enabled of azurerm_virtual_machine_extension.
func (vme virtualMachineExtensionAttributes) AutomaticUpgradeEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(vme.ref.Append("automatic_upgrade_enabled"))
}

// FailureSuppressionEnabled returns a reference to field failure_suppression_enabled of azurerm_virtual_machine_extension.
func (vme virtualMachineExtensionAttributes) FailureSuppressionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(vme.ref.Append("failure_suppression_enabled"))
}

// Id returns a reference to field id of azurerm_virtual_machine_extension.
func (vme virtualMachineExtensionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vme.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_virtual_machine_extension.
func (vme virtualMachineExtensionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vme.ref.Append("name"))
}

// ProtectedSettings returns a reference to field protected_settings of azurerm_virtual_machine_extension.
func (vme virtualMachineExtensionAttributes) ProtectedSettings() terra.StringValue {
	return terra.ReferenceAsString(vme.ref.Append("protected_settings"))
}

// Publisher returns a reference to field publisher of azurerm_virtual_machine_extension.
func (vme virtualMachineExtensionAttributes) Publisher() terra.StringValue {
	return terra.ReferenceAsString(vme.ref.Append("publisher"))
}

// Settings returns a reference to field settings of azurerm_virtual_machine_extension.
func (vme virtualMachineExtensionAttributes) Settings() terra.StringValue {
	return terra.ReferenceAsString(vme.ref.Append("settings"))
}

// Tags returns a reference to field tags of azurerm_virtual_machine_extension.
func (vme virtualMachineExtensionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vme.ref.Append("tags"))
}

// Type returns a reference to field type of azurerm_virtual_machine_extension.
func (vme virtualMachineExtensionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vme.ref.Append("type"))
}

// TypeHandlerVersion returns a reference to field type_handler_version of azurerm_virtual_machine_extension.
func (vme virtualMachineExtensionAttributes) TypeHandlerVersion() terra.StringValue {
	return terra.ReferenceAsString(vme.ref.Append("type_handler_version"))
}

// VirtualMachineId returns a reference to field virtual_machine_id of azurerm_virtual_machine_extension.
func (vme virtualMachineExtensionAttributes) VirtualMachineId() terra.StringValue {
	return terra.ReferenceAsString(vme.ref.Append("virtual_machine_id"))
}

func (vme virtualMachineExtensionAttributes) ProtectedSettingsFromKeyVault() terra.ListValue[virtualmachineextension.ProtectedSettingsFromKeyVaultAttributes] {
	return terra.ReferenceAsList[virtualmachineextension.ProtectedSettingsFromKeyVaultAttributes](vme.ref.Append("protected_settings_from_key_vault"))
}

func (vme virtualMachineExtensionAttributes) Timeouts() virtualmachineextension.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualmachineextension.TimeoutsAttributes](vme.ref.Append("timeouts"))
}

type virtualMachineExtensionState struct {
	AutoUpgradeMinorVersion       bool                                                         `json:"auto_upgrade_minor_version"`
	AutomaticUpgradeEnabled       bool                                                         `json:"automatic_upgrade_enabled"`
	FailureSuppressionEnabled     bool                                                         `json:"failure_suppression_enabled"`
	Id                            string                                                       `json:"id"`
	Name                          string                                                       `json:"name"`
	ProtectedSettings             string                                                       `json:"protected_settings"`
	Publisher                     string                                                       `json:"publisher"`
	Settings                      string                                                       `json:"settings"`
	Tags                          map[string]string                                            `json:"tags"`
	Type                          string                                                       `json:"type"`
	TypeHandlerVersion            string                                                       `json:"type_handler_version"`
	VirtualMachineId              string                                                       `json:"virtual_machine_id"`
	ProtectedSettingsFromKeyVault []virtualmachineextension.ProtectedSettingsFromKeyVaultState `json:"protected_settings_from_key_vault"`
	Timeouts                      *virtualmachineextension.TimeoutsState                       `json:"timeouts"`
}