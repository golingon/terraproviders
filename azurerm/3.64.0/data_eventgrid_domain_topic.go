// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataeventgriddomaintopic "github.com/golingon/terraproviders/azurerm/3.64.0/dataeventgriddomaintopic"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEventgridDomainTopic creates a new instance of [DataEventgridDomainTopic].
func NewDataEventgridDomainTopic(name string, args DataEventgridDomainTopicArgs) *DataEventgridDomainTopic {
	return &DataEventgridDomainTopic{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEventgridDomainTopic)(nil)

// DataEventgridDomainTopic represents the Terraform data resource azurerm_eventgrid_domain_topic.
type DataEventgridDomainTopic struct {
	Name string
	Args DataEventgridDomainTopicArgs
}

// DataSource returns the Terraform object type for [DataEventgridDomainTopic].
func (edt *DataEventgridDomainTopic) DataSource() string {
	return "azurerm_eventgrid_domain_topic"
}

// LocalName returns the local name for [DataEventgridDomainTopic].
func (edt *DataEventgridDomainTopic) LocalName() string {
	return edt.Name
}

// Configuration returns the configuration (args) for [DataEventgridDomainTopic].
func (edt *DataEventgridDomainTopic) Configuration() interface{} {
	return edt.Args
}

// Attributes returns the attributes for [DataEventgridDomainTopic].
func (edt *DataEventgridDomainTopic) Attributes() dataEventgridDomainTopicAttributes {
	return dataEventgridDomainTopicAttributes{ref: terra.ReferenceDataResource(edt)}
}

// DataEventgridDomainTopicArgs contains the configurations for azurerm_eventgrid_domain_topic.
type DataEventgridDomainTopicArgs struct {
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataeventgriddomaintopic.Timeouts `hcl:"timeouts,block"`
}
type dataEventgridDomainTopicAttributes struct {
	ref terra.Reference
}

// DomainName returns a reference to field domain_name of azurerm_eventgrid_domain_topic.
func (edt dataEventgridDomainTopicAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(edt.ref.Append("domain_name"))
}

// Id returns a reference to field id of azurerm_eventgrid_domain_topic.
func (edt dataEventgridDomainTopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(edt.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_eventgrid_domain_topic.
func (edt dataEventgridDomainTopicAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(edt.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventgrid_domain_topic.
func (edt dataEventgridDomainTopicAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(edt.ref.Append("resource_group_name"))
}

func (edt dataEventgridDomainTopicAttributes) Timeouts() dataeventgriddomaintopic.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataeventgriddomaintopic.TimeoutsAttributes](edt.ref.Append("timeouts"))
}
