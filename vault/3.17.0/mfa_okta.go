// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMfaOkta creates a new instance of [MfaOkta].
func NewMfaOkta(name string, args MfaOktaArgs) *MfaOkta {
	return &MfaOkta{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MfaOkta)(nil)

// MfaOkta represents the Terraform resource vault_mfa_okta.
type MfaOkta struct {
	Name      string
	Args      MfaOktaArgs
	state     *mfaOktaState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MfaOkta].
func (mo *MfaOkta) Type() string {
	return "vault_mfa_okta"
}

// LocalName returns the local name for [MfaOkta].
func (mo *MfaOkta) LocalName() string {
	return mo.Name
}

// Configuration returns the configuration (args) for [MfaOkta].
func (mo *MfaOkta) Configuration() interface{} {
	return mo.Args
}

// DependOn is used for other resources to depend on [MfaOkta].
func (mo *MfaOkta) DependOn() terra.Reference {
	return terra.ReferenceResource(mo)
}

// Dependencies returns the list of resources [MfaOkta] depends_on.
func (mo *MfaOkta) Dependencies() terra.Dependencies {
	return mo.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MfaOkta].
func (mo *MfaOkta) LifecycleManagement() *terra.Lifecycle {
	return mo.Lifecycle
}

// Attributes returns the attributes for [MfaOkta].
func (mo *MfaOkta) Attributes() mfaOktaAttributes {
	return mfaOktaAttributes{ref: terra.ReferenceResource(mo)}
}

// ImportState imports the given attribute values into [MfaOkta]'s state.
func (mo *MfaOkta) ImportState(av io.Reader) error {
	mo.state = &mfaOktaState{}
	if err := json.NewDecoder(av).Decode(mo.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mo.Type(), mo.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MfaOkta] has state.
func (mo *MfaOkta) State() (*mfaOktaState, bool) {
	return mo.state, mo.state != nil
}

// StateMust returns the state for [MfaOkta]. Panics if the state is nil.
func (mo *MfaOkta) StateMust() *mfaOktaState {
	if mo.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mo.Type(), mo.LocalName()))
	}
	return mo.state
}

// MfaOktaArgs contains the configurations for vault_mfa_okta.
type MfaOktaArgs struct {
	// ApiToken: string, required
	ApiToken terra.StringValue `hcl:"api_token,attr" validate:"required"`
	// BaseUrl: string, optional
	BaseUrl terra.StringValue `hcl:"base_url,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MountAccessor: string, required
	MountAccessor terra.StringValue `hcl:"mount_accessor,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// OrgName: string, required
	OrgName terra.StringValue `hcl:"org_name,attr" validate:"required"`
	// PrimaryEmail: bool, optional
	PrimaryEmail terra.BoolValue `hcl:"primary_email,attr"`
	// UsernameFormat: string, optional
	UsernameFormat terra.StringValue `hcl:"username_format,attr"`
}
type mfaOktaAttributes struct {
	ref terra.Reference
}

// ApiToken returns a reference to field api_token of vault_mfa_okta.
func (mo mfaOktaAttributes) ApiToken() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("api_token"))
}

// BaseUrl returns a reference to field base_url of vault_mfa_okta.
func (mo mfaOktaAttributes) BaseUrl() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("base_url"))
}

// Id returns a reference to field id of vault_mfa_okta.
func (mo mfaOktaAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("id"))
}

// MountAccessor returns a reference to field mount_accessor of vault_mfa_okta.
func (mo mfaOktaAttributes) MountAccessor() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("mount_accessor"))
}

// Name returns a reference to field name of vault_mfa_okta.
func (mo mfaOktaAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_mfa_okta.
func (mo mfaOktaAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("namespace"))
}

// OrgName returns a reference to field org_name of vault_mfa_okta.
func (mo mfaOktaAttributes) OrgName() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("org_name"))
}

// PrimaryEmail returns a reference to field primary_email of vault_mfa_okta.
func (mo mfaOktaAttributes) PrimaryEmail() terra.BoolValue {
	return terra.ReferenceAsBool(mo.ref.Append("primary_email"))
}

// UsernameFormat returns a reference to field username_format of vault_mfa_okta.
func (mo mfaOktaAttributes) UsernameFormat() terra.StringValue {
	return terra.ReferenceAsString(mo.ref.Append("username_format"))
}

type mfaOktaState struct {
	ApiToken       string `json:"api_token"`
	BaseUrl        string `json:"base_url"`
	Id             string `json:"id"`
	MountAccessor  string `json:"mount_accessor"`
	Name           string `json:"name"`
	Namespace      string `json:"namespace"`
	OrgName        string `json:"org_name"`
	PrimaryEmail   bool   `json:"primary_email"`
	UsernameFormat string `json:"username_format"`
}