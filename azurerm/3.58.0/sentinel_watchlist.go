// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentinelwatchlist "github.com/golingon/terraproviders/azurerm/3.58.0/sentinelwatchlist"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelWatchlist creates a new instance of [SentinelWatchlist].
func NewSentinelWatchlist(name string, args SentinelWatchlistArgs) *SentinelWatchlist {
	return &SentinelWatchlist{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelWatchlist)(nil)

// SentinelWatchlist represents the Terraform resource azurerm_sentinel_watchlist.
type SentinelWatchlist struct {
	Name      string
	Args      SentinelWatchlistArgs
	state     *sentinelWatchlistState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelWatchlist].
func (sw *SentinelWatchlist) Type() string {
	return "azurerm_sentinel_watchlist"
}

// LocalName returns the local name for [SentinelWatchlist].
func (sw *SentinelWatchlist) LocalName() string {
	return sw.Name
}

// Configuration returns the configuration (args) for [SentinelWatchlist].
func (sw *SentinelWatchlist) Configuration() interface{} {
	return sw.Args
}

// DependOn is used for other resources to depend on [SentinelWatchlist].
func (sw *SentinelWatchlist) DependOn() terra.Reference {
	return terra.ReferenceResource(sw)
}

// Dependencies returns the list of resources [SentinelWatchlist] depends_on.
func (sw *SentinelWatchlist) Dependencies() terra.Dependencies {
	return sw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelWatchlist].
func (sw *SentinelWatchlist) LifecycleManagement() *terra.Lifecycle {
	return sw.Lifecycle
}

// Attributes returns the attributes for [SentinelWatchlist].
func (sw *SentinelWatchlist) Attributes() sentinelWatchlistAttributes {
	return sentinelWatchlistAttributes{ref: terra.ReferenceResource(sw)}
}

// ImportState imports the given attribute values into [SentinelWatchlist]'s state.
func (sw *SentinelWatchlist) ImportState(av io.Reader) error {
	sw.state = &sentinelWatchlistState{}
	if err := json.NewDecoder(av).Decode(sw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sw.Type(), sw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelWatchlist] has state.
func (sw *SentinelWatchlist) State() (*sentinelWatchlistState, bool) {
	return sw.state, sw.state != nil
}

// StateMust returns the state for [SentinelWatchlist]. Panics if the state is nil.
func (sw *SentinelWatchlist) StateMust() *sentinelWatchlistState {
	if sw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sw.Type(), sw.LocalName()))
	}
	return sw.state
}

// SentinelWatchlistArgs contains the configurations for azurerm_sentinel_watchlist.
type SentinelWatchlistArgs struct {
	// DefaultDuration: string, optional
	DefaultDuration terra.StringValue `hcl:"default_duration,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ItemSearchKey: string, required
	ItemSearchKey terra.StringValue `hcl:"item_search_key,attr" validate:"required"`
	// Labels: list of string, optional
	Labels terra.ListValue[terra.StringValue] `hcl:"labels,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *sentinelwatchlist.Timeouts `hcl:"timeouts,block"`
}
type sentinelWatchlistAttributes struct {
	ref terra.Reference
}

// DefaultDuration returns a reference to field default_duration of azurerm_sentinel_watchlist.
func (sw sentinelWatchlistAttributes) DefaultDuration() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("default_duration"))
}

// Description returns a reference to field description of azurerm_sentinel_watchlist.
func (sw sentinelWatchlistAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_sentinel_watchlist.
func (sw sentinelWatchlistAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_sentinel_watchlist.
func (sw sentinelWatchlistAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("id"))
}

// ItemSearchKey returns a reference to field item_search_key of azurerm_sentinel_watchlist.
func (sw sentinelWatchlistAttributes) ItemSearchKey() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("item_search_key"))
}

// Labels returns a reference to field labels of azurerm_sentinel_watchlist.
func (sw sentinelWatchlistAttributes) Labels() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sw.ref.Append("labels"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_watchlist.
func (sw sentinelWatchlistAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_watchlist.
func (sw sentinelWatchlistAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("name"))
}

func (sw sentinelWatchlistAttributes) Timeouts() sentinelwatchlist.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentinelwatchlist.TimeoutsAttributes](sw.ref.Append("timeouts"))
}

type sentinelWatchlistState struct {
	DefaultDuration         string                           `json:"default_duration"`
	Description             string                           `json:"description"`
	DisplayName             string                           `json:"display_name"`
	Id                      string                           `json:"id"`
	ItemSearchKey           string                           `json:"item_search_key"`
	Labels                  []string                         `json:"labels"`
	LogAnalyticsWorkspaceId string                           `json:"log_analytics_workspace_id"`
	Name                    string                           `json:"name"`
	Timeouts                *sentinelwatchlist.TimeoutsState `json:"timeouts"`
}
