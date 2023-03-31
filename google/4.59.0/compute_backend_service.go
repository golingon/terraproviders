// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computebackendservice "github.com/golingon/terraproviders/google/4.59.0/computebackendservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeBackendService(name string, args ComputeBackendServiceArgs) *ComputeBackendService {
	return &ComputeBackendService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeBackendService)(nil)

type ComputeBackendService struct {
	Name  string
	Args  ComputeBackendServiceArgs
	state *computeBackendServiceState
}

func (cbs *ComputeBackendService) Type() string {
	return "google_compute_backend_service"
}

func (cbs *ComputeBackendService) LocalName() string {
	return cbs.Name
}

func (cbs *ComputeBackendService) Configuration() interface{} {
	return cbs.Args
}

func (cbs *ComputeBackendService) Attributes() computeBackendServiceAttributes {
	return computeBackendServiceAttributes{ref: terra.ReferenceResource(cbs)}
}

func (cbs *ComputeBackendService) ImportState(av io.Reader) error {
	cbs.state = &computeBackendServiceState{}
	if err := json.NewDecoder(av).Decode(cbs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cbs.Type(), cbs.LocalName(), err)
	}
	return nil
}

func (cbs *ComputeBackendService) State() (*computeBackendServiceState, bool) {
	return cbs.state, cbs.state != nil
}

func (cbs *ComputeBackendService) StateMust() *computeBackendServiceState {
	if cbs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cbs.Type(), cbs.LocalName()))
	}
	return cbs.state
}

func (cbs *ComputeBackendService) DependOn() terra.Reference {
	return terra.ReferenceResource(cbs)
}

type ComputeBackendServiceArgs struct {
	// AffinityCookieTtlSec: number, optional
	AffinityCookieTtlSec terra.NumberValue `hcl:"affinity_cookie_ttl_sec,attr"`
	// CompressionMode: string, optional
	CompressionMode terra.StringValue `hcl:"compression_mode,attr"`
	// ConnectionDrainingTimeoutSec: number, optional
	ConnectionDrainingTimeoutSec terra.NumberValue `hcl:"connection_draining_timeout_sec,attr"`
	// CustomRequestHeaders: set of string, optional
	CustomRequestHeaders terra.SetValue[terra.StringValue] `hcl:"custom_request_headers,attr"`
	// CustomResponseHeaders: set of string, optional
	CustomResponseHeaders terra.SetValue[terra.StringValue] `hcl:"custom_response_headers,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EdgeSecurityPolicy: string, optional
	EdgeSecurityPolicy terra.StringValue `hcl:"edge_security_policy,attr"`
	// EnableCdn: bool, optional
	EnableCdn terra.BoolValue `hcl:"enable_cdn,attr"`
	// HealthChecks: set of string, optional
	HealthChecks terra.SetValue[terra.StringValue] `hcl:"health_checks,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LoadBalancingScheme: string, optional
	LoadBalancingScheme terra.StringValue `hcl:"load_balancing_scheme,attr"`
	// LocalityLbPolicy: string, optional
	LocalityLbPolicy terra.StringValue `hcl:"locality_lb_policy,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PortName: string, optional
	PortName terra.StringValue `hcl:"port_name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// SecurityPolicy: string, optional
	SecurityPolicy terra.StringValue `hcl:"security_policy,attr"`
	// SessionAffinity: string, optional
	SessionAffinity terra.StringValue `hcl:"session_affinity,attr"`
	// TimeoutSec: number, optional
	TimeoutSec terra.NumberValue `hcl:"timeout_sec,attr"`
	// Backend: min=0
	Backend []computebackendservice.Backend `hcl:"backend,block" validate:"min=0"`
	// CdnPolicy: optional
	CdnPolicy *computebackendservice.CdnPolicy `hcl:"cdn_policy,block"`
	// CircuitBreakers: optional
	CircuitBreakers *computebackendservice.CircuitBreakers `hcl:"circuit_breakers,block"`
	// ConsistentHash: optional
	ConsistentHash *computebackendservice.ConsistentHash `hcl:"consistent_hash,block"`
	// Iap: optional
	Iap *computebackendservice.Iap `hcl:"iap,block"`
	// LocalityLbPolicies: min=0
	LocalityLbPolicies []computebackendservice.LocalityLbPolicies `hcl:"locality_lb_policies,block" validate:"min=0"`
	// LogConfig: optional
	LogConfig *computebackendservice.LogConfig `hcl:"log_config,block"`
	// OutlierDetection: optional
	OutlierDetection *computebackendservice.OutlierDetection `hcl:"outlier_detection,block"`
	// SecuritySettings: optional
	SecuritySettings *computebackendservice.SecuritySettings `hcl:"security_settings,block"`
	// Timeouts: optional
	Timeouts *computebackendservice.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ComputeBackendService depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeBackendServiceAttributes struct {
	ref terra.Reference
}

func (cbs computeBackendServiceAttributes) AffinityCookieTtlSec() terra.NumberValue {
	return terra.ReferenceNumber(cbs.ref.Append("affinity_cookie_ttl_sec"))
}

func (cbs computeBackendServiceAttributes) CompressionMode() terra.StringValue {
	return terra.ReferenceString(cbs.ref.Append("compression_mode"))
}

func (cbs computeBackendServiceAttributes) ConnectionDrainingTimeoutSec() terra.NumberValue {
	return terra.ReferenceNumber(cbs.ref.Append("connection_draining_timeout_sec"))
}

func (cbs computeBackendServiceAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceString(cbs.ref.Append("creation_timestamp"))
}

func (cbs computeBackendServiceAttributes) CustomRequestHeaders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cbs.ref.Append("custom_request_headers"))
}

func (cbs computeBackendServiceAttributes) CustomResponseHeaders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cbs.ref.Append("custom_response_headers"))
}

func (cbs computeBackendServiceAttributes) Description() terra.StringValue {
	return terra.ReferenceString(cbs.ref.Append("description"))
}

func (cbs computeBackendServiceAttributes) EdgeSecurityPolicy() terra.StringValue {
	return terra.ReferenceString(cbs.ref.Append("edge_security_policy"))
}

func (cbs computeBackendServiceAttributes) EnableCdn() terra.BoolValue {
	return terra.ReferenceBool(cbs.ref.Append("enable_cdn"))
}

func (cbs computeBackendServiceAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceString(cbs.ref.Append("fingerprint"))
}

func (cbs computeBackendServiceAttributes) GeneratedId() terra.NumberValue {
	return terra.ReferenceNumber(cbs.ref.Append("generated_id"))
}

func (cbs computeBackendServiceAttributes) HealthChecks() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cbs.ref.Append("health_checks"))
}

func (cbs computeBackendServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cbs.ref.Append("id"))
}

func (cbs computeBackendServiceAttributes) LoadBalancingScheme() terra.StringValue {
	return terra.ReferenceString(cbs.ref.Append("load_balancing_scheme"))
}

func (cbs computeBackendServiceAttributes) LocalityLbPolicy() terra.StringValue {
	return terra.ReferenceString(cbs.ref.Append("locality_lb_policy"))
}

func (cbs computeBackendServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cbs.ref.Append("name"))
}

func (cbs computeBackendServiceAttributes) PortName() terra.StringValue {
	return terra.ReferenceString(cbs.ref.Append("port_name"))
}

func (cbs computeBackendServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceString(cbs.ref.Append("project"))
}

func (cbs computeBackendServiceAttributes) Protocol() terra.StringValue {
	return terra.ReferenceString(cbs.ref.Append("protocol"))
}

func (cbs computeBackendServiceAttributes) SecurityPolicy() terra.StringValue {
	return terra.ReferenceString(cbs.ref.Append("security_policy"))
}

func (cbs computeBackendServiceAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(cbs.ref.Append("self_link"))
}

func (cbs computeBackendServiceAttributes) SessionAffinity() terra.StringValue {
	return terra.ReferenceString(cbs.ref.Append("session_affinity"))
}

func (cbs computeBackendServiceAttributes) TimeoutSec() terra.NumberValue {
	return terra.ReferenceNumber(cbs.ref.Append("timeout_sec"))
}

func (cbs computeBackendServiceAttributes) Backend() terra.SetValue[computebackendservice.BackendAttributes] {
	return terra.ReferenceSet[computebackendservice.BackendAttributes](cbs.ref.Append("backend"))
}

func (cbs computeBackendServiceAttributes) CdnPolicy() terra.ListValue[computebackendservice.CdnPolicyAttributes] {
	return terra.ReferenceList[computebackendservice.CdnPolicyAttributes](cbs.ref.Append("cdn_policy"))
}

func (cbs computeBackendServiceAttributes) CircuitBreakers() terra.ListValue[computebackendservice.CircuitBreakersAttributes] {
	return terra.ReferenceList[computebackendservice.CircuitBreakersAttributes](cbs.ref.Append("circuit_breakers"))
}

func (cbs computeBackendServiceAttributes) ConsistentHash() terra.ListValue[computebackendservice.ConsistentHashAttributes] {
	return terra.ReferenceList[computebackendservice.ConsistentHashAttributes](cbs.ref.Append("consistent_hash"))
}

func (cbs computeBackendServiceAttributes) Iap() terra.ListValue[computebackendservice.IapAttributes] {
	return terra.ReferenceList[computebackendservice.IapAttributes](cbs.ref.Append("iap"))
}

func (cbs computeBackendServiceAttributes) LocalityLbPolicies() terra.ListValue[computebackendservice.LocalityLbPoliciesAttributes] {
	return terra.ReferenceList[computebackendservice.LocalityLbPoliciesAttributes](cbs.ref.Append("locality_lb_policies"))
}

func (cbs computeBackendServiceAttributes) LogConfig() terra.ListValue[computebackendservice.LogConfigAttributes] {
	return terra.ReferenceList[computebackendservice.LogConfigAttributes](cbs.ref.Append("log_config"))
}

func (cbs computeBackendServiceAttributes) OutlierDetection() terra.ListValue[computebackendservice.OutlierDetectionAttributes] {
	return terra.ReferenceList[computebackendservice.OutlierDetectionAttributes](cbs.ref.Append("outlier_detection"))
}

func (cbs computeBackendServiceAttributes) SecuritySettings() terra.ListValue[computebackendservice.SecuritySettingsAttributes] {
	return terra.ReferenceList[computebackendservice.SecuritySettingsAttributes](cbs.ref.Append("security_settings"))
}

func (cbs computeBackendServiceAttributes) Timeouts() computebackendservice.TimeoutsAttributes {
	return terra.ReferenceSingle[computebackendservice.TimeoutsAttributes](cbs.ref.Append("timeouts"))
}

type computeBackendServiceState struct {
	AffinityCookieTtlSec         float64                                         `json:"affinity_cookie_ttl_sec"`
	CompressionMode              string                                          `json:"compression_mode"`
	ConnectionDrainingTimeoutSec float64                                         `json:"connection_draining_timeout_sec"`
	CreationTimestamp            string                                          `json:"creation_timestamp"`
	CustomRequestHeaders         []string                                        `json:"custom_request_headers"`
	CustomResponseHeaders        []string                                        `json:"custom_response_headers"`
	Description                  string                                          `json:"description"`
	EdgeSecurityPolicy           string                                          `json:"edge_security_policy"`
	EnableCdn                    bool                                            `json:"enable_cdn"`
	Fingerprint                  string                                          `json:"fingerprint"`
	GeneratedId                  float64                                         `json:"generated_id"`
	HealthChecks                 []string                                        `json:"health_checks"`
	Id                           string                                          `json:"id"`
	LoadBalancingScheme          string                                          `json:"load_balancing_scheme"`
	LocalityLbPolicy             string                                          `json:"locality_lb_policy"`
	Name                         string                                          `json:"name"`
	PortName                     string                                          `json:"port_name"`
	Project                      string                                          `json:"project"`
	Protocol                     string                                          `json:"protocol"`
	SecurityPolicy               string                                          `json:"security_policy"`
	SelfLink                     string                                          `json:"self_link"`
	SessionAffinity              string                                          `json:"session_affinity"`
	TimeoutSec                   float64                                         `json:"timeout_sec"`
	Backend                      []computebackendservice.BackendState            `json:"backend"`
	CdnPolicy                    []computebackendservice.CdnPolicyState          `json:"cdn_policy"`
	CircuitBreakers              []computebackendservice.CircuitBreakersState    `json:"circuit_breakers"`
	ConsistentHash               []computebackendservice.ConsistentHashState     `json:"consistent_hash"`
	Iap                          []computebackendservice.IapState                `json:"iap"`
	LocalityLbPolicies           []computebackendservice.LocalityLbPoliciesState `json:"locality_lb_policies"`
	LogConfig                    []computebackendservice.LogConfigState          `json:"log_config"`
	OutlierDetection             []computebackendservice.OutlierDetectionState   `json:"outlier_detection"`
	SecuritySettings             []computebackendservice.SecuritySettingsState   `json:"security_settings"`
	Timeouts                     *computebackendservice.TimeoutsState            `json:"timeouts"`
}
