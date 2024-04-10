// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	streamanalyticsoutputblob "github.com/golingon/terraproviders/azurerm/3.98.0/streamanalyticsoutputblob"
	"io"
)

// NewStreamAnalyticsOutputBlob creates a new instance of [StreamAnalyticsOutputBlob].
func NewStreamAnalyticsOutputBlob(name string, args StreamAnalyticsOutputBlobArgs) *StreamAnalyticsOutputBlob {
	return &StreamAnalyticsOutputBlob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsOutputBlob)(nil)

// StreamAnalyticsOutputBlob represents the Terraform resource azurerm_stream_analytics_output_blob.
type StreamAnalyticsOutputBlob struct {
	Name      string
	Args      StreamAnalyticsOutputBlobArgs
	state     *streamAnalyticsOutputBlobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsOutputBlob].
func (saob *StreamAnalyticsOutputBlob) Type() string {
	return "azurerm_stream_analytics_output_blob"
}

// LocalName returns the local name for [StreamAnalyticsOutputBlob].
func (saob *StreamAnalyticsOutputBlob) LocalName() string {
	return saob.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsOutputBlob].
func (saob *StreamAnalyticsOutputBlob) Configuration() interface{} {
	return saob.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsOutputBlob].
func (saob *StreamAnalyticsOutputBlob) DependOn() terra.Reference {
	return terra.ReferenceResource(saob)
}

// Dependencies returns the list of resources [StreamAnalyticsOutputBlob] depends_on.
func (saob *StreamAnalyticsOutputBlob) Dependencies() terra.Dependencies {
	return saob.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsOutputBlob].
func (saob *StreamAnalyticsOutputBlob) LifecycleManagement() *terra.Lifecycle {
	return saob.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsOutputBlob].
func (saob *StreamAnalyticsOutputBlob) Attributes() streamAnalyticsOutputBlobAttributes {
	return streamAnalyticsOutputBlobAttributes{ref: terra.ReferenceResource(saob)}
}

// ImportState imports the given attribute values into [StreamAnalyticsOutputBlob]'s state.
func (saob *StreamAnalyticsOutputBlob) ImportState(av io.Reader) error {
	saob.state = &streamAnalyticsOutputBlobState{}
	if err := json.NewDecoder(av).Decode(saob.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", saob.Type(), saob.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsOutputBlob] has state.
func (saob *StreamAnalyticsOutputBlob) State() (*streamAnalyticsOutputBlobState, bool) {
	return saob.state, saob.state != nil
}

// StateMust returns the state for [StreamAnalyticsOutputBlob]. Panics if the state is nil.
func (saob *StreamAnalyticsOutputBlob) StateMust() *streamAnalyticsOutputBlobState {
	if saob.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", saob.Type(), saob.LocalName()))
	}
	return saob.state
}

// StreamAnalyticsOutputBlobArgs contains the configurations for azurerm_stream_analytics_output_blob.
type StreamAnalyticsOutputBlobArgs struct {
	// AuthenticationMode: string, optional
	AuthenticationMode terra.StringValue `hcl:"authentication_mode,attr"`
	// BatchMaxWaitTime: string, optional
	BatchMaxWaitTime terra.StringValue `hcl:"batch_max_wait_time,attr"`
	// BatchMinRows: number, optional
	BatchMinRows terra.NumberValue `hcl:"batch_min_rows,attr"`
	// BlobWriteMode: string, optional
	BlobWriteMode terra.StringValue `hcl:"blob_write_mode,attr"`
	// DateFormat: string, required
	DateFormat terra.StringValue `hcl:"date_format,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PathPattern: string, required
	PathPattern terra.StringValue `hcl:"path_pattern,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StorageAccountKey: string, optional
	StorageAccountKey terra.StringValue `hcl:"storage_account_key,attr"`
	// StorageAccountName: string, required
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr" validate:"required"`
	// StorageContainerName: string, required
	StorageContainerName terra.StringValue `hcl:"storage_container_name,attr" validate:"required"`
	// StreamAnalyticsJobName: string, required
	StreamAnalyticsJobName terra.StringValue `hcl:"stream_analytics_job_name,attr" validate:"required"`
	// TimeFormat: string, required
	TimeFormat terra.StringValue `hcl:"time_format,attr" validate:"required"`
	// Serialization: required
	Serialization *streamanalyticsoutputblob.Serialization `hcl:"serialization,block" validate:"required"`
	// Timeouts: optional
	Timeouts *streamanalyticsoutputblob.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsOutputBlobAttributes struct {
	ref terra.Reference
}

// AuthenticationMode returns a reference to field authentication_mode of azurerm_stream_analytics_output_blob.
func (saob streamAnalyticsOutputBlobAttributes) AuthenticationMode() terra.StringValue {
	return terra.ReferenceAsString(saob.ref.Append("authentication_mode"))
}

// BatchMaxWaitTime returns a reference to field batch_max_wait_time of azurerm_stream_analytics_output_blob.
func (saob streamAnalyticsOutputBlobAttributes) BatchMaxWaitTime() terra.StringValue {
	return terra.ReferenceAsString(saob.ref.Append("batch_max_wait_time"))
}

// BatchMinRows returns a reference to field batch_min_rows of azurerm_stream_analytics_output_blob.
func (saob streamAnalyticsOutputBlobAttributes) BatchMinRows() terra.NumberValue {
	return terra.ReferenceAsNumber(saob.ref.Append("batch_min_rows"))
}

// BlobWriteMode returns a reference to field blob_write_mode of azurerm_stream_analytics_output_blob.
func (saob streamAnalyticsOutputBlobAttributes) BlobWriteMode() terra.StringValue {
	return terra.ReferenceAsString(saob.ref.Append("blob_write_mode"))
}

// DateFormat returns a reference to field date_format of azurerm_stream_analytics_output_blob.
func (saob streamAnalyticsOutputBlobAttributes) DateFormat() terra.StringValue {
	return terra.ReferenceAsString(saob.ref.Append("date_format"))
}

// Id returns a reference to field id of azurerm_stream_analytics_output_blob.
func (saob streamAnalyticsOutputBlobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(saob.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_stream_analytics_output_blob.
func (saob streamAnalyticsOutputBlobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(saob.ref.Append("name"))
}

// PathPattern returns a reference to field path_pattern of azurerm_stream_analytics_output_blob.
func (saob streamAnalyticsOutputBlobAttributes) PathPattern() terra.StringValue {
	return terra.ReferenceAsString(saob.ref.Append("path_pattern"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stream_analytics_output_blob.
func (saob streamAnalyticsOutputBlobAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(saob.ref.Append("resource_group_name"))
}

// StorageAccountKey returns a reference to field storage_account_key of azurerm_stream_analytics_output_blob.
func (saob streamAnalyticsOutputBlobAttributes) StorageAccountKey() terra.StringValue {
	return terra.ReferenceAsString(saob.ref.Append("storage_account_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_stream_analytics_output_blob.
func (saob streamAnalyticsOutputBlobAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(saob.ref.Append("storage_account_name"))
}

// StorageContainerName returns a reference to field storage_container_name of azurerm_stream_analytics_output_blob.
func (saob streamAnalyticsOutputBlobAttributes) StorageContainerName() terra.StringValue {
	return terra.ReferenceAsString(saob.ref.Append("storage_container_name"))
}

// StreamAnalyticsJobName returns a reference to field stream_analytics_job_name of azurerm_stream_analytics_output_blob.
func (saob streamAnalyticsOutputBlobAttributes) StreamAnalyticsJobName() terra.StringValue {
	return terra.ReferenceAsString(saob.ref.Append("stream_analytics_job_name"))
}

// TimeFormat returns a reference to field time_format of azurerm_stream_analytics_output_blob.
func (saob streamAnalyticsOutputBlobAttributes) TimeFormat() terra.StringValue {
	return terra.ReferenceAsString(saob.ref.Append("time_format"))
}

func (saob streamAnalyticsOutputBlobAttributes) Serialization() terra.ListValue[streamanalyticsoutputblob.SerializationAttributes] {
	return terra.ReferenceAsList[streamanalyticsoutputblob.SerializationAttributes](saob.ref.Append("serialization"))
}

func (saob streamAnalyticsOutputBlobAttributes) Timeouts() streamanalyticsoutputblob.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticsoutputblob.TimeoutsAttributes](saob.ref.Append("timeouts"))
}

type streamAnalyticsOutputBlobState struct {
	AuthenticationMode     string                                         `json:"authentication_mode"`
	BatchMaxWaitTime       string                                         `json:"batch_max_wait_time"`
	BatchMinRows           float64                                        `json:"batch_min_rows"`
	BlobWriteMode          string                                         `json:"blob_write_mode"`
	DateFormat             string                                         `json:"date_format"`
	Id                     string                                         `json:"id"`
	Name                   string                                         `json:"name"`
	PathPattern            string                                         `json:"path_pattern"`
	ResourceGroupName      string                                         `json:"resource_group_name"`
	StorageAccountKey      string                                         `json:"storage_account_key"`
	StorageAccountName     string                                         `json:"storage_account_name"`
	StorageContainerName   string                                         `json:"storage_container_name"`
	StreamAnalyticsJobName string                                         `json:"stream_analytics_job_name"`
	TimeFormat             string                                         `json:"time_format"`
	Serialization          []streamanalyticsoutputblob.SerializationState `json:"serialization"`
	Timeouts               *streamanalyticsoutputblob.TimeoutsState       `json:"timeouts"`
}
