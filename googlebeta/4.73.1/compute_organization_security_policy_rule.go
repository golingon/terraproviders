// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeorganizationsecuritypolicyrule "github.com/golingon/terraproviders/googlebeta/4.73.1/computeorganizationsecuritypolicyrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeOrganizationSecurityPolicyRule creates a new instance of [ComputeOrganizationSecurityPolicyRule].
func NewComputeOrganizationSecurityPolicyRule(name string, args ComputeOrganizationSecurityPolicyRuleArgs) *ComputeOrganizationSecurityPolicyRule {
	return &ComputeOrganizationSecurityPolicyRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeOrganizationSecurityPolicyRule)(nil)

// ComputeOrganizationSecurityPolicyRule represents the Terraform resource google_compute_organization_security_policy_rule.
type ComputeOrganizationSecurityPolicyRule struct {
	Name      string
	Args      ComputeOrganizationSecurityPolicyRuleArgs
	state     *computeOrganizationSecurityPolicyRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeOrganizationSecurityPolicyRule].
func (cospr *ComputeOrganizationSecurityPolicyRule) Type() string {
	return "google_compute_organization_security_policy_rule"
}

// LocalName returns the local name for [ComputeOrganizationSecurityPolicyRule].
func (cospr *ComputeOrganizationSecurityPolicyRule) LocalName() string {
	return cospr.Name
}

// Configuration returns the configuration (args) for [ComputeOrganizationSecurityPolicyRule].
func (cospr *ComputeOrganizationSecurityPolicyRule) Configuration() interface{} {
	return cospr.Args
}

// DependOn is used for other resources to depend on [ComputeOrganizationSecurityPolicyRule].
func (cospr *ComputeOrganizationSecurityPolicyRule) DependOn() terra.Reference {
	return terra.ReferenceResource(cospr)
}

// Dependencies returns the list of resources [ComputeOrganizationSecurityPolicyRule] depends_on.
func (cospr *ComputeOrganizationSecurityPolicyRule) Dependencies() terra.Dependencies {
	return cospr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeOrganizationSecurityPolicyRule].
func (cospr *ComputeOrganizationSecurityPolicyRule) LifecycleManagement() *terra.Lifecycle {
	return cospr.Lifecycle
}

// Attributes returns the attributes for [ComputeOrganizationSecurityPolicyRule].
func (cospr *ComputeOrganizationSecurityPolicyRule) Attributes() computeOrganizationSecurityPolicyRuleAttributes {
	return computeOrganizationSecurityPolicyRuleAttributes{ref: terra.ReferenceResource(cospr)}
}

// ImportState imports the given attribute values into [ComputeOrganizationSecurityPolicyRule]'s state.
func (cospr *ComputeOrganizationSecurityPolicyRule) ImportState(av io.Reader) error {
	cospr.state = &computeOrganizationSecurityPolicyRuleState{}
	if err := json.NewDecoder(av).Decode(cospr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cospr.Type(), cospr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeOrganizationSecurityPolicyRule] has state.
func (cospr *ComputeOrganizationSecurityPolicyRule) State() (*computeOrganizationSecurityPolicyRuleState, bool) {
	return cospr.state, cospr.state != nil
}

// StateMust returns the state for [ComputeOrganizationSecurityPolicyRule]. Panics if the state is nil.
func (cospr *ComputeOrganizationSecurityPolicyRule) StateMust() *computeOrganizationSecurityPolicyRuleState {
	if cospr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cospr.Type(), cospr.LocalName()))
	}
	return cospr.state
}

// ComputeOrganizationSecurityPolicyRuleArgs contains the configurations for google_compute_organization_security_policy_rule.
type ComputeOrganizationSecurityPolicyRuleArgs struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Direction: string, optional
	Direction terra.StringValue `hcl:"direction,attr"`
	// EnableLogging: bool, optional
	EnableLogging terra.BoolValue `hcl:"enable_logging,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyId: string, required
	PolicyId terra.StringValue `hcl:"policy_id,attr" validate:"required"`
	// Preview: bool, optional
	Preview terra.BoolValue `hcl:"preview,attr"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// TargetResources: list of string, optional
	TargetResources terra.ListValue[terra.StringValue] `hcl:"target_resources,attr"`
	// TargetServiceAccounts: list of string, optional
	TargetServiceAccounts terra.ListValue[terra.StringValue] `hcl:"target_service_accounts,attr"`
	// Match: required
	Match *computeorganizationsecuritypolicyrule.Match `hcl:"match,block" validate:"required"`
	// Timeouts: optional
	Timeouts *computeorganizationsecuritypolicyrule.Timeouts `hcl:"timeouts,block"`
}
type computeOrganizationSecurityPolicyRuleAttributes struct {
	ref terra.Reference
}

// Action returns a reference to field action of google_compute_organization_security_policy_rule.
func (cospr computeOrganizationSecurityPolicyRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(cospr.ref.Append("action"))
}

// Description returns a reference to field description of google_compute_organization_security_policy_rule.
func (cospr computeOrganizationSecurityPolicyRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cospr.ref.Append("description"))
}

// Direction returns a reference to field direction of google_compute_organization_security_policy_rule.
func (cospr computeOrganizationSecurityPolicyRuleAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(cospr.ref.Append("direction"))
}

// EnableLogging returns a reference to field enable_logging of google_compute_organization_security_policy_rule.
func (cospr computeOrganizationSecurityPolicyRuleAttributes) EnableLogging() terra.BoolValue {
	return terra.ReferenceAsBool(cospr.ref.Append("enable_logging"))
}

// Id returns a reference to field id of google_compute_organization_security_policy_rule.
func (cospr computeOrganizationSecurityPolicyRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cospr.ref.Append("id"))
}

// PolicyId returns a reference to field policy_id of google_compute_organization_security_policy_rule.
func (cospr computeOrganizationSecurityPolicyRuleAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceAsString(cospr.ref.Append("policy_id"))
}

// Preview returns a reference to field preview of google_compute_organization_security_policy_rule.
func (cospr computeOrganizationSecurityPolicyRuleAttributes) Preview() terra.BoolValue {
	return terra.ReferenceAsBool(cospr.ref.Append("preview"))
}

// Priority returns a reference to field priority of google_compute_organization_security_policy_rule.
func (cospr computeOrganizationSecurityPolicyRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(cospr.ref.Append("priority"))
}

// TargetResources returns a reference to field target_resources of google_compute_organization_security_policy_rule.
func (cospr computeOrganizationSecurityPolicyRuleAttributes) TargetResources() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cospr.ref.Append("target_resources"))
}

// TargetServiceAccounts returns a reference to field target_service_accounts of google_compute_organization_security_policy_rule.
func (cospr computeOrganizationSecurityPolicyRuleAttributes) TargetServiceAccounts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cospr.ref.Append("target_service_accounts"))
}

func (cospr computeOrganizationSecurityPolicyRuleAttributes) Match() terra.ListValue[computeorganizationsecuritypolicyrule.MatchAttributes] {
	return terra.ReferenceAsList[computeorganizationsecuritypolicyrule.MatchAttributes](cospr.ref.Append("match"))
}

func (cospr computeOrganizationSecurityPolicyRuleAttributes) Timeouts() computeorganizationsecuritypolicyrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeorganizationsecuritypolicyrule.TimeoutsAttributes](cospr.ref.Append("timeouts"))
}

type computeOrganizationSecurityPolicyRuleState struct {
	Action                string                                               `json:"action"`
	Description           string                                               `json:"description"`
	Direction             string                                               `json:"direction"`
	EnableLogging         bool                                                 `json:"enable_logging"`
	Id                    string                                               `json:"id"`
	PolicyId              string                                               `json:"policy_id"`
	Preview               bool                                                 `json:"preview"`
	Priority              float64                                              `json:"priority"`
	TargetResources       []string                                             `json:"target_resources"`
	TargetServiceAccounts []string                                             `json:"target_service_accounts"`
	Match                 []computeorganizationsecuritypolicyrule.MatchState   `json:"match"`
	Timeouts              *computeorganizationsecuritypolicyrule.TimeoutsState `json:"timeouts"`
}
