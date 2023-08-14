// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataKmsCustomKeyStore creates a new instance of [DataKmsCustomKeyStore].
func NewDataKmsCustomKeyStore(name string, args DataKmsCustomKeyStoreArgs) *DataKmsCustomKeyStore {
	return &DataKmsCustomKeyStore{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKmsCustomKeyStore)(nil)

// DataKmsCustomKeyStore represents the Terraform data resource aws_kms_custom_key_store.
type DataKmsCustomKeyStore struct {
	Name string
	Args DataKmsCustomKeyStoreArgs
}

// DataSource returns the Terraform object type for [DataKmsCustomKeyStore].
func (kcks *DataKmsCustomKeyStore) DataSource() string {
	return "aws_kms_custom_key_store"
}

// LocalName returns the local name for [DataKmsCustomKeyStore].
func (kcks *DataKmsCustomKeyStore) LocalName() string {
	return kcks.Name
}

// Configuration returns the configuration (args) for [DataKmsCustomKeyStore].
func (kcks *DataKmsCustomKeyStore) Configuration() interface{} {
	return kcks.Args
}

// Attributes returns the attributes for [DataKmsCustomKeyStore].
func (kcks *DataKmsCustomKeyStore) Attributes() dataKmsCustomKeyStoreAttributes {
	return dataKmsCustomKeyStoreAttributes{ref: terra.ReferenceDataResource(kcks)}
}

// DataKmsCustomKeyStoreArgs contains the configurations for aws_kms_custom_key_store.
type DataKmsCustomKeyStoreArgs struct {
	// CustomKeyStoreId: string, optional
	CustomKeyStoreId terra.StringValue `hcl:"custom_key_store_id,attr"`
	// CustomKeyStoreName: string, optional
	CustomKeyStoreName terra.StringValue `hcl:"custom_key_store_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataKmsCustomKeyStoreAttributes struct {
	ref terra.Reference
}

// CloudHsmClusterId returns a reference to field cloud_hsm_cluster_id of aws_kms_custom_key_store.
func (kcks dataKmsCustomKeyStoreAttributes) CloudHsmClusterId() terra.StringValue {
	return terra.ReferenceAsString(kcks.ref.Append("cloud_hsm_cluster_id"))
}

// ConnectionState returns a reference to field connection_state of aws_kms_custom_key_store.
func (kcks dataKmsCustomKeyStoreAttributes) ConnectionState() terra.StringValue {
	return terra.ReferenceAsString(kcks.ref.Append("connection_state"))
}

// CreationDate returns a reference to field creation_date of aws_kms_custom_key_store.
func (kcks dataKmsCustomKeyStoreAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceAsString(kcks.ref.Append("creation_date"))
}

// CustomKeyStoreId returns a reference to field custom_key_store_id of aws_kms_custom_key_store.
func (kcks dataKmsCustomKeyStoreAttributes) CustomKeyStoreId() terra.StringValue {
	return terra.ReferenceAsString(kcks.ref.Append("custom_key_store_id"))
}

// CustomKeyStoreName returns a reference to field custom_key_store_name of aws_kms_custom_key_store.
func (kcks dataKmsCustomKeyStoreAttributes) CustomKeyStoreName() terra.StringValue {
	return terra.ReferenceAsString(kcks.ref.Append("custom_key_store_name"))
}

// Id returns a reference to field id of aws_kms_custom_key_store.
func (kcks dataKmsCustomKeyStoreAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kcks.ref.Append("id"))
}

// TrustAnchorCertificate returns a reference to field trust_anchor_certificate of aws_kms_custom_key_store.
func (kcks dataKmsCustomKeyStoreAttributes) TrustAnchorCertificate() terra.StringValue {
	return terra.ReferenceAsString(kcks.ref.Append("trust_anchor_certificate"))
}