// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeregionbackendservice "github.com/golingon/terraproviders/googlebeta/4.74.0/computeregionbackendservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionBackendService creates a new instance of [ComputeRegionBackendService].
func NewComputeRegionBackendService(name string, args ComputeRegionBackendServiceArgs) *ComputeRegionBackendService {
	return &ComputeRegionBackendService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionBackendService)(nil)

// ComputeRegionBackendService represents the Terraform resource google_compute_region_backend_service.
type ComputeRegionBackendService struct {
	Name      string
	Args      ComputeRegionBackendServiceArgs
	state     *computeRegionBackendServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionBackendService].
func (crbs *ComputeRegionBackendService) Type() string {
	return "google_compute_region_backend_service"
}

// LocalName returns the local name for [ComputeRegionBackendService].
func (crbs *ComputeRegionBackendService) LocalName() string {
	return crbs.Name
}

// Configuration returns the configuration (args) for [ComputeRegionBackendService].
func (crbs *ComputeRegionBackendService) Configuration() interface{} {
	return crbs.Args
}

// DependOn is used for other resources to depend on [ComputeRegionBackendService].
func (crbs *ComputeRegionBackendService) DependOn() terra.Reference {
	return terra.ReferenceResource(crbs)
}

// Dependencies returns the list of resources [ComputeRegionBackendService] depends_on.
func (crbs *ComputeRegionBackendService) Dependencies() terra.Dependencies {
	return crbs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionBackendService].
func (crbs *ComputeRegionBackendService) LifecycleManagement() *terra.Lifecycle {
	return crbs.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionBackendService].
func (crbs *ComputeRegionBackendService) Attributes() computeRegionBackendServiceAttributes {
	return computeRegionBackendServiceAttributes{ref: terra.ReferenceResource(crbs)}
}

// ImportState imports the given attribute values into [ComputeRegionBackendService]'s state.
func (crbs *ComputeRegionBackendService) ImportState(av io.Reader) error {
	crbs.state = &computeRegionBackendServiceState{}
	if err := json.NewDecoder(av).Decode(crbs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crbs.Type(), crbs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionBackendService] has state.
func (crbs *ComputeRegionBackendService) State() (*computeRegionBackendServiceState, bool) {
	return crbs.state, crbs.state != nil
}

// StateMust returns the state for [ComputeRegionBackendService]. Panics if the state is nil.
func (crbs *ComputeRegionBackendService) StateMust() *computeRegionBackendServiceState {
	if crbs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crbs.Type(), crbs.LocalName()))
	}
	return crbs.state
}

// ComputeRegionBackendServiceArgs contains the configurations for google_compute_region_backend_service.
type ComputeRegionBackendServiceArgs struct {
	// AffinityCookieTtlSec: number, optional
	AffinityCookieTtlSec terra.NumberValue `hcl:"affinity_cookie_ttl_sec,attr"`
	// ConnectionDrainingTimeoutSec: number, optional
	ConnectionDrainingTimeoutSec terra.NumberValue `hcl:"connection_draining_timeout_sec,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
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
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// PortName: string, optional
	PortName terra.StringValue `hcl:"port_name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// SessionAffinity: string, optional
	SessionAffinity terra.StringValue `hcl:"session_affinity,attr"`
	// TimeoutSec: number, optional
	TimeoutSec terra.NumberValue `hcl:"timeout_sec,attr"`
	// Backend: min=0
	Backend []computeregionbackendservice.Backend `hcl:"backend,block" validate:"min=0"`
	// CdnPolicy: optional
	CdnPolicy *computeregionbackendservice.CdnPolicy `hcl:"cdn_policy,block"`
	// CircuitBreakers: optional
	CircuitBreakers *computeregionbackendservice.CircuitBreakers `hcl:"circuit_breakers,block"`
	// ConnectionTrackingPolicy: optional
	ConnectionTrackingPolicy *computeregionbackendservice.ConnectionTrackingPolicy `hcl:"connection_tracking_policy,block"`
	// ConsistentHash: optional
	ConsistentHash *computeregionbackendservice.ConsistentHash `hcl:"consistent_hash,block"`
	// FailoverPolicy: optional
	FailoverPolicy *computeregionbackendservice.FailoverPolicy `hcl:"failover_policy,block"`
	// Iap: optional
	Iap *computeregionbackendservice.Iap `hcl:"iap,block"`
	// LogConfig: optional
	LogConfig *computeregionbackendservice.LogConfig `hcl:"log_config,block"`
	// OutlierDetection: optional
	OutlierDetection *computeregionbackendservice.OutlierDetection `hcl:"outlier_detection,block"`
	// Subsetting: optional
	Subsetting *computeregionbackendservice.Subsetting `hcl:"subsetting,block"`
	// Timeouts: optional
	Timeouts *computeregionbackendservice.Timeouts `hcl:"timeouts,block"`
}
type computeRegionBackendServiceAttributes struct {
	ref terra.Reference
}

// AffinityCookieTtlSec returns a reference to field affinity_cookie_ttl_sec of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) AffinityCookieTtlSec() terra.NumberValue {
	return terra.ReferenceAsNumber(crbs.ref.Append("affinity_cookie_ttl_sec"))
}

// ConnectionDrainingTimeoutSec returns a reference to field connection_draining_timeout_sec of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) ConnectionDrainingTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(crbs.ref.Append("connection_draining_timeout_sec"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(crbs.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crbs.ref.Append("description"))
}

// EnableCdn returns a reference to field enable_cdn of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) EnableCdn() terra.BoolValue {
	return terra.ReferenceAsBool(crbs.ref.Append("enable_cdn"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(crbs.ref.Append("fingerprint"))
}

// HealthChecks returns a reference to field health_checks of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) HealthChecks() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crbs.ref.Append("health_checks"))
}

// Id returns a reference to field id of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crbs.ref.Append("id"))
}

// LoadBalancingScheme returns a reference to field load_balancing_scheme of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) LoadBalancingScheme() terra.StringValue {
	return terra.ReferenceAsString(crbs.ref.Append("load_balancing_scheme"))
}

// LocalityLbPolicy returns a reference to field locality_lb_policy of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) LocalityLbPolicy() terra.StringValue {
	return terra.ReferenceAsString(crbs.ref.Append("locality_lb_policy"))
}

// Name returns a reference to field name of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crbs.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(crbs.ref.Append("network"))
}

// PortName returns a reference to field port_name of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) PortName() terra.StringValue {
	return terra.ReferenceAsString(crbs.ref.Append("port_name"))
}

// Project returns a reference to field project of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crbs.ref.Append("project"))
}

// Protocol returns a reference to field protocol of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(crbs.ref.Append("protocol"))
}

// Region returns a reference to field region of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crbs.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crbs.ref.Append("self_link"))
}

// SessionAffinity returns a reference to field session_affinity of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) SessionAffinity() terra.StringValue {
	return terra.ReferenceAsString(crbs.ref.Append("session_affinity"))
}

// TimeoutSec returns a reference to field timeout_sec of google_compute_region_backend_service.
func (crbs computeRegionBackendServiceAttributes) TimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(crbs.ref.Append("timeout_sec"))
}

func (crbs computeRegionBackendServiceAttributes) Backend() terra.SetValue[computeregionbackendservice.BackendAttributes] {
	return terra.ReferenceAsSet[computeregionbackendservice.BackendAttributes](crbs.ref.Append("backend"))
}

func (crbs computeRegionBackendServiceAttributes) CdnPolicy() terra.ListValue[computeregionbackendservice.CdnPolicyAttributes] {
	return terra.ReferenceAsList[computeregionbackendservice.CdnPolicyAttributes](crbs.ref.Append("cdn_policy"))
}

func (crbs computeRegionBackendServiceAttributes) CircuitBreakers() terra.ListValue[computeregionbackendservice.CircuitBreakersAttributes] {
	return terra.ReferenceAsList[computeregionbackendservice.CircuitBreakersAttributes](crbs.ref.Append("circuit_breakers"))
}

func (crbs computeRegionBackendServiceAttributes) ConnectionTrackingPolicy() terra.ListValue[computeregionbackendservice.ConnectionTrackingPolicyAttributes] {
	return terra.ReferenceAsList[computeregionbackendservice.ConnectionTrackingPolicyAttributes](crbs.ref.Append("connection_tracking_policy"))
}

func (crbs computeRegionBackendServiceAttributes) ConsistentHash() terra.ListValue[computeregionbackendservice.ConsistentHashAttributes] {
	return terra.ReferenceAsList[computeregionbackendservice.ConsistentHashAttributes](crbs.ref.Append("consistent_hash"))
}

func (crbs computeRegionBackendServiceAttributes) FailoverPolicy() terra.ListValue[computeregionbackendservice.FailoverPolicyAttributes] {
	return terra.ReferenceAsList[computeregionbackendservice.FailoverPolicyAttributes](crbs.ref.Append("failover_policy"))
}

func (crbs computeRegionBackendServiceAttributes) Iap() terra.ListValue[computeregionbackendservice.IapAttributes] {
	return terra.ReferenceAsList[computeregionbackendservice.IapAttributes](crbs.ref.Append("iap"))
}

func (crbs computeRegionBackendServiceAttributes) LogConfig() terra.ListValue[computeregionbackendservice.LogConfigAttributes] {
	return terra.ReferenceAsList[computeregionbackendservice.LogConfigAttributes](crbs.ref.Append("log_config"))
}

func (crbs computeRegionBackendServiceAttributes) OutlierDetection() terra.ListValue[computeregionbackendservice.OutlierDetectionAttributes] {
	return terra.ReferenceAsList[computeregionbackendservice.OutlierDetectionAttributes](crbs.ref.Append("outlier_detection"))
}

func (crbs computeRegionBackendServiceAttributes) Subsetting() terra.ListValue[computeregionbackendservice.SubsettingAttributes] {
	return terra.ReferenceAsList[computeregionbackendservice.SubsettingAttributes](crbs.ref.Append("subsetting"))
}

func (crbs computeRegionBackendServiceAttributes) Timeouts() computeregionbackendservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregionbackendservice.TimeoutsAttributes](crbs.ref.Append("timeouts"))
}

type computeRegionBackendServiceState struct {
	AffinityCookieTtlSec         float64                                                     `json:"affinity_cookie_ttl_sec"`
	ConnectionDrainingTimeoutSec float64                                                     `json:"connection_draining_timeout_sec"`
	CreationTimestamp            string                                                      `json:"creation_timestamp"`
	Description                  string                                                      `json:"description"`
	EnableCdn                    bool                                                        `json:"enable_cdn"`
	Fingerprint                  string                                                      `json:"fingerprint"`
	HealthChecks                 []string                                                    `json:"health_checks"`
	Id                           string                                                      `json:"id"`
	LoadBalancingScheme          string                                                      `json:"load_balancing_scheme"`
	LocalityLbPolicy             string                                                      `json:"locality_lb_policy"`
	Name                         string                                                      `json:"name"`
	Network                      string                                                      `json:"network"`
	PortName                     string                                                      `json:"port_name"`
	Project                      string                                                      `json:"project"`
	Protocol                     string                                                      `json:"protocol"`
	Region                       string                                                      `json:"region"`
	SelfLink                     string                                                      `json:"self_link"`
	SessionAffinity              string                                                      `json:"session_affinity"`
	TimeoutSec                   float64                                                     `json:"timeout_sec"`
	Backend                      []computeregionbackendservice.BackendState                  `json:"backend"`
	CdnPolicy                    []computeregionbackendservice.CdnPolicyState                `json:"cdn_policy"`
	CircuitBreakers              []computeregionbackendservice.CircuitBreakersState          `json:"circuit_breakers"`
	ConnectionTrackingPolicy     []computeregionbackendservice.ConnectionTrackingPolicyState `json:"connection_tracking_policy"`
	ConsistentHash               []computeregionbackendservice.ConsistentHashState           `json:"consistent_hash"`
	FailoverPolicy               []computeregionbackendservice.FailoverPolicyState           `json:"failover_policy"`
	Iap                          []computeregionbackendservice.IapState                      `json:"iap"`
	LogConfig                    []computeregionbackendservice.LogConfigState                `json:"log_config"`
	OutlierDetection             []computeregionbackendservice.OutlierDetectionState         `json:"outlier_detection"`
	Subsetting                   []computeregionbackendservice.SubsettingState               `json:"subsetting"`
	Timeouts                     *computeregionbackendservice.TimeoutsState                  `json:"timeouts"`
}
