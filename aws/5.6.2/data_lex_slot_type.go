// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datalexslottype "github.com/golingon/terraproviders/aws/5.6.2/datalexslottype"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLexSlotType creates a new instance of [DataLexSlotType].
func NewDataLexSlotType(name string, args DataLexSlotTypeArgs) *DataLexSlotType {
	return &DataLexSlotType{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLexSlotType)(nil)

// DataLexSlotType represents the Terraform data resource aws_lex_slot_type.
type DataLexSlotType struct {
	Name string
	Args DataLexSlotTypeArgs
}

// DataSource returns the Terraform object type for [DataLexSlotType].
func (lst *DataLexSlotType) DataSource() string {
	return "aws_lex_slot_type"
}

// LocalName returns the local name for [DataLexSlotType].
func (lst *DataLexSlotType) LocalName() string {
	return lst.Name
}

// Configuration returns the configuration (args) for [DataLexSlotType].
func (lst *DataLexSlotType) Configuration() interface{} {
	return lst.Args
}

// Attributes returns the attributes for [DataLexSlotType].
func (lst *DataLexSlotType) Attributes() dataLexSlotTypeAttributes {
	return dataLexSlotTypeAttributes{ref: terra.ReferenceDataResource(lst)}
}

// DataLexSlotTypeArgs contains the configurations for aws_lex_slot_type.
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

// Checksum returns a reference to field checksum of aws_lex_slot_type.
func (lst dataLexSlotTypeAttributes) Checksum() terra.StringValue {
	return terra.ReferenceAsString(lst.ref.Append("checksum"))
}

// CreatedDate returns a reference to field created_date of aws_lex_slot_type.
func (lst dataLexSlotTypeAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(lst.ref.Append("created_date"))
}

// Description returns a reference to field description of aws_lex_slot_type.
func (lst dataLexSlotTypeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lst.ref.Append("description"))
}

// Id returns a reference to field id of aws_lex_slot_type.
func (lst dataLexSlotTypeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lst.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_lex_slot_type.
func (lst dataLexSlotTypeAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(lst.ref.Append("last_updated_date"))
}

// Name returns a reference to field name of aws_lex_slot_type.
func (lst dataLexSlotTypeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lst.ref.Append("name"))
}

// ValueSelectionStrategy returns a reference to field value_selection_strategy of aws_lex_slot_type.
func (lst dataLexSlotTypeAttributes) ValueSelectionStrategy() terra.StringValue {
	return terra.ReferenceAsString(lst.ref.Append("value_selection_strategy"))
}

// Version returns a reference to field version of aws_lex_slot_type.
func (lst dataLexSlotTypeAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(lst.ref.Append("version"))
}

func (lst dataLexSlotTypeAttributes) EnumerationValue() terra.SetValue[datalexslottype.EnumerationValueAttributes] {
	return terra.ReferenceAsSet[datalexslottype.EnumerationValueAttributes](lst.ref.Append("enumeration_value"))
}
