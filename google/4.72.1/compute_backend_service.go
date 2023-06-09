// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computebackendservice "github.com/golingon/terraproviders/google/4.72.1/computebackendservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeBackendService creates a new instance of [ComputeBackendService].
func NewComputeBackendService(name string, args ComputeBackendServiceArgs) *ComputeBackendService {
	return &ComputeBackendService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeBackendService)(nil)

// ComputeBackendService represents the Terraform resource google_compute_backend_service.
type ComputeBackendService struct {
	Name      string
	Args      ComputeBackendServiceArgs
	state     *computeBackendServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeBackendService].
func (cbs *ComputeBackendService) Type() string {
	return "google_compute_backend_service"
}

// LocalName returns the local name for [ComputeBackendService].
func (cbs *ComputeBackendService) LocalName() string {
	return cbs.Name
}

// Configuration returns the configuration (args) for [ComputeBackendService].
func (cbs *ComputeBackendService) Configuration() interface{} {
	return cbs.Args
}

// DependOn is used for other resources to depend on [ComputeBackendService].
func (cbs *ComputeBackendService) DependOn() terra.Reference {
	return terra.ReferenceResource(cbs)
}

// Dependencies returns the list of resources [ComputeBackendService] depends_on.
func (cbs *ComputeBackendService) Dependencies() terra.Dependencies {
	return cbs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeBackendService].
func (cbs *ComputeBackendService) LifecycleManagement() *terra.Lifecycle {
	return cbs.Lifecycle
}

// Attributes returns the attributes for [ComputeBackendService].
func (cbs *ComputeBackendService) Attributes() computeBackendServiceAttributes {
	return computeBackendServiceAttributes{ref: terra.ReferenceResource(cbs)}
}

// ImportState imports the given attribute values into [ComputeBackendService]'s state.
func (cbs *ComputeBackendService) ImportState(av io.Reader) error {
	cbs.state = &computeBackendServiceState{}
	if err := json.NewDecoder(av).Decode(cbs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cbs.Type(), cbs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeBackendService] has state.
func (cbs *ComputeBackendService) State() (*computeBackendServiceState, bool) {
	return cbs.state, cbs.state != nil
}

// StateMust returns the state for [ComputeBackendService]. Panics if the state is nil.
func (cbs *ComputeBackendService) StateMust() *computeBackendServiceState {
	if cbs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cbs.Type(), cbs.LocalName()))
	}
	return cbs.state
}

// ComputeBackendServiceArgs contains the configurations for google_compute_backend_service.
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
}
type computeBackendServiceAttributes struct {
	ref terra.Reference
}

// AffinityCookieTtlSec returns a reference to field affinity_cookie_ttl_sec of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) AffinityCookieTtlSec() terra.NumberValue {
	return terra.ReferenceAsNumber(cbs.ref.Append("affinity_cookie_ttl_sec"))
}

// CompressionMode returns a reference to field compression_mode of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) CompressionMode() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("compression_mode"))
}

// ConnectionDrainingTimeoutSec returns a reference to field connection_draining_timeout_sec of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) ConnectionDrainingTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(cbs.ref.Append("connection_draining_timeout_sec"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("creation_timestamp"))
}

// CustomRequestHeaders returns a reference to field custom_request_headers of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) CustomRequestHeaders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cbs.ref.Append("custom_request_headers"))
}

// CustomResponseHeaders returns a reference to field custom_response_headers of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) CustomResponseHeaders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cbs.ref.Append("custom_response_headers"))
}

// Description returns a reference to field description of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("description"))
}

// EdgeSecurityPolicy returns a reference to field edge_security_policy of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) EdgeSecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("edge_security_policy"))
}

// EnableCdn returns a reference to field enable_cdn of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) EnableCdn() terra.BoolValue {
	return terra.ReferenceAsBool(cbs.ref.Append("enable_cdn"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("fingerprint"))
}

// GeneratedId returns a reference to field generated_id of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) GeneratedId() terra.NumberValue {
	return terra.ReferenceAsNumber(cbs.ref.Append("generated_id"))
}

// HealthChecks returns a reference to field health_checks of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) HealthChecks() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cbs.ref.Append("health_checks"))
}

// Id returns a reference to field id of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("id"))
}

// LoadBalancingScheme returns a reference to field load_balancing_scheme of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) LoadBalancingScheme() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("load_balancing_scheme"))
}

// LocalityLbPolicy returns a reference to field locality_lb_policy of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) LocalityLbPolicy() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("locality_lb_policy"))
}

// Name returns a reference to field name of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("name"))
}

// PortName returns a reference to field port_name of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) PortName() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("port_name"))
}

// Project returns a reference to field project of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("project"))
}

// Protocol returns a reference to field protocol of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("protocol"))
}

// SecurityPolicy returns a reference to field security_policy of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) SecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("security_policy"))
}

// SelfLink returns a reference to field self_link of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("self_link"))
}

// SessionAffinity returns a reference to field session_affinity of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) SessionAffinity() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("session_affinity"))
}

// TimeoutSec returns a reference to field timeout_sec of google_compute_backend_service.
func (cbs computeBackendServiceAttributes) TimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(cbs.ref.Append("timeout_sec"))
}

func (cbs computeBackendServiceAttributes) Backend() terra.SetValue[computebackendservice.BackendAttributes] {
	return terra.ReferenceAsSet[computebackendservice.BackendAttributes](cbs.ref.Append("backend"))
}

func (cbs computeBackendServiceAttributes) CdnPolicy() terra.ListValue[computebackendservice.CdnPolicyAttributes] {
	return terra.ReferenceAsList[computebackendservice.CdnPolicyAttributes](cbs.ref.Append("cdn_policy"))
}

func (cbs computeBackendServiceAttributes) CircuitBreakers() terra.ListValue[computebackendservice.CircuitBreakersAttributes] {
	return terra.ReferenceAsList[computebackendservice.CircuitBreakersAttributes](cbs.ref.Append("circuit_breakers"))
}

func (cbs computeBackendServiceAttributes) ConsistentHash() terra.ListValue[computebackendservice.ConsistentHashAttributes] {
	return terra.ReferenceAsList[computebackendservice.ConsistentHashAttributes](cbs.ref.Append("consistent_hash"))
}

func (cbs computeBackendServiceAttributes) Iap() terra.ListValue[computebackendservice.IapAttributes] {
	return terra.ReferenceAsList[computebackendservice.IapAttributes](cbs.ref.Append("iap"))
}

func (cbs computeBackendServiceAttributes) LocalityLbPolicies() terra.ListValue[computebackendservice.LocalityLbPoliciesAttributes] {
	return terra.ReferenceAsList[computebackendservice.LocalityLbPoliciesAttributes](cbs.ref.Append("locality_lb_policies"))
}

func (cbs computeBackendServiceAttributes) LogConfig() terra.ListValue[computebackendservice.LogConfigAttributes] {
	return terra.ReferenceAsList[computebackendservice.LogConfigAttributes](cbs.ref.Append("log_config"))
}

func (cbs computeBackendServiceAttributes) OutlierDetection() terra.ListValue[computebackendservice.OutlierDetectionAttributes] {
	return terra.ReferenceAsList[computebackendservice.OutlierDetectionAttributes](cbs.ref.Append("outlier_detection"))
}

func (cbs computeBackendServiceAttributes) SecuritySettings() terra.ListValue[computebackendservice.SecuritySettingsAttributes] {
	return terra.ReferenceAsList[computebackendservice.SecuritySettingsAttributes](cbs.ref.Append("security_settings"))
}

func (cbs computeBackendServiceAttributes) Timeouts() computebackendservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computebackendservice.TimeoutsAttributes](cbs.ref.Append("timeouts"))
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
