// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computesecuritypolicy "github.com/golingon/terraproviders/google/4.59.0/computesecuritypolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeSecurityPolicy creates a new instance of [ComputeSecurityPolicy].
func NewComputeSecurityPolicy(name string, args ComputeSecurityPolicyArgs) *ComputeSecurityPolicy {
	return &ComputeSecurityPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeSecurityPolicy)(nil)

// ComputeSecurityPolicy represents the Terraform resource google_compute_security_policy.
type ComputeSecurityPolicy struct {
	Name      string
	Args      ComputeSecurityPolicyArgs
	state     *computeSecurityPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeSecurityPolicy].
func (csp *ComputeSecurityPolicy) Type() string {
	return "google_compute_security_policy"
}

// LocalName returns the local name for [ComputeSecurityPolicy].
func (csp *ComputeSecurityPolicy) LocalName() string {
	return csp.Name
}

// Configuration returns the configuration (args) for [ComputeSecurityPolicy].
func (csp *ComputeSecurityPolicy) Configuration() interface{} {
	return csp.Args
}

// DependOn is used for other resources to depend on [ComputeSecurityPolicy].
func (csp *ComputeSecurityPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(csp)
}

// Dependencies returns the list of resources [ComputeSecurityPolicy] depends_on.
func (csp *ComputeSecurityPolicy) Dependencies() terra.Dependencies {
	return csp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeSecurityPolicy].
func (csp *ComputeSecurityPolicy) LifecycleManagement() *terra.Lifecycle {
	return csp.Lifecycle
}

// Attributes returns the attributes for [ComputeSecurityPolicy].
func (csp *ComputeSecurityPolicy) Attributes() computeSecurityPolicyAttributes {
	return computeSecurityPolicyAttributes{ref: terra.ReferenceResource(csp)}
}

// ImportState imports the given attribute values into [ComputeSecurityPolicy]'s state.
func (csp *ComputeSecurityPolicy) ImportState(av io.Reader) error {
	csp.state = &computeSecurityPolicyState{}
	if err := json.NewDecoder(av).Decode(csp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csp.Type(), csp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeSecurityPolicy] has state.
func (csp *ComputeSecurityPolicy) State() (*computeSecurityPolicyState, bool) {
	return csp.state, csp.state != nil
}

// StateMust returns the state for [ComputeSecurityPolicy]. Panics if the state is nil.
func (csp *ComputeSecurityPolicy) StateMust() *computeSecurityPolicyState {
	if csp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csp.Type(), csp.LocalName()))
	}
	return csp.state
}

// ComputeSecurityPolicyArgs contains the configurations for google_compute_security_policy.
type ComputeSecurityPolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// AdaptiveProtectionConfig: optional
	AdaptiveProtectionConfig *computesecuritypolicy.AdaptiveProtectionConfig `hcl:"adaptive_protection_config,block"`
	// AdvancedOptionsConfig: optional
	AdvancedOptionsConfig *computesecuritypolicy.AdvancedOptionsConfig `hcl:"advanced_options_config,block"`
	// RecaptchaOptionsConfig: optional
	RecaptchaOptionsConfig *computesecuritypolicy.RecaptchaOptionsConfig `hcl:"recaptcha_options_config,block"`
	// Rule: min=0
	Rule []computesecuritypolicy.Rule `hcl:"rule,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *computesecuritypolicy.Timeouts `hcl:"timeouts,block"`
}
type computeSecurityPolicyAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_security_policy.
func (csp computeSecurityPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("description"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_security_policy.
func (csp computeSecurityPolicyAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_compute_security_policy.
func (csp computeSecurityPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_security_policy.
func (csp computeSecurityPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_security_policy.
func (csp computeSecurityPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_security_policy.
func (csp computeSecurityPolicyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("self_link"))
}

// Type returns a reference to field type of google_compute_security_policy.
func (csp computeSecurityPolicyAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("type"))
}

func (csp computeSecurityPolicyAttributes) AdaptiveProtectionConfig() terra.ListValue[computesecuritypolicy.AdaptiveProtectionConfigAttributes] {
	return terra.ReferenceAsList[computesecuritypolicy.AdaptiveProtectionConfigAttributes](csp.ref.Append("adaptive_protection_config"))
}

func (csp computeSecurityPolicyAttributes) AdvancedOptionsConfig() terra.ListValue[computesecuritypolicy.AdvancedOptionsConfigAttributes] {
	return terra.ReferenceAsList[computesecuritypolicy.AdvancedOptionsConfigAttributes](csp.ref.Append("advanced_options_config"))
}

func (csp computeSecurityPolicyAttributes) RecaptchaOptionsConfig() terra.ListValue[computesecuritypolicy.RecaptchaOptionsConfigAttributes] {
	return terra.ReferenceAsList[computesecuritypolicy.RecaptchaOptionsConfigAttributes](csp.ref.Append("recaptcha_options_config"))
}

func (csp computeSecurityPolicyAttributes) Rule() terra.SetValue[computesecuritypolicy.RuleAttributes] {
	return terra.ReferenceAsSet[computesecuritypolicy.RuleAttributes](csp.ref.Append("rule"))
}

func (csp computeSecurityPolicyAttributes) Timeouts() computesecuritypolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computesecuritypolicy.TimeoutsAttributes](csp.ref.Append("timeouts"))
}

type computeSecurityPolicyState struct {
	Description              string                                                `json:"description"`
	Fingerprint              string                                                `json:"fingerprint"`
	Id                       string                                                `json:"id"`
	Name                     string                                                `json:"name"`
	Project                  string                                                `json:"project"`
	SelfLink                 string                                                `json:"self_link"`
	Type                     string                                                `json:"type"`
	AdaptiveProtectionConfig []computesecuritypolicy.AdaptiveProtectionConfigState `json:"adaptive_protection_config"`
	AdvancedOptionsConfig    []computesecuritypolicy.AdvancedOptionsConfigState    `json:"advanced_options_config"`
	RecaptchaOptionsConfig   []computesecuritypolicy.RecaptchaOptionsConfigState   `json:"recaptcha_options_config"`
	Rule                     []computesecuritypolicy.RuleState                     `json:"rule"`
	Timeouts                 *computesecuritypolicy.TimeoutsState                  `json:"timeouts"`
}
