// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeregionurlmap "github.com/golingon/terraproviders/googlebeta/4.66.0/computeregionurlmap"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionUrlMap creates a new instance of [ComputeRegionUrlMap].
func NewComputeRegionUrlMap(name string, args ComputeRegionUrlMapArgs) *ComputeRegionUrlMap {
	return &ComputeRegionUrlMap{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionUrlMap)(nil)

// ComputeRegionUrlMap represents the Terraform resource google_compute_region_url_map.
type ComputeRegionUrlMap struct {
	Name      string
	Args      ComputeRegionUrlMapArgs
	state     *computeRegionUrlMapState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionUrlMap].
func (crum *ComputeRegionUrlMap) Type() string {
	return "google_compute_region_url_map"
}

// LocalName returns the local name for [ComputeRegionUrlMap].
func (crum *ComputeRegionUrlMap) LocalName() string {
	return crum.Name
}

// Configuration returns the configuration (args) for [ComputeRegionUrlMap].
func (crum *ComputeRegionUrlMap) Configuration() interface{} {
	return crum.Args
}

// DependOn is used for other resources to depend on [ComputeRegionUrlMap].
func (crum *ComputeRegionUrlMap) DependOn() terra.Reference {
	return terra.ReferenceResource(crum)
}

// Dependencies returns the list of resources [ComputeRegionUrlMap] depends_on.
func (crum *ComputeRegionUrlMap) Dependencies() terra.Dependencies {
	return crum.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionUrlMap].
func (crum *ComputeRegionUrlMap) LifecycleManagement() *terra.Lifecycle {
	return crum.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionUrlMap].
func (crum *ComputeRegionUrlMap) Attributes() computeRegionUrlMapAttributes {
	return computeRegionUrlMapAttributes{ref: terra.ReferenceResource(crum)}
}

// ImportState imports the given attribute values into [ComputeRegionUrlMap]'s state.
func (crum *ComputeRegionUrlMap) ImportState(av io.Reader) error {
	crum.state = &computeRegionUrlMapState{}
	if err := json.NewDecoder(av).Decode(crum.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crum.Type(), crum.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionUrlMap] has state.
func (crum *ComputeRegionUrlMap) State() (*computeRegionUrlMapState, bool) {
	return crum.state, crum.state != nil
}

// StateMust returns the state for [ComputeRegionUrlMap]. Panics if the state is nil.
func (crum *ComputeRegionUrlMap) StateMust() *computeRegionUrlMapState {
	if crum.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crum.Type(), crum.LocalName()))
	}
	return crum.state
}

// ComputeRegionUrlMapArgs contains the configurations for google_compute_region_url_map.
type ComputeRegionUrlMapArgs struct {
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
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// DefaultRouteAction: optional
	DefaultRouteAction *computeregionurlmap.DefaultRouteAction `hcl:"default_route_action,block"`
	// DefaultUrlRedirect: optional
	DefaultUrlRedirect *computeregionurlmap.DefaultUrlRedirect `hcl:"default_url_redirect,block"`
	// HostRule: min=0
	HostRule []computeregionurlmap.HostRule `hcl:"host_rule,block" validate:"min=0"`
	// PathMatcher: min=0
	PathMatcher []computeregionurlmap.PathMatcher `hcl:"path_matcher,block" validate:"min=0"`
	// Test: min=0
	Test []computeregionurlmap.Test `hcl:"test,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *computeregionurlmap.Timeouts `hcl:"timeouts,block"`
}
type computeRegionUrlMapAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_region_url_map.
func (crum computeRegionUrlMapAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(crum.ref.Append("creation_timestamp"))
}

// DefaultService returns a reference to field default_service of google_compute_region_url_map.
func (crum computeRegionUrlMapAttributes) DefaultService() terra.StringValue {
	return terra.ReferenceAsString(crum.ref.Append("default_service"))
}

// Description returns a reference to field description of google_compute_region_url_map.
func (crum computeRegionUrlMapAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crum.ref.Append("description"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_region_url_map.
func (crum computeRegionUrlMapAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(crum.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_compute_region_url_map.
func (crum computeRegionUrlMapAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crum.ref.Append("id"))
}

// MapId returns a reference to field map_id of google_compute_region_url_map.
func (crum computeRegionUrlMapAttributes) MapId() terra.NumberValue {
	return terra.ReferenceAsNumber(crum.ref.Append("map_id"))
}

// Name returns a reference to field name of google_compute_region_url_map.
func (crum computeRegionUrlMapAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crum.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_url_map.
func (crum computeRegionUrlMapAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crum.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_url_map.
func (crum computeRegionUrlMapAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crum.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_url_map.
func (crum computeRegionUrlMapAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crum.ref.Append("self_link"))
}

func (crum computeRegionUrlMapAttributes) DefaultRouteAction() terra.ListValue[computeregionurlmap.DefaultRouteActionAttributes] {
	return terra.ReferenceAsList[computeregionurlmap.DefaultRouteActionAttributes](crum.ref.Append("default_route_action"))
}

func (crum computeRegionUrlMapAttributes) DefaultUrlRedirect() terra.ListValue[computeregionurlmap.DefaultUrlRedirectAttributes] {
	return terra.ReferenceAsList[computeregionurlmap.DefaultUrlRedirectAttributes](crum.ref.Append("default_url_redirect"))
}

func (crum computeRegionUrlMapAttributes) HostRule() terra.SetValue[computeregionurlmap.HostRuleAttributes] {
	return terra.ReferenceAsSet[computeregionurlmap.HostRuleAttributes](crum.ref.Append("host_rule"))
}

func (crum computeRegionUrlMapAttributes) PathMatcher() terra.ListValue[computeregionurlmap.PathMatcherAttributes] {
	return terra.ReferenceAsList[computeregionurlmap.PathMatcherAttributes](crum.ref.Append("path_matcher"))
}

func (crum computeRegionUrlMapAttributes) Test() terra.ListValue[computeregionurlmap.TestAttributes] {
	return terra.ReferenceAsList[computeregionurlmap.TestAttributes](crum.ref.Append("test"))
}

func (crum computeRegionUrlMapAttributes) Timeouts() computeregionurlmap.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregionurlmap.TimeoutsAttributes](crum.ref.Append("timeouts"))
}

type computeRegionUrlMapState struct {
	CreationTimestamp  string                                        `json:"creation_timestamp"`
	DefaultService     string                                        `json:"default_service"`
	Description        string                                        `json:"description"`
	Fingerprint        string                                        `json:"fingerprint"`
	Id                 string                                        `json:"id"`
	MapId              float64                                       `json:"map_id"`
	Name               string                                        `json:"name"`
	Project            string                                        `json:"project"`
	Region             string                                        `json:"region"`
	SelfLink           string                                        `json:"self_link"`
	DefaultRouteAction []computeregionurlmap.DefaultRouteActionState `json:"default_route_action"`
	DefaultUrlRedirect []computeregionurlmap.DefaultUrlRedirectState `json:"default_url_redirect"`
	HostRule           []computeregionurlmap.HostRuleState           `json:"host_rule"`
	PathMatcher        []computeregionurlmap.PathMatcherState        `json:"path_matcher"`
	Test               []computeregionurlmap.TestState               `json:"test"`
	Timeouts           *computeregionurlmap.TimeoutsState            `json:"timeouts"`
}
