// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datakmscryptokey "github.com/golingon/terraproviders/googlebeta/4.74.0/datakmscryptokey"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKmsCryptoKey creates a new instance of [DataKmsCryptoKey].
func NewDataKmsCryptoKey(name string, args DataKmsCryptoKeyArgs) *DataKmsCryptoKey {
	return &DataKmsCryptoKey{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKmsCryptoKey)(nil)

// DataKmsCryptoKey represents the Terraform data resource google_kms_crypto_key.
type DataKmsCryptoKey struct {
	Name string
	Args DataKmsCryptoKeyArgs
}

// DataSource returns the Terraform object type for [DataKmsCryptoKey].
func (kck *DataKmsCryptoKey) DataSource() string {
	return "google_kms_crypto_key"
}

// LocalName returns the local name for [DataKmsCryptoKey].
func (kck *DataKmsCryptoKey) LocalName() string {
	return kck.Name
}

// Configuration returns the configuration (args) for [DataKmsCryptoKey].
func (kck *DataKmsCryptoKey) Configuration() interface{} {
	return kck.Args
}

// Attributes returns the attributes for [DataKmsCryptoKey].
func (kck *DataKmsCryptoKey) Attributes() dataKmsCryptoKeyAttributes {
	return dataKmsCryptoKeyAttributes{ref: terra.ReferenceDataResource(kck)}
}

// DataKmsCryptoKeyArgs contains the configurations for google_kms_crypto_key.
type DataKmsCryptoKeyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyRing: string, required
	KeyRing terra.StringValue `hcl:"key_ring,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// VersionTemplate: min=0
	VersionTemplate []datakmscryptokey.VersionTemplate `hcl:"version_template,block" validate:"min=0"`
}
type dataKmsCryptoKeyAttributes struct {
	ref terra.Reference
}

// DestroyScheduledDuration returns a reference to field destroy_scheduled_duration of google_kms_crypto_key.
func (kck dataKmsCryptoKeyAttributes) DestroyScheduledDuration() terra.StringValue {
	return terra.ReferenceAsString(kck.ref.Append("destroy_scheduled_duration"))
}

// Id returns a reference to field id of google_kms_crypto_key.
func (kck dataKmsCryptoKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kck.ref.Append("id"))
}

// ImportOnly returns a reference to field import_only of google_kms_crypto_key.
func (kck dataKmsCryptoKeyAttributes) ImportOnly() terra.BoolValue {
	return terra.ReferenceAsBool(kck.ref.Append("import_only"))
}

// KeyRing returns a reference to field key_ring of google_kms_crypto_key.
func (kck dataKmsCryptoKeyAttributes) KeyRing() terra.StringValue {
	return terra.ReferenceAsString(kck.ref.Append("key_ring"))
}

// Labels returns a reference to field labels of google_kms_crypto_key.
func (kck dataKmsCryptoKeyAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kck.ref.Append("labels"))
}

// Name returns a reference to field name of google_kms_crypto_key.
func (kck dataKmsCryptoKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kck.ref.Append("name"))
}

// Purpose returns a reference to field purpose of google_kms_crypto_key.
func (kck dataKmsCryptoKeyAttributes) Purpose() terra.StringValue {
	return terra.ReferenceAsString(kck.ref.Append("purpose"))
}

// RotationPeriod returns a reference to field rotation_period of google_kms_crypto_key.
func (kck dataKmsCryptoKeyAttributes) RotationPeriod() terra.StringValue {
	return terra.ReferenceAsString(kck.ref.Append("rotation_period"))
}

// SkipInitialVersionCreation returns a reference to field skip_initial_version_creation of google_kms_crypto_key.
func (kck dataKmsCryptoKeyAttributes) SkipInitialVersionCreation() terra.BoolValue {
	return terra.ReferenceAsBool(kck.ref.Append("skip_initial_version_creation"))
}

func (kck dataKmsCryptoKeyAttributes) VersionTemplate() terra.ListValue[datakmscryptokey.VersionTemplateAttributes] {
	return terra.ReferenceAsList[datakmscryptokey.VersionTemplateAttributes](kck.ref.Append("version_template"))
}
