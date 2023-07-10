// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataContainerAttachedVersions creates a new instance of [DataContainerAttachedVersions].
func NewDataContainerAttachedVersions(name string, args DataContainerAttachedVersionsArgs) *DataContainerAttachedVersions {
	return &DataContainerAttachedVersions{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataContainerAttachedVersions)(nil)

// DataContainerAttachedVersions represents the Terraform data resource google_container_attached_versions.
type DataContainerAttachedVersions struct {
	Name string
	Args DataContainerAttachedVersionsArgs
}

// DataSource returns the Terraform object type for [DataContainerAttachedVersions].
func (cav *DataContainerAttachedVersions) DataSource() string {
	return "google_container_attached_versions"
}

// LocalName returns the local name for [DataContainerAttachedVersions].
func (cav *DataContainerAttachedVersions) LocalName() string {
	return cav.Name
}

// Configuration returns the configuration (args) for [DataContainerAttachedVersions].
func (cav *DataContainerAttachedVersions) Configuration() interface{} {
	return cav.Args
}

// Attributes returns the attributes for [DataContainerAttachedVersions].
func (cav *DataContainerAttachedVersions) Attributes() dataContainerAttachedVersionsAttributes {
	return dataContainerAttachedVersionsAttributes{ref: terra.ReferenceDataResource(cav)}
}

// DataContainerAttachedVersionsArgs contains the configurations for google_container_attached_versions.
type DataContainerAttachedVersionsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
}
type dataContainerAttachedVersionsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_container_attached_versions.
func (cav dataContainerAttachedVersionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cav.ref.Append("id"))
}

// Location returns a reference to field location of google_container_attached_versions.
func (cav dataContainerAttachedVersionsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cav.ref.Append("location"))
}

// Project returns a reference to field project of google_container_attached_versions.
func (cav dataContainerAttachedVersionsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cav.ref.Append("project"))
}

// ValidVersions returns a reference to field valid_versions of google_container_attached_versions.
func (cav dataContainerAttachedVersionsAttributes) ValidVersions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cav.ref.Append("valid_versions"))
}
