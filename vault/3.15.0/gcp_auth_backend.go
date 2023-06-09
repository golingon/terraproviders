// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	gcpauthbackend "github.com/golingon/terraproviders/vault/3.15.0/gcpauthbackend"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGcpAuthBackend creates a new instance of [GcpAuthBackend].
func NewGcpAuthBackend(name string, args GcpAuthBackendArgs) *GcpAuthBackend {
	return &GcpAuthBackend{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GcpAuthBackend)(nil)

// GcpAuthBackend represents the Terraform resource vault_gcp_auth_backend.
type GcpAuthBackend struct {
	Name      string
	Args      GcpAuthBackendArgs
	state     *gcpAuthBackendState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GcpAuthBackend].
func (gab *GcpAuthBackend) Type() string {
	return "vault_gcp_auth_backend"
}

// LocalName returns the local name for [GcpAuthBackend].
func (gab *GcpAuthBackend) LocalName() string {
	return gab.Name
}

// Configuration returns the configuration (args) for [GcpAuthBackend].
func (gab *GcpAuthBackend) Configuration() interface{} {
	return gab.Args
}

// DependOn is used for other resources to depend on [GcpAuthBackend].
func (gab *GcpAuthBackend) DependOn() terra.Reference {
	return terra.ReferenceResource(gab)
}

// Dependencies returns the list of resources [GcpAuthBackend] depends_on.
func (gab *GcpAuthBackend) Dependencies() terra.Dependencies {
	return gab.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GcpAuthBackend].
func (gab *GcpAuthBackend) LifecycleManagement() *terra.Lifecycle {
	return gab.Lifecycle
}

// Attributes returns the attributes for [GcpAuthBackend].
func (gab *GcpAuthBackend) Attributes() gcpAuthBackendAttributes {
	return gcpAuthBackendAttributes{ref: terra.ReferenceResource(gab)}
}

// ImportState imports the given attribute values into [GcpAuthBackend]'s state.
func (gab *GcpAuthBackend) ImportState(av io.Reader) error {
	gab.state = &gcpAuthBackendState{}
	if err := json.NewDecoder(av).Decode(gab.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gab.Type(), gab.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GcpAuthBackend] has state.
func (gab *GcpAuthBackend) State() (*gcpAuthBackendState, bool) {
	return gab.state, gab.state != nil
}

// StateMust returns the state for [GcpAuthBackend]. Panics if the state is nil.
func (gab *GcpAuthBackend) StateMust() *gcpAuthBackendState {
	if gab.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gab.Type(), gab.LocalName()))
	}
	return gab.state
}

// GcpAuthBackendArgs contains the configurations for vault_gcp_auth_backend.
type GcpAuthBackendArgs struct {
	// ClientEmail: string, optional
	ClientEmail terra.StringValue `hcl:"client_email,attr"`
	// ClientId: string, optional
	ClientId terra.StringValue `hcl:"client_id,attr"`
	// Credentials: string, optional
	Credentials terra.StringValue `hcl:"credentials,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisableRemount: bool, optional
	DisableRemount terra.BoolValue `hcl:"disable_remount,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Local: bool, optional
	Local terra.BoolValue `hcl:"local,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// PrivateKeyId: string, optional
	PrivateKeyId terra.StringValue `hcl:"private_key_id,attr"`
	// ProjectId: string, optional
	ProjectId terra.StringValue `hcl:"project_id,attr"`
	// CustomEndpoint: optional
	CustomEndpoint *gcpauthbackend.CustomEndpoint `hcl:"custom_endpoint,block"`
}
type gcpAuthBackendAttributes struct {
	ref terra.Reference
}

// ClientEmail returns a reference to field client_email of vault_gcp_auth_backend.
func (gab gcpAuthBackendAttributes) ClientEmail() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("client_email"))
}

// ClientId returns a reference to field client_id of vault_gcp_auth_backend.
func (gab gcpAuthBackendAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("client_id"))
}

// Credentials returns a reference to field credentials of vault_gcp_auth_backend.
func (gab gcpAuthBackendAttributes) Credentials() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("credentials"))
}

// Description returns a reference to field description of vault_gcp_auth_backend.
func (gab gcpAuthBackendAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("description"))
}

// DisableRemount returns a reference to field disable_remount of vault_gcp_auth_backend.
func (gab gcpAuthBackendAttributes) DisableRemount() terra.BoolValue {
	return terra.ReferenceAsBool(gab.ref.Append("disable_remount"))
}

// Id returns a reference to field id of vault_gcp_auth_backend.
func (gab gcpAuthBackendAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("id"))
}

// Local returns a reference to field local of vault_gcp_auth_backend.
func (gab gcpAuthBackendAttributes) Local() terra.BoolValue {
	return terra.ReferenceAsBool(gab.ref.Append("local"))
}

// Namespace returns a reference to field namespace of vault_gcp_auth_backend.
func (gab gcpAuthBackendAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_gcp_auth_backend.
func (gab gcpAuthBackendAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("path"))
}

// PrivateKeyId returns a reference to field private_key_id of vault_gcp_auth_backend.
func (gab gcpAuthBackendAttributes) PrivateKeyId() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("private_key_id"))
}

// ProjectId returns a reference to field project_id of vault_gcp_auth_backend.
func (gab gcpAuthBackendAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceAsString(gab.ref.Append("project_id"))
}

func (gab gcpAuthBackendAttributes) CustomEndpoint() terra.ListValue[gcpauthbackend.CustomEndpointAttributes] {
	return terra.ReferenceAsList[gcpauthbackend.CustomEndpointAttributes](gab.ref.Append("custom_endpoint"))
}

type gcpAuthBackendState struct {
	ClientEmail    string                               `json:"client_email"`
	ClientId       string                               `json:"client_id"`
	Credentials    string                               `json:"credentials"`
	Description    string                               `json:"description"`
	DisableRemount bool                                 `json:"disable_remount"`
	Id             string                               `json:"id"`
	Local          bool                                 `json:"local"`
	Namespace      string                               `json:"namespace"`
	Path           string                               `json:"path"`
	PrivateKeyId   string                               `json:"private_key_id"`
	ProjectId      string                               `json:"project_id"`
	CustomEndpoint []gcpauthbackend.CustomEndpointState `json:"custom_endpoint"`
}
