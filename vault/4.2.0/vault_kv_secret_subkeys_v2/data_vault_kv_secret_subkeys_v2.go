// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_kv_secret_subkeys_v2

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource vault_kv_secret_subkeys_v2.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (vkssv *DataSource) DataSource() string {
	return "vault_kv_secret_subkeys_v2"
}

// LocalName returns the local name for [DataSource].
func (vkssv *DataSource) LocalName() string {
	return vkssv.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (vkssv *DataSource) Configuration() interface{} {
	return vkssv.Args
}

// Attributes returns the attributes for [DataSource].
func (vkssv *DataSource) Attributes() dataVaultKvSecretSubkeysV2Attributes {
	return dataVaultKvSecretSubkeysV2Attributes{ref: terra.ReferenceDataSource(vkssv)}
}

// DataArgs contains the configurations for vault_kv_secret_subkeys_v2.
type DataArgs struct {
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

type dataVaultKvSecretSubkeysV2Attributes struct {
	ref terra.Reference
}

// Data returns a reference to field data of vault_kv_secret_subkeys_v2.
func (vkssv dataVaultKvSecretSubkeysV2Attributes) Data() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vkssv.ref.Append("data"))
}

// DataJson returns a reference to field data_json of vault_kv_secret_subkeys_v2.
func (vkssv dataVaultKvSecretSubkeysV2Attributes) DataJson() terra.StringValue {
	return terra.ReferenceAsString(vkssv.ref.Append("data_json"))
}

// Depth returns a reference to field depth of vault_kv_secret_subkeys_v2.
func (vkssv dataVaultKvSecretSubkeysV2Attributes) Depth() terra.NumberValue {
	return terra.ReferenceAsNumber(vkssv.ref.Append("depth"))
}

// Id returns a reference to field id of vault_kv_secret_subkeys_v2.
func (vkssv dataVaultKvSecretSubkeysV2Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vkssv.ref.Append("id"))
}

// Mount returns a reference to field mount of vault_kv_secret_subkeys_v2.
func (vkssv dataVaultKvSecretSubkeysV2Attributes) Mount() terra.StringValue {
	return terra.ReferenceAsString(vkssv.ref.Append("mount"))
}

// Name returns a reference to field name of vault_kv_secret_subkeys_v2.
func (vkssv dataVaultKvSecretSubkeysV2Attributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vkssv.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_kv_secret_subkeys_v2.
func (vkssv dataVaultKvSecretSubkeysV2Attributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vkssv.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_kv_secret_subkeys_v2.
func (vkssv dataVaultKvSecretSubkeysV2Attributes) Path() terra.StringValue {
	return terra.ReferenceAsString(vkssv.ref.Append("path"))
}

// Version returns a reference to field version of vault_kv_secret_subkeys_v2.
func (vkssv dataVaultKvSecretSubkeysV2Attributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(vkssv.ref.Append("version"))
}
