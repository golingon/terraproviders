// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	apigatewayapiconfig "github.com/golingon/terraproviders/googlebeta/4.62.0/apigatewayapiconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayApiConfig creates a new instance of [ApiGatewayApiConfig].
func NewApiGatewayApiConfig(name string, args ApiGatewayApiConfigArgs) *ApiGatewayApiConfig {
	return &ApiGatewayApiConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayApiConfig)(nil)

// ApiGatewayApiConfig represents the Terraform resource google_api_gateway_api_config.
type ApiGatewayApiConfig struct {
	Name      string
	Args      ApiGatewayApiConfigArgs
	state     *apiGatewayApiConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayApiConfig].
func (agac *ApiGatewayApiConfig) Type() string {
	return "google_api_gateway_api_config"
}

// LocalName returns the local name for [ApiGatewayApiConfig].
func (agac *ApiGatewayApiConfig) LocalName() string {
	return agac.Name
}

// Configuration returns the configuration (args) for [ApiGatewayApiConfig].
func (agac *ApiGatewayApiConfig) Configuration() interface{} {
	return agac.Args
}

// DependOn is used for other resources to depend on [ApiGatewayApiConfig].
func (agac *ApiGatewayApiConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(agac)
}

// Dependencies returns the list of resources [ApiGatewayApiConfig] depends_on.
func (agac *ApiGatewayApiConfig) Dependencies() terra.Dependencies {
	return agac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayApiConfig].
func (agac *ApiGatewayApiConfig) LifecycleManagement() *terra.Lifecycle {
	return agac.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayApiConfig].
func (agac *ApiGatewayApiConfig) Attributes() apiGatewayApiConfigAttributes {
	return apiGatewayApiConfigAttributes{ref: terra.ReferenceResource(agac)}
}

// ImportState imports the given attribute values into [ApiGatewayApiConfig]'s state.
func (agac *ApiGatewayApiConfig) ImportState(av io.Reader) error {
	agac.state = &apiGatewayApiConfigState{}
	if err := json.NewDecoder(av).Decode(agac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agac.Type(), agac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayApiConfig] has state.
func (agac *ApiGatewayApiConfig) State() (*apiGatewayApiConfigState, bool) {
	return agac.state, agac.state != nil
}

// StateMust returns the state for [ApiGatewayApiConfig]. Panics if the state is nil.
func (agac *ApiGatewayApiConfig) StateMust() *apiGatewayApiConfigState {
	if agac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agac.Type(), agac.LocalName()))
	}
	return agac.state
}

// ApiGatewayApiConfigArgs contains the configurations for google_api_gateway_api_config.
type ApiGatewayApiConfigArgs struct {
	// Api: string, required
	Api terra.StringValue `hcl:"api,attr" validate:"required"`
	// ApiConfigId: string, optional
	ApiConfigId terra.StringValue `hcl:"api_config_id,attr"`
	// ApiConfigIdPrefix: string, optional
	ApiConfigIdPrefix terra.StringValue `hcl:"api_config_id_prefix,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// GatewayConfig: optional
	GatewayConfig *apigatewayapiconfig.GatewayConfig `hcl:"gateway_config,block"`
	// GrpcServices: min=0
	GrpcServices []apigatewayapiconfig.GrpcServices `hcl:"grpc_services,block" validate:"min=0"`
	// ManagedServiceConfigs: min=0
	ManagedServiceConfigs []apigatewayapiconfig.ManagedServiceConfigs `hcl:"managed_service_configs,block" validate:"min=0"`
	// OpenapiDocuments: min=0
	OpenapiDocuments []apigatewayapiconfig.OpenapiDocuments `hcl:"openapi_documents,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *apigatewayapiconfig.Timeouts `hcl:"timeouts,block"`
}
type apiGatewayApiConfigAttributes struct {
	ref terra.Reference
}

// Api returns a reference to field api of google_api_gateway_api_config.
func (agac apiGatewayApiConfigAttributes) Api() terra.StringValue {
	return terra.ReferenceAsString(agac.ref.Append("api"))
}

// ApiConfigId returns a reference to field api_config_id of google_api_gateway_api_config.
func (agac apiGatewayApiConfigAttributes) ApiConfigId() terra.StringValue {
	return terra.ReferenceAsString(agac.ref.Append("api_config_id"))
}

// ApiConfigIdPrefix returns a reference to field api_config_id_prefix of google_api_gateway_api_config.
func (agac apiGatewayApiConfigAttributes) ApiConfigIdPrefix() terra.StringValue {
	return terra.ReferenceAsString(agac.ref.Append("api_config_id_prefix"))
}

// DisplayName returns a reference to field display_name of google_api_gateway_api_config.
func (agac apiGatewayApiConfigAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(agac.ref.Append("display_name"))
}

// Id returns a reference to field id of google_api_gateway_api_config.
func (agac apiGatewayApiConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agac.ref.Append("id"))
}

// Labels returns a reference to field labels of google_api_gateway_api_config.
func (agac apiGatewayApiConfigAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agac.ref.Append("labels"))
}

// Name returns a reference to field name of google_api_gateway_api_config.
func (agac apiGatewayApiConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(agac.ref.Append("name"))
}

// Project returns a reference to field project of google_api_gateway_api_config.
func (agac apiGatewayApiConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(agac.ref.Append("project"))
}

// ServiceConfigId returns a reference to field service_config_id of google_api_gateway_api_config.
func (agac apiGatewayApiConfigAttributes) ServiceConfigId() terra.StringValue {
	return terra.ReferenceAsString(agac.ref.Append("service_config_id"))
}

func (agac apiGatewayApiConfigAttributes) GatewayConfig() terra.ListValue[apigatewayapiconfig.GatewayConfigAttributes] {
	return terra.ReferenceAsList[apigatewayapiconfig.GatewayConfigAttributes](agac.ref.Append("gateway_config"))
}

func (agac apiGatewayApiConfigAttributes) GrpcServices() terra.ListValue[apigatewayapiconfig.GrpcServicesAttributes] {
	return terra.ReferenceAsList[apigatewayapiconfig.GrpcServicesAttributes](agac.ref.Append("grpc_services"))
}

func (agac apiGatewayApiConfigAttributes) ManagedServiceConfigs() terra.ListValue[apigatewayapiconfig.ManagedServiceConfigsAttributes] {
	return terra.ReferenceAsList[apigatewayapiconfig.ManagedServiceConfigsAttributes](agac.ref.Append("managed_service_configs"))
}

func (agac apiGatewayApiConfigAttributes) OpenapiDocuments() terra.ListValue[apigatewayapiconfig.OpenapiDocumentsAttributes] {
	return terra.ReferenceAsList[apigatewayapiconfig.OpenapiDocumentsAttributes](agac.ref.Append("openapi_documents"))
}

func (agac apiGatewayApiConfigAttributes) Timeouts() apigatewayapiconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigatewayapiconfig.TimeoutsAttributes](agac.ref.Append("timeouts"))
}

type apiGatewayApiConfigState struct {
	Api                   string                                           `json:"api"`
	ApiConfigId           string                                           `json:"api_config_id"`
	ApiConfigIdPrefix     string                                           `json:"api_config_id_prefix"`
	DisplayName           string                                           `json:"display_name"`
	Id                    string                                           `json:"id"`
	Labels                map[string]string                                `json:"labels"`
	Name                  string                                           `json:"name"`
	Project               string                                           `json:"project"`
	ServiceConfigId       string                                           `json:"service_config_id"`
	GatewayConfig         []apigatewayapiconfig.GatewayConfigState         `json:"gateway_config"`
	GrpcServices          []apigatewayapiconfig.GrpcServicesState          `json:"grpc_services"`
	ManagedServiceConfigs []apigatewayapiconfig.ManagedServiceConfigsState `json:"managed_service_configs"`
	OpenapiDocuments      []apigatewayapiconfig.OpenapiDocumentsState      `json:"openapi_documents"`
	Timeouts              *apigatewayapiconfig.TimeoutsState               `json:"timeouts"`
}
