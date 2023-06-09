// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasharedimage "github.com/golingon/terraproviders/azurerm/3.63.0/datasharedimage"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSharedImage creates a new instance of [DataSharedImage].
func NewDataSharedImage(name string, args DataSharedImageArgs) *DataSharedImage {
	return &DataSharedImage{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSharedImage)(nil)

// DataSharedImage represents the Terraform data resource azurerm_shared_image.
type DataSharedImage struct {
	Name string
	Args DataSharedImageArgs
}

// DataSource returns the Terraform object type for [DataSharedImage].
func (si *DataSharedImage) DataSource() string {
	return "azurerm_shared_image"
}

// LocalName returns the local name for [DataSharedImage].
func (si *DataSharedImage) LocalName() string {
	return si.Name
}

// Configuration returns the configuration (args) for [DataSharedImage].
func (si *DataSharedImage) Configuration() interface{} {
	return si.Args
}

// Attributes returns the attributes for [DataSharedImage].
func (si *DataSharedImage) Attributes() dataSharedImageAttributes {
	return dataSharedImageAttributes{ref: terra.ReferenceDataResource(si)}
}

// DataSharedImageArgs contains the configurations for azurerm_shared_image.
type DataSharedImageArgs struct {
	// GalleryName: string, required
	GalleryName terra.StringValue `hcl:"gallery_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identifier: min=0
	Identifier []datasharedimage.Identifier `hcl:"identifier,block" validate:"min=0"`
	// PurchasePlan: min=0
	PurchasePlan []datasharedimage.PurchasePlan `hcl:"purchase_plan,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datasharedimage.Timeouts `hcl:"timeouts,block"`
}
type dataSharedImageAttributes struct {
	ref terra.Reference
}

// Architecture returns a reference to field architecture of azurerm_shared_image.
func (si dataSharedImageAttributes) Architecture() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("architecture"))
}

// Description returns a reference to field description of azurerm_shared_image.
func (si dataSharedImageAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("description"))
}

// Eula returns a reference to field eula of azurerm_shared_image.
func (si dataSharedImageAttributes) Eula() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("eula"))
}

// GalleryName returns a reference to field gallery_name of azurerm_shared_image.
func (si dataSharedImageAttributes) GalleryName() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("gallery_name"))
}

// HyperVGeneration returns a reference to field hyper_v_generation of azurerm_shared_image.
func (si dataSharedImageAttributes) HyperVGeneration() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("hyper_v_generation"))
}

// Id returns a reference to field id of azurerm_shared_image.
func (si dataSharedImageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_shared_image.
func (si dataSharedImageAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_shared_image.
func (si dataSharedImageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("name"))
}

// OsType returns a reference to field os_type of azurerm_shared_image.
func (si dataSharedImageAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("os_type"))
}

// PrivacyStatementUri returns a reference to field privacy_statement_uri of azurerm_shared_image.
func (si dataSharedImageAttributes) PrivacyStatementUri() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("privacy_statement_uri"))
}

// ReleaseNoteUri returns a reference to field release_note_uri of azurerm_shared_image.
func (si dataSharedImageAttributes) ReleaseNoteUri() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("release_note_uri"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_shared_image.
func (si dataSharedImageAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("resource_group_name"))
}

// Specialized returns a reference to field specialized of azurerm_shared_image.
func (si dataSharedImageAttributes) Specialized() terra.BoolValue {
	return terra.ReferenceAsBool(si.ref.Append("specialized"))
}

// Tags returns a reference to field tags of azurerm_shared_image.
func (si dataSharedImageAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](si.ref.Append("tags"))
}

func (si dataSharedImageAttributes) Identifier() terra.ListValue[datasharedimage.IdentifierAttributes] {
	return terra.ReferenceAsList[datasharedimage.IdentifierAttributes](si.ref.Append("identifier"))
}

func (si dataSharedImageAttributes) PurchasePlan() terra.ListValue[datasharedimage.PurchasePlanAttributes] {
	return terra.ReferenceAsList[datasharedimage.PurchasePlanAttributes](si.ref.Append("purchase_plan"))
}

func (si dataSharedImageAttributes) Timeouts() datasharedimage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasharedimage.TimeoutsAttributes](si.ref.Append("timeouts"))
}
