// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_network_peering

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_compute_network_peering.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gcnp *DataSource) DataSource() string {
	return "google_compute_network_peering"
}

// LocalName returns the local name for [DataSource].
func (gcnp *DataSource) LocalName() string {
	return gcnp.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gcnp *DataSource) Configuration() interface{} {
	return gcnp.Args
}

// Attributes returns the attributes for [DataSource].
func (gcnp *DataSource) Attributes() dataGoogleComputeNetworkPeeringAttributes {
	return dataGoogleComputeNetworkPeeringAttributes{ref: terra.ReferenceDataSource(gcnp)}
}

// DataArgs contains the configurations for google_compute_network_peering.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataGoogleComputeNetworkPeeringAttributes struct {
	ref terra.Reference
}

// ExportCustomRoutes returns a reference to field export_custom_routes of google_compute_network_peering.
func (gcnp dataGoogleComputeNetworkPeeringAttributes) ExportCustomRoutes() terra.BoolValue {
	return terra.ReferenceAsBool(gcnp.ref.Append("export_custom_routes"))
}

// ExportSubnetRoutesWithPublicIp returns a reference to field export_subnet_routes_with_public_ip of google_compute_network_peering.
func (gcnp dataGoogleComputeNetworkPeeringAttributes) ExportSubnetRoutesWithPublicIp() terra.BoolValue {
	return terra.ReferenceAsBool(gcnp.ref.Append("export_subnet_routes_with_public_ip"))
}

// Id returns a reference to field id of google_compute_network_peering.
func (gcnp dataGoogleComputeNetworkPeeringAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcnp.ref.Append("id"))
}

// ImportCustomRoutes returns a reference to field import_custom_routes of google_compute_network_peering.
func (gcnp dataGoogleComputeNetworkPeeringAttributes) ImportCustomRoutes() terra.BoolValue {
	return terra.ReferenceAsBool(gcnp.ref.Append("import_custom_routes"))
}

// ImportSubnetRoutesWithPublicIp returns a reference to field import_subnet_routes_with_public_ip of google_compute_network_peering.
func (gcnp dataGoogleComputeNetworkPeeringAttributes) ImportSubnetRoutesWithPublicIp() terra.BoolValue {
	return terra.ReferenceAsBool(gcnp.ref.Append("import_subnet_routes_with_public_ip"))
}

// Name returns a reference to field name of google_compute_network_peering.
func (gcnp dataGoogleComputeNetworkPeeringAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcnp.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_network_peering.
func (gcnp dataGoogleComputeNetworkPeeringAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(gcnp.ref.Append("network"))
}

// PeerNetwork returns a reference to field peer_network of google_compute_network_peering.
func (gcnp dataGoogleComputeNetworkPeeringAttributes) PeerNetwork() terra.StringValue {
	return terra.ReferenceAsString(gcnp.ref.Append("peer_network"))
}

// StackType returns a reference to field stack_type of google_compute_network_peering.
func (gcnp dataGoogleComputeNetworkPeeringAttributes) StackType() terra.StringValue {
	return terra.ReferenceAsString(gcnp.ref.Append("stack_type"))
}

// State returns a reference to field state of google_compute_network_peering.
func (gcnp dataGoogleComputeNetworkPeeringAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(gcnp.ref.Append("state"))
}

// StateDetails returns a reference to field state_details of google_compute_network_peering.
func (gcnp dataGoogleComputeNetworkPeeringAttributes) StateDetails() terra.StringValue {
	return terra.ReferenceAsString(gcnp.ref.Append("state_details"))
}

func (gcnp dataGoogleComputeNetworkPeeringAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](gcnp.ref.Append("timeouts"))
}
