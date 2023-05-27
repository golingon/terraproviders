// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataNetworkmanagerGlobalNetwork creates a new instance of [DataNetworkmanagerGlobalNetwork].
func NewDataNetworkmanagerGlobalNetwork(name string, args DataNetworkmanagerGlobalNetworkArgs) *DataNetworkmanagerGlobalNetwork {
	return &DataNetworkmanagerGlobalNetwork{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkmanagerGlobalNetwork)(nil)

// DataNetworkmanagerGlobalNetwork represents the Terraform data resource aws_networkmanager_global_network.
type DataNetworkmanagerGlobalNetwork struct {
	Name string
	Args DataNetworkmanagerGlobalNetworkArgs
}

// DataSource returns the Terraform object type for [DataNetworkmanagerGlobalNetwork].
func (ngn *DataNetworkmanagerGlobalNetwork) DataSource() string {
	return "aws_networkmanager_global_network"
}

// LocalName returns the local name for [DataNetworkmanagerGlobalNetwork].
func (ngn *DataNetworkmanagerGlobalNetwork) LocalName() string {
	return ngn.Name
}

// Configuration returns the configuration (args) for [DataNetworkmanagerGlobalNetwork].
func (ngn *DataNetworkmanagerGlobalNetwork) Configuration() interface{} {
	return ngn.Args
}

// Attributes returns the attributes for [DataNetworkmanagerGlobalNetwork].
func (ngn *DataNetworkmanagerGlobalNetwork) Attributes() dataNetworkmanagerGlobalNetworkAttributes {
	return dataNetworkmanagerGlobalNetworkAttributes{ref: terra.ReferenceDataResource(ngn)}
}

// DataNetworkmanagerGlobalNetworkArgs contains the configurations for aws_networkmanager_global_network.
type DataNetworkmanagerGlobalNetworkArgs struct {
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataNetworkmanagerGlobalNetworkAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkmanager_global_network.
func (ngn dataNetworkmanagerGlobalNetworkAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ngn.ref.Append("arn"))
}

// Description returns a reference to field description of aws_networkmanager_global_network.
func (ngn dataNetworkmanagerGlobalNetworkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ngn.ref.Append("description"))
}

// GlobalNetworkId returns a reference to field global_network_id of aws_networkmanager_global_network.
func (ngn dataNetworkmanagerGlobalNetworkAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceAsString(ngn.ref.Append("global_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_global_network.
func (ngn dataNetworkmanagerGlobalNetworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ngn.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_networkmanager_global_network.
func (ngn dataNetworkmanagerGlobalNetworkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ngn.ref.Append("tags"))
}
