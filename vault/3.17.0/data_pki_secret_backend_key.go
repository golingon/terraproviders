// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataPkiSecretBackendKey creates a new instance of [DataPkiSecretBackendKey].
func NewDataPkiSecretBackendKey(name string, args DataPkiSecretBackendKeyArgs) *DataPkiSecretBackendKey {
	return &DataPkiSecretBackendKey{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPkiSecretBackendKey)(nil)

// DataPkiSecretBackendKey represents the Terraform data resource vault_pki_secret_backend_key.
type DataPkiSecretBackendKey struct {
	Name string
	Args DataPkiSecretBackendKeyArgs
}

// DataSource returns the Terraform object type for [DataPkiSecretBackendKey].
func (psbk *DataPkiSecretBackendKey) DataSource() string {
	return "vault_pki_secret_backend_key"
}

// LocalName returns the local name for [DataPkiSecretBackendKey].
func (psbk *DataPkiSecretBackendKey) LocalName() string {
	return psbk.Name
}

// Configuration returns the configuration (args) for [DataPkiSecretBackendKey].
func (psbk *DataPkiSecretBackendKey) Configuration() interface{} {
	return psbk.Args
}

// Attributes returns the attributes for [DataPkiSecretBackendKey].
func (psbk *DataPkiSecretBackendKey) Attributes() dataPkiSecretBackendKeyAttributes {
	return dataPkiSecretBackendKeyAttributes{ref: terra.ReferenceDataResource(psbk)}
}

// DataPkiSecretBackendKeyArgs contains the configurations for vault_pki_secret_backend_key.
type DataPkiSecretBackendKeyArgs struct {
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyRef: string, required
	KeyRef terra.StringValue `hcl:"key_ref,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
}
type dataPkiSecretBackendKeyAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_pki_secret_backend_key.
func (psbk dataPkiSecretBackendKeyAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(psbk.ref.Append("backend"))
}

// Id returns a reference to field id of vault_pki_secret_backend_key.
func (psbk dataPkiSecretBackendKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(psbk.ref.Append("id"))
}

// KeyId returns a reference to field key_id of vault_pki_secret_backend_key.
func (psbk dataPkiSecretBackendKeyAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(psbk.ref.Append("key_id"))
}

// KeyName returns a reference to field key_name of vault_pki_secret_backend_key.
func (psbk dataPkiSecretBackendKeyAttributes) KeyName() terra.StringValue {
	return terra.ReferenceAsString(psbk.ref.Append("key_name"))
}

// KeyRef returns a reference to field key_ref of vault_pki_secret_backend_key.
func (psbk dataPkiSecretBackendKeyAttributes) KeyRef() terra.StringValue {
	return terra.ReferenceAsString(psbk.ref.Append("key_ref"))
}

// KeyType returns a reference to field key_type of vault_pki_secret_backend_key.
func (psbk dataPkiSecretBackendKeyAttributes) KeyType() terra.StringValue {
	return terra.ReferenceAsString(psbk.ref.Append("key_type"))
}

// Namespace returns a reference to field namespace of vault_pki_secret_backend_key.
func (psbk dataPkiSecretBackendKeyAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(psbk.ref.Append("namespace"))
}
