// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datakmscryptokeyversion "github.com/golingon/terraproviders/googlebeta/4.62.0/datakmscryptokeyversion"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKmsCryptoKeyVersion creates a new instance of [DataKmsCryptoKeyVersion].
func NewDataKmsCryptoKeyVersion(name string, args DataKmsCryptoKeyVersionArgs) *DataKmsCryptoKeyVersion {
	return &DataKmsCryptoKeyVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKmsCryptoKeyVersion)(nil)

// DataKmsCryptoKeyVersion represents the Terraform data resource google_kms_crypto_key_version.
type DataKmsCryptoKeyVersion struct {
	Name string
	Args DataKmsCryptoKeyVersionArgs
}

// DataSource returns the Terraform object type for [DataKmsCryptoKeyVersion].
func (kckv *DataKmsCryptoKeyVersion) DataSource() string {
	return "google_kms_crypto_key_version"
}

// LocalName returns the local name for [DataKmsCryptoKeyVersion].
func (kckv *DataKmsCryptoKeyVersion) LocalName() string {
	return kckv.Name
}

// Configuration returns the configuration (args) for [DataKmsCryptoKeyVersion].
func (kckv *DataKmsCryptoKeyVersion) Configuration() interface{} {
	return kckv.Args
}

// Attributes returns the attributes for [DataKmsCryptoKeyVersion].
func (kckv *DataKmsCryptoKeyVersion) Attributes() dataKmsCryptoKeyVersionAttributes {
	return dataKmsCryptoKeyVersionAttributes{ref: terra.ReferenceDataResource(kckv)}
}

// DataKmsCryptoKeyVersionArgs contains the configurations for google_kms_crypto_key_version.
type DataKmsCryptoKeyVersionArgs struct {
	// CryptoKey: string, required
	CryptoKey terra.StringValue `hcl:"crypto_key,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Version: number, optional
	Version terra.NumberValue `hcl:"version,attr"`
	// PublicKey: min=0
	PublicKey []datakmscryptokeyversion.PublicKey `hcl:"public_key,block" validate:"min=0"`
}
type dataKmsCryptoKeyVersionAttributes struct {
	ref terra.Reference
}

// Algorithm returns a reference to field algorithm of google_kms_crypto_key_version.
func (kckv dataKmsCryptoKeyVersionAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceAsString(kckv.ref.Append("algorithm"))
}

// CryptoKey returns a reference to field crypto_key of google_kms_crypto_key_version.
func (kckv dataKmsCryptoKeyVersionAttributes) CryptoKey() terra.StringValue {
	return terra.ReferenceAsString(kckv.ref.Append("crypto_key"))
}

// Id returns a reference to field id of google_kms_crypto_key_version.
func (kckv dataKmsCryptoKeyVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kckv.ref.Append("id"))
}

// Name returns a reference to field name of google_kms_crypto_key_version.
func (kckv dataKmsCryptoKeyVersionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kckv.ref.Append("name"))
}

// ProtectionLevel returns a reference to field protection_level of google_kms_crypto_key_version.
func (kckv dataKmsCryptoKeyVersionAttributes) ProtectionLevel() terra.StringValue {
	return terra.ReferenceAsString(kckv.ref.Append("protection_level"))
}

// State returns a reference to field state of google_kms_crypto_key_version.
func (kckv dataKmsCryptoKeyVersionAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(kckv.ref.Append("state"))
}

// Version returns a reference to field version of google_kms_crypto_key_version.
func (kckv dataKmsCryptoKeyVersionAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(kckv.ref.Append("version"))
}

func (kckv dataKmsCryptoKeyVersionAttributes) PublicKey() terra.ListValue[datakmscryptokeyversion.PublicKeyAttributes] {
	return terra.ReferenceAsList[datakmscryptokeyversion.PublicKeyAttributes](kckv.ref.Append("public_key"))
}
