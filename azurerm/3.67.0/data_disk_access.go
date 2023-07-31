// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadiskaccess "github.com/golingon/terraproviders/azurerm/3.67.0/datadiskaccess"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDiskAccess creates a new instance of [DataDiskAccess].
func NewDataDiskAccess(name string, args DataDiskAccessArgs) *DataDiskAccess {
	return &DataDiskAccess{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDiskAccess)(nil)

// DataDiskAccess represents the Terraform data resource azurerm_disk_access.
type DataDiskAccess struct {
	Name string
	Args DataDiskAccessArgs
}

// DataSource returns the Terraform object type for [DataDiskAccess].
func (da *DataDiskAccess) DataSource() string {
	return "azurerm_disk_access"
}

// LocalName returns the local name for [DataDiskAccess].
func (da *DataDiskAccess) LocalName() string {
	return da.Name
}

// Configuration returns the configuration (args) for [DataDiskAccess].
func (da *DataDiskAccess) Configuration() interface{} {
	return da.Args
}

// Attributes returns the attributes for [DataDiskAccess].
func (da *DataDiskAccess) Attributes() dataDiskAccessAttributes {
	return dataDiskAccessAttributes{ref: terra.ReferenceDataResource(da)}
}

// DataDiskAccessArgs contains the configurations for azurerm_disk_access.
type DataDiskAccessArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datadiskaccess.Timeouts `hcl:"timeouts,block"`
}
type dataDiskAccessAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_disk_access.
func (da dataDiskAccessAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_disk_access.
func (da dataDiskAccessAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_disk_access.
func (da dataDiskAccessAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_disk_access.
func (da dataDiskAccessAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](da.ref.Append("tags"))
}

func (da dataDiskAccessAttributes) Timeouts() datadiskaccess.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadiskaccess.TimeoutsAttributes](da.ref.Append("timeouts"))
}
