// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataimage "github.com/golingon/terraproviders/azurerm/3.58.0/dataimage"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataImage creates a new instance of [DataImage].
func NewDataImage(name string, args DataImageArgs) *DataImage {
	return &DataImage{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataImage)(nil)

// DataImage represents the Terraform data resource azurerm_image.
type DataImage struct {
	Name string
	Args DataImageArgs
}

// DataSource returns the Terraform object type for [DataImage].
func (i *DataImage) DataSource() string {
	return "azurerm_image"
}

// LocalName returns the local name for [DataImage].
func (i *DataImage) LocalName() string {
	return i.Name
}

// Configuration returns the configuration (args) for [DataImage].
func (i *DataImage) Configuration() interface{} {
	return i.Args
}

// Attributes returns the attributes for [DataImage].
func (i *DataImage) Attributes() dataImageAttributes {
	return dataImageAttributes{ref: terra.ReferenceDataResource(i)}
}

// DataImageArgs contains the configurations for azurerm_image.
type DataImageArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NameRegex: string, optional
	NameRegex terra.StringValue `hcl:"name_regex,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SortDescending: bool, optional
	SortDescending terra.BoolValue `hcl:"sort_descending,attr"`
	// DataDisk: min=0
	DataDisk []dataimage.DataDisk `hcl:"data_disk,block" validate:"min=0"`
	// OsDisk: min=0
	OsDisk []dataimage.OsDisk `hcl:"os_disk,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataimage.Timeouts `hcl:"timeouts,block"`
}
type dataImageAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_image.
func (i dataImageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_image.
func (i dataImageAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_image.
func (i dataImageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("name"))
}

// NameRegex returns a reference to field name_regex of azurerm_image.
func (i dataImageAttributes) NameRegex() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("name_regex"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_image.
func (i dataImageAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("resource_group_name"))
}

// SortDescending returns a reference to field sort_descending of azurerm_image.
func (i dataImageAttributes) SortDescending() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("sort_descending"))
}

// Tags returns a reference to field tags of azurerm_image.
func (i dataImageAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](i.ref.Append("tags"))
}

// ZoneResilient returns a reference to field zone_resilient of azurerm_image.
func (i dataImageAttributes) ZoneResilient() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("zone_resilient"))
}

func (i dataImageAttributes) DataDisk() terra.ListValue[dataimage.DataDiskAttributes] {
	return terra.ReferenceAsList[dataimage.DataDiskAttributes](i.ref.Append("data_disk"))
}

func (i dataImageAttributes) OsDisk() terra.ListValue[dataimage.OsDiskAttributes] {
	return terra.ReferenceAsList[dataimage.OsDiskAttributes](i.ref.Append("os_disk"))
}

func (i dataImageAttributes) Timeouts() dataimage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataimage.TimeoutsAttributes](i.ref.Append("timeouts"))
}
