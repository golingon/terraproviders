// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ramresourceshareaccepter "github.com/golingon/terraproviders/aws/5.8.0/ramresourceshareaccepter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRamResourceShareAccepter creates a new instance of [RamResourceShareAccepter].
func NewRamResourceShareAccepter(name string, args RamResourceShareAccepterArgs) *RamResourceShareAccepter {
	return &RamResourceShareAccepter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RamResourceShareAccepter)(nil)

// RamResourceShareAccepter represents the Terraform resource aws_ram_resource_share_accepter.
type RamResourceShareAccepter struct {
	Name      string
	Args      RamResourceShareAccepterArgs
	state     *ramResourceShareAccepterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RamResourceShareAccepter].
func (rrsa *RamResourceShareAccepter) Type() string {
	return "aws_ram_resource_share_accepter"
}

// LocalName returns the local name for [RamResourceShareAccepter].
func (rrsa *RamResourceShareAccepter) LocalName() string {
	return rrsa.Name
}

// Configuration returns the configuration (args) for [RamResourceShareAccepter].
func (rrsa *RamResourceShareAccepter) Configuration() interface{} {
	return rrsa.Args
}

// DependOn is used for other resources to depend on [RamResourceShareAccepter].
func (rrsa *RamResourceShareAccepter) DependOn() terra.Reference {
	return terra.ReferenceResource(rrsa)
}

// Dependencies returns the list of resources [RamResourceShareAccepter] depends_on.
func (rrsa *RamResourceShareAccepter) Dependencies() terra.Dependencies {
	return rrsa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RamResourceShareAccepter].
func (rrsa *RamResourceShareAccepter) LifecycleManagement() *terra.Lifecycle {
	return rrsa.Lifecycle
}

// Attributes returns the attributes for [RamResourceShareAccepter].
func (rrsa *RamResourceShareAccepter) Attributes() ramResourceShareAccepterAttributes {
	return ramResourceShareAccepterAttributes{ref: terra.ReferenceResource(rrsa)}
}

// ImportState imports the given attribute values into [RamResourceShareAccepter]'s state.
func (rrsa *RamResourceShareAccepter) ImportState(av io.Reader) error {
	rrsa.state = &ramResourceShareAccepterState{}
	if err := json.NewDecoder(av).Decode(rrsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rrsa.Type(), rrsa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RamResourceShareAccepter] has state.
func (rrsa *RamResourceShareAccepter) State() (*ramResourceShareAccepterState, bool) {
	return rrsa.state, rrsa.state != nil
}

// StateMust returns the state for [RamResourceShareAccepter]. Panics if the state is nil.
func (rrsa *RamResourceShareAccepter) StateMust() *ramResourceShareAccepterState {
	if rrsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rrsa.Type(), rrsa.LocalName()))
	}
	return rrsa.state
}

// RamResourceShareAccepterArgs contains the configurations for aws_ram_resource_share_accepter.
type RamResourceShareAccepterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ShareArn: string, required
	ShareArn terra.StringValue `hcl:"share_arn,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *ramresourceshareaccepter.Timeouts `hcl:"timeouts,block"`
}
type ramResourceShareAccepterAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ram_resource_share_accepter.
func (rrsa ramResourceShareAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rrsa.ref.Append("id"))
}

// InvitationArn returns a reference to field invitation_arn of aws_ram_resource_share_accepter.
func (rrsa ramResourceShareAccepterAttributes) InvitationArn() terra.StringValue {
	return terra.ReferenceAsString(rrsa.ref.Append("invitation_arn"))
}

// ReceiverAccountId returns a reference to field receiver_account_id of aws_ram_resource_share_accepter.
func (rrsa ramResourceShareAccepterAttributes) ReceiverAccountId() terra.StringValue {
	return terra.ReferenceAsString(rrsa.ref.Append("receiver_account_id"))
}

// Resources returns a reference to field resources of aws_ram_resource_share_accepter.
func (rrsa ramResourceShareAccepterAttributes) Resources() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rrsa.ref.Append("resources"))
}

// SenderAccountId returns a reference to field sender_account_id of aws_ram_resource_share_accepter.
func (rrsa ramResourceShareAccepterAttributes) SenderAccountId() terra.StringValue {
	return terra.ReferenceAsString(rrsa.ref.Append("sender_account_id"))
}

// ShareArn returns a reference to field share_arn of aws_ram_resource_share_accepter.
func (rrsa ramResourceShareAccepterAttributes) ShareArn() terra.StringValue {
	return terra.ReferenceAsString(rrsa.ref.Append("share_arn"))
}

// ShareId returns a reference to field share_id of aws_ram_resource_share_accepter.
func (rrsa ramResourceShareAccepterAttributes) ShareId() terra.StringValue {
	return terra.ReferenceAsString(rrsa.ref.Append("share_id"))
}

// ShareName returns a reference to field share_name of aws_ram_resource_share_accepter.
func (rrsa ramResourceShareAccepterAttributes) ShareName() terra.StringValue {
	return terra.ReferenceAsString(rrsa.ref.Append("share_name"))
}

// Status returns a reference to field status of aws_ram_resource_share_accepter.
func (rrsa ramResourceShareAccepterAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(rrsa.ref.Append("status"))
}

func (rrsa ramResourceShareAccepterAttributes) Timeouts() ramresourceshareaccepter.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ramresourceshareaccepter.TimeoutsAttributes](rrsa.ref.Append("timeouts"))
}

type ramResourceShareAccepterState struct {
	Id                string                                  `json:"id"`
	InvitationArn     string                                  `json:"invitation_arn"`
	ReceiverAccountId string                                  `json:"receiver_account_id"`
	Resources         []string                                `json:"resources"`
	SenderAccountId   string                                  `json:"sender_account_id"`
	ShareArn          string                                  `json:"share_arn"`
	ShareId           string                                  `json:"share_id"`
	ShareName         string                                  `json:"share_name"`
	Status            string                                  `json:"status"`
	Timeouts          *ramresourceshareaccepter.TimeoutsState `json:"timeouts"`
}
