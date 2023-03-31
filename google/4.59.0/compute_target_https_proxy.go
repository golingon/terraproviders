// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computetargethttpsproxy "github.com/golingon/terraproviders/google/4.59.0/computetargethttpsproxy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeTargetHttpsProxy(name string, args ComputeTargetHttpsProxyArgs) *ComputeTargetHttpsProxy {
	return &ComputeTargetHttpsProxy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeTargetHttpsProxy)(nil)

type ComputeTargetHttpsProxy struct {
	Name  string
	Args  ComputeTargetHttpsProxyArgs
	state *computeTargetHttpsProxyState
}

func (cthp *ComputeTargetHttpsProxy) Type() string {
	return "google_compute_target_https_proxy"
}

func (cthp *ComputeTargetHttpsProxy) LocalName() string {
	return cthp.Name
}

func (cthp *ComputeTargetHttpsProxy) Configuration() interface{} {
	return cthp.Args
}

func (cthp *ComputeTargetHttpsProxy) Attributes() computeTargetHttpsProxyAttributes {
	return computeTargetHttpsProxyAttributes{ref: terra.ReferenceResource(cthp)}
}

func (cthp *ComputeTargetHttpsProxy) ImportState(av io.Reader) error {
	cthp.state = &computeTargetHttpsProxyState{}
	if err := json.NewDecoder(av).Decode(cthp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cthp.Type(), cthp.LocalName(), err)
	}
	return nil
}

func (cthp *ComputeTargetHttpsProxy) State() (*computeTargetHttpsProxyState, bool) {
	return cthp.state, cthp.state != nil
}

func (cthp *ComputeTargetHttpsProxy) StateMust() *computeTargetHttpsProxyState {
	if cthp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cthp.Type(), cthp.LocalName()))
	}
	return cthp.state
}

func (cthp *ComputeTargetHttpsProxy) DependOn() terra.Reference {
	return terra.ReferenceResource(cthp)
}

type ComputeTargetHttpsProxyArgs struct {
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
	// ProxyBind: bool, optional
	ProxyBind terra.BoolValue `hcl:"proxy_bind,attr"`
	// QuicOverride: string, optional
	QuicOverride terra.StringValue `hcl:"quic_override,attr"`
	// SslCertificates: list of string, optional
	SslCertificates terra.ListValue[terra.StringValue] `hcl:"ssl_certificates,attr"`
	// SslPolicy: string, optional
	SslPolicy terra.StringValue `hcl:"ssl_policy,attr"`
	// UrlMap: string, required
	UrlMap terra.StringValue `hcl:"url_map,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *computetargethttpsproxy.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ComputeTargetHttpsProxy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeTargetHttpsProxyAttributes struct {
	ref terra.Reference
}

func (cthp computeTargetHttpsProxyAttributes) CertificateMap() terra.StringValue {
	return terra.ReferenceString(cthp.ref.Append("certificate_map"))
}

func (cthp computeTargetHttpsProxyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceString(cthp.ref.Append("creation_timestamp"))
}

func (cthp computeTargetHttpsProxyAttributes) Description() terra.StringValue {
	return terra.ReferenceString(cthp.ref.Append("description"))
}

func (cthp computeTargetHttpsProxyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cthp.ref.Append("id"))
}

func (cthp computeTargetHttpsProxyAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cthp.ref.Append("name"))
}

func (cthp computeTargetHttpsProxyAttributes) Project() terra.StringValue {
	return terra.ReferenceString(cthp.ref.Append("project"))
}

func (cthp computeTargetHttpsProxyAttributes) ProxyBind() terra.BoolValue {
	return terra.ReferenceBool(cthp.ref.Append("proxy_bind"))
}

func (cthp computeTargetHttpsProxyAttributes) ProxyId() terra.NumberValue {
	return terra.ReferenceNumber(cthp.ref.Append("proxy_id"))
}

func (cthp computeTargetHttpsProxyAttributes) QuicOverride() terra.StringValue {
	return terra.ReferenceString(cthp.ref.Append("quic_override"))
}

func (cthp computeTargetHttpsProxyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(cthp.ref.Append("self_link"))
}

func (cthp computeTargetHttpsProxyAttributes) SslCertificates() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](cthp.ref.Append("ssl_certificates"))
}

func (cthp computeTargetHttpsProxyAttributes) SslPolicy() terra.StringValue {
	return terra.ReferenceString(cthp.ref.Append("ssl_policy"))
}

func (cthp computeTargetHttpsProxyAttributes) UrlMap() terra.StringValue {
	return terra.ReferenceString(cthp.ref.Append("url_map"))
}

func (cthp computeTargetHttpsProxyAttributes) Timeouts() computetargethttpsproxy.TimeoutsAttributes {
	return terra.ReferenceSingle[computetargethttpsproxy.TimeoutsAttributes](cthp.ref.Append("timeouts"))
}

type computeTargetHttpsProxyState struct {
	CertificateMap    string                                 `json:"certificate_map"`
	CreationTimestamp string                                 `json:"creation_timestamp"`
	Description       string                                 `json:"description"`
	Id                string                                 `json:"id"`
	Name              string                                 `json:"name"`
	Project           string                                 `json:"project"`
	ProxyBind         bool                                   `json:"proxy_bind"`
	ProxyId           float64                                `json:"proxy_id"`
	QuicOverride      string                                 `json:"quic_override"`
	SelfLink          string                                 `json:"self_link"`
	SslCertificates   []string                               `json:"ssl_certificates"`
	SslPolicy         string                                 `json:"ssl_policy"`
	UrlMap            string                                 `json:"url_map"`
	Timeouts          *computetargethttpsproxy.TimeoutsState `json:"timeouts"`
}
