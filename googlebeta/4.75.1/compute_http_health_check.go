// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computehttphealthcheck "github.com/golingon/terraproviders/googlebeta/4.75.1/computehttphealthcheck"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeHttpHealthCheck creates a new instance of [ComputeHttpHealthCheck].
func NewComputeHttpHealthCheck(name string, args ComputeHttpHealthCheckArgs) *ComputeHttpHealthCheck {
	return &ComputeHttpHealthCheck{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeHttpHealthCheck)(nil)

// ComputeHttpHealthCheck represents the Terraform resource google_compute_http_health_check.
type ComputeHttpHealthCheck struct {
	Name      string
	Args      ComputeHttpHealthCheckArgs
	state     *computeHttpHealthCheckState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeHttpHealthCheck].
func (chhc *ComputeHttpHealthCheck) Type() string {
	return "google_compute_http_health_check"
}

// LocalName returns the local name for [ComputeHttpHealthCheck].
func (chhc *ComputeHttpHealthCheck) LocalName() string {
	return chhc.Name
}

// Configuration returns the configuration (args) for [ComputeHttpHealthCheck].
func (chhc *ComputeHttpHealthCheck) Configuration() interface{} {
	return chhc.Args
}

// DependOn is used for other resources to depend on [ComputeHttpHealthCheck].
func (chhc *ComputeHttpHealthCheck) DependOn() terra.Reference {
	return terra.ReferenceResource(chhc)
}

// Dependencies returns the list of resources [ComputeHttpHealthCheck] depends_on.
func (chhc *ComputeHttpHealthCheck) Dependencies() terra.Dependencies {
	return chhc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeHttpHealthCheck].
func (chhc *ComputeHttpHealthCheck) LifecycleManagement() *terra.Lifecycle {
	return chhc.Lifecycle
}

// Attributes returns the attributes for [ComputeHttpHealthCheck].
func (chhc *ComputeHttpHealthCheck) Attributes() computeHttpHealthCheckAttributes {
	return computeHttpHealthCheckAttributes{ref: terra.ReferenceResource(chhc)}
}

// ImportState imports the given attribute values into [ComputeHttpHealthCheck]'s state.
func (chhc *ComputeHttpHealthCheck) ImportState(av io.Reader) error {
	chhc.state = &computeHttpHealthCheckState{}
	if err := json.NewDecoder(av).Decode(chhc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", chhc.Type(), chhc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeHttpHealthCheck] has state.
func (chhc *ComputeHttpHealthCheck) State() (*computeHttpHealthCheckState, bool) {
	return chhc.state, chhc.state != nil
}

// StateMust returns the state for [ComputeHttpHealthCheck]. Panics if the state is nil.
func (chhc *ComputeHttpHealthCheck) StateMust() *computeHttpHealthCheckState {
	if chhc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", chhc.Type(), chhc.LocalName()))
	}
	return chhc.state
}

// ComputeHttpHealthCheckArgs contains the configurations for google_compute_http_health_check.
type ComputeHttpHealthCheckArgs struct {
	// CheckIntervalSec: number, optional
	CheckIntervalSec terra.NumberValue `hcl:"check_interval_sec,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// HealthyThreshold: number, optional
	HealthyThreshold terra.NumberValue `hcl:"healthy_threshold,attr"`
	// Host: string, optional
	Host terra.StringValue `hcl:"host,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RequestPath: string, optional
	RequestPath terra.StringValue `hcl:"request_path,attr"`
	// TimeoutSec: number, optional
	TimeoutSec terra.NumberValue `hcl:"timeout_sec,attr"`
	// UnhealthyThreshold: number, optional
	UnhealthyThreshold terra.NumberValue `hcl:"unhealthy_threshold,attr"`
	// Timeouts: optional
	Timeouts *computehttphealthcheck.Timeouts `hcl:"timeouts,block"`
}
type computeHttpHealthCheckAttributes struct {
	ref terra.Reference
}

// CheckIntervalSec returns a reference to field check_interval_sec of google_compute_http_health_check.
func (chhc computeHttpHealthCheckAttributes) CheckIntervalSec() terra.NumberValue {
	return terra.ReferenceAsNumber(chhc.ref.Append("check_interval_sec"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_http_health_check.
func (chhc computeHttpHealthCheckAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(chhc.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_http_health_check.
func (chhc computeHttpHealthCheckAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(chhc.ref.Append("description"))
}

// HealthyThreshold returns a reference to field healthy_threshold of google_compute_http_health_check.
func (chhc computeHttpHealthCheckAttributes) HealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(chhc.ref.Append("healthy_threshold"))
}

// Host returns a reference to field host of google_compute_http_health_check.
func (chhc computeHttpHealthCheckAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(chhc.ref.Append("host"))
}

// Id returns a reference to field id of google_compute_http_health_check.
func (chhc computeHttpHealthCheckAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(chhc.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_http_health_check.
func (chhc computeHttpHealthCheckAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(chhc.ref.Append("name"))
}

// Port returns a reference to field port of google_compute_http_health_check.
func (chhc computeHttpHealthCheckAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(chhc.ref.Append("port"))
}

// Project returns a reference to field project of google_compute_http_health_check.
func (chhc computeHttpHealthCheckAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(chhc.ref.Append("project"))
}

// RequestPath returns a reference to field request_path of google_compute_http_health_check.
func (chhc computeHttpHealthCheckAttributes) RequestPath() terra.StringValue {
	return terra.ReferenceAsString(chhc.ref.Append("request_path"))
}

// SelfLink returns a reference to field self_link of google_compute_http_health_check.
func (chhc computeHttpHealthCheckAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(chhc.ref.Append("self_link"))
}

// TimeoutSec returns a reference to field timeout_sec of google_compute_http_health_check.
func (chhc computeHttpHealthCheckAttributes) TimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(chhc.ref.Append("timeout_sec"))
}

// UnhealthyThreshold returns a reference to field unhealthy_threshold of google_compute_http_health_check.
func (chhc computeHttpHealthCheckAttributes) UnhealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(chhc.ref.Append("unhealthy_threshold"))
}

func (chhc computeHttpHealthCheckAttributes) Timeouts() computehttphealthcheck.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computehttphealthcheck.TimeoutsAttributes](chhc.ref.Append("timeouts"))
}

type computeHttpHealthCheckState struct {
	CheckIntervalSec   float64                               `json:"check_interval_sec"`
	CreationTimestamp  string                                `json:"creation_timestamp"`
	Description        string                                `json:"description"`
	HealthyThreshold   float64                               `json:"healthy_threshold"`
	Host               string                                `json:"host"`
	Id                 string                                `json:"id"`
	Name               string                                `json:"name"`
	Port               float64                               `json:"port"`
	Project            string                                `json:"project"`
	RequestPath        string                                `json:"request_path"`
	SelfLink           string                                `json:"self_link"`
	TimeoutSec         float64                               `json:"timeout_sec"`
	UnhealthyThreshold float64                               `json:"unhealthy_threshold"`
	Timeouts           *computehttphealthcheck.TimeoutsState `json:"timeouts"`
}
