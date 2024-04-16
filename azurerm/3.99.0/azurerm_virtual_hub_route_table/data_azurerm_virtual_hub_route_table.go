// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_virtual_hub_route_table

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_virtual_hub_route_table.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (avhrt *DataSource) DataSource() string {
	return "azurerm_virtual_hub_route_table"
}

// LocalName returns the local name for [DataSource].
func (avhrt *DataSource) LocalName() string {
	return avhrt.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (avhrt *DataSource) Configuration() interface{} {
	return avhrt.Args
}

// Attributes returns the attributes for [DataSource].
func (avhrt *DataSource) Attributes() dataAzurermVirtualHubRouteTableAttributes {
	return dataAzurermVirtualHubRouteTableAttributes{ref: terra.ReferenceDataSource(avhrt)}
}

// DataArgs contains the configurations for azurerm_virtual_hub_route_table.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// VirtualHubName: string, required
	VirtualHubName terra.StringValue `hcl:"virtual_hub_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermVirtualHubRouteTableAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_virtual_hub_route_table.
func (avhrt dataAzurermVirtualHubRouteTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avhrt.ref.Append("id"))
}

// Labels returns a reference to field labels of azurerm_virtual_hub_route_table.
func (avhrt dataAzurermVirtualHubRouteTableAttributes) Labels() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](avhrt.ref.Append("labels"))
}

// Name returns a reference to field name of azurerm_virtual_hub_route_table.
func (avhrt dataAzurermVirtualHubRouteTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avhrt.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_hub_route_table.
func (avhrt dataAzurermVirtualHubRouteTableAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(avhrt.ref.Append("resource_group_name"))
}

// VirtualHubId returns a reference to field virtual_hub_id of azurerm_virtual_hub_route_table.
func (avhrt dataAzurermVirtualHubRouteTableAttributes) VirtualHubId() terra.StringValue {
	return terra.ReferenceAsString(avhrt.ref.Append("virtual_hub_id"))
}

// VirtualHubName returns a reference to field virtual_hub_name of azurerm_virtual_hub_route_table.
func (avhrt dataAzurermVirtualHubRouteTableAttributes) VirtualHubName() terra.StringValue {
	return terra.ReferenceAsString(avhrt.ref.Append("virtual_hub_name"))
}

func (avhrt dataAzurermVirtualHubRouteTableAttributes) Route() terra.ListValue[DataRouteAttributes] {
	return terra.ReferenceAsList[DataRouteAttributes](avhrt.ref.Append("route"))
}

func (avhrt dataAzurermVirtualHubRouteTableAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](avhrt.ref.Append("timeouts"))
}
