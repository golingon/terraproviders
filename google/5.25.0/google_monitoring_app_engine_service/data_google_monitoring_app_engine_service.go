// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_monitoring_app_engine_service

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_monitoring_app_engine_service.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gmaes *DataSource) DataSource() string {
	return "google_monitoring_app_engine_service"
}

// LocalName returns the local name for [DataSource].
func (gmaes *DataSource) LocalName() string {
	return gmaes.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gmaes *DataSource) Configuration() interface{} {
	return gmaes.Args
}

// Attributes returns the attributes for [DataSource].
func (gmaes *DataSource) Attributes() dataGoogleMonitoringAppEngineServiceAttributes {
	return dataGoogleMonitoringAppEngineServiceAttributes{ref: terra.ReferenceDataSource(gmaes)}
}

// DataArgs contains the configurations for google_monitoring_app_engine_service.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ModuleId: string, required
	ModuleId terra.StringValue `hcl:"module_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}

type dataGoogleMonitoringAppEngineServiceAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_monitoring_app_engine_service.
func (gmaes dataGoogleMonitoringAppEngineServiceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gmaes.ref.Append("display_name"))
}

// Id returns a reference to field id of google_monitoring_app_engine_service.
func (gmaes dataGoogleMonitoringAppEngineServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gmaes.ref.Append("id"))
}

// ModuleId returns a reference to field module_id of google_monitoring_app_engine_service.
func (gmaes dataGoogleMonitoringAppEngineServiceAttributes) ModuleId() terra.StringValue {
	return terra.ReferenceAsString(gmaes.ref.Append("module_id"))
}

// Name returns a reference to field name of google_monitoring_app_engine_service.
func (gmaes dataGoogleMonitoringAppEngineServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gmaes.ref.Append("name"))
}

// Project returns a reference to field project of google_monitoring_app_engine_service.
func (gmaes dataGoogleMonitoringAppEngineServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gmaes.ref.Append("project"))
}

// ServiceId returns a reference to field service_id of google_monitoring_app_engine_service.
func (gmaes dataGoogleMonitoringAppEngineServiceAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(gmaes.ref.Append("service_id"))
}

// UserLabels returns a reference to field user_labels of google_monitoring_app_engine_service.
func (gmaes dataGoogleMonitoringAppEngineServiceAttributes) UserLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gmaes.ref.Append("user_labels"))
}

func (gmaes dataGoogleMonitoringAppEngineServiceAttributes) Telemetry() terra.ListValue[DataTelemetryAttributes] {
	return terra.ReferenceAsList[DataTelemetryAttributes](gmaes.ref.Append("telemetry"))
}
