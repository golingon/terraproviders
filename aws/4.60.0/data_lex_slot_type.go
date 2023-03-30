// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datalexslottype "github.com/golingon/terraproviders/aws/4.60.0/datalexslottype"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataLexSlotType(name string, args DataLexSlotTypeArgs) *DataLexSlotType {
	return &DataLexSlotType{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLexSlotType)(nil)

type DataLexSlotType struct {
	Name string
	Args DataLexSlotTypeArgs
}

func (lst *DataLexSlotType) DataSource() string {
	return "aws_lex_slot_type"
}

func (lst *DataLexSlotType) LocalName() string {
	return lst.Name
}

func (lst *DataLexSlotType) Configuration() interface{} {
	return lst.Args
}

func (lst *DataLexSlotType) Attributes() dataLexSlotTypeAttributes {
	return dataLexSlotTypeAttributes{ref: terra.ReferenceDataResource(lst)}
}

type DataLexSlotTypeArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// EnumerationValue: min=0
	EnumerationValue []datalexslottype.EnumerationValue `hcl:"enumeration_value,block" validate:"min=0"`
}
type dataLexSlotTypeAttributes struct {
	ref terra.Reference
}

func (lst dataLexSlotTypeAttributes) Checksum() terra.StringValue {
	return terra.ReferenceString(lst.ref.Append("checksum"))
}

func (lst dataLexSlotTypeAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceString(lst.ref.Append("created_date"))
}

func (lst dataLexSlotTypeAttributes) Description() terra.StringValue {
	return terra.ReferenceString(lst.ref.Append("description"))
}

func (lst dataLexSlotTypeAttributes) Id() terra.StringValue {
	return terra.ReferenceString(lst.ref.Append("id"))
}

func (lst dataLexSlotTypeAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceString(lst.ref.Append("last_updated_date"))
}

func (lst dataLexSlotTypeAttributes) Name() terra.StringValue {
	return terra.ReferenceString(lst.ref.Append("name"))
}

func (lst dataLexSlotTypeAttributes) ValueSelectionStrategy() terra.StringValue {
	return terra.ReferenceString(lst.ref.Append("value_selection_strategy"))
}

func (lst dataLexSlotTypeAttributes) Version() terra.StringValue {
	return terra.ReferenceString(lst.ref.Append("version"))
}

func (lst dataLexSlotTypeAttributes) EnumerationValue() terra.SetValue[datalexslottype.EnumerationValueAttributes] {
	return terra.ReferenceSet[datalexslottype.EnumerationValueAttributes](lst.ref.Append("enumeration_value"))
}
