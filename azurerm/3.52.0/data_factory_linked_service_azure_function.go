// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedserviceazurefunction "github.com/golingon/terraproviders/azurerm/3.52.0/datafactorylinkedserviceazurefunction"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedServiceAzureFunction creates a new instance of [DataFactoryLinkedServiceAzureFunction].
func NewDataFactoryLinkedServiceAzureFunction(name string, args DataFactoryLinkedServiceAzureFunctionArgs) *DataFactoryLinkedServiceAzureFunction {
	return &DataFactoryLinkedServiceAzureFunction{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceAzureFunction)(nil)

// DataFactoryLinkedServiceAzureFunction represents the Terraform resource azurerm_data_factory_linked_service_azure_function.
type DataFactoryLinkedServiceAzureFunction struct {
	Name      string
	Args      DataFactoryLinkedServiceAzureFunctionArgs
	state     *dataFactoryLinkedServiceAzureFunctionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedServiceAzureFunction].
func (dflsaf *DataFactoryLinkedServiceAzureFunction) Type() string {
	return "azurerm_data_factory_linked_service_azure_function"
}

// LocalName returns the local name for [DataFactoryLinkedServiceAzureFunction].
func (dflsaf *DataFactoryLinkedServiceAzureFunction) LocalName() string {
	return dflsaf.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedServiceAzureFunction].
func (dflsaf *DataFactoryLinkedServiceAzureFunction) Configuration() interface{} {
	return dflsaf.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedServiceAzureFunction].
func (dflsaf *DataFactoryLinkedServiceAzureFunction) DependOn() terra.Reference {
	return terra.ReferenceResource(dflsaf)
}

// Dependencies returns the list of resources [DataFactoryLinkedServiceAzureFunction] depends_on.
func (dflsaf *DataFactoryLinkedServiceAzureFunction) Dependencies() terra.Dependencies {
	return dflsaf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedServiceAzureFunction].
func (dflsaf *DataFactoryLinkedServiceAzureFunction) LifecycleManagement() *terra.Lifecycle {
	return dflsaf.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedServiceAzureFunction].
func (dflsaf *DataFactoryLinkedServiceAzureFunction) Attributes() dataFactoryLinkedServiceAzureFunctionAttributes {
	return dataFactoryLinkedServiceAzureFunctionAttributes{ref: terra.ReferenceResource(dflsaf)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedServiceAzureFunction]'s state.
func (dflsaf *DataFactoryLinkedServiceAzureFunction) ImportState(av io.Reader) error {
	dflsaf.state = &dataFactoryLinkedServiceAzureFunctionState{}
	if err := json.NewDecoder(av).Decode(dflsaf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflsaf.Type(), dflsaf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedServiceAzureFunction] has state.
func (dflsaf *DataFactoryLinkedServiceAzureFunction) State() (*dataFactoryLinkedServiceAzureFunctionState, bool) {
	return dflsaf.state, dflsaf.state != nil
}

// StateMust returns the state for [DataFactoryLinkedServiceAzureFunction]. Panics if the state is nil.
func (dflsaf *DataFactoryLinkedServiceAzureFunction) StateMust() *dataFactoryLinkedServiceAzureFunctionState {
	if dflsaf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflsaf.Type(), dflsaf.LocalName()))
	}
	return dflsaf.state
}

// DataFactoryLinkedServiceAzureFunctionArgs contains the configurations for azurerm_data_factory_linked_service_azure_function.
type DataFactoryLinkedServiceAzureFunctionArgs struct {
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
	// IntegrationRuntimeName: string, optional
	IntegrationRuntimeName terra.StringValue `hcl:"integration_runtime_name,attr"`
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// Url: string, required
	Url terra.StringValue `hcl:"url,attr" validate:"required"`
	// KeyVaultKey: optional
	KeyVaultKey *datafactorylinkedserviceazurefunction.KeyVaultKey `hcl:"key_vault_key,block"`
	// Timeouts: optional
	Timeouts *datafactorylinkedserviceazurefunction.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedServiceAzureFunctionAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_azure_function.
func (dflsaf dataFactoryLinkedServiceAzureFunctionAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsaf.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_azure_function.
func (dflsaf dataFactoryLinkedServiceAzureFunctionAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflsaf.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_azure_function.
func (dflsaf dataFactoryLinkedServiceAzureFunctionAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflsaf.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_azure_function.
func (dflsaf dataFactoryLinkedServiceAzureFunctionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflsaf.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_azure_function.
func (dflsaf dataFactoryLinkedServiceAzureFunctionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflsaf.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_azure_function.
func (dflsaf dataFactoryLinkedServiceAzureFunctionAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(dflsaf.ref.Append("integration_runtime_name"))
}

// Key returns a reference to field key of azurerm_data_factory_linked_service_azure_function.
func (dflsaf dataFactoryLinkedServiceAzureFunctionAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(dflsaf.ref.Append("key"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_azure_function.
func (dflsaf dataFactoryLinkedServiceAzureFunctionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflsaf.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_azure_function.
func (dflsaf dataFactoryLinkedServiceAzureFunctionAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsaf.ref.Append("parameters"))
}

// Url returns a reference to field url of azurerm_data_factory_linked_service_azure_function.
func (dflsaf dataFactoryLinkedServiceAzureFunctionAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(dflsaf.ref.Append("url"))
}

func (dflsaf dataFactoryLinkedServiceAzureFunctionAttributes) KeyVaultKey() terra.ListValue[datafactorylinkedserviceazurefunction.KeyVaultKeyAttributes] {
	return terra.ReferenceAsList[datafactorylinkedserviceazurefunction.KeyVaultKeyAttributes](dflsaf.ref.Append("key_vault_key"))
}

func (dflsaf dataFactoryLinkedServiceAzureFunctionAttributes) Timeouts() datafactorylinkedserviceazurefunction.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedserviceazurefunction.TimeoutsAttributes](dflsaf.ref.Append("timeouts"))
}

type dataFactoryLinkedServiceAzureFunctionState struct {
	AdditionalProperties   map[string]string                                        `json:"additional_properties"`
	Annotations            []string                                                 `json:"annotations"`
	DataFactoryId          string                                                   `json:"data_factory_id"`
	Description            string                                                   `json:"description"`
	Id                     string                                                   `json:"id"`
	IntegrationRuntimeName string                                                   `json:"integration_runtime_name"`
	Key                    string                                                   `json:"key"`
	Name                   string                                                   `json:"name"`
	Parameters             map[string]string                                        `json:"parameters"`
	Url                    string                                                   `json:"url"`
	KeyVaultKey            []datafactorylinkedserviceazurefunction.KeyVaultKeyState `json:"key_vault_key"`
	Timeouts               *datafactorylinkedserviceazurefunction.TimeoutsState     `json:"timeouts"`
}
