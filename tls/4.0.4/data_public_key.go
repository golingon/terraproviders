// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package tls

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataPublicKey creates a new instance of [DataPublicKey].
func NewDataPublicKey(name string, args DataPublicKeyArgs) *DataPublicKey {
	return &DataPublicKey{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPublicKey)(nil)

// DataPublicKey represents the Terraform data resource tls_public_key.
type DataPublicKey struct {
	Name string
	Args DataPublicKeyArgs
}

// DataSource returns the Terraform object type for [DataPublicKey].
func (pk *DataPublicKey) DataSource() string {
	return "tls_public_key"
}

// LocalName returns the local name for [DataPublicKey].
func (pk *DataPublicKey) LocalName() string {
	return pk.Name
}

// Configuration returns the configuration (args) for [DataPublicKey].
func (pk *DataPublicKey) Configuration() interface{} {
	return pk.Args
}

// Attributes returns the attributes for [DataPublicKey].
func (pk *DataPublicKey) Attributes() dataPublicKeyAttributes {
	return dataPublicKeyAttributes{ref: terra.ReferenceDataResource(pk)}
}

// DataPublicKeyArgs contains the configurations for tls_public_key.
type DataPublicKeyArgs struct {
	// PrivateKeyOpenssh: string, optional
	PrivateKeyOpenssh terra.StringValue `hcl:"private_key_openssh,attr"`
	// PrivateKeyPem: string, optional
	PrivateKeyPem terra.StringValue `hcl:"private_key_pem,attr"`
}
type dataPublicKeyAttributes struct {
	ref terra.Reference
}

// Algorithm returns a reference to field algorithm of tls_public_key.
func (pk dataPublicKeyAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("algorithm"))
}

// Id returns a reference to field id of tls_public_key.
func (pk dataPublicKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("id"))
}

// PrivateKeyOpenssh returns a reference to field private_key_openssh of tls_public_key.
func (pk dataPublicKeyAttributes) PrivateKeyOpenssh() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("private_key_openssh"))
}

// PrivateKeyPem returns a reference to field private_key_pem of tls_public_key.
func (pk dataPublicKeyAttributes) PrivateKeyPem() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("private_key_pem"))
}

// PublicKeyFingerprintMd5 returns a reference to field public_key_fingerprint_md5 of tls_public_key.
func (pk dataPublicKeyAttributes) PublicKeyFingerprintMd5() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("public_key_fingerprint_md5"))
}

// PublicKeyFingerprintSha256 returns a reference to field public_key_fingerprint_sha256 of tls_public_key.
func (pk dataPublicKeyAttributes) PublicKeyFingerprintSha256() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("public_key_fingerprint_sha256"))
}

// PublicKeyOpenssh returns a reference to field public_key_openssh of tls_public_key.
func (pk dataPublicKeyAttributes) PublicKeyOpenssh() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("public_key_openssh"))
}

// PublicKeyPem returns a reference to field public_key_pem of tls_public_key.
func (pk dataPublicKeyAttributes) PublicKeyPem() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("public_key_pem"))
}