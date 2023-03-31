// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataecrrepository "github.com/golingon/terraproviders/aws/4.60.0/dataecrrepository"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEcrRepository creates a new instance of [DataEcrRepository].
func NewDataEcrRepository(name string, args DataEcrRepositoryArgs) *DataEcrRepository {
	return &DataEcrRepository{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEcrRepository)(nil)

// DataEcrRepository represents the Terraform data resource aws_ecr_repository.
type DataEcrRepository struct {
	Name string
	Args DataEcrRepositoryArgs
}

// DataSource returns the Terraform object type for [DataEcrRepository].
func (er *DataEcrRepository) DataSource() string {
	return "aws_ecr_repository"
}

// LocalName returns the local name for [DataEcrRepository].
func (er *DataEcrRepository) LocalName() string {
	return er.Name
}

// Configuration returns the configuration (args) for [DataEcrRepository].
func (er *DataEcrRepository) Configuration() interface{} {
	return er.Args
}

// Attributes returns the attributes for [DataEcrRepository].
func (er *DataEcrRepository) Attributes() dataEcrRepositoryAttributes {
	return dataEcrRepositoryAttributes{ref: terra.ReferenceDataResource(er)}
}

// DataEcrRepositoryArgs contains the configurations for aws_ecr_repository.
type DataEcrRepositoryArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RegistryId: string, optional
	RegistryId terra.StringValue `hcl:"registry_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// EncryptionConfiguration: min=0
	EncryptionConfiguration []dataecrrepository.EncryptionConfiguration `hcl:"encryption_configuration,block" validate:"min=0"`
	// ImageScanningConfiguration: min=0
	ImageScanningConfiguration []dataecrrepository.ImageScanningConfiguration `hcl:"image_scanning_configuration,block" validate:"min=0"`
}
type dataEcrRepositoryAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ecr_repository.
func (er dataEcrRepositoryAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("arn"))
}

// Id returns a reference to field id of aws_ecr_repository.
func (er dataEcrRepositoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("id"))
}

// ImageTagMutability returns a reference to field image_tag_mutability of aws_ecr_repository.
func (er dataEcrRepositoryAttributes) ImageTagMutability() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("image_tag_mutability"))
}

// MostRecentImageTags returns a reference to field most_recent_image_tags of aws_ecr_repository.
func (er dataEcrRepositoryAttributes) MostRecentImageTags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](er.ref.Append("most_recent_image_tags"))
}

// Name returns a reference to field name of aws_ecr_repository.
func (er dataEcrRepositoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("name"))
}

// RegistryId returns a reference to field registry_id of aws_ecr_repository.
func (er dataEcrRepositoryAttributes) RegistryId() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("registry_id"))
}

// RepositoryUrl returns a reference to field repository_url of aws_ecr_repository.
func (er dataEcrRepositoryAttributes) RepositoryUrl() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("repository_url"))
}

// Tags returns a reference to field tags of aws_ecr_repository.
func (er dataEcrRepositoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](er.ref.Append("tags"))
}

func (er dataEcrRepositoryAttributes) EncryptionConfiguration() terra.ListValue[dataecrrepository.EncryptionConfigurationAttributes] {
	return terra.ReferenceAsList[dataecrrepository.EncryptionConfigurationAttributes](er.ref.Append("encryption_configuration"))
}

func (er dataEcrRepositoryAttributes) ImageScanningConfiguration() terra.ListValue[dataecrrepository.ImageScanningConfigurationAttributes] {
	return terra.ReferenceAsList[dataecrrepository.ImageScanningConfigurationAttributes](er.ref.Append("image_scanning_configuration"))
}
