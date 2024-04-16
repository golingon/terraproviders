// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_imagebuilder_image_recipes

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_imagebuilder_image_recipes.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aiir *DataSource) DataSource() string {
	return "aws_imagebuilder_image_recipes"
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
func (aiir *DataSource) Attributes() dataAwsImagebuilderImageRecipesAttributes {
	return dataAwsImagebuilderImageRecipesAttributes{ref: terra.ReferenceDataSource(aiir)}
}

// DataArgs contains the configurations for aws_imagebuilder_image_recipes.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Owner: string, optional
	Owner terra.StringValue `hcl:"owner,attr"`
	// Filter: min=0
	Filter []DataFilter `hcl:"filter,block" validate:"min=0"`
}

type dataAwsImagebuilderImageRecipesAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_imagebuilder_image_recipes.
func (aiir dataAwsImagebuilderImageRecipesAttributes) Arns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aiir.ref.Append("arns"))
}

// Id returns a reference to field id of aws_imagebuilder_image_recipes.
func (aiir dataAwsImagebuilderImageRecipesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aiir.ref.Append("id"))
}

// Names returns a reference to field names of aws_imagebuilder_image_recipes.
func (aiir dataAwsImagebuilderImageRecipesAttributes) Names() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aiir.ref.Append("names"))
}

// Owner returns a reference to field owner of aws_imagebuilder_image_recipes.
func (aiir dataAwsImagebuilderImageRecipesAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(aiir.ref.Append("owner"))
}

func (aiir dataAwsImagebuilderImageRecipesAttributes) Filter() terra.SetValue[DataFilterAttributes] {
	return terra.ReferenceAsSet[DataFilterAttributes](aiir.ref.Append("filter"))
}
