// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	binaryauthorizationpolicy "github.com/golingon/terraproviders/google/5.24.0/binaryauthorizationpolicy"
	"io"
)

// NewBinaryAuthorizationPolicy creates a new instance of [BinaryAuthorizationPolicy].
func NewBinaryAuthorizationPolicy(name string, args BinaryAuthorizationPolicyArgs) *BinaryAuthorizationPolicy {
	return &BinaryAuthorizationPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BinaryAuthorizationPolicy)(nil)

// BinaryAuthorizationPolicy represents the Terraform resource google_binary_authorization_policy.
type BinaryAuthorizationPolicy struct {
	Name      string
	Args      BinaryAuthorizationPolicyArgs
	state     *binaryAuthorizationPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BinaryAuthorizationPolicy].
func (bap *BinaryAuthorizationPolicy) Type() string {
	return "google_binary_authorization_policy"
}

// LocalName returns the local name for [BinaryAuthorizationPolicy].
func (bap *BinaryAuthorizationPolicy) LocalName() string {
	return bap.Name
}

// Configuration returns the configuration (args) for [BinaryAuthorizationPolicy].
func (bap *BinaryAuthorizationPolicy) Configuration() interface{} {
	return bap.Args
}

// DependOn is used for other resources to depend on [BinaryAuthorizationPolicy].
func (bap *BinaryAuthorizationPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(bap)
}

// Dependencies returns the list of resources [BinaryAuthorizationPolicy] depends_on.
func (bap *BinaryAuthorizationPolicy) Dependencies() terra.Dependencies {
	return bap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BinaryAuthorizationPolicy].
func (bap *BinaryAuthorizationPolicy) LifecycleManagement() *terra.Lifecycle {
	return bap.Lifecycle
}

// Attributes returns the attributes for [BinaryAuthorizationPolicy].
func (bap *BinaryAuthorizationPolicy) Attributes() binaryAuthorizationPolicyAttributes {
	return binaryAuthorizationPolicyAttributes{ref: terra.ReferenceResource(bap)}
}

// ImportState imports the given attribute values into [BinaryAuthorizationPolicy]'s state.
func (bap *BinaryAuthorizationPolicy) ImportState(av io.Reader) error {
	bap.state = &binaryAuthorizationPolicyState{}
	if err := json.NewDecoder(av).Decode(bap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bap.Type(), bap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BinaryAuthorizationPolicy] has state.
func (bap *BinaryAuthorizationPolicy) State() (*binaryAuthorizationPolicyState, bool) {
	return bap.state, bap.state != nil
}

// StateMust returns the state for [BinaryAuthorizationPolicy]. Panics if the state is nil.
func (bap *BinaryAuthorizationPolicy) StateMust() *binaryAuthorizationPolicyState {
	if bap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bap.Type(), bap.LocalName()))
	}
	return bap.state
}

// BinaryAuthorizationPolicyArgs contains the configurations for google_binary_authorization_policy.
type BinaryAuthorizationPolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// GlobalPolicyEvaluationMode: string, optional
	GlobalPolicyEvaluationMode terra.StringValue `hcl:"global_policy_evaluation_mode,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// AdmissionWhitelistPatterns: min=0
	AdmissionWhitelistPatterns []binaryauthorizationpolicy.AdmissionWhitelistPatterns `hcl:"admission_whitelist_patterns,block" validate:"min=0"`
	// ClusterAdmissionRules: min=0
	ClusterAdmissionRules []binaryauthorizationpolicy.ClusterAdmissionRules `hcl:"cluster_admission_rules,block" validate:"min=0"`
	// DefaultAdmissionRule: required
	DefaultAdmissionRule *binaryauthorizationpolicy.DefaultAdmissionRule `hcl:"default_admission_rule,block" validate:"required"`
	// Timeouts: optional
	Timeouts *binaryauthorizationpolicy.Timeouts `hcl:"timeouts,block"`
}
type binaryAuthorizationPolicyAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_binary_authorization_policy.
func (bap binaryAuthorizationPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(bap.ref.Append("description"))
}

// GlobalPolicyEvaluationMode returns a reference to field global_policy_evaluation_mode of google_binary_authorization_policy.
func (bap binaryAuthorizationPolicyAttributes) GlobalPolicyEvaluationMode() terra.StringValue {
	return terra.ReferenceAsString(bap.ref.Append("global_policy_evaluation_mode"))
}

// Id returns a reference to field id of google_binary_authorization_policy.
func (bap binaryAuthorizationPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bap.ref.Append("id"))
}

// Project returns a reference to field project of google_binary_authorization_policy.
func (bap binaryAuthorizationPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bap.ref.Append("project"))
}

func (bap binaryAuthorizationPolicyAttributes) AdmissionWhitelistPatterns() terra.ListValue[binaryauthorizationpolicy.AdmissionWhitelistPatternsAttributes] {
	return terra.ReferenceAsList[binaryauthorizationpolicy.AdmissionWhitelistPatternsAttributes](bap.ref.Append("admission_whitelist_patterns"))
}

func (bap binaryAuthorizationPolicyAttributes) ClusterAdmissionRules() terra.SetValue[binaryauthorizationpolicy.ClusterAdmissionRulesAttributes] {
	return terra.ReferenceAsSet[binaryauthorizationpolicy.ClusterAdmissionRulesAttributes](bap.ref.Append("cluster_admission_rules"))
}

func (bap binaryAuthorizationPolicyAttributes) DefaultAdmissionRule() terra.ListValue[binaryauthorizationpolicy.DefaultAdmissionRuleAttributes] {
	return terra.ReferenceAsList[binaryauthorizationpolicy.DefaultAdmissionRuleAttributes](bap.ref.Append("default_admission_rule"))
}

func (bap binaryAuthorizationPolicyAttributes) Timeouts() binaryauthorizationpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[binaryauthorizationpolicy.TimeoutsAttributes](bap.ref.Append("timeouts"))
}

type binaryAuthorizationPolicyState struct {
	Description                string                                                      `json:"description"`
	GlobalPolicyEvaluationMode string                                                      `json:"global_policy_evaluation_mode"`
	Id                         string                                                      `json:"id"`
	Project                    string                                                      `json:"project"`
	AdmissionWhitelistPatterns []binaryauthorizationpolicy.AdmissionWhitelistPatternsState `json:"admission_whitelist_patterns"`
	ClusterAdmissionRules      []binaryauthorizationpolicy.ClusterAdmissionRulesState      `json:"cluster_admission_rules"`
	DefaultAdmissionRule       []binaryauthorizationpolicy.DefaultAdmissionRuleState       `json:"default_admission_rule"`
	Timeouts                   *binaryauthorizationpolicy.TimeoutsState                    `json:"timeouts"`
}
