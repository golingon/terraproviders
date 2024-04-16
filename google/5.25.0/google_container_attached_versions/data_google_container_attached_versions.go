// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_container_attached_versions

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_container_attached_versions.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gcav *DataSource) DataSource() string {
	return "google_container_attached_versions"
}

// LocalName returns the local name for [DataSource].
func (gcav *DataSource) LocalName() string {
	return gcav.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gcav *DataSource) Configuration() interface{} {
	return gcav.Args
}

// Attributes returns the attributes for [DataSource].
func (gcav *DataSource) Attributes() dataGoogleContainerAttachedVersionsAttributes {
	return dataGoogleContainerAttachedVersionsAttributes{ref: terra.ReferenceDataSource(gcav)}
}

// DataArgs contains the configurations for google_container_attached_versions.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
}

type dataGoogleContainerAttachedVersionsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_container_attached_versions.
func (gcav dataGoogleContainerAttachedVersionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcav.ref.Append("id"))
}

// Location returns a reference to field location of google_container_attached_versions.
func (gcav dataGoogleContainerAttachedVersionsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gcav.ref.Append("location"))
}

// Project returns a reference to field project of google_container_attached_versions.
func (gcav dataGoogleContainerAttachedVersionsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcav.ref.Append("project"))
}

// ValidVersions returns a reference to field valid_versions of google_container_attached_versions.
func (gcav dataGoogleContainerAttachedVersionsAttributes) ValidVersions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gcav.ref.Append("valid_versions"))
}
