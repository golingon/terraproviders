// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataPkiSecretBackendIssuers creates a new instance of [DataPkiSecretBackendIssuers].
func NewDataPkiSecretBackendIssuers(name string, args DataPkiSecretBackendIssuersArgs) *DataPkiSecretBackendIssuers {
	return &DataPkiSecretBackendIssuers{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPkiSecretBackendIssuers)(nil)

// DataPkiSecretBackendIssuers represents the Terraform data resource vault_pki_secret_backend_issuers.
type DataPkiSecretBackendIssuers struct {
	Name string
	Args DataPkiSecretBackendIssuersArgs
}

// DataSource returns the Terraform object type for [DataPkiSecretBackendIssuers].
func (psbi *DataPkiSecretBackendIssuers) DataSource() string {
	return "vault_pki_secret_backend_issuers"
}

// LocalName returns the local name for [DataPkiSecretBackendIssuers].
func (psbi *DataPkiSecretBackendIssuers) LocalName() string {
	return psbi.Name
}

// Configuration returns the configuration (args) for [DataPkiSecretBackendIssuers].
func (psbi *DataPkiSecretBackendIssuers) Configuration() interface{} {
	return psbi.Args
}

// Attributes returns the attributes for [DataPkiSecretBackendIssuers].
func (psbi *DataPkiSecretBackendIssuers) Attributes() dataPkiSecretBackendIssuersAttributes {
	return dataPkiSecretBackendIssuersAttributes{ref: terra.ReferenceDataResource(psbi)}
}

// DataPkiSecretBackendIssuersArgs contains the configurations for vault_pki_secret_backend_issuers.
type DataPkiSecretBackendIssuersArgs struct {
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
}
type dataPkiSecretBackendIssuersAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_pki_secret_backend_issuers.
func (psbi dataPkiSecretBackendIssuersAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(psbi.ref.Append("backend"))
}

// Id returns a reference to field id of vault_pki_secret_backend_issuers.
func (psbi dataPkiSecretBackendIssuersAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(psbi.ref.Append("id"))
}

// KeyInfo returns a reference to field key_info of vault_pki_secret_backend_issuers.
func (psbi dataPkiSecretBackendIssuersAttributes) KeyInfo() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](psbi.ref.Append("key_info"))
}

// KeyInfoJson returns a reference to field key_info_json of vault_pki_secret_backend_issuers.
func (psbi dataPkiSecretBackendIssuersAttributes) KeyInfoJson() terra.StringValue {
	return terra.ReferenceAsString(psbi.ref.Append("key_info_json"))
}

// Keys returns a reference to field keys of vault_pki_secret_backend_issuers.
func (psbi dataPkiSecretBackendIssuersAttributes) Keys() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](psbi.ref.Append("keys"))
}

// Namespace returns a reference to field namespace of vault_pki_secret_backend_issuers.
func (psbi dataPkiSecretBackendIssuersAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(psbi.ref.Append("namespace"))
}