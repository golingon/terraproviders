// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_gcp_secret_impersonated_account

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

// Resource represents the Terraform resource vault_gcp_secret_impersonated_account.
type Resource struct {
	Name      string
	Args      Args
	state     *vaultGcpSecretImpersonatedAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (vgsia *Resource) Type() string {
	return "vault_gcp_secret_impersonated_account"
}

// LocalName returns the local name for [Resource].
func (vgsia *Resource) LocalName() string {
	return vgsia.Name
}

// Configuration returns the configuration (args) for [Resource].
func (vgsia *Resource) Configuration() interface{} {
	return vgsia.Args
}

// DependOn is used for other resources to depend on [Resource].
func (vgsia *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(vgsia)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (vgsia *Resource) Dependencies() terra.Dependencies {
	return vgsia.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (vgsia *Resource) LifecycleManagement() *terra.Lifecycle {
	return vgsia.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (vgsia *Resource) Attributes() vaultGcpSecretImpersonatedAccountAttributes {
	return vaultGcpSecretImpersonatedAccountAttributes{ref: terra.ReferenceResource(vgsia)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (vgsia *Resource) ImportState(state io.Reader) error {
	vgsia.state = &vaultGcpSecretImpersonatedAccountState{}
	if err := json.NewDecoder(state).Decode(vgsia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vgsia.Type(), vgsia.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (vgsia *Resource) State() (*vaultGcpSecretImpersonatedAccountState, bool) {
	return vgsia.state, vgsia.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (vgsia *Resource) StateMust() *vaultGcpSecretImpersonatedAccountState {
	if vgsia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vgsia.Type(), vgsia.LocalName()))
	}
	return vgsia.state
}

// Args contains the configurations for vault_gcp_secret_impersonated_account.
type Args struct {
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImpersonatedAccount: string, required
	ImpersonatedAccount terra.StringValue `hcl:"impersonated_account,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// ServiceAccountEmail: string, required
	ServiceAccountEmail terra.StringValue `hcl:"service_account_email,attr" validate:"required"`
	// TokenScopes: set of string, optional
	TokenScopes terra.SetValue[terra.StringValue] `hcl:"token_scopes,attr"`
}

type vaultGcpSecretImpersonatedAccountAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_gcp_secret_impersonated_account.
func (vgsia vaultGcpSecretImpersonatedAccountAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(vgsia.ref.Append("backend"))
}

// Id returns a reference to field id of vault_gcp_secret_impersonated_account.
func (vgsia vaultGcpSecretImpersonatedAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vgsia.ref.Append("id"))
}

// ImpersonatedAccount returns a reference to field impersonated_account of vault_gcp_secret_impersonated_account.
func (vgsia vaultGcpSecretImpersonatedAccountAttributes) ImpersonatedAccount() terra.StringValue {
	return terra.ReferenceAsString(vgsia.ref.Append("impersonated_account"))
}

// Namespace returns a reference to field namespace of vault_gcp_secret_impersonated_account.
func (vgsia vaultGcpSecretImpersonatedAccountAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vgsia.ref.Append("namespace"))
}

// ServiceAccountEmail returns a reference to field service_account_email of vault_gcp_secret_impersonated_account.
func (vgsia vaultGcpSecretImpersonatedAccountAttributes) ServiceAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(vgsia.ref.Append("service_account_email"))
}

// ServiceAccountProject returns a reference to field service_account_project of vault_gcp_secret_impersonated_account.
func (vgsia vaultGcpSecretImpersonatedAccountAttributes) ServiceAccountProject() terra.StringValue {
	return terra.ReferenceAsString(vgsia.ref.Append("service_account_project"))
}

// TokenScopes returns a reference to field token_scopes of vault_gcp_secret_impersonated_account.
func (vgsia vaultGcpSecretImpersonatedAccountAttributes) TokenScopes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgsia.ref.Append("token_scopes"))
}

type vaultGcpSecretImpersonatedAccountState struct {
	Backend               string   `json:"backend"`
	Id                    string   `json:"id"`
	ImpersonatedAccount   string   `json:"impersonated_account"`
	Namespace             string   `json:"namespace"`
	ServiceAccountEmail   string   `json:"service_account_email"`
	ServiceAccountProject string   `json:"service_account_project"`
	TokenScopes           []string `json:"token_scopes"`
}
