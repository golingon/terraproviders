// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import "github.com/golingon/lingon/pkg/terra"

// NewDataKubernetesAuthBackendRole creates a new instance of [DataKubernetesAuthBackendRole].
func NewDataKubernetesAuthBackendRole(name string, args DataKubernetesAuthBackendRoleArgs) *DataKubernetesAuthBackendRole {
	return &DataKubernetesAuthBackendRole{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKubernetesAuthBackendRole)(nil)

// DataKubernetesAuthBackendRole represents the Terraform data resource vault_kubernetes_auth_backend_role.
type DataKubernetesAuthBackendRole struct {
	Name string
	Args DataKubernetesAuthBackendRoleArgs
}

// DataSource returns the Terraform object type for [DataKubernetesAuthBackendRole].
func (kabr *DataKubernetesAuthBackendRole) DataSource() string {
	return "vault_kubernetes_auth_backend_role"
}

// LocalName returns the local name for [DataKubernetesAuthBackendRole].
func (kabr *DataKubernetesAuthBackendRole) LocalName() string {
	return kabr.Name
}

// Configuration returns the configuration (args) for [DataKubernetesAuthBackendRole].
func (kabr *DataKubernetesAuthBackendRole) Configuration() interface{} {
	return kabr.Args
}

// Attributes returns the attributes for [DataKubernetesAuthBackendRole].
func (kabr *DataKubernetesAuthBackendRole) Attributes() dataKubernetesAuthBackendRoleAttributes {
	return dataKubernetesAuthBackendRoleAttributes{ref: terra.ReferenceDataResource(kabr)}
}

// DataKubernetesAuthBackendRoleArgs contains the configurations for vault_kubernetes_auth_backend_role.
type DataKubernetesAuthBackendRoleArgs struct {
	// Audience: string, optional
	Audience terra.StringValue `hcl:"audience,attr"`
	// Backend: string, optional
	Backend terra.StringValue `hcl:"backend,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// RoleName: string, required
	RoleName terra.StringValue `hcl:"role_name,attr" validate:"required"`
	// TokenBoundCidrs: set of string, optional
	TokenBoundCidrs terra.SetValue[terra.StringValue] `hcl:"token_bound_cidrs,attr"`
	// TokenExplicitMaxTtl: number, optional
	TokenExplicitMaxTtl terra.NumberValue `hcl:"token_explicit_max_ttl,attr"`
	// TokenMaxTtl: number, optional
	TokenMaxTtl terra.NumberValue `hcl:"token_max_ttl,attr"`
	// TokenNoDefaultPolicy: bool, optional
	TokenNoDefaultPolicy terra.BoolValue `hcl:"token_no_default_policy,attr"`
	// TokenNumUses: number, optional
	TokenNumUses terra.NumberValue `hcl:"token_num_uses,attr"`
	// TokenPeriod: number, optional
	TokenPeriod terra.NumberValue `hcl:"token_period,attr"`
	// TokenPolicies: set of string, optional
	TokenPolicies terra.SetValue[terra.StringValue] `hcl:"token_policies,attr"`
	// TokenTtl: number, optional
	TokenTtl terra.NumberValue `hcl:"token_ttl,attr"`
	// TokenType: string, optional
	TokenType terra.StringValue `hcl:"token_type,attr"`
}
type dataKubernetesAuthBackendRoleAttributes struct {
	ref terra.Reference
}

// AliasNameSource returns a reference to field alias_name_source of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) AliasNameSource() terra.StringValue {
	return terra.ReferenceAsString(kabr.ref.Append("alias_name_source"))
}

// Audience returns a reference to field audience of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) Audience() terra.StringValue {
	return terra.ReferenceAsString(kabr.ref.Append("audience"))
}

// Backend returns a reference to field backend of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(kabr.ref.Append("backend"))
}

// BoundServiceAccountNames returns a reference to field bound_service_account_names of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) BoundServiceAccountNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kabr.ref.Append("bound_service_account_names"))
}

// BoundServiceAccountNamespaces returns a reference to field bound_service_account_namespaces of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) BoundServiceAccountNamespaces() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kabr.ref.Append("bound_service_account_namespaces"))
}

// Id returns a reference to field id of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kabr.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(kabr.ref.Append("namespace"))
}

// RoleName returns a reference to field role_name of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) RoleName() terra.StringValue {
	return terra.ReferenceAsString(kabr.ref.Append("role_name"))
}

// TokenBoundCidrs returns a reference to field token_bound_cidrs of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) TokenBoundCidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kabr.ref.Append("token_bound_cidrs"))
}

// TokenExplicitMaxTtl returns a reference to field token_explicit_max_ttl of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) TokenExplicitMaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(kabr.ref.Append("token_explicit_max_ttl"))
}

// TokenMaxTtl returns a reference to field token_max_ttl of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) TokenMaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(kabr.ref.Append("token_max_ttl"))
}

// TokenNoDefaultPolicy returns a reference to field token_no_default_policy of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) TokenNoDefaultPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(kabr.ref.Append("token_no_default_policy"))
}

// TokenNumUses returns a reference to field token_num_uses of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) TokenNumUses() terra.NumberValue {
	return terra.ReferenceAsNumber(kabr.ref.Append("token_num_uses"))
}

// TokenPeriod returns a reference to field token_period of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) TokenPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(kabr.ref.Append("token_period"))
}

// TokenPolicies returns a reference to field token_policies of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) TokenPolicies() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kabr.ref.Append("token_policies"))
}

// TokenTtl returns a reference to field token_ttl of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) TokenTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(kabr.ref.Append("token_ttl"))
}

// TokenType returns a reference to field token_type of vault_kubernetes_auth_backend_role.
func (kabr dataKubernetesAuthBackendRoleAttributes) TokenType() terra.StringValue {
	return terra.ReferenceAsString(kabr.ref.Append("token_type"))
}
