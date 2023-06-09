// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataServiceAccountKey creates a new instance of [DataServiceAccountKey].
func NewDataServiceAccountKey(name string, args DataServiceAccountKeyArgs) *DataServiceAccountKey {
	return &DataServiceAccountKey{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServiceAccountKey)(nil)

// DataServiceAccountKey represents the Terraform data resource google_service_account_key.
type DataServiceAccountKey struct {
	Name string
	Args DataServiceAccountKeyArgs
}

// DataSource returns the Terraform object type for [DataServiceAccountKey].
func (sak *DataServiceAccountKey) DataSource() string {
	return "google_service_account_key"
}

// LocalName returns the local name for [DataServiceAccountKey].
func (sak *DataServiceAccountKey) LocalName() string {
	return sak.Name
}

// Configuration returns the configuration (args) for [DataServiceAccountKey].
func (sak *DataServiceAccountKey) Configuration() interface{} {
	return sak.Args
}

// Attributes returns the attributes for [DataServiceAccountKey].
func (sak *DataServiceAccountKey) Attributes() dataServiceAccountKeyAttributes {
	return dataServiceAccountKeyAttributes{ref: terra.ReferenceDataResource(sak)}
}

// DataServiceAccountKeyArgs contains the configurations for google_service_account_key.
type DataServiceAccountKeyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// PublicKeyType: string, optional
	PublicKeyType terra.StringValue `hcl:"public_key_type,attr"`
}
type dataServiceAccountKeyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_service_account_key.
func (sak dataServiceAccountKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sak.ref.Append("id"))
}

// KeyAlgorithm returns a reference to field key_algorithm of google_service_account_key.
func (sak dataServiceAccountKeyAttributes) KeyAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(sak.ref.Append("key_algorithm"))
}

// Name returns a reference to field name of google_service_account_key.
func (sak dataServiceAccountKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sak.ref.Append("name"))
}

// Project returns a reference to field project of google_service_account_key.
func (sak dataServiceAccountKeyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sak.ref.Append("project"))
}

// PublicKey returns a reference to field public_key of google_service_account_key.
func (sak dataServiceAccountKeyAttributes) PublicKey() terra.StringValue {
	return terra.ReferenceAsString(sak.ref.Append("public_key"))
}

// PublicKeyType returns a reference to field public_key_type of google_service_account_key.
func (sak dataServiceAccountKeyAttributes) PublicKeyType() terra.StringValue {
	return terra.ReferenceAsString(sak.ref.Append("public_key_type"))
}
