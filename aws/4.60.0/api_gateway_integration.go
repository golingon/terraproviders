// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	apigatewayintegration "github.com/golingon/terraproviders/aws/4.60.0/apigatewayintegration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewApiGatewayIntegration(name string, args ApiGatewayIntegrationArgs) *ApiGatewayIntegration {
	return &ApiGatewayIntegration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayIntegration)(nil)

type ApiGatewayIntegration struct {
	Name  string
	Args  ApiGatewayIntegrationArgs
	state *apiGatewayIntegrationState
}

func (agi *ApiGatewayIntegration) Type() string {
	return "aws_api_gateway_integration"
}

func (agi *ApiGatewayIntegration) LocalName() string {
	return agi.Name
}

func (agi *ApiGatewayIntegration) Configuration() interface{} {
	return agi.Args
}

func (agi *ApiGatewayIntegration) Attributes() apiGatewayIntegrationAttributes {
	return apiGatewayIntegrationAttributes{ref: terra.ReferenceResource(agi)}
}

func (agi *ApiGatewayIntegration) ImportState(av io.Reader) error {
	agi.state = &apiGatewayIntegrationState{}
	if err := json.NewDecoder(av).Decode(agi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agi.Type(), agi.LocalName(), err)
	}
	return nil
}

func (agi *ApiGatewayIntegration) State() (*apiGatewayIntegrationState, bool) {
	return agi.state, agi.state != nil
}

func (agi *ApiGatewayIntegration) StateMust() *apiGatewayIntegrationState {
	if agi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agi.Type(), agi.LocalName()))
	}
	return agi.state
}

func (agi *ApiGatewayIntegration) DependOn() terra.Reference {
	return terra.ReferenceResource(agi)
}

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
	// DependsOn contains resources that ApiGatewayIntegration depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type apiGatewayIntegrationAttributes struct {
	ref terra.Reference
}

func (agi apiGatewayIntegrationAttributes) CacheKeyParameters() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](agi.ref.Append("cache_key_parameters"))
}

func (agi apiGatewayIntegrationAttributes) CacheNamespace() terra.StringValue {
	return terra.ReferenceString(agi.ref.Append("cache_namespace"))
}

func (agi apiGatewayIntegrationAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceString(agi.ref.Append("connection_id"))
}

func (agi apiGatewayIntegrationAttributes) ConnectionType() terra.StringValue {
	return terra.ReferenceString(agi.ref.Append("connection_type"))
}

func (agi apiGatewayIntegrationAttributes) ContentHandling() terra.StringValue {
	return terra.ReferenceString(agi.ref.Append("content_handling"))
}

func (agi apiGatewayIntegrationAttributes) Credentials() terra.StringValue {
	return terra.ReferenceString(agi.ref.Append("credentials"))
}

func (agi apiGatewayIntegrationAttributes) HttpMethod() terra.StringValue {
	return terra.ReferenceString(agi.ref.Append("http_method"))
}

func (agi apiGatewayIntegrationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(agi.ref.Append("id"))
}

func (agi apiGatewayIntegrationAttributes) IntegrationHttpMethod() terra.StringValue {
	return terra.ReferenceString(agi.ref.Append("integration_http_method"))
}

func (agi apiGatewayIntegrationAttributes) PassthroughBehavior() terra.StringValue {
	return terra.ReferenceString(agi.ref.Append("passthrough_behavior"))
}

func (agi apiGatewayIntegrationAttributes) RequestParameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](agi.ref.Append("request_parameters"))
}

func (agi apiGatewayIntegrationAttributes) RequestTemplates() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](agi.ref.Append("request_templates"))
}

func (agi apiGatewayIntegrationAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceString(agi.ref.Append("resource_id"))
}

func (agi apiGatewayIntegrationAttributes) RestApiId() terra.StringValue {
	return terra.ReferenceString(agi.ref.Append("rest_api_id"))
}

func (agi apiGatewayIntegrationAttributes) TimeoutMilliseconds() terra.NumberValue {
	return terra.ReferenceNumber(agi.ref.Append("timeout_milliseconds"))
}

func (agi apiGatewayIntegrationAttributes) Type() terra.StringValue {
	return terra.ReferenceString(agi.ref.Append("type"))
}

func (agi apiGatewayIntegrationAttributes) Uri() terra.StringValue {
	return terra.ReferenceString(agi.ref.Append("uri"))
}

func (agi apiGatewayIntegrationAttributes) TlsConfig() terra.ListValue[apigatewayintegration.TlsConfigAttributes] {
	return terra.ReferenceList[apigatewayintegration.TlsConfigAttributes](agi.ref.Append("tls_config"))
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
