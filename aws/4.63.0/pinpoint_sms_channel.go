// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPinpointSmsChannel creates a new instance of [PinpointSmsChannel].
func NewPinpointSmsChannel(name string, args PinpointSmsChannelArgs) *PinpointSmsChannel {
	return &PinpointSmsChannel{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PinpointSmsChannel)(nil)

// PinpointSmsChannel represents the Terraform resource aws_pinpoint_sms_channel.
type PinpointSmsChannel struct {
	Name      string
	Args      PinpointSmsChannelArgs
	state     *pinpointSmsChannelState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PinpointSmsChannel].
func (psc *PinpointSmsChannel) Type() string {
	return "aws_pinpoint_sms_channel"
}

// LocalName returns the local name for [PinpointSmsChannel].
func (psc *PinpointSmsChannel) LocalName() string {
	return psc.Name
}

// Configuration returns the configuration (args) for [PinpointSmsChannel].
func (psc *PinpointSmsChannel) Configuration() interface{} {
	return psc.Args
}

// DependOn is used for other resources to depend on [PinpointSmsChannel].
func (psc *PinpointSmsChannel) DependOn() terra.Reference {
	return terra.ReferenceResource(psc)
}

// Dependencies returns the list of resources [PinpointSmsChannel] depends_on.
func (psc *PinpointSmsChannel) Dependencies() terra.Dependencies {
	return psc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PinpointSmsChannel].
func (psc *PinpointSmsChannel) LifecycleManagement() *terra.Lifecycle {
	return psc.Lifecycle
}

// Attributes returns the attributes for [PinpointSmsChannel].
func (psc *PinpointSmsChannel) Attributes() pinpointSmsChannelAttributes {
	return pinpointSmsChannelAttributes{ref: terra.ReferenceResource(psc)}
}

// ImportState imports the given attribute values into [PinpointSmsChannel]'s state.
func (psc *PinpointSmsChannel) ImportState(av io.Reader) error {
	psc.state = &pinpointSmsChannelState{}
	if err := json.NewDecoder(av).Decode(psc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", psc.Type(), psc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PinpointSmsChannel] has state.
func (psc *PinpointSmsChannel) State() (*pinpointSmsChannelState, bool) {
	return psc.state, psc.state != nil
}

// StateMust returns the state for [PinpointSmsChannel]. Panics if the state is nil.
func (psc *PinpointSmsChannel) StateMust() *pinpointSmsChannelState {
	if psc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", psc.Type(), psc.LocalName()))
	}
	return psc.state
}

// PinpointSmsChannelArgs contains the configurations for aws_pinpoint_sms_channel.
type PinpointSmsChannelArgs struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SenderId: string, optional
	SenderId terra.StringValue `hcl:"sender_id,attr"`
	// ShortCode: string, optional
	ShortCode terra.StringValue `hcl:"short_code,attr"`
}
type pinpointSmsChannelAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of aws_pinpoint_sms_channel.
func (psc pinpointSmsChannelAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(psc.ref.Append("application_id"))
}

// Enabled returns a reference to field enabled of aws_pinpoint_sms_channel.
func (psc pinpointSmsChannelAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(psc.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_pinpoint_sms_channel.
func (psc pinpointSmsChannelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(psc.ref.Append("id"))
}

// PromotionalMessagesPerSecond returns a reference to field promotional_messages_per_second of aws_pinpoint_sms_channel.
func (psc pinpointSmsChannelAttributes) PromotionalMessagesPerSecond() terra.NumberValue {
	return terra.ReferenceAsNumber(psc.ref.Append("promotional_messages_per_second"))
}

// SenderId returns a reference to field sender_id of aws_pinpoint_sms_channel.
func (psc pinpointSmsChannelAttributes) SenderId() terra.StringValue {
	return terra.ReferenceAsString(psc.ref.Append("sender_id"))
}

// ShortCode returns a reference to field short_code of aws_pinpoint_sms_channel.
func (psc pinpointSmsChannelAttributes) ShortCode() terra.StringValue {
	return terra.ReferenceAsString(psc.ref.Append("short_code"))
}

// TransactionalMessagesPerSecond returns a reference to field transactional_messages_per_second of aws_pinpoint_sms_channel.
func (psc pinpointSmsChannelAttributes) TransactionalMessagesPerSecond() terra.NumberValue {
	return terra.ReferenceAsNumber(psc.ref.Append("transactional_messages_per_second"))
}

type pinpointSmsChannelState struct {
	ApplicationId                  string  `json:"application_id"`
	Enabled                        bool    `json:"enabled"`
	Id                             string  `json:"id"`
	PromotionalMessagesPerSecond   float64 `json:"promotional_messages_per_second"`
	SenderId                       string  `json:"sender_id"`
	ShortCode                      string  `json:"short_code"`
	TransactionalMessagesPerSecond float64 `json:"transactional_messages_per_second"`
}
