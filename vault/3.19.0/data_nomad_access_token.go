// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataNomadAccessToken creates a new instance of [DataNomadAccessToken].
func NewDataNomadAccessToken(name string, args DataNomadAccessTokenArgs) *DataNomadAccessToken {
	return &DataNomadAccessToken{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNomadAccessToken)(nil)

// DataNomadAccessToken represents the Terraform data resource vault_nomad_access_token.
type DataNomadAccessToken struct {
	Name string
	Args DataNomadAccessTokenArgs
}

// DataSource returns the Terraform object type for [DataNomadAccessToken].
func (nat *DataNomadAccessToken) DataSource() string {
	return "vault_nomad_access_token"
}

// LocalName returns the local name for [DataNomadAccessToken].
func (nat *DataNomadAccessToken) LocalName() string {
	return nat.Name
}

// Configuration returns the configuration (args) for [DataNomadAccessToken].
func (nat *DataNomadAccessToken) Configuration() interface{} {
	return nat.Args
}

// Attributes returns the attributes for [DataNomadAccessToken].
func (nat *DataNomadAccessToken) Attributes() dataNomadAccessTokenAttributes {
	return dataNomadAccessTokenAttributes{ref: terra.ReferenceDataResource(nat)}
}

// DataNomadAccessTokenArgs contains the configurations for vault_nomad_access_token.
type DataNomadAccessTokenArgs struct {
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
}
type dataNomadAccessTokenAttributes struct {
	ref terra.Reference
}

// AccessorId returns a reference to field accessor_id of vault_nomad_access_token.
func (nat dataNomadAccessTokenAttributes) AccessorId() terra.StringValue {
	return terra.ReferenceAsString(nat.ref.Append("accessor_id"))
}

// Backend returns a reference to field backend of vault_nomad_access_token.
func (nat dataNomadAccessTokenAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(nat.ref.Append("backend"))
}

// Id returns a reference to field id of vault_nomad_access_token.
func (nat dataNomadAccessTokenAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nat.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_nomad_access_token.
func (nat dataNomadAccessTokenAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(nat.ref.Append("namespace"))
}

// Role returns a reference to field role of vault_nomad_access_token.
func (nat dataNomadAccessTokenAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(nat.ref.Append("role"))
}

// SecretId returns a reference to field secret_id of vault_nomad_access_token.
func (nat dataNomadAccessTokenAttributes) SecretId() terra.StringValue {
	return terra.ReferenceAsString(nat.ref.Append("secret_id"))
}