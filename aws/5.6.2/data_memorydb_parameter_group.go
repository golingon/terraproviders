// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datamemorydbparametergroup "github.com/golingon/terraproviders/aws/5.6.2/datamemorydbparametergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMemorydbParameterGroup creates a new instance of [DataMemorydbParameterGroup].
func NewDataMemorydbParameterGroup(name string, args DataMemorydbParameterGroupArgs) *DataMemorydbParameterGroup {
	return &DataMemorydbParameterGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMemorydbParameterGroup)(nil)

// DataMemorydbParameterGroup represents the Terraform data resource aws_memorydb_parameter_group.
type DataMemorydbParameterGroup struct {
	Name string
	Args DataMemorydbParameterGroupArgs
}

// DataSource returns the Terraform object type for [DataMemorydbParameterGroup].
func (mpg *DataMemorydbParameterGroup) DataSource() string {
	return "aws_memorydb_parameter_group"
}

// LocalName returns the local name for [DataMemorydbParameterGroup].
func (mpg *DataMemorydbParameterGroup) LocalName() string {
	return mpg.Name
}

// Configuration returns the configuration (args) for [DataMemorydbParameterGroup].
func (mpg *DataMemorydbParameterGroup) Configuration() interface{} {
	return mpg.Args
}

// Attributes returns the attributes for [DataMemorydbParameterGroup].
func (mpg *DataMemorydbParameterGroup) Attributes() dataMemorydbParameterGroupAttributes {
	return dataMemorydbParameterGroupAttributes{ref: terra.ReferenceDataResource(mpg)}
}

// DataMemorydbParameterGroupArgs contains the configurations for aws_memorydb_parameter_group.
type DataMemorydbParameterGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Parameter: min=0
	Parameter []datamemorydbparametergroup.Parameter `hcl:"parameter,block" validate:"min=0"`
}
type dataMemorydbParameterGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_memorydb_parameter_group.
func (mpg dataMemorydbParameterGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(mpg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_memorydb_parameter_group.
func (mpg dataMemorydbParameterGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mpg.ref.Append("description"))
}

// Family returns a reference to field family of aws_memorydb_parameter_group.
func (mpg dataMemorydbParameterGroupAttributes) Family() terra.StringValue {
	return terra.ReferenceAsString(mpg.ref.Append("family"))
}

// Id returns a reference to field id of aws_memorydb_parameter_group.
func (mpg dataMemorydbParameterGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mpg.ref.Append("id"))
}

// Name returns a reference to field name of aws_memorydb_parameter_group.
func (mpg dataMemorydbParameterGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mpg.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_memorydb_parameter_group.
func (mpg dataMemorydbParameterGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mpg.ref.Append("tags"))
}

func (mpg dataMemorydbParameterGroupAttributes) Parameter() terra.SetValue[datamemorydbparametergroup.ParameterAttributes] {
	return terra.ReferenceAsSet[datamemorydbparametergroup.ParameterAttributes](mpg.ref.Append("parameter"))
}
