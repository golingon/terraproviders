// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datamonitoringmeshistioservice "github.com/golingon/terraproviders/googlebeta/4.66.0/datamonitoringmeshistioservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMonitoringMeshIstioService creates a new instance of [DataMonitoringMeshIstioService].
func NewDataMonitoringMeshIstioService(name string, args DataMonitoringMeshIstioServiceArgs) *DataMonitoringMeshIstioService {
	return &DataMonitoringMeshIstioService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMonitoringMeshIstioService)(nil)

// DataMonitoringMeshIstioService represents the Terraform data resource google_monitoring_mesh_istio_service.
type DataMonitoringMeshIstioService struct {
	Name string
	Args DataMonitoringMeshIstioServiceArgs
}

// DataSource returns the Terraform object type for [DataMonitoringMeshIstioService].
func (mmis *DataMonitoringMeshIstioService) DataSource() string {
	return "google_monitoring_mesh_istio_service"
}

// LocalName returns the local name for [DataMonitoringMeshIstioService].
func (mmis *DataMonitoringMeshIstioService) LocalName() string {
	return mmis.Name
}

// Configuration returns the configuration (args) for [DataMonitoringMeshIstioService].
func (mmis *DataMonitoringMeshIstioService) Configuration() interface{} {
	return mmis.Args
}

// Attributes returns the attributes for [DataMonitoringMeshIstioService].
func (mmis *DataMonitoringMeshIstioService) Attributes() dataMonitoringMeshIstioServiceAttributes {
	return dataMonitoringMeshIstioServiceAttributes{ref: terra.ReferenceDataResource(mmis)}
}

// DataMonitoringMeshIstioServiceArgs contains the configurations for google_monitoring_mesh_istio_service.
type DataMonitoringMeshIstioServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MeshUid: string, required
	MeshUid terra.StringValue `hcl:"mesh_uid,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// ServiceNamespace: string, required
	ServiceNamespace terra.StringValue `hcl:"service_namespace,attr" validate:"required"`
	// Telemetry: min=0
	Telemetry []datamonitoringmeshistioservice.Telemetry `hcl:"telemetry,block" validate:"min=0"`
}
type dataMonitoringMeshIstioServiceAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_monitoring_mesh_istio_service.
func (mmis dataMonitoringMeshIstioServiceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(mmis.ref.Append("display_name"))
}

// Id returns a reference to field id of google_monitoring_mesh_istio_service.
func (mmis dataMonitoringMeshIstioServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mmis.ref.Append("id"))
}

// MeshUid returns a reference to field mesh_uid of google_monitoring_mesh_istio_service.
func (mmis dataMonitoringMeshIstioServiceAttributes) MeshUid() terra.StringValue {
	return terra.ReferenceAsString(mmis.ref.Append("mesh_uid"))
}

// Name returns a reference to field name of google_monitoring_mesh_istio_service.
func (mmis dataMonitoringMeshIstioServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mmis.ref.Append("name"))
}

// Project returns a reference to field project of google_monitoring_mesh_istio_service.
func (mmis dataMonitoringMeshIstioServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(mmis.ref.Append("project"))
}

// ServiceId returns a reference to field service_id of google_monitoring_mesh_istio_service.
func (mmis dataMonitoringMeshIstioServiceAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(mmis.ref.Append("service_id"))
}

// ServiceName returns a reference to field service_name of google_monitoring_mesh_istio_service.
func (mmis dataMonitoringMeshIstioServiceAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(mmis.ref.Append("service_name"))
}

// ServiceNamespace returns a reference to field service_namespace of google_monitoring_mesh_istio_service.
func (mmis dataMonitoringMeshIstioServiceAttributes) ServiceNamespace() terra.StringValue {
	return terra.ReferenceAsString(mmis.ref.Append("service_namespace"))
}

// UserLabels returns a reference to field user_labels of google_monitoring_mesh_istio_service.
func (mmis dataMonitoringMeshIstioServiceAttributes) UserLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mmis.ref.Append("user_labels"))
}

func (mmis dataMonitoringMeshIstioServiceAttributes) Telemetry() terra.ListValue[datamonitoringmeshistioservice.TelemetryAttributes] {
	return terra.ReferenceAsList[datamonitoringmeshistioservice.TelemetryAttributes](mmis.ref.Append("telemetry"))
}
