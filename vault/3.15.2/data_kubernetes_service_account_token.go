// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataKubernetesServiceAccountToken creates a new instance of [DataKubernetesServiceAccountToken].
func NewDataKubernetesServiceAccountToken(name string, args DataKubernetesServiceAccountTokenArgs) *DataKubernetesServiceAccountToken {
	return &DataKubernetesServiceAccountToken{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKubernetesServiceAccountToken)(nil)

// DataKubernetesServiceAccountToken represents the Terraform data resource vault_kubernetes_service_account_token.
type DataKubernetesServiceAccountToken struct {
	Name string
	Args DataKubernetesServiceAccountTokenArgs
}

// DataSource returns the Terraform object type for [DataKubernetesServiceAccountToken].
func (ksat *DataKubernetesServiceAccountToken) DataSource() string {
	return "vault_kubernetes_service_account_token"
}

// LocalName returns the local name for [DataKubernetesServiceAccountToken].
func (ksat *DataKubernetesServiceAccountToken) LocalName() string {
	return ksat.Name
}

// Configuration returns the configuration (args) for [DataKubernetesServiceAccountToken].
func (ksat *DataKubernetesServiceAccountToken) Configuration() interface{} {
	return ksat.Args
}

// Attributes returns the attributes for [DataKubernetesServiceAccountToken].
func (ksat *DataKubernetesServiceAccountToken) Attributes() dataKubernetesServiceAccountTokenAttributes {
	return dataKubernetesServiceAccountTokenAttributes{ref: terra.ReferenceDataResource(ksat)}
}

// DataKubernetesServiceAccountTokenArgs contains the configurations for vault_kubernetes_service_account_token.
type DataKubernetesServiceAccountTokenArgs struct {
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// ClusterRoleBinding: bool, optional
	ClusterRoleBinding terra.BoolValue `hcl:"cluster_role_binding,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KubernetesNamespace: string, required
	KubernetesNamespace terra.StringValue `hcl:"kubernetes_namespace,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Ttl: string, optional
	Ttl terra.StringValue `hcl:"ttl,attr"`
}
type dataKubernetesServiceAccountTokenAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_kubernetes_service_account_token.
func (ksat dataKubernetesServiceAccountTokenAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(ksat.ref.Append("backend"))
}

// ClusterRoleBinding returns a reference to field cluster_role_binding of vault_kubernetes_service_account_token.
func (ksat dataKubernetesServiceAccountTokenAttributes) ClusterRoleBinding() terra.BoolValue {
	return terra.ReferenceAsBool(ksat.ref.Append("cluster_role_binding"))
}

// Id returns a reference to field id of vault_kubernetes_service_account_token.
func (ksat dataKubernetesServiceAccountTokenAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ksat.ref.Append("id"))
}

// KubernetesNamespace returns a reference to field kubernetes_namespace of vault_kubernetes_service_account_token.
func (ksat dataKubernetesServiceAccountTokenAttributes) KubernetesNamespace() terra.StringValue {
	return terra.ReferenceAsString(ksat.ref.Append("kubernetes_namespace"))
}

// LeaseDuration returns a reference to field lease_duration of vault_kubernetes_service_account_token.
func (ksat dataKubernetesServiceAccountTokenAttributes) LeaseDuration() terra.NumberValue {
	return terra.ReferenceAsNumber(ksat.ref.Append("lease_duration"))
}

// LeaseId returns a reference to field lease_id of vault_kubernetes_service_account_token.
func (ksat dataKubernetesServiceAccountTokenAttributes) LeaseId() terra.StringValue {
	return terra.ReferenceAsString(ksat.ref.Append("lease_id"))
}

// LeaseRenewable returns a reference to field lease_renewable of vault_kubernetes_service_account_token.
func (ksat dataKubernetesServiceAccountTokenAttributes) LeaseRenewable() terra.BoolValue {
	return terra.ReferenceAsBool(ksat.ref.Append("lease_renewable"))
}

// Namespace returns a reference to field namespace of vault_kubernetes_service_account_token.
func (ksat dataKubernetesServiceAccountTokenAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(ksat.ref.Append("namespace"))
}

// Role returns a reference to field role of vault_kubernetes_service_account_token.
func (ksat dataKubernetesServiceAccountTokenAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ksat.ref.Append("role"))
}

// ServiceAccountName returns a reference to field service_account_name of vault_kubernetes_service_account_token.
func (ksat dataKubernetesServiceAccountTokenAttributes) ServiceAccountName() terra.StringValue {
	return terra.ReferenceAsString(ksat.ref.Append("service_account_name"))
}

// ServiceAccountNamespace returns a reference to field service_account_namespace of vault_kubernetes_service_account_token.
func (ksat dataKubernetesServiceAccountTokenAttributes) ServiceAccountNamespace() terra.StringValue {
	return terra.ReferenceAsString(ksat.ref.Append("service_account_namespace"))
}

// ServiceAccountToken returns a reference to field service_account_token of vault_kubernetes_service_account_token.
func (ksat dataKubernetesServiceAccountTokenAttributes) ServiceAccountToken() terra.StringValue {
	return terra.ReferenceAsString(ksat.ref.Append("service_account_token"))
}

// Ttl returns a reference to field ttl of vault_kubernetes_service_account_token.
func (ksat dataKubernetesServiceAccountTokenAttributes) Ttl() terra.StringValue {
	return terra.ReferenceAsString(ksat.ref.Append("ttl"))
}
