// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_application_insights_workbook_template

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

// Resource represents the Terraform resource azurerm_application_insights_workbook_template.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermApplicationInsightsWorkbookTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aaiwt *Resource) Type() string {
	return "azurerm_application_insights_workbook_template"
}

// LocalName returns the local name for [Resource].
func (aaiwt *Resource) LocalName() string {
	return aaiwt.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aaiwt *Resource) Configuration() interface{} {
	return aaiwt.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aaiwt *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aaiwt)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aaiwt *Resource) Dependencies() terra.Dependencies {
	return aaiwt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aaiwt *Resource) LifecycleManagement() *terra.Lifecycle {
	return aaiwt.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aaiwt *Resource) Attributes() azurermApplicationInsightsWorkbookTemplateAttributes {
	return azurermApplicationInsightsWorkbookTemplateAttributes{ref: terra.ReferenceResource(aaiwt)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aaiwt *Resource) ImportState(state io.Reader) error {
	aaiwt.state = &azurermApplicationInsightsWorkbookTemplateState{}
	if err := json.NewDecoder(state).Decode(aaiwt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aaiwt.Type(), aaiwt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aaiwt *Resource) State() (*azurermApplicationInsightsWorkbookTemplateState, bool) {
	return aaiwt.state, aaiwt.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aaiwt *Resource) StateMust() *azurermApplicationInsightsWorkbookTemplateState {
	if aaiwt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aaiwt.Type(), aaiwt.LocalName()))
	}
	return aaiwt.state
}

// Args contains the configurations for azurerm_application_insights_workbook_template.
type Args struct {
	// Author: string, optional
	Author terra.StringValue `hcl:"author,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Localized: string, optional
	Localized terra.StringValue `hcl:"localized,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TemplateData: string, required
	TemplateData terra.StringValue `hcl:"template_data,attr" validate:"required"`
	// Galleries: min=1
	Galleries []Galleries `hcl:"galleries,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermApplicationInsightsWorkbookTemplateAttributes struct {
	ref terra.Reference
}

// Author returns a reference to field author of azurerm_application_insights_workbook_template.
func (aaiwt azurermApplicationInsightsWorkbookTemplateAttributes) Author() terra.StringValue {
	return terra.ReferenceAsString(aaiwt.ref.Append("author"))
}

// Id returns a reference to field id of azurerm_application_insights_workbook_template.
func (aaiwt azurermApplicationInsightsWorkbookTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aaiwt.ref.Append("id"))
}

// Localized returns a reference to field localized of azurerm_application_insights_workbook_template.
func (aaiwt azurermApplicationInsightsWorkbookTemplateAttributes) Localized() terra.StringValue {
	return terra.ReferenceAsString(aaiwt.ref.Append("localized"))
}

// Location returns a reference to field location of azurerm_application_insights_workbook_template.
func (aaiwt azurermApplicationInsightsWorkbookTemplateAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aaiwt.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_application_insights_workbook_template.
func (aaiwt azurermApplicationInsightsWorkbookTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aaiwt.ref.Append("name"))
}

// Priority returns a reference to field priority of azurerm_application_insights_workbook_template.
func (aaiwt azurermApplicationInsightsWorkbookTemplateAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(aaiwt.ref.Append("priority"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_application_insights_workbook_template.
func (aaiwt azurermApplicationInsightsWorkbookTemplateAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aaiwt.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_application_insights_workbook_template.
func (aaiwt azurermApplicationInsightsWorkbookTemplateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aaiwt.ref.Append("tags"))
}

// TemplateData returns a reference to field template_data of azurerm_application_insights_workbook_template.
func (aaiwt azurermApplicationInsightsWorkbookTemplateAttributes) TemplateData() terra.StringValue {
	return terra.ReferenceAsString(aaiwt.ref.Append("template_data"))
}

func (aaiwt azurermApplicationInsightsWorkbookTemplateAttributes) Galleries() terra.ListValue[GalleriesAttributes] {
	return terra.ReferenceAsList[GalleriesAttributes](aaiwt.ref.Append("galleries"))
}

func (aaiwt azurermApplicationInsightsWorkbookTemplateAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aaiwt.ref.Append("timeouts"))
}

type azurermApplicationInsightsWorkbookTemplateState struct {
	Author            string            `json:"author"`
	Id                string            `json:"id"`
	Localized         string            `json:"localized"`
	Location          string            `json:"location"`
	Name              string            `json:"name"`
	Priority          float64           `json:"priority"`
	ResourceGroupName string            `json:"resource_group_name"`
	Tags              map[string]string `json:"tags"`
	TemplateData      string            `json:"template_data"`
	Galleries         []GalleriesState  `json:"galleries"`
	Timeouts          *TimeoutsState    `json:"timeouts"`
}
