// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataEcrpublicAuthorizationToken(name string, args DataEcrpublicAuthorizationTokenArgs) *DataEcrpublicAuthorizationToken {
	return &DataEcrpublicAuthorizationToken{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEcrpublicAuthorizationToken)(nil)

type DataEcrpublicAuthorizationToken struct {
	Name string
	Args DataEcrpublicAuthorizationTokenArgs
}

func (eat *DataEcrpublicAuthorizationToken) DataSource() string {
	return "aws_ecrpublic_authorization_token"
}

func (eat *DataEcrpublicAuthorizationToken) LocalName() string {
	return eat.Name
}

func (eat *DataEcrpublicAuthorizationToken) Configuration() interface{} {
	return eat.Args
}

func (eat *DataEcrpublicAuthorizationToken) Attributes() dataEcrpublicAuthorizationTokenAttributes {
	return dataEcrpublicAuthorizationTokenAttributes{ref: terra.ReferenceDataResource(eat)}
}

type DataEcrpublicAuthorizationTokenArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataEcrpublicAuthorizationTokenAttributes struct {
	ref terra.Reference
}

func (eat dataEcrpublicAuthorizationTokenAttributes) AuthorizationToken() terra.StringValue {
	return terra.ReferenceString(eat.ref.Append("authorization_token"))
}

func (eat dataEcrpublicAuthorizationTokenAttributes) ExpiresAt() terra.StringValue {
	return terra.ReferenceString(eat.ref.Append("expires_at"))
}

func (eat dataEcrpublicAuthorizationTokenAttributes) Id() terra.StringValue {
	return terra.ReferenceString(eat.ref.Append("id"))
}

func (eat dataEcrpublicAuthorizationTokenAttributes) Password() terra.StringValue {
	return terra.ReferenceString(eat.ref.Append("password"))
}

func (eat dataEcrpublicAuthorizationTokenAttributes) UserName() terra.StringValue {
	return terra.ReferenceString(eat.ref.Append("user_name"))
}
