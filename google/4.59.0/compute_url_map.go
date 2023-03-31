// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeurlmap "github.com/golingon/terraproviders/google/4.59.0/computeurlmap"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeUrlMap(name string, args ComputeUrlMapArgs) *ComputeUrlMap {
	return &ComputeUrlMap{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeUrlMap)(nil)

type ComputeUrlMap struct {
	Name  string
	Args  ComputeUrlMapArgs
	state *computeUrlMapState
}

func (cum *ComputeUrlMap) Type() string {
	return "google_compute_url_map"
}

func (cum *ComputeUrlMap) LocalName() string {
	return cum.Name
}

func (cum *ComputeUrlMap) Configuration() interface{} {
	return cum.Args
}

func (cum *ComputeUrlMap) Attributes() computeUrlMapAttributes {
	return computeUrlMapAttributes{ref: terra.ReferenceResource(cum)}
}

func (cum *ComputeUrlMap) ImportState(av io.Reader) error {
	cum.state = &computeUrlMapState{}
	if err := json.NewDecoder(av).Decode(cum.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cum.Type(), cum.LocalName(), err)
	}
	return nil
}

func (cum *ComputeUrlMap) State() (*computeUrlMapState, bool) {
	return cum.state, cum.state != nil
}

func (cum *ComputeUrlMap) StateMust() *computeUrlMapState {
	if cum.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cum.Type(), cum.LocalName()))
	}
	return cum.state
}

func (cum *ComputeUrlMap) DependOn() terra.Reference {
	return terra.ReferenceResource(cum)
}

type ComputeUrlMapArgs struct {
	// DefaultService: string, optional
	DefaultService terra.StringValue `hcl:"default_service,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// DefaultRouteAction: optional
	DefaultRouteAction *computeurlmap.DefaultRouteAction `hcl:"default_route_action,block"`
	// DefaultUrlRedirect: optional
	DefaultUrlRedirect *computeurlmap.DefaultUrlRedirect `hcl:"default_url_redirect,block"`
	// HeaderAction: optional
	HeaderAction *computeurlmap.HeaderAction `hcl:"header_action,block"`
	// HostRule: min=0
	HostRule []computeurlmap.HostRule `hcl:"host_rule,block" validate:"min=0"`
	// PathMatcher: min=0
	PathMatcher []computeurlmap.PathMatcher `hcl:"path_matcher,block" validate:"min=0"`
	// Test: min=0
	Test []computeurlmap.Test `hcl:"test,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *computeurlmap.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ComputeUrlMap depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeUrlMapAttributes struct {
	ref terra.Reference
}

func (cum computeUrlMapAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceString(cum.ref.Append("creation_timestamp"))
}

func (cum computeUrlMapAttributes) DefaultService() terra.StringValue {
	return terra.ReferenceString(cum.ref.Append("default_service"))
}

func (cum computeUrlMapAttributes) Description() terra.StringValue {
	return terra.ReferenceString(cum.ref.Append("description"))
}

func (cum computeUrlMapAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceString(cum.ref.Append("fingerprint"))
}

func (cum computeUrlMapAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cum.ref.Append("id"))
}

func (cum computeUrlMapAttributes) MapId() terra.NumberValue {
	return terra.ReferenceNumber(cum.ref.Append("map_id"))
}

func (cum computeUrlMapAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cum.ref.Append("name"))
}

func (cum computeUrlMapAttributes) Project() terra.StringValue {
	return terra.ReferenceString(cum.ref.Append("project"))
}

func (cum computeUrlMapAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(cum.ref.Append("self_link"))
}

func (cum computeUrlMapAttributes) DefaultRouteAction() terra.ListValue[computeurlmap.DefaultRouteActionAttributes] {
	return terra.ReferenceList[computeurlmap.DefaultRouteActionAttributes](cum.ref.Append("default_route_action"))
}

func (cum computeUrlMapAttributes) DefaultUrlRedirect() terra.ListValue[computeurlmap.DefaultUrlRedirectAttributes] {
	return terra.ReferenceList[computeurlmap.DefaultUrlRedirectAttributes](cum.ref.Append("default_url_redirect"))
}

func (cum computeUrlMapAttributes) HeaderAction() terra.ListValue[computeurlmap.HeaderActionAttributes] {
	return terra.ReferenceList[computeurlmap.HeaderActionAttributes](cum.ref.Append("header_action"))
}

func (cum computeUrlMapAttributes) HostRule() terra.SetValue[computeurlmap.HostRuleAttributes] {
	return terra.ReferenceSet[computeurlmap.HostRuleAttributes](cum.ref.Append("host_rule"))
}

func (cum computeUrlMapAttributes) PathMatcher() terra.ListValue[computeurlmap.PathMatcherAttributes] {
	return terra.ReferenceList[computeurlmap.PathMatcherAttributes](cum.ref.Append("path_matcher"))
}

func (cum computeUrlMapAttributes) Test() terra.ListValue[computeurlmap.TestAttributes] {
	return terra.ReferenceList[computeurlmap.TestAttributes](cum.ref.Append("test"))
}

func (cum computeUrlMapAttributes) Timeouts() computeurlmap.TimeoutsAttributes {
	return terra.ReferenceSingle[computeurlmap.TimeoutsAttributes](cum.ref.Append("timeouts"))
}

type computeUrlMapState struct {
	CreationTimestamp  string                                  `json:"creation_timestamp"`
	DefaultService     string                                  `json:"default_service"`
	Description        string                                  `json:"description"`
	Fingerprint        string                                  `json:"fingerprint"`
	Id                 string                                  `json:"id"`
	MapId              float64                                 `json:"map_id"`
	Name               string                                  `json:"name"`
	Project            string                                  `json:"project"`
	SelfLink           string                                  `json:"self_link"`
	DefaultRouteAction []computeurlmap.DefaultRouteActionState `json:"default_route_action"`
	DefaultUrlRedirect []computeurlmap.DefaultUrlRedirectState `json:"default_url_redirect"`
	HeaderAction       []computeurlmap.HeaderActionState       `json:"header_action"`
	HostRule           []computeurlmap.HostRuleState           `json:"host_rule"`
	PathMatcher        []computeurlmap.PathMatcherState        `json:"path_matcher"`
	Test               []computeurlmap.TestState               `json:"test"`
	Timeouts           *computeurlmap.TimeoutsState            `json:"timeouts"`
}
