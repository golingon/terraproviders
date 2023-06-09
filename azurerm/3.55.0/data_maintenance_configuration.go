// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamaintenanceconfiguration "github.com/golingon/terraproviders/azurerm/3.55.0/datamaintenanceconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMaintenanceConfiguration creates a new instance of [DataMaintenanceConfiguration].
func NewDataMaintenanceConfiguration(name string, args DataMaintenanceConfigurationArgs) *DataMaintenanceConfiguration {
	return &DataMaintenanceConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMaintenanceConfiguration)(nil)

// DataMaintenanceConfiguration represents the Terraform data resource azurerm_maintenance_configuration.
type DataMaintenanceConfiguration struct {
	Name string
	Args DataMaintenanceConfigurationArgs
}

// DataSource returns the Terraform object type for [DataMaintenanceConfiguration].
func (mc *DataMaintenanceConfiguration) DataSource() string {
	return "azurerm_maintenance_configuration"
}

// LocalName returns the local name for [DataMaintenanceConfiguration].
func (mc *DataMaintenanceConfiguration) LocalName() string {
	return mc.Name
}

// Configuration returns the configuration (args) for [DataMaintenanceConfiguration].
func (mc *DataMaintenanceConfiguration) Configuration() interface{} {
	return mc.Args
}

// Attributes returns the attributes for [DataMaintenanceConfiguration].
func (mc *DataMaintenanceConfiguration) Attributes() dataMaintenanceConfigurationAttributes {
	return dataMaintenanceConfigurationAttributes{ref: terra.ReferenceDataResource(mc)}
}

// DataMaintenanceConfigurationArgs contains the configurations for azurerm_maintenance_configuration.
type DataMaintenanceConfigurationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// InstallPatches: min=0
	InstallPatches []datamaintenanceconfiguration.InstallPatches `hcl:"install_patches,block" validate:"min=0"`
	// Window: min=0
	Window []datamaintenanceconfiguration.Window `hcl:"window,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datamaintenanceconfiguration.Timeouts `hcl:"timeouts,block"`
}
type dataMaintenanceConfigurationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_maintenance_configuration.
func (mc dataMaintenanceConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("id"))
}

// InGuestUserPatchMode returns a reference to field in_guest_user_patch_mode of azurerm_maintenance_configuration.
func (mc dataMaintenanceConfigurationAttributes) InGuestUserPatchMode() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("in_guest_user_patch_mode"))
}

// Location returns a reference to field location of azurerm_maintenance_configuration.
func (mc dataMaintenanceConfigurationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_maintenance_configuration.
func (mc dataMaintenanceConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("name"))
}

// Properties returns a reference to field properties of azurerm_maintenance_configuration.
func (mc dataMaintenanceConfigurationAttributes) Properties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mc.ref.Append("properties"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_maintenance_configuration.
func (mc dataMaintenanceConfigurationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("resource_group_name"))
}

// Scope returns a reference to field scope of azurerm_maintenance_configuration.
func (mc dataMaintenanceConfigurationAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("scope"))
}

// Tags returns a reference to field tags of azurerm_maintenance_configuration.
func (mc dataMaintenanceConfigurationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mc.ref.Append("tags"))
}

// Visibility returns a reference to field visibility of azurerm_maintenance_configuration.
func (mc dataMaintenanceConfigurationAttributes) Visibility() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("visibility"))
}

func (mc dataMaintenanceConfigurationAttributes) InstallPatches() terra.ListValue[datamaintenanceconfiguration.InstallPatchesAttributes] {
	return terra.ReferenceAsList[datamaintenanceconfiguration.InstallPatchesAttributes](mc.ref.Append("install_patches"))
}

func (mc dataMaintenanceConfigurationAttributes) Window() terra.ListValue[datamaintenanceconfiguration.WindowAttributes] {
	return terra.ReferenceAsList[datamaintenanceconfiguration.WindowAttributes](mc.ref.Append("window"))
}

func (mc dataMaintenanceConfigurationAttributes) Timeouts() datamaintenanceconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamaintenanceconfiguration.TimeoutsAttributes](mc.ref.Append("timeouts"))
}
