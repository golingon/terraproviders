// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	kustoiothubdataconnection "github.com/golingon/terraproviders/azurerm/3.49.0/kustoiothubdataconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKustoIothubDataConnection creates a new instance of [KustoIothubDataConnection].
func NewKustoIothubDataConnection(name string, args KustoIothubDataConnectionArgs) *KustoIothubDataConnection {
	return &KustoIothubDataConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KustoIothubDataConnection)(nil)

// KustoIothubDataConnection represents the Terraform resource azurerm_kusto_iothub_data_connection.
type KustoIothubDataConnection struct {
	Name      string
	Args      KustoIothubDataConnectionArgs
	state     *kustoIothubDataConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KustoIothubDataConnection].
func (kidc *KustoIothubDataConnection) Type() string {
	return "azurerm_kusto_iothub_data_connection"
}

// LocalName returns the local name for [KustoIothubDataConnection].
func (kidc *KustoIothubDataConnection) LocalName() string {
	return kidc.Name
}

// Configuration returns the configuration (args) for [KustoIothubDataConnection].
func (kidc *KustoIothubDataConnection) Configuration() interface{} {
	return kidc.Args
}

// DependOn is used for other resources to depend on [KustoIothubDataConnection].
func (kidc *KustoIothubDataConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(kidc)
}

// Dependencies returns the list of resources [KustoIothubDataConnection] depends_on.
func (kidc *KustoIothubDataConnection) Dependencies() terra.Dependencies {
	return kidc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KustoIothubDataConnection].
func (kidc *KustoIothubDataConnection) LifecycleManagement() *terra.Lifecycle {
	return kidc.Lifecycle
}

// Attributes returns the attributes for [KustoIothubDataConnection].
func (kidc *KustoIothubDataConnection) Attributes() kustoIothubDataConnectionAttributes {
	return kustoIothubDataConnectionAttributes{ref: terra.ReferenceResource(kidc)}
}

// ImportState imports the given attribute values into [KustoIothubDataConnection]'s state.
func (kidc *KustoIothubDataConnection) ImportState(av io.Reader) error {
	kidc.state = &kustoIothubDataConnectionState{}
	if err := json.NewDecoder(av).Decode(kidc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kidc.Type(), kidc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KustoIothubDataConnection] has state.
func (kidc *KustoIothubDataConnection) State() (*kustoIothubDataConnectionState, bool) {
	return kidc.state, kidc.state != nil
}

// StateMust returns the state for [KustoIothubDataConnection]. Panics if the state is nil.
func (kidc *KustoIothubDataConnection) StateMust() *kustoIothubDataConnectionState {
	if kidc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kidc.Type(), kidc.LocalName()))
	}
	return kidc.state
}

// KustoIothubDataConnectionArgs contains the configurations for azurerm_kusto_iothub_data_connection.
type KustoIothubDataConnectionArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// ConsumerGroup: string, required
	ConsumerGroup terra.StringValue `hcl:"consumer_group,attr" validate:"required"`
	// DataFormat: string, optional
	DataFormat terra.StringValue `hcl:"data_format,attr"`
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// DatabaseRoutingType: string, optional
	DatabaseRoutingType terra.StringValue `hcl:"database_routing_type,attr"`
	// EventSystemProperties: set of string, optional
	EventSystemProperties terra.SetValue[terra.StringValue] `hcl:"event_system_properties,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IothubId: string, required
	IothubId terra.StringValue `hcl:"iothub_id,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MappingRuleName: string, optional
	MappingRuleName terra.StringValue `hcl:"mapping_rule_name,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SharedAccessPolicyName: string, required
	SharedAccessPolicyName terra.StringValue `hcl:"shared_access_policy_name,attr" validate:"required"`
	// TableName: string, optional
	TableName terra.StringValue `hcl:"table_name,attr"`
	// Timeouts: optional
	Timeouts *kustoiothubdataconnection.Timeouts `hcl:"timeouts,block"`
}
type kustoIothubDataConnectionAttributes struct {
	ref terra.Reference
}

// ClusterName returns a reference to field cluster_name of azurerm_kusto_iothub_data_connection.
func (kidc kustoIothubDataConnectionAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(kidc.ref.Append("cluster_name"))
}

// ConsumerGroup returns a reference to field consumer_group of azurerm_kusto_iothub_data_connection.
func (kidc kustoIothubDataConnectionAttributes) ConsumerGroup() terra.StringValue {
	return terra.ReferenceAsString(kidc.ref.Append("consumer_group"))
}

// DataFormat returns a reference to field data_format of azurerm_kusto_iothub_data_connection.
func (kidc kustoIothubDataConnectionAttributes) DataFormat() terra.StringValue {
	return terra.ReferenceAsString(kidc.ref.Append("data_format"))
}

// DatabaseName returns a reference to field database_name of azurerm_kusto_iothub_data_connection.
func (kidc kustoIothubDataConnectionAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(kidc.ref.Append("database_name"))
}

// DatabaseRoutingType returns a reference to field database_routing_type of azurerm_kusto_iothub_data_connection.
func (kidc kustoIothubDataConnectionAttributes) DatabaseRoutingType() terra.StringValue {
	return terra.ReferenceAsString(kidc.ref.Append("database_routing_type"))
}

// EventSystemProperties returns a reference to field event_system_properties of azurerm_kusto_iothub_data_connection.
func (kidc kustoIothubDataConnectionAttributes) EventSystemProperties() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kidc.ref.Append("event_system_properties"))
}

// Id returns a reference to field id of azurerm_kusto_iothub_data_connection.
func (kidc kustoIothubDataConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kidc.ref.Append("id"))
}

// IothubId returns a reference to field iothub_id of azurerm_kusto_iothub_data_connection.
func (kidc kustoIothubDataConnectionAttributes) IothubId() terra.StringValue {
	return terra.ReferenceAsString(kidc.ref.Append("iothub_id"))
}

// Location returns a reference to field location of azurerm_kusto_iothub_data_connection.
func (kidc kustoIothubDataConnectionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(kidc.ref.Append("location"))
}

// MappingRuleName returns a reference to field mapping_rule_name of azurerm_kusto_iothub_data_connection.
func (kidc kustoIothubDataConnectionAttributes) MappingRuleName() terra.StringValue {
	return terra.ReferenceAsString(kidc.ref.Append("mapping_rule_name"))
}

// Name returns a reference to field name of azurerm_kusto_iothub_data_connection.
func (kidc kustoIothubDataConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kidc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_kusto_iothub_data_connection.
func (kidc kustoIothubDataConnectionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(kidc.ref.Append("resource_group_name"))
}

// SharedAccessPolicyName returns a reference to field shared_access_policy_name of azurerm_kusto_iothub_data_connection.
func (kidc kustoIothubDataConnectionAttributes) SharedAccessPolicyName() terra.StringValue {
	return terra.ReferenceAsString(kidc.ref.Append("shared_access_policy_name"))
}

// TableName returns a reference to field table_name of azurerm_kusto_iothub_data_connection.
func (kidc kustoIothubDataConnectionAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(kidc.ref.Append("table_name"))
}

func (kidc kustoIothubDataConnectionAttributes) Timeouts() kustoiothubdataconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kustoiothubdataconnection.TimeoutsAttributes](kidc.ref.Append("timeouts"))
}

type kustoIothubDataConnectionState struct {
	ClusterName            string                                   `json:"cluster_name"`
	ConsumerGroup          string                                   `json:"consumer_group"`
	DataFormat             string                                   `json:"data_format"`
	DatabaseName           string                                   `json:"database_name"`
	DatabaseRoutingType    string                                   `json:"database_routing_type"`
	EventSystemProperties  []string                                 `json:"event_system_properties"`
	Id                     string                                   `json:"id"`
	IothubId               string                                   `json:"iothub_id"`
	Location               string                                   `json:"location"`
	MappingRuleName        string                                   `json:"mapping_rule_name"`
	Name                   string                                   `json:"name"`
	ResourceGroupName      string                                   `json:"resource_group_name"`
	SharedAccessPolicyName string                                   `json:"shared_access_policy_name"`
	TableName              string                                   `json:"table_name"`
	Timeouts               *kustoiothubdataconnection.TimeoutsState `json:"timeouts"`
}
