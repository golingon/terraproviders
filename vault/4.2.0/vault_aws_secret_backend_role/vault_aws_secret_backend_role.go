// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_aws_secret_backend_role

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

// Resource represents the Terraform resource vault_aws_secret_backend_role.
type Resource struct {
	Name      string
	Args      Args
	state     *vaultAwsSecretBackendRoleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (vasbr *Resource) Type() string {
	return "vault_aws_secret_backend_role"
}

// LocalName returns the local name for [Resource].
func (vasbr *Resource) LocalName() string {
	return vasbr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (vasbr *Resource) Configuration() interface{} {
	return vasbr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (vasbr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(vasbr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (vasbr *Resource) Dependencies() terra.Dependencies {
	return vasbr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (vasbr *Resource) LifecycleManagement() *terra.Lifecycle {
	return vasbr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (vasbr *Resource) Attributes() vaultAwsSecretBackendRoleAttributes {
	return vaultAwsSecretBackendRoleAttributes{ref: terra.ReferenceResource(vasbr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (vasbr *Resource) ImportState(state io.Reader) error {
	vasbr.state = &vaultAwsSecretBackendRoleState{}
	if err := json.NewDecoder(state).Decode(vasbr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vasbr.Type(), vasbr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (vasbr *Resource) State() (*vaultAwsSecretBackendRoleState, bool) {
	return vasbr.state, vasbr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (vasbr *Resource) StateMust() *vaultAwsSecretBackendRoleState {
	if vasbr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vasbr.Type(), vasbr.LocalName()))
	}
	return vasbr.state
}

// Args contains the configurations for vault_aws_secret_backend_role.
type Args struct {
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// CredentialType: string, required
	CredentialType terra.StringValue `hcl:"credential_type,attr" validate:"required"`
	// DefaultStsTtl: number, optional
	DefaultStsTtl terra.NumberValue `hcl:"default_sts_ttl,attr"`
	// IamGroups: set of string, optional
	IamGroups terra.SetValue[terra.StringValue] `hcl:"iam_groups,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxStsTtl: number, optional
	MaxStsTtl terra.NumberValue `hcl:"max_sts_ttl,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// PermissionsBoundaryArn: string, optional
	PermissionsBoundaryArn terra.StringValue `hcl:"permissions_boundary_arn,attr"`
	// PolicyArns: set of string, optional
	PolicyArns terra.SetValue[terra.StringValue] `hcl:"policy_arns,attr"`
	// PolicyDocument: string, optional
	PolicyDocument terra.StringValue `hcl:"policy_document,attr"`
	// RoleArns: set of string, optional
	RoleArns terra.SetValue[terra.StringValue] `hcl:"role_arns,attr"`
	// UserPath: string, optional
	UserPath terra.StringValue `hcl:"user_path,attr"`
}

type vaultAwsSecretBackendRoleAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_aws_secret_backend_role.
func (vasbr vaultAwsSecretBackendRoleAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(vasbr.ref.Append("backend"))
}

// CredentialType returns a reference to field credential_type of vault_aws_secret_backend_role.
func (vasbr vaultAwsSecretBackendRoleAttributes) CredentialType() terra.StringValue {
	return terra.ReferenceAsString(vasbr.ref.Append("credential_type"))
}

// DefaultStsTtl returns a reference to field default_sts_ttl of vault_aws_secret_backend_role.
func (vasbr vaultAwsSecretBackendRoleAttributes) DefaultStsTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(vasbr.ref.Append("default_sts_ttl"))
}

// IamGroups returns a reference to field iam_groups of vault_aws_secret_backend_role.
func (vasbr vaultAwsSecretBackendRoleAttributes) IamGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vasbr.ref.Append("iam_groups"))
}

// Id returns a reference to field id of vault_aws_secret_backend_role.
func (vasbr vaultAwsSecretBackendRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vasbr.ref.Append("id"))
}

// MaxStsTtl returns a reference to field max_sts_ttl of vault_aws_secret_backend_role.
func (vasbr vaultAwsSecretBackendRoleAttributes) MaxStsTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(vasbr.ref.Append("max_sts_ttl"))
}

// Name returns a reference to field name of vault_aws_secret_backend_role.
func (vasbr vaultAwsSecretBackendRoleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vasbr.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_aws_secret_backend_role.
func (vasbr vaultAwsSecretBackendRoleAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vasbr.ref.Append("namespace"))
}

// PermissionsBoundaryArn returns a reference to field permissions_boundary_arn of vault_aws_secret_backend_role.
func (vasbr vaultAwsSecretBackendRoleAttributes) PermissionsBoundaryArn() terra.StringValue {
	return terra.ReferenceAsString(vasbr.ref.Append("permissions_boundary_arn"))
}

// PolicyArns returns a reference to field policy_arns of vault_aws_secret_backend_role.
func (vasbr vaultAwsSecretBackendRoleAttributes) PolicyArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vasbr.ref.Append("policy_arns"))
}

// PolicyDocument returns a reference to field policy_document of vault_aws_secret_backend_role.
func (vasbr vaultAwsSecretBackendRoleAttributes) PolicyDocument() terra.StringValue {
	return terra.ReferenceAsString(vasbr.ref.Append("policy_document"))
}

// RoleArns returns a reference to field role_arns of vault_aws_secret_backend_role.
func (vasbr vaultAwsSecretBackendRoleAttributes) RoleArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vasbr.ref.Append("role_arns"))
}

// UserPath returns a reference to field user_path of vault_aws_secret_backend_role.
func (vasbr vaultAwsSecretBackendRoleAttributes) UserPath() terra.StringValue {
	return terra.ReferenceAsString(vasbr.ref.Append("user_path"))
}

type vaultAwsSecretBackendRoleState struct {
	Backend                string   `json:"backend"`
	CredentialType         string   `json:"credential_type"`
	DefaultStsTtl          float64  `json:"default_sts_ttl"`
	IamGroups              []string `json:"iam_groups"`
	Id                     string   `json:"id"`
	MaxStsTtl              float64  `json:"max_sts_ttl"`
	Name                   string   `json:"name"`
	Namespace              string   `json:"namespace"`
	PermissionsBoundaryArn string   `json:"permissions_boundary_arn"`
	PolicyArns             []string `json:"policy_arns"`
	PolicyDocument         string   `json:"policy_document"`
	RoleArns               []string `json:"role_arns"`
	UserPath               string   `json:"user_path"`
}
