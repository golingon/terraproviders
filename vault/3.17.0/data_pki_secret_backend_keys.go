// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataPkiSecretBackendKeys creates a new instance of [DataPkiSecretBackendKeys].
func NewDataPkiSecretBackendKeys(name string, args DataPkiSecretBackendKeysArgs) *DataPkiSecretBackendKeys {
	return &DataPkiSecretBackendKeys{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPkiSecretBackendKeys)(nil)

// DataPkiSecretBackendKeys represents the Terraform data resource vault_pki_secret_backend_keys.
type DataPkiSecretBackendKeys struct {
	Name string
	Args DataPkiSecretBackendKeysArgs
}

// DataSource returns the Terraform object type for [DataPkiSecretBackendKeys].
func (psbk *DataPkiSecretBackendKeys) DataSource() string {
	return "vault_pki_secret_backend_keys"
}

// LocalName returns the local name for [DataPkiSecretBackendKeys].
func (psbk *DataPkiSecretBackendKeys) LocalName() string {
	return psbk.Name
}

// Configuration returns the configuration (args) for [DataPkiSecretBackendKeys].
func (psbk *DataPkiSecretBackendKeys) Configuration() interface{} {
	return psbk.Args
}

// Attributes returns the attributes for [DataPkiSecretBackendKeys].
func (psbk *DataPkiSecretBackendKeys) Attributes() dataPkiSecretBackendKeysAttributes {
	return dataPkiSecretBackendKeysAttributes{ref: terra.ReferenceDataResource(psbk)}
}

// DataPkiSecretBackendKeysArgs contains the configurations for vault_pki_secret_backend_keys.
type DataPkiSecretBackendKeysArgs struct {
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
}
type dataPkiSecretBackendKeysAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_pki_secret_backend_keys.
func (psbk dataPkiSecretBackendKeysAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(psbk.ref.Append("backend"))
}

// Id returns a reference to field id of vault_pki_secret_backend_keys.
func (psbk dataPkiSecretBackendKeysAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(psbk.ref.Append("id"))
}

// KeyInfo returns a reference to field key_info of vault_pki_secret_backend_keys.
func (psbk dataPkiSecretBackendKeysAttributes) KeyInfo() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](psbk.ref.Append("key_info"))
}

// KeyInfoJson returns a reference to field key_info_json of vault_pki_secret_backend_keys.
func (psbk dataPkiSecretBackendKeysAttributes) KeyInfoJson() terra.StringValue {
	return terra.ReferenceAsString(psbk.ref.Append("key_info_json"))
}

// Keys returns a reference to field keys of vault_pki_secret_backend_keys.
func (psbk dataPkiSecretBackendKeysAttributes) Keys() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](psbk.ref.Append("keys"))
}

// Namespace returns a reference to field namespace of vault_pki_secret_backend_keys.
func (psbk dataPkiSecretBackendKeysAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(psbk.ref.Append("namespace"))
}
