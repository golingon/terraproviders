// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataNetworkmanagerLinks creates a new instance of [DataNetworkmanagerLinks].
func NewDataNetworkmanagerLinks(name string, args DataNetworkmanagerLinksArgs) *DataNetworkmanagerLinks {
	return &DataNetworkmanagerLinks{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkmanagerLinks)(nil)

// DataNetworkmanagerLinks represents the Terraform data resource aws_networkmanager_links.
type DataNetworkmanagerLinks struct {
	Name string
	Args DataNetworkmanagerLinksArgs
}

// DataSource returns the Terraform object type for [DataNetworkmanagerLinks].
func (nl *DataNetworkmanagerLinks) DataSource() string {
	return "aws_networkmanager_links"
}

// LocalName returns the local name for [DataNetworkmanagerLinks].
func (nl *DataNetworkmanagerLinks) LocalName() string {
	return nl.Name
}

// Configuration returns the configuration (args) for [DataNetworkmanagerLinks].
func (nl *DataNetworkmanagerLinks) Configuration() interface{} {
	return nl.Args
}

// Attributes returns the attributes for [DataNetworkmanagerLinks].
func (nl *DataNetworkmanagerLinks) Attributes() dataNetworkmanagerLinksAttributes {
	return dataNetworkmanagerLinksAttributes{ref: terra.ReferenceDataResource(nl)}
}

// DataNetworkmanagerLinksArgs contains the configurations for aws_networkmanager_links.
type DataNetworkmanagerLinksArgs struct {
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProviderName: string, optional
	ProviderName terra.StringValue `hcl:"provider_name,attr"`
	// SiteId: string, optional
	SiteId terra.StringValue `hcl:"site_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}
type dataNetworkmanagerLinksAttributes struct {
	ref terra.Reference
}

// GlobalNetworkId returns a reference to field global_network_id of aws_networkmanager_links.
func (nl dataNetworkmanagerLinksAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("global_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_links.
func (nl dataNetworkmanagerLinksAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_networkmanager_links.
func (nl dataNetworkmanagerLinksAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nl.ref.Append("ids"))
}

// ProviderName returns a reference to field provider_name of aws_networkmanager_links.
func (nl dataNetworkmanagerLinksAttributes) ProviderName() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("provider_name"))
}

// SiteId returns a reference to field site_id of aws_networkmanager_links.
func (nl dataNetworkmanagerLinksAttributes) SiteId() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("site_id"))
}

// Tags returns a reference to field tags of aws_networkmanager_links.
func (nl dataNetworkmanagerLinksAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nl.ref.Append("tags"))
}

// Type returns a reference to field type of aws_networkmanager_links.
func (nl dataNetworkmanagerLinksAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("type"))
}
