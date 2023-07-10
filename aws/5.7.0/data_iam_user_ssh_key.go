// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataIamUserSshKey creates a new instance of [DataIamUserSshKey].
func NewDataIamUserSshKey(name string, args DataIamUserSshKeyArgs) *DataIamUserSshKey {
	return &DataIamUserSshKey{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamUserSshKey)(nil)

// DataIamUserSshKey represents the Terraform data resource aws_iam_user_ssh_key.
type DataIamUserSshKey struct {
	Name string
	Args DataIamUserSshKeyArgs
}

// DataSource returns the Terraform object type for [DataIamUserSshKey].
func (iusk *DataIamUserSshKey) DataSource() string {
	return "aws_iam_user_ssh_key"
}

// LocalName returns the local name for [DataIamUserSshKey].
func (iusk *DataIamUserSshKey) LocalName() string {
	return iusk.Name
}

// Configuration returns the configuration (args) for [DataIamUserSshKey].
func (iusk *DataIamUserSshKey) Configuration() interface{} {
	return iusk.Args
}

// Attributes returns the attributes for [DataIamUserSshKey].
func (iusk *DataIamUserSshKey) Attributes() dataIamUserSshKeyAttributes {
	return dataIamUserSshKeyAttributes{ref: terra.ReferenceDataResource(iusk)}
}

// DataIamUserSshKeyArgs contains the configurations for aws_iam_user_ssh_key.
type DataIamUserSshKeyArgs struct {
	// Encoding: string, required
	Encoding terra.StringValue `hcl:"encoding,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SshPublicKeyId: string, required
	SshPublicKeyId terra.StringValue `hcl:"ssh_public_key_id,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}
type dataIamUserSshKeyAttributes struct {
	ref terra.Reference
}

// Encoding returns a reference to field encoding of aws_iam_user_ssh_key.
func (iusk dataIamUserSshKeyAttributes) Encoding() terra.StringValue {
	return terra.ReferenceAsString(iusk.ref.Append("encoding"))
}

// Fingerprint returns a reference to field fingerprint of aws_iam_user_ssh_key.
func (iusk dataIamUserSshKeyAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(iusk.ref.Append("fingerprint"))
}

// Id returns a reference to field id of aws_iam_user_ssh_key.
func (iusk dataIamUserSshKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iusk.ref.Append("id"))
}

// PublicKey returns a reference to field public_key of aws_iam_user_ssh_key.
func (iusk dataIamUserSshKeyAttributes) PublicKey() terra.StringValue {
	return terra.ReferenceAsString(iusk.ref.Append("public_key"))
}

// SshPublicKeyId returns a reference to field ssh_public_key_id of aws_iam_user_ssh_key.
func (iusk dataIamUserSshKeyAttributes) SshPublicKeyId() terra.StringValue {
	return terra.ReferenceAsString(iusk.ref.Append("ssh_public_key_id"))
}

// Status returns a reference to field status of aws_iam_user_ssh_key.
func (iusk dataIamUserSshKeyAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(iusk.ref.Append("status"))
}

// Username returns a reference to field username of aws_iam_user_ssh_key.
func (iusk dataIamUserSshKeyAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(iusk.ref.Append("username"))
}
