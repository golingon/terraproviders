// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataebsencryptionbydefault "github.com/golingon/terraproviders/aws/5.7.0/dataebsencryptionbydefault"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEbsEncryptionByDefault creates a new instance of [DataEbsEncryptionByDefault].
func NewDataEbsEncryptionByDefault(name string, args DataEbsEncryptionByDefaultArgs) *DataEbsEncryptionByDefault {
	return &DataEbsEncryptionByDefault{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEbsEncryptionByDefault)(nil)

// DataEbsEncryptionByDefault represents the Terraform data resource aws_ebs_encryption_by_default.
type DataEbsEncryptionByDefault struct {
	Name string
	Args DataEbsEncryptionByDefaultArgs
}

// DataSource returns the Terraform object type for [DataEbsEncryptionByDefault].
func (eebd *DataEbsEncryptionByDefault) DataSource() string {
	return "aws_ebs_encryption_by_default"
}

// LocalName returns the local name for [DataEbsEncryptionByDefault].
func (eebd *DataEbsEncryptionByDefault) LocalName() string {
	return eebd.Name
}

// Configuration returns the configuration (args) for [DataEbsEncryptionByDefault].
func (eebd *DataEbsEncryptionByDefault) Configuration() interface{} {
	return eebd.Args
}

// Attributes returns the attributes for [DataEbsEncryptionByDefault].
func (eebd *DataEbsEncryptionByDefault) Attributes() dataEbsEncryptionByDefaultAttributes {
	return dataEbsEncryptionByDefaultAttributes{ref: terra.ReferenceDataResource(eebd)}
}

// DataEbsEncryptionByDefaultArgs contains the configurations for aws_ebs_encryption_by_default.
type DataEbsEncryptionByDefaultArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *dataebsencryptionbydefault.Timeouts `hcl:"timeouts,block"`
}
type dataEbsEncryptionByDefaultAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of aws_ebs_encryption_by_default.
func (eebd dataEbsEncryptionByDefaultAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(eebd.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_ebs_encryption_by_default.
func (eebd dataEbsEncryptionByDefaultAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eebd.ref.Append("id"))
}

func (eebd dataEbsEncryptionByDefaultAttributes) Timeouts() dataebsencryptionbydefault.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataebsencryptionbydefault.TimeoutsAttributes](eebd.ref.Append("timeouts"))
}
