// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	dataimages "github.com/golingon/terraproviders/azurerm/3.98.0/dataimages"
)

// NewDataImages creates a new instance of [DataImages].
func NewDataImages(name string, args DataImagesArgs) *DataImages {
	return &DataImages{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataImages)(nil)

// DataImages represents the Terraform data resource azurerm_images.
type DataImages struct {
	Name string
	Args DataImagesArgs
}

// DataSource returns the Terraform object type for [DataImages].
func (i *DataImages) DataSource() string {
	return "azurerm_images"
}

// LocalName returns the local name for [DataImages].
func (i *DataImages) LocalName() string {
	return i.Name
}

// Configuration returns the configuration (args) for [DataImages].
func (i *DataImages) Configuration() interface{} {
	return i.Args
}

// Attributes returns the attributes for [DataImages].
func (i *DataImages) Attributes() dataImagesAttributes {
	return dataImagesAttributes{ref: terra.ReferenceDataResource(i)}
}

// DataImagesArgs contains the configurations for azurerm_images.
type DataImagesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// TagsFilter: map of string, optional
	TagsFilter terra.MapValue[terra.StringValue] `hcl:"tags_filter,attr"`
	// Images: min=0
	Images []dataimages.Images `hcl:"images,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataimages.Timeouts `hcl:"timeouts,block"`
}
type dataImagesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_images.
func (i dataImagesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_images.
func (i dataImagesAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("resource_group_name"))
}

// TagsFilter returns a reference to field tags_filter of azurerm_images.
func (i dataImagesAttributes) TagsFilter() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](i.ref.Append("tags_filter"))
}

func (i dataImagesAttributes) Images() terra.ListValue[dataimages.ImagesAttributes] {
	return terra.ReferenceAsList[dataimages.ImagesAttributes](i.ref.Append("images"))
}

func (i dataImagesAttributes) Timeouts() dataimages.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataimages.TimeoutsAttributes](i.ref.Append("timeouts"))
}
