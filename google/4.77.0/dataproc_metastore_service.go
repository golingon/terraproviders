// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dataprocmetastoreservice "github.com/golingon/terraproviders/google/4.77.0/dataprocmetastoreservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataprocMetastoreService creates a new instance of [DataprocMetastoreService].
func NewDataprocMetastoreService(name string, args DataprocMetastoreServiceArgs) *DataprocMetastoreService {
	return &DataprocMetastoreService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocMetastoreService)(nil)

// DataprocMetastoreService represents the Terraform resource google_dataproc_metastore_service.
type DataprocMetastoreService struct {
	Name      string
	Args      DataprocMetastoreServiceArgs
	state     *dataprocMetastoreServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocMetastoreService].
func (dms *DataprocMetastoreService) Type() string {
	return "google_dataproc_metastore_service"
}

// LocalName returns the local name for [DataprocMetastoreService].
func (dms *DataprocMetastoreService) LocalName() string {
	return dms.Name
}

// Configuration returns the configuration (args) for [DataprocMetastoreService].
func (dms *DataprocMetastoreService) Configuration() interface{} {
	return dms.Args
}

// DependOn is used for other resources to depend on [DataprocMetastoreService].
func (dms *DataprocMetastoreService) DependOn() terra.Reference {
	return terra.ReferenceResource(dms)
}

// Dependencies returns the list of resources [DataprocMetastoreService] depends_on.
func (dms *DataprocMetastoreService) Dependencies() terra.Dependencies {
	return dms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocMetastoreService].
func (dms *DataprocMetastoreService) LifecycleManagement() *terra.Lifecycle {
	return dms.Lifecycle
}

// Attributes returns the attributes for [DataprocMetastoreService].
func (dms *DataprocMetastoreService) Attributes() dataprocMetastoreServiceAttributes {
	return dataprocMetastoreServiceAttributes{ref: terra.ReferenceResource(dms)}
}

// ImportState imports the given attribute values into [DataprocMetastoreService]'s state.
func (dms *DataprocMetastoreService) ImportState(av io.Reader) error {
	dms.state = &dataprocMetastoreServiceState{}
	if err := json.NewDecoder(av).Decode(dms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dms.Type(), dms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocMetastoreService] has state.
func (dms *DataprocMetastoreService) State() (*dataprocMetastoreServiceState, bool) {
	return dms.state, dms.state != nil
}

// StateMust returns the state for [DataprocMetastoreService]. Panics if the state is nil.
func (dms *DataprocMetastoreService) StateMust() *dataprocMetastoreServiceState {
	if dms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dms.Type(), dms.LocalName()))
	}
	return dms.state
}

// DataprocMetastoreServiceArgs contains the configurations for google_dataproc_metastore_service.
type DataprocMetastoreServiceArgs struct {
	// DatabaseType: string, optional
	DatabaseType terra.StringValue `hcl:"database_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ReleaseChannel: string, optional
	ReleaseChannel terra.StringValue `hcl:"release_channel,attr"`
	// ServiceId: string, required
	ServiceId terra.StringValue `hcl:"service_id,attr" validate:"required"`
	// Tier: string, optional
	Tier terra.StringValue `hcl:"tier,attr"`
	// EncryptionConfig: optional
	EncryptionConfig *dataprocmetastoreservice.EncryptionConfig `hcl:"encryption_config,block"`
	// HiveMetastoreConfig: optional
	HiveMetastoreConfig *dataprocmetastoreservice.HiveMetastoreConfig `hcl:"hive_metastore_config,block"`
	// MaintenanceWindow: optional
	MaintenanceWindow *dataprocmetastoreservice.MaintenanceWindow `hcl:"maintenance_window,block"`
	// NetworkConfig: optional
	NetworkConfig *dataprocmetastoreservice.NetworkConfig `hcl:"network_config,block"`
	// TelemetryConfig: optional
	TelemetryConfig *dataprocmetastoreservice.TelemetryConfig `hcl:"telemetry_config,block"`
	// Timeouts: optional
	Timeouts *dataprocmetastoreservice.Timeouts `hcl:"timeouts,block"`
}
type dataprocMetastoreServiceAttributes struct {
	ref terra.Reference
}

// ArtifactGcsUri returns a reference to field artifact_gcs_uri of google_dataproc_metastore_service.
func (dms dataprocMetastoreServiceAttributes) ArtifactGcsUri() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("artifact_gcs_uri"))
}

// DatabaseType returns a reference to field database_type of google_dataproc_metastore_service.
func (dms dataprocMetastoreServiceAttributes) DatabaseType() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("database_type"))
}

// EndpointUri returns a reference to field endpoint_uri of google_dataproc_metastore_service.
func (dms dataprocMetastoreServiceAttributes) EndpointUri() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("endpoint_uri"))
}

// Id returns a reference to field id of google_dataproc_metastore_service.
func (dms dataprocMetastoreServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("id"))
}

// Labels returns a reference to field labels of google_dataproc_metastore_service.
func (dms dataprocMetastoreServiceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dms.ref.Append("labels"))
}

// Location returns a reference to field location of google_dataproc_metastore_service.
func (dms dataprocMetastoreServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("location"))
}

// Name returns a reference to field name of google_dataproc_metastore_service.
func (dms dataprocMetastoreServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("name"))
}

// Network returns a reference to field network of google_dataproc_metastore_service.
func (dms dataprocMetastoreServiceAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("network"))
}

// Port returns a reference to field port of google_dataproc_metastore_service.
func (dms dataprocMetastoreServiceAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(dms.ref.Append("port"))
}

// Project returns a reference to field project of google_dataproc_metastore_service.
func (dms dataprocMetastoreServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("project"))
}

// ReleaseChannel returns a reference to field release_channel of google_dataproc_metastore_service.
func (dms dataprocMetastoreServiceAttributes) ReleaseChannel() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("release_channel"))
}

// ServiceId returns a reference to field service_id of google_dataproc_metastore_service.
func (dms dataprocMetastoreServiceAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("service_id"))
}

// State returns a reference to field state of google_dataproc_metastore_service.
func (dms dataprocMetastoreServiceAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("state"))
}

// StateMessage returns a reference to field state_message of google_dataproc_metastore_service.
func (dms dataprocMetastoreServiceAttributes) StateMessage() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("state_message"))
}

// Tier returns a reference to field tier of google_dataproc_metastore_service.
func (dms dataprocMetastoreServiceAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("tier"))
}

// Uid returns a reference to field uid of google_dataproc_metastore_service.
func (dms dataprocMetastoreServiceAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("uid"))
}

func (dms dataprocMetastoreServiceAttributes) EncryptionConfig() terra.ListValue[dataprocmetastoreservice.EncryptionConfigAttributes] {
	return terra.ReferenceAsList[dataprocmetastoreservice.EncryptionConfigAttributes](dms.ref.Append("encryption_config"))
}

func (dms dataprocMetastoreServiceAttributes) HiveMetastoreConfig() terra.ListValue[dataprocmetastoreservice.HiveMetastoreConfigAttributes] {
	return terra.ReferenceAsList[dataprocmetastoreservice.HiveMetastoreConfigAttributes](dms.ref.Append("hive_metastore_config"))
}

func (dms dataprocMetastoreServiceAttributes) MaintenanceWindow() terra.ListValue[dataprocmetastoreservice.MaintenanceWindowAttributes] {
	return terra.ReferenceAsList[dataprocmetastoreservice.MaintenanceWindowAttributes](dms.ref.Append("maintenance_window"))
}

func (dms dataprocMetastoreServiceAttributes) NetworkConfig() terra.ListValue[dataprocmetastoreservice.NetworkConfigAttributes] {
	return terra.ReferenceAsList[dataprocmetastoreservice.NetworkConfigAttributes](dms.ref.Append("network_config"))
}

func (dms dataprocMetastoreServiceAttributes) TelemetryConfig() terra.ListValue[dataprocmetastoreservice.TelemetryConfigAttributes] {
	return terra.ReferenceAsList[dataprocmetastoreservice.TelemetryConfigAttributes](dms.ref.Append("telemetry_config"))
}

func (dms dataprocMetastoreServiceAttributes) Timeouts() dataprocmetastoreservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprocmetastoreservice.TimeoutsAttributes](dms.ref.Append("timeouts"))
}

type dataprocMetastoreServiceState struct {
	ArtifactGcsUri      string                                              `json:"artifact_gcs_uri"`
	DatabaseType        string                                              `json:"database_type"`
	EndpointUri         string                                              `json:"endpoint_uri"`
	Id                  string                                              `json:"id"`
	Labels              map[string]string                                   `json:"labels"`
	Location            string                                              `json:"location"`
	Name                string                                              `json:"name"`
	Network             string                                              `json:"network"`
	Port                float64                                             `json:"port"`
	Project             string                                              `json:"project"`
	ReleaseChannel      string                                              `json:"release_channel"`
	ServiceId           string                                              `json:"service_id"`
	State               string                                              `json:"state"`
	StateMessage        string                                              `json:"state_message"`
	Tier                string                                              `json:"tier"`
	Uid                 string                                              `json:"uid"`
	EncryptionConfig    []dataprocmetastoreservice.EncryptionConfigState    `json:"encryption_config"`
	HiveMetastoreConfig []dataprocmetastoreservice.HiveMetastoreConfigState `json:"hive_metastore_config"`
	MaintenanceWindow   []dataprocmetastoreservice.MaintenanceWindowState   `json:"maintenance_window"`
	NetworkConfig       []dataprocmetastoreservice.NetworkConfigState       `json:"network_config"`
	TelemetryConfig     []dataprocmetastoreservice.TelemetryConfigState     `json:"telemetry_config"`
	Timeouts            *dataprocmetastoreservice.TimeoutsState             `json:"timeouts"`
}