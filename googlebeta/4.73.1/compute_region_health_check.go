// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeregionhealthcheck "github.com/golingon/terraproviders/googlebeta/4.73.1/computeregionhealthcheck"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionHealthCheck creates a new instance of [ComputeRegionHealthCheck].
func NewComputeRegionHealthCheck(name string, args ComputeRegionHealthCheckArgs) *ComputeRegionHealthCheck {
	return &ComputeRegionHealthCheck{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionHealthCheck)(nil)

// ComputeRegionHealthCheck represents the Terraform resource google_compute_region_health_check.
type ComputeRegionHealthCheck struct {
	Name      string
	Args      ComputeRegionHealthCheckArgs
	state     *computeRegionHealthCheckState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionHealthCheck].
func (crhc *ComputeRegionHealthCheck) Type() string {
	return "google_compute_region_health_check"
}

// LocalName returns the local name for [ComputeRegionHealthCheck].
func (crhc *ComputeRegionHealthCheck) LocalName() string {
	return crhc.Name
}

// Configuration returns the configuration (args) for [ComputeRegionHealthCheck].
func (crhc *ComputeRegionHealthCheck) Configuration() interface{} {
	return crhc.Args
}

// DependOn is used for other resources to depend on [ComputeRegionHealthCheck].
func (crhc *ComputeRegionHealthCheck) DependOn() terra.Reference {
	return terra.ReferenceResource(crhc)
}

// Dependencies returns the list of resources [ComputeRegionHealthCheck] depends_on.
func (crhc *ComputeRegionHealthCheck) Dependencies() terra.Dependencies {
	return crhc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionHealthCheck].
func (crhc *ComputeRegionHealthCheck) LifecycleManagement() *terra.Lifecycle {
	return crhc.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionHealthCheck].
func (crhc *ComputeRegionHealthCheck) Attributes() computeRegionHealthCheckAttributes {
	return computeRegionHealthCheckAttributes{ref: terra.ReferenceResource(crhc)}
}

// ImportState imports the given attribute values into [ComputeRegionHealthCheck]'s state.
func (crhc *ComputeRegionHealthCheck) ImportState(av io.Reader) error {
	crhc.state = &computeRegionHealthCheckState{}
	if err := json.NewDecoder(av).Decode(crhc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crhc.Type(), crhc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionHealthCheck] has state.
func (crhc *ComputeRegionHealthCheck) State() (*computeRegionHealthCheckState, bool) {
	return crhc.state, crhc.state != nil
}

// StateMust returns the state for [ComputeRegionHealthCheck]. Panics if the state is nil.
func (crhc *ComputeRegionHealthCheck) StateMust() *computeRegionHealthCheckState {
	if crhc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crhc.Type(), crhc.LocalName()))
	}
	return crhc.state
}

// ComputeRegionHealthCheckArgs contains the configurations for google_compute_region_health_check.
type ComputeRegionHealthCheckArgs struct {
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
	GrpcHealthCheck *computeregionhealthcheck.GrpcHealthCheck `hcl:"grpc_health_check,block"`
	// Http2HealthCheck: optional
	Http2HealthCheck *computeregionhealthcheck.Http2HealthCheck `hcl:"http2_health_check,block"`
	// HttpHealthCheck: optional
	HttpHealthCheck *computeregionhealthcheck.HttpHealthCheck `hcl:"http_health_check,block"`
	// HttpsHealthCheck: optional
	HttpsHealthCheck *computeregionhealthcheck.HttpsHealthCheck `hcl:"https_health_check,block"`
	// LogConfig: optional
	LogConfig *computeregionhealthcheck.LogConfig `hcl:"log_config,block"`
	// SslHealthCheck: optional
	SslHealthCheck *computeregionhealthcheck.SslHealthCheck `hcl:"ssl_health_check,block"`
	// TcpHealthCheck: optional
	TcpHealthCheck *computeregionhealthcheck.TcpHealthCheck `hcl:"tcp_health_check,block"`
	// Timeouts: optional
	Timeouts *computeregionhealthcheck.Timeouts `hcl:"timeouts,block"`
}
type computeRegionHealthCheckAttributes struct {
	ref terra.Reference
}

// CheckIntervalSec returns a reference to field check_interval_sec of google_compute_region_health_check.
func (crhc computeRegionHealthCheckAttributes) CheckIntervalSec() terra.NumberValue {
	return terra.ReferenceAsNumber(crhc.ref.Append("check_interval_sec"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_region_health_check.
func (crhc computeRegionHealthCheckAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(crhc.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_region_health_check.
func (crhc computeRegionHealthCheckAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crhc.ref.Append("description"))
}

// HealthyThreshold returns a reference to field healthy_threshold of google_compute_region_health_check.
func (crhc computeRegionHealthCheckAttributes) HealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(crhc.ref.Append("healthy_threshold"))
}

// Id returns a reference to field id of google_compute_region_health_check.
func (crhc computeRegionHealthCheckAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crhc.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_health_check.
func (crhc computeRegionHealthCheckAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crhc.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_health_check.
func (crhc computeRegionHealthCheckAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crhc.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_health_check.
func (crhc computeRegionHealthCheckAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crhc.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_health_check.
func (crhc computeRegionHealthCheckAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crhc.ref.Append("self_link"))
}

// TimeoutSec returns a reference to field timeout_sec of google_compute_region_health_check.
func (crhc computeRegionHealthCheckAttributes) TimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(crhc.ref.Append("timeout_sec"))
}

// Type returns a reference to field type of google_compute_region_health_check.
func (crhc computeRegionHealthCheckAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(crhc.ref.Append("type"))
}

// UnhealthyThreshold returns a reference to field unhealthy_threshold of google_compute_region_health_check.
func (crhc computeRegionHealthCheckAttributes) UnhealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(crhc.ref.Append("unhealthy_threshold"))
}

func (crhc computeRegionHealthCheckAttributes) GrpcHealthCheck() terra.ListValue[computeregionhealthcheck.GrpcHealthCheckAttributes] {
	return terra.ReferenceAsList[computeregionhealthcheck.GrpcHealthCheckAttributes](crhc.ref.Append("grpc_health_check"))
}

func (crhc computeRegionHealthCheckAttributes) Http2HealthCheck() terra.ListValue[computeregionhealthcheck.Http2HealthCheckAttributes] {
	return terra.ReferenceAsList[computeregionhealthcheck.Http2HealthCheckAttributes](crhc.ref.Append("http2_health_check"))
}

func (crhc computeRegionHealthCheckAttributes) HttpHealthCheck() terra.ListValue[computeregionhealthcheck.HttpHealthCheckAttributes] {
	return terra.ReferenceAsList[computeregionhealthcheck.HttpHealthCheckAttributes](crhc.ref.Append("http_health_check"))
}

func (crhc computeRegionHealthCheckAttributes) HttpsHealthCheck() terra.ListValue[computeregionhealthcheck.HttpsHealthCheckAttributes] {
	return terra.ReferenceAsList[computeregionhealthcheck.HttpsHealthCheckAttributes](crhc.ref.Append("https_health_check"))
}

func (crhc computeRegionHealthCheckAttributes) LogConfig() terra.ListValue[computeregionhealthcheck.LogConfigAttributes] {
	return terra.ReferenceAsList[computeregionhealthcheck.LogConfigAttributes](crhc.ref.Append("log_config"))
}

func (crhc computeRegionHealthCheckAttributes) SslHealthCheck() terra.ListValue[computeregionhealthcheck.SslHealthCheckAttributes] {
	return terra.ReferenceAsList[computeregionhealthcheck.SslHealthCheckAttributes](crhc.ref.Append("ssl_health_check"))
}

func (crhc computeRegionHealthCheckAttributes) TcpHealthCheck() terra.ListValue[computeregionhealthcheck.TcpHealthCheckAttributes] {
	return terra.ReferenceAsList[computeregionhealthcheck.TcpHealthCheckAttributes](crhc.ref.Append("tcp_health_check"))
}

func (crhc computeRegionHealthCheckAttributes) Timeouts() computeregionhealthcheck.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregionhealthcheck.TimeoutsAttributes](crhc.ref.Append("timeouts"))
}

type computeRegionHealthCheckState struct {
	CheckIntervalSec   float64                                          `json:"check_interval_sec"`
	CreationTimestamp  string                                           `json:"creation_timestamp"`
	Description        string                                           `json:"description"`
	HealthyThreshold   float64                                          `json:"healthy_threshold"`
	Id                 string                                           `json:"id"`
	Name               string                                           `json:"name"`
	Project            string                                           `json:"project"`
	Region             string                                           `json:"region"`
	SelfLink           string                                           `json:"self_link"`
	TimeoutSec         float64                                          `json:"timeout_sec"`
	Type               string                                           `json:"type"`
	UnhealthyThreshold float64                                          `json:"unhealthy_threshold"`
	GrpcHealthCheck    []computeregionhealthcheck.GrpcHealthCheckState  `json:"grpc_health_check"`
	Http2HealthCheck   []computeregionhealthcheck.Http2HealthCheckState `json:"http2_health_check"`
	HttpHealthCheck    []computeregionhealthcheck.HttpHealthCheckState  `json:"http_health_check"`
	HttpsHealthCheck   []computeregionhealthcheck.HttpsHealthCheckState `json:"https_health_check"`
	LogConfig          []computeregionhealthcheck.LogConfigState        `json:"log_config"`
	SslHealthCheck     []computeregionhealthcheck.SslHealthCheckState   `json:"ssl_health_check"`
	TcpHealthCheck     []computeregionhealthcheck.TcpHealthCheckState   `json:"tcp_health_check"`
	Timeouts           *computeregionhealthcheck.TimeoutsState          `json:"timeouts"`
}
