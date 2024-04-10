// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewStorageNotification creates a new instance of [StorageNotification].
func NewStorageNotification(name string, args StorageNotificationArgs) *StorageNotification {
	return &StorageNotification{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageNotification)(nil)

// StorageNotification represents the Terraform resource google_storage_notification.
type StorageNotification struct {
	Name      string
	Args      StorageNotificationArgs
	state     *storageNotificationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageNotification].
func (sn *StorageNotification) Type() string {
	return "google_storage_notification"
}

// LocalName returns the local name for [StorageNotification].
func (sn *StorageNotification) LocalName() string {
	return sn.Name
}

// Configuration returns the configuration (args) for [StorageNotification].
func (sn *StorageNotification) Configuration() interface{} {
	return sn.Args
}

// DependOn is used for other resources to depend on [StorageNotification].
func (sn *StorageNotification) DependOn() terra.Reference {
	return terra.ReferenceResource(sn)
}

// Dependencies returns the list of resources [StorageNotification] depends_on.
func (sn *StorageNotification) Dependencies() terra.Dependencies {
	return sn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageNotification].
func (sn *StorageNotification) LifecycleManagement() *terra.Lifecycle {
	return sn.Lifecycle
}

// Attributes returns the attributes for [StorageNotification].
func (sn *StorageNotification) Attributes() storageNotificationAttributes {
	return storageNotificationAttributes{ref: terra.ReferenceResource(sn)}
}

// ImportState imports the given attribute values into [StorageNotification]'s state.
func (sn *StorageNotification) ImportState(av io.Reader) error {
	sn.state = &storageNotificationState{}
	if err := json.NewDecoder(av).Decode(sn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sn.Type(), sn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageNotification] has state.
func (sn *StorageNotification) State() (*storageNotificationState, bool) {
	return sn.state, sn.state != nil
}

// StateMust returns the state for [StorageNotification]. Panics if the state is nil.
func (sn *StorageNotification) StateMust() *storageNotificationState {
	if sn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sn.Type(), sn.LocalName()))
	}
	return sn.state
}

// StorageNotificationArgs contains the configurations for google_storage_notification.
type StorageNotificationArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// CustomAttributes: map of string, optional
	CustomAttributes terra.MapValue[terra.StringValue] `hcl:"custom_attributes,attr"`
	// EventTypes: set of string, optional
	EventTypes terra.SetValue[terra.StringValue] `hcl:"event_types,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ObjectNamePrefix: string, optional
	ObjectNamePrefix terra.StringValue `hcl:"object_name_prefix,attr"`
	// PayloadFormat: string, required
	PayloadFormat terra.StringValue `hcl:"payload_format,attr" validate:"required"`
	// Topic: string, required
	Topic terra.StringValue `hcl:"topic,attr" validate:"required"`
}
type storageNotificationAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of google_storage_notification.
func (sn storageNotificationAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("bucket"))
}

// CustomAttributes returns a reference to field custom_attributes of google_storage_notification.
func (sn storageNotificationAttributes) CustomAttributes() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sn.ref.Append("custom_attributes"))
}

// EventTypes returns a reference to field event_types of google_storage_notification.
func (sn storageNotificationAttributes) EventTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sn.ref.Append("event_types"))
}

// Id returns a reference to field id of google_storage_notification.
func (sn storageNotificationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("id"))
}

// NotificationId returns a reference to field notification_id of google_storage_notification.
func (sn storageNotificationAttributes) NotificationId() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("notification_id"))
}

// ObjectNamePrefix returns a reference to field object_name_prefix of google_storage_notification.
func (sn storageNotificationAttributes) ObjectNamePrefix() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("object_name_prefix"))
}

// PayloadFormat returns a reference to field payload_format of google_storage_notification.
func (sn storageNotificationAttributes) PayloadFormat() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("payload_format"))
}

// SelfLink returns a reference to field self_link of google_storage_notification.
func (sn storageNotificationAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("self_link"))
}

// Topic returns a reference to field topic of google_storage_notification.
func (sn storageNotificationAttributes) Topic() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("topic"))
}

type storageNotificationState struct {
	Bucket           string            `json:"bucket"`
	CustomAttributes map[string]string `json:"custom_attributes"`
	EventTypes       []string          `json:"event_types"`
	Id               string            `json:"id"`
	NotificationId   string            `json:"notification_id"`
	ObjectNamePrefix string            `json:"object_name_prefix"`
	PayloadFormat    string            `json:"payload_format"`
	SelfLink         string            `json:"self_link"`
	Topic            string            `json:"topic"`
}
