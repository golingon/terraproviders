// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_cognito_identity_pool

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataCognitoIdentityProvidersAttributes struct {
	ref terra.Reference
}

func (cip DataCognitoIdentityProvidersAttributes) InternalRef() (terra.Reference, error) {
	return cip.ref, nil
}

func (cip DataCognitoIdentityProvidersAttributes) InternalWithRef(ref terra.Reference) DataCognitoIdentityProvidersAttributes {
	return DataCognitoIdentityProvidersAttributes{ref: ref}
}

func (cip DataCognitoIdentityProvidersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cip.ref.InternalTokens()
}

func (cip DataCognitoIdentityProvidersAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(cip.ref.Append("client_id"))
}

func (cip DataCognitoIdentityProvidersAttributes) ProviderName() terra.StringValue {
	return terra.ReferenceAsString(cip.ref.Append("provider_name"))
}

func (cip DataCognitoIdentityProvidersAttributes) ServerSideTokenCheck() terra.BoolValue {
	return terra.ReferenceAsBool(cip.ref.Append("server_side_token_check"))
}

type DataCognitoIdentityProvidersState struct {
	ClientId             string `json:"client_id"`
	ProviderName         string `json:"provider_name"`
	ServerSideTokenCheck bool   `json:"server_side_token_check"`
}
