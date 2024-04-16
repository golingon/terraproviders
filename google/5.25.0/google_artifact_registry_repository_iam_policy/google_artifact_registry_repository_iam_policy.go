// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_artifact_registry_repository_iam_policy

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_artifact_registry_repository_iam_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleArtifactRegistryRepositoryIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (garrip *Resource) Type() string {
	return "google_artifact_registry_repository_iam_policy"
}

// LocalName returns the local name for [Resource].
func (garrip *Resource) LocalName() string {
	return garrip.Name
}

// Configuration returns the configuration (args) for [Resource].
func (garrip *Resource) Configuration() interface{} {
	return garrip.Args
}

// DependOn is used for other resources to depend on [Resource].
func (garrip *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(garrip)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (garrip *Resource) Dependencies() terra.Dependencies {
	return garrip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (garrip *Resource) LifecycleManagement() *terra.Lifecycle {
	return garrip.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (garrip *Resource) Attributes() googleArtifactRegistryRepositoryIamPolicyAttributes {
	return googleArtifactRegistryRepositoryIamPolicyAttributes{ref: terra.ReferenceResource(garrip)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (garrip *Resource) ImportState(state io.Reader) error {
	garrip.state = &googleArtifactRegistryRepositoryIamPolicyState{}
	if err := json.NewDecoder(state).Decode(garrip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", garrip.Type(), garrip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (garrip *Resource) State() (*googleArtifactRegistryRepositoryIamPolicyState, bool) {
	return garrip.state, garrip.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (garrip *Resource) StateMust() *googleArtifactRegistryRepositoryIamPolicyState {
	if garrip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", garrip.Type(), garrip.LocalName()))
	}
	return garrip.state
}

// Args contains the configurations for google_artifact_registry_repository_iam_policy.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Repository: string, required
	Repository terra.StringValue `hcl:"repository,attr" validate:"required"`
}

type googleArtifactRegistryRepositoryIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_artifact_registry_repository_iam_policy.
func (garrip googleArtifactRegistryRepositoryIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(garrip.ref.Append("etag"))
}

// Id returns a reference to field id of google_artifact_registry_repository_iam_policy.
func (garrip googleArtifactRegistryRepositoryIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(garrip.ref.Append("id"))
}

// Location returns a reference to field location of google_artifact_registry_repository_iam_policy.
func (garrip googleArtifactRegistryRepositoryIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(garrip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_artifact_registry_repository_iam_policy.
func (garrip googleArtifactRegistryRepositoryIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(garrip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_artifact_registry_repository_iam_policy.
func (garrip googleArtifactRegistryRepositoryIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(garrip.ref.Append("project"))
}

// Repository returns a reference to field repository of google_artifact_registry_repository_iam_policy.
func (garrip googleArtifactRegistryRepositoryIamPolicyAttributes) Repository() terra.StringValue {
	return terra.ReferenceAsString(garrip.ref.Append("repository"))
}

type googleArtifactRegistryRepositoryIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Location   string `json:"location"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	Repository string `json:"repository"`
}
