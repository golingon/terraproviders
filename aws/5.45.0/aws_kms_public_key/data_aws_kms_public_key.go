// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_kms_public_key

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_kms_public_key.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (akpk *DataSource) DataSource() string {
	return "aws_kms_public_key"
}

// LocalName returns the local name for [DataSource].
func (akpk *DataSource) LocalName() string {
	return akpk.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (akpk *DataSource) Configuration() interface{} {
	return akpk.Args
}

// Attributes returns the attributes for [DataSource].
func (akpk *DataSource) Attributes() dataAwsKmsPublicKeyAttributes {
	return dataAwsKmsPublicKeyAttributes{ref: terra.ReferenceDataSource(akpk)}
}

// DataArgs contains the configurations for aws_kms_public_key.
type DataArgs struct {
	// GrantTokens: list of string, optional
	GrantTokens terra.ListValue[terra.StringValue] `hcl:"grant_tokens,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyId: string, required
	KeyId terra.StringValue `hcl:"key_id,attr" validate:"required"`
}

type dataAwsKmsPublicKeyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kms_public_key.
func (akpk dataAwsKmsPublicKeyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(akpk.ref.Append("arn"))
}

// CustomerMasterKeySpec returns a reference to field customer_master_key_spec of aws_kms_public_key.
func (akpk dataAwsKmsPublicKeyAttributes) CustomerMasterKeySpec() terra.StringValue {
	return terra.ReferenceAsString(akpk.ref.Append("customer_master_key_spec"))
}

// EncryptionAlgorithms returns a reference to field encryption_algorithms of aws_kms_public_key.
func (akpk dataAwsKmsPublicKeyAttributes) EncryptionAlgorithms() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](akpk.ref.Append("encryption_algorithms"))
}

// GrantTokens returns a reference to field grant_tokens of aws_kms_public_key.
func (akpk dataAwsKmsPublicKeyAttributes) GrantTokens() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](akpk.ref.Append("grant_tokens"))
}

// Id returns a reference to field id of aws_kms_public_key.
func (akpk dataAwsKmsPublicKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akpk.ref.Append("id"))
}

// KeyId returns a reference to field key_id of aws_kms_public_key.
func (akpk dataAwsKmsPublicKeyAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(akpk.ref.Append("key_id"))
}

// KeyUsage returns a reference to field key_usage of aws_kms_public_key.
func (akpk dataAwsKmsPublicKeyAttributes) KeyUsage() terra.StringValue {
	return terra.ReferenceAsString(akpk.ref.Append("key_usage"))
}

// PublicKey returns a reference to field public_key of aws_kms_public_key.
func (akpk dataAwsKmsPublicKeyAttributes) PublicKey() terra.StringValue {
	return terra.ReferenceAsString(akpk.ref.Append("public_key"))
}

// PublicKeyPem returns a reference to field public_key_pem of aws_kms_public_key.
func (akpk dataAwsKmsPublicKeyAttributes) PublicKeyPem() terra.StringValue {
	return terra.ReferenceAsString(akpk.ref.Append("public_key_pem"))
}

// SigningAlgorithms returns a reference to field signing_algorithms of aws_kms_public_key.
func (akpk dataAwsKmsPublicKeyAttributes) SigningAlgorithms() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](akpk.ref.Append("signing_algorithms"))
}
