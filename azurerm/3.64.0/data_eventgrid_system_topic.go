// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataeventgridsystemtopic "github.com/golingon/terraproviders/azurerm/3.64.0/dataeventgridsystemtopic"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEventgridSystemTopic creates a new instance of [DataEventgridSystemTopic].
func NewDataEventgridSystemTopic(name string, args DataEventgridSystemTopicArgs) *DataEventgridSystemTopic {
	return &DataEventgridSystemTopic{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEventgridSystemTopic)(nil)

// DataEventgridSystemTopic represents the Terraform data resource azurerm_eventgrid_system_topic.
type DataEventgridSystemTopic struct {
	Name string
	Args DataEventgridSystemTopicArgs
}

// DataSource returns the Terraform object type for [DataEventgridSystemTopic].
func (est *DataEventgridSystemTopic) DataSource() string {
	return "azurerm_eventgrid_system_topic"
}

// LocalName returns the local name for [DataEventgridSystemTopic].
func (est *DataEventgridSystemTopic) LocalName() string {
	return est.Name
}

// Configuration returns the configuration (args) for [DataEventgridSystemTopic].
func (est *DataEventgridSystemTopic) Configuration() interface{} {
	return est.Args
}

// Attributes returns the attributes for [DataEventgridSystemTopic].
func (est *DataEventgridSystemTopic) Attributes() dataEventgridSystemTopicAttributes {
	return dataEventgridSystemTopicAttributes{ref: terra.ReferenceDataResource(est)}
}

// DataEventgridSystemTopicArgs contains the configurations for azurerm_eventgrid_system_topic.
type DataEventgridSystemTopicArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []dataeventgridsystemtopic.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataeventgridsystemtopic.Timeouts `hcl:"timeouts,block"`
}
type dataEventgridSystemTopicAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_eventgrid_system_topic.
func (est dataEventgridSystemTopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(est.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_eventgrid_system_topic.
func (est dataEventgridSystemTopicAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(est.ref.Append("location"))
}

// MetricArmResourceId returns a reference to field metric_arm_resource_id of azurerm_eventgrid_system_topic.
func (est dataEventgridSystemTopicAttributes) MetricArmResourceId() terra.StringValue {
	return terra.ReferenceAsString(est.ref.Append("metric_arm_resource_id"))
}

// Name returns a reference to field name of azurerm_eventgrid_system_topic.
func (est dataEventgridSystemTopicAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(est.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventgrid_system_topic.
func (est dataEventgridSystemTopicAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(est.ref.Append("resource_group_name"))
}

// SourceArmResourceId returns a reference to field source_arm_resource_id of azurerm_eventgrid_system_topic.
func (est dataEventgridSystemTopicAttributes) SourceArmResourceId() terra.StringValue {
	return terra.ReferenceAsString(est.ref.Append("source_arm_resource_id"))
}

// Tags returns a reference to field tags of azurerm_eventgrid_system_topic.
func (est dataEventgridSystemTopicAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](est.ref.Append("tags"))
}

// TopicType returns a reference to field topic_type of azurerm_eventgrid_system_topic.
func (est dataEventgridSystemTopicAttributes) TopicType() terra.StringValue {
	return terra.ReferenceAsString(est.ref.Append("topic_type"))
}

func (est dataEventgridSystemTopicAttributes) Identity() terra.ListValue[dataeventgridsystemtopic.IdentityAttributes] {
	return terra.ReferenceAsList[dataeventgridsystemtopic.IdentityAttributes](est.ref.Append("identity"))
}

func (est dataEventgridSystemTopicAttributes) Timeouts() dataeventgridsystemtopic.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataeventgridsystemtopic.TimeoutsAttributes](est.ref.Append("timeouts"))
}
