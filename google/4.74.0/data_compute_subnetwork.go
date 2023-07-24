// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datacomputesubnetwork "github.com/golingon/terraproviders/google/4.74.0/datacomputesubnetwork"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeSubnetwork creates a new instance of [DataComputeSubnetwork].
func NewDataComputeSubnetwork(name string, args DataComputeSubnetworkArgs) *DataComputeSubnetwork {
	return &DataComputeSubnetwork{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeSubnetwork)(nil)

// DataComputeSubnetwork represents the Terraform data resource google_compute_subnetwork.
type DataComputeSubnetwork struct {
	Name string
	Args DataComputeSubnetworkArgs
}

// DataSource returns the Terraform object type for [DataComputeSubnetwork].
func (cs *DataComputeSubnetwork) DataSource() string {
	return "google_compute_subnetwork"
}

// LocalName returns the local name for [DataComputeSubnetwork].
func (cs *DataComputeSubnetwork) LocalName() string {
	return cs.Name
}

// Configuration returns the configuration (args) for [DataComputeSubnetwork].
func (cs *DataComputeSubnetwork) Configuration() interface{} {
	return cs.Args
}

// Attributes returns the attributes for [DataComputeSubnetwork].
func (cs *DataComputeSubnetwork) Attributes() dataComputeSubnetworkAttributes {
	return dataComputeSubnetworkAttributes{ref: terra.ReferenceDataResource(cs)}
}

// DataComputeSubnetworkArgs contains the configurations for google_compute_subnetwork.
type DataComputeSubnetworkArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// SelfLink: string, optional
	SelfLink terra.StringValue `hcl:"self_link,attr"`
	// SecondaryIpRange: min=0
	SecondaryIpRange []datacomputesubnetwork.SecondaryIpRange `hcl:"secondary_ip_range,block" validate:"min=0"`
}
type dataComputeSubnetworkAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_subnetwork.
func (cs dataComputeSubnetworkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("description"))
}

// GatewayAddress returns a reference to field gateway_address of google_compute_subnetwork.
func (cs dataComputeSubnetworkAttributes) GatewayAddress() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("gateway_address"))
}

// Id returns a reference to field id of google_compute_subnetwork.
func (cs dataComputeSubnetworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("id"))
}

// IpCidrRange returns a reference to field ip_cidr_range of google_compute_subnetwork.
func (cs dataComputeSubnetworkAttributes) IpCidrRange() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("ip_cidr_range"))
}

// Name returns a reference to field name of google_compute_subnetwork.
func (cs dataComputeSubnetworkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_subnetwork.
func (cs dataComputeSubnetworkAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("network"))
}

// PrivateIpGoogleAccess returns a reference to field private_ip_google_access of google_compute_subnetwork.
func (cs dataComputeSubnetworkAttributes) PrivateIpGoogleAccess() terra.BoolValue {
	return terra.ReferenceAsBool(cs.ref.Append("private_ip_google_access"))
}

// Project returns a reference to field project of google_compute_subnetwork.
func (cs dataComputeSubnetworkAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_subnetwork.
func (cs dataComputeSubnetworkAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_subnetwork.
func (cs dataComputeSubnetworkAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("self_link"))
}

func (cs dataComputeSubnetworkAttributes) SecondaryIpRange() terra.ListValue[datacomputesubnetwork.SecondaryIpRangeAttributes] {
	return terra.ReferenceAsList[datacomputesubnetwork.SecondaryIpRangeAttributes](cs.ref.Append("secondary_ip_range"))
}
