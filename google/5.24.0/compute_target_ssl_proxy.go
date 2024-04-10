// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	computetargetsslproxy "github.com/golingon/terraproviders/google/5.24.0/computetargetsslproxy"
	"io"
)

// NewComputeTargetSslProxy creates a new instance of [ComputeTargetSslProxy].
func NewComputeTargetSslProxy(name string, args ComputeTargetSslProxyArgs) *ComputeTargetSslProxy {
	return &ComputeTargetSslProxy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeTargetSslProxy)(nil)

// ComputeTargetSslProxy represents the Terraform resource google_compute_target_ssl_proxy.
type ComputeTargetSslProxy struct {
	Name      string
	Args      ComputeTargetSslProxyArgs
	state     *computeTargetSslProxyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeTargetSslProxy].
func (ctsp *ComputeTargetSslProxy) Type() string {
	return "google_compute_target_ssl_proxy"
}

// LocalName returns the local name for [ComputeTargetSslProxy].
func (ctsp *ComputeTargetSslProxy) LocalName() string {
	return ctsp.Name
}

// Configuration returns the configuration (args) for [ComputeTargetSslProxy].
func (ctsp *ComputeTargetSslProxy) Configuration() interface{} {
	return ctsp.Args
}

// DependOn is used for other resources to depend on [ComputeTargetSslProxy].
func (ctsp *ComputeTargetSslProxy) DependOn() terra.Reference {
	return terra.ReferenceResource(ctsp)
}

// Dependencies returns the list of resources [ComputeTargetSslProxy] depends_on.
func (ctsp *ComputeTargetSslProxy) Dependencies() terra.Dependencies {
	return ctsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeTargetSslProxy].
func (ctsp *ComputeTargetSslProxy) LifecycleManagement() *terra.Lifecycle {
	return ctsp.Lifecycle
}

// Attributes returns the attributes for [ComputeTargetSslProxy].
func (ctsp *ComputeTargetSslProxy) Attributes() computeTargetSslProxyAttributes {
	return computeTargetSslProxyAttributes{ref: terra.ReferenceResource(ctsp)}
}

// ImportState imports the given attribute values into [ComputeTargetSslProxy]'s state.
func (ctsp *ComputeTargetSslProxy) ImportState(av io.Reader) error {
	ctsp.state = &computeTargetSslProxyState{}
	if err := json.NewDecoder(av).Decode(ctsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ctsp.Type(), ctsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeTargetSslProxy] has state.
func (ctsp *ComputeTargetSslProxy) State() (*computeTargetSslProxyState, bool) {
	return ctsp.state, ctsp.state != nil
}

// StateMust returns the state for [ComputeTargetSslProxy]. Panics if the state is nil.
func (ctsp *ComputeTargetSslProxy) StateMust() *computeTargetSslProxyState {
	if ctsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ctsp.Type(), ctsp.LocalName()))
	}
	return ctsp.state
}

// ComputeTargetSslProxyArgs contains the configurations for google_compute_target_ssl_proxy.
type ComputeTargetSslProxyArgs struct {
	// BackendService: string, required
	BackendService terra.StringValue `hcl:"backend_service,attr" validate:"required"`
	// CertificateMap: string, optional
	CertificateMap terra.StringValue `hcl:"certificate_map,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ProxyHeader: string, optional
	ProxyHeader terra.StringValue `hcl:"proxy_header,attr"`
	// SslCertificates: list of string, optional
	SslCertificates terra.ListValue[terra.StringValue] `hcl:"ssl_certificates,attr"`
	// SslPolicy: string, optional
	SslPolicy terra.StringValue `hcl:"ssl_policy,attr"`
	// Timeouts: optional
	Timeouts *computetargetsslproxy.Timeouts `hcl:"timeouts,block"`
}
type computeTargetSslProxyAttributes struct {
	ref terra.Reference
}

// BackendService returns a reference to field backend_service of google_compute_target_ssl_proxy.
func (ctsp computeTargetSslProxyAttributes) BackendService() terra.StringValue {
	return terra.ReferenceAsString(ctsp.ref.Append("backend_service"))
}

// CertificateMap returns a reference to field certificate_map of google_compute_target_ssl_proxy.
func (ctsp computeTargetSslProxyAttributes) CertificateMap() terra.StringValue {
	return terra.ReferenceAsString(ctsp.ref.Append("certificate_map"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_target_ssl_proxy.
func (ctsp computeTargetSslProxyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(ctsp.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_target_ssl_proxy.
func (ctsp computeTargetSslProxyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ctsp.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_target_ssl_proxy.
func (ctsp computeTargetSslProxyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ctsp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_target_ssl_proxy.
func (ctsp computeTargetSslProxyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ctsp.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_target_ssl_proxy.
func (ctsp computeTargetSslProxyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ctsp.ref.Append("project"))
}

// ProxyHeader returns a reference to field proxy_header of google_compute_target_ssl_proxy.
func (ctsp computeTargetSslProxyAttributes) ProxyHeader() terra.StringValue {
	return terra.ReferenceAsString(ctsp.ref.Append("proxy_header"))
}

// ProxyId returns a reference to field proxy_id of google_compute_target_ssl_proxy.
func (ctsp computeTargetSslProxyAttributes) ProxyId() terra.NumberValue {
	return terra.ReferenceAsNumber(ctsp.ref.Append("proxy_id"))
}

// SelfLink returns a reference to field self_link of google_compute_target_ssl_proxy.
func (ctsp computeTargetSslProxyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(ctsp.ref.Append("self_link"))
}

// SslCertificates returns a reference to field ssl_certificates of google_compute_target_ssl_proxy.
func (ctsp computeTargetSslProxyAttributes) SslCertificates() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ctsp.ref.Append("ssl_certificates"))
}

// SslPolicy returns a reference to field ssl_policy of google_compute_target_ssl_proxy.
func (ctsp computeTargetSslProxyAttributes) SslPolicy() terra.StringValue {
	return terra.ReferenceAsString(ctsp.ref.Append("ssl_policy"))
}

func (ctsp computeTargetSslProxyAttributes) Timeouts() computetargetsslproxy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computetargetsslproxy.TimeoutsAttributes](ctsp.ref.Append("timeouts"))
}

type computeTargetSslProxyState struct {
	BackendService    string                               `json:"backend_service"`
	CertificateMap    string                               `json:"certificate_map"`
	CreationTimestamp string                               `json:"creation_timestamp"`
	Description       string                               `json:"description"`
	Id                string                               `json:"id"`
	Name              string                               `json:"name"`
	Project           string                               `json:"project"`
	ProxyHeader       string                               `json:"proxy_header"`
	ProxyId           float64                              `json:"proxy_id"`
	SelfLink          string                               `json:"self_link"`
	SslCertificates   []string                             `json:"ssl_certificates"`
	SslPolicy         string                               `json:"ssl_policy"`
	Timeouts          *computetargetsslproxy.TimeoutsState `json:"timeouts"`
}
