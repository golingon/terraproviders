// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	eventarcgooglechannelconfig "github.com/golingon/terraproviders/google/4.71.0/eventarcgooglechannelconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEventarcGoogleChannelConfig creates a new instance of [EventarcGoogleChannelConfig].
func NewEventarcGoogleChannelConfig(name string, args EventarcGoogleChannelConfigArgs) *EventarcGoogleChannelConfig {
	return &EventarcGoogleChannelConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EventarcGoogleChannelConfig)(nil)

// EventarcGoogleChannelConfig represents the Terraform resource google_eventarc_google_channel_config.
type EventarcGoogleChannelConfig struct {
	Name      string
	Args      EventarcGoogleChannelConfigArgs
	state     *eventarcGoogleChannelConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EventarcGoogleChannelConfig].
func (egcc *EventarcGoogleChannelConfig) Type() string {
	return "google_eventarc_google_channel_config"
}

// LocalName returns the local name for [EventarcGoogleChannelConfig].
func (egcc *EventarcGoogleChannelConfig) LocalName() string {
	return egcc.Name
}

// Configuration returns the configuration (args) for [EventarcGoogleChannelConfig].
func (egcc *EventarcGoogleChannelConfig) Configuration() interface{} {
	return egcc.Args
}

// DependOn is used for other resources to depend on [EventarcGoogleChannelConfig].
func (egcc *EventarcGoogleChannelConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(egcc)
}

// Dependencies returns the list of resources [EventarcGoogleChannelConfig] depends_on.
func (egcc *EventarcGoogleChannelConfig) Dependencies() terra.Dependencies {
	return egcc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EventarcGoogleChannelConfig].
func (egcc *EventarcGoogleChannelConfig) LifecycleManagement() *terra.Lifecycle {
	return egcc.Lifecycle
}

// Attributes returns the attributes for [EventarcGoogleChannelConfig].
func (egcc *EventarcGoogleChannelConfig) Attributes() eventarcGoogleChannelConfigAttributes {
	return eventarcGoogleChannelConfigAttributes{ref: terra.ReferenceResource(egcc)}
}

// ImportState imports the given attribute values into [EventarcGoogleChannelConfig]'s state.
func (egcc *EventarcGoogleChannelConfig) ImportState(av io.Reader) error {
	egcc.state = &eventarcGoogleChannelConfigState{}
	if err := json.NewDecoder(av).Decode(egcc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", egcc.Type(), egcc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EventarcGoogleChannelConfig] has state.
func (egcc *EventarcGoogleChannelConfig) State() (*eventarcGoogleChannelConfigState, bool) {
	return egcc.state, egcc.state != nil
}

// StateMust returns the state for [EventarcGoogleChannelConfig]. Panics if the state is nil.
func (egcc *EventarcGoogleChannelConfig) StateMust() *eventarcGoogleChannelConfigState {
	if egcc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", egcc.Type(), egcc.LocalName()))
	}
	return egcc.state
}

// EventarcGoogleChannelConfigArgs contains the configurations for google_eventarc_google_channel_config.
type EventarcGoogleChannelConfigArgs struct {
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
	// Timeouts: optional
	Timeouts *eventarcgooglechannelconfig.Timeouts `hcl:"timeouts,block"`
}
type eventarcGoogleChannelConfigAttributes struct {
	ref terra.Reference
}

// CryptoKeyName returns a reference to field crypto_key_name of google_eventarc_google_channel_config.
func (egcc eventarcGoogleChannelConfigAttributes) CryptoKeyName() terra.StringValue {
	return terra.ReferenceAsString(egcc.ref.Append("crypto_key_name"))
}

// Id returns a reference to field id of google_eventarc_google_channel_config.
func (egcc eventarcGoogleChannelConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(egcc.ref.Append("id"))
}

// Location returns a reference to field location of google_eventarc_google_channel_config.
func (egcc eventarcGoogleChannelConfigAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(egcc.ref.Append("location"))
}

// Name returns a reference to field name of google_eventarc_google_channel_config.
func (egcc eventarcGoogleChannelConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(egcc.ref.Append("name"))
}

// Project returns a reference to field project of google_eventarc_google_channel_config.
func (egcc eventarcGoogleChannelConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(egcc.ref.Append("project"))
}

// UpdateTime returns a reference to field update_time of google_eventarc_google_channel_config.
func (egcc eventarcGoogleChannelConfigAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(egcc.ref.Append("update_time"))
}

func (egcc eventarcGoogleChannelConfigAttributes) Timeouts() eventarcgooglechannelconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eventarcgooglechannelconfig.TimeoutsAttributes](egcc.ref.Append("timeouts"))
}

type eventarcGoogleChannelConfigState struct {
	CryptoKeyName string                                     `json:"crypto_key_name"`
	Id            string                                     `json:"id"`
	Location      string                                     `json:"location"`
	Name          string                                     `json:"name"`
	Project       string                                     `json:"project"`
	UpdateTime    string                                     `json:"update_time"`
	Timeouts      *eventarcgooglechannelconfig.TimeoutsState `json:"timeouts"`
}
