// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iothubfileupload "github.com/golingon/terraproviders/azurerm/3.63.0/iothubfileupload"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIothubFileUpload creates a new instance of [IothubFileUpload].
func NewIothubFileUpload(name string, args IothubFileUploadArgs) *IothubFileUpload {
	return &IothubFileUpload{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IothubFileUpload)(nil)

// IothubFileUpload represents the Terraform resource azurerm_iothub_file_upload.
type IothubFileUpload struct {
	Name      string
	Args      IothubFileUploadArgs
	state     *iothubFileUploadState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IothubFileUpload].
func (ifu *IothubFileUpload) Type() string {
	return "azurerm_iothub_file_upload"
}

// LocalName returns the local name for [IothubFileUpload].
func (ifu *IothubFileUpload) LocalName() string {
	return ifu.Name
}

// Configuration returns the configuration (args) for [IothubFileUpload].
func (ifu *IothubFileUpload) Configuration() interface{} {
	return ifu.Args
}

// DependOn is used for other resources to depend on [IothubFileUpload].
func (ifu *IothubFileUpload) DependOn() terra.Reference {
	return terra.ReferenceResource(ifu)
}

// Dependencies returns the list of resources [IothubFileUpload] depends_on.
func (ifu *IothubFileUpload) Dependencies() terra.Dependencies {
	return ifu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IothubFileUpload].
func (ifu *IothubFileUpload) LifecycleManagement() *terra.Lifecycle {
	return ifu.Lifecycle
}

// Attributes returns the attributes for [IothubFileUpload].
func (ifu *IothubFileUpload) Attributes() iothubFileUploadAttributes {
	return iothubFileUploadAttributes{ref: terra.ReferenceResource(ifu)}
}

// ImportState imports the given attribute values into [IothubFileUpload]'s state.
func (ifu *IothubFileUpload) ImportState(av io.Reader) error {
	ifu.state = &iothubFileUploadState{}
	if err := json.NewDecoder(av).Decode(ifu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ifu.Type(), ifu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IothubFileUpload] has state.
func (ifu *IothubFileUpload) State() (*iothubFileUploadState, bool) {
	return ifu.state, ifu.state != nil
}

// StateMust returns the state for [IothubFileUpload]. Panics if the state is nil.
func (ifu *IothubFileUpload) StateMust() *iothubFileUploadState {
	if ifu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ifu.Type(), ifu.LocalName()))
	}
	return ifu.state
}

// IothubFileUploadArgs contains the configurations for azurerm_iothub_file_upload.
type IothubFileUploadArgs struct {
	// AuthenticationType: string, optional
	AuthenticationType terra.StringValue `hcl:"authentication_type,attr"`
	// ConnectionString: string, required
	ConnectionString terra.StringValue `hcl:"connection_string,attr" validate:"required"`
	// ContainerName: string, required
	ContainerName terra.StringValue `hcl:"container_name,attr" validate:"required"`
	// DefaultTtl: string, optional
	DefaultTtl terra.StringValue `hcl:"default_ttl,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityId: string, optional
	IdentityId terra.StringValue `hcl:"identity_id,attr"`
	// IothubId: string, required
	IothubId terra.StringValue `hcl:"iothub_id,attr" validate:"required"`
	// LockDuration: string, optional
	LockDuration terra.StringValue `hcl:"lock_duration,attr"`
	// MaxDeliveryCount: number, optional
	MaxDeliveryCount terra.NumberValue `hcl:"max_delivery_count,attr"`
	// NotificationsEnabled: bool, optional
	NotificationsEnabled terra.BoolValue `hcl:"notifications_enabled,attr"`
	// SasTtl: string, optional
	SasTtl terra.StringValue `hcl:"sas_ttl,attr"`
	// Timeouts: optional
	Timeouts *iothubfileupload.Timeouts `hcl:"timeouts,block"`
}
type iothubFileUploadAttributes struct {
	ref terra.Reference
}

// AuthenticationType returns a reference to field authentication_type of azurerm_iothub_file_upload.
func (ifu iothubFileUploadAttributes) AuthenticationType() terra.StringValue {
	return terra.ReferenceAsString(ifu.ref.Append("authentication_type"))
}

// ConnectionString returns a reference to field connection_string of azurerm_iothub_file_upload.
func (ifu iothubFileUploadAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(ifu.ref.Append("connection_string"))
}

// ContainerName returns a reference to field container_name of azurerm_iothub_file_upload.
func (ifu iothubFileUploadAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(ifu.ref.Append("container_name"))
}

// DefaultTtl returns a reference to field default_ttl of azurerm_iothub_file_upload.
func (ifu iothubFileUploadAttributes) DefaultTtl() terra.StringValue {
	return terra.ReferenceAsString(ifu.ref.Append("default_ttl"))
}

// Id returns a reference to field id of azurerm_iothub_file_upload.
func (ifu iothubFileUploadAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ifu.ref.Append("id"))
}

// IdentityId returns a reference to field identity_id of azurerm_iothub_file_upload.
func (ifu iothubFileUploadAttributes) IdentityId() terra.StringValue {
	return terra.ReferenceAsString(ifu.ref.Append("identity_id"))
}

// IothubId returns a reference to field iothub_id of azurerm_iothub_file_upload.
func (ifu iothubFileUploadAttributes) IothubId() terra.StringValue {
	return terra.ReferenceAsString(ifu.ref.Append("iothub_id"))
}

// LockDuration returns a reference to field lock_duration of azurerm_iothub_file_upload.
func (ifu iothubFileUploadAttributes) LockDuration() terra.StringValue {
	return terra.ReferenceAsString(ifu.ref.Append("lock_duration"))
}

// MaxDeliveryCount returns a reference to field max_delivery_count of azurerm_iothub_file_upload.
func (ifu iothubFileUploadAttributes) MaxDeliveryCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ifu.ref.Append("max_delivery_count"))
}

// NotificationsEnabled returns a reference to field notifications_enabled of azurerm_iothub_file_upload.
func (ifu iothubFileUploadAttributes) NotificationsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ifu.ref.Append("notifications_enabled"))
}

// SasTtl returns a reference to field sas_ttl of azurerm_iothub_file_upload.
func (ifu iothubFileUploadAttributes) SasTtl() terra.StringValue {
	return terra.ReferenceAsString(ifu.ref.Append("sas_ttl"))
}

func (ifu iothubFileUploadAttributes) Timeouts() iothubfileupload.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iothubfileupload.TimeoutsAttributes](ifu.ref.Append("timeouts"))
}

type iothubFileUploadState struct {
	AuthenticationType   string                          `json:"authentication_type"`
	ConnectionString     string                          `json:"connection_string"`
	ContainerName        string                          `json:"container_name"`
	DefaultTtl           string                          `json:"default_ttl"`
	Id                   string                          `json:"id"`
	IdentityId           string                          `json:"identity_id"`
	IothubId             string                          `json:"iothub_id"`
	LockDuration         string                          `json:"lock_duration"`
	MaxDeliveryCount     float64                         `json:"max_delivery_count"`
	NotificationsEnabled bool                            `json:"notifications_enabled"`
	SasTtl               string                          `json:"sas_ttl"`
	Timeouts             *iothubfileupload.TimeoutsState `json:"timeouts"`
}
