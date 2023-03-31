// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	artifactregistryrepository "github.com/golingon/terraproviders/google/4.59.0/artifactregistryrepository"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewArtifactRegistryRepository(name string, args ArtifactRegistryRepositoryArgs) *ArtifactRegistryRepository {
	return &ArtifactRegistryRepository{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ArtifactRegistryRepository)(nil)

type ArtifactRegistryRepository struct {
	Name  string
	Args  ArtifactRegistryRepositoryArgs
	state *artifactRegistryRepositoryState
}

func (arr *ArtifactRegistryRepository) Type() string {
	return "google_artifact_registry_repository"
}

func (arr *ArtifactRegistryRepository) LocalName() string {
	return arr.Name
}

func (arr *ArtifactRegistryRepository) Configuration() interface{} {
	return arr.Args
}

func (arr *ArtifactRegistryRepository) Attributes() artifactRegistryRepositoryAttributes {
	return artifactRegistryRepositoryAttributes{ref: terra.ReferenceResource(arr)}
}

func (arr *ArtifactRegistryRepository) ImportState(av io.Reader) error {
	arr.state = &artifactRegistryRepositoryState{}
	if err := json.NewDecoder(av).Decode(arr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", arr.Type(), arr.LocalName(), err)
	}
	return nil
}

func (arr *ArtifactRegistryRepository) State() (*artifactRegistryRepositoryState, bool) {
	return arr.state, arr.state != nil
}

func (arr *ArtifactRegistryRepository) StateMust() *artifactRegistryRepositoryState {
	if arr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", arr.Type(), arr.LocalName()))
	}
	return arr.state
}

func (arr *ArtifactRegistryRepository) DependOn() terra.Reference {
	return terra.ReferenceResource(arr)
}

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
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RepositoryId: string, required
	RepositoryId terra.StringValue `hcl:"repository_id,attr" validate:"required"`
	// MavenConfig: optional
	MavenConfig *artifactregistryrepository.MavenConfig `hcl:"maven_config,block"`
	// Timeouts: optional
	Timeouts *artifactregistryrepository.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ArtifactRegistryRepository depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type artifactRegistryRepositoryAttributes struct {
	ref terra.Reference
}

func (arr artifactRegistryRepositoryAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("create_time"))
}

func (arr artifactRegistryRepositoryAttributes) Description() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("description"))
}

func (arr artifactRegistryRepositoryAttributes) Format() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("format"))
}

func (arr artifactRegistryRepositoryAttributes) Id() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("id"))
}

func (arr artifactRegistryRepositoryAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("kms_key_name"))
}

func (arr artifactRegistryRepositoryAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](arr.ref.Append("labels"))
}

func (arr artifactRegistryRepositoryAttributes) Location() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("location"))
}

func (arr artifactRegistryRepositoryAttributes) Name() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("name"))
}

func (arr artifactRegistryRepositoryAttributes) Project() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("project"))
}

func (arr artifactRegistryRepositoryAttributes) RepositoryId() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("repository_id"))
}

func (arr artifactRegistryRepositoryAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceString(arr.ref.Append("update_time"))
}

func (arr artifactRegistryRepositoryAttributes) MavenConfig() terra.ListValue[artifactregistryrepository.MavenConfigAttributes] {
	return terra.ReferenceList[artifactregistryrepository.MavenConfigAttributes](arr.ref.Append("maven_config"))
}

func (arr artifactRegistryRepositoryAttributes) Timeouts() artifactregistryrepository.TimeoutsAttributes {
	return terra.ReferenceSingle[artifactregistryrepository.TimeoutsAttributes](arr.ref.Append("timeouts"))
}

type artifactRegistryRepositoryState struct {
	CreateTime   string                                        `json:"create_time"`
	Description  string                                        `json:"description"`
	Format       string                                        `json:"format"`
	Id           string                                        `json:"id"`
	KmsKeyName   string                                        `json:"kms_key_name"`
	Labels       map[string]string                             `json:"labels"`
	Location     string                                        `json:"location"`
	Name         string                                        `json:"name"`
	Project      string                                        `json:"project"`
	RepositoryId string                                        `json:"repository_id"`
	UpdateTime   string                                        `json:"update_time"`
	MavenConfig  []artifactregistryrepository.MavenConfigState `json:"maven_config"`
	Timeouts     *artifactregistryrepository.TimeoutsState     `json:"timeouts"`
}
