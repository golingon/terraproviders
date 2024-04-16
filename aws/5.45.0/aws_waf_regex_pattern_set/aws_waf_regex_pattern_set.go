// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_waf_regex_pattern_set

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

// Resource represents the Terraform resource aws_waf_regex_pattern_set.
type Resource struct {
	Name      string
	Args      Args
	state     *awsWafRegexPatternSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (awrps *Resource) Type() string {
	return "aws_waf_regex_pattern_set"
}

// LocalName returns the local name for [Resource].
func (awrps *Resource) LocalName() string {
	return awrps.Name
}

// Configuration returns the configuration (args) for [Resource].
func (awrps *Resource) Configuration() interface{} {
	return awrps.Args
}

// DependOn is used for other resources to depend on [Resource].
func (awrps *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(awrps)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (awrps *Resource) Dependencies() terra.Dependencies {
	return awrps.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (awrps *Resource) LifecycleManagement() *terra.Lifecycle {
	return awrps.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (awrps *Resource) Attributes() awsWafRegexPatternSetAttributes {
	return awsWafRegexPatternSetAttributes{ref: terra.ReferenceResource(awrps)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (awrps *Resource) ImportState(state io.Reader) error {
	awrps.state = &awsWafRegexPatternSetState{}
	if err := json.NewDecoder(state).Decode(awrps.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", awrps.Type(), awrps.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (awrps *Resource) State() (*awsWafRegexPatternSetState, bool) {
	return awrps.state, awrps.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (awrps *Resource) StateMust() *awsWafRegexPatternSetState {
	if awrps.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", awrps.Type(), awrps.LocalName()))
	}
	return awrps.state
}

// Args contains the configurations for aws_waf_regex_pattern_set.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RegexPatternStrings: set of string, optional
	RegexPatternStrings terra.SetValue[terra.StringValue] `hcl:"regex_pattern_strings,attr"`
}

type awsWafRegexPatternSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_waf_regex_pattern_set.
func (awrps awsWafRegexPatternSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(awrps.ref.Append("arn"))
}

// Id returns a reference to field id of aws_waf_regex_pattern_set.
func (awrps awsWafRegexPatternSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(awrps.ref.Append("id"))
}

// Name returns a reference to field name of aws_waf_regex_pattern_set.
func (awrps awsWafRegexPatternSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(awrps.ref.Append("name"))
}

// RegexPatternStrings returns a reference to field regex_pattern_strings of aws_waf_regex_pattern_set.
func (awrps awsWafRegexPatternSetAttributes) RegexPatternStrings() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](awrps.ref.Append("regex_pattern_strings"))
}

type awsWafRegexPatternSetState struct {
	Arn                 string   `json:"arn"`
	Id                  string   `json:"id"`
	Name                string   `json:"name"`
	RegexPatternStrings []string `json:"regex_pattern_strings"`
}
