// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	apigatewayv2stage "github.com/golingon/terraproviders/aws/5.7.0/apigatewayv2stage"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigatewayv2Stage creates a new instance of [Apigatewayv2Stage].
func NewApigatewayv2Stage(name string, args Apigatewayv2StageArgs) *Apigatewayv2Stage {
	return &Apigatewayv2Stage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Apigatewayv2Stage)(nil)

// Apigatewayv2Stage represents the Terraform resource aws_apigatewayv2_stage.
type Apigatewayv2Stage struct {
	Name      string
	Args      Apigatewayv2StageArgs
	state     *apigatewayv2StageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Apigatewayv2Stage].
func (as *Apigatewayv2Stage) Type() string {
	return "aws_apigatewayv2_stage"
}

// LocalName returns the local name for [Apigatewayv2Stage].
func (as *Apigatewayv2Stage) LocalName() string {
	return as.Name
}

// Configuration returns the configuration (args) for [Apigatewayv2Stage].
func (as *Apigatewayv2Stage) Configuration() interface{} {
	return as.Args
}

// DependOn is used for other resources to depend on [Apigatewayv2Stage].
func (as *Apigatewayv2Stage) DependOn() terra.Reference {
	return terra.ReferenceResource(as)
}

// Dependencies returns the list of resources [Apigatewayv2Stage] depends_on.
func (as *Apigatewayv2Stage) Dependencies() terra.Dependencies {
	return as.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Apigatewayv2Stage].
func (as *Apigatewayv2Stage) LifecycleManagement() *terra.Lifecycle {
	return as.Lifecycle
}

// Attributes returns the attributes for [Apigatewayv2Stage].
func (as *Apigatewayv2Stage) Attributes() apigatewayv2StageAttributes {
	return apigatewayv2StageAttributes{ref: terra.ReferenceResource(as)}
}

// ImportState imports the given attribute values into [Apigatewayv2Stage]'s state.
func (as *Apigatewayv2Stage) ImportState(av io.Reader) error {
	as.state = &apigatewayv2StageState{}
	if err := json.NewDecoder(av).Decode(as.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", as.Type(), as.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Apigatewayv2Stage] has state.
func (as *Apigatewayv2Stage) State() (*apigatewayv2StageState, bool) {
	return as.state, as.state != nil
}

// StateMust returns the state for [Apigatewayv2Stage]. Panics if the state is nil.
func (as *Apigatewayv2Stage) StateMust() *apigatewayv2StageState {
	if as.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", as.Type(), as.LocalName()))
	}
	return as.state
}

// Apigatewayv2StageArgs contains the configurations for aws_apigatewayv2_stage.
type Apigatewayv2StageArgs struct {
	// ApiId: string, required
	ApiId terra.StringValue `hcl:"api_id,attr" validate:"required"`
	// AutoDeploy: bool, optional
	AutoDeploy terra.BoolValue `hcl:"auto_deploy,attr"`
	// ClientCertificateId: string, optional
	ClientCertificateId terra.StringValue `hcl:"client_certificate_id,attr"`
	// DeploymentId: string, optional
	DeploymentId terra.StringValue `hcl:"deployment_id,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StageVariables: map of string, optional
	StageVariables terra.MapValue[terra.StringValue] `hcl:"stage_variables,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// AccessLogSettings: optional
	AccessLogSettings *apigatewayv2stage.AccessLogSettings `hcl:"access_log_settings,block"`
	// DefaultRouteSettings: optional
	DefaultRouteSettings *apigatewayv2stage.DefaultRouteSettings `hcl:"default_route_settings,block"`
	// RouteSettings: min=0
	RouteSettings []apigatewayv2stage.RouteSettings `hcl:"route_settings,block" validate:"min=0"`
}
type apigatewayv2StageAttributes struct {
	ref terra.Reference
}

// ApiId returns a reference to field api_id of aws_apigatewayv2_stage.
func (as apigatewayv2StageAttributes) ApiId() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("api_id"))
}

// Arn returns a reference to field arn of aws_apigatewayv2_stage.
func (as apigatewayv2StageAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("arn"))
}

// AutoDeploy returns a reference to field auto_deploy of aws_apigatewayv2_stage.
func (as apigatewayv2StageAttributes) AutoDeploy() terra.BoolValue {
	return terra.ReferenceAsBool(as.ref.Append("auto_deploy"))
}

// ClientCertificateId returns a reference to field client_certificate_id of aws_apigatewayv2_stage.
func (as apigatewayv2StageAttributes) ClientCertificateId() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("client_certificate_id"))
}

// DeploymentId returns a reference to field deployment_id of aws_apigatewayv2_stage.
func (as apigatewayv2StageAttributes) DeploymentId() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("deployment_id"))
}

// Description returns a reference to field description of aws_apigatewayv2_stage.
func (as apigatewayv2StageAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("description"))
}

// ExecutionArn returns a reference to field execution_arn of aws_apigatewayv2_stage.
func (as apigatewayv2StageAttributes) ExecutionArn() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("execution_arn"))
}

// Id returns a reference to field id of aws_apigatewayv2_stage.
func (as apigatewayv2StageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("id"))
}

// InvokeUrl returns a reference to field invoke_url of aws_apigatewayv2_stage.
func (as apigatewayv2StageAttributes) InvokeUrl() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("invoke_url"))
}

// Name returns a reference to field name of aws_apigatewayv2_stage.
func (as apigatewayv2StageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("name"))
}

// StageVariables returns a reference to field stage_variables of aws_apigatewayv2_stage.
func (as apigatewayv2StageAttributes) StageVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](as.ref.Append("stage_variables"))
}

// Tags returns a reference to field tags of aws_apigatewayv2_stage.
func (as apigatewayv2StageAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](as.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_apigatewayv2_stage.
func (as apigatewayv2StageAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](as.ref.Append("tags_all"))
}

func (as apigatewayv2StageAttributes) AccessLogSettings() terra.ListValue[apigatewayv2stage.AccessLogSettingsAttributes] {
	return terra.ReferenceAsList[apigatewayv2stage.AccessLogSettingsAttributes](as.ref.Append("access_log_settings"))
}

func (as apigatewayv2StageAttributes) DefaultRouteSettings() terra.ListValue[apigatewayv2stage.DefaultRouteSettingsAttributes] {
	return terra.ReferenceAsList[apigatewayv2stage.DefaultRouteSettingsAttributes](as.ref.Append("default_route_settings"))
}

func (as apigatewayv2StageAttributes) RouteSettings() terra.SetValue[apigatewayv2stage.RouteSettingsAttributes] {
	return terra.ReferenceAsSet[apigatewayv2stage.RouteSettingsAttributes](as.ref.Append("route_settings"))
}

type apigatewayv2StageState struct {
	ApiId                string                                        `json:"api_id"`
	Arn                  string                                        `json:"arn"`
	AutoDeploy           bool                                          `json:"auto_deploy"`
	ClientCertificateId  string                                        `json:"client_certificate_id"`
	DeploymentId         string                                        `json:"deployment_id"`
	Description          string                                        `json:"description"`
	ExecutionArn         string                                        `json:"execution_arn"`
	Id                   string                                        `json:"id"`
	InvokeUrl            string                                        `json:"invoke_url"`
	Name                 string                                        `json:"name"`
	StageVariables       map[string]string                             `json:"stage_variables"`
	Tags                 map[string]string                             `json:"tags"`
	TagsAll              map[string]string                             `json:"tags_all"`
	AccessLogSettings    []apigatewayv2stage.AccessLogSettingsState    `json:"access_log_settings"`
	DefaultRouteSettings []apigatewayv2stage.DefaultRouteSettingsState `json:"default_route_settings"`
	RouteSettings        []apigatewayv2stage.RouteSettingsState        `json:"route_settings"`
}
