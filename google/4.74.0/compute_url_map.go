// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeurlmap "github.com/golingon/terraproviders/google/4.74.0/computeurlmap"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeUrlMap creates a new instance of [ComputeUrlMap].
func NewComputeUrlMap(name string, args ComputeUrlMapArgs) *ComputeUrlMap {
	return &ComputeUrlMap{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeUrlMap)(nil)

// ComputeUrlMap represents the Terraform resource google_compute_url_map.
type ComputeUrlMap struct {
	Name      string
	Args      ComputeUrlMapArgs
	state     *computeUrlMapState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeUrlMap].
func (cum *ComputeUrlMap) Type() string {
	return "google_compute_url_map"
}

// LocalName returns the local name for [ComputeUrlMap].
func (cum *ComputeUrlMap) LocalName() string {
	return cum.Name
}

// Configuration returns the configuration (args) for [ComputeUrlMap].
func (cum *ComputeUrlMap) Configuration() interface{} {
	return cum.Args
}

// DependOn is used for other resources to depend on [ComputeUrlMap].
func (cum *ComputeUrlMap) DependOn() terra.Reference {
	return terra.ReferenceResource(cum)
}

// Dependencies returns the list of resources [ComputeUrlMap] depends_on.
func (cum *ComputeUrlMap) Dependencies() terra.Dependencies {
	return cum.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeUrlMap].
func (cum *ComputeUrlMap) LifecycleManagement() *terra.Lifecycle {
	return cum.Lifecycle
}

// Attributes returns the attributes for [ComputeUrlMap].
func (cum *ComputeUrlMap) Attributes() computeUrlMapAttributes {
	return computeUrlMapAttributes{ref: terra.ReferenceResource(cum)}
}

// ImportState imports the given attribute values into [ComputeUrlMap]'s state.
func (cum *ComputeUrlMap) ImportState(av io.Reader) error {
	cum.state = &computeUrlMapState{}
	if err := json.NewDecoder(av).Decode(cum.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cum.Type(), cum.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeUrlMap] has state.
func (cum *ComputeUrlMap) State() (*computeUrlMapState, bool) {
	return cum.state, cum.state != nil
}

// StateMust returns the state for [ComputeUrlMap]. Panics if the state is nil.
func (cum *ComputeUrlMap) StateMust() *computeUrlMapState {
	if cum.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cum.Type(), cum.LocalName()))
	}
	return cum.state
}

// ComputeUrlMapArgs contains the configurations for google_compute_url_map.
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
}
type computeUrlMapAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_url_map.
func (cum computeUrlMapAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cum.ref.Append("creation_timestamp"))
}

// DefaultService returns a reference to field default_service of google_compute_url_map.
func (cum computeUrlMapAttributes) DefaultService() terra.StringValue {
	return terra.ReferenceAsString(cum.ref.Append("default_service"))
}

// Description returns a reference to field description of google_compute_url_map.
func (cum computeUrlMapAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cum.ref.Append("description"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_url_map.
func (cum computeUrlMapAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(cum.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_compute_url_map.
func (cum computeUrlMapAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cum.ref.Append("id"))
}

// MapId returns a reference to field map_id of google_compute_url_map.
func (cum computeUrlMapAttributes) MapId() terra.NumberValue {
	return terra.ReferenceAsNumber(cum.ref.Append("map_id"))
}

// Name returns a reference to field name of google_compute_url_map.
func (cum computeUrlMapAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cum.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_url_map.
func (cum computeUrlMapAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cum.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_url_map.
func (cum computeUrlMapAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cum.ref.Append("self_link"))
}

func (cum computeUrlMapAttributes) DefaultRouteAction() terra.ListValue[computeurlmap.DefaultRouteActionAttributes] {
	return terra.ReferenceAsList[computeurlmap.DefaultRouteActionAttributes](cum.ref.Append("default_route_action"))
}

func (cum computeUrlMapAttributes) DefaultUrlRedirect() terra.ListValue[computeurlmap.DefaultUrlRedirectAttributes] {
	return terra.ReferenceAsList[computeurlmap.DefaultUrlRedirectAttributes](cum.ref.Append("default_url_redirect"))
}

func (cum computeUrlMapAttributes) HeaderAction() terra.ListValue[computeurlmap.HeaderActionAttributes] {
	return terra.ReferenceAsList[computeurlmap.HeaderActionAttributes](cum.ref.Append("header_action"))
}

func (cum computeUrlMapAttributes) HostRule() terra.SetValue[computeurlmap.HostRuleAttributes] {
	return terra.ReferenceAsSet[computeurlmap.HostRuleAttributes](cum.ref.Append("host_rule"))
}

func (cum computeUrlMapAttributes) PathMatcher() terra.ListValue[computeurlmap.PathMatcherAttributes] {
	return terra.ReferenceAsList[computeurlmap.PathMatcherAttributes](cum.ref.Append("path_matcher"))
}

func (cum computeUrlMapAttributes) Test() terra.ListValue[computeurlmap.TestAttributes] {
	return terra.ReferenceAsList[computeurlmap.TestAttributes](cum.ref.Append("test"))
}

func (cum computeUrlMapAttributes) Timeouts() computeurlmap.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeurlmap.TimeoutsAttributes](cum.ref.Append("timeouts"))
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
