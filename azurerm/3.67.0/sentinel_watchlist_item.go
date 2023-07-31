// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentinelwatchlistitem "github.com/golingon/terraproviders/azurerm/3.67.0/sentinelwatchlistitem"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelWatchlistItem creates a new instance of [SentinelWatchlistItem].
func NewSentinelWatchlistItem(name string, args SentinelWatchlistItemArgs) *SentinelWatchlistItem {
	return &SentinelWatchlistItem{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelWatchlistItem)(nil)

// SentinelWatchlistItem represents the Terraform resource azurerm_sentinel_watchlist_item.
type SentinelWatchlistItem struct {
	Name      string
	Args      SentinelWatchlistItemArgs
	state     *sentinelWatchlistItemState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelWatchlistItem].
func (swi *SentinelWatchlistItem) Type() string {
	return "azurerm_sentinel_watchlist_item"
}

// LocalName returns the local name for [SentinelWatchlistItem].
func (swi *SentinelWatchlistItem) LocalName() string {
	return swi.Name
}

// Configuration returns the configuration (args) for [SentinelWatchlistItem].
func (swi *SentinelWatchlistItem) Configuration() interface{} {
	return swi.Args
}

// DependOn is used for other resources to depend on [SentinelWatchlistItem].
func (swi *SentinelWatchlistItem) DependOn() terra.Reference {
	return terra.ReferenceResource(swi)
}

// Dependencies returns the list of resources [SentinelWatchlistItem] depends_on.
func (swi *SentinelWatchlistItem) Dependencies() terra.Dependencies {
	return swi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelWatchlistItem].
func (swi *SentinelWatchlistItem) LifecycleManagement() *terra.Lifecycle {
	return swi.Lifecycle
}

// Attributes returns the attributes for [SentinelWatchlistItem].
func (swi *SentinelWatchlistItem) Attributes() sentinelWatchlistItemAttributes {
	return sentinelWatchlistItemAttributes{ref: terra.ReferenceResource(swi)}
}

// ImportState imports the given attribute values into [SentinelWatchlistItem]'s state.
func (swi *SentinelWatchlistItem) ImportState(av io.Reader) error {
	swi.state = &sentinelWatchlistItemState{}
	if err := json.NewDecoder(av).Decode(swi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", swi.Type(), swi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelWatchlistItem] has state.
func (swi *SentinelWatchlistItem) State() (*sentinelWatchlistItemState, bool) {
	return swi.state, swi.state != nil
}

// StateMust returns the state for [SentinelWatchlistItem]. Panics if the state is nil.
func (swi *SentinelWatchlistItem) StateMust() *sentinelWatchlistItemState {
	if swi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", swi.Type(), swi.LocalName()))
	}
	return swi.state
}

// SentinelWatchlistItemArgs contains the configurations for azurerm_sentinel_watchlist_item.
type SentinelWatchlistItemArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Properties: map of string, required
	Properties terra.MapValue[terra.StringValue] `hcl:"properties,attr" validate:"required"`
	// WatchlistId: string, required
	WatchlistId terra.StringValue `hcl:"watchlist_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *sentinelwatchlistitem.Timeouts `hcl:"timeouts,block"`
}
type sentinelWatchlistItemAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sentinel_watchlist_item.
func (swi sentinelWatchlistItemAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(swi.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_sentinel_watchlist_item.
func (swi sentinelWatchlistItemAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(swi.ref.Append("name"))
}

// Properties returns a reference to field properties of azurerm_sentinel_watchlist_item.
func (swi sentinelWatchlistItemAttributes) Properties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](swi.ref.Append("properties"))
}

// WatchlistId returns a reference to field watchlist_id of azurerm_sentinel_watchlist_item.
func (swi sentinelWatchlistItemAttributes) WatchlistId() terra.StringValue {
	return terra.ReferenceAsString(swi.ref.Append("watchlist_id"))
}

func (swi sentinelWatchlistItemAttributes) Timeouts() sentinelwatchlistitem.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentinelwatchlistitem.TimeoutsAttributes](swi.ref.Append("timeouts"))
}

type sentinelWatchlistItemState struct {
	Id          string                               `json:"id"`
	Name        string                               `json:"name"`
	Properties  map[string]string                    `json:"properties"`
	WatchlistId string                               `json:"watchlist_id"`
	Timeouts    *sentinelwatchlistitem.TimeoutsState `json:"timeouts"`
}
