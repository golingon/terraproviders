// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataregions "github.com/golingon/terraproviders/aws/5.9.0/dataregions"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataRegions creates a new instance of [DataRegions].
func NewDataRegions(name string, args DataRegionsArgs) *DataRegions {
	return &DataRegions{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRegions)(nil)

// DataRegions represents the Terraform data resource aws_regions.
type DataRegions struct {
	Name string
	Args DataRegionsArgs
}

// DataSource returns the Terraform object type for [DataRegions].
func (r *DataRegions) DataSource() string {
	return "aws_regions"
}

// LocalName returns the local name for [DataRegions].
func (r *DataRegions) LocalName() string {
	return r.Name
}

// Configuration returns the configuration (args) for [DataRegions].
func (r *DataRegions) Configuration() interface{} {
	return r.Args
}

// Attributes returns the attributes for [DataRegions].
func (r *DataRegions) Attributes() dataRegionsAttributes {
	return dataRegionsAttributes{ref: terra.ReferenceDataResource(r)}
}

// DataRegionsArgs contains the configurations for aws_regions.
type DataRegionsArgs struct {
	// AllRegions: bool, optional
	AllRegions terra.BoolValue `hcl:"all_regions,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Filter: min=0
	Filter []dataregions.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataRegionsAttributes struct {
	ref terra.Reference
}

// AllRegions returns a reference to field all_regions of aws_regions.
func (r dataRegionsAttributes) AllRegions() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("all_regions"))
}

// Id returns a reference to field id of aws_regions.
func (r dataRegionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("id"))
}

// Names returns a reference to field names of aws_regions.
func (r dataRegionsAttributes) Names() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](r.ref.Append("names"))
}

func (r dataRegionsAttributes) Filter() terra.SetValue[dataregions.FilterAttributes] {
	return terra.ReferenceAsSet[dataregions.FilterAttributes](r.ref.Append("filter"))
}
