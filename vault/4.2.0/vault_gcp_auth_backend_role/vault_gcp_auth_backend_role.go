// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_gcp_auth_backend_role

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

// Resource represents the Terraform resource vault_gcp_auth_backend_role.
type Resource struct {
	Name      string
	Args      Args
	state     *vaultGcpAuthBackendRoleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (vgabr *Resource) Type() string {
	return "vault_gcp_auth_backend_role"
}

// LocalName returns the local name for [Resource].
func (vgabr *Resource) LocalName() string {
	return vgabr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (vgabr *Resource) Configuration() interface{} {
	return vgabr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (vgabr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(vgabr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (vgabr *Resource) Dependencies() terra.Dependencies {
	return vgabr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (vgabr *Resource) LifecycleManagement() *terra.Lifecycle {
	return vgabr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (vgabr *Resource) Attributes() vaultGcpAuthBackendRoleAttributes {
	return vaultGcpAuthBackendRoleAttributes{ref: terra.ReferenceResource(vgabr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (vgabr *Resource) ImportState(state io.Reader) error {
	vgabr.state = &vaultGcpAuthBackendRoleState{}
	if err := json.NewDecoder(state).Decode(vgabr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vgabr.Type(), vgabr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (vgabr *Resource) State() (*vaultGcpAuthBackendRoleState, bool) {
	return vgabr.state, vgabr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (vgabr *Resource) StateMust() *vaultGcpAuthBackendRoleState {
	if vgabr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vgabr.Type(), vgabr.LocalName()))
	}
	return vgabr.state
}

// Args contains the configurations for vault_gcp_auth_backend_role.
type Args struct {
	// AddGroupAliases: bool, optional
	AddGroupAliases terra.BoolValue `hcl:"add_group_aliases,attr"`
	// AllowGceInference: bool, optional
	AllowGceInference terra.BoolValue `hcl:"allow_gce_inference,attr"`
	// Backend: string, optional
	Backend terra.StringValue `hcl:"backend,attr"`
	// BoundInstanceGroups: set of string, optional
	BoundInstanceGroups terra.SetValue[terra.StringValue] `hcl:"bound_instance_groups,attr"`
	// BoundLabels: set of string, optional
	BoundLabels terra.SetValue[terra.StringValue] `hcl:"bound_labels,attr"`
	// BoundProjects: set of string, optional
	BoundProjects terra.SetValue[terra.StringValue] `hcl:"bound_projects,attr"`
	// BoundRegions: set of string, optional
	BoundRegions terra.SetValue[terra.StringValue] `hcl:"bound_regions,attr"`
	// BoundServiceAccounts: set of string, optional
	BoundServiceAccounts terra.SetValue[terra.StringValue] `hcl:"bound_service_accounts,attr"`
	// BoundZones: set of string, optional
	BoundZones terra.SetValue[terra.StringValue] `hcl:"bound_zones,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxJwtExp: string, optional
	MaxJwtExp terra.StringValue `hcl:"max_jwt_exp,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
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
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type vaultGcpAuthBackendRoleAttributes struct {
	ref terra.Reference
}

// AddGroupAliases returns a reference to field add_group_aliases of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) AddGroupAliases() terra.BoolValue {
	return terra.ReferenceAsBool(vgabr.ref.Append("add_group_aliases"))
}

// AllowGceInference returns a reference to field allow_gce_inference of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) AllowGceInference() terra.BoolValue {
	return terra.ReferenceAsBool(vgabr.ref.Append("allow_gce_inference"))
}

// Backend returns a reference to field backend of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(vgabr.ref.Append("backend"))
}

// BoundInstanceGroups returns a reference to field bound_instance_groups of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) BoundInstanceGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgabr.ref.Append("bound_instance_groups"))
}

// BoundLabels returns a reference to field bound_labels of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) BoundLabels() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgabr.ref.Append("bound_labels"))
}

// BoundProjects returns a reference to field bound_projects of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) BoundProjects() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgabr.ref.Append("bound_projects"))
}

// BoundRegions returns a reference to field bound_regions of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) BoundRegions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgabr.ref.Append("bound_regions"))
}

// BoundServiceAccounts returns a reference to field bound_service_accounts of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) BoundServiceAccounts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgabr.ref.Append("bound_service_accounts"))
}

// BoundZones returns a reference to field bound_zones of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) BoundZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgabr.ref.Append("bound_zones"))
}

// Id returns a reference to field id of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vgabr.ref.Append("id"))
}

// MaxJwtExp returns a reference to field max_jwt_exp of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) MaxJwtExp() terra.StringValue {
	return terra.ReferenceAsString(vgabr.ref.Append("max_jwt_exp"))
}

// Namespace returns a reference to field namespace of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vgabr.ref.Append("namespace"))
}

// Role returns a reference to field role of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(vgabr.ref.Append("role"))
}

// TokenBoundCidrs returns a reference to field token_bound_cidrs of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) TokenBoundCidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgabr.ref.Append("token_bound_cidrs"))
}

// TokenExplicitMaxTtl returns a reference to field token_explicit_max_ttl of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) TokenExplicitMaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(vgabr.ref.Append("token_explicit_max_ttl"))
}

// TokenMaxTtl returns a reference to field token_max_ttl of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) TokenMaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(vgabr.ref.Append("token_max_ttl"))
}

// TokenNoDefaultPolicy returns a reference to field token_no_default_policy of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) TokenNoDefaultPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(vgabr.ref.Append("token_no_default_policy"))
}

// TokenNumUses returns a reference to field token_num_uses of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) TokenNumUses() terra.NumberValue {
	return terra.ReferenceAsNumber(vgabr.ref.Append("token_num_uses"))
}

// TokenPeriod returns a reference to field token_period of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) TokenPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(vgabr.ref.Append("token_period"))
}

// TokenPolicies returns a reference to field token_policies of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) TokenPolicies() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vgabr.ref.Append("token_policies"))
}

// TokenTtl returns a reference to field token_ttl of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) TokenTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(vgabr.ref.Append("token_ttl"))
}

// TokenType returns a reference to field token_type of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) TokenType() terra.StringValue {
	return terra.ReferenceAsString(vgabr.ref.Append("token_type"))
}

// Type returns a reference to field type of vault_gcp_auth_backend_role.
func (vgabr vaultGcpAuthBackendRoleAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vgabr.ref.Append("type"))
}

type vaultGcpAuthBackendRoleState struct {
	AddGroupAliases      bool     `json:"add_group_aliases"`
	AllowGceInference    bool     `json:"allow_gce_inference"`
	Backend              string   `json:"backend"`
	BoundInstanceGroups  []string `json:"bound_instance_groups"`
	BoundLabels          []string `json:"bound_labels"`
	BoundProjects        []string `json:"bound_projects"`
	BoundRegions         []string `json:"bound_regions"`
	BoundServiceAccounts []string `json:"bound_service_accounts"`
	BoundZones           []string `json:"bound_zones"`
	Id                   string   `json:"id"`
	MaxJwtExp            string   `json:"max_jwt_exp"`
	Namespace            string   `json:"namespace"`
	Role                 string   `json:"role"`
	TokenBoundCidrs      []string `json:"token_bound_cidrs"`
	TokenExplicitMaxTtl  float64  `json:"token_explicit_max_ttl"`
	TokenMaxTtl          float64  `json:"token_max_ttl"`
	TokenNoDefaultPolicy bool     `json:"token_no_default_policy"`
	TokenNumUses         float64  `json:"token_num_uses"`
	TokenPeriod          float64  `json:"token_period"`
	TokenPolicies        []string `json:"token_policies"`
	TokenTtl             float64  `json:"token_ttl"`
	TokenType            string   `json:"token_type"`
	Type                 string   `json:"type"`
}
