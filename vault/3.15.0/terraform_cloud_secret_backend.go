// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTerraformCloudSecretBackend creates a new instance of [TerraformCloudSecretBackend].
func NewTerraformCloudSecretBackend(name string, args TerraformCloudSecretBackendArgs) *TerraformCloudSecretBackend {
	return &TerraformCloudSecretBackend{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TerraformCloudSecretBackend)(nil)

// TerraformCloudSecretBackend represents the Terraform resource vault_terraform_cloud_secret_backend.
type TerraformCloudSecretBackend struct {
	Name      string
	Args      TerraformCloudSecretBackendArgs
	state     *terraformCloudSecretBackendState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TerraformCloudSecretBackend].
func (tcsb *TerraformCloudSecretBackend) Type() string {
	return "vault_terraform_cloud_secret_backend"
}

// LocalName returns the local name for [TerraformCloudSecretBackend].
func (tcsb *TerraformCloudSecretBackend) LocalName() string {
	return tcsb.Name
}

// Configuration returns the configuration (args) for [TerraformCloudSecretBackend].
func (tcsb *TerraformCloudSecretBackend) Configuration() interface{} {
	return tcsb.Args
}

// DependOn is used for other resources to depend on [TerraformCloudSecretBackend].
func (tcsb *TerraformCloudSecretBackend) DependOn() terra.Reference {
	return terra.ReferenceResource(tcsb)
}

// Dependencies returns the list of resources [TerraformCloudSecretBackend] depends_on.
func (tcsb *TerraformCloudSecretBackend) Dependencies() terra.Dependencies {
	return tcsb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TerraformCloudSecretBackend].
func (tcsb *TerraformCloudSecretBackend) LifecycleManagement() *terra.Lifecycle {
	return tcsb.Lifecycle
}

// Attributes returns the attributes for [TerraformCloudSecretBackend].
func (tcsb *TerraformCloudSecretBackend) Attributes() terraformCloudSecretBackendAttributes {
	return terraformCloudSecretBackendAttributes{ref: terra.ReferenceResource(tcsb)}
}

// ImportState imports the given attribute values into [TerraformCloudSecretBackend]'s state.
func (tcsb *TerraformCloudSecretBackend) ImportState(av io.Reader) error {
	tcsb.state = &terraformCloudSecretBackendState{}
	if err := json.NewDecoder(av).Decode(tcsb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tcsb.Type(), tcsb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TerraformCloudSecretBackend] has state.
func (tcsb *TerraformCloudSecretBackend) State() (*terraformCloudSecretBackendState, bool) {
	return tcsb.state, tcsb.state != nil
}

// StateMust returns the state for [TerraformCloudSecretBackend]. Panics if the state is nil.
func (tcsb *TerraformCloudSecretBackend) StateMust() *terraformCloudSecretBackendState {
	if tcsb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tcsb.Type(), tcsb.LocalName()))
	}
	return tcsb.state
}

// TerraformCloudSecretBackendArgs contains the configurations for vault_terraform_cloud_secret_backend.
type TerraformCloudSecretBackendArgs struct {
	// Address: string, optional
	Address terra.StringValue `hcl:"address,attr"`
	// Backend: string, optional
	Backend terra.StringValue `hcl:"backend,attr"`
	// BasePath: string, optional
	BasePath terra.StringValue `hcl:"base_path,attr"`
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
	// Token: string, optional
	Token terra.StringValue `hcl:"token,attr"`
}
type terraformCloudSecretBackendAttributes struct {
	ref terra.Reference
}

// Address returns a reference to field address of vault_terraform_cloud_secret_backend.
func (tcsb terraformCloudSecretBackendAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(tcsb.ref.Append("address"))
}

// Backend returns a reference to field backend of vault_terraform_cloud_secret_backend.
func (tcsb terraformCloudSecretBackendAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(tcsb.ref.Append("backend"))
}

// BasePath returns a reference to field base_path of vault_terraform_cloud_secret_backend.
func (tcsb terraformCloudSecretBackendAttributes) BasePath() terra.StringValue {
	return terra.ReferenceAsString(tcsb.ref.Append("base_path"))
}

// DefaultLeaseTtlSeconds returns a reference to field default_lease_ttl_seconds of vault_terraform_cloud_secret_backend.
func (tcsb terraformCloudSecretBackendAttributes) DefaultLeaseTtlSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(tcsb.ref.Append("default_lease_ttl_seconds"))
}

// Description returns a reference to field description of vault_terraform_cloud_secret_backend.
func (tcsb terraformCloudSecretBackendAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(tcsb.ref.Append("description"))
}

// DisableRemount returns a reference to field disable_remount of vault_terraform_cloud_secret_backend.
func (tcsb terraformCloudSecretBackendAttributes) DisableRemount() terra.BoolValue {
	return terra.ReferenceAsBool(tcsb.ref.Append("disable_remount"))
}

// Id returns a reference to field id of vault_terraform_cloud_secret_backend.
func (tcsb terraformCloudSecretBackendAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tcsb.ref.Append("id"))
}

// MaxLeaseTtlSeconds returns a reference to field max_lease_ttl_seconds of vault_terraform_cloud_secret_backend.
func (tcsb terraformCloudSecretBackendAttributes) MaxLeaseTtlSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(tcsb.ref.Append("max_lease_ttl_seconds"))
}

// Namespace returns a reference to field namespace of vault_terraform_cloud_secret_backend.
func (tcsb terraformCloudSecretBackendAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(tcsb.ref.Append("namespace"))
}

// Token returns a reference to field token of vault_terraform_cloud_secret_backend.
func (tcsb terraformCloudSecretBackendAttributes) Token() terra.StringValue {
	return terra.ReferenceAsString(tcsb.ref.Append("token"))
}

type terraformCloudSecretBackendState struct {
	Address                string  `json:"address"`
	Backend                string  `json:"backend"`
	BasePath               string  `json:"base_path"`
	DefaultLeaseTtlSeconds float64 `json:"default_lease_ttl_seconds"`
	Description            string  `json:"description"`
	DisableRemount         bool    `json:"disable_remount"`
	Id                     string  `json:"id"`
	MaxLeaseTtlSeconds     float64 `json:"max_lease_ttl_seconds"`
	Namespace              string  `json:"namespace"`
	Token                  string  `json:"token"`
}