// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeregiontargethttpsproxy "github.com/golingon/terraproviders/google/4.73.1/computeregiontargethttpsproxy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionTargetHttpsProxy creates a new instance of [ComputeRegionTargetHttpsProxy].
func NewComputeRegionTargetHttpsProxy(name string, args ComputeRegionTargetHttpsProxyArgs) *ComputeRegionTargetHttpsProxy {
	return &ComputeRegionTargetHttpsProxy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionTargetHttpsProxy)(nil)

// ComputeRegionTargetHttpsProxy represents the Terraform resource google_compute_region_target_https_proxy.
type ComputeRegionTargetHttpsProxy struct {
	Name      string
	Args      ComputeRegionTargetHttpsProxyArgs
	state     *computeRegionTargetHttpsProxyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionTargetHttpsProxy].
func (crthp *ComputeRegionTargetHttpsProxy) Type() string {
	return "google_compute_region_target_https_proxy"
}

// LocalName returns the local name for [ComputeRegionTargetHttpsProxy].
func (crthp *ComputeRegionTargetHttpsProxy) LocalName() string {
	return crthp.Name
}

// Configuration returns the configuration (args) for [ComputeRegionTargetHttpsProxy].
func (crthp *ComputeRegionTargetHttpsProxy) Configuration() interface{} {
	return crthp.Args
}

// DependOn is used for other resources to depend on [ComputeRegionTargetHttpsProxy].
func (crthp *ComputeRegionTargetHttpsProxy) DependOn() terra.Reference {
	return terra.ReferenceResource(crthp)
}

// Dependencies returns the list of resources [ComputeRegionTargetHttpsProxy] depends_on.
func (crthp *ComputeRegionTargetHttpsProxy) Dependencies() terra.Dependencies {
	return crthp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionTargetHttpsProxy].
func (crthp *ComputeRegionTargetHttpsProxy) LifecycleManagement() *terra.Lifecycle {
	return crthp.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionTargetHttpsProxy].
func (crthp *ComputeRegionTargetHttpsProxy) Attributes() computeRegionTargetHttpsProxyAttributes {
	return computeRegionTargetHttpsProxyAttributes{ref: terra.ReferenceResource(crthp)}
}

// ImportState imports the given attribute values into [ComputeRegionTargetHttpsProxy]'s state.
func (crthp *ComputeRegionTargetHttpsProxy) ImportState(av io.Reader) error {
	crthp.state = &computeRegionTargetHttpsProxyState{}
	if err := json.NewDecoder(av).Decode(crthp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crthp.Type(), crthp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionTargetHttpsProxy] has state.
func (crthp *ComputeRegionTargetHttpsProxy) State() (*computeRegionTargetHttpsProxyState, bool) {
	return crthp.state, crthp.state != nil
}

// StateMust returns the state for [ComputeRegionTargetHttpsProxy]. Panics if the state is nil.
func (crthp *ComputeRegionTargetHttpsProxy) StateMust() *computeRegionTargetHttpsProxyState {
	if crthp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crthp.Type(), crthp.LocalName()))
	}
	return crthp.state
}

// ComputeRegionTargetHttpsProxyArgs contains the configurations for google_compute_region_target_https_proxy.
type ComputeRegionTargetHttpsProxyArgs struct {
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
	// SslCertificates: list of string, required
	SslCertificates terra.ListValue[terra.StringValue] `hcl:"ssl_certificates,attr" validate:"required"`
	// UrlMap: string, required
	UrlMap terra.StringValue `hcl:"url_map,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *computeregiontargethttpsproxy.Timeouts `hcl:"timeouts,block"`
}
type computeRegionTargetHttpsProxyAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_region_target_https_proxy.
func (crthp computeRegionTargetHttpsProxyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(crthp.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_region_target_https_proxy.
func (crthp computeRegionTargetHttpsProxyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crthp.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_region_target_https_proxy.
func (crthp computeRegionTargetHttpsProxyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crthp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_target_https_proxy.
func (crthp computeRegionTargetHttpsProxyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crthp.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_target_https_proxy.
func (crthp computeRegionTargetHttpsProxyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crthp.ref.Append("project"))
}

// ProxyId returns a reference to field proxy_id of google_compute_region_target_https_proxy.
func (crthp computeRegionTargetHttpsProxyAttributes) ProxyId() terra.NumberValue {
	return terra.ReferenceAsNumber(crthp.ref.Append("proxy_id"))
}

// Region returns a reference to field region of google_compute_region_target_https_proxy.
func (crthp computeRegionTargetHttpsProxyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crthp.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_target_https_proxy.
func (crthp computeRegionTargetHttpsProxyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crthp.ref.Append("self_link"))
}

// SslCertificates returns a reference to field ssl_certificates of google_compute_region_target_https_proxy.
func (crthp computeRegionTargetHttpsProxyAttributes) SslCertificates() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](crthp.ref.Append("ssl_certificates"))
}

// UrlMap returns a reference to field url_map of google_compute_region_target_https_proxy.
func (crthp computeRegionTargetHttpsProxyAttributes) UrlMap() terra.StringValue {
	return terra.ReferenceAsString(crthp.ref.Append("url_map"))
}

func (crthp computeRegionTargetHttpsProxyAttributes) Timeouts() computeregiontargethttpsproxy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregiontargethttpsproxy.TimeoutsAttributes](crthp.ref.Append("timeouts"))
}

type computeRegionTargetHttpsProxyState struct {
	CreationTimestamp string                                       `json:"creation_timestamp"`
	Description       string                                       `json:"description"`
	Id                string                                       `json:"id"`
	Name              string                                       `json:"name"`
	Project           string                                       `json:"project"`
	ProxyId           float64                                      `json:"proxy_id"`
	Region            string                                       `json:"region"`
	SelfLink          string                                       `json:"self_link"`
	SslCertificates   []string                                     `json:"ssl_certificates"`
	UrlMap            string                                       `json:"url_map"`
	Timeouts          *computeregiontargethttpsproxy.TimeoutsState `json:"timeouts"`
}
