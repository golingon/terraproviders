// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mediaassetfilter "github.com/golingon/terraproviders/azurerm/3.52.0/mediaassetfilter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMediaAssetFilter creates a new instance of [MediaAssetFilter].
func NewMediaAssetFilter(name string, args MediaAssetFilterArgs) *MediaAssetFilter {
	return &MediaAssetFilter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MediaAssetFilter)(nil)

// MediaAssetFilter represents the Terraform resource azurerm_media_asset_filter.
type MediaAssetFilter struct {
	Name      string
	Args      MediaAssetFilterArgs
	state     *mediaAssetFilterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MediaAssetFilter].
func (maf *MediaAssetFilter) Type() string {
	return "azurerm_media_asset_filter"
}

// LocalName returns the local name for [MediaAssetFilter].
func (maf *MediaAssetFilter) LocalName() string {
	return maf.Name
}

// Configuration returns the configuration (args) for [MediaAssetFilter].
func (maf *MediaAssetFilter) Configuration() interface{} {
	return maf.Args
}

// DependOn is used for other resources to depend on [MediaAssetFilter].
func (maf *MediaAssetFilter) DependOn() terra.Reference {
	return terra.ReferenceResource(maf)
}

// Dependencies returns the list of resources [MediaAssetFilter] depends_on.
func (maf *MediaAssetFilter) Dependencies() terra.Dependencies {
	return maf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MediaAssetFilter].
func (maf *MediaAssetFilter) LifecycleManagement() *terra.Lifecycle {
	return maf.Lifecycle
}

// Attributes returns the attributes for [MediaAssetFilter].
func (maf *MediaAssetFilter) Attributes() mediaAssetFilterAttributes {
	return mediaAssetFilterAttributes{ref: terra.ReferenceResource(maf)}
}

// ImportState imports the given attribute values into [MediaAssetFilter]'s state.
func (maf *MediaAssetFilter) ImportState(av io.Reader) error {
	maf.state = &mediaAssetFilterState{}
	if err := json.NewDecoder(av).Decode(maf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", maf.Type(), maf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MediaAssetFilter] has state.
func (maf *MediaAssetFilter) State() (*mediaAssetFilterState, bool) {
	return maf.state, maf.state != nil
}

// StateMust returns the state for [MediaAssetFilter]. Panics if the state is nil.
func (maf *MediaAssetFilter) StateMust() *mediaAssetFilterState {
	if maf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", maf.Type(), maf.LocalName()))
	}
	return maf.state
}

// MediaAssetFilterArgs contains the configurations for azurerm_media_asset_filter.
type MediaAssetFilterArgs struct {
	// AssetId: string, required
	AssetId terra.StringValue `hcl:"asset_id,attr" validate:"required"`
	// FirstQualityBitrate: number, optional
	FirstQualityBitrate terra.NumberValue `hcl:"first_quality_bitrate,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PresentationTimeRange: optional
	PresentationTimeRange *mediaassetfilter.PresentationTimeRange `hcl:"presentation_time_range,block"`
	// Timeouts: optional
	Timeouts *mediaassetfilter.Timeouts `hcl:"timeouts,block"`
	// TrackSelection: min=0
	TrackSelection []mediaassetfilter.TrackSelection `hcl:"track_selection,block" validate:"min=0"`
}
type mediaAssetFilterAttributes struct {
	ref terra.Reference
}

// AssetId returns a reference to field asset_id of azurerm_media_asset_filter.
func (maf mediaAssetFilterAttributes) AssetId() terra.StringValue {
	return terra.ReferenceAsString(maf.ref.Append("asset_id"))
}

// FirstQualityBitrate returns a reference to field first_quality_bitrate of azurerm_media_asset_filter.
func (maf mediaAssetFilterAttributes) FirstQualityBitrate() terra.NumberValue {
	return terra.ReferenceAsNumber(maf.ref.Append("first_quality_bitrate"))
}

// Id returns a reference to field id of azurerm_media_asset_filter.
func (maf mediaAssetFilterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(maf.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_media_asset_filter.
func (maf mediaAssetFilterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(maf.ref.Append("name"))
}

func (maf mediaAssetFilterAttributes) PresentationTimeRange() terra.ListValue[mediaassetfilter.PresentationTimeRangeAttributes] {
	return terra.ReferenceAsList[mediaassetfilter.PresentationTimeRangeAttributes](maf.ref.Append("presentation_time_range"))
}

func (maf mediaAssetFilterAttributes) Timeouts() mediaassetfilter.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mediaassetfilter.TimeoutsAttributes](maf.ref.Append("timeouts"))
}

func (maf mediaAssetFilterAttributes) TrackSelection() terra.ListValue[mediaassetfilter.TrackSelectionAttributes] {
	return terra.ReferenceAsList[mediaassetfilter.TrackSelectionAttributes](maf.ref.Append("track_selection"))
}

type mediaAssetFilterState struct {
	AssetId               string                                        `json:"asset_id"`
	FirstQualityBitrate   float64                                       `json:"first_quality_bitrate"`
	Id                    string                                        `json:"id"`
	Name                  string                                        `json:"name"`
	PresentationTimeRange []mediaassetfilter.PresentationTimeRangeState `json:"presentation_time_range"`
	Timeouts              *mediaassetfilter.TimeoutsState               `json:"timeouts"`
	TrackSelection        []mediaassetfilter.TrackSelectionState        `json:"track_selection"`
}
