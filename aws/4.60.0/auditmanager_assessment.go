// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	auditmanagerassessment "github.com/golingon/terraproviders/aws/4.60.0/auditmanagerassessment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAuditmanagerAssessment creates a new instance of [AuditmanagerAssessment].
func NewAuditmanagerAssessment(name string, args AuditmanagerAssessmentArgs) *AuditmanagerAssessment {
	return &AuditmanagerAssessment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AuditmanagerAssessment)(nil)

// AuditmanagerAssessment represents the Terraform resource aws_auditmanager_assessment.
type AuditmanagerAssessment struct {
	Name      string
	Args      AuditmanagerAssessmentArgs
	state     *auditmanagerAssessmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AuditmanagerAssessment].
func (aa *AuditmanagerAssessment) Type() string {
	return "aws_auditmanager_assessment"
}

// LocalName returns the local name for [AuditmanagerAssessment].
func (aa *AuditmanagerAssessment) LocalName() string {
	return aa.Name
}

// Configuration returns the configuration (args) for [AuditmanagerAssessment].
func (aa *AuditmanagerAssessment) Configuration() interface{} {
	return aa.Args
}

// DependOn is used for other resources to depend on [AuditmanagerAssessment].
func (aa *AuditmanagerAssessment) DependOn() terra.Reference {
	return terra.ReferenceResource(aa)
}

// Dependencies returns the list of resources [AuditmanagerAssessment] depends_on.
func (aa *AuditmanagerAssessment) Dependencies() terra.Dependencies {
	return aa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AuditmanagerAssessment].
func (aa *AuditmanagerAssessment) LifecycleManagement() *terra.Lifecycle {
	return aa.Lifecycle
}

// Attributes returns the attributes for [AuditmanagerAssessment].
func (aa *AuditmanagerAssessment) Attributes() auditmanagerAssessmentAttributes {
	return auditmanagerAssessmentAttributes{ref: terra.ReferenceResource(aa)}
}

// ImportState imports the given attribute values into [AuditmanagerAssessment]'s state.
func (aa *AuditmanagerAssessment) ImportState(av io.Reader) error {
	aa.state = &auditmanagerAssessmentState{}
	if err := json.NewDecoder(av).Decode(aa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aa.Type(), aa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AuditmanagerAssessment] has state.
func (aa *AuditmanagerAssessment) State() (*auditmanagerAssessmentState, bool) {
	return aa.state, aa.state != nil
}

// StateMust returns the state for [AuditmanagerAssessment]. Panics if the state is nil.
func (aa *AuditmanagerAssessment) StateMust() *auditmanagerAssessmentState {
	if aa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aa.Type(), aa.LocalName()))
	}
	return aa.state
}

// AuditmanagerAssessmentArgs contains the configurations for aws_auditmanager_assessment.
type AuditmanagerAssessmentArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FrameworkId: string, required
	FrameworkId terra.StringValue `hcl:"framework_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Roles: min=0
	Roles []auditmanagerassessment.Roles `hcl:"roles,block" validate:"min=0"`
	// RolesAll: min=0
	RolesAll []auditmanagerassessment.RolesAll `hcl:"roles_all,block" validate:"min=0"`
	// AssessmentReportsDestination: min=0
	AssessmentReportsDestination []auditmanagerassessment.AssessmentReportsDestination `hcl:"assessment_reports_destination,block" validate:"min=0"`
	// Scope: min=0
	Scope []auditmanagerassessment.Scope `hcl:"scope,block" validate:"min=0"`
}
type auditmanagerAssessmentAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_auditmanager_assessment.
func (aa auditmanagerAssessmentAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("arn"))
}

// Description returns a reference to field description of aws_auditmanager_assessment.
func (aa auditmanagerAssessmentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("description"))
}

// FrameworkId returns a reference to field framework_id of aws_auditmanager_assessment.
func (aa auditmanagerAssessmentAttributes) FrameworkId() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("framework_id"))
}

// Id returns a reference to field id of aws_auditmanager_assessment.
func (aa auditmanagerAssessmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("id"))
}

// Name returns a reference to field name of aws_auditmanager_assessment.
func (aa auditmanagerAssessmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("name"))
}

// Status returns a reference to field status of aws_auditmanager_assessment.
func (aa auditmanagerAssessmentAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_auditmanager_assessment.
func (aa auditmanagerAssessmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aa.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_auditmanager_assessment.
func (aa auditmanagerAssessmentAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aa.ref.Append("tags_all"))
}

func (aa auditmanagerAssessmentAttributes) Roles() terra.SetValue[auditmanagerassessment.RolesAttributes] {
	return terra.ReferenceAsSet[auditmanagerassessment.RolesAttributes](aa.ref.Append("roles"))
}

func (aa auditmanagerAssessmentAttributes) RolesAll() terra.SetValue[auditmanagerassessment.RolesAllAttributes] {
	return terra.ReferenceAsSet[auditmanagerassessment.RolesAllAttributes](aa.ref.Append("roles_all"))
}

func (aa auditmanagerAssessmentAttributes) AssessmentReportsDestination() terra.ListValue[auditmanagerassessment.AssessmentReportsDestinationAttributes] {
	return terra.ReferenceAsList[auditmanagerassessment.AssessmentReportsDestinationAttributes](aa.ref.Append("assessment_reports_destination"))
}

func (aa auditmanagerAssessmentAttributes) Scope() terra.ListValue[auditmanagerassessment.ScopeAttributes] {
	return terra.ReferenceAsList[auditmanagerassessment.ScopeAttributes](aa.ref.Append("scope"))
}

type auditmanagerAssessmentState struct {
	Arn                          string                                                     `json:"arn"`
	Description                  string                                                     `json:"description"`
	FrameworkId                  string                                                     `json:"framework_id"`
	Id                           string                                                     `json:"id"`
	Name                         string                                                     `json:"name"`
	Status                       string                                                     `json:"status"`
	Tags                         map[string]string                                          `json:"tags"`
	TagsAll                      map[string]string                                          `json:"tags_all"`
	Roles                        []auditmanagerassessment.RolesState                        `json:"roles"`
	RolesAll                     []auditmanagerassessment.RolesAllState                     `json:"roles_all"`
	AssessmentReportsDestination []auditmanagerassessment.AssessmentReportsDestinationState `json:"assessment_reports_destination"`
	Scope                        []auditmanagerassessment.ScopeState                        `json:"scope"`
}
