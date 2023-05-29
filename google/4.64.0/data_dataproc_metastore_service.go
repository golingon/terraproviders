// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datadataprocmetastoreservice "github.com/golingon/terraproviders/google/4.64.0/datadataprocmetastoreservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDataprocMetastoreService creates a new instance of [DataDataprocMetastoreService].
func NewDataDataprocMetastoreService(name string, args DataDataprocMetastoreServiceArgs) *DataDataprocMetastoreService {
	return &DataDataprocMetastoreService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDataprocMetastoreService)(nil)

// DataDataprocMetastoreService represents the Terraform data resource google_dataproc_metastore_service.
type DataDataprocMetastoreService struct {
	Name string
	Args DataDataprocMetastoreServiceArgs
}

// DataSource returns the Terraform object type for [DataDataprocMetastoreService].
func (dms *DataDataprocMetastoreService) DataSource() string {
	return "google_dataproc_metastore_service"
}

// LocalName returns the local name for [DataDataprocMetastoreService].
func (dms *DataDataprocMetastoreService) LocalName() string {
	return dms.Name
}

// Configuration returns the configuration (args) for [DataDataprocMetastoreService].
func (dms *DataDataprocMetastoreService) Configuration() interface{} {
	return dms.Args
}

// Attributes returns the attributes for [DataDataprocMetastoreService].
func (dms *DataDataprocMetastoreService) Attributes() dataDataprocMetastoreServiceAttributes {
	return dataDataprocMetastoreServiceAttributes{ref: terra.ReferenceDataResource(dms)}
}

// DataDataprocMetastoreServiceArgs contains the configurations for google_dataproc_metastore_service.
type DataDataprocMetastoreServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ServiceId: string, required
	ServiceId terra.StringValue `hcl:"service_id,attr" validate:"required"`
	// EncryptionConfig: min=0
	EncryptionConfig []datadataprocmetastoreservice.EncryptionConfig `hcl:"encryption_config,block" validate:"min=0"`
	// HiveMetastoreConfig: min=0
	HiveMetastoreConfig []datadataprocmetastoreservice.HiveMetastoreConfig `hcl:"hive_metastore_config,block" validate:"min=0"`
	// MaintenanceWindow: min=0
	MaintenanceWindow []datadataprocmetastoreservice.MaintenanceWindow `hcl:"maintenance_window,block" validate:"min=0"`
	// NetworkConfig: min=0
	NetworkConfig []datadataprocmetastoreservice.NetworkConfig `hcl:"network_config,block" validate:"min=0"`
	// TelemetryConfig: min=0
	TelemetryConfig []datadataprocmetastoreservice.TelemetryConfig `hcl:"telemetry_config,block" validate:"min=0"`
}
type dataDataprocMetastoreServiceAttributes struct {
	ref terra.Reference
}

// ArtifactGcsUri returns a reference to field artifact_gcs_uri of google_dataproc_metastore_service.
func (dms dataDataprocMetastoreServiceAttributes) ArtifactGcsUri() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("artifact_gcs_uri"))
}

// DatabaseType returns a reference to field database_type of google_dataproc_metastore_service.
func (dms dataDataprocMetastoreServiceAttributes) DatabaseType() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("database_type"))
}

// EndpointUri returns a reference to field endpoint_uri of google_dataproc_metastore_service.
func (dms dataDataprocMetastoreServiceAttributes) EndpointUri() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("endpoint_uri"))
}

// Id returns a reference to field id of google_dataproc_metastore_service.
func (dms dataDataprocMetastoreServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("id"))
}

// Labels returns a reference to field labels of google_dataproc_metastore_service.
func (dms dataDataprocMetastoreServiceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dms.ref.Append("labels"))
}

// Location returns a reference to field location of google_dataproc_metastore_service.
func (dms dataDataprocMetastoreServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("location"))
}

// Name returns a reference to field name of google_dataproc_metastore_service.
func (dms dataDataprocMetastoreServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("name"))
}

// Network returns a reference to field network of google_dataproc_metastore_service.
func (dms dataDataprocMetastoreServiceAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("network"))
}

// Port returns a reference to field port of google_dataproc_metastore_service.
func (dms dataDataprocMetastoreServiceAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(dms.ref.Append("port"))
}

// Project returns a reference to field project of google_dataproc_metastore_service.
func (dms dataDataprocMetastoreServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("project"))
}

// ReleaseChannel returns a reference to field release_channel of google_dataproc_metastore_service.
func (dms dataDataprocMetastoreServiceAttributes) ReleaseChannel() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("release_channel"))
}

// ServiceId returns a reference to field service_id of google_dataproc_metastore_service.
func (dms dataDataprocMetastoreServiceAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("service_id"))
}

// State returns a reference to field state of google_dataproc_metastore_service.
func (dms dataDataprocMetastoreServiceAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("state"))
}

// StateMessage returns a reference to field state_message of google_dataproc_metastore_service.
func (dms dataDataprocMetastoreServiceAttributes) StateMessage() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("state_message"))
}

// Tier returns a reference to field tier of google_dataproc_metastore_service.
func (dms dataDataprocMetastoreServiceAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("tier"))
}

// Uid returns a reference to field uid of google_dataproc_metastore_service.
func (dms dataDataprocMetastoreServiceAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("uid"))
}

func (dms dataDataprocMetastoreServiceAttributes) EncryptionConfig() terra.ListValue[datadataprocmetastoreservice.EncryptionConfigAttributes] {
	return terra.ReferenceAsList[datadataprocmetastoreservice.EncryptionConfigAttributes](dms.ref.Append("encryption_config"))
}

func (dms dataDataprocMetastoreServiceAttributes) HiveMetastoreConfig() terra.ListValue[datadataprocmetastoreservice.HiveMetastoreConfigAttributes] {
	return terra.ReferenceAsList[datadataprocmetastoreservice.HiveMetastoreConfigAttributes](dms.ref.Append("hive_metastore_config"))
}

func (dms dataDataprocMetastoreServiceAttributes) MaintenanceWindow() terra.ListValue[datadataprocmetastoreservice.MaintenanceWindowAttributes] {
	return terra.ReferenceAsList[datadataprocmetastoreservice.MaintenanceWindowAttributes](dms.ref.Append("maintenance_window"))
}

func (dms dataDataprocMetastoreServiceAttributes) NetworkConfig() terra.ListValue[datadataprocmetastoreservice.NetworkConfigAttributes] {
	return terra.ReferenceAsList[datadataprocmetastoreservice.NetworkConfigAttributes](dms.ref.Append("network_config"))
}

func (dms dataDataprocMetastoreServiceAttributes) TelemetryConfig() terra.ListValue[datadataprocmetastoreservice.TelemetryConfigAttributes] {
	return terra.ReferenceAsList[datadataprocmetastoreservice.TelemetryConfigAttributes](dms.ref.Append("telemetry_config"))
}