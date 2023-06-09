// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataQldbLedger creates a new instance of [DataQldbLedger].
func NewDataQldbLedger(name string, args DataQldbLedgerArgs) *DataQldbLedger {
	return &DataQldbLedger{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataQldbLedger)(nil)

// DataQldbLedger represents the Terraform data resource aws_qldb_ledger.
type DataQldbLedger struct {
	Name string
	Args DataQldbLedgerArgs
}

// DataSource returns the Terraform object type for [DataQldbLedger].
func (ql *DataQldbLedger) DataSource() string {
	return "aws_qldb_ledger"
}

// LocalName returns the local name for [DataQldbLedger].
func (ql *DataQldbLedger) LocalName() string {
	return ql.Name
}

// Configuration returns the configuration (args) for [DataQldbLedger].
func (ql *DataQldbLedger) Configuration() interface{} {
	return ql.Args
}

// Attributes returns the attributes for [DataQldbLedger].
func (ql *DataQldbLedger) Attributes() dataQldbLedgerAttributes {
	return dataQldbLedgerAttributes{ref: terra.ReferenceDataResource(ql)}
}

// DataQldbLedgerArgs contains the configurations for aws_qldb_ledger.
type DataQldbLedgerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataQldbLedgerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_qldb_ledger.
func (ql dataQldbLedgerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ql.ref.Append("arn"))
}

// DeletionProtection returns a reference to field deletion_protection of aws_qldb_ledger.
func (ql dataQldbLedgerAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(ql.ref.Append("deletion_protection"))
}

// Id returns a reference to field id of aws_qldb_ledger.
func (ql dataQldbLedgerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ql.ref.Append("id"))
}

// KmsKey returns a reference to field kms_key of aws_qldb_ledger.
func (ql dataQldbLedgerAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(ql.ref.Append("kms_key"))
}

// Name returns a reference to field name of aws_qldb_ledger.
func (ql dataQldbLedgerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ql.ref.Append("name"))
}

// PermissionsMode returns a reference to field permissions_mode of aws_qldb_ledger.
func (ql dataQldbLedgerAttributes) PermissionsMode() terra.StringValue {
	return terra.ReferenceAsString(ql.ref.Append("permissions_mode"))
}

// Tags returns a reference to field tags of aws_qldb_ledger.
func (ql dataQldbLedgerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ql.ref.Append("tags"))
}
