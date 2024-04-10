// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewTerraformCloudSecretRole creates a new instance of [TerraformCloudSecretRole].
func NewTerraformCloudSecretRole(name string, args TerraformCloudSecretRoleArgs) *TerraformCloudSecretRole {
	return &TerraformCloudSecretRole{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TerraformCloudSecretRole)(nil)

// TerraformCloudSecretRole represents the Terraform resource vault_terraform_cloud_secret_role.
type TerraformCloudSecretRole struct {
	Name      string
	Args      TerraformCloudSecretRoleArgs
	state     *terraformCloudSecretRoleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TerraformCloudSecretRole].
func (tcsr *TerraformCloudSecretRole) Type() string {
	return "vault_terraform_cloud_secret_role"
}

// LocalName returns the local name for [TerraformCloudSecretRole].
func (tcsr *TerraformCloudSecretRole) LocalName() string {
	return tcsr.Name
}

// Configuration returns the configuration (args) for [TerraformCloudSecretRole].
func (tcsr *TerraformCloudSecretRole) Configuration() interface{} {
	return tcsr.Args
}

// DependOn is used for other resources to depend on [TerraformCloudSecretRole].
func (tcsr *TerraformCloudSecretRole) DependOn() terra.Reference {
	return terra.ReferenceResource(tcsr)
}

// Dependencies returns the list of resources [TerraformCloudSecretRole] depends_on.
func (tcsr *TerraformCloudSecretRole) Dependencies() terra.Dependencies {
	return tcsr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TerraformCloudSecretRole].
func (tcsr *TerraformCloudSecretRole) LifecycleManagement() *terra.Lifecycle {
	return tcsr.Lifecycle
}

// Attributes returns the attributes for [TerraformCloudSecretRole].
func (tcsr *TerraformCloudSecretRole) Attributes() terraformCloudSecretRoleAttributes {
	return terraformCloudSecretRoleAttributes{ref: terra.ReferenceResource(tcsr)}
}

// ImportState imports the given attribute values into [TerraformCloudSecretRole]'s state.
func (tcsr *TerraformCloudSecretRole) ImportState(av io.Reader) error {
	tcsr.state = &terraformCloudSecretRoleState{}
	if err := json.NewDecoder(av).Decode(tcsr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tcsr.Type(), tcsr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TerraformCloudSecretRole] has state.
func (tcsr *TerraformCloudSecretRole) State() (*terraformCloudSecretRoleState, bool) {
	return tcsr.state, tcsr.state != nil
}

// StateMust returns the state for [TerraformCloudSecretRole]. Panics if the state is nil.
func (tcsr *TerraformCloudSecretRole) StateMust() *terraformCloudSecretRoleState {
	if tcsr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tcsr.Type(), tcsr.LocalName()))
	}
	return tcsr.state
}

// TerraformCloudSecretRoleArgs contains the configurations for vault_terraform_cloud_secret_role.
type TerraformCloudSecretRoleArgs struct {
	// Backend: string, optional
	Backend terra.StringValue `hcl:"backend,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxTtl: number, optional
	MaxTtl terra.NumberValue `hcl:"max_ttl,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Organization: string, optional
	Organization terra.StringValue `hcl:"organization,attr"`
	// TeamId: string, optional
	TeamId terra.StringValue `hcl:"team_id,attr"`
	// Ttl: number, optional
	Ttl terra.NumberValue `hcl:"ttl,attr"`
	// UserId: string, optional
	UserId terra.StringValue `hcl:"user_id,attr"`
}
type terraformCloudSecretRoleAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_terraform_cloud_secret_role.
func (tcsr terraformCloudSecretRoleAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(tcsr.ref.Append("backend"))
}

// Id returns a reference to field id of vault_terraform_cloud_secret_role.
func (tcsr terraformCloudSecretRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tcsr.ref.Append("id"))
}

// MaxTtl returns a reference to field max_ttl of vault_terraform_cloud_secret_role.
func (tcsr terraformCloudSecretRoleAttributes) MaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(tcsr.ref.Append("max_ttl"))
}

// Name returns a reference to field name of vault_terraform_cloud_secret_role.
func (tcsr terraformCloudSecretRoleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tcsr.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_terraform_cloud_secret_role.
func (tcsr terraformCloudSecretRoleAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(tcsr.ref.Append("namespace"))
}

// Organization returns a reference to field organization of vault_terraform_cloud_secret_role.
func (tcsr terraformCloudSecretRoleAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(tcsr.ref.Append("organization"))
}

// TeamId returns a reference to field team_id of vault_terraform_cloud_secret_role.
func (tcsr terraformCloudSecretRoleAttributes) TeamId() terra.StringValue {
	return terra.ReferenceAsString(tcsr.ref.Append("team_id"))
}

// Ttl returns a reference to field ttl of vault_terraform_cloud_secret_role.
func (tcsr terraformCloudSecretRoleAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(tcsr.ref.Append("ttl"))
}

// UserId returns a reference to field user_id of vault_terraform_cloud_secret_role.
func (tcsr terraformCloudSecretRoleAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(tcsr.ref.Append("user_id"))
}

type terraformCloudSecretRoleState struct {
	Backend      string  `json:"backend"`
	Id           string  `json:"id"`
	MaxTtl       float64 `json:"max_ttl"`
	Name         string  `json:"name"`
	Namespace    string  `json:"namespace"`
	Organization string  `json:"organization"`
	TeamId       string  `json:"team_id"`
	Ttl          float64 `json:"ttl"`
	UserId       string  `json:"user_id"`
}
