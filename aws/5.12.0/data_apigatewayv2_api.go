// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataapigatewayv2api "github.com/golingon/terraproviders/aws/5.12.0/dataapigatewayv2api"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataApigatewayv2Api creates a new instance of [DataApigatewayv2Api].
func NewDataApigatewayv2Api(name string, args DataApigatewayv2ApiArgs) *DataApigatewayv2Api {
	return &DataApigatewayv2Api{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApigatewayv2Api)(nil)

// DataApigatewayv2Api represents the Terraform data resource aws_apigatewayv2_api.
type DataApigatewayv2Api struct {
	Name string
	Args DataApigatewayv2ApiArgs
}

// DataSource returns the Terraform object type for [DataApigatewayv2Api].
func (aa *DataApigatewayv2Api) DataSource() string {
	return "aws_apigatewayv2_api"
}

// LocalName returns the local name for [DataApigatewayv2Api].
func (aa *DataApigatewayv2Api) LocalName() string {
	return aa.Name
}

// Configuration returns the configuration (args) for [DataApigatewayv2Api].
func (aa *DataApigatewayv2Api) Configuration() interface{} {
	return aa.Args
}

// Attributes returns the attributes for [DataApigatewayv2Api].
func (aa *DataApigatewayv2Api) Attributes() dataApigatewayv2ApiAttributes {
	return dataApigatewayv2ApiAttributes{ref: terra.ReferenceDataResource(aa)}
}

// DataApigatewayv2ApiArgs contains the configurations for aws_apigatewayv2_api.
type DataApigatewayv2ApiArgs struct {
	// ApiId: string, required
	ApiId terra.StringValue `hcl:"api_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// CorsConfiguration: min=0
	CorsConfiguration []dataapigatewayv2api.CorsConfiguration `hcl:"cors_configuration,block" validate:"min=0"`
}
type dataApigatewayv2ApiAttributes struct {
	ref terra.Reference
}

// ApiEndpoint returns a reference to field api_endpoint of aws_apigatewayv2_api.
func (aa dataApigatewayv2ApiAttributes) ApiEndpoint() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("api_endpoint"))
}

// ApiId returns a reference to field api_id of aws_apigatewayv2_api.
func (aa dataApigatewayv2ApiAttributes) ApiId() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("api_id"))
}

// ApiKeySelectionExpression returns a reference to field api_key_selection_expression of aws_apigatewayv2_api.
func (aa dataApigatewayv2ApiAttributes) ApiKeySelectionExpression() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("api_key_selection_expression"))
}

// Arn returns a reference to field arn of aws_apigatewayv2_api.
func (aa dataApigatewayv2ApiAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("arn"))
}

// Description returns a reference to field description of aws_apigatewayv2_api.
func (aa dataApigatewayv2ApiAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("description"))
}

// DisableExecuteApiEndpoint returns a reference to field disable_execute_api_endpoint of aws_apigatewayv2_api.
func (aa dataApigatewayv2ApiAttributes) DisableExecuteApiEndpoint() terra.BoolValue {
	return terra.ReferenceAsBool(aa.ref.Append("disable_execute_api_endpoint"))
}

// ExecutionArn returns a reference to field execution_arn of aws_apigatewayv2_api.
func (aa dataApigatewayv2ApiAttributes) ExecutionArn() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("execution_arn"))
}

// Id returns a reference to field id of aws_apigatewayv2_api.
func (aa dataApigatewayv2ApiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("id"))
}

// Name returns a reference to field name of aws_apigatewayv2_api.
func (aa dataApigatewayv2ApiAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("name"))
}

// ProtocolType returns a reference to field protocol_type of aws_apigatewayv2_api.
func (aa dataApigatewayv2ApiAttributes) ProtocolType() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("protocol_type"))
}

// RouteSelectionExpression returns a reference to field route_selection_expression of aws_apigatewayv2_api.
func (aa dataApigatewayv2ApiAttributes) RouteSelectionExpression() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("route_selection_expression"))
}

// Tags returns a reference to field tags of aws_apigatewayv2_api.
func (aa dataApigatewayv2ApiAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aa.ref.Append("tags"))
}

// Version returns a reference to field version of aws_apigatewayv2_api.
func (aa dataApigatewayv2ApiAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("version"))
}

func (aa dataApigatewayv2ApiAttributes) CorsConfiguration() terra.ListValue[dataapigatewayv2api.CorsConfigurationAttributes] {
	return terra.ReferenceAsList[dataapigatewayv2api.CorsConfigurationAttributes](aa.ref.Append("cors_configuration"))
}
