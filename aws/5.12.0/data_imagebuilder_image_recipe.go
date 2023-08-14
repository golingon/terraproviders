// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataimagebuilderimagerecipe "github.com/golingon/terraproviders/aws/5.12.0/dataimagebuilderimagerecipe"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataImagebuilderImageRecipe creates a new instance of [DataImagebuilderImageRecipe].
func NewDataImagebuilderImageRecipe(name string, args DataImagebuilderImageRecipeArgs) *DataImagebuilderImageRecipe {
	return &DataImagebuilderImageRecipe{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataImagebuilderImageRecipe)(nil)

// DataImagebuilderImageRecipe represents the Terraform data resource aws_imagebuilder_image_recipe.
type DataImagebuilderImageRecipe struct {
	Name string
	Args DataImagebuilderImageRecipeArgs
}

// DataSource returns the Terraform object type for [DataImagebuilderImageRecipe].
func (iir *DataImagebuilderImageRecipe) DataSource() string {
	return "aws_imagebuilder_image_recipe"
}

// LocalName returns the local name for [DataImagebuilderImageRecipe].
func (iir *DataImagebuilderImageRecipe) LocalName() string {
	return iir.Name
}

// Configuration returns the configuration (args) for [DataImagebuilderImageRecipe].
func (iir *DataImagebuilderImageRecipe) Configuration() interface{} {
	return iir.Args
}

// Attributes returns the attributes for [DataImagebuilderImageRecipe].
func (iir *DataImagebuilderImageRecipe) Attributes() dataImagebuilderImageRecipeAttributes {
	return dataImagebuilderImageRecipeAttributes{ref: terra.ReferenceDataResource(iir)}
}

// DataImagebuilderImageRecipeArgs contains the configurations for aws_imagebuilder_image_recipe.
type DataImagebuilderImageRecipeArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// BlockDeviceMapping: min=0
	BlockDeviceMapping []dataimagebuilderimagerecipe.BlockDeviceMapping `hcl:"block_device_mapping,block" validate:"min=0"`
	// Component: min=0
	Component []dataimagebuilderimagerecipe.Component `hcl:"component,block" validate:"min=0"`
}
type dataImagebuilderImageRecipeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_imagebuilder_image_recipe.
func (iir dataImagebuilderImageRecipeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(iir.ref.Append("arn"))
}

// DateCreated returns a reference to field date_created of aws_imagebuilder_image_recipe.
func (iir dataImagebuilderImageRecipeAttributes) DateCreated() terra.StringValue {
	return terra.ReferenceAsString(iir.ref.Append("date_created"))
}

// Description returns a reference to field description of aws_imagebuilder_image_recipe.
func (iir dataImagebuilderImageRecipeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(iir.ref.Append("description"))
}

// Id returns a reference to field id of aws_imagebuilder_image_recipe.
func (iir dataImagebuilderImageRecipeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iir.ref.Append("id"))
}

// Name returns a reference to field name of aws_imagebuilder_image_recipe.
func (iir dataImagebuilderImageRecipeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iir.ref.Append("name"))
}

// Owner returns a reference to field owner of aws_imagebuilder_image_recipe.
func (iir dataImagebuilderImageRecipeAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(iir.ref.Append("owner"))
}

// ParentImage returns a reference to field parent_image of aws_imagebuilder_image_recipe.
func (iir dataImagebuilderImageRecipeAttributes) ParentImage() terra.StringValue {
	return terra.ReferenceAsString(iir.ref.Append("parent_image"))
}

// Platform returns a reference to field platform of aws_imagebuilder_image_recipe.
func (iir dataImagebuilderImageRecipeAttributes) Platform() terra.StringValue {
	return terra.ReferenceAsString(iir.ref.Append("platform"))
}

// Tags returns a reference to field tags of aws_imagebuilder_image_recipe.
func (iir dataImagebuilderImageRecipeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iir.ref.Append("tags"))
}

// UserDataBase64 returns a reference to field user_data_base64 of aws_imagebuilder_image_recipe.
func (iir dataImagebuilderImageRecipeAttributes) UserDataBase64() terra.StringValue {
	return terra.ReferenceAsString(iir.ref.Append("user_data_base64"))
}

// Version returns a reference to field version of aws_imagebuilder_image_recipe.
func (iir dataImagebuilderImageRecipeAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(iir.ref.Append("version"))
}

// WorkingDirectory returns a reference to field working_directory of aws_imagebuilder_image_recipe.
func (iir dataImagebuilderImageRecipeAttributes) WorkingDirectory() terra.StringValue {
	return terra.ReferenceAsString(iir.ref.Append("working_directory"))
}

func (iir dataImagebuilderImageRecipeAttributes) BlockDeviceMapping() terra.SetValue[dataimagebuilderimagerecipe.BlockDeviceMappingAttributes] {
	return terra.ReferenceAsSet[dataimagebuilderimagerecipe.BlockDeviceMappingAttributes](iir.ref.Append("block_device_mapping"))
}

func (iir dataImagebuilderImageRecipeAttributes) Component() terra.ListValue[dataimagebuilderimagerecipe.ComponentAttributes] {
	return terra.ReferenceAsList[dataimagebuilderimagerecipe.ComponentAttributes](iir.ref.Append("component"))
}
