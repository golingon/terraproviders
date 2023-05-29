// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datacomputerouter "github.com/golingon/terraproviders/google/4.63.1/datacomputerouter"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeRouter creates a new instance of [DataComputeRouter].
func NewDataComputeRouter(name string, args DataComputeRouterArgs) *DataComputeRouter {
	return &DataComputeRouter{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeRouter)(nil)

// DataComputeRouter represents the Terraform data resource google_compute_router.
type DataComputeRouter struct {
	Name string
	Args DataComputeRouterArgs
}

// DataSource returns the Terraform object type for [DataComputeRouter].
func (cr *DataComputeRouter) DataSource() string {
	return "google_compute_router"
}

// LocalName returns the local name for [DataComputeRouter].
func (cr *DataComputeRouter) LocalName() string {
	return cr.Name
}

// Configuration returns the configuration (args) for [DataComputeRouter].
func (cr *DataComputeRouter) Configuration() interface{} {
	return cr.Args
}

// Attributes returns the attributes for [DataComputeRouter].
func (cr *DataComputeRouter) Attributes() dataComputeRouterAttributes {
	return dataComputeRouterAttributes{ref: terra.ReferenceDataResource(cr)}
}

// DataComputeRouterArgs contains the configurations for google_compute_router.
type DataComputeRouterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Bgp: min=0
	Bgp []datacomputerouter.Bgp `hcl:"bgp,block" validate:"min=0"`
}
type dataComputeRouterAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_router.
func (cr dataComputeRouterAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_router.
func (cr dataComputeRouterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("description"))
}

// EncryptedInterconnectRouter returns a reference to field encrypted_interconnect_router of google_compute_router.
func (cr dataComputeRouterAttributes) EncryptedInterconnectRouter() terra.BoolValue {
	return terra.ReferenceAsBool(cr.ref.Append("encrypted_interconnect_router"))
}

// Id returns a reference to field id of google_compute_router.
func (cr dataComputeRouterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_router.
func (cr dataComputeRouterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_router.
func (cr dataComputeRouterAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("network"))
}

// Project returns a reference to field project of google_compute_router.
func (cr dataComputeRouterAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_router.
func (cr dataComputeRouterAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_router.
func (cr dataComputeRouterAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("self_link"))
}

func (cr dataComputeRouterAttributes) Bgp() terra.ListValue[datacomputerouter.BgpAttributes] {
	return terra.ReferenceAsList[datacomputerouter.BgpAttributes](cr.ref.Append("bgp"))
}