// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataSagemakerPrebuiltEcrImage creates a new instance of [DataSagemakerPrebuiltEcrImage].
func NewDataSagemakerPrebuiltEcrImage(name string, args DataSagemakerPrebuiltEcrImageArgs) *DataSagemakerPrebuiltEcrImage {
	return &DataSagemakerPrebuiltEcrImage{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSagemakerPrebuiltEcrImage)(nil)

// DataSagemakerPrebuiltEcrImage represents the Terraform data resource aws_sagemaker_prebuilt_ecr_image.
type DataSagemakerPrebuiltEcrImage struct {
	Name string
	Args DataSagemakerPrebuiltEcrImageArgs
}

// DataSource returns the Terraform object type for [DataSagemakerPrebuiltEcrImage].
func (spei *DataSagemakerPrebuiltEcrImage) DataSource() string {
	return "aws_sagemaker_prebuilt_ecr_image"
}

// LocalName returns the local name for [DataSagemakerPrebuiltEcrImage].
func (spei *DataSagemakerPrebuiltEcrImage) LocalName() string {
	return spei.Name
}

// Configuration returns the configuration (args) for [DataSagemakerPrebuiltEcrImage].
func (spei *DataSagemakerPrebuiltEcrImage) Configuration() interface{} {
	return spei.Args
}

// Attributes returns the attributes for [DataSagemakerPrebuiltEcrImage].
func (spei *DataSagemakerPrebuiltEcrImage) Attributes() dataSagemakerPrebuiltEcrImageAttributes {
	return dataSagemakerPrebuiltEcrImageAttributes{ref: terra.ReferenceDataResource(spei)}
}

// DataSagemakerPrebuiltEcrImageArgs contains the configurations for aws_sagemaker_prebuilt_ecr_image.
type DataSagemakerPrebuiltEcrImageArgs struct {
	// DnsSuffix: string, optional
	DnsSuffix terra.StringValue `hcl:"dns_suffix,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImageTag: string, optional
	ImageTag terra.StringValue `hcl:"image_tag,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// RepositoryName: string, required
	RepositoryName terra.StringValue `hcl:"repository_name,attr" validate:"required"`
}
type dataSagemakerPrebuiltEcrImageAttributes struct {
	ref terra.Reference
}

// DnsSuffix returns a reference to field dns_suffix of aws_sagemaker_prebuilt_ecr_image.
func (spei dataSagemakerPrebuiltEcrImageAttributes) DnsSuffix() terra.StringValue {
	return terra.ReferenceAsString(spei.ref.Append("dns_suffix"))
}

// Id returns a reference to field id of aws_sagemaker_prebuilt_ecr_image.
func (spei dataSagemakerPrebuiltEcrImageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spei.ref.Append("id"))
}

// ImageTag returns a reference to field image_tag of aws_sagemaker_prebuilt_ecr_image.
func (spei dataSagemakerPrebuiltEcrImageAttributes) ImageTag() terra.StringValue {
	return terra.ReferenceAsString(spei.ref.Append("image_tag"))
}

// Region returns a reference to field region of aws_sagemaker_prebuilt_ecr_image.
func (spei dataSagemakerPrebuiltEcrImageAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(spei.ref.Append("region"))
}

// RegistryId returns a reference to field registry_id of aws_sagemaker_prebuilt_ecr_image.
func (spei dataSagemakerPrebuiltEcrImageAttributes) RegistryId() terra.StringValue {
	return terra.ReferenceAsString(spei.ref.Append("registry_id"))
}

// RegistryPath returns a reference to field registry_path of aws_sagemaker_prebuilt_ecr_image.
func (spei dataSagemakerPrebuiltEcrImageAttributes) RegistryPath() terra.StringValue {
	return terra.ReferenceAsString(spei.ref.Append("registry_path"))
}

// RepositoryName returns a reference to field repository_name of aws_sagemaker_prebuilt_ecr_image.
func (spei dataSagemakerPrebuiltEcrImageAttributes) RepositoryName() terra.StringValue {
	return terra.ReferenceAsString(spei.ref.Append("repository_name"))
}