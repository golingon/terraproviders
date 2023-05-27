// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datacomputehavpngateway "github.com/golingon/terraproviders/google/4.66.0/datacomputehavpngateway"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeHaVpnGateway creates a new instance of [DataComputeHaVpnGateway].
func NewDataComputeHaVpnGateway(name string, args DataComputeHaVpnGatewayArgs) *DataComputeHaVpnGateway {
	return &DataComputeHaVpnGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeHaVpnGateway)(nil)

// DataComputeHaVpnGateway represents the Terraform data resource google_compute_ha_vpn_gateway.
type DataComputeHaVpnGateway struct {
	Name string
	Args DataComputeHaVpnGatewayArgs
}

// DataSource returns the Terraform object type for [DataComputeHaVpnGateway].
func (chvg *DataComputeHaVpnGateway) DataSource() string {
	return "google_compute_ha_vpn_gateway"
}

// LocalName returns the local name for [DataComputeHaVpnGateway].
func (chvg *DataComputeHaVpnGateway) LocalName() string {
	return chvg.Name
}

// Configuration returns the configuration (args) for [DataComputeHaVpnGateway].
func (chvg *DataComputeHaVpnGateway) Configuration() interface{} {
	return chvg.Args
}

// Attributes returns the attributes for [DataComputeHaVpnGateway].
func (chvg *DataComputeHaVpnGateway) Attributes() dataComputeHaVpnGatewayAttributes {
	return dataComputeHaVpnGatewayAttributes{ref: terra.ReferenceDataResource(chvg)}
}

// DataComputeHaVpnGatewayArgs contains the configurations for google_compute_ha_vpn_gateway.
type DataComputeHaVpnGatewayArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// VpnInterfaces: min=0
	VpnInterfaces []datacomputehavpngateway.VpnInterfaces `hcl:"vpn_interfaces,block" validate:"min=0"`
}
type dataComputeHaVpnGatewayAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_ha_vpn_gateway.
func (chvg dataComputeHaVpnGatewayAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(chvg.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_ha_vpn_gateway.
func (chvg dataComputeHaVpnGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(chvg.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_ha_vpn_gateway.
func (chvg dataComputeHaVpnGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(chvg.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_ha_vpn_gateway.
func (chvg dataComputeHaVpnGatewayAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(chvg.ref.Append("network"))
}

// Project returns a reference to field project of google_compute_ha_vpn_gateway.
func (chvg dataComputeHaVpnGatewayAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(chvg.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_ha_vpn_gateway.
func (chvg dataComputeHaVpnGatewayAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(chvg.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_ha_vpn_gateway.
func (chvg dataComputeHaVpnGatewayAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(chvg.ref.Append("self_link"))
}

// StackType returns a reference to field stack_type of google_compute_ha_vpn_gateway.
func (chvg dataComputeHaVpnGatewayAttributes) StackType() terra.StringValue {
	return terra.ReferenceAsString(chvg.ref.Append("stack_type"))
}

func (chvg dataComputeHaVpnGatewayAttributes) VpnInterfaces() terra.ListValue[datacomputehavpngateway.VpnInterfacesAttributes] {
	return terra.ReferenceAsList[datacomputehavpngateway.VpnInterfacesAttributes](chvg.ref.Append("vpn_interfaces"))
}
