// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	securitycenterassessmentpolicy "github.com/golingon/terraproviders/azurerm/3.64.0/securitycenterassessmentpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSecurityCenterAssessmentPolicy creates a new instance of [SecurityCenterAssessmentPolicy].
func NewSecurityCenterAssessmentPolicy(name string, args SecurityCenterAssessmentPolicyArgs) *SecurityCenterAssessmentPolicy {
	return &SecurityCenterAssessmentPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecurityCenterAssessmentPolicy)(nil)

// SecurityCenterAssessmentPolicy represents the Terraform resource azurerm_security_center_assessment_policy.
type SecurityCenterAssessmentPolicy struct {
	Name      string
	Args      SecurityCenterAssessmentPolicyArgs
	state     *securityCenterAssessmentPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecurityCenterAssessmentPolicy].
func (scap *SecurityCenterAssessmentPolicy) Type() string {
	return "azurerm_security_center_assessment_policy"
}

// LocalName returns the local name for [SecurityCenterAssessmentPolicy].
func (scap *SecurityCenterAssessmentPolicy) LocalName() string {
	return scap.Name
}

// Configuration returns the configuration (args) for [SecurityCenterAssessmentPolicy].
func (scap *SecurityCenterAssessmentPolicy) Configuration() interface{} {
	return scap.Args
}

// DependOn is used for other resources to depend on [SecurityCenterAssessmentPolicy].
func (scap *SecurityCenterAssessmentPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(scap)
}

// Dependencies returns the list of resources [SecurityCenterAssessmentPolicy] depends_on.
func (scap *SecurityCenterAssessmentPolicy) Dependencies() terra.Dependencies {
	return scap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecurityCenterAssessmentPolicy].
func (scap *SecurityCenterAssessmentPolicy) LifecycleManagement() *terra.Lifecycle {
	return scap.Lifecycle
}

// Attributes returns the attributes for [SecurityCenterAssessmentPolicy].
func (scap *SecurityCenterAssessmentPolicy) Attributes() securityCenterAssessmentPolicyAttributes {
	return securityCenterAssessmentPolicyAttributes{ref: terra.ReferenceResource(scap)}
}

// ImportState imports the given attribute values into [SecurityCenterAssessmentPolicy]'s state.
func (scap *SecurityCenterAssessmentPolicy) ImportState(av io.Reader) error {
	scap.state = &securityCenterAssessmentPolicyState{}
	if err := json.NewDecoder(av).Decode(scap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scap.Type(), scap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecurityCenterAssessmentPolicy] has state.
func (scap *SecurityCenterAssessmentPolicy) State() (*securityCenterAssessmentPolicyState, bool) {
	return scap.state, scap.state != nil
}

// StateMust returns the state for [SecurityCenterAssessmentPolicy]. Panics if the state is nil.
func (scap *SecurityCenterAssessmentPolicy) StateMust() *securityCenterAssessmentPolicyState {
	if scap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scap.Type(), scap.LocalName()))
	}
	return scap.state
}

// SecurityCenterAssessmentPolicyArgs contains the configurations for azurerm_security_center_assessment_policy.
type SecurityCenterAssessmentPolicyArgs struct {
	// Categories: set of string, optional
	Categories terra.SetValue[terra.StringValue] `hcl:"categories,attr"`
	// Description: string, required
	Description terra.StringValue `hcl:"description,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImplementationEffort: string, optional
	ImplementationEffort terra.StringValue `hcl:"implementation_effort,attr"`
	// RemediationDescription: string, optional
	RemediationDescription terra.StringValue `hcl:"remediation_description,attr"`
	// Severity: string, optional
	Severity terra.StringValue `hcl:"severity,attr"`
	// Threats: set of string, optional
	Threats terra.SetValue[terra.StringValue] `hcl:"threats,attr"`
	// UserImpact: string, optional
	UserImpact terra.StringValue `hcl:"user_impact,attr"`
	// Timeouts: optional
	Timeouts *securitycenterassessmentpolicy.Timeouts `hcl:"timeouts,block"`
}
type securityCenterAssessmentPolicyAttributes struct {
	ref terra.Reference
}

// Categories returns a reference to field categories of azurerm_security_center_assessment_policy.
func (scap securityCenterAssessmentPolicyAttributes) Categories() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](scap.ref.Append("categories"))
}

// Description returns a reference to field description of azurerm_security_center_assessment_policy.
func (scap securityCenterAssessmentPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(scap.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_security_center_assessment_policy.
func (scap securityCenterAssessmentPolicyAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(scap.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_security_center_assessment_policy.
func (scap securityCenterAssessmentPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scap.ref.Append("id"))
}

// ImplementationEffort returns a reference to field implementation_effort of azurerm_security_center_assessment_policy.
func (scap securityCenterAssessmentPolicyAttributes) ImplementationEffort() terra.StringValue {
	return terra.ReferenceAsString(scap.ref.Append("implementation_effort"))
}

// Name returns a reference to field name of azurerm_security_center_assessment_policy.
func (scap securityCenterAssessmentPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scap.ref.Append("name"))
}

// RemediationDescription returns a reference to field remediation_description of azurerm_security_center_assessment_policy.
func (scap securityCenterAssessmentPolicyAttributes) RemediationDescription() terra.StringValue {
	return terra.ReferenceAsString(scap.ref.Append("remediation_description"))
}

// Severity returns a reference to field severity of azurerm_security_center_assessment_policy.
func (scap securityCenterAssessmentPolicyAttributes) Severity() terra.StringValue {
	return terra.ReferenceAsString(scap.ref.Append("severity"))
}

// Threats returns a reference to field threats of azurerm_security_center_assessment_policy.
func (scap securityCenterAssessmentPolicyAttributes) Threats() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](scap.ref.Append("threats"))
}

// UserImpact returns a reference to field user_impact of azurerm_security_center_assessment_policy.
func (scap securityCenterAssessmentPolicyAttributes) UserImpact() terra.StringValue {
	return terra.ReferenceAsString(scap.ref.Append("user_impact"))
}

func (scap securityCenterAssessmentPolicyAttributes) Timeouts() securitycenterassessmentpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[securitycenterassessmentpolicy.TimeoutsAttributes](scap.ref.Append("timeouts"))
}

type securityCenterAssessmentPolicyState struct {
	Categories             []string                                      `json:"categories"`
	Description            string                                        `json:"description"`
	DisplayName            string                                        `json:"display_name"`
	Id                     string                                        `json:"id"`
	ImplementationEffort   string                                        `json:"implementation_effort"`
	Name                   string                                        `json:"name"`
	RemediationDescription string                                        `json:"remediation_description"`
	Severity               string                                        `json:"severity"`
	Threats                []string                                      `json:"threats"`
	UserImpact             string                                        `json:"user_impact"`
	Timeouts               *securitycenterassessmentpolicy.TimeoutsState `json:"timeouts"`
}
