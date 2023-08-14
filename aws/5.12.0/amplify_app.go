// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	amplifyapp "github.com/golingon/terraproviders/aws/5.12.0/amplifyapp"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAmplifyApp creates a new instance of [AmplifyApp].
func NewAmplifyApp(name string, args AmplifyAppArgs) *AmplifyApp {
	return &AmplifyApp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AmplifyApp)(nil)

// AmplifyApp represents the Terraform resource aws_amplify_app.
type AmplifyApp struct {
	Name      string
	Args      AmplifyAppArgs
	state     *amplifyAppState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AmplifyApp].
func (aa *AmplifyApp) Type() string {
	return "aws_amplify_app"
}

// LocalName returns the local name for [AmplifyApp].
func (aa *AmplifyApp) LocalName() string {
	return aa.Name
}

// Configuration returns the configuration (args) for [AmplifyApp].
func (aa *AmplifyApp) Configuration() interface{} {
	return aa.Args
}

// DependOn is used for other resources to depend on [AmplifyApp].
func (aa *AmplifyApp) DependOn() terra.Reference {
	return terra.ReferenceResource(aa)
}

// Dependencies returns the list of resources [AmplifyApp] depends_on.
func (aa *AmplifyApp) Dependencies() terra.Dependencies {
	return aa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AmplifyApp].
func (aa *AmplifyApp) LifecycleManagement() *terra.Lifecycle {
	return aa.Lifecycle
}

// Attributes returns the attributes for [AmplifyApp].
func (aa *AmplifyApp) Attributes() amplifyAppAttributes {
	return amplifyAppAttributes{ref: terra.ReferenceResource(aa)}
}

// ImportState imports the given attribute values into [AmplifyApp]'s state.
func (aa *AmplifyApp) ImportState(av io.Reader) error {
	aa.state = &amplifyAppState{}
	if err := json.NewDecoder(av).Decode(aa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aa.Type(), aa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AmplifyApp] has state.
func (aa *AmplifyApp) State() (*amplifyAppState, bool) {
	return aa.state, aa.state != nil
}

// StateMust returns the state for [AmplifyApp]. Panics if the state is nil.
func (aa *AmplifyApp) StateMust() *amplifyAppState {
	if aa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aa.Type(), aa.LocalName()))
	}
	return aa.state
}

// AmplifyAppArgs contains the configurations for aws_amplify_app.
type AmplifyAppArgs struct {
	// AccessToken: string, optional
	AccessToken terra.StringValue `hcl:"access_token,attr"`
	// AutoBranchCreationPatterns: set of string, optional
	AutoBranchCreationPatterns terra.SetValue[terra.StringValue] `hcl:"auto_branch_creation_patterns,attr"`
	// BasicAuthCredentials: string, optional
	BasicAuthCredentials terra.StringValue `hcl:"basic_auth_credentials,attr"`
	// BuildSpec: string, optional
	BuildSpec terra.StringValue `hcl:"build_spec,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EnableAutoBranchCreation: bool, optional
	EnableAutoBranchCreation terra.BoolValue `hcl:"enable_auto_branch_creation,attr"`
	// EnableBasicAuth: bool, optional
	EnableBasicAuth terra.BoolValue `hcl:"enable_basic_auth,attr"`
	// EnableBranchAutoBuild: bool, optional
	EnableBranchAutoBuild terra.BoolValue `hcl:"enable_branch_auto_build,attr"`
	// EnableBranchAutoDeletion: bool, optional
	EnableBranchAutoDeletion terra.BoolValue `hcl:"enable_branch_auto_deletion,attr"`
	// EnvironmentVariables: map of string, optional
	EnvironmentVariables terra.MapValue[terra.StringValue] `hcl:"environment_variables,attr"`
	// IamServiceRoleArn: string, optional
	IamServiceRoleArn terra.StringValue `hcl:"iam_service_role_arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OauthToken: string, optional
	OauthToken terra.StringValue `hcl:"oauth_token,attr"`
	// Platform: string, optional
	Platform terra.StringValue `hcl:"platform,attr"`
	// Repository: string, optional
	Repository terra.StringValue `hcl:"repository,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ProductionBranch: min=0
	ProductionBranch []amplifyapp.ProductionBranch `hcl:"production_branch,block" validate:"min=0"`
	// AutoBranchCreationConfig: optional
	AutoBranchCreationConfig *amplifyapp.AutoBranchCreationConfig `hcl:"auto_branch_creation_config,block"`
	// CustomRule: min=0
	CustomRule []amplifyapp.CustomRule `hcl:"custom_rule,block" validate:"min=0"`
}
type amplifyAppAttributes struct {
	ref terra.Reference
}

// AccessToken returns a reference to field access_token of aws_amplify_app.
func (aa amplifyAppAttributes) AccessToken() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("access_token"))
}

// Arn returns a reference to field arn of aws_amplify_app.
func (aa amplifyAppAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("arn"))
}

// AutoBranchCreationPatterns returns a reference to field auto_branch_creation_patterns of aws_amplify_app.
func (aa amplifyAppAttributes) AutoBranchCreationPatterns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aa.ref.Append("auto_branch_creation_patterns"))
}

// BasicAuthCredentials returns a reference to field basic_auth_credentials of aws_amplify_app.
func (aa amplifyAppAttributes) BasicAuthCredentials() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("basic_auth_credentials"))
}

// BuildSpec returns a reference to field build_spec of aws_amplify_app.
func (aa amplifyAppAttributes) BuildSpec() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("build_spec"))
}

// DefaultDomain returns a reference to field default_domain of aws_amplify_app.
func (aa amplifyAppAttributes) DefaultDomain() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("default_domain"))
}

// Description returns a reference to field description of aws_amplify_app.
func (aa amplifyAppAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("description"))
}

// EnableAutoBranchCreation returns a reference to field enable_auto_branch_creation of aws_amplify_app.
func (aa amplifyAppAttributes) EnableAutoBranchCreation() terra.BoolValue {
	return terra.ReferenceAsBool(aa.ref.Append("enable_auto_branch_creation"))
}

// EnableBasicAuth returns a reference to field enable_basic_auth of aws_amplify_app.
func (aa amplifyAppAttributes) EnableBasicAuth() terra.BoolValue {
	return terra.ReferenceAsBool(aa.ref.Append("enable_basic_auth"))
}

// EnableBranchAutoBuild returns a reference to field enable_branch_auto_build of aws_amplify_app.
func (aa amplifyAppAttributes) EnableBranchAutoBuild() terra.BoolValue {
	return terra.ReferenceAsBool(aa.ref.Append("enable_branch_auto_build"))
}

// EnableBranchAutoDeletion returns a reference to field enable_branch_auto_deletion of aws_amplify_app.
func (aa amplifyAppAttributes) EnableBranchAutoDeletion() terra.BoolValue {
	return terra.ReferenceAsBool(aa.ref.Append("enable_branch_auto_deletion"))
}

// EnvironmentVariables returns a reference to field environment_variables of aws_amplify_app.
func (aa amplifyAppAttributes) EnvironmentVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aa.ref.Append("environment_variables"))
}

// IamServiceRoleArn returns a reference to field iam_service_role_arn of aws_amplify_app.
func (aa amplifyAppAttributes) IamServiceRoleArn() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("iam_service_role_arn"))
}

// Id returns a reference to field id of aws_amplify_app.
func (aa amplifyAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("id"))
}

// Name returns a reference to field name of aws_amplify_app.
func (aa amplifyAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("name"))
}

// OauthToken returns a reference to field oauth_token of aws_amplify_app.
func (aa amplifyAppAttributes) OauthToken() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("oauth_token"))
}

// Platform returns a reference to field platform of aws_amplify_app.
func (aa amplifyAppAttributes) Platform() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("platform"))
}

// Repository returns a reference to field repository of aws_amplify_app.
func (aa amplifyAppAttributes) Repository() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("repository"))
}

// Tags returns a reference to field tags of aws_amplify_app.
func (aa amplifyAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aa.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_amplify_app.
func (aa amplifyAppAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aa.ref.Append("tags_all"))
}

func (aa amplifyAppAttributes) ProductionBranch() terra.ListValue[amplifyapp.ProductionBranchAttributes] {
	return terra.ReferenceAsList[amplifyapp.ProductionBranchAttributes](aa.ref.Append("production_branch"))
}

func (aa amplifyAppAttributes) AutoBranchCreationConfig() terra.ListValue[amplifyapp.AutoBranchCreationConfigAttributes] {
	return terra.ReferenceAsList[amplifyapp.AutoBranchCreationConfigAttributes](aa.ref.Append("auto_branch_creation_config"))
}

func (aa amplifyAppAttributes) CustomRule() terra.ListValue[amplifyapp.CustomRuleAttributes] {
	return terra.ReferenceAsList[amplifyapp.CustomRuleAttributes](aa.ref.Append("custom_rule"))
}

type amplifyAppState struct {
	AccessToken                string                                     `json:"access_token"`
	Arn                        string                                     `json:"arn"`
	AutoBranchCreationPatterns []string                                   `json:"auto_branch_creation_patterns"`
	BasicAuthCredentials       string                                     `json:"basic_auth_credentials"`
	BuildSpec                  string                                     `json:"build_spec"`
	DefaultDomain              string                                     `json:"default_domain"`
	Description                string                                     `json:"description"`
	EnableAutoBranchCreation   bool                                       `json:"enable_auto_branch_creation"`
	EnableBasicAuth            bool                                       `json:"enable_basic_auth"`
	EnableBranchAutoBuild      bool                                       `json:"enable_branch_auto_build"`
	EnableBranchAutoDeletion   bool                                       `json:"enable_branch_auto_deletion"`
	EnvironmentVariables       map[string]string                          `json:"environment_variables"`
	IamServiceRoleArn          string                                     `json:"iam_service_role_arn"`
	Id                         string                                     `json:"id"`
	Name                       string                                     `json:"name"`
	OauthToken                 string                                     `json:"oauth_token"`
	Platform                   string                                     `json:"platform"`
	Repository                 string                                     `json:"repository"`
	Tags                       map[string]string                          `json:"tags"`
	TagsAll                    map[string]string                          `json:"tags_all"`
	ProductionBranch           []amplifyapp.ProductionBranchState         `json:"production_branch"`
	AutoBranchCreationConfig   []amplifyapp.AutoBranchCreationConfigState `json:"auto_branch_creation_config"`
	CustomRule                 []amplifyapp.CustomRuleState               `json:"custom_rule"`
}
