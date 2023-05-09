// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeregiontargettcpproxy "github.com/golingon/terraproviders/googlebeta/4.63.1/computeregiontargettcpproxy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionTargetTcpProxy creates a new instance of [ComputeRegionTargetTcpProxy].
func NewComputeRegionTargetTcpProxy(name string, args ComputeRegionTargetTcpProxyArgs) *ComputeRegionTargetTcpProxy {
	return &ComputeRegionTargetTcpProxy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionTargetTcpProxy)(nil)

// ComputeRegionTargetTcpProxy represents the Terraform resource google_compute_region_target_tcp_proxy.
type ComputeRegionTargetTcpProxy struct {
	Name      string
	Args      ComputeRegionTargetTcpProxyArgs
	state     *computeRegionTargetTcpProxyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionTargetTcpProxy].
func (crttp *ComputeRegionTargetTcpProxy) Type() string {
	return "google_compute_region_target_tcp_proxy"
}

// LocalName returns the local name for [ComputeRegionTargetTcpProxy].
func (crttp *ComputeRegionTargetTcpProxy) LocalName() string {
	return crttp.Name
}

// Configuration returns the configuration (args) for [ComputeRegionTargetTcpProxy].
func (crttp *ComputeRegionTargetTcpProxy) Configuration() interface{} {
	return crttp.Args
}

// DependOn is used for other resources to depend on [ComputeRegionTargetTcpProxy].
func (crttp *ComputeRegionTargetTcpProxy) DependOn() terra.Reference {
	return terra.ReferenceResource(crttp)
}

// Dependencies returns the list of resources [ComputeRegionTargetTcpProxy] depends_on.
func (crttp *ComputeRegionTargetTcpProxy) Dependencies() terra.Dependencies {
	return crttp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionTargetTcpProxy].
func (crttp *ComputeRegionTargetTcpProxy) LifecycleManagement() *terra.Lifecycle {
	return crttp.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionTargetTcpProxy].
func (crttp *ComputeRegionTargetTcpProxy) Attributes() computeRegionTargetTcpProxyAttributes {
	return computeRegionTargetTcpProxyAttributes{ref: terra.ReferenceResource(crttp)}
}

// ImportState imports the given attribute values into [ComputeRegionTargetTcpProxy]'s state.
func (crttp *ComputeRegionTargetTcpProxy) ImportState(av io.Reader) error {
	crttp.state = &computeRegionTargetTcpProxyState{}
	if err := json.NewDecoder(av).Decode(crttp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crttp.Type(), crttp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionTargetTcpProxy] has state.
func (crttp *ComputeRegionTargetTcpProxy) State() (*computeRegionTargetTcpProxyState, bool) {
	return crttp.state, crttp.state != nil
}

// StateMust returns the state for [ComputeRegionTargetTcpProxy]. Panics if the state is nil.
func (crttp *ComputeRegionTargetTcpProxy) StateMust() *computeRegionTargetTcpProxyState {
	if crttp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crttp.Type(), crttp.LocalName()))
	}
	return crttp.state
}

// ComputeRegionTargetTcpProxyArgs contains the configurations for google_compute_region_target_tcp_proxy.
type ComputeRegionTargetTcpProxyArgs struct {
	// BackendService: string, required
	BackendService terra.StringValue `hcl:"backend_service,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ProxyBind: bool, optional
	ProxyBind terra.BoolValue `hcl:"proxy_bind,attr"`
	// ProxyHeader: string, optional
	ProxyHeader terra.StringValue `hcl:"proxy_header,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Timeouts: optional
	Timeouts *computeregiontargettcpproxy.Timeouts `hcl:"timeouts,block"`
}
type computeRegionTargetTcpProxyAttributes struct {
	ref terra.Reference
}

// BackendService returns a reference to field backend_service of google_compute_region_target_tcp_proxy.
func (crttp computeRegionTargetTcpProxyAttributes) BackendService() terra.StringValue {
	return terra.ReferenceAsString(crttp.ref.Append("backend_service"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_region_target_tcp_proxy.
func (crttp computeRegionTargetTcpProxyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(crttp.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_region_target_tcp_proxy.
func (crttp computeRegionTargetTcpProxyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crttp.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_region_target_tcp_proxy.
func (crttp computeRegionTargetTcpProxyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crttp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_target_tcp_proxy.
func (crttp computeRegionTargetTcpProxyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crttp.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_target_tcp_proxy.
func (crttp computeRegionTargetTcpProxyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crttp.ref.Append("project"))
}

// ProxyBind returns a reference to field proxy_bind of google_compute_region_target_tcp_proxy.
func (crttp computeRegionTargetTcpProxyAttributes) ProxyBind() terra.BoolValue {
	return terra.ReferenceAsBool(crttp.ref.Append("proxy_bind"))
}

// ProxyHeader returns a reference to field proxy_header of google_compute_region_target_tcp_proxy.
func (crttp computeRegionTargetTcpProxyAttributes) ProxyHeader() terra.StringValue {
	return terra.ReferenceAsString(crttp.ref.Append("proxy_header"))
}

// ProxyId returns a reference to field proxy_id of google_compute_region_target_tcp_proxy.
func (crttp computeRegionTargetTcpProxyAttributes) ProxyId() terra.NumberValue {
	return terra.ReferenceAsNumber(crttp.ref.Append("proxy_id"))
}

// Region returns a reference to field region of google_compute_region_target_tcp_proxy.
func (crttp computeRegionTargetTcpProxyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crttp.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_target_tcp_proxy.
func (crttp computeRegionTargetTcpProxyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crttp.ref.Append("self_link"))
}

func (crttp computeRegionTargetTcpProxyAttributes) Timeouts() computeregiontargettcpproxy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregiontargettcpproxy.TimeoutsAttributes](crttp.ref.Append("timeouts"))
}

type computeRegionTargetTcpProxyState struct {
	BackendService    string                                     `json:"backend_service"`
	CreationTimestamp string                                     `json:"creation_timestamp"`
	Description       string                                     `json:"description"`
	Id                string                                     `json:"id"`
	Name              string                                     `json:"name"`
	Project           string                                     `json:"project"`
	ProxyBind         bool                                       `json:"proxy_bind"`
	ProxyHeader       string                                     `json:"proxy_header"`
	ProxyId           float64                                    `json:"proxy_id"`
	Region            string                                     `json:"region"`
	SelfLink          string                                     `json:"self_link"`
	Timeouts          *computeregiontargettcpproxy.TimeoutsState `json:"timeouts"`
}
