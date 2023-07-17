// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataGenericSecret creates a new instance of [DataGenericSecret].
func NewDataGenericSecret(name string, args DataGenericSecretArgs) *DataGenericSecret {
	return &DataGenericSecret{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataGenericSecret)(nil)

// DataGenericSecret represents the Terraform data resource vault_generic_secret.
type DataGenericSecret struct {
	Name string
	Args DataGenericSecretArgs
}

// DataSource returns the Terraform object type for [DataGenericSecret].
func (gs *DataGenericSecret) DataSource() string {
	return "vault_generic_secret"
}

// LocalName returns the local name for [DataGenericSecret].
func (gs *DataGenericSecret) LocalName() string {
	return gs.Name
}

// Configuration returns the configuration (args) for [DataGenericSecret].
func (gs *DataGenericSecret) Configuration() interface{} {
	return gs.Args
}

// Attributes returns the attributes for [DataGenericSecret].
func (gs *DataGenericSecret) Attributes() dataGenericSecretAttributes {
	return dataGenericSecretAttributes{ref: terra.ReferenceDataResource(gs)}
}

// DataGenericSecretArgs contains the configurations for vault_generic_secret.
type DataGenericSecretArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
	// Version: number, optional
	Version terra.NumberValue `hcl:"version,attr"`
	// WithLeaseStartTime: bool, optional
	WithLeaseStartTime terra.BoolValue `hcl:"with_lease_start_time,attr"`
}
type dataGenericSecretAttributes struct {
	ref terra.Reference
}

// Data returns a reference to field data of vault_generic_secret.
func (gs dataGenericSecretAttributes) Data() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gs.ref.Append("data"))
}

// DataJson returns a reference to field data_json of vault_generic_secret.
func (gs dataGenericSecretAttributes) DataJson() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("data_json"))
}

// Id returns a reference to field id of vault_generic_secret.
func (gs dataGenericSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("id"))
}

// LeaseDuration returns a reference to field lease_duration of vault_generic_secret.
func (gs dataGenericSecretAttributes) LeaseDuration() terra.NumberValue {
	return terra.ReferenceAsNumber(gs.ref.Append("lease_duration"))
}

// LeaseId returns a reference to field lease_id of vault_generic_secret.
func (gs dataGenericSecretAttributes) LeaseId() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("lease_id"))
}

// LeaseRenewable returns a reference to field lease_renewable of vault_generic_secret.
func (gs dataGenericSecretAttributes) LeaseRenewable() terra.BoolValue {
	return terra.ReferenceAsBool(gs.ref.Append("lease_renewable"))
}

// LeaseStartTime returns a reference to field lease_start_time of vault_generic_secret.
func (gs dataGenericSecretAttributes) LeaseStartTime() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("lease_start_time"))
}

// Namespace returns a reference to field namespace of vault_generic_secret.
func (gs dataGenericSecretAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_generic_secret.
func (gs dataGenericSecretAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("path"))
}

// Version returns a reference to field version of vault_generic_secret.
func (gs dataGenericSecretAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(gs.ref.Append("version"))
}

// WithLeaseStartTime returns a reference to field with_lease_start_time of vault_generic_secret.
func (gs dataGenericSecretAttributes) WithLeaseStartTime() terra.BoolValue {
	return terra.ReferenceAsBool(gs.ref.Append("with_lease_start_time"))
}
