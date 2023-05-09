// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appsyncgraphqlapi "github.com/golingon/terraproviders/aws/4.66.1/appsyncgraphqlapi"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppsyncGraphqlApi creates a new instance of [AppsyncGraphqlApi].
func NewAppsyncGraphqlApi(name string, args AppsyncGraphqlApiArgs) *AppsyncGraphqlApi {
	return &AppsyncGraphqlApi{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppsyncGraphqlApi)(nil)

// AppsyncGraphqlApi represents the Terraform resource aws_appsync_graphql_api.
type AppsyncGraphqlApi struct {
	Name      string
	Args      AppsyncGraphqlApiArgs
	state     *appsyncGraphqlApiState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppsyncGraphqlApi].
func (aga *AppsyncGraphqlApi) Type() string {
	return "aws_appsync_graphql_api"
}

// LocalName returns the local name for [AppsyncGraphqlApi].
func (aga *AppsyncGraphqlApi) LocalName() string {
	return aga.Name
}

// Configuration returns the configuration (args) for [AppsyncGraphqlApi].
func (aga *AppsyncGraphqlApi) Configuration() interface{} {
	return aga.Args
}

// DependOn is used for other resources to depend on [AppsyncGraphqlApi].
func (aga *AppsyncGraphqlApi) DependOn() terra.Reference {
	return terra.ReferenceResource(aga)
}

// Dependencies returns the list of resources [AppsyncGraphqlApi] depends_on.
func (aga *AppsyncGraphqlApi) Dependencies() terra.Dependencies {
	return aga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppsyncGraphqlApi].
func (aga *AppsyncGraphqlApi) LifecycleManagement() *terra.Lifecycle {
	return aga.Lifecycle
}

// Attributes returns the attributes for [AppsyncGraphqlApi].
func (aga *AppsyncGraphqlApi) Attributes() appsyncGraphqlApiAttributes {
	return appsyncGraphqlApiAttributes{ref: terra.ReferenceResource(aga)}
}

// ImportState imports the given attribute values into [AppsyncGraphqlApi]'s state.
func (aga *AppsyncGraphqlApi) ImportState(av io.Reader) error {
	aga.state = &appsyncGraphqlApiState{}
	if err := json.NewDecoder(av).Decode(aga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aga.Type(), aga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppsyncGraphqlApi] has state.
func (aga *AppsyncGraphqlApi) State() (*appsyncGraphqlApiState, bool) {
	return aga.state, aga.state != nil
}

// StateMust returns the state for [AppsyncGraphqlApi]. Panics if the state is nil.
func (aga *AppsyncGraphqlApi) StateMust() *appsyncGraphqlApiState {
	if aga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aga.Type(), aga.LocalName()))
	}
	return aga.state
}

// AppsyncGraphqlApiArgs contains the configurations for aws_appsync_graphql_api.
type AppsyncGraphqlApiArgs struct {
	// AuthenticationType: string, required
	AuthenticationType terra.StringValue `hcl:"authentication_type,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Schema: string, optional
	Schema terra.StringValue `hcl:"schema,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// XrayEnabled: bool, optional
	XrayEnabled terra.BoolValue `hcl:"xray_enabled,attr"`
	// AdditionalAuthenticationProvider: min=0
	AdditionalAuthenticationProvider []appsyncgraphqlapi.AdditionalAuthenticationProvider `hcl:"additional_authentication_provider,block" validate:"min=0"`
	// LambdaAuthorizerConfig: optional
	LambdaAuthorizerConfig *appsyncgraphqlapi.LambdaAuthorizerConfig `hcl:"lambda_authorizer_config,block"`
	// LogConfig: optional
	LogConfig *appsyncgraphqlapi.LogConfig `hcl:"log_config,block"`
	// OpenidConnectConfig: optional
	OpenidConnectConfig *appsyncgraphqlapi.OpenidConnectConfig `hcl:"openid_connect_config,block"`
	// UserPoolConfig: optional
	UserPoolConfig *appsyncgraphqlapi.UserPoolConfig `hcl:"user_pool_config,block"`
}
type appsyncGraphqlApiAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appsync_graphql_api.
func (aga appsyncGraphqlApiAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("arn"))
}

// AuthenticationType returns a reference to field authentication_type of aws_appsync_graphql_api.
func (aga appsyncGraphqlApiAttributes) AuthenticationType() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("authentication_type"))
}

// Id returns a reference to field id of aws_appsync_graphql_api.
func (aga appsyncGraphqlApiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("id"))
}

// Name returns a reference to field name of aws_appsync_graphql_api.
func (aga appsyncGraphqlApiAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("name"))
}

// Schema returns a reference to field schema of aws_appsync_graphql_api.
func (aga appsyncGraphqlApiAttributes) Schema() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("schema"))
}

// Tags returns a reference to field tags of aws_appsync_graphql_api.
func (aga appsyncGraphqlApiAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aga.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appsync_graphql_api.
func (aga appsyncGraphqlApiAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aga.ref.Append("tags_all"))
}

// Uris returns a reference to field uris of aws_appsync_graphql_api.
func (aga appsyncGraphqlApiAttributes) Uris() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aga.ref.Append("uris"))
}

// XrayEnabled returns a reference to field xray_enabled of aws_appsync_graphql_api.
func (aga appsyncGraphqlApiAttributes) XrayEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aga.ref.Append("xray_enabled"))
}

func (aga appsyncGraphqlApiAttributes) AdditionalAuthenticationProvider() terra.ListValue[appsyncgraphqlapi.AdditionalAuthenticationProviderAttributes] {
	return terra.ReferenceAsList[appsyncgraphqlapi.AdditionalAuthenticationProviderAttributes](aga.ref.Append("additional_authentication_provider"))
}

func (aga appsyncGraphqlApiAttributes) LambdaAuthorizerConfig() terra.ListValue[appsyncgraphqlapi.LambdaAuthorizerConfigAttributes] {
	return terra.ReferenceAsList[appsyncgraphqlapi.LambdaAuthorizerConfigAttributes](aga.ref.Append("lambda_authorizer_config"))
}

func (aga appsyncGraphqlApiAttributes) LogConfig() terra.ListValue[appsyncgraphqlapi.LogConfigAttributes] {
	return terra.ReferenceAsList[appsyncgraphqlapi.LogConfigAttributes](aga.ref.Append("log_config"))
}

func (aga appsyncGraphqlApiAttributes) OpenidConnectConfig() terra.ListValue[appsyncgraphqlapi.OpenidConnectConfigAttributes] {
	return terra.ReferenceAsList[appsyncgraphqlapi.OpenidConnectConfigAttributes](aga.ref.Append("openid_connect_config"))
}

func (aga appsyncGraphqlApiAttributes) UserPoolConfig() terra.ListValue[appsyncgraphqlapi.UserPoolConfigAttributes] {
	return terra.ReferenceAsList[appsyncgraphqlapi.UserPoolConfigAttributes](aga.ref.Append("user_pool_config"))
}

type appsyncGraphqlApiState struct {
	Arn                              string                                                    `json:"arn"`
	AuthenticationType               string                                                    `json:"authentication_type"`
	Id                               string                                                    `json:"id"`
	Name                             string                                                    `json:"name"`
	Schema                           string                                                    `json:"schema"`
	Tags                             map[string]string                                         `json:"tags"`
	TagsAll                          map[string]string                                         `json:"tags_all"`
	Uris                             map[string]string                                         `json:"uris"`
	XrayEnabled                      bool                                                      `json:"xray_enabled"`
	AdditionalAuthenticationProvider []appsyncgraphqlapi.AdditionalAuthenticationProviderState `json:"additional_authentication_provider"`
	LambdaAuthorizerConfig           []appsyncgraphqlapi.LambdaAuthorizerConfigState           `json:"lambda_authorizer_config"`
	LogConfig                        []appsyncgraphqlapi.LogConfigState                        `json:"log_config"`
	OpenidConnectConfig              []appsyncgraphqlapi.OpenidConnectConfigState              `json:"openid_connect_config"`
	UserPoolConfig                   []appsyncgraphqlapi.UserPoolConfigState                   `json:"user_pool_config"`
}
