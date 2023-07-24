// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	artifactregistryrepositoryiammember "github.com/golingon/terraproviders/google/4.74.0/artifactregistryrepositoryiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewArtifactRegistryRepositoryIamMember creates a new instance of [ArtifactRegistryRepositoryIamMember].
func NewArtifactRegistryRepositoryIamMember(name string, args ArtifactRegistryRepositoryIamMemberArgs) *ArtifactRegistryRepositoryIamMember {
	return &ArtifactRegistryRepositoryIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ArtifactRegistryRepositoryIamMember)(nil)

// ArtifactRegistryRepositoryIamMember represents the Terraform resource google_artifact_registry_repository_iam_member.
type ArtifactRegistryRepositoryIamMember struct {
	Name      string
	Args      ArtifactRegistryRepositoryIamMemberArgs
	state     *artifactRegistryRepositoryIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ArtifactRegistryRepositoryIamMember].
func (arrim *ArtifactRegistryRepositoryIamMember) Type() string {
	return "google_artifact_registry_repository_iam_member"
}

// LocalName returns the local name for [ArtifactRegistryRepositoryIamMember].
func (arrim *ArtifactRegistryRepositoryIamMember) LocalName() string {
	return arrim.Name
}

// Configuration returns the configuration (args) for [ArtifactRegistryRepositoryIamMember].
func (arrim *ArtifactRegistryRepositoryIamMember) Configuration() interface{} {
	return arrim.Args
}

// DependOn is used for other resources to depend on [ArtifactRegistryRepositoryIamMember].
func (arrim *ArtifactRegistryRepositoryIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(arrim)
}

// Dependencies returns the list of resources [ArtifactRegistryRepositoryIamMember] depends_on.
func (arrim *ArtifactRegistryRepositoryIamMember) Dependencies() terra.Dependencies {
	return arrim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ArtifactRegistryRepositoryIamMember].
func (arrim *ArtifactRegistryRepositoryIamMember) LifecycleManagement() *terra.Lifecycle {
	return arrim.Lifecycle
}

// Attributes returns the attributes for [ArtifactRegistryRepositoryIamMember].
func (arrim *ArtifactRegistryRepositoryIamMember) Attributes() artifactRegistryRepositoryIamMemberAttributes {
	return artifactRegistryRepositoryIamMemberAttributes{ref: terra.ReferenceResource(arrim)}
}

// ImportState imports the given attribute values into [ArtifactRegistryRepositoryIamMember]'s state.
func (arrim *ArtifactRegistryRepositoryIamMember) ImportState(av io.Reader) error {
	arrim.state = &artifactRegistryRepositoryIamMemberState{}
	if err := json.NewDecoder(av).Decode(arrim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", arrim.Type(), arrim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ArtifactRegistryRepositoryIamMember] has state.
func (arrim *ArtifactRegistryRepositoryIamMember) State() (*artifactRegistryRepositoryIamMemberState, bool) {
	return arrim.state, arrim.state != nil
}

// StateMust returns the state for [ArtifactRegistryRepositoryIamMember]. Panics if the state is nil.
func (arrim *ArtifactRegistryRepositoryIamMember) StateMust() *artifactRegistryRepositoryIamMemberState {
	if arrim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", arrim.Type(), arrim.LocalName()))
	}
	return arrim.state
}

// ArtifactRegistryRepositoryIamMemberArgs contains the configurations for google_artifact_registry_repository_iam_member.
type ArtifactRegistryRepositoryIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Repository: string, required
	Repository terra.StringValue `hcl:"repository,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *artifactregistryrepositoryiammember.Condition `hcl:"condition,block"`
}
type artifactRegistryRepositoryIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_artifact_registry_repository_iam_member.
func (arrim artifactRegistryRepositoryIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(arrim.ref.Append("etag"))
}

// Id returns a reference to field id of google_artifact_registry_repository_iam_member.
func (arrim artifactRegistryRepositoryIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(arrim.ref.Append("id"))
}

// Location returns a reference to field location of google_artifact_registry_repository_iam_member.
func (arrim artifactRegistryRepositoryIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(arrim.ref.Append("location"))
}

// Member returns a reference to field member of google_artifact_registry_repository_iam_member.
func (arrim artifactRegistryRepositoryIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(arrim.ref.Append("member"))
}

// Project returns a reference to field project of google_artifact_registry_repository_iam_member.
func (arrim artifactRegistryRepositoryIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(arrim.ref.Append("project"))
}

// Repository returns a reference to field repository of google_artifact_registry_repository_iam_member.
func (arrim artifactRegistryRepositoryIamMemberAttributes) Repository() terra.StringValue {
	return terra.ReferenceAsString(arrim.ref.Append("repository"))
}

// Role returns a reference to field role of google_artifact_registry_repository_iam_member.
func (arrim artifactRegistryRepositoryIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(arrim.ref.Append("role"))
}

func (arrim artifactRegistryRepositoryIamMemberAttributes) Condition() terra.ListValue[artifactregistryrepositoryiammember.ConditionAttributes] {
	return terra.ReferenceAsList[artifactregistryrepositoryiammember.ConditionAttributes](arrim.ref.Append("condition"))
}

type artifactRegistryRepositoryIamMemberState struct {
	Etag       string                                               `json:"etag"`
	Id         string                                               `json:"id"`
	Location   string                                               `json:"location"`
	Member     string                                               `json:"member"`
	Project    string                                               `json:"project"`
	Repository string                                               `json:"repository"`
	Role       string                                               `json:"role"`
	Condition  []artifactregistryrepositoryiammember.ConditionState `json:"condition"`
}
