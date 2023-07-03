// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	filestorebackup "github.com/golingon/terraproviders/googlebeta/4.71.0/filestorebackup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFilestoreBackup creates a new instance of [FilestoreBackup].
func NewFilestoreBackup(name string, args FilestoreBackupArgs) *FilestoreBackup {
	return &FilestoreBackup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FilestoreBackup)(nil)

// FilestoreBackup represents the Terraform resource google_filestore_backup.
type FilestoreBackup struct {
	Name      string
	Args      FilestoreBackupArgs
	state     *filestoreBackupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FilestoreBackup].
func (fb *FilestoreBackup) Type() string {
	return "google_filestore_backup"
}

// LocalName returns the local name for [FilestoreBackup].
func (fb *FilestoreBackup) LocalName() string {
	return fb.Name
}

// Configuration returns the configuration (args) for [FilestoreBackup].
func (fb *FilestoreBackup) Configuration() interface{} {
	return fb.Args
}

// DependOn is used for other resources to depend on [FilestoreBackup].
func (fb *FilestoreBackup) DependOn() terra.Reference {
	return terra.ReferenceResource(fb)
}

// Dependencies returns the list of resources [FilestoreBackup] depends_on.
func (fb *FilestoreBackup) Dependencies() terra.Dependencies {
	return fb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FilestoreBackup].
func (fb *FilestoreBackup) LifecycleManagement() *terra.Lifecycle {
	return fb.Lifecycle
}

// Attributes returns the attributes for [FilestoreBackup].
func (fb *FilestoreBackup) Attributes() filestoreBackupAttributes {
	return filestoreBackupAttributes{ref: terra.ReferenceResource(fb)}
}

// ImportState imports the given attribute values into [FilestoreBackup]'s state.
func (fb *FilestoreBackup) ImportState(av io.Reader) error {
	fb.state = &filestoreBackupState{}
	if err := json.NewDecoder(av).Decode(fb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fb.Type(), fb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FilestoreBackup] has state.
func (fb *FilestoreBackup) State() (*filestoreBackupState, bool) {
	return fb.state, fb.state != nil
}

// StateMust returns the state for [FilestoreBackup]. Panics if the state is nil.
func (fb *FilestoreBackup) StateMust() *filestoreBackupState {
	if fb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fb.Type(), fb.LocalName()))
	}
	return fb.state
}

// FilestoreBackupArgs contains the configurations for google_filestore_backup.
type FilestoreBackupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SourceFileShare: string, required
	SourceFileShare terra.StringValue `hcl:"source_file_share,attr" validate:"required"`
	// SourceInstance: string, required
	SourceInstance terra.StringValue `hcl:"source_instance,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *filestorebackup.Timeouts `hcl:"timeouts,block"`
}
type filestoreBackupAttributes struct {
	ref terra.Reference
}

// CapacityGb returns a reference to field capacity_gb of google_filestore_backup.
func (fb filestoreBackupAttributes) CapacityGb() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("capacity_gb"))
}

// CreateTime returns a reference to field create_time of google_filestore_backup.
func (fb filestoreBackupAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("create_time"))
}

// Description returns a reference to field description of google_filestore_backup.
func (fb filestoreBackupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("description"))
}

// DownloadBytes returns a reference to field download_bytes of google_filestore_backup.
func (fb filestoreBackupAttributes) DownloadBytes() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("download_bytes"))
}

// Id returns a reference to field id of google_filestore_backup.
func (fb filestoreBackupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("id"))
}

// KmsKeyName returns a reference to field kms_key_name of google_filestore_backup.
func (fb filestoreBackupAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("kms_key_name"))
}

// Labels returns a reference to field labels of google_filestore_backup.
func (fb filestoreBackupAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fb.ref.Append("labels"))
}

// Location returns a reference to field location of google_filestore_backup.
func (fb filestoreBackupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("location"))
}

// Name returns a reference to field name of google_filestore_backup.
func (fb filestoreBackupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("name"))
}

// Project returns a reference to field project of google_filestore_backup.
func (fb filestoreBackupAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("project"))
}

// SourceFileShare returns a reference to field source_file_share of google_filestore_backup.
func (fb filestoreBackupAttributes) SourceFileShare() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("source_file_share"))
}

// SourceInstance returns a reference to field source_instance of google_filestore_backup.
func (fb filestoreBackupAttributes) SourceInstance() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("source_instance"))
}

// SourceInstanceTier returns a reference to field source_instance_tier of google_filestore_backup.
func (fb filestoreBackupAttributes) SourceInstanceTier() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("source_instance_tier"))
}

// State returns a reference to field state of google_filestore_backup.
func (fb filestoreBackupAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("state"))
}

// StorageBytes returns a reference to field storage_bytes of google_filestore_backup.
func (fb filestoreBackupAttributes) StorageBytes() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("storage_bytes"))
}

func (fb filestoreBackupAttributes) Timeouts() filestorebackup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[filestorebackup.TimeoutsAttributes](fb.ref.Append("timeouts"))
}

type filestoreBackupState struct {
	CapacityGb         string                         `json:"capacity_gb"`
	CreateTime         string                         `json:"create_time"`
	Description        string                         `json:"description"`
	DownloadBytes      string                         `json:"download_bytes"`
	Id                 string                         `json:"id"`
	KmsKeyName         string                         `json:"kms_key_name"`
	Labels             map[string]string              `json:"labels"`
	Location           string                         `json:"location"`
	Name               string                         `json:"name"`
	Project            string                         `json:"project"`
	SourceFileShare    string                         `json:"source_file_share"`
	SourceInstance     string                         `json:"source_instance"`
	SourceInstanceTier string                         `json:"source_instance_tier"`
	State              string                         `json:"state"`
	StorageBytes       string                         `json:"storage_bytes"`
	Timeouts           *filestorebackup.TimeoutsState `json:"timeouts"`
}
