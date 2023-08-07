// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datavmwareenginenetwork "github.com/golingon/terraproviders/googlebeta/4.76.0/datavmwareenginenetwork"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVmwareengineNetwork creates a new instance of [DataVmwareengineNetwork].
func NewDataVmwareengineNetwork(name string, args DataVmwareengineNetworkArgs) *DataVmwareengineNetwork {
	return &DataVmwareengineNetwork{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVmwareengineNetwork)(nil)

// DataVmwareengineNetwork represents the Terraform data resource google_vmwareengine_network.
type DataVmwareengineNetwork struct {
	Name string
	Args DataVmwareengineNetworkArgs
}

// DataSource returns the Terraform object type for [DataVmwareengineNetwork].
func (vn *DataVmwareengineNetwork) DataSource() string {
	return "google_vmwareengine_network"
}

// LocalName returns the local name for [DataVmwareengineNetwork].
func (vn *DataVmwareengineNetwork) LocalName() string {
	return vn.Name
}

// Configuration returns the configuration (args) for [DataVmwareengineNetwork].
func (vn *DataVmwareengineNetwork) Configuration() interface{} {
	return vn.Args
}

// Attributes returns the attributes for [DataVmwareengineNetwork].
func (vn *DataVmwareengineNetwork) Attributes() dataVmwareengineNetworkAttributes {
	return dataVmwareengineNetworkAttributes{ref: terra.ReferenceDataResource(vn)}
}

// DataVmwareengineNetworkArgs contains the configurations for google_vmwareengine_network.
type DataVmwareengineNetworkArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// VpcNetworks: min=0
	VpcNetworks []datavmwareenginenetwork.VpcNetworks `hcl:"vpc_networks,block" validate:"min=0"`
}
type dataVmwareengineNetworkAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_vmwareengine_network.
func (vn dataVmwareengineNetworkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("description"))
}

// Id returns a reference to field id of google_vmwareengine_network.
func (vn dataVmwareengineNetworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("id"))
}

// Location returns a reference to field location of google_vmwareengine_network.
func (vn dataVmwareengineNetworkAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("location"))
}

// Name returns a reference to field name of google_vmwareengine_network.
func (vn dataVmwareengineNetworkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("name"))
}

// Project returns a reference to field project of google_vmwareengine_network.
func (vn dataVmwareengineNetworkAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("project"))
}

// State returns a reference to field state of google_vmwareengine_network.
func (vn dataVmwareengineNetworkAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("state"))
}

// Type returns a reference to field type of google_vmwareengine_network.
func (vn dataVmwareengineNetworkAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("type"))
}

// Uid returns a reference to field uid of google_vmwareengine_network.
func (vn dataVmwareengineNetworkAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("uid"))
}

func (vn dataVmwareengineNetworkAttributes) VpcNetworks() terra.ListValue[datavmwareenginenetwork.VpcNetworksAttributes] {
	return terra.ReferenceAsList[datavmwareenginenetwork.VpcNetworksAttributes](vn.ref.Append("vpc_networks"))
}
