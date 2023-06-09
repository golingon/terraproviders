// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataimagebuildercontainerrecipe "github.com/golingon/terraproviders/aws/5.6.2/dataimagebuildercontainerrecipe"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataImagebuilderContainerRecipe creates a new instance of [DataImagebuilderContainerRecipe].
func NewDataImagebuilderContainerRecipe(name string, args DataImagebuilderContainerRecipeArgs) *DataImagebuilderContainerRecipe {
	return &DataImagebuilderContainerRecipe{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataImagebuilderContainerRecipe)(nil)

// DataImagebuilderContainerRecipe represents the Terraform data resource aws_imagebuilder_container_recipe.
type DataImagebuilderContainerRecipe struct {
	Name string
	Args DataImagebuilderContainerRecipeArgs
}

// DataSource returns the Terraform object type for [DataImagebuilderContainerRecipe].
func (icr *DataImagebuilderContainerRecipe) DataSource() string {
	return "aws_imagebuilder_container_recipe"
}

// LocalName returns the local name for [DataImagebuilderContainerRecipe].
func (icr *DataImagebuilderContainerRecipe) LocalName() string {
	return icr.Name
}

// Configuration returns the configuration (args) for [DataImagebuilderContainerRecipe].
func (icr *DataImagebuilderContainerRecipe) Configuration() interface{} {
	return icr.Args
}

// Attributes returns the attributes for [DataImagebuilderContainerRecipe].
func (icr *DataImagebuilderContainerRecipe) Attributes() dataImagebuilderContainerRecipeAttributes {
	return dataImagebuilderContainerRecipeAttributes{ref: terra.ReferenceDataResource(icr)}
}

// DataImagebuilderContainerRecipeArgs contains the configurations for aws_imagebuilder_container_recipe.
type DataImagebuilderContainerRecipeArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Component: min=0
	Component []dataimagebuildercontainerrecipe.Component `hcl:"component,block" validate:"min=0"`
	// InstanceConfiguration: min=0
	InstanceConfiguration []dataimagebuildercontainerrecipe.InstanceConfiguration `hcl:"instance_configuration,block" validate:"min=0"`
	// TargetRepository: min=0
	TargetRepository []dataimagebuildercontainerrecipe.TargetRepository `hcl:"target_repository,block" validate:"min=0"`
}
type dataImagebuilderContainerRecipeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_imagebuilder_container_recipe.
func (icr dataImagebuilderContainerRecipeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("arn"))
}

// ContainerType returns a reference to field container_type of aws_imagebuilder_container_recipe.
func (icr dataImagebuilderContainerRecipeAttributes) ContainerType() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("container_type"))
}

// DateCreated returns a reference to field date_created of aws_imagebuilder_container_recipe.
func (icr dataImagebuilderContainerRecipeAttributes) DateCreated() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("date_created"))
}

// Description returns a reference to field description of aws_imagebuilder_container_recipe.
func (icr dataImagebuilderContainerRecipeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("description"))
}

// DockerfileTemplateData returns a reference to field dockerfile_template_data of aws_imagebuilder_container_recipe.
func (icr dataImagebuilderContainerRecipeAttributes) DockerfileTemplateData() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("dockerfile_template_data"))
}

// Encrypted returns a reference to field encrypted of aws_imagebuilder_container_recipe.
func (icr dataImagebuilderContainerRecipeAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(icr.ref.Append("encrypted"))
}

// Id returns a reference to field id of aws_imagebuilder_container_recipe.
func (icr dataImagebuilderContainerRecipeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_imagebuilder_container_recipe.
func (icr dataImagebuilderContainerRecipeAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("kms_key_id"))
}

// Name returns a reference to field name of aws_imagebuilder_container_recipe.
func (icr dataImagebuilderContainerRecipeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("name"))
}

// Owner returns a reference to field owner of aws_imagebuilder_container_recipe.
func (icr dataImagebuilderContainerRecipeAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("owner"))
}

// ParentImage returns a reference to field parent_image of aws_imagebuilder_container_recipe.
func (icr dataImagebuilderContainerRecipeAttributes) ParentImage() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("parent_image"))
}

// Platform returns a reference to field platform of aws_imagebuilder_container_recipe.
func (icr dataImagebuilderContainerRecipeAttributes) Platform() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("platform"))
}

// Tags returns a reference to field tags of aws_imagebuilder_container_recipe.
func (icr dataImagebuilderContainerRecipeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](icr.ref.Append("tags"))
}

// Version returns a reference to field version of aws_imagebuilder_container_recipe.
func (icr dataImagebuilderContainerRecipeAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("version"))
}

// WorkingDirectory returns a reference to field working_directory of aws_imagebuilder_container_recipe.
func (icr dataImagebuilderContainerRecipeAttributes) WorkingDirectory() terra.StringValue {
	return terra.ReferenceAsString(icr.ref.Append("working_directory"))
}

func (icr dataImagebuilderContainerRecipeAttributes) Component() terra.ListValue[dataimagebuildercontainerrecipe.ComponentAttributes] {
	return terra.ReferenceAsList[dataimagebuildercontainerrecipe.ComponentAttributes](icr.ref.Append("component"))
}

func (icr dataImagebuilderContainerRecipeAttributes) InstanceConfiguration() terra.ListValue[dataimagebuildercontainerrecipe.InstanceConfigurationAttributes] {
	return terra.ReferenceAsList[dataimagebuildercontainerrecipe.InstanceConfigurationAttributes](icr.ref.Append("instance_configuration"))
}

func (icr dataImagebuilderContainerRecipeAttributes) TargetRepository() terra.ListValue[dataimagebuildercontainerrecipe.TargetRepositoryAttributes] {
	return terra.ReferenceAsList[dataimagebuildercontainerrecipe.TargetRepositoryAttributes](icr.ref.Append("target_repository"))
}
