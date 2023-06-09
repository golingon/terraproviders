// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataGcpAuthBackendRole creates a new instance of [DataGcpAuthBackendRole].
func NewDataGcpAuthBackendRole(name string, args DataGcpAuthBackendRoleArgs) *DataGcpAuthBackendRole {
	return &DataGcpAuthBackendRole{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataGcpAuthBackendRole)(nil)

// DataGcpAuthBackendRole represents the Terraform data resource vault_gcp_auth_backend_role.
type DataGcpAuthBackendRole struct {
	Name string
	Args DataGcpAuthBackendRoleArgs
}

// DataSource returns the Terraform object type for [DataGcpAuthBackendRole].
func (gabr *DataGcpAuthBackendRole) DataSource() string {
	return "vault_gcp_auth_backend_role"
}

// LocalName returns the local name for [DataGcpAuthBackendRole].
func (gabr *DataGcpAuthBackendRole) LocalName() string {
	return gabr.Name
}

// Configuration returns the configuration (args) for [DataGcpAuthBackendRole].
func (gabr *DataGcpAuthBackendRole) Configuration() interface{} {
	return gabr.Args
}

// Attributes returns the attributes for [DataGcpAuthBackendRole].
func (gabr *DataGcpAuthBackendRole) Attributes() dataGcpAuthBackendRoleAttributes {
	return dataGcpAuthBackendRoleAttributes{ref: terra.ReferenceDataResource(gabr)}
}

// DataGcpAuthBackendRoleArgs contains the configurations for vault_gcp_auth_backend_role.
type DataGcpAuthBackendRoleArgs struct {
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
type dataGcpAuthBackendRoleAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(gabr.ref.Append("backend"))
}

// BoundInstanceGroups returns a reference to field bound_instance_groups of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) BoundInstanceGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gabr.ref.Append("bound_instance_groups"))
}

// BoundLabels returns a reference to field bound_labels of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) BoundLabels() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gabr.ref.Append("bound_labels"))
}

// BoundProjects returns a reference to field bound_projects of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) BoundProjects() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gabr.ref.Append("bound_projects"))
}

// BoundRegions returns a reference to field bound_regions of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) BoundRegions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gabr.ref.Append("bound_regions"))
}

// BoundServiceAccounts returns a reference to field bound_service_accounts of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) BoundServiceAccounts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gabr.ref.Append("bound_service_accounts"))
}

// BoundZones returns a reference to field bound_zones of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) BoundZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gabr.ref.Append("bound_zones"))
}

// Id returns a reference to field id of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gabr.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(gabr.ref.Append("namespace"))
}

// RoleId returns a reference to field role_id of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) RoleId() terra.StringValue {
	return terra.ReferenceAsString(gabr.ref.Append("role_id"))
}

// RoleName returns a reference to field role_name of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) RoleName() terra.StringValue {
	return terra.ReferenceAsString(gabr.ref.Append("role_name"))
}

// TokenBoundCidrs returns a reference to field token_bound_cidrs of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) TokenBoundCidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gabr.ref.Append("token_bound_cidrs"))
}

// TokenExplicitMaxTtl returns a reference to field token_explicit_max_ttl of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) TokenExplicitMaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(gabr.ref.Append("token_explicit_max_ttl"))
}

// TokenMaxTtl returns a reference to field token_max_ttl of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) TokenMaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(gabr.ref.Append("token_max_ttl"))
}

// TokenNoDefaultPolicy returns a reference to field token_no_default_policy of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) TokenNoDefaultPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(gabr.ref.Append("token_no_default_policy"))
}

// TokenNumUses returns a reference to field token_num_uses of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) TokenNumUses() terra.NumberValue {
	return terra.ReferenceAsNumber(gabr.ref.Append("token_num_uses"))
}

// TokenPeriod returns a reference to field token_period of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) TokenPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(gabr.ref.Append("token_period"))
}

// TokenPolicies returns a reference to field token_policies of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) TokenPolicies() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gabr.ref.Append("token_policies"))
}

// TokenTtl returns a reference to field token_ttl of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) TokenTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(gabr.ref.Append("token_ttl"))
}

// TokenType returns a reference to field token_type of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) TokenType() terra.StringValue {
	return terra.ReferenceAsString(gabr.ref.Append("token_type"))
}

// Type returns a reference to field type of vault_gcp_auth_backend_role.
func (gabr dataGcpAuthBackendRoleAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(gabr.ref.Append("type"))
}
