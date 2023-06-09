// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataKvSecretSubkeysV2 creates a new instance of [DataKvSecretSubkeysV2].
func NewDataKvSecretSubkeysV2(name string, args DataKvSecretSubkeysV2Args) *DataKvSecretSubkeysV2 {
	return &DataKvSecretSubkeysV2{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKvSecretSubkeysV2)(nil)

// DataKvSecretSubkeysV2 represents the Terraform data resource vault_kv_secret_subkeys_v2.
type DataKvSecretSubkeysV2 struct {
	Name string
	Args DataKvSecretSubkeysV2Args
}

// DataSource returns the Terraform object type for [DataKvSecretSubkeysV2].
func (kssv *DataKvSecretSubkeysV2) DataSource() string {
	return "vault_kv_secret_subkeys_v2"
}

// LocalName returns the local name for [DataKvSecretSubkeysV2].
func (kssv *DataKvSecretSubkeysV2) LocalName() string {
	return kssv.Name
}

// Configuration returns the configuration (args) for [DataKvSecretSubkeysV2].
func (kssv *DataKvSecretSubkeysV2) Configuration() interface{} {
	return kssv.Args
}

// Attributes returns the attributes for [DataKvSecretSubkeysV2].
func (kssv *DataKvSecretSubkeysV2) Attributes() dataKvSecretSubkeysV2Attributes {
	return dataKvSecretSubkeysV2Attributes{ref: terra.ReferenceDataResource(kssv)}
}

// DataKvSecretSubkeysV2Args contains the configurations for vault_kv_secret_subkeys_v2.
type DataKvSecretSubkeysV2Args struct {
	// Depth: number, optional
	Depth terra.NumberValue `hcl:"depth,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Mount: string, required
	Mount terra.StringValue `hcl:"mount,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Version: number, optional
	Version terra.NumberValue `hcl:"version,attr"`
}
type dataKvSecretSubkeysV2Attributes struct {
	ref terra.Reference
}

// Data returns a reference to field data of vault_kv_secret_subkeys_v2.
func (kssv dataKvSecretSubkeysV2Attributes) Data() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kssv.ref.Append("data"))
}

// DataJson returns a reference to field data_json of vault_kv_secret_subkeys_v2.
func (kssv dataKvSecretSubkeysV2Attributes) DataJson() terra.StringValue {
	return terra.ReferenceAsString(kssv.ref.Append("data_json"))
}

// Depth returns a reference to field depth of vault_kv_secret_subkeys_v2.
func (kssv dataKvSecretSubkeysV2Attributes) Depth() terra.NumberValue {
	return terra.ReferenceAsNumber(kssv.ref.Append("depth"))
}

// Id returns a reference to field id of vault_kv_secret_subkeys_v2.
func (kssv dataKvSecretSubkeysV2Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kssv.ref.Append("id"))
}

// Mount returns a reference to field mount of vault_kv_secret_subkeys_v2.
func (kssv dataKvSecretSubkeysV2Attributes) Mount() terra.StringValue {
	return terra.ReferenceAsString(kssv.ref.Append("mount"))
}

// Name returns a reference to field name of vault_kv_secret_subkeys_v2.
func (kssv dataKvSecretSubkeysV2Attributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kssv.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_kv_secret_subkeys_v2.
func (kssv dataKvSecretSubkeysV2Attributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(kssv.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_kv_secret_subkeys_v2.
func (kssv dataKvSecretSubkeysV2Attributes) Path() terra.StringValue {
	return terra.ReferenceAsString(kssv.ref.Append("path"))
}

// Version returns a reference to field version of vault_kv_secret_subkeys_v2.
func (kssv dataKvSecretSubkeysV2Attributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(kssv.ref.Append("version"))
}
