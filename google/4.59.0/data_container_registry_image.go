// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataContainerRegistryImage creates a new instance of [DataContainerRegistryImage].
func NewDataContainerRegistryImage(name string, args DataContainerRegistryImageArgs) *DataContainerRegistryImage {
	return &DataContainerRegistryImage{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataContainerRegistryImage)(nil)

// DataContainerRegistryImage represents the Terraform data resource google_container_registry_image.
type DataContainerRegistryImage struct {
	Name string
	Args DataContainerRegistryImageArgs
}

// DataSource returns the Terraform object type for [DataContainerRegistryImage].
func (cri *DataContainerRegistryImage) DataSource() string {
	return "google_container_registry_image"
}

// LocalName returns the local name for [DataContainerRegistryImage].
func (cri *DataContainerRegistryImage) LocalName() string {
	return cri.Name
}

// Configuration returns the configuration (args) for [DataContainerRegistryImage].
func (cri *DataContainerRegistryImage) Configuration() interface{} {
	return cri.Args
}

// Attributes returns the attributes for [DataContainerRegistryImage].
func (cri *DataContainerRegistryImage) Attributes() dataContainerRegistryImageAttributes {
	return dataContainerRegistryImageAttributes{ref: terra.ReferenceDataResource(cri)}
}

// DataContainerRegistryImageArgs contains the configurations for google_container_registry_image.
type DataContainerRegistryImageArgs struct {
	// Digest: string, optional
	Digest terra.StringValue `hcl:"digest,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Tag: string, optional
	Tag terra.StringValue `hcl:"tag,attr"`
}
type dataContainerRegistryImageAttributes struct {
	ref terra.Reference
}

// Digest returns a reference to field digest of google_container_registry_image.
func (cri dataContainerRegistryImageAttributes) Digest() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("digest"))
}

// Id returns a reference to field id of google_container_registry_image.
func (cri dataContainerRegistryImageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("id"))
}

// ImageUrl returns a reference to field image_url of google_container_registry_image.
func (cri dataContainerRegistryImageAttributes) ImageUrl() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("image_url"))
}

// Name returns a reference to field name of google_container_registry_image.
func (cri dataContainerRegistryImageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("name"))
}

// Project returns a reference to field project of google_container_registry_image.
func (cri dataContainerRegistryImageAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("project"))
}

// Region returns a reference to field region of google_container_registry_image.
func (cri dataContainerRegistryImageAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("region"))
}

// Tag returns a reference to field tag of google_container_registry_image.
func (cri dataContainerRegistryImageAttributes) Tag() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("tag"))
}