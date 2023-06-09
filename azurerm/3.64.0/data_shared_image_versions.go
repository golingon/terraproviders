// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasharedimageversions "github.com/golingon/terraproviders/azurerm/3.64.0/datasharedimageversions"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSharedImageVersions creates a new instance of [DataSharedImageVersions].
func NewDataSharedImageVersions(name string, args DataSharedImageVersionsArgs) *DataSharedImageVersions {
	return &DataSharedImageVersions{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSharedImageVersions)(nil)

// DataSharedImageVersions represents the Terraform data resource azurerm_shared_image_versions.
type DataSharedImageVersions struct {
	Name string
	Args DataSharedImageVersionsArgs
}

// DataSource returns the Terraform object type for [DataSharedImageVersions].
func (siv *DataSharedImageVersions) DataSource() string {
	return "azurerm_shared_image_versions"
}

// LocalName returns the local name for [DataSharedImageVersions].
func (siv *DataSharedImageVersions) LocalName() string {
	return siv.Name
}

// Configuration returns the configuration (args) for [DataSharedImageVersions].
func (siv *DataSharedImageVersions) Configuration() interface{} {
	return siv.Args
}

// Attributes returns the attributes for [DataSharedImageVersions].
func (siv *DataSharedImageVersions) Attributes() dataSharedImageVersionsAttributes {
	return dataSharedImageVersionsAttributes{ref: terra.ReferenceDataResource(siv)}
}

// DataSharedImageVersionsArgs contains the configurations for azurerm_shared_image_versions.
type DataSharedImageVersionsArgs struct {
	// GalleryName: string, required
	GalleryName terra.StringValue `hcl:"gallery_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImageName: string, required
	ImageName terra.StringValue `hcl:"image_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// TagsFilter: map of string, optional
	TagsFilter terra.MapValue[terra.StringValue] `hcl:"tags_filter,attr"`
	// Images: min=0
	Images []datasharedimageversions.Images `hcl:"images,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datasharedimageversions.Timeouts `hcl:"timeouts,block"`
}
type dataSharedImageVersionsAttributes struct {
	ref terra.Reference
}

// GalleryName returns a reference to field gallery_name of azurerm_shared_image_versions.
func (siv dataSharedImageVersionsAttributes) GalleryName() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("gallery_name"))
}

// Id returns a reference to field id of azurerm_shared_image_versions.
func (siv dataSharedImageVersionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("id"))
}

// ImageName returns a reference to field image_name of azurerm_shared_image_versions.
func (siv dataSharedImageVersionsAttributes) ImageName() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("image_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_shared_image_versions.
func (siv dataSharedImageVersionsAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("resource_group_name"))
}

// TagsFilter returns a reference to field tags_filter of azurerm_shared_image_versions.
func (siv dataSharedImageVersionsAttributes) TagsFilter() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](siv.ref.Append("tags_filter"))
}

func (siv dataSharedImageVersionsAttributes) Images() terra.ListValue[datasharedimageversions.ImagesAttributes] {
	return terra.ReferenceAsList[datasharedimageversions.ImagesAttributes](siv.ref.Append("images"))
}

func (siv dataSharedImageVersionsAttributes) Timeouts() datasharedimageversions.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasharedimageversions.TimeoutsAttributes](siv.ref.Append("timeouts"))
}
