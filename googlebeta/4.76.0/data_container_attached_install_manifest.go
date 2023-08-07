// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataContainerAttachedInstallManifest creates a new instance of [DataContainerAttachedInstallManifest].
func NewDataContainerAttachedInstallManifest(name string, args DataContainerAttachedInstallManifestArgs) *DataContainerAttachedInstallManifest {
	return &DataContainerAttachedInstallManifest{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataContainerAttachedInstallManifest)(nil)

// DataContainerAttachedInstallManifest represents the Terraform data resource google_container_attached_install_manifest.
type DataContainerAttachedInstallManifest struct {
	Name string
	Args DataContainerAttachedInstallManifestArgs
}

// DataSource returns the Terraform object type for [DataContainerAttachedInstallManifest].
func (caim *DataContainerAttachedInstallManifest) DataSource() string {
	return "google_container_attached_install_manifest"
}

// LocalName returns the local name for [DataContainerAttachedInstallManifest].
func (caim *DataContainerAttachedInstallManifest) LocalName() string {
	return caim.Name
}

// Configuration returns the configuration (args) for [DataContainerAttachedInstallManifest].
func (caim *DataContainerAttachedInstallManifest) Configuration() interface{} {
	return caim.Args
}

// Attributes returns the attributes for [DataContainerAttachedInstallManifest].
func (caim *DataContainerAttachedInstallManifest) Attributes() dataContainerAttachedInstallManifestAttributes {
	return dataContainerAttachedInstallManifestAttributes{ref: terra.ReferenceDataResource(caim)}
}

// DataContainerAttachedInstallManifestArgs contains the configurations for google_container_attached_install_manifest.
type DataContainerAttachedInstallManifestArgs struct {
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// PlatformVersion: string, required
	PlatformVersion terra.StringValue `hcl:"platform_version,attr" validate:"required"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
}
type dataContainerAttachedInstallManifestAttributes struct {
	ref terra.Reference
}

// ClusterId returns a reference to field cluster_id of google_container_attached_install_manifest.
func (caim dataContainerAttachedInstallManifestAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(caim.ref.Append("cluster_id"))
}

// Id returns a reference to field id of google_container_attached_install_manifest.
func (caim dataContainerAttachedInstallManifestAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(caim.ref.Append("id"))
}

// Location returns a reference to field location of google_container_attached_install_manifest.
func (caim dataContainerAttachedInstallManifestAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(caim.ref.Append("location"))
}

// Manifest returns a reference to field manifest of google_container_attached_install_manifest.
func (caim dataContainerAttachedInstallManifestAttributes) Manifest() terra.StringValue {
	return terra.ReferenceAsString(caim.ref.Append("manifest"))
}

// PlatformVersion returns a reference to field platform_version of google_container_attached_install_manifest.
func (caim dataContainerAttachedInstallManifestAttributes) PlatformVersion() terra.StringValue {
	return terra.ReferenceAsString(caim.ref.Append("platform_version"))
}

// Project returns a reference to field project of google_container_attached_install_manifest.
func (caim dataContainerAttachedInstallManifestAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(caim.ref.Append("project"))
}
