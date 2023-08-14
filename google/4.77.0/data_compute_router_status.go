// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datacomputerouterstatus "github.com/golingon/terraproviders/google/4.77.0/datacomputerouterstatus"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeRouterStatus creates a new instance of [DataComputeRouterStatus].
func NewDataComputeRouterStatus(name string, args DataComputeRouterStatusArgs) *DataComputeRouterStatus {
	return &DataComputeRouterStatus{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeRouterStatus)(nil)

// DataComputeRouterStatus represents the Terraform data resource google_compute_router_status.
type DataComputeRouterStatus struct {
	Name string
	Args DataComputeRouterStatusArgs
}

// DataSource returns the Terraform object type for [DataComputeRouterStatus].
func (crs *DataComputeRouterStatus) DataSource() string {
	return "google_compute_router_status"
}

// LocalName returns the local name for [DataComputeRouterStatus].
func (crs *DataComputeRouterStatus) LocalName() string {
	return crs.Name
}

// Configuration returns the configuration (args) for [DataComputeRouterStatus].
func (crs *DataComputeRouterStatus) Configuration() interface{} {
	return crs.Args
}

// Attributes returns the attributes for [DataComputeRouterStatus].
func (crs *DataComputeRouterStatus) Attributes() dataComputeRouterStatusAttributes {
	return dataComputeRouterStatusAttributes{ref: terra.ReferenceDataResource(crs)}
}

// DataComputeRouterStatusArgs contains the configurations for google_compute_router_status.
type DataComputeRouterStatusArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// BestRoutes: min=0
	BestRoutes []datacomputerouterstatus.BestRoutes `hcl:"best_routes,block" validate:"min=0"`
	// BestRoutesForRouter: min=0
	BestRoutesForRouter []datacomputerouterstatus.BestRoutesForRouter `hcl:"best_routes_for_router,block" validate:"min=0"`
}
type dataComputeRouterStatusAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_compute_router_status.
func (crs dataComputeRouterStatusAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_router_status.
func (crs dataComputeRouterStatusAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_router_status.
func (crs dataComputeRouterStatusAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("network"))
}

// Project returns a reference to field project of google_compute_router_status.
func (crs dataComputeRouterStatusAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_router_status.
func (crs dataComputeRouterStatusAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crs.ref.Append("region"))
}

func (crs dataComputeRouterStatusAttributes) BestRoutes() terra.ListValue[datacomputerouterstatus.BestRoutesAttributes] {
	return terra.ReferenceAsList[datacomputerouterstatus.BestRoutesAttributes](crs.ref.Append("best_routes"))
}

func (crs dataComputeRouterStatusAttributes) BestRoutesForRouter() terra.ListValue[datacomputerouterstatus.BestRoutesForRouterAttributes] {
	return terra.ReferenceAsList[datacomputerouterstatus.BestRoutesForRouterAttributes](crs.ref.Append("best_routes_for_router"))
}
