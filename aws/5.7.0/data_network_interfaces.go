// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datanetworkinterfaces "github.com/golingon/terraproviders/aws/5.7.0/datanetworkinterfaces"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataNetworkInterfaces creates a new instance of [DataNetworkInterfaces].
func NewDataNetworkInterfaces(name string, args DataNetworkInterfacesArgs) *DataNetworkInterfaces {
	return &DataNetworkInterfaces{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkInterfaces)(nil)

// DataNetworkInterfaces represents the Terraform data resource aws_network_interfaces.
type DataNetworkInterfaces struct {
	Name string
	Args DataNetworkInterfacesArgs
}

// DataSource returns the Terraform object type for [DataNetworkInterfaces].
func (ni *DataNetworkInterfaces) DataSource() string {
	return "aws_network_interfaces"
}

// LocalName returns the local name for [DataNetworkInterfaces].
func (ni *DataNetworkInterfaces) LocalName() string {
	return ni.Name
}

// Configuration returns the configuration (args) for [DataNetworkInterfaces].
func (ni *DataNetworkInterfaces) Configuration() interface{} {
	return ni.Args
}

// Attributes returns the attributes for [DataNetworkInterfaces].
func (ni *DataNetworkInterfaces) Attributes() dataNetworkInterfacesAttributes {
	return dataNetworkInterfacesAttributes{ref: terra.ReferenceDataResource(ni)}
}

// DataNetworkInterfacesArgs contains the configurations for aws_network_interfaces.
type DataNetworkInterfacesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []datanetworkinterfaces.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datanetworkinterfaces.Timeouts `hcl:"timeouts,block"`
}
type dataNetworkInterfacesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_network_interfaces.
func (ni dataNetworkInterfacesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_network_interfaces.
func (ni dataNetworkInterfacesAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ni.ref.Append("ids"))
}

// Tags returns a reference to field tags of aws_network_interfaces.
func (ni dataNetworkInterfacesAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ni.ref.Append("tags"))
}

func (ni dataNetworkInterfacesAttributes) Filter() terra.SetValue[datanetworkinterfaces.FilterAttributes] {
	return terra.ReferenceAsSet[datanetworkinterfaces.FilterAttributes](ni.ref.Append("filter"))
}

func (ni dataNetworkInterfacesAttributes) Timeouts() datanetworkinterfaces.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datanetworkinterfaces.TimeoutsAttributes](ni.ref.Append("timeouts"))
}
