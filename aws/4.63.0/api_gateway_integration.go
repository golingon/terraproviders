// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	apigatewayintegration "github.com/golingon/terraproviders/aws/4.63.0/apigatewayintegration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayIntegration creates a new instance of [ApiGatewayIntegration].
func NewApiGatewayIntegration(name string, args ApiGatewayIntegrationArgs) *ApiGatewayIntegration {
	return &ApiGatewayIntegration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayIntegration)(nil)

// ApiGatewayIntegration represents the Terraform resource aws_api_gateway_integration.
type ApiGatewayIntegration struct {
	Name      string
	Args      ApiGatewayIntegrationArgs
	state     *apiGatewayIntegrationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayIntegration].
func (agi *ApiGatewayIntegration) Type() string {
	return "aws_api_gateway_integration"
}

// LocalName returns the local name for [ApiGatewayIntegration].
func (agi *ApiGatewayIntegration) LocalName() string {
	return agi.Name
}

// Configuration returns the configuration (args) for [ApiGatewayIntegration].
func (agi *ApiGatewayIntegration) Configuration() interface{} {
	return agi.Args
}

// DependOn is used for other resources to depend on [ApiGatewayIntegration].
func (agi *ApiGatewayIntegration) DependOn() terra.Reference {
	return terra.ReferenceResource(agi)
}

// Dependencies returns the list of resources [ApiGatewayIntegration] depends_on.
func (agi *ApiGatewayIntegration) Dependencies() terra.Dependencies {
	return agi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayIntegration].
func (agi *ApiGatewayIntegration) LifecycleManagement() *terra.Lifecycle {
	return agi.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayIntegration].
func (agi *ApiGatewayIntegration) Attributes() apiGatewayIntegrationAttributes {
	return apiGatewayIntegrationAttributes{ref: terra.ReferenceResource(agi)}
}

// ImportState imports the given attribute values into [ApiGatewayIntegration]'s state.
func (agi *ApiGatewayIntegration) ImportState(av io.Reader) error {
	agi.state = &apiGatewayIntegrationState{}
	if err := json.NewDecoder(av).Decode(agi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agi.Type(), agi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayIntegration] has state.
func (agi *ApiGatewayIntegration) State() (*apiGatewayIntegrationState, bool) {
	return agi.state, agi.state != nil
}

// StateMust returns the state for [ApiGatewayIntegration]. Panics if the state is nil.
func (agi *ApiGatewayIntegration) StateMust() *apiGatewayIntegrationState {
	if agi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agi.Type(), agi.LocalName()))
	}
	return agi.state
}

// ApiGatewayIntegrationArgs contains the configurations for aws_api_gateway_integration.
type ApiGatewayIntegrationArgs struct {
	// CacheKeyParameters: set of string, optional
	CacheKeyParameters terra.SetValue[terra.StringValue] `hcl:"cache_key_parameters,attr"`
	// CacheNamespace: string, optional
	CacheNamespace terra.StringValue `hcl:"cache_namespace,attr"`
	// ConnectionId: string, optional
	ConnectionId terra.StringValue `hcl:"connection_id,attr"`
	// ConnectionType: string, optional
	ConnectionType terra.StringValue `hcl:"connection_type,attr"`
	// ContentHandling: string, optional
	ContentHandling terra.StringValue `hcl:"content_handling,attr"`
	// Credentials: string, optional
	Credentials terra.StringValue `hcl:"credentials,attr"`
	// HttpMethod: string, required
	HttpMethod terra.StringValue `hcl:"http_method,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntegrationHttpMethod: string, optional
	IntegrationHttpMethod terra.StringValue `hcl:"integration_http_method,attr"`
	// PassthroughBehavior: string, optional
	PassthroughBehavior terra.StringValue `hcl:"passthrough_behavior,attr"`
	// RequestParameters: map of string, optional
	RequestParameters terra.MapValue[terra.StringValue] `hcl:"request_parameters,attr"`
	// RequestTemplates: map of string, optional
	RequestTemplates terra.MapValue[terra.StringValue] `hcl:"request_templates,attr"`
	// ResourceId: string, required
	ResourceId terra.StringValue `hcl:"resource_id,attr" validate:"required"`
	// RestApiId: string, required
	RestApiId terra.StringValue `hcl:"rest_api_id,attr" validate:"required"`
	// TimeoutMilliseconds: number, optional
	TimeoutMilliseconds terra.NumberValue `hcl:"timeout_milliseconds,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Uri: string, optional
	Uri terra.StringValue `hcl:"uri,attr"`
	// TlsConfig: optional
	TlsConfig *apigatewayintegration.TlsConfig `hcl:"tls_config,block"`
}
type apiGatewayIntegrationAttributes struct {
	ref terra.Reference
}

// CacheKeyParameters returns a reference to field cache_key_parameters of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) CacheKeyParameters() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](agi.ref.Append("cache_key_parameters"))
}

// CacheNamespace returns a reference to field cache_namespace of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) CacheNamespace() terra.StringValue {
	return terra.ReferenceAsString(agi.ref.Append("cache_namespace"))
}

// ConnectionId returns a reference to field connection_id of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(agi.ref.Append("connection_id"))
}

// ConnectionType returns a reference to field connection_type of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) ConnectionType() terra.StringValue {
	return terra.ReferenceAsString(agi.ref.Append("connection_type"))
}

// ContentHandling returns a reference to field content_handling of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) ContentHandling() terra.StringValue {
	return terra.ReferenceAsString(agi.ref.Append("content_handling"))
}

// Credentials returns a reference to field credentials of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) Credentials() terra.StringValue {
	return terra.ReferenceAsString(agi.ref.Append("credentials"))
}

// HttpMethod returns a reference to field http_method of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) HttpMethod() terra.StringValue {
	return terra.ReferenceAsString(agi.ref.Append("http_method"))
}

// Id returns a reference to field id of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agi.ref.Append("id"))
}

// IntegrationHttpMethod returns a reference to field integration_http_method of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) IntegrationHttpMethod() terra.StringValue {
	return terra.ReferenceAsString(agi.ref.Append("integration_http_method"))
}

// PassthroughBehavior returns a reference to field passthrough_behavior of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) PassthroughBehavior() terra.StringValue {
	return terra.ReferenceAsString(agi.ref.Append("passthrough_behavior"))
}

// RequestParameters returns a reference to field request_parameters of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) RequestParameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agi.ref.Append("request_parameters"))
}

// RequestTemplates returns a reference to field request_templates of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) RequestTemplates() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agi.ref.Append("request_templates"))
}

// ResourceId returns a reference to field resource_id of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(agi.ref.Append("resource_id"))
}

// RestApiId returns a reference to field rest_api_id of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) RestApiId() terra.StringValue {
	return terra.ReferenceAsString(agi.ref.Append("rest_api_id"))
}

// TimeoutMilliseconds returns a reference to field timeout_milliseconds of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) TimeoutMilliseconds() terra.NumberValue {
	return terra.ReferenceAsNumber(agi.ref.Append("timeout_milliseconds"))
}

// Type returns a reference to field type of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(agi.ref.Append("type"))
}

// Uri returns a reference to field uri of aws_api_gateway_integration.
func (agi apiGatewayIntegrationAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(agi.ref.Append("uri"))
}

func (agi apiGatewayIntegrationAttributes) TlsConfig() terra.ListValue[apigatewayintegration.TlsConfigAttributes] {
	return terra.ReferenceAsList[apigatewayintegration.TlsConfigAttributes](agi.ref.Append("tls_config"))
}

type apiGatewayIntegrationState struct {
	CacheKeyParameters    []string                               `json:"cache_key_parameters"`
	CacheNamespace        string                                 `json:"cache_namespace"`
	ConnectionId          string                                 `json:"connection_id"`
	ConnectionType        string                                 `json:"connection_type"`
	ContentHandling       string                                 `json:"content_handling"`
	Credentials           string                                 `json:"credentials"`
	HttpMethod            string                                 `json:"http_method"`
	Id                    string                                 `json:"id"`
	IntegrationHttpMethod string                                 `json:"integration_http_method"`
	PassthroughBehavior   string                                 `json:"passthrough_behavior"`
	RequestParameters     map[string]string                      `json:"request_parameters"`
	RequestTemplates      map[string]string                      `json:"request_templates"`
	ResourceId            string                                 `json:"resource_id"`
	RestApiId             string                                 `json:"rest_api_id"`
	TimeoutMilliseconds   float64                                `json:"timeout_milliseconds"`
	Type                  string                                 `json:"type"`
	Uri                   string                                 `json:"uri"`
	TlsConfig             []apigatewayintegration.TlsConfigState `json:"tls_config"`
}
