// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"github.com/golingon/lingon/pkg/terra"
	datanetworkmanagerlink "github.com/golingon/terraproviders/aws/5.44.0/datanetworkmanagerlink"
)

// NewDataNetworkmanagerLink creates a new instance of [DataNetworkmanagerLink].
func NewDataNetworkmanagerLink(name string, args DataNetworkmanagerLinkArgs) *DataNetworkmanagerLink {
	return &DataNetworkmanagerLink{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkmanagerLink)(nil)

// DataNetworkmanagerLink represents the Terraform data resource aws_networkmanager_link.
type DataNetworkmanagerLink struct {
	Name string
	Args DataNetworkmanagerLinkArgs
}

// DataSource returns the Terraform object type for [DataNetworkmanagerLink].
func (nl *DataNetworkmanagerLink) DataSource() string {
	return "aws_networkmanager_link"
}

// LocalName returns the local name for [DataNetworkmanagerLink].
func (nl *DataNetworkmanagerLink) LocalName() string {
	return nl.Name
}

// Configuration returns the configuration (args) for [DataNetworkmanagerLink].
func (nl *DataNetworkmanagerLink) Configuration() interface{} {
	return nl.Args
}

// Attributes returns the attributes for [DataNetworkmanagerLink].
func (nl *DataNetworkmanagerLink) Attributes() dataNetworkmanagerLinkAttributes {
	return dataNetworkmanagerLinkAttributes{ref: terra.ReferenceDataResource(nl)}
}

// DataNetworkmanagerLinkArgs contains the configurations for aws_networkmanager_link.
type DataNetworkmanagerLinkArgs struct {
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkId: string, required
	LinkId terra.StringValue `hcl:"link_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Bandwidth: min=0
	Bandwidth []datanetworkmanagerlink.Bandwidth `hcl:"bandwidth,block" validate:"min=0"`
}
type dataNetworkmanagerLinkAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkmanager_link.
func (nl dataNetworkmanagerLinkAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("arn"))
}

// Description returns a reference to field description of aws_networkmanager_link.
func (nl dataNetworkmanagerLinkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("description"))
}

// GlobalNetworkId returns a reference to field global_network_id of aws_networkmanager_link.
func (nl dataNetworkmanagerLinkAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("global_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_link.
func (nl dataNetworkmanagerLinkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("id"))
}

// LinkId returns a reference to field link_id of aws_networkmanager_link.
func (nl dataNetworkmanagerLinkAttributes) LinkId() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("link_id"))
}

// ProviderName returns a reference to field provider_name of aws_networkmanager_link.
func (nl dataNetworkmanagerLinkAttributes) ProviderName() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("provider_name"))
}

// SiteId returns a reference to field site_id of aws_networkmanager_link.
func (nl dataNetworkmanagerLinkAttributes) SiteId() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("site_id"))
}

// Tags returns a reference to field tags of aws_networkmanager_link.
func (nl dataNetworkmanagerLinkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nl.ref.Append("tags"))
}

// Type returns a reference to field type of aws_networkmanager_link.
func (nl dataNetworkmanagerLinkAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("type"))
}

func (nl dataNetworkmanagerLinkAttributes) Bandwidth() terra.ListValue[datanetworkmanagerlink.BandwidthAttributes] {
	return terra.ReferenceAsList[datanetworkmanagerlink.BandwidthAttributes](nl.ref.Append("bandwidth"))
}
