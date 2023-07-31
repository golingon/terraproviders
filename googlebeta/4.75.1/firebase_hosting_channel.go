// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	firebasehostingchannel "github.com/golingon/terraproviders/googlebeta/4.75.1/firebasehostingchannel"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirebaseHostingChannel creates a new instance of [FirebaseHostingChannel].
func NewFirebaseHostingChannel(name string, args FirebaseHostingChannelArgs) *FirebaseHostingChannel {
	return &FirebaseHostingChannel{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirebaseHostingChannel)(nil)

// FirebaseHostingChannel represents the Terraform resource google_firebase_hosting_channel.
type FirebaseHostingChannel struct {
	Name      string
	Args      FirebaseHostingChannelArgs
	state     *firebaseHostingChannelState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirebaseHostingChannel].
func (fhc *FirebaseHostingChannel) Type() string {
	return "google_firebase_hosting_channel"
}

// LocalName returns the local name for [FirebaseHostingChannel].
func (fhc *FirebaseHostingChannel) LocalName() string {
	return fhc.Name
}

// Configuration returns the configuration (args) for [FirebaseHostingChannel].
func (fhc *FirebaseHostingChannel) Configuration() interface{} {
	return fhc.Args
}

// DependOn is used for other resources to depend on [FirebaseHostingChannel].
func (fhc *FirebaseHostingChannel) DependOn() terra.Reference {
	return terra.ReferenceResource(fhc)
}

// Dependencies returns the list of resources [FirebaseHostingChannel] depends_on.
func (fhc *FirebaseHostingChannel) Dependencies() terra.Dependencies {
	return fhc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirebaseHostingChannel].
func (fhc *FirebaseHostingChannel) LifecycleManagement() *terra.Lifecycle {
	return fhc.Lifecycle
}

// Attributes returns the attributes for [FirebaseHostingChannel].
func (fhc *FirebaseHostingChannel) Attributes() firebaseHostingChannelAttributes {
	return firebaseHostingChannelAttributes{ref: terra.ReferenceResource(fhc)}
}

// ImportState imports the given attribute values into [FirebaseHostingChannel]'s state.
func (fhc *FirebaseHostingChannel) ImportState(av io.Reader) error {
	fhc.state = &firebaseHostingChannelState{}
	if err := json.NewDecoder(av).Decode(fhc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fhc.Type(), fhc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirebaseHostingChannel] has state.
func (fhc *FirebaseHostingChannel) State() (*firebaseHostingChannelState, bool) {
	return fhc.state, fhc.state != nil
}

// StateMust returns the state for [FirebaseHostingChannel]. Panics if the state is nil.
func (fhc *FirebaseHostingChannel) StateMust() *firebaseHostingChannelState {
	if fhc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fhc.Type(), fhc.LocalName()))
	}
	return fhc.state
}

// FirebaseHostingChannelArgs contains the configurations for google_firebase_hosting_channel.
type FirebaseHostingChannelArgs struct {
	// ChannelId: string, required
	ChannelId terra.StringValue `hcl:"channel_id,attr" validate:"required"`
	// ExpireTime: string, optional
	ExpireTime terra.StringValue `hcl:"expire_time,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// RetainedReleaseCount: number, optional
	RetainedReleaseCount terra.NumberValue `hcl:"retained_release_count,attr"`
	// SiteId: string, required
	SiteId terra.StringValue `hcl:"site_id,attr" validate:"required"`
	// Ttl: string, optional
	Ttl terra.StringValue `hcl:"ttl,attr"`
	// Timeouts: optional
	Timeouts *firebasehostingchannel.Timeouts `hcl:"timeouts,block"`
}
type firebaseHostingChannelAttributes struct {
	ref terra.Reference
}

// ChannelId returns a reference to field channel_id of google_firebase_hosting_channel.
func (fhc firebaseHostingChannelAttributes) ChannelId() terra.StringValue {
	return terra.ReferenceAsString(fhc.ref.Append("channel_id"))
}

// ExpireTime returns a reference to field expire_time of google_firebase_hosting_channel.
func (fhc firebaseHostingChannelAttributes) ExpireTime() terra.StringValue {
	return terra.ReferenceAsString(fhc.ref.Append("expire_time"))
}

// Id returns a reference to field id of google_firebase_hosting_channel.
func (fhc firebaseHostingChannelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fhc.ref.Append("id"))
}

// Labels returns a reference to field labels of google_firebase_hosting_channel.
func (fhc firebaseHostingChannelAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fhc.ref.Append("labels"))
}

// Name returns a reference to field name of google_firebase_hosting_channel.
func (fhc firebaseHostingChannelAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fhc.ref.Append("name"))
}

// RetainedReleaseCount returns a reference to field retained_release_count of google_firebase_hosting_channel.
func (fhc firebaseHostingChannelAttributes) RetainedReleaseCount() terra.NumberValue {
	return terra.ReferenceAsNumber(fhc.ref.Append("retained_release_count"))
}

// SiteId returns a reference to field site_id of google_firebase_hosting_channel.
func (fhc firebaseHostingChannelAttributes) SiteId() terra.StringValue {
	return terra.ReferenceAsString(fhc.ref.Append("site_id"))
}

// Ttl returns a reference to field ttl of google_firebase_hosting_channel.
func (fhc firebaseHostingChannelAttributes) Ttl() terra.StringValue {
	return terra.ReferenceAsString(fhc.ref.Append("ttl"))
}

func (fhc firebaseHostingChannelAttributes) Timeouts() firebasehostingchannel.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firebasehostingchannel.TimeoutsAttributes](fhc.ref.Append("timeouts"))
}

type firebaseHostingChannelState struct {
	ChannelId            string                                `json:"channel_id"`
	ExpireTime           string                                `json:"expire_time"`
	Id                   string                                `json:"id"`
	Labels               map[string]string                     `json:"labels"`
	Name                 string                                `json:"name"`
	RetainedReleaseCount float64                               `json:"retained_release_count"`
	SiteId               string                                `json:"site_id"`
	Ttl                  string                                `json:"ttl"`
	Timeouts             *firebasehostingchannel.TimeoutsState `json:"timeouts"`
}
