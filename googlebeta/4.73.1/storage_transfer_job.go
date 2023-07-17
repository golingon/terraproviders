// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	storagetransferjob "github.com/golingon/terraproviders/googlebeta/4.73.1/storagetransferjob"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageTransferJob creates a new instance of [StorageTransferJob].
func NewStorageTransferJob(name string, args StorageTransferJobArgs) *StorageTransferJob {
	return &StorageTransferJob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageTransferJob)(nil)

// StorageTransferJob represents the Terraform resource google_storage_transfer_job.
type StorageTransferJob struct {
	Name      string
	Args      StorageTransferJobArgs
	state     *storageTransferJobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageTransferJob].
func (stj *StorageTransferJob) Type() string {
	return "google_storage_transfer_job"
}

// LocalName returns the local name for [StorageTransferJob].
func (stj *StorageTransferJob) LocalName() string {
	return stj.Name
}

// Configuration returns the configuration (args) for [StorageTransferJob].
func (stj *StorageTransferJob) Configuration() interface{} {
	return stj.Args
}

// DependOn is used for other resources to depend on [StorageTransferJob].
func (stj *StorageTransferJob) DependOn() terra.Reference {
	return terra.ReferenceResource(stj)
}

// Dependencies returns the list of resources [StorageTransferJob] depends_on.
func (stj *StorageTransferJob) Dependencies() terra.Dependencies {
	return stj.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageTransferJob].
func (stj *StorageTransferJob) LifecycleManagement() *terra.Lifecycle {
	return stj.Lifecycle
}

// Attributes returns the attributes for [StorageTransferJob].
func (stj *StorageTransferJob) Attributes() storageTransferJobAttributes {
	return storageTransferJobAttributes{ref: terra.ReferenceResource(stj)}
}

// ImportState imports the given attribute values into [StorageTransferJob]'s state.
func (stj *StorageTransferJob) ImportState(av io.Reader) error {
	stj.state = &storageTransferJobState{}
	if err := json.NewDecoder(av).Decode(stj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", stj.Type(), stj.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageTransferJob] has state.
func (stj *StorageTransferJob) State() (*storageTransferJobState, bool) {
	return stj.state, stj.state != nil
}

// StateMust returns the state for [StorageTransferJob]. Panics if the state is nil.
func (stj *StorageTransferJob) StateMust() *storageTransferJobState {
	if stj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", stj.Type(), stj.LocalName()))
	}
	return stj.state
}

// StorageTransferJobArgs contains the configurations for google_storage_transfer_job.
type StorageTransferJobArgs struct {
	// Description: string, required
	Description terra.StringValue `hcl:"description,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// NotificationConfig: optional
	NotificationConfig *storagetransferjob.NotificationConfig `hcl:"notification_config,block"`
	// Schedule: optional
	Schedule *storagetransferjob.Schedule `hcl:"schedule,block"`
	// TransferSpec: required
	TransferSpec *storagetransferjob.TransferSpec `hcl:"transfer_spec,block" validate:"required"`
}
type storageTransferJobAttributes struct {
	ref terra.Reference
}

// CreationTime returns a reference to field creation_time of google_storage_transfer_job.
func (stj storageTransferJobAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(stj.ref.Append("creation_time"))
}

// DeletionTime returns a reference to field deletion_time of google_storage_transfer_job.
func (stj storageTransferJobAttributes) DeletionTime() terra.StringValue {
	return terra.ReferenceAsString(stj.ref.Append("deletion_time"))
}

// Description returns a reference to field description of google_storage_transfer_job.
func (stj storageTransferJobAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(stj.ref.Append("description"))
}

// Id returns a reference to field id of google_storage_transfer_job.
func (stj storageTransferJobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(stj.ref.Append("id"))
}

// LastModificationTime returns a reference to field last_modification_time of google_storage_transfer_job.
func (stj storageTransferJobAttributes) LastModificationTime() terra.StringValue {
	return terra.ReferenceAsString(stj.ref.Append("last_modification_time"))
}

// Name returns a reference to field name of google_storage_transfer_job.
func (stj storageTransferJobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(stj.ref.Append("name"))
}

// Project returns a reference to field project of google_storage_transfer_job.
func (stj storageTransferJobAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(stj.ref.Append("project"))
}

// Status returns a reference to field status of google_storage_transfer_job.
func (stj storageTransferJobAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(stj.ref.Append("status"))
}

func (stj storageTransferJobAttributes) NotificationConfig() terra.ListValue[storagetransferjob.NotificationConfigAttributes] {
	return terra.ReferenceAsList[storagetransferjob.NotificationConfigAttributes](stj.ref.Append("notification_config"))
}

func (stj storageTransferJobAttributes) Schedule() terra.ListValue[storagetransferjob.ScheduleAttributes] {
	return terra.ReferenceAsList[storagetransferjob.ScheduleAttributes](stj.ref.Append("schedule"))
}

func (stj storageTransferJobAttributes) TransferSpec() terra.ListValue[storagetransferjob.TransferSpecAttributes] {
	return terra.ReferenceAsList[storagetransferjob.TransferSpecAttributes](stj.ref.Append("transfer_spec"))
}

type storageTransferJobState struct {
	CreationTime         string                                       `json:"creation_time"`
	DeletionTime         string                                       `json:"deletion_time"`
	Description          string                                       `json:"description"`
	Id                   string                                       `json:"id"`
	LastModificationTime string                                       `json:"last_modification_time"`
	Name                 string                                       `json:"name"`
	Project              string                                       `json:"project"`
	Status               string                                       `json:"status"`
	NotificationConfig   []storagetransferjob.NotificationConfigState `json:"notification_config"`
	Schedule             []storagetransferjob.ScheduleState           `json:"schedule"`
	TransferSpec         []storagetransferjob.TransferSpecState       `json:"transfer_spec"`
}
