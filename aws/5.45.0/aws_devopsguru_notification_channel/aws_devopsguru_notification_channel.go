// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_devopsguru_notification_channel

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

// Resource represents the Terraform resource aws_devopsguru_notification_channel.
type Resource struct {
	Name      string
	Args      Args
	state     *awsDevopsguruNotificationChannelState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adnc *Resource) Type() string {
	return "aws_devopsguru_notification_channel"
}

// LocalName returns the local name for [Resource].
func (adnc *Resource) LocalName() string {
	return adnc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adnc *Resource) Configuration() interface{} {
	return adnc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adnc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adnc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adnc *Resource) Dependencies() terra.Dependencies {
	return adnc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adnc *Resource) LifecycleManagement() *terra.Lifecycle {
	return adnc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adnc *Resource) Attributes() awsDevopsguruNotificationChannelAttributes {
	return awsDevopsguruNotificationChannelAttributes{ref: terra.ReferenceResource(adnc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adnc *Resource) ImportState(state io.Reader) error {
	adnc.state = &awsDevopsguruNotificationChannelState{}
	if err := json.NewDecoder(state).Decode(adnc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adnc.Type(), adnc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adnc *Resource) State() (*awsDevopsguruNotificationChannelState, bool) {
	return adnc.state, adnc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adnc *Resource) StateMust() *awsDevopsguruNotificationChannelState {
	if adnc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adnc.Type(), adnc.LocalName()))
	}
	return adnc.state
}

// Args contains the configurations for aws_devopsguru_notification_channel.
type Args struct {
	// Filters: min=0
	Filters []Filters `hcl:"filters,block" validate:"min=0"`
	// Sns: min=0
	Sns []Sns `hcl:"sns,block" validate:"min=0"`
}

type awsDevopsguruNotificationChannelAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_devopsguru_notification_channel.
func (adnc awsDevopsguruNotificationChannelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adnc.ref.Append("id"))
}

func (adnc awsDevopsguruNotificationChannelAttributes) Filters() terra.ListValue[FiltersAttributes] {
	return terra.ReferenceAsList[FiltersAttributes](adnc.ref.Append("filters"))
}

func (adnc awsDevopsguruNotificationChannelAttributes) Sns() terra.ListValue[SnsAttributes] {
	return terra.ReferenceAsList[SnsAttributes](adnc.ref.Append("sns"))
}

type awsDevopsguruNotificationChannelState struct {
	Id      string         `json:"id"`
	Filters []FiltersState `json:"filters"`
	Sns     []SnsState     `json:"sns"`
}
