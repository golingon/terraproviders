// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedserviceazuresearch "github.com/golingon/terraproviders/azurerm/3.49.0/datafactorylinkedserviceazuresearch"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedServiceAzureSearch creates a new instance of [DataFactoryLinkedServiceAzureSearch].
func NewDataFactoryLinkedServiceAzureSearch(name string, args DataFactoryLinkedServiceAzureSearchArgs) *DataFactoryLinkedServiceAzureSearch {
	return &DataFactoryLinkedServiceAzureSearch{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceAzureSearch)(nil)

// DataFactoryLinkedServiceAzureSearch represents the Terraform resource azurerm_data_factory_linked_service_azure_search.
type DataFactoryLinkedServiceAzureSearch struct {
	Name      string
	Args      DataFactoryLinkedServiceAzureSearchArgs
	state     *dataFactoryLinkedServiceAzureSearchState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedServiceAzureSearch].
func (dflsas *DataFactoryLinkedServiceAzureSearch) Type() string {
	return "azurerm_data_factory_linked_service_azure_search"
}

// LocalName returns the local name for [DataFactoryLinkedServiceAzureSearch].
func (dflsas *DataFactoryLinkedServiceAzureSearch) LocalName() string {
	return dflsas.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedServiceAzureSearch].
func (dflsas *DataFactoryLinkedServiceAzureSearch) Configuration() interface{} {
	return dflsas.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedServiceAzureSearch].
func (dflsas *DataFactoryLinkedServiceAzureSearch) DependOn() terra.Reference {
	return terra.ReferenceResource(dflsas)
}

// Dependencies returns the list of resources [DataFactoryLinkedServiceAzureSearch] depends_on.
func (dflsas *DataFactoryLinkedServiceAzureSearch) Dependencies() terra.Dependencies {
	return dflsas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedServiceAzureSearch].
func (dflsas *DataFactoryLinkedServiceAzureSearch) LifecycleManagement() *terra.Lifecycle {
	return dflsas.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedServiceAzureSearch].
func (dflsas *DataFactoryLinkedServiceAzureSearch) Attributes() dataFactoryLinkedServiceAzureSearchAttributes {
	return dataFactoryLinkedServiceAzureSearchAttributes{ref: terra.ReferenceResource(dflsas)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedServiceAzureSearch]'s state.
func (dflsas *DataFactoryLinkedServiceAzureSearch) ImportState(av io.Reader) error {
	dflsas.state = &dataFactoryLinkedServiceAzureSearchState{}
	if err := json.NewDecoder(av).Decode(dflsas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflsas.Type(), dflsas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedServiceAzureSearch] has state.
func (dflsas *DataFactoryLinkedServiceAzureSearch) State() (*dataFactoryLinkedServiceAzureSearchState, bool) {
	return dflsas.state, dflsas.state != nil
}

// StateMust returns the state for [DataFactoryLinkedServiceAzureSearch]. Panics if the state is nil.
func (dflsas *DataFactoryLinkedServiceAzureSearch) StateMust() *dataFactoryLinkedServiceAzureSearchState {
	if dflsas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflsas.Type(), dflsas.LocalName()))
	}
	return dflsas.state
}

// DataFactoryLinkedServiceAzureSearchArgs contains the configurations for azurerm_data_factory_linked_service_azure_search.
type DataFactoryLinkedServiceAzureSearchArgs struct {
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
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// SearchServiceKey: string, required
	SearchServiceKey terra.StringValue `hcl:"search_service_key,attr" validate:"required"`
	// Url: string, required
	Url terra.StringValue `hcl:"url,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datafactorylinkedserviceazuresearch.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedServiceAzureSearchAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_azure_search.
func (dflsas dataFactoryLinkedServiceAzureSearchAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsas.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_azure_search.
func (dflsas dataFactoryLinkedServiceAzureSearchAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflsas.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_azure_search.
func (dflsas dataFactoryLinkedServiceAzureSearchAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflsas.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_azure_search.
func (dflsas dataFactoryLinkedServiceAzureSearchAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflsas.ref.Append("description"))
}

// EncryptedCredential returns a reference to field encrypted_credential of azurerm_data_factory_linked_service_azure_search.
func (dflsas dataFactoryLinkedServiceAzureSearchAttributes) EncryptedCredential() terra.StringValue {
	return terra.ReferenceAsString(dflsas.ref.Append("encrypted_credential"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_azure_search.
func (dflsas dataFactoryLinkedServiceAzureSearchAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflsas.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_azure_search.
func (dflsas dataFactoryLinkedServiceAzureSearchAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(dflsas.ref.Append("integration_runtime_name"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_azure_search.
func (dflsas dataFactoryLinkedServiceAzureSearchAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflsas.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_azure_search.
func (dflsas dataFactoryLinkedServiceAzureSearchAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsas.ref.Append("parameters"))
}

// SearchServiceKey returns a reference to field search_service_key of azurerm_data_factory_linked_service_azure_search.
func (dflsas dataFactoryLinkedServiceAzureSearchAttributes) SearchServiceKey() terra.StringValue {
	return terra.ReferenceAsString(dflsas.ref.Append("search_service_key"))
}

// Url returns a reference to field url of azurerm_data_factory_linked_service_azure_search.
func (dflsas dataFactoryLinkedServiceAzureSearchAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(dflsas.ref.Append("url"))
}

func (dflsas dataFactoryLinkedServiceAzureSearchAttributes) Timeouts() datafactorylinkedserviceazuresearch.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedserviceazuresearch.TimeoutsAttributes](dflsas.ref.Append("timeouts"))
}

type dataFactoryLinkedServiceAzureSearchState struct {
	AdditionalProperties   map[string]string                                  `json:"additional_properties"`
	Annotations            []string                                           `json:"annotations"`
	DataFactoryId          string                                             `json:"data_factory_id"`
	Description            string                                             `json:"description"`
	EncryptedCredential    string                                             `json:"encrypted_credential"`
	Id                     string                                             `json:"id"`
	IntegrationRuntimeName string                                             `json:"integration_runtime_name"`
	Name                   string                                             `json:"name"`
	Parameters             map[string]string                                  `json:"parameters"`
	SearchServiceKey       string                                             `json:"search_service_key"`
	Url                    string                                             `json:"url"`
	Timeouts               *datafactorylinkedserviceazuresearch.TimeoutsState `json:"timeouts"`
}
