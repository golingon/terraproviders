// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_auditmanager_assessment_report

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

// Resource represents the Terraform resource aws_auditmanager_assessment_report.
type Resource struct {
	Name      string
	Args      Args
	state     *awsAuditmanagerAssessmentReportState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aaar *Resource) Type() string {
	return "aws_auditmanager_assessment_report"
}

// LocalName returns the local name for [Resource].
func (aaar *Resource) LocalName() string {
	return aaar.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aaar *Resource) Configuration() interface{} {
	return aaar.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aaar *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aaar)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aaar *Resource) Dependencies() terra.Dependencies {
	return aaar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aaar *Resource) LifecycleManagement() *terra.Lifecycle {
	return aaar.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aaar *Resource) Attributes() awsAuditmanagerAssessmentReportAttributes {
	return awsAuditmanagerAssessmentReportAttributes{ref: terra.ReferenceResource(aaar)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aaar *Resource) ImportState(state io.Reader) error {
	aaar.state = &awsAuditmanagerAssessmentReportState{}
	if err := json.NewDecoder(state).Decode(aaar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aaar.Type(), aaar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aaar *Resource) State() (*awsAuditmanagerAssessmentReportState, bool) {
	return aaar.state, aaar.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aaar *Resource) StateMust() *awsAuditmanagerAssessmentReportState {
	if aaar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aaar.Type(), aaar.LocalName()))
	}
	return aaar.state
}

// Args contains the configurations for aws_auditmanager_assessment_report.
type Args struct {
	// AssessmentId: string, required
	AssessmentId terra.StringValue `hcl:"assessment_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type awsAuditmanagerAssessmentReportAttributes struct {
	ref terra.Reference
}

// AssessmentId returns a reference to field assessment_id of aws_auditmanager_assessment_report.
func (aaar awsAuditmanagerAssessmentReportAttributes) AssessmentId() terra.StringValue {
	return terra.ReferenceAsString(aaar.ref.Append("assessment_id"))
}

// Author returns a reference to field author of aws_auditmanager_assessment_report.
func (aaar awsAuditmanagerAssessmentReportAttributes) Author() terra.StringValue {
	return terra.ReferenceAsString(aaar.ref.Append("author"))
}

// Description returns a reference to field description of aws_auditmanager_assessment_report.
func (aaar awsAuditmanagerAssessmentReportAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aaar.ref.Append("description"))
}

// Id returns a reference to field id of aws_auditmanager_assessment_report.
func (aaar awsAuditmanagerAssessmentReportAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aaar.ref.Append("id"))
}

// Name returns a reference to field name of aws_auditmanager_assessment_report.
func (aaar awsAuditmanagerAssessmentReportAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aaar.ref.Append("name"))
}

// Status returns a reference to field status of aws_auditmanager_assessment_report.
func (aaar awsAuditmanagerAssessmentReportAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(aaar.ref.Append("status"))
}

type awsAuditmanagerAssessmentReportState struct {
	AssessmentId string `json:"assessment_id"`
	Author       string `json:"author"`
	Description  string `json:"description"`
	Id           string `json:"id"`
	Name         string `json:"name"`
	Status       string `json:"status"`
}
