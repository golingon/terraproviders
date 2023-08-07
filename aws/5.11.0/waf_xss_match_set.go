// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafxssmatchset "github.com/golingon/terraproviders/aws/5.11.0/wafxssmatchset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafXssMatchSet creates a new instance of [WafXssMatchSet].
func NewWafXssMatchSet(name string, args WafXssMatchSetArgs) *WafXssMatchSet {
	return &WafXssMatchSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafXssMatchSet)(nil)

// WafXssMatchSet represents the Terraform resource aws_waf_xss_match_set.
type WafXssMatchSet struct {
	Name      string
	Args      WafXssMatchSetArgs
	state     *wafXssMatchSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafXssMatchSet].
func (wxms *WafXssMatchSet) Type() string {
	return "aws_waf_xss_match_set"
}

// LocalName returns the local name for [WafXssMatchSet].
func (wxms *WafXssMatchSet) LocalName() string {
	return wxms.Name
}

// Configuration returns the configuration (args) for [WafXssMatchSet].
func (wxms *WafXssMatchSet) Configuration() interface{} {
	return wxms.Args
}

// DependOn is used for other resources to depend on [WafXssMatchSet].
func (wxms *WafXssMatchSet) DependOn() terra.Reference {
	return terra.ReferenceResource(wxms)
}

// Dependencies returns the list of resources [WafXssMatchSet] depends_on.
func (wxms *WafXssMatchSet) Dependencies() terra.Dependencies {
	return wxms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafXssMatchSet].
func (wxms *WafXssMatchSet) LifecycleManagement() *terra.Lifecycle {
	return wxms.Lifecycle
}

// Attributes returns the attributes for [WafXssMatchSet].
func (wxms *WafXssMatchSet) Attributes() wafXssMatchSetAttributes {
	return wafXssMatchSetAttributes{ref: terra.ReferenceResource(wxms)}
}

// ImportState imports the given attribute values into [WafXssMatchSet]'s state.
func (wxms *WafXssMatchSet) ImportState(av io.Reader) error {
	wxms.state = &wafXssMatchSetState{}
	if err := json.NewDecoder(av).Decode(wxms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wxms.Type(), wxms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafXssMatchSet] has state.
func (wxms *WafXssMatchSet) State() (*wafXssMatchSetState, bool) {
	return wxms.state, wxms.state != nil
}

// StateMust returns the state for [WafXssMatchSet]. Panics if the state is nil.
func (wxms *WafXssMatchSet) StateMust() *wafXssMatchSetState {
	if wxms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wxms.Type(), wxms.LocalName()))
	}
	return wxms.state
}

// WafXssMatchSetArgs contains the configurations for aws_waf_xss_match_set.
type WafXssMatchSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// XssMatchTuples: min=0
	XssMatchTuples []wafxssmatchset.XssMatchTuples `hcl:"xss_match_tuples,block" validate:"min=0"`
}
type wafXssMatchSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_waf_xss_match_set.
func (wxms wafXssMatchSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wxms.ref.Append("arn"))
}

// Id returns a reference to field id of aws_waf_xss_match_set.
func (wxms wafXssMatchSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wxms.ref.Append("id"))
}

// Name returns a reference to field name of aws_waf_xss_match_set.
func (wxms wafXssMatchSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wxms.ref.Append("name"))
}

func (wxms wafXssMatchSetAttributes) XssMatchTuples() terra.SetValue[wafxssmatchset.XssMatchTuplesAttributes] {
	return terra.ReferenceAsSet[wafxssmatchset.XssMatchTuplesAttributes](wxms.ref.Append("xss_match_tuples"))
}

type wafXssMatchSetState struct {
	Arn            string                               `json:"arn"`
	Id             string                               `json:"id"`
	Name           string                               `json:"name"`
	XssMatchTuples []wafxssmatchset.XssMatchTuplesState `json:"xss_match_tuples"`
}
