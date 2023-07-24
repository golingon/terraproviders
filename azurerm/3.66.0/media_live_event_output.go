// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	medialiveeventoutput "github.com/golingon/terraproviders/azurerm/3.66.0/medialiveeventoutput"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMediaLiveEventOutput creates a new instance of [MediaLiveEventOutput].
func NewMediaLiveEventOutput(name string, args MediaLiveEventOutputArgs) *MediaLiveEventOutput {
	return &MediaLiveEventOutput{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MediaLiveEventOutput)(nil)

// MediaLiveEventOutput represents the Terraform resource azurerm_media_live_event_output.
type MediaLiveEventOutput struct {
	Name      string
	Args      MediaLiveEventOutputArgs
	state     *mediaLiveEventOutputState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MediaLiveEventOutput].
func (mleo *MediaLiveEventOutput) Type() string {
	return "azurerm_media_live_event_output"
}

// LocalName returns the local name for [MediaLiveEventOutput].
func (mleo *MediaLiveEventOutput) LocalName() string {
	return mleo.Name
}

// Configuration returns the configuration (args) for [MediaLiveEventOutput].
func (mleo *MediaLiveEventOutput) Configuration() interface{} {
	return mleo.Args
}

// DependOn is used for other resources to depend on [MediaLiveEventOutput].
func (mleo *MediaLiveEventOutput) DependOn() terra.Reference {
	return terra.ReferenceResource(mleo)
}

// Dependencies returns the list of resources [MediaLiveEventOutput] depends_on.
func (mleo *MediaLiveEventOutput) Dependencies() terra.Dependencies {
	return mleo.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MediaLiveEventOutput].
func (mleo *MediaLiveEventOutput) LifecycleManagement() *terra.Lifecycle {
	return mleo.Lifecycle
}

// Attributes returns the attributes for [MediaLiveEventOutput].
func (mleo *MediaLiveEventOutput) Attributes() mediaLiveEventOutputAttributes {
	return mediaLiveEventOutputAttributes{ref: terra.ReferenceResource(mleo)}
}

// ImportState imports the given attribute values into [MediaLiveEventOutput]'s state.
func (mleo *MediaLiveEventOutput) ImportState(av io.Reader) error {
	mleo.state = &mediaLiveEventOutputState{}
	if err := json.NewDecoder(av).Decode(mleo.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mleo.Type(), mleo.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MediaLiveEventOutput] has state.
func (mleo *MediaLiveEventOutput) State() (*mediaLiveEventOutputState, bool) {
	return mleo.state, mleo.state != nil
}

// StateMust returns the state for [MediaLiveEventOutput]. Panics if the state is nil.
func (mleo *MediaLiveEventOutput) StateMust() *mediaLiveEventOutputState {
	if mleo.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mleo.Type(), mleo.LocalName()))
	}
	return mleo.state
}

// MediaLiveEventOutputArgs contains the configurations for azurerm_media_live_event_output.
type MediaLiveEventOutputArgs struct {
	// ArchiveWindowDuration: string, required
	ArchiveWindowDuration terra.StringValue `hcl:"archive_window_duration,attr" validate:"required"`
	// AssetName: string, required
	AssetName terra.StringValue `hcl:"asset_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// HlsFragmentsPerTsSegment: number, optional
	HlsFragmentsPerTsSegment terra.NumberValue `hcl:"hls_fragments_per_ts_segment,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LiveEventId: string, required
	LiveEventId terra.StringValue `hcl:"live_event_id,attr" validate:"required"`
	// ManifestName: string, optional
	ManifestName terra.StringValue `hcl:"manifest_name,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OutputSnapTimeInSeconds: number, optional
	OutputSnapTimeInSeconds terra.NumberValue `hcl:"output_snap_time_in_seconds,attr"`
	// RewindWindowDuration: string, optional
	RewindWindowDuration terra.StringValue `hcl:"rewind_window_duration,attr"`
	// Timeouts: optional
	Timeouts *medialiveeventoutput.Timeouts `hcl:"timeouts,block"`
}
type mediaLiveEventOutputAttributes struct {
	ref terra.Reference
}

// ArchiveWindowDuration returns a reference to field archive_window_duration of azurerm_media_live_event_output.
func (mleo mediaLiveEventOutputAttributes) ArchiveWindowDuration() terra.StringValue {
	return terra.ReferenceAsString(mleo.ref.Append("archive_window_duration"))
}

// AssetName returns a reference to field asset_name of azurerm_media_live_event_output.
func (mleo mediaLiveEventOutputAttributes) AssetName() terra.StringValue {
	return terra.ReferenceAsString(mleo.ref.Append("asset_name"))
}

// Description returns a reference to field description of azurerm_media_live_event_output.
func (mleo mediaLiveEventOutputAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mleo.ref.Append("description"))
}

// HlsFragmentsPerTsSegment returns a reference to field hls_fragments_per_ts_segment of azurerm_media_live_event_output.
func (mleo mediaLiveEventOutputAttributes) HlsFragmentsPerTsSegment() terra.NumberValue {
	return terra.ReferenceAsNumber(mleo.ref.Append("hls_fragments_per_ts_segment"))
}

// Id returns a reference to field id of azurerm_media_live_event_output.
func (mleo mediaLiveEventOutputAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mleo.ref.Append("id"))
}

// LiveEventId returns a reference to field live_event_id of azurerm_media_live_event_output.
func (mleo mediaLiveEventOutputAttributes) LiveEventId() terra.StringValue {
	return terra.ReferenceAsString(mleo.ref.Append("live_event_id"))
}

// ManifestName returns a reference to field manifest_name of azurerm_media_live_event_output.
func (mleo mediaLiveEventOutputAttributes) ManifestName() terra.StringValue {
	return terra.ReferenceAsString(mleo.ref.Append("manifest_name"))
}

// Name returns a reference to field name of azurerm_media_live_event_output.
func (mleo mediaLiveEventOutputAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mleo.ref.Append("name"))
}

// OutputSnapTimeInSeconds returns a reference to field output_snap_time_in_seconds of azurerm_media_live_event_output.
func (mleo mediaLiveEventOutputAttributes) OutputSnapTimeInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(mleo.ref.Append("output_snap_time_in_seconds"))
}

// RewindWindowDuration returns a reference to field rewind_window_duration of azurerm_media_live_event_output.
func (mleo mediaLiveEventOutputAttributes) RewindWindowDuration() terra.StringValue {
	return terra.ReferenceAsString(mleo.ref.Append("rewind_window_duration"))
}

func (mleo mediaLiveEventOutputAttributes) Timeouts() medialiveeventoutput.TimeoutsAttributes {
	return terra.ReferenceAsSingle[medialiveeventoutput.TimeoutsAttributes](mleo.ref.Append("timeouts"))
}

type mediaLiveEventOutputState struct {
	ArchiveWindowDuration    string                              `json:"archive_window_duration"`
	AssetName                string                              `json:"asset_name"`
	Description              string                              `json:"description"`
	HlsFragmentsPerTsSegment float64                             `json:"hls_fragments_per_ts_segment"`
	Id                       string                              `json:"id"`
	LiveEventId              string                              `json:"live_event_id"`
	ManifestName             string                              `json:"manifest_name"`
	Name                     string                              `json:"name"`
	OutputSnapTimeInSeconds  float64                             `json:"output_snap_time_in_seconds"`
	RewindWindowDuration     string                              `json:"rewind_window_duration"`
	Timeouts                 *medialiveeventoutput.TimeoutsState `json:"timeouts"`
}
