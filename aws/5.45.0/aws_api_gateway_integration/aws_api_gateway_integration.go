// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_api_gateway_integration

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

// Resource represents the Terraform resource aws_api_gateway_integration.
type Resource struct {
	Name      string
	Args      Args
	state     *awsApiGatewayIntegrationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aagi *Resource) Type() string {
	return "aws_api_gateway_integration"
}

// LocalName returns the local name for [Resource].
func (aagi *Resource) LocalName() string {
	return aagi.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aagi *Resource) Configuration() interface{} {
	return aagi.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aagi *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aagi)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aagi *Resource) Dependencies() terra.Dependencies {
	return aagi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aagi *Resource) LifecycleManagement() *terra.Lifecycle {
	return aagi.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aagi *Resource) Attributes() awsApiGatewayIntegrationAttributes {
	return awsApiGatewayIntegrationAttributes{ref: terra.ReferenceResource(aagi)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aagi *Resource) ImportState(state io.Reader) error {
	aagi.state = &awsApiGatewayIntegrationState{}
	if err := json.NewDecoder(state).Decode(aagi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aagi.Type(), aagi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aagi *Resource) State() (*awsApiGatewayIntegrationState, bool) {
	return aagi.state, aagi.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aagi *Resource) StateMust() *awsApiGatewayIntegrationState {
	if aagi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aagi.Type(), aagi.LocalName()))
	}
	return aagi.state
}

// Args contains the configurations for aws_api_gateway_integration.
type Args struct {
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
	TlsConfig *TlsConfig `hcl:"tls_config,block"`
}

type awsApiGatewayIntegrationAttributes struct {
	ref terra.Reference
}

// CacheKeyParameters returns a reference to field cache_key_parameters of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) CacheKeyParameters() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aagi.ref.Append("cache_key_parameters"))
}

// CacheNamespace returns a reference to field cache_namespace of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) CacheNamespace() terra.StringValue {
	return terra.ReferenceAsString(aagi.ref.Append("cache_namespace"))
}

// ConnectionId returns a reference to field connection_id of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(aagi.ref.Append("connection_id"))
}

// ConnectionType returns a reference to field connection_type of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) ConnectionType() terra.StringValue {
	return terra.ReferenceAsString(aagi.ref.Append("connection_type"))
}

// ContentHandling returns a reference to field content_handling of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) ContentHandling() terra.StringValue {
	return terra.ReferenceAsString(aagi.ref.Append("content_handling"))
}

// Credentials returns a reference to field credentials of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) Credentials() terra.StringValue {
	return terra.ReferenceAsString(aagi.ref.Append("credentials"))
}

// HttpMethod returns a reference to field http_method of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) HttpMethod() terra.StringValue {
	return terra.ReferenceAsString(aagi.ref.Append("http_method"))
}

// Id returns a reference to field id of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aagi.ref.Append("id"))
}

// IntegrationHttpMethod returns a reference to field integration_http_method of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) IntegrationHttpMethod() terra.StringValue {
	return terra.ReferenceAsString(aagi.ref.Append("integration_http_method"))
}

// PassthroughBehavior returns a reference to field passthrough_behavior of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) PassthroughBehavior() terra.StringValue {
	return terra.ReferenceAsString(aagi.ref.Append("passthrough_behavior"))
}

// RequestParameters returns a reference to field request_parameters of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) RequestParameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aagi.ref.Append("request_parameters"))
}

// RequestTemplates returns a reference to field request_templates of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) RequestTemplates() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aagi.ref.Append("request_templates"))
}

// ResourceId returns a reference to field resource_id of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(aagi.ref.Append("resource_id"))
}

// RestApiId returns a reference to field rest_api_id of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) RestApiId() terra.StringValue {
	return terra.ReferenceAsString(aagi.ref.Append("rest_api_id"))
}

// TimeoutMilliseconds returns a reference to field timeout_milliseconds of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) TimeoutMilliseconds() terra.NumberValue {
	return terra.ReferenceAsNumber(aagi.ref.Append("timeout_milliseconds"))
}

// Type returns a reference to field type of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(aagi.ref.Append("type"))
}

// Uri returns a reference to field uri of aws_api_gateway_integration.
func (aagi awsApiGatewayIntegrationAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(aagi.ref.Append("uri"))
}

func (aagi awsApiGatewayIntegrationAttributes) TlsConfig() terra.ListValue[TlsConfigAttributes] {
	return terra.ReferenceAsList[TlsConfigAttributes](aagi.ref.Append("tls_config"))
}

type awsApiGatewayIntegrationState struct {
	CacheKeyParameters    []string          `json:"cache_key_parameters"`
	CacheNamespace        string            `json:"cache_namespace"`
	ConnectionId          string            `json:"connection_id"`
	ConnectionType        string            `json:"connection_type"`
	ContentHandling       string            `json:"content_handling"`
	Credentials           string            `json:"credentials"`
	HttpMethod            string            `json:"http_method"`
	Id                    string            `json:"id"`
	IntegrationHttpMethod string            `json:"integration_http_method"`
	PassthroughBehavior   string            `json:"passthrough_behavior"`
	RequestParameters     map[string]string `json:"request_parameters"`
	RequestTemplates      map[string]string `json:"request_templates"`
	ResourceId            string            `json:"resource_id"`
	RestApiId             string            `json:"rest_api_id"`
	TimeoutMilliseconds   float64           `json:"timeout_milliseconds"`
	Type                  string            `json:"type"`
	Uri                   string            `json:"uri"`
	TlsConfig             []TlsConfigState  `json:"tls_config"`
}
