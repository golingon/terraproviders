// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafv2regexpatternset "github.com/golingon/terraproviders/aws/4.60.0/wafv2regexpatternset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafv2RegexPatternSet creates a new instance of [Wafv2RegexPatternSet].
func NewWafv2RegexPatternSet(name string, args Wafv2RegexPatternSetArgs) *Wafv2RegexPatternSet {
	return &Wafv2RegexPatternSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Wafv2RegexPatternSet)(nil)

// Wafv2RegexPatternSet represents the Terraform resource aws_wafv2_regex_pattern_set.
type Wafv2RegexPatternSet struct {
	Name      string
	Args      Wafv2RegexPatternSetArgs
	state     *wafv2RegexPatternSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Wafv2RegexPatternSet].
func (wrps *Wafv2RegexPatternSet) Type() string {
	return "aws_wafv2_regex_pattern_set"
}

// LocalName returns the local name for [Wafv2RegexPatternSet].
func (wrps *Wafv2RegexPatternSet) LocalName() string {
	return wrps.Name
}

// Configuration returns the configuration (args) for [Wafv2RegexPatternSet].
func (wrps *Wafv2RegexPatternSet) Configuration() interface{} {
	return wrps.Args
}

// DependOn is used for other resources to depend on [Wafv2RegexPatternSet].
func (wrps *Wafv2RegexPatternSet) DependOn() terra.Reference {
	return terra.ReferenceResource(wrps)
}

// Dependencies returns the list of resources [Wafv2RegexPatternSet] depends_on.
func (wrps *Wafv2RegexPatternSet) Dependencies() terra.Dependencies {
	return wrps.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Wafv2RegexPatternSet].
func (wrps *Wafv2RegexPatternSet) LifecycleManagement() *terra.Lifecycle {
	return wrps.Lifecycle
}

// Attributes returns the attributes for [Wafv2RegexPatternSet].
func (wrps *Wafv2RegexPatternSet) Attributes() wafv2RegexPatternSetAttributes {
	return wafv2RegexPatternSetAttributes{ref: terra.ReferenceResource(wrps)}
}

// ImportState imports the given attribute values into [Wafv2RegexPatternSet]'s state.
func (wrps *Wafv2RegexPatternSet) ImportState(av io.Reader) error {
	wrps.state = &wafv2RegexPatternSetState{}
	if err := json.NewDecoder(av).Decode(wrps.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wrps.Type(), wrps.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Wafv2RegexPatternSet] has state.
func (wrps *Wafv2RegexPatternSet) State() (*wafv2RegexPatternSetState, bool) {
	return wrps.state, wrps.state != nil
}

// StateMust returns the state for [Wafv2RegexPatternSet]. Panics if the state is nil.
func (wrps *Wafv2RegexPatternSet) StateMust() *wafv2RegexPatternSetState {
	if wrps.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wrps.Type(), wrps.LocalName()))
	}
	return wrps.state
}

// Wafv2RegexPatternSetArgs contains the configurations for aws_wafv2_regex_pattern_set.
type Wafv2RegexPatternSetArgs struct {
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
	// RegularExpression: min=0,max=10
	RegularExpression []wafv2regexpatternset.RegularExpression `hcl:"regular_expression,block" validate:"min=0,max=10"`
}
type wafv2RegexPatternSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_wafv2_regex_pattern_set.
func (wrps wafv2RegexPatternSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wrps.ref.Append("arn"))
}

// Description returns a reference to field description of aws_wafv2_regex_pattern_set.
func (wrps wafv2RegexPatternSetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(wrps.ref.Append("description"))
}

// Id returns a reference to field id of aws_wafv2_regex_pattern_set.
func (wrps wafv2RegexPatternSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wrps.ref.Append("id"))
}

// LockToken returns a reference to field lock_token of aws_wafv2_regex_pattern_set.
func (wrps wafv2RegexPatternSetAttributes) LockToken() terra.StringValue {
	return terra.ReferenceAsString(wrps.ref.Append("lock_token"))
}

// Name returns a reference to field name of aws_wafv2_regex_pattern_set.
func (wrps wafv2RegexPatternSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wrps.ref.Append("name"))
}

// Scope returns a reference to field scope of aws_wafv2_regex_pattern_set.
func (wrps wafv2RegexPatternSetAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(wrps.ref.Append("scope"))
}

// Tags returns a reference to field tags of aws_wafv2_regex_pattern_set.
func (wrps wafv2RegexPatternSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wrps.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_wafv2_regex_pattern_set.
func (wrps wafv2RegexPatternSetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wrps.ref.Append("tags_all"))
}

func (wrps wafv2RegexPatternSetAttributes) RegularExpression() terra.SetValue[wafv2regexpatternset.RegularExpressionAttributes] {
	return terra.ReferenceAsSet[wafv2regexpatternset.RegularExpressionAttributes](wrps.ref.Append("regular_expression"))
}

type wafv2RegexPatternSetState struct {
	Arn               string                                        `json:"arn"`
	Description       string                                        `json:"description"`
	Id                string                                        `json:"id"`
	LockToken         string                                        `json:"lock_token"`
	Name              string                                        `json:"name"`
	Scope             string                                        `json:"scope"`
	Tags              map[string]string                             `json:"tags"`
	TagsAll           map[string]string                             `json:"tags_all"`
	RegularExpression []wafv2regexpatternset.RegularExpressionState `json:"regular_expression"`
}
