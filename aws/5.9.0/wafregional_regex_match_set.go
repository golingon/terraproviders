// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafregionalregexmatchset "github.com/golingon/terraproviders/aws/5.9.0/wafregionalregexmatchset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafregionalRegexMatchSet creates a new instance of [WafregionalRegexMatchSet].
func NewWafregionalRegexMatchSet(name string, args WafregionalRegexMatchSetArgs) *WafregionalRegexMatchSet {
	return &WafregionalRegexMatchSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafregionalRegexMatchSet)(nil)

// WafregionalRegexMatchSet represents the Terraform resource aws_wafregional_regex_match_set.
type WafregionalRegexMatchSet struct {
	Name      string
	Args      WafregionalRegexMatchSetArgs
	state     *wafregionalRegexMatchSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafregionalRegexMatchSet].
func (wrms *WafregionalRegexMatchSet) Type() string {
	return "aws_wafregional_regex_match_set"
}

// LocalName returns the local name for [WafregionalRegexMatchSet].
func (wrms *WafregionalRegexMatchSet) LocalName() string {
	return wrms.Name
}

// Configuration returns the configuration (args) for [WafregionalRegexMatchSet].
func (wrms *WafregionalRegexMatchSet) Configuration() interface{} {
	return wrms.Args
}

// DependOn is used for other resources to depend on [WafregionalRegexMatchSet].
func (wrms *WafregionalRegexMatchSet) DependOn() terra.Reference {
	return terra.ReferenceResource(wrms)
}

// Dependencies returns the list of resources [WafregionalRegexMatchSet] depends_on.
func (wrms *WafregionalRegexMatchSet) Dependencies() terra.Dependencies {
	return wrms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafregionalRegexMatchSet].
func (wrms *WafregionalRegexMatchSet) LifecycleManagement() *terra.Lifecycle {
	return wrms.Lifecycle
}

// Attributes returns the attributes for [WafregionalRegexMatchSet].
func (wrms *WafregionalRegexMatchSet) Attributes() wafregionalRegexMatchSetAttributes {
	return wafregionalRegexMatchSetAttributes{ref: terra.ReferenceResource(wrms)}
}

// ImportState imports the given attribute values into [WafregionalRegexMatchSet]'s state.
func (wrms *WafregionalRegexMatchSet) ImportState(av io.Reader) error {
	wrms.state = &wafregionalRegexMatchSetState{}
	if err := json.NewDecoder(av).Decode(wrms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wrms.Type(), wrms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafregionalRegexMatchSet] has state.
func (wrms *WafregionalRegexMatchSet) State() (*wafregionalRegexMatchSetState, bool) {
	return wrms.state, wrms.state != nil
}

// StateMust returns the state for [WafregionalRegexMatchSet]. Panics if the state is nil.
func (wrms *WafregionalRegexMatchSet) StateMust() *wafregionalRegexMatchSetState {
	if wrms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wrms.Type(), wrms.LocalName()))
	}
	return wrms.state
}

// WafregionalRegexMatchSetArgs contains the configurations for aws_wafregional_regex_match_set.
type WafregionalRegexMatchSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RegexMatchTuple: min=0
	RegexMatchTuple []wafregionalregexmatchset.RegexMatchTuple `hcl:"regex_match_tuple,block" validate:"min=0"`
}
type wafregionalRegexMatchSetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_wafregional_regex_match_set.
func (wrms wafregionalRegexMatchSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wrms.ref.Append("id"))
}

// Name returns a reference to field name of aws_wafregional_regex_match_set.
func (wrms wafregionalRegexMatchSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wrms.ref.Append("name"))
}

func (wrms wafregionalRegexMatchSetAttributes) RegexMatchTuple() terra.SetValue[wafregionalregexmatchset.RegexMatchTupleAttributes] {
	return terra.ReferenceAsSet[wafregionalregexmatchset.RegexMatchTupleAttributes](wrms.ref.Append("regex_match_tuple"))
}

type wafregionalRegexMatchSetState struct {
	Id              string                                          `json:"id"`
	Name            string                                          `json:"name"`
	RegexMatchTuple []wafregionalregexmatchset.RegexMatchTupleState `json:"regex_match_tuple"`
}
