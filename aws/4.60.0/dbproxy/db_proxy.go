// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dbproxy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Auth struct {
	// AuthScheme: string, optional
	AuthScheme terra.StringValue `hcl:"auth_scheme,attr"`
	// ClientPasswordAuthType: string, optional
	ClientPasswordAuthType terra.StringValue `hcl:"client_password_auth_type,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// IamAuth: string, optional
	IamAuth terra.StringValue `hcl:"iam_auth,attr"`
	// SecretArn: string, optional
	SecretArn terra.StringValue `hcl:"secret_arn,attr"`
	// Username: string, optional
	Username terra.StringValue `hcl:"username,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AuthAttributes struct {
	ref terra.Reference
}

func (a AuthAttributes) InternalRef() terra.Reference {
	return a.ref
}

func (a AuthAttributes) InternalWithRef(ref terra.Reference) AuthAttributes {
	return AuthAttributes{ref: ref}
}

func (a AuthAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a AuthAttributes) AuthScheme() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("auth_scheme"))
}

func (a AuthAttributes) ClientPasswordAuthType() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("client_password_auth_type"))
}

func (a AuthAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("description"))
}

func (a AuthAttributes) IamAuth() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("iam_auth"))
}

func (a AuthAttributes) SecretArn() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("secret_arn"))
}

func (a AuthAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("username"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AuthState struct {
	AuthScheme             string `json:"auth_scheme"`
	ClientPasswordAuthType string `json:"client_password_auth_type"`
	Description            string `json:"description"`
	IamAuth                string `json:"iam_auth"`
	SecretArn              string `json:"secret_arn"`
	Username               string `json:"username"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
