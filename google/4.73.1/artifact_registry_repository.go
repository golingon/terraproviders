// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	artifactregistryrepository "github.com/golingon/terraproviders/google/4.73.1/artifactregistryrepository"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewArtifactRegistryRepository creates a new instance of [ArtifactRegistryRepository].
func NewArtifactRegistryRepository(name string, args ArtifactRegistryRepositoryArgs) *ArtifactRegistryRepository {
	return &ArtifactRegistryRepository{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ArtifactRegistryRepository)(nil)

// ArtifactRegistryRepository represents the Terraform resource google_artifact_registry_repository.
type ArtifactRegistryRepository struct {
	Name      string
	Args      ArtifactRegistryRepositoryArgs
	state     *artifactRegistryRepositoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ArtifactRegistryRepository].
func (arr *ArtifactRegistryRepository) Type() string {
	return "google_artifact_registry_repository"
}

// LocalName returns the local name for [ArtifactRegistryRepository].
func (arr *ArtifactRegistryRepository) LocalName() string {
	return arr.Name
}

// Configuration returns the configuration (args) for [ArtifactRegistryRepository].
func (arr *ArtifactRegistryRepository) Configuration() interface{} {
	return arr.Args
}

// DependOn is used for other resources to depend on [ArtifactRegistryRepository].
func (arr *ArtifactRegistryRepository) DependOn() terra.Reference {
	return terra.ReferenceResource(arr)
}

// Dependencies returns the list of resources [ArtifactRegistryRepository] depends_on.
func (arr *ArtifactRegistryRepository) Dependencies() terra.Dependencies {
	return arr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ArtifactRegistryRepository].
func (arr *ArtifactRegistryRepository) LifecycleManagement() *terra.Lifecycle {
	return arr.Lifecycle
}

// Attributes returns the attributes for [ArtifactRegistryRepository].
func (arr *ArtifactRegistryRepository) Attributes() artifactRegistryRepositoryAttributes {
	return artifactRegistryRepositoryAttributes{ref: terra.ReferenceResource(arr)}
}

// ImportState imports the given attribute values into [ArtifactRegistryRepository]'s state.
func (arr *ArtifactRegistryRepository) ImportState(av io.Reader) error {
	arr.state = &artifactRegistryRepositoryState{}
	if err := json.NewDecoder(av).Decode(arr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", arr.Type(), arr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ArtifactRegistryRepository] has state.
func (arr *ArtifactRegistryRepository) State() (*artifactRegistryRepositoryState, bool) {
	return arr.state, arr.state != nil
}

// StateMust returns the state for [ArtifactRegistryRepository]. Panics if the state is nil.
func (arr *ArtifactRegistryRepository) StateMust() *artifactRegistryRepositoryState {
	if arr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", arr.Type(), arr.LocalName()))
	}
	return arr.state
}

// ArtifactRegistryRepositoryArgs contains the configurations for google_artifact_registry_repository.
type ArtifactRegistryRepositoryArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Format: string, required
	Format terra.StringValue `hcl:"format,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyName: string, optional
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RepositoryId: string, required
	RepositoryId terra.StringValue `hcl:"repository_id,attr" validate:"required"`
	// DockerConfig: optional
	DockerConfig *artifactregistryrepository.DockerConfig `hcl:"docker_config,block"`
	// MavenConfig: optional
	MavenConfig *artifactregistryrepository.MavenConfig `hcl:"maven_config,block"`
	// RemoteRepositoryConfig: optional
	RemoteRepositoryConfig *artifactregistryrepository.RemoteRepositoryConfig `hcl:"remote_repository_config,block"`
	// Timeouts: optional
	Timeouts *artifactregistryrepository.Timeouts `hcl:"timeouts,block"`
	// VirtualRepositoryConfig: optional
	VirtualRepositoryConfig *artifactregistryrepository.VirtualRepositoryConfig `hcl:"virtual_repository_config,block"`
}
type artifactRegistryRepositoryAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_artifact_registry_repository.
func (arr artifactRegistryRepositoryAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("create_time"))
}

// Description returns a reference to field description of google_artifact_registry_repository.
func (arr artifactRegistryRepositoryAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("description"))
}

// Format returns a reference to field format of google_artifact_registry_repository.
func (arr artifactRegistryRepositoryAttributes) Format() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("format"))
}

// Id returns a reference to field id of google_artifact_registry_repository.
func (arr artifactRegistryRepositoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("id"))
}

// KmsKeyName returns a reference to field kms_key_name of google_artifact_registry_repository.
func (arr artifactRegistryRepositoryAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("kms_key_name"))
}

// Labels returns a reference to field labels of google_artifact_registry_repository.
func (arr artifactRegistryRepositoryAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](arr.ref.Append("labels"))
}

// Location returns a reference to field location of google_artifact_registry_repository.
func (arr artifactRegistryRepositoryAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("location"))
}

// Mode returns a reference to field mode of google_artifact_registry_repository.
func (arr artifactRegistryRepositoryAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("mode"))
}

// Name returns a reference to field name of google_artifact_registry_repository.
func (arr artifactRegistryRepositoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("name"))
}

// Project returns a reference to field project of google_artifact_registry_repository.
func (arr artifactRegistryRepositoryAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("project"))
}

// RepositoryId returns a reference to field repository_id of google_artifact_registry_repository.
func (arr artifactRegistryRepositoryAttributes) RepositoryId() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("repository_id"))
}

// UpdateTime returns a reference to field update_time of google_artifact_registry_repository.
func (arr artifactRegistryRepositoryAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("update_time"))
}

func (arr artifactRegistryRepositoryAttributes) DockerConfig() terra.ListValue[artifactregistryrepository.DockerConfigAttributes] {
	return terra.ReferenceAsList[artifactregistryrepository.DockerConfigAttributes](arr.ref.Append("docker_config"))
}

func (arr artifactRegistryRepositoryAttributes) MavenConfig() terra.ListValue[artifactregistryrepository.MavenConfigAttributes] {
	return terra.ReferenceAsList[artifactregistryrepository.MavenConfigAttributes](arr.ref.Append("maven_config"))
}

func (arr artifactRegistryRepositoryAttributes) RemoteRepositoryConfig() terra.ListValue[artifactregistryrepository.RemoteRepositoryConfigAttributes] {
	return terra.ReferenceAsList[artifactregistryrepository.RemoteRepositoryConfigAttributes](arr.ref.Append("remote_repository_config"))
}

func (arr artifactRegistryRepositoryAttributes) Timeouts() artifactregistryrepository.TimeoutsAttributes {
	return terra.ReferenceAsSingle[artifactregistryrepository.TimeoutsAttributes](arr.ref.Append("timeouts"))
}

func (arr artifactRegistryRepositoryAttributes) VirtualRepositoryConfig() terra.ListValue[artifactregistryrepository.VirtualRepositoryConfigAttributes] {
	return terra.ReferenceAsList[artifactregistryrepository.VirtualRepositoryConfigAttributes](arr.ref.Append("virtual_repository_config"))
}

type artifactRegistryRepositoryState struct {
	CreateTime              string                                                    `json:"create_time"`
	Description             string                                                    `json:"description"`
	Format                  string                                                    `json:"format"`
	Id                      string                                                    `json:"id"`
	KmsKeyName              string                                                    `json:"kms_key_name"`
	Labels                  map[string]string                                         `json:"labels"`
	Location                string                                                    `json:"location"`
	Mode                    string                                                    `json:"mode"`
	Name                    string                                                    `json:"name"`
	Project                 string                                                    `json:"project"`
	RepositoryId            string                                                    `json:"repository_id"`
	UpdateTime              string                                                    `json:"update_time"`
	DockerConfig            []artifactregistryrepository.DockerConfigState            `json:"docker_config"`
	MavenConfig             []artifactregistryrepository.MavenConfigState             `json:"maven_config"`
	RemoteRepositoryConfig  []artifactregistryrepository.RemoteRepositoryConfigState  `json:"remote_repository_config"`
	Timeouts                *artifactregistryrepository.TimeoutsState                 `json:"timeouts"`
	VirtualRepositoryConfig []artifactregistryrepository.VirtualRepositoryConfigState `json:"virtual_repository_config"`
}
