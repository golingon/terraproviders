// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computetargetsslproxy "github.com/golingon/terraproviders/google/4.59.0/computetargetsslproxy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeTargetSslProxy(name string, args ComputeTargetSslProxyArgs) *ComputeTargetSslProxy {
	return &ComputeTargetSslProxy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeTargetSslProxy)(nil)

type ComputeTargetSslProxy struct {
	Name  string
	Args  ComputeTargetSslProxyArgs
	state *computeTargetSslProxyState
}

func (ctsp *ComputeTargetSslProxy) Type() string {
	return "google_compute_target_ssl_proxy"
}

func (ctsp *ComputeTargetSslProxy) LocalName() string {
	return ctsp.Name
}

func (ctsp *ComputeTargetSslProxy) Configuration() interface{} {
	return ctsp.Args
}

func (ctsp *ComputeTargetSslProxy) Attributes() computeTargetSslProxyAttributes {
	return computeTargetSslProxyAttributes{ref: terra.ReferenceResource(ctsp)}
}

func (ctsp *ComputeTargetSslProxy) ImportState(av io.Reader) error {
	ctsp.state = &computeTargetSslProxyState{}
	if err := json.NewDecoder(av).Decode(ctsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ctsp.Type(), ctsp.LocalName(), err)
	}
	return nil
}

func (ctsp *ComputeTargetSslProxy) State() (*computeTargetSslProxyState, bool) {
	return ctsp.state, ctsp.state != nil
}

func (ctsp *ComputeTargetSslProxy) StateMust() *computeTargetSslProxyState {
	if ctsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ctsp.Type(), ctsp.LocalName()))
	}
	return ctsp.state
}

func (ctsp *ComputeTargetSslProxy) DependOn() terra.Reference {
	return terra.ReferenceResource(ctsp)
}

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
	// DependsOn contains resources that ComputeTargetSslProxy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeTargetSslProxyAttributes struct {
	ref terra.Reference
}

func (ctsp computeTargetSslProxyAttributes) BackendService() terra.StringValue {
	return terra.ReferenceString(ctsp.ref.Append("backend_service"))
}

func (ctsp computeTargetSslProxyAttributes) CertificateMap() terra.StringValue {
	return terra.ReferenceString(ctsp.ref.Append("certificate_map"))
}

func (ctsp computeTargetSslProxyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceString(ctsp.ref.Append("creation_timestamp"))
}

func (ctsp computeTargetSslProxyAttributes) Description() terra.StringValue {
	return terra.ReferenceString(ctsp.ref.Append("description"))
}

func (ctsp computeTargetSslProxyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ctsp.ref.Append("id"))
}

func (ctsp computeTargetSslProxyAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ctsp.ref.Append("name"))
}

func (ctsp computeTargetSslProxyAttributes) Project() terra.StringValue {
	return terra.ReferenceString(ctsp.ref.Append("project"))
}

func (ctsp computeTargetSslProxyAttributes) ProxyHeader() terra.StringValue {
	return terra.ReferenceString(ctsp.ref.Append("proxy_header"))
}

func (ctsp computeTargetSslProxyAttributes) ProxyId() terra.NumberValue {
	return terra.ReferenceNumber(ctsp.ref.Append("proxy_id"))
}

func (ctsp computeTargetSslProxyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(ctsp.ref.Append("self_link"))
}

func (ctsp computeTargetSslProxyAttributes) SslCertificates() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](ctsp.ref.Append("ssl_certificates"))
}

func (ctsp computeTargetSslProxyAttributes) SslPolicy() terra.StringValue {
	return terra.ReferenceString(ctsp.ref.Append("ssl_policy"))
}

func (ctsp computeTargetSslProxyAttributes) Timeouts() computetargetsslproxy.TimeoutsAttributes {
	return terra.ReferenceSingle[computetargetsslproxy.TimeoutsAttributes](ctsp.ref.Append("timeouts"))
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
