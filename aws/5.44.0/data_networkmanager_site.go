// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"github.com/golingon/lingon/pkg/terra"
	datanetworkmanagersite "github.com/golingon/terraproviders/aws/5.44.0/datanetworkmanagersite"
)

// NewDataNetworkmanagerSite creates a new instance of [DataNetworkmanagerSite].
func NewDataNetworkmanagerSite(name string, args DataNetworkmanagerSiteArgs) *DataNetworkmanagerSite {
	return &DataNetworkmanagerSite{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkmanagerSite)(nil)

// DataNetworkmanagerSite represents the Terraform data resource aws_networkmanager_site.
type DataNetworkmanagerSite struct {
	Name string
	Args DataNetworkmanagerSiteArgs
}

// DataSource returns the Terraform object type for [DataNetworkmanagerSite].
func (ns *DataNetworkmanagerSite) DataSource() string {
	return "aws_networkmanager_site"
}

// LocalName returns the local name for [DataNetworkmanagerSite].
func (ns *DataNetworkmanagerSite) LocalName() string {
	return ns.Name
}

// Configuration returns the configuration (args) for [DataNetworkmanagerSite].
func (ns *DataNetworkmanagerSite) Configuration() interface{} {
	return ns.Args
}

// Attributes returns the attributes for [DataNetworkmanagerSite].
func (ns *DataNetworkmanagerSite) Attributes() dataNetworkmanagerSiteAttributes {
	return dataNetworkmanagerSiteAttributes{ref: terra.ReferenceDataResource(ns)}
}

// DataNetworkmanagerSiteArgs contains the configurations for aws_networkmanager_site.
type DataNetworkmanagerSiteArgs struct {
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SiteId: string, required
	SiteId terra.StringValue `hcl:"site_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Location: min=0
	Location []datanetworkmanagersite.Location `hcl:"location,block" validate:"min=0"`
}
type dataNetworkmanagerSiteAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkmanager_site.
func (ns dataNetworkmanagerSiteAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("arn"))
}

// Description returns a reference to field description of aws_networkmanager_site.
func (ns dataNetworkmanagerSiteAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("description"))
}

// GlobalNetworkId returns a reference to field global_network_id of aws_networkmanager_site.
func (ns dataNetworkmanagerSiteAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("global_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_site.
func (ns dataNetworkmanagerSiteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("id"))
}

// SiteId returns a reference to field site_id of aws_networkmanager_site.
func (ns dataNetworkmanagerSiteAttributes) SiteId() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("site_id"))
}

// Tags returns a reference to field tags of aws_networkmanager_site.
func (ns dataNetworkmanagerSiteAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ns.ref.Append("tags"))
}

func (ns dataNetworkmanagerSiteAttributes) Location() terra.ListValue[datanetworkmanagersite.LocationAttributes] {
	return terra.ReferenceAsList[datanetworkmanagersite.LocationAttributes](ns.ref.Append("location"))
}
