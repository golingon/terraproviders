// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadedicatedhost "github.com/golingon/terraproviders/azurerm/3.58.0/datadedicatedhost"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDedicatedHost creates a new instance of [DataDedicatedHost].
func NewDataDedicatedHost(name string, args DataDedicatedHostArgs) *DataDedicatedHost {
	return &DataDedicatedHost{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDedicatedHost)(nil)

// DataDedicatedHost represents the Terraform data resource azurerm_dedicated_host.
type DataDedicatedHost struct {
	Name string
	Args DataDedicatedHostArgs
}

// DataSource returns the Terraform object type for [DataDedicatedHost].
func (dh *DataDedicatedHost) DataSource() string {
	return "azurerm_dedicated_host"
}

// LocalName returns the local name for [DataDedicatedHost].
func (dh *DataDedicatedHost) LocalName() string {
	return dh.Name
}

// Configuration returns the configuration (args) for [DataDedicatedHost].
func (dh *DataDedicatedHost) Configuration() interface{} {
	return dh.Args
}

// Attributes returns the attributes for [DataDedicatedHost].
func (dh *DataDedicatedHost) Attributes() dataDedicatedHostAttributes {
	return dataDedicatedHostAttributes{ref: terra.ReferenceDataResource(dh)}
}

// DataDedicatedHostArgs contains the configurations for azurerm_dedicated_host.
type DataDedicatedHostArgs struct {
	// DedicatedHostGroupName: string, required
	DedicatedHostGroupName terra.StringValue `hcl:"dedicated_host_group_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datadedicatedhost.Timeouts `hcl:"timeouts,block"`
}
type dataDedicatedHostAttributes struct {
	ref terra.Reference
}

// DedicatedHostGroupName returns a reference to field dedicated_host_group_name of azurerm_dedicated_host.
func (dh dataDedicatedHostAttributes) DedicatedHostGroupName() terra.StringValue {
	return terra.ReferenceAsString(dh.ref.Append("dedicated_host_group_name"))
}

// Id returns a reference to field id of azurerm_dedicated_host.
func (dh dataDedicatedHostAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dh.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_dedicated_host.
func (dh dataDedicatedHostAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dh.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_dedicated_host.
func (dh dataDedicatedHostAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dh.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dedicated_host.
func (dh dataDedicatedHostAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dh.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dedicated_host.
func (dh dataDedicatedHostAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dh.ref.Append("tags"))
}

func (dh dataDedicatedHostAttributes) Timeouts() datadedicatedhost.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadedicatedhost.TimeoutsAttributes](dh.ref.Append("timeouts"))
}
