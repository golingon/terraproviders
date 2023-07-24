// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computetargettcpproxy "github.com/golingon/terraproviders/googlebeta/4.74.0/computetargettcpproxy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeTargetTcpProxy creates a new instance of [ComputeTargetTcpProxy].
func NewComputeTargetTcpProxy(name string, args ComputeTargetTcpProxyArgs) *ComputeTargetTcpProxy {
	return &ComputeTargetTcpProxy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeTargetTcpProxy)(nil)

// ComputeTargetTcpProxy represents the Terraform resource google_compute_target_tcp_proxy.
type ComputeTargetTcpProxy struct {
	Name      string
	Args      ComputeTargetTcpProxyArgs
	state     *computeTargetTcpProxyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeTargetTcpProxy].
func (cttp *ComputeTargetTcpProxy) Type() string {
	return "google_compute_target_tcp_proxy"
}

// LocalName returns the local name for [ComputeTargetTcpProxy].
func (cttp *ComputeTargetTcpProxy) LocalName() string {
	return cttp.Name
}

// Configuration returns the configuration (args) for [ComputeTargetTcpProxy].
func (cttp *ComputeTargetTcpProxy) Configuration() interface{} {
	return cttp.Args
}

// DependOn is used for other resources to depend on [ComputeTargetTcpProxy].
func (cttp *ComputeTargetTcpProxy) DependOn() terra.Reference {
	return terra.ReferenceResource(cttp)
}

// Dependencies returns the list of resources [ComputeTargetTcpProxy] depends_on.
func (cttp *ComputeTargetTcpProxy) Dependencies() terra.Dependencies {
	return cttp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeTargetTcpProxy].
func (cttp *ComputeTargetTcpProxy) LifecycleManagement() *terra.Lifecycle {
	return cttp.Lifecycle
}

// Attributes returns the attributes for [ComputeTargetTcpProxy].
func (cttp *ComputeTargetTcpProxy) Attributes() computeTargetTcpProxyAttributes {
	return computeTargetTcpProxyAttributes{ref: terra.ReferenceResource(cttp)}
}

// ImportState imports the given attribute values into [ComputeTargetTcpProxy]'s state.
func (cttp *ComputeTargetTcpProxy) ImportState(av io.Reader) error {
	cttp.state = &computeTargetTcpProxyState{}
	if err := json.NewDecoder(av).Decode(cttp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cttp.Type(), cttp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeTargetTcpProxy] has state.
func (cttp *ComputeTargetTcpProxy) State() (*computeTargetTcpProxyState, bool) {
	return cttp.state, cttp.state != nil
}

// StateMust returns the state for [ComputeTargetTcpProxy]. Panics if the state is nil.
func (cttp *ComputeTargetTcpProxy) StateMust() *computeTargetTcpProxyState {
	if cttp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cttp.Type(), cttp.LocalName()))
	}
	return cttp.state
}

// ComputeTargetTcpProxyArgs contains the configurations for google_compute_target_tcp_proxy.
type ComputeTargetTcpProxyArgs struct {
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
	// Timeouts: optional
	Timeouts *computetargettcpproxy.Timeouts `hcl:"timeouts,block"`
}
type computeTargetTcpProxyAttributes struct {
	ref terra.Reference
}

// BackendService returns a reference to field backend_service of google_compute_target_tcp_proxy.
func (cttp computeTargetTcpProxyAttributes) BackendService() terra.StringValue {
	return terra.ReferenceAsString(cttp.ref.Append("backend_service"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_target_tcp_proxy.
func (cttp computeTargetTcpProxyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cttp.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_target_tcp_proxy.
func (cttp computeTargetTcpProxyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cttp.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_target_tcp_proxy.
func (cttp computeTargetTcpProxyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cttp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_target_tcp_proxy.
func (cttp computeTargetTcpProxyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cttp.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_target_tcp_proxy.
func (cttp computeTargetTcpProxyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cttp.ref.Append("project"))
}

// ProxyBind returns a reference to field proxy_bind of google_compute_target_tcp_proxy.
func (cttp computeTargetTcpProxyAttributes) ProxyBind() terra.BoolValue {
	return terra.ReferenceAsBool(cttp.ref.Append("proxy_bind"))
}

// ProxyHeader returns a reference to field proxy_header of google_compute_target_tcp_proxy.
func (cttp computeTargetTcpProxyAttributes) ProxyHeader() terra.StringValue {
	return terra.ReferenceAsString(cttp.ref.Append("proxy_header"))
}

// ProxyId returns a reference to field proxy_id of google_compute_target_tcp_proxy.
func (cttp computeTargetTcpProxyAttributes) ProxyId() terra.NumberValue {
	return terra.ReferenceAsNumber(cttp.ref.Append("proxy_id"))
}

// SelfLink returns a reference to field self_link of google_compute_target_tcp_proxy.
func (cttp computeTargetTcpProxyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cttp.ref.Append("self_link"))
}

func (cttp computeTargetTcpProxyAttributes) Timeouts() computetargettcpproxy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computetargettcpproxy.TimeoutsAttributes](cttp.ref.Append("timeouts"))
}

type computeTargetTcpProxyState struct {
	BackendService    string                               `json:"backend_service"`
	CreationTimestamp string                               `json:"creation_timestamp"`
	Description       string                               `json:"description"`
	Id                string                               `json:"id"`
	Name              string                               `json:"name"`
	Project           string                               `json:"project"`
	ProxyBind         bool                                 `json:"proxy_bind"`
	ProxyHeader       string                               `json:"proxy_header"`
	ProxyId           float64                              `json:"proxy_id"`
	SelfLink          string                               `json:"self_link"`
	Timeouts          *computetargettcpproxy.TimeoutsState `json:"timeouts"`
}
