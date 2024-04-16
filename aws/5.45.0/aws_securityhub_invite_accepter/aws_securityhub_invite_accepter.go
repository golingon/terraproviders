// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_securityhub_invite_accepter

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

// Resource represents the Terraform resource aws_securityhub_invite_accepter.
type Resource struct {
	Name      string
	Args      Args
	state     *awsSecurityhubInviteAccepterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asia *Resource) Type() string {
	return "aws_securityhub_invite_accepter"
}

// LocalName returns the local name for [Resource].
func (asia *Resource) LocalName() string {
	return asia.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asia *Resource) Configuration() interface{} {
	return asia.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asia *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asia)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asia *Resource) Dependencies() terra.Dependencies {
	return asia.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asia *Resource) LifecycleManagement() *terra.Lifecycle {
	return asia.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asia *Resource) Attributes() awsSecurityhubInviteAccepterAttributes {
	return awsSecurityhubInviteAccepterAttributes{ref: terra.ReferenceResource(asia)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asia *Resource) ImportState(state io.Reader) error {
	asia.state = &awsSecurityhubInviteAccepterState{}
	if err := json.NewDecoder(state).Decode(asia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asia.Type(), asia.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asia *Resource) State() (*awsSecurityhubInviteAccepterState, bool) {
	return asia.state, asia.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asia *Resource) StateMust() *awsSecurityhubInviteAccepterState {
	if asia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asia.Type(), asia.LocalName()))
	}
	return asia.state
}

// Args contains the configurations for aws_securityhub_invite_accepter.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MasterId: string, required
	MasterId terra.StringValue `hcl:"master_id,attr" validate:"required"`
}

type awsSecurityhubInviteAccepterAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_securityhub_invite_accepter.
func (asia awsSecurityhubInviteAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asia.ref.Append("id"))
}

// InvitationId returns a reference to field invitation_id of aws_securityhub_invite_accepter.
func (asia awsSecurityhubInviteAccepterAttributes) InvitationId() terra.StringValue {
	return terra.ReferenceAsString(asia.ref.Append("invitation_id"))
}

// MasterId returns a reference to field master_id of aws_securityhub_invite_accepter.
func (asia awsSecurityhubInviteAccepterAttributes) MasterId() terra.StringValue {
	return terra.ReferenceAsString(asia.ref.Append("master_id"))
}

type awsSecurityhubInviteAccepterState struct {
	Id           string `json:"id"`
	InvitationId string `json:"invitation_id"`
	MasterId     string `json:"master_id"`
}
