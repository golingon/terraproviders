// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataEcrImage creates a new instance of [DataEcrImage].
func NewDataEcrImage(name string, args DataEcrImageArgs) *DataEcrImage {
	return &DataEcrImage{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEcrImage)(nil)

// DataEcrImage represents the Terraform data resource aws_ecr_image.
type DataEcrImage struct {
	Name string
	Args DataEcrImageArgs
}

// DataSource returns the Terraform object type for [DataEcrImage].
func (ei *DataEcrImage) DataSource() string {
	return "aws_ecr_image"
}

// LocalName returns the local name for [DataEcrImage].
func (ei *DataEcrImage) LocalName() string {
	return ei.Name
}

// Configuration returns the configuration (args) for [DataEcrImage].
func (ei *DataEcrImage) Configuration() interface{} {
	return ei.Args
}

// Attributes returns the attributes for [DataEcrImage].
func (ei *DataEcrImage) Attributes() dataEcrImageAttributes {
	return dataEcrImageAttributes{ref: terra.ReferenceDataResource(ei)}
}

// DataEcrImageArgs contains the configurations for aws_ecr_image.
type DataEcrImageArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImageDigest: string, optional
	ImageDigest terra.StringValue `hcl:"image_digest,attr"`
	// ImageTag: string, optional
	ImageTag terra.StringValue `hcl:"image_tag,attr"`
	// MostRecent: bool, optional
	MostRecent terra.BoolValue `hcl:"most_recent,attr"`
	// RegistryId: string, optional
	RegistryId terra.StringValue `hcl:"registry_id,attr"`
	// RepositoryName: string, required
	RepositoryName terra.StringValue `hcl:"repository_name,attr" validate:"required"`
}
type dataEcrImageAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ecr_image.
func (ei dataEcrImageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("id"))
}

// ImageDigest returns a reference to field image_digest of aws_ecr_image.
func (ei dataEcrImageAttributes) ImageDigest() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("image_digest"))
}

// ImagePushedAt returns a reference to field image_pushed_at of aws_ecr_image.
func (ei dataEcrImageAttributes) ImagePushedAt() terra.NumberValue {
	return terra.ReferenceAsNumber(ei.ref.Append("image_pushed_at"))
}

// ImageSizeInBytes returns a reference to field image_size_in_bytes of aws_ecr_image.
func (ei dataEcrImageAttributes) ImageSizeInBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(ei.ref.Append("image_size_in_bytes"))
}

// ImageTag returns a reference to field image_tag of aws_ecr_image.
func (ei dataEcrImageAttributes) ImageTag() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("image_tag"))
}

// ImageTags returns a reference to field image_tags of aws_ecr_image.
func (ei dataEcrImageAttributes) ImageTags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ei.ref.Append("image_tags"))
}

// MostRecent returns a reference to field most_recent of aws_ecr_image.
func (ei dataEcrImageAttributes) MostRecent() terra.BoolValue {
	return terra.ReferenceAsBool(ei.ref.Append("most_recent"))
}

// RegistryId returns a reference to field registry_id of aws_ecr_image.
func (ei dataEcrImageAttributes) RegistryId() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("registry_id"))
}

// RepositoryName returns a reference to field repository_name of aws_ecr_image.
func (ei dataEcrImageAttributes) RepositoryName() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("repository_name"))
}
