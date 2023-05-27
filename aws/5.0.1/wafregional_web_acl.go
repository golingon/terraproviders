// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafregionalwebacl "github.com/golingon/terraproviders/aws/5.0.1/wafregionalwebacl"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafregionalWebAcl creates a new instance of [WafregionalWebAcl].
func NewWafregionalWebAcl(name string, args WafregionalWebAclArgs) *WafregionalWebAcl {
	return &WafregionalWebAcl{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafregionalWebAcl)(nil)

// WafregionalWebAcl represents the Terraform resource aws_wafregional_web_acl.
type WafregionalWebAcl struct {
	Name      string
	Args      WafregionalWebAclArgs
	state     *wafregionalWebAclState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafregionalWebAcl].
func (wwa *WafregionalWebAcl) Type() string {
	return "aws_wafregional_web_acl"
}

// LocalName returns the local name for [WafregionalWebAcl].
func (wwa *WafregionalWebAcl) LocalName() string {
	return wwa.Name
}

// Configuration returns the configuration (args) for [WafregionalWebAcl].
func (wwa *WafregionalWebAcl) Configuration() interface{} {
	return wwa.Args
}

// DependOn is used for other resources to depend on [WafregionalWebAcl].
func (wwa *WafregionalWebAcl) DependOn() terra.Reference {
	return terra.ReferenceResource(wwa)
}

// Dependencies returns the list of resources [WafregionalWebAcl] depends_on.
func (wwa *WafregionalWebAcl) Dependencies() terra.Dependencies {
	return wwa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafregionalWebAcl].
func (wwa *WafregionalWebAcl) LifecycleManagement() *terra.Lifecycle {
	return wwa.Lifecycle
}

// Attributes returns the attributes for [WafregionalWebAcl].
func (wwa *WafregionalWebAcl) Attributes() wafregionalWebAclAttributes {
	return wafregionalWebAclAttributes{ref: terra.ReferenceResource(wwa)}
}

// ImportState imports the given attribute values into [WafregionalWebAcl]'s state.
func (wwa *WafregionalWebAcl) ImportState(av io.Reader) error {
	wwa.state = &wafregionalWebAclState{}
	if err := json.NewDecoder(av).Decode(wwa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wwa.Type(), wwa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafregionalWebAcl] has state.
func (wwa *WafregionalWebAcl) State() (*wafregionalWebAclState, bool) {
	return wwa.state, wwa.state != nil
}

// StateMust returns the state for [WafregionalWebAcl]. Panics if the state is nil.
func (wwa *WafregionalWebAcl) StateMust() *wafregionalWebAclState {
	if wwa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wwa.Type(), wwa.LocalName()))
	}
	return wwa.state
}

// WafregionalWebAclArgs contains the configurations for aws_wafregional_web_acl.
type WafregionalWebAclArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MetricName: string, required
	MetricName terra.StringValue `hcl:"metric_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DefaultAction: required
	DefaultAction *wafregionalwebacl.DefaultAction `hcl:"default_action,block" validate:"required"`
	// LoggingConfiguration: optional
	LoggingConfiguration *wafregionalwebacl.LoggingConfiguration `hcl:"logging_configuration,block"`
	// Rule: min=0
	Rule []wafregionalwebacl.Rule `hcl:"rule,block" validate:"min=0"`
}
type wafregionalWebAclAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_wafregional_web_acl.
func (wwa wafregionalWebAclAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("arn"))
}

// Id returns a reference to field id of aws_wafregional_web_acl.
func (wwa wafregionalWebAclAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("id"))
}

// MetricName returns a reference to field metric_name of aws_wafregional_web_acl.
func (wwa wafregionalWebAclAttributes) MetricName() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("metric_name"))
}

// Name returns a reference to field name of aws_wafregional_web_acl.
func (wwa wafregionalWebAclAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_wafregional_web_acl.
func (wwa wafregionalWebAclAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wwa.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_wafregional_web_acl.
func (wwa wafregionalWebAclAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wwa.ref.Append("tags_all"))
}

func (wwa wafregionalWebAclAttributes) DefaultAction() terra.ListValue[wafregionalwebacl.DefaultActionAttributes] {
	return terra.ReferenceAsList[wafregionalwebacl.DefaultActionAttributes](wwa.ref.Append("default_action"))
}

func (wwa wafregionalWebAclAttributes) LoggingConfiguration() terra.ListValue[wafregionalwebacl.LoggingConfigurationAttributes] {
	return terra.ReferenceAsList[wafregionalwebacl.LoggingConfigurationAttributes](wwa.ref.Append("logging_configuration"))
}

func (wwa wafregionalWebAclAttributes) Rule() terra.SetValue[wafregionalwebacl.RuleAttributes] {
	return terra.ReferenceAsSet[wafregionalwebacl.RuleAttributes](wwa.ref.Append("rule"))
}

type wafregionalWebAclState struct {
	Arn                  string                                        `json:"arn"`
	Id                   string                                        `json:"id"`
	MetricName           string                                        `json:"metric_name"`
	Name                 string                                        `json:"name"`
	Tags                 map[string]string                             `json:"tags"`
	TagsAll              map[string]string                             `json:"tags_all"`
	DefaultAction        []wafregionalwebacl.DefaultActionState        `json:"default_action"`
	LoggingConfiguration []wafregionalwebacl.LoggingConfigurationState `json:"logging_configuration"`
	Rule                 []wafregionalwebacl.RuleState                 `json:"rule"`
}
