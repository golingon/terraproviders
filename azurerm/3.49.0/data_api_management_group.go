// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataapimanagementgroup "github.com/golingon/terraproviders/azurerm/3.49.0/dataapimanagementgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataApiManagementGroup creates a new instance of [DataApiManagementGroup].
func NewDataApiManagementGroup(name string, args DataApiManagementGroupArgs) *DataApiManagementGroup {
	return &DataApiManagementGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApiManagementGroup)(nil)

// DataApiManagementGroup represents the Terraform data resource azurerm_api_management_group.
type DataApiManagementGroup struct {
	Name string
	Args DataApiManagementGroupArgs
}

// DataSource returns the Terraform object type for [DataApiManagementGroup].
func (amg *DataApiManagementGroup) DataSource() string {
	return "azurerm_api_management_group"
}

// LocalName returns the local name for [DataApiManagementGroup].
func (amg *DataApiManagementGroup) LocalName() string {
	return amg.Name
}

// Configuration returns the configuration (args) for [DataApiManagementGroup].
func (amg *DataApiManagementGroup) Configuration() interface{} {
	return amg.Args
}

// Attributes returns the attributes for [DataApiManagementGroup].
func (amg *DataApiManagementGroup) Attributes() dataApiManagementGroupAttributes {
	return dataApiManagementGroupAttributes{ref: terra.ReferenceDataResource(amg)}
}

// DataApiManagementGroupArgs contains the configurations for azurerm_api_management_group.
type DataApiManagementGroupArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataapimanagementgroup.Timeouts `hcl:"timeouts,block"`
}
type dataApiManagementGroupAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_group.
func (amg dataApiManagementGroupAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amg.ref.Append("api_management_name"))
}

// Description returns a reference to field description of azurerm_api_management_group.
func (amg dataApiManagementGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amg.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_group.
func (amg dataApiManagementGroupAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(amg.ref.Append("display_name"))
}

// ExternalId returns a reference to field external_id of azurerm_api_management_group.
func (amg dataApiManagementGroupAttributes) ExternalId() terra.StringValue {
	return terra.ReferenceAsString(amg.ref.Append("external_id"))
}

// Id returns a reference to field id of azurerm_api_management_group.
func (amg dataApiManagementGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_api_management_group.
func (amg dataApiManagementGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amg.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_group.
func (amg dataApiManagementGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amg.ref.Append("resource_group_name"))
}

// Type returns a reference to field type of azurerm_api_management_group.
func (amg dataApiManagementGroupAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(amg.ref.Append("type"))
}

func (amg dataApiManagementGroupAttributes) Timeouts() dataapimanagementgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataapimanagementgroup.TimeoutsAttributes](amg.ref.Append("timeouts"))
}
