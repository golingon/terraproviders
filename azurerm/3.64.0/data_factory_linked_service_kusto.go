// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedservicekusto "github.com/golingon/terraproviders/azurerm/3.64.0/datafactorylinkedservicekusto"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedServiceKusto creates a new instance of [DataFactoryLinkedServiceKusto].
func NewDataFactoryLinkedServiceKusto(name string, args DataFactoryLinkedServiceKustoArgs) *DataFactoryLinkedServiceKusto {
	return &DataFactoryLinkedServiceKusto{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceKusto)(nil)

// DataFactoryLinkedServiceKusto represents the Terraform resource azurerm_data_factory_linked_service_kusto.
type DataFactoryLinkedServiceKusto struct {
	Name      string
	Args      DataFactoryLinkedServiceKustoArgs
	state     *dataFactoryLinkedServiceKustoState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedServiceKusto].
func (dflsk *DataFactoryLinkedServiceKusto) Type() string {
	return "azurerm_data_factory_linked_service_kusto"
}

// LocalName returns the local name for [DataFactoryLinkedServiceKusto].
func (dflsk *DataFactoryLinkedServiceKusto) LocalName() string {
	return dflsk.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedServiceKusto].
func (dflsk *DataFactoryLinkedServiceKusto) Configuration() interface{} {
	return dflsk.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedServiceKusto].
func (dflsk *DataFactoryLinkedServiceKusto) DependOn() terra.Reference {
	return terra.ReferenceResource(dflsk)
}

// Dependencies returns the list of resources [DataFactoryLinkedServiceKusto] depends_on.
func (dflsk *DataFactoryLinkedServiceKusto) Dependencies() terra.Dependencies {
	return dflsk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedServiceKusto].
func (dflsk *DataFactoryLinkedServiceKusto) LifecycleManagement() *terra.Lifecycle {
	return dflsk.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedServiceKusto].
func (dflsk *DataFactoryLinkedServiceKusto) Attributes() dataFactoryLinkedServiceKustoAttributes {
	return dataFactoryLinkedServiceKustoAttributes{ref: terra.ReferenceResource(dflsk)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedServiceKusto]'s state.
func (dflsk *DataFactoryLinkedServiceKusto) ImportState(av io.Reader) error {
	dflsk.state = &dataFactoryLinkedServiceKustoState{}
	if err := json.NewDecoder(av).Decode(dflsk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflsk.Type(), dflsk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedServiceKusto] has state.
func (dflsk *DataFactoryLinkedServiceKusto) State() (*dataFactoryLinkedServiceKustoState, bool) {
	return dflsk.state, dflsk.state != nil
}

// StateMust returns the state for [DataFactoryLinkedServiceKusto]. Panics if the state is nil.
func (dflsk *DataFactoryLinkedServiceKusto) StateMust() *dataFactoryLinkedServiceKustoState {
	if dflsk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflsk.Type(), dflsk.LocalName()))
	}
	return dflsk.state
}

// DataFactoryLinkedServiceKustoArgs contains the configurations for azurerm_data_factory_linked_service_kusto.
type DataFactoryLinkedServiceKustoArgs struct {
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
	// KustoDatabaseName: string, required
	KustoDatabaseName terra.StringValue `hcl:"kusto_database_name,attr" validate:"required"`
	// KustoEndpoint: string, required
	KustoEndpoint terra.StringValue `hcl:"kusto_endpoint,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// ServicePrincipalId: string, optional
	ServicePrincipalId terra.StringValue `hcl:"service_principal_id,attr"`
	// ServicePrincipalKey: string, optional
	ServicePrincipalKey terra.StringValue `hcl:"service_principal_key,attr"`
	// Tenant: string, optional
	Tenant terra.StringValue `hcl:"tenant,attr"`
	// UseManagedIdentity: bool, optional
	UseManagedIdentity terra.BoolValue `hcl:"use_managed_identity,attr"`
	// Timeouts: optional
	Timeouts *datafactorylinkedservicekusto.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedServiceKustoAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_kusto.
func (dflsk dataFactoryLinkedServiceKustoAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsk.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_kusto.
func (dflsk dataFactoryLinkedServiceKustoAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflsk.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_kusto.
func (dflsk dataFactoryLinkedServiceKustoAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflsk.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_kusto.
func (dflsk dataFactoryLinkedServiceKustoAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflsk.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_kusto.
func (dflsk dataFactoryLinkedServiceKustoAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflsk.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_kusto.
func (dflsk dataFactoryLinkedServiceKustoAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(dflsk.ref.Append("integration_runtime_name"))
}

// KustoDatabaseName returns a reference to field kusto_database_name of azurerm_data_factory_linked_service_kusto.
func (dflsk dataFactoryLinkedServiceKustoAttributes) KustoDatabaseName() terra.StringValue {
	return terra.ReferenceAsString(dflsk.ref.Append("kusto_database_name"))
}

// KustoEndpoint returns a reference to field kusto_endpoint of azurerm_data_factory_linked_service_kusto.
func (dflsk dataFactoryLinkedServiceKustoAttributes) KustoEndpoint() terra.StringValue {
	return terra.ReferenceAsString(dflsk.ref.Append("kusto_endpoint"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_kusto.
func (dflsk dataFactoryLinkedServiceKustoAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflsk.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_kusto.
func (dflsk dataFactoryLinkedServiceKustoAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflsk.ref.Append("parameters"))
}

// ServicePrincipalId returns a reference to field service_principal_id of azurerm_data_factory_linked_service_kusto.
func (dflsk dataFactoryLinkedServiceKustoAttributes) ServicePrincipalId() terra.StringValue {
	return terra.ReferenceAsString(dflsk.ref.Append("service_principal_id"))
}

// ServicePrincipalKey returns a reference to field service_principal_key of azurerm_data_factory_linked_service_kusto.
func (dflsk dataFactoryLinkedServiceKustoAttributes) ServicePrincipalKey() terra.StringValue {
	return terra.ReferenceAsString(dflsk.ref.Append("service_principal_key"))
}

// Tenant returns a reference to field tenant of azurerm_data_factory_linked_service_kusto.
func (dflsk dataFactoryLinkedServiceKustoAttributes) Tenant() terra.StringValue {
	return terra.ReferenceAsString(dflsk.ref.Append("tenant"))
}

// UseManagedIdentity returns a reference to field use_managed_identity of azurerm_data_factory_linked_service_kusto.
func (dflsk dataFactoryLinkedServiceKustoAttributes) UseManagedIdentity() terra.BoolValue {
	return terra.ReferenceAsBool(dflsk.ref.Append("use_managed_identity"))
}

func (dflsk dataFactoryLinkedServiceKustoAttributes) Timeouts() datafactorylinkedservicekusto.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedservicekusto.TimeoutsAttributes](dflsk.ref.Append("timeouts"))
}

type dataFactoryLinkedServiceKustoState struct {
	AdditionalProperties   map[string]string                            `json:"additional_properties"`
	Annotations            []string                                     `json:"annotations"`
	DataFactoryId          string                                       `json:"data_factory_id"`
	Description            string                                       `json:"description"`
	Id                     string                                       `json:"id"`
	IntegrationRuntimeName string                                       `json:"integration_runtime_name"`
	KustoDatabaseName      string                                       `json:"kusto_database_name"`
	KustoEndpoint          string                                       `json:"kusto_endpoint"`
	Name                   string                                       `json:"name"`
	Parameters             map[string]string                            `json:"parameters"`
	ServicePrincipalId     string                                       `json:"service_principal_id"`
	ServicePrincipalKey    string                                       `json:"service_principal_key"`
	Tenant                 string                                       `json:"tenant"`
	UseManagedIdentity     bool                                         `json:"use_managed_identity"`
	Timeouts               *datafactorylinkedservicekusto.TimeoutsState `json:"timeouts"`
}
