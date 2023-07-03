// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataimagebuilderimagepipelines "github.com/golingon/terraproviders/aws/5.6.2/dataimagebuilderimagepipelines"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataImagebuilderImagePipelines creates a new instance of [DataImagebuilderImagePipelines].
func NewDataImagebuilderImagePipelines(name string, args DataImagebuilderImagePipelinesArgs) *DataImagebuilderImagePipelines {
	return &DataImagebuilderImagePipelines{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataImagebuilderImagePipelines)(nil)

// DataImagebuilderImagePipelines represents the Terraform data resource aws_imagebuilder_image_pipelines.
type DataImagebuilderImagePipelines struct {
	Name string
	Args DataImagebuilderImagePipelinesArgs
}

// DataSource returns the Terraform object type for [DataImagebuilderImagePipelines].
func (iip *DataImagebuilderImagePipelines) DataSource() string {
	return "aws_imagebuilder_image_pipelines"
}

// LocalName returns the local name for [DataImagebuilderImagePipelines].
func (iip *DataImagebuilderImagePipelines) LocalName() string {
	return iip.Name
}

// Configuration returns the configuration (args) for [DataImagebuilderImagePipelines].
func (iip *DataImagebuilderImagePipelines) Configuration() interface{} {
	return iip.Args
}

// Attributes returns the attributes for [DataImagebuilderImagePipelines].
func (iip *DataImagebuilderImagePipelines) Attributes() dataImagebuilderImagePipelinesAttributes {
	return dataImagebuilderImagePipelinesAttributes{ref: terra.ReferenceDataResource(iip)}
}

// DataImagebuilderImagePipelinesArgs contains the configurations for aws_imagebuilder_image_pipelines.
type DataImagebuilderImagePipelinesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Filter: min=0
	Filter []dataimagebuilderimagepipelines.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataImagebuilderImagePipelinesAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_imagebuilder_image_pipelines.
func (iip dataImagebuilderImagePipelinesAttributes) Arns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iip.ref.Append("arns"))
}

// Id returns a reference to field id of aws_imagebuilder_image_pipelines.
func (iip dataImagebuilderImagePipelinesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("id"))
}

// Names returns a reference to field names of aws_imagebuilder_image_pipelines.
func (iip dataImagebuilderImagePipelinesAttributes) Names() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iip.ref.Append("names"))
}

func (iip dataImagebuilderImagePipelinesAttributes) Filter() terra.SetValue[dataimagebuilderimagepipelines.FilterAttributes] {
	return terra.ReferenceAsSet[dataimagebuilderimagepipelines.FilterAttributes](iip.ref.Append("filter"))
}
