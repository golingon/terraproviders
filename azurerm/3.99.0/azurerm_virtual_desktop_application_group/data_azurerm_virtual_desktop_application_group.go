// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_virtual_desktop_application_group

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_virtual_desktop_application_group.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (avdag *DataSource) DataSource() string {
	return "azurerm_virtual_desktop_application_group"
}

// LocalName returns the local name for [DataSource].
func (avdag *DataSource) LocalName() string {
	return avdag.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (avdag *DataSource) Configuration() interface{} {
	return avdag.Args
}

// Attributes returns the attributes for [DataSource].
func (avdag *DataSource) Attributes() dataAzurermVirtualDesktopApplicationGroupAttributes {
	return dataAzurermVirtualDesktopApplicationGroupAttributes{ref: terra.ReferenceDataSource(avdag)}
}

// DataArgs contains the configurations for azurerm_virtual_desktop_application_group.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermVirtualDesktopApplicationGroupAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_virtual_desktop_application_group.
func (avdag dataAzurermVirtualDesktopApplicationGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(avdag.ref.Append("description"))
}

// FriendlyName returns a reference to field friendly_name of azurerm_virtual_desktop_application_group.
func (avdag dataAzurermVirtualDesktopApplicationGroupAttributes) FriendlyName() terra.StringValue {
	return terra.ReferenceAsString(avdag.ref.Append("friendly_name"))
}

// HostPoolId returns a reference to field host_pool_id of azurerm_virtual_desktop_application_group.
func (avdag dataAzurermVirtualDesktopApplicationGroupAttributes) HostPoolId() terra.StringValue {
	return terra.ReferenceAsString(avdag.ref.Append("host_pool_id"))
}

// Id returns a reference to field id of azurerm_virtual_desktop_application_group.
func (avdag dataAzurermVirtualDesktopApplicationGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avdag.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_virtual_desktop_application_group.
func (avdag dataAzurermVirtualDesktopApplicationGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(avdag.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_virtual_desktop_application_group.
func (avdag dataAzurermVirtualDesktopApplicationGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avdag.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_desktop_application_group.
func (avdag dataAzurermVirtualDesktopApplicationGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(avdag.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_virtual_desktop_application_group.
func (avdag dataAzurermVirtualDesktopApplicationGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avdag.ref.Append("tags"))
}

// Type returns a reference to field type of azurerm_virtual_desktop_application_group.
func (avdag dataAzurermVirtualDesktopApplicationGroupAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(avdag.ref.Append("type"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_virtual_desktop_application_group.
func (avdag dataAzurermVirtualDesktopApplicationGroupAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(avdag.ref.Append("workspace_id"))
}

func (avdag dataAzurermVirtualDesktopApplicationGroupAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](avdag.ref.Append("timeouts"))
}
