// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_imagebuilder_image_recipe

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_imagebuilder_image_recipe.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aiir *DataSource) DataSource() string {
	return "aws_imagebuilder_image_recipe"
}

// LocalName returns the local name for [DataSource].
func (aiir *DataSource) LocalName() string {
	return aiir.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aiir *DataSource) Configuration() interface{} {
	return aiir.Args
}

// Attributes returns the attributes for [DataSource].
func (aiir *DataSource) Attributes() dataAwsImagebuilderImageRecipeAttributes {
	return dataAwsImagebuilderImageRecipeAttributes{ref: terra.ReferenceDataSource(aiir)}
}

// DataArgs contains the configurations for aws_imagebuilder_image_recipe.
type DataArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}

type dataAwsImagebuilderImageRecipeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_imagebuilder_image_recipe.
func (aiir dataAwsImagebuilderImageRecipeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aiir.ref.Append("arn"))
}

// DateCreated returns a reference to field date_created of aws_imagebuilder_image_recipe.
func (aiir dataAwsImagebuilderImageRecipeAttributes) DateCreated() terra.StringValue {
	return terra.ReferenceAsString(aiir.ref.Append("date_created"))
}

// Description returns a reference to field description of aws_imagebuilder_image_recipe.
func (aiir dataAwsImagebuilderImageRecipeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aiir.ref.Append("description"))
}

// Id returns a reference to field id of aws_imagebuilder_image_recipe.
func (aiir dataAwsImagebuilderImageRecipeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aiir.ref.Append("id"))
}

// Name returns a reference to field name of aws_imagebuilder_image_recipe.
func (aiir dataAwsImagebuilderImageRecipeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aiir.ref.Append("name"))
}

// Owner returns a reference to field owner of aws_imagebuilder_image_recipe.
func (aiir dataAwsImagebuilderImageRecipeAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(aiir.ref.Append("owner"))
}

// ParentImage returns a reference to field parent_image of aws_imagebuilder_image_recipe.
func (aiir dataAwsImagebuilderImageRecipeAttributes) ParentImage() terra.StringValue {
	return terra.ReferenceAsString(aiir.ref.Append("parent_image"))
}

// Platform returns a reference to field platform of aws_imagebuilder_image_recipe.
func (aiir dataAwsImagebuilderImageRecipeAttributes) Platform() terra.StringValue {
	return terra.ReferenceAsString(aiir.ref.Append("platform"))
}

// Tags returns a reference to field tags of aws_imagebuilder_image_recipe.
func (aiir dataAwsImagebuilderImageRecipeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aiir.ref.Append("tags"))
}

// UserDataBase64 returns a reference to field user_data_base64 of aws_imagebuilder_image_recipe.
func (aiir dataAwsImagebuilderImageRecipeAttributes) UserDataBase64() terra.StringValue {
	return terra.ReferenceAsString(aiir.ref.Append("user_data_base64"))
}

// Version returns a reference to field version of aws_imagebuilder_image_recipe.
func (aiir dataAwsImagebuilderImageRecipeAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(aiir.ref.Append("version"))
}

// WorkingDirectory returns a reference to field working_directory of aws_imagebuilder_image_recipe.
func (aiir dataAwsImagebuilderImageRecipeAttributes) WorkingDirectory() terra.StringValue {
	return terra.ReferenceAsString(aiir.ref.Append("working_directory"))
}

func (aiir dataAwsImagebuilderImageRecipeAttributes) BlockDeviceMapping() terra.SetValue[DataBlockDeviceMappingAttributes] {
	return terra.ReferenceAsSet[DataBlockDeviceMappingAttributes](aiir.ref.Append("block_device_mapping"))
}

func (aiir dataAwsImagebuilderImageRecipeAttributes) Component() terra.ListValue[DataComponentAttributes] {
	return terra.ReferenceAsList[DataComponentAttributes](aiir.ref.Append("component"))
}
