// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	kustocosmosdbdataconnection "github.com/golingon/terraproviders/azurerm/3.98.0/kustocosmosdbdataconnection"
	"io"
)

// NewKustoCosmosdbDataConnection creates a new instance of [KustoCosmosdbDataConnection].
func NewKustoCosmosdbDataConnection(name string, args KustoCosmosdbDataConnectionArgs) *KustoCosmosdbDataConnection {
	return &KustoCosmosdbDataConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KustoCosmosdbDataConnection)(nil)

// KustoCosmosdbDataConnection represents the Terraform resource azurerm_kusto_cosmosdb_data_connection.
type KustoCosmosdbDataConnection struct {
	Name      string
	Args      KustoCosmosdbDataConnectionArgs
	state     *kustoCosmosdbDataConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KustoCosmosdbDataConnection].
func (kcdc *KustoCosmosdbDataConnection) Type() string {
	return "azurerm_kusto_cosmosdb_data_connection"
}

// LocalName returns the local name for [KustoCosmosdbDataConnection].
func (kcdc *KustoCosmosdbDataConnection) LocalName() string {
	return kcdc.Name
}

// Configuration returns the configuration (args) for [KustoCosmosdbDataConnection].
func (kcdc *KustoCosmosdbDataConnection) Configuration() interface{} {
	return kcdc.Args
}

// DependOn is used for other resources to depend on [KustoCosmosdbDataConnection].
func (kcdc *KustoCosmosdbDataConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(kcdc)
}

// Dependencies returns the list of resources [KustoCosmosdbDataConnection] depends_on.
func (kcdc *KustoCosmosdbDataConnection) Dependencies() terra.Dependencies {
	return kcdc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KustoCosmosdbDataConnection].
func (kcdc *KustoCosmosdbDataConnection) LifecycleManagement() *terra.Lifecycle {
	return kcdc.Lifecycle
}

// Attributes returns the attributes for [KustoCosmosdbDataConnection].
func (kcdc *KustoCosmosdbDataConnection) Attributes() kustoCosmosdbDataConnectionAttributes {
	return kustoCosmosdbDataConnectionAttributes{ref: terra.ReferenceResource(kcdc)}
}

// ImportState imports the given attribute values into [KustoCosmosdbDataConnection]'s state.
func (kcdc *KustoCosmosdbDataConnection) ImportState(av io.Reader) error {
	kcdc.state = &kustoCosmosdbDataConnectionState{}
	if err := json.NewDecoder(av).Decode(kcdc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kcdc.Type(), kcdc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KustoCosmosdbDataConnection] has state.
func (kcdc *KustoCosmosdbDataConnection) State() (*kustoCosmosdbDataConnectionState, bool) {
	return kcdc.state, kcdc.state != nil
}

// StateMust returns the state for [KustoCosmosdbDataConnection]. Panics if the state is nil.
func (kcdc *KustoCosmosdbDataConnection) StateMust() *kustoCosmosdbDataConnectionState {
	if kcdc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kcdc.Type(), kcdc.LocalName()))
	}
	return kcdc.state
}

// KustoCosmosdbDataConnectionArgs contains the configurations for azurerm_kusto_cosmosdb_data_connection.
type KustoCosmosdbDataConnectionArgs struct {
	// CosmosdbContainerId: string, required
	CosmosdbContainerId terra.StringValue `hcl:"cosmosdb_container_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KustoDatabaseId: string, required
	KustoDatabaseId terra.StringValue `hcl:"kusto_database_id,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ManagedIdentityId: string, required
	ManagedIdentityId terra.StringValue `hcl:"managed_identity_id,attr" validate:"required"`
	// MappingRuleName: string, optional
	MappingRuleName terra.StringValue `hcl:"mapping_rule_name,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RetrievalStartDate: string, optional
	RetrievalStartDate terra.StringValue `hcl:"retrieval_start_date,attr"`
	// TableName: string, required
	TableName terra.StringValue `hcl:"table_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *kustocosmosdbdataconnection.Timeouts `hcl:"timeouts,block"`
}
type kustoCosmosdbDataConnectionAttributes struct {
	ref terra.Reference
}

// CosmosdbContainerId returns a reference to field cosmosdb_container_id of azurerm_kusto_cosmosdb_data_connection.
func (kcdc kustoCosmosdbDataConnectionAttributes) CosmosdbContainerId() terra.StringValue {
	return terra.ReferenceAsString(kcdc.ref.Append("cosmosdb_container_id"))
}

// Id returns a reference to field id of azurerm_kusto_cosmosdb_data_connection.
func (kcdc kustoCosmosdbDataConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kcdc.ref.Append("id"))
}

// KustoDatabaseId returns a reference to field kusto_database_id of azurerm_kusto_cosmosdb_data_connection.
func (kcdc kustoCosmosdbDataConnectionAttributes) KustoDatabaseId() terra.StringValue {
	return terra.ReferenceAsString(kcdc.ref.Append("kusto_database_id"))
}

// Location returns a reference to field location of azurerm_kusto_cosmosdb_data_connection.
func (kcdc kustoCosmosdbDataConnectionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(kcdc.ref.Append("location"))
}

// ManagedIdentityId returns a reference to field managed_identity_id of azurerm_kusto_cosmosdb_data_connection.
func (kcdc kustoCosmosdbDataConnectionAttributes) ManagedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(kcdc.ref.Append("managed_identity_id"))
}

// MappingRuleName returns a reference to field mapping_rule_name of azurerm_kusto_cosmosdb_data_connection.
func (kcdc kustoCosmosdbDataConnectionAttributes) MappingRuleName() terra.StringValue {
	return terra.ReferenceAsString(kcdc.ref.Append("mapping_rule_name"))
}

// Name returns a reference to field name of azurerm_kusto_cosmosdb_data_connection.
func (kcdc kustoCosmosdbDataConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kcdc.ref.Append("name"))
}

// RetrievalStartDate returns a reference to field retrieval_start_date of azurerm_kusto_cosmosdb_data_connection.
func (kcdc kustoCosmosdbDataConnectionAttributes) RetrievalStartDate() terra.StringValue {
	return terra.ReferenceAsString(kcdc.ref.Append("retrieval_start_date"))
}

// TableName returns a reference to field table_name of azurerm_kusto_cosmosdb_data_connection.
func (kcdc kustoCosmosdbDataConnectionAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(kcdc.ref.Append("table_name"))
}

func (kcdc kustoCosmosdbDataConnectionAttributes) Timeouts() kustocosmosdbdataconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kustocosmosdbdataconnection.TimeoutsAttributes](kcdc.ref.Append("timeouts"))
}

type kustoCosmosdbDataConnectionState struct {
	CosmosdbContainerId string                                     `json:"cosmosdb_container_id"`
	Id                  string                                     `json:"id"`
	KustoDatabaseId     string                                     `json:"kusto_database_id"`
	Location            string                                     `json:"location"`
	ManagedIdentityId   string                                     `json:"managed_identity_id"`
	MappingRuleName     string                                     `json:"mapping_rule_name"`
	Name                string                                     `json:"name"`
	RetrievalStartDate  string                                     `json:"retrieval_start_date"`
	TableName           string                                     `json:"table_name"`
	Timeouts            *kustocosmosdbdataconnection.TimeoutsState `json:"timeouts"`
}
