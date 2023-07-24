// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafv2webacl "github.com/golingon/terraproviders/aws/5.9.0/wafv2webacl"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafv2WebAcl creates a new instance of [Wafv2WebAcl].
func NewWafv2WebAcl(name string, args Wafv2WebAclArgs) *Wafv2WebAcl {
	return &Wafv2WebAcl{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Wafv2WebAcl)(nil)

// Wafv2WebAcl represents the Terraform resource aws_wafv2_web_acl.
type Wafv2WebAcl struct {
	Name      string
	Args      Wafv2WebAclArgs
	state     *wafv2WebAclState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Wafv2WebAcl].
func (wwa *Wafv2WebAcl) Type() string {
	return "aws_wafv2_web_acl"
}

// LocalName returns the local name for [Wafv2WebAcl].
func (wwa *Wafv2WebAcl) LocalName() string {
	return wwa.Name
}

// Configuration returns the configuration (args) for [Wafv2WebAcl].
func (wwa *Wafv2WebAcl) Configuration() interface{} {
	return wwa.Args
}

// DependOn is used for other resources to depend on [Wafv2WebAcl].
func (wwa *Wafv2WebAcl) DependOn() terra.Reference {
	return terra.ReferenceResource(wwa)
}

// Dependencies returns the list of resources [Wafv2WebAcl] depends_on.
func (wwa *Wafv2WebAcl) Dependencies() terra.Dependencies {
	return wwa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Wafv2WebAcl].
func (wwa *Wafv2WebAcl) LifecycleManagement() *terra.Lifecycle {
	return wwa.Lifecycle
}

// Attributes returns the attributes for [Wafv2WebAcl].
func (wwa *Wafv2WebAcl) Attributes() wafv2WebAclAttributes {
	return wafv2WebAclAttributes{ref: terra.ReferenceResource(wwa)}
}

// ImportState imports the given attribute values into [Wafv2WebAcl]'s state.
func (wwa *Wafv2WebAcl) ImportState(av io.Reader) error {
	wwa.state = &wafv2WebAclState{}
	if err := json.NewDecoder(av).Decode(wwa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wwa.Type(), wwa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Wafv2WebAcl] has state.
func (wwa *Wafv2WebAcl) State() (*wafv2WebAclState, bool) {
	return wwa.state, wwa.state != nil
}

// StateMust returns the state for [Wafv2WebAcl]. Panics if the state is nil.
func (wwa *Wafv2WebAcl) StateMust() *wafv2WebAclState {
	if wwa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wwa.Type(), wwa.LocalName()))
	}
	return wwa.state
}

// Wafv2WebAclArgs contains the configurations for aws_wafv2_web_acl.
type Wafv2WebAclArgs struct {
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
	AssociationConfig *wafv2webacl.AssociationConfig `hcl:"association_config,block"`
	// CaptchaConfig: optional
	CaptchaConfig *wafv2webacl.CaptchaConfig `hcl:"captcha_config,block"`
	// CustomResponseBody: min=0
	CustomResponseBody []wafv2webacl.CustomResponseBody `hcl:"custom_response_body,block" validate:"min=0"`
	// DefaultAction: required
	DefaultAction *wafv2webacl.DefaultAction `hcl:"default_action,block" validate:"required"`
	// Rule: min=0
	Rule []wafv2webacl.Rule `hcl:"rule,block" validate:"min=0"`
	// VisibilityConfig: required
	VisibilityConfig *wafv2webacl.VisibilityConfig `hcl:"visibility_config,block" validate:"required"`
}
type wafv2WebAclAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_wafv2_web_acl.
func (wwa wafv2WebAclAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("arn"))
}

// Capacity returns a reference to field capacity of aws_wafv2_web_acl.
func (wwa wafv2WebAclAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(wwa.ref.Append("capacity"))
}

// Description returns a reference to field description of aws_wafv2_web_acl.
func (wwa wafv2WebAclAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("description"))
}

// Id returns a reference to field id of aws_wafv2_web_acl.
func (wwa wafv2WebAclAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("id"))
}

// LockToken returns a reference to field lock_token of aws_wafv2_web_acl.
func (wwa wafv2WebAclAttributes) LockToken() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("lock_token"))
}

// Name returns a reference to field name of aws_wafv2_web_acl.
func (wwa wafv2WebAclAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("name"))
}

// Scope returns a reference to field scope of aws_wafv2_web_acl.
func (wwa wafv2WebAclAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("scope"))
}

// Tags returns a reference to field tags of aws_wafv2_web_acl.
func (wwa wafv2WebAclAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wwa.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_wafv2_web_acl.
func (wwa wafv2WebAclAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wwa.ref.Append("tags_all"))
}

// TokenDomains returns a reference to field token_domains of aws_wafv2_web_acl.
func (wwa wafv2WebAclAttributes) TokenDomains() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](wwa.ref.Append("token_domains"))
}

func (wwa wafv2WebAclAttributes) AssociationConfig() terra.ListValue[wafv2webacl.AssociationConfigAttributes] {
	return terra.ReferenceAsList[wafv2webacl.AssociationConfigAttributes](wwa.ref.Append("association_config"))
}

func (wwa wafv2WebAclAttributes) CaptchaConfig() terra.ListValue[wafv2webacl.CaptchaConfigAttributes] {
	return terra.ReferenceAsList[wafv2webacl.CaptchaConfigAttributes](wwa.ref.Append("captcha_config"))
}

func (wwa wafv2WebAclAttributes) CustomResponseBody() terra.SetValue[wafv2webacl.CustomResponseBodyAttributes] {
	return terra.ReferenceAsSet[wafv2webacl.CustomResponseBodyAttributes](wwa.ref.Append("custom_response_body"))
}

func (wwa wafv2WebAclAttributes) DefaultAction() terra.ListValue[wafv2webacl.DefaultActionAttributes] {
	return terra.ReferenceAsList[wafv2webacl.DefaultActionAttributes](wwa.ref.Append("default_action"))
}

func (wwa wafv2WebAclAttributes) Rule() terra.SetValue[wafv2webacl.RuleAttributes] {
	return terra.ReferenceAsSet[wafv2webacl.RuleAttributes](wwa.ref.Append("rule"))
}

func (wwa wafv2WebAclAttributes) VisibilityConfig() terra.ListValue[wafv2webacl.VisibilityConfigAttributes] {
	return terra.ReferenceAsList[wafv2webacl.VisibilityConfigAttributes](wwa.ref.Append("visibility_config"))
}

type wafv2WebAclState struct {
	Arn                string                                `json:"arn"`
	Capacity           float64                               `json:"capacity"`
	Description        string                                `json:"description"`
	Id                 string                                `json:"id"`
	LockToken          string                                `json:"lock_token"`
	Name               string                                `json:"name"`
	Scope              string                                `json:"scope"`
	Tags               map[string]string                     `json:"tags"`
	TagsAll            map[string]string                     `json:"tags_all"`
	TokenDomains       []string                              `json:"token_domains"`
	AssociationConfig  []wafv2webacl.AssociationConfigState  `json:"association_config"`
	CaptchaConfig      []wafv2webacl.CaptchaConfigState      `json:"captcha_config"`
	CustomResponseBody []wafv2webacl.CustomResponseBodyState `json:"custom_response_body"`
	DefaultAction      []wafv2webacl.DefaultActionState      `json:"default_action"`
	Rule               []wafv2webacl.RuleState               `json:"rule"`
	VisibilityConfig   []wafv2webacl.VisibilityConfigState   `json:"visibility_config"`
}
