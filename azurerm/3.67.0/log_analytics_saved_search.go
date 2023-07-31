// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	loganalyticssavedsearch "github.com/golingon/terraproviders/azurerm/3.67.0/loganalyticssavedsearch"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogAnalyticsSavedSearch creates a new instance of [LogAnalyticsSavedSearch].
func NewLogAnalyticsSavedSearch(name string, args LogAnalyticsSavedSearchArgs) *LogAnalyticsSavedSearch {
	return &LogAnalyticsSavedSearch{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogAnalyticsSavedSearch)(nil)

// LogAnalyticsSavedSearch represents the Terraform resource azurerm_log_analytics_saved_search.
type LogAnalyticsSavedSearch struct {
	Name      string
	Args      LogAnalyticsSavedSearchArgs
	state     *logAnalyticsSavedSearchState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogAnalyticsSavedSearch].
func (lass *LogAnalyticsSavedSearch) Type() string {
	return "azurerm_log_analytics_saved_search"
}

// LocalName returns the local name for [LogAnalyticsSavedSearch].
func (lass *LogAnalyticsSavedSearch) LocalName() string {
	return lass.Name
}

// Configuration returns the configuration (args) for [LogAnalyticsSavedSearch].
func (lass *LogAnalyticsSavedSearch) Configuration() interface{} {
	return lass.Args
}

// DependOn is used for other resources to depend on [LogAnalyticsSavedSearch].
func (lass *LogAnalyticsSavedSearch) DependOn() terra.Reference {
	return terra.ReferenceResource(lass)
}

// Dependencies returns the list of resources [LogAnalyticsSavedSearch] depends_on.
func (lass *LogAnalyticsSavedSearch) Dependencies() terra.Dependencies {
	return lass.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogAnalyticsSavedSearch].
func (lass *LogAnalyticsSavedSearch) LifecycleManagement() *terra.Lifecycle {
	return lass.Lifecycle
}

// Attributes returns the attributes for [LogAnalyticsSavedSearch].
func (lass *LogAnalyticsSavedSearch) Attributes() logAnalyticsSavedSearchAttributes {
	return logAnalyticsSavedSearchAttributes{ref: terra.ReferenceResource(lass)}
}

// ImportState imports the given attribute values into [LogAnalyticsSavedSearch]'s state.
func (lass *LogAnalyticsSavedSearch) ImportState(av io.Reader) error {
	lass.state = &logAnalyticsSavedSearchState{}
	if err := json.NewDecoder(av).Decode(lass.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lass.Type(), lass.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogAnalyticsSavedSearch] has state.
func (lass *LogAnalyticsSavedSearch) State() (*logAnalyticsSavedSearchState, bool) {
	return lass.state, lass.state != nil
}

// StateMust returns the state for [LogAnalyticsSavedSearch]. Panics if the state is nil.
func (lass *LogAnalyticsSavedSearch) StateMust() *logAnalyticsSavedSearchState {
	if lass.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lass.Type(), lass.LocalName()))
	}
	return lass.state
}

// LogAnalyticsSavedSearchArgs contains the configurations for azurerm_log_analytics_saved_search.
type LogAnalyticsSavedSearchArgs struct {
	// Category: string, required
	Category terra.StringValue `hcl:"category,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// FunctionAlias: string, optional
	FunctionAlias terra.StringValue `hcl:"function_alias,attr"`
	// FunctionParameters: set of string, optional
	FunctionParameters terra.SetValue[terra.StringValue] `hcl:"function_parameters,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Query: string, required
	Query terra.StringValue `hcl:"query,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *loganalyticssavedsearch.Timeouts `hcl:"timeouts,block"`
}
type logAnalyticsSavedSearchAttributes struct {
	ref terra.Reference
}

// Category returns a reference to field category of azurerm_log_analytics_saved_search.
func (lass logAnalyticsSavedSearchAttributes) Category() terra.StringValue {
	return terra.ReferenceAsString(lass.ref.Append("category"))
}

// DisplayName returns a reference to field display_name of azurerm_log_analytics_saved_search.
func (lass logAnalyticsSavedSearchAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(lass.ref.Append("display_name"))
}

// FunctionAlias returns a reference to field function_alias of azurerm_log_analytics_saved_search.
func (lass logAnalyticsSavedSearchAttributes) FunctionAlias() terra.StringValue {
	return terra.ReferenceAsString(lass.ref.Append("function_alias"))
}

// FunctionParameters returns a reference to field function_parameters of azurerm_log_analytics_saved_search.
func (lass logAnalyticsSavedSearchAttributes) FunctionParameters() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](lass.ref.Append("function_parameters"))
}

// Id returns a reference to field id of azurerm_log_analytics_saved_search.
func (lass logAnalyticsSavedSearchAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lass.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_log_analytics_saved_search.
func (lass logAnalyticsSavedSearchAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(lass.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_log_analytics_saved_search.
func (lass logAnalyticsSavedSearchAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lass.ref.Append("name"))
}

// Query returns a reference to field query of azurerm_log_analytics_saved_search.
func (lass logAnalyticsSavedSearchAttributes) Query() terra.StringValue {
	return terra.ReferenceAsString(lass.ref.Append("query"))
}

// Tags returns a reference to field tags of azurerm_log_analytics_saved_search.
func (lass logAnalyticsSavedSearchAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lass.ref.Append("tags"))
}

func (lass logAnalyticsSavedSearchAttributes) Timeouts() loganalyticssavedsearch.TimeoutsAttributes {
	return terra.ReferenceAsSingle[loganalyticssavedsearch.TimeoutsAttributes](lass.ref.Append("timeouts"))
}

type logAnalyticsSavedSearchState struct {
	Category                string                                 `json:"category"`
	DisplayName             string                                 `json:"display_name"`
	FunctionAlias           string                                 `json:"function_alias"`
	FunctionParameters      []string                               `json:"function_parameters"`
	Id                      string                                 `json:"id"`
	LogAnalyticsWorkspaceId string                                 `json:"log_analytics_workspace_id"`
	Name                    string                                 `json:"name"`
	Query                   string                                 `json:"query"`
	Tags                    map[string]string                      `json:"tags"`
	Timeouts                *loganalyticssavedsearch.TimeoutsState `json:"timeouts"`
}
