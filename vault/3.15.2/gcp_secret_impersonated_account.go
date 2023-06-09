// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGcpSecretImpersonatedAccount creates a new instance of [GcpSecretImpersonatedAccount].
func NewGcpSecretImpersonatedAccount(name string, args GcpSecretImpersonatedAccountArgs) *GcpSecretImpersonatedAccount {
	return &GcpSecretImpersonatedAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GcpSecretImpersonatedAccount)(nil)

// GcpSecretImpersonatedAccount represents the Terraform resource vault_gcp_secret_impersonated_account.
type GcpSecretImpersonatedAccount struct {
	Name      string
	Args      GcpSecretImpersonatedAccountArgs
	state     *gcpSecretImpersonatedAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GcpSecretImpersonatedAccount].
func (gsia *GcpSecretImpersonatedAccount) Type() string {
	return "vault_gcp_secret_impersonated_account"
}

// LocalName returns the local name for [GcpSecretImpersonatedAccount].
func (gsia *GcpSecretImpersonatedAccount) LocalName() string {
	return gsia.Name
}

// Configuration returns the configuration (args) for [GcpSecretImpersonatedAccount].
func (gsia *GcpSecretImpersonatedAccount) Configuration() interface{} {
	return gsia.Args
}

// DependOn is used for other resources to depend on [GcpSecretImpersonatedAccount].
func (gsia *GcpSecretImpersonatedAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(gsia)
}

// Dependencies returns the list of resources [GcpSecretImpersonatedAccount] depends_on.
func (gsia *GcpSecretImpersonatedAccount) Dependencies() terra.Dependencies {
	return gsia.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GcpSecretImpersonatedAccount].
func (gsia *GcpSecretImpersonatedAccount) LifecycleManagement() *terra.Lifecycle {
	return gsia.Lifecycle
}

// Attributes returns the attributes for [GcpSecretImpersonatedAccount].
func (gsia *GcpSecretImpersonatedAccount) Attributes() gcpSecretImpersonatedAccountAttributes {
	return gcpSecretImpersonatedAccountAttributes{ref: terra.ReferenceResource(gsia)}
}

// ImportState imports the given attribute values into [GcpSecretImpersonatedAccount]'s state.
func (gsia *GcpSecretImpersonatedAccount) ImportState(av io.Reader) error {
	gsia.state = &gcpSecretImpersonatedAccountState{}
	if err := json.NewDecoder(av).Decode(gsia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gsia.Type(), gsia.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GcpSecretImpersonatedAccount] has state.
func (gsia *GcpSecretImpersonatedAccount) State() (*gcpSecretImpersonatedAccountState, bool) {
	return gsia.state, gsia.state != nil
}

// StateMust returns the state for [GcpSecretImpersonatedAccount]. Panics if the state is nil.
func (gsia *GcpSecretImpersonatedAccount) StateMust() *gcpSecretImpersonatedAccountState {
	if gsia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gsia.Type(), gsia.LocalName()))
	}
	return gsia.state
}

// GcpSecretImpersonatedAccountArgs contains the configurations for vault_gcp_secret_impersonated_account.
type GcpSecretImpersonatedAccountArgs struct {
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
type gcpSecretImpersonatedAccountAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_gcp_secret_impersonated_account.
func (gsia gcpSecretImpersonatedAccountAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(gsia.ref.Append("backend"))
}

// Id returns a reference to field id of vault_gcp_secret_impersonated_account.
func (gsia gcpSecretImpersonatedAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gsia.ref.Append("id"))
}

// ImpersonatedAccount returns a reference to field impersonated_account of vault_gcp_secret_impersonated_account.
func (gsia gcpSecretImpersonatedAccountAttributes) ImpersonatedAccount() terra.StringValue {
	return terra.ReferenceAsString(gsia.ref.Append("impersonated_account"))
}

// Namespace returns a reference to field namespace of vault_gcp_secret_impersonated_account.
func (gsia gcpSecretImpersonatedAccountAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(gsia.ref.Append("namespace"))
}

// ServiceAccountEmail returns a reference to field service_account_email of vault_gcp_secret_impersonated_account.
func (gsia gcpSecretImpersonatedAccountAttributes) ServiceAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(gsia.ref.Append("service_account_email"))
}

// ServiceAccountProject returns a reference to field service_account_project of vault_gcp_secret_impersonated_account.
func (gsia gcpSecretImpersonatedAccountAttributes) ServiceAccountProject() terra.StringValue {
	return terra.ReferenceAsString(gsia.ref.Append("service_account_project"))
}

// TokenScopes returns a reference to field token_scopes of vault_gcp_secret_impersonated_account.
func (gsia gcpSecretImpersonatedAccountAttributes) TokenScopes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gsia.ref.Append("token_scopes"))
}

type gcpSecretImpersonatedAccountState struct {
	Backend               string   `json:"backend"`
	Id                    string   `json:"id"`
	ImpersonatedAccount   string   `json:"impersonated_account"`
	Namespace             string   `json:"namespace"`
	ServiceAccountEmail   string   `json:"service_account_email"`
	ServiceAccountProject string   `json:"service_account_project"`
	TokenScopes           []string `json:"token_scopes"`
}
