// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ram_resource_share_accepter

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

// Resource represents the Terraform resource aws_ram_resource_share_accepter.
type Resource struct {
	Name      string
	Args      Args
	state     *awsRamResourceShareAccepterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (arrsa *Resource) Type() string {
	return "aws_ram_resource_share_accepter"
}

// LocalName returns the local name for [Resource].
func (arrsa *Resource) LocalName() string {
	return arrsa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (arrsa *Resource) Configuration() interface{} {
	return arrsa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (arrsa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(arrsa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (arrsa *Resource) Dependencies() terra.Dependencies {
	return arrsa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (arrsa *Resource) LifecycleManagement() *terra.Lifecycle {
	return arrsa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (arrsa *Resource) Attributes() awsRamResourceShareAccepterAttributes {
	return awsRamResourceShareAccepterAttributes{ref: terra.ReferenceResource(arrsa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (arrsa *Resource) ImportState(state io.Reader) error {
	arrsa.state = &awsRamResourceShareAccepterState{}
	if err := json.NewDecoder(state).Decode(arrsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", arrsa.Type(), arrsa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (arrsa *Resource) State() (*awsRamResourceShareAccepterState, bool) {
	return arrsa.state, arrsa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (arrsa *Resource) StateMust() *awsRamResourceShareAccepterState {
	if arrsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", arrsa.Type(), arrsa.LocalName()))
	}
	return arrsa.state
}

// Args contains the configurations for aws_ram_resource_share_accepter.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ShareArn: string, required
	ShareArn terra.StringValue `hcl:"share_arn,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsRamResourceShareAccepterAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ram_resource_share_accepter.
func (arrsa awsRamResourceShareAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(arrsa.ref.Append("id"))
}

// InvitationArn returns a reference to field invitation_arn of aws_ram_resource_share_accepter.
func (arrsa awsRamResourceShareAccepterAttributes) InvitationArn() terra.StringValue {
	return terra.ReferenceAsString(arrsa.ref.Append("invitation_arn"))
}

// ReceiverAccountId returns a reference to field receiver_account_id of aws_ram_resource_share_accepter.
func (arrsa awsRamResourceShareAccepterAttributes) ReceiverAccountId() terra.StringValue {
	return terra.ReferenceAsString(arrsa.ref.Append("receiver_account_id"))
}

// Resources returns a reference to field resources of aws_ram_resource_share_accepter.
func (arrsa awsRamResourceShareAccepterAttributes) Resources() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](arrsa.ref.Append("resources"))
}

// SenderAccountId returns a reference to field sender_account_id of aws_ram_resource_share_accepter.
func (arrsa awsRamResourceShareAccepterAttributes) SenderAccountId() terra.StringValue {
	return terra.ReferenceAsString(arrsa.ref.Append("sender_account_id"))
}

// ShareArn returns a reference to field share_arn of aws_ram_resource_share_accepter.
func (arrsa awsRamResourceShareAccepterAttributes) ShareArn() terra.StringValue {
	return terra.ReferenceAsString(arrsa.ref.Append("share_arn"))
}

// ShareId returns a reference to field share_id of aws_ram_resource_share_accepter.
func (arrsa awsRamResourceShareAccepterAttributes) ShareId() terra.StringValue {
	return terra.ReferenceAsString(arrsa.ref.Append("share_id"))
}

// ShareName returns a reference to field share_name of aws_ram_resource_share_accepter.
func (arrsa awsRamResourceShareAccepterAttributes) ShareName() terra.StringValue {
	return terra.ReferenceAsString(arrsa.ref.Append("share_name"))
}

// Status returns a reference to field status of aws_ram_resource_share_accepter.
func (arrsa awsRamResourceShareAccepterAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(arrsa.ref.Append("status"))
}

func (arrsa awsRamResourceShareAccepterAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](arrsa.ref.Append("timeouts"))
}

type awsRamResourceShareAccepterState struct {
	Id                string         `json:"id"`
	InvitationArn     string         `json:"invitation_arn"`
	ReceiverAccountId string         `json:"receiver_account_id"`
	Resources         []string       `json:"resources"`
	SenderAccountId   string         `json:"sender_account_id"`
	ShareArn          string         `json:"share_arn"`
	ShareId           string         `json:"share_id"`
	ShareName         string         `json:"share_name"`
	Status            string         `json:"status"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
