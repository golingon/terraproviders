// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_eventhub_consumer_group

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_eventhub_consumer_group.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aecg *DataSource) DataSource() string {
	return "azurerm_eventhub_consumer_group"
}

// LocalName returns the local name for [DataSource].
func (aecg *DataSource) LocalName() string {
	return aecg.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aecg *DataSource) Configuration() interface{} {
	return aecg.Args
}

// Attributes returns the attributes for [DataSource].
func (aecg *DataSource) Attributes() dataAzurermEventhubConsumerGroupAttributes {
	return dataAzurermEventhubConsumerGroupAttributes{ref: terra.ReferenceDataSource(aecg)}
}

// DataArgs contains the configurations for azurerm_eventhub_consumer_group.
type DataArgs struct {
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
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermEventhubConsumerGroupAttributes struct {
	ref terra.Reference
}

// EventhubName returns a reference to field eventhub_name of azurerm_eventhub_consumer_group.
func (aecg dataAzurermEventhubConsumerGroupAttributes) EventhubName() terra.StringValue {
	return terra.ReferenceAsString(aecg.ref.Append("eventhub_name"))
}

// Id returns a reference to field id of azurerm_eventhub_consumer_group.
func (aecg dataAzurermEventhubConsumerGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aecg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_eventhub_consumer_group.
func (aecg dataAzurermEventhubConsumerGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aecg.ref.Append("name"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_eventhub_consumer_group.
func (aecg dataAzurermEventhubConsumerGroupAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(aecg.ref.Append("namespace_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventhub_consumer_group.
func (aecg dataAzurermEventhubConsumerGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aecg.ref.Append("resource_group_name"))
}

// UserMetadata returns a reference to field user_metadata of azurerm_eventhub_consumer_group.
func (aecg dataAzurermEventhubConsumerGroupAttributes) UserMetadata() terra.StringValue {
	return terra.ReferenceAsString(aecg.ref.Append("user_metadata"))
}

func (aecg dataAzurermEventhubConsumerGroupAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](aecg.ref.Append("timeouts"))
}
