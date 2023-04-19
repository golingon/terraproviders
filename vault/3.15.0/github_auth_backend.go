// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	githubauthbackend "github.com/golingon/terraproviders/vault/3.15.0/githubauthbackend"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGithubAuthBackend creates a new instance of [GithubAuthBackend].
func NewGithubAuthBackend(name string, args GithubAuthBackendArgs) *GithubAuthBackend {
	return &GithubAuthBackend{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GithubAuthBackend)(nil)

// GithubAuthBackend represents the Terraform resource vault_github_auth_backend.
type GithubAuthBackend struct {
	Name      string
	Args      GithubAuthBackendArgs
	state     *githubAuthBackendState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GithubAuthBackend].
func (gab *GithubAuthBackend) Type() string {
	return "vault_github_auth_backend"
}

// LocalName returns the local name for [GithubAuthBackend].
func (gab *GithubAuthBackend) LocalName() string {
	return gab.Name
}

// Configuration returns the configuration (args) for [GithubAuthBackend].
func (gab *GithubAuthBackend) Configuration() interface{} {
	return gab.Args
}

// DependOn is used for other resources to depend on [GithubAuthBackend].
func (gab *GithubAuthBackend) DependOn() terra.Reference {
	return terra.ReferenceResource(gab)
}

// Dependencies returns the list of resources [GithubAuthBackend] depends_on.
func (gab *GithubAuthBackend) Dependencies() terra.Dependencies {
	return gab.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GithubAuthBackend].
func (gab *GithubAuthBackend) LifecycleManagement() *terra.Lifecycle {
	return gab.Lifecycle
}

// Attributes returns the attributes for [GithubAuthBackend].
func (gab *GithubAuthBackend) Attributes() githubAuthBackendAttributes {
	return githubAuthBackendAttributes{ref: terra.ReferenceResource(gab)}
}

// ImportState imports the given attribute values into [GithubAuthBackend]'s state.
func (gab *GithubAuthBackend) ImportState(av io.Reader) error {
	gab.state = &githubAuthBackendState{}
	if err := json.NewDecoder(av).Decode(gab.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gab.Type(), gab.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GithubAuthBackend] has state.
func (gab *GithubAuthBackend) State() (*githubAuthBackendState, bool) {
	return gab.state, gab.state != nil
}

// StateMust returns the state for [GithubAuthBackend]. Panics if the state is nil.
func (gab *GithubAuthBackend) StateMust() *githubAuthBackendState {
	if gab.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gab.Type(), gab.LocalName()))
	}
	return gab.state
}

// GithubAuthBackendArgs contains the configurations for vault_github_auth_backend.
type GithubAuthBackendArgs struct {
	// BaseUrl: string, optional
	BaseUrl terra.StringValue `hcl:"base_url,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisableRemount: bool, optional
	DisableRemount terra.BoolValue `hcl:"disable_remount,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Organization: string, required
	Organization terra.StringValue `hcl:"organization,attr" validate:"required"`
	// OrganizationId: number, optional
	OrganizationId terra.NumberValue `hcl:"organization_id,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// TokenBoundCidrs: set of string, optional
	TokenBoundCidrs terra.SetValue[terra.StringValue] `hcl:"token_bound_cidrs,attr"`
	// TokenExplicitMaxTtl: number, optional
	TokenExplicitMaxTtl terra.NumberValue `hcl:"token_explicit_max_ttl,attr"`
	// TokenMaxTtl: number, optional
	TokenMaxTtl terra.NumberValue `hcl:"token_max_ttl,attr"`
	// TokenNoDefaultPolicy: bool, optional
	TokenNoDefaultPolicy terra.BoolValue `hcl:"token_no_default_policy,attr"`
	// TokenNumUses: number, optional
	TokenNumUses terra.NumberValue `hcl:"token_num_uses,attr"`
	// TokenPeriod: number, optional
	TokenPeriod terra.NumberValue `hcl:"token_period,attr"`
	// TokenPolicies: set of string, optional
	TokenPolicies terra.SetValue[terra.StringValue] `hcl:"token_policies,attr"`
	// TokenTtl: number, optional
	TokenTtl terra.NumberValue `hcl:"token_ttl,attr"`
	// TokenType: string, optional
	TokenType terra.StringValue `hcl:"token_type,attr"`
	// Tune: min=0
	Tune []githubauthbackend.Tune `hcl:"tune,block" validate:"min=0"`
}
type githubAuthBackendAttributes struct {
	ref terra.Reference
}

// Accessor returns a reference to field accessor of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) Accessor() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("accessor"))
}

// BaseUrl returns a reference to field base_url of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) BaseUrl() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("base_url"))
}

// Description returns a reference to field description of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("description"))
}

// DisableRemount returns a reference to field disable_remount of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) DisableRemount() terra.BoolValue {
	return terra.ReferenceAsBool(gab.ref.Append("disable_remount"))
}

// Id returns a reference to field id of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("namespace"))
}

// Organization returns a reference to field organization of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("organization"))
}

// OrganizationId returns a reference to field organization_id of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) OrganizationId() terra.NumberValue {
	return terra.ReferenceAsNumber(gab.ref.Append("organization_id"))
}

// Path returns a reference to field path of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("path"))
}

// TokenBoundCidrs returns a reference to field token_bound_cidrs of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) TokenBoundCidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gab.ref.Append("token_bound_cidrs"))
}

// TokenExplicitMaxTtl returns a reference to field token_explicit_max_ttl of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) TokenExplicitMaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(gab.ref.Append("token_explicit_max_ttl"))
}

// TokenMaxTtl returns a reference to field token_max_ttl of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) TokenMaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(gab.ref.Append("token_max_ttl"))
}

// TokenNoDefaultPolicy returns a reference to field token_no_default_policy of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) TokenNoDefaultPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(gab.ref.Append("token_no_default_policy"))
}

// TokenNumUses returns a reference to field token_num_uses of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) TokenNumUses() terra.NumberValue {
	return terra.ReferenceAsNumber(gab.ref.Append("token_num_uses"))
}

// TokenPeriod returns a reference to field token_period of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) TokenPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(gab.ref.Append("token_period"))
}

// TokenPolicies returns a reference to field token_policies of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) TokenPolicies() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gab.ref.Append("token_policies"))
}

// TokenTtl returns a reference to field token_ttl of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) TokenTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(gab.ref.Append("token_ttl"))
}

// TokenType returns a reference to field token_type of vault_github_auth_backend.
func (gab githubAuthBackendAttributes) TokenType() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("token_type"))
}

func (gab githubAuthBackendAttributes) Tune() terra.SetValue[githubauthbackend.TuneAttributes] {
	return terra.ReferenceAsSet[githubauthbackend.TuneAttributes](gab.ref.Append("tune"))
}

type githubAuthBackendState struct {
	Accessor             string                        `json:"accessor"`
	BaseUrl              string                        `json:"base_url"`
	Description          string                        `json:"description"`
	DisableRemount       bool                          `json:"disable_remount"`
	Id                   string                        `json:"id"`
	Namespace            string                        `json:"namespace"`
	Organization         string                        `json:"organization"`
	OrganizationId       float64                       `json:"organization_id"`
	Path                 string                        `json:"path"`
	TokenBoundCidrs      []string                      `json:"token_bound_cidrs"`
	TokenExplicitMaxTtl  float64                       `json:"token_explicit_max_ttl"`
	TokenMaxTtl          float64                       `json:"token_max_ttl"`
	TokenNoDefaultPolicy bool                          `json:"token_no_default_policy"`
	TokenNumUses         float64                       `json:"token_num_uses"`
	TokenPeriod          float64                       `json:"token_period"`
	TokenPolicies        []string                      `json:"token_policies"`
	TokenTtl             float64                       `json:"token_ttl"`
	TokenType            string                        `json:"token_type"`
	Tune                 []githubauthbackend.TuneState `json:"tune"`
}