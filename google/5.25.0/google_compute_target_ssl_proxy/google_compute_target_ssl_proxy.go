// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_target_ssl_proxy

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

// Resource represents the Terraform resource google_compute_target_ssl_proxy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeTargetSslProxyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gctsp *Resource) Type() string {
	return "google_compute_target_ssl_proxy"
}

// LocalName returns the local name for [Resource].
func (gctsp *Resource) LocalName() string {
	return gctsp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gctsp *Resource) Configuration() interface{} {
	return gctsp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gctsp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gctsp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gctsp *Resource) Dependencies() terra.Dependencies {
	return gctsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gctsp *Resource) LifecycleManagement() *terra.Lifecycle {
	return gctsp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gctsp *Resource) Attributes() googleComputeTargetSslProxyAttributes {
	return googleComputeTargetSslProxyAttributes{ref: terra.ReferenceResource(gctsp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gctsp *Resource) ImportState(state io.Reader) error {
	gctsp.state = &googleComputeTargetSslProxyState{}
	if err := json.NewDecoder(state).Decode(gctsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gctsp.Type(), gctsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gctsp *Resource) State() (*googleComputeTargetSslProxyState, bool) {
	return gctsp.state, gctsp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gctsp *Resource) StateMust() *googleComputeTargetSslProxyState {
	if gctsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gctsp.Type(), gctsp.LocalName()))
	}
	return gctsp.state
}

// Args contains the configurations for google_compute_target_ssl_proxy.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleComputeTargetSslProxyAttributes struct {
	ref terra.Reference
}

// BackendService returns a reference to field backend_service of google_compute_target_ssl_proxy.
func (gctsp googleComputeTargetSslProxyAttributes) BackendService() terra.StringValue {
	return terra.ReferenceAsString(gctsp.ref.Append("backend_service"))
}

// CertificateMap returns a reference to field certificate_map of google_compute_target_ssl_proxy.
func (gctsp googleComputeTargetSslProxyAttributes) CertificateMap() terra.StringValue {
	return terra.ReferenceAsString(gctsp.ref.Append("certificate_map"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_target_ssl_proxy.
func (gctsp googleComputeTargetSslProxyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(gctsp.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_target_ssl_proxy.
func (gctsp googleComputeTargetSslProxyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gctsp.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_target_ssl_proxy.
func (gctsp googleComputeTargetSslProxyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gctsp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_target_ssl_proxy.
func (gctsp googleComputeTargetSslProxyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gctsp.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_target_ssl_proxy.
func (gctsp googleComputeTargetSslProxyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gctsp.ref.Append("project"))
}

// ProxyHeader returns a reference to field proxy_header of google_compute_target_ssl_proxy.
func (gctsp googleComputeTargetSslProxyAttributes) ProxyHeader() terra.StringValue {
	return terra.ReferenceAsString(gctsp.ref.Append("proxy_header"))
}

// ProxyId returns a reference to field proxy_id of google_compute_target_ssl_proxy.
func (gctsp googleComputeTargetSslProxyAttributes) ProxyId() terra.NumberValue {
	return terra.ReferenceAsNumber(gctsp.ref.Append("proxy_id"))
}

// SelfLink returns a reference to field self_link of google_compute_target_ssl_proxy.
func (gctsp googleComputeTargetSslProxyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gctsp.ref.Append("self_link"))
}

// SslCertificates returns a reference to field ssl_certificates of google_compute_target_ssl_proxy.
func (gctsp googleComputeTargetSslProxyAttributes) SslCertificates() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gctsp.ref.Append("ssl_certificates"))
}

// SslPolicy returns a reference to field ssl_policy of google_compute_target_ssl_proxy.
func (gctsp googleComputeTargetSslProxyAttributes) SslPolicy() terra.StringValue {
	return terra.ReferenceAsString(gctsp.ref.Append("ssl_policy"))
}

func (gctsp googleComputeTargetSslProxyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gctsp.ref.Append("timeouts"))
}

type googleComputeTargetSslProxyState struct {
	BackendService    string         `json:"backend_service"`
	CertificateMap    string         `json:"certificate_map"`
	CreationTimestamp string         `json:"creation_timestamp"`
	Description       string         `json:"description"`
	Id                string         `json:"id"`
	Name              string         `json:"name"`
	Project           string         `json:"project"`
	ProxyHeader       string         `json:"proxy_header"`
	ProxyId           float64        `json:"proxy_id"`
	SelfLink          string         `json:"self_link"`
	SslCertificates   []string       `json:"ssl_certificates"`
	SslPolicy         string         `json:"ssl_policy"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
