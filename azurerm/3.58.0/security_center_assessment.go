// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	securitycenterassessment "github.com/golingon/terraproviders/azurerm/3.58.0/securitycenterassessment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSecurityCenterAssessment creates a new instance of [SecurityCenterAssessment].
func NewSecurityCenterAssessment(name string, args SecurityCenterAssessmentArgs) *SecurityCenterAssessment {
	return &SecurityCenterAssessment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecurityCenterAssessment)(nil)

// SecurityCenterAssessment represents the Terraform resource azurerm_security_center_assessment.
type SecurityCenterAssessment struct {
	Name      string
	Args      SecurityCenterAssessmentArgs
	state     *securityCenterAssessmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecurityCenterAssessment].
func (sca *SecurityCenterAssessment) Type() string {
	return "azurerm_security_center_assessment"
}

// LocalName returns the local name for [SecurityCenterAssessment].
func (sca *SecurityCenterAssessment) LocalName() string {
	return sca.Name
}

// Configuration returns the configuration (args) for [SecurityCenterAssessment].
func (sca *SecurityCenterAssessment) Configuration() interface{} {
	return sca.Args
}

// DependOn is used for other resources to depend on [SecurityCenterAssessment].
func (sca *SecurityCenterAssessment) DependOn() terra.Reference {
	return terra.ReferenceResource(sca)
}

// Dependencies returns the list of resources [SecurityCenterAssessment] depends_on.
func (sca *SecurityCenterAssessment) Dependencies() terra.Dependencies {
	return sca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecurityCenterAssessment].
func (sca *SecurityCenterAssessment) LifecycleManagement() *terra.Lifecycle {
	return sca.Lifecycle
}

// Attributes returns the attributes for [SecurityCenterAssessment].
func (sca *SecurityCenterAssessment) Attributes() securityCenterAssessmentAttributes {
	return securityCenterAssessmentAttributes{ref: terra.ReferenceResource(sca)}
}

// ImportState imports the given attribute values into [SecurityCenterAssessment]'s state.
func (sca *SecurityCenterAssessment) ImportState(av io.Reader) error {
	sca.state = &securityCenterAssessmentState{}
	if err := json.NewDecoder(av).Decode(sca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sca.Type(), sca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecurityCenterAssessment] has state.
func (sca *SecurityCenterAssessment) State() (*securityCenterAssessmentState, bool) {
	return sca.state, sca.state != nil
}

// StateMust returns the state for [SecurityCenterAssessment]. Panics if the state is nil.
func (sca *SecurityCenterAssessment) StateMust() *securityCenterAssessmentState {
	if sca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sca.Type(), sca.LocalName()))
	}
	return sca.state
}

// SecurityCenterAssessmentArgs contains the configurations for azurerm_security_center_assessment.
type SecurityCenterAssessmentArgs struct {
	// AdditionalData: map of string, optional
	AdditionalData terra.MapValue[terra.StringValue] `hcl:"additional_data,attr"`
	// AssessmentPolicyId: string, required
	AssessmentPolicyId terra.StringValue `hcl:"assessment_policy_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TargetResourceId: string, required
	TargetResourceId terra.StringValue `hcl:"target_resource_id,attr" validate:"required"`
	// Status: required
	Status *securitycenterassessment.Status `hcl:"status,block" validate:"required"`
	// Timeouts: optional
	Timeouts *securitycenterassessment.Timeouts `hcl:"timeouts,block"`
}
type securityCenterAssessmentAttributes struct {
	ref terra.Reference
}

// AdditionalData returns a reference to field additional_data of azurerm_security_center_assessment.
func (sca securityCenterAssessmentAttributes) AdditionalData() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sca.ref.Append("additional_data"))
}

// AssessmentPolicyId returns a reference to field assessment_policy_id of azurerm_security_center_assessment.
func (sca securityCenterAssessmentAttributes) AssessmentPolicyId() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("assessment_policy_id"))
}

// Id returns a reference to field id of azurerm_security_center_assessment.
func (sca securityCenterAssessmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("id"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_security_center_assessment.
func (sca securityCenterAssessmentAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("target_resource_id"))
}

func (sca securityCenterAssessmentAttributes) Status() terra.ListValue[securitycenterassessment.StatusAttributes] {
	return terra.ReferenceAsList[securitycenterassessment.StatusAttributes](sca.ref.Append("status"))
}

func (sca securityCenterAssessmentAttributes) Timeouts() securitycenterassessment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[securitycenterassessment.TimeoutsAttributes](sca.ref.Append("timeouts"))
}

type securityCenterAssessmentState struct {
	AdditionalData     map[string]string                       `json:"additional_data"`
	AssessmentPolicyId string                                  `json:"assessment_policy_id"`
	Id                 string                                  `json:"id"`
	TargetResourceId   string                                  `json:"target_resource_id"`
	Status             []securitycenterassessment.StatusState  `json:"status"`
	Timeouts           *securitycenterassessment.TimeoutsState `json:"timeouts"`
}
