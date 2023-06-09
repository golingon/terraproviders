// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataComputeNetwork creates a new instance of [DataComputeNetwork].
func NewDataComputeNetwork(name string, args DataComputeNetworkArgs) *DataComputeNetwork {
	return &DataComputeNetwork{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeNetwork)(nil)

// DataComputeNetwork represents the Terraform data resource google_compute_network.
type DataComputeNetwork struct {
	Name string
	Args DataComputeNetworkArgs
}

// DataSource returns the Terraform object type for [DataComputeNetwork].
func (cn *DataComputeNetwork) DataSource() string {
	return "google_compute_network"
}

// LocalName returns the local name for [DataComputeNetwork].
func (cn *DataComputeNetwork) LocalName() string {
	return cn.Name
}

// Configuration returns the configuration (args) for [DataComputeNetwork].
func (cn *DataComputeNetwork) Configuration() interface{} {
	return cn.Args
}

// Attributes returns the attributes for [DataComputeNetwork].
func (cn *DataComputeNetwork) Attributes() dataComputeNetworkAttributes {
	return dataComputeNetworkAttributes{ref: terra.ReferenceDataResource(cn)}
}

// DataComputeNetworkArgs contains the configurations for google_compute_network.
type DataComputeNetworkArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataComputeNetworkAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_network.
func (cn dataComputeNetworkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("description"))
}

// GatewayIpv4 returns a reference to field gateway_ipv4 of google_compute_network.
func (cn dataComputeNetworkAttributes) GatewayIpv4() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("gateway_ipv4"))
}

// Id returns a reference to field id of google_compute_network.
func (cn dataComputeNetworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_network.
func (cn dataComputeNetworkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_network.
func (cn dataComputeNetworkAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_network.
func (cn dataComputeNetworkAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("self_link"))
}

// SubnetworksSelfLinks returns a reference to field subnetworks_self_links of google_compute_network.
func (cn dataComputeNetworkAttributes) SubnetworksSelfLinks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cn.ref.Append("subnetworks_self_links"))
}
