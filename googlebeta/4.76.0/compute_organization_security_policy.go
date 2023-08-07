// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeorganizationsecuritypolicy "github.com/golingon/terraproviders/googlebeta/4.76.0/computeorganizationsecuritypolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeOrganizationSecurityPolicy creates a new instance of [ComputeOrganizationSecurityPolicy].
func NewComputeOrganizationSecurityPolicy(name string, args ComputeOrganizationSecurityPolicyArgs) *ComputeOrganizationSecurityPolicy {
	return &ComputeOrganizationSecurityPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeOrganizationSecurityPolicy)(nil)

// ComputeOrganizationSecurityPolicy represents the Terraform resource google_compute_organization_security_policy.
type ComputeOrganizationSecurityPolicy struct {
	Name      string
	Args      ComputeOrganizationSecurityPolicyArgs
	state     *computeOrganizationSecurityPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeOrganizationSecurityPolicy].
func (cosp *ComputeOrganizationSecurityPolicy) Type() string {
	return "google_compute_organization_security_policy"
}

// LocalName returns the local name for [ComputeOrganizationSecurityPolicy].
func (cosp *ComputeOrganizationSecurityPolicy) LocalName() string {
	return cosp.Name
}

// Configuration returns the configuration (args) for [ComputeOrganizationSecurityPolicy].
func (cosp *ComputeOrganizationSecurityPolicy) Configuration() interface{} {
	return cosp.Args
}

// DependOn is used for other resources to depend on [ComputeOrganizationSecurityPolicy].
func (cosp *ComputeOrganizationSecurityPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(cosp)
}

// Dependencies returns the list of resources [ComputeOrganizationSecurityPolicy] depends_on.
func (cosp *ComputeOrganizationSecurityPolicy) Dependencies() terra.Dependencies {
	return cosp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeOrganizationSecurityPolicy].
func (cosp *ComputeOrganizationSecurityPolicy) LifecycleManagement() *terra.Lifecycle {
	return cosp.Lifecycle
}

// Attributes returns the attributes for [ComputeOrganizationSecurityPolicy].
func (cosp *ComputeOrganizationSecurityPolicy) Attributes() computeOrganizationSecurityPolicyAttributes {
	return computeOrganizationSecurityPolicyAttributes{ref: terra.ReferenceResource(cosp)}
}

// ImportState imports the given attribute values into [ComputeOrganizationSecurityPolicy]'s state.
func (cosp *ComputeOrganizationSecurityPolicy) ImportState(av io.Reader) error {
	cosp.state = &computeOrganizationSecurityPolicyState{}
	if err := json.NewDecoder(av).Decode(cosp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cosp.Type(), cosp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeOrganizationSecurityPolicy] has state.
func (cosp *ComputeOrganizationSecurityPolicy) State() (*computeOrganizationSecurityPolicyState, bool) {
	return cosp.state, cosp.state != nil
}

// StateMust returns the state for [ComputeOrganizationSecurityPolicy]. Panics if the state is nil.
func (cosp *ComputeOrganizationSecurityPolicy) StateMust() *computeOrganizationSecurityPolicyState {
	if cosp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cosp.Type(), cosp.LocalName()))
	}
	return cosp.state
}

// ComputeOrganizationSecurityPolicyArgs contains the configurations for google_compute_organization_security_policy.
type ComputeOrganizationSecurityPolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Timeouts: optional
	Timeouts *computeorganizationsecuritypolicy.Timeouts `hcl:"timeouts,block"`
}
type computeOrganizationSecurityPolicyAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_organization_security_policy.
func (cosp computeOrganizationSecurityPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cosp.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_compute_organization_security_policy.
func (cosp computeOrganizationSecurityPolicyAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(cosp.ref.Append("display_name"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_organization_security_policy.
func (cosp computeOrganizationSecurityPolicyAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(cosp.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_compute_organization_security_policy.
func (cosp computeOrganizationSecurityPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cosp.ref.Append("id"))
}

// Parent returns a reference to field parent of google_compute_organization_security_policy.
func (cosp computeOrganizationSecurityPolicyAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(cosp.ref.Append("parent"))
}

// PolicyId returns a reference to field policy_id of google_compute_organization_security_policy.
func (cosp computeOrganizationSecurityPolicyAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceAsString(cosp.ref.Append("policy_id"))
}

// Type returns a reference to field type of google_compute_organization_security_policy.
func (cosp computeOrganizationSecurityPolicyAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(cosp.ref.Append("type"))
}

func (cosp computeOrganizationSecurityPolicyAttributes) Timeouts() computeorganizationsecuritypolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeorganizationsecuritypolicy.TimeoutsAttributes](cosp.ref.Append("timeouts"))
}

type computeOrganizationSecurityPolicyState struct {
	Description string                                           `json:"description"`
	DisplayName string                                           `json:"display_name"`
	Fingerprint string                                           `json:"fingerprint"`
	Id          string                                           `json:"id"`
	Parent      string                                           `json:"parent"`
	PolicyId    string                                           `json:"policy_id"`
	Type        string                                           `json:"type"`
	Timeouts    *computeorganizationsecuritypolicy.TimeoutsState `json:"timeouts"`
}
