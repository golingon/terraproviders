// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewWafregionalRegexPatternSet(name string, args WafregionalRegexPatternSetArgs) *WafregionalRegexPatternSet {
	return &WafregionalRegexPatternSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafregionalRegexPatternSet)(nil)

type WafregionalRegexPatternSet struct {
	Name  string
	Args  WafregionalRegexPatternSetArgs
	state *wafregionalRegexPatternSetState
}

func (wrps *WafregionalRegexPatternSet) Type() string {
	return "aws_wafregional_regex_pattern_set"
}

func (wrps *WafregionalRegexPatternSet) LocalName() string {
	return wrps.Name
}

func (wrps *WafregionalRegexPatternSet) Configuration() interface{} {
	return wrps.Args
}

func (wrps *WafregionalRegexPatternSet) Attributes() wafregionalRegexPatternSetAttributes {
	return wafregionalRegexPatternSetAttributes{ref: terra.ReferenceResource(wrps)}
}

func (wrps *WafregionalRegexPatternSet) ImportState(av io.Reader) error {
	wrps.state = &wafregionalRegexPatternSetState{}
	if err := json.NewDecoder(av).Decode(wrps.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wrps.Type(), wrps.LocalName(), err)
	}
	return nil
}

func (wrps *WafregionalRegexPatternSet) State() (*wafregionalRegexPatternSetState, bool) {
	return wrps.state, wrps.state != nil
}

func (wrps *WafregionalRegexPatternSet) StateMust() *wafregionalRegexPatternSetState {
	if wrps.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wrps.Type(), wrps.LocalName()))
	}
	return wrps.state
}

func (wrps *WafregionalRegexPatternSet) DependOn() terra.Reference {
	return terra.ReferenceResource(wrps)
}

type WafregionalRegexPatternSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RegexPatternStrings: set of string, optional
	RegexPatternStrings terra.SetValue[terra.StringValue] `hcl:"regex_pattern_strings,attr"`
	// DependsOn contains resources that WafregionalRegexPatternSet depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type wafregionalRegexPatternSetAttributes struct {
	ref terra.Reference
}

func (wrps wafregionalRegexPatternSetAttributes) Id() terra.StringValue {
	return terra.ReferenceString(wrps.ref.Append("id"))
}

func (wrps wafregionalRegexPatternSetAttributes) Name() terra.StringValue {
	return terra.ReferenceString(wrps.ref.Append("name"))
}

func (wrps wafregionalRegexPatternSetAttributes) RegexPatternStrings() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](wrps.ref.Append("regex_pattern_strings"))
}

type wafregionalRegexPatternSetState struct {
	Id                  string   `json:"id"`
	Name                string   `json:"name"`
	RegexPatternStrings []string `json:"regex_pattern_strings"`
}
