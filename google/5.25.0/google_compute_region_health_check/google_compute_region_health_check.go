// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_region_health_check

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

// Resource represents the Terraform resource google_compute_region_health_check.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeRegionHealthCheckState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcrhc *Resource) Type() string {
	return "google_compute_region_health_check"
}

// LocalName returns the local name for [Resource].
func (gcrhc *Resource) LocalName() string {
	return gcrhc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcrhc *Resource) Configuration() interface{} {
	return gcrhc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcrhc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcrhc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcrhc *Resource) Dependencies() terra.Dependencies {
	return gcrhc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcrhc *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcrhc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcrhc *Resource) Attributes() googleComputeRegionHealthCheckAttributes {
	return googleComputeRegionHealthCheckAttributes{ref: terra.ReferenceResource(gcrhc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcrhc *Resource) ImportState(state io.Reader) error {
	gcrhc.state = &googleComputeRegionHealthCheckState{}
	if err := json.NewDecoder(state).Decode(gcrhc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcrhc.Type(), gcrhc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcrhc *Resource) State() (*googleComputeRegionHealthCheckState, bool) {
	return gcrhc.state, gcrhc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcrhc *Resource) StateMust() *googleComputeRegionHealthCheckState {
	if gcrhc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcrhc.Type(), gcrhc.LocalName()))
	}
	return gcrhc.state
}

// Args contains the configurations for google_compute_region_health_check.
type Args struct {
	// CheckIntervalSec: number, optional
	CheckIntervalSec terra.NumberValue `hcl:"check_interval_sec,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// HealthyThreshold: number, optional
	HealthyThreshold terra.NumberValue `hcl:"healthy_threshold,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// TimeoutSec: number, optional
	TimeoutSec terra.NumberValue `hcl:"timeout_sec,attr"`
	// UnhealthyThreshold: number, optional
	UnhealthyThreshold terra.NumberValue `hcl:"unhealthy_threshold,attr"`
	// GrpcHealthCheck: optional
	GrpcHealthCheck *GrpcHealthCheck `hcl:"grpc_health_check,block"`
	// Http2HealthCheck: optional
	Http2HealthCheck *Http2HealthCheck `hcl:"http2_health_check,block"`
	// HttpHealthCheck: optional
	HttpHealthCheck *HttpHealthCheck `hcl:"http_health_check,block"`
	// HttpsHealthCheck: optional
	HttpsHealthCheck *HttpsHealthCheck `hcl:"https_health_check,block"`
	// LogConfig: optional
	LogConfig *LogConfig `hcl:"log_config,block"`
	// SslHealthCheck: optional
	SslHealthCheck *SslHealthCheck `hcl:"ssl_health_check,block"`
	// TcpHealthCheck: optional
	TcpHealthCheck *TcpHealthCheck `hcl:"tcp_health_check,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleComputeRegionHealthCheckAttributes struct {
	ref terra.Reference
}

// CheckIntervalSec returns a reference to field check_interval_sec of google_compute_region_health_check.
func (gcrhc googleComputeRegionHealthCheckAttributes) CheckIntervalSec() terra.NumberValue {
	return terra.ReferenceAsNumber(gcrhc.ref.Append("check_interval_sec"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_region_health_check.
func (gcrhc googleComputeRegionHealthCheckAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(gcrhc.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_region_health_check.
func (gcrhc googleComputeRegionHealthCheckAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gcrhc.ref.Append("description"))
}

// HealthyThreshold returns a reference to field healthy_threshold of google_compute_region_health_check.
func (gcrhc googleComputeRegionHealthCheckAttributes) HealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(gcrhc.ref.Append("healthy_threshold"))
}

// Id returns a reference to field id of google_compute_region_health_check.
func (gcrhc googleComputeRegionHealthCheckAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcrhc.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_health_check.
func (gcrhc googleComputeRegionHealthCheckAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcrhc.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_health_check.
func (gcrhc googleComputeRegionHealthCheckAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcrhc.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_health_check.
func (gcrhc googleComputeRegionHealthCheckAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gcrhc.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_health_check.
func (gcrhc googleComputeRegionHealthCheckAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gcrhc.ref.Append("self_link"))
}

// TimeoutSec returns a reference to field timeout_sec of google_compute_region_health_check.
func (gcrhc googleComputeRegionHealthCheckAttributes) TimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(gcrhc.ref.Append("timeout_sec"))
}

// Type returns a reference to field type of google_compute_region_health_check.
func (gcrhc googleComputeRegionHealthCheckAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(gcrhc.ref.Append("type"))
}

// UnhealthyThreshold returns a reference to field unhealthy_threshold of google_compute_region_health_check.
func (gcrhc googleComputeRegionHealthCheckAttributes) UnhealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(gcrhc.ref.Append("unhealthy_threshold"))
}

func (gcrhc googleComputeRegionHealthCheckAttributes) GrpcHealthCheck() terra.ListValue[GrpcHealthCheckAttributes] {
	return terra.ReferenceAsList[GrpcHealthCheckAttributes](gcrhc.ref.Append("grpc_health_check"))
}

func (gcrhc googleComputeRegionHealthCheckAttributes) Http2HealthCheck() terra.ListValue[Http2HealthCheckAttributes] {
	return terra.ReferenceAsList[Http2HealthCheckAttributes](gcrhc.ref.Append("http2_health_check"))
}

func (gcrhc googleComputeRegionHealthCheckAttributes) HttpHealthCheck() terra.ListValue[HttpHealthCheckAttributes] {
	return terra.ReferenceAsList[HttpHealthCheckAttributes](gcrhc.ref.Append("http_health_check"))
}

func (gcrhc googleComputeRegionHealthCheckAttributes) HttpsHealthCheck() terra.ListValue[HttpsHealthCheckAttributes] {
	return terra.ReferenceAsList[HttpsHealthCheckAttributes](gcrhc.ref.Append("https_health_check"))
}

func (gcrhc googleComputeRegionHealthCheckAttributes) LogConfig() terra.ListValue[LogConfigAttributes] {
	return terra.ReferenceAsList[LogConfigAttributes](gcrhc.ref.Append("log_config"))
}

func (gcrhc googleComputeRegionHealthCheckAttributes) SslHealthCheck() terra.ListValue[SslHealthCheckAttributes] {
	return terra.ReferenceAsList[SslHealthCheckAttributes](gcrhc.ref.Append("ssl_health_check"))
}

func (gcrhc googleComputeRegionHealthCheckAttributes) TcpHealthCheck() terra.ListValue[TcpHealthCheckAttributes] {
	return terra.ReferenceAsList[TcpHealthCheckAttributes](gcrhc.ref.Append("tcp_health_check"))
}

func (gcrhc googleComputeRegionHealthCheckAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcrhc.ref.Append("timeouts"))
}

type googleComputeRegionHealthCheckState struct {
	CheckIntervalSec   float64                 `json:"check_interval_sec"`
	CreationTimestamp  string                  `json:"creation_timestamp"`
	Description        string                  `json:"description"`
	HealthyThreshold   float64                 `json:"healthy_threshold"`
	Id                 string                  `json:"id"`
	Name               string                  `json:"name"`
	Project            string                  `json:"project"`
	Region             string                  `json:"region"`
	SelfLink           string                  `json:"self_link"`
	TimeoutSec         float64                 `json:"timeout_sec"`
	Type               string                  `json:"type"`
	UnhealthyThreshold float64                 `json:"unhealthy_threshold"`
	GrpcHealthCheck    []GrpcHealthCheckState  `json:"grpc_health_check"`
	Http2HealthCheck   []Http2HealthCheckState `json:"http2_health_check"`
	HttpHealthCheck    []HttpHealthCheckState  `json:"http_health_check"`
	HttpsHealthCheck   []HttpsHealthCheckState `json:"https_health_check"`
	LogConfig          []LogConfigState        `json:"log_config"`
	SslHealthCheck     []SslHealthCheckState   `json:"ssl_health_check"`
	TcpHealthCheck     []TcpHealthCheckState   `json:"tcp_health_check"`
	Timeouts           *TimeoutsState          `json:"timeouts"`
}
