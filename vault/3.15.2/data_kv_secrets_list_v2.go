// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataKvSecretsListV2 creates a new instance of [DataKvSecretsListV2].
func NewDataKvSecretsListV2(name string, args DataKvSecretsListV2Args) *DataKvSecretsListV2 {
	return &DataKvSecretsListV2{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKvSecretsListV2)(nil)

// DataKvSecretsListV2 represents the Terraform data resource vault_kv_secrets_list_v2.
type DataKvSecretsListV2 struct {
	Name string
	Args DataKvSecretsListV2Args
}

// DataSource returns the Terraform object type for [DataKvSecretsListV2].
func (kslv *DataKvSecretsListV2) DataSource() string {
	return "vault_kv_secrets_list_v2"
}

// LocalName returns the local name for [DataKvSecretsListV2].
func (kslv *DataKvSecretsListV2) LocalName() string {
	return kslv.Name
}

// Configuration returns the configuration (args) for [DataKvSecretsListV2].
func (kslv *DataKvSecretsListV2) Configuration() interface{} {
	return kslv.Args
}

// Attributes returns the attributes for [DataKvSecretsListV2].
func (kslv *DataKvSecretsListV2) Attributes() dataKvSecretsListV2Attributes {
	return dataKvSecretsListV2Attributes{ref: terra.ReferenceDataResource(kslv)}
}

// DataKvSecretsListV2Args contains the configurations for vault_kv_secrets_list_v2.
type DataKvSecretsListV2Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Mount: string, required
	Mount terra.StringValue `hcl:"mount,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
}
type dataKvSecretsListV2Attributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of vault_kv_secrets_list_v2.
func (kslv dataKvSecretsListV2Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kslv.ref.Append("id"))
}

// Mount returns a reference to field mount of vault_kv_secrets_list_v2.
func (kslv dataKvSecretsListV2Attributes) Mount() terra.StringValue {
	return terra.ReferenceAsString(kslv.ref.Append("mount"))
}

// Name returns a reference to field name of vault_kv_secrets_list_v2.
func (kslv dataKvSecretsListV2Attributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kslv.ref.Append("name"))
}

// Names returns a reference to field names of vault_kv_secrets_list_v2.
func (kslv dataKvSecretsListV2Attributes) Names() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kslv.ref.Append("names"))
}

// Namespace returns a reference to field namespace of vault_kv_secrets_list_v2.
func (kslv dataKvSecretsListV2Attributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(kslv.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_kv_secrets_list_v2.
func (kslv dataKvSecretsListV2Attributes) Path() terra.StringValue {
	return terra.ReferenceAsString(kslv.ref.Append("path"))
}
