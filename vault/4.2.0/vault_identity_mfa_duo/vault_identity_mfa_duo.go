// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_identity_mfa_duo

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

// Resource represents the Terraform resource vault_identity_mfa_duo.
type Resource struct {
	Name      string
	Args      Args
	state     *vaultIdentityMfaDuoState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (vimd *Resource) Type() string {
	return "vault_identity_mfa_duo"
}

// LocalName returns the local name for [Resource].
func (vimd *Resource) LocalName() string {
	return vimd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (vimd *Resource) Configuration() interface{} {
	return vimd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (vimd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(vimd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (vimd *Resource) Dependencies() terra.Dependencies {
	return vimd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (vimd *Resource) LifecycleManagement() *terra.Lifecycle {
	return vimd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (vimd *Resource) Attributes() vaultIdentityMfaDuoAttributes {
	return vaultIdentityMfaDuoAttributes{ref: terra.ReferenceResource(vimd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (vimd *Resource) ImportState(state io.Reader) error {
	vimd.state = &vaultIdentityMfaDuoState{}
	if err := json.NewDecoder(state).Decode(vimd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vimd.Type(), vimd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (vimd *Resource) State() (*vaultIdentityMfaDuoState, bool) {
	return vimd.state, vimd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (vimd *Resource) StateMust() *vaultIdentityMfaDuoState {
	if vimd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vimd.Type(), vimd.LocalName()))
	}
	return vimd.state
}

// Args contains the configurations for vault_identity_mfa_duo.
type Args struct {
	// ApiHostname: string, required
	ApiHostname terra.StringValue `hcl:"api_hostname,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntegrationKey: string, required
	IntegrationKey terra.StringValue `hcl:"integration_key,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// PushInfo: string, optional
	PushInfo terra.StringValue `hcl:"push_info,attr"`
	// SecretKey: string, required
	SecretKey terra.StringValue `hcl:"secret_key,attr" validate:"required"`
	// UsePasscode: bool, optional
	UsePasscode terra.BoolValue `hcl:"use_passcode,attr"`
	// UsernameFormat: string, optional
	UsernameFormat terra.StringValue `hcl:"username_format,attr"`
}

type vaultIdentityMfaDuoAttributes struct {
	ref terra.Reference
}

// ApiHostname returns a reference to field api_hostname of vault_identity_mfa_duo.
func (vimd vaultIdentityMfaDuoAttributes) ApiHostname() terra.StringValue {
	return terra.ReferenceAsString(vimd.ref.Append("api_hostname"))
}

// Id returns a reference to field id of vault_identity_mfa_duo.
func (vimd vaultIdentityMfaDuoAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vimd.ref.Append("id"))
}

// IntegrationKey returns a reference to field integration_key of vault_identity_mfa_duo.
func (vimd vaultIdentityMfaDuoAttributes) IntegrationKey() terra.StringValue {
	return terra.ReferenceAsString(vimd.ref.Append("integration_key"))
}

// MethodId returns a reference to field method_id of vault_identity_mfa_duo.
func (vimd vaultIdentityMfaDuoAttributes) MethodId() terra.StringValue {
	return terra.ReferenceAsString(vimd.ref.Append("method_id"))
}

// MountAccessor returns a reference to field mount_accessor of vault_identity_mfa_duo.
func (vimd vaultIdentityMfaDuoAttributes) MountAccessor() terra.StringValue {
	return terra.ReferenceAsString(vimd.ref.Append("mount_accessor"))
}

// Name returns a reference to field name of vault_identity_mfa_duo.
func (vimd vaultIdentityMfaDuoAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vimd.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_identity_mfa_duo.
func (vimd vaultIdentityMfaDuoAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vimd.ref.Append("namespace"))
}

// NamespaceId returns a reference to field namespace_id of vault_identity_mfa_duo.
func (vimd vaultIdentityMfaDuoAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(vimd.ref.Append("namespace_id"))
}

// NamespacePath returns a reference to field namespace_path of vault_identity_mfa_duo.
func (vimd vaultIdentityMfaDuoAttributes) NamespacePath() terra.StringValue {
	return terra.ReferenceAsString(vimd.ref.Append("namespace_path"))
}

// PushInfo returns a reference to field push_info of vault_identity_mfa_duo.
func (vimd vaultIdentityMfaDuoAttributes) PushInfo() terra.StringValue {
	return terra.ReferenceAsString(vimd.ref.Append("push_info"))
}

// SecretKey returns a reference to field secret_key of vault_identity_mfa_duo.
func (vimd vaultIdentityMfaDuoAttributes) SecretKey() terra.StringValue {
	return terra.ReferenceAsString(vimd.ref.Append("secret_key"))
}

// Type returns a reference to field type of vault_identity_mfa_duo.
func (vimd vaultIdentityMfaDuoAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vimd.ref.Append("type"))
}

// UsePasscode returns a reference to field use_passcode of vault_identity_mfa_duo.
func (vimd vaultIdentityMfaDuoAttributes) UsePasscode() terra.BoolValue {
	return terra.ReferenceAsBool(vimd.ref.Append("use_passcode"))
}

// UsernameFormat returns a reference to field username_format of vault_identity_mfa_duo.
func (vimd vaultIdentityMfaDuoAttributes) UsernameFormat() terra.StringValue {
	return terra.ReferenceAsString(vimd.ref.Append("username_format"))
}

// Uuid returns a reference to field uuid of vault_identity_mfa_duo.
func (vimd vaultIdentityMfaDuoAttributes) Uuid() terra.StringValue {
	return terra.ReferenceAsString(vimd.ref.Append("uuid"))
}

type vaultIdentityMfaDuoState struct {
	ApiHostname    string `json:"api_hostname"`
	Id             string `json:"id"`
	IntegrationKey string `json:"integration_key"`
	MethodId       string `json:"method_id"`
	MountAccessor  string `json:"mount_accessor"`
	Name           string `json:"name"`
	Namespace      string `json:"namespace"`
	NamespaceId    string `json:"namespace_id"`
	NamespacePath  string `json:"namespace_path"`
	PushInfo       string `json:"push_info"`
	SecretKey      string `json:"secret_key"`
	Type           string `json:"type"`
	UsePasscode    bool   `json:"use_passcode"`
	UsernameFormat string `json:"username_format"`
	Uuid           string `json:"uuid"`
}
