// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	macie2invitationaccepter "github.com/golingon/terraproviders/aws/5.44.0/macie2invitationaccepter"
	"io"
)

// NewMacie2InvitationAccepter creates a new instance of [Macie2InvitationAccepter].
func NewMacie2InvitationAccepter(name string, args Macie2InvitationAccepterArgs) *Macie2InvitationAccepter {
	return &Macie2InvitationAccepter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Macie2InvitationAccepter)(nil)

// Macie2InvitationAccepter represents the Terraform resource aws_macie2_invitation_accepter.
type Macie2InvitationAccepter struct {
	Name      string
	Args      Macie2InvitationAccepterArgs
	state     *macie2InvitationAccepterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Macie2InvitationAccepter].
func (mia *Macie2InvitationAccepter) Type() string {
	return "aws_macie2_invitation_accepter"
}

// LocalName returns the local name for [Macie2InvitationAccepter].
func (mia *Macie2InvitationAccepter) LocalName() string {
	return mia.Name
}

// Configuration returns the configuration (args) for [Macie2InvitationAccepter].
func (mia *Macie2InvitationAccepter) Configuration() interface{} {
	return mia.Args
}

// DependOn is used for other resources to depend on [Macie2InvitationAccepter].
func (mia *Macie2InvitationAccepter) DependOn() terra.Reference {
	return terra.ReferenceResource(mia)
}

// Dependencies returns the list of resources [Macie2InvitationAccepter] depends_on.
func (mia *Macie2InvitationAccepter) Dependencies() terra.Dependencies {
	return mia.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Macie2InvitationAccepter].
func (mia *Macie2InvitationAccepter) LifecycleManagement() *terra.Lifecycle {
	return mia.Lifecycle
}

// Attributes returns the attributes for [Macie2InvitationAccepter].
func (mia *Macie2InvitationAccepter) Attributes() macie2InvitationAccepterAttributes {
	return macie2InvitationAccepterAttributes{ref: terra.ReferenceResource(mia)}
}

// ImportState imports the given attribute values into [Macie2InvitationAccepter]'s state.
func (mia *Macie2InvitationAccepter) ImportState(av io.Reader) error {
	mia.state = &macie2InvitationAccepterState{}
	if err := json.NewDecoder(av).Decode(mia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mia.Type(), mia.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Macie2InvitationAccepter] has state.
func (mia *Macie2InvitationAccepter) State() (*macie2InvitationAccepterState, bool) {
	return mia.state, mia.state != nil
}

// StateMust returns the state for [Macie2InvitationAccepter]. Panics if the state is nil.
func (mia *Macie2InvitationAccepter) StateMust() *macie2InvitationAccepterState {
	if mia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mia.Type(), mia.LocalName()))
	}
	return mia.state
}

// Macie2InvitationAccepterArgs contains the configurations for aws_macie2_invitation_accepter.
type Macie2InvitationAccepterArgs struct {
	// AdministratorAccountId: string, required
	AdministratorAccountId terra.StringValue `hcl:"administrator_account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *macie2invitationaccepter.Timeouts `hcl:"timeouts,block"`
}
type macie2InvitationAccepterAttributes struct {
	ref terra.Reference
}

// AdministratorAccountId returns a reference to field administrator_account_id of aws_macie2_invitation_accepter.
func (mia macie2InvitationAccepterAttributes) AdministratorAccountId() terra.StringValue {
	return terra.ReferenceAsString(mia.ref.Append("administrator_account_id"))
}

// Id returns a reference to field id of aws_macie2_invitation_accepter.
func (mia macie2InvitationAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mia.ref.Append("id"))
}

// InvitationId returns a reference to field invitation_id of aws_macie2_invitation_accepter.
func (mia macie2InvitationAccepterAttributes) InvitationId() terra.StringValue {
	return terra.ReferenceAsString(mia.ref.Append("invitation_id"))
}

func (mia macie2InvitationAccepterAttributes) Timeouts() macie2invitationaccepter.TimeoutsAttributes {
	return terra.ReferenceAsSingle[macie2invitationaccepter.TimeoutsAttributes](mia.ref.Append("timeouts"))
}

type macie2InvitationAccepterState struct {
	AdministratorAccountId string                                  `json:"administrator_account_id"`
	Id                     string                                  `json:"id"`
	InvitationId           string                                  `json:"invitation_id"`
	Timeouts               *macie2invitationaccepter.TimeoutsState `json:"timeouts"`
}
