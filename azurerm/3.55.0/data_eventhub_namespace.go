// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataeventhubnamespace "github.com/golingon/terraproviders/azurerm/3.55.0/dataeventhubnamespace"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEventhubNamespace creates a new instance of [DataEventhubNamespace].
func NewDataEventhubNamespace(name string, args DataEventhubNamespaceArgs) *DataEventhubNamespace {
	return &DataEventhubNamespace{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEventhubNamespace)(nil)

// DataEventhubNamespace represents the Terraform data resource azurerm_eventhub_namespace.
type DataEventhubNamespace struct {
	Name string
	Args DataEventhubNamespaceArgs
}

// DataSource returns the Terraform object type for [DataEventhubNamespace].
func (en *DataEventhubNamespace) DataSource() string {
	return "azurerm_eventhub_namespace"
}

// LocalName returns the local name for [DataEventhubNamespace].
func (en *DataEventhubNamespace) LocalName() string {
	return en.Name
}

// Configuration returns the configuration (args) for [DataEventhubNamespace].
func (en *DataEventhubNamespace) Configuration() interface{} {
	return en.Args
}

// Attributes returns the attributes for [DataEventhubNamespace].
func (en *DataEventhubNamespace) Attributes() dataEventhubNamespaceAttributes {
	return dataEventhubNamespaceAttributes{ref: terra.ReferenceDataResource(en)}
}

// DataEventhubNamespaceArgs contains the configurations for azurerm_eventhub_namespace.
type DataEventhubNamespaceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataeventhubnamespace.Timeouts `hcl:"timeouts,block"`
}
type dataEventhubNamespaceAttributes struct {
	ref terra.Reference
}

// AutoInflateEnabled returns a reference to field auto_inflate_enabled of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) AutoInflateEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(en.ref.Append("auto_inflate_enabled"))
}

// Capacity returns a reference to field capacity of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(en.ref.Append("capacity"))
}

// DedicatedClusterId returns a reference to field dedicated_cluster_id of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) DedicatedClusterId() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("dedicated_cluster_id"))
}

// DefaultPrimaryConnectionString returns a reference to field default_primary_connection_string of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) DefaultPrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("default_primary_connection_string"))
}

// DefaultPrimaryConnectionStringAlias returns a reference to field default_primary_connection_string_alias of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) DefaultPrimaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("default_primary_connection_string_alias"))
}

// DefaultPrimaryKey returns a reference to field default_primary_key of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) DefaultPrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("default_primary_key"))
}

// DefaultSecondaryConnectionString returns a reference to field default_secondary_connection_string of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) DefaultSecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("default_secondary_connection_string"))
}

// DefaultSecondaryConnectionStringAlias returns a reference to field default_secondary_connection_string_alias of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) DefaultSecondaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("default_secondary_connection_string_alias"))
}

// DefaultSecondaryKey returns a reference to field default_secondary_key of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) DefaultSecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("default_secondary_key"))
}

// Id returns a reference to field id of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("id"))
}

// KafkaEnabled returns a reference to field kafka_enabled of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) KafkaEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(en.ref.Append("kafka_enabled"))
}

// Location returns a reference to field location of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("location"))
}

// MaximumThroughputUnits returns a reference to field maximum_throughput_units of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) MaximumThroughputUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(en.ref.Append("maximum_throughput_units"))
}

// Name returns a reference to field name of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(en.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](en.ref.Append("tags"))
}

// ZoneRedundant returns a reference to field zone_redundant of azurerm_eventhub_namespace.
func (en dataEventhubNamespaceAttributes) ZoneRedundant() terra.BoolValue {
	return terra.ReferenceAsBool(en.ref.Append("zone_redundant"))
}

func (en dataEventhubNamespaceAttributes) Timeouts() dataeventhubnamespace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataeventhubnamespace.TimeoutsAttributes](en.ref.Append("timeouts"))
}
