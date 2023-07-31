// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datamonitoringclusteristioservice "github.com/golingon/terraproviders/google/4.75.1/datamonitoringclusteristioservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMonitoringClusterIstioService creates a new instance of [DataMonitoringClusterIstioService].
func NewDataMonitoringClusterIstioService(name string, args DataMonitoringClusterIstioServiceArgs) *DataMonitoringClusterIstioService {
	return &DataMonitoringClusterIstioService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMonitoringClusterIstioService)(nil)

// DataMonitoringClusterIstioService represents the Terraform data resource google_monitoring_cluster_istio_service.
type DataMonitoringClusterIstioService struct {
	Name string
	Args DataMonitoringClusterIstioServiceArgs
}

// DataSource returns the Terraform object type for [DataMonitoringClusterIstioService].
func (mcis *DataMonitoringClusterIstioService) DataSource() string {
	return "google_monitoring_cluster_istio_service"
}

// LocalName returns the local name for [DataMonitoringClusterIstioService].
func (mcis *DataMonitoringClusterIstioService) LocalName() string {
	return mcis.Name
}

// Configuration returns the configuration (args) for [DataMonitoringClusterIstioService].
func (mcis *DataMonitoringClusterIstioService) Configuration() interface{} {
	return mcis.Args
}

// Attributes returns the attributes for [DataMonitoringClusterIstioService].
func (mcis *DataMonitoringClusterIstioService) Attributes() dataMonitoringClusterIstioServiceAttributes {
	return dataMonitoringClusterIstioServiceAttributes{ref: terra.ReferenceDataResource(mcis)}
}

// DataMonitoringClusterIstioServiceArgs contains the configurations for google_monitoring_cluster_istio_service.
type DataMonitoringClusterIstioServiceArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// ServiceNamespace: string, required
	ServiceNamespace terra.StringValue `hcl:"service_namespace,attr" validate:"required"`
	// Telemetry: min=0
	Telemetry []datamonitoringclusteristioservice.Telemetry `hcl:"telemetry,block" validate:"min=0"`
}
type dataMonitoringClusterIstioServiceAttributes struct {
	ref terra.Reference
}

// ClusterName returns a reference to field cluster_name of google_monitoring_cluster_istio_service.
func (mcis dataMonitoringClusterIstioServiceAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(mcis.ref.Append("cluster_name"))
}

// DisplayName returns a reference to field display_name of google_monitoring_cluster_istio_service.
func (mcis dataMonitoringClusterIstioServiceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(mcis.ref.Append("display_name"))
}

// Id returns a reference to field id of google_monitoring_cluster_istio_service.
func (mcis dataMonitoringClusterIstioServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mcis.ref.Append("id"))
}

// Location returns a reference to field location of google_monitoring_cluster_istio_service.
func (mcis dataMonitoringClusterIstioServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mcis.ref.Append("location"))
}

// Name returns a reference to field name of google_monitoring_cluster_istio_service.
func (mcis dataMonitoringClusterIstioServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mcis.ref.Append("name"))
}

// Project returns a reference to field project of google_monitoring_cluster_istio_service.
func (mcis dataMonitoringClusterIstioServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(mcis.ref.Append("project"))
}

// ServiceId returns a reference to field service_id of google_monitoring_cluster_istio_service.
func (mcis dataMonitoringClusterIstioServiceAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(mcis.ref.Append("service_id"))
}

// ServiceName returns a reference to field service_name of google_monitoring_cluster_istio_service.
func (mcis dataMonitoringClusterIstioServiceAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(mcis.ref.Append("service_name"))
}

// ServiceNamespace returns a reference to field service_namespace of google_monitoring_cluster_istio_service.
func (mcis dataMonitoringClusterIstioServiceAttributes) ServiceNamespace() terra.StringValue {
	return terra.ReferenceAsString(mcis.ref.Append("service_namespace"))
}

// UserLabels returns a reference to field user_labels of google_monitoring_cluster_istio_service.
func (mcis dataMonitoringClusterIstioServiceAttributes) UserLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mcis.ref.Append("user_labels"))
}

func (mcis dataMonitoringClusterIstioServiceAttributes) Telemetry() terra.ListValue[datamonitoringclusteristioservice.TelemetryAttributes] {
	return terra.ReferenceAsList[datamonitoringclusteristioservice.TelemetryAttributes](mcis.ref.Append("telemetry"))
}
