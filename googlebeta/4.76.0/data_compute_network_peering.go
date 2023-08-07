// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datacomputenetworkpeering "github.com/golingon/terraproviders/googlebeta/4.76.0/datacomputenetworkpeering"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeNetworkPeering creates a new instance of [DataComputeNetworkPeering].
func NewDataComputeNetworkPeering(name string, args DataComputeNetworkPeeringArgs) *DataComputeNetworkPeering {
	return &DataComputeNetworkPeering{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeNetworkPeering)(nil)

// DataComputeNetworkPeering represents the Terraform data resource google_compute_network_peering.
type DataComputeNetworkPeering struct {
	Name string
	Args DataComputeNetworkPeeringArgs
}

// DataSource returns the Terraform object type for [DataComputeNetworkPeering].
func (cnp *DataComputeNetworkPeering) DataSource() string {
	return "google_compute_network_peering"
}

// LocalName returns the local name for [DataComputeNetworkPeering].
func (cnp *DataComputeNetworkPeering) LocalName() string {
	return cnp.Name
}

// Configuration returns the configuration (args) for [DataComputeNetworkPeering].
func (cnp *DataComputeNetworkPeering) Configuration() interface{} {
	return cnp.Args
}

// Attributes returns the attributes for [DataComputeNetworkPeering].
func (cnp *DataComputeNetworkPeering) Attributes() dataComputeNetworkPeeringAttributes {
	return dataComputeNetworkPeeringAttributes{ref: terra.ReferenceDataResource(cnp)}
}

// DataComputeNetworkPeeringArgs contains the configurations for google_compute_network_peering.
type DataComputeNetworkPeeringArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datacomputenetworkpeering.Timeouts `hcl:"timeouts,block"`
}
type dataComputeNetworkPeeringAttributes struct {
	ref terra.Reference
}

// ExportCustomRoutes returns a reference to field export_custom_routes of google_compute_network_peering.
func (cnp dataComputeNetworkPeeringAttributes) ExportCustomRoutes() terra.BoolValue {
	return terra.ReferenceAsBool(cnp.ref.Append("export_custom_routes"))
}

// ExportSubnetRoutesWithPublicIp returns a reference to field export_subnet_routes_with_public_ip of google_compute_network_peering.
func (cnp dataComputeNetworkPeeringAttributes) ExportSubnetRoutesWithPublicIp() terra.BoolValue {
	return terra.ReferenceAsBool(cnp.ref.Append("export_subnet_routes_with_public_ip"))
}

// Id returns a reference to field id of google_compute_network_peering.
func (cnp dataComputeNetworkPeeringAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cnp.ref.Append("id"))
}

// ImportCustomRoutes returns a reference to field import_custom_routes of google_compute_network_peering.
func (cnp dataComputeNetworkPeeringAttributes) ImportCustomRoutes() terra.BoolValue {
	return terra.ReferenceAsBool(cnp.ref.Append("import_custom_routes"))
}

// ImportSubnetRoutesWithPublicIp returns a reference to field import_subnet_routes_with_public_ip of google_compute_network_peering.
func (cnp dataComputeNetworkPeeringAttributes) ImportSubnetRoutesWithPublicIp() terra.BoolValue {
	return terra.ReferenceAsBool(cnp.ref.Append("import_subnet_routes_with_public_ip"))
}

// Name returns a reference to field name of google_compute_network_peering.
func (cnp dataComputeNetworkPeeringAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cnp.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_network_peering.
func (cnp dataComputeNetworkPeeringAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cnp.ref.Append("network"))
}

// PeerNetwork returns a reference to field peer_network of google_compute_network_peering.
func (cnp dataComputeNetworkPeeringAttributes) PeerNetwork() terra.StringValue {
	return terra.ReferenceAsString(cnp.ref.Append("peer_network"))
}

// StackType returns a reference to field stack_type of google_compute_network_peering.
func (cnp dataComputeNetworkPeeringAttributes) StackType() terra.StringValue {
	return terra.ReferenceAsString(cnp.ref.Append("stack_type"))
}

// State returns a reference to field state of google_compute_network_peering.
func (cnp dataComputeNetworkPeeringAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(cnp.ref.Append("state"))
}

// StateDetails returns a reference to field state_details of google_compute_network_peering.
func (cnp dataComputeNetworkPeeringAttributes) StateDetails() terra.StringValue {
	return terra.ReferenceAsString(cnp.ref.Append("state_details"))
}

func (cnp dataComputeNetworkPeeringAttributes) Timeouts() datacomputenetworkpeering.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacomputenetworkpeering.TimeoutsAttributes](cnp.ref.Append("timeouts"))
}
