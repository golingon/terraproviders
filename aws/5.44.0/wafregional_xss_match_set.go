// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	wafregionalxssmatchset "github.com/golingon/terraproviders/aws/5.44.0/wafregionalxssmatchset"
	"io"
)

// NewWafregionalXssMatchSet creates a new instance of [WafregionalXssMatchSet].
func NewWafregionalXssMatchSet(name string, args WafregionalXssMatchSetArgs) *WafregionalXssMatchSet {
	return &WafregionalXssMatchSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafregionalXssMatchSet)(nil)

// WafregionalXssMatchSet represents the Terraform resource aws_wafregional_xss_match_set.
type WafregionalXssMatchSet struct {
	Name      string
	Args      WafregionalXssMatchSetArgs
	state     *wafregionalXssMatchSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafregionalXssMatchSet].
func (wxms *WafregionalXssMatchSet) Type() string {
	return "aws_wafregional_xss_match_set"
}

// LocalName returns the local name for [WafregionalXssMatchSet].
func (wxms *WafregionalXssMatchSet) LocalName() string {
	return wxms.Name
}

// Configuration returns the configuration (args) for [WafregionalXssMatchSet].
func (wxms *WafregionalXssMatchSet) Configuration() interface{} {
	return wxms.Args
}

// DependOn is used for other resources to depend on [WafregionalXssMatchSet].
func (wxms *WafregionalXssMatchSet) DependOn() terra.Reference {
	return terra.ReferenceResource(wxms)
}

// Dependencies returns the list of resources [WafregionalXssMatchSet] depends_on.
func (wxms *WafregionalXssMatchSet) Dependencies() terra.Dependencies {
	return wxms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafregionalXssMatchSet].
func (wxms *WafregionalXssMatchSet) LifecycleManagement() *terra.Lifecycle {
	return wxms.Lifecycle
}

// Attributes returns the attributes for [WafregionalXssMatchSet].
func (wxms *WafregionalXssMatchSet) Attributes() wafregionalXssMatchSetAttributes {
	return wafregionalXssMatchSetAttributes{ref: terra.ReferenceResource(wxms)}
}

// ImportState imports the given attribute values into [WafregionalXssMatchSet]'s state.
func (wxms *WafregionalXssMatchSet) ImportState(av io.Reader) error {
	wxms.state = &wafregionalXssMatchSetState{}
	if err := json.NewDecoder(av).Decode(wxms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wxms.Type(), wxms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafregionalXssMatchSet] has state.
func (wxms *WafregionalXssMatchSet) State() (*wafregionalXssMatchSetState, bool) {
	return wxms.state, wxms.state != nil
}

// StateMust returns the state for [WafregionalXssMatchSet]. Panics if the state is nil.
func (wxms *WafregionalXssMatchSet) StateMust() *wafregionalXssMatchSetState {
	if wxms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wxms.Type(), wxms.LocalName()))
	}
	return wxms.state
}

// WafregionalXssMatchSetArgs contains the configurations for aws_wafregional_xss_match_set.
type WafregionalXssMatchSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// XssMatchTuple: min=0
	XssMatchTuple []wafregionalxssmatchset.XssMatchTuple `hcl:"xss_match_tuple,block" validate:"min=0"`
}
type wafregionalXssMatchSetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_wafregional_xss_match_set.
func (wxms wafregionalXssMatchSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wxms.ref.Append("id"))
}

// Name returns a reference to field name of aws_wafregional_xss_match_set.
func (wxms wafregionalXssMatchSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wxms.ref.Append("name"))
}

func (wxms wafregionalXssMatchSetAttributes) XssMatchTuple() terra.SetValue[wafregionalxssmatchset.XssMatchTupleAttributes] {
	return terra.ReferenceAsSet[wafregionalxssmatchset.XssMatchTupleAttributes](wxms.ref.Append("xss_match_tuple"))
}

type wafregionalXssMatchSetState struct {
	Id            string                                      `json:"id"`
	Name          string                                      `json:"name"`
	XssMatchTuple []wafregionalxssmatchset.XssMatchTupleState `json:"xss_match_tuple"`
}
