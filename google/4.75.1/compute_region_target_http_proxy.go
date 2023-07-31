// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeregiontargethttpproxy "github.com/golingon/terraproviders/google/4.75.1/computeregiontargethttpproxy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionTargetHttpProxy creates a new instance of [ComputeRegionTargetHttpProxy].
func NewComputeRegionTargetHttpProxy(name string, args ComputeRegionTargetHttpProxyArgs) *ComputeRegionTargetHttpProxy {
	return &ComputeRegionTargetHttpProxy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionTargetHttpProxy)(nil)

// ComputeRegionTargetHttpProxy represents the Terraform resource google_compute_region_target_http_proxy.
type ComputeRegionTargetHttpProxy struct {
	Name      string
	Args      ComputeRegionTargetHttpProxyArgs
	state     *computeRegionTargetHttpProxyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionTargetHttpProxy].
func (crthp *ComputeRegionTargetHttpProxy) Type() string {
	return "google_compute_region_target_http_proxy"
}

// LocalName returns the local name for [ComputeRegionTargetHttpProxy].
func (crthp *ComputeRegionTargetHttpProxy) LocalName() string {
	return crthp.Name
}

// Configuration returns the configuration (args) for [ComputeRegionTargetHttpProxy].
func (crthp *ComputeRegionTargetHttpProxy) Configuration() interface{} {
	return crthp.Args
}

// DependOn is used for other resources to depend on [ComputeRegionTargetHttpProxy].
func (crthp *ComputeRegionTargetHttpProxy) DependOn() terra.Reference {
	return terra.ReferenceResource(crthp)
}

// Dependencies returns the list of resources [ComputeRegionTargetHttpProxy] depends_on.
func (crthp *ComputeRegionTargetHttpProxy) Dependencies() terra.Dependencies {
	return crthp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionTargetHttpProxy].
func (crthp *ComputeRegionTargetHttpProxy) LifecycleManagement() *terra.Lifecycle {
	return crthp.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionTargetHttpProxy].
func (crthp *ComputeRegionTargetHttpProxy) Attributes() computeRegionTargetHttpProxyAttributes {
	return computeRegionTargetHttpProxyAttributes{ref: terra.ReferenceResource(crthp)}
}

// ImportState imports the given attribute values into [ComputeRegionTargetHttpProxy]'s state.
func (crthp *ComputeRegionTargetHttpProxy) ImportState(av io.Reader) error {
	crthp.state = &computeRegionTargetHttpProxyState{}
	if err := json.NewDecoder(av).Decode(crthp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crthp.Type(), crthp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionTargetHttpProxy] has state.
func (crthp *ComputeRegionTargetHttpProxy) State() (*computeRegionTargetHttpProxyState, bool) {
	return crthp.state, crthp.state != nil
}

// StateMust returns the state for [ComputeRegionTargetHttpProxy]. Panics if the state is nil.
func (crthp *ComputeRegionTargetHttpProxy) StateMust() *computeRegionTargetHttpProxyState {
	if crthp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crthp.Type(), crthp.LocalName()))
	}
	return crthp.state
}

// ComputeRegionTargetHttpProxyArgs contains the configurations for google_compute_region_target_http_proxy.
type ComputeRegionTargetHttpProxyArgs struct {
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
	// UrlMap: string, required
	UrlMap terra.StringValue `hcl:"url_map,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *computeregiontargethttpproxy.Timeouts `hcl:"timeouts,block"`
}
type computeRegionTargetHttpProxyAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_region_target_http_proxy.
func (crthp computeRegionTargetHttpProxyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(crthp.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_region_target_http_proxy.
func (crthp computeRegionTargetHttpProxyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crthp.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_region_target_http_proxy.
func (crthp computeRegionTargetHttpProxyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crthp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_target_http_proxy.
func (crthp computeRegionTargetHttpProxyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crthp.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_target_http_proxy.
func (crthp computeRegionTargetHttpProxyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crthp.ref.Append("project"))
}

// ProxyId returns a reference to field proxy_id of google_compute_region_target_http_proxy.
func (crthp computeRegionTargetHttpProxyAttributes) ProxyId() terra.NumberValue {
	return terra.ReferenceAsNumber(crthp.ref.Append("proxy_id"))
}

// Region returns a reference to field region of google_compute_region_target_http_proxy.
func (crthp computeRegionTargetHttpProxyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crthp.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_target_http_proxy.
func (crthp computeRegionTargetHttpProxyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crthp.ref.Append("self_link"))
}

// UrlMap returns a reference to field url_map of google_compute_region_target_http_proxy.
func (crthp computeRegionTargetHttpProxyAttributes) UrlMap() terra.StringValue {
	return terra.ReferenceAsString(crthp.ref.Append("url_map"))
}

func (crthp computeRegionTargetHttpProxyAttributes) Timeouts() computeregiontargethttpproxy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregiontargethttpproxy.TimeoutsAttributes](crthp.ref.Append("timeouts"))
}

type computeRegionTargetHttpProxyState struct {
	CreationTimestamp string                                      `json:"creation_timestamp"`
	Description       string                                      `json:"description"`
	Id                string                                      `json:"id"`
	Name              string                                      `json:"name"`
	Project           string                                      `json:"project"`
	ProxyId           float64                                     `json:"proxy_id"`
	Region            string                                      `json:"region"`
	SelfLink          string                                      `json:"self_link"`
	UrlMap            string                                      `json:"url_map"`
	Timeouts          *computeregiontargethttpproxy.TimeoutsState `json:"timeouts"`
}
