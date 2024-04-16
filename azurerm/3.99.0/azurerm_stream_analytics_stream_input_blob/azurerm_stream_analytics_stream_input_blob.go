// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_stream_analytics_stream_input_blob

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_stream_analytics_stream_input_blob.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermStreamAnalyticsStreamInputBlobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asasib *Resource) Type() string {
	return "azurerm_stream_analytics_stream_input_blob"
}

// LocalName returns the local name for [Resource].
func (asasib *Resource) LocalName() string {
	return asasib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asasib *Resource) Configuration() interface{} {
	return asasib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asasib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asasib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asasib *Resource) Dependencies() terra.Dependencies {
	return asasib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asasib *Resource) LifecycleManagement() *terra.Lifecycle {
	return asasib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asasib *Resource) Attributes() azurermStreamAnalyticsStreamInputBlobAttributes {
	return azurermStreamAnalyticsStreamInputBlobAttributes{ref: terra.ReferenceResource(asasib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asasib *Resource) ImportState(state io.Reader) error {
	asasib.state = &azurermStreamAnalyticsStreamInputBlobState{}
	if err := json.NewDecoder(state).Decode(asasib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asasib.Type(), asasib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asasib *Resource) State() (*azurermStreamAnalyticsStreamInputBlobState, bool) {
	return asasib.state, asasib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asasib *Resource) StateMust() *azurermStreamAnalyticsStreamInputBlobState {
	if asasib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asasib.Type(), asasib.LocalName()))
	}
	return asasib.state
}

// Args contains the configurations for azurerm_stream_analytics_stream_input_blob.
type Args struct {
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
	// StorageAccountKey: string, required
	StorageAccountKey terra.StringValue `hcl:"storage_account_key,attr" validate:"required"`
	// StorageAccountName: string, required
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr" validate:"required"`
	// StorageContainerName: string, required
	StorageContainerName terra.StringValue `hcl:"storage_container_name,attr" validate:"required"`
	// StreamAnalyticsJobName: string, required
	StreamAnalyticsJobName terra.StringValue `hcl:"stream_analytics_job_name,attr" validate:"required"`
	// TimeFormat: string, required
	TimeFormat terra.StringValue `hcl:"time_format,attr" validate:"required"`
	// Serialization: required
	Serialization *Serialization `hcl:"serialization,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermStreamAnalyticsStreamInputBlobAttributes struct {
	ref terra.Reference
}

// DateFormat returns a reference to field date_format of azurerm_stream_analytics_stream_input_blob.
func (asasib azurermStreamAnalyticsStreamInputBlobAttributes) DateFormat() terra.StringValue {
	return terra.ReferenceAsString(asasib.ref.Append("date_format"))
}

// Id returns a reference to field id of azurerm_stream_analytics_stream_input_blob.
func (asasib azurermStreamAnalyticsStreamInputBlobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asasib.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_stream_analytics_stream_input_blob.
func (asasib azurermStreamAnalyticsStreamInputBlobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asasib.ref.Append("name"))
}

// PathPattern returns a reference to field path_pattern of azurerm_stream_analytics_stream_input_blob.
func (asasib azurermStreamAnalyticsStreamInputBlobAttributes) PathPattern() terra.StringValue {
	return terra.ReferenceAsString(asasib.ref.Append("path_pattern"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stream_analytics_stream_input_blob.
func (asasib azurermStreamAnalyticsStreamInputBlobAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(asasib.ref.Append("resource_group_name"))
}

// StorageAccountKey returns a reference to field storage_account_key of azurerm_stream_analytics_stream_input_blob.
func (asasib azurermStreamAnalyticsStreamInputBlobAttributes) StorageAccountKey() terra.StringValue {
	return terra.ReferenceAsString(asasib.ref.Append("storage_account_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_stream_analytics_stream_input_blob.
func (asasib azurermStreamAnalyticsStreamInputBlobAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(asasib.ref.Append("storage_account_name"))
}

// StorageContainerName returns a reference to field storage_container_name of azurerm_stream_analytics_stream_input_blob.
func (asasib azurermStreamAnalyticsStreamInputBlobAttributes) StorageContainerName() terra.StringValue {
	return terra.ReferenceAsString(asasib.ref.Append("storage_container_name"))
}

// StreamAnalyticsJobName returns a reference to field stream_analytics_job_name of azurerm_stream_analytics_stream_input_blob.
func (asasib azurermStreamAnalyticsStreamInputBlobAttributes) StreamAnalyticsJobName() terra.StringValue {
	return terra.ReferenceAsString(asasib.ref.Append("stream_analytics_job_name"))
}

// TimeFormat returns a reference to field time_format of azurerm_stream_analytics_stream_input_blob.
func (asasib azurermStreamAnalyticsStreamInputBlobAttributes) TimeFormat() terra.StringValue {
	return terra.ReferenceAsString(asasib.ref.Append("time_format"))
}

func (asasib azurermStreamAnalyticsStreamInputBlobAttributes) Serialization() terra.ListValue[SerializationAttributes] {
	return terra.ReferenceAsList[SerializationAttributes](asasib.ref.Append("serialization"))
}

func (asasib azurermStreamAnalyticsStreamInputBlobAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](asasib.ref.Append("timeouts"))
}

type azurermStreamAnalyticsStreamInputBlobState struct {
	DateFormat             string               `json:"date_format"`
	Id                     string               `json:"id"`
	Name                   string               `json:"name"`
	PathPattern            string               `json:"path_pattern"`
	ResourceGroupName      string               `json:"resource_group_name"`
	StorageAccountKey      string               `json:"storage_account_key"`
	StorageAccountName     string               `json:"storage_account_name"`
	StorageContainerName   string               `json:"storage_container_name"`
	StreamAnalyticsJobName string               `json:"stream_analytics_job_name"`
	TimeFormat             string               `json:"time_format"`
	Serialization          []SerializationState `json:"serialization"`
	Timeouts               *TimeoutsState       `json:"timeouts"`
}
