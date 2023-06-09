// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualmachinescalesetextension "github.com/golingon/terraproviders/azurerm/3.52.0/virtualmachinescalesetextension"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualMachineScaleSetExtension creates a new instance of [VirtualMachineScaleSetExtension].
func NewVirtualMachineScaleSetExtension(name string, args VirtualMachineScaleSetExtensionArgs) *VirtualMachineScaleSetExtension {
	return &VirtualMachineScaleSetExtension{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualMachineScaleSetExtension)(nil)

// VirtualMachineScaleSetExtension represents the Terraform resource azurerm_virtual_machine_scale_set_extension.
type VirtualMachineScaleSetExtension struct {
	Name      string
	Args      VirtualMachineScaleSetExtensionArgs
	state     *virtualMachineScaleSetExtensionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualMachineScaleSetExtension].
func (vmsse *VirtualMachineScaleSetExtension) Type() string {
	return "azurerm_virtual_machine_scale_set_extension"
}

// LocalName returns the local name for [VirtualMachineScaleSetExtension].
func (vmsse *VirtualMachineScaleSetExtension) LocalName() string {
	return vmsse.Name
}

// Configuration returns the configuration (args) for [VirtualMachineScaleSetExtension].
func (vmsse *VirtualMachineScaleSetExtension) Configuration() interface{} {
	return vmsse.Args
}

// DependOn is used for other resources to depend on [VirtualMachineScaleSetExtension].
func (vmsse *VirtualMachineScaleSetExtension) DependOn() terra.Reference {
	return terra.ReferenceResource(vmsse)
}

// Dependencies returns the list of resources [VirtualMachineScaleSetExtension] depends_on.
func (vmsse *VirtualMachineScaleSetExtension) Dependencies() terra.Dependencies {
	return vmsse.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualMachineScaleSetExtension].
func (vmsse *VirtualMachineScaleSetExtension) LifecycleManagement() *terra.Lifecycle {
	return vmsse.Lifecycle
}

// Attributes returns the attributes for [VirtualMachineScaleSetExtension].
func (vmsse *VirtualMachineScaleSetExtension) Attributes() virtualMachineScaleSetExtensionAttributes {
	return virtualMachineScaleSetExtensionAttributes{ref: terra.ReferenceResource(vmsse)}
}

// ImportState imports the given attribute values into [VirtualMachineScaleSetExtension]'s state.
func (vmsse *VirtualMachineScaleSetExtension) ImportState(av io.Reader) error {
	vmsse.state = &virtualMachineScaleSetExtensionState{}
	if err := json.NewDecoder(av).Decode(vmsse.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vmsse.Type(), vmsse.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualMachineScaleSetExtension] has state.
func (vmsse *VirtualMachineScaleSetExtension) State() (*virtualMachineScaleSetExtensionState, bool) {
	return vmsse.state, vmsse.state != nil
}

// StateMust returns the state for [VirtualMachineScaleSetExtension]. Panics if the state is nil.
func (vmsse *VirtualMachineScaleSetExtension) StateMust() *virtualMachineScaleSetExtensionState {
	if vmsse.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vmsse.Type(), vmsse.LocalName()))
	}
	return vmsse.state
}

// VirtualMachineScaleSetExtensionArgs contains the configurations for azurerm_virtual_machine_scale_set_extension.
type VirtualMachineScaleSetExtensionArgs struct {
	// AutoUpgradeMinorVersion: bool, optional
	AutoUpgradeMinorVersion terra.BoolValue `hcl:"auto_upgrade_minor_version,attr"`
	// AutomaticUpgradeEnabled: bool, optional
	AutomaticUpgradeEnabled terra.BoolValue `hcl:"automatic_upgrade_enabled,attr"`
	// FailureSuppressionEnabled: bool, optional
	FailureSuppressionEnabled terra.BoolValue `hcl:"failure_suppression_enabled,attr"`
	// ForceUpdateTag: string, optional
	ForceUpdateTag terra.StringValue `hcl:"force_update_tag,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProtectedSettings: string, optional
	ProtectedSettings terra.StringValue `hcl:"protected_settings,attr"`
	// ProvisionAfterExtensions: list of string, optional
	ProvisionAfterExtensions terra.ListValue[terra.StringValue] `hcl:"provision_after_extensions,attr"`
	// Publisher: string, required
	Publisher terra.StringValue `hcl:"publisher,attr" validate:"required"`
	// Settings: string, optional
	Settings terra.StringValue `hcl:"settings,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// TypeHandlerVersion: string, required
	TypeHandlerVersion terra.StringValue `hcl:"type_handler_version,attr" validate:"required"`
	// VirtualMachineScaleSetId: string, required
	VirtualMachineScaleSetId terra.StringValue `hcl:"virtual_machine_scale_set_id,attr" validate:"required"`
	// ProtectedSettingsFromKeyVault: optional
	ProtectedSettingsFromKeyVault *virtualmachinescalesetextension.ProtectedSettingsFromKeyVault `hcl:"protected_settings_from_key_vault,block"`
	// Timeouts: optional
	Timeouts *virtualmachinescalesetextension.Timeouts `hcl:"timeouts,block"`
}
type virtualMachineScaleSetExtensionAttributes struct {
	ref terra.Reference
}

// AutoUpgradeMinorVersion returns a reference to field auto_upgrade_minor_version of azurerm_virtual_machine_scale_set_extension.
func (vmsse virtualMachineScaleSetExtensionAttributes) AutoUpgradeMinorVersion() terra.BoolValue {
	return terra.ReferenceAsBool(vmsse.ref.Append("auto_upgrade_minor_version"))
}

// AutomaticUpgradeEnabled returns a reference to field automatic_upgrade_enabled of azurerm_virtual_machine_scale_set_extension.
func (vmsse virtualMachineScaleSetExtensionAttributes) AutomaticUpgradeEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(vmsse.ref.Append("automatic_upgrade_enabled"))
}

// FailureSuppressionEnabled returns a reference to field failure_suppression_enabled of azurerm_virtual_machine_scale_set_extension.
func (vmsse virtualMachineScaleSetExtensionAttributes) FailureSuppressionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(vmsse.ref.Append("failure_suppression_enabled"))
}

// ForceUpdateTag returns a reference to field force_update_tag of azurerm_virtual_machine_scale_set_extension.
func (vmsse virtualMachineScaleSetExtensionAttributes) ForceUpdateTag() terra.StringValue {
	return terra.ReferenceAsString(vmsse.ref.Append("force_update_tag"))
}

// Id returns a reference to field id of azurerm_virtual_machine_scale_set_extension.
func (vmsse virtualMachineScaleSetExtensionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vmsse.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_virtual_machine_scale_set_extension.
func (vmsse virtualMachineScaleSetExtensionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vmsse.ref.Append("name"))
}

// ProtectedSettings returns a reference to field protected_settings of azurerm_virtual_machine_scale_set_extension.
func (vmsse virtualMachineScaleSetExtensionAttributes) ProtectedSettings() terra.StringValue {
	return terra.ReferenceAsString(vmsse.ref.Append("protected_settings"))
}

// ProvisionAfterExtensions returns a reference to field provision_after_extensions of azurerm_virtual_machine_scale_set_extension.
func (vmsse virtualMachineScaleSetExtensionAttributes) ProvisionAfterExtensions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vmsse.ref.Append("provision_after_extensions"))
}

// Publisher returns a reference to field publisher of azurerm_virtual_machine_scale_set_extension.
func (vmsse virtualMachineScaleSetExtensionAttributes) Publisher() terra.StringValue {
	return terra.ReferenceAsString(vmsse.ref.Append("publisher"))
}

// Settings returns a reference to field settings of azurerm_virtual_machine_scale_set_extension.
func (vmsse virtualMachineScaleSetExtensionAttributes) Settings() terra.StringValue {
	return terra.ReferenceAsString(vmsse.ref.Append("settings"))
}

// Type returns a reference to field type of azurerm_virtual_machine_scale_set_extension.
func (vmsse virtualMachineScaleSetExtensionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vmsse.ref.Append("type"))
}

// TypeHandlerVersion returns a reference to field type_handler_version of azurerm_virtual_machine_scale_set_extension.
func (vmsse virtualMachineScaleSetExtensionAttributes) TypeHandlerVersion() terra.StringValue {
	return terra.ReferenceAsString(vmsse.ref.Append("type_handler_version"))
}

// VirtualMachineScaleSetId returns a reference to field virtual_machine_scale_set_id of azurerm_virtual_machine_scale_set_extension.
func (vmsse virtualMachineScaleSetExtensionAttributes) VirtualMachineScaleSetId() terra.StringValue {
	return terra.ReferenceAsString(vmsse.ref.Append("virtual_machine_scale_set_id"))
}

func (vmsse virtualMachineScaleSetExtensionAttributes) ProtectedSettingsFromKeyVault() terra.ListValue[virtualmachinescalesetextension.ProtectedSettingsFromKeyVaultAttributes] {
	return terra.ReferenceAsList[virtualmachinescalesetextension.ProtectedSettingsFromKeyVaultAttributes](vmsse.ref.Append("protected_settings_from_key_vault"))
}

func (vmsse virtualMachineScaleSetExtensionAttributes) Timeouts() virtualmachinescalesetextension.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualmachinescalesetextension.TimeoutsAttributes](vmsse.ref.Append("timeouts"))
}

type virtualMachineScaleSetExtensionState struct {
	AutoUpgradeMinorVersion       bool                                                                 `json:"auto_upgrade_minor_version"`
	AutomaticUpgradeEnabled       bool                                                                 `json:"automatic_upgrade_enabled"`
	FailureSuppressionEnabled     bool                                                                 `json:"failure_suppression_enabled"`
	ForceUpdateTag                string                                                               `json:"force_update_tag"`
	Id                            string                                                               `json:"id"`
	Name                          string                                                               `json:"name"`
	ProtectedSettings             string                                                               `json:"protected_settings"`
	ProvisionAfterExtensions      []string                                                             `json:"provision_after_extensions"`
	Publisher                     string                                                               `json:"publisher"`
	Settings                      string                                                               `json:"settings"`
	Type                          string                                                               `json:"type"`
	TypeHandlerVersion            string                                                               `json:"type_handler_version"`
	VirtualMachineScaleSetId      string                                                               `json:"virtual_machine_scale_set_id"`
	ProtectedSettingsFromKeyVault []virtualmachinescalesetextension.ProtectedSettingsFromKeyVaultState `json:"protected_settings_from_key_vault"`
	Timeouts                      *virtualmachinescalesetextension.TimeoutsState                       `json:"timeouts"`
}
