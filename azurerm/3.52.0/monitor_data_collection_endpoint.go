// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitordatacollectionendpoint "github.com/golingon/terraproviders/azurerm/3.52.0/monitordatacollectionendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorDataCollectionEndpoint creates a new instance of [MonitorDataCollectionEndpoint].
func NewMonitorDataCollectionEndpoint(name string, args MonitorDataCollectionEndpointArgs) *MonitorDataCollectionEndpoint {
	return &MonitorDataCollectionEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorDataCollectionEndpoint)(nil)

// MonitorDataCollectionEndpoint represents the Terraform resource azurerm_monitor_data_collection_endpoint.
type MonitorDataCollectionEndpoint struct {
	Name      string
	Args      MonitorDataCollectionEndpointArgs
	state     *monitorDataCollectionEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorDataCollectionEndpoint].
func (mdce *MonitorDataCollectionEndpoint) Type() string {
	return "azurerm_monitor_data_collection_endpoint"
}

// LocalName returns the local name for [MonitorDataCollectionEndpoint].
func (mdce *MonitorDataCollectionEndpoint) LocalName() string {
	return mdce.Name
}

// Configuration returns the configuration (args) for [MonitorDataCollectionEndpoint].
func (mdce *MonitorDataCollectionEndpoint) Configuration() interface{} {
	return mdce.Args
}

// DependOn is used for other resources to depend on [MonitorDataCollectionEndpoint].
func (mdce *MonitorDataCollectionEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(mdce)
}

// Dependencies returns the list of resources [MonitorDataCollectionEndpoint] depends_on.
func (mdce *MonitorDataCollectionEndpoint) Dependencies() terra.Dependencies {
	return mdce.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorDataCollectionEndpoint].
func (mdce *MonitorDataCollectionEndpoint) LifecycleManagement() *terra.Lifecycle {
	return mdce.Lifecycle
}

// Attributes returns the attributes for [MonitorDataCollectionEndpoint].
func (mdce *MonitorDataCollectionEndpoint) Attributes() monitorDataCollectionEndpointAttributes {
	return monitorDataCollectionEndpointAttributes{ref: terra.ReferenceResource(mdce)}
}

// ImportState imports the given attribute values into [MonitorDataCollectionEndpoint]'s state.
func (mdce *MonitorDataCollectionEndpoint) ImportState(av io.Reader) error {
	mdce.state = &monitorDataCollectionEndpointState{}
	if err := json.NewDecoder(av).Decode(mdce.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mdce.Type(), mdce.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorDataCollectionEndpoint] has state.
func (mdce *MonitorDataCollectionEndpoint) State() (*monitorDataCollectionEndpointState, bool) {
	return mdce.state, mdce.state != nil
}

// StateMust returns the state for [MonitorDataCollectionEndpoint]. Panics if the state is nil.
func (mdce *MonitorDataCollectionEndpoint) StateMust() *monitorDataCollectionEndpointState {
	if mdce.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mdce.Type(), mdce.LocalName()))
	}
	return mdce.state
}

// MonitorDataCollectionEndpointArgs contains the configurations for azurerm_monitor_data_collection_endpoint.
type MonitorDataCollectionEndpointArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Kind: string, optional
	Kind terra.StringValue `hcl:"kind,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *monitordatacollectionendpoint.Timeouts `hcl:"timeouts,block"`
}
type monitorDataCollectionEndpointAttributes struct {
	ref terra.Reference
}

// ConfigurationAccessEndpoint returns a reference to field configuration_access_endpoint of azurerm_monitor_data_collection_endpoint.
func (mdce monitorDataCollectionEndpointAttributes) ConfigurationAccessEndpoint() terra.StringValue {
	return terra.ReferenceAsString(mdce.ref.Append("configuration_access_endpoint"))
}

// Description returns a reference to field description of azurerm_monitor_data_collection_endpoint.
func (mdce monitorDataCollectionEndpointAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mdce.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_monitor_data_collection_endpoint.
func (mdce monitorDataCollectionEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mdce.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_monitor_data_collection_endpoint.
func (mdce monitorDataCollectionEndpointAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(mdce.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_monitor_data_collection_endpoint.
func (mdce monitorDataCollectionEndpointAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mdce.ref.Append("location"))
}

// LogsIngestionEndpoint returns a reference to field logs_ingestion_endpoint of azurerm_monitor_data_collection_endpoint.
func (mdce monitorDataCollectionEndpointAttributes) LogsIngestionEndpoint() terra.StringValue {
	return terra.ReferenceAsString(mdce.ref.Append("logs_ingestion_endpoint"))
}

// Name returns a reference to field name of azurerm_monitor_data_collection_endpoint.
func (mdce monitorDataCollectionEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mdce.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_monitor_data_collection_endpoint.
func (mdce monitorDataCollectionEndpointAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mdce.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_data_collection_endpoint.
func (mdce monitorDataCollectionEndpointAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mdce.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_monitor_data_collection_endpoint.
func (mdce monitorDataCollectionEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mdce.ref.Append("tags"))
}

func (mdce monitorDataCollectionEndpointAttributes) Timeouts() monitordatacollectionendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitordatacollectionendpoint.TimeoutsAttributes](mdce.ref.Append("timeouts"))
}

type monitorDataCollectionEndpointState struct {
	ConfigurationAccessEndpoint string                                       `json:"configuration_access_endpoint"`
	Description                 string                                       `json:"description"`
	Id                          string                                       `json:"id"`
	Kind                        string                                       `json:"kind"`
	Location                    string                                       `json:"location"`
	LogsIngestionEndpoint       string                                       `json:"logs_ingestion_endpoint"`
	Name                        string                                       `json:"name"`
	PublicNetworkAccessEnabled  bool                                         `json:"public_network_access_enabled"`
	ResourceGroupName           string                                       `json:"resource_group_name"`
	Tags                        map[string]string                            `json:"tags"`
	Timeouts                    *monitordatacollectionendpoint.TimeoutsState `json:"timeouts"`
}
