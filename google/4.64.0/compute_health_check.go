// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computehealthcheck "github.com/golingon/terraproviders/google/4.64.0/computehealthcheck"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeHealthCheck creates a new instance of [ComputeHealthCheck].
func NewComputeHealthCheck(name string, args ComputeHealthCheckArgs) *ComputeHealthCheck {
	return &ComputeHealthCheck{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeHealthCheck)(nil)

// ComputeHealthCheck represents the Terraform resource google_compute_health_check.
type ComputeHealthCheck struct {
	Name      string
	Args      ComputeHealthCheckArgs
	state     *computeHealthCheckState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeHealthCheck].
func (chc *ComputeHealthCheck) Type() string {
	return "google_compute_health_check"
}

// LocalName returns the local name for [ComputeHealthCheck].
func (chc *ComputeHealthCheck) LocalName() string {
	return chc.Name
}

// Configuration returns the configuration (args) for [ComputeHealthCheck].
func (chc *ComputeHealthCheck) Configuration() interface{} {
	return chc.Args
}

// DependOn is used for other resources to depend on [ComputeHealthCheck].
func (chc *ComputeHealthCheck) DependOn() terra.Reference {
	return terra.ReferenceResource(chc)
}

// Dependencies returns the list of resources [ComputeHealthCheck] depends_on.
func (chc *ComputeHealthCheck) Dependencies() terra.Dependencies {
	return chc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeHealthCheck].
func (chc *ComputeHealthCheck) LifecycleManagement() *terra.Lifecycle {
	return chc.Lifecycle
}

// Attributes returns the attributes for [ComputeHealthCheck].
func (chc *ComputeHealthCheck) Attributes() computeHealthCheckAttributes {
	return computeHealthCheckAttributes{ref: terra.ReferenceResource(chc)}
}

// ImportState imports the given attribute values into [ComputeHealthCheck]'s state.
func (chc *ComputeHealthCheck) ImportState(av io.Reader) error {
	chc.state = &computeHealthCheckState{}
	if err := json.NewDecoder(av).Decode(chc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", chc.Type(), chc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeHealthCheck] has state.
func (chc *ComputeHealthCheck) State() (*computeHealthCheckState, bool) {
	return chc.state, chc.state != nil
}

// StateMust returns the state for [ComputeHealthCheck]. Panics if the state is nil.
func (chc *ComputeHealthCheck) StateMust() *computeHealthCheckState {
	if chc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", chc.Type(), chc.LocalName()))
	}
	return chc.state
}

// ComputeHealthCheckArgs contains the configurations for google_compute_health_check.
type ComputeHealthCheckArgs struct {
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
	// TimeoutSec: number, optional
	TimeoutSec terra.NumberValue `hcl:"timeout_sec,attr"`
	// UnhealthyThreshold: number, optional
	UnhealthyThreshold terra.NumberValue `hcl:"unhealthy_threshold,attr"`
	// GrpcHealthCheck: optional
	GrpcHealthCheck *computehealthcheck.GrpcHealthCheck `hcl:"grpc_health_check,block"`
	// Http2HealthCheck: optional
	Http2HealthCheck *computehealthcheck.Http2HealthCheck `hcl:"http2_health_check,block"`
	// HttpHealthCheck: optional
	HttpHealthCheck *computehealthcheck.HttpHealthCheck `hcl:"http_health_check,block"`
	// HttpsHealthCheck: optional
	HttpsHealthCheck *computehealthcheck.HttpsHealthCheck `hcl:"https_health_check,block"`
	// LogConfig: optional
	LogConfig *computehealthcheck.LogConfig `hcl:"log_config,block"`
	// SslHealthCheck: optional
	SslHealthCheck *computehealthcheck.SslHealthCheck `hcl:"ssl_health_check,block"`
	// TcpHealthCheck: optional
	TcpHealthCheck *computehealthcheck.TcpHealthCheck `hcl:"tcp_health_check,block"`
	// Timeouts: optional
	Timeouts *computehealthcheck.Timeouts `hcl:"timeouts,block"`
}
type computeHealthCheckAttributes struct {
	ref terra.Reference
}

// CheckIntervalSec returns a reference to field check_interval_sec of google_compute_health_check.
func (chc computeHealthCheckAttributes) CheckIntervalSec() terra.NumberValue {
	return terra.ReferenceAsNumber(chc.ref.Append("check_interval_sec"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_health_check.
func (chc computeHealthCheckAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(chc.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_health_check.
func (chc computeHealthCheckAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(chc.ref.Append("description"))
}

// HealthyThreshold returns a reference to field healthy_threshold of google_compute_health_check.
func (chc computeHealthCheckAttributes) HealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(chc.ref.Append("healthy_threshold"))
}

// Id returns a reference to field id of google_compute_health_check.
func (chc computeHealthCheckAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(chc.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_health_check.
func (chc computeHealthCheckAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(chc.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_health_check.
func (chc computeHealthCheckAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(chc.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_health_check.
func (chc computeHealthCheckAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(chc.ref.Append("self_link"))
}

// TimeoutSec returns a reference to field timeout_sec of google_compute_health_check.
func (chc computeHealthCheckAttributes) TimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(chc.ref.Append("timeout_sec"))
}

// Type returns a reference to field type of google_compute_health_check.
func (chc computeHealthCheckAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(chc.ref.Append("type"))
}

// UnhealthyThreshold returns a reference to field unhealthy_threshold of google_compute_health_check.
func (chc computeHealthCheckAttributes) UnhealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(chc.ref.Append("unhealthy_threshold"))
}

func (chc computeHealthCheckAttributes) GrpcHealthCheck() terra.ListValue[computehealthcheck.GrpcHealthCheckAttributes] {
	return terra.ReferenceAsList[computehealthcheck.GrpcHealthCheckAttributes](chc.ref.Append("grpc_health_check"))
}

func (chc computeHealthCheckAttributes) Http2HealthCheck() terra.ListValue[computehealthcheck.Http2HealthCheckAttributes] {
	return terra.ReferenceAsList[computehealthcheck.Http2HealthCheckAttributes](chc.ref.Append("http2_health_check"))
}

func (chc computeHealthCheckAttributes) HttpHealthCheck() terra.ListValue[computehealthcheck.HttpHealthCheckAttributes] {
	return terra.ReferenceAsList[computehealthcheck.HttpHealthCheckAttributes](chc.ref.Append("http_health_check"))
}

func (chc computeHealthCheckAttributes) HttpsHealthCheck() terra.ListValue[computehealthcheck.HttpsHealthCheckAttributes] {
	return terra.ReferenceAsList[computehealthcheck.HttpsHealthCheckAttributes](chc.ref.Append("https_health_check"))
}

func (chc computeHealthCheckAttributes) LogConfig() terra.ListValue[computehealthcheck.LogConfigAttributes] {
	return terra.ReferenceAsList[computehealthcheck.LogConfigAttributes](chc.ref.Append("log_config"))
}

func (chc computeHealthCheckAttributes) SslHealthCheck() terra.ListValue[computehealthcheck.SslHealthCheckAttributes] {
	return terra.ReferenceAsList[computehealthcheck.SslHealthCheckAttributes](chc.ref.Append("ssl_health_check"))
}

func (chc computeHealthCheckAttributes) TcpHealthCheck() terra.ListValue[computehealthcheck.TcpHealthCheckAttributes] {
	return terra.ReferenceAsList[computehealthcheck.TcpHealthCheckAttributes](chc.ref.Append("tcp_health_check"))
}

func (chc computeHealthCheckAttributes) Timeouts() computehealthcheck.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computehealthcheck.TimeoutsAttributes](chc.ref.Append("timeouts"))
}

type computeHealthCheckState struct {
	CheckIntervalSec   float64                                    `json:"check_interval_sec"`
	CreationTimestamp  string                                     `json:"creation_timestamp"`
	Description        string                                     `json:"description"`
	HealthyThreshold   float64                                    `json:"healthy_threshold"`
	Id                 string                                     `json:"id"`
	Name               string                                     `json:"name"`
	Project            string                                     `json:"project"`
	SelfLink           string                                     `json:"self_link"`
	TimeoutSec         float64                                    `json:"timeout_sec"`
	Type               string                                     `json:"type"`
	UnhealthyThreshold float64                                    `json:"unhealthy_threshold"`
	GrpcHealthCheck    []computehealthcheck.GrpcHealthCheckState  `json:"grpc_health_check"`
	Http2HealthCheck   []computehealthcheck.Http2HealthCheckState `json:"http2_health_check"`
	HttpHealthCheck    []computehealthcheck.HttpHealthCheckState  `json:"http_health_check"`
	HttpsHealthCheck   []computehealthcheck.HttpsHealthCheckState `json:"https_health_check"`
	LogConfig          []computehealthcheck.LogConfigState        `json:"log_config"`
	SslHealthCheck     []computehealthcheck.SslHealthCheckState   `json:"ssl_health_check"`
	TcpHealthCheck     []computehealthcheck.TcpHealthCheckState   `json:"tcp_health_check"`
	Timeouts           *computehealthcheck.TimeoutsState          `json:"timeouts"`
}
