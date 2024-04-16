// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_imagebuilder_container_recipe

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_imagebuilder_container_recipe.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aicr *DataSource) DataSource() string {
	return "aws_imagebuilder_container_recipe"
}

// LocalName returns the local name for [DataSource].
func (aicr *DataSource) LocalName() string {
	return aicr.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aicr *DataSource) Configuration() interface{} {
	return aicr.Args
}

// Attributes returns the attributes for [DataSource].
func (aicr *DataSource) Attributes() dataAwsImagebuilderContainerRecipeAttributes {
	return dataAwsImagebuilderContainerRecipeAttributes{ref: terra.ReferenceDataSource(aicr)}
}

// DataArgs contains the configurations for aws_imagebuilder_container_recipe.
type DataArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}

type dataAwsImagebuilderContainerRecipeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_imagebuilder_container_recipe.
func (aicr dataAwsImagebuilderContainerRecipeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aicr.ref.Append("arn"))
}

// ContainerType returns a reference to field container_type of aws_imagebuilder_container_recipe.
func (aicr dataAwsImagebuilderContainerRecipeAttributes) ContainerType() terra.StringValue {
	return terra.ReferenceAsString(aicr.ref.Append("container_type"))
}

// DateCreated returns a reference to field date_created of aws_imagebuilder_container_recipe.
func (aicr dataAwsImagebuilderContainerRecipeAttributes) DateCreated() terra.StringValue {
	return terra.ReferenceAsString(aicr.ref.Append("date_created"))
}

// Description returns a reference to field description of aws_imagebuilder_container_recipe.
func (aicr dataAwsImagebuilderContainerRecipeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aicr.ref.Append("description"))
}

// DockerfileTemplateData returns a reference to field dockerfile_template_data of aws_imagebuilder_container_recipe.
func (aicr dataAwsImagebuilderContainerRecipeAttributes) DockerfileTemplateData() terra.StringValue {
	return terra.ReferenceAsString(aicr.ref.Append("dockerfile_template_data"))
}

// Encrypted returns a reference to field encrypted of aws_imagebuilder_container_recipe.
func (aicr dataAwsImagebuilderContainerRecipeAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(aicr.ref.Append("encrypted"))
}

// Id returns a reference to field id of aws_imagebuilder_container_recipe.
func (aicr dataAwsImagebuilderContainerRecipeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aicr.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_imagebuilder_container_recipe.
func (aicr dataAwsImagebuilderContainerRecipeAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(aicr.ref.Append("kms_key_id"))
}

// Name returns a reference to field name of aws_imagebuilder_container_recipe.
func (aicr dataAwsImagebuilderContainerRecipeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aicr.ref.Append("name"))
}

// Owner returns a reference to field owner of aws_imagebuilder_container_recipe.
func (aicr dataAwsImagebuilderContainerRecipeAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(aicr.ref.Append("owner"))
}

// ParentImage returns a reference to field parent_image of aws_imagebuilder_container_recipe.
func (aicr dataAwsImagebuilderContainerRecipeAttributes) ParentImage() terra.StringValue {
	return terra.ReferenceAsString(aicr.ref.Append("parent_image"))
}

// Platform returns a reference to field platform of aws_imagebuilder_container_recipe.
func (aicr dataAwsImagebuilderContainerRecipeAttributes) Platform() terra.StringValue {
	return terra.ReferenceAsString(aicr.ref.Append("platform"))
}

// Tags returns a reference to field tags of aws_imagebuilder_container_recipe.
func (aicr dataAwsImagebuilderContainerRecipeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aicr.ref.Append("tags"))
}

// Version returns a reference to field version of aws_imagebuilder_container_recipe.
func (aicr dataAwsImagebuilderContainerRecipeAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(aicr.ref.Append("version"))
}

// WorkingDirectory returns a reference to field working_directory of aws_imagebuilder_container_recipe.
func (aicr dataAwsImagebuilderContainerRecipeAttributes) WorkingDirectory() terra.StringValue {
	return terra.ReferenceAsString(aicr.ref.Append("working_directory"))
}

func (aicr dataAwsImagebuilderContainerRecipeAttributes) Component() terra.ListValue[DataComponentAttributes] {
	return terra.ReferenceAsList[DataComponentAttributes](aicr.ref.Append("component"))
}

func (aicr dataAwsImagebuilderContainerRecipeAttributes) InstanceConfiguration() terra.ListValue[DataInstanceConfigurationAttributes] {
	return terra.ReferenceAsList[DataInstanceConfigurationAttributes](aicr.ref.Append("instance_configuration"))
}

func (aicr dataAwsImagebuilderContainerRecipeAttributes) TargetRepository() terra.ListValue[DataTargetRepositoryAttributes] {
	return terra.ReferenceAsList[DataTargetRepositoryAttributes](aicr.ref.Append("target_repository"))
}
