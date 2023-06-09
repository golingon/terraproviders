// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataComputeVpnGateway creates a new instance of [DataComputeVpnGateway].
func NewDataComputeVpnGateway(name string, args DataComputeVpnGatewayArgs) *DataComputeVpnGateway {
	return &DataComputeVpnGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeVpnGateway)(nil)

// DataComputeVpnGateway represents the Terraform data resource google_compute_vpn_gateway.
type DataComputeVpnGateway struct {
	Name string
	Args DataComputeVpnGatewayArgs
}

// DataSource returns the Terraform object type for [DataComputeVpnGateway].
func (cvg *DataComputeVpnGateway) DataSource() string {
	return "google_compute_vpn_gateway"
}

// LocalName returns the local name for [DataComputeVpnGateway].
func (cvg *DataComputeVpnGateway) LocalName() string {
	return cvg.Name
}

// Configuration returns the configuration (args) for [DataComputeVpnGateway].
func (cvg *DataComputeVpnGateway) Configuration() interface{} {
	return cvg.Args
}

// Attributes returns the attributes for [DataComputeVpnGateway].
func (cvg *DataComputeVpnGateway) Attributes() dataComputeVpnGatewayAttributes {
	return dataComputeVpnGatewayAttributes{ref: terra.ReferenceDataResource(cvg)}
}

// DataComputeVpnGatewayArgs contains the configurations for google_compute_vpn_gateway.
type DataComputeVpnGatewayArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}
type dataComputeVpnGatewayAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_vpn_gateway.
func (cvg dataComputeVpnGatewayAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cvg.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_vpn_gateway.
func (cvg dataComputeVpnGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cvg.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_vpn_gateway.
func (cvg dataComputeVpnGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cvg.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_vpn_gateway.
func (cvg dataComputeVpnGatewayAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cvg.ref.Append("network"))
}

// Project returns a reference to field project of google_compute_vpn_gateway.
func (cvg dataComputeVpnGatewayAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cvg.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_vpn_gateway.
func (cvg dataComputeVpnGatewayAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cvg.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_vpn_gateway.
func (cvg dataComputeVpnGatewayAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cvg.ref.Append("self_link"))
}
