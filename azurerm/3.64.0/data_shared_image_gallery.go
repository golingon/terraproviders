// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasharedimagegallery "github.com/golingon/terraproviders/azurerm/3.64.0/datasharedimagegallery"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSharedImageGallery creates a new instance of [DataSharedImageGallery].
func NewDataSharedImageGallery(name string, args DataSharedImageGalleryArgs) *DataSharedImageGallery {
	return &DataSharedImageGallery{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSharedImageGallery)(nil)

// DataSharedImageGallery represents the Terraform data resource azurerm_shared_image_gallery.
type DataSharedImageGallery struct {
	Name string
	Args DataSharedImageGalleryArgs
}

// DataSource returns the Terraform object type for [DataSharedImageGallery].
func (sig *DataSharedImageGallery) DataSource() string {
	return "azurerm_shared_image_gallery"
}

// LocalName returns the local name for [DataSharedImageGallery].
func (sig *DataSharedImageGallery) LocalName() string {
	return sig.Name
}

// Configuration returns the configuration (args) for [DataSharedImageGallery].
func (sig *DataSharedImageGallery) Configuration() interface{} {
	return sig.Args
}

// Attributes returns the attributes for [DataSharedImageGallery].
func (sig *DataSharedImageGallery) Attributes() dataSharedImageGalleryAttributes {
	return dataSharedImageGalleryAttributes{ref: terra.ReferenceDataResource(sig)}
}

// DataSharedImageGalleryArgs contains the configurations for azurerm_shared_image_gallery.
type DataSharedImageGalleryArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datasharedimagegallery.Timeouts `hcl:"timeouts,block"`
}
type dataSharedImageGalleryAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_shared_image_gallery.
func (sig dataSharedImageGalleryAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sig.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_shared_image_gallery.
func (sig dataSharedImageGalleryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sig.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_shared_image_gallery.
func (sig dataSharedImageGalleryAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sig.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_shared_image_gallery.
func (sig dataSharedImageGalleryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sig.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_shared_image_gallery.
func (sig dataSharedImageGalleryAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sig.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_shared_image_gallery.
func (sig dataSharedImageGalleryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sig.ref.Append("tags"))
}

// UniqueName returns a reference to field unique_name of azurerm_shared_image_gallery.
func (sig dataSharedImageGalleryAttributes) UniqueName() terra.StringValue {
	return terra.ReferenceAsString(sig.ref.Append("unique_name"))
}

func (sig dataSharedImageGalleryAttributes) Timeouts() datasharedimagegallery.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasharedimagegallery.TimeoutsAttributes](sig.ref.Append("timeouts"))
}