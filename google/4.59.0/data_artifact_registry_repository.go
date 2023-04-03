// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	dataartifactregistryrepository "github.com/golingon/terraproviders/google/4.59.0/dataartifactregistryrepository"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataArtifactRegistryRepository creates a new instance of [DataArtifactRegistryRepository].
func NewDataArtifactRegistryRepository(name string, args DataArtifactRegistryRepositoryArgs) *DataArtifactRegistryRepository {
	return &DataArtifactRegistryRepository{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataArtifactRegistryRepository)(nil)

// DataArtifactRegistryRepository represents the Terraform data resource google_artifact_registry_repository.
type DataArtifactRegistryRepository struct {
	Name string
	Args DataArtifactRegistryRepositoryArgs
}

// DataSource returns the Terraform object type for [DataArtifactRegistryRepository].
func (arr *DataArtifactRegistryRepository) DataSource() string {
	return "google_artifact_registry_repository"
}

// LocalName returns the local name for [DataArtifactRegistryRepository].
func (arr *DataArtifactRegistryRepository) LocalName() string {
	return arr.Name
}

// Configuration returns the configuration (args) for [DataArtifactRegistryRepository].
func (arr *DataArtifactRegistryRepository) Configuration() interface{} {
	return arr.Args
}

// Attributes returns the attributes for [DataArtifactRegistryRepository].
func (arr *DataArtifactRegistryRepository) Attributes() dataArtifactRegistryRepositoryAttributes {
	return dataArtifactRegistryRepositoryAttributes{ref: terra.ReferenceDataResource(arr)}
}

// DataArtifactRegistryRepositoryArgs contains the configurations for google_artifact_registry_repository.
type DataArtifactRegistryRepositoryArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RepositoryId: string, required
	RepositoryId terra.StringValue `hcl:"repository_id,attr" validate:"required"`
	// MavenConfig: min=0
	MavenConfig []dataartifactregistryrepository.MavenConfig `hcl:"maven_config,block" validate:"min=0"`
}
type dataArtifactRegistryRepositoryAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_artifact_registry_repository.
func (arr dataArtifactRegistryRepositoryAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("create_time"))
}

// Description returns a reference to field description of google_artifact_registry_repository.
func (arr dataArtifactRegistryRepositoryAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("description"))
}

// Format returns a reference to field format of google_artifact_registry_repository.
func (arr dataArtifactRegistryRepositoryAttributes) Format() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("format"))
}

// Id returns a reference to field id of google_artifact_registry_repository.
func (arr dataArtifactRegistryRepositoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("id"))
}

// KmsKeyName returns a reference to field kms_key_name of google_artifact_registry_repository.
func (arr dataArtifactRegistryRepositoryAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("kms_key_name"))
}

// Labels returns a reference to field labels of google_artifact_registry_repository.
func (arr dataArtifactRegistryRepositoryAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](arr.ref.Append("labels"))
}

// Location returns a reference to field location of google_artifact_registry_repository.
func (arr dataArtifactRegistryRepositoryAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("location"))
}

// Name returns a reference to field name of google_artifact_registry_repository.
func (arr dataArtifactRegistryRepositoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("name"))
}

// Project returns a reference to field project of google_artifact_registry_repository.
func (arr dataArtifactRegistryRepositoryAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("project"))
}

// RepositoryId returns a reference to field repository_id of google_artifact_registry_repository.
func (arr dataArtifactRegistryRepositoryAttributes) RepositoryId() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("repository_id"))
}

// UpdateTime returns a reference to field update_time of google_artifact_registry_repository.
func (arr dataArtifactRegistryRepositoryAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("update_time"))
}

func (arr dataArtifactRegistryRepositoryAttributes) MavenConfig() terra.ListValue[dataartifactregistryrepository.MavenConfigAttributes] {
	return terra.ReferenceAsList[dataartifactregistryrepository.MavenConfigAttributes](arr.ref.Append("maven_config"))
}
