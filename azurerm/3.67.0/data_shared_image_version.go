// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasharedimageversion "github.com/golingon/terraproviders/azurerm/3.67.0/datasharedimageversion"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSharedImageVersion creates a new instance of [DataSharedImageVersion].
func NewDataSharedImageVersion(name string, args DataSharedImageVersionArgs) *DataSharedImageVersion {
	return &DataSharedImageVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSharedImageVersion)(nil)

// DataSharedImageVersion represents the Terraform data resource azurerm_shared_image_version.
type DataSharedImageVersion struct {
	Name string
	Args DataSharedImageVersionArgs
}

// DataSource returns the Terraform object type for [DataSharedImageVersion].
func (siv *DataSharedImageVersion) DataSource() string {
	return "azurerm_shared_image_version"
}

// LocalName returns the local name for [DataSharedImageVersion].
func (siv *DataSharedImageVersion) LocalName() string {
	return siv.Name
}

// Configuration returns the configuration (args) for [DataSharedImageVersion].
func (siv *DataSharedImageVersion) Configuration() interface{} {
	return siv.Args
}

// Attributes returns the attributes for [DataSharedImageVersion].
func (siv *DataSharedImageVersion) Attributes() dataSharedImageVersionAttributes {
	return dataSharedImageVersionAttributes{ref: terra.ReferenceDataResource(siv)}
}

// DataSharedImageVersionArgs contains the configurations for azurerm_shared_image_version.
type DataSharedImageVersionArgs struct {
	// GalleryName: string, required
	GalleryName terra.StringValue `hcl:"gallery_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImageName: string, required
	ImageName terra.StringValue `hcl:"image_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SortVersionsBySemver: bool, optional
	SortVersionsBySemver terra.BoolValue `hcl:"sort_versions_by_semver,attr"`
	// TargetRegion: min=0
	TargetRegion []datasharedimageversion.TargetRegion `hcl:"target_region,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datasharedimageversion.Timeouts `hcl:"timeouts,block"`
}
type dataSharedImageVersionAttributes struct {
	ref terra.Reference
}

// ExcludeFromLatest returns a reference to field exclude_from_latest of azurerm_shared_image_version.
func (siv dataSharedImageVersionAttributes) ExcludeFromLatest() terra.BoolValue {
	return terra.ReferenceAsBool(siv.ref.Append("exclude_from_latest"))
}

// GalleryName returns a reference to field gallery_name of azurerm_shared_image_version.
func (siv dataSharedImageVersionAttributes) GalleryName() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("gallery_name"))
}

// Id returns a reference to field id of azurerm_shared_image_version.
func (siv dataSharedImageVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("id"))
}

// ImageName returns a reference to field image_name of azurerm_shared_image_version.
func (siv dataSharedImageVersionAttributes) ImageName() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("image_name"))
}

// Location returns a reference to field location of azurerm_shared_image_version.
func (siv dataSharedImageVersionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("location"))
}

// ManagedImageId returns a reference to field managed_image_id of azurerm_shared_image_version.
func (siv dataSharedImageVersionAttributes) ManagedImageId() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("managed_image_id"))
}

// Name returns a reference to field name of azurerm_shared_image_version.
func (siv dataSharedImageVersionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("name"))
}

// OsDiskImageSizeGb returns a reference to field os_disk_image_size_gb of azurerm_shared_image_version.
func (siv dataSharedImageVersionAttributes) OsDiskImageSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(siv.ref.Append("os_disk_image_size_gb"))
}

// OsDiskSnapshotId returns a reference to field os_disk_snapshot_id of azurerm_shared_image_version.
func (siv dataSharedImageVersionAttributes) OsDiskSnapshotId() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("os_disk_snapshot_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_shared_image_version.
func (siv dataSharedImageVersionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("resource_group_name"))
}

// SortVersionsBySemver returns a reference to field sort_versions_by_semver of azurerm_shared_image_version.
func (siv dataSharedImageVersionAttributes) SortVersionsBySemver() terra.BoolValue {
	return terra.ReferenceAsBool(siv.ref.Append("sort_versions_by_semver"))
}

// Tags returns a reference to field tags of azurerm_shared_image_version.
func (siv dataSharedImageVersionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](siv.ref.Append("tags"))
}

func (siv dataSharedImageVersionAttributes) TargetRegion() terra.ListValue[datasharedimageversion.TargetRegionAttributes] {
	return terra.ReferenceAsList[datasharedimageversion.TargetRegionAttributes](siv.ref.Append("target_region"))
}

func (siv dataSharedImageVersionAttributes) Timeouts() datasharedimageversion.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasharedimageversion.TimeoutsAttributes](siv.ref.Append("timeouts"))
}
