// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementapioperation "github.com/golingon/terraproviders/azurerm/3.67.0/apimanagementapioperation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementApiOperation creates a new instance of [ApiManagementApiOperation].
func NewApiManagementApiOperation(name string, args ApiManagementApiOperationArgs) *ApiManagementApiOperation {
	return &ApiManagementApiOperation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementApiOperation)(nil)

// ApiManagementApiOperation represents the Terraform resource azurerm_api_management_api_operation.
type ApiManagementApiOperation struct {
	Name      string
	Args      ApiManagementApiOperationArgs
	state     *apiManagementApiOperationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementApiOperation].
func (amao *ApiManagementApiOperation) Type() string {
	return "azurerm_api_management_api_operation"
}

// LocalName returns the local name for [ApiManagementApiOperation].
func (amao *ApiManagementApiOperation) LocalName() string {
	return amao.Name
}

// Configuration returns the configuration (args) for [ApiManagementApiOperation].
func (amao *ApiManagementApiOperation) Configuration() interface{} {
	return amao.Args
}

// DependOn is used for other resources to depend on [ApiManagementApiOperation].
func (amao *ApiManagementApiOperation) DependOn() terra.Reference {
	return terra.ReferenceResource(amao)
}

// Dependencies returns the list of resources [ApiManagementApiOperation] depends_on.
func (amao *ApiManagementApiOperation) Dependencies() terra.Dependencies {
	return amao.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementApiOperation].
func (amao *ApiManagementApiOperation) LifecycleManagement() *terra.Lifecycle {
	return amao.Lifecycle
}

// Attributes returns the attributes for [ApiManagementApiOperation].
func (amao *ApiManagementApiOperation) Attributes() apiManagementApiOperationAttributes {
	return apiManagementApiOperationAttributes{ref: terra.ReferenceResource(amao)}
}

// ImportState imports the given attribute values into [ApiManagementApiOperation]'s state.
func (amao *ApiManagementApiOperation) ImportState(av io.Reader) error {
	amao.state = &apiManagementApiOperationState{}
	if err := json.NewDecoder(av).Decode(amao.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amao.Type(), amao.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementApiOperation] has state.
func (amao *ApiManagementApiOperation) State() (*apiManagementApiOperationState, bool) {
	return amao.state, amao.state != nil
}

// StateMust returns the state for [ApiManagementApiOperation]. Panics if the state is nil.
func (amao *ApiManagementApiOperation) StateMust() *apiManagementApiOperationState {
	if amao.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amao.Type(), amao.LocalName()))
	}
	return amao.state
}

// ApiManagementApiOperationArgs contains the configurations for azurerm_api_management_api_operation.
type ApiManagementApiOperationArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// ApiName: string, required
	ApiName terra.StringValue `hcl:"api_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Method: string, required
	Method terra.StringValue `hcl:"method,attr" validate:"required"`
	// OperationId: string, required
	OperationId terra.StringValue `hcl:"operation_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// UrlTemplate: string, required
	UrlTemplate terra.StringValue `hcl:"url_template,attr" validate:"required"`
	// Request: optional
	Request *apimanagementapioperation.Request `hcl:"request,block"`
	// Response: min=0
	Response []apimanagementapioperation.Response `hcl:"response,block" validate:"min=0"`
	// TemplateParameter: min=0
	TemplateParameter []apimanagementapioperation.TemplateParameter `hcl:"template_parameter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *apimanagementapioperation.Timeouts `hcl:"timeouts,block"`
}
type apiManagementApiOperationAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_api_operation.
func (amao apiManagementApiOperationAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amao.ref.Append("api_management_name"))
}

// ApiName returns a reference to field api_name of azurerm_api_management_api_operation.
func (amao apiManagementApiOperationAttributes) ApiName() terra.StringValue {
	return terra.ReferenceAsString(amao.ref.Append("api_name"))
}

// Description returns a reference to field description of azurerm_api_management_api_operation.
func (amao apiManagementApiOperationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amao.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_api_operation.
func (amao apiManagementApiOperationAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(amao.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_api_management_api_operation.
func (amao apiManagementApiOperationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amao.ref.Append("id"))
}

// Method returns a reference to field method of azurerm_api_management_api_operation.
func (amao apiManagementApiOperationAttributes) Method() terra.StringValue {
	return terra.ReferenceAsString(amao.ref.Append("method"))
}

// OperationId returns a reference to field operation_id of azurerm_api_management_api_operation.
func (amao apiManagementApiOperationAttributes) OperationId() terra.StringValue {
	return terra.ReferenceAsString(amao.ref.Append("operation_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_api_operation.
func (amao apiManagementApiOperationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amao.ref.Append("resource_group_name"))
}

// UrlTemplate returns a reference to field url_template of azurerm_api_management_api_operation.
func (amao apiManagementApiOperationAttributes) UrlTemplate() terra.StringValue {
	return terra.ReferenceAsString(amao.ref.Append("url_template"))
}

func (amao apiManagementApiOperationAttributes) Request() terra.ListValue[apimanagementapioperation.RequestAttributes] {
	return terra.ReferenceAsList[apimanagementapioperation.RequestAttributes](amao.ref.Append("request"))
}

func (amao apiManagementApiOperationAttributes) Response() terra.ListValue[apimanagementapioperation.ResponseAttributes] {
	return terra.ReferenceAsList[apimanagementapioperation.ResponseAttributes](amao.ref.Append("response"))
}

func (amao apiManagementApiOperationAttributes) TemplateParameter() terra.ListValue[apimanagementapioperation.TemplateParameterAttributes] {
	return terra.ReferenceAsList[apimanagementapioperation.TemplateParameterAttributes](amao.ref.Append("template_parameter"))
}

func (amao apiManagementApiOperationAttributes) Timeouts() apimanagementapioperation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementapioperation.TimeoutsAttributes](amao.ref.Append("timeouts"))
}

type apiManagementApiOperationState struct {
	ApiManagementName string                                             `json:"api_management_name"`
	ApiName           string                                             `json:"api_name"`
	Description       string                                             `json:"description"`
	DisplayName       string                                             `json:"display_name"`
	Id                string                                             `json:"id"`
	Method            string                                             `json:"method"`
	OperationId       string                                             `json:"operation_id"`
	ResourceGroupName string                                             `json:"resource_group_name"`
	UrlTemplate       string                                             `json:"url_template"`
	Request           []apimanagementapioperation.RequestState           `json:"request"`
	Response          []apimanagementapioperation.ResponseState          `json:"response"`
	TemplateParameter []apimanagementapioperation.TemplateParameterState `json:"template_parameter"`
	Timeouts          *apimanagementapioperation.TimeoutsState           `json:"timeouts"`
}
