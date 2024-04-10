// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import "github.com/golingon/lingon/pkg/terra"

// NewDataAdAccessCredentials creates a new instance of [DataAdAccessCredentials].
func NewDataAdAccessCredentials(name string, args DataAdAccessCredentialsArgs) *DataAdAccessCredentials {
	return &DataAdAccessCredentials{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAdAccessCredentials)(nil)

// DataAdAccessCredentials represents the Terraform data resource vault_ad_access_credentials.
type DataAdAccessCredentials struct {
	Name string
	Args DataAdAccessCredentialsArgs
}

// DataSource returns the Terraform object type for [DataAdAccessCredentials].
func (aac *DataAdAccessCredentials) DataSource() string {
	return "vault_ad_access_credentials"
}

// LocalName returns the local name for [DataAdAccessCredentials].
func (aac *DataAdAccessCredentials) LocalName() string {
	return aac.Name
}

// Configuration returns the configuration (args) for [DataAdAccessCredentials].
func (aac *DataAdAccessCredentials) Configuration() interface{} {
	return aac.Args
}

// Attributes returns the attributes for [DataAdAccessCredentials].
func (aac *DataAdAccessCredentials) Attributes() dataAdAccessCredentialsAttributes {
	return dataAdAccessCredentialsAttributes{ref: terra.ReferenceDataResource(aac)}
}

// DataAdAccessCredentialsArgs contains the configurations for vault_ad_access_credentials.
type DataAdAccessCredentialsArgs struct {
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
}
type dataAdAccessCredentialsAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_ad_access_credentials.
func (aac dataAdAccessCredentialsAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("backend"))
}

// CurrentPassword returns a reference to field current_password of vault_ad_access_credentials.
func (aac dataAdAccessCredentialsAttributes) CurrentPassword() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("current_password"))
}

// Id returns a reference to field id of vault_ad_access_credentials.
func (aac dataAdAccessCredentialsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("id"))
}

// LastPassword returns a reference to field last_password of vault_ad_access_credentials.
func (aac dataAdAccessCredentialsAttributes) LastPassword() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("last_password"))
}

// Namespace returns a reference to field namespace of vault_ad_access_credentials.
func (aac dataAdAccessCredentialsAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("namespace"))
}

// Role returns a reference to field role of vault_ad_access_credentials.
func (aac dataAdAccessCredentialsAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("role"))
}

// Username returns a reference to field username of vault_ad_access_credentials.
func (aac dataAdAccessCredentialsAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("username"))
}
