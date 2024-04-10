// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	streamanalyticsreferenceinputblob "github.com/golingon/terraproviders/azurerm/3.98.0/streamanalyticsreferenceinputblob"
	"io"
)

// NewStreamAnalyticsReferenceInputBlob creates a new instance of [StreamAnalyticsReferenceInputBlob].
func NewStreamAnalyticsReferenceInputBlob(name string, args StreamAnalyticsReferenceInputBlobArgs) *StreamAnalyticsReferenceInputBlob {
	return &StreamAnalyticsReferenceInputBlob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsReferenceInputBlob)(nil)

// StreamAnalyticsReferenceInputBlob represents the Terraform resource azurerm_stream_analytics_reference_input_blob.
type StreamAnalyticsReferenceInputBlob struct {
	Name      string
	Args      StreamAnalyticsReferenceInputBlobArgs
	state     *streamAnalyticsReferenceInputBlobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsReferenceInputBlob].
func (sarib *StreamAnalyticsReferenceInputBlob) Type() string {
	return "azurerm_stream_analytics_reference_input_blob"
}

// LocalName returns the local name for [StreamAnalyticsReferenceInputBlob].
func (sarib *StreamAnalyticsReferenceInputBlob) LocalName() string {
	return sarib.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsReferenceInputBlob].
func (sarib *StreamAnalyticsReferenceInputBlob) Configuration() interface{} {
	return sarib.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsReferenceInputBlob].
func (sarib *StreamAnalyticsReferenceInputBlob) DependOn() terra.Reference {
	return terra.ReferenceResource(sarib)
}

// Dependencies returns the list of resources [StreamAnalyticsReferenceInputBlob] depends_on.
func (sarib *StreamAnalyticsReferenceInputBlob) Dependencies() terra.Dependencies {
	return sarib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsReferenceInputBlob].
func (sarib *StreamAnalyticsReferenceInputBlob) LifecycleManagement() *terra.Lifecycle {
	return sarib.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsReferenceInputBlob].
func (sarib *StreamAnalyticsReferenceInputBlob) Attributes() streamAnalyticsReferenceInputBlobAttributes {
	return streamAnalyticsReferenceInputBlobAttributes{ref: terra.ReferenceResource(sarib)}
}

// ImportState imports the given attribute values into [StreamAnalyticsReferenceInputBlob]'s state.
func (sarib *StreamAnalyticsReferenceInputBlob) ImportState(av io.Reader) error {
	sarib.state = &streamAnalyticsReferenceInputBlobState{}
	if err := json.NewDecoder(av).Decode(sarib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sarib.Type(), sarib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsReferenceInputBlob] has state.
func (sarib *StreamAnalyticsReferenceInputBlob) State() (*streamAnalyticsReferenceInputBlobState, bool) {
	return sarib.state, sarib.state != nil
}

// StateMust returns the state for [StreamAnalyticsReferenceInputBlob]. Panics if the state is nil.
func (sarib *StreamAnalyticsReferenceInputBlob) StateMust() *streamAnalyticsReferenceInputBlobState {
	if sarib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sarib.Type(), sarib.LocalName()))
	}
	return sarib.state
}

// StreamAnalyticsReferenceInputBlobArgs contains the configurations for azurerm_stream_analytics_reference_input_blob.
type StreamAnalyticsReferenceInputBlobArgs struct {
	// AuthenticationMode: string, optional
	AuthenticationMode terra.StringValue `hcl:"authentication_mode,attr"`
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
	Serialization *streamanalyticsreferenceinputblob.Serialization `hcl:"serialization,block" validate:"required"`
	// Timeouts: optional
	Timeouts *streamanalyticsreferenceinputblob.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsReferenceInputBlobAttributes struct {
	ref terra.Reference
}

// AuthenticationMode returns a reference to field authentication_mode of azurerm_stream_analytics_reference_input_blob.
func (sarib streamAnalyticsReferenceInputBlobAttributes) AuthenticationMode() terra.StringValue {
	return terra.ReferenceAsString(sarib.ref.Append("authentication_mode"))
}

// DateFormat returns a reference to field date_format of azurerm_stream_analytics_reference_input_blob.
func (sarib streamAnalyticsReferenceInputBlobAttributes) DateFormat() terra.StringValue {
	return terra.ReferenceAsString(sarib.ref.Append("date_format"))
}

// Id returns a reference to field id of azurerm_stream_analytics_reference_input_blob.
func (sarib streamAnalyticsReferenceInputBlobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sarib.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_stream_analytics_reference_input_blob.
func (sarib streamAnalyticsReferenceInputBlobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sarib.ref.Append("name"))
}

// PathPattern returns a reference to field path_pattern of azurerm_stream_analytics_reference_input_blob.
func (sarib streamAnalyticsReferenceInputBlobAttributes) PathPattern() terra.StringValue {
	return terra.ReferenceAsString(sarib.ref.Append("path_pattern"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stream_analytics_reference_input_blob.
func (sarib streamAnalyticsReferenceInputBlobAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sarib.ref.Append("resource_group_name"))
}

// StorageAccountKey returns a reference to field storage_account_key of azurerm_stream_analytics_reference_input_blob.
func (sarib streamAnalyticsReferenceInputBlobAttributes) StorageAccountKey() terra.StringValue {
	return terra.ReferenceAsString(sarib.ref.Append("storage_account_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_stream_analytics_reference_input_blob.
func (sarib streamAnalyticsReferenceInputBlobAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(sarib.ref.Append("storage_account_name"))
}

// StorageContainerName returns a reference to field storage_container_name of azurerm_stream_analytics_reference_input_blob.
func (sarib streamAnalyticsReferenceInputBlobAttributes) StorageContainerName() terra.StringValue {
	return terra.ReferenceAsString(sarib.ref.Append("storage_container_name"))
}

// StreamAnalyticsJobName returns a reference to field stream_analytics_job_name of azurerm_stream_analytics_reference_input_blob.
func (sarib streamAnalyticsReferenceInputBlobAttributes) StreamAnalyticsJobName() terra.StringValue {
	return terra.ReferenceAsString(sarib.ref.Append("stream_analytics_job_name"))
}

// TimeFormat returns a reference to field time_format of azurerm_stream_analytics_reference_input_blob.
func (sarib streamAnalyticsReferenceInputBlobAttributes) TimeFormat() terra.StringValue {
	return terra.ReferenceAsString(sarib.ref.Append("time_format"))
}

func (sarib streamAnalyticsReferenceInputBlobAttributes) Serialization() terra.ListValue[streamanalyticsreferenceinputblob.SerializationAttributes] {
	return terra.ReferenceAsList[streamanalyticsreferenceinputblob.SerializationAttributes](sarib.ref.Append("serialization"))
}

func (sarib streamAnalyticsReferenceInputBlobAttributes) Timeouts() streamanalyticsreferenceinputblob.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticsreferenceinputblob.TimeoutsAttributes](sarib.ref.Append("timeouts"))
}

type streamAnalyticsReferenceInputBlobState struct {
	AuthenticationMode     string                                                 `json:"authentication_mode"`
	DateFormat             string                                                 `json:"date_format"`
	Id                     string                                                 `json:"id"`
	Name                   string                                                 `json:"name"`
	PathPattern            string                                                 `json:"path_pattern"`
	ResourceGroupName      string                                                 `json:"resource_group_name"`
	StorageAccountKey      string                                                 `json:"storage_account_key"`
	StorageAccountName     string                                                 `json:"storage_account_name"`
	StorageContainerName   string                                                 `json:"storage_container_name"`
	StreamAnalyticsJobName string                                                 `json:"stream_analytics_job_name"`
	TimeFormat             string                                                 `json:"time_format"`
	Serialization          []streamanalyticsreferenceinputblob.SerializationState `json:"serialization"`
	Timeouts               *streamanalyticsreferenceinputblob.TimeoutsState       `json:"timeouts"`
}
