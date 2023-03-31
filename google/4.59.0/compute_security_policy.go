// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computesecuritypolicy "github.com/golingon/terraproviders/google/4.59.0/computesecuritypolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeSecurityPolicy(name string, args ComputeSecurityPolicyArgs) *ComputeSecurityPolicy {
	return &ComputeSecurityPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeSecurityPolicy)(nil)

type ComputeSecurityPolicy struct {
	Name  string
	Args  ComputeSecurityPolicyArgs
	state *computeSecurityPolicyState
}

func (csp *ComputeSecurityPolicy) Type() string {
	return "google_compute_security_policy"
}

func (csp *ComputeSecurityPolicy) LocalName() string {
	return csp.Name
}

func (csp *ComputeSecurityPolicy) Configuration() interface{} {
	return csp.Args
}

func (csp *ComputeSecurityPolicy) Attributes() computeSecurityPolicyAttributes {
	return computeSecurityPolicyAttributes{ref: terra.ReferenceResource(csp)}
}

func (csp *ComputeSecurityPolicy) ImportState(av io.Reader) error {
	csp.state = &computeSecurityPolicyState{}
	if err := json.NewDecoder(av).Decode(csp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csp.Type(), csp.LocalName(), err)
	}
	return nil
}

func (csp *ComputeSecurityPolicy) State() (*computeSecurityPolicyState, bool) {
	return csp.state, csp.state != nil
}

func (csp *ComputeSecurityPolicy) StateMust() *computeSecurityPolicyState {
	if csp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csp.Type(), csp.LocalName()))
	}
	return csp.state
}

func (csp *ComputeSecurityPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(csp)
}

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
	// DependsOn contains resources that ComputeSecurityPolicy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeSecurityPolicyAttributes struct {
	ref terra.Reference
}

func (csp computeSecurityPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceString(csp.ref.Append("description"))
}

func (csp computeSecurityPolicyAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceString(csp.ref.Append("fingerprint"))
}

func (csp computeSecurityPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(csp.ref.Append("id"))
}

func (csp computeSecurityPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceString(csp.ref.Append("name"))
}

func (csp computeSecurityPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceString(csp.ref.Append("project"))
}

func (csp computeSecurityPolicyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(csp.ref.Append("self_link"))
}

func (csp computeSecurityPolicyAttributes) Type() terra.StringValue {
	return terra.ReferenceString(csp.ref.Append("type"))
}

func (csp computeSecurityPolicyAttributes) AdaptiveProtectionConfig() terra.ListValue[computesecuritypolicy.AdaptiveProtectionConfigAttributes] {
	return terra.ReferenceList[computesecuritypolicy.AdaptiveProtectionConfigAttributes](csp.ref.Append("adaptive_protection_config"))
}

func (csp computeSecurityPolicyAttributes) AdvancedOptionsConfig() terra.ListValue[computesecuritypolicy.AdvancedOptionsConfigAttributes] {
	return terra.ReferenceList[computesecuritypolicy.AdvancedOptionsConfigAttributes](csp.ref.Append("advanced_options_config"))
}

func (csp computeSecurityPolicyAttributes) RecaptchaOptionsConfig() terra.ListValue[computesecuritypolicy.RecaptchaOptionsConfigAttributes] {
	return terra.ReferenceList[computesecuritypolicy.RecaptchaOptionsConfigAttributes](csp.ref.Append("recaptcha_options_config"))
}

func (csp computeSecurityPolicyAttributes) Rule() terra.SetValue[computesecuritypolicy.RuleAttributes] {
	return terra.ReferenceSet[computesecuritypolicy.RuleAttributes](csp.ref.Append("rule"))
}

func (csp computeSecurityPolicyAttributes) Timeouts() computesecuritypolicy.TimeoutsAttributes {
	return terra.ReferenceSingle[computesecuritypolicy.TimeoutsAttributes](csp.ref.Append("timeouts"))
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
