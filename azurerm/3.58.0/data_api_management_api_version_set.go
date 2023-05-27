// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataapimanagementapiversionset "github.com/golingon/terraproviders/azurerm/3.58.0/dataapimanagementapiversionset"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataApiManagementApiVersionSet creates a new instance of [DataApiManagementApiVersionSet].
func NewDataApiManagementApiVersionSet(name string, args DataApiManagementApiVersionSetArgs) *DataApiManagementApiVersionSet {
	return &DataApiManagementApiVersionSet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApiManagementApiVersionSet)(nil)

// DataApiManagementApiVersionSet represents the Terraform data resource azurerm_api_management_api_version_set.
type DataApiManagementApiVersionSet struct {
	Name string
	Args DataApiManagementApiVersionSetArgs
}

// DataSource returns the Terraform object type for [DataApiManagementApiVersionSet].
func (amavs *DataApiManagementApiVersionSet) DataSource() string {
	return "azurerm_api_management_api_version_set"
}

// LocalName returns the local name for [DataApiManagementApiVersionSet].
func (amavs *DataApiManagementApiVersionSet) LocalName() string {
	return amavs.Name
}

// Configuration returns the configuration (args) for [DataApiManagementApiVersionSet].
func (amavs *DataApiManagementApiVersionSet) Configuration() interface{} {
	return amavs.Args
}

// Attributes returns the attributes for [DataApiManagementApiVersionSet].
func (amavs *DataApiManagementApiVersionSet) Attributes() dataApiManagementApiVersionSetAttributes {
	return dataApiManagementApiVersionSetAttributes{ref: terra.ReferenceDataResource(amavs)}
}

// DataApiManagementApiVersionSetArgs contains the configurations for azurerm_api_management_api_version_set.
type DataApiManagementApiVersionSetArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataapimanagementapiversionset.Timeouts `hcl:"timeouts,block"`
}
type dataApiManagementApiVersionSetAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_api_version_set.
func (amavs dataApiManagementApiVersionSetAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("api_management_name"))
}

// Description returns a reference to field description of azurerm_api_management_api_version_set.
func (amavs dataApiManagementApiVersionSetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_api_version_set.
func (amavs dataApiManagementApiVersionSetAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_api_management_api_version_set.
func (amavs dataApiManagementApiVersionSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_api_management_api_version_set.
func (amavs dataApiManagementApiVersionSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_api_version_set.
func (amavs dataApiManagementApiVersionSetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("resource_group_name"))
}

// VersionHeaderName returns a reference to field version_header_name of azurerm_api_management_api_version_set.
func (amavs dataApiManagementApiVersionSetAttributes) VersionHeaderName() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("version_header_name"))
}

// VersionQueryName returns a reference to field version_query_name of azurerm_api_management_api_version_set.
func (amavs dataApiManagementApiVersionSetAttributes) VersionQueryName() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("version_query_name"))
}

// VersioningScheme returns a reference to field versioning_scheme of azurerm_api_management_api_version_set.
func (amavs dataApiManagementApiVersionSetAttributes) VersioningScheme() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("versioning_scheme"))
}

func (amavs dataApiManagementApiVersionSetAttributes) Timeouts() dataapimanagementapiversionset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataapimanagementapiversionset.TimeoutsAttributes](amavs.ref.Append("timeouts"))
}
