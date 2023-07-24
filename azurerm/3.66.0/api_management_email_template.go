// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementemailtemplate "github.com/golingon/terraproviders/azurerm/3.66.0/apimanagementemailtemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementEmailTemplate creates a new instance of [ApiManagementEmailTemplate].
func NewApiManagementEmailTemplate(name string, args ApiManagementEmailTemplateArgs) *ApiManagementEmailTemplate {
	return &ApiManagementEmailTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementEmailTemplate)(nil)

// ApiManagementEmailTemplate represents the Terraform resource azurerm_api_management_email_template.
type ApiManagementEmailTemplate struct {
	Name      string
	Args      ApiManagementEmailTemplateArgs
	state     *apiManagementEmailTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementEmailTemplate].
func (amet *ApiManagementEmailTemplate) Type() string {
	return "azurerm_api_management_email_template"
}

// LocalName returns the local name for [ApiManagementEmailTemplate].
func (amet *ApiManagementEmailTemplate) LocalName() string {
	return amet.Name
}

// Configuration returns the configuration (args) for [ApiManagementEmailTemplate].
func (amet *ApiManagementEmailTemplate) Configuration() interface{} {
	return amet.Args
}

// DependOn is used for other resources to depend on [ApiManagementEmailTemplate].
func (amet *ApiManagementEmailTemplate) DependOn() terra.Reference {
	return terra.ReferenceResource(amet)
}

// Dependencies returns the list of resources [ApiManagementEmailTemplate] depends_on.
func (amet *ApiManagementEmailTemplate) Dependencies() terra.Dependencies {
	return amet.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementEmailTemplate].
func (amet *ApiManagementEmailTemplate) LifecycleManagement() *terra.Lifecycle {
	return amet.Lifecycle
}

// Attributes returns the attributes for [ApiManagementEmailTemplate].
func (amet *ApiManagementEmailTemplate) Attributes() apiManagementEmailTemplateAttributes {
	return apiManagementEmailTemplateAttributes{ref: terra.ReferenceResource(amet)}
}

// ImportState imports the given attribute values into [ApiManagementEmailTemplate]'s state.
func (amet *ApiManagementEmailTemplate) ImportState(av io.Reader) error {
	amet.state = &apiManagementEmailTemplateState{}
	if err := json.NewDecoder(av).Decode(amet.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amet.Type(), amet.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementEmailTemplate] has state.
func (amet *ApiManagementEmailTemplate) State() (*apiManagementEmailTemplateState, bool) {
	return amet.state, amet.state != nil
}

// StateMust returns the state for [ApiManagementEmailTemplate]. Panics if the state is nil.
func (amet *ApiManagementEmailTemplate) StateMust() *apiManagementEmailTemplateState {
	if amet.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amet.Type(), amet.LocalName()))
	}
	return amet.state
}

// ApiManagementEmailTemplateArgs contains the configurations for azurerm_api_management_email_template.
type ApiManagementEmailTemplateArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// Body: string, required
	Body terra.StringValue `hcl:"body,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Subject: string, required
	Subject terra.StringValue `hcl:"subject,attr" validate:"required"`
	// TemplateName: string, required
	TemplateName terra.StringValue `hcl:"template_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apimanagementemailtemplate.Timeouts `hcl:"timeouts,block"`
}
type apiManagementEmailTemplateAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_email_template.
func (amet apiManagementEmailTemplateAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amet.ref.Append("api_management_name"))
}

// Body returns a reference to field body of azurerm_api_management_email_template.
func (amet apiManagementEmailTemplateAttributes) Body() terra.StringValue {
	return terra.ReferenceAsString(amet.ref.Append("body"))
}

// Description returns a reference to field description of azurerm_api_management_email_template.
func (amet apiManagementEmailTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amet.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_api_management_email_template.
func (amet apiManagementEmailTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amet.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_email_template.
func (amet apiManagementEmailTemplateAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amet.ref.Append("resource_group_name"))
}

// Subject returns a reference to field subject of azurerm_api_management_email_template.
func (amet apiManagementEmailTemplateAttributes) Subject() terra.StringValue {
	return terra.ReferenceAsString(amet.ref.Append("subject"))
}

// TemplateName returns a reference to field template_name of azurerm_api_management_email_template.
func (amet apiManagementEmailTemplateAttributes) TemplateName() terra.StringValue {
	return terra.ReferenceAsString(amet.ref.Append("template_name"))
}

// Title returns a reference to field title of azurerm_api_management_email_template.
func (amet apiManagementEmailTemplateAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(amet.ref.Append("title"))
}

func (amet apiManagementEmailTemplateAttributes) Timeouts() apimanagementemailtemplate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementemailtemplate.TimeoutsAttributes](amet.ref.Append("timeouts"))
}

type apiManagementEmailTemplateState struct {
	ApiManagementName string                                    `json:"api_management_name"`
	Body              string                                    `json:"body"`
	Description       string                                    `json:"description"`
	Id                string                                    `json:"id"`
	ResourceGroupName string                                    `json:"resource_group_name"`
	Subject           string                                    `json:"subject"`
	TemplateName      string                                    `json:"template_name"`
	Title             string                                    `json:"title"`
	Timeouts          *apimanagementemailtemplate.TimeoutsState `json:"timeouts"`
}
