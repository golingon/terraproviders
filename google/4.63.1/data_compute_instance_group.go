// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datacomputeinstancegroup "github.com/golingon/terraproviders/google/4.63.1/datacomputeinstancegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeInstanceGroup creates a new instance of [DataComputeInstanceGroup].
func NewDataComputeInstanceGroup(name string, args DataComputeInstanceGroupArgs) *DataComputeInstanceGroup {
	return &DataComputeInstanceGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeInstanceGroup)(nil)

// DataComputeInstanceGroup represents the Terraform data resource google_compute_instance_group.
type DataComputeInstanceGroup struct {
	Name string
	Args DataComputeInstanceGroupArgs
}

// DataSource returns the Terraform object type for [DataComputeInstanceGroup].
func (cig *DataComputeInstanceGroup) DataSource() string {
	return "google_compute_instance_group"
}

// LocalName returns the local name for [DataComputeInstanceGroup].
func (cig *DataComputeInstanceGroup) LocalName() string {
	return cig.Name
}

// Configuration returns the configuration (args) for [DataComputeInstanceGroup].
func (cig *DataComputeInstanceGroup) Configuration() interface{} {
	return cig.Args
}

// Attributes returns the attributes for [DataComputeInstanceGroup].
func (cig *DataComputeInstanceGroup) Attributes() dataComputeInstanceGroupAttributes {
	return dataComputeInstanceGroupAttributes{ref: terra.ReferenceDataResource(cig)}
}

// DataComputeInstanceGroupArgs contains the configurations for google_compute_instance_group.
type DataComputeInstanceGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SelfLink: string, optional
	SelfLink terra.StringValue `hcl:"self_link,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// NamedPort: min=0
	NamedPort []datacomputeinstancegroup.NamedPort `hcl:"named_port,block" validate:"min=0"`
}
type dataComputeInstanceGroupAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_instance_group.
func (cig dataComputeInstanceGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_instance_group.
func (cig dataComputeInstanceGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("id"))
}

// Instances returns a reference to field instances of google_compute_instance_group.
func (cig dataComputeInstanceGroupAttributes) Instances() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cig.ref.Append("instances"))
}

// Name returns a reference to field name of google_compute_instance_group.
func (cig dataComputeInstanceGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_instance_group.
func (cig dataComputeInstanceGroupAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("network"))
}

// Project returns a reference to field project of google_compute_instance_group.
func (cig dataComputeInstanceGroupAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_instance_group.
func (cig dataComputeInstanceGroupAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("self_link"))
}

// Size returns a reference to field size of google_compute_instance_group.
func (cig dataComputeInstanceGroupAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(cig.ref.Append("size"))
}

// Zone returns a reference to field zone of google_compute_instance_group.
func (cig dataComputeInstanceGroupAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cig.ref.Append("zone"))
}

func (cig dataComputeInstanceGroupAttributes) NamedPort() terra.ListValue[datacomputeinstancegroup.NamedPortAttributes] {
	return terra.ReferenceAsList[datacomputeinstancegroup.NamedPortAttributes](cig.ref.Append("named_port"))
}
