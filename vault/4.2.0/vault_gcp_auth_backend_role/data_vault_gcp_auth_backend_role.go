// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_gcp_auth_backend_role

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource vault_gcp_auth_backend_role.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (vgabr *DataSource) DataSource() string {
	return "vault_gcp_auth_backend_role"
}

// LocalName returns the local name for [DataSource].
func (vgabr *DataSource) LocalName() string {
	return vgabr.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (vgabr *DataSource) Configuration() interface{} {
	return vgabr.Args
}

// Attributes returns the attributes for [DataSource].
func (vgabr *DataSource) Attributes() dataVaultGcpAuthBackendRoleAttributes {
	return dataVaultGcpAuthBackendRoleAttributes{ref: terra.ReferenceDataSource(vgabr)}
}

// DataArgs contains the configurations for vault_gcp_auth_backend_role.
type DataArgs struct {
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

type dataVaultGcpAuthBackendRoleAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(vgabr.ref.Append("backend"))
}

// BoundInstanceGroups returns a reference to field bound_instance_groups of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) BoundInstanceGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgabr.ref.Append("bound_instance_groups"))
}

// BoundLabels returns a reference to field bound_labels of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) BoundLabels() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgabr.ref.Append("bound_labels"))
}

// BoundProjects returns a reference to field bound_projects of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) BoundProjects() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgabr.ref.Append("bound_projects"))
}

// BoundRegions returns a reference to field bound_regions of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) BoundRegions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgabr.ref.Append("bound_regions"))
}

// BoundServiceAccounts returns a reference to field bound_service_accounts of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) BoundServiceAccounts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgabr.ref.Append("bound_service_accounts"))
}

// BoundZones returns a reference to field bound_zones of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) BoundZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgabr.ref.Append("bound_zones"))
}

// Id returns a reference to field id of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vgabr.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vgabr.ref.Append("namespace"))
}

// RoleId returns a reference to field role_id of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) RoleId() terra.StringValue {
	return terra.ReferenceAsString(vgabr.ref.Append("role_id"))
}

// RoleName returns a reference to field role_name of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) RoleName() terra.StringValue {
	return terra.ReferenceAsString(vgabr.ref.Append("role_name"))
}

// TokenBoundCidrs returns a reference to field token_bound_cidrs of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) TokenBoundCidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgabr.ref.Append("token_bound_cidrs"))
}

// TokenExplicitMaxTtl returns a reference to field token_explicit_max_ttl of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) TokenExplicitMaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(vgabr.ref.Append("token_explicit_max_ttl"))
}

// TokenMaxTtl returns a reference to field token_max_ttl of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) TokenMaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(vgabr.ref.Append("token_max_ttl"))
}

// TokenNoDefaultPolicy returns a reference to field token_no_default_policy of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) TokenNoDefaultPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(vgabr.ref.Append("token_no_default_policy"))
}

// TokenNumUses returns a reference to field token_num_uses of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) TokenNumUses() terra.NumberValue {
	return terra.ReferenceAsNumber(vgabr.ref.Append("token_num_uses"))
}

// TokenPeriod returns a reference to field token_period of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) TokenPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(vgabr.ref.Append("token_period"))
}

// TokenPolicies returns a reference to field token_policies of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) TokenPolicies() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgabr.ref.Append("token_policies"))
}

// TokenTtl returns a reference to field token_ttl of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) TokenTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(vgabr.ref.Append("token_ttl"))
}

// TokenType returns a reference to field token_type of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) TokenType() terra.StringValue {
	return terra.ReferenceAsString(vgabr.ref.Append("token_type"))
}

// Type returns a reference to field type of vault_gcp_auth_backend_role.
func (vgabr dataVaultGcpAuthBackendRoleAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vgabr.ref.Append("type"))
}
