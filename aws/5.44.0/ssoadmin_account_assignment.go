// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	ssoadminaccountassignment "github.com/golingon/terraproviders/aws/5.44.0/ssoadminaccountassignment"
	"io"
)

// NewSsoadminAccountAssignment creates a new instance of [SsoadminAccountAssignment].
func NewSsoadminAccountAssignment(name string, args SsoadminAccountAssignmentArgs) *SsoadminAccountAssignment {
	return &SsoadminAccountAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SsoadminAccountAssignment)(nil)

// SsoadminAccountAssignment represents the Terraform resource aws_ssoadmin_account_assignment.
type SsoadminAccountAssignment struct {
	Name      string
	Args      SsoadminAccountAssignmentArgs
	state     *ssoadminAccountAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SsoadminAccountAssignment].
func (saa *SsoadminAccountAssignment) Type() string {
	return "aws_ssoadmin_account_assignment"
}

// LocalName returns the local name for [SsoadminAccountAssignment].
func (saa *SsoadminAccountAssignment) LocalName() string {
	return saa.Name
}

// Configuration returns the configuration (args) for [SsoadminAccountAssignment].
func (saa *SsoadminAccountAssignment) Configuration() interface{} {
	return saa.Args
}

// DependOn is used for other resources to depend on [SsoadminAccountAssignment].
func (saa *SsoadminAccountAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(saa)
}

// Dependencies returns the list of resources [SsoadminAccountAssignment] depends_on.
func (saa *SsoadminAccountAssignment) Dependencies() terra.Dependencies {
	return saa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SsoadminAccountAssignment].
func (saa *SsoadminAccountAssignment) LifecycleManagement() *terra.Lifecycle {
	return saa.Lifecycle
}

// Attributes returns the attributes for [SsoadminAccountAssignment].
func (saa *SsoadminAccountAssignment) Attributes() ssoadminAccountAssignmentAttributes {
	return ssoadminAccountAssignmentAttributes{ref: terra.ReferenceResource(saa)}
}

// ImportState imports the given attribute values into [SsoadminAccountAssignment]'s state.
func (saa *SsoadminAccountAssignment) ImportState(av io.Reader) error {
	saa.state = &ssoadminAccountAssignmentState{}
	if err := json.NewDecoder(av).Decode(saa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", saa.Type(), saa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SsoadminAccountAssignment] has state.
func (saa *SsoadminAccountAssignment) State() (*ssoadminAccountAssignmentState, bool) {
	return saa.state, saa.state != nil
}

// StateMust returns the state for [SsoadminAccountAssignment]. Panics if the state is nil.
func (saa *SsoadminAccountAssignment) StateMust() *ssoadminAccountAssignmentState {
	if saa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", saa.Type(), saa.LocalName()))
	}
	return saa.state
}

// SsoadminAccountAssignmentArgs contains the configurations for aws_ssoadmin_account_assignment.
type SsoadminAccountAssignmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceArn: string, required
	InstanceArn terra.StringValue `hcl:"instance_arn,attr" validate:"required"`
	// PermissionSetArn: string, required
	PermissionSetArn terra.StringValue `hcl:"permission_set_arn,attr" validate:"required"`
	// PrincipalId: string, required
	PrincipalId terra.StringValue `hcl:"principal_id,attr" validate:"required"`
	// PrincipalType: string, required
	PrincipalType terra.StringValue `hcl:"principal_type,attr" validate:"required"`
	// TargetId: string, required
	TargetId terra.StringValue `hcl:"target_id,attr" validate:"required"`
	// TargetType: string, optional
	TargetType terra.StringValue `hcl:"target_type,attr"`
	// Timeouts: optional
	Timeouts *ssoadminaccountassignment.Timeouts `hcl:"timeouts,block"`
}
type ssoadminAccountAssignmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ssoadmin_account_assignment.
func (saa ssoadminAccountAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("id"))
}

// InstanceArn returns a reference to field instance_arn of aws_ssoadmin_account_assignment.
func (saa ssoadminAccountAssignmentAttributes) InstanceArn() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("instance_arn"))
}

// PermissionSetArn returns a reference to field permission_set_arn of aws_ssoadmin_account_assignment.
func (saa ssoadminAccountAssignmentAttributes) PermissionSetArn() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("permission_set_arn"))
}

// PrincipalId returns a reference to field principal_id of aws_ssoadmin_account_assignment.
func (saa ssoadminAccountAssignmentAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("principal_id"))
}

// PrincipalType returns a reference to field principal_type of aws_ssoadmin_account_assignment.
func (saa ssoadminAccountAssignmentAttributes) PrincipalType() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("principal_type"))
}

// TargetId returns a reference to field target_id of aws_ssoadmin_account_assignment.
func (saa ssoadminAccountAssignmentAttributes) TargetId() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("target_id"))
}

// TargetType returns a reference to field target_type of aws_ssoadmin_account_assignment.
func (saa ssoadminAccountAssignmentAttributes) TargetType() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("target_type"))
}

func (saa ssoadminAccountAssignmentAttributes) Timeouts() ssoadminaccountassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ssoadminaccountassignment.TimeoutsAttributes](saa.ref.Append("timeouts"))
}

type ssoadminAccountAssignmentState struct {
	Id               string                                   `json:"id"`
	InstanceArn      string                                   `json:"instance_arn"`
	PermissionSetArn string                                   `json:"permission_set_arn"`
	PrincipalId      string                                   `json:"principal_id"`
	PrincipalType    string                                   `json:"principal_type"`
	TargetId         string                                   `json:"target_id"`
	TargetType       string                                   `json:"target_type"`
	Timeouts         *ssoadminaccountassignment.TimeoutsState `json:"timeouts"`
}
