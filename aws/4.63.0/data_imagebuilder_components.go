// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataimagebuildercomponents "github.com/golingon/terraproviders/aws/4.63.0/dataimagebuildercomponents"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataImagebuilderComponents creates a new instance of [DataImagebuilderComponents].
func NewDataImagebuilderComponents(name string, args DataImagebuilderComponentsArgs) *DataImagebuilderComponents {
	return &DataImagebuilderComponents{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataImagebuilderComponents)(nil)

// DataImagebuilderComponents represents the Terraform data resource aws_imagebuilder_components.
type DataImagebuilderComponents struct {
	Name string
	Args DataImagebuilderComponentsArgs
}

// DataSource returns the Terraform object type for [DataImagebuilderComponents].
func (ic *DataImagebuilderComponents) DataSource() string {
	return "aws_imagebuilder_components"
}

// LocalName returns the local name for [DataImagebuilderComponents].
func (ic *DataImagebuilderComponents) LocalName() string {
	return ic.Name
}

// Configuration returns the configuration (args) for [DataImagebuilderComponents].
func (ic *DataImagebuilderComponents) Configuration() interface{} {
	return ic.Args
}

// Attributes returns the attributes for [DataImagebuilderComponents].
func (ic *DataImagebuilderComponents) Attributes() dataImagebuilderComponentsAttributes {
	return dataImagebuilderComponentsAttributes{ref: terra.ReferenceDataResource(ic)}
}

// DataImagebuilderComponentsArgs contains the configurations for aws_imagebuilder_components.
type DataImagebuilderComponentsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Owner: string, optional
	Owner terra.StringValue `hcl:"owner,attr"`
	// Filter: min=0
	Filter []dataimagebuildercomponents.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataImagebuilderComponentsAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_imagebuilder_components.
func (ic dataImagebuilderComponentsAttributes) Arns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ic.ref.Append("arns"))
}

// Id returns a reference to field id of aws_imagebuilder_components.
func (ic dataImagebuilderComponentsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("id"))
}

// Names returns a reference to field names of aws_imagebuilder_components.
func (ic dataImagebuilderComponentsAttributes) Names() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ic.ref.Append("names"))
}

// Owner returns a reference to field owner of aws_imagebuilder_components.
func (ic dataImagebuilderComponentsAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("owner"))
}

func (ic dataImagebuilderComponentsAttributes) Filter() terra.SetValue[dataimagebuildercomponents.FilterAttributes] {
	return terra.ReferenceAsSet[dataimagebuildercomponents.FilterAttributes](ic.ref.Append("filter"))
}
