// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDetectiveInvitationAccepter creates a new instance of [DetectiveInvitationAccepter].
func NewDetectiveInvitationAccepter(name string, args DetectiveInvitationAccepterArgs) *DetectiveInvitationAccepter {
	return &DetectiveInvitationAccepter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DetectiveInvitationAccepter)(nil)

// DetectiveInvitationAccepter represents the Terraform resource aws_detective_invitation_accepter.
type DetectiveInvitationAccepter struct {
	Name      string
	Args      DetectiveInvitationAccepterArgs
	state     *detectiveInvitationAccepterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DetectiveInvitationAccepter].
func (dia *DetectiveInvitationAccepter) Type() string {
	return "aws_detective_invitation_accepter"
}

// LocalName returns the local name for [DetectiveInvitationAccepter].
func (dia *DetectiveInvitationAccepter) LocalName() string {
	return dia.Name
}

// Configuration returns the configuration (args) for [DetectiveInvitationAccepter].
func (dia *DetectiveInvitationAccepter) Configuration() interface{} {
	return dia.Args
}

// DependOn is used for other resources to depend on [DetectiveInvitationAccepter].
func (dia *DetectiveInvitationAccepter) DependOn() terra.Reference {
	return terra.ReferenceResource(dia)
}

// Dependencies returns the list of resources [DetectiveInvitationAccepter] depends_on.
func (dia *DetectiveInvitationAccepter) Dependencies() terra.Dependencies {
	return dia.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DetectiveInvitationAccepter].
func (dia *DetectiveInvitationAccepter) LifecycleManagement() *terra.Lifecycle {
	return dia.Lifecycle
}

// Attributes returns the attributes for [DetectiveInvitationAccepter].
func (dia *DetectiveInvitationAccepter) Attributes() detectiveInvitationAccepterAttributes {
	return detectiveInvitationAccepterAttributes{ref: terra.ReferenceResource(dia)}
}

// ImportState imports the given attribute values into [DetectiveInvitationAccepter]'s state.
func (dia *DetectiveInvitationAccepter) ImportState(av io.Reader) error {
	dia.state = &detectiveInvitationAccepterState{}
	if err := json.NewDecoder(av).Decode(dia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dia.Type(), dia.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DetectiveInvitationAccepter] has state.
func (dia *DetectiveInvitationAccepter) State() (*detectiveInvitationAccepterState, bool) {
	return dia.state, dia.state != nil
}

// StateMust returns the state for [DetectiveInvitationAccepter]. Panics if the state is nil.
func (dia *DetectiveInvitationAccepter) StateMust() *detectiveInvitationAccepterState {
	if dia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dia.Type(), dia.LocalName()))
	}
	return dia.state
}

// DetectiveInvitationAccepterArgs contains the configurations for aws_detective_invitation_accepter.
type DetectiveInvitationAccepterArgs struct {
	// GraphArn: string, required
	GraphArn terra.StringValue `hcl:"graph_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type detectiveInvitationAccepterAttributes struct {
	ref terra.Reference
}

// GraphArn returns a reference to field graph_arn of aws_detective_invitation_accepter.
func (dia detectiveInvitationAccepterAttributes) GraphArn() terra.StringValue {
	return terra.ReferenceAsString(dia.ref.Append("graph_arn"))
}

// Id returns a reference to field id of aws_detective_invitation_accepter.
func (dia detectiveInvitationAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dia.ref.Append("id"))
}

type detectiveInvitationAccepterState struct {
	GraphArn string `json:"graph_arn"`
	Id       string `json:"id"`
}
