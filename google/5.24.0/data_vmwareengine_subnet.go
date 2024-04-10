// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"github.com/golingon/lingon/pkg/terra"
	datavmwareenginesubnet "github.com/golingon/terraproviders/google/5.24.0/datavmwareenginesubnet"
)

// NewDataVmwareengineSubnet creates a new instance of [DataVmwareengineSubnet].
func NewDataVmwareengineSubnet(name string, args DataVmwareengineSubnetArgs) *DataVmwareengineSubnet {
	return &DataVmwareengineSubnet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVmwareengineSubnet)(nil)

// DataVmwareengineSubnet represents the Terraform data resource google_vmwareengine_subnet.
type DataVmwareengineSubnet struct {
	Name string
	Args DataVmwareengineSubnetArgs
}

// DataSource returns the Terraform object type for [DataVmwareengineSubnet].
func (vs *DataVmwareengineSubnet) DataSource() string {
	return "google_vmwareengine_subnet"
}

// LocalName returns the local name for [DataVmwareengineSubnet].
func (vs *DataVmwareengineSubnet) LocalName() string {
	return vs.Name
}

// Configuration returns the configuration (args) for [DataVmwareengineSubnet].
func (vs *DataVmwareengineSubnet) Configuration() interface{} {
	return vs.Args
}

// Attributes returns the attributes for [DataVmwareengineSubnet].
func (vs *DataVmwareengineSubnet) Attributes() dataVmwareengineSubnetAttributes {
	return dataVmwareengineSubnetAttributes{ref: terra.ReferenceDataResource(vs)}
}

// DataVmwareengineSubnetArgs contains the configurations for google_vmwareengine_subnet.
type DataVmwareengineSubnetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// DhcpAddressRanges: min=0
	DhcpAddressRanges []datavmwareenginesubnet.DhcpAddressRanges `hcl:"dhcp_address_ranges,block" validate:"min=0"`
}
type dataVmwareengineSubnetAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_vmwareengine_subnet.
func (vs dataVmwareengineSubnetAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("create_time"))
}

// GatewayId returns a reference to field gateway_id of google_vmwareengine_subnet.
func (vs dataVmwareengineSubnetAttributes) GatewayId() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("gateway_id"))
}

// GatewayIp returns a reference to field gateway_ip of google_vmwareengine_subnet.
func (vs dataVmwareengineSubnetAttributes) GatewayIp() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("gateway_ip"))
}

// Id returns a reference to field id of google_vmwareengine_subnet.
func (vs dataVmwareengineSubnetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("id"))
}

// IpCidrRange returns a reference to field ip_cidr_range of google_vmwareengine_subnet.
func (vs dataVmwareengineSubnetAttributes) IpCidrRange() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("ip_cidr_range"))
}

// Name returns a reference to field name of google_vmwareengine_subnet.
func (vs dataVmwareengineSubnetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("name"))
}

// Parent returns a reference to field parent of google_vmwareengine_subnet.
func (vs dataVmwareengineSubnetAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("parent"))
}

// StandardConfig returns a reference to field standard_config of google_vmwareengine_subnet.
func (vs dataVmwareengineSubnetAttributes) StandardConfig() terra.BoolValue {
	return terra.ReferenceAsBool(vs.ref.Append("standard_config"))
}

// State returns a reference to field state of google_vmwareengine_subnet.
func (vs dataVmwareengineSubnetAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("state"))
}

// Type returns a reference to field type of google_vmwareengine_subnet.
func (vs dataVmwareengineSubnetAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("type"))
}

// Uid returns a reference to field uid of google_vmwareengine_subnet.
func (vs dataVmwareengineSubnetAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_vmwareengine_subnet.
func (vs dataVmwareengineSubnetAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("update_time"))
}

// VlanId returns a reference to field vlan_id of google_vmwareengine_subnet.
func (vs dataVmwareengineSubnetAttributes) VlanId() terra.NumberValue {
	return terra.ReferenceAsNumber(vs.ref.Append("vlan_id"))
}

func (vs dataVmwareengineSubnetAttributes) DhcpAddressRanges() terra.ListValue[datavmwareenginesubnet.DhcpAddressRangesAttributes] {
	return terra.ReferenceAsList[datavmwareenginesubnet.DhcpAddressRangesAttributes](vs.ref.Append("dhcp_address_ranges"))
}
