// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafregionalgeomatchset "github.com/golingon/terraproviders/aws/4.60.0/wafregionalgeomatchset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewWafregionalGeoMatchSet(name string, args WafregionalGeoMatchSetArgs) *WafregionalGeoMatchSet {
	return &WafregionalGeoMatchSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafregionalGeoMatchSet)(nil)

type WafregionalGeoMatchSet struct {
	Name  string
	Args  WafregionalGeoMatchSetArgs
	state *wafregionalGeoMatchSetState
}

func (wgms *WafregionalGeoMatchSet) Type() string {
	return "aws_wafregional_geo_match_set"
}

func (wgms *WafregionalGeoMatchSet) LocalName() string {
	return wgms.Name
}

func (wgms *WafregionalGeoMatchSet) Configuration() interface{} {
	return wgms.Args
}

func (wgms *WafregionalGeoMatchSet) Attributes() wafregionalGeoMatchSetAttributes {
	return wafregionalGeoMatchSetAttributes{ref: terra.ReferenceResource(wgms)}
}

func (wgms *WafregionalGeoMatchSet) ImportState(av io.Reader) error {
	wgms.state = &wafregionalGeoMatchSetState{}
	if err := json.NewDecoder(av).Decode(wgms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wgms.Type(), wgms.LocalName(), err)
	}
	return nil
}

func (wgms *WafregionalGeoMatchSet) State() (*wafregionalGeoMatchSetState, bool) {
	return wgms.state, wgms.state != nil
}

func (wgms *WafregionalGeoMatchSet) StateMust() *wafregionalGeoMatchSetState {
	if wgms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wgms.Type(), wgms.LocalName()))
	}
	return wgms.state
}

func (wgms *WafregionalGeoMatchSet) DependOn() terra.Reference {
	return terra.ReferenceResource(wgms)
}

type WafregionalGeoMatchSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// GeoMatchConstraint: min=0
	GeoMatchConstraint []wafregionalgeomatchset.GeoMatchConstraint `hcl:"geo_match_constraint,block" validate:"min=0"`
	// DependsOn contains resources that WafregionalGeoMatchSet depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type wafregionalGeoMatchSetAttributes struct {
	ref terra.Reference
}

func (wgms wafregionalGeoMatchSetAttributes) Id() terra.StringValue {
	return terra.ReferenceString(wgms.ref.Append("id"))
}

func (wgms wafregionalGeoMatchSetAttributes) Name() terra.StringValue {
	return terra.ReferenceString(wgms.ref.Append("name"))
}

func (wgms wafregionalGeoMatchSetAttributes) GeoMatchConstraint() terra.SetValue[wafregionalgeomatchset.GeoMatchConstraintAttributes] {
	return terra.ReferenceSet[wafregionalgeomatchset.GeoMatchConstraintAttributes](wgms.ref.Append("geo_match_constraint"))
}

type wafregionalGeoMatchSetState struct {
	Id                 string                                           `json:"id"`
	Name               string                                           `json:"name"`
	GeoMatchConstraint []wafregionalgeomatchset.GeoMatchConstraintState `json:"geo_match_constraint"`
}
