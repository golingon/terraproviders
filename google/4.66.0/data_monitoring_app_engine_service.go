// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datamonitoringappengineservice "github.com/golingon/terraproviders/google/4.66.0/datamonitoringappengineservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMonitoringAppEngineService creates a new instance of [DataMonitoringAppEngineService].
func NewDataMonitoringAppEngineService(name string, args DataMonitoringAppEngineServiceArgs) *DataMonitoringAppEngineService {
	return &DataMonitoringAppEngineService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMonitoringAppEngineService)(nil)

// DataMonitoringAppEngineService represents the Terraform data resource google_monitoring_app_engine_service.
type DataMonitoringAppEngineService struct {
	Name string
	Args DataMonitoringAppEngineServiceArgs
}

// DataSource returns the Terraform object type for [DataMonitoringAppEngineService].
func (maes *DataMonitoringAppEngineService) DataSource() string {
	return "google_monitoring_app_engine_service"
}

// LocalName returns the local name for [DataMonitoringAppEngineService].
func (maes *DataMonitoringAppEngineService) LocalName() string {
	return maes.Name
}

// Configuration returns the configuration (args) for [DataMonitoringAppEngineService].
func (maes *DataMonitoringAppEngineService) Configuration() interface{} {
	return maes.Args
}

// Attributes returns the attributes for [DataMonitoringAppEngineService].
func (maes *DataMonitoringAppEngineService) Attributes() dataMonitoringAppEngineServiceAttributes {
	return dataMonitoringAppEngineServiceAttributes{ref: terra.ReferenceDataResource(maes)}
}

// DataMonitoringAppEngineServiceArgs contains the configurations for google_monitoring_app_engine_service.
type DataMonitoringAppEngineServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ModuleId: string, required
	ModuleId terra.StringValue `hcl:"module_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Telemetry: min=0
	Telemetry []datamonitoringappengineservice.Telemetry `hcl:"telemetry,block" validate:"min=0"`
}
type dataMonitoringAppEngineServiceAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_monitoring_app_engine_service.
func (maes dataMonitoringAppEngineServiceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(maes.ref.Append("display_name"))
}

// Id returns a reference to field id of google_monitoring_app_engine_service.
func (maes dataMonitoringAppEngineServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(maes.ref.Append("id"))
}

// ModuleId returns a reference to field module_id of google_monitoring_app_engine_service.
func (maes dataMonitoringAppEngineServiceAttributes) ModuleId() terra.StringValue {
	return terra.ReferenceAsString(maes.ref.Append("module_id"))
}

// Name returns a reference to field name of google_monitoring_app_engine_service.
func (maes dataMonitoringAppEngineServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(maes.ref.Append("name"))
}

// Project returns a reference to field project of google_monitoring_app_engine_service.
func (maes dataMonitoringAppEngineServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(maes.ref.Append("project"))
}

// ServiceId returns a reference to field service_id of google_monitoring_app_engine_service.
func (maes dataMonitoringAppEngineServiceAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(maes.ref.Append("service_id"))
}

// UserLabels returns a reference to field user_labels of google_monitoring_app_engine_service.
func (maes dataMonitoringAppEngineServiceAttributes) UserLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](maes.ref.Append("user_labels"))
}

func (maes dataMonitoringAppEngineServiceAttributes) Telemetry() terra.ListValue[datamonitoringappengineservice.TelemetryAttributes] {
	return terra.ReferenceAsList[datamonitoringappengineservice.TelemetryAttributes](maes.ref.Append("telemetry"))
}
