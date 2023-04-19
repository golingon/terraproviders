// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamonitordatacollectionendpoint "github.com/golingon/terraproviders/azurerm/3.52.0/datamonitordatacollectionendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMonitorDataCollectionEndpoint creates a new instance of [DataMonitorDataCollectionEndpoint].
func NewDataMonitorDataCollectionEndpoint(name string, args DataMonitorDataCollectionEndpointArgs) *DataMonitorDataCollectionEndpoint {
	return &DataMonitorDataCollectionEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMonitorDataCollectionEndpoint)(nil)

// DataMonitorDataCollectionEndpoint represents the Terraform data resource azurerm_monitor_data_collection_endpoint.
type DataMonitorDataCollectionEndpoint struct {
	Name string
	Args DataMonitorDataCollectionEndpointArgs
}

// DataSource returns the Terraform object type for [DataMonitorDataCollectionEndpoint].
func (mdce *DataMonitorDataCollectionEndpoint) DataSource() string {
	return "azurerm_monitor_data_collection_endpoint"
}

// LocalName returns the local name for [DataMonitorDataCollectionEndpoint].
func (mdce *DataMonitorDataCollectionEndpoint) LocalName() string {
	return mdce.Name
}

// Configuration returns the configuration (args) for [DataMonitorDataCollectionEndpoint].
func (mdce *DataMonitorDataCollectionEndpoint) Configuration() interface{} {
	return mdce.Args
}

// Attributes returns the attributes for [DataMonitorDataCollectionEndpoint].
func (mdce *DataMonitorDataCollectionEndpoint) Attributes() dataMonitorDataCollectionEndpointAttributes {
	return dataMonitorDataCollectionEndpointAttributes{ref: terra.ReferenceDataResource(mdce)}
}

// DataMonitorDataCollectionEndpointArgs contains the configurations for azurerm_monitor_data_collection_endpoint.
type DataMonitorDataCollectionEndpointArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datamonitordatacollectionendpoint.Timeouts `hcl:"timeouts,block"`
}
type dataMonitorDataCollectionEndpointAttributes struct {
	ref terra.Reference
}

// ConfigurationAccessEndpoint returns a reference to field configuration_access_endpoint of azurerm_monitor_data_collection_endpoint.
func (mdce dataMonitorDataCollectionEndpointAttributes) ConfigurationAccessEndpoint() terra.StringValue {
	return terra.ReferenceAsString(mdce.ref.Append("configuration_access_endpoint"))
}

// Description returns a reference to field description of azurerm_monitor_data_collection_endpoint.
func (mdce dataMonitorDataCollectionEndpointAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mdce.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_monitor_data_collection_endpoint.
func (mdce dataMonitorDataCollectionEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mdce.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_monitor_data_collection_endpoint.
func (mdce dataMonitorDataCollectionEndpointAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(mdce.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_monitor_data_collection_endpoint.
func (mdce dataMonitorDataCollectionEndpointAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mdce.ref.Append("location"))
}

// LogsIngestionEndpoint returns a reference to field logs_ingestion_endpoint of azurerm_monitor_data_collection_endpoint.
func (mdce dataMonitorDataCollectionEndpointAttributes) LogsIngestionEndpoint() terra.StringValue {
	return terra.ReferenceAsString(mdce.ref.Append("logs_ingestion_endpoint"))
}

// Name returns a reference to field name of azurerm_monitor_data_collection_endpoint.
func (mdce dataMonitorDataCollectionEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mdce.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_monitor_data_collection_endpoint.
func (mdce dataMonitorDataCollectionEndpointAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mdce.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_data_collection_endpoint.
func (mdce dataMonitorDataCollectionEndpointAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mdce.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_monitor_data_collection_endpoint.
func (mdce dataMonitorDataCollectionEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mdce.ref.Append("tags"))
}

func (mdce dataMonitorDataCollectionEndpointAttributes) Timeouts() datamonitordatacollectionendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamonitordatacollectionendpoint.TimeoutsAttributes](mdce.ref.Append("timeouts"))
}
