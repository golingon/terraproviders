// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	applicationinsightsanalyticsitem "github.com/golingon/terraproviders/azurerm/3.58.0/applicationinsightsanalyticsitem"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApplicationInsightsAnalyticsItem creates a new instance of [ApplicationInsightsAnalyticsItem].
func NewApplicationInsightsAnalyticsItem(name string, args ApplicationInsightsAnalyticsItemArgs) *ApplicationInsightsAnalyticsItem {
	return &ApplicationInsightsAnalyticsItem{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApplicationInsightsAnalyticsItem)(nil)

// ApplicationInsightsAnalyticsItem represents the Terraform resource azurerm_application_insights_analytics_item.
type ApplicationInsightsAnalyticsItem struct {
	Name      string
	Args      ApplicationInsightsAnalyticsItemArgs
	state     *applicationInsightsAnalyticsItemState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApplicationInsightsAnalyticsItem].
func (aiai *ApplicationInsightsAnalyticsItem) Type() string {
	return "azurerm_application_insights_analytics_item"
}

// LocalName returns the local name for [ApplicationInsightsAnalyticsItem].
func (aiai *ApplicationInsightsAnalyticsItem) LocalName() string {
	return aiai.Name
}

// Configuration returns the configuration (args) for [ApplicationInsightsAnalyticsItem].
func (aiai *ApplicationInsightsAnalyticsItem) Configuration() interface{} {
	return aiai.Args
}

// DependOn is used for other resources to depend on [ApplicationInsightsAnalyticsItem].
func (aiai *ApplicationInsightsAnalyticsItem) DependOn() terra.Reference {
	return terra.ReferenceResource(aiai)
}

// Dependencies returns the list of resources [ApplicationInsightsAnalyticsItem] depends_on.
func (aiai *ApplicationInsightsAnalyticsItem) Dependencies() terra.Dependencies {
	return aiai.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApplicationInsightsAnalyticsItem].
func (aiai *ApplicationInsightsAnalyticsItem) LifecycleManagement() *terra.Lifecycle {
	return aiai.Lifecycle
}

// Attributes returns the attributes for [ApplicationInsightsAnalyticsItem].
func (aiai *ApplicationInsightsAnalyticsItem) Attributes() applicationInsightsAnalyticsItemAttributes {
	return applicationInsightsAnalyticsItemAttributes{ref: terra.ReferenceResource(aiai)}
}

// ImportState imports the given attribute values into [ApplicationInsightsAnalyticsItem]'s state.
func (aiai *ApplicationInsightsAnalyticsItem) ImportState(av io.Reader) error {
	aiai.state = &applicationInsightsAnalyticsItemState{}
	if err := json.NewDecoder(av).Decode(aiai.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aiai.Type(), aiai.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApplicationInsightsAnalyticsItem] has state.
func (aiai *ApplicationInsightsAnalyticsItem) State() (*applicationInsightsAnalyticsItemState, bool) {
	return aiai.state, aiai.state != nil
}

// StateMust returns the state for [ApplicationInsightsAnalyticsItem]. Panics if the state is nil.
func (aiai *ApplicationInsightsAnalyticsItem) StateMust() *applicationInsightsAnalyticsItemState {
	if aiai.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aiai.Type(), aiai.LocalName()))
	}
	return aiai.state
}

// ApplicationInsightsAnalyticsItemArgs contains the configurations for azurerm_application_insights_analytics_item.
type ApplicationInsightsAnalyticsItemArgs struct {
	// ApplicationInsightsId: string, required
	ApplicationInsightsId terra.StringValue `hcl:"application_insights_id,attr" validate:"required"`
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// FunctionAlias: string, optional
	FunctionAlias terra.StringValue `hcl:"function_alias,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *applicationinsightsanalyticsitem.Timeouts `hcl:"timeouts,block"`
}
type applicationInsightsAnalyticsItemAttributes struct {
	ref terra.Reference
}

// ApplicationInsightsId returns a reference to field application_insights_id of azurerm_application_insights_analytics_item.
func (aiai applicationInsightsAnalyticsItemAttributes) ApplicationInsightsId() terra.StringValue {
	return terra.ReferenceAsString(aiai.ref.Append("application_insights_id"))
}

// Content returns a reference to field content of azurerm_application_insights_analytics_item.
func (aiai applicationInsightsAnalyticsItemAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(aiai.ref.Append("content"))
}

// FunctionAlias returns a reference to field function_alias of azurerm_application_insights_analytics_item.
func (aiai applicationInsightsAnalyticsItemAttributes) FunctionAlias() terra.StringValue {
	return terra.ReferenceAsString(aiai.ref.Append("function_alias"))
}

// Id returns a reference to field id of azurerm_application_insights_analytics_item.
func (aiai applicationInsightsAnalyticsItemAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aiai.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_application_insights_analytics_item.
func (aiai applicationInsightsAnalyticsItemAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aiai.ref.Append("name"))
}

// Scope returns a reference to field scope of azurerm_application_insights_analytics_item.
func (aiai applicationInsightsAnalyticsItemAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(aiai.ref.Append("scope"))
}

// TimeCreated returns a reference to field time_created of azurerm_application_insights_analytics_item.
func (aiai applicationInsightsAnalyticsItemAttributes) TimeCreated() terra.StringValue {
	return terra.ReferenceAsString(aiai.ref.Append("time_created"))
}

// TimeModified returns a reference to field time_modified of azurerm_application_insights_analytics_item.
func (aiai applicationInsightsAnalyticsItemAttributes) TimeModified() terra.StringValue {
	return terra.ReferenceAsString(aiai.ref.Append("time_modified"))
}

// Type returns a reference to field type of azurerm_application_insights_analytics_item.
func (aiai applicationInsightsAnalyticsItemAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(aiai.ref.Append("type"))
}

// Version returns a reference to field version of azurerm_application_insights_analytics_item.
func (aiai applicationInsightsAnalyticsItemAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(aiai.ref.Append("version"))
}

func (aiai applicationInsightsAnalyticsItemAttributes) Timeouts() applicationinsightsanalyticsitem.TimeoutsAttributes {
	return terra.ReferenceAsSingle[applicationinsightsanalyticsitem.TimeoutsAttributes](aiai.ref.Append("timeouts"))
}

type applicationInsightsAnalyticsItemState struct {
	ApplicationInsightsId string                                          `json:"application_insights_id"`
	Content               string                                          `json:"content"`
	FunctionAlias         string                                          `json:"function_alias"`
	Id                    string                                          `json:"id"`
	Name                  string                                          `json:"name"`
	Scope                 string                                          `json:"scope"`
	TimeCreated           string                                          `json:"time_created"`
	TimeModified          string                                          `json:"time_modified"`
	Type                  string                                          `json:"type"`
	Version               string                                          `json:"version"`
	Timeouts              *applicationinsightsanalyticsitem.TimeoutsState `json:"timeouts"`
}
