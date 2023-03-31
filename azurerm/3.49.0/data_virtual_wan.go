// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datavirtualwan "github.com/golingon/terraproviders/azurerm/3.49.0/datavirtualwan"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataVirtualWan(name string, args DataVirtualWanArgs) *DataVirtualWan {
	return &DataVirtualWan{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVirtualWan)(nil)

type DataVirtualWan struct {
	Name string
	Args DataVirtualWanArgs
}

func (vw *DataVirtualWan) DataSource() string {
	return "azurerm_virtual_wan"
}

func (vw *DataVirtualWan) LocalName() string {
	return vw.Name
}

func (vw *DataVirtualWan) Configuration() interface{} {
	return vw.Args
}

func (vw *DataVirtualWan) Attributes() dataVirtualWanAttributes {
	return dataVirtualWanAttributes{ref: terra.ReferenceDataResource(vw)}
}

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

func (vw dataVirtualWanAttributes) AllowBranchToBranchTraffic() terra.BoolValue {
	return terra.ReferenceBool(vw.ref.Append("allow_branch_to_branch_traffic"))
}

func (vw dataVirtualWanAttributes) DisableVpnEncryption() terra.BoolValue {
	return terra.ReferenceBool(vw.ref.Append("disable_vpn_encryption"))
}

func (vw dataVirtualWanAttributes) Id() terra.StringValue {
	return terra.ReferenceString(vw.ref.Append("id"))
}

func (vw dataVirtualWanAttributes) Location() terra.StringValue {
	return terra.ReferenceString(vw.ref.Append("location"))
}

func (vw dataVirtualWanAttributes) Name() terra.StringValue {
	return terra.ReferenceString(vw.ref.Append("name"))
}

func (vw dataVirtualWanAttributes) Office365LocalBreakoutCategory() terra.StringValue {
	return terra.ReferenceString(vw.ref.Append("office365_local_breakout_category"))
}

func (vw dataVirtualWanAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(vw.ref.Append("resource_group_name"))
}

func (vw dataVirtualWanAttributes) Sku() terra.StringValue {
	return terra.ReferenceString(vw.ref.Append("sku"))
}

func (vw dataVirtualWanAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](vw.ref.Append("tags"))
}

func (vw dataVirtualWanAttributes) VirtualHubIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](vw.ref.Append("virtual_hub_ids"))
}

func (vw dataVirtualWanAttributes) VpnSiteIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](vw.ref.Append("vpn_site_ids"))
}

func (vw dataVirtualWanAttributes) Timeouts() datavirtualwan.TimeoutsAttributes {
	return terra.ReferenceSingle[datavirtualwan.TimeoutsAttributes](vw.ref.Append("timeouts"))
}
