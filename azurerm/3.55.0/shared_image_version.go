// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sharedimageversion "github.com/golingon/terraproviders/azurerm/3.55.0/sharedimageversion"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSharedImageVersion creates a new instance of [SharedImageVersion].
func NewSharedImageVersion(name string, args SharedImageVersionArgs) *SharedImageVersion {
	return &SharedImageVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SharedImageVersion)(nil)

// SharedImageVersion represents the Terraform resource azurerm_shared_image_version.
type SharedImageVersion struct {
	Name      string
	Args      SharedImageVersionArgs
	state     *sharedImageVersionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SharedImageVersion].
func (siv *SharedImageVersion) Type() string {
	return "azurerm_shared_image_version"
}

// LocalName returns the local name for [SharedImageVersion].
func (siv *SharedImageVersion) LocalName() string {
	return siv.Name
}

// Configuration returns the configuration (args) for [SharedImageVersion].
func (siv *SharedImageVersion) Configuration() interface{} {
	return siv.Args
}

// DependOn is used for other resources to depend on [SharedImageVersion].
func (siv *SharedImageVersion) DependOn() terra.Reference {
	return terra.ReferenceResource(siv)
}

// Dependencies returns the list of resources [SharedImageVersion] depends_on.
func (siv *SharedImageVersion) Dependencies() terra.Dependencies {
	return siv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SharedImageVersion].
func (siv *SharedImageVersion) LifecycleManagement() *terra.Lifecycle {
	return siv.Lifecycle
}

// Attributes returns the attributes for [SharedImageVersion].
func (siv *SharedImageVersion) Attributes() sharedImageVersionAttributes {
	return sharedImageVersionAttributes{ref: terra.ReferenceResource(siv)}
}

// ImportState imports the given attribute values into [SharedImageVersion]'s state.
func (siv *SharedImageVersion) ImportState(av io.Reader) error {
	siv.state = &sharedImageVersionState{}
	if err := json.NewDecoder(av).Decode(siv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", siv.Type(), siv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SharedImageVersion] has state.
func (siv *SharedImageVersion) State() (*sharedImageVersionState, bool) {
	return siv.state, siv.state != nil
}

// StateMust returns the state for [SharedImageVersion]. Panics if the state is nil.
func (siv *SharedImageVersion) StateMust() *sharedImageVersionState {
	if siv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", siv.Type(), siv.LocalName()))
	}
	return siv.state
}

// SharedImageVersionArgs contains the configurations for azurerm_shared_image_version.
type SharedImageVersionArgs struct {
	// BlobUri: string, optional
	BlobUri terra.StringValue `hcl:"blob_uri,attr"`
	// EndOfLifeDate: string, optional
	EndOfLifeDate terra.StringValue `hcl:"end_of_life_date,attr"`
	// ExcludeFromLatest: bool, optional
	ExcludeFromLatest terra.BoolValue `hcl:"exclude_from_latest,attr"`
	// GalleryName: string, required
	GalleryName terra.StringValue `hcl:"gallery_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImageName: string, required
	ImageName terra.StringValue `hcl:"image_name,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ManagedImageId: string, optional
	ManagedImageId terra.StringValue `hcl:"managed_image_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OsDiskSnapshotId: string, optional
	OsDiskSnapshotId terra.StringValue `hcl:"os_disk_snapshot_id,attr"`
	// ReplicationMode: string, optional
	ReplicationMode terra.StringValue `hcl:"replication_mode,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StorageAccountId: string, optional
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TargetRegion: min=1
	TargetRegion []sharedimageversion.TargetRegion `hcl:"target_region,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *sharedimageversion.Timeouts `hcl:"timeouts,block"`
}
type sharedImageVersionAttributes struct {
	ref terra.Reference
}

// BlobUri returns a reference to field blob_uri of azurerm_shared_image_version.
func (siv sharedImageVersionAttributes) BlobUri() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("blob_uri"))
}

// EndOfLifeDate returns a reference to field end_of_life_date of azurerm_shared_image_version.
func (siv sharedImageVersionAttributes) EndOfLifeDate() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("end_of_life_date"))
}

// ExcludeFromLatest returns a reference to field exclude_from_latest of azurerm_shared_image_version.
func (siv sharedImageVersionAttributes) ExcludeFromLatest() terra.BoolValue {
	return terra.ReferenceAsBool(siv.ref.Append("exclude_from_latest"))
}

// GalleryName returns a reference to field gallery_name of azurerm_shared_image_version.
func (siv sharedImageVersionAttributes) GalleryName() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("gallery_name"))
}

// Id returns a reference to field id of azurerm_shared_image_version.
func (siv sharedImageVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("id"))
}

// ImageName returns a reference to field image_name of azurerm_shared_image_version.
func (siv sharedImageVersionAttributes) ImageName() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("image_name"))
}

// Location returns a reference to field location of azurerm_shared_image_version.
func (siv sharedImageVersionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("location"))
}

// ManagedImageId returns a reference to field managed_image_id of azurerm_shared_image_version.
func (siv sharedImageVersionAttributes) ManagedImageId() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("managed_image_id"))
}

// Name returns a reference to field name of azurerm_shared_image_version.
func (siv sharedImageVersionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("name"))
}

// OsDiskSnapshotId returns a reference to field os_disk_snapshot_id of azurerm_shared_image_version.
func (siv sharedImageVersionAttributes) OsDiskSnapshotId() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("os_disk_snapshot_id"))
}

// ReplicationMode returns a reference to field replication_mode of azurerm_shared_image_version.
func (siv sharedImageVersionAttributes) ReplicationMode() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("replication_mode"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_shared_image_version.
func (siv sharedImageVersionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("resource_group_name"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_shared_image_version.
func (siv sharedImageVersionAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(siv.ref.Append("storage_account_id"))
}

// Tags returns a reference to field tags of azurerm_shared_image_version.
func (siv sharedImageVersionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](siv.ref.Append("tags"))
}

func (siv sharedImageVersionAttributes) TargetRegion() terra.ListValue[sharedimageversion.TargetRegionAttributes] {
	return terra.ReferenceAsList[sharedimageversion.TargetRegionAttributes](siv.ref.Append("target_region"))
}

func (siv sharedImageVersionAttributes) Timeouts() sharedimageversion.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sharedimageversion.TimeoutsAttributes](siv.ref.Append("timeouts"))
}

type sharedImageVersionState struct {
	BlobUri           string                                 `json:"blob_uri"`
	EndOfLifeDate     string                                 `json:"end_of_life_date"`
	ExcludeFromLatest bool                                   `json:"exclude_from_latest"`
	GalleryName       string                                 `json:"gallery_name"`
	Id                string                                 `json:"id"`
	ImageName         string                                 `json:"image_name"`
	Location          string                                 `json:"location"`
	ManagedImageId    string                                 `json:"managed_image_id"`
	Name              string                                 `json:"name"`
	OsDiskSnapshotId  string                                 `json:"os_disk_snapshot_id"`
	ReplicationMode   string                                 `json:"replication_mode"`
	ResourceGroupName string                                 `json:"resource_group_name"`
	StorageAccountId  string                                 `json:"storage_account_id"`
	Tags              map[string]string                      `json:"tags"`
	TargetRegion      []sharedimageversion.TargetRegionState `json:"target_region"`
	Timeouts          *sharedimageversion.TimeoutsState      `json:"timeouts"`
}
