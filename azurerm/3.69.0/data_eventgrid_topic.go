// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataeventgridtopic "github.com/golingon/terraproviders/azurerm/3.69.0/dataeventgridtopic"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEventgridTopic creates a new instance of [DataEventgridTopic].
func NewDataEventgridTopic(name string, args DataEventgridTopicArgs) *DataEventgridTopic {
	return &DataEventgridTopic{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEventgridTopic)(nil)

// DataEventgridTopic represents the Terraform data resource azurerm_eventgrid_topic.
type DataEventgridTopic struct {
	Name string
	Args DataEventgridTopicArgs
}

// DataSource returns the Terraform object type for [DataEventgridTopic].
func (et *DataEventgridTopic) DataSource() string {
	return "azurerm_eventgrid_topic"
}

// LocalName returns the local name for [DataEventgridTopic].
func (et *DataEventgridTopic) LocalName() string {
	return et.Name
}

// Configuration returns the configuration (args) for [DataEventgridTopic].
func (et *DataEventgridTopic) Configuration() interface{} {
	return et.Args
}

// Attributes returns the attributes for [DataEventgridTopic].
func (et *DataEventgridTopic) Attributes() dataEventgridTopicAttributes {
	return dataEventgridTopicAttributes{ref: terra.ReferenceDataResource(et)}
}

// DataEventgridTopicArgs contains the configurations for azurerm_eventgrid_topic.
type DataEventgridTopicArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataeventgridtopic.Timeouts `hcl:"timeouts,block"`
}
type dataEventgridTopicAttributes struct {
	ref terra.Reference
}

// Endpoint returns a reference to field endpoint of azurerm_eventgrid_topic.
func (et dataEventgridTopicAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("endpoint"))
}

// Id returns a reference to field id of azurerm_eventgrid_topic.
func (et dataEventgridTopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_eventgrid_topic.
func (et dataEventgridTopicAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_eventgrid_topic.
func (et dataEventgridTopicAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("name"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_eventgrid_topic.
func (et dataEventgridTopicAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("primary_access_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventgrid_topic.
func (et dataEventgridTopicAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_eventgrid_topic.
func (et dataEventgridTopicAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(et.ref.Append("secondary_access_key"))
}

// Tags returns a reference to field tags of azurerm_eventgrid_topic.
func (et dataEventgridTopicAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](et.ref.Append("tags"))
}

func (et dataEventgridTopicAttributes) Timeouts() dataeventgridtopic.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataeventgridtopic.TimeoutsAttributes](et.ref.Append("timeouts"))
}
