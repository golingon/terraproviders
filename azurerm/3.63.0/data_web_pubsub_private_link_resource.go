// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datawebpubsubprivatelinkresource "github.com/golingon/terraproviders/azurerm/3.63.0/datawebpubsubprivatelinkresource"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataWebPubsubPrivateLinkResource creates a new instance of [DataWebPubsubPrivateLinkResource].
func NewDataWebPubsubPrivateLinkResource(name string, args DataWebPubsubPrivateLinkResourceArgs) *DataWebPubsubPrivateLinkResource {
	return &DataWebPubsubPrivateLinkResource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataWebPubsubPrivateLinkResource)(nil)

// DataWebPubsubPrivateLinkResource represents the Terraform data resource azurerm_web_pubsub_private_link_resource.
type DataWebPubsubPrivateLinkResource struct {
	Name string
	Args DataWebPubsubPrivateLinkResourceArgs
}

// DataSource returns the Terraform object type for [DataWebPubsubPrivateLinkResource].
func (wpplr *DataWebPubsubPrivateLinkResource) DataSource() string {
	return "azurerm_web_pubsub_private_link_resource"
}

// LocalName returns the local name for [DataWebPubsubPrivateLinkResource].
func (wpplr *DataWebPubsubPrivateLinkResource) LocalName() string {
	return wpplr.Name
}

// Configuration returns the configuration (args) for [DataWebPubsubPrivateLinkResource].
func (wpplr *DataWebPubsubPrivateLinkResource) Configuration() interface{} {
	return wpplr.Args
}

// Attributes returns the attributes for [DataWebPubsubPrivateLinkResource].
func (wpplr *DataWebPubsubPrivateLinkResource) Attributes() dataWebPubsubPrivateLinkResourceAttributes {
	return dataWebPubsubPrivateLinkResourceAttributes{ref: terra.ReferenceDataResource(wpplr)}
}

// DataWebPubsubPrivateLinkResourceArgs contains the configurations for azurerm_web_pubsub_private_link_resource.
type DataWebPubsubPrivateLinkResourceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// WebPubsubId: string, required
	WebPubsubId terra.StringValue `hcl:"web_pubsub_id,attr" validate:"required"`
	// SharedPrivateLinkResourceTypes: min=0
	SharedPrivateLinkResourceTypes []datawebpubsubprivatelinkresource.SharedPrivateLinkResourceTypes `hcl:"shared_private_link_resource_types,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datawebpubsubprivatelinkresource.Timeouts `hcl:"timeouts,block"`
}
type dataWebPubsubPrivateLinkResourceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_web_pubsub_private_link_resource.
func (wpplr dataWebPubsubPrivateLinkResourceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wpplr.ref.Append("id"))
}

// WebPubsubId returns a reference to field web_pubsub_id of azurerm_web_pubsub_private_link_resource.
func (wpplr dataWebPubsubPrivateLinkResourceAttributes) WebPubsubId() terra.StringValue {
	return terra.ReferenceAsString(wpplr.ref.Append("web_pubsub_id"))
}

func (wpplr dataWebPubsubPrivateLinkResourceAttributes) SharedPrivateLinkResourceTypes() terra.ListValue[datawebpubsubprivatelinkresource.SharedPrivateLinkResourceTypesAttributes] {
	return terra.ReferenceAsList[datawebpubsubprivatelinkresource.SharedPrivateLinkResourceTypesAttributes](wpplr.ref.Append("shared_private_link_resource_types"))
}

func (wpplr dataWebPubsubPrivateLinkResourceAttributes) Timeouts() datawebpubsubprivatelinkresource.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datawebpubsubprivatelinkresource.TimeoutsAttributes](wpplr.ref.Append("timeouts"))
}
