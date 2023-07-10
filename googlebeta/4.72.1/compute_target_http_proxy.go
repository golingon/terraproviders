// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computetargethttpproxy "github.com/golingon/terraproviders/googlebeta/4.72.1/computetargethttpproxy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeTargetHttpProxy creates a new instance of [ComputeTargetHttpProxy].
func NewComputeTargetHttpProxy(name string, args ComputeTargetHttpProxyArgs) *ComputeTargetHttpProxy {
	return &ComputeTargetHttpProxy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeTargetHttpProxy)(nil)

// ComputeTargetHttpProxy represents the Terraform resource google_compute_target_http_proxy.
type ComputeTargetHttpProxy struct {
	Name      string
	Args      ComputeTargetHttpProxyArgs
	state     *computeTargetHttpProxyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeTargetHttpProxy].
func (cthp *ComputeTargetHttpProxy) Type() string {
	return "google_compute_target_http_proxy"
}

// LocalName returns the local name for [ComputeTargetHttpProxy].
func (cthp *ComputeTargetHttpProxy) LocalName() string {
	return cthp.Name
}

// Configuration returns the configuration (args) for [ComputeTargetHttpProxy].
func (cthp *ComputeTargetHttpProxy) Configuration() interface{} {
	return cthp.Args
}

// DependOn is used for other resources to depend on [ComputeTargetHttpProxy].
func (cthp *ComputeTargetHttpProxy) DependOn() terra.Reference {
	return terra.ReferenceResource(cthp)
}

// Dependencies returns the list of resources [ComputeTargetHttpProxy] depends_on.
func (cthp *ComputeTargetHttpProxy) Dependencies() terra.Dependencies {
	return cthp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeTargetHttpProxy].
func (cthp *ComputeTargetHttpProxy) LifecycleManagement() *terra.Lifecycle {
	return cthp.Lifecycle
}

// Attributes returns the attributes for [ComputeTargetHttpProxy].
func (cthp *ComputeTargetHttpProxy) Attributes() computeTargetHttpProxyAttributes {
	return computeTargetHttpProxyAttributes{ref: terra.ReferenceResource(cthp)}
}

// ImportState imports the given attribute values into [ComputeTargetHttpProxy]'s state.
func (cthp *ComputeTargetHttpProxy) ImportState(av io.Reader) error {
	cthp.state = &computeTargetHttpProxyState{}
	if err := json.NewDecoder(av).Decode(cthp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cthp.Type(), cthp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeTargetHttpProxy] has state.
func (cthp *ComputeTargetHttpProxy) State() (*computeTargetHttpProxyState, bool) {
	return cthp.state, cthp.state != nil
}

// StateMust returns the state for [ComputeTargetHttpProxy]. Panics if the state is nil.
func (cthp *ComputeTargetHttpProxy) StateMust() *computeTargetHttpProxyState {
	if cthp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cthp.Type(), cthp.LocalName()))
	}
	return cthp.state
}

// ComputeTargetHttpProxyArgs contains the configurations for google_compute_target_http_proxy.
type ComputeTargetHttpProxyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// HttpKeepAliveTimeoutSec: number, optional
	HttpKeepAliveTimeoutSec terra.NumberValue `hcl:"http_keep_alive_timeout_sec,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ProxyBind: bool, optional
	ProxyBind terra.BoolValue `hcl:"proxy_bind,attr"`
	// UrlMap: string, required
	UrlMap terra.StringValue `hcl:"url_map,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *computetargethttpproxy.Timeouts `hcl:"timeouts,block"`
}
type computeTargetHttpProxyAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_target_http_proxy.
func (cthp computeTargetHttpProxyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_target_http_proxy.
func (cthp computeTargetHttpProxyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("description"))
}

// HttpKeepAliveTimeoutSec returns a reference to field http_keep_alive_timeout_sec of google_compute_target_http_proxy.
func (cthp computeTargetHttpProxyAttributes) HttpKeepAliveTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(cthp.ref.Append("http_keep_alive_timeout_sec"))
}

// Id returns a reference to field id of google_compute_target_http_proxy.
func (cthp computeTargetHttpProxyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_target_http_proxy.
func (cthp computeTargetHttpProxyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_target_http_proxy.
func (cthp computeTargetHttpProxyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("project"))
}

// ProxyBind returns a reference to field proxy_bind of google_compute_target_http_proxy.
func (cthp computeTargetHttpProxyAttributes) ProxyBind() terra.BoolValue {
	return terra.ReferenceAsBool(cthp.ref.Append("proxy_bind"))
}

// ProxyId returns a reference to field proxy_id of google_compute_target_http_proxy.
func (cthp computeTargetHttpProxyAttributes) ProxyId() terra.NumberValue {
	return terra.ReferenceAsNumber(cthp.ref.Append("proxy_id"))
}

// SelfLink returns a reference to field self_link of google_compute_target_http_proxy.
func (cthp computeTargetHttpProxyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("self_link"))
}

// UrlMap returns a reference to field url_map of google_compute_target_http_proxy.
func (cthp computeTargetHttpProxyAttributes) UrlMap() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("url_map"))
}

func (cthp computeTargetHttpProxyAttributes) Timeouts() computetargethttpproxy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computetargethttpproxy.TimeoutsAttributes](cthp.ref.Append("timeouts"))
}

type computeTargetHttpProxyState struct {
	CreationTimestamp       string                                `json:"creation_timestamp"`
	Description             string                                `json:"description"`
	HttpKeepAliveTimeoutSec float64                               `json:"http_keep_alive_timeout_sec"`
	Id                      string                                `json:"id"`
	Name                    string                                `json:"name"`
	Project                 string                                `json:"project"`
	ProxyBind               bool                                  `json:"proxy_bind"`
	ProxyId                 float64                               `json:"proxy_id"`
	SelfLink                string                                `json:"self_link"`
	UrlMap                  string                                `json:"url_map"`
	Timeouts                *computetargethttpproxy.TimeoutsState `json:"timeouts"`
}
