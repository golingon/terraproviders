// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_iam_user_ssh_key

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_iam_user_ssh_key.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aiusk *DataSource) DataSource() string {
	return "aws_iam_user_ssh_key"
}

// LocalName returns the local name for [DataSource].
func (aiusk *DataSource) LocalName() string {
	return aiusk.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aiusk *DataSource) Configuration() interface{} {
	return aiusk.Args
}

// Attributes returns the attributes for [DataSource].
func (aiusk *DataSource) Attributes() dataAwsIamUserSshKeyAttributes {
	return dataAwsIamUserSshKeyAttributes{ref: terra.ReferenceDataSource(aiusk)}
}

// DataArgs contains the configurations for aws_iam_user_ssh_key.
type DataArgs struct {
	// Encoding: string, required
	Encoding terra.StringValue `hcl:"encoding,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SshPublicKeyId: string, required
	SshPublicKeyId terra.StringValue `hcl:"ssh_public_key_id,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type dataAwsIamUserSshKeyAttributes struct {
	ref terra.Reference
}

// Encoding returns a reference to field encoding of aws_iam_user_ssh_key.
func (aiusk dataAwsIamUserSshKeyAttributes) Encoding() terra.StringValue {
	return terra.ReferenceAsString(aiusk.ref.Append("encoding"))
}

// Fingerprint returns a reference to field fingerprint of aws_iam_user_ssh_key.
func (aiusk dataAwsIamUserSshKeyAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(aiusk.ref.Append("fingerprint"))
}

// Id returns a reference to field id of aws_iam_user_ssh_key.
func (aiusk dataAwsIamUserSshKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aiusk.ref.Append("id"))
}

// PublicKey returns a reference to field public_key of aws_iam_user_ssh_key.
func (aiusk dataAwsIamUserSshKeyAttributes) PublicKey() terra.StringValue {
	return terra.ReferenceAsString(aiusk.ref.Append("public_key"))
}

// SshPublicKeyId returns a reference to field ssh_public_key_id of aws_iam_user_ssh_key.
func (aiusk dataAwsIamUserSshKeyAttributes) SshPublicKeyId() terra.StringValue {
	return terra.ReferenceAsString(aiusk.ref.Append("ssh_public_key_id"))
}

// Status returns a reference to field status of aws_iam_user_ssh_key.
func (aiusk dataAwsIamUserSshKeyAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(aiusk.ref.Append("status"))
}

// Username returns a reference to field username of aws_iam_user_ssh_key.
func (aiusk dataAwsIamUserSshKeyAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(aiusk.ref.Append("username"))
}
