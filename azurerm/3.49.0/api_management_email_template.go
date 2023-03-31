// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementemailtemplate "github.com/golingon/terraproviders/azurerm/3.49.0/apimanagementemailtemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewApiManagementEmailTemplate(name string, args ApiManagementEmailTemplateArgs) *ApiManagementEmailTemplate {
	return &ApiManagementEmailTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementEmailTemplate)(nil)

type ApiManagementEmailTemplate struct {
	Name  string
	Args  ApiManagementEmailTemplateArgs
	state *apiManagementEmailTemplateState
}

func (amet *ApiManagementEmailTemplate) Type() string {
	return "azurerm_api_management_email_template"
}

func (amet *ApiManagementEmailTemplate) LocalName() string {
	return amet.Name
}

func (amet *ApiManagementEmailTemplate) Configuration() interface{} {
	return amet.Args
}

func (amet *ApiManagementEmailTemplate) Attributes() apiManagementEmailTemplateAttributes {
	return apiManagementEmailTemplateAttributes{ref: terra.ReferenceResource(amet)}
}

func (amet *ApiManagementEmailTemplate) ImportState(av io.Reader) error {
	amet.state = &apiManagementEmailTemplateState{}
	if err := json.NewDecoder(av).Decode(amet.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amet.Type(), amet.LocalName(), err)
	}
	return nil
}

func (amet *ApiManagementEmailTemplate) State() (*apiManagementEmailTemplateState, bool) {
	return amet.state, amet.state != nil
}

func (amet *ApiManagementEmailTemplate) StateMust() *apiManagementEmailTemplateState {
	if amet.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amet.Type(), amet.LocalName()))
	}
	return amet.state
}

func (amet *ApiManagementEmailTemplate) DependOn() terra.Reference {
	return terra.ReferenceResource(amet)
}

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
	// DependsOn contains resources that ApiManagementEmailTemplate depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type apiManagementEmailTemplateAttributes struct {
	ref terra.Reference
}

func (amet apiManagementEmailTemplateAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceString(amet.ref.Append("api_management_name"))
}

func (amet apiManagementEmailTemplateAttributes) Body() terra.StringValue {
	return terra.ReferenceString(amet.ref.Append("body"))
}

func (amet apiManagementEmailTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceString(amet.ref.Append("description"))
}

func (amet apiManagementEmailTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceString(amet.ref.Append("id"))
}

func (amet apiManagementEmailTemplateAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(amet.ref.Append("resource_group_name"))
}

func (amet apiManagementEmailTemplateAttributes) Subject() terra.StringValue {
	return terra.ReferenceString(amet.ref.Append("subject"))
}

func (amet apiManagementEmailTemplateAttributes) TemplateName() terra.StringValue {
	return terra.ReferenceString(amet.ref.Append("template_name"))
}

func (amet apiManagementEmailTemplateAttributes) Title() terra.StringValue {
	return terra.ReferenceString(amet.ref.Append("title"))
}

func (amet apiManagementEmailTemplateAttributes) Timeouts() apimanagementemailtemplate.TimeoutsAttributes {
	return terra.ReferenceSingle[apimanagementemailtemplate.TimeoutsAttributes](amet.ref.Append("timeouts"))
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
