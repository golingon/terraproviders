// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_github_team

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

// Resource represents the Terraform resource vault_github_team.
type Resource struct {
	Name      string
	Args      Args
	state     *vaultGithubTeamState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (vgt *Resource) Type() string {
	return "vault_github_team"
}

// LocalName returns the local name for [Resource].
func (vgt *Resource) LocalName() string {
	return vgt.Name
}

// Configuration returns the configuration (args) for [Resource].
func (vgt *Resource) Configuration() interface{} {
	return vgt.Args
}

// DependOn is used for other resources to depend on [Resource].
func (vgt *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(vgt)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (vgt *Resource) Dependencies() terra.Dependencies {
	return vgt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (vgt *Resource) LifecycleManagement() *terra.Lifecycle {
	return vgt.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (vgt *Resource) Attributes() vaultGithubTeamAttributes {
	return vaultGithubTeamAttributes{ref: terra.ReferenceResource(vgt)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (vgt *Resource) ImportState(state io.Reader) error {
	vgt.state = &vaultGithubTeamState{}
	if err := json.NewDecoder(state).Decode(vgt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vgt.Type(), vgt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (vgt *Resource) State() (*vaultGithubTeamState, bool) {
	return vgt.state, vgt.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (vgt *Resource) StateMust() *vaultGithubTeamState {
	if vgt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vgt.Type(), vgt.LocalName()))
	}
	return vgt.state
}

// Args contains the configurations for vault_github_team.
type Args struct {
	// Backend: string, optional
	Backend terra.StringValue `hcl:"backend,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Policies: list of string, optional
	Policies terra.ListValue[terra.StringValue] `hcl:"policies,attr"`
	// Team: string, required
	Team terra.StringValue `hcl:"team,attr" validate:"required"`
}

type vaultGithubTeamAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_github_team.
func (vgt vaultGithubTeamAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(vgt.ref.Append("backend"))
}

// Id returns a reference to field id of vault_github_team.
func (vgt vaultGithubTeamAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vgt.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_github_team.
func (vgt vaultGithubTeamAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vgt.ref.Append("namespace"))
}

// Policies returns a reference to field policies of vault_github_team.
func (vgt vaultGithubTeamAttributes) Policies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vgt.ref.Append("policies"))
}

// Team returns a reference to field team of vault_github_team.
func (vgt vaultGithubTeamAttributes) Team() terra.StringValue {
	return terra.ReferenceAsString(vgt.ref.Append("team"))
}

type vaultGithubTeamState struct {
	Backend   string   `json:"backend"`
	Id        string   `json:"id"`
	Namespace string   `json:"namespace"`
	Policies  []string `json:"policies"`
	Team      string   `json:"team"`
}
