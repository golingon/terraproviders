// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	kustoeventgriddataconnection "github.com/golingon/terraproviders/azurerm/3.66.0/kustoeventgriddataconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKustoEventgridDataConnection creates a new instance of [KustoEventgridDataConnection].
func NewKustoEventgridDataConnection(name string, args KustoEventgridDataConnectionArgs) *KustoEventgridDataConnection {
	return &KustoEventgridDataConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KustoEventgridDataConnection)(nil)

// KustoEventgridDataConnection represents the Terraform resource azurerm_kusto_eventgrid_data_connection.
type KustoEventgridDataConnection struct {
	Name      string
	Args      KustoEventgridDataConnectionArgs
	state     *kustoEventgridDataConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KustoEventgridDataConnection].
func (kedc *KustoEventgridDataConnection) Type() string {
	return "azurerm_kusto_eventgrid_data_connection"
}

// LocalName returns the local name for [KustoEventgridDataConnection].
func (kedc *KustoEventgridDataConnection) LocalName() string {
	return kedc.Name
}

// Configuration returns the configuration (args) for [KustoEventgridDataConnection].
func (kedc *KustoEventgridDataConnection) Configuration() interface{} {
	return kedc.Args
}

// DependOn is used for other resources to depend on [KustoEventgridDataConnection].
func (kedc *KustoEventgridDataConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(kedc)
}

// Dependencies returns the list of resources [KustoEventgridDataConnection] depends_on.
func (kedc *KustoEventgridDataConnection) Dependencies() terra.Dependencies {
	return kedc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KustoEventgridDataConnection].
func (kedc *KustoEventgridDataConnection) LifecycleManagement() *terra.Lifecycle {
	return kedc.Lifecycle
}

// Attributes returns the attributes for [KustoEventgridDataConnection].
func (kedc *KustoEventgridDataConnection) Attributes() kustoEventgridDataConnectionAttributes {
	return kustoEventgridDataConnectionAttributes{ref: terra.ReferenceResource(kedc)}
}

// ImportState imports the given attribute values into [KustoEventgridDataConnection]'s state.
func (kedc *KustoEventgridDataConnection) ImportState(av io.Reader) error {
	kedc.state = &kustoEventgridDataConnectionState{}
	if err := json.NewDecoder(av).Decode(kedc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kedc.Type(), kedc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KustoEventgridDataConnection] has state.
func (kedc *KustoEventgridDataConnection) State() (*kustoEventgridDataConnectionState, bool) {
	return kedc.state, kedc.state != nil
}

// StateMust returns the state for [KustoEventgridDataConnection]. Panics if the state is nil.
func (kedc *KustoEventgridDataConnection) StateMust() *kustoEventgridDataConnectionState {
	if kedc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kedc.Type(), kedc.LocalName()))
	}
	return kedc.state
}

// KustoEventgridDataConnectionArgs contains the configurations for azurerm_kusto_eventgrid_data_connection.
type KustoEventgridDataConnectionArgs struct {
	// BlobStorageEventType: string, optional
	BlobStorageEventType terra.StringValue `hcl:"blob_storage_event_type,attr"`
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// DataFormat: string, optional
	DataFormat terra.StringValue `hcl:"data_format,attr"`
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// DatabaseRoutingType: string, optional
	DatabaseRoutingType terra.StringValue `hcl:"database_routing_type,attr"`
	// EventgridResourceId: string, optional
	EventgridResourceId terra.StringValue `hcl:"eventgrid_resource_id,attr"`
	// EventhubConsumerGroupName: string, required
	EventhubConsumerGroupName terra.StringValue `hcl:"eventhub_consumer_group_name,attr" validate:"required"`
	// EventhubId: string, required
	EventhubId terra.StringValue `hcl:"eventhub_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ManagedIdentityResourceId: string, optional
	ManagedIdentityResourceId terra.StringValue `hcl:"managed_identity_resource_id,attr"`
	// MappingRuleName: string, optional
	MappingRuleName terra.StringValue `hcl:"mapping_rule_name,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkipFirstRecord: bool, optional
	SkipFirstRecord terra.BoolValue `hcl:"skip_first_record,attr"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// TableName: string, optional
	TableName terra.StringValue `hcl:"table_name,attr"`
	// Timeouts: optional
	Timeouts *kustoeventgriddataconnection.Timeouts `hcl:"timeouts,block"`
}
type kustoEventgridDataConnectionAttributes struct {
	ref terra.Reference
}

// BlobStorageEventType returns a reference to field blob_storage_event_type of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) BlobStorageEventType() terra.StringValue {
	return terra.ReferenceAsString(kedc.ref.Append("blob_storage_event_type"))
}

// ClusterName returns a reference to field cluster_name of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(kedc.ref.Append("cluster_name"))
}

// DataFormat returns a reference to field data_format of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) DataFormat() terra.StringValue {
	return terra.ReferenceAsString(kedc.ref.Append("data_format"))
}

// DatabaseName returns a reference to field database_name of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(kedc.ref.Append("database_name"))
}

// DatabaseRoutingType returns a reference to field database_routing_type of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) DatabaseRoutingType() terra.StringValue {
	return terra.ReferenceAsString(kedc.ref.Append("database_routing_type"))
}

// EventgridResourceId returns a reference to field eventgrid_resource_id of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) EventgridResourceId() terra.StringValue {
	return terra.ReferenceAsString(kedc.ref.Append("eventgrid_resource_id"))
}

// EventhubConsumerGroupName returns a reference to field eventhub_consumer_group_name of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) EventhubConsumerGroupName() terra.StringValue {
	return terra.ReferenceAsString(kedc.ref.Append("eventhub_consumer_group_name"))
}

// EventhubId returns a reference to field eventhub_id of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) EventhubId() terra.StringValue {
	return terra.ReferenceAsString(kedc.ref.Append("eventhub_id"))
}

// Id returns a reference to field id of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kedc.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(kedc.ref.Append("location"))
}

// ManagedIdentityResourceId returns a reference to field managed_identity_resource_id of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) ManagedIdentityResourceId() terra.StringValue {
	return terra.ReferenceAsString(kedc.ref.Append("managed_identity_resource_id"))
}

// MappingRuleName returns a reference to field mapping_rule_name of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) MappingRuleName() terra.StringValue {
	return terra.ReferenceAsString(kedc.ref.Append("mapping_rule_name"))
}

// Name returns a reference to field name of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kedc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(kedc.ref.Append("resource_group_name"))
}

// SkipFirstRecord returns a reference to field skip_first_record of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) SkipFirstRecord() terra.BoolValue {
	return terra.ReferenceAsBool(kedc.ref.Append("skip_first_record"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(kedc.ref.Append("storage_account_id"))
}

// TableName returns a reference to field table_name of azurerm_kusto_eventgrid_data_connection.
func (kedc kustoEventgridDataConnectionAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(kedc.ref.Append("table_name"))
}

func (kedc kustoEventgridDataConnectionAttributes) Timeouts() kustoeventgriddataconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kustoeventgriddataconnection.TimeoutsAttributes](kedc.ref.Append("timeouts"))
}

type kustoEventgridDataConnectionState struct {
	BlobStorageEventType      string                                      `json:"blob_storage_event_type"`
	ClusterName               string                                      `json:"cluster_name"`
	DataFormat                string                                      `json:"data_format"`
	DatabaseName              string                                      `json:"database_name"`
	DatabaseRoutingType       string                                      `json:"database_routing_type"`
	EventgridResourceId       string                                      `json:"eventgrid_resource_id"`
	EventhubConsumerGroupName string                                      `json:"eventhub_consumer_group_name"`
	EventhubId                string                                      `json:"eventhub_id"`
	Id                        string                                      `json:"id"`
	Location                  string                                      `json:"location"`
	ManagedIdentityResourceId string                                      `json:"managed_identity_resource_id"`
	MappingRuleName           string                                      `json:"mapping_rule_name"`
	Name                      string                                      `json:"name"`
	ResourceGroupName         string                                      `json:"resource_group_name"`
	SkipFirstRecord           bool                                        `json:"skip_first_record"`
	StorageAccountId          string                                      `json:"storage_account_id"`
	TableName                 string                                      `json:"table_name"`
	Timeouts                  *kustoeventgriddataconnection.TimeoutsState `json:"timeouts"`
}