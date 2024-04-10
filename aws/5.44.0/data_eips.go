// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"github.com/golingon/lingon/pkg/terra"
	dataeips "github.com/golingon/terraproviders/aws/5.44.0/dataeips"
)

// NewDataEips creates a new instance of [DataEips].
func NewDataEips(name string, args DataEipsArgs) *DataEips {
	return &DataEips{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEips)(nil)

// DataEips represents the Terraform data resource aws_eips.
type DataEips struct {
	Name string
	Args DataEipsArgs
}

// DataSource returns the Terraform object type for [DataEips].
func (e *DataEips) DataSource() string {
	return "aws_eips"
}

// LocalName returns the local name for [DataEips].
func (e *DataEips) LocalName() string {
	return e.Name
}

// Configuration returns the configuration (args) for [DataEips].
func (e *DataEips) Configuration() interface{} {
	return e.Args
}

// Attributes returns the attributes for [DataEips].
func (e *DataEips) Attributes() dataEipsAttributes {
	return dataEipsAttributes{ref: terra.ReferenceDataResource(e)}
}

// DataEipsArgs contains the configurations for aws_eips.
type DataEipsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataeips.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataeips.Timeouts `hcl:"timeouts,block"`
}
type dataEipsAttributes struct {
	ref terra.Reference
}

// AllocationIds returns a reference to field allocation_ids of aws_eips.
func (e dataEipsAttributes) AllocationIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](e.ref.Append("allocation_ids"))
}

// Id returns a reference to field id of aws_eips.
func (e dataEipsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("id"))
}

// PublicIps returns a reference to field public_ips of aws_eips.
func (e dataEipsAttributes) PublicIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](e.ref.Append("public_ips"))
}

// Tags returns a reference to field tags of aws_eips.
func (e dataEipsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](e.ref.Append("tags"))
}

func (e dataEipsAttributes) Filter() terra.SetValue[dataeips.FilterAttributes] {
	return terra.ReferenceAsSet[dataeips.FilterAttributes](e.ref.Append("filter"))
}

func (e dataEipsAttributes) Timeouts() dataeips.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataeips.TimeoutsAttributes](e.ref.Append("timeouts"))
}
