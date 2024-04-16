// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_log_analytics_query_pack_query

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

// Resource represents the Terraform resource azurerm_log_analytics_query_pack_query.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermLogAnalyticsQueryPackQueryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (alaqpq *Resource) Type() string {
	return "azurerm_log_analytics_query_pack_query"
}

// LocalName returns the local name for [Resource].
func (alaqpq *Resource) LocalName() string {
	return alaqpq.Name
}

// Configuration returns the configuration (args) for [Resource].
func (alaqpq *Resource) Configuration() interface{} {
	return alaqpq.Args
}

// DependOn is used for other resources to depend on [Resource].
func (alaqpq *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(alaqpq)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (alaqpq *Resource) Dependencies() terra.Dependencies {
	return alaqpq.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (alaqpq *Resource) LifecycleManagement() *terra.Lifecycle {
	return alaqpq.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (alaqpq *Resource) Attributes() azurermLogAnalyticsQueryPackQueryAttributes {
	return azurermLogAnalyticsQueryPackQueryAttributes{ref: terra.ReferenceResource(alaqpq)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (alaqpq *Resource) ImportState(state io.Reader) error {
	alaqpq.state = &azurermLogAnalyticsQueryPackQueryState{}
	if err := json.NewDecoder(state).Decode(alaqpq.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", alaqpq.Type(), alaqpq.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (alaqpq *Resource) State() (*azurermLogAnalyticsQueryPackQueryState, bool) {
	return alaqpq.state, alaqpq.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (alaqpq *Resource) StateMust() *azurermLogAnalyticsQueryPackQueryState {
	if alaqpq.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", alaqpq.Type(), alaqpq.LocalName()))
	}
	return alaqpq.state
}

// Args contains the configurations for azurerm_log_analytics_query_pack_query.
type Args struct {
	// AdditionalSettingsJson: string, optional
	AdditionalSettingsJson terra.StringValue `hcl:"additional_settings_json,attr"`
	// Body: string, required
	Body terra.StringValue `hcl:"body,attr" validate:"required"`
	// Categories: list of string, optional
	Categories terra.ListValue[terra.StringValue] `hcl:"categories,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// QueryPackId: string, required
	QueryPackId terra.StringValue `hcl:"query_pack_id,attr" validate:"required"`
	// ResourceTypes: list of string, optional
	ResourceTypes terra.ListValue[terra.StringValue] `hcl:"resource_types,attr"`
	// Solutions: list of string, optional
	Solutions terra.ListValue[terra.StringValue] `hcl:"solutions,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermLogAnalyticsQueryPackQueryAttributes struct {
	ref terra.Reference
}

// AdditionalSettingsJson returns a reference to field additional_settings_json of azurerm_log_analytics_query_pack_query.
func (alaqpq azurermLogAnalyticsQueryPackQueryAttributes) AdditionalSettingsJson() terra.StringValue {
	return terra.ReferenceAsString(alaqpq.ref.Append("additional_settings_json"))
}

// Body returns a reference to field body of azurerm_log_analytics_query_pack_query.
func (alaqpq azurermLogAnalyticsQueryPackQueryAttributes) Body() terra.StringValue {
	return terra.ReferenceAsString(alaqpq.ref.Append("body"))
}

// Categories returns a reference to field categories of azurerm_log_analytics_query_pack_query.
func (alaqpq azurermLogAnalyticsQueryPackQueryAttributes) Categories() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](alaqpq.ref.Append("categories"))
}

// Description returns a reference to field description of azurerm_log_analytics_query_pack_query.
func (alaqpq azurermLogAnalyticsQueryPackQueryAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(alaqpq.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_log_analytics_query_pack_query.
func (alaqpq azurermLogAnalyticsQueryPackQueryAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(alaqpq.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_log_analytics_query_pack_query.
func (alaqpq azurermLogAnalyticsQueryPackQueryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(alaqpq.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_log_analytics_query_pack_query.
func (alaqpq azurermLogAnalyticsQueryPackQueryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(alaqpq.ref.Append("name"))
}

// QueryPackId returns a reference to field query_pack_id of azurerm_log_analytics_query_pack_query.
func (alaqpq azurermLogAnalyticsQueryPackQueryAttributes) QueryPackId() terra.StringValue {
	return terra.ReferenceAsString(alaqpq.ref.Append("query_pack_id"))
}

// ResourceTypes returns a reference to field resource_types of azurerm_log_analytics_query_pack_query.
func (alaqpq azurermLogAnalyticsQueryPackQueryAttributes) ResourceTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](alaqpq.ref.Append("resource_types"))
}

// Solutions returns a reference to field solutions of azurerm_log_analytics_query_pack_query.
func (alaqpq azurermLogAnalyticsQueryPackQueryAttributes) Solutions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](alaqpq.ref.Append("solutions"))
}

// Tags returns a reference to field tags of azurerm_log_analytics_query_pack_query.
func (alaqpq azurermLogAnalyticsQueryPackQueryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](alaqpq.ref.Append("tags"))
}

func (alaqpq azurermLogAnalyticsQueryPackQueryAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](alaqpq.ref.Append("timeouts"))
}

type azurermLogAnalyticsQueryPackQueryState struct {
	AdditionalSettingsJson string            `json:"additional_settings_json"`
	Body                   string            `json:"body"`
	Categories             []string          `json:"categories"`
	Description            string            `json:"description"`
	DisplayName            string            `json:"display_name"`
	Id                     string            `json:"id"`
	Name                   string            `json:"name"`
	QueryPackId            string            `json:"query_pack_id"`
	ResourceTypes          []string          `json:"resource_types"`
	Solutions              []string          `json:"solutions"`
	Tags                   map[string]string `json:"tags"`
	Timeouts               *TimeoutsState    `json:"timeouts"`
}
