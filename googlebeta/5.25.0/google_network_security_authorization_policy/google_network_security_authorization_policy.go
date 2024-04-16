// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_network_security_authorization_policy

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

// Resource represents the Terraform resource google_network_security_authorization_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleNetworkSecurityAuthorizationPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gnsap *Resource) Type() string {
	return "google_network_security_authorization_policy"
}

// LocalName returns the local name for [Resource].
func (gnsap *Resource) LocalName() string {
	return gnsap.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gnsap *Resource) Configuration() interface{} {
	return gnsap.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gnsap *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gnsap)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gnsap *Resource) Dependencies() terra.Dependencies {
	return gnsap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gnsap *Resource) LifecycleManagement() *terra.Lifecycle {
	return gnsap.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gnsap *Resource) Attributes() googleNetworkSecurityAuthorizationPolicyAttributes {
	return googleNetworkSecurityAuthorizationPolicyAttributes{ref: terra.ReferenceResource(gnsap)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gnsap *Resource) ImportState(state io.Reader) error {
	gnsap.state = &googleNetworkSecurityAuthorizationPolicyState{}
	if err := json.NewDecoder(state).Decode(gnsap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gnsap.Type(), gnsap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gnsap *Resource) State() (*googleNetworkSecurityAuthorizationPolicyState, bool) {
	return gnsap.state, gnsap.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gnsap *Resource) StateMust() *googleNetworkSecurityAuthorizationPolicyState {
	if gnsap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gnsap.Type(), gnsap.LocalName()))
	}
	return gnsap.state
}

// Args contains the configurations for google_network_security_authorization_policy.
type Args struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Rules: min=0
	Rules []Rules `hcl:"rules,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleNetworkSecurityAuthorizationPolicyAttributes struct {
	ref terra.Reference
}

// Action returns a reference to field action of google_network_security_authorization_policy.
func (gnsap googleNetworkSecurityAuthorizationPolicyAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(gnsap.ref.Append("action"))
}

// CreateTime returns a reference to field create_time of google_network_security_authorization_policy.
func (gnsap googleNetworkSecurityAuthorizationPolicyAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gnsap.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_security_authorization_policy.
func (gnsap googleNetworkSecurityAuthorizationPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gnsap.ref.Append("description"))
}

// EffectiveLabels returns a reference to field effective_labels of google_network_security_authorization_policy.
func (gnsap googleNetworkSecurityAuthorizationPolicyAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gnsap.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_network_security_authorization_policy.
func (gnsap googleNetworkSecurityAuthorizationPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gnsap.ref.Append("id"))
}

// Labels returns a reference to field labels of google_network_security_authorization_policy.
func (gnsap googleNetworkSecurityAuthorizationPolicyAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gnsap.ref.Append("labels"))
}

// Location returns a reference to field location of google_network_security_authorization_policy.
func (gnsap googleNetworkSecurityAuthorizationPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gnsap.ref.Append("location"))
}

// Name returns a reference to field name of google_network_security_authorization_policy.
func (gnsap googleNetworkSecurityAuthorizationPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gnsap.ref.Append("name"))
}

// Project returns a reference to field project of google_network_security_authorization_policy.
func (gnsap googleNetworkSecurityAuthorizationPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gnsap.ref.Append("project"))
}

// TerraformLabels returns a reference to field terraform_labels of google_network_security_authorization_policy.
func (gnsap googleNetworkSecurityAuthorizationPolicyAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gnsap.ref.Append("terraform_labels"))
}

// UpdateTime returns a reference to field update_time of google_network_security_authorization_policy.
func (gnsap googleNetworkSecurityAuthorizationPolicyAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gnsap.ref.Append("update_time"))
}

func (gnsap googleNetworkSecurityAuthorizationPolicyAttributes) Rules() terra.ListValue[RulesAttributes] {
	return terra.ReferenceAsList[RulesAttributes](gnsap.ref.Append("rules"))
}

func (gnsap googleNetworkSecurityAuthorizationPolicyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gnsap.ref.Append("timeouts"))
}

type googleNetworkSecurityAuthorizationPolicyState struct {
	Action          string            `json:"action"`
	CreateTime      string            `json:"create_time"`
	Description     string            `json:"description"`
	EffectiveLabels map[string]string `json:"effective_labels"`
	Id              string            `json:"id"`
	Labels          map[string]string `json:"labels"`
	Location        string            `json:"location"`
	Name            string            `json:"name"`
	Project         string            `json:"project"`
	TerraformLabels map[string]string `json:"terraform_labels"`
	UpdateTime      string            `json:"update_time"`
	Rules           []RulesState      `json:"rules"`
	Timeouts        *TimeoutsState    `json:"timeouts"`
}
