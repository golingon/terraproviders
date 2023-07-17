// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datacomputebackendservice "github.com/golingon/terraproviders/googlebeta/4.73.1/datacomputebackendservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeBackendService creates a new instance of [DataComputeBackendService].
func NewDataComputeBackendService(name string, args DataComputeBackendServiceArgs) *DataComputeBackendService {
	return &DataComputeBackendService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeBackendService)(nil)

// DataComputeBackendService represents the Terraform data resource google_compute_backend_service.
type DataComputeBackendService struct {
	Name string
	Args DataComputeBackendServiceArgs
}

// DataSource returns the Terraform object type for [DataComputeBackendService].
func (cbs *DataComputeBackendService) DataSource() string {
	return "google_compute_backend_service"
}

// LocalName returns the local name for [DataComputeBackendService].
func (cbs *DataComputeBackendService) LocalName() string {
	return cbs.Name
}

// Configuration returns the configuration (args) for [DataComputeBackendService].
func (cbs *DataComputeBackendService) Configuration() interface{} {
	return cbs.Args
}

// Attributes returns the attributes for [DataComputeBackendService].
func (cbs *DataComputeBackendService) Attributes() dataComputeBackendServiceAttributes {
	return dataComputeBackendServiceAttributes{ref: terra.ReferenceDataResource(cbs)}
}

// DataComputeBackendServiceArgs contains the configurations for google_compute_backend_service.
type DataComputeBackendServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Backend: min=0
	Backend []datacomputebackendservice.Backend `hcl:"backend,block" validate:"min=0"`
	// CdnPolicy: min=0
	CdnPolicy []datacomputebackendservice.CdnPolicy `hcl:"cdn_policy,block" validate:"min=0"`
	// CircuitBreakers: min=0
	CircuitBreakers []datacomputebackendservice.CircuitBreakers `hcl:"circuit_breakers,block" validate:"min=0"`
	// ConsistentHash: min=0
	ConsistentHash []datacomputebackendservice.ConsistentHash `hcl:"consistent_hash,block" validate:"min=0"`
	// Iap: min=0
	Iap []datacomputebackendservice.Iap `hcl:"iap,block" validate:"min=0"`
	// LocalityLbPolicies: min=0
	LocalityLbPolicies []datacomputebackendservice.LocalityLbPolicies `hcl:"locality_lb_policies,block" validate:"min=0"`
	// LogConfig: min=0
	LogConfig []datacomputebackendservice.LogConfig `hcl:"log_config,block" validate:"min=0"`
	// OutlierDetection: min=0
	OutlierDetection []datacomputebackendservice.OutlierDetection `hcl:"outlier_detection,block" validate:"min=0"`
	// SecuritySettings: min=0
	SecuritySettings []datacomputebackendservice.SecuritySettings `hcl:"security_settings,block" validate:"min=0"`
}
type dataComputeBackendServiceAttributes struct {
	ref terra.Reference
}

// AffinityCookieTtlSec returns a reference to field affinity_cookie_ttl_sec of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) AffinityCookieTtlSec() terra.NumberValue {
	return terra.ReferenceAsNumber(cbs.ref.Append("affinity_cookie_ttl_sec"))
}

// CompressionMode returns a reference to field compression_mode of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) CompressionMode() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("compression_mode"))
}

// ConnectionDrainingTimeoutSec returns a reference to field connection_draining_timeout_sec of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) ConnectionDrainingTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(cbs.ref.Append("connection_draining_timeout_sec"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("creation_timestamp"))
}

// CustomRequestHeaders returns a reference to field custom_request_headers of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) CustomRequestHeaders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cbs.ref.Append("custom_request_headers"))
}

// CustomResponseHeaders returns a reference to field custom_response_headers of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) CustomResponseHeaders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cbs.ref.Append("custom_response_headers"))
}

// Description returns a reference to field description of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("description"))
}

// EdgeSecurityPolicy returns a reference to field edge_security_policy of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) EdgeSecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("edge_security_policy"))
}

// EnableCdn returns a reference to field enable_cdn of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) EnableCdn() terra.BoolValue {
	return terra.ReferenceAsBool(cbs.ref.Append("enable_cdn"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("fingerprint"))
}

// GeneratedId returns a reference to field generated_id of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) GeneratedId() terra.NumberValue {
	return terra.ReferenceAsNumber(cbs.ref.Append("generated_id"))
}

// HealthChecks returns a reference to field health_checks of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) HealthChecks() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cbs.ref.Append("health_checks"))
}

// Id returns a reference to field id of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("id"))
}

// LoadBalancingScheme returns a reference to field load_balancing_scheme of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) LoadBalancingScheme() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("load_balancing_scheme"))
}

// LocalityLbPolicy returns a reference to field locality_lb_policy of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) LocalityLbPolicy() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("locality_lb_policy"))
}

// Name returns a reference to field name of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("name"))
}

// PortName returns a reference to field port_name of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) PortName() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("port_name"))
}

// Project returns a reference to field project of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("project"))
}

// Protocol returns a reference to field protocol of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("protocol"))
}

// SecurityPolicy returns a reference to field security_policy of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) SecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("security_policy"))
}

// SelfLink returns a reference to field self_link of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("self_link"))
}

// SessionAffinity returns a reference to field session_affinity of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) SessionAffinity() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("session_affinity"))
}

// TimeoutSec returns a reference to field timeout_sec of google_compute_backend_service.
func (cbs dataComputeBackendServiceAttributes) TimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(cbs.ref.Append("timeout_sec"))
}

func (cbs dataComputeBackendServiceAttributes) Backend() terra.SetValue[datacomputebackendservice.BackendAttributes] {
	return terra.ReferenceAsSet[datacomputebackendservice.BackendAttributes](cbs.ref.Append("backend"))
}

func (cbs dataComputeBackendServiceAttributes) CdnPolicy() terra.ListValue[datacomputebackendservice.CdnPolicyAttributes] {
	return terra.ReferenceAsList[datacomputebackendservice.CdnPolicyAttributes](cbs.ref.Append("cdn_policy"))
}

func (cbs dataComputeBackendServiceAttributes) CircuitBreakers() terra.ListValue[datacomputebackendservice.CircuitBreakersAttributes] {
	return terra.ReferenceAsList[datacomputebackendservice.CircuitBreakersAttributes](cbs.ref.Append("circuit_breakers"))
}

func (cbs dataComputeBackendServiceAttributes) ConsistentHash() terra.ListValue[datacomputebackendservice.ConsistentHashAttributes] {
	return terra.ReferenceAsList[datacomputebackendservice.ConsistentHashAttributes](cbs.ref.Append("consistent_hash"))
}

func (cbs dataComputeBackendServiceAttributes) Iap() terra.ListValue[datacomputebackendservice.IapAttributes] {
	return terra.ReferenceAsList[datacomputebackendservice.IapAttributes](cbs.ref.Append("iap"))
}

func (cbs dataComputeBackendServiceAttributes) LocalityLbPolicies() terra.ListValue[datacomputebackendservice.LocalityLbPoliciesAttributes] {
	return terra.ReferenceAsList[datacomputebackendservice.LocalityLbPoliciesAttributes](cbs.ref.Append("locality_lb_policies"))
}

func (cbs dataComputeBackendServiceAttributes) LogConfig() terra.ListValue[datacomputebackendservice.LogConfigAttributes] {
	return terra.ReferenceAsList[datacomputebackendservice.LogConfigAttributes](cbs.ref.Append("log_config"))
}

func (cbs dataComputeBackendServiceAttributes) OutlierDetection() terra.ListValue[datacomputebackendservice.OutlierDetectionAttributes] {
	return terra.ReferenceAsList[datacomputebackendservice.OutlierDetectionAttributes](cbs.ref.Append("outlier_detection"))
}

func (cbs dataComputeBackendServiceAttributes) SecuritySettings() terra.ListValue[datacomputebackendservice.SecuritySettingsAttributes] {
	return terra.ReferenceAsList[datacomputebackendservice.SecuritySettingsAttributes](cbs.ref.Append("security_settings"))
}
