// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataKubernetesAuthBackendConfig creates a new instance of [DataKubernetesAuthBackendConfig].
func NewDataKubernetesAuthBackendConfig(name string, args DataKubernetesAuthBackendConfigArgs) *DataKubernetesAuthBackendConfig {
	return &DataKubernetesAuthBackendConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKubernetesAuthBackendConfig)(nil)

// DataKubernetesAuthBackendConfig represents the Terraform data resource vault_kubernetes_auth_backend_config.
type DataKubernetesAuthBackendConfig struct {
	Name string
	Args DataKubernetesAuthBackendConfigArgs
}

// DataSource returns the Terraform object type for [DataKubernetesAuthBackendConfig].
func (kabc *DataKubernetesAuthBackendConfig) DataSource() string {
	return "vault_kubernetes_auth_backend_config"
}

// LocalName returns the local name for [DataKubernetesAuthBackendConfig].
func (kabc *DataKubernetesAuthBackendConfig) LocalName() string {
	return kabc.Name
}

// Configuration returns the configuration (args) for [DataKubernetesAuthBackendConfig].
func (kabc *DataKubernetesAuthBackendConfig) Configuration() interface{} {
	return kabc.Args
}

// Attributes returns the attributes for [DataKubernetesAuthBackendConfig].
func (kabc *DataKubernetesAuthBackendConfig) Attributes() dataKubernetesAuthBackendConfigAttributes {
	return dataKubernetesAuthBackendConfigAttributes{ref: terra.ReferenceDataResource(kabc)}
}

// DataKubernetesAuthBackendConfigArgs contains the configurations for vault_kubernetes_auth_backend_config.
type DataKubernetesAuthBackendConfigArgs struct {
	// Backend: string, optional
	Backend terra.StringValue `hcl:"backend,attr"`
	// DisableIssValidation: bool, optional
	DisableIssValidation terra.BoolValue `hcl:"disable_iss_validation,attr"`
	// DisableLocalCaJwt: bool, optional
	DisableLocalCaJwt terra.BoolValue `hcl:"disable_local_ca_jwt,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Issuer: string, optional
	Issuer terra.StringValue `hcl:"issuer,attr"`
	// KubernetesCaCert: string, optional
	KubernetesCaCert terra.StringValue `hcl:"kubernetes_ca_cert,attr"`
	// KubernetesHost: string, optional
	KubernetesHost terra.StringValue `hcl:"kubernetes_host,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// PemKeys: list of string, optional
	PemKeys terra.ListValue[terra.StringValue] `hcl:"pem_keys,attr"`
}
type dataKubernetesAuthBackendConfigAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_kubernetes_auth_backend_config.
func (kabc dataKubernetesAuthBackendConfigAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(kabc.ref.Append("backend"))
}

// DisableIssValidation returns a reference to field disable_iss_validation of vault_kubernetes_auth_backend_config.
func (kabc dataKubernetesAuthBackendConfigAttributes) DisableIssValidation() terra.BoolValue {
	return terra.ReferenceAsBool(kabc.ref.Append("disable_iss_validation"))
}

// DisableLocalCaJwt returns a reference to field disable_local_ca_jwt of vault_kubernetes_auth_backend_config.
func (kabc dataKubernetesAuthBackendConfigAttributes) DisableLocalCaJwt() terra.BoolValue {
	return terra.ReferenceAsBool(kabc.ref.Append("disable_local_ca_jwt"))
}

// Id returns a reference to field id of vault_kubernetes_auth_backend_config.
func (kabc dataKubernetesAuthBackendConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kabc.ref.Append("id"))
}

// Issuer returns a reference to field issuer of vault_kubernetes_auth_backend_config.
func (kabc dataKubernetesAuthBackendConfigAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(kabc.ref.Append("issuer"))
}

// KubernetesCaCert returns a reference to field kubernetes_ca_cert of vault_kubernetes_auth_backend_config.
func (kabc dataKubernetesAuthBackendConfigAttributes) KubernetesCaCert() terra.StringValue {
	return terra.ReferenceAsString(kabc.ref.Append("kubernetes_ca_cert"))
}

// KubernetesHost returns a reference to field kubernetes_host of vault_kubernetes_auth_backend_config.
func (kabc dataKubernetesAuthBackendConfigAttributes) KubernetesHost() terra.StringValue {
	return terra.ReferenceAsString(kabc.ref.Append("kubernetes_host"))
}

// Namespace returns a reference to field namespace of vault_kubernetes_auth_backend_config.
func (kabc dataKubernetesAuthBackendConfigAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(kabc.ref.Append("namespace"))
}

// PemKeys returns a reference to field pem_keys of vault_kubernetes_auth_backend_config.
func (kabc dataKubernetesAuthBackendConfigAttributes) PemKeys() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kabc.ref.Append("pem_keys"))
}
