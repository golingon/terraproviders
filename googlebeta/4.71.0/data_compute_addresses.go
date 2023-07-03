// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datacomputeaddresses "github.com/golingon/terraproviders/googlebeta/4.71.0/datacomputeaddresses"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeAddresses creates a new instance of [DataComputeAddresses].
func NewDataComputeAddresses(name string, args DataComputeAddressesArgs) *DataComputeAddresses {
	return &DataComputeAddresses{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeAddresses)(nil)

// DataComputeAddresses represents the Terraform data resource google_compute_addresses.
type DataComputeAddresses struct {
	Name string
	Args DataComputeAddressesArgs
}

// DataSource returns the Terraform object type for [DataComputeAddresses].
func (ca *DataComputeAddresses) DataSource() string {
	return "google_compute_addresses"
}

// LocalName returns the local name for [DataComputeAddresses].
func (ca *DataComputeAddresses) LocalName() string {
	return ca.Name
}

// Configuration returns the configuration (args) for [DataComputeAddresses].
func (ca *DataComputeAddresses) Configuration() interface{} {
	return ca.Args
}

// Attributes returns the attributes for [DataComputeAddresses].
func (ca *DataComputeAddresses) Attributes() dataComputeAddressesAttributes {
	return dataComputeAddressesAttributes{ref: terra.ReferenceDataResource(ca)}
}

// DataComputeAddressesArgs contains the configurations for google_compute_addresses.
type DataComputeAddressesArgs struct {
	// Filter: string, optional
	Filter terra.StringValue `hcl:"filter,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Addresses: min=0
	Addresses []datacomputeaddresses.Addresses `hcl:"addresses,block" validate:"min=0"`
}
type dataComputeAddressesAttributes struct {
	ref terra.Reference
}

// Filter returns a reference to field filter of google_compute_addresses.
func (ca dataComputeAddressesAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("filter"))
}

// Id returns a reference to field id of google_compute_addresses.
func (ca dataComputeAddressesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("id"))
}

// Project returns a reference to field project of google_compute_addresses.
func (ca dataComputeAddressesAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_addresses.
func (ca dataComputeAddressesAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("region"))
}

func (ca dataComputeAddressesAttributes) Addresses() terra.ListValue[datacomputeaddresses.AddressesAttributes] {
	return terra.ReferenceAsList[datacomputeaddresses.AddressesAttributes](ca.ref.Append("addresses"))
}
