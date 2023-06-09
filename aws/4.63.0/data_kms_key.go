// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datakmskey "github.com/golingon/terraproviders/aws/4.63.0/datakmskey"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKmsKey creates a new instance of [DataKmsKey].
func NewDataKmsKey(name string, args DataKmsKeyArgs) *DataKmsKey {
	return &DataKmsKey{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKmsKey)(nil)

// DataKmsKey represents the Terraform data resource aws_kms_key.
type DataKmsKey struct {
	Name string
	Args DataKmsKeyArgs
}

// DataSource returns the Terraform object type for [DataKmsKey].
func (kk *DataKmsKey) DataSource() string {
	return "aws_kms_key"
}

// LocalName returns the local name for [DataKmsKey].
func (kk *DataKmsKey) LocalName() string {
	return kk.Name
}

// Configuration returns the configuration (args) for [DataKmsKey].
func (kk *DataKmsKey) Configuration() interface{} {
	return kk.Args
}

// Attributes returns the attributes for [DataKmsKey].
func (kk *DataKmsKey) Attributes() dataKmsKeyAttributes {
	return dataKmsKeyAttributes{ref: terra.ReferenceDataResource(kk)}
}

// DataKmsKeyArgs contains the configurations for aws_kms_key.
type DataKmsKeyArgs struct {
	// GrantTokens: list of string, optional
	GrantTokens terra.ListValue[terra.StringValue] `hcl:"grant_tokens,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyId: string, required
	KeyId terra.StringValue `hcl:"key_id,attr" validate:"required"`
	// MultiRegionConfiguration: min=0
	MultiRegionConfiguration []datakmskey.MultiRegionConfiguration `hcl:"multi_region_configuration,block" validate:"min=0"`
}
type dataKmsKeyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kms_key.
func (kk dataKmsKeyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("arn"))
}

// AwsAccountId returns a reference to field aws_account_id of aws_kms_key.
func (kk dataKmsKeyAttributes) AwsAccountId() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("aws_account_id"))
}

// CreationDate returns a reference to field creation_date of aws_kms_key.
func (kk dataKmsKeyAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("creation_date"))
}

// CustomerMasterKeySpec returns a reference to field customer_master_key_spec of aws_kms_key.
func (kk dataKmsKeyAttributes) CustomerMasterKeySpec() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("customer_master_key_spec"))
}

// DeletionDate returns a reference to field deletion_date of aws_kms_key.
func (kk dataKmsKeyAttributes) DeletionDate() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("deletion_date"))
}

// Description returns a reference to field description of aws_kms_key.
func (kk dataKmsKeyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("description"))
}

// Enabled returns a reference to field enabled of aws_kms_key.
func (kk dataKmsKeyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(kk.ref.Append("enabled"))
}

// ExpirationModel returns a reference to field expiration_model of aws_kms_key.
func (kk dataKmsKeyAttributes) ExpirationModel() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("expiration_model"))
}

// GrantTokens returns a reference to field grant_tokens of aws_kms_key.
func (kk dataKmsKeyAttributes) GrantTokens() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kk.ref.Append("grant_tokens"))
}

// Id returns a reference to field id of aws_kms_key.
func (kk dataKmsKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("id"))
}

// KeyId returns a reference to field key_id of aws_kms_key.
func (kk dataKmsKeyAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("key_id"))
}

// KeyManager returns a reference to field key_manager of aws_kms_key.
func (kk dataKmsKeyAttributes) KeyManager() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("key_manager"))
}

// KeyState returns a reference to field key_state of aws_kms_key.
func (kk dataKmsKeyAttributes) KeyState() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("key_state"))
}

// KeyUsage returns a reference to field key_usage of aws_kms_key.
func (kk dataKmsKeyAttributes) KeyUsage() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("key_usage"))
}

// MultiRegion returns a reference to field multi_region of aws_kms_key.
func (kk dataKmsKeyAttributes) MultiRegion() terra.BoolValue {
	return terra.ReferenceAsBool(kk.ref.Append("multi_region"))
}

// Origin returns a reference to field origin of aws_kms_key.
func (kk dataKmsKeyAttributes) Origin() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("origin"))
}

// ValidTo returns a reference to field valid_to of aws_kms_key.
func (kk dataKmsKeyAttributes) ValidTo() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("valid_to"))
}

func (kk dataKmsKeyAttributes) MultiRegionConfiguration() terra.ListValue[datakmskey.MultiRegionConfigurationAttributes] {
	return terra.ReferenceAsList[datakmskey.MultiRegionConfigurationAttributes](kk.ref.Append("multi_region_configuration"))
}
