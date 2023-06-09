// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedcustomservice "github.com/golingon/terraproviders/azurerm/3.52.0/datafactorylinkedcustomservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedCustomService creates a new instance of [DataFactoryLinkedCustomService].
func NewDataFactoryLinkedCustomService(name string, args DataFactoryLinkedCustomServiceArgs) *DataFactoryLinkedCustomService {
	return &DataFactoryLinkedCustomService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedCustomService)(nil)

// DataFactoryLinkedCustomService represents the Terraform resource azurerm_data_factory_linked_custom_service.
type DataFactoryLinkedCustomService struct {
	Name      string
	Args      DataFactoryLinkedCustomServiceArgs
	state     *dataFactoryLinkedCustomServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedCustomService].
func (dflcs *DataFactoryLinkedCustomService) Type() string {
	return "azurerm_data_factory_linked_custom_service"
}

// LocalName returns the local name for [DataFactoryLinkedCustomService].
func (dflcs *DataFactoryLinkedCustomService) LocalName() string {
	return dflcs.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedCustomService].
func (dflcs *DataFactoryLinkedCustomService) Configuration() interface{} {
	return dflcs.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedCustomService].
func (dflcs *DataFactoryLinkedCustomService) DependOn() terra.Reference {
	return terra.ReferenceResource(dflcs)
}

// Dependencies returns the list of resources [DataFactoryLinkedCustomService] depends_on.
func (dflcs *DataFactoryLinkedCustomService) Dependencies() terra.Dependencies {
	return dflcs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedCustomService].
func (dflcs *DataFactoryLinkedCustomService) LifecycleManagement() *terra.Lifecycle {
	return dflcs.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedCustomService].
func (dflcs *DataFactoryLinkedCustomService) Attributes() dataFactoryLinkedCustomServiceAttributes {
	return dataFactoryLinkedCustomServiceAttributes{ref: terra.ReferenceResource(dflcs)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedCustomService]'s state.
func (dflcs *DataFactoryLinkedCustomService) ImportState(av io.Reader) error {
	dflcs.state = &dataFactoryLinkedCustomServiceState{}
	if err := json.NewDecoder(av).Decode(dflcs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflcs.Type(), dflcs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedCustomService] has state.
func (dflcs *DataFactoryLinkedCustomService) State() (*dataFactoryLinkedCustomServiceState, bool) {
	return dflcs.state, dflcs.state != nil
}

// StateMust returns the state for [DataFactoryLinkedCustomService]. Panics if the state is nil.
func (dflcs *DataFactoryLinkedCustomService) StateMust() *dataFactoryLinkedCustomServiceState {
	if dflcs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflcs.Type(), dflcs.LocalName()))
	}
	return dflcs.state
}

// DataFactoryLinkedCustomServiceArgs contains the configurations for azurerm_data_factory_linked_custom_service.
type DataFactoryLinkedCustomServiceArgs struct {
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// TypePropertiesJson: string, required
	TypePropertiesJson terra.StringValue `hcl:"type_properties_json,attr" validate:"required"`
	// IntegrationRuntime: optional
	IntegrationRuntime *datafactorylinkedcustomservice.IntegrationRuntime `hcl:"integration_runtime,block"`
	// Timeouts: optional
	Timeouts *datafactorylinkedcustomservice.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedCustomServiceAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_custom_service.
func (dflcs dataFactoryLinkedCustomServiceAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflcs.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_custom_service.
func (dflcs dataFactoryLinkedCustomServiceAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflcs.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_custom_service.
func (dflcs dataFactoryLinkedCustomServiceAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflcs.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_custom_service.
func (dflcs dataFactoryLinkedCustomServiceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflcs.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_custom_service.
func (dflcs dataFactoryLinkedCustomServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflcs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_custom_service.
func (dflcs dataFactoryLinkedCustomServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflcs.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_custom_service.
func (dflcs dataFactoryLinkedCustomServiceAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflcs.ref.Append("parameters"))
}

// Type returns a reference to field type of azurerm_data_factory_linked_custom_service.
func (dflcs dataFactoryLinkedCustomServiceAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(dflcs.ref.Append("type"))
}

// TypePropertiesJson returns a reference to field type_properties_json of azurerm_data_factory_linked_custom_service.
func (dflcs dataFactoryLinkedCustomServiceAttributes) TypePropertiesJson() terra.StringValue {
	return terra.ReferenceAsString(dflcs.ref.Append("type_properties_json"))
}

func (dflcs dataFactoryLinkedCustomServiceAttributes) IntegrationRuntime() terra.ListValue[datafactorylinkedcustomservice.IntegrationRuntimeAttributes] {
	return terra.ReferenceAsList[datafactorylinkedcustomservice.IntegrationRuntimeAttributes](dflcs.ref.Append("integration_runtime"))
}

func (dflcs dataFactoryLinkedCustomServiceAttributes) Timeouts() datafactorylinkedcustomservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedcustomservice.TimeoutsAttributes](dflcs.ref.Append("timeouts"))
}

type dataFactoryLinkedCustomServiceState struct {
	AdditionalProperties map[string]string                                        `json:"additional_properties"`
	Annotations          []string                                                 `json:"annotations"`
	DataFactoryId        string                                                   `json:"data_factory_id"`
	Description          string                                                   `json:"description"`
	Id                   string                                                   `json:"id"`
	Name                 string                                                   `json:"name"`
	Parameters           map[string]string                                        `json:"parameters"`
	Type                 string                                                   `json:"type"`
	TypePropertiesJson   string                                                   `json:"type_properties_json"`
	IntegrationRuntime   []datafactorylinkedcustomservice.IntegrationRuntimeState `json:"integration_runtime"`
	Timeouts             *datafactorylinkedcustomservice.TimeoutsState            `json:"timeouts"`
}
