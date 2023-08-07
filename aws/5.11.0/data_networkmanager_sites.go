// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataNetworkmanagerSites creates a new instance of [DataNetworkmanagerSites].
func NewDataNetworkmanagerSites(name string, args DataNetworkmanagerSitesArgs) *DataNetworkmanagerSites {
	return &DataNetworkmanagerSites{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkmanagerSites)(nil)

// DataNetworkmanagerSites represents the Terraform data resource aws_networkmanager_sites.
type DataNetworkmanagerSites struct {
	Name string
	Args DataNetworkmanagerSitesArgs
}

// DataSource returns the Terraform object type for [DataNetworkmanagerSites].
func (ns *DataNetworkmanagerSites) DataSource() string {
	return "aws_networkmanager_sites"
}

// LocalName returns the local name for [DataNetworkmanagerSites].
func (ns *DataNetworkmanagerSites) LocalName() string {
	return ns.Name
}

// Configuration returns the configuration (args) for [DataNetworkmanagerSites].
func (ns *DataNetworkmanagerSites) Configuration() interface{} {
	return ns.Args
}

// Attributes returns the attributes for [DataNetworkmanagerSites].
func (ns *DataNetworkmanagerSites) Attributes() dataNetworkmanagerSitesAttributes {
	return dataNetworkmanagerSitesAttributes{ref: terra.ReferenceDataResource(ns)}
}

// DataNetworkmanagerSitesArgs contains the configurations for aws_networkmanager_sites.
type DataNetworkmanagerSitesArgs struct {
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataNetworkmanagerSitesAttributes struct {
	ref terra.Reference
}

// GlobalNetworkId returns a reference to field global_network_id of aws_networkmanager_sites.
func (ns dataNetworkmanagerSitesAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("global_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_sites.
func (ns dataNetworkmanagerSitesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_networkmanager_sites.
func (ns dataNetworkmanagerSitesAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ns.ref.Append("ids"))
}

// Tags returns a reference to field tags of aws_networkmanager_sites.
func (ns dataNetworkmanagerSitesAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ns.ref.Append("tags"))
}
