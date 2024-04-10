// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewCertAuthBackendRole creates a new instance of [CertAuthBackendRole].
func NewCertAuthBackendRole(name string, args CertAuthBackendRoleArgs) *CertAuthBackendRole {
	return &CertAuthBackendRole{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CertAuthBackendRole)(nil)

// CertAuthBackendRole represents the Terraform resource vault_cert_auth_backend_role.
type CertAuthBackendRole struct {
	Name      string
	Args      CertAuthBackendRoleArgs
	state     *certAuthBackendRoleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CertAuthBackendRole].
func (cabr *CertAuthBackendRole) Type() string {
	return "vault_cert_auth_backend_role"
}

// LocalName returns the local name for [CertAuthBackendRole].
func (cabr *CertAuthBackendRole) LocalName() string {
	return cabr.Name
}

// Configuration returns the configuration (args) for [CertAuthBackendRole].
func (cabr *CertAuthBackendRole) Configuration() interface{} {
	return cabr.Args
}

// DependOn is used for other resources to depend on [CertAuthBackendRole].
func (cabr *CertAuthBackendRole) DependOn() terra.Reference {
	return terra.ReferenceResource(cabr)
}

// Dependencies returns the list of resources [CertAuthBackendRole] depends_on.
func (cabr *CertAuthBackendRole) Dependencies() terra.Dependencies {
	return cabr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CertAuthBackendRole].
func (cabr *CertAuthBackendRole) LifecycleManagement() *terra.Lifecycle {
	return cabr.Lifecycle
}

// Attributes returns the attributes for [CertAuthBackendRole].
func (cabr *CertAuthBackendRole) Attributes() certAuthBackendRoleAttributes {
	return certAuthBackendRoleAttributes{ref: terra.ReferenceResource(cabr)}
}

// ImportState imports the given attribute values into [CertAuthBackendRole]'s state.
func (cabr *CertAuthBackendRole) ImportState(av io.Reader) error {
	cabr.state = &certAuthBackendRoleState{}
	if err := json.NewDecoder(av).Decode(cabr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cabr.Type(), cabr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CertAuthBackendRole] has state.
func (cabr *CertAuthBackendRole) State() (*certAuthBackendRoleState, bool) {
	return cabr.state, cabr.state != nil
}

// StateMust returns the state for [CertAuthBackendRole]. Panics if the state is nil.
func (cabr *CertAuthBackendRole) StateMust() *certAuthBackendRoleState {
	if cabr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cabr.Type(), cabr.LocalName()))
	}
	return cabr.state
}

// CertAuthBackendRoleArgs contains the configurations for vault_cert_auth_backend_role.
type CertAuthBackendRoleArgs struct {
	// AllowedCommonNames: set of string, optional
	AllowedCommonNames terra.SetValue[terra.StringValue] `hcl:"allowed_common_names,attr"`
	// AllowedDnsSans: set of string, optional
	AllowedDnsSans terra.SetValue[terra.StringValue] `hcl:"allowed_dns_sans,attr"`
	// AllowedEmailSans: set of string, optional
	AllowedEmailSans terra.SetValue[terra.StringValue] `hcl:"allowed_email_sans,attr"`
	// AllowedNames: set of string, optional
	AllowedNames terra.SetValue[terra.StringValue] `hcl:"allowed_names,attr"`
	// AllowedOrganizationalUnits: set of string, optional
	AllowedOrganizationalUnits terra.SetValue[terra.StringValue] `hcl:"allowed_organizational_units,attr"`
	// AllowedUriSans: set of string, optional
	AllowedUriSans terra.SetValue[terra.StringValue] `hcl:"allowed_uri_sans,attr"`
	// Backend: string, optional
	Backend terra.StringValue `hcl:"backend,attr"`
	// Certificate: string, required
	Certificate terra.StringValue `hcl:"certificate,attr" validate:"required"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// OcspCaCertificates: string, optional
	OcspCaCertificates terra.StringValue `hcl:"ocsp_ca_certificates,attr"`
	// OcspEnabled: bool, optional
	OcspEnabled terra.BoolValue `hcl:"ocsp_enabled,attr"`
	// OcspFailOpen: bool, optional
	OcspFailOpen terra.BoolValue `hcl:"ocsp_fail_open,attr"`
	// OcspQueryAllServers: bool, optional
	OcspQueryAllServers terra.BoolValue `hcl:"ocsp_query_all_servers,attr"`
	// OcspServersOverride: set of string, optional
	OcspServersOverride terra.SetValue[terra.StringValue] `hcl:"ocsp_servers_override,attr"`
	// RequiredExtensions: set of string, optional
	RequiredExtensions terra.SetValue[terra.StringValue] `hcl:"required_extensions,attr"`
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
type certAuthBackendRoleAttributes struct {
	ref terra.Reference
}

// AllowedCommonNames returns a reference to field allowed_common_names of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) AllowedCommonNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cabr.ref.Append("allowed_common_names"))
}

// AllowedDnsSans returns a reference to field allowed_dns_sans of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) AllowedDnsSans() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cabr.ref.Append("allowed_dns_sans"))
}

// AllowedEmailSans returns a reference to field allowed_email_sans of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) AllowedEmailSans() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cabr.ref.Append("allowed_email_sans"))
}

// AllowedNames returns a reference to field allowed_names of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) AllowedNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cabr.ref.Append("allowed_names"))
}

// AllowedOrganizationalUnits returns a reference to field allowed_organizational_units of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) AllowedOrganizationalUnits() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cabr.ref.Append("allowed_organizational_units"))
}

// AllowedUriSans returns a reference to field allowed_uri_sans of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) AllowedUriSans() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cabr.ref.Append("allowed_uri_sans"))
}

// Backend returns a reference to field backend of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(cabr.ref.Append("backend"))
}

// Certificate returns a reference to field certificate of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(cabr.ref.Append("certificate"))
}

// DisplayName returns a reference to field display_name of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(cabr.ref.Append("display_name"))
}

// Id returns a reference to field id of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cabr.ref.Append("id"))
}

// Name returns a reference to field name of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cabr.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(cabr.ref.Append("namespace"))
}

// OcspCaCertificates returns a reference to field ocsp_ca_certificates of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) OcspCaCertificates() terra.StringValue {
	return terra.ReferenceAsString(cabr.ref.Append("ocsp_ca_certificates"))
}

// OcspEnabled returns a reference to field ocsp_enabled of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) OcspEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cabr.ref.Append("ocsp_enabled"))
}

// OcspFailOpen returns a reference to field ocsp_fail_open of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) OcspFailOpen() terra.BoolValue {
	return terra.ReferenceAsBool(cabr.ref.Append("ocsp_fail_open"))
}

// OcspQueryAllServers returns a reference to field ocsp_query_all_servers of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) OcspQueryAllServers() terra.BoolValue {
	return terra.ReferenceAsBool(cabr.ref.Append("ocsp_query_all_servers"))
}

// OcspServersOverride returns a reference to field ocsp_servers_override of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) OcspServersOverride() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cabr.ref.Append("ocsp_servers_override"))
}

// RequiredExtensions returns a reference to field required_extensions of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) RequiredExtensions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cabr.ref.Append("required_extensions"))
}

// TokenBoundCidrs returns a reference to field token_bound_cidrs of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) TokenBoundCidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cabr.ref.Append("token_bound_cidrs"))
}

// TokenExplicitMaxTtl returns a reference to field token_explicit_max_ttl of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) TokenExplicitMaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(cabr.ref.Append("token_explicit_max_ttl"))
}

// TokenMaxTtl returns a reference to field token_max_ttl of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) TokenMaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(cabr.ref.Append("token_max_ttl"))
}

// TokenNoDefaultPolicy returns a reference to field token_no_default_policy of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) TokenNoDefaultPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(cabr.ref.Append("token_no_default_policy"))
}

// TokenNumUses returns a reference to field token_num_uses of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) TokenNumUses() terra.NumberValue {
	return terra.ReferenceAsNumber(cabr.ref.Append("token_num_uses"))
}

// TokenPeriod returns a reference to field token_period of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) TokenPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(cabr.ref.Append("token_period"))
}

// TokenPolicies returns a reference to field token_policies of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) TokenPolicies() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cabr.ref.Append("token_policies"))
}

// TokenTtl returns a reference to field token_ttl of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) TokenTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(cabr.ref.Append("token_ttl"))
}

// TokenType returns a reference to field token_type of vault_cert_auth_backend_role.
func (cabr certAuthBackendRoleAttributes) TokenType() terra.StringValue {
	return terra.ReferenceAsString(cabr.ref.Append("token_type"))
}

type certAuthBackendRoleState struct {
	AllowedCommonNames         []string `json:"allowed_common_names"`
	AllowedDnsSans             []string `json:"allowed_dns_sans"`
	AllowedEmailSans           []string `json:"allowed_email_sans"`
	AllowedNames               []string `json:"allowed_names"`
	AllowedOrganizationalUnits []string `json:"allowed_organizational_units"`
	AllowedUriSans             []string `json:"allowed_uri_sans"`
	Backend                    string   `json:"backend"`
	Certificate                string   `json:"certificate"`
	DisplayName                string   `json:"display_name"`
	Id                         string   `json:"id"`
	Name                       string   `json:"name"`
	Namespace                  string   `json:"namespace"`
	OcspCaCertificates         string   `json:"ocsp_ca_certificates"`
	OcspEnabled                bool     `json:"ocsp_enabled"`
	OcspFailOpen               bool     `json:"ocsp_fail_open"`
	OcspQueryAllServers        bool     `json:"ocsp_query_all_servers"`
	OcspServersOverride        []string `json:"ocsp_servers_override"`
	RequiredExtensions         []string `json:"required_extensions"`
	TokenBoundCidrs            []string `json:"token_bound_cidrs"`
	TokenExplicitMaxTtl        float64  `json:"token_explicit_max_ttl"`
	TokenMaxTtl                float64  `json:"token_max_ttl"`
	TokenNoDefaultPolicy       bool     `json:"token_no_default_policy"`
	TokenNumUses               float64  `json:"token_num_uses"`
	TokenPeriod                float64  `json:"token_period"`
	TokenPolicies              []string `json:"token_policies"`
	TokenTtl                   float64  `json:"token_ttl"`
	TokenType                  string   `json:"token_type"`
}
