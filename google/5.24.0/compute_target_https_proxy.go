// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	computetargethttpsproxy "github.com/golingon/terraproviders/google/5.24.0/computetargethttpsproxy"
	"io"
)

// NewComputeTargetHttpsProxy creates a new instance of [ComputeTargetHttpsProxy].
func NewComputeTargetHttpsProxy(name string, args ComputeTargetHttpsProxyArgs) *ComputeTargetHttpsProxy {
	return &ComputeTargetHttpsProxy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeTargetHttpsProxy)(nil)

// ComputeTargetHttpsProxy represents the Terraform resource google_compute_target_https_proxy.
type ComputeTargetHttpsProxy struct {
	Name      string
	Args      ComputeTargetHttpsProxyArgs
	state     *computeTargetHttpsProxyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeTargetHttpsProxy].
func (cthp *ComputeTargetHttpsProxy) Type() string {
	return "google_compute_target_https_proxy"
}

// LocalName returns the local name for [ComputeTargetHttpsProxy].
func (cthp *ComputeTargetHttpsProxy) LocalName() string {
	return cthp.Name
}

// Configuration returns the configuration (args) for [ComputeTargetHttpsProxy].
func (cthp *ComputeTargetHttpsProxy) Configuration() interface{} {
	return cthp.Args
}

// DependOn is used for other resources to depend on [ComputeTargetHttpsProxy].
func (cthp *ComputeTargetHttpsProxy) DependOn() terra.Reference {
	return terra.ReferenceResource(cthp)
}

// Dependencies returns the list of resources [ComputeTargetHttpsProxy] depends_on.
func (cthp *ComputeTargetHttpsProxy) Dependencies() terra.Dependencies {
	return cthp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeTargetHttpsProxy].
func (cthp *ComputeTargetHttpsProxy) LifecycleManagement() *terra.Lifecycle {
	return cthp.Lifecycle
}

// Attributes returns the attributes for [ComputeTargetHttpsProxy].
func (cthp *ComputeTargetHttpsProxy) Attributes() computeTargetHttpsProxyAttributes {
	return computeTargetHttpsProxyAttributes{ref: terra.ReferenceResource(cthp)}
}

// ImportState imports the given attribute values into [ComputeTargetHttpsProxy]'s state.
func (cthp *ComputeTargetHttpsProxy) ImportState(av io.Reader) error {
	cthp.state = &computeTargetHttpsProxyState{}
	if err := json.NewDecoder(av).Decode(cthp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cthp.Type(), cthp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeTargetHttpsProxy] has state.
func (cthp *ComputeTargetHttpsProxy) State() (*computeTargetHttpsProxyState, bool) {
	return cthp.state, cthp.state != nil
}

// StateMust returns the state for [ComputeTargetHttpsProxy]. Panics if the state is nil.
func (cthp *ComputeTargetHttpsProxy) StateMust() *computeTargetHttpsProxyState {
	if cthp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cthp.Type(), cthp.LocalName()))
	}
	return cthp.state
}

// ComputeTargetHttpsProxyArgs contains the configurations for google_compute_target_https_proxy.
type ComputeTargetHttpsProxyArgs struct {
	// CertificateManagerCertificates: list of string, optional
	CertificateManagerCertificates terra.ListValue[terra.StringValue] `hcl:"certificate_manager_certificates,attr"`
	// CertificateMap: string, optional
	CertificateMap terra.StringValue `hcl:"certificate_map,attr"`
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
	// QuicOverride: string, optional
	QuicOverride terra.StringValue `hcl:"quic_override,attr"`
	// ServerTlsPolicy: string, optional
	ServerTlsPolicy terra.StringValue `hcl:"server_tls_policy,attr"`
	// SslCertificates: list of string, optional
	SslCertificates terra.ListValue[terra.StringValue] `hcl:"ssl_certificates,attr"`
	// SslPolicy: string, optional
	SslPolicy terra.StringValue `hcl:"ssl_policy,attr"`
	// UrlMap: string, required
	UrlMap terra.StringValue `hcl:"url_map,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *computetargethttpsproxy.Timeouts `hcl:"timeouts,block"`
}
type computeTargetHttpsProxyAttributes struct {
	ref terra.Reference
}

// CertificateManagerCertificates returns a reference to field certificate_manager_certificates of google_compute_target_https_proxy.
func (cthp computeTargetHttpsProxyAttributes) CertificateManagerCertificates() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cthp.ref.Append("certificate_manager_certificates"))
}

// CertificateMap returns a reference to field certificate_map of google_compute_target_https_proxy.
func (cthp computeTargetHttpsProxyAttributes) CertificateMap() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("certificate_map"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_target_https_proxy.
func (cthp computeTargetHttpsProxyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_target_https_proxy.
func (cthp computeTargetHttpsProxyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("description"))
}

// HttpKeepAliveTimeoutSec returns a reference to field http_keep_alive_timeout_sec of google_compute_target_https_proxy.
func (cthp computeTargetHttpsProxyAttributes) HttpKeepAliveTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(cthp.ref.Append("http_keep_alive_timeout_sec"))
}

// Id returns a reference to field id of google_compute_target_https_proxy.
func (cthp computeTargetHttpsProxyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_target_https_proxy.
func (cthp computeTargetHttpsProxyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_target_https_proxy.
func (cthp computeTargetHttpsProxyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("project"))
}

// ProxyBind returns a reference to field proxy_bind of google_compute_target_https_proxy.
func (cthp computeTargetHttpsProxyAttributes) ProxyBind() terra.BoolValue {
	return terra.ReferenceAsBool(cthp.ref.Append("proxy_bind"))
}

// ProxyId returns a reference to field proxy_id of google_compute_target_https_proxy.
func (cthp computeTargetHttpsProxyAttributes) ProxyId() terra.NumberValue {
	return terra.ReferenceAsNumber(cthp.ref.Append("proxy_id"))
}

// QuicOverride returns a reference to field quic_override of google_compute_target_https_proxy.
func (cthp computeTargetHttpsProxyAttributes) QuicOverride() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("quic_override"))
}

// SelfLink returns a reference to field self_link of google_compute_target_https_proxy.
func (cthp computeTargetHttpsProxyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("self_link"))
}

// ServerTlsPolicy returns a reference to field server_tls_policy of google_compute_target_https_proxy.
func (cthp computeTargetHttpsProxyAttributes) ServerTlsPolicy() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("server_tls_policy"))
}

// SslCertificates returns a reference to field ssl_certificates of google_compute_target_https_proxy.
func (cthp computeTargetHttpsProxyAttributes) SslCertificates() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cthp.ref.Append("ssl_certificates"))
}

// SslPolicy returns a reference to field ssl_policy of google_compute_target_https_proxy.
func (cthp computeTargetHttpsProxyAttributes) SslPolicy() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("ssl_policy"))
}

// UrlMap returns a reference to field url_map of google_compute_target_https_proxy.
func (cthp computeTargetHttpsProxyAttributes) UrlMap() terra.StringValue {
	return terra.ReferenceAsString(cthp.ref.Append("url_map"))
}

func (cthp computeTargetHttpsProxyAttributes) Timeouts() computetargethttpsproxy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computetargethttpsproxy.TimeoutsAttributes](cthp.ref.Append("timeouts"))
}

type computeTargetHttpsProxyState struct {
	CertificateManagerCertificates []string                               `json:"certificate_manager_certificates"`
	CertificateMap                 string                                 `json:"certificate_map"`
	CreationTimestamp              string                                 `json:"creation_timestamp"`
	Description                    string                                 `json:"description"`
	HttpKeepAliveTimeoutSec        float64                                `json:"http_keep_alive_timeout_sec"`
	Id                             string                                 `json:"id"`
	Name                           string                                 `json:"name"`
	Project                        string                                 `json:"project"`
	ProxyBind                      bool                                   `json:"proxy_bind"`
	ProxyId                        float64                                `json:"proxy_id"`
	QuicOverride                   string                                 `json:"quic_override"`
	SelfLink                       string                                 `json:"self_link"`
	ServerTlsPolicy                string                                 `json:"server_tls_policy"`
	SslCertificates                []string                               `json:"ssl_certificates"`
	SslPolicy                      string                                 `json:"ssl_policy"`
	UrlMap                         string                                 `json:"url_map"`
	Timeouts                       *computetargethttpsproxy.TimeoutsState `json:"timeouts"`
}
