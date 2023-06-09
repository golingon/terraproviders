// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKmipSecretBackend creates a new instance of [KmipSecretBackend].
func NewKmipSecretBackend(name string, args KmipSecretBackendArgs) *KmipSecretBackend {
	return &KmipSecretBackend{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KmipSecretBackend)(nil)

// KmipSecretBackend represents the Terraform resource vault_kmip_secret_backend.
type KmipSecretBackend struct {
	Name      string
	Args      KmipSecretBackendArgs
	state     *kmipSecretBackendState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KmipSecretBackend].
func (ksb *KmipSecretBackend) Type() string {
	return "vault_kmip_secret_backend"
}

// LocalName returns the local name for [KmipSecretBackend].
func (ksb *KmipSecretBackend) LocalName() string {
	return ksb.Name
}

// Configuration returns the configuration (args) for [KmipSecretBackend].
func (ksb *KmipSecretBackend) Configuration() interface{} {
	return ksb.Args
}

// DependOn is used for other resources to depend on [KmipSecretBackend].
func (ksb *KmipSecretBackend) DependOn() terra.Reference {
	return terra.ReferenceResource(ksb)
}

// Dependencies returns the list of resources [KmipSecretBackend] depends_on.
func (ksb *KmipSecretBackend) Dependencies() terra.Dependencies {
	return ksb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KmipSecretBackend].
func (ksb *KmipSecretBackend) LifecycleManagement() *terra.Lifecycle {
	return ksb.Lifecycle
}

// Attributes returns the attributes for [KmipSecretBackend].
func (ksb *KmipSecretBackend) Attributes() kmipSecretBackendAttributes {
	return kmipSecretBackendAttributes{ref: terra.ReferenceResource(ksb)}
}

// ImportState imports the given attribute values into [KmipSecretBackend]'s state.
func (ksb *KmipSecretBackend) ImportState(av io.Reader) error {
	ksb.state = &kmipSecretBackendState{}
	if err := json.NewDecoder(av).Decode(ksb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ksb.Type(), ksb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KmipSecretBackend] has state.
func (ksb *KmipSecretBackend) State() (*kmipSecretBackendState, bool) {
	return ksb.state, ksb.state != nil
}

// StateMust returns the state for [KmipSecretBackend]. Panics if the state is nil.
func (ksb *KmipSecretBackend) StateMust() *kmipSecretBackendState {
	if ksb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ksb.Type(), ksb.LocalName()))
	}
	return ksb.state
}

// KmipSecretBackendArgs contains the configurations for vault_kmip_secret_backend.
type KmipSecretBackendArgs struct {
	// DefaultTlsClientKeyBits: number, optional
	DefaultTlsClientKeyBits terra.NumberValue `hcl:"default_tls_client_key_bits,attr"`
	// DefaultTlsClientKeyType: string, optional
	DefaultTlsClientKeyType terra.StringValue `hcl:"default_tls_client_key_type,attr"`
	// DefaultTlsClientTtl: number, optional
	DefaultTlsClientTtl terra.NumberValue `hcl:"default_tls_client_ttl,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisableRemount: bool, optional
	DisableRemount terra.BoolValue `hcl:"disable_remount,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ListenAddrs: set of string, optional
	ListenAddrs terra.SetValue[terra.StringValue] `hcl:"listen_addrs,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
	// ServerHostnames: set of string, optional
	ServerHostnames terra.SetValue[terra.StringValue] `hcl:"server_hostnames,attr"`
	// ServerIps: set of string, optional
	ServerIps terra.SetValue[terra.StringValue] `hcl:"server_ips,attr"`
	// TlsCaKeyBits: number, optional
	TlsCaKeyBits terra.NumberValue `hcl:"tls_ca_key_bits,attr"`
	// TlsCaKeyType: string, optional
	TlsCaKeyType terra.StringValue `hcl:"tls_ca_key_type,attr"`
	// TlsMinVersion: string, optional
	TlsMinVersion terra.StringValue `hcl:"tls_min_version,attr"`
}
type kmipSecretBackendAttributes struct {
	ref terra.Reference
}

// DefaultTlsClientKeyBits returns a reference to field default_tls_client_key_bits of vault_kmip_secret_backend.
func (ksb kmipSecretBackendAttributes) DefaultTlsClientKeyBits() terra.NumberValue {
	return terra.ReferenceAsNumber(ksb.ref.Append("default_tls_client_key_bits"))
}

// DefaultTlsClientKeyType returns a reference to field default_tls_client_key_type of vault_kmip_secret_backend.
func (ksb kmipSecretBackendAttributes) DefaultTlsClientKeyType() terra.StringValue {
	return terra.ReferenceAsString(ksb.ref.Append("default_tls_client_key_type"))
}

// DefaultTlsClientTtl returns a reference to field default_tls_client_ttl of vault_kmip_secret_backend.
func (ksb kmipSecretBackendAttributes) DefaultTlsClientTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(ksb.ref.Append("default_tls_client_ttl"))
}

// Description returns a reference to field description of vault_kmip_secret_backend.
func (ksb kmipSecretBackendAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ksb.ref.Append("description"))
}

// DisableRemount returns a reference to field disable_remount of vault_kmip_secret_backend.
func (ksb kmipSecretBackendAttributes) DisableRemount() terra.BoolValue {
	return terra.ReferenceAsBool(ksb.ref.Append("disable_remount"))
}

// Id returns a reference to field id of vault_kmip_secret_backend.
func (ksb kmipSecretBackendAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ksb.ref.Append("id"))
}

// ListenAddrs returns a reference to field listen_addrs of vault_kmip_secret_backend.
func (ksb kmipSecretBackendAttributes) ListenAddrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ksb.ref.Append("listen_addrs"))
}

// Namespace returns a reference to field namespace of vault_kmip_secret_backend.
func (ksb kmipSecretBackendAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(ksb.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_kmip_secret_backend.
func (ksb kmipSecretBackendAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(ksb.ref.Append("path"))
}

// ServerHostnames returns a reference to field server_hostnames of vault_kmip_secret_backend.
func (ksb kmipSecretBackendAttributes) ServerHostnames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ksb.ref.Append("server_hostnames"))
}

// ServerIps returns a reference to field server_ips of vault_kmip_secret_backend.
func (ksb kmipSecretBackendAttributes) ServerIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ksb.ref.Append("server_ips"))
}

// TlsCaKeyBits returns a reference to field tls_ca_key_bits of vault_kmip_secret_backend.
func (ksb kmipSecretBackendAttributes) TlsCaKeyBits() terra.NumberValue {
	return terra.ReferenceAsNumber(ksb.ref.Append("tls_ca_key_bits"))
}

// TlsCaKeyType returns a reference to field tls_ca_key_type of vault_kmip_secret_backend.
func (ksb kmipSecretBackendAttributes) TlsCaKeyType() terra.StringValue {
	return terra.ReferenceAsString(ksb.ref.Append("tls_ca_key_type"))
}

// TlsMinVersion returns a reference to field tls_min_version of vault_kmip_secret_backend.
func (ksb kmipSecretBackendAttributes) TlsMinVersion() terra.StringValue {
	return terra.ReferenceAsString(ksb.ref.Append("tls_min_version"))
}

type kmipSecretBackendState struct {
	DefaultTlsClientKeyBits float64  `json:"default_tls_client_key_bits"`
	DefaultTlsClientKeyType string   `json:"default_tls_client_key_type"`
	DefaultTlsClientTtl     float64  `json:"default_tls_client_ttl"`
	Description             string   `json:"description"`
	DisableRemount          bool     `json:"disable_remount"`
	Id                      string   `json:"id"`
	ListenAddrs             []string `json:"listen_addrs"`
	Namespace               string   `json:"namespace"`
	Path                    string   `json:"path"`
	ServerHostnames         []string `json:"server_hostnames"`
	ServerIps               []string `json:"server_ips"`
	TlsCaKeyBits            float64  `json:"tls_ca_key_bits"`
	TlsCaKeyType            string   `json:"tls_ca_key_type"`
	TlsMinVersion           string   `json:"tls_min_version"`
}
