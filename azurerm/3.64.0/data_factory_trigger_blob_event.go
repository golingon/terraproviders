// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorytriggerblobevent "github.com/golingon/terraproviders/azurerm/3.64.0/datafactorytriggerblobevent"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryTriggerBlobEvent creates a new instance of [DataFactoryTriggerBlobEvent].
func NewDataFactoryTriggerBlobEvent(name string, args DataFactoryTriggerBlobEventArgs) *DataFactoryTriggerBlobEvent {
	return &DataFactoryTriggerBlobEvent{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryTriggerBlobEvent)(nil)

// DataFactoryTriggerBlobEvent represents the Terraform resource azurerm_data_factory_trigger_blob_event.
type DataFactoryTriggerBlobEvent struct {
	Name      string
	Args      DataFactoryTriggerBlobEventArgs
	state     *dataFactoryTriggerBlobEventState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryTriggerBlobEvent].
func (dftbe *DataFactoryTriggerBlobEvent) Type() string {
	return "azurerm_data_factory_trigger_blob_event"
}

// LocalName returns the local name for [DataFactoryTriggerBlobEvent].
func (dftbe *DataFactoryTriggerBlobEvent) LocalName() string {
	return dftbe.Name
}

// Configuration returns the configuration (args) for [DataFactoryTriggerBlobEvent].
func (dftbe *DataFactoryTriggerBlobEvent) Configuration() interface{} {
	return dftbe.Args
}

// DependOn is used for other resources to depend on [DataFactoryTriggerBlobEvent].
func (dftbe *DataFactoryTriggerBlobEvent) DependOn() terra.Reference {
	return terra.ReferenceResource(dftbe)
}

// Dependencies returns the list of resources [DataFactoryTriggerBlobEvent] depends_on.
func (dftbe *DataFactoryTriggerBlobEvent) Dependencies() terra.Dependencies {
	return dftbe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryTriggerBlobEvent].
func (dftbe *DataFactoryTriggerBlobEvent) LifecycleManagement() *terra.Lifecycle {
	return dftbe.Lifecycle
}

// Attributes returns the attributes for [DataFactoryTriggerBlobEvent].
func (dftbe *DataFactoryTriggerBlobEvent) Attributes() dataFactoryTriggerBlobEventAttributes {
	return dataFactoryTriggerBlobEventAttributes{ref: terra.ReferenceResource(dftbe)}
}

// ImportState imports the given attribute values into [DataFactoryTriggerBlobEvent]'s state.
func (dftbe *DataFactoryTriggerBlobEvent) ImportState(av io.Reader) error {
	dftbe.state = &dataFactoryTriggerBlobEventState{}
	if err := json.NewDecoder(av).Decode(dftbe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dftbe.Type(), dftbe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryTriggerBlobEvent] has state.
func (dftbe *DataFactoryTriggerBlobEvent) State() (*dataFactoryTriggerBlobEventState, bool) {
	return dftbe.state, dftbe.state != nil
}

// StateMust returns the state for [DataFactoryTriggerBlobEvent]. Panics if the state is nil.
func (dftbe *DataFactoryTriggerBlobEvent) StateMust() *dataFactoryTriggerBlobEventState {
	if dftbe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dftbe.Type(), dftbe.LocalName()))
	}
	return dftbe.state
}

// DataFactoryTriggerBlobEventArgs contains the configurations for azurerm_data_factory_trigger_blob_event.
type DataFactoryTriggerBlobEventArgs struct {
	// Activated: bool, optional
	Activated terra.BoolValue `hcl:"activated,attr"`
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// BlobPathBeginsWith: string, optional
	BlobPathBeginsWith terra.StringValue `hcl:"blob_path_begins_with,attr"`
	// BlobPathEndsWith: string, optional
	BlobPathEndsWith terra.StringValue `hcl:"blob_path_ends_with,attr"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Events: set of string, required
	Events terra.SetValue[terra.StringValue] `hcl:"events,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IgnoreEmptyBlobs: bool, optional
	IgnoreEmptyBlobs terra.BoolValue `hcl:"ignore_empty_blobs,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// Pipeline: min=1
	Pipeline []datafactorytriggerblobevent.Pipeline `hcl:"pipeline,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *datafactorytriggerblobevent.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryTriggerBlobEventAttributes struct {
	ref terra.Reference
}

// Activated returns a reference to field activated of azurerm_data_factory_trigger_blob_event.
func (dftbe dataFactoryTriggerBlobEventAttributes) Activated() terra.BoolValue {
	return terra.ReferenceAsBool(dftbe.ref.Append("activated"))
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_trigger_blob_event.
func (dftbe dataFactoryTriggerBlobEventAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dftbe.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_trigger_blob_event.
func (dftbe dataFactoryTriggerBlobEventAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dftbe.ref.Append("annotations"))
}

// BlobPathBeginsWith returns a reference to field blob_path_begins_with of azurerm_data_factory_trigger_blob_event.
func (dftbe dataFactoryTriggerBlobEventAttributes) BlobPathBeginsWith() terra.StringValue {
	return terra.ReferenceAsString(dftbe.ref.Append("blob_path_begins_with"))
}

// BlobPathEndsWith returns a reference to field blob_path_ends_with of azurerm_data_factory_trigger_blob_event.
func (dftbe dataFactoryTriggerBlobEventAttributes) BlobPathEndsWith() terra.StringValue {
	return terra.ReferenceAsString(dftbe.ref.Append("blob_path_ends_with"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_trigger_blob_event.
func (dftbe dataFactoryTriggerBlobEventAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dftbe.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_trigger_blob_event.
func (dftbe dataFactoryTriggerBlobEventAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dftbe.ref.Append("description"))
}

// Events returns a reference to field events of azurerm_data_factory_trigger_blob_event.
func (dftbe dataFactoryTriggerBlobEventAttributes) Events() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dftbe.ref.Append("events"))
}

// Id returns a reference to field id of azurerm_data_factory_trigger_blob_event.
func (dftbe dataFactoryTriggerBlobEventAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dftbe.ref.Append("id"))
}

// IgnoreEmptyBlobs returns a reference to field ignore_empty_blobs of azurerm_data_factory_trigger_blob_event.
func (dftbe dataFactoryTriggerBlobEventAttributes) IgnoreEmptyBlobs() terra.BoolValue {
	return terra.ReferenceAsBool(dftbe.ref.Append("ignore_empty_blobs"))
}

// Name returns a reference to field name of azurerm_data_factory_trigger_blob_event.
func (dftbe dataFactoryTriggerBlobEventAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dftbe.ref.Append("name"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_data_factory_trigger_blob_event.
func (dftbe dataFactoryTriggerBlobEventAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(dftbe.ref.Append("storage_account_id"))
}

func (dftbe dataFactoryTriggerBlobEventAttributes) Pipeline() terra.SetValue[datafactorytriggerblobevent.PipelineAttributes] {
	return terra.ReferenceAsSet[datafactorytriggerblobevent.PipelineAttributes](dftbe.ref.Append("pipeline"))
}

func (dftbe dataFactoryTriggerBlobEventAttributes) Timeouts() datafactorytriggerblobevent.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorytriggerblobevent.TimeoutsAttributes](dftbe.ref.Append("timeouts"))
}

type dataFactoryTriggerBlobEventState struct {
	Activated            bool                                        `json:"activated"`
	AdditionalProperties map[string]string                           `json:"additional_properties"`
	Annotations          []string                                    `json:"annotations"`
	BlobPathBeginsWith   string                                      `json:"blob_path_begins_with"`
	BlobPathEndsWith     string                                      `json:"blob_path_ends_with"`
	DataFactoryId        string                                      `json:"data_factory_id"`
	Description          string                                      `json:"description"`
	Events               []string                                    `json:"events"`
	Id                   string                                      `json:"id"`
	IgnoreEmptyBlobs     bool                                        `json:"ignore_empty_blobs"`
	Name                 string                                      `json:"name"`
	StorageAccountId     string                                      `json:"storage_account_id"`
	Pipeline             []datafactorytriggerblobevent.PipelineState `json:"pipeline"`
	Timeouts             *datafactorytriggerblobevent.TimeoutsState  `json:"timeouts"`
}
