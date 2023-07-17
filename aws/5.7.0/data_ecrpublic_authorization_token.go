// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataEcrpublicAuthorizationToken creates a new instance of [DataEcrpublicAuthorizationToken].
func NewDataEcrpublicAuthorizationToken(name string, args DataEcrpublicAuthorizationTokenArgs) *DataEcrpublicAuthorizationToken {
	return &DataEcrpublicAuthorizationToken{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEcrpublicAuthorizationToken)(nil)

// DataEcrpublicAuthorizationToken represents the Terraform data resource aws_ecrpublic_authorization_token.
type DataEcrpublicAuthorizationToken struct {
	Name string
	Args DataEcrpublicAuthorizationTokenArgs
}

// DataSource returns the Terraform object type for [DataEcrpublicAuthorizationToken].
func (eat *DataEcrpublicAuthorizationToken) DataSource() string {
	return "aws_ecrpublic_authorization_token"
}

// LocalName returns the local name for [DataEcrpublicAuthorizationToken].
func (eat *DataEcrpublicAuthorizationToken) LocalName() string {
	return eat.Name
}

// Configuration returns the configuration (args) for [DataEcrpublicAuthorizationToken].
func (eat *DataEcrpublicAuthorizationToken) Configuration() interface{} {
	return eat.Args
}

// Attributes returns the attributes for [DataEcrpublicAuthorizationToken].
func (eat *DataEcrpublicAuthorizationToken) Attributes() dataEcrpublicAuthorizationTokenAttributes {
	return dataEcrpublicAuthorizationTokenAttributes{ref: terra.ReferenceDataResource(eat)}
}

// DataEcrpublicAuthorizationTokenArgs contains the configurations for aws_ecrpublic_authorization_token.
type DataEcrpublicAuthorizationTokenArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataEcrpublicAuthorizationTokenAttributes struct {
	ref terra.Reference
}

// AuthorizationToken returns a reference to field authorization_token of aws_ecrpublic_authorization_token.
func (eat dataEcrpublicAuthorizationTokenAttributes) AuthorizationToken() terra.StringValue {
	return terra.ReferenceAsString(eat.ref.Append("authorization_token"))
}

// ExpiresAt returns a reference to field expires_at of aws_ecrpublic_authorization_token.
func (eat dataEcrpublicAuthorizationTokenAttributes) ExpiresAt() terra.StringValue {
	return terra.ReferenceAsString(eat.ref.Append("expires_at"))
}

// Id returns a reference to field id of aws_ecrpublic_authorization_token.
func (eat dataEcrpublicAuthorizationTokenAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eat.ref.Append("id"))
}

// Password returns a reference to field password of aws_ecrpublic_authorization_token.
func (eat dataEcrpublicAuthorizationTokenAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(eat.ref.Append("password"))
}

// UserName returns a reference to field user_name of aws_ecrpublic_authorization_token.
func (eat dataEcrpublicAuthorizationTokenAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(eat.ref.Append("user_name"))
}