// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	medialiveevent "github.com/golingon/terraproviders/azurerm/3.49.0/medialiveevent"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMediaLiveEvent creates a new instance of [MediaLiveEvent].
func NewMediaLiveEvent(name string, args MediaLiveEventArgs) *MediaLiveEvent {
	return &MediaLiveEvent{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MediaLiveEvent)(nil)

// MediaLiveEvent represents the Terraform resource azurerm_media_live_event.
type MediaLiveEvent struct {
	Name      string
	Args      MediaLiveEventArgs
	state     *mediaLiveEventState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MediaLiveEvent].
func (mle *MediaLiveEvent) Type() string {
	return "azurerm_media_live_event"
}

// LocalName returns the local name for [MediaLiveEvent].
func (mle *MediaLiveEvent) LocalName() string {
	return mle.Name
}

// Configuration returns the configuration (args) for [MediaLiveEvent].
func (mle *MediaLiveEvent) Configuration() interface{} {
	return mle.Args
}

// DependOn is used for other resources to depend on [MediaLiveEvent].
func (mle *MediaLiveEvent) DependOn() terra.Reference {
	return terra.ReferenceResource(mle)
}

// Dependencies returns the list of resources [MediaLiveEvent] depends_on.
func (mle *MediaLiveEvent) Dependencies() terra.Dependencies {
	return mle.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MediaLiveEvent].
func (mle *MediaLiveEvent) LifecycleManagement() *terra.Lifecycle {
	return mle.Lifecycle
}

// Attributes returns the attributes for [MediaLiveEvent].
func (mle *MediaLiveEvent) Attributes() mediaLiveEventAttributes {
	return mediaLiveEventAttributes{ref: terra.ReferenceResource(mle)}
}

// ImportState imports the given attribute values into [MediaLiveEvent]'s state.
func (mle *MediaLiveEvent) ImportState(av io.Reader) error {
	mle.state = &mediaLiveEventState{}
	if err := json.NewDecoder(av).Decode(mle.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mle.Type(), mle.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MediaLiveEvent] has state.
func (mle *MediaLiveEvent) State() (*mediaLiveEventState, bool) {
	return mle.state, mle.state != nil
}

// StateMust returns the state for [MediaLiveEvent]. Panics if the state is nil.
func (mle *MediaLiveEvent) StateMust() *mediaLiveEventState {
	if mle.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mle.Type(), mle.LocalName()))
	}
	return mle.state
}

// MediaLiveEventArgs contains the configurations for azurerm_media_live_event.
type MediaLiveEventArgs struct {
	// AutoStartEnabled: bool, optional
	AutoStartEnabled terra.BoolValue `hcl:"auto_start_enabled,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// HostnamePrefix: string, optional
	HostnamePrefix terra.StringValue `hcl:"hostname_prefix,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MediaServicesAccountName: string, required
	MediaServicesAccountName terra.StringValue `hcl:"media_services_account_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StreamOptions: list of string, optional
	StreamOptions terra.ListValue[terra.StringValue] `hcl:"stream_options,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TranscriptionLanguages: list of string, optional
	TranscriptionLanguages terra.ListValue[terra.StringValue] `hcl:"transcription_languages,attr"`
	// UseStaticHostname: bool, optional
	UseStaticHostname terra.BoolValue `hcl:"use_static_hostname,attr"`
	// CrossSiteAccessPolicy: optional
	CrossSiteAccessPolicy *medialiveevent.CrossSiteAccessPolicy `hcl:"cross_site_access_policy,block"`
	// Encoding: optional
	Encoding *medialiveevent.Encoding `hcl:"encoding,block"`
	// Input: required
	Input *medialiveevent.Input `hcl:"input,block" validate:"required"`
	// Preview: optional
	Preview *medialiveevent.Preview `hcl:"preview,block"`
	// Timeouts: optional
	Timeouts *medialiveevent.Timeouts `hcl:"timeouts,block"`
}
type mediaLiveEventAttributes struct {
	ref terra.Reference
}

// AutoStartEnabled returns a reference to field auto_start_enabled of azurerm_media_live_event.
func (mle mediaLiveEventAttributes) AutoStartEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mle.ref.Append("auto_start_enabled"))
}

// Description returns a reference to field description of azurerm_media_live_event.
func (mle mediaLiveEventAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mle.ref.Append("description"))
}

// HostnamePrefix returns a reference to field hostname_prefix of azurerm_media_live_event.
func (mle mediaLiveEventAttributes) HostnamePrefix() terra.StringValue {
	return terra.ReferenceAsString(mle.ref.Append("hostname_prefix"))
}

// Id returns a reference to field id of azurerm_media_live_event.
func (mle mediaLiveEventAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mle.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_media_live_event.
func (mle mediaLiveEventAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mle.ref.Append("location"))
}

// MediaServicesAccountName returns a reference to field media_services_account_name of azurerm_media_live_event.
func (mle mediaLiveEventAttributes) MediaServicesAccountName() terra.StringValue {
	return terra.ReferenceAsString(mle.ref.Append("media_services_account_name"))
}

// Name returns a reference to field name of azurerm_media_live_event.
func (mle mediaLiveEventAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mle.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_media_live_event.
func (mle mediaLiveEventAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mle.ref.Append("resource_group_name"))
}

// StreamOptions returns a reference to field stream_options of azurerm_media_live_event.
func (mle mediaLiveEventAttributes) StreamOptions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mle.ref.Append("stream_options"))
}

// Tags returns a reference to field tags of azurerm_media_live_event.
func (mle mediaLiveEventAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mle.ref.Append("tags"))
}

// TranscriptionLanguages returns a reference to field transcription_languages of azurerm_media_live_event.
func (mle mediaLiveEventAttributes) TranscriptionLanguages() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mle.ref.Append("transcription_languages"))
}

// UseStaticHostname returns a reference to field use_static_hostname of azurerm_media_live_event.
func (mle mediaLiveEventAttributes) UseStaticHostname() terra.BoolValue {
	return terra.ReferenceAsBool(mle.ref.Append("use_static_hostname"))
}

func (mle mediaLiveEventAttributes) CrossSiteAccessPolicy() terra.ListValue[medialiveevent.CrossSiteAccessPolicyAttributes] {
	return terra.ReferenceAsList[medialiveevent.CrossSiteAccessPolicyAttributes](mle.ref.Append("cross_site_access_policy"))
}

func (mle mediaLiveEventAttributes) Encoding() terra.ListValue[medialiveevent.EncodingAttributes] {
	return terra.ReferenceAsList[medialiveevent.EncodingAttributes](mle.ref.Append("encoding"))
}

func (mle mediaLiveEventAttributes) Input() terra.ListValue[medialiveevent.InputAttributes] {
	return terra.ReferenceAsList[medialiveevent.InputAttributes](mle.ref.Append("input"))
}

func (mle mediaLiveEventAttributes) Preview() terra.ListValue[medialiveevent.PreviewAttributes] {
	return terra.ReferenceAsList[medialiveevent.PreviewAttributes](mle.ref.Append("preview"))
}

func (mle mediaLiveEventAttributes) Timeouts() medialiveevent.TimeoutsAttributes {
	return terra.ReferenceAsSingle[medialiveevent.TimeoutsAttributes](mle.ref.Append("timeouts"))
}

type mediaLiveEventState struct {
	AutoStartEnabled         bool                                        `json:"auto_start_enabled"`
	Description              string                                      `json:"description"`
	HostnamePrefix           string                                      `json:"hostname_prefix"`
	Id                       string                                      `json:"id"`
	Location                 string                                      `json:"location"`
	MediaServicesAccountName string                                      `json:"media_services_account_name"`
	Name                     string                                      `json:"name"`
	ResourceGroupName        string                                      `json:"resource_group_name"`
	StreamOptions            []string                                    `json:"stream_options"`
	Tags                     map[string]string                           `json:"tags"`
	TranscriptionLanguages   []string                                    `json:"transcription_languages"`
	UseStaticHostname        bool                                        `json:"use_static_hostname"`
	CrossSiteAccessPolicy    []medialiveevent.CrossSiteAccessPolicyState `json:"cross_site_access_policy"`
	Encoding                 []medialiveevent.EncodingState              `json:"encoding"`
	Input                    []medialiveevent.InputState                 `json:"input"`
	Preview                  []medialiveevent.PreviewState               `json:"preview"`
	Timeouts                 *medialiveevent.TimeoutsState               `json:"timeouts"`
}
