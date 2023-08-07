// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataS3AccountPublicAccessBlock creates a new instance of [DataS3AccountPublicAccessBlock].
func NewDataS3AccountPublicAccessBlock(name string, args DataS3AccountPublicAccessBlockArgs) *DataS3AccountPublicAccessBlock {
	return &DataS3AccountPublicAccessBlock{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataS3AccountPublicAccessBlock)(nil)

// DataS3AccountPublicAccessBlock represents the Terraform data resource aws_s3_account_public_access_block.
type DataS3AccountPublicAccessBlock struct {
	Name string
	Args DataS3AccountPublicAccessBlockArgs
}

// DataSource returns the Terraform object type for [DataS3AccountPublicAccessBlock].
func (sapab *DataS3AccountPublicAccessBlock) DataSource() string {
	return "aws_s3_account_public_access_block"
}

// LocalName returns the local name for [DataS3AccountPublicAccessBlock].
func (sapab *DataS3AccountPublicAccessBlock) LocalName() string {
	return sapab.Name
}

// Configuration returns the configuration (args) for [DataS3AccountPublicAccessBlock].
func (sapab *DataS3AccountPublicAccessBlock) Configuration() interface{} {
	return sapab.Args
}

// Attributes returns the attributes for [DataS3AccountPublicAccessBlock].
func (sapab *DataS3AccountPublicAccessBlock) Attributes() dataS3AccountPublicAccessBlockAttributes {
	return dataS3AccountPublicAccessBlockAttributes{ref: terra.ReferenceDataResource(sapab)}
}

// DataS3AccountPublicAccessBlockArgs contains the configurations for aws_s3_account_public_access_block.
type DataS3AccountPublicAccessBlockArgs struct {
	// AccountId: string, optional
	AccountId terra.StringValue `hcl:"account_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataS3AccountPublicAccessBlockAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_s3_account_public_access_block.
func (sapab dataS3AccountPublicAccessBlockAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(sapab.ref.Append("account_id"))
}

// BlockPublicAcls returns a reference to field block_public_acls of aws_s3_account_public_access_block.
func (sapab dataS3AccountPublicAccessBlockAttributes) BlockPublicAcls() terra.BoolValue {
	return terra.ReferenceAsBool(sapab.ref.Append("block_public_acls"))
}

// BlockPublicPolicy returns a reference to field block_public_policy of aws_s3_account_public_access_block.
func (sapab dataS3AccountPublicAccessBlockAttributes) BlockPublicPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(sapab.ref.Append("block_public_policy"))
}

// Id returns a reference to field id of aws_s3_account_public_access_block.
func (sapab dataS3AccountPublicAccessBlockAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sapab.ref.Append("id"))
}

// IgnorePublicAcls returns a reference to field ignore_public_acls of aws_s3_account_public_access_block.
func (sapab dataS3AccountPublicAccessBlockAttributes) IgnorePublicAcls() terra.BoolValue {
	return terra.ReferenceAsBool(sapab.ref.Append("ignore_public_acls"))
}

// RestrictPublicBuckets returns a reference to field restrict_public_buckets of aws_s3_account_public_access_block.
func (sapab dataS3AccountPublicAccessBlockAttributes) RestrictPublicBuckets() terra.BoolValue {
	return terra.ReferenceAsBool(sapab.ref.Append("restrict_public_buckets"))
}
