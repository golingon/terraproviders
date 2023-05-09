// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	apigatewaystage "github.com/golingon/terraproviders/aws/4.66.1/apigatewaystage"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayStage creates a new instance of [ApiGatewayStage].
func NewApiGatewayStage(name string, args ApiGatewayStageArgs) *ApiGatewayStage {
	return &ApiGatewayStage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayStage)(nil)

// ApiGatewayStage represents the Terraform resource aws_api_gateway_stage.
type ApiGatewayStage struct {
	Name      string
	Args      ApiGatewayStageArgs
	state     *apiGatewayStageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayStage].
func (ags *ApiGatewayStage) Type() string {
	return "aws_api_gateway_stage"
}

// LocalName returns the local name for [ApiGatewayStage].
func (ags *ApiGatewayStage) LocalName() string {
	return ags.Name
}

// Configuration returns the configuration (args) for [ApiGatewayStage].
func (ags *ApiGatewayStage) Configuration() interface{} {
	return ags.Args
}

// DependOn is used for other resources to depend on [ApiGatewayStage].
func (ags *ApiGatewayStage) DependOn() terra.Reference {
	return terra.ReferenceResource(ags)
}

// Dependencies returns the list of resources [ApiGatewayStage] depends_on.
func (ags *ApiGatewayStage) Dependencies() terra.Dependencies {
	return ags.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayStage].
func (ags *ApiGatewayStage) LifecycleManagement() *terra.Lifecycle {
	return ags.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayStage].
func (ags *ApiGatewayStage) Attributes() apiGatewayStageAttributes {
	return apiGatewayStageAttributes{ref: terra.ReferenceResource(ags)}
}

// ImportState imports the given attribute values into [ApiGatewayStage]'s state.
func (ags *ApiGatewayStage) ImportState(av io.Reader) error {
	ags.state = &apiGatewayStageState{}
	if err := json.NewDecoder(av).Decode(ags.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ags.Type(), ags.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayStage] has state.
func (ags *ApiGatewayStage) State() (*apiGatewayStageState, bool) {
	return ags.state, ags.state != nil
}

// StateMust returns the state for [ApiGatewayStage]. Panics if the state is nil.
func (ags *ApiGatewayStage) StateMust() *apiGatewayStageState {
	if ags.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ags.Type(), ags.LocalName()))
	}
	return ags.state
}

// ApiGatewayStageArgs contains the configurations for aws_api_gateway_stage.
type ApiGatewayStageArgs struct {
	// CacheClusterEnabled: bool, optional
	CacheClusterEnabled terra.BoolValue `hcl:"cache_cluster_enabled,attr"`
	// CacheClusterSize: string, optional
	CacheClusterSize terra.StringValue `hcl:"cache_cluster_size,attr"`
	// ClientCertificateId: string, optional
	ClientCertificateId terra.StringValue `hcl:"client_certificate_id,attr"`
	// DeploymentId: string, required
	DeploymentId terra.StringValue `hcl:"deployment_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DocumentationVersion: string, optional
	DocumentationVersion terra.StringValue `hcl:"documentation_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RestApiId: string, required
	RestApiId terra.StringValue `hcl:"rest_api_id,attr" validate:"required"`
	// StageName: string, required
	StageName terra.StringValue `hcl:"stage_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Variables: map of string, optional
	Variables terra.MapValue[terra.StringValue] `hcl:"variables,attr"`
	// XrayTracingEnabled: bool, optional
	XrayTracingEnabled terra.BoolValue `hcl:"xray_tracing_enabled,attr"`
	// AccessLogSettings: optional
	AccessLogSettings *apigatewaystage.AccessLogSettings `hcl:"access_log_settings,block"`
	// CanarySettings: optional
	CanarySettings *apigatewaystage.CanarySettings `hcl:"canary_settings,block"`
}
type apiGatewayStageAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("arn"))
}

// CacheClusterEnabled returns a reference to field cache_cluster_enabled of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) CacheClusterEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ags.ref.Append("cache_cluster_enabled"))
}

// CacheClusterSize returns a reference to field cache_cluster_size of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) CacheClusterSize() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("cache_cluster_size"))
}

// ClientCertificateId returns a reference to field client_certificate_id of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) ClientCertificateId() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("client_certificate_id"))
}

// DeploymentId returns a reference to field deployment_id of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) DeploymentId() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("deployment_id"))
}

// Description returns a reference to field description of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("description"))
}

// DocumentationVersion returns a reference to field documentation_version of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) DocumentationVersion() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("documentation_version"))
}

// ExecutionArn returns a reference to field execution_arn of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) ExecutionArn() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("execution_arn"))
}

// Id returns a reference to field id of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("id"))
}

// InvokeUrl returns a reference to field invoke_url of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) InvokeUrl() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("invoke_url"))
}

// RestApiId returns a reference to field rest_api_id of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) RestApiId() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("rest_api_id"))
}

// StageName returns a reference to field stage_name of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) StageName() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("stage_name"))
}

// Tags returns a reference to field tags of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ags.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ags.ref.Append("tags_all"))
}

// Variables returns a reference to field variables of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) Variables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ags.ref.Append("variables"))
}

// WebAclArn returns a reference to field web_acl_arn of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) WebAclArn() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("web_acl_arn"))
}

// XrayTracingEnabled returns a reference to field xray_tracing_enabled of aws_api_gateway_stage.
func (ags apiGatewayStageAttributes) XrayTracingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ags.ref.Append("xray_tracing_enabled"))
}

func (ags apiGatewayStageAttributes) AccessLogSettings() terra.ListValue[apigatewaystage.AccessLogSettingsAttributes] {
	return terra.ReferenceAsList[apigatewaystage.AccessLogSettingsAttributes](ags.ref.Append("access_log_settings"))
}

func (ags apiGatewayStageAttributes) CanarySettings() terra.ListValue[apigatewaystage.CanarySettingsAttributes] {
	return terra.ReferenceAsList[apigatewaystage.CanarySettingsAttributes](ags.ref.Append("canary_settings"))
}

type apiGatewayStageState struct {
	Arn                  string                                   `json:"arn"`
	CacheClusterEnabled  bool                                     `json:"cache_cluster_enabled"`
	CacheClusterSize     string                                   `json:"cache_cluster_size"`
	ClientCertificateId  string                                   `json:"client_certificate_id"`
	DeploymentId         string                                   `json:"deployment_id"`
	Description          string                                   `json:"description"`
	DocumentationVersion string                                   `json:"documentation_version"`
	ExecutionArn         string                                   `json:"execution_arn"`
	Id                   string                                   `json:"id"`
	InvokeUrl            string                                   `json:"invoke_url"`
	RestApiId            string                                   `json:"rest_api_id"`
	StageName            string                                   `json:"stage_name"`
	Tags                 map[string]string                        `json:"tags"`
	TagsAll              map[string]string                        `json:"tags_all"`
	Variables            map[string]string                        `json:"variables"`
	WebAclArn            string                                   `json:"web_acl_arn"`
	XrayTracingEnabled   bool                                     `json:"xray_tracing_enabled"`
	AccessLogSettings    []apigatewaystage.AccessLogSettingsState `json:"access_log_settings"`
	CanarySettings       []apigatewaystage.CanarySettingsState    `json:"canary_settings"`
}
