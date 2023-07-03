// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConsulSecretBackendRole creates a new instance of [ConsulSecretBackendRole].
func NewConsulSecretBackendRole(name string, args ConsulSecretBackendRoleArgs) *ConsulSecretBackendRole {
	return &ConsulSecretBackendRole{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConsulSecretBackendRole)(nil)

// ConsulSecretBackendRole represents the Terraform resource vault_consul_secret_backend_role.
type ConsulSecretBackendRole struct {
	Name      string
	Args      ConsulSecretBackendRoleArgs
	state     *consulSecretBackendRoleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConsulSecretBackendRole].
func (csbr *ConsulSecretBackendRole) Type() string {
	return "vault_consul_secret_backend_role"
}

// LocalName returns the local name for [ConsulSecretBackendRole].
func (csbr *ConsulSecretBackendRole) LocalName() string {
	return csbr.Name
}

// Configuration returns the configuration (args) for [ConsulSecretBackendRole].
func (csbr *ConsulSecretBackendRole) Configuration() interface{} {
	return csbr.Args
}

// DependOn is used for other resources to depend on [ConsulSecretBackendRole].
func (csbr *ConsulSecretBackendRole) DependOn() terra.Reference {
	return terra.ReferenceResource(csbr)
}

// Dependencies returns the list of resources [ConsulSecretBackendRole] depends_on.
func (csbr *ConsulSecretBackendRole) Dependencies() terra.Dependencies {
	return csbr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConsulSecretBackendRole].
func (csbr *ConsulSecretBackendRole) LifecycleManagement() *terra.Lifecycle {
	return csbr.Lifecycle
}

// Attributes returns the attributes for [ConsulSecretBackendRole].
func (csbr *ConsulSecretBackendRole) Attributes() consulSecretBackendRoleAttributes {
	return consulSecretBackendRoleAttributes{ref: terra.ReferenceResource(csbr)}
}

// ImportState imports the given attribute values into [ConsulSecretBackendRole]'s state.
func (csbr *ConsulSecretBackendRole) ImportState(av io.Reader) error {
	csbr.state = &consulSecretBackendRoleState{}
	if err := json.NewDecoder(av).Decode(csbr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csbr.Type(), csbr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConsulSecretBackendRole] has state.
func (csbr *ConsulSecretBackendRole) State() (*consulSecretBackendRoleState, bool) {
	return csbr.state, csbr.state != nil
}

// StateMust returns the state for [ConsulSecretBackendRole]. Panics if the state is nil.
func (csbr *ConsulSecretBackendRole) StateMust() *consulSecretBackendRoleState {
	if csbr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csbr.Type(), csbr.LocalName()))
	}
	return csbr.state
}

// ConsulSecretBackendRoleArgs contains the configurations for vault_consul_secret_backend_role.
type ConsulSecretBackendRoleArgs struct {
	// Backend: string, optional
	Backend terra.StringValue `hcl:"backend,attr"`
	// ConsulNamespace: string, optional
	ConsulNamespace terra.StringValue `hcl:"consul_namespace,attr"`
	// ConsulPolicies: set of string, optional
	ConsulPolicies terra.SetValue[terra.StringValue] `hcl:"consul_policies,attr"`
	// ConsulRoles: set of string, optional
	ConsulRoles terra.SetValue[terra.StringValue] `hcl:"consul_roles,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Local: bool, optional
	Local terra.BoolValue `hcl:"local,attr"`
	// MaxTtl: number, optional
	MaxTtl terra.NumberValue `hcl:"max_ttl,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// NodeIdentities: set of string, optional
	NodeIdentities terra.SetValue[terra.StringValue] `hcl:"node_identities,attr"`
	// Partition: string, optional
	Partition terra.StringValue `hcl:"partition,attr"`
	// Policies: list of string, optional
	Policies terra.ListValue[terra.StringValue] `hcl:"policies,attr"`
	// ServiceIdentities: set of string, optional
	ServiceIdentities terra.SetValue[terra.StringValue] `hcl:"service_identities,attr"`
	// TokenType: string, optional
	TokenType terra.StringValue `hcl:"token_type,attr"`
	// Ttl: number, optional
	Ttl terra.NumberValue `hcl:"ttl,attr"`
}
type consulSecretBackendRoleAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_consul_secret_backend_role.
func (csbr consulSecretBackendRoleAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(csbr.ref.Append("backend"))
}

// ConsulNamespace returns a reference to field consul_namespace of vault_consul_secret_backend_role.
func (csbr consulSecretBackendRoleAttributes) ConsulNamespace() terra.StringValue {
	return terra.ReferenceAsString(csbr.ref.Append("consul_namespace"))
}

// ConsulPolicies returns a reference to field consul_policies of vault_consul_secret_backend_role.
func (csbr consulSecretBackendRoleAttributes) ConsulPolicies() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](csbr.ref.Append("consul_policies"))
}

// ConsulRoles returns a reference to field consul_roles of vault_consul_secret_backend_role.
func (csbr consulSecretBackendRoleAttributes) ConsulRoles() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](csbr.ref.Append("consul_roles"))
}

// Id returns a reference to field id of vault_consul_secret_backend_role.
func (csbr consulSecretBackendRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csbr.ref.Append("id"))
}

// Local returns a reference to field local of vault_consul_secret_backend_role.
func (csbr consulSecretBackendRoleAttributes) Local() terra.BoolValue {
	return terra.ReferenceAsBool(csbr.ref.Append("local"))
}

// MaxTtl returns a reference to field max_ttl of vault_consul_secret_backend_role.
func (csbr consulSecretBackendRoleAttributes) MaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(csbr.ref.Append("max_ttl"))
}

// Name returns a reference to field name of vault_consul_secret_backend_role.
func (csbr consulSecretBackendRoleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(csbr.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_consul_secret_backend_role.
func (csbr consulSecretBackendRoleAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(csbr.ref.Append("namespace"))
}

// NodeIdentities returns a reference to field node_identities of vault_consul_secret_backend_role.
func (csbr consulSecretBackendRoleAttributes) NodeIdentities() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](csbr.ref.Append("node_identities"))
}

// Partition returns a reference to field partition of vault_consul_secret_backend_role.
func (csbr consulSecretBackendRoleAttributes) Partition() terra.StringValue {
	return terra.ReferenceAsString(csbr.ref.Append("partition"))
}

// Policies returns a reference to field policies of vault_consul_secret_backend_role.
func (csbr consulSecretBackendRoleAttributes) Policies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](csbr.ref.Append("policies"))
}

// ServiceIdentities returns a reference to field service_identities of vault_consul_secret_backend_role.
func (csbr consulSecretBackendRoleAttributes) ServiceIdentities() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](csbr.ref.Append("service_identities"))
}

// TokenType returns a reference to field token_type of vault_consul_secret_backend_role.
func (csbr consulSecretBackendRoleAttributes) TokenType() terra.StringValue {
	return terra.ReferenceAsString(csbr.ref.Append("token_type"))
}

// Ttl returns a reference to field ttl of vault_consul_secret_backend_role.
func (csbr consulSecretBackendRoleAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(csbr.ref.Append("ttl"))
}

type consulSecretBackendRoleState struct {
	Backend           string   `json:"backend"`
	ConsulNamespace   string   `json:"consul_namespace"`
	ConsulPolicies    []string `json:"consul_policies"`
	ConsulRoles       []string `json:"consul_roles"`
	Id                string   `json:"id"`
	Local             bool     `json:"local"`
	MaxTtl            float64  `json:"max_ttl"`
	Name              string   `json:"name"`
	Namespace         string   `json:"namespace"`
	NodeIdentities    []string `json:"node_identities"`
	Partition         string   `json:"partition"`
	Policies          []string `json:"policies"`
	ServiceIdentities []string `json:"service_identities"`
	TokenType         string   `json:"token_type"`
	Ttl               float64  `json:"ttl"`
}
