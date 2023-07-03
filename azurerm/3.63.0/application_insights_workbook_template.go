// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	applicationinsightsworkbooktemplate "github.com/golingon/terraproviders/azurerm/3.63.0/applicationinsightsworkbooktemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApplicationInsightsWorkbookTemplate creates a new instance of [ApplicationInsightsWorkbookTemplate].
func NewApplicationInsightsWorkbookTemplate(name string, args ApplicationInsightsWorkbookTemplateArgs) *ApplicationInsightsWorkbookTemplate {
	return &ApplicationInsightsWorkbookTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApplicationInsightsWorkbookTemplate)(nil)

// ApplicationInsightsWorkbookTemplate represents the Terraform resource azurerm_application_insights_workbook_template.
type ApplicationInsightsWorkbookTemplate struct {
	Name      string
	Args      ApplicationInsightsWorkbookTemplateArgs
	state     *applicationInsightsWorkbookTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApplicationInsightsWorkbookTemplate].
func (aiwt *ApplicationInsightsWorkbookTemplate) Type() string {
	return "azurerm_application_insights_workbook_template"
}

// LocalName returns the local name for [ApplicationInsightsWorkbookTemplate].
func (aiwt *ApplicationInsightsWorkbookTemplate) LocalName() string {
	return aiwt.Name
}

// Configuration returns the configuration (args) for [ApplicationInsightsWorkbookTemplate].
func (aiwt *ApplicationInsightsWorkbookTemplate) Configuration() interface{} {
	return aiwt.Args
}

// DependOn is used for other resources to depend on [ApplicationInsightsWorkbookTemplate].
func (aiwt *ApplicationInsightsWorkbookTemplate) DependOn() terra.Reference {
	return terra.ReferenceResource(aiwt)
}

// Dependencies returns the list of resources [ApplicationInsightsWorkbookTemplate] depends_on.
func (aiwt *ApplicationInsightsWorkbookTemplate) Dependencies() terra.Dependencies {
	return aiwt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApplicationInsightsWorkbookTemplate].
func (aiwt *ApplicationInsightsWorkbookTemplate) LifecycleManagement() *terra.Lifecycle {
	return aiwt.Lifecycle
}

// Attributes returns the attributes for [ApplicationInsightsWorkbookTemplate].
func (aiwt *ApplicationInsightsWorkbookTemplate) Attributes() applicationInsightsWorkbookTemplateAttributes {
	return applicationInsightsWorkbookTemplateAttributes{ref: terra.ReferenceResource(aiwt)}
}

// ImportState imports the given attribute values into [ApplicationInsightsWorkbookTemplate]'s state.
func (aiwt *ApplicationInsightsWorkbookTemplate) ImportState(av io.Reader) error {
	aiwt.state = &applicationInsightsWorkbookTemplateState{}
	if err := json.NewDecoder(av).Decode(aiwt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aiwt.Type(), aiwt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApplicationInsightsWorkbookTemplate] has state.
func (aiwt *ApplicationInsightsWorkbookTemplate) State() (*applicationInsightsWorkbookTemplateState, bool) {
	return aiwt.state, aiwt.state != nil
}

// StateMust returns the state for [ApplicationInsightsWorkbookTemplate]. Panics if the state is nil.
func (aiwt *ApplicationInsightsWorkbookTemplate) StateMust() *applicationInsightsWorkbookTemplateState {
	if aiwt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aiwt.Type(), aiwt.LocalName()))
	}
	return aiwt.state
}

// ApplicationInsightsWorkbookTemplateArgs contains the configurations for azurerm_application_insights_workbook_template.
type ApplicationInsightsWorkbookTemplateArgs struct {
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
	Galleries []applicationinsightsworkbooktemplate.Galleries `hcl:"galleries,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *applicationinsightsworkbooktemplate.Timeouts `hcl:"timeouts,block"`
}
type applicationInsightsWorkbookTemplateAttributes struct {
	ref terra.Reference
}

// Author returns a reference to field author of azurerm_application_insights_workbook_template.
func (aiwt applicationInsightsWorkbookTemplateAttributes) Author() terra.StringValue {
	return terra.ReferenceAsString(aiwt.ref.Append("author"))
}

// Id returns a reference to field id of azurerm_application_insights_workbook_template.
func (aiwt applicationInsightsWorkbookTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aiwt.ref.Append("id"))
}

// Localized returns a reference to field localized of azurerm_application_insights_workbook_template.
func (aiwt applicationInsightsWorkbookTemplateAttributes) Localized() terra.StringValue {
	return terra.ReferenceAsString(aiwt.ref.Append("localized"))
}

// Location returns a reference to field location of azurerm_application_insights_workbook_template.
func (aiwt applicationInsightsWorkbookTemplateAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aiwt.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_application_insights_workbook_template.
func (aiwt applicationInsightsWorkbookTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aiwt.ref.Append("name"))
}

// Priority returns a reference to field priority of azurerm_application_insights_workbook_template.
func (aiwt applicationInsightsWorkbookTemplateAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(aiwt.ref.Append("priority"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_application_insights_workbook_template.
func (aiwt applicationInsightsWorkbookTemplateAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aiwt.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_application_insights_workbook_template.
func (aiwt applicationInsightsWorkbookTemplateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aiwt.ref.Append("tags"))
}

// TemplateData returns a reference to field template_data of azurerm_application_insights_workbook_template.
func (aiwt applicationInsightsWorkbookTemplateAttributes) TemplateData() terra.StringValue {
	return terra.ReferenceAsString(aiwt.ref.Append("template_data"))
}

func (aiwt applicationInsightsWorkbookTemplateAttributes) Galleries() terra.ListValue[applicationinsightsworkbooktemplate.GalleriesAttributes] {
	return terra.ReferenceAsList[applicationinsightsworkbooktemplate.GalleriesAttributes](aiwt.ref.Append("galleries"))
}

func (aiwt applicationInsightsWorkbookTemplateAttributes) Timeouts() applicationinsightsworkbooktemplate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[applicationinsightsworkbooktemplate.TimeoutsAttributes](aiwt.ref.Append("timeouts"))
}

type applicationInsightsWorkbookTemplateState struct {
	Author            string                                               `json:"author"`
	Id                string                                               `json:"id"`
	Localized         string                                               `json:"localized"`
	Location          string                                               `json:"location"`
	Name              string                                               `json:"name"`
	Priority          float64                                              `json:"priority"`
	ResourceGroupName string                                               `json:"resource_group_name"`
	Tags              map[string]string                                    `json:"tags"`
	TemplateData      string                                               `json:"template_data"`
	Galleries         []applicationinsightsworkbooktemplate.GalleriesState `json:"galleries"`
	Timeouts          *applicationinsightsworkbooktemplate.TimeoutsState   `json:"timeouts"`
}
