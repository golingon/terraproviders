// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataeventhubconsumergroup "github.com/golingon/terraproviders/azurerm/3.49.0/dataeventhubconsumergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEventhubConsumerGroup creates a new instance of [DataEventhubConsumerGroup].
func NewDataEventhubConsumerGroup(name string, args DataEventhubConsumerGroupArgs) *DataEventhubConsumerGroup {
	return &DataEventhubConsumerGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEventhubConsumerGroup)(nil)

// DataEventhubConsumerGroup represents the Terraform data resource azurerm_eventhub_consumer_group.
type DataEventhubConsumerGroup struct {
	Name string
	Args DataEventhubConsumerGroupArgs
}

// DataSource returns the Terraform object type for [DataEventhubConsumerGroup].
func (ecg *DataEventhubConsumerGroup) DataSource() string {
	return "azurerm_eventhub_consumer_group"
}

// LocalName returns the local name for [DataEventhubConsumerGroup].
func (ecg *DataEventhubConsumerGroup) LocalName() string {
	return ecg.Name
}

// Configuration returns the configuration (args) for [DataEventhubConsumerGroup].
func (ecg *DataEventhubConsumerGroup) Configuration() interface{} {
	return ecg.Args
}

// Attributes returns the attributes for [DataEventhubConsumerGroup].
func (ecg *DataEventhubConsumerGroup) Attributes() dataEventhubConsumerGroupAttributes {
	return dataEventhubConsumerGroupAttributes{ref: terra.ReferenceDataResource(ecg)}
}

// DataEventhubConsumerGroupArgs contains the configurations for azurerm_eventhub_consumer_group.
type DataEventhubConsumerGroupArgs struct {
	// EventhubName: string, required
	EventhubName terra.StringValue `hcl:"eventhub_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceName: string, required
	NamespaceName terra.StringValue `hcl:"namespace_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataeventhubconsumergroup.Timeouts `hcl:"timeouts,block"`
}
type dataEventhubConsumerGroupAttributes struct {
	ref terra.Reference
}

// EventhubName returns a reference to field eventhub_name of azurerm_eventhub_consumer_group.
func (ecg dataEventhubConsumerGroupAttributes) EventhubName() terra.StringValue {
	return terra.ReferenceAsString(ecg.ref.Append("eventhub_name"))
}

// Id returns a reference to field id of azurerm_eventhub_consumer_group.
func (ecg dataEventhubConsumerGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ecg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_eventhub_consumer_group.
func (ecg dataEventhubConsumerGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ecg.ref.Append("name"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_eventhub_consumer_group.
func (ecg dataEventhubConsumerGroupAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(ecg.ref.Append("namespace_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventhub_consumer_group.
func (ecg dataEventhubConsumerGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ecg.ref.Append("resource_group_name"))
}

// UserMetadata returns a reference to field user_metadata of azurerm_eventhub_consumer_group.
func (ecg dataEventhubConsumerGroupAttributes) UserMetadata() terra.StringValue {
	return terra.ReferenceAsString(ecg.ref.Append("user_metadata"))
}

func (ecg dataEventhubConsumerGroupAttributes) Timeouts() dataeventhubconsumergroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataeventhubconsumergroup.TimeoutsAttributes](ecg.ref.Append("timeouts"))
}
