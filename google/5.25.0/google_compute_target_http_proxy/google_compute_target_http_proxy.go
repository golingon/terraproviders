// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_target_http_proxy

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_compute_target_http_proxy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeTargetHttpProxyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcthp *Resource) Type() string {
	return "google_compute_target_http_proxy"
}

// LocalName returns the local name for [Resource].
func (gcthp *Resource) LocalName() string {
	return gcthp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcthp *Resource) Configuration() interface{} {
	return gcthp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcthp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcthp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcthp *Resource) Dependencies() terra.Dependencies {
	return gcthp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcthp *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcthp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcthp *Resource) Attributes() googleComputeTargetHttpProxyAttributes {
	return googleComputeTargetHttpProxyAttributes{ref: terra.ReferenceResource(gcthp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcthp *Resource) ImportState(state io.Reader) error {
	gcthp.state = &googleComputeTargetHttpProxyState{}
	if err := json.NewDecoder(state).Decode(gcthp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcthp.Type(), gcthp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcthp *Resource) State() (*googleComputeTargetHttpProxyState, bool) {
	return gcthp.state, gcthp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcthp *Resource) StateMust() *googleComputeTargetHttpProxyState {
	if gcthp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcthp.Type(), gcthp.LocalName()))
	}
	return gcthp.state
}

// Args contains the configurations for google_compute_target_http_proxy.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleComputeTargetHttpProxyAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_target_http_proxy.
func (gcthp googleComputeTargetHttpProxyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(gcthp.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_target_http_proxy.
func (gcthp googleComputeTargetHttpProxyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gcthp.ref.Append("description"))
}

// HttpKeepAliveTimeoutSec returns a reference to field http_keep_alive_timeout_sec of google_compute_target_http_proxy.
func (gcthp googleComputeTargetHttpProxyAttributes) HttpKeepAliveTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(gcthp.ref.Append("http_keep_alive_timeout_sec"))
}

// Id returns a reference to field id of google_compute_target_http_proxy.
func (gcthp googleComputeTargetHttpProxyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcthp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_target_http_proxy.
func (gcthp googleComputeTargetHttpProxyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcthp.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_target_http_proxy.
func (gcthp googleComputeTargetHttpProxyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcthp.ref.Append("project"))
}

// ProxyBind returns a reference to field proxy_bind of google_compute_target_http_proxy.
func (gcthp googleComputeTargetHttpProxyAttributes) ProxyBind() terra.BoolValue {
	return terra.ReferenceAsBool(gcthp.ref.Append("proxy_bind"))
}

// ProxyId returns a reference to field proxy_id of google_compute_target_http_proxy.
func (gcthp googleComputeTargetHttpProxyAttributes) ProxyId() terra.NumberValue {
	return terra.ReferenceAsNumber(gcthp.ref.Append("proxy_id"))
}

// SelfLink returns a reference to field self_link of google_compute_target_http_proxy.
func (gcthp googleComputeTargetHttpProxyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gcthp.ref.Append("self_link"))
}

// UrlMap returns a reference to field url_map of google_compute_target_http_proxy.
func (gcthp googleComputeTargetHttpProxyAttributes) UrlMap() terra.StringValue {
	return terra.ReferenceAsString(gcthp.ref.Append("url_map"))
}

func (gcthp googleComputeTargetHttpProxyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcthp.ref.Append("timeouts"))
}

type googleComputeTargetHttpProxyState struct {
	CreationTimestamp       string         `json:"creation_timestamp"`
	Description             string         `json:"description"`
	HttpKeepAliveTimeoutSec float64        `json:"http_keep_alive_timeout_sec"`
	Id                      string         `json:"id"`
	Name                    string         `json:"name"`
	Project                 string         `json:"project"`
	ProxyBind               bool           `json:"proxy_bind"`
	ProxyId                 float64        `json:"proxy_id"`
	SelfLink                string         `json:"self_link"`
	UrlMap                  string         `json:"url_map"`
	Timeouts                *TimeoutsState `json:"timeouts"`
}
