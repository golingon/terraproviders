// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	eventarcchannel "github.com/golingon/terraproviders/googlebeta/4.71.0/eventarcchannel"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEventarcChannel creates a new instance of [EventarcChannel].
func NewEventarcChannel(name string, args EventarcChannelArgs) *EventarcChannel {
	return &EventarcChannel{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EventarcChannel)(nil)

// EventarcChannel represents the Terraform resource google_eventarc_channel.
type EventarcChannel struct {
	Name      string
	Args      EventarcChannelArgs
	state     *eventarcChannelState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EventarcChannel].
func (ec *EventarcChannel) Type() string {
	return "google_eventarc_channel"
}

// LocalName returns the local name for [EventarcChannel].
func (ec *EventarcChannel) LocalName() string {
	return ec.Name
}

// Configuration returns the configuration (args) for [EventarcChannel].
func (ec *EventarcChannel) Configuration() interface{} {
	return ec.Args
}

// DependOn is used for other resources to depend on [EventarcChannel].
func (ec *EventarcChannel) DependOn() terra.Reference {
	return terra.ReferenceResource(ec)
}

// Dependencies returns the list of resources [EventarcChannel] depends_on.
func (ec *EventarcChannel) Dependencies() terra.Dependencies {
	return ec.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EventarcChannel].
func (ec *EventarcChannel) LifecycleManagement() *terra.Lifecycle {
	return ec.Lifecycle
}

// Attributes returns the attributes for [EventarcChannel].
func (ec *EventarcChannel) Attributes() eventarcChannelAttributes {
	return eventarcChannelAttributes{ref: terra.ReferenceResource(ec)}
}

// ImportState imports the given attribute values into [EventarcChannel]'s state.
func (ec *EventarcChannel) ImportState(av io.Reader) error {
	ec.state = &eventarcChannelState{}
	if err := json.NewDecoder(av).Decode(ec.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ec.Type(), ec.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EventarcChannel] has state.
func (ec *EventarcChannel) State() (*eventarcChannelState, bool) {
	return ec.state, ec.state != nil
}

// StateMust returns the state for [EventarcChannel]. Panics if the state is nil.
func (ec *EventarcChannel) StateMust() *eventarcChannelState {
	if ec.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ec.Type(), ec.LocalName()))
	}
	return ec.state
}

// EventarcChannelArgs contains the configurations for google_eventarc_channel.
type EventarcChannelArgs struct {
	// CryptoKeyName: string, optional
	CryptoKeyName terra.StringValue `hcl:"crypto_key_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ThirdPartyProvider: string, optional
	ThirdPartyProvider terra.StringValue `hcl:"third_party_provider,attr"`
	// Timeouts: optional
	Timeouts *eventarcchannel.Timeouts `hcl:"timeouts,block"`
}
type eventarcChannelAttributes struct {
	ref terra.Reference
}

// ActivationToken returns a reference to field activation_token of google_eventarc_channel.
func (ec eventarcChannelAttributes) ActivationToken() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("activation_token"))
}

// CreateTime returns a reference to field create_time of google_eventarc_channel.
func (ec eventarcChannelAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("create_time"))
}

// CryptoKeyName returns a reference to field crypto_key_name of google_eventarc_channel.
func (ec eventarcChannelAttributes) CryptoKeyName() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("crypto_key_name"))
}

// Id returns a reference to field id of google_eventarc_channel.
func (ec eventarcChannelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("id"))
}

// Location returns a reference to field location of google_eventarc_channel.
func (ec eventarcChannelAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("location"))
}

// Name returns a reference to field name of google_eventarc_channel.
func (ec eventarcChannelAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("name"))
}

// Project returns a reference to field project of google_eventarc_channel.
func (ec eventarcChannelAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("project"))
}

// PubsubTopic returns a reference to field pubsub_topic of google_eventarc_channel.
func (ec eventarcChannelAttributes) PubsubTopic() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("pubsub_topic"))
}

// State returns a reference to field state of google_eventarc_channel.
func (ec eventarcChannelAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("state"))
}

// ThirdPartyProvider returns a reference to field third_party_provider of google_eventarc_channel.
func (ec eventarcChannelAttributes) ThirdPartyProvider() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("third_party_provider"))
}

// Uid returns a reference to field uid of google_eventarc_channel.
func (ec eventarcChannelAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_eventarc_channel.
func (ec eventarcChannelAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("update_time"))
}

func (ec eventarcChannelAttributes) Timeouts() eventarcchannel.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eventarcchannel.TimeoutsAttributes](ec.ref.Append("timeouts"))
}

type eventarcChannelState struct {
	ActivationToken    string                         `json:"activation_token"`
	CreateTime         string                         `json:"create_time"`
	CryptoKeyName      string                         `json:"crypto_key_name"`
	Id                 string                         `json:"id"`
	Location           string                         `json:"location"`
	Name               string                         `json:"name"`
	Project            string                         `json:"project"`
	PubsubTopic        string                         `json:"pubsub_topic"`
	State              string                         `json:"state"`
	ThirdPartyProvider string                         `json:"third_party_provider"`
	Uid                string                         `json:"uid"`
	UpdateTime         string                         `json:"update_time"`
	Timeouts           *eventarcchannel.TimeoutsState `json:"timeouts"`
}
