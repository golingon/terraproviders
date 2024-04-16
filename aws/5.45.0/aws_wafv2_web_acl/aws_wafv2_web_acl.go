// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_wafv2_web_acl

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

// Resource represents the Terraform resource aws_wafv2_web_acl.
type Resource struct {
	Name      string
	Args      Args
	state     *awsWafv2WebAclState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (awwa *Resource) Type() string {
	return "aws_wafv2_web_acl"
}

// LocalName returns the local name for [Resource].
func (awwa *Resource) LocalName() string {
	return awwa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (awwa *Resource) Configuration() interface{} {
	return awwa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (awwa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(awwa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (awwa *Resource) Dependencies() terra.Dependencies {
	return awwa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (awwa *Resource) LifecycleManagement() *terra.Lifecycle {
	return awwa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (awwa *Resource) Attributes() awsWafv2WebAclAttributes {
	return awsWafv2WebAclAttributes{ref: terra.ReferenceResource(awwa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (awwa *Resource) ImportState(state io.Reader) error {
	awwa.state = &awsWafv2WebAclState{}
	if err := json.NewDecoder(state).Decode(awwa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", awwa.Type(), awwa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (awwa *Resource) State() (*awsWafv2WebAclState, bool) {
	return awwa.state, awwa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (awwa *Resource) StateMust() *awsWafv2WebAclState {
	if awwa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", awwa.Type(), awwa.LocalName()))
	}
	return awwa.state
}

// Args contains the configurations for aws_wafv2_web_acl.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TokenDomains: set of string, optional
	TokenDomains terra.SetValue[terra.StringValue] `hcl:"token_domains,attr"`
	// AssociationConfig: optional
	AssociationConfig *AssociationConfig `hcl:"association_config,block"`
	// CaptchaConfig: optional
	CaptchaConfig *CaptchaConfig `hcl:"captcha_config,block"`
	// ChallengeConfig: optional
	ChallengeConfig *ChallengeConfig `hcl:"challenge_config,block"`
	// CustomResponseBody: min=0
	CustomResponseBody []CustomResponseBody `hcl:"custom_response_body,block" validate:"min=0"`
	// DefaultAction: required
	DefaultAction *DefaultAction `hcl:"default_action,block" validate:"required"`
	// Rule: min=0
	Rule []Rule `hcl:"rule,block" validate:"min=0"`
	// VisibilityConfig: required
	VisibilityConfig *VisibilityConfig `hcl:"visibility_config,block" validate:"required"`
}

type awsWafv2WebAclAttributes struct {
	ref terra.Reference
}

// ApplicationIntegrationUrl returns a reference to field application_integration_url of aws_wafv2_web_acl.
func (awwa awsWafv2WebAclAttributes) ApplicationIntegrationUrl() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("application_integration_url"))
}

// Arn returns a reference to field arn of aws_wafv2_web_acl.
func (awwa awsWafv2WebAclAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("arn"))
}

// Capacity returns a reference to field capacity of aws_wafv2_web_acl.
func (awwa awsWafv2WebAclAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(awwa.ref.Append("capacity"))
}

// Description returns a reference to field description of aws_wafv2_web_acl.
func (awwa awsWafv2WebAclAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("description"))
}

// Id returns a reference to field id of aws_wafv2_web_acl.
func (awwa awsWafv2WebAclAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("id"))
}

// LockToken returns a reference to field lock_token of aws_wafv2_web_acl.
func (awwa awsWafv2WebAclAttributes) LockToken() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("lock_token"))
}

// Name returns a reference to field name of aws_wafv2_web_acl.
func (awwa awsWafv2WebAclAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("name"))
}

// Scope returns a reference to field scope of aws_wafv2_web_acl.
func (awwa awsWafv2WebAclAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(awwa.ref.Append("scope"))
}

// Tags returns a reference to field tags of aws_wafv2_web_acl.
func (awwa awsWafv2WebAclAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](awwa.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_wafv2_web_acl.
func (awwa awsWafv2WebAclAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](awwa.ref.Append("tags_all"))
}

// TokenDomains returns a reference to field token_domains of aws_wafv2_web_acl.
func (awwa awsWafv2WebAclAttributes) TokenDomains() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](awwa.ref.Append("token_domains"))
}

func (awwa awsWafv2WebAclAttributes) AssociationConfig() terra.ListValue[AssociationConfigAttributes] {
	return terra.ReferenceAsList[AssociationConfigAttributes](awwa.ref.Append("association_config"))
}

func (awwa awsWafv2WebAclAttributes) CaptchaConfig() terra.ListValue[CaptchaConfigAttributes] {
	return terra.ReferenceAsList[CaptchaConfigAttributes](awwa.ref.Append("captcha_config"))
}

func (awwa awsWafv2WebAclAttributes) ChallengeConfig() terra.ListValue[ChallengeConfigAttributes] {
	return terra.ReferenceAsList[ChallengeConfigAttributes](awwa.ref.Append("challenge_config"))
}

func (awwa awsWafv2WebAclAttributes) CustomResponseBody() terra.SetValue[CustomResponseBodyAttributes] {
	return terra.ReferenceAsSet[CustomResponseBodyAttributes](awwa.ref.Append("custom_response_body"))
}

func (awwa awsWafv2WebAclAttributes) DefaultAction() terra.ListValue[DefaultActionAttributes] {
	return terra.ReferenceAsList[DefaultActionAttributes](awwa.ref.Append("default_action"))
}

func (awwa awsWafv2WebAclAttributes) Rule() terra.SetValue[RuleAttributes] {
	return terra.ReferenceAsSet[RuleAttributes](awwa.ref.Append("rule"))
}

func (awwa awsWafv2WebAclAttributes) VisibilityConfig() terra.ListValue[VisibilityConfigAttributes] {
	return terra.ReferenceAsList[VisibilityConfigAttributes](awwa.ref.Append("visibility_config"))
}

type awsWafv2WebAclState struct {
	ApplicationIntegrationUrl string                    `json:"application_integration_url"`
	Arn                       string                    `json:"arn"`
	Capacity                  float64                   `json:"capacity"`
	Description               string                    `json:"description"`
	Id                        string                    `json:"id"`
	LockToken                 string                    `json:"lock_token"`
	Name                      string                    `json:"name"`
	Scope                     string                    `json:"scope"`
	Tags                      map[string]string         `json:"tags"`
	TagsAll                   map[string]string         `json:"tags_all"`
	TokenDomains              []string                  `json:"token_domains"`
	AssociationConfig         []AssociationConfigState  `json:"association_config"`
	CaptchaConfig             []CaptchaConfigState      `json:"captcha_config"`
	ChallengeConfig           []ChallengeConfigState    `json:"challenge_config"`
	CustomResponseBody        []CustomResponseBodyState `json:"custom_response_body"`
	DefaultAction             []DefaultActionState      `json:"default_action"`
	Rule                      []RuleState               `json:"rule"`
	VisibilityConfig          []VisibilityConfigState   `json:"visibility_config"`
}
