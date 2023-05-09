// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datacomputeregionnetworkendpointgroup "github.com/golingon/terraproviders/google/4.64.0/datacomputeregionnetworkendpointgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeRegionNetworkEndpointGroup creates a new instance of [DataComputeRegionNetworkEndpointGroup].
func NewDataComputeRegionNetworkEndpointGroup(name string, args DataComputeRegionNetworkEndpointGroupArgs) *DataComputeRegionNetworkEndpointGroup {
	return &DataComputeRegionNetworkEndpointGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeRegionNetworkEndpointGroup)(nil)

// DataComputeRegionNetworkEndpointGroup represents the Terraform data resource google_compute_region_network_endpoint_group.
type DataComputeRegionNetworkEndpointGroup struct {
	Name string
	Args DataComputeRegionNetworkEndpointGroupArgs
}

// DataSource returns the Terraform object type for [DataComputeRegionNetworkEndpointGroup].
func (crneg *DataComputeRegionNetworkEndpointGroup) DataSource() string {
	return "google_compute_region_network_endpoint_group"
}

// LocalName returns the local name for [DataComputeRegionNetworkEndpointGroup].
func (crneg *DataComputeRegionNetworkEndpointGroup) LocalName() string {
	return crneg.Name
}

// Configuration returns the configuration (args) for [DataComputeRegionNetworkEndpointGroup].
func (crneg *DataComputeRegionNetworkEndpointGroup) Configuration() interface{} {
	return crneg.Args
}

// Attributes returns the attributes for [DataComputeRegionNetworkEndpointGroup].
func (crneg *DataComputeRegionNetworkEndpointGroup) Attributes() dataComputeRegionNetworkEndpointGroupAttributes {
	return dataComputeRegionNetworkEndpointGroupAttributes{ref: terra.ReferenceDataResource(crneg)}
}

// DataComputeRegionNetworkEndpointGroupArgs contains the configurations for google_compute_region_network_endpoint_group.
type DataComputeRegionNetworkEndpointGroupArgs struct {
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
	// AppEngine: min=0
	AppEngine []datacomputeregionnetworkendpointgroup.AppEngine `hcl:"app_engine,block" validate:"min=0"`
	// CloudFunction: min=0
	CloudFunction []datacomputeregionnetworkendpointgroup.CloudFunction `hcl:"cloud_function,block" validate:"min=0"`
	// CloudRun: min=0
	CloudRun []datacomputeregionnetworkendpointgroup.CloudRun `hcl:"cloud_run,block" validate:"min=0"`
}
type dataComputeRegionNetworkEndpointGroupAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_region_network_endpoint_group.
func (crneg dataComputeRegionNetworkEndpointGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_region_network_endpoint_group.
func (crneg dataComputeRegionNetworkEndpointGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_network_endpoint_group.
func (crneg dataComputeRegionNetworkEndpointGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_region_network_endpoint_group.
func (crneg dataComputeRegionNetworkEndpointGroupAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("network"))
}

// NetworkEndpointType returns a reference to field network_endpoint_type of google_compute_region_network_endpoint_group.
func (crneg dataComputeRegionNetworkEndpointGroupAttributes) NetworkEndpointType() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("network_endpoint_type"))
}

// Project returns a reference to field project of google_compute_region_network_endpoint_group.
func (crneg dataComputeRegionNetworkEndpointGroupAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("project"))
}

// PscTargetService returns a reference to field psc_target_service of google_compute_region_network_endpoint_group.
func (crneg dataComputeRegionNetworkEndpointGroupAttributes) PscTargetService() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("psc_target_service"))
}

// Region returns a reference to field region of google_compute_region_network_endpoint_group.
func (crneg dataComputeRegionNetworkEndpointGroupAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_network_endpoint_group.
func (crneg dataComputeRegionNetworkEndpointGroupAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("self_link"))
}

// Subnetwork returns a reference to field subnetwork of google_compute_region_network_endpoint_group.
func (crneg dataComputeRegionNetworkEndpointGroupAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("subnetwork"))
}

func (crneg dataComputeRegionNetworkEndpointGroupAttributes) AppEngine() terra.ListValue[datacomputeregionnetworkendpointgroup.AppEngineAttributes] {
	return terra.ReferenceAsList[datacomputeregionnetworkendpointgroup.AppEngineAttributes](crneg.ref.Append("app_engine"))
}

func (crneg dataComputeRegionNetworkEndpointGroupAttributes) CloudFunction() terra.ListValue[datacomputeregionnetworkendpointgroup.CloudFunctionAttributes] {
	return terra.ReferenceAsList[datacomputeregionnetworkendpointgroup.CloudFunctionAttributes](crneg.ref.Append("cloud_function"))
}

func (crneg dataComputeRegionNetworkEndpointGroupAttributes) CloudRun() terra.ListValue[datacomputeregionnetworkendpointgroup.CloudRunAttributes] {
	return terra.ReferenceAsList[datacomputeregionnetworkendpointgroup.CloudRunAttributes](crneg.ref.Append("cloud_run"))
}
