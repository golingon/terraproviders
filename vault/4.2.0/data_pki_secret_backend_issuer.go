// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import "github.com/golingon/lingon/pkg/terra"

// NewDataPkiSecretBackendIssuer creates a new instance of [DataPkiSecretBackendIssuer].
func NewDataPkiSecretBackendIssuer(name string, args DataPkiSecretBackendIssuerArgs) *DataPkiSecretBackendIssuer {
	return &DataPkiSecretBackendIssuer{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPkiSecretBackendIssuer)(nil)

// DataPkiSecretBackendIssuer represents the Terraform data resource vault_pki_secret_backend_issuer.
type DataPkiSecretBackendIssuer struct {
	Name string
	Args DataPkiSecretBackendIssuerArgs
}

// DataSource returns the Terraform object type for [DataPkiSecretBackendIssuer].
func (psbi *DataPkiSecretBackendIssuer) DataSource() string {
	return "vault_pki_secret_backend_issuer"
}

// LocalName returns the local name for [DataPkiSecretBackendIssuer].
func (psbi *DataPkiSecretBackendIssuer) LocalName() string {
	return psbi.Name
}

// Configuration returns the configuration (args) for [DataPkiSecretBackendIssuer].
func (psbi *DataPkiSecretBackendIssuer) Configuration() interface{} {
	return psbi.Args
}

// Attributes returns the attributes for [DataPkiSecretBackendIssuer].
func (psbi *DataPkiSecretBackendIssuer) Attributes() dataPkiSecretBackendIssuerAttributes {
	return dataPkiSecretBackendIssuerAttributes{ref: terra.ReferenceDataResource(psbi)}
}

// DataPkiSecretBackendIssuerArgs contains the configurations for vault_pki_secret_backend_issuer.
type DataPkiSecretBackendIssuerArgs struct {
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IssuerRef: string, required
	IssuerRef terra.StringValue `hcl:"issuer_ref,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
}
type dataPkiSecretBackendIssuerAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_pki_secret_backend_issuer.
func (psbi dataPkiSecretBackendIssuerAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(psbi.ref.Append("backend"))
}

// CaChain returns a reference to field ca_chain of vault_pki_secret_backend_issuer.
func (psbi dataPkiSecretBackendIssuerAttributes) CaChain() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](psbi.ref.Append("ca_chain"))
}

// Certificate returns a reference to field certificate of vault_pki_secret_backend_issuer.
func (psbi dataPkiSecretBackendIssuerAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(psbi.ref.Append("certificate"))
}

// Id returns a reference to field id of vault_pki_secret_backend_issuer.
func (psbi dataPkiSecretBackendIssuerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(psbi.ref.Append("id"))
}

// IssuerId returns a reference to field issuer_id of vault_pki_secret_backend_issuer.
func (psbi dataPkiSecretBackendIssuerAttributes) IssuerId() terra.StringValue {
	return terra.ReferenceAsString(psbi.ref.Append("issuer_id"))
}

// IssuerName returns a reference to field issuer_name of vault_pki_secret_backend_issuer.
func (psbi dataPkiSecretBackendIssuerAttributes) IssuerName() terra.StringValue {
	return terra.ReferenceAsString(psbi.ref.Append("issuer_name"))
}

// IssuerRef returns a reference to field issuer_ref of vault_pki_secret_backend_issuer.
func (psbi dataPkiSecretBackendIssuerAttributes) IssuerRef() terra.StringValue {
	return terra.ReferenceAsString(psbi.ref.Append("issuer_ref"))
}

// KeyId returns a reference to field key_id of vault_pki_secret_backend_issuer.
func (psbi dataPkiSecretBackendIssuerAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(psbi.ref.Append("key_id"))
}

// LeafNotAfterBehavior returns a reference to field leaf_not_after_behavior of vault_pki_secret_backend_issuer.
func (psbi dataPkiSecretBackendIssuerAttributes) LeafNotAfterBehavior() terra.StringValue {
	return terra.ReferenceAsString(psbi.ref.Append("leaf_not_after_behavior"))
}

// ManualChain returns a reference to field manual_chain of vault_pki_secret_backend_issuer.
func (psbi dataPkiSecretBackendIssuerAttributes) ManualChain() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](psbi.ref.Append("manual_chain"))
}

// Namespace returns a reference to field namespace of vault_pki_secret_backend_issuer.
func (psbi dataPkiSecretBackendIssuerAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(psbi.ref.Append("namespace"))
}

// Usage returns a reference to field usage of vault_pki_secret_backend_issuer.
func (psbi dataPkiSecretBackendIssuerAttributes) Usage() terra.StringValue {
	return terra.ReferenceAsString(psbi.ref.Append("usage"))
}
