// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRabbitmqSecretBackend creates a new instance of [RabbitmqSecretBackend].
func NewRabbitmqSecretBackend(name string, args RabbitmqSecretBackendArgs) *RabbitmqSecretBackend {
	return &RabbitmqSecretBackend{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RabbitmqSecretBackend)(nil)

// RabbitmqSecretBackend represents the Terraform resource vault_rabbitmq_secret_backend.
type RabbitmqSecretBackend struct {
	Name      string
	Args      RabbitmqSecretBackendArgs
	state     *rabbitmqSecretBackendState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RabbitmqSecretBackend].
func (rsb *RabbitmqSecretBackend) Type() string {
	return "vault_rabbitmq_secret_backend"
}

// LocalName returns the local name for [RabbitmqSecretBackend].
func (rsb *RabbitmqSecretBackend) LocalName() string {
	return rsb.Name
}

// Configuration returns the configuration (args) for [RabbitmqSecretBackend].
func (rsb *RabbitmqSecretBackend) Configuration() interface{} {
	return rsb.Args
}

// DependOn is used for other resources to depend on [RabbitmqSecretBackend].
func (rsb *RabbitmqSecretBackend) DependOn() terra.Reference {
	return terra.ReferenceResource(rsb)
}

// Dependencies returns the list of resources [RabbitmqSecretBackend] depends_on.
func (rsb *RabbitmqSecretBackend) Dependencies() terra.Dependencies {
	return rsb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RabbitmqSecretBackend].
func (rsb *RabbitmqSecretBackend) LifecycleManagement() *terra.Lifecycle {
	return rsb.Lifecycle
}

// Attributes returns the attributes for [RabbitmqSecretBackend].
func (rsb *RabbitmqSecretBackend) Attributes() rabbitmqSecretBackendAttributes {
	return rabbitmqSecretBackendAttributes{ref: terra.ReferenceResource(rsb)}
}

// ImportState imports the given attribute values into [RabbitmqSecretBackend]'s state.
func (rsb *RabbitmqSecretBackend) ImportState(av io.Reader) error {
	rsb.state = &rabbitmqSecretBackendState{}
	if err := json.NewDecoder(av).Decode(rsb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rsb.Type(), rsb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RabbitmqSecretBackend] has state.
func (rsb *RabbitmqSecretBackend) State() (*rabbitmqSecretBackendState, bool) {
	return rsb.state, rsb.state != nil
}

// StateMust returns the state for [RabbitmqSecretBackend]. Panics if the state is nil.
func (rsb *RabbitmqSecretBackend) StateMust() *rabbitmqSecretBackendState {
	if rsb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rsb.Type(), rsb.LocalName()))
	}
	return rsb.state
}

// RabbitmqSecretBackendArgs contains the configurations for vault_rabbitmq_secret_backend.
type RabbitmqSecretBackendArgs struct {
	// ConnectionUri: string, required
	ConnectionUri terra.StringValue `hcl:"connection_uri,attr" validate:"required"`
	// DefaultLeaseTtlSeconds: number, optional
	DefaultLeaseTtlSeconds terra.NumberValue `hcl:"default_lease_ttl_seconds,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisableRemount: bool, optional
	DisableRemount terra.BoolValue `hcl:"disable_remount,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxLeaseTtlSeconds: number, optional
	MaxLeaseTtlSeconds terra.NumberValue `hcl:"max_lease_ttl_seconds,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// PasswordPolicy: string, optional
	PasswordPolicy terra.StringValue `hcl:"password_policy,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
	// UsernameTemplate: string, optional
	UsernameTemplate terra.StringValue `hcl:"username_template,attr"`
	// VerifyConnection: bool, optional
	VerifyConnection terra.BoolValue `hcl:"verify_connection,attr"`
}
type rabbitmqSecretBackendAttributes struct {
	ref terra.Reference
}

// ConnectionUri returns a reference to field connection_uri of vault_rabbitmq_secret_backend.
func (rsb rabbitmqSecretBackendAttributes) ConnectionUri() terra.StringValue {
	return terra.ReferenceAsString(rsb.ref.Append("connection_uri"))
}

// DefaultLeaseTtlSeconds returns a reference to field default_lease_ttl_seconds of vault_rabbitmq_secret_backend.
func (rsb rabbitmqSecretBackendAttributes) DefaultLeaseTtlSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(rsb.ref.Append("default_lease_ttl_seconds"))
}

// Description returns a reference to field description of vault_rabbitmq_secret_backend.
func (rsb rabbitmqSecretBackendAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(rsb.ref.Append("description"))
}

// DisableRemount returns a reference to field disable_remount of vault_rabbitmq_secret_backend.
func (rsb rabbitmqSecretBackendAttributes) DisableRemount() terra.BoolValue {
	return terra.ReferenceAsBool(rsb.ref.Append("disable_remount"))
}

// Id returns a reference to field id of vault_rabbitmq_secret_backend.
func (rsb rabbitmqSecretBackendAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rsb.ref.Append("id"))
}

// MaxLeaseTtlSeconds returns a reference to field max_lease_ttl_seconds of vault_rabbitmq_secret_backend.
func (rsb rabbitmqSecretBackendAttributes) MaxLeaseTtlSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(rsb.ref.Append("max_lease_ttl_seconds"))
}

// Namespace returns a reference to field namespace of vault_rabbitmq_secret_backend.
func (rsb rabbitmqSecretBackendAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(rsb.ref.Append("namespace"))
}

// Password returns a reference to field password of vault_rabbitmq_secret_backend.
func (rsb rabbitmqSecretBackendAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(rsb.ref.Append("password"))
}

// PasswordPolicy returns a reference to field password_policy of vault_rabbitmq_secret_backend.
func (rsb rabbitmqSecretBackendAttributes) PasswordPolicy() terra.StringValue {
	return terra.ReferenceAsString(rsb.ref.Append("password_policy"))
}

// Path returns a reference to field path of vault_rabbitmq_secret_backend.
func (rsb rabbitmqSecretBackendAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(rsb.ref.Append("path"))
}

// Username returns a reference to field username of vault_rabbitmq_secret_backend.
func (rsb rabbitmqSecretBackendAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(rsb.ref.Append("username"))
}

// UsernameTemplate returns a reference to field username_template of vault_rabbitmq_secret_backend.
func (rsb rabbitmqSecretBackendAttributes) UsernameTemplate() terra.StringValue {
	return terra.ReferenceAsString(rsb.ref.Append("username_template"))
}

// VerifyConnection returns a reference to field verify_connection of vault_rabbitmq_secret_backend.
func (rsb rabbitmqSecretBackendAttributes) VerifyConnection() terra.BoolValue {
	return terra.ReferenceAsBool(rsb.ref.Append("verify_connection"))
}

type rabbitmqSecretBackendState struct {
	ConnectionUri          string  `json:"connection_uri"`
	DefaultLeaseTtlSeconds float64 `json:"default_lease_ttl_seconds"`
	Description            string  `json:"description"`
	DisableRemount         bool    `json:"disable_remount"`
	Id                     string  `json:"id"`
	MaxLeaseTtlSeconds     float64 `json:"max_lease_ttl_seconds"`
	Namespace              string  `json:"namespace"`
	Password               string  `json:"password"`
	PasswordPolicy         string  `json:"password_policy"`
	Path                   string  `json:"path"`
	Username               string  `json:"username"`
	UsernameTemplate       string  `json:"username_template"`
	VerifyConnection       bool    `json:"verify_connection"`
}