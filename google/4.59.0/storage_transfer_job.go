// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	storagetransferjob "github.com/golingon/terraproviders/google/4.59.0/storagetransferjob"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewStorageTransferJob(name string, args StorageTransferJobArgs) *StorageTransferJob {
	return &StorageTransferJob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageTransferJob)(nil)

type StorageTransferJob struct {
	Name  string
	Args  StorageTransferJobArgs
	state *storageTransferJobState
}

func (stj *StorageTransferJob) Type() string {
	return "google_storage_transfer_job"
}

func (stj *StorageTransferJob) LocalName() string {
	return stj.Name
}

func (stj *StorageTransferJob) Configuration() interface{} {
	return stj.Args
}

func (stj *StorageTransferJob) Attributes() storageTransferJobAttributes {
	return storageTransferJobAttributes{ref: terra.ReferenceResource(stj)}
}

func (stj *StorageTransferJob) ImportState(av io.Reader) error {
	stj.state = &storageTransferJobState{}
	if err := json.NewDecoder(av).Decode(stj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", stj.Type(), stj.LocalName(), err)
	}
	return nil
}

func (stj *StorageTransferJob) State() (*storageTransferJobState, bool) {
	return stj.state, stj.state != nil
}

func (stj *StorageTransferJob) StateMust() *storageTransferJobState {
	if stj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", stj.Type(), stj.LocalName()))
	}
	return stj.state
}

func (stj *StorageTransferJob) DependOn() terra.Reference {
	return terra.ReferenceResource(stj)
}

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
	// DependsOn contains resources that StorageTransferJob depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type storageTransferJobAttributes struct {
	ref terra.Reference
}

func (stj storageTransferJobAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceString(stj.ref.Append("creation_time"))
}

func (stj storageTransferJobAttributes) DeletionTime() terra.StringValue {
	return terra.ReferenceString(stj.ref.Append("deletion_time"))
}

func (stj storageTransferJobAttributes) Description() terra.StringValue {
	return terra.ReferenceString(stj.ref.Append("description"))
}

func (stj storageTransferJobAttributes) Id() terra.StringValue {
	return terra.ReferenceString(stj.ref.Append("id"))
}

func (stj storageTransferJobAttributes) LastModificationTime() terra.StringValue {
	return terra.ReferenceString(stj.ref.Append("last_modification_time"))
}

func (stj storageTransferJobAttributes) Name() terra.StringValue {
	return terra.ReferenceString(stj.ref.Append("name"))
}

func (stj storageTransferJobAttributes) Project() terra.StringValue {
	return terra.ReferenceString(stj.ref.Append("project"))
}

func (stj storageTransferJobAttributes) Status() terra.StringValue {
	return terra.ReferenceString(stj.ref.Append("status"))
}

func (stj storageTransferJobAttributes) NotificationConfig() terra.ListValue[storagetransferjob.NotificationConfigAttributes] {
	return terra.ReferenceList[storagetransferjob.NotificationConfigAttributes](stj.ref.Append("notification_config"))
}

func (stj storageTransferJobAttributes) Schedule() terra.ListValue[storagetransferjob.ScheduleAttributes] {
	return terra.ReferenceList[storagetransferjob.ScheduleAttributes](stj.ref.Append("schedule"))
}

func (stj storageTransferJobAttributes) TransferSpec() terra.ListValue[storagetransferjob.TransferSpecAttributes] {
	return terra.ReferenceList[storagetransferjob.TransferSpecAttributes](stj.ref.Append("transfer_spec"))
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
