// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	digitaltwinstimeseriesdatabaseconnection "github.com/golingon/terraproviders/azurerm/3.49.0/digitaltwinstimeseriesdatabaseconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDigitalTwinsTimeSeriesDatabaseConnection(name string, args DigitalTwinsTimeSeriesDatabaseConnectionArgs) *DigitalTwinsTimeSeriesDatabaseConnection {
	return &DigitalTwinsTimeSeriesDatabaseConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DigitalTwinsTimeSeriesDatabaseConnection)(nil)

type DigitalTwinsTimeSeriesDatabaseConnection struct {
	Name  string
	Args  DigitalTwinsTimeSeriesDatabaseConnectionArgs
	state *digitalTwinsTimeSeriesDatabaseConnectionState
}

func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) Type() string {
	return "azurerm_digital_twins_time_series_database_connection"
}

func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) LocalName() string {
	return dttsdc.Name
}

func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) Configuration() interface{} {
	return dttsdc.Args
}

func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) Attributes() digitalTwinsTimeSeriesDatabaseConnectionAttributes {
	return digitalTwinsTimeSeriesDatabaseConnectionAttributes{ref: terra.ReferenceResource(dttsdc)}
}

func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) ImportState(av io.Reader) error {
	dttsdc.state = &digitalTwinsTimeSeriesDatabaseConnectionState{}
	if err := json.NewDecoder(av).Decode(dttsdc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dttsdc.Type(), dttsdc.LocalName(), err)
	}
	return nil
}

func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) State() (*digitalTwinsTimeSeriesDatabaseConnectionState, bool) {
	return dttsdc.state, dttsdc.state != nil
}

func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) StateMust() *digitalTwinsTimeSeriesDatabaseConnectionState {
	if dttsdc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dttsdc.Type(), dttsdc.LocalName()))
	}
	return dttsdc.state
}

func (dttsdc *DigitalTwinsTimeSeriesDatabaseConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(dttsdc)
}

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
	// DependsOn contains resources that DigitalTwinsTimeSeriesDatabaseConnection depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type digitalTwinsTimeSeriesDatabaseConnectionAttributes struct {
	ref terra.Reference
}

func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) DigitalTwinsId() terra.StringValue {
	return terra.ReferenceString(dttsdc.ref.Append("digital_twins_id"))
}

func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) EventhubConsumerGroupName() terra.StringValue {
	return terra.ReferenceString(dttsdc.ref.Append("eventhub_consumer_group_name"))
}

func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) EventhubName() terra.StringValue {
	return terra.ReferenceString(dttsdc.ref.Append("eventhub_name"))
}

func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) EventhubNamespaceEndpointUri() terra.StringValue {
	return terra.ReferenceString(dttsdc.ref.Append("eventhub_namespace_endpoint_uri"))
}

func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) EventhubNamespaceId() terra.StringValue {
	return terra.ReferenceString(dttsdc.ref.Append("eventhub_namespace_id"))
}

func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dttsdc.ref.Append("id"))
}

func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) KustoClusterId() terra.StringValue {
	return terra.ReferenceString(dttsdc.ref.Append("kusto_cluster_id"))
}

func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) KustoClusterUri() terra.StringValue {
	return terra.ReferenceString(dttsdc.ref.Append("kusto_cluster_uri"))
}

func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) KustoDatabaseName() terra.StringValue {
	return terra.ReferenceString(dttsdc.ref.Append("kusto_database_name"))
}

func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) KustoTableName() terra.StringValue {
	return terra.ReferenceString(dttsdc.ref.Append("kusto_table_name"))
}

func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dttsdc.ref.Append("name"))
}

func (dttsdc digitalTwinsTimeSeriesDatabaseConnectionAttributes) Timeouts() digitaltwinstimeseriesdatabaseconnection.TimeoutsAttributes {
	return terra.ReferenceSingle[digitaltwinstimeseriesdatabaseconnection.TimeoutsAttributes](dttsdc.ref.Append("timeouts"))
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
