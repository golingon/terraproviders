// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	guarddutyinviteaccepter "github.com/golingon/terraproviders/aws/5.7.0/guarddutyinviteaccepter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGuarddutyInviteAccepter creates a new instance of [GuarddutyInviteAccepter].
func NewGuarddutyInviteAccepter(name string, args GuarddutyInviteAccepterArgs) *GuarddutyInviteAccepter {
	return &GuarddutyInviteAccepter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GuarddutyInviteAccepter)(nil)

// GuarddutyInviteAccepter represents the Terraform resource aws_guardduty_invite_accepter.
type GuarddutyInviteAccepter struct {
	Name      string
	Args      GuarddutyInviteAccepterArgs
	state     *guarddutyInviteAccepterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GuarddutyInviteAccepter].
func (gia *GuarddutyInviteAccepter) Type() string {
	return "aws_guardduty_invite_accepter"
}

// LocalName returns the local name for [GuarddutyInviteAccepter].
func (gia *GuarddutyInviteAccepter) LocalName() string {
	return gia.Name
}

// Configuration returns the configuration (args) for [GuarddutyInviteAccepter].
func (gia *GuarddutyInviteAccepter) Configuration() interface{} {
	return gia.Args
}

// DependOn is used for other resources to depend on [GuarddutyInviteAccepter].
func (gia *GuarddutyInviteAccepter) DependOn() terra.Reference {
	return terra.ReferenceResource(gia)
}

// Dependencies returns the list of resources [GuarddutyInviteAccepter] depends_on.
func (gia *GuarddutyInviteAccepter) Dependencies() terra.Dependencies {
	return gia.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GuarddutyInviteAccepter].
func (gia *GuarddutyInviteAccepter) LifecycleManagement() *terra.Lifecycle {
	return gia.Lifecycle
}

// Attributes returns the attributes for [GuarddutyInviteAccepter].
func (gia *GuarddutyInviteAccepter) Attributes() guarddutyInviteAccepterAttributes {
	return guarddutyInviteAccepterAttributes{ref: terra.ReferenceResource(gia)}
}

// ImportState imports the given attribute values into [GuarddutyInviteAccepter]'s state.
func (gia *GuarddutyInviteAccepter) ImportState(av io.Reader) error {
	gia.state = &guarddutyInviteAccepterState{}
	if err := json.NewDecoder(av).Decode(gia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gia.Type(), gia.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GuarddutyInviteAccepter] has state.
func (gia *GuarddutyInviteAccepter) State() (*guarddutyInviteAccepterState, bool) {
	return gia.state, gia.state != nil
}

// StateMust returns the state for [GuarddutyInviteAccepter]. Panics if the state is nil.
func (gia *GuarddutyInviteAccepter) StateMust() *guarddutyInviteAccepterState {
	if gia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gia.Type(), gia.LocalName()))
	}
	return gia.state
}

// GuarddutyInviteAccepterArgs contains the configurations for aws_guardduty_invite_accepter.
type GuarddutyInviteAccepterArgs struct {
	// DetectorId: string, required
	DetectorId terra.StringValue `hcl:"detector_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MasterAccountId: string, required
	MasterAccountId terra.StringValue `hcl:"master_account_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *guarddutyinviteaccepter.Timeouts `hcl:"timeouts,block"`
}
type guarddutyInviteAccepterAttributes struct {
	ref terra.Reference
}

// DetectorId returns a reference to field detector_id of aws_guardduty_invite_accepter.
func (gia guarddutyInviteAccepterAttributes) DetectorId() terra.StringValue {
	return terra.ReferenceAsString(gia.ref.Append("detector_id"))
}

// Id returns a reference to field id of aws_guardduty_invite_accepter.
func (gia guarddutyInviteAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gia.ref.Append("id"))
}

// MasterAccountId returns a reference to field master_account_id of aws_guardduty_invite_accepter.
func (gia guarddutyInviteAccepterAttributes) MasterAccountId() terra.StringValue {
	return terra.ReferenceAsString(gia.ref.Append("master_account_id"))
}

func (gia guarddutyInviteAccepterAttributes) Timeouts() guarddutyinviteaccepter.TimeoutsAttributes {
	return terra.ReferenceAsSingle[guarddutyinviteaccepter.TimeoutsAttributes](gia.ref.Append("timeouts"))
}

type guarddutyInviteAccepterState struct {
	DetectorId      string                                 `json:"detector_id"`
	Id              string                                 `json:"id"`
	MasterAccountId string                                 `json:"master_account_id"`
	Timeouts        *guarddutyinviteaccepter.TimeoutsState `json:"timeouts"`
}
