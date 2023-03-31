// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasharedimage "github.com/golingon/terraproviders/azurerm/3.49.0/datasharedimage"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataSharedImage(name string, args DataSharedImageArgs) *DataSharedImage {
	return &DataSharedImage{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSharedImage)(nil)

type DataSharedImage struct {
	Name string
	Args DataSharedImageArgs
}

func (si *DataSharedImage) DataSource() string {
	return "azurerm_shared_image"
}

func (si *DataSharedImage) LocalName() string {
	return si.Name
}

func (si *DataSharedImage) Configuration() interface{} {
	return si.Args
}

func (si *DataSharedImage) Attributes() dataSharedImageAttributes {
	return dataSharedImageAttributes{ref: terra.ReferenceDataResource(si)}
}

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

func (si dataSharedImageAttributes) Architecture() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("architecture"))
}

func (si dataSharedImageAttributes) Description() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("description"))
}

func (si dataSharedImageAttributes) Eula() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("eula"))
}

func (si dataSharedImageAttributes) GalleryName() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("gallery_name"))
}

func (si dataSharedImageAttributes) HyperVGeneration() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("hyper_v_generation"))
}

func (si dataSharedImageAttributes) Id() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("id"))
}

func (si dataSharedImageAttributes) Location() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("location"))
}

func (si dataSharedImageAttributes) Name() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("name"))
}

func (si dataSharedImageAttributes) OsType() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("os_type"))
}

func (si dataSharedImageAttributes) PrivacyStatementUri() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("privacy_statement_uri"))
}

func (si dataSharedImageAttributes) ReleaseNoteUri() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("release_note_uri"))
}

func (si dataSharedImageAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("resource_group_name"))
}

func (si dataSharedImageAttributes) Specialized() terra.BoolValue {
	return terra.ReferenceBool(si.ref.Append("specialized"))
}

func (si dataSharedImageAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](si.ref.Append("tags"))
}

func (si dataSharedImageAttributes) Identifier() terra.ListValue[datasharedimage.IdentifierAttributes] {
	return terra.ReferenceList[datasharedimage.IdentifierAttributes](si.ref.Append("identifier"))
}

func (si dataSharedImageAttributes) PurchasePlan() terra.ListValue[datasharedimage.PurchasePlanAttributes] {
	return terra.ReferenceList[datasharedimage.PurchasePlanAttributes](si.ref.Append("purchase_plan"))
}

func (si dataSharedImageAttributes) Timeouts() datasharedimage.TimeoutsAttributes {
	return terra.ReferenceSingle[datasharedimage.TimeoutsAttributes](si.ref.Append("timeouts"))
}
