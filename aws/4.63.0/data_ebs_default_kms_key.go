// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataebsdefaultkmskey "github.com/golingon/terraproviders/aws/4.63.0/dataebsdefaultkmskey"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEbsDefaultKmsKey creates a new instance of [DataEbsDefaultKmsKey].
func NewDataEbsDefaultKmsKey(name string, args DataEbsDefaultKmsKeyArgs) *DataEbsDefaultKmsKey {
	return &DataEbsDefaultKmsKey{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEbsDefaultKmsKey)(nil)

// DataEbsDefaultKmsKey represents the Terraform data resource aws_ebs_default_kms_key.
type DataEbsDefaultKmsKey struct {
	Name string
	Args DataEbsDefaultKmsKeyArgs
}

// DataSource returns the Terraform object type for [DataEbsDefaultKmsKey].
func (edkk *DataEbsDefaultKmsKey) DataSource() string {
	return "aws_ebs_default_kms_key"
}

// LocalName returns the local name for [DataEbsDefaultKmsKey].
func (edkk *DataEbsDefaultKmsKey) LocalName() string {
	return edkk.Name
}

// Configuration returns the configuration (args) for [DataEbsDefaultKmsKey].
func (edkk *DataEbsDefaultKmsKey) Configuration() interface{} {
	return edkk.Args
}

// Attributes returns the attributes for [DataEbsDefaultKmsKey].
func (edkk *DataEbsDefaultKmsKey) Attributes() dataEbsDefaultKmsKeyAttributes {
	return dataEbsDefaultKmsKeyAttributes{ref: terra.ReferenceDataResource(edkk)}
}

// DataEbsDefaultKmsKeyArgs contains the configurations for aws_ebs_default_kms_key.
type DataEbsDefaultKmsKeyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *dataebsdefaultkmskey.Timeouts `hcl:"timeouts,block"`
}
type dataEbsDefaultKmsKeyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ebs_default_kms_key.
func (edkk dataEbsDefaultKmsKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(edkk.ref.Append("id"))
}

// KeyArn returns a reference to field key_arn of aws_ebs_default_kms_key.
func (edkk dataEbsDefaultKmsKeyAttributes) KeyArn() terra.StringValue {
	return terra.ReferenceAsString(edkk.ref.Append("key_arn"))
}

func (edkk dataEbsDefaultKmsKeyAttributes) Timeouts() dataebsdefaultkmskey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataebsdefaultkmskey.TimeoutsAttributes](edkk.ref.Append("timeouts"))
}