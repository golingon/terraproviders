// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datavpcs "github.com/golingon/terraproviders/aws/5.7.0/datavpcs"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVpcs creates a new instance of [DataVpcs].
func NewDataVpcs(name string, args DataVpcsArgs) *DataVpcs {
	return &DataVpcs{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpcs)(nil)

// DataVpcs represents the Terraform data resource aws_vpcs.
type DataVpcs struct {
	Name string
	Args DataVpcsArgs
}

// DataSource returns the Terraform object type for [DataVpcs].
func (v *DataVpcs) DataSource() string {
	return "aws_vpcs"
}

// LocalName returns the local name for [DataVpcs].
func (v *DataVpcs) LocalName() string {
	return v.Name
}

// Configuration returns the configuration (args) for [DataVpcs].
func (v *DataVpcs) Configuration() interface{} {
	return v.Args
}

// Attributes returns the attributes for [DataVpcs].
func (v *DataVpcs) Attributes() dataVpcsAttributes {
	return dataVpcsAttributes{ref: terra.ReferenceDataResource(v)}
}

// DataVpcsArgs contains the configurations for aws_vpcs.
type DataVpcsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []datavpcs.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datavpcs.Timeouts `hcl:"timeouts,block"`
}
type dataVpcsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_vpcs.
func (v dataVpcsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_vpcs.
func (v dataVpcsAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](v.ref.Append("ids"))
}

// Tags returns a reference to field tags of aws_vpcs.
func (v dataVpcsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](v.ref.Append("tags"))
}

func (v dataVpcsAttributes) Filter() terra.SetValue[datavpcs.FilterAttributes] {
	return terra.ReferenceAsSet[datavpcs.FilterAttributes](v.ref.Append("filter"))
}

func (v dataVpcsAttributes) Timeouts() datavpcs.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavpcs.TimeoutsAttributes](v.ref.Append("timeouts"))
}
