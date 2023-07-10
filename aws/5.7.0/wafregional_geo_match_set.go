// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafregionalgeomatchset "github.com/golingon/terraproviders/aws/5.7.0/wafregionalgeomatchset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafregionalGeoMatchSet creates a new instance of [WafregionalGeoMatchSet].
func NewWafregionalGeoMatchSet(name string, args WafregionalGeoMatchSetArgs) *WafregionalGeoMatchSet {
	return &WafregionalGeoMatchSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafregionalGeoMatchSet)(nil)

// WafregionalGeoMatchSet represents the Terraform resource aws_wafregional_geo_match_set.
type WafregionalGeoMatchSet struct {
	Name      string
	Args      WafregionalGeoMatchSetArgs
	state     *wafregionalGeoMatchSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafregionalGeoMatchSet].
func (wgms *WafregionalGeoMatchSet) Type() string {
	return "aws_wafregional_geo_match_set"
}

// LocalName returns the local name for [WafregionalGeoMatchSet].
func (wgms *WafregionalGeoMatchSet) LocalName() string {
	return wgms.Name
}

// Configuration returns the configuration (args) for [WafregionalGeoMatchSet].
func (wgms *WafregionalGeoMatchSet) Configuration() interface{} {
	return wgms.Args
}

// DependOn is used for other resources to depend on [WafregionalGeoMatchSet].
func (wgms *WafregionalGeoMatchSet) DependOn() terra.Reference {
	return terra.ReferenceResource(wgms)
}

// Dependencies returns the list of resources [WafregionalGeoMatchSet] depends_on.
func (wgms *WafregionalGeoMatchSet) Dependencies() terra.Dependencies {
	return wgms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafregionalGeoMatchSet].
func (wgms *WafregionalGeoMatchSet) LifecycleManagement() *terra.Lifecycle {
	return wgms.Lifecycle
}

// Attributes returns the attributes for [WafregionalGeoMatchSet].
func (wgms *WafregionalGeoMatchSet) Attributes() wafregionalGeoMatchSetAttributes {
	return wafregionalGeoMatchSetAttributes{ref: terra.ReferenceResource(wgms)}
}

// ImportState imports the given attribute values into [WafregionalGeoMatchSet]'s state.
func (wgms *WafregionalGeoMatchSet) ImportState(av io.Reader) error {
	wgms.state = &wafregionalGeoMatchSetState{}
	if err := json.NewDecoder(av).Decode(wgms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wgms.Type(), wgms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafregionalGeoMatchSet] has state.
func (wgms *WafregionalGeoMatchSet) State() (*wafregionalGeoMatchSetState, bool) {
	return wgms.state, wgms.state != nil
}

// StateMust returns the state for [WafregionalGeoMatchSet]. Panics if the state is nil.
func (wgms *WafregionalGeoMatchSet) StateMust() *wafregionalGeoMatchSetState {
	if wgms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wgms.Type(), wgms.LocalName()))
	}
	return wgms.state
}

// WafregionalGeoMatchSetArgs contains the configurations for aws_wafregional_geo_match_set.
type WafregionalGeoMatchSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// GeoMatchConstraint: min=0
	GeoMatchConstraint []wafregionalgeomatchset.GeoMatchConstraint `hcl:"geo_match_constraint,block" validate:"min=0"`
}
type wafregionalGeoMatchSetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_wafregional_geo_match_set.
func (wgms wafregionalGeoMatchSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wgms.ref.Append("id"))
}

// Name returns a reference to field name of aws_wafregional_geo_match_set.
func (wgms wafregionalGeoMatchSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wgms.ref.Append("name"))
}

func (wgms wafregionalGeoMatchSetAttributes) GeoMatchConstraint() terra.SetValue[wafregionalgeomatchset.GeoMatchConstraintAttributes] {
	return terra.ReferenceAsSet[wafregionalgeomatchset.GeoMatchConstraintAttributes](wgms.ref.Append("geo_match_constraint"))
}

type wafregionalGeoMatchSetState struct {
	Id                 string                                           `json:"id"`
	Name               string                                           `json:"name"`
	GeoMatchConstraint []wafregionalgeomatchset.GeoMatchConstraintState `json:"geo_match_constraint"`
}
