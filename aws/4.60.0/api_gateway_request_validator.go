// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewApiGatewayRequestValidator(name string, args ApiGatewayRequestValidatorArgs) *ApiGatewayRequestValidator {
	return &ApiGatewayRequestValidator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayRequestValidator)(nil)

type ApiGatewayRequestValidator struct {
	Name  string
	Args  ApiGatewayRequestValidatorArgs
	state *apiGatewayRequestValidatorState
}

func (agrv *ApiGatewayRequestValidator) Type() string {
	return "aws_api_gateway_request_validator"
}

func (agrv *ApiGatewayRequestValidator) LocalName() string {
	return agrv.Name
}

func (agrv *ApiGatewayRequestValidator) Configuration() interface{} {
	return agrv.Args
}

func (agrv *ApiGatewayRequestValidator) Attributes() apiGatewayRequestValidatorAttributes {
	return apiGatewayRequestValidatorAttributes{ref: terra.ReferenceResource(agrv)}
}

func (agrv *ApiGatewayRequestValidator) ImportState(av io.Reader) error {
	agrv.state = &apiGatewayRequestValidatorState{}
	if err := json.NewDecoder(av).Decode(agrv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agrv.Type(), agrv.LocalName(), err)
	}
	return nil
}

func (agrv *ApiGatewayRequestValidator) State() (*apiGatewayRequestValidatorState, bool) {
	return agrv.state, agrv.state != nil
}

func (agrv *ApiGatewayRequestValidator) StateMust() *apiGatewayRequestValidatorState {
	if agrv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agrv.Type(), agrv.LocalName()))
	}
	return agrv.state
}

func (agrv *ApiGatewayRequestValidator) DependOn() terra.Reference {
	return terra.ReferenceResource(agrv)
}

type ApiGatewayRequestValidatorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RestApiId: string, required
	RestApiId terra.StringValue `hcl:"rest_api_id,attr" validate:"required"`
	// ValidateRequestBody: bool, optional
	ValidateRequestBody terra.BoolValue `hcl:"validate_request_body,attr"`
	// ValidateRequestParameters: bool, optional
	ValidateRequestParameters terra.BoolValue `hcl:"validate_request_parameters,attr"`
	// DependsOn contains resources that ApiGatewayRequestValidator depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type apiGatewayRequestValidatorAttributes struct {
	ref terra.Reference
}

func (agrv apiGatewayRequestValidatorAttributes) Id() terra.StringValue {
	return terra.ReferenceString(agrv.ref.Append("id"))
}

func (agrv apiGatewayRequestValidatorAttributes) Name() terra.StringValue {
	return terra.ReferenceString(agrv.ref.Append("name"))
}

func (agrv apiGatewayRequestValidatorAttributes) RestApiId() terra.StringValue {
	return terra.ReferenceString(agrv.ref.Append("rest_api_id"))
}

func (agrv apiGatewayRequestValidatorAttributes) ValidateRequestBody() terra.BoolValue {
	return terra.ReferenceBool(agrv.ref.Append("validate_request_body"))
}

func (agrv apiGatewayRequestValidatorAttributes) ValidateRequestParameters() terra.BoolValue {
	return terra.ReferenceBool(agrv.ref.Append("validate_request_parameters"))
}

type apiGatewayRequestValidatorState struct {
	Id                        string `json:"id"`
	Name                      string `json:"name"`
	RestApiId                 string `json:"rest_api_id"`
	ValidateRequestBody       bool   `json:"validate_request_body"`
	ValidateRequestParameters bool   `json:"validate_request_parameters"`
}
