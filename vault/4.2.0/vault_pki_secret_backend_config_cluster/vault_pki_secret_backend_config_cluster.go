// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_pki_secret_backend_config_cluster

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

// Resource represents the Terraform resource vault_pki_secret_backend_config_cluster.
type Resource struct {
	Name      string
	Args      Args
	state     *vaultPkiSecretBackendConfigClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (vpsbcc *Resource) Type() string {
	return "vault_pki_secret_backend_config_cluster"
}

// LocalName returns the local name for [Resource].
func (vpsbcc *Resource) LocalName() string {
	return vpsbcc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (vpsbcc *Resource) Configuration() interface{} {
	return vpsbcc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (vpsbcc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(vpsbcc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (vpsbcc *Resource) Dependencies() terra.Dependencies {
	return vpsbcc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (vpsbcc *Resource) LifecycleManagement() *terra.Lifecycle {
	return vpsbcc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (vpsbcc *Resource) Attributes() vaultPkiSecretBackendConfigClusterAttributes {
	return vaultPkiSecretBackendConfigClusterAttributes{ref: terra.ReferenceResource(vpsbcc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (vpsbcc *Resource) ImportState(state io.Reader) error {
	vpsbcc.state = &vaultPkiSecretBackendConfigClusterState{}
	if err := json.NewDecoder(state).Decode(vpsbcc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vpsbcc.Type(), vpsbcc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (vpsbcc *Resource) State() (*vaultPkiSecretBackendConfigClusterState, bool) {
	return vpsbcc.state, vpsbcc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (vpsbcc *Resource) StateMust() *vaultPkiSecretBackendConfigClusterState {
	if vpsbcc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vpsbcc.Type(), vpsbcc.LocalName()))
	}
	return vpsbcc.state
}

// Args contains the configurations for vault_pki_secret_backend_config_cluster.
type Args struct {
	// AiaPath: string, optional
	AiaPath terra.StringValue `hcl:"aia_path,attr"`
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
}

type vaultPkiSecretBackendConfigClusterAttributes struct {
	ref terra.Reference
}

// AiaPath returns a reference to field aia_path of vault_pki_secret_backend_config_cluster.
func (vpsbcc vaultPkiSecretBackendConfigClusterAttributes) AiaPath() terra.StringValue {
	return terra.ReferenceAsString(vpsbcc.ref.Append("aia_path"))
}

// Backend returns a reference to field backend of vault_pki_secret_backend_config_cluster.
func (vpsbcc vaultPkiSecretBackendConfigClusterAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(vpsbcc.ref.Append("backend"))
}

// Id returns a reference to field id of vault_pki_secret_backend_config_cluster.
func (vpsbcc vaultPkiSecretBackendConfigClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vpsbcc.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_pki_secret_backend_config_cluster.
func (vpsbcc vaultPkiSecretBackendConfigClusterAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vpsbcc.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_pki_secret_backend_config_cluster.
func (vpsbcc vaultPkiSecretBackendConfigClusterAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(vpsbcc.ref.Append("path"))
}

type vaultPkiSecretBackendConfigClusterState struct {
	AiaPath   string `json:"aia_path"`
	Backend   string `json:"backend"`
	Id        string `json:"id"`
	Namespace string `json:"namespace"`
	Path      string `json:"path"`
}
