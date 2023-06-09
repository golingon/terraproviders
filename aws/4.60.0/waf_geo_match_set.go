// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafgeomatchset "github.com/golingon/terraproviders/aws/4.60.0/wafgeomatchset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafGeoMatchSet creates a new instance of [WafGeoMatchSet].
func NewWafGeoMatchSet(name string, args WafGeoMatchSetArgs) *WafGeoMatchSet {
	return &WafGeoMatchSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafGeoMatchSet)(nil)

// WafGeoMatchSet represents the Terraform resource aws_waf_geo_match_set.
type WafGeoMatchSet struct {
	Name      string
	Args      WafGeoMatchSetArgs
	state     *wafGeoMatchSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafGeoMatchSet].
func (wgms *WafGeoMatchSet) Type() string {
	return "aws_waf_geo_match_set"
}

// LocalName returns the local name for [WafGeoMatchSet].
func (wgms *WafGeoMatchSet) LocalName() string {
	return wgms.Name
}

// Configuration returns the configuration (args) for [WafGeoMatchSet].
func (wgms *WafGeoMatchSet) Configuration() interface{} {
	return wgms.Args
}

// DependOn is used for other resources to depend on [WafGeoMatchSet].
func (wgms *WafGeoMatchSet) DependOn() terra.Reference {
	return terra.ReferenceResource(wgms)
}

// Dependencies returns the list of resources [WafGeoMatchSet] depends_on.
func (wgms *WafGeoMatchSet) Dependencies() terra.Dependencies {
	return wgms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafGeoMatchSet].
func (wgms *WafGeoMatchSet) LifecycleManagement() *terra.Lifecycle {
	return wgms.Lifecycle
}

// Attributes returns the attributes for [WafGeoMatchSet].
func (wgms *WafGeoMatchSet) Attributes() wafGeoMatchSetAttributes {
	return wafGeoMatchSetAttributes{ref: terra.ReferenceResource(wgms)}
}

// ImportState imports the given attribute values into [WafGeoMatchSet]'s state.
func (wgms *WafGeoMatchSet) ImportState(av io.Reader) error {
	wgms.state = &wafGeoMatchSetState{}
	if err := json.NewDecoder(av).Decode(wgms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wgms.Type(), wgms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafGeoMatchSet] has state.
func (wgms *WafGeoMatchSet) State() (*wafGeoMatchSetState, bool) {
	return wgms.state, wgms.state != nil
}

// StateMust returns the state for [WafGeoMatchSet]. Panics if the state is nil.
func (wgms *WafGeoMatchSet) StateMust() *wafGeoMatchSetState {
	if wgms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wgms.Type(), wgms.LocalName()))
	}
	return wgms.state
}

// WafGeoMatchSetArgs contains the configurations for aws_waf_geo_match_set.
type WafGeoMatchSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// GeoMatchConstraint: min=0
	GeoMatchConstraint []wafgeomatchset.GeoMatchConstraint `hcl:"geo_match_constraint,block" validate:"min=0"`
}
type wafGeoMatchSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_waf_geo_match_set.
func (wgms wafGeoMatchSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wgms.ref.Append("arn"))
}

// Id returns a reference to field id of aws_waf_geo_match_set.
func (wgms wafGeoMatchSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wgms.ref.Append("id"))
}

// Name returns a reference to field name of aws_waf_geo_match_set.
func (wgms wafGeoMatchSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wgms.ref.Append("name"))
}

func (wgms wafGeoMatchSetAttributes) GeoMatchConstraint() terra.SetValue[wafgeomatchset.GeoMatchConstraintAttributes] {
	return terra.ReferenceAsSet[wafgeomatchset.GeoMatchConstraintAttributes](wgms.ref.Append("geo_match_constraint"))
}

type wafGeoMatchSetState struct {
	Arn                string                                   `json:"arn"`
	Id                 string                                   `json:"id"`
	Name               string                                   `json:"name"`
	GeoMatchConstraint []wafgeomatchset.GeoMatchConstraintState `json:"geo_match_constraint"`
}
