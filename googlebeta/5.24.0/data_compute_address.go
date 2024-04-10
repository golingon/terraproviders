// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataComputeAddress creates a new instance of [DataComputeAddress].
func NewDataComputeAddress(name string, args DataComputeAddressArgs) *DataComputeAddress {
	return &DataComputeAddress{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeAddress)(nil)

// DataComputeAddress represents the Terraform data resource google_compute_address.
type DataComputeAddress struct {
	Name string
	Args DataComputeAddressArgs
}

// DataSource returns the Terraform object type for [DataComputeAddress].
func (ca *DataComputeAddress) DataSource() string {
	return "google_compute_address"
}

// LocalName returns the local name for [DataComputeAddress].
func (ca *DataComputeAddress) LocalName() string {
	return ca.Name
}

// Configuration returns the configuration (args) for [DataComputeAddress].
func (ca *DataComputeAddress) Configuration() interface{} {
	return ca.Args
}

// Attributes returns the attributes for [DataComputeAddress].
func (ca *DataComputeAddress) Attributes() dataComputeAddressAttributes {
	return dataComputeAddressAttributes{ref: terra.ReferenceDataResource(ca)}
}

// DataComputeAddressArgs contains the configurations for google_compute_address.
type DataComputeAddressArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}
type dataComputeAddressAttributes struct {
	ref terra.Reference
}

// Address returns a reference to field address of google_compute_address.
func (ca dataComputeAddressAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("address"))
}

// AddressType returns a reference to field address_type of google_compute_address.
func (ca dataComputeAddressAttributes) AddressType() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("address_type"))
}

// Id returns a reference to field id of google_compute_address.
func (ca dataComputeAddressAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_address.
func (ca dataComputeAddressAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_address.
func (ca dataComputeAddressAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("network"))
}

// NetworkTier returns a reference to field network_tier of google_compute_address.
func (ca dataComputeAddressAttributes) NetworkTier() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("network_tier"))
}

// PrefixLength returns a reference to field prefix_length of google_compute_address.
func (ca dataComputeAddressAttributes) PrefixLength() terra.NumberValue {
	return terra.ReferenceAsNumber(ca.ref.Append("prefix_length"))
}

// Project returns a reference to field project of google_compute_address.
func (ca dataComputeAddressAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("project"))
}

// Purpose returns a reference to field purpose of google_compute_address.
func (ca dataComputeAddressAttributes) Purpose() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("purpose"))
}

// Region returns a reference to field region of google_compute_address.
func (ca dataComputeAddressAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_address.
func (ca dataComputeAddressAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("self_link"))
}

// Status returns a reference to field status of google_compute_address.
func (ca dataComputeAddressAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("status"))
}

// Subnetwork returns a reference to field subnetwork of google_compute_address.
func (ca dataComputeAddressAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("subnetwork"))
}

// Users returns a reference to field users of google_compute_address.
func (ca dataComputeAddressAttributes) Users() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("users"))
}
