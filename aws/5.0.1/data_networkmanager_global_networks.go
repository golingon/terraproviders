// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataNetworkmanagerGlobalNetworks creates a new instance of [DataNetworkmanagerGlobalNetworks].
func NewDataNetworkmanagerGlobalNetworks(name string, args DataNetworkmanagerGlobalNetworksArgs) *DataNetworkmanagerGlobalNetworks {
	return &DataNetworkmanagerGlobalNetworks{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkmanagerGlobalNetworks)(nil)

// DataNetworkmanagerGlobalNetworks represents the Terraform data resource aws_networkmanager_global_networks.
type DataNetworkmanagerGlobalNetworks struct {
	Name string
	Args DataNetworkmanagerGlobalNetworksArgs
}

// DataSource returns the Terraform object type for [DataNetworkmanagerGlobalNetworks].
func (ngn *DataNetworkmanagerGlobalNetworks) DataSource() string {
	return "aws_networkmanager_global_networks"
}

// LocalName returns the local name for [DataNetworkmanagerGlobalNetworks].
func (ngn *DataNetworkmanagerGlobalNetworks) LocalName() string {
	return ngn.Name
}

// Configuration returns the configuration (args) for [DataNetworkmanagerGlobalNetworks].
func (ngn *DataNetworkmanagerGlobalNetworks) Configuration() interface{} {
	return ngn.Args
}

// Attributes returns the attributes for [DataNetworkmanagerGlobalNetworks].
func (ngn *DataNetworkmanagerGlobalNetworks) Attributes() dataNetworkmanagerGlobalNetworksAttributes {
	return dataNetworkmanagerGlobalNetworksAttributes{ref: terra.ReferenceDataResource(ngn)}
}

// DataNetworkmanagerGlobalNetworksArgs contains the configurations for aws_networkmanager_global_networks.
type DataNetworkmanagerGlobalNetworksArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataNetworkmanagerGlobalNetworksAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_networkmanager_global_networks.
func (ngn dataNetworkmanagerGlobalNetworksAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ngn.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_networkmanager_global_networks.
func (ngn dataNetworkmanagerGlobalNetworksAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ngn.ref.Append("ids"))
}

// Tags returns a reference to field tags of aws_networkmanager_global_networks.
func (ngn dataNetworkmanagerGlobalNetworksAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ngn.ref.Append("tags"))
}
