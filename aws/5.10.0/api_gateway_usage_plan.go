// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	apigatewayusageplan "github.com/golingon/terraproviders/aws/5.10.0/apigatewayusageplan"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayUsagePlan creates a new instance of [ApiGatewayUsagePlan].
func NewApiGatewayUsagePlan(name string, args ApiGatewayUsagePlanArgs) *ApiGatewayUsagePlan {
	return &ApiGatewayUsagePlan{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayUsagePlan)(nil)

// ApiGatewayUsagePlan represents the Terraform resource aws_api_gateway_usage_plan.
type ApiGatewayUsagePlan struct {
	Name      string
	Args      ApiGatewayUsagePlanArgs
	state     *apiGatewayUsagePlanState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayUsagePlan].
func (agup *ApiGatewayUsagePlan) Type() string {
	return "aws_api_gateway_usage_plan"
}

// LocalName returns the local name for [ApiGatewayUsagePlan].
func (agup *ApiGatewayUsagePlan) LocalName() string {
	return agup.Name
}

// Configuration returns the configuration (args) for [ApiGatewayUsagePlan].
func (agup *ApiGatewayUsagePlan) Configuration() interface{} {
	return agup.Args
}

// DependOn is used for other resources to depend on [ApiGatewayUsagePlan].
func (agup *ApiGatewayUsagePlan) DependOn() terra.Reference {
	return terra.ReferenceResource(agup)
}

// Dependencies returns the list of resources [ApiGatewayUsagePlan] depends_on.
func (agup *ApiGatewayUsagePlan) Dependencies() terra.Dependencies {
	return agup.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayUsagePlan].
func (agup *ApiGatewayUsagePlan) LifecycleManagement() *terra.Lifecycle {
	return agup.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayUsagePlan].
func (agup *ApiGatewayUsagePlan) Attributes() apiGatewayUsagePlanAttributes {
	return apiGatewayUsagePlanAttributes{ref: terra.ReferenceResource(agup)}
}

// ImportState imports the given attribute values into [ApiGatewayUsagePlan]'s state.
func (agup *ApiGatewayUsagePlan) ImportState(av io.Reader) error {
	agup.state = &apiGatewayUsagePlanState{}
	if err := json.NewDecoder(av).Decode(agup.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agup.Type(), agup.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayUsagePlan] has state.
func (agup *ApiGatewayUsagePlan) State() (*apiGatewayUsagePlanState, bool) {
	return agup.state, agup.state != nil
}

// StateMust returns the state for [ApiGatewayUsagePlan]. Panics if the state is nil.
func (agup *ApiGatewayUsagePlan) StateMust() *apiGatewayUsagePlanState {
	if agup.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agup.Type(), agup.LocalName()))
	}
	return agup.state
}

// ApiGatewayUsagePlanArgs contains the configurations for aws_api_gateway_usage_plan.
type ApiGatewayUsagePlanArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProductCode: string, optional
	ProductCode terra.StringValue `hcl:"product_code,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ApiStages: min=0
	ApiStages []apigatewayusageplan.ApiStages `hcl:"api_stages,block" validate:"min=0"`
	// QuotaSettings: optional
	QuotaSettings *apigatewayusageplan.QuotaSettings `hcl:"quota_settings,block"`
	// ThrottleSettings: optional
	ThrottleSettings *apigatewayusageplan.ThrottleSettings `hcl:"throttle_settings,block"`
}
type apiGatewayUsagePlanAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_api_gateway_usage_plan.
func (agup apiGatewayUsagePlanAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(agup.ref.Append("arn"))
}

// Description returns a reference to field description of aws_api_gateway_usage_plan.
func (agup apiGatewayUsagePlanAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(agup.ref.Append("description"))
}

// Id returns a reference to field id of aws_api_gateway_usage_plan.
func (agup apiGatewayUsagePlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agup.ref.Append("id"))
}

// Name returns a reference to field name of aws_api_gateway_usage_plan.
func (agup apiGatewayUsagePlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(agup.ref.Append("name"))
}

// ProductCode returns a reference to field product_code of aws_api_gateway_usage_plan.
func (agup apiGatewayUsagePlanAttributes) ProductCode() terra.StringValue {
	return terra.ReferenceAsString(agup.ref.Append("product_code"))
}

// Tags returns a reference to field tags of aws_api_gateway_usage_plan.
func (agup apiGatewayUsagePlanAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agup.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_api_gateway_usage_plan.
func (agup apiGatewayUsagePlanAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agup.ref.Append("tags_all"))
}

func (agup apiGatewayUsagePlanAttributes) ApiStages() terra.SetValue[apigatewayusageplan.ApiStagesAttributes] {
	return terra.ReferenceAsSet[apigatewayusageplan.ApiStagesAttributes](agup.ref.Append("api_stages"))
}

func (agup apiGatewayUsagePlanAttributes) QuotaSettings() terra.ListValue[apigatewayusageplan.QuotaSettingsAttributes] {
	return terra.ReferenceAsList[apigatewayusageplan.QuotaSettingsAttributes](agup.ref.Append("quota_settings"))
}

func (agup apiGatewayUsagePlanAttributes) ThrottleSettings() terra.ListValue[apigatewayusageplan.ThrottleSettingsAttributes] {
	return terra.ReferenceAsList[apigatewayusageplan.ThrottleSettingsAttributes](agup.ref.Append("throttle_settings"))
}

type apiGatewayUsagePlanState struct {
	Arn              string                                      `json:"arn"`
	Description      string                                      `json:"description"`
	Id               string                                      `json:"id"`
	Name             string                                      `json:"name"`
	ProductCode      string                                      `json:"product_code"`
	Tags             map[string]string                           `json:"tags"`
	TagsAll          map[string]string                           `json:"tags_all"`
	ApiStages        []apigatewayusageplan.ApiStagesState        `json:"api_stages"`
	QuotaSettings    []apigatewayusageplan.QuotaSettingsState    `json:"quota_settings"`
	ThrottleSettings []apigatewayusageplan.ThrottleSettingsState `json:"throttle_settings"`
}
