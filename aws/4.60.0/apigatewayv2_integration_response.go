// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewApigatewayv2IntegrationResponse(name string, args Apigatewayv2IntegrationResponseArgs) *Apigatewayv2IntegrationResponse {
	return &Apigatewayv2IntegrationResponse{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Apigatewayv2IntegrationResponse)(nil)

type Apigatewayv2IntegrationResponse struct {
	Name  string
	Args  Apigatewayv2IntegrationResponseArgs
	state *apigatewayv2IntegrationResponseState
}

func (air *Apigatewayv2IntegrationResponse) Type() string {
	return "aws_apigatewayv2_integration_response"
}

func (air *Apigatewayv2IntegrationResponse) LocalName() string {
	return air.Name
}

func (air *Apigatewayv2IntegrationResponse) Configuration() interface{} {
	return air.Args
}

func (air *Apigatewayv2IntegrationResponse) Attributes() apigatewayv2IntegrationResponseAttributes {
	return apigatewayv2IntegrationResponseAttributes{ref: terra.ReferenceResource(air)}
}

func (air *Apigatewayv2IntegrationResponse) ImportState(av io.Reader) error {
	air.state = &apigatewayv2IntegrationResponseState{}
	if err := json.NewDecoder(av).Decode(air.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", air.Type(), air.LocalName(), err)
	}
	return nil
}

func (air *Apigatewayv2IntegrationResponse) State() (*apigatewayv2IntegrationResponseState, bool) {
	return air.state, air.state != nil
}

func (air *Apigatewayv2IntegrationResponse) StateMust() *apigatewayv2IntegrationResponseState {
	if air.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", air.Type(), air.LocalName()))
	}
	return air.state
}

func (air *Apigatewayv2IntegrationResponse) DependOn() terra.Reference {
	return terra.ReferenceResource(air)
}

type Apigatewayv2IntegrationResponseArgs struct {
	// ApiId: string, required
	ApiId terra.StringValue `hcl:"api_id,attr" validate:"required"`
	// ContentHandlingStrategy: string, optional
	ContentHandlingStrategy terra.StringValue `hcl:"content_handling_strategy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntegrationId: string, required
	IntegrationId terra.StringValue `hcl:"integration_id,attr" validate:"required"`
	// IntegrationResponseKey: string, required
	IntegrationResponseKey terra.StringValue `hcl:"integration_response_key,attr" validate:"required"`
	// ResponseTemplates: map of string, optional
	ResponseTemplates terra.MapValue[terra.StringValue] `hcl:"response_templates,attr"`
	// TemplateSelectionExpression: string, optional
	TemplateSelectionExpression terra.StringValue `hcl:"template_selection_expression,attr"`
	// DependsOn contains resources that Apigatewayv2IntegrationResponse depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type apigatewayv2IntegrationResponseAttributes struct {
	ref terra.Reference
}

func (air apigatewayv2IntegrationResponseAttributes) ApiId() terra.StringValue {
	return terra.ReferenceString(air.ref.Append("api_id"))
}

func (air apigatewayv2IntegrationResponseAttributes) ContentHandlingStrategy() terra.StringValue {
	return terra.ReferenceString(air.ref.Append("content_handling_strategy"))
}

func (air apigatewayv2IntegrationResponseAttributes) Id() terra.StringValue {
	return terra.ReferenceString(air.ref.Append("id"))
}

func (air apigatewayv2IntegrationResponseAttributes) IntegrationId() terra.StringValue {
	return terra.ReferenceString(air.ref.Append("integration_id"))
}

func (air apigatewayv2IntegrationResponseAttributes) IntegrationResponseKey() terra.StringValue {
	return terra.ReferenceString(air.ref.Append("integration_response_key"))
}

func (air apigatewayv2IntegrationResponseAttributes) ResponseTemplates() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](air.ref.Append("response_templates"))
}

func (air apigatewayv2IntegrationResponseAttributes) TemplateSelectionExpression() terra.StringValue {
	return terra.ReferenceString(air.ref.Append("template_selection_expression"))
}

type apigatewayv2IntegrationResponseState struct {
	ApiId                       string            `json:"api_id"`
	ContentHandlingStrategy     string            `json:"content_handling_strategy"`
	Id                          string            `json:"id"`
	IntegrationId               string            `json:"integration_id"`
	IntegrationResponseKey      string            `json:"integration_response_key"`
	ResponseTemplates           map[string]string `json:"response_templates"`
	TemplateSelectionExpression string            `json:"template_selection_expression"`
}
