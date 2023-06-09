// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datavirtualwan "github.com/golingon/terraproviders/azurerm/3.49.0/datavirtualwan"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVirtualWan creates a new instance of [DataVirtualWan].
func NewDataVirtualWan(name string, args DataVirtualWanArgs) *DataVirtualWan {
	return &DataVirtualWan{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVirtualWan)(nil)

// DataVirtualWan represents the Terraform data resource azurerm_virtual_wan.
type DataVirtualWan struct {
	Name string
	Args DataVirtualWanArgs
}

// DataSource returns the Terraform object type for [DataVirtualWan].
func (vw *DataVirtualWan) DataSource() string {
	return "azurerm_virtual_wan"
}

// LocalName returns the local name for [DataVirtualWan].
func (vw *DataVirtualWan) LocalName() string {
	return vw.Name
}

// Configuration returns the configuration (args) for [DataVirtualWan].
func (vw *DataVirtualWan) Configuration() interface{} {
	return vw.Args
}

// Attributes returns the attributes for [DataVirtualWan].
func (vw *DataVirtualWan) Attributes() dataVirtualWanAttributes {
	return dataVirtualWanAttributes{ref: terra.ReferenceDataResource(vw)}
}

// DataVirtualWanArgs contains the configurations for azurerm_virtual_wan.
type DataVirtualWanArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datavirtualwan.Timeouts `hcl:"timeouts,block"`
}
type dataVirtualWanAttributes struct {
	ref terra.Reference
}

// AllowBranchToBranchTraffic returns a reference to field allow_branch_to_branch_traffic of azurerm_virtual_wan.
func (vw dataVirtualWanAttributes) AllowBranchToBranchTraffic() terra.BoolValue {
	return terra.ReferenceAsBool(vw.ref.Append("allow_branch_to_branch_traffic"))
}

// DisableVpnEncryption returns a reference to field disable_vpn_encryption of azurerm_virtual_wan.
func (vw dataVirtualWanAttributes) DisableVpnEncryption() terra.BoolValue {
	return terra.ReferenceAsBool(vw.ref.Append("disable_vpn_encryption"))
}

// Id returns a reference to field id of azurerm_virtual_wan.
func (vw dataVirtualWanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vw.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_virtual_wan.
func (vw dataVirtualWanAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vw.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_virtual_wan.
func (vw dataVirtualWanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vw.ref.Append("name"))
}

// Office365LocalBreakoutCategory returns a reference to field office365_local_breakout_category of azurerm_virtual_wan.
func (vw dataVirtualWanAttributes) Office365LocalBreakoutCategory() terra.StringValue {
	return terra.ReferenceAsString(vw.ref.Append("office365_local_breakout_category"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_wan.
func (vw dataVirtualWanAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vw.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_virtual_wan.
func (vw dataVirtualWanAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(vw.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_virtual_wan.
func (vw dataVirtualWanAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vw.ref.Append("tags"))
}

// VirtualHubIds returns a reference to field virtual_hub_ids of azurerm_virtual_wan.
func (vw dataVirtualWanAttributes) VirtualHubIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vw.ref.Append("virtual_hub_ids"))
}

// VpnSiteIds returns a reference to field vpn_site_ids of azurerm_virtual_wan.
func (vw dataVirtualWanAttributes) VpnSiteIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vw.ref.Append("vpn_site_ids"))
}

func (vw dataVirtualWanAttributes) Timeouts() datavirtualwan.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavirtualwan.TimeoutsAttributes](vw.ref.Append("timeouts"))
}
