// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewSsoadminApplicationAssignment creates a new instance of [SsoadminApplicationAssignment].
func NewSsoadminApplicationAssignment(name string, args SsoadminApplicationAssignmentArgs) *SsoadminApplicationAssignment {
	return &SsoadminApplicationAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SsoadminApplicationAssignment)(nil)

// SsoadminApplicationAssignment represents the Terraform resource aws_ssoadmin_application_assignment.
type SsoadminApplicationAssignment struct {
	Name      string
	Args      SsoadminApplicationAssignmentArgs
	state     *ssoadminApplicationAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SsoadminApplicationAssignment].
func (saa *SsoadminApplicationAssignment) Type() string {
	return "aws_ssoadmin_application_assignment"
}

// LocalName returns the local name for [SsoadminApplicationAssignment].
func (saa *SsoadminApplicationAssignment) LocalName() string {
	return saa.Name
}

// Configuration returns the configuration (args) for [SsoadminApplicationAssignment].
func (saa *SsoadminApplicationAssignment) Configuration() interface{} {
	return saa.Args
}

// DependOn is used for other resources to depend on [SsoadminApplicationAssignment].
func (saa *SsoadminApplicationAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(saa)
}

// Dependencies returns the list of resources [SsoadminApplicationAssignment] depends_on.
func (saa *SsoadminApplicationAssignment) Dependencies() terra.Dependencies {
	return saa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SsoadminApplicationAssignment].
func (saa *SsoadminApplicationAssignment) LifecycleManagement() *terra.Lifecycle {
	return saa.Lifecycle
}

// Attributes returns the attributes for [SsoadminApplicationAssignment].
func (saa *SsoadminApplicationAssignment) Attributes() ssoadminApplicationAssignmentAttributes {
	return ssoadminApplicationAssignmentAttributes{ref: terra.ReferenceResource(saa)}
}

// ImportState imports the given attribute values into [SsoadminApplicationAssignment]'s state.
func (saa *SsoadminApplicationAssignment) ImportState(av io.Reader) error {
	saa.state = &ssoadminApplicationAssignmentState{}
	if err := json.NewDecoder(av).Decode(saa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", saa.Type(), saa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SsoadminApplicationAssignment] has state.
func (saa *SsoadminApplicationAssignment) State() (*ssoadminApplicationAssignmentState, bool) {
	return saa.state, saa.state != nil
}

// StateMust returns the state for [SsoadminApplicationAssignment]. Panics if the state is nil.
func (saa *SsoadminApplicationAssignment) StateMust() *ssoadminApplicationAssignmentState {
	if saa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", saa.Type(), saa.LocalName()))
	}
	return saa.state
}

// SsoadminApplicationAssignmentArgs contains the configurations for aws_ssoadmin_application_assignment.
type SsoadminApplicationAssignmentArgs struct {
	// ApplicationArn: string, required
	ApplicationArn terra.StringValue `hcl:"application_arn,attr" validate:"required"`
	// PrincipalId: string, required
	PrincipalId terra.StringValue `hcl:"principal_id,attr" validate:"required"`
	// PrincipalType: string, required
	PrincipalType terra.StringValue `hcl:"principal_type,attr" validate:"required"`
}
type ssoadminApplicationAssignmentAttributes struct {
	ref terra.Reference
}

// ApplicationArn returns a reference to field application_arn of aws_ssoadmin_application_assignment.
func (saa ssoadminApplicationAssignmentAttributes) ApplicationArn() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("application_arn"))
}

// Id returns a reference to field id of aws_ssoadmin_application_assignment.
func (saa ssoadminApplicationAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("id"))
}

// PrincipalId returns a reference to field principal_id of aws_ssoadmin_application_assignment.
func (saa ssoadminApplicationAssignmentAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("principal_id"))
}

// PrincipalType returns a reference to field principal_type of aws_ssoadmin_application_assignment.
func (saa ssoadminApplicationAssignmentAttributes) PrincipalType() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("principal_type"))
}

type ssoadminApplicationAssignmentState struct {
	ApplicationArn string `json:"application_arn"`
	Id             string `json:"id"`
	PrincipalId    string `json:"principal_id"`
	PrincipalType  string `json:"principal_type"`
}
