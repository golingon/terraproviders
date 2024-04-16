// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_apigatewayv2_api

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_apigatewayv2_api.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aaa *DataSource) DataSource() string {
	return "aws_apigatewayv2_api"
}

// LocalName returns the local name for [DataSource].
func (aaa *DataSource) LocalName() string {
	return aaa.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aaa *DataSource) Configuration() interface{} {
	return aaa.Args
}

// Attributes returns the attributes for [DataSource].
func (aaa *DataSource) Attributes() dataAwsApigatewayv2ApiAttributes {
	return dataAwsApigatewayv2ApiAttributes{ref: terra.ReferenceDataSource(aaa)}
}

// DataArgs contains the configurations for aws_apigatewayv2_api.
type DataArgs struct {
	// ApiId: string, required
	ApiId terra.StringValue `hcl:"api_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}

type dataAwsApigatewayv2ApiAttributes struct {
	ref terra.Reference
}

// ApiEndpoint returns a reference to field api_endpoint of aws_apigatewayv2_api.
func (aaa dataAwsApigatewayv2ApiAttributes) ApiEndpoint() terra.StringValue {
	return terra.ReferenceAsString(aaa.ref.Append("api_endpoint"))
}

// ApiId returns a reference to field api_id of aws_apigatewayv2_api.
func (aaa dataAwsApigatewayv2ApiAttributes) ApiId() terra.StringValue {
	return terra.ReferenceAsString(aaa.ref.Append("api_id"))
}

// ApiKeySelectionExpression returns a reference to field api_key_selection_expression of aws_apigatewayv2_api.
func (aaa dataAwsApigatewayv2ApiAttributes) ApiKeySelectionExpression() terra.StringValue {
	return terra.ReferenceAsString(aaa.ref.Append("api_key_selection_expression"))
}

// Arn returns a reference to field arn of aws_apigatewayv2_api.
func (aaa dataAwsApigatewayv2ApiAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aaa.ref.Append("arn"))
}

// Description returns a reference to field description of aws_apigatewayv2_api.
func (aaa dataAwsApigatewayv2ApiAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aaa.ref.Append("description"))
}

// DisableExecuteApiEndpoint returns a reference to field disable_execute_api_endpoint of aws_apigatewayv2_api.
func (aaa dataAwsApigatewayv2ApiAttributes) DisableExecuteApiEndpoint() terra.BoolValue {
	return terra.ReferenceAsBool(aaa.ref.Append("disable_execute_api_endpoint"))
}

// ExecutionArn returns a reference to field execution_arn of aws_apigatewayv2_api.
func (aaa dataAwsApigatewayv2ApiAttributes) ExecutionArn() terra.StringValue {
	return terra.ReferenceAsString(aaa.ref.Append("execution_arn"))
}

// Id returns a reference to field id of aws_apigatewayv2_api.
func (aaa dataAwsApigatewayv2ApiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aaa.ref.Append("id"))
}

// Name returns a reference to field name of aws_apigatewayv2_api.
func (aaa dataAwsApigatewayv2ApiAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aaa.ref.Append("name"))
}

// ProtocolType returns a reference to field protocol_type of aws_apigatewayv2_api.
func (aaa dataAwsApigatewayv2ApiAttributes) ProtocolType() terra.StringValue {
	return terra.ReferenceAsString(aaa.ref.Append("protocol_type"))
}

// RouteSelectionExpression returns a reference to field route_selection_expression of aws_apigatewayv2_api.
func (aaa dataAwsApigatewayv2ApiAttributes) RouteSelectionExpression() terra.StringValue {
	return terra.ReferenceAsString(aaa.ref.Append("route_selection_expression"))
}

// Tags returns a reference to field tags of aws_apigatewayv2_api.
func (aaa dataAwsApigatewayv2ApiAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aaa.ref.Append("tags"))
}

// Version returns a reference to field version of aws_apigatewayv2_api.
func (aaa dataAwsApigatewayv2ApiAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(aaa.ref.Append("version"))
}

func (aaa dataAwsApigatewayv2ApiAttributes) CorsConfiguration() terra.ListValue[DataCorsConfigurationAttributes] {
	return terra.ReferenceAsList[DataCorsConfigurationAttributes](aaa.ref.Append("cors_configuration"))
}
