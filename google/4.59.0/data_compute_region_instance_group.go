// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datacomputeregioninstancegroup "github.com/golingon/terraproviders/google/4.59.0/datacomputeregioninstancegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeRegionInstanceGroup creates a new instance of [DataComputeRegionInstanceGroup].
func NewDataComputeRegionInstanceGroup(name string, args DataComputeRegionInstanceGroupArgs) *DataComputeRegionInstanceGroup {
	return &DataComputeRegionInstanceGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeRegionInstanceGroup)(nil)

// DataComputeRegionInstanceGroup represents the Terraform data resource google_compute_region_instance_group.
type DataComputeRegionInstanceGroup struct {
	Name string
	Args DataComputeRegionInstanceGroupArgs
}

// DataSource returns the Terraform object type for [DataComputeRegionInstanceGroup].
func (crig *DataComputeRegionInstanceGroup) DataSource() string {
	return "google_compute_region_instance_group"
}

// LocalName returns the local name for [DataComputeRegionInstanceGroup].
func (crig *DataComputeRegionInstanceGroup) LocalName() string {
	return crig.Name
}

// Configuration returns the configuration (args) for [DataComputeRegionInstanceGroup].
func (crig *DataComputeRegionInstanceGroup) Configuration() interface{} {
	return crig.Args
}

// Attributes returns the attributes for [DataComputeRegionInstanceGroup].
func (crig *DataComputeRegionInstanceGroup) Attributes() dataComputeRegionInstanceGroupAttributes {
	return dataComputeRegionInstanceGroupAttributes{ref: terra.ReferenceDataResource(crig)}
}

// DataComputeRegionInstanceGroupArgs contains the configurations for google_compute_region_instance_group.
type DataComputeRegionInstanceGroupArgs struct {
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
	// Instances: min=0
	Instances []datacomputeregioninstancegroup.Instances `hcl:"instances,block" validate:"min=0"`
}
type dataComputeRegionInstanceGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_compute_region_instance_group.
func (crig dataComputeRegionInstanceGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crig.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_instance_group.
func (crig dataComputeRegionInstanceGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crig.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_instance_group.
func (crig dataComputeRegionInstanceGroupAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crig.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_instance_group.
func (crig dataComputeRegionInstanceGroupAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crig.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_instance_group.
func (crig dataComputeRegionInstanceGroupAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crig.ref.Append("self_link"))
}

// Size returns a reference to field size of google_compute_region_instance_group.
func (crig dataComputeRegionInstanceGroupAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(crig.ref.Append("size"))
}

func (crig dataComputeRegionInstanceGroupAttributes) Instances() terra.ListValue[datacomputeregioninstancegroup.InstancesAttributes] {
	return terra.ReferenceAsList[datacomputeregioninstancegroup.InstancesAttributes](crig.ref.Append("instances"))
}
