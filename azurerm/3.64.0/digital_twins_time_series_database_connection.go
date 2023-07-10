// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	digitaltwinstimeseriesdatabaseconnection "github.com/golingon/terraproviders/azurerm/3.64.0/digitaltwinstimeseriesdatabaseconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDigitalTwinsTimeSeriesDatabaseConnection creates a new instance of [DigitalTwinsTimeSeriesDatabaseConnection].
func NewDigitalTwinsTimeSeriesDatabaseConnection(name string, args DigitalTwinsTimeSeriesDatabaseConnectionArgs) *DigitalTwinsTimeSeriesDatabaseConnection {
	return &DigitalTwinsTimeSeriesDatabaseConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DigitalTwinsTimeSeriesDatabaseConnection)(nil)

// DigitalTwinsTimeSeriesDatabaseConnection represents the Terraform resource azurerm_digital_twins_time_series_database_connection.
type DigitalTwinsTimeSeriesDatabaseConnection struct {
	Name      string
	Args      DigitalTwinsTimeSeriesDatabaseConnectionArgs
	state     *digitalTwinsTimeSeriesDatabaseConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DigitalTwinsTimeSeriesDatabaseConnection].
func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) Type() string {
	return "azurerm_digital_twins_time_series_database_connection"
}

// LocalName returns the local name for [DigitalTwinsTimeSeriesDatabaseConnection].
func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) LocalName() string {
	return dttsdc.Name
}

// Configuration returns the configuration (args) for [DigitalTwinsTimeSeriesDatabaseConnection].
func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) Configuration() interface{} {
	return dttsdc.Args
}

// DependOn is used for other resources to depend on [DigitalTwinsTimeSeriesDatabaseConnection].
func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(dttsdc)
}

// Dependencies returns the list of resources [DigitalTwinsTimeSeriesDatabaseConnection] depends_on.
func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) Dependencies() terra.Dependencies {
	return dttsdc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DigitalTwinsTimeSeriesDatabaseConnection].
func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) LifecycleManagement() *terra.Lifecycle {
	return dttsdc.Lifecycle
}

// Attributes returns the attributes for [DigitalTwinsTimeSeriesDatabaseConnection].
func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) Attributes() digitalTwinsTimeSeriesDatabaseConnectionAttributes {
	return digitalTwinsTimeSeriesDatabaseConnectionAttributes{ref: terra.ReferenceResource(dttsdc)}
}

// ImportState imports the given attribute values into [DigitalTwinsTimeSeriesDatabaseConnection]'s state.
func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) ImportState(av io.Reader) error {
	dttsdc.state = &digitalTwinsTimeSeriesDatabaseConnectionState{}
	if err := json.NewDecoder(av).Decode(dttsdc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dttsdc.Type(), dttsdc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DigitalTwinsTimeSeriesDatabaseConnection] has state.
func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) State() (*digitalTwinsTimeSeriesDatabaseConnectionState, bool) {
	return dttsdc.state, dttsdc.state != nil
}

// StateMust returns the state for [DigitalTwinsTimeSeriesDatabaseConnection]. Panics if the state is nil.
func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) StateMust() *digitalTwinsTimeSeriesDatabaseConnectionState {
	if dttsdc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dttsdc.Type(), dttsdc.LocalName()))
	}
	return dttsdc.state
}

// DigitalTwinsTimeSeriesDatabaseConnectionArgs contains the configurations for azurerm_digital_twins_time_series_database_connection.
type DigitalTwinsTimeSeriesDatabaseConnectionArgs struct {
	// DigitalTwinsId: string, required
	DigitalTwinsId terra.StringValue `hcl:"digital_twins_id,attr" validate:"required"`
	// EventhubConsumerGroupName: string, optional
	EventhubConsumerGroupName terra.StringValue `hcl:"eventhub_consumer_group_name,attr"`
	// EventhubName: string, required
	EventhubName terra.StringValue `hcl:"eventhub_name,attr" validate:"required"`
	// EventhubNamespaceEndpointUri: string, required
	EventhubNamespaceEndpointUri terra.StringValue `hcl:"eventhub_namespace_endpoint_uri,attr" validate:"required"`
	// EventhubNamespaceId: string, required
	EventhubNamespaceId terra.StringValue `hcl:"eventhub_namespace_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KustoClusterId: string, required
	KustoClusterId terra.StringValue `hcl:"kusto_cluster_id,attr" validate:"required"`
	// KustoClusterUri: string, required
	KustoClusterUri terra.StringValue `hcl:"kusto_cluster_uri,attr" validate:"required"`
	// KustoDatabaseName: string, required
	KustoDatabaseName terra.StringValue `hcl:"kusto_database_name,attr" validate:"required"`
	// KustoTableName: string, optional
	KustoTableName terra.StringValue `hcl:"kusto_table_name,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *digitaltwinstimeseriesdatabaseconnection.Timeouts `hcl:"timeouts,block"`
}
type digitalTwinsTimeSeriesDatabaseConnectionAttributes struct {
	ref terra.Reference
}

// DigitalTwinsId returns a reference to field digital_twins_id of azurerm_digital_twins_time_series_database_connection.
func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) DigitalTwinsId() terra.StringValue {
	return terra.ReferenceAsString(dttsdc.ref.Append("digital_twins_id"))
}

// EventhubConsumerGroupName returns a reference to field eventhub_consumer_group_name of azurerm_digital_twins_time_series_database_connection.
func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) EventhubConsumerGroupName() terra.StringValue {
	return terra.ReferenceAsString(dttsdc.ref.Append("eventhub_consumer_group_name"))
}

// EventhubName returns a reference to field eventhub_name of azurerm_digital_twins_time_series_database_connection.
func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) EventhubName() terra.StringValue {
	return terra.ReferenceAsString(dttsdc.ref.Append("eventhub_name"))
}

// EventhubNamespaceEndpointUri returns a reference to field eventhub_namespace_endpoint_uri of azurerm_digital_twins_time_series_database_connection.
func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) EventhubNamespaceEndpointUri() terra.StringValue {
	return terra.ReferenceAsString(dttsdc.ref.Append("eventhub_namespace_endpoint_uri"))
}

// EventhubNamespaceId returns a reference to field eventhub_namespace_id of azurerm_digital_twins_time_series_database_connection.
func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) EventhubNamespaceId() terra.StringValue {
	return terra.ReferenceAsString(dttsdc.ref.Append("eventhub_namespace_id"))
}

// Id returns a reference to field id of azurerm_digital_twins_time_series_database_connection.
func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dttsdc.ref.Append("id"))
}

// KustoClusterId returns a reference to field kusto_cluster_id of azurerm_digital_twins_time_series_database_connection.
func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) KustoClusterId() terra.StringValue {
	return terra.ReferenceAsString(dttsdc.ref.Append("kusto_cluster_id"))
}

// KustoClusterUri returns a reference to field kusto_cluster_uri of azurerm_digital_twins_time_series_database_connection.
func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) KustoClusterUri() terra.StringValue {
	return terra.ReferenceAsString(dttsdc.ref.Append("kusto_cluster_uri"))
}

// KustoDatabaseName returns a reference to field kusto_database_name of azurerm_digital_twins_time_series_database_connection.
func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) KustoDatabaseName() terra.StringValue {
	return terra.ReferenceAsString(dttsdc.ref.Append("kusto_database_name"))
}

// KustoTableName returns a reference to field kusto_table_name of azurerm_digital_twins_time_series_database_connection.
func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) KustoTableName() terra.StringValue {
	return terra.ReferenceAsString(dttsdc.ref.Append("kusto_table_name"))
}

// Name returns a reference to field name of azurerm_digital_twins_time_series_database_connection.
func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dttsdc.ref.Append("name"))
}

func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) Timeouts() digitaltwinstimeseriesdatabaseconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[digitaltwinstimeseriesdatabaseconnection.TimeoutsAttributes](dttsdc.ref.Append("timeouts"))
}

type digitalTwinsTimeSeriesDatabaseConnectionState struct {
	DigitalTwinsId               string                                                  `json:"digital_twins_id"`
	EventhubConsumerGroupName    string                                                  `json:"eventhub_consumer_group_name"`
	EventhubName                 string                                                  `json:"eventhub_name"`
	EventhubNamespaceEndpointUri string                                                  `json:"eventhub_namespace_endpoint_uri"`
	EventhubNamespaceId          string                                                  `json:"eventhub_namespace_id"`
	Id                           string                                                  `json:"id"`
	KustoClusterId               string                                                  `json:"kusto_cluster_id"`
	KustoClusterUri              string                                                  `json:"kusto_cluster_uri"`
	KustoDatabaseName            string                                                  `json:"kusto_database_name"`
	KustoTableName               string                                                  `json:"kusto_table_name"`
	Name                         string                                                  `json:"name"`
	Timeouts                     *digitaltwinstimeseriesdatabaseconnection.TimeoutsState `json:"timeouts"`
}
