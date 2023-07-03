// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataKvSecret creates a new instance of [DataKvSecret].
func NewDataKvSecret(name string, args DataKvSecretArgs) *DataKvSecret {
	return &DataKvSecret{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKvSecret)(nil)

// DataKvSecret represents the Terraform data resource vault_kv_secret.
type DataKvSecret struct {
	Name string
	Args DataKvSecretArgs
}

// DataSource returns the Terraform object type for [DataKvSecret].
func (ks *DataKvSecret) DataSource() string {
	return "vault_kv_secret"
}

// LocalName returns the local name for [DataKvSecret].
func (ks *DataKvSecret) LocalName() string {
	return ks.Name
}

// Configuration returns the configuration (args) for [DataKvSecret].
func (ks *DataKvSecret) Configuration() interface{} {
	return ks.Args
}

// Attributes returns the attributes for [DataKvSecret].
func (ks *DataKvSecret) Attributes() dataKvSecretAttributes {
	return dataKvSecretAttributes{ref: terra.ReferenceDataResource(ks)}
}

// DataKvSecretArgs contains the configurations for vault_kv_secret.
type DataKvSecretArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
}
type dataKvSecretAttributes struct {
	ref terra.Reference
}

// Data returns a reference to field data of vault_kv_secret.
func (ks dataKvSecretAttributes) Data() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ks.ref.Append("data"))
}

// DataJson returns a reference to field data_json of vault_kv_secret.
func (ks dataKvSecretAttributes) DataJson() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("data_json"))
}

// Id returns a reference to field id of vault_kv_secret.
func (ks dataKvSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("id"))
}

// LeaseDuration returns a reference to field lease_duration of vault_kv_secret.
func (ks dataKvSecretAttributes) LeaseDuration() terra.NumberValue {
	return terra.ReferenceAsNumber(ks.ref.Append("lease_duration"))
}

// LeaseId returns a reference to field lease_id of vault_kv_secret.
func (ks dataKvSecretAttributes) LeaseId() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("lease_id"))
}

// LeaseRenewable returns a reference to field lease_renewable of vault_kv_secret.
func (ks dataKvSecretAttributes) LeaseRenewable() terra.BoolValue {
	return terra.ReferenceAsBool(ks.ref.Append("lease_renewable"))
}

// Namespace returns a reference to field namespace of vault_kv_secret.
func (ks dataKvSecretAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_kv_secret.
func (ks dataKvSecretAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("path"))
}
