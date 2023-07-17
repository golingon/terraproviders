// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafregexmatchset "github.com/golingon/terraproviders/aws/5.8.0/wafregexmatchset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafRegexMatchSet creates a new instance of [WafRegexMatchSet].
func NewWafRegexMatchSet(name string, args WafRegexMatchSetArgs) *WafRegexMatchSet {
	return &WafRegexMatchSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafRegexMatchSet)(nil)

// WafRegexMatchSet represents the Terraform resource aws_waf_regex_match_set.
type WafRegexMatchSet struct {
	Name      string
	Args      WafRegexMatchSetArgs
	state     *wafRegexMatchSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafRegexMatchSet].
func (wrms *WafRegexMatchSet) Type() string {
	return "aws_waf_regex_match_set"
}

// LocalName returns the local name for [WafRegexMatchSet].
func (wrms *WafRegexMatchSet) LocalName() string {
	return wrms.Name
}

// Configuration returns the configuration (args) for [WafRegexMatchSet].
func (wrms *WafRegexMatchSet) Configuration() interface{} {
	return wrms.Args
}

// DependOn is used for other resources to depend on [WafRegexMatchSet].
func (wrms *WafRegexMatchSet) DependOn() terra.Reference {
	return terra.ReferenceResource(wrms)
}

// Dependencies returns the list of resources [WafRegexMatchSet] depends_on.
func (wrms *WafRegexMatchSet) Dependencies() terra.Dependencies {
	return wrms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafRegexMatchSet].
func (wrms *WafRegexMatchSet) LifecycleManagement() *terra.Lifecycle {
	return wrms.Lifecycle
}

// Attributes returns the attributes for [WafRegexMatchSet].
func (wrms *WafRegexMatchSet) Attributes() wafRegexMatchSetAttributes {
	return wafRegexMatchSetAttributes{ref: terra.ReferenceResource(wrms)}
}

// ImportState imports the given attribute values into [WafRegexMatchSet]'s state.
func (wrms *WafRegexMatchSet) ImportState(av io.Reader) error {
	wrms.state = &wafRegexMatchSetState{}
	if err := json.NewDecoder(av).Decode(wrms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wrms.Type(), wrms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafRegexMatchSet] has state.
func (wrms *WafRegexMatchSet) State() (*wafRegexMatchSetState, bool) {
	return wrms.state, wrms.state != nil
}

// StateMust returns the state for [WafRegexMatchSet]. Panics if the state is nil.
func (wrms *WafRegexMatchSet) StateMust() *wafRegexMatchSetState {
	if wrms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wrms.Type(), wrms.LocalName()))
	}
	return wrms.state
}

// WafRegexMatchSetArgs contains the configurations for aws_waf_regex_match_set.
type WafRegexMatchSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RegexMatchTuple: min=0
	RegexMatchTuple []wafregexmatchset.RegexMatchTuple `hcl:"regex_match_tuple,block" validate:"min=0"`
}
type wafRegexMatchSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_waf_regex_match_set.
func (wrms wafRegexMatchSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wrms.ref.Append("arn"))
}

// Id returns a reference to field id of aws_waf_regex_match_set.
func (wrms wafRegexMatchSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wrms.ref.Append("id"))
}

// Name returns a reference to field name of aws_waf_regex_match_set.
func (wrms wafRegexMatchSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wrms.ref.Append("name"))
}

func (wrms wafRegexMatchSetAttributes) RegexMatchTuple() terra.SetValue[wafregexmatchset.RegexMatchTupleAttributes] {
	return terra.ReferenceAsSet[wafregexmatchset.RegexMatchTupleAttributes](wrms.ref.Append("regex_match_tuple"))
}

type wafRegexMatchSetState struct {
	Arn             string                                  `json:"arn"`
	Id              string                                  `json:"id"`
	Name            string                                  `json:"name"`
	RegexMatchTuple []wafregexmatchset.RegexMatchTupleState `json:"regex_match_tuple"`
}
