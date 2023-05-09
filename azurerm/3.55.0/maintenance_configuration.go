// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	maintenanceconfiguration "github.com/golingon/terraproviders/azurerm/3.55.0/maintenanceconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMaintenanceConfiguration creates a new instance of [MaintenanceConfiguration].
func NewMaintenanceConfiguration(name string, args MaintenanceConfigurationArgs) *MaintenanceConfiguration {
	return &MaintenanceConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MaintenanceConfiguration)(nil)

// MaintenanceConfiguration represents the Terraform resource azurerm_maintenance_configuration.
type MaintenanceConfiguration struct {
	Name      string
	Args      MaintenanceConfigurationArgs
	state     *maintenanceConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MaintenanceConfiguration].
func (mc *MaintenanceConfiguration) Type() string {
	return "azurerm_maintenance_configuration"
}

// LocalName returns the local name for [MaintenanceConfiguration].
func (mc *MaintenanceConfiguration) LocalName() string {
	return mc.Name
}

// Configuration returns the configuration (args) for [MaintenanceConfiguration].
func (mc *MaintenanceConfiguration) Configuration() interface{} {
	return mc.Args
}

// DependOn is used for other resources to depend on [MaintenanceConfiguration].
func (mc *MaintenanceConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(mc)
}

// Dependencies returns the list of resources [MaintenanceConfiguration] depends_on.
func (mc *MaintenanceConfiguration) Dependencies() terra.Dependencies {
	return mc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MaintenanceConfiguration].
func (mc *MaintenanceConfiguration) LifecycleManagement() *terra.Lifecycle {
	return mc.Lifecycle
}

// Attributes returns the attributes for [MaintenanceConfiguration].
func (mc *MaintenanceConfiguration) Attributes() maintenanceConfigurationAttributes {
	return maintenanceConfigurationAttributes{ref: terra.ReferenceResource(mc)}
}

// ImportState imports the given attribute values into [MaintenanceConfiguration]'s state.
func (mc *MaintenanceConfiguration) ImportState(av io.Reader) error {
	mc.state = &maintenanceConfigurationState{}
	if err := json.NewDecoder(av).Decode(mc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mc.Type(), mc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MaintenanceConfiguration] has state.
func (mc *MaintenanceConfiguration) State() (*maintenanceConfigurationState, bool) {
	return mc.state, mc.state != nil
}

// StateMust returns the state for [MaintenanceConfiguration]. Panics if the state is nil.
func (mc *MaintenanceConfiguration) StateMust() *maintenanceConfigurationState {
	if mc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mc.Type(), mc.LocalName()))
	}
	return mc.state
}

// MaintenanceConfigurationArgs contains the configurations for azurerm_maintenance_configuration.
type MaintenanceConfigurationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InGuestUserPatchMode: string, optional
	InGuestUserPatchMode terra.StringValue `hcl:"in_guest_user_patch_mode,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Properties: map of string, optional
	Properties terra.MapValue[terra.StringValue] `hcl:"properties,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Visibility: string, optional
	Visibility terra.StringValue `hcl:"visibility,attr"`
	// InstallPatches: optional
	InstallPatches *maintenanceconfiguration.InstallPatches `hcl:"install_patches,block"`
	// Timeouts: optional
	Timeouts *maintenanceconfiguration.Timeouts `hcl:"timeouts,block"`
	// Window: optional
	Window *maintenanceconfiguration.Window `hcl:"window,block"`
}
type maintenanceConfigurationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_maintenance_configuration.
func (mc maintenanceConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("id"))
}

// InGuestUserPatchMode returns a reference to field in_guest_user_patch_mode of azurerm_maintenance_configuration.
func (mc maintenanceConfigurationAttributes) InGuestUserPatchMode() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("in_guest_user_patch_mode"))
}

// Location returns a reference to field location of azurerm_maintenance_configuration.
func (mc maintenanceConfigurationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_maintenance_configuration.
func (mc maintenanceConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("name"))
}

// Properties returns a reference to field properties of azurerm_maintenance_configuration.
func (mc maintenanceConfigurationAttributes) Properties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mc.ref.Append("properties"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_maintenance_configuration.
func (mc maintenanceConfigurationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("resource_group_name"))
}

// Scope returns a reference to field scope of azurerm_maintenance_configuration.
func (mc maintenanceConfigurationAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("scope"))
}

// Tags returns a reference to field tags of azurerm_maintenance_configuration.
func (mc maintenanceConfigurationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mc.ref.Append("tags"))
}

// Visibility returns a reference to field visibility of azurerm_maintenance_configuration.
func (mc maintenanceConfigurationAttributes) Visibility() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("visibility"))
}

func (mc maintenanceConfigurationAttributes) InstallPatches() terra.ListValue[maintenanceconfiguration.InstallPatchesAttributes] {
	return terra.ReferenceAsList[maintenanceconfiguration.InstallPatchesAttributes](mc.ref.Append("install_patches"))
}

func (mc maintenanceConfigurationAttributes) Timeouts() maintenanceconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[maintenanceconfiguration.TimeoutsAttributes](mc.ref.Append("timeouts"))
}

func (mc maintenanceConfigurationAttributes) Window() terra.ListValue[maintenanceconfiguration.WindowAttributes] {
	return terra.ReferenceAsList[maintenanceconfiguration.WindowAttributes](mc.ref.Append("window"))
}

type maintenanceConfigurationState struct {
	Id                   string                                         `json:"id"`
	InGuestUserPatchMode string                                         `json:"in_guest_user_patch_mode"`
	Location             string                                         `json:"location"`
	Name                 string                                         `json:"name"`
	Properties           map[string]string                              `json:"properties"`
	ResourceGroupName    string                                         `json:"resource_group_name"`
	Scope                string                                         `json:"scope"`
	Tags                 map[string]string                              `json:"tags"`
	Visibility           string                                         `json:"visibility"`
	InstallPatches       []maintenanceconfiguration.InstallPatchesState `json:"install_patches"`
	Timeouts             *maintenanceconfiguration.TimeoutsState        `json:"timeouts"`
	Window               []maintenanceconfiguration.WindowState         `json:"window"`
}
