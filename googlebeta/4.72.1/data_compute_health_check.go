// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datacomputehealthcheck "github.com/golingon/terraproviders/googlebeta/4.72.1/datacomputehealthcheck"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeHealthCheck creates a new instance of [DataComputeHealthCheck].
func NewDataComputeHealthCheck(name string, args DataComputeHealthCheckArgs) *DataComputeHealthCheck {
	return &DataComputeHealthCheck{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeHealthCheck)(nil)

// DataComputeHealthCheck represents the Terraform data resource google_compute_health_check.
type DataComputeHealthCheck struct {
	Name string
	Args DataComputeHealthCheckArgs
}

// DataSource returns the Terraform object type for [DataComputeHealthCheck].
func (chc *DataComputeHealthCheck) DataSource() string {
	return "google_compute_health_check"
}

// LocalName returns the local name for [DataComputeHealthCheck].
func (chc *DataComputeHealthCheck) LocalName() string {
	return chc.Name
}

// Configuration returns the configuration (args) for [DataComputeHealthCheck].
func (chc *DataComputeHealthCheck) Configuration() interface{} {
	return chc.Args
}

// Attributes returns the attributes for [DataComputeHealthCheck].
func (chc *DataComputeHealthCheck) Attributes() dataComputeHealthCheckAttributes {
	return dataComputeHealthCheckAttributes{ref: terra.ReferenceDataResource(chc)}
}

// DataComputeHealthCheckArgs contains the configurations for google_compute_health_check.
type DataComputeHealthCheckArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// GrpcHealthCheck: min=0
	GrpcHealthCheck []datacomputehealthcheck.GrpcHealthCheck `hcl:"grpc_health_check,block" validate:"min=0"`
	// Http2HealthCheck: min=0
	Http2HealthCheck []datacomputehealthcheck.Http2HealthCheck `hcl:"http2_health_check,block" validate:"min=0"`
	// HttpHealthCheck: min=0
	HttpHealthCheck []datacomputehealthcheck.HttpHealthCheck `hcl:"http_health_check,block" validate:"min=0"`
	// HttpsHealthCheck: min=0
	HttpsHealthCheck []datacomputehealthcheck.HttpsHealthCheck `hcl:"https_health_check,block" validate:"min=0"`
	// LogConfig: min=0
	LogConfig []datacomputehealthcheck.LogConfig `hcl:"log_config,block" validate:"min=0"`
	// SslHealthCheck: min=0
	SslHealthCheck []datacomputehealthcheck.SslHealthCheck `hcl:"ssl_health_check,block" validate:"min=0"`
	// TcpHealthCheck: min=0
	TcpHealthCheck []datacomputehealthcheck.TcpHealthCheck `hcl:"tcp_health_check,block" validate:"min=0"`
}
type dataComputeHealthCheckAttributes struct {
	ref terra.Reference
}

// CheckIntervalSec returns a reference to field check_interval_sec of google_compute_health_check.
func (chc dataComputeHealthCheckAttributes) CheckIntervalSec() terra.NumberValue {
	return terra.ReferenceAsNumber(chc.ref.Append("check_interval_sec"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_health_check.
func (chc dataComputeHealthCheckAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(chc.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_health_check.
func (chc dataComputeHealthCheckAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(chc.ref.Append("description"))
}

// HealthyThreshold returns a reference to field healthy_threshold of google_compute_health_check.
func (chc dataComputeHealthCheckAttributes) HealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(chc.ref.Append("healthy_threshold"))
}

// Id returns a reference to field id of google_compute_health_check.
func (chc dataComputeHealthCheckAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(chc.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_health_check.
func (chc dataComputeHealthCheckAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(chc.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_health_check.
func (chc dataComputeHealthCheckAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(chc.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_health_check.
func (chc dataComputeHealthCheckAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(chc.ref.Append("self_link"))
}

// TimeoutSec returns a reference to field timeout_sec of google_compute_health_check.
func (chc dataComputeHealthCheckAttributes) TimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(chc.ref.Append("timeout_sec"))
}

// Type returns a reference to field type of google_compute_health_check.
func (chc dataComputeHealthCheckAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(chc.ref.Append("type"))
}

// UnhealthyThreshold returns a reference to field unhealthy_threshold of google_compute_health_check.
func (chc dataComputeHealthCheckAttributes) UnhealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(chc.ref.Append("unhealthy_threshold"))
}

func (chc dataComputeHealthCheckAttributes) GrpcHealthCheck() terra.ListValue[datacomputehealthcheck.GrpcHealthCheckAttributes] {
	return terra.ReferenceAsList[datacomputehealthcheck.GrpcHealthCheckAttributes](chc.ref.Append("grpc_health_check"))
}

func (chc dataComputeHealthCheckAttributes) Http2HealthCheck() terra.ListValue[datacomputehealthcheck.Http2HealthCheckAttributes] {
	return terra.ReferenceAsList[datacomputehealthcheck.Http2HealthCheckAttributes](chc.ref.Append("http2_health_check"))
}

func (chc dataComputeHealthCheckAttributes) HttpHealthCheck() terra.ListValue[datacomputehealthcheck.HttpHealthCheckAttributes] {
	return terra.ReferenceAsList[datacomputehealthcheck.HttpHealthCheckAttributes](chc.ref.Append("http_health_check"))
}

func (chc dataComputeHealthCheckAttributes) HttpsHealthCheck() terra.ListValue[datacomputehealthcheck.HttpsHealthCheckAttributes] {
	return terra.ReferenceAsList[datacomputehealthcheck.HttpsHealthCheckAttributes](chc.ref.Append("https_health_check"))
}

func (chc dataComputeHealthCheckAttributes) LogConfig() terra.ListValue[datacomputehealthcheck.LogConfigAttributes] {
	return terra.ReferenceAsList[datacomputehealthcheck.LogConfigAttributes](chc.ref.Append("log_config"))
}

func (chc dataComputeHealthCheckAttributes) SslHealthCheck() terra.ListValue[datacomputehealthcheck.SslHealthCheckAttributes] {
	return terra.ReferenceAsList[datacomputehealthcheck.SslHealthCheckAttributes](chc.ref.Append("ssl_health_check"))
}

func (chc dataComputeHealthCheckAttributes) TcpHealthCheck() terra.ListValue[datacomputehealthcheck.TcpHealthCheckAttributes] {
	return terra.ReferenceAsList[datacomputehealthcheck.TcpHealthCheckAttributes](chc.ref.Append("tcp_health_check"))
}
