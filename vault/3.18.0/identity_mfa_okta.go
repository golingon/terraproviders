// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIdentityMfaOkta creates a new instance of [IdentityMfaOkta].
func NewIdentityMfaOkta(name string, args IdentityMfaOktaArgs) *IdentityMfaOkta {
	return &IdentityMfaOkta{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityMfaOkta)(nil)

// IdentityMfaOkta represents the Terraform resource vault_identity_mfa_okta.
type IdentityMfaOkta struct {
	Name      string
	Args      IdentityMfaOktaArgs
	state     *identityMfaOktaState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentityMfaOkta].
func (imo *IdentityMfaOkta) Type() string {
	return "vault_identity_mfa_okta"
}

// LocalName returns the local name for [IdentityMfaOkta].
func (imo *IdentityMfaOkta) LocalName() string {
	return imo.Name
}

// Configuration returns the configuration (args) for [IdentityMfaOkta].
func (imo *IdentityMfaOkta) Configuration() interface{} {
	return imo.Args
}

// DependOn is used for other resources to depend on [IdentityMfaOkta].
func (imo *IdentityMfaOkta) DependOn() terra.Reference {
	return terra.ReferenceResource(imo)
}

// Dependencies returns the list of resources [IdentityMfaOkta] depends_on.
func (imo *IdentityMfaOkta) Dependencies() terra.Dependencies {
	return imo.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentityMfaOkta].
func (imo *IdentityMfaOkta) LifecycleManagement() *terra.Lifecycle {
	return imo.Lifecycle
}

// Attributes returns the attributes for [IdentityMfaOkta].
func (imo *IdentityMfaOkta) Attributes() identityMfaOktaAttributes {
	return identityMfaOktaAttributes{ref: terra.ReferenceResource(imo)}
}

// ImportState imports the given attribute values into [IdentityMfaOkta]'s state.
func (imo *IdentityMfaOkta) ImportState(av io.Reader) error {
	imo.state = &identityMfaOktaState{}
	if err := json.NewDecoder(av).Decode(imo.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", imo.Type(), imo.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentityMfaOkta] has state.
func (imo *IdentityMfaOkta) State() (*identityMfaOktaState, bool) {
	return imo.state, imo.state != nil
}

// StateMust returns the state for [IdentityMfaOkta]. Panics if the state is nil.
func (imo *IdentityMfaOkta) StateMust() *identityMfaOktaState {
	if imo.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", imo.Type(), imo.LocalName()))
	}
	return imo.state
}

// IdentityMfaOktaArgs contains the configurations for vault_identity_mfa_okta.
type IdentityMfaOktaArgs struct {
	// ApiToken: string, required
	ApiToken terra.StringValue `hcl:"api_token,attr" validate:"required"`
	// BaseUrl: string, optional
	BaseUrl terra.StringValue `hcl:"base_url,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// OrgName: string, required
	OrgName terra.StringValue `hcl:"org_name,attr" validate:"required"`
	// PrimaryEmail: bool, optional
	PrimaryEmail terra.BoolValue `hcl:"primary_email,attr"`
	// UsernameFormat: string, optional
	UsernameFormat terra.StringValue `hcl:"username_format,attr"`
}
type identityMfaOktaAttributes struct {
	ref terra.Reference
}

// ApiToken returns a reference to field api_token of vault_identity_mfa_okta.
func (imo identityMfaOktaAttributes) ApiToken() terra.StringValue {
	return terra.ReferenceAsString(imo.ref.Append("api_token"))
}

// BaseUrl returns a reference to field base_url of vault_identity_mfa_okta.
func (imo identityMfaOktaAttributes) BaseUrl() terra.StringValue {
	return terra.ReferenceAsString(imo.ref.Append("base_url"))
}

// Id returns a reference to field id of vault_identity_mfa_okta.
func (imo identityMfaOktaAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(imo.ref.Append("id"))
}

// MethodId returns a reference to field method_id of vault_identity_mfa_okta.
func (imo identityMfaOktaAttributes) MethodId() terra.StringValue {
	return terra.ReferenceAsString(imo.ref.Append("method_id"))
}

// MountAccessor returns a reference to field mount_accessor of vault_identity_mfa_okta.
func (imo identityMfaOktaAttributes) MountAccessor() terra.StringValue {
	return terra.ReferenceAsString(imo.ref.Append("mount_accessor"))
}

// Name returns a reference to field name of vault_identity_mfa_okta.
func (imo identityMfaOktaAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(imo.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_identity_mfa_okta.
func (imo identityMfaOktaAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(imo.ref.Append("namespace"))
}

// NamespaceId returns a reference to field namespace_id of vault_identity_mfa_okta.
func (imo identityMfaOktaAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(imo.ref.Append("namespace_id"))
}

// NamespacePath returns a reference to field namespace_path of vault_identity_mfa_okta.
func (imo identityMfaOktaAttributes) NamespacePath() terra.StringValue {
	return terra.ReferenceAsString(imo.ref.Append("namespace_path"))
}

// OrgName returns a reference to field org_name of vault_identity_mfa_okta.
func (imo identityMfaOktaAttributes) OrgName() terra.StringValue {
	return terra.ReferenceAsString(imo.ref.Append("org_name"))
}

// PrimaryEmail returns a reference to field primary_email of vault_identity_mfa_okta.
func (imo identityMfaOktaAttributes) PrimaryEmail() terra.BoolValue {
	return terra.ReferenceAsBool(imo.ref.Append("primary_email"))
}

// Type returns a reference to field type of vault_identity_mfa_okta.
func (imo identityMfaOktaAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(imo.ref.Append("type"))
}

// UsernameFormat returns a reference to field username_format of vault_identity_mfa_okta.
func (imo identityMfaOktaAttributes) UsernameFormat() terra.StringValue {
	return terra.ReferenceAsString(imo.ref.Append("username_format"))
}

// Uuid returns a reference to field uuid of vault_identity_mfa_okta.
func (imo identityMfaOktaAttributes) Uuid() terra.StringValue {
	return terra.ReferenceAsString(imo.ref.Append("uuid"))
}

type identityMfaOktaState struct {
	ApiToken       string `json:"api_token"`
	BaseUrl        string `json:"base_url"`
	Id             string `json:"id"`
	MethodId       string `json:"method_id"`
	MountAccessor  string `json:"mount_accessor"`
	Name           string `json:"name"`
	Namespace      string `json:"namespace"`
	NamespaceId    string `json:"namespace_id"`
	NamespacePath  string `json:"namespace_path"`
	OrgName        string `json:"org_name"`
	PrimaryEmail   bool   `json:"primary_email"`
	Type           string `json:"type"`
	UsernameFormat string `json:"username_format"`
	Uuid           string `json:"uuid"`
}